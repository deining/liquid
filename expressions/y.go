// Code generated by goyacc expressions.y. DO NOT EDIT.

//line expressions.y:2
package expressions

import __yyfmt__ "fmt"

//line expressions.y:2
import (
	"fmt"
	"github.com/osteele/liquid/values"
)

func init() {
	// This allows adding and removing references to fmt in the rules below,
	// without having to comment and un-comment the import statement above.
	_ = ""
}

//line expressions.y:15
type yySymType struct {
	yys           int
	name          string
	val           interface{}
	f             func(Context) values.Value
	s             string
	ss            []string
	exprs         []Expression
	cycle         Cycle
	cyclefn       func(string) Cycle
	loop          Loop
	loopmods      loopModifiers
	filter_params []valueFn
}

const LITERAL = 57346
const IDENTIFIER = 57347
const KEYWORD = 57348
const PROPERTY = 57349
const ASSIGN = 57350
const CYCLE = 57351
const LOOP = 57352
const WHEN = 57353
const EQ = 57354
const NEQ = 57355
const GE = 57356
const LE = 57357
const IN = 57358
const AND = 57359
const OR = 57360
const CONTAINS = 57361
const DOTDOT = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LITERAL",
	"IDENTIFIER",
	"KEYWORD",
	"PROPERTY",
	"ASSIGN",
	"CYCLE",
	"LOOP",
	"WHEN",
	"EQ",
	"NEQ",
	"GE",
	"LE",
	"IN",
	"AND",
	"OR",
	"CONTAINS",
	"DOTDOT",
	"'.'",
	"'|'",
	"'<'",
	"'>'",
	"';'",
	"'='",
	"':'",
	"','",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 75,
	20, 18,
	-2, 23,
	-1, 76,
	20, 19,
	-2, 24,
}

const yyPrivate = 57344

const yyLast = 104

var yyAct = [...]int8{
	9, 74, 46, 41, 8, 87, 78, 23, 14, 15,
	18, 10, 11, 25, 42, 3, 4, 5, 6, 25,
	37, 58, 10, 11, 40, 42, 45, 50, 51, 52,
	53, 54, 55, 56, 57, 43, 12, 26, 60, 38,
	81, 24, 59, 26, 69, 60, 25, 12, 66, 65,
	68, 61, 24, 62, 44, 70, 25, 79, 80, 75,
	76, 27, 28, 31, 32, 71, 72, 47, 33, 77,
	26, 7, 30, 29, 10, 11, 21, 14, 15, 82,
	26, 16, 83, 86, 12, 13, 35, 36, 48, 49,
	84, 85, 19, 34, 2, 1, 73, 20, 39, 64,
	17, 22, 67, 63,
}

var yyPact = [...]int16{
	7, -1000, 60, 76, 88, 71, 18, -1000, 19, 49,
	-1000, -1000, 18, -1000, 18, 18, -6, 14, -3, -1000,
	10, 38, 1, 39, 83, -1000, 18, 18, 18, 18,
	18, 18, 18, 18, -9, -1000, -1000, 18, -1000, -1000,
	88, -1000, 88, -1000, 70, -1000, -1000, 18, -1000, 18,
	12, 6, 6, 6, 6, 6, 6, 6, -1000, 30,
	6, -14, -14, -1000, 55, 19, 39, -22, 6, -1000,
	-1000, -1000, -1000, 52, 20, -1000, -1000, -1000, 18, -1000,
	86, 86, 6, -1000, -1000, -1000, -25, -1000,
}

var yyPgo = [...]int8{
	0, 0, 71, 4, 93, 1, 103, 102, 101, 2,
	100, 98, 3, 97, 96, 10, 95,
}

var yyR1 = [...]int8{
	0, 16, 16, 16, 16, 16, 10, 11, 11, 12,
	12, 8, 9, 9, 15, 13, 6, 6, 5, 5,
	14, 14, 14, 1, 1, 1, 1, 1, 3, 3,
	3, 7, 7, 2, 2, 2, 2, 2, 2, 2,
	2, 4, 4, 4,
}

var yyR2 = [...]int8{
	0, 2, 5, 3, 3, 3, 2, 3, 1, 0,
	3, 2, 0, 3, 1, 4, 5, 1, 1, 1,
	0, 2, 3, 1, 1, 2, 4, 3, 1, 3,
	4, 1, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 1, 3, 3,
}

var yyChk = [...]int16{
	-1000, -16, -4, 8, 9, 10, 11, -2, -3, -1,
	4, 5, 29, 25, 17, 18, 5, -10, -15, 4,
	-13, 5, -8, -1, 22, 7, 31, 12, 13, 24,
	23, 14, 15, 19, -4, -2, -2, 26, 25, -11,
	27, -12, 28, 25, 16, 25, -9, 28, 5, 6,
	-1, -1, -1, -1, -1, -1, -1, -1, 30, -3,
	-1, -15, -15, -6, 29, -3, -1, -7, -1, 32,
	25, -12, -12, -14, -5, 4, 5, -9, 28, 5,
	6, 20, -1, -5, 4, 5, -5, 30,
}

var yyDef = [...]int8{
	0, -2, 0, 0, 0, 0, 0, 41, 33, 28,
	23, 24, 0, 1, 0, 0, 0, 0, 9, 14,
	0, 0, 0, 12, 0, 25, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 42, 43, 0, 3, 6,
	0, 8, 0, 4, 0, 5, 11, 0, 29, 0,
	0, 34, 35, 36, 37, 38, 39, 40, 27, 0,
	28, 9, 9, 20, 0, 17, 12, 30, 31, 26,
	2, 7, 10, 15, 0, -2, -2, 13, 0, 21,
	0, 0, 32, 22, 18, 19, 0, 16,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	29, 30, 3, 3, 28, 3, 21, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 27, 25,
	23, 26, 24, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 31, 3, 32, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 22,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:45
		{
			yylex.(*lexer).val = yyDollar[1].f
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions.y:46
		{
			yylex.(*lexer).Assignment = Assignment{yyDollar[2].name, &expression{yyDollar[4].f}}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:49
		{
			yylex.(*lexer).Cycle = yyDollar[2].cycle
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:50
		{
			yylex.(*lexer).Loop = yyDollar[2].loop
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:51
		{
			yylex.(*lexer).When = When{yyDollar[2].exprs}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:54
		{
			yyVAL.cycle = yyDollar[2].cyclefn(yyDollar[1].s)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:57
		{
			h, t := yyDollar[2].s, yyDollar[3].ss
			yyVAL.cyclefn = func(g string) Cycle { return Cycle{g, append([]string{h}, t...)} }
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:61
		{
			vals := yyDollar[1].ss
			yyVAL.cyclefn = func(h string) Cycle { return Cycle{Values: append([]string{h}, vals...)} }
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:68
		{
			yyVAL.ss = []string{}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:69
		{
			yyVAL.ss = append([]string{yyDollar[2].s}, yyDollar[3].ss...)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:72
		{
			yyVAL.exprs = append([]Expression{&expression{yyDollar[1].f}}, yyDollar[2].exprs...)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:74
		{
			yyVAL.exprs = []Expression{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:75
		{
			yyVAL.exprs = append([]Expression{&expression{yyDollar[2].f}}, yyDollar[3].exprs...)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:78
		{
			s, ok := yyDollar[1].val.(string)
			if !ok {
				panic(SyntaxError(fmt.Sprintf("expected a string for %q", yyDollar[1].val)))
			}
			yyVAL.s = s
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:86
		{
			name, expr, mods := yyDollar[1].name, yyDollar[3].f, yyDollar[4].loopmods
			yyVAL.loop = Loop{name, &expression{expr}, mods}
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
//line expressions.y:92
		{
			yyVAL.f = makeRangeExpr(yyDollar[2].f, yyDollar[4].f)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:100
		{
			val := yyDollar[1].val
			yyVAL.f = func(Context) values.Value { return values.ValueOf(val) }
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:101
		{
			name := yyDollar[1].name
			yyVAL.f = func(ctx Context) values.Value { return values.ValueOf(ctx.Get(name)) }
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
//line expressions.y:104
		{
			yyVAL.loopmods = loopModifiers{}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:105
		{
			switch yyDollar[2].name {
			case "reversed":
				yyDollar[1].loopmods.Reversed = true
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:114
		{
			switch yyDollar[2].name {
			case "cols":
				yyDollar[1].loopmods.Cols = &expression{yyDollar[3].f}
			case "limit":
				yyDollar[1].loopmods.Limit = &expression{yyDollar[3].f}
			case "offset":
				yyDollar[1].loopmods.Offset = &expression{yyDollar[3].f}
			default:
				panic(SyntaxError(fmt.Sprintf("undefined loop modifier %q", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:130
		{
			val := yyDollar[1].val
			yyVAL.f = func(Context) values.Value { return values.ValueOf(val) }
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:131
		{
			name := yyDollar[1].name
			yyVAL.f = func(ctx Context) values.Value { return values.ValueOf(ctx.Get(name)) }
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line expressions.y:132
		{
			yyVAL.f = makeObjectPropertyExpr(yyDollar[1].f, yyDollar[2].name)
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:133
		{
			yyVAL.f = makeIndexExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:134
		{
			yyVAL.f = yyDollar[2].f
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:139
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, nil)
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line expressions.y:140
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, yyDollar[4].filter_params)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line expressions.y:144
		{
			yyVAL.filter_params = []valueFn{yyDollar[1].f}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:146
		{
			yyVAL.filter_params = append(yyDollar[1].filter_params, yyDollar[3].f)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:150
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Equal(b))
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:157
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(!a.Equal(b))
			}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:164
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(b.Less(a))
			}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:171
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Less(b))
			}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:178
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(b.Less(a) || a.Equal(b))
			}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:185
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				a, b := fa(ctx), fb(ctx)
				return values.ValueOf(a.Less(b) || a.Equal(b))
			}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:192
		{
			yyVAL.f = makeContainsExpr(yyDollar[1].f, yyDollar[3].f)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:197
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				return values.ValueOf(fa(ctx).Test() && fb(ctx).Test())
			}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line expressions.y:203
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) values.Value {
				return values.ValueOf(fa(ctx).Test() || fb(ctx).Test())
			}
		}
	}
	goto yystack /* stack new state and value */
}
