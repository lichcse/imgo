package response

import (
	"imgo/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// RestResponseStatusSuccess var
	RestResponseStatusSuccess int = 1
	// RestResponseStatusFail var
	RestResponseStatusFail int = 0
)

type rest struct {
	language            utils.IMLanguage
	errorMessageMapping CodeMessageMapping
}

// NewRestResponse func
func NewRestResponse(language utils.IMLanguage, errorMessageMapping CodeMessageMapping) IMResponse {
	return &rest{language: language, errorMessageMapping: errorMessageMapping}
}

func (r *rest) Out(ctx *gin.Context, err error, data interface{}) {
	r.setLang(ctx)
	result := RestResponse{}
	result.Code = ""
	result.Status = RestResponseStatusSuccess
	code := ""
	status := http.StatusOK
	if err == nil {
		code = "success"
		result.Data = data
	} else {
		code = err.Error()
		result.Status = RestResponseStatusFail
	}
	result.Message = r.language.GetMessage(code)
	if codeStatus, ok := r.errorMessageMapping[code]; ok {
		result.Code = codeStatus.Code
		status = codeStatus.Status
	}
	ctx.JSON(status, result)
}

func (r *rest) setLang(ctx *gin.Context) {
	r.language.SetLanguage(ctx.DefaultQuery("lang", "en"))
}
