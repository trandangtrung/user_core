package global

import (
	mail "strongbody-api/utility/gmail"
	"strongbody-api/utility/token"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	Logger *glog.Logger
	Token  token.Maker
	Gmail  mail.EmailSender
)
