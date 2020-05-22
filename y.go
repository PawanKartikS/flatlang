// Code generated by goyacc parse.y. DO NOT EDIT.

//line parse.y:2
package flatlang

import __yyfmt__ "fmt"

//line parse.y:2

//line parse.y:5
type yySymType struct {
	yys   int
	token int
	node  *Node
}

const interp = 57346
const comment = 57347
const ident = 57348
const bool_ = 57349
const int_ = 57350
const float = 57351
const text = 57352
const negate = 57353
const lte = 57354
const gte = 57355
const pipe = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"interp",
	"comment",
	"ident",
	"bool_",
	"int_",
	"float",
	"text",
	"','",
	"':'",
	"';'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"'='",
	"'\\''",
	"'\"'",
	"'`'",
	"'@'",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'/'",
	"'*'",
	"negate",
	"'!'",
	"'>'",
	"'<'",
	"lte",
	"gte",
	"pipe",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 195

var yyAct = [...]int{

	9, 8, 31, 35, 32, 33, 5, 77, 34, 39,
	38, 36, 37, 76, 42, 43, 44, 45, 46, 47,
	48, 82, 32, 54, 50, 6, 75, 41, 40, 39,
	38, 36, 37, 68, 4, 36, 37, 62, 63, 64,
	65, 66, 67, 41, 40, 39, 38, 36, 37, 41,
	40, 39, 38, 36, 37, 73, 70, 71, 60, 41,
	40, 39, 34, 36, 37, 57, 73, 74, 61, 56,
	78, 7, 29, 27, 28, 55, 7, 72, 80, 13,
	81, 22, 52, 20, 7, 30, 24, 25, 26, 49,
	51, 58, 79, 19, 69, 1, 21, 18, 16, 14,
	15, 17, 7, 29, 27, 28, 10, 23, 11, 12,
	13, 59, 22, 3, 20, 2, 0, 24, 25, 26,
	0, 0, 0, 0, 19, 0, 0, 0, 18, 0,
	14, 15, 17, 7, 29, 27, 28, 0, 0, 0,
	0, 13, 0, 0, 53, 0, 0, 0, 24, 25,
	26, 0, 0, 0, 0, 19, 0, 0, 0, 18,
	16, 14, 15, 17, 7, 29, 27, 28, 0, 0,
	0, 0, 13, 0, 0, 0, 0, 0, 0, 24,
	25, 26, 0, 0, 0, 0, 19, 0, 0, 0,
	18, 16, 14, 15, 17,
}
var yyPact = [...]int{

	-1000, -1000, 65, 72, -11, -15, 96, -1000, -1000, 34,
	-1000, -1000, -1000, 158, 158, 158, 158, 158, 158, 158,
	70, 71, 127, 58, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 65, 65, -1000, -1000, 158, 158, 158, 158,
	158, 158, 18, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	82, -1000, 78, -1000, 24, -1000, 158, 56, 45, 3,
	96, -29, -1000, -1000, 6, 6, -18, -18, -1000, 158,
	80, 24, -1000, -1000, -1000, -1000, -1000, 158, 24, 158,
	2, 24, -1000,
}
var yyPgo = [...]int{

	0, 115, 113, 34, 25, 1, 3, 0, 65, 111,
	109, 108, 107, 106, 96, 95,
}
var yyR1 = [...]int{

	0, 15, 1, 1, 1, 6, 2, 3, 3, 4,
	4, 5, 5, 5, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 10, 10,
	10, 10, 10, 10, 10, 8, 8, 9, 9, 9,
	11, 11, 12, 12, 13, 13, 14, 14,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 3, 1, 3, 1, 3, 1,
	2, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 2, 2, 1, 3,
	3, 3, 1, 1, 1, 0, 2, 0, 2, 4,
	2, 2, 2, 3, 2, 2, 4, 5,
}
var yyChk = [...]int{

	-1000, -15, -1, -2, -3, -6, -4, 6, -5, -7,
	-13, -11, -10, 14, 34, 35, 33, 36, 32, 28,
	18, -14, 16, -12, 21, 22, 23, 8, 9, 7,
	13, 13, 33, 20, -5, -6, 29, 30, 28, 27,
	26, 25, -7, -7, -7, -7, -7, -7, -7, 19,
	-6, 19, 11, 17, -7, 17, 11, -8, -8, -9,
	-4, -3, -7, -7, -7, -7, -7, -7, 15, 12,
	-6, -7, 21, 10, 22, 23, 10, 4, -7, 12,
	-7, -7, 19,
}
var yyDef = [...]int{

	2, -2, 1, 0, 0, 28, 7, 5, 9, 11,
	12, 13, 14, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 35, 37, 32, 33, 34,
	3, 4, 0, 0, 10, 28, 0, 0, 0, 0,
	0, 0, 0, 22, 23, 24, 25, 26, 27, 44,
	0, 45, 0, 40, 42, 41, 0, 0, 0, 0,
	8, 6, 16, 17, 18, 19, 20, 21, 15, 0,
	0, 43, 29, 36, 30, 31, 38, 0, 46, 0,
	0, 47, 39,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 32, 22, 3, 3, 3, 26, 21,
	14, 15, 30, 27, 11, 28, 3, 29, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 12, 13,
	34, 20, 33, 3, 24, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 16, 3, 17, 3, 3, 23, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 25, 19,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 31,
	35, 36, 37,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	p := yylex.(*Parser)

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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:32
		{
			p.Result = yyDollar[1].node
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:35
		{
			yyVAL.node = NewNode(ProgramNode)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:36
		{
			yyDollar[1].node.N1(yyDollar[2].node).T1(yyDollar[3].token)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:37
		{
			yyDollar[1].node.N(yyDollar[2].node.Nodes...).T1(yyDollar[3].token)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:40
		{
			yyVAL.node = NewNode(IdentNode, yyDollar[1].token)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:43
		{
			yyVAL.node = NewNode(VarNode, yyDollar[2].token).N1(yyDollar[1].node).N(yyDollar[3].node.Nodes...)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:47
		{
			yyVAL.node = NewNode(ValNode).N1(yyDollar[1].node)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:48
		{
			yyDollar[1].node.T1(yyDollar[2].token).N1(yyDollar[3].node)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:52
		{
			yyVAL.node = NewNode(ValNode).N1(yyDollar[1].node)
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:53
		{
			yyDollar[1].node.N1(yyDollar[2].node)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:60
		{
			yyVAL.node = yyDollar[2].node
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:62
		{
			yyVAL.node = NewOpNode('/', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:63
		{
			yyVAL.node = NewOpNode('*', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:64
		{
			yyVAL.node = NewOpNode('-', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:65
		{
			yyVAL.node = NewOpNode('+', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:66
		{
			yyVAL.node = NewOpNode('&', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:67
		{
			yyVAL.node = NewOpNode('|', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:69
		{
			yyVAL.node = NewOpNode('<', yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:70
		{
			yyVAL.node = NewOpNode(lte, yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:71
		{
			yyVAL.node = NewOpNode('>', yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:72
		{
			yyVAL.node = NewOpNode(gte, yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:73
		{
			yyVAL.node = NewOpNode('!', yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:74
		{
			yyVAL.node = NewOpNode(negate, yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:79
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:80
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:81
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:82
		{
			yyVAL.node = NewNode(IntNode, yyDollar[1].token)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:83
		{
			yyVAL.node = NewNode(FloatNode, yyDollar[1].token)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:84
		{
			yyVAL.node = NewNode(BoolNode, yyDollar[1].token)
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:88
		{
			yyVAL.node = NewNode(StringNode)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:89
		{
			yyVAL.node = yyDollar[1].node.N1(NewNode(TextNode, yyDollar[2].token))
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:93
		{
			yyVAL.node = NewNode(StringNode)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:94
		{
			yyVAL.node = yyDollar[1].node.N1(NewNode(TextNode, yyDollar[2].token))
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parse.y:95
		{
			yyVAL.node = yyDollar[1].node.N1(NewNode(InterpNode, yyDollar[2].token, yyDollar[4].token).N1(yyDollar[3].node))
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:99
		{
			yyVAL.node = NewNode(ListNode, yyDollar[1].token, yyDollar[2].token)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:100
		{
			yyVAL.node = yyDollar[1].node.T1(yyDollar[2].token)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:104
		{
			yyVAL.node = NewNode(ListNode, yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:105
		{
			yyDollar[1].node.T1(yyDollar[2].token).N1(yyDollar[3].node)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:109
		{
			yyVAL.node = NewNode(MapNode, yyDollar[1].token, yyDollar[2].token)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:110
		{
			yyVAL.node = yyDollar[1].node.T1(yyDollar[2].token)
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parse.y:114
		{
			yyVAL.node = NewNode(MapNode, yyDollar[1].token).N1(yyDollar[2].node).T1(yyDollar[3].token).N1(yyDollar[4].node)
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parse.y:115
		{
			yyDollar[1].node.T1(yyDollar[2].token).N1(yyDollar[3].node).T1(yyDollar[4].token).N1(yyDollar[5].node)
		}
	}
	goto yystack /* stack new state and value */
}