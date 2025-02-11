// Code generated from Predikit.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Predikit
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PredikitParser struct {
	*antlr.BaseParser
}

var PredikitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func predikitParserInit() {
	staticData := &PredikitParserStaticData
	staticData.LiteralNames = []string{
		"", "'all'", "'any'", "'none'", "'test'", "'not'", "'tool'", "'@'",
		"':'", "'$'", "'{'", "'}'", "'('", "')'", "'true'", "'false'", "'='",
		"'!='", "'=~'", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "PK_ALL", "PK_ANY", "PK_NONE", "PK_TEST", "PK_NOT", "PK_TOOL", "PK_RETRYING",
		"PK_COLON", "PK_DOLLAR", "PK_LCURLY", "PK_RCURLY", "PK_LPAREN", "PK_RPAREN",
		"PK_TRUE", "PK_FALSE", "PK_CMP_EQ", "PK_CMP_NEQ", "PK_CMP_RE", "PK_CMP_GT",
		"PK_CMP_GTE", "PK_CMP_LT", "PK_CMP_LTE", "PK_FN_LIT", "PK_STRING_LIT",
		"PK_ID", "PK_TYPENAME", "PK_INT", "LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"pk_toplevel", "pk_toplevel_child", "pk_group", "pk_group_child", "pk_group_agg",
		"pk_test", "pk_tool", "pk_tool_child", "pk_tool_metaparam", "pk_actual_param",
		"pk_actual_param_value", "pk_tool_actual_param", "pk_tool_actual_param_value",
		"pk_conversion_fn", "pk_bool",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 123, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 4, 0,
		32, 8, 0, 11, 0, 12, 0, 33, 1, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 2, 1,
		2, 1, 2, 4, 2, 44, 8, 2, 11, 2, 12, 2, 45, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		3, 3, 53, 8, 3, 1, 4, 1, 4, 1, 5, 3, 5, 58, 8, 5, 1, 5, 1, 5, 3, 5, 62,
		8, 5, 1, 5, 1, 5, 1, 5, 4, 5, 67, 8, 5, 11, 5, 12, 5, 68, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 4, 6, 77, 8, 6, 11, 6, 12, 6, 78, 1, 6, 1, 6, 1, 7,
		1, 7, 3, 7, 85, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 91, 8, 8, 11, 8, 12,
		8, 92, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 105, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 116, 8, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 2,
		1, 0, 1, 3, 1, 0, 14, 15, 126, 0, 31, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4,
		40, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 57, 1, 0, 0,
		0, 12, 72, 1, 0, 0, 0, 14, 84, 1, 0, 0, 0, 16, 86, 1, 0, 0, 0, 18, 96,
		1, 0, 0, 0, 20, 104, 1, 0, 0, 0, 22, 106, 1, 0, 0, 0, 24, 115, 1, 0, 0,
		0, 26, 117, 1, 0, 0, 0, 28, 120, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30,
		1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0,
		34, 1, 1, 0, 0, 0, 35, 39, 3, 4, 2, 0, 36, 39, 3, 10, 5, 0, 37, 39, 3,
		12, 6, 0, 38, 35, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39,
		3, 1, 0, 0, 0, 40, 41, 3, 8, 4, 0, 41, 43, 5, 10, 0, 0, 42, 44, 3, 6, 3,
		0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 5, 11, 0, 0, 48, 5, 1, 0, 0, 0,
		49, 53, 3, 10, 5, 0, 50, 53, 3, 4, 2, 0, 51, 53, 3, 18, 9, 0, 52, 49, 1,
		0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0, 0, 53, 7, 1, 0, 0, 0, 54,
		55, 7, 0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 58, 5, 7, 0, 0, 57, 56, 1, 0, 0,
		0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 5, 4, 0, 0, 60, 62,
		5, 5, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 64, 5, 25, 0, 0, 64, 66, 5, 10, 0, 0, 65, 67, 3, 18, 9, 0, 66, 65,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 71, 5, 11, 0, 0, 71, 11, 1, 0, 0, 0, 72, 73, 5,
		6, 0, 0, 73, 74, 5, 25, 0, 0, 74, 76, 5, 10, 0, 0, 75, 77, 3, 14, 7, 0,
		76, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 5, 11, 0, 0, 81, 13, 1, 0, 0, 0, 82,
		85, 3, 18, 9, 0, 83, 85, 3, 16, 8, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0,
		0, 0, 85, 15, 1, 0, 0, 0, 86, 87, 5, 9, 0, 0, 87, 88, 5, 25, 0, 0, 88,
		90, 5, 10, 0, 0, 89, 91, 3, 22, 11, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0,
		0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		5, 11, 0, 0, 95, 17, 1, 0, 0, 0, 96, 97, 5, 25, 0, 0, 97, 98, 5, 8, 0,
		0, 98, 99, 3, 20, 10, 0, 99, 19, 1, 0, 0, 0, 100, 105, 5, 27, 0, 0, 101,
		105, 5, 24, 0, 0, 102, 105, 3, 28, 14, 0, 103, 105, 3, 26, 13, 0, 104,
		100, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 103,
		1, 0, 0, 0, 105, 21, 1, 0, 0, 0, 106, 107, 5, 25, 0, 0, 107, 108, 5, 8,
		0, 0, 108, 109, 3, 24, 12, 0, 109, 23, 1, 0, 0, 0, 110, 116, 5, 27, 0,
		0, 111, 116, 5, 24, 0, 0, 112, 116, 3, 28, 14, 0, 113, 116, 3, 26, 13,
		0, 114, 116, 5, 26, 0, 0, 115, 110, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115,
		112, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 25, 1,
		0, 0, 0, 117, 118, 5, 25, 0, 0, 118, 119, 5, 23, 0, 0, 119, 27, 1, 0, 0,
		0, 120, 121, 7, 1, 0, 0, 121, 29, 1, 0, 0, 0, 12, 33, 38, 45, 52, 57, 61,
		68, 78, 84, 92, 104, 115,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PredikitParserInit initializes any static state used to implement PredikitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPredikitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PredikitParserInit() {
	staticData := &PredikitParserStaticData
	staticData.once.Do(predikitParserInit)
}

// NewPredikitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPredikitParser(input antlr.TokenStream) *PredikitParser {
	PredikitParserInit()
	this := new(PredikitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PredikitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Predikit.g4"

	return this
}

// PredikitParser tokens.
const (
	PredikitParserEOF           = antlr.TokenEOF
	PredikitParserPK_ALL        = 1
	PredikitParserPK_ANY        = 2
	PredikitParserPK_NONE       = 3
	PredikitParserPK_TEST       = 4
	PredikitParserPK_NOT        = 5
	PredikitParserPK_TOOL       = 6
	PredikitParserPK_RETRYING   = 7
	PredikitParserPK_COLON      = 8
	PredikitParserPK_DOLLAR     = 9
	PredikitParserPK_LCURLY     = 10
	PredikitParserPK_RCURLY     = 11
	PredikitParserPK_LPAREN     = 12
	PredikitParserPK_RPAREN     = 13
	PredikitParserPK_TRUE       = 14
	PredikitParserPK_FALSE      = 15
	PredikitParserPK_CMP_EQ     = 16
	PredikitParserPK_CMP_NEQ    = 17
	PredikitParserPK_CMP_RE     = 18
	PredikitParserPK_CMP_GT     = 19
	PredikitParserPK_CMP_GTE    = 20
	PredikitParserPK_CMP_LT     = 21
	PredikitParserPK_CMP_LTE    = 22
	PredikitParserPK_FN_LIT     = 23
	PredikitParserPK_STRING_LIT = 24
	PredikitParserPK_ID         = 25
	PredikitParserPK_TYPENAME   = 26
	PredikitParserPK_INT        = 27
	PredikitParserLINE_COMMENT  = 28
	PredikitParserCOMMENT       = 29
	PredikitParserWS            = 30
)

// PredikitParser rules.
const (
	PredikitParserRULE_pk_toplevel                = 0
	PredikitParserRULE_pk_toplevel_child          = 1
	PredikitParserRULE_pk_group                   = 2
	PredikitParserRULE_pk_group_child             = 3
	PredikitParserRULE_pk_group_agg               = 4
	PredikitParserRULE_pk_test                    = 5
	PredikitParserRULE_pk_tool                    = 6
	PredikitParserRULE_pk_tool_child              = 7
	PredikitParserRULE_pk_tool_metaparam          = 8
	PredikitParserRULE_pk_actual_param            = 9
	PredikitParserRULE_pk_actual_param_value      = 10
	PredikitParserRULE_pk_tool_actual_param       = 11
	PredikitParserRULE_pk_tool_actual_param_value = 12
	PredikitParserRULE_pk_conversion_fn           = 13
	PredikitParserRULE_pk_bool                    = 14
)

// IPk_toplevelContext is an interface to support dynamic dispatch.
type IPk_toplevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_pk_toplevel_child returns the _pk_toplevel_child rule contexts.
	Get_pk_toplevel_child() IPk_toplevel_childContext

	// Set_pk_toplevel_child sets the _pk_toplevel_child rule contexts.
	Set_pk_toplevel_child(IPk_toplevel_childContext)

	// GetKids returns the kids rule context list.
	GetKids() []IPk_toplevel_childContext

	// SetKids sets the kids rule context list.
	SetKids([]IPk_toplevel_childContext)

	// Getter signatures
	AllPk_toplevel_child() []IPk_toplevel_childContext
	Pk_toplevel_child(i int) IPk_toplevel_childContext

	// IsPk_toplevelContext differentiates from other interfaces.
	IsPk_toplevelContext()
}

type Pk_toplevelContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	_pk_toplevel_child IPk_toplevel_childContext
	kids               []IPk_toplevel_childContext
}

func NewEmptyPk_toplevelContext() *Pk_toplevelContext {
	var p = new(Pk_toplevelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_toplevel
	return p
}

func InitEmptyPk_toplevelContext(p *Pk_toplevelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_toplevel
}

func (*Pk_toplevelContext) IsPk_toplevelContext() {}

func NewPk_toplevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_toplevelContext {
	var p = new(Pk_toplevelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_toplevel

	return p
}

func (s *Pk_toplevelContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_toplevelContext) Get_pk_toplevel_child() IPk_toplevel_childContext {
	return s._pk_toplevel_child
}

func (s *Pk_toplevelContext) Set_pk_toplevel_child(v IPk_toplevel_childContext) {
	s._pk_toplevel_child = v
}

func (s *Pk_toplevelContext) GetKids() []IPk_toplevel_childContext { return s.kids }

func (s *Pk_toplevelContext) SetKids(v []IPk_toplevel_childContext) { s.kids = v }

func (s *Pk_toplevelContext) AllPk_toplevel_child() []IPk_toplevel_childContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPk_toplevel_childContext); ok {
			len++
		}
	}

	tst := make([]IPk_toplevel_childContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPk_toplevel_childContext); ok {
			tst[i] = t.(IPk_toplevel_childContext)
			i++
		}
	}

	return tst
}

func (s *Pk_toplevelContext) Pk_toplevel_child(i int) IPk_toplevel_childContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_toplevel_childContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_toplevel_childContext)
}

func (s *Pk_toplevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_toplevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_toplevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_toplevel(s)
	}
}

func (s *Pk_toplevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_toplevel(s)
	}
}

func (p *PredikitParser) Pk_toplevel() (localctx IPk_toplevelContext) {
	localctx = NewPk_toplevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PredikitParserRULE_pk_toplevel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&222) != 0) {
		{
			p.SetState(30)

			var _x = p.Pk_toplevel_child()

			localctx.(*Pk_toplevelContext)._pk_toplevel_child = _x
		}
		localctx.(*Pk_toplevelContext).kids = append(localctx.(*Pk_toplevelContext).kids, localctx.(*Pk_toplevelContext)._pk_toplevel_child)

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_toplevel_childContext is an interface to support dynamic dispatch.
type IPk_toplevel_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGroup returns the group rule contexts.
	GetGroup() IPk_groupContext

	// GetTest returns the test rule contexts.
	GetTest() IPk_testContext

	// GetTool returns the tool rule contexts.
	GetTool() IPk_toolContext

	// SetGroup sets the group rule contexts.
	SetGroup(IPk_groupContext)

	// SetTest sets the test rule contexts.
	SetTest(IPk_testContext)

	// SetTool sets the tool rule contexts.
	SetTool(IPk_toolContext)

	// Getter signatures
	Pk_group() IPk_groupContext
	Pk_test() IPk_testContext
	Pk_tool() IPk_toolContext

	// IsPk_toplevel_childContext differentiates from other interfaces.
	IsPk_toplevel_childContext()
}

type Pk_toplevel_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	group  IPk_groupContext
	test   IPk_testContext
	tool   IPk_toolContext
}

func NewEmptyPk_toplevel_childContext() *Pk_toplevel_childContext {
	var p = new(Pk_toplevel_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_toplevel_child
	return p
}

func InitEmptyPk_toplevel_childContext(p *Pk_toplevel_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_toplevel_child
}

func (*Pk_toplevel_childContext) IsPk_toplevel_childContext() {}

func NewPk_toplevel_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_toplevel_childContext {
	var p = new(Pk_toplevel_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_toplevel_child

	return p
}

func (s *Pk_toplevel_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_toplevel_childContext) GetGroup() IPk_groupContext { return s.group }

func (s *Pk_toplevel_childContext) GetTest() IPk_testContext { return s.test }

func (s *Pk_toplevel_childContext) GetTool() IPk_toolContext { return s.tool }

func (s *Pk_toplevel_childContext) SetGroup(v IPk_groupContext) { s.group = v }

func (s *Pk_toplevel_childContext) SetTest(v IPk_testContext) { s.test = v }

func (s *Pk_toplevel_childContext) SetTool(v IPk_toolContext) { s.tool = v }

func (s *Pk_toplevel_childContext) Pk_group() IPk_groupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_groupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_groupContext)
}

func (s *Pk_toplevel_childContext) Pk_test() IPk_testContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_testContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_testContext)
}

func (s *Pk_toplevel_childContext) Pk_tool() IPk_toolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_toolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_toolContext)
}

func (s *Pk_toplevel_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_toplevel_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_toplevel_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_toplevel_child(s)
	}
}

func (s *Pk_toplevel_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_toplevel_child(s)
	}
}

func (p *PredikitParser) Pk_toplevel_child() (localctx IPk_toplevel_childContext) {
	localctx = NewPk_toplevel_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PredikitParserRULE_pk_toplevel_child)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PredikitParserPK_ALL, PredikitParserPK_ANY, PredikitParserPK_NONE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)

			var _x = p.Pk_group()

			localctx.(*Pk_toplevel_childContext).group = _x
		}

	case PredikitParserPK_TEST, PredikitParserPK_RETRYING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)

			var _x = p.Pk_test()

			localctx.(*Pk_toplevel_childContext).test = _x
		}

	case PredikitParserPK_TOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)

			var _x = p.Pk_tool()

			localctx.(*Pk_toplevel_childContext).tool = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_groupContext is an interface to support dynamic dispatch.
type IPk_groupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAgg_fn returns the agg_fn rule contexts.
	GetAgg_fn() IPk_group_aggContext

	// Get_pk_group_child returns the _pk_group_child rule contexts.
	Get_pk_group_child() IPk_group_childContext

	// SetAgg_fn sets the agg_fn rule contexts.
	SetAgg_fn(IPk_group_aggContext)

	// Set_pk_group_child sets the _pk_group_child rule contexts.
	Set_pk_group_child(IPk_group_childContext)

	// GetGroup_children returns the group_children rule context list.
	GetGroup_children() []IPk_group_childContext

	// SetGroup_children sets the group_children rule context list.
	SetGroup_children([]IPk_group_childContext)

	// Getter signatures
	PK_LCURLY() antlr.TerminalNode
	PK_RCURLY() antlr.TerminalNode
	Pk_group_agg() IPk_group_aggContext
	AllPk_group_child() []IPk_group_childContext
	Pk_group_child(i int) IPk_group_childContext

	// IsPk_groupContext differentiates from other interfaces.
	IsPk_groupContext()
}

type Pk_groupContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	agg_fn          IPk_group_aggContext
	_pk_group_child IPk_group_childContext
	group_children  []IPk_group_childContext
}

func NewEmptyPk_groupContext() *Pk_groupContext {
	var p = new(Pk_groupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group
	return p
}

func InitEmptyPk_groupContext(p *Pk_groupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group
}

func (*Pk_groupContext) IsPk_groupContext() {}

func NewPk_groupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_groupContext {
	var p = new(Pk_groupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_group

	return p
}

func (s *Pk_groupContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_groupContext) GetAgg_fn() IPk_group_aggContext { return s.agg_fn }

func (s *Pk_groupContext) Get_pk_group_child() IPk_group_childContext { return s._pk_group_child }

func (s *Pk_groupContext) SetAgg_fn(v IPk_group_aggContext) { s.agg_fn = v }

func (s *Pk_groupContext) Set_pk_group_child(v IPk_group_childContext) { s._pk_group_child = v }

func (s *Pk_groupContext) GetGroup_children() []IPk_group_childContext { return s.group_children }

func (s *Pk_groupContext) SetGroup_children(v []IPk_group_childContext) { s.group_children = v }

func (s *Pk_groupContext) PK_LCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_LCURLY, 0)
}

func (s *Pk_groupContext) PK_RCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_RCURLY, 0)
}

func (s *Pk_groupContext) Pk_group_agg() IPk_group_aggContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_group_aggContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_group_aggContext)
}

func (s *Pk_groupContext) AllPk_group_child() []IPk_group_childContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPk_group_childContext); ok {
			len++
		}
	}

	tst := make([]IPk_group_childContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPk_group_childContext); ok {
			tst[i] = t.(IPk_group_childContext)
			i++
		}
	}

	return tst
}

func (s *Pk_groupContext) Pk_group_child(i int) IPk_group_childContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_group_childContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_group_childContext)
}

func (s *Pk_groupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_groupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_groupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_group(s)
	}
}

func (s *Pk_groupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_group(s)
	}
}

func (p *PredikitParser) Pk_group() (localctx IPk_groupContext) {
	localctx = NewPk_groupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PredikitParserRULE_pk_group)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)

		var _x = p.Pk_group_agg()

		localctx.(*Pk_groupContext).agg_fn = _x
	}
	{
		p.SetState(41)
		p.Match(PredikitParserPK_LCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33554590) != 0) {
		{
			p.SetState(42)

			var _x = p.Pk_group_child()

			localctx.(*Pk_groupContext)._pk_group_child = _x
		}
		localctx.(*Pk_groupContext).group_children = append(localctx.(*Pk_groupContext).group_children, localctx.(*Pk_groupContext)._pk_group_child)

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(PredikitParserPK_RCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_group_childContext is an interface to support dynamic dispatch.
type IPk_group_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTest returns the test rule contexts.
	GetTest() IPk_testContext

	// GetGroup returns the group rule contexts.
	GetGroup() IPk_groupContext

	// GetActual_param returns the actual_param rule contexts.
	GetActual_param() IPk_actual_paramContext

	// SetTest sets the test rule contexts.
	SetTest(IPk_testContext)

	// SetGroup sets the group rule contexts.
	SetGroup(IPk_groupContext)

	// SetActual_param sets the actual_param rule contexts.
	SetActual_param(IPk_actual_paramContext)

	// Getter signatures
	Pk_test() IPk_testContext
	Pk_group() IPk_groupContext
	Pk_actual_param() IPk_actual_paramContext

	// IsPk_group_childContext differentiates from other interfaces.
	IsPk_group_childContext()
}

type Pk_group_childContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	test         IPk_testContext
	group        IPk_groupContext
	actual_param IPk_actual_paramContext
}

func NewEmptyPk_group_childContext() *Pk_group_childContext {
	var p = new(Pk_group_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group_child
	return p
}

func InitEmptyPk_group_childContext(p *Pk_group_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group_child
}

func (*Pk_group_childContext) IsPk_group_childContext() {}

func NewPk_group_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_group_childContext {
	var p = new(Pk_group_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_group_child

	return p
}

func (s *Pk_group_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_group_childContext) GetTest() IPk_testContext { return s.test }

func (s *Pk_group_childContext) GetGroup() IPk_groupContext { return s.group }

func (s *Pk_group_childContext) GetActual_param() IPk_actual_paramContext { return s.actual_param }

func (s *Pk_group_childContext) SetTest(v IPk_testContext) { s.test = v }

func (s *Pk_group_childContext) SetGroup(v IPk_groupContext) { s.group = v }

func (s *Pk_group_childContext) SetActual_param(v IPk_actual_paramContext) { s.actual_param = v }

func (s *Pk_group_childContext) Pk_test() IPk_testContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_testContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_testContext)
}

func (s *Pk_group_childContext) Pk_group() IPk_groupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_groupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_groupContext)
}

func (s *Pk_group_childContext) Pk_actual_param() IPk_actual_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_actual_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_actual_paramContext)
}

func (s *Pk_group_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_group_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_group_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_group_child(s)
	}
}

func (s *Pk_group_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_group_child(s)
	}
}

func (p *PredikitParser) Pk_group_child() (localctx IPk_group_childContext) {
	localctx = NewPk_group_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PredikitParserRULE_pk_group_child)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PredikitParserPK_TEST, PredikitParserPK_RETRYING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _x = p.Pk_test()

			localctx.(*Pk_group_childContext).test = _x
		}

	case PredikitParserPK_ALL, PredikitParserPK_ANY, PredikitParserPK_NONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _x = p.Pk_group()

			localctx.(*Pk_group_childContext).group = _x
		}

	case PredikitParserPK_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)

			var _x = p.Pk_actual_param()

			localctx.(*Pk_group_childContext).actual_param = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_group_aggContext is an interface to support dynamic dispatch.
type IPk_group_aggContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PK_ALL() antlr.TerminalNode
	PK_ANY() antlr.TerminalNode
	PK_NONE() antlr.TerminalNode

	// IsPk_group_aggContext differentiates from other interfaces.
	IsPk_group_aggContext()
}

type Pk_group_aggContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPk_group_aggContext() *Pk_group_aggContext {
	var p = new(Pk_group_aggContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group_agg
	return p
}

func InitEmptyPk_group_aggContext(p *Pk_group_aggContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_group_agg
}

func (*Pk_group_aggContext) IsPk_group_aggContext() {}

func NewPk_group_aggContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_group_aggContext {
	var p = new(Pk_group_aggContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_group_agg

	return p
}

func (s *Pk_group_aggContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_group_aggContext) PK_ALL() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ALL, 0)
}

func (s *Pk_group_aggContext) PK_ANY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ANY, 0)
}

func (s *Pk_group_aggContext) PK_NONE() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_NONE, 0)
}

func (s *Pk_group_aggContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_group_aggContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_group_aggContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_group_agg(s)
	}
}

func (s *Pk_group_aggContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_group_agg(s)
	}
}

func (p *PredikitParser) Pk_group_agg() (localctx IPk_group_aggContext) {
	localctx = NewPk_group_aggContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PredikitParserRULE_pk_group_agg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_testContext is an interface to support dynamic dispatch.
type IPk_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTestname returns the testname token.
	GetTestname() antlr.Token

	// SetTestname sets the testname token.
	SetTestname(antlr.Token)

	// Get_pk_actual_param returns the _pk_actual_param rule contexts.
	Get_pk_actual_param() IPk_actual_paramContext

	// Set_pk_actual_param sets the _pk_actual_param rule contexts.
	Set_pk_actual_param(IPk_actual_paramContext)

	// GetAps returns the aps rule context list.
	GetAps() []IPk_actual_paramContext

	// SetAps sets the aps rule context list.
	SetAps([]IPk_actual_paramContext)

	// Getter signatures
	PK_TEST() antlr.TerminalNode
	PK_LCURLY() antlr.TerminalNode
	PK_RCURLY() antlr.TerminalNode
	PK_ID() antlr.TerminalNode
	PK_RETRYING() antlr.TerminalNode
	PK_NOT() antlr.TerminalNode
	AllPk_actual_param() []IPk_actual_paramContext
	Pk_actual_param(i int) IPk_actual_paramContext

	// IsPk_testContext differentiates from other interfaces.
	IsPk_testContext()
}

type Pk_testContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	testname         antlr.Token
	_pk_actual_param IPk_actual_paramContext
	aps              []IPk_actual_paramContext
}

func NewEmptyPk_testContext() *Pk_testContext {
	var p = new(Pk_testContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_test
	return p
}

func InitEmptyPk_testContext(p *Pk_testContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_test
}

func (*Pk_testContext) IsPk_testContext() {}

func NewPk_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_testContext {
	var p = new(Pk_testContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_test

	return p
}

func (s *Pk_testContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_testContext) GetTestname() antlr.Token { return s.testname }

func (s *Pk_testContext) SetTestname(v antlr.Token) { s.testname = v }

func (s *Pk_testContext) Get_pk_actual_param() IPk_actual_paramContext { return s._pk_actual_param }

func (s *Pk_testContext) Set_pk_actual_param(v IPk_actual_paramContext) { s._pk_actual_param = v }

func (s *Pk_testContext) GetAps() []IPk_actual_paramContext { return s.aps }

func (s *Pk_testContext) SetAps(v []IPk_actual_paramContext) { s.aps = v }

func (s *Pk_testContext) PK_TEST() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_TEST, 0)
}

func (s *Pk_testContext) PK_LCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_LCURLY, 0)
}

func (s *Pk_testContext) PK_RCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_RCURLY, 0)
}

func (s *Pk_testContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_testContext) PK_RETRYING() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_RETRYING, 0)
}

func (s *Pk_testContext) PK_NOT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_NOT, 0)
}

func (s *Pk_testContext) AllPk_actual_param() []IPk_actual_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPk_actual_paramContext); ok {
			len++
		}
	}

	tst := make([]IPk_actual_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPk_actual_paramContext); ok {
			tst[i] = t.(IPk_actual_paramContext)
			i++
		}
	}

	return tst
}

func (s *Pk_testContext) Pk_actual_param(i int) IPk_actual_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_actual_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_actual_paramContext)
}

func (s *Pk_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_test(s)
	}
}

func (s *Pk_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_test(s)
	}
}

func (p *PredikitParser) Pk_test() (localctx IPk_testContext) {
	localctx = NewPk_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PredikitParserRULE_pk_test)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PredikitParserPK_RETRYING {
		{
			p.SetState(56)
			p.Match(PredikitParserPK_RETRYING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(59)
		p.Match(PredikitParserPK_TEST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PredikitParserPK_NOT {
		{
			p.SetState(60)
			p.Match(PredikitParserPK_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(63)

		var _m = p.Match(PredikitParserPK_ID)

		localctx.(*Pk_testContext).testname = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(PredikitParserPK_LCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PredikitParserPK_ID {
		{
			p.SetState(65)

			var _x = p.Pk_actual_param()

			localctx.(*Pk_testContext)._pk_actual_param = _x
		}
		localctx.(*Pk_testContext).aps = append(localctx.(*Pk_testContext).aps, localctx.(*Pk_testContext)._pk_actual_param)

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(PredikitParserPK_RCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_toolContext is an interface to support dynamic dispatch.
type IPk_toolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTool_name returns the tool_name token.
	GetTool_name() antlr.Token

	// SetTool_name sets the tool_name token.
	SetTool_name(antlr.Token)

	// Get_pk_tool_child returns the _pk_tool_child rule contexts.
	Get_pk_tool_child() IPk_tool_childContext

	// Set_pk_tool_child sets the _pk_tool_child rule contexts.
	Set_pk_tool_child(IPk_tool_childContext)

	// GetKids returns the kids rule context list.
	GetKids() []IPk_tool_childContext

	// SetKids sets the kids rule context list.
	SetKids([]IPk_tool_childContext)

	// Getter signatures
	PK_TOOL() antlr.TerminalNode
	PK_LCURLY() antlr.TerminalNode
	PK_RCURLY() antlr.TerminalNode
	PK_ID() antlr.TerminalNode
	AllPk_tool_child() []IPk_tool_childContext
	Pk_tool_child(i int) IPk_tool_childContext

	// IsPk_toolContext differentiates from other interfaces.
	IsPk_toolContext()
}

type Pk_toolContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	tool_name      antlr.Token
	_pk_tool_child IPk_tool_childContext
	kids           []IPk_tool_childContext
}

func NewEmptyPk_toolContext() *Pk_toolContext {
	var p = new(Pk_toolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool
	return p
}

func InitEmptyPk_toolContext(p *Pk_toolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool
}

func (*Pk_toolContext) IsPk_toolContext() {}

func NewPk_toolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_toolContext {
	var p = new(Pk_toolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_tool

	return p
}

func (s *Pk_toolContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_toolContext) GetTool_name() antlr.Token { return s.tool_name }

func (s *Pk_toolContext) SetTool_name(v antlr.Token) { s.tool_name = v }

func (s *Pk_toolContext) Get_pk_tool_child() IPk_tool_childContext { return s._pk_tool_child }

func (s *Pk_toolContext) Set_pk_tool_child(v IPk_tool_childContext) { s._pk_tool_child = v }

func (s *Pk_toolContext) GetKids() []IPk_tool_childContext { return s.kids }

func (s *Pk_toolContext) SetKids(v []IPk_tool_childContext) { s.kids = v }

func (s *Pk_toolContext) PK_TOOL() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_TOOL, 0)
}

func (s *Pk_toolContext) PK_LCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_LCURLY, 0)
}

func (s *Pk_toolContext) PK_RCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_RCURLY, 0)
}

func (s *Pk_toolContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_toolContext) AllPk_tool_child() []IPk_tool_childContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPk_tool_childContext); ok {
			len++
		}
	}

	tst := make([]IPk_tool_childContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPk_tool_childContext); ok {
			tst[i] = t.(IPk_tool_childContext)
			i++
		}
	}

	return tst
}

func (s *Pk_toolContext) Pk_tool_child(i int) IPk_tool_childContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_tool_childContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_tool_childContext)
}

func (s *Pk_toolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_toolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_toolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_tool(s)
	}
}

func (s *Pk_toolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_tool(s)
	}
}

func (p *PredikitParser) Pk_tool() (localctx IPk_toolContext) {
	localctx = NewPk_toolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PredikitParserRULE_pk_tool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(PredikitParserPK_TOOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)

		var _m = p.Match(PredikitParserPK_ID)

		localctx.(*Pk_toolContext).tool_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(PredikitParserPK_LCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PredikitParserPK_DOLLAR || _la == PredikitParserPK_ID {
		{
			p.SetState(75)

			var _x = p.Pk_tool_child()

			localctx.(*Pk_toolContext)._pk_tool_child = _x
		}
		localctx.(*Pk_toolContext).kids = append(localctx.(*Pk_toolContext).kids, localctx.(*Pk_toolContext)._pk_tool_child)

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(PredikitParserPK_RCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_tool_childContext is an interface to support dynamic dispatch.
type IPk_tool_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pk_actual_param() IPk_actual_paramContext
	Pk_tool_metaparam() IPk_tool_metaparamContext

	// IsPk_tool_childContext differentiates from other interfaces.
	IsPk_tool_childContext()
}

type Pk_tool_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPk_tool_childContext() *Pk_tool_childContext {
	var p = new(Pk_tool_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_child
	return p
}

func InitEmptyPk_tool_childContext(p *Pk_tool_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_child
}

func (*Pk_tool_childContext) IsPk_tool_childContext() {}

func NewPk_tool_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_tool_childContext {
	var p = new(Pk_tool_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_tool_child

	return p
}

func (s *Pk_tool_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_tool_childContext) Pk_actual_param() IPk_actual_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_actual_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_actual_paramContext)
}

func (s *Pk_tool_childContext) Pk_tool_metaparam() IPk_tool_metaparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_tool_metaparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_tool_metaparamContext)
}

func (s *Pk_tool_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_tool_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_tool_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_tool_child(s)
	}
}

func (s *Pk_tool_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_tool_child(s)
	}
}

func (p *PredikitParser) Pk_tool_child() (localctx IPk_tool_childContext) {
	localctx = NewPk_tool_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PredikitParserRULE_pk_tool_child)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PredikitParserPK_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Pk_actual_param()
		}

	case PredikitParserPK_DOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Pk_tool_metaparam()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_tool_metaparamContext is an interface to support dynamic dispatch.
type IPk_tool_metaparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTool_param_name returns the tool_param_name token.
	GetTool_param_name() antlr.Token

	// SetTool_param_name sets the tool_param_name token.
	SetTool_param_name(antlr.Token)

	// Get_pk_tool_actual_param returns the _pk_tool_actual_param rule contexts.
	Get_pk_tool_actual_param() IPk_tool_actual_paramContext

	// Set_pk_tool_actual_param sets the _pk_tool_actual_param rule contexts.
	Set_pk_tool_actual_param(IPk_tool_actual_paramContext)

	// GetTool_actual_params returns the tool_actual_params rule context list.
	GetTool_actual_params() []IPk_tool_actual_paramContext

	// SetTool_actual_params sets the tool_actual_params rule context list.
	SetTool_actual_params([]IPk_tool_actual_paramContext)

	// Getter signatures
	PK_DOLLAR() antlr.TerminalNode
	PK_LCURLY() antlr.TerminalNode
	PK_RCURLY() antlr.TerminalNode
	PK_ID() antlr.TerminalNode
	AllPk_tool_actual_param() []IPk_tool_actual_paramContext
	Pk_tool_actual_param(i int) IPk_tool_actual_paramContext

	// IsPk_tool_metaparamContext differentiates from other interfaces.
	IsPk_tool_metaparamContext()
}

type Pk_tool_metaparamContext struct {
	antlr.BaseParserRuleContext
	parser                antlr.Parser
	tool_param_name       antlr.Token
	_pk_tool_actual_param IPk_tool_actual_paramContext
	tool_actual_params    []IPk_tool_actual_paramContext
}

func NewEmptyPk_tool_metaparamContext() *Pk_tool_metaparamContext {
	var p = new(Pk_tool_metaparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_metaparam
	return p
}

func InitEmptyPk_tool_metaparamContext(p *Pk_tool_metaparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_metaparam
}

func (*Pk_tool_metaparamContext) IsPk_tool_metaparamContext() {}

func NewPk_tool_metaparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_tool_metaparamContext {
	var p = new(Pk_tool_metaparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_tool_metaparam

	return p
}

func (s *Pk_tool_metaparamContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_tool_metaparamContext) GetTool_param_name() antlr.Token { return s.tool_param_name }

func (s *Pk_tool_metaparamContext) SetTool_param_name(v antlr.Token) { s.tool_param_name = v }

func (s *Pk_tool_metaparamContext) Get_pk_tool_actual_param() IPk_tool_actual_paramContext {
	return s._pk_tool_actual_param
}

func (s *Pk_tool_metaparamContext) Set_pk_tool_actual_param(v IPk_tool_actual_paramContext) {
	s._pk_tool_actual_param = v
}

func (s *Pk_tool_metaparamContext) GetTool_actual_params() []IPk_tool_actual_paramContext {
	return s.tool_actual_params
}

func (s *Pk_tool_metaparamContext) SetTool_actual_params(v []IPk_tool_actual_paramContext) {
	s.tool_actual_params = v
}

func (s *Pk_tool_metaparamContext) PK_DOLLAR() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_DOLLAR, 0)
}

func (s *Pk_tool_metaparamContext) PK_LCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_LCURLY, 0)
}

func (s *Pk_tool_metaparamContext) PK_RCURLY() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_RCURLY, 0)
}

func (s *Pk_tool_metaparamContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_tool_metaparamContext) AllPk_tool_actual_param() []IPk_tool_actual_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPk_tool_actual_paramContext); ok {
			len++
		}
	}

	tst := make([]IPk_tool_actual_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPk_tool_actual_paramContext); ok {
			tst[i] = t.(IPk_tool_actual_paramContext)
			i++
		}
	}

	return tst
}

func (s *Pk_tool_metaparamContext) Pk_tool_actual_param(i int) IPk_tool_actual_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_tool_actual_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_tool_actual_paramContext)
}

func (s *Pk_tool_metaparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_tool_metaparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_tool_metaparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_tool_metaparam(s)
	}
}

func (s *Pk_tool_metaparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_tool_metaparam(s)
	}
}

func (p *PredikitParser) Pk_tool_metaparam() (localctx IPk_tool_metaparamContext) {
	localctx = NewPk_tool_metaparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PredikitParserRULE_pk_tool_metaparam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(PredikitParserPK_DOLLAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)

		var _m = p.Match(PredikitParserPK_ID)

		localctx.(*Pk_tool_metaparamContext).tool_param_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(PredikitParserPK_LCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PredikitParserPK_ID {
		{
			p.SetState(89)

			var _x = p.Pk_tool_actual_param()

			localctx.(*Pk_tool_metaparamContext)._pk_tool_actual_param = _x
		}
		localctx.(*Pk_tool_metaparamContext).tool_actual_params = append(localctx.(*Pk_tool_metaparamContext).tool_actual_params, localctx.(*Pk_tool_metaparamContext)._pk_tool_actual_param)

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(PredikitParserPK_RCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_actual_paramContext is an interface to support dynamic dispatch.
type IPk_actual_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParam_name returns the param_name token.
	GetParam_name() antlr.Token

	// SetParam_name sets the param_name token.
	SetParam_name(antlr.Token)

	// GetParam_value returns the param_value rule contexts.
	GetParam_value() IPk_actual_param_valueContext

	// SetParam_value sets the param_value rule contexts.
	SetParam_value(IPk_actual_param_valueContext)

	// Getter signatures
	PK_COLON() antlr.TerminalNode
	PK_ID() antlr.TerminalNode
	Pk_actual_param_value() IPk_actual_param_valueContext

	// IsPk_actual_paramContext differentiates from other interfaces.
	IsPk_actual_paramContext()
}

type Pk_actual_paramContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	param_name  antlr.Token
	param_value IPk_actual_param_valueContext
}

func NewEmptyPk_actual_paramContext() *Pk_actual_paramContext {
	var p = new(Pk_actual_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_actual_param
	return p
}

func InitEmptyPk_actual_paramContext(p *Pk_actual_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_actual_param
}

func (*Pk_actual_paramContext) IsPk_actual_paramContext() {}

func NewPk_actual_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_actual_paramContext {
	var p = new(Pk_actual_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_actual_param

	return p
}

func (s *Pk_actual_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_actual_paramContext) GetParam_name() antlr.Token { return s.param_name }

func (s *Pk_actual_paramContext) SetParam_name(v antlr.Token) { s.param_name = v }

func (s *Pk_actual_paramContext) GetParam_value() IPk_actual_param_valueContext { return s.param_value }

func (s *Pk_actual_paramContext) SetParam_value(v IPk_actual_param_valueContext) { s.param_value = v }

func (s *Pk_actual_paramContext) PK_COLON() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_COLON, 0)
}

func (s *Pk_actual_paramContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_actual_paramContext) Pk_actual_param_value() IPk_actual_param_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_actual_param_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_actual_param_valueContext)
}

func (s *Pk_actual_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_actual_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_actual_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_actual_param(s)
	}
}

func (s *Pk_actual_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_actual_param(s)
	}
}

func (p *PredikitParser) Pk_actual_param() (localctx IPk_actual_paramContext) {
	localctx = NewPk_actual_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PredikitParserRULE_pk_actual_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)

		var _m = p.Match(PredikitParserPK_ID)

		localctx.(*Pk_actual_paramContext).param_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(PredikitParserPK_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)

		var _x = p.Pk_actual_param_value()

		localctx.(*Pk_actual_paramContext).param_value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_actual_param_valueContext is an interface to support dynamic dispatch.
type IPk_actual_param_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVi returns the vi token.
	GetVi() antlr.Token

	// GetVs returns the vs token.
	GetVs() antlr.Token

	// SetVi sets the vi token.
	SetVi(antlr.Token)

	// SetVs sets the vs token.
	SetVs(antlr.Token)

	// GetVb returns the vb rule contexts.
	GetVb() IPk_boolContext

	// GetVc returns the vc rule contexts.
	GetVc() IPk_conversion_fnContext

	// SetVb sets the vb rule contexts.
	SetVb(IPk_boolContext)

	// SetVc sets the vc rule contexts.
	SetVc(IPk_conversion_fnContext)

	// Getter signatures
	PK_INT() antlr.TerminalNode
	PK_STRING_LIT() antlr.TerminalNode
	Pk_bool() IPk_boolContext
	Pk_conversion_fn() IPk_conversion_fnContext

	// IsPk_actual_param_valueContext differentiates from other interfaces.
	IsPk_actual_param_valueContext()
}

type Pk_actual_param_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	vi     antlr.Token
	vs     antlr.Token
	vb     IPk_boolContext
	vc     IPk_conversion_fnContext
}

func NewEmptyPk_actual_param_valueContext() *Pk_actual_param_valueContext {
	var p = new(Pk_actual_param_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_actual_param_value
	return p
}

func InitEmptyPk_actual_param_valueContext(p *Pk_actual_param_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_actual_param_value
}

func (*Pk_actual_param_valueContext) IsPk_actual_param_valueContext() {}

func NewPk_actual_param_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_actual_param_valueContext {
	var p = new(Pk_actual_param_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_actual_param_value

	return p
}

func (s *Pk_actual_param_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_actual_param_valueContext) GetVi() antlr.Token { return s.vi }

func (s *Pk_actual_param_valueContext) GetVs() antlr.Token { return s.vs }

func (s *Pk_actual_param_valueContext) SetVi(v antlr.Token) { s.vi = v }

func (s *Pk_actual_param_valueContext) SetVs(v antlr.Token) { s.vs = v }

func (s *Pk_actual_param_valueContext) GetVb() IPk_boolContext { return s.vb }

func (s *Pk_actual_param_valueContext) GetVc() IPk_conversion_fnContext { return s.vc }

func (s *Pk_actual_param_valueContext) SetVb(v IPk_boolContext) { s.vb = v }

func (s *Pk_actual_param_valueContext) SetVc(v IPk_conversion_fnContext) { s.vc = v }

func (s *Pk_actual_param_valueContext) PK_INT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_INT, 0)
}

func (s *Pk_actual_param_valueContext) PK_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_STRING_LIT, 0)
}

func (s *Pk_actual_param_valueContext) Pk_bool() IPk_boolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_boolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_boolContext)
}

func (s *Pk_actual_param_valueContext) Pk_conversion_fn() IPk_conversion_fnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_conversion_fnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_conversion_fnContext)
}

func (s *Pk_actual_param_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_actual_param_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_actual_param_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_actual_param_value(s)
	}
}

func (s *Pk_actual_param_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_actual_param_value(s)
	}
}

func (p *PredikitParser) Pk_actual_param_value() (localctx IPk_actual_param_valueContext) {
	localctx = NewPk_actual_param_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PredikitParserRULE_pk_actual_param_value)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PredikitParserPK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)

			var _m = p.Match(PredikitParserPK_INT)

			localctx.(*Pk_actual_param_valueContext).vi = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PredikitParserPK_STRING_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)

			var _m = p.Match(PredikitParserPK_STRING_LIT)

			localctx.(*Pk_actual_param_valueContext).vs = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PredikitParserPK_TRUE, PredikitParserPK_FALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)

			var _x = p.Pk_bool()

			localctx.(*Pk_actual_param_valueContext).vb = _x
		}

	case PredikitParserPK_ID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)

			var _x = p.Pk_conversion_fn()

			localctx.(*Pk_actual_param_valueContext).vc = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_tool_actual_paramContext is an interface to support dynamic dispatch.
type IPk_tool_actual_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParam_name returns the param_name token.
	GetParam_name() antlr.Token

	// SetParam_name sets the param_name token.
	SetParam_name(antlr.Token)

	// GetParam_value returns the param_value rule contexts.
	GetParam_value() IPk_tool_actual_param_valueContext

	// SetParam_value sets the param_value rule contexts.
	SetParam_value(IPk_tool_actual_param_valueContext)

	// Getter signatures
	PK_COLON() antlr.TerminalNode
	PK_ID() antlr.TerminalNode
	Pk_tool_actual_param_value() IPk_tool_actual_param_valueContext

	// IsPk_tool_actual_paramContext differentiates from other interfaces.
	IsPk_tool_actual_paramContext()
}

type Pk_tool_actual_paramContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	param_name  antlr.Token
	param_value IPk_tool_actual_param_valueContext
}

func NewEmptyPk_tool_actual_paramContext() *Pk_tool_actual_paramContext {
	var p = new(Pk_tool_actual_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param
	return p
}

func InitEmptyPk_tool_actual_paramContext(p *Pk_tool_actual_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param
}

func (*Pk_tool_actual_paramContext) IsPk_tool_actual_paramContext() {}

func NewPk_tool_actual_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_tool_actual_paramContext {
	var p = new(Pk_tool_actual_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param

	return p
}

func (s *Pk_tool_actual_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_tool_actual_paramContext) GetParam_name() antlr.Token { return s.param_name }

func (s *Pk_tool_actual_paramContext) SetParam_name(v antlr.Token) { s.param_name = v }

func (s *Pk_tool_actual_paramContext) GetParam_value() IPk_tool_actual_param_valueContext {
	return s.param_value
}

func (s *Pk_tool_actual_paramContext) SetParam_value(v IPk_tool_actual_param_valueContext) {
	s.param_value = v
}

func (s *Pk_tool_actual_paramContext) PK_COLON() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_COLON, 0)
}

func (s *Pk_tool_actual_paramContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_tool_actual_paramContext) Pk_tool_actual_param_value() IPk_tool_actual_param_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_tool_actual_param_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_tool_actual_param_valueContext)
}

func (s *Pk_tool_actual_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_tool_actual_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_tool_actual_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_tool_actual_param(s)
	}
}

func (s *Pk_tool_actual_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_tool_actual_param(s)
	}
}

func (p *PredikitParser) Pk_tool_actual_param() (localctx IPk_tool_actual_paramContext) {
	localctx = NewPk_tool_actual_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PredikitParserRULE_pk_tool_actual_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)

		var _m = p.Match(PredikitParserPK_ID)

		localctx.(*Pk_tool_actual_paramContext).param_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(PredikitParserPK_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)

		var _x = p.Pk_tool_actual_param_value()

		localctx.(*Pk_tool_actual_paramContext).param_value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_tool_actual_param_valueContext is an interface to support dynamic dispatch.
type IPk_tool_actual_param_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVi returns the vi token.
	GetVi() antlr.Token

	// GetVs returns the vs token.
	GetVs() antlr.Token

	// GetVtn returns the vtn token.
	GetVtn() antlr.Token

	// SetVi sets the vi token.
	SetVi(antlr.Token)

	// SetVs sets the vs token.
	SetVs(antlr.Token)

	// SetVtn sets the vtn token.
	SetVtn(antlr.Token)

	// GetVb returns the vb rule contexts.
	GetVb() IPk_boolContext

	// GetVc returns the vc rule contexts.
	GetVc() IPk_conversion_fnContext

	// SetVb sets the vb rule contexts.
	SetVb(IPk_boolContext)

	// SetVc sets the vc rule contexts.
	SetVc(IPk_conversion_fnContext)

	// Getter signatures
	PK_INT() antlr.TerminalNode
	PK_STRING_LIT() antlr.TerminalNode
	Pk_bool() IPk_boolContext
	Pk_conversion_fn() IPk_conversion_fnContext
	PK_TYPENAME() antlr.TerminalNode

	// IsPk_tool_actual_param_valueContext differentiates from other interfaces.
	IsPk_tool_actual_param_valueContext()
}

type Pk_tool_actual_param_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	vi     antlr.Token
	vs     antlr.Token
	vb     IPk_boolContext
	vc     IPk_conversion_fnContext
	vtn    antlr.Token
}

func NewEmptyPk_tool_actual_param_valueContext() *Pk_tool_actual_param_valueContext {
	var p = new(Pk_tool_actual_param_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param_value
	return p
}

func InitEmptyPk_tool_actual_param_valueContext(p *Pk_tool_actual_param_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param_value
}

func (*Pk_tool_actual_param_valueContext) IsPk_tool_actual_param_valueContext() {}

func NewPk_tool_actual_param_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_tool_actual_param_valueContext {
	var p = new(Pk_tool_actual_param_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_tool_actual_param_value

	return p
}

func (s *Pk_tool_actual_param_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_tool_actual_param_valueContext) GetVi() antlr.Token { return s.vi }

func (s *Pk_tool_actual_param_valueContext) GetVs() antlr.Token { return s.vs }

func (s *Pk_tool_actual_param_valueContext) GetVtn() antlr.Token { return s.vtn }

func (s *Pk_tool_actual_param_valueContext) SetVi(v antlr.Token) { s.vi = v }

func (s *Pk_tool_actual_param_valueContext) SetVs(v antlr.Token) { s.vs = v }

func (s *Pk_tool_actual_param_valueContext) SetVtn(v antlr.Token) { s.vtn = v }

func (s *Pk_tool_actual_param_valueContext) GetVb() IPk_boolContext { return s.vb }

func (s *Pk_tool_actual_param_valueContext) GetVc() IPk_conversion_fnContext { return s.vc }

func (s *Pk_tool_actual_param_valueContext) SetVb(v IPk_boolContext) { s.vb = v }

func (s *Pk_tool_actual_param_valueContext) SetVc(v IPk_conversion_fnContext) { s.vc = v }

func (s *Pk_tool_actual_param_valueContext) PK_INT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_INT, 0)
}

func (s *Pk_tool_actual_param_valueContext) PK_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_STRING_LIT, 0)
}

func (s *Pk_tool_actual_param_valueContext) Pk_bool() IPk_boolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_boolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_boolContext)
}

func (s *Pk_tool_actual_param_valueContext) Pk_conversion_fn() IPk_conversion_fnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPk_conversion_fnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPk_conversion_fnContext)
}

func (s *Pk_tool_actual_param_valueContext) PK_TYPENAME() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_TYPENAME, 0)
}

func (s *Pk_tool_actual_param_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_tool_actual_param_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_tool_actual_param_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_tool_actual_param_value(s)
	}
}

func (s *Pk_tool_actual_param_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_tool_actual_param_value(s)
	}
}

func (p *PredikitParser) Pk_tool_actual_param_value() (localctx IPk_tool_actual_param_valueContext) {
	localctx = NewPk_tool_actual_param_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PredikitParserRULE_pk_tool_actual_param_value)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PredikitParserPK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)

			var _m = p.Match(PredikitParserPK_INT)

			localctx.(*Pk_tool_actual_param_valueContext).vi = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PredikitParserPK_STRING_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)

			var _m = p.Match(PredikitParserPK_STRING_LIT)

			localctx.(*Pk_tool_actual_param_valueContext).vs = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PredikitParserPK_TRUE, PredikitParserPK_FALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)

			var _x = p.Pk_bool()

			localctx.(*Pk_tool_actual_param_valueContext).vb = _x
		}

	case PredikitParserPK_ID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)

			var _x = p.Pk_conversion_fn()

			localctx.(*Pk_tool_actual_param_valueContext).vc = _x
		}

	case PredikitParserPK_TYPENAME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)

			var _m = p.Match(PredikitParserPK_TYPENAME)

			localctx.(*Pk_tool_actual_param_valueContext).vtn = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_conversion_fnContext is an interface to support dynamic dispatch.
type IPk_conversion_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PK_ID() antlr.TerminalNode
	PK_FN_LIT() antlr.TerminalNode

	// IsPk_conversion_fnContext differentiates from other interfaces.
	IsPk_conversion_fnContext()
}

type Pk_conversion_fnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPk_conversion_fnContext() *Pk_conversion_fnContext {
	var p = new(Pk_conversion_fnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_conversion_fn
	return p
}

func InitEmptyPk_conversion_fnContext(p *Pk_conversion_fnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_conversion_fn
}

func (*Pk_conversion_fnContext) IsPk_conversion_fnContext() {}

func NewPk_conversion_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_conversion_fnContext {
	var p = new(Pk_conversion_fnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_conversion_fn

	return p
}

func (s *Pk_conversion_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_conversion_fnContext) PK_ID() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_ID, 0)
}

func (s *Pk_conversion_fnContext) PK_FN_LIT() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_FN_LIT, 0)
}

func (s *Pk_conversion_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_conversion_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_conversion_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_conversion_fn(s)
	}
}

func (s *Pk_conversion_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_conversion_fn(s)
	}
}

func (p *PredikitParser) Pk_conversion_fn() (localctx IPk_conversion_fnContext) {
	localctx = NewPk_conversion_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PredikitParserRULE_pk_conversion_fn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(PredikitParserPK_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(PredikitParserPK_FN_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPk_boolContext is an interface to support dynamic dispatch.
type IPk_boolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PK_TRUE() antlr.TerminalNode
	PK_FALSE() antlr.TerminalNode

	// IsPk_boolContext differentiates from other interfaces.
	IsPk_boolContext()
}

type Pk_boolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPk_boolContext() *Pk_boolContext {
	var p = new(Pk_boolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_bool
	return p
}

func InitEmptyPk_boolContext(p *Pk_boolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PredikitParserRULE_pk_bool
}

func (*Pk_boolContext) IsPk_boolContext() {}

func NewPk_boolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pk_boolContext {
	var p = new(Pk_boolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PredikitParserRULE_pk_bool

	return p
}

func (s *Pk_boolContext) GetParser() antlr.Parser { return s.parser }

func (s *Pk_boolContext) PK_TRUE() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_TRUE, 0)
}

func (s *Pk_boolContext) PK_FALSE() antlr.TerminalNode {
	return s.GetToken(PredikitParserPK_FALSE, 0)
}

func (s *Pk_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pk_boolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pk_boolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.EnterPk_bool(s)
	}
}

func (s *Pk_boolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PredikitListener); ok {
		listenerT.ExitPk_bool(s)
	}
}

func (p *PredikitParser) Pk_bool() (localctx IPk_boolContext) {
	localctx = NewPk_boolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PredikitParserRULE_pk_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PredikitParserPK_TRUE || _la == PredikitParserPK_FALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
