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

	status := http.StatusOK
	if err != nil {
		result.Message = r.language.GetMessage(err.Error())
		if codeStatus, ok := r.errorMessageMapping[err.Error()]; ok {
			result.Code = codeStatus.Code
			status = codeStatus.Status
		}
		result.Status = RestResponseStatusFail
	} else {
		result.Data = data
	}
	ctx.JSON(status, result)
}

func (r *rest) setLang(ctx *gin.Context) {
	r.language.SetLanguage(ctx.DefaultQuery("lang", "en"))
}
