package comp

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/bookshelfdave/gpredikit/parser"

	rt "github.com/bookshelfdave/gpredikit/runtime"
	//"golang.org/x/exp/slices"
	"github.com/davecgh/go-spew/spew"
)

type ParseTreeProperty = map[any]any

type TreeShapeListener struct {
	*parser.BasePredikitListener
	TreeProps ParseTreeProperty

	// used to generate ContentAddress values
	Filename string

	// the return values from walking the tree
	TopLevelChecks []*AstChkInstance
	ToolDefs       []*AstToolDef
	Errors         []error
}

func NewTreeShapeListener(filename string) *TreeShapeListener {
	return &TreeShapeListener{TreeProps: make(map[any]any), Filename: filename}
}

func (tsl *TreeShapeListener) AddError(err error) {
	tsl.Errors = append(tsl.Errors, err)
}

func (tsl *TreeShapeListener) MakeAddress(t antlr.Token) rt.ContentAddress {
	return rt.ContentAddress{
		Line:     t.GetLine(),
		Col:      t.GetColumn(),
		Filename: tsl.Filename,
	}
}

func (tsl *TreeShapeListener) ExitPk_toplevel(ctx *parser.Pk_toplevelContext) {
	// Children is already a method in Antlr
	for _, child := range ctx.GetKids() {
		v := tsl.TreeProps[child]
		switch v := v.(type) {
		case *AstChkInstance:
			tsl.TopLevelChecks = append(tsl.TopLevelChecks, v)
		case *AstToolDef:
			tsl.ToolDefs = append(tsl.ToolDefs, v)
		default:
			msg := fmt.Sprintf("Unexpected top level element %+v", v)
			panic(msg)
		}
	}

	fmt.Printf("%d checks\n", len(tsl.TopLevelChecks))
	fmt.Printf("%d tools\n", len(tsl.ToolDefs))
}

func (tsl *TreeShapeListener) ExitPk_toplevel_child(ctx *parser.Pk_toplevel_childContext) {
	if ctx.GetTest() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetTest()]
	} else if ctx.GetGroup() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetGroup()]
	} else if ctx.GetTool() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetTool()]
	}
}

func (tsl *TreeShapeListener) ExitPk_group_agg(ctx *parser.Pk_group_aggContext) {
	var t antlr.Token
	var s string
	if ctx.PK_ALL() != nil {
		s = "all"
		t = ctx.PK_ALL().GetSymbol()
	} else if ctx.PK_ANY() != nil {
		s = "any"
		t = ctx.PK_ANY().GetSymbol()
	} else if ctx.PK_NONE() != nil {
		s = "none"
		t = ctx.PK_NONE().GetSymbol()
	} else {
		panic("Invalid group function")
	}
	as := &AddressableString{Address: tsl.MakeAddress(t), V: s}
	tsl.TreeProps[ctx] = as
}

func (tsl *TreeShapeListener) ExitPk_group(ctx *parser.Pk_groupContext) {
	// it's harder to get the content address here because we only have non-terminals
	// so use an "AddressableString" to get the first position of "all" | "any" | "none"
	aggFn := tsl.TreeProps[ctx.GetAgg_fn()].(*AddressableString)

	actualParams := make(map[string]*rt.ActualParam, 0)
	children := make([]*AstChkInstance, 0)
	gcs := ctx.GetGroup_children()

	for _, gc := range gcs {
		v := tsl.TreeProps[gc]
		switch v := v.(type) {
		case *AstChkInstance:
			cd := v
			children = append(children, cd)
		case *rt.ActualParam:
			app := v
			if _, ok := actualParams[app.Name]; ok {
				// parameter defined multiple times!
				tsl.AddError(fmt.Errorf("Parameter %s appears multiple times in %s",
					app.Name, aggFn.Address))

			}
			actualParams[app.Name] = app
		default:
			fmt.Printf("Found unexpected group element! %+v\n", v)
			panic("Found unexpected group element")
		}
	}

	isNegated := false
	isRetrying := false
	isGroup := true
	cd := &AstChkInstance{
		FnName:       aggFn.V,
		ActualParams: actualParams,
		Children:     children,
		Address:      aggFn.Address,
		IsNegated:    isNegated,
		IsRetrying:   isRetrying,
		IsGroup:      isGroup,
	}
	tsl.TreeProps[ctx] = cd
}

func (tsl *TreeShapeListener) ExitPk_group_child(ctx *parser.Pk_group_childContext) {
	if ctx.GetTest() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetTest()]
	} else if ctx.GetGroup() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetGroup()]
	} else if ctx.GetActual_param() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.GetActual_param()]
	}
}

// func (tsl *TreeShapeListener) ExitPk_group_agg(ctx *parser.Pk_group_aggContext) {}

func (tsl *TreeShapeListener) ExitPk_actual_param(ctx *parser.Pk_actual_paramContext) {
	v := tsl.TreeProps[ctx.Pk_actual_param_value()].(*rt.ActualParam)
	v.Name = ctx.GetParam_name().GetText()
	tsl.TreeProps[ctx] = v
}

func (tsl *TreeShapeListener) ExitPk_actual_param_value(ctx *parser.Pk_actual_param_valueContext) {
	var retval *rt.ActualParam
	if ctx.GetVs() != nil {
		// String will alwaytsl come from Antlr with double quotetsl on each end
		v := StripFirstAndLast(ctx.GetVs().GetText())
		retval = rt.NewUnnamedParamString(v)

	} else if ctx.GetVi() != nil {
		v, err := strconv.Atoi(ctx.GetVi().GetText())
		if err != nil {
			// this would happen if you pass a massive number in as an int
			// TODO: check min/max parseable values
			msg := fmt.Sprintf("Cannot parse number at %s line %d col %d",
				tsl.Filename, ctx.GetVi().GetLine(), ctx.GetVi().GetColumn())
			tsl.AddError(errors.New(msg))
			// return a valid value, the upstream compiler code will check for an empty
			// list of errors to see if it can continue
			retval = rt.NewUnnamedParamInt(0)
		} else {
			retval = rt.NewUnnamedParamInt(v)
		}
	} else if ctx.GetVb() != nil {
		v, err := strconv.ParseBool(ctx.GetVb().GetText())
		if err != nil {
			// this shouldn't happen, antlr will return either true or false... If you get here, something
			// is _really_ broken
			panic("Cannot parse bool")
		}
		retval = rt.NewUnnamedParamBool(v)
	} else if ctx.GetVc() != nil {
		// NOT IMPLEMENTED, THIS JUST RETURNS THE ENTIRE STRING
		convStr := ctx.GetVc().GetText()
		fmt.Printf("Conversion string [%s]\n", convStr)

		retval = rt.NewUnnamedParamString(ctx.GetVc().GetText())
	}
	tsl.TreeProps[ctx] = retval
}

func (tsl *TreeShapeListener) ExitPk_test(ctx *parser.Pk_testContext) {
	aps := ctx.GetAps()
	address := tsl.MakeAddress(ctx.GetTestname())
	actualParams := make(map[string]*rt.ActualParam, len(aps))
	for _, ap := range aps {
		v := tsl.TreeProps[ap]
		app := v.(*rt.ActualParam)
		if _, ok := actualParams[app.Name]; ok {
			tsl.AddError(fmt.Errorf("Parameter %s appears multiple times in %s",
				app.Name, address))
		} else {
			actualParams[app.Name] = app
		}
	}

	fnname := ctx.GetTestname().GetText()
	isNegated := ctx.PK_NOT() != nil
	isRetrying := ctx.PK_RETRYING() != nil
	isGroup := false
	cd := &AstChkInstance{
		FnName:       fnname,
		ActualParams: actualParams,
		Children:     []*AstChkInstance{},
		Address:      address,
		IsNegated:    isNegated,
		IsRetrying:   isRetrying,
		IsGroup:      isGroup,
	}
	tsl.TreeProps[ctx] = cd
}

func (tsl *TreeShapeListener) ExitPk_tool(ctx *parser.Pk_toolContext) {
	td := &AstToolDef{
		ToolName:         ctx.GetTool_name().GetText(),
		DesignTimeParams: []*rt.ActualParam{},
		RuntimeParams:    map[string]map[string]*rt.ActualParam{},
		Address:          tsl.MakeAddress(ctx.GetTool_name()),
	}

	for _, toolChild := range ctx.GetKids() {
		v := tsl.TreeProps[toolChild]
		switch v := v.(type) {
		case *rt.ActualParam:
			td.DesignTimeParams = append(td.DesignTimeParams, v)
		case *AstToolInstanceParam:
			ip := v
			td.RuntimeParams[ip.ParamName] = ip.ParamsProps
		}
	}
	tsl.TreeProps[ctx] = td
}

func (tsl *TreeShapeListener) ExitPk_tool_child(ctx *parser.Pk_tool_childContext) {
	if ctx.Pk_actual_param() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.Pk_actual_param()]
	} else if ctx.Pk_tool_metaparam() != nil {
		tsl.TreeProps[ctx] = tsl.TreeProps[ctx.Pk_tool_metaparam()]
	} else {
		panic("Invalid state detected in tool_child")
	}
}

func (tsl *TreeShapeListener) ExitPk_tool_metaparam(ctx *parser.Pk_tool_metaparamContext) {
	paramName := ctx.GetTool_param_name().GetText()
	paramProps := make(map[string]*rt.ActualParam)
	for _, tap := range ctx.GetTool_actual_params() {
		ap := tsl.TreeProps[tap].(*rt.ActualParam)
		paramProps[ap.Name] = ap
	}
	ips := &AstToolInstanceParam{
		ParamName:   paramName,
		ParamsProps: paramProps,
		Address:     tsl.MakeAddress(ctx.GetTool_param_name()),
	}
	tsl.TreeProps[ctx] = ips
}

func (tsl *TreeShapeListener) ExitPk_tool_actual_param(ctx *parser.Pk_tool_actual_paramContext) {
	v := tsl.TreeProps[ctx.GetParam_value()].(*rt.ActualParam)
	v.Name = ctx.GetParam_name().GetText()
	v.Address = tsl.MakeAddress(ctx.GetParam_name())
	tsl.TreeProps[ctx] = v
}

func (tsl *TreeShapeListener) ExitPk_tool_actual_param_value(ctx *parser.Pk_tool_actual_param_valueContext) {
	// tool_actual_params have type names in addition to string|int|bool|etc
	// TODO: maybe reuse the code between actual_param and tool_actual_param?
	// might not be possible due to different parent types
	var retval *rt.ActualParam
	if ctx.GetVs() != nil {
		// String will alwaytsl come from Antlr with double quotetsl on each end
		v := StripFirstAndLast(ctx.GetVs().GetText())
		retval = rt.NewUnnamedParamString(v)
	} else if ctx.GetVi() != nil {
		v, err := strconv.Atoi(ctx.GetVi().GetText())
		if err != nil {
			// this would happen if you pass a massive number in as an int
			msg := fmt.Sprintf("Cannot parse number at %d:%d", ctx.GetVi().GetLine(), ctx.GetVi().GetColumn())
			tsl.AddError(errors.New(msg))
			// return a valid value, the upstream compiler code will check for an empty
			// list of errors to see if it can continue
			retval = rt.NewUnnamedParamInt(0)
		} else {
			retval = rt.NewUnnamedParamInt(v)
		}
	} else if ctx.GetVb() != nil {
		v, err := strconv.ParseBool(ctx.GetVb().GetText())
		if err != nil {
			// this shouldn't happen, antlr will return either true or false... If you get here, something
			// is _really_ broken
			panic("Cannot parse bool")
		}
		retval = rt.NewUnnamedParamBool(v)
	} else if ctx.GetVc() != nil {
		// NOT IMPLEMENTED, THIS JUST RETURNS THE ENTIRE STRING
		retval = rt.NewUnnamedParamString(ctx.GetVc().GetText())
	} else if ctx.GetVtn() != nil {
		tn, err := rt.TypeNameToType(ctx.GetVtn().GetText())
		if err != nil {
			msg := fmt.Errorf("Invalid type name at %s line %d col %d (%w)",
				tsl.Filename, ctx.GetVtn().GetLine(), ctx.GetVtn().GetColumn(), err)
			tsl.AddError(msg)
			// return a dummy value just to make the tree happy
			retval = rt.NewUnnamedParamInt(0)
		} else {
			retval = rt.NewUnnamedParamTypeName(tn)
		}
	}
	tsl.TreeProps[ctx] = retval
}

func DumpAstObject(o any) {
	spew.Config.Indent = "    "
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Dump(o)
}

func Walk(filename string) (*AstFile, error) {
	fmt.Printf("Compiling %s\n", filename)

	input, _ := antlr.NewFileStream(filename)
	lexer := parser.NewPredikitLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPredikitParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Pk_toplevel()
	tsl := NewTreeShapeListener(filename)

	antlr.ParseTreeWalkerDefault.Walk(tsl, tree)

	if len(tsl.Errors) > 0 {
		for _, e := range tsl.Errors {
			// TODO: figure out where / how to write errors
			fmt.Printf("Compile error: %s\n", e)
		}
		return nil, fmt.Errorf("Compilation failed")
	}
	return &AstFile{
		Filename: filename,
		Checks:   tsl.TopLevelChecks,
		Tools:    tsl.ToolDefs,
	}, nil
}
