// Code generated from Predikit.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Predikit
import "github.com/antlr4-go/antlr/v4"

// PredikitListener is a complete listener for a parse tree produced by PredikitParser.
type PredikitListener interface {
	antlr.ParseTreeListener

	// EnterPk_toplevel is called when entering the pk_toplevel production.
	EnterPk_toplevel(c *Pk_toplevelContext)

	// EnterPk_toplevel_child is called when entering the pk_toplevel_child production.
	EnterPk_toplevel_child(c *Pk_toplevel_childContext)

	// EnterPk_group is called when entering the pk_group production.
	EnterPk_group(c *Pk_groupContext)

	// EnterPk_group_child is called when entering the pk_group_child production.
	EnterPk_group_child(c *Pk_group_childContext)

	// EnterPk_group_agg is called when entering the pk_group_agg production.
	EnterPk_group_agg(c *Pk_group_aggContext)

	// EnterPk_test is called when entering the pk_test production.
	EnterPk_test(c *Pk_testContext)

	// EnterPk_test_pred is called when entering the pk_test_pred production.
	EnterPk_test_pred(c *Pk_test_predContext)

	// EnterPk_tool is called when entering the pk_tool production.
	EnterPk_tool(c *Pk_toolContext)

	// EnterPk_tool_child is called when entering the pk_tool_child production.
	EnterPk_tool_child(c *Pk_tool_childContext)

	// EnterPk_tool_metaparam is called when entering the pk_tool_metaparam production.
	EnterPk_tool_metaparam(c *Pk_tool_metaparamContext)

	// EnterPk_actual_param is called when entering the pk_actual_param production.
	EnterPk_actual_param(c *Pk_actual_paramContext)

	// EnterPk_actual_param_value is called when entering the pk_actual_param_value production.
	EnterPk_actual_param_value(c *Pk_actual_param_valueContext)

	// EnterPk_tool_actual_param is called when entering the pk_tool_actual_param production.
	EnterPk_tool_actual_param(c *Pk_tool_actual_paramContext)

	// EnterPk_tool_actual_param_value is called when entering the pk_tool_actual_param_value production.
	EnterPk_tool_actual_param_value(c *Pk_tool_actual_param_valueContext)

	// EnterPk_conversion_fn is called when entering the pk_conversion_fn production.
	EnterPk_conversion_fn(c *Pk_conversion_fnContext)

	// EnterPk_bool is called when entering the pk_bool production.
	EnterPk_bool(c *Pk_boolContext)

	// ExitPk_toplevel is called when exiting the pk_toplevel production.
	ExitPk_toplevel(c *Pk_toplevelContext)

	// ExitPk_toplevel_child is called when exiting the pk_toplevel_child production.
	ExitPk_toplevel_child(c *Pk_toplevel_childContext)

	// ExitPk_group is called when exiting the pk_group production.
	ExitPk_group(c *Pk_groupContext)

	// ExitPk_group_child is called when exiting the pk_group_child production.
	ExitPk_group_child(c *Pk_group_childContext)

	// ExitPk_group_agg is called when exiting the pk_group_agg production.
	ExitPk_group_agg(c *Pk_group_aggContext)

	// ExitPk_test is called when exiting the pk_test production.
	ExitPk_test(c *Pk_testContext)

	// ExitPk_test_pred is called when exiting the pk_test_pred production.
	ExitPk_test_pred(c *Pk_test_predContext)

	// ExitPk_tool is called when exiting the pk_tool production.
	ExitPk_tool(c *Pk_toolContext)

	// ExitPk_tool_child is called when exiting the pk_tool_child production.
	ExitPk_tool_child(c *Pk_tool_childContext)

	// ExitPk_tool_metaparam is called when exiting the pk_tool_metaparam production.
	ExitPk_tool_metaparam(c *Pk_tool_metaparamContext)

	// ExitPk_actual_param is called when exiting the pk_actual_param production.
	ExitPk_actual_param(c *Pk_actual_paramContext)

	// ExitPk_actual_param_value is called when exiting the pk_actual_param_value production.
	ExitPk_actual_param_value(c *Pk_actual_param_valueContext)

	// ExitPk_tool_actual_param is called when exiting the pk_tool_actual_param production.
	ExitPk_tool_actual_param(c *Pk_tool_actual_paramContext)

	// ExitPk_tool_actual_param_value is called when exiting the pk_tool_actual_param_value production.
	ExitPk_tool_actual_param_value(c *Pk_tool_actual_param_valueContext)

	// ExitPk_conversion_fn is called when exiting the pk_conversion_fn production.
	ExitPk_conversion_fn(c *Pk_conversion_fnContext)

	// ExitPk_bool is called when exiting the pk_bool production.
	ExitPk_bool(c *Pk_boolContext)
}
