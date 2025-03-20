package runtime

import (
	"fmt"
	"time"
)

func runCheckNoRetry(inst *ChkInstance, runEnv *RunEnv) {
	inst.RunHookIfDefined("on_init", runEnv)

	result := inst.Run(runEnv)
	inst.Result = result
	// TODO: RUN PASS / FAIL / ERROR hooks
	if inst.Result.Err != nil {
		inst.RunHookIfDefined("on_error", runEnv)
	} else {
		if inst.Result.PassFail {
			inst.RunHookIfDefined("on_pass", runEnv)
		} else {
			inst.RunHookIfDefined("on_fail", runEnv)
		}
	}
	inst.RunHookIfDefined("on_term", runEnv)
}

func runWithRetry(inst *ChkInstance, runEnv *RunEnv) {
	inst.RunHookIfDefined("on_init", runEnv)

	pass := false
	retries, _ := inst.ActualParams.GetIntOrDefault("retries")
	retryDelay, _ := inst.ActualParams.GetDurationOrDefault("retry_delay")

	for attemptNum := 0; attemptNum < retries && !pass; attemptNum += 1 {
		result := inst.Run(runEnv)
		inst.Result = result
		if inst.Result.PassFail {
			pass = true
		}

		// TODO: RUN PASS / FAIL / ERROR hooks
		if inst.Result.Err != nil {
			inst.RunHookIfDefined("on_error", runEnv)
		} else {
			if inst.Result.PassFail {
				inst.RunHookIfDefined("on_pass", runEnv)
			} else {
				inst.RunHookIfDefined("on_fail", runEnv)
			}
		}
		retryInDelayInSeconds := retryDelay / time.Second
		fmt.Printf("\tSleep for %d seconds before attempt %d\n", retryInDelayInSeconds, attemptNum)
		time.Sleep(retryDelay)
	}

	inst.RunHookIfDefined("on_term", runEnv)

}

func RunCheck(inst *ChkInstance, runEnv *RunEnv) {
	runEnv.Formatter.CheckStart(inst)
	if !inst.IsRetrying {
		runCheckNoRetry(inst, runEnv)
	} else {
		fmt.Println("Retries not implemented, sorry!")
		runCheckNoRetry(inst, runEnv)
	}
	runEnv.Formatter.CheckEnd(inst)
	return
}
