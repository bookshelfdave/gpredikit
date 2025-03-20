package comp

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/bookshelfdave/gpredikit/functions"
	rt "github.com/bookshelfdave/gpredikit/runtime"
)

type CompileErrors struct {
	Errors []error
}

func NewCompilerErrors() *CompileErrors {
	return &CompileErrors{
		Errors: make([]error, 0),
	}
}

func (ce *CompileErrors) AddError(e error) {
	ce.Errors = append(ce.Errors, e)
}

func typecheckParams(inst *rt.ChkInstance) []error {
	var errors []error
	// 1) check for required params
	// TODO? sort not found errors

	for _, fp := range inst.Def.FormalParams {
		if fp.Required {
			if _, ok := inst.ActualParams.Get(fp.Name); !ok {
				err := fmt.Errorf("Missing required parameter %s for check %s on line %d col %d", fp.Name, inst.Name, inst.Address.Line, inst.Address.Col)
				errors = append(errors, err)
			}
		}
	}

	// 2) look for unexpected / unknown params
	for k := range inst.ActualParams.Params {
		if _, ok := inst.Def.FormalParams[k]; !ok {
			err := fmt.Errorf("Unknown parameter %s for check %s on line %d col %d", k, inst.Name, inst.Address.Line, inst.Address.Col)
			errors = append(errors, err)
		}
	}

	// 3) check for param types
	for actualParamValue, apValue := range inst.ActualParams.Params {
		if formalParam, ok := inst.Def.FormalParams[actualParamValue]; ok {
			if formalParam.CoercibleFromString() && apValue.CoercibleToString() {
				fmt.Println("Duration to string")
				// pass
			} else if formalParam.ParamType != apValue.ParamType {
				err := fmt.Errorf("Expected type %s for parameter %s of check %s, got %s instead on line %d col %d",
					formalParam.ParamType, formalParam.Name, inst.Name, apValue.ParamType, inst.Address.Line, inst.Address.Col)
				errors = append(errors, err)
			}
		} else {
			// TODO: add more info here
			slog.Warn("Unhandled type param case")
		}
	}

	for _, child := range inst.Children {
		childErrors := typecheckParams(child)
		if len(childErrors) > 0 {
			errors = append(errors, childErrors...)
		}
	}

	return errors
}

func instantiateDefaultParams(inst *rt.ChkInstance) []error {
	var errors []error

	for _, fp := range inst.Def.FormalParams {
		if _, ok := inst.ActualParams.Get(fp.Name); !ok {
			if !fp.Required {
				if fp.Default == nil {
					continue
				}
				// only instantiate default values for optional params

				var ap *rt.ActualParam
				switch fp.ParamType {
				case rt.ParamTypeBool:
					ap = rt.NewUnnamedParamString(fp.Default.(string))
				case rt.ParamTypeDuration:
					ap = rt.NewUnnamedParamDuration(fp.Default.(time.Duration))
				case rt.ParamTypeInt:
					ap = rt.NewUnnamedParamInt(fp.Default.(int))
				case rt.ParamTypeString:
					ap = rt.NewUnnamedParamString(fp.Default.(string))
				case rt.ParamTypeTypeName:
					ap = rt.NewUnnamedParamTypeName(fp.Default.(rt.ParamTypeName))
				default:
					errors = append(errors, fmt.Errorf("unexpected runtime.ParamTypeName: %#v", fp.ParamType))
				}
				ap.Name = fp.Name
				inst.ActualParams.Params[fp.Name] = ap

			} /* else { errors for missing required params are caught in another function } */
		}
	}

	// recursively instantiate default params for children
	for _, child := range inst.Children {
		childErrors := instantiateDefaultParams(child)
		if len(childErrors) > 0 {
			errors = append(errors, childErrors...)
		}
	}

	return errors
}

func isValidFunctionName(chk *AstChkInstance, reg *rt.ChkRegistry) error {
	name := chk.FnName
	if !reg.IsValidFn(name) {
		err := fmt.Errorf("Unknown function '%s' specified at %s",
			chk.FnName, chk.Address)
		return err
	}

	return nil
}

func makeCheckInstance(chk *AstChkInstance, reg *rt.ChkRegistry, parent *rt.ChkInstance) (*rt.ChkInstance, error) {

	if err := isValidFunctionName(chk, reg); err != nil {
		return nil, err
	}

	newInst := &rt.ChkInstance{}
	newInst.Name = chk.FnName
	newInst.Def = reg.GetFn(chk.FnName)
	newInst.Parent = parent
	newInst.ActualParams = rt.NewActualParams(newInst.Def, chk.ActualParams)
	newInst.Parent = parent
	newInst.IsRetrying = chk.IsRetrying
	newInst.IsNegated = chk.IsNegated
	newInst.Address = chk.Address

	for _, child := range chk.Children {
		childInst, err := makeCheckInstance(child, reg, newInst)
		if err != nil {
			return nil, err
		}
		newInst.Children = append(newInst.Children, childInst)
	}

	return newInst, nil
}

func materializeFormalParams(cd *rt.ChkDef) []error {
	var errors []error
	addMatParam := func(mp rt.FormalParam) {
		if _, ok := cd.FormalParams[mp.Name]; ok {
			err := fmt.Errorf("Meta param key %s already exists in %s", mp.Name, cd.Name)
			errors = append(errors, err)
		}
		cd.FormalParams[mp.Name] = &mp
	}
	for _, mp := range META_PARAMS {
		addMatParam(mp)
	}

	for _, mp := range RETRY_PARAMS {
		addMatParam(mp)
	}
	return errors
}

func _collectSetOfDefs(node *rt.ChkInstance, defs map[string]*rt.ChkDef) {
	defs[node.Def.Name] = node.Def
	for _, child := range node.Children {
		_collectSetOfDefs(child, defs)
	}
}

func collectSetOfDefs(root *rt.ChkInstance) map[string]*rt.ChkDef {
	defs := map[string]*rt.ChkDef{}
	_collectSetOfDefs(root, defs)
	return defs
}

func CompileTools(tools []*AstToolDef) []error {
	allErrors := []error{}

	for _, toolDef := range tools {
		fmt.Printf("COMPILING TOOL %+v\n", toolDef)
		for paramName, paramParams := range toolDef.RuntimeParams {
			// runtime params get turned into FormalParams
			fmt.Printf("---%s\n", paramName)
			fmt.Printf("%+v\n", paramParams)
			required := false

			// make sure "type" is specified
			if _, ok := paramParams["type"]; !ok {
				err := fmt.Errorf("Parameter 'type' for %v required at %s", paramName, toolDef.Address.Render())
				allErrors = append(allErrors, err)
				continue
			}

			// make sure "type" is a TypeName
			if paramParams["type"].ParamType != rt.ParamTypeTypeName {
				err := fmt.Errorf("Invalid type for parameter 'type' at %s",
					paramParams["type"].Address.Render())
				allErrors = append(allErrors, err)
				continue
			}

			theType, _ := paramParams["type"].GetTypeName()

			// the "required" parameter value is optional
			if _, ok := paramParams["required"]; ok {
				// but make sure it's a Bool
				if paramParams["required"].ParamType != rt.ParamTypeBool {
					err := fmt.Errorf("Invalid type for parameter 'required' at %s",
						paramParams["required"].Address.Render())
					allErrors = append(allErrors, err)
					continue
				} else {
					required, _ = paramParams["required"].GetBool()
				}
			}
			fmt.Printf("Required = %t\n", required)
			fmt.Printf("The type = %+v\n", theType)
		}

		// td := rt.ToolDef{
		// 	Name:             toolDef.ToolName,
		// 	RuntimeParams:    map[string]*rt.FormalParam{},
		// 	DesignTimeParams: map[string]*rt.ActualParam{},
		// }
	}

	for _, err := range allErrors {
		fmt.Println(err)
	}
	return allErrors
}

func CompileChecksToAsts(astFiles []*AstFile, reg *rt.ChkRegistry) []*rt.CompiledFileOut {
	allCfos := make([]*rt.CompiledFileOut, len(astFiles))

	for n, astFile := range astFiles {
		cfo := rt.NewCompiledFileOut(astFile.Filename)
		allCfos[n] = cfo
		for _, checkDef := range astFile.Checks {
			inst, err := makeCheckInstance(checkDef, reg, nil)
			if err != nil {
				cfo.AddError(err)
				continue
			}
			setOfDefs := collectSetOfDefs(inst)
			for _, def := range setOfDefs {
				if errors := materializeFormalParams(def); errors != nil {
					cfo.AddErrors(errors)
				}
			}
			// TODO: THIS NEEDS TO BE RECURSIVE
			defaultParamErrors := instantiateDefaultParams(inst)

			cfo.AddErrors(defaultParamErrors)
			typecheckErrors := typecheckParams(inst)

			cfo.AddErrors(typecheckErrors)

			cfo.AddChkInstance(inst)
		}

		CompileTools(astFile.Tools)
	}
	return allCfos
}

// TODO move this somewhere else
func RunChecks(compiledAstFiles []*rt.CompiledFileOut, formatter rt.OutputFormatter) {
	runEnv := &rt.RunEnv{Formatter: formatter}
	runEnv.Formatter.Init()
	for _, cfo := range compiledAstFiles {
		for _, inst := range cfo.Instances {
			rt.RunCheck(inst, runEnv)
		}
	}
	runEnv.Formatter.Term()
}

func CompileInputFiles(inFiles []string) []*rt.CompiledFileOut {
	allAstFiles := make([]*AstFile, len(inFiles))
	for n, inFile := range inFiles {
		// build an ast for each file, collect them in allAstFiles
		astFile, err := Walk(inFile)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		// keep gathering compilation errors without exiting
		allAstFiles[n] = astFile
	}

	reg := rt.NewChkRegistry()
	functions.RegisterBuiltins(reg)
	compiledAstFiles := CompileChecksToAsts(allAstFiles, reg)

	errorsDetected := false
	for _, cfo := range compiledAstFiles {
		for _, err := range cfo.Errors {
			fmt.Printf("Error: %s\n", err)
		}
		if len(cfo.Errors) != 0 {
			// compilation can't finish, but show all the compilation errors before exiting
			errorsDetected = true
		}
	}

	if errorsDetected {
		return nil
	}

	return compiledAstFiles

}
