// Code generated from Predikit.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Predikit
import "github.com/antlr4-go/antlr/v4"

// BasePredikitListener is a complete listener for a parse tree produced by PredikitParser.
type BasePredikitListener struct{}

var _ PredikitListener = &BasePredikitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePredikitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePredikitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePredikitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePredikitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPk_toplevel is called when production pk_toplevel is entered.
func (s *BasePredikitListener) EnterPk_toplevel(ctx *Pk_toplevelContext) {}

// ExitPk_toplevel is called when production pk_toplevel is exited.
func (s *BasePredikitListener) ExitPk_toplevel(ctx *Pk_toplevelContext) {}

// EnterPk_toplevel_child is called when production pk_toplevel_child is entered.
func (s *BasePredikitListener) EnterPk_toplevel_child(ctx *Pk_toplevel_childContext) {}

// ExitPk_toplevel_child is called when production pk_toplevel_child is exited.
func (s *BasePredikitListener) ExitPk_toplevel_child(ctx *Pk_toplevel_childContext) {}

// EnterPk_group is called when production pk_group is entered.
func (s *BasePredikitListener) EnterPk_group(ctx *Pk_groupContext) {}

// ExitPk_group is called when production pk_group is exited.
func (s *BasePredikitListener) ExitPk_group(ctx *Pk_groupContext) {}

// EnterPk_group_child is called when production pk_group_child is entered.
func (s *BasePredikitListener) EnterPk_group_child(ctx *Pk_group_childContext) {}

// ExitPk_group_child is called when production pk_group_child is exited.
func (s *BasePredikitListener) ExitPk_group_child(ctx *Pk_group_childContext) {}

// EnterPk_group_agg is called when production pk_group_agg is entered.
func (s *BasePredikitListener) EnterPk_group_agg(ctx *Pk_group_aggContext) {}

// ExitPk_group_agg is called when production pk_group_agg is exited.
func (s *BasePredikitListener) ExitPk_group_agg(ctx *Pk_group_aggContext) {}

// EnterPk_test is called when production pk_test is entered.
func (s *BasePredikitListener) EnterPk_test(ctx *Pk_testContext) {}

// ExitPk_test is called when production pk_test is exited.
func (s *BasePredikitListener) ExitPk_test(ctx *Pk_testContext) {}

// EnterPk_tool is called when production pk_tool is entered.
func (s *BasePredikitListener) EnterPk_tool(ctx *Pk_toolContext) {}

// ExitPk_tool is called when production pk_tool is exited.
func (s *BasePredikitListener) ExitPk_tool(ctx *Pk_toolContext) {}

// EnterPk_tool_child is called when production pk_tool_child is entered.
func (s *BasePredikitListener) EnterPk_tool_child(ctx *Pk_tool_childContext) {}

// ExitPk_tool_child is called when production pk_tool_child is exited.
func (s *BasePredikitListener) ExitPk_tool_child(ctx *Pk_tool_childContext) {}

// EnterPk_tool_metaparam is called when production pk_tool_metaparam is entered.
func (s *BasePredikitListener) EnterPk_tool_metaparam(ctx *Pk_tool_metaparamContext) {}

// ExitPk_tool_metaparam is called when production pk_tool_metaparam is exited.
func (s *BasePredikitListener) ExitPk_tool_metaparam(ctx *Pk_tool_metaparamContext) {}

// EnterPk_actual_param is called when production pk_actual_param is entered.
func (s *BasePredikitListener) EnterPk_actual_param(ctx *Pk_actual_paramContext) {}

// ExitPk_actual_param is called when production pk_actual_param is exited.
func (s *BasePredikitListener) ExitPk_actual_param(ctx *Pk_actual_paramContext) {}

// EnterPk_actual_param_value is called when production pk_actual_param_value is entered.
func (s *BasePredikitListener) EnterPk_actual_param_value(ctx *Pk_actual_param_valueContext) {}

// ExitPk_actual_param_value is called when production pk_actual_param_value is exited.
func (s *BasePredikitListener) ExitPk_actual_param_value(ctx *Pk_actual_param_valueContext) {}

// EnterPk_tool_actual_param is called when production pk_tool_actual_param is entered.
func (s *BasePredikitListener) EnterPk_tool_actual_param(ctx *Pk_tool_actual_paramContext) {}

// ExitPk_tool_actual_param is called when production pk_tool_actual_param is exited.
func (s *BasePredikitListener) ExitPk_tool_actual_param(ctx *Pk_tool_actual_paramContext) {}

// EnterPk_tool_actual_param_value is called when production pk_tool_actual_param_value is entered.
func (s *BasePredikitListener) EnterPk_tool_actual_param_value(ctx *Pk_tool_actual_param_valueContext) {
}

// ExitPk_tool_actual_param_value is called when production pk_tool_actual_param_value is exited.
func (s *BasePredikitListener) ExitPk_tool_actual_param_value(ctx *Pk_tool_actual_param_valueContext) {
}

// EnterPk_conversion_fn is called when production pk_conversion_fn is entered.
func (s *BasePredikitListener) EnterPk_conversion_fn(ctx *Pk_conversion_fnContext) {}

// ExitPk_conversion_fn is called when production pk_conversion_fn is exited.
func (s *BasePredikitListener) ExitPk_conversion_fn(ctx *Pk_conversion_fnContext) {}

// EnterPk_bool is called when production pk_bool is entered.
func (s *BasePredikitListener) EnterPk_bool(ctx *Pk_boolContext) {}

// ExitPk_bool is called when production pk_bool is exited.
func (s *BasePredikitListener) ExitPk_bool(ctx *Pk_boolContext) {}
