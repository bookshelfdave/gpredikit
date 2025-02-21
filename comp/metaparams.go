package comp

import (
	"time"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

var META_PARAMS = map[string]rt.FormalParam{
	"title": {
		Name:      "title",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},

	"on_pass": {
		Name:      "on_pass",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},

	"on_fail": {
		Name:      "on_fail",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},

	"on_error": {
		Name:      "on_error",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},

	"on_init": {
		Name:      "on_init",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},

	"on_term": {
		Name:      "on_term",
		Required:  false,
		ParamType: rt.ParamTypeString,
		Default:   nil,
	},
}

var RETRY_PARAMS = map[string]rt.FormalParam{
	"retry_count": {
		Name:      "retry_count",
		Required:  false,
		ParamType: rt.ParamTypeInt,
		Default:   3,
	},
	"retry_delay": {
		Name:      "retry_delay",
		Required:  false,
		ParamType: rt.ParamTypeDuration,
		Default:   5 * time.Second,
	},
}
