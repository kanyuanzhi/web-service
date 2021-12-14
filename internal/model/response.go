package model

type Response struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code uint, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    20000,
		Message: "operate successfully",
		Data:    data,
	}
}

type FilesResponse struct {
	Avatar string `json:"avatar"`
}
