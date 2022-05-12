package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorResponseBody = ErrorResponse{HttpSC: 400, Error: Err{Error: "ErrResponseBody", ErrorCode: "001"}}
	ErrorNotAuthUser  = ErrorResponse{HttpSC: 401, Error: Err{Error: "ErrorNotAuthUser", ErrorCode: "002"}}
	ErrorDbOps        = ErrorResponse{HttpSC: 500, Error: Err{Error: "ErrorDbOps", ErrorCode: "003"}}
	ErrorInternal     = ErrorResponse{HttpSC: 500, Error: Err{Error: "ErrorInternal", ErrorCode: "004"}}
)
