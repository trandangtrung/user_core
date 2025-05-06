package consts

const (
	AuthorizationHeader = "Authorization"
	AuthorizationScope  = "Scope"
	AuthorizationType   = "Bearer"
	AuthorizationKey    = "authorization_payload"
)

var CONFIG_PERMISSIONS = map[string]string{
	"admin":   "admin",
	"network": "network",
	"shop":    "shop",
	"multi":   "multi",
}

var CONFIG_SCOPE = map[string]string{
	"network": "network",
	"shop":    "shop",
	"multi":   "multi",
}


