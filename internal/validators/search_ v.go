package validators

import (
	"errors"
	"soulvent/internal/dto"
	"soulvent/internal/model"
)

func  ValidateSearchRequest(req *dto.SearchRequest) error {
	if req.Query == "" {
		return errors.New("search query is required")
	}
	if len(req.Query) > 100 {
		return errors.New("search query too long")
	}
	if req.Page < 1 {
		return errors.New("page must be greater than 0")
	}
	if req.Limit < 1 || req.Limit > 50 {
		return errors.New("limit must be between 1 and 50")
	}
	return nil
}

func ValidateAddSearch(req *model.Search)error{
	if req.Query == "" {
		return errors.New("search query is required")
	}
	if len(req.Query) > 100 {
		return errors.New("search query too long")
	}
	return nil
}