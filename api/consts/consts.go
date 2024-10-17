package consts

const (
	SUCCESS      = "200"
	SYSTEM_ERROR = "500"
	BAD_REQUEST  = "400"

	PARAMS_INVALID = "300"
	TOKEN_INVALID  = "301"
	LOGIN_FAILED   = "302"
	TOO_FREQUENTLY = "303"
	NO_DATA        = "304"
	DELETE_FAILED  = "305"
)

const (
	USER_ID = "userID"
)

var MsgFlags = map[string]string{
	SUCCESS:      "SUCCESS",
	SYSTEM_ERROR: "SYSTEM_ERROR",
	BAD_REQUEST:  "BAD_REQUEST",

	PARAMS_INVALID: "PARAMS_INVALID",
	TOKEN_INVALID:  "TOKEN_INVALID",
	LOGIN_FAILED:   "LOGIN_FAILED",
	TOO_FREQUENTLY: "TOO_FREQUENTLY",
	NO_DATA:        "NO_DATA",
	DELETE_FAILED:  "DELETE_FAILED",
}

type ApiErr struct {
	Code string
	Msg  string
}

func (e ApiErr) Error() string {
	return e.Msg
}
