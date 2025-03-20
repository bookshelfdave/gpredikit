package runtime

import (
	"fmt"

	"github.com/hoisie/mustache"
)

type ToolDef struct {
	Name string
	// Tool prefix to help a little bit with keeping param types clear
	RuntimeParams    map[string]*FormalParam
	DesignTimeParams map[string]*ActualParam
}

func (td *ToolDef) MakeChkDefFromTool() *ChkDef {
	cmdTemplate := td.DesignTimeParams["cmd_template"]
	fmt.Printf("cmd_template = [%s]\n", cmdTemplate)

	var formalParams map[string]*FormalParam
	var templateParams map[string]*ActualParam

	return &ChkDef{
		Name: td.Name,
		CheckFunction: func(actualParams *ActualParams, runEnv *RunEnv, chkInst *ChkInstance) *ChkResult {
			// get the cmd_template
			// render actual params in the template
			// exec the command
			// return the ChkResult
			//
			data := mustache.Render("hello {{c}}", map[string]string{"c": "world"})
			fmt.Println(data)

			return nil
		},
		FormalParams:    formalParams,
		TemplateParams:  templateParams,
		AcceptsChildren: false,
		IsGroup:         false,
		IsQuery:         false,
	}
}
