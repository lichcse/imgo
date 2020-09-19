package http

import (
	"errors"
	"imgo/app/common/identity/v1/dto"
	"imgo/app/modules/identity/v1/service"
	"imgo/app/modules/identity/v1/validation"
	"imgo/app/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler struct
type UserHandler struct {
	userService    service.UserService
	response       response.IMResponse
	userValidation validation.UserValidation
}

// NewUserHandler func
func NewUserHandler(
	userService service.UserService,
	response response.IMResponse,
	userValidation validation.UserValidation,
) *UserHandler {
	return &UserHandler{
		userService:    userService,
		response:       response,
		userValidation: userValidation,
	}
}

// Add func godoc
// @Summary Add a new user
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param UserAddRequest body dto.UserAddRequest true "Add a new user body"
// @Success 200 {object} dto.UserDetailResponse "success"
// @Router /identity/v1/user [post]
func (u *UserHandler) Add(ctx *gin.Context) {
	userAddRequest := &dto.UserAddRequest{}
	err := ctx.BindJSON(userAddRequest)
	if err != nil {
		u.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	err = u.userValidation.Add(userAddRequest)
	if err != nil {
		u.response.Out(ctx, err, nil)
		return
	}

	userDetailResponse, err := u.userService.Add(userAddRequest)
	if err != nil {
		u.response.Out(ctx, err, nil)
		return
	}
	u.response.Out(ctx, err, userDetailResponse)
}

// Detail func godoc
// @Summary Detail info of user
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param id path int true "number"
// @Success 200 {object} dto.UserDetailResponse "success"
// @Router /identity/v1/user/{id} [get]
func (u *UserHandler) Detail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		u.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	userDetailResponse, err := u.userService.Detail(uint64(id))
	u.response.Out(ctx, err, userDetailResponse)
	return
}

// Update func
func (u *UserHandler) Update(ctx *gin.Context) {
	// TODO
}

// Delete func
func (u *UserHandler) Delete(ctx *gin.Context) {
	// TODO
}
