package app

import e "github.com/coy102/go-starter/utils/e"

// MsgFlags for mapping constant code
var MsgFlags = map[int]string{
	e.Success:      "Ok",
	e.Error:        "Fail",
	e.InvalidParam: "Invalid Parameters",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[e.Error]
}