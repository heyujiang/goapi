package errno

var (
	// Common errors
	OK      = &Errno{Code: 200, Message: "OK"}
	SUCCESS = &Errno{Code: 200, Message: "Success."}

	NoUsername = &Errno{Code: 100, Message: "Username is no found."}

	ErrTokenInvalid       = &Errno{Code: 20004, Message: "The token was invalid."}
	ErrMissingTokenString = &Errno{Code: 20005, Message: "The length of the `Authorization` header is zero."}

	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrVaildation       = &Errno{Code: 20003, Message: "Validation failed."}
	ErrEncrypt          = &Errno{Code: 20004, Message: "Password Encrypt failed."}
	ErrCreateUser       = &Errno{Code: 20005, Message: "创建用户失败."}
	ErrDeleteUser       = &Errno{Code: 20006, Message: "删除用户失败"}

	ErrData = &Errno{Code: 20009, Message: "无数据"}
)
