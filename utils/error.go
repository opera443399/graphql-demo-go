package utils

import "fmt"

//ErrorCode error code
type ErrorCode int

const (
	//ErrCodeBase error base number
	ErrCodeBase = 10000
)

//ErrorCode sys=10000+
const (
	ErrInternalException ErrorCode = 1*ErrCodeBase + iota // value --> 0
)

//ErrorCode service=20000+
const (
	ErrTokenInvalid ErrorCode = 2*ErrCodeBase + iota // value --> 0
	ErrTokenNotFound
)

var errCodeDesc = map[ErrorCode]string{
	//10000+
	ErrInternalException:       "系统内部错误",

	//20000+
	ErrTokenInvalid:            "Token无效或者已过期",
	ErrTokenNotFound: 			"非法请求",
}

//String ErrorCode
func (f ErrorCode) String() string {
	if value, ok := errCodeDesc[f]; ok {
		return value
	}
	return "UNKNOWN"
}

//Error ErrorCode
func (f ErrorCode) Error() string {
	//file, line := WhereAmI(6)
	return fmt.Sprintf("%d:%s", int(f), f.String())
}
