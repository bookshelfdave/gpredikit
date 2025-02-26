package runtime

import "fmt"

func RunCheckNoRetry(inst *ChkInstance, runEnv *RunEnv) {
	// TODO: RUN HOOK INIT

	result := inst.Run(runEnv)
	inst.Result = result
	// TODO: RUN PASS / FAIL / ERROR hooks

	// TODO: RUN HOOK TERM
}

func RunCheckMaybeRetry(inst *ChkInstance, runEnv *RunEnv) {
	runEnv.Formatter.CheckStart(inst)
	if !inst.IsRetrying {
		RunCheckNoRetry(inst, runEnv)
	} else {
		fmt.Println("Retries not implemented, sorry!")
		RunCheckNoRetry(inst, runEnv)
	}
	runEnv.Formatter.CheckEnd(inst)
	return
}
