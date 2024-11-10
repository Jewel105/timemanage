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
	CODE_INVALID   = "306"
)

const (
	USER_ID      = "userID"
	TOKEN_ID     = "tokenID"
	EQUIPMENT_ID = "equipmentID"
	LANG         = "lang"
)

type ApiErr struct {
	Code string
	Msg  string
}

func (e ApiErr) Error() string {
	return e.Msg
}
