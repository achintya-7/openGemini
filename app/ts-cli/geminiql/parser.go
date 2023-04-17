// Code generated by goyacc -o parser.go -p QL parser.y. DO NOT EDIT.

//line parser.y:2
/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package geminiql

import (
	__yyfmt__ "fmt"
	"strconv"
)

func updateStmt(QLlex interface{}, stmt Statement) {
	QLlex.(*QLLexerImpl).UpdateStmt(stmt)
}

//line parser.y:32
type QLSymType struct {
	yys      int
	stmts    []Statement
	stmt     Statement
	str      string
	strslice []string
	integer  int64
	decimal  float64
	pair     Pair
	pairs    Pairs
}

const INSERT = 57346
const INTO = 57347
const USE = 57348
const SET = 57349
const CHUNKED = 57350
const CHUNK_SIZE = 57351
const DOT = 57352
const COMMA = 57353
const EQ = 57354
const IDENT = 57355
const INTEGER = 57356
const DECIMAL = 57357
const STRING = 57358
const RAW = 57359

var QLToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INSERT",
	"INTO",
	"USE",
	"SET",
	"CHUNKED",
	"CHUNK_SIZE",
	"DOT",
	"COMMA",
	"EQ",
	"IDENT",
	"INTEGER",
	"DECIMAL",
	"STRING",
	"RAW",
}

var QLStatenames = [...]string{}

const QLEofCode = 1
const QLErrCode = 2
const QLInitialStackSize = 16

//line parser.y:245

//line yacctab:1
var QLExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	13, 13,
	-2, 27,
}

const QLPrivate = 57344

const QLLast = 48

var QLAct = [...]int8{
	34, 19, 13, 17, 39, 41, 42, 40, 28, 47,
	23, 12, 36, 21, 18, 25, 24, 16, 26, 16,
	7, 45, 8, 9, 10, 11, 32, 33, 44, 31,
	29, 30, 20, 38, 37, 43, 22, 27, 35, 15,
	14, 6, 5, 4, 3, 46, 2, 1,
}

var QLPact = [...]int16{
	16, -1000, -1000, -1000, -1000, -1000, -1000, 6, 1, 0,
	-1000, -4, 5, -1000, -6, 19, -1000, -1000, 21, -1000,
	18, 14, -1000, -1000, 4, -1000, 21, -1000, -1000, -1,
	1, 0, -9, -1000, -1, 17, 9, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1, -8, -1000, -1000,
}

var QLPgo = [...]int8{
	0, 47, 46, 44, 43, 42, 41, 2, 40, 39,
	38, 0, 37, 36, 3, 32, 1,
}

var QLR1 = [...]int8{
	0, 1, 1, 1, 1, 1, 4, 3, 2, 2,
	2, 5, 6, 14, 14, 7, 7, 8, 15, 15,
	15, 15, 16, 16, 11, 11, 10, 9, 12, 13,
}

var QLR2 = [...]int8{
	0, 1, 1, 1, 1, 1, 2, 2, 4, 3,
	2, 1, 2, 1, 3, 1, 2, 4, 3, 3,
	3, 3, 1, 3, 1, 3, 3, 1, 1, 1,
}

var QLChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, 4, 6, 7,
	8, 9, 5, -7, -8, -9, 13, -14, 13, -16,
	-15, 13, -13, 14, -14, -7, 13, -12, 14, 11,
	10, 11, 12, -7, -11, -10, 13, -14, -16, 13,
	16, 14, 15, -11, 11, 12, -11, 17,
}

var QLDef = [...]int8{
	0, -2, 1, 2, 3, 4, 5, 0, 0, 0,
	11, 0, 0, 10, 15, 0, 27, 7, 13, 6,
	22, 0, 12, 29, 0, 9, -2, 16, 28, 0,
	0, 0, 0, 8, 0, 24, 0, 14, 23, 18,
	19, 20, 21, 17, 0, 0, 25, 26,
}

var QLTok1 = [...]int8{
	1,
}

var QLTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}

var QLTok3 = [...]int8{
	0,
}

var QLErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	QLDebug        = 0
	QLErrorVerbose = false
)

type QLLexer interface {
	Lex(lval *QLSymType) int
	Error(s string)
}

type QLParser interface {
	Parse(QLLexer) int
	Lookahead() int
}

type QLParserImpl struct {
	lval  QLSymType
	stack [QLInitialStackSize]QLSymType
	char  int
}

func (p *QLParserImpl) Lookahead() int {
	return p.char
}

func QLNewParser() QLParser {
	return &QLParserImpl{}
}

const QLFlag = -1000

func QLTokname(c int) string {
	if c >= 1 && c-1 < len(QLToknames) {
		if QLToknames[c-1] != "" {
			return QLToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func QLStatname(s int) string {
	if s >= 0 && s < len(QLStatenames) {
		if QLStatenames[s] != "" {
			return QLStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func QLErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !QLErrorVerbose {
		return "syntax error"
	}

	for _, e := range QLErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + QLTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(QLPact[state])
	for tok := TOKSTART; tok-1 < len(QLToknames); tok++ {
		if n := base + tok; n >= 0 && n < QLLast && int(QLChk[int(QLAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if QLDef[state] == -2 {
		i := 0
		for QLExca[i] != -1 || int(QLExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; QLExca[i] >= 0; i += 2 {
			tok := int(QLExca[i])
			if tok < TOKSTART || QLExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if QLExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += QLTokname(tok)
	}
	return res
}

func QLlex1(lex QLLexer, lval *QLSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(QLTok1[0])
		goto out
	}
	if char < len(QLTok1) {
		token = int(QLTok1[char])
		goto out
	}
	if char >= QLPrivate {
		if char < QLPrivate+len(QLTok2) {
			token = int(QLTok2[char-QLPrivate])
			goto out
		}
	}
	for i := 0; i < len(QLTok3); i += 2 {
		token = int(QLTok3[i+0])
		if token == char {
			token = int(QLTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(QLTok2[1]) /* unknown char */
	}
	if QLDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", QLTokname(token), uint(char))
	}
	return char, token
}

func QLParse(QLlex QLLexer) int {
	return QLNewParser().Parse(QLlex)
}

func (QLrcvr *QLParserImpl) Parse(QLlex QLLexer) int {
	var QLn int
	var QLVAL QLSymType
	var QLDollar []QLSymType
	_ = QLDollar // silence set and not used
	QLS := QLrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	QLstate := 0
	QLrcvr.char = -1
	QLtoken := -1 // QLrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		QLstate = -1
		QLrcvr.char = -1
		QLtoken = -1
	}()
	QLp := -1
	goto QLstack

ret0:
	return 0

ret1:
	return 1

QLstack:
	/* put a state and value onto the stack */
	if QLDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", QLTokname(QLtoken), QLStatname(QLstate))
	}

	QLp++
	if QLp >= len(QLS) {
		nyys := make([]QLSymType, len(QLS)*2)
		copy(nyys, QLS)
		QLS = nyys
	}
	QLS[QLp] = QLVAL
	QLS[QLp].yys = QLstate

QLnewstate:
	QLn = int(QLPact[QLstate])
	if QLn <= QLFlag {
		goto QLdefault /* simple state */
	}
	if QLrcvr.char < 0 {
		QLrcvr.char, QLtoken = QLlex1(QLlex, &QLrcvr.lval)
	}
	QLn += QLtoken
	if QLn < 0 || QLn >= QLLast {
		goto QLdefault
	}
	QLn = int(QLAct[QLn])
	if int(QLChk[QLn]) == QLtoken { /* valid shift */
		QLrcvr.char = -1
		QLtoken = -1
		QLVAL = QLrcvr.lval
		QLstate = QLn
		if Errflag > 0 {
			Errflag--
		}
		goto QLstack
	}

QLdefault:
	/* default state action */
	QLn = int(QLDef[QLstate])
	if QLn == -2 {
		if QLrcvr.char < 0 {
			QLrcvr.char, QLtoken = QLlex1(QLlex, &QLrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if QLExca[xi+0] == -1 && int(QLExca[xi+1]) == QLstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			QLn = int(QLExca[xi+0])
			if QLn < 0 || QLn == QLtoken {
				break
			}
		}
		QLn = int(QLExca[xi+1])
		if QLn < 0 {
			goto ret0
		}
	}
	if QLn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			QLlex.Error(QLErrorMessage(QLstate, QLtoken))
			Nerrs++
			if QLDebug >= 1 {
				__yyfmt__.Printf("%s", QLStatname(QLstate))
				__yyfmt__.Printf(" saw %s\n", QLTokname(QLtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for QLp >= 0 {
				QLn = int(QLPact[QLS[QLp].yys]) + QLErrCode
				if QLn >= 0 && QLn < QLLast {
					QLstate = int(QLAct[QLn]) /* simulate a shift of "error" */
					if int(QLChk[QLstate]) == QLErrCode {
						goto QLstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if QLDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", QLS[QLp].yys)
				}
				QLp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if QLDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", QLTokname(QLtoken))
			}
			if QLtoken == QLEofCode {
				goto ret1
			}
			QLrcvr.char = -1
			QLtoken = -1
			goto QLnewstate /* try again in the same state */
		}
	}

	/* reduction by production QLn */
	if QLDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", QLn, QLStatname(QLstate))
	}

	QLnt := QLn
	QLpt := QLp
	_ = QLpt // guard against "declared and not used"

	QLp -= int(QLR2[QLn])
	// QLp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if QLp+1 >= len(QLS) {
		nyys := make([]QLSymType, len(QLS)*2)
		copy(nyys, QLS)
		QLS = nyys
	}
	QLVAL = QLS[QLp+1]

	/* consult goto table to find next state */
	QLn = int(QLR1[QLn])
	QLg := int(QLPgo[QLn])
	QLj := QLg + QLS[QLp].yys + 1

	if QLj >= QLLast {
		QLstate = int(QLAct[QLg])
	} else {
		QLstate = int(QLAct[QLj])
		if int(QLChk[QLstate]) != -QLn {
			QLstate = int(QLAct[QLg])
		}
	}
	// dummy call; replaced with literal code
	switch QLnt {

	case 1:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:67
		{
			updateStmt(QLlex, QLDollar[1].stmt)
		}
	case 2:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:71
		{
			updateStmt(QLlex, QLDollar[1].stmt)
		}
	case 3:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:75
		{
			updateStmt(QLlex, QLDollar[1].stmt)
		}
	case 4:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:79
		{
			updateStmt(QLlex, QLDollar[1].stmt)
		}
	case 5:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:83
		{
			updateStmt(QLlex, QLDollar[1].stmt)
		}
	case 6:
		QLDollar = QLS[QLpt-2 : QLpt+1]
//line parser.y:89
		{
			stmt := &SetStatement{}
			stmt.KVS = QLDollar[2].pairs
			QLVAL.stmt = stmt
		}
	case 7:
		QLDollar = QLS[QLpt-2 : QLpt+1]
//line parser.y:97
		{
			stmt := &UseStatement{}
			if len(QLDollar[2].strslice) == 1 {
				stmt.DB = QLDollar[2].strslice[0]
				QLVAL.stmt = stmt
			} else if len(QLDollar[2].strslice) == 2 {
				stmt.DB = QLDollar[2].strslice[0]
				stmt.RP = QLDollar[2].strslice[1]
				QLVAL.stmt = stmt
			} else {
				QLlex.Error("namespace must be <db>.<rp>")
			}
		}
	case 8:
		QLDollar = QLS[QLpt-4 : QLpt+1]
//line parser.y:113
		{
			stmt := &InsertStatement{}
			stmt.LineProtocol = QLDollar[4].str

			if len(QLDollar[3].strslice) != 2 {
				QLlex.Error("namespace must be <db>.<rp>")
			} else {
				stmt.DB = QLDollar[3].strslice[0]
				stmt.RP = QLDollar[3].strslice[1]
				QLVAL.stmt = stmt
			}
		}
	case 9:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:126
		{
			stmt := &InsertStatement{}
			stmt.LineProtocol = QLDollar[3].str
			QLVAL.stmt = stmt
		}
	case 10:
		QLDollar = QLS[QLpt-2 : QLpt+1]
//line parser.y:132
		{
			stmt := &InsertStatement{}
			stmt.LineProtocol = QLDollar[2].str
			QLVAL.stmt = stmt
		}
	case 11:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:140
		{
			stmt := &ChunkedStatement{}
			QLVAL.stmt = stmt
		}
	case 12:
		QLDollar = QLS[QLpt-2 : QLpt+1]
//line parser.y:147
		{
			stmt := &ChunkSizeStatement{}
			stmt.Size = QLDollar[2].integer
			QLVAL.stmt = stmt
		}
	case 13:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:155
		{
			QLVAL.strslice = []string{QLDollar[1].str}
		}
	case 14:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:159
		{
			ns := []string{QLDollar[1].str}
			QLVAL.strslice = append(ns, QLDollar[3].strslice...)
		}
	case 15:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:166
		{
			QLVAL.str = QLDollar[1].str
		}
	case 16:
		QLDollar = QLS[QLpt-2 : QLpt+1]
//line parser.y:170
		{
			QLVAL.str = QLDollar[1].str + " " + QLDollar[2].str
		}
	case 17:
		QLDollar = QLS[QLpt-4 : QLpt+1]
//line parser.y:176
		{
			QLVAL.str = QLDollar[1].str + QLDollar[2].str + QLDollar[3].str + " " + QLDollar[4].str
		}
	case 18:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:182
		{
			p := NewPair(QLDollar[1].str, QLDollar[3].str)
			QLVAL.pair = *p
		}
	case 19:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:187
		{
			p := NewPair(QLDollar[1].str, QLDollar[3].str)
			QLVAL.pair = *p
		}
	case 20:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:192
		{
			p := NewPair(QLDollar[1].str, QLDollar[3].integer)
			QLVAL.pair = *p
		}
	case 21:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:197
		{
			p := NewPair(QLDollar[1].str, QLDollar[3].decimal)
			QLVAL.pair = *p
		}
	case 22:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:204
		{
			QLVAL.pairs = Pairs{QLDollar[1].pair}
		}
	case 23:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:208
		{
			QLVAL.pairs = append(QLDollar[3].pairs, QLDollar[1].pair)
		}
	case 24:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:214
		{
			QLVAL.str = QLDollar[1].str
		}
	case 25:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:218
		{
			QLVAL.str = QLDollar[1].str + QLDollar[2].str + QLDollar[3].str
		}
	case 26:
		QLDollar = QLS[QLpt-3 : QLpt+1]
//line parser.y:224
		{
			QLVAL.str = QLDollar[1].str + QLDollar[2].str + QLDollar[3].str
		}
	case 27:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:230
		{
			QLVAL.str = QLDollar[1].str
		}
	case 28:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:236
		{
			QLVAL.str = strconv.FormatInt(QLDollar[1].integer, 10)
		}
	case 29:
		QLDollar = QLS[QLpt-1 : QLpt+1]
//line parser.y:242
		{
			QLVAL.integer = QLDollar[1].integer
		}
	}
	goto QLstack /* stack new state and value */
}
