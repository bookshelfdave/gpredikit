package functions

import (
	"fmt"
	"net"
	"time"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func defPortOpen() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "port_open?",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			fmt.Printf("NET ACTUAL PARAMS:\n%+v\n", actualParams)
			addr, _ := actualParams.GetStringOrDefault("addr")

			// TODO: make this a parameter
			timeout := 10 * time.Millisecond
			conn, err := net.DialTimeout("tcp", addr, timeout)
			if err != nil {
				return rt.ResFail()
			}
			if conn != nil {
				defer conn.Close()
				return rt.ResPass()
			} else {
				return rt.ResFail()
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("addr").
			Build(),
		IsGroup: false,
	}
}
