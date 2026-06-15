package http

import (
	"errors"
	"imgo/app/modules/identity/v1/service"
	"imgo/app/modules/identity/v1/validation"
	"imgo/app/response"
	schema "imgo/app/schema/identity/v1"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostHandler struct
type PostHandler struct {
	postService    service.PostService
	response       response.IMResponse
	postValidation validation.PostValidation
}

// NewPostHandler func new post handler
func NewPostHandler(
	postService service.PostService,
	response response.IMResponse,
	postValidation validation.PostValidation,
) *PostHandler {
	return &PostHandler{
		postService:    postService,
		response:       response,
		postValidation: postValidation,
	}
}

// Add func godoc
// @Summary Add a new post
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param PostAddRequest body schema.PostAddRequest true "Add a new post body"
// @Success 200 {object} schema.PostDetailResponse "success"
// @Router /identity/v1/post [post]
func (p *PostHandler) Add(ctx *gin.Context) {
	postAddRequest := &schema.PostAddRequest{}
	err := ctx.BindJSON(postAddRequest)
	if err != nil {
		p.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	err = p.postValidation.Add(postAddRequest)
	if err != nil {
		p.response.Out(ctx, err, nil)
		return
	}

	postDetailResponse, err := p.postService.Add(postAddRequest)
	if err != nil {
		p.response.Out(ctx, err, nil)
		return
	}
	p.response.Out(ctx, err, postDetailResponse)
}

// Detail func godoc
// @Summary Detail info of post
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param id path int true "number"
// @Success 200 {object} schema.PostDetailResponse "success"
// @Router /identity/v1/post/{id} [get]
func (p *PostHandler) Detail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		p.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	postDetailResponse, err := p.postService.Detail(uint64(id))
	p.response.Out(ctx, err, postDetailResponse)
	return
}

// Update func godoc
// @Summary Update a post
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param id path int true "number"
// @Param PostUpdateRequest body schema.PostUpdateRequest true "Update a post body"
// @Success 200 {object} response.IMResponse "success"
// @Router /identity/v1/post/{id} [put]
func (p *PostHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		p.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	postUpdateRequest := &schema.PostUpdateRequest{}
	err := ctx.BindJSON(postUpdateRequest)
	if err != nil {
		p.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	err = p.postValidation.Update(postUpdateRequest)
	if err != nil {
		p.response.Out(ctx, err, nil)
		return
	}

	err = p.postService.Update(id, postUpdateRequest)
	p.response.Out(ctx, err, nil)
}

// Delete func godoc
// @Summary Delete a post
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param id path int true "number"
// @Success 200 {object} response.IMResponse "success"
// @Router /identity/v1/post/{id} [delete]
func (p *PostHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		p.response.Out(ctx, errors.New("not_allow"), nil)
		return
	}

	err := p.postService.Delete(id)
	p.response.Out(ctx, err, nil)
}

// List func godoc
// @Summary List posts by user
// @Description Author: LichTV
// @Tags identity
// @Accept json
// @Produce json
// @Param lang query string false "string" enums(en, vi)
// @Param user_id query int false "user id"
// @Param page query int false "page number"
// @Param page_size query int false "page size"
// @Success 200 {object} schema.PostListResponse "success"
// @Router /identity/v1/post [get]
func (p *PostHandler) List(ctx *gin.Context) {
	userID, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	postListResponse, err := p.postService.List(userID, page, pageSize)
	p.response.Out(ctx, err, postListResponse)
}