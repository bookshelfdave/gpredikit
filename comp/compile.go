package comp

import (
	"fmt"
	"log/slog"
	"time"

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
	// 		- sort not found errors

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
	for apName, apValue := range inst.ActualParams.Params {
		if formalParam, ok := inst.Def.FormalParams[apName]; ok {
			if formalParam.ParamType != apValue.ParamType {
				err := fmt.Errorf("Expected type %s for parameter %s of check %s, got %s instead on line %d col %d",
					formalParam.ParamType, formalParam.Name, inst.Name, apValue.ParamType, inst.Address.Line, inst.Address.Col)
				errors = append(errors, err)
			}
		} else {
			// TODO: add more info here
			slog.Warn("Unhandled type param case")
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
					ap = rt.NewUnnamedParamTypeName(fp.Default.(string))
				default:
					errors = append(errors, fmt.Errorf("unexpected runtime.ParamTypeName: %#v", fp.ParamType))
				}
				ap.Name = fp.Name
				inst.ActualParams.Params[fp.Name] = ap

			} /* else { errors for missing required params are caught in another function } */
		}
	}
	return errors
}

func isValidFunctionName(chk *AstChkInstance, reg *rt.ChkRegistry) error {
	name := chk.FnName
	if !reg.IsValidFn(name) {
		err := fmt.Errorf("Unknown function '%s' specified at %s:%d:%d",
			chk.FnName, chk.Address.Filename, chk.Address.Line, chk.Address.Col)
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

	//newInst.IsQuery = chk.Is
	//newInst.InstanceID
	//newInst.Title

	// try to make all children first
	for _, child := range chk.Children {
		childInst, err := makeCheckInstance(child, reg, newInst)
		if err != nil {
			return nil, err
		}
		newInst.Children = append(newInst.Children, childInst)
	}

	return newInst, nil
}

func materializeFormalParams(inst *rt.ChkInstance) []error {
	var errors []error
	for _, mp := range META_PARAMS {
		if _, ok := inst.Def.FormalParams[mp.Name]; ok {
			err := fmt.Errorf("Meta param key %s already exists in %s", mp.Name, inst.Name)
			errors = append(errors, err)
		}
		inst.Def.FormalParams[mp.Name] = &mp
	}
	return errors
}

func CompileChecksToAsts(astFiles []*AstFile, reg *rt.ChkRegistry) []*rt.CompiledFileOut {
	allCfos := []*rt.CompiledFileOut{}

	for _, astFile := range astFiles {
		cfo := rt.NewCompiledFileOut(astFile.Filename)
		allCfos = append(allCfos, cfo)
		for _, checkDef := range astFile.Checks {
			inst, err := makeCheckInstance(checkDef, reg, nil)
			if err != nil {
				cfo.AddError(err)
				continue
			}
			if errors := materializeFormalParams(inst); errors != nil {
				cfo.AddErrors(errors)
			}
			defaultParamErrors := instantiateDefaultParams(inst)
			cfo.AddErrors(defaultParamErrors)
			typecheckErrors := typecheckParams(inst)
			cfo.AddErrors(typecheckErrors)

			cfo.AddChkInstance(inst)
		}
	}
	return allCfos
}
