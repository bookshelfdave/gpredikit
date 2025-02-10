package comp

import (
	"fmt"

	//"github.com/bookshelfdave/gpredikit/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/bookshelfdave/gpredikit/parser"
)

type TreeShapeListener struct {
	*parser.BasePredikitListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// }

// VisitTerminal is called when a terminal node is visited.
// func (s *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// // VisitErrorNode is called when an error node is visited.
// func (s *TreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// // EnterEveryRule is called when any rule is entered.
// func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// // ExitEveryRule is called when any rule is exited.
// func (s *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPk_toplevel is called when production pk_toplevel is entered.
func (s *TreeShapeListener) EnterPk_toplevel(ctx *parser.Pk_toplevelContext) {
	fmt.Println("Enter toplevel")

}

// ExitPk_toplevel is called when production pk_toplevel is exited.
func (s *TreeShapeListener) ExitPk_toplevel(ctx *parser.Pk_toplevelContext) {
	fmt.Println("Exit toplevel")
}

// func (s *TreeShapeListener) EnterPk_toplevel_child(ctx *parser.Pk_toplevel_childContext) {}

// func (s *TreeShapeListener) ExitPk_toplevel_child(ctx *parser.Pk_toplevel_childContext) {}

func (s *TreeShapeListener) EnterPk_group(ctx *parser.Pk_groupContext) {
	fmt.Println("Enter group")
}

func (s *TreeShapeListener) ExitPk_group(ctx *parser.Pk_groupContext) {
	fmt.Println("Exit group")
}

// func (s *TreeShapeListener) EnterPk_group_child(ctx *parser.Pk_group_childContext) {}

// func (s *TreeShapeListener) ExitPk_group_child(ctx *parser.Pk_group_childContext) {}

// func (s *TreeShapeListener) EnterPk_group_agg(ctx *parser.Pk_group_aggContext) {}

// func (s *TreeShapeListener) ExitPk_group_agg(ctx *parser.Pk_group_aggContext) {}

func (s *TreeShapeListener) EnterPk_test(ctx *parser.Pk_testContext) {
	fmt.Println("Enter test")
}

func (s *TreeShapeListener) ExitPk_test(ctx *parser.Pk_testContext) {
	fmt.Println("Exit test")
}

// func (s *TreeShapeListener) EnterPk_tool(ctx *parser.Pk_toolContext) {}

// func (s *TreeShapeListener) ExitPk_tool(ctx *parser.Pk_toolContext) {}

// func (s *TreeShapeListener) EnterPk_tool_child(ctx *parser.Pk_tool_childContext) {}

// func (s *TreeShapeListener) ExitPk_tool_child(ctx *parser.Pk_tool_childContext) {}

// func (s *TreeShapeListener) EnterPk_tool_metaparam(ctx *parser.Pk_tool_metaparamContext) {}

// func (s *TreeShapeListener) ExitPk_tool_metaparam(ctx *parser.Pk_tool_metaparamContext) {}

// func (s *TreeShapeListener) EnterPk_actual_param(ctx *parser.Pk_actual_paramContext) {}

// func (s *TreeShapeListener) ExitPk_actual_param(ctx *parser.Pk_actual_paramContext) {}

// func (s *TreeShapeListener) EnterPk_actual_param_value(ctx *parser.Pk_actual_param_valueContext) {}

// func (s *TreeShapeListener) ExitPk_actual_param_value(ctx *parser.Pk_actual_param_valueContext) {}

// func (s *TreeShapeListener) EnterPk_bool(ctx *parser.Pk_boolContext) {}

// func (s *TreeShapeListener) ExitPk_bool(ctx *parser.Pk_boolContext) {}

func Run() {
	fmt.Println("Compiling...")

	input, _ := antlr.NewFileStream("./checks/first.pk")
	lexer := parser.NewPredikitLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPredikitParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Pk_toplevel()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	//p.AddErrorListener(

}
