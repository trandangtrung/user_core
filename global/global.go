package global

import (
	mail "demo/utility/gmail"
	"demo/utility/token"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	Logger *glog.Logger
	Token  token.Maker
	Gmail  mail.EmailSender
)
