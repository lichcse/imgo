package response

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// IMResponseMock struct
type IMResponseMock struct {
	mock.Mock
}

// Out func mock set response
func (r *IMResponseMock) Out(ctx *gin.Context, err error, data interface{}) {
	r.Called(ctx, err, data)
}
