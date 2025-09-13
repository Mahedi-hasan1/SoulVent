package dto

type CreatePostRequest struct {
	//UserID    string   `json:"user_id" validate:"required,uuid4"`
	Content   string   `json:"content" validate:"max=10000"`
	ImageURLs []string `json:"image_urls,omitempty" validate:"omitempty,dive,url,max=100"` // max 100 images, each a valid URL

}