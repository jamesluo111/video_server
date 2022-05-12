package main

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized   = Err{Error: "api not recognized", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "ErrorRequestBodyParseFailed", ErrorCode: "002"}
	IternalError                = Err{Error: "IternalError", ErrorCode: "003"}
)
