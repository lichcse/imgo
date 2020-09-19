package response

import "github.com/gin-gonic/gin"

// Response struct
type Response struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RestResponse struct
type RestResponse struct {
	Response
	Data interface{} `json:"data"`
}

// CodeStatus struct
type CodeStatus struct {
	Code   string
	Status int
}

// CodeMessageMapping struct
type CodeMessageMapping map[string]CodeStatus

// IMResponse interface
type IMResponse interface {
	Out(ctx *gin.Context, err error, data interface{})
}
