package service

import (
	"imgo/app/modules/identity/v1/entity"
	"imgo/app/modules/identity/v1/repository"
	schema "imgo/app/schema/identity/v1"
	"imgo/app/utils"
)

// PostService interface of post service object
type PostService interface {
	Add(postAddRequest *schema.PostAddRequest) (*schema.PostDetailResponse, error)
	Detail(id uint64) (*schema.PostDetailResponse, error)
	Update(id string, postUpdate *schema.PostUpdateRequest) error
	Delete(id string) error
	List(userID uint64, page, pageSize int) (*schema.PostListResponse, error)
}

type postService struct {
	postRepo  repository.PostRepository
	convert   utils.IMConvert
}

// NewPostService func new post service object
func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
		convert:  utils.NewIMConvert(),
	}
}

// Add func add new post
func (p *postService) Add(postAddRequest *schema.PostAddRequest) (*schema.PostDetailResponse, error) {
	result := &schema.PostDetailResponse{}
	post := &entity.Post{
		UserID:  postAddRequest.UserID,
		Title:   postAddRequest.Title,
		Content: postAddRequest.Content,
		Status:  entity.PostStatusDefault,
	}

	err := p.postRepo.Add(post)
	if err != nil {
		return result, err
	}

	err = p.convert.Object(post, &result)
	return result, err
}

// Detail func get detail post info
func (p *postService) Detail(id uint64) (*schema.PostDetailResponse, error) {
	result := &schema.PostDetailResponse{}
	post, err := p.postRepo.Detail(id)
	if err != nil {
		return result, p.convert.DatabaseError(err)
	}

	err = p.convert.Object(post, &result)
	return result, err
}

// Update func update post info
func (p *postService) Update(id string, postUpdate *schema.PostUpdateRequest) error {
	post := &entity.Post{
		Title:   postUpdate.Title,
		Content: postUpdate.Content,
		Status:  postUpdate.Status,
	}
	return p.postRepo.Update(id, post)
}

// Delete func delete post info
func (p *postService) Delete(id string) error {
	return p.postRepo.Delete(id)
}

// List func get list post by user
func (p *postService) List(userID uint64, page, pageSize int) (*schema.PostListResponse, error) {
	result := &schema.PostListResponse{}
	posts, total, err := p.postRepo.List(userID, page, pageSize)
	if err != nil {
		return result, p.convert.DatabaseError(err)
	}

	var postResponses []schema.PostDetailResponse
	for _, post := range posts {
		var postResp schema.PostDetailResponse
		p.convert.Object(post, &postResp)
		postResponses = append(postResponses, postResp)
	}

	result.Posts = postResponses
	result.Total = total
	result.Page = page
	result.PageSize = pageSize

	return result, nil
}