package service

import (
	"soulvent/internal/dto"
	"soulvent/internal/model"
	"soulvent/internal/repository"
)

func SearchUsers(userId string , searchReq *dto.SearchRequest)(*dto.SearchResponse, error){
	// Get users with pagination
	users, err := repository.SearchUsers(searchReq.Query, userId, searchReq.Page, searchReq.Limit)
	if err != nil {
		return nil, err
	}

	// Build response
	response := &dto.SearchResponse{
		Query:       searchReq.Query,
		CurrentPage: searchReq.Page,
		Users:       users,
	}
	return response, nil
}

func AddSearch(addSearchReq *model.Search) error{
	return repository.AddSearch(addSearchReq)
}

func GetSearches(userID string, limit int)([]*model.Search, error){
	return repository.GetSearches(userID,limit)
}