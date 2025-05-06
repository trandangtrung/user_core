package global

import (
	mail "github.com/quannv/strongbody-api/utility/gmail"
	"github.com/quannv/strongbody-api/utility/template"
	"github.com/quannv/strongbody-api/utility/token"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	Logger   *glog.Logger
	Token    token.Maker
	Gmail    mail.EmailSender
	Template template.ITemplate
)
