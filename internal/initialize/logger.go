package initialize

import (
	"strongbody-api/global"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// Custom to required
// var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
//     jsonForLogger := JsonOutputsForLogger{
//         Time:    in.TimeFormat,
//         Level:   gstr.Trim(in.LevelFormat, "[]"),
//         Content: gstr.Trim(in.Content),
//     }
//     jsonBytes, err := json.Marshal(jsonForLogger)
//     if err != nil {
//         _, _ = os.Stderr.WriteString(err.Error())
//         return
//     }
//     in.Buffer.Write(jsonBytes)
//     in.Buffer.WriteString("\n")
//     in.Next(ctx)
// }

func InitLogger(env string) {
	glog.SetDefaultHandler(glog.HandlerJson)

	valid := map[string]string{
		"dev":        "dev",
		"production": "production",
		"test":       "test",
	}

	name, ok := valid[env]
	if !ok {
		name = "dev"
	}

	global.Logger = g.Log(name)
}
