package rescode

import "github.com/gogf/gf/v2/errors/gcode"

var (
	// HTTP
	Success       = gcode.New(200, "SUCCESS", nil)
	ActionSuccess = gcode.New(201, "CREATE_SUCCESS", nil)
	BadRequest    = gcode.New(400, "BAD_REQUEST", nil)
	Unauthorized  = gcode.New(401, "UNAUTHORIZED", nil)
	Forbidden     = gcode.New(403, "FORBIDDEN", nil)
	NotFound      = gcode.New(404, "NOT_FOUND", nil)
	AlreadyExist  = gcode.New(409, "ALREADY_EXIST", nil)
	InternalError = gcode.New(500, "INTERNAL_ERROR", nil)

	// AUTH
	TokenExpired = gcode.New(40100, "TOKEN_EXPIRED", nil)
	TokenInvalid = gcode.New(40101, "TOKEN_INVALID", nil)
	TokenMissing = gcode.New(40102, "TOKEN_MISSING", nil)

	RefreshTokenExpired = gcode.New(40104, "REFRESH_TOKEN_EXPIRED", nil)
	InvalidRefreshToken = gcode.New(40103, "INVALID_REFRESH_TOKEN", nil)

	AccessTokenCreationFailed  = gcode.New(50000, "ACCESS_TOKEN_CREATION_FAILED", nil)
	RefreshTokenCreationFailed = gcode.New(50001, "Refresh_TOKEN_CREATION_FAILED", nil)

	PasswordResetTokenExpired = gcode.New(40105, "PASSWORD_RESET_TOKEN_EXPIRED", nil)
	LoginRequired             = gcode.New(40106, "LOGIN_REQUIRED", nil)

	AccountNotVerified = gcode.New(40300, "ACCOUNT_NOT_VERIFIED", nil)
	AccountLocked      = gcode.New(40301, "ACCOUNT_LOCKED", nil)
	AccountDisabled    = gcode.New(40302, "ACCOUNT_DISABLED", nil)
	PermissionDenied   = gcode.New(40303, "PERMISSION_DENIED", nil)

	EmailAlreadyVerified  = gcode.New(40900, "EMAIL_ALREADY_VERIFIED", nil)
	MobileAlreadyVerified = gcode.New(40901, "MOBILE_ALREADY_VERIFIED", nil)
	HashPasswordFailed    = gcode.New(50002, "HASH_PASSWORD_FAILED", nil)

	// USER
	UserNotFound          = gcode.New(40410, "USER_NOT_FOUND", nil)
	UsernameRequired      = gcode.New(40010, "USERNAME_REQUIRED", nil)
	UsernameInvalid       = gcode.New(40011, "USERNAME_INVALID", nil)
	UsernameAlreadyExists = gcode.New(40911, "USERNAME_ALREADY_EXISTS", nil)

	EmailRequired      = gcode.New(40012, "EMAIL_REQUIRED", nil)
	EmailNotFound      = gcode.New(40012, "EMAIL_NOT_FOUND", nil)
	EmailInvalid       = gcode.New(40013, "EMAIL_INVALID", nil)
	EmailAlreadyExists = gcode.New(40912, "EMAIL_ALREADY_EXISTS", nil)

	PasswordRequired = gcode.New(40014, "PASSWORD_REQUIRED", nil)
	PasswordTooShort = gcode.New(40015, "PASSWORD_TOO_SHORT", nil)
	PasswordMismatch = gcode.New(40016, "PASSWORD_MISMATCH", nil)
	PasswordWrong    = gcode.New(40017, "PASSWORD_Wrong", nil)

	UserCreateFailed = gcode.New(50010, "USER_CREATE_FAILED", nil)
	UserUpdateFailed = gcode.New(50011, "USER_UPDATE_FAILED", nil)
	UserDeleteFailed = gcode.New(50012, "USER_DELETE_FAILED", nil)
	UserGetFailed    = gcode.New(50013, "USER_GET_FAILED", nil)

	// PRODUCT
	ProductNotFound      = gcode.New(40420, "PRODUCT_NOT_FOUND", nil)
	ProductNameRequired  = gcode.New(40020, "PRODUCT_NAME_REQUIRED", nil)
	ProductPriceInvalid  = gcode.New(40021, "PRODUCT_PRICE_INVALID", nil)
	ProductAlreadyExists = gcode.New(40920, "PRODUCT_ALREADY_EXISTS", nil)

	// APP
	AppNotFound      = gcode.New(40430, "APP_NOT_FOUND", nil)
	AppNameRequired  = gcode.New(40030, "APP_NAME_REQUIRED", nil)
	AppNameInvalid   = gcode.New(40031, "APP_NAME_INVALID", nil)
	AppAlreadyExists = gcode.New(40930, "APP_ALREADY_EXISTS", nil)

	AppCreateFailed = gcode.New(50030, "APP_CREATE_FAILED", nil)
	AppUpdateFailed = gcode.New(50031, "APP_UPDATE_FAILED", nil)
	AppDeleteFailed = gcode.New(50032, "APP_DELETE_FAILED", nil)
	AppGetFailed    = gcode.New(50033, "APP_GET_FAILED", nil)

	AppCreateSuccess = gcode.New(20130, "APP_CREATE_SUCCESS", nil)
	AppUpdateSuccess = gcode.New(20031, "APP_UPDATE_SUCCESS", nil)
	AppDeleteSuccess = gcode.New(20032, "APP_DELETE_SUCCESS", nil)
	AppGetSuccess    = gcode.New(20033, "APP_GET_SUCCESS", nil)

	// ROLE
	RoleGetSuccess    = gcode.New(20043, "ROLE_GET_SUCCESS", nil)
	RoleCreateSuccess = gcode.New(20140, "ROLE_CREATE_SUCCESS", nil)
	RoleUpdateSuccess = gcode.New(20041, "ROLE_UPDATE_SUCCESS", nil)
	RoleDeleteSuccess = gcode.New(20042, "ROLE_DELETE_SUCCESS", nil)

	RolePermissionDenied = gcode.New(40340, "ROLE_PERMISSION_DENIED", nil)

	RoleNotFound      = gcode.New(40440, "ROLE_NOT_FOUND", nil)
	RoleAppIDRequired = gcode.New(40043, "ROLE_APP_ID_REQUIRED", nil)
	RoleAppIDInvalid  = gcode.New(40044, "ROLE_APP_ID_INVALID", nil)

	RoleNameAlreadyExists = gcode.New(40940, "ROLE_NAME_ALREADY_EXISTS", nil)

	RoleCreateFailed = gcode.New(50040, "ROLE_CREATE_FAILED", nil)
	RoleUpdateFailed = gcode.New(50041, "ROLE_UPDATE_FAILED", nil)
	RoleDeleteFailed = gcode.New(50042, "ROLE_DELETE_FAILED", nil)
	RoleGetFailed    = gcode.New(50043, "ROLE_GET_FAILED", nil)
)
