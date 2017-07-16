package tags

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	"github.com/osteele/liquid/parser"
	"github.com/osteele/liquid/render"
	"github.com/stretchr/testify/require"
)

var iterationTests = []struct{ in, expected string }{
	{`{% for a in array %}{{ a }} {% endfor %}`, "first second third "},

	// loop modifiers
	{`{% for a in array reversed %}{{ a }}.{% endfor %}`, "third.second.first."},
	{`{% for a in array limit:2 %}{{ a }}.{% endfor %}`, "first.second."},
	{`{% for a in array offset:1 %}{{ a }}.{% endfor %}`, "second.third."},
	{`{% for a in array reversed limit:1 %}{{ a }}.{% endfor %}`, "third."},
	// TODO investigate how these combine; does it depend on the order?
	// {`{% for a in array reversed offset:1 %}{{ a }}.{% endfor %}`, "second.first."},
	// {`{% for a in array limit:1 offset:1 %}{{ a }}.{% endfor %}`, "second."},
	// {`{% for a in array reversed limit:1 offset:1 %}{{ a }}.{% endfor %}`, "second."},

	// loop variables
	{`{% for a in array %}{{ forloop.first }}.{% endfor %}`, "true.false.false."},
	{`{% for a in array %}{{ forloop.last }}.{% endfor %}`, "false.false.true."},
	{`{% for a in array %}{{ forloop.index }}.{% endfor %}`, "1.2.3."},
	{`{% for a in array %}{{ forloop.index0 }}.{% endfor %}`, "0.1.2."},
	{`{% for a in array %}{{ forloop.rindex }}.{% endfor %}`, "3.2.1."},
	{`{% for a in array %}{{ forloop.rindex0 }}.{% endfor %}`, "2.1.0."},
	{`{% for a in array %}{{ forloop.length }}.{% endfor %}`, "3.3.3."},

	{`{% for i in array %}{{ forloop.index }}[{% for j in array %}{{ forloop.index }}{% endfor %}]{{ forloop.index }}{% endfor %}`,
		"1[123]12[123]23[123]3"},

	{`{% for a in array reversed %}{{ forloop.first }}.{% endfor %}`, "true.false.false."},
	{`{% for a in array reversed %}{{ forloop.last }}.{% endfor %}`, "false.false.true."},
	{`{% for a in array reversed %}{{ forloop.index }}.{% endfor %}`, "1.2.3."},
	{`{% for a in array reversed %}{{ forloop.rindex }}.{% endfor %}`, "3.2.1."},
	{`{% for a in array reversed %}{{ forloop.length }}.{% endfor %}`, "3.3.3."},

	{`{% for a in array limit:2 %}{{ forloop.index }}.{% endfor %}`, "1.2."},
	{`{% for a in array limit:2 %}{{ forloop.rindex }}.{% endfor %}`, "2.1."},
	{`{% for a in array limit:2 %}{{ forloop.first }}.{% endfor %}`, "true.false."},
	{`{% for a in array limit:2 %}{{ forloop.last }}.{% endfor %}`, "false.true."},
	{`{% for a in array limit:2 %}{{ forloop.length }}.{% endfor %}`, "2.2."},

	{`{% for a in array offset:1 %}{{ forloop.index }}.{% endfor %}`, "1.2."},
	{`{% for a in array offset:1 %}{{ forloop.rindex }}.{% endfor %}`, "2.1."},
	{`{% for a in array offset:1 %}{{ forloop.first }}.{% endfor %}`, "true.false."},
	{`{% for a in array offset:1 %}{{ forloop.last }}.{% endfor %}`, "false.true."},
	{`{% for a in array offset:1 %}{{ forloop.length }}.{% endfor %}`, "2.2."},

	{`{% for a in array %}{% if a == 'second' %}{% break %}{% endif %}{{ a }}{% endfor %}`, "first"},
	{`{% for a in array %}{% if a == 'second' %}{% continue %}{% endif %}{{ a }}.{% endfor %}`, "first.third."},

	// hash
	{`{% for a in hash %}{{ a }}{% endfor %}`, "a"},

	// cycle
	{`{% for a in array %}{% cycle 'even', 'odd' %}.{% endfor %}`, "even.odd.even."},
	{`{% for a in array %}{% cycle '0', '1' %},{% cycle '0', '1' %}.{% endfor %}`, "0,1.0,1.0,1."},
	// {`{% for a in array %}{% cycle group: 'a', '0', '1' %},{% cycle '0', '1' %}.{% endfor %}`, "0,1.0,1.0,1."},

	// range
	{`{% for i in (3 .. 5) %}{{i}}.{% endfor %}`, "3.4.5."},
	{`{% for i in (3..5) %}{{i}}.{% endfor %}`, "3.4.5."},

	// tablerow
	{`{% tablerow product in products %}{{ product }}{% endtablerow %}`,
		`<tr class="row1"><td class="col1">Cool Shirt</td>
	<td class="col2">Alien Poster</td>
	<td class="col3">Batman Poster</td>
	<td class="col4">Bullseye Shirt</td>
	<td class="col5">Another Classic Vinyl</td>
	<td class="col6">Awesome Jeans</td></tr>`},

	{`{% tablerow product in products cols:2 %}{{ product }}{% endtablerow %}`,
		`<tr class="row1"><td class="col1">Cool Shirt</td><td class="col2">Alien Poster</td></tr>
		 <tr class="row2"><td class="col1">Batman Poster</td><td class="col2">Bullseye Shirt</td></tr>
	  	 <tr class="row3"><td class="col1">Another Classic Vinyl</td><td class="col2">Awesome Jeans</td></tr>`},
}

var iterationSyntaxErrorTests = []struct{ in, expected string }{
	{`{% for a b c %}{% endfor %}`, "parse error"},
	{`{% for a in array offset %}{% endfor %}`, "undefined loop modifier"},
	{`{% cycle %}`, "parse error"},
}

var iterationErrorTests = []struct{ in, expected string }{
	{`{% break %}`, "break outside a loop"},
	{`{% continue %}`, "continue outside a loop"},
	{`{% cycle 'a', 'b' %}`, "cycle must be within a forloop"},
}

var iterationTestBindings = map[string]interface{}{
	"array": []string{"first", "second", "third"},
	"hash":  map[string]interface{}{"a": 1},
	"products": []string{
		"Cool Shirt", "Alien Poster", "Batman Poster", "Bullseye Shirt", "Another Classic Vinyl", "Awesome Jeans",
	},
}

func TestIterationTags(t *testing.T) {
	config := render.NewConfig()
	AddStandardTags(config)
	for i, test := range iterationTests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			root, err := config.Compile(test.in, parser.SourceLoc{})
			require.NoErrorf(t, err, test.in)
			buf := new(bytes.Buffer)
			err = render.Render(root, buf, iterationTestBindings, config)
			require.NoErrorf(t, err, test.in)
			actual := buf.String()
			if strings.Contains(test.in, "{% tablerow") {
				replaceWS := regexp.MustCompile(`\n\s*`).ReplaceAllString
				actual = replaceWS(actual, "")
				test.expected = replaceWS(test.expected, "")
			}
			require.Equalf(t, test.expected, actual, test.in)
		})
	}
}

func TestLoopTag_errors(t *testing.T) {
	config := render.NewConfig()
	AddStandardTags(config)

	for i, test := range iterationSyntaxErrorTests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			_, err := config.Compile(test.in, parser.SourceLoc{})
			require.Errorf(t, err, test.in)
			require.Containsf(t, err.Error(), test.expected, test.in)
		})
	}

	for i, test := range iterationErrorTests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			root, err := config.Compile(test.in, parser.SourceLoc{})
			require.NoErrorf(t, err, test.in)
			err = render.Render(root, ioutil.Discard, iterationTestBindings, config)
			require.Errorf(t, err, test.in)
			require.Containsf(t, err.Error(), test.expected, test.in)
		})
	}
}