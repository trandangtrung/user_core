package global

import (
	"demo/utility/token"
	"time"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	Logger      *glog.Logger
	Token       token.Maker
	VariableEnv struct {
		SecretKey   string
		TimeAccess  time.Duration
		TimeRefresh time.Duration
	}
)
