package repository

import (
	"imgo/app/modules/identity/v1/entity"
	"imgo/app/utils"

	"gorm.io/gorm"
)

// PostRepository interface of post repository object
type PostRepository interface {
	Add(post *entity.Post) error
	Detail(id uint64) (*entity.Post, error)
	Update(id string, post *entity.Post) error
	Delete(id string) error
	List(userID uint64, page, pageSize int) ([]entity.Post, int64, error)
}

type postRepository struct {
	db     *gorm.DB
	imTime utils.IMTime
}

// NewPostRepository func new post repository object
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db, imTime: utils.NewIMTime()}
}

// Add func add new post
func (p *postRepository) Add(post *entity.Post) error {
	post.CreatedAt = p.imTime.TimeDB()
	post.ModifiedAt = p.imTime.TimeDB()
	return p.db.Create(&post).Error
}

// Detail func get detail post info
func (p *postRepository) Detail(id uint64) (*entity.Post, error) {
	result := &entity.Post{}
	err := p.db.Where("id = ?", id).First(result).Error
	return result, err
}

// Update func update post info
func (p *postRepository) Update(id string, post *entity.Post) error {
	post.ModifiedAt = p.imTime.TimeDB()
	return p.db.Model(&entity.Post{}).Where("id = ?", id).Updates(post).Error
}

// Delete func delete post info
func (p *postRepository) Delete(id string) error {
	return p.db.Where("id = ?", id).Delete(&entity.Post{}).Error
}

// List func get list post by user
func (p *postRepository) List(userID uint64, page, pageSize int) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var total int64

	query := p.db.Model(&entity.Post{})
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return posts, total, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&posts).Error

	return posts, total, err
}