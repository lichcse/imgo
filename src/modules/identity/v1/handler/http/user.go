package http

import (
	"errors"
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/service"
	"imgo/src/modules/identity/v1/validation"
	"imgo/src/response"
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
) *UserHandler {
	return &UserHandler{
		userService:    userService,
		response:       response,
		userValidation: validation.NewUserValidation(),
	}
}

// Add func godoc
// @Summary Add a new user
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param UserAddDTO body entity.UserAddDTO true "Add a new user body"
// @Success 200 {object} entity.UserResponseDTO "success"
// @Router /identity/v1/user [post]
func (m *UserHandler) Add(ctx *gin.Context) {
	userAddDTO := entity.UserAddDTO{}
	err := ctx.BindJSON(&userAddDTO)
	if err != nil {
		m.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	err = m.userValidation.Add(userAddDTO)
	if err != nil {
		m.response.Out(ctx, err, nil)
		return
	}

	err = m.userService.Add(userAddDTO)
	m.response.Out(ctx, err, nil)
	return
}

// Detail func godoc
// @Summary Detail info of user
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param user_id path int true "number"
// @Success 200 {object} entity.UserResponseDTO "success"
// @Router /identity/v1/user/{user_id} [get]
func (m *UserHandler) Detail(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		m.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	user, err := m.userService.Detail(uint64(userID))
	m.response.Out(ctx, err, user)
	return
}

// Update func
func (m *UserHandler) Update(ctx *gin.Context) {
	// TODO
}

// Delete func
func (m *UserHandler) Delete(ctx *gin.Context) {
	// TODO
}
