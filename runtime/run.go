package runtime

import "fmt"

func RunCheckNoRetry(inst *ChkInstance, runEnv *RunEnv) *ChkResult {
	// TODO: RUN HOOK INIT

	result := inst.Def.CheckFunction(inst.ActualParams, runEnv, inst)

	// TODO: RUN PASS / FAIL / ERROR hooks

	// TODO: RUN HOOK TERM
	return result
}

func RunCheckMaybeRetry(inst *ChkInstance, runEnv *RunEnv) *ChkResult {
	if !inst.IsRetrying {
		RunCheckNoRetry(inst, runEnv)
	} else {
		fmt.Println("Retries not implemented, sorry!")
		RunCheckNoRetry(inst, runEnv)
	}
	return ResOk(true)
}
