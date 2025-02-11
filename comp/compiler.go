package comp

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/bookshelfdave/gpredikit/parser"
)

type ParseTreeProperty = map[any]any

type TreeShapeListener struct {
	*parser.BasePredikitListener
	TreeProps ParseTreeProperty
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{TreeProps: make(map[any]any)}
}

// func (thitsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// }

// func (tsl *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}
// func (tsl *TreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}
// func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}
// func (tsl *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (tsl *TreeShapeListener) EnterPk_toplevel(ctx *parser.Pk_toplevelContext) {
	fmt.Println("Enter toplevel")

}

func (tsl *TreeShapeListener) ExitPk_toplevel(ctx *parser.Pk_toplevelContext) {
	fmt.Println("Exit toplevel")
}

// func (tsl *TreeShapeListener) ExitPk_toplevel_child(ctx *parser.Pk_toplevel_childContext) {}

func (tsl *TreeShapeListener) ExitPk_group(ctx *parser.Pk_groupContext) {
	agg_fn := ctx.GetAgg_fn().GetText()
	fmt.Printf("AGG FN = [%s]\n", agg_fn)
	_ = ctx.GetGroup_children()
}

func (tsl *TreeShapeListener) ExitPk_group_child(ctx *parser.Pk_group_childContext) {
	// test | group | actual_param
}

// func (tsl *TreeShapeListener) ExitPk_group_agg(ctx *parser.Pk_group_aggContext) {}

func (tsl *TreeShapeListener) ExitPk_test(ctx *parser.Pk_testContext) {
	aps := ctx.GetAps()
	actualParams := make([]*ActualParam, len(aps))
	for _, ap := range aps {
		v := tsl.TreeProps[ap]
		app := v.(*ActualParam)
		actualParams = append(actualParams, app)
	}

	fnname := ctx.GetTestname().GetText()
	is_negated := ctx.PK_NOT() != nil
	is_retrying := ctx.PK_RETRYING() != nil
	is_group := false
	cd := AstCheckDef{
		FnName:         fnname,
		ActualParams:   actualParams,
		Children:       []*AstCheckDef{},
		ContentAddress: "",
		IsNegated:      is_negated,
		IsRetrying:     is_retrying,
		IsGroup:        is_group,
	}
	fmt.Printf("\nCheck %+v\n", cd)
	tsl.TreeProps[ctx] = cd
}

// func (tsl *TreeShapeListener) ExitPk_tool(ctx *parser.Pk_toolContext) {}

// func (tsl *TreeShapeListener) ExitPk_tool_child(ctx *parser.Pk_tool_childContext) {}

// func (tsl *TreeShapeListener) ExitPk_tool_metaparam(ctx *parser.Pk_tool_metaparamContext) {}

func (tsl *TreeShapeListener) ExitPk_actual_param(ctx *parser.Pk_actual_paramContext) {
	fmt.Printf("%s => %s", ctx.GetParam_name().GetText(), ctx.GetParam_value().GetText())

	v := tsl.TreeProps[ctx.Pk_actual_param_value()].(*ActualParam)
	v.Name = ctx.GetParam_name().GetText()

	tsl.TreeProps[ctx] = v
}

func (tsl *TreeShapeListener) ExitPk_actual_param_value(ctx *parser.Pk_actual_param_valueContext) {
	var retval *ActualParam
	if ctx.GetVs() != nil {
		// String will alwaytsl come from Antlr with double quotetsl on each end
		v := StripFirstAndLast(ctx.GetVs().GetText())
		retval = NewUnnamedParamString(v)

	} else if ctx.GetVi() != nil {
		v, err := strconv.Atoi(ctx.GetVi().GetText())
		if err != nil {
			// TODO add these to a list of compile errors
			panic("Cannot parse number")
		}
		retval = NewUnnamedParamInt(v)
	} else if ctx.GetVb() != nil {
		v, err := strconv.ParseBool(ctx.GetVb().GetText())
		if err != nil {
			// TODO add these to a list of compile errors
			panic("Cannot parse bool")
		}
		retval = NewUnnamedParamBool(v)
	} else if ctx.GetVc() != nil {
		// NOT IMPLEMENTED, THItsl JUST RETURNtsl THE ENTIRE STRING
		retval = NewUnnamedParamString(ctx.GetVc().GetText())
	}
	tsl.TreeProps[ctx] = retval
}

// func (tsl *TreeShapeListener) ExitPk_bool(ctx *parser.Pk_boolContext) {}

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
