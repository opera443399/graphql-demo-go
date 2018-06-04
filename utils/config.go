package utils

//PortDefault default
var PortDefault = 2000
//ServicesCode all the service
type ServicesCode int

///ServicesCode all the service
const (
	SvcIgnore ServicesCode = iota //只应该在单元测试中使用
	SvcGateway
	SvcToken
)

var serviceDesc = map[ServicesCode]string{
	SvcIgnore:  "ignore",
	SvcGateway: "gateway",
	SvcToken: "token",
}

//String code to String
func (f ServicesCode) String() string {
	if value, ok := serviceDesc[f]; ok {
		return value
	}
	return "UNKNOWN"
}
