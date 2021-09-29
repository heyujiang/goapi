package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrVaildation = &Errno{Code: 20003, Message: "Validation failed."}
	ErrEncrypt    = &Errno{Code: 20004, Message: "Password Encrypt failed."}
	ErrCreateUser = &Errno{Code: 20005, Message: "创建用户失败."}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
