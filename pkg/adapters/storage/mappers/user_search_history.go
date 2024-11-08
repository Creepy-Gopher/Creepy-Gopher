package mappers

import (
	history "creepy/internal/user_search_history"
	"creepy/pkg/adapters/storage/entities"
)

func UserSearchHistoryEntityToDomain(h *entities.UserSearchHistory) *history.UserSearchHistory {
	user := UserEntityToDomain(&h.User)
	filter := FilterEntityToDomain(&h.Filter)
	return &history.UserSearchHistory{
		ID:       h.ID,
		UserName: h.UserName,
		User:     *user,
		FilterID: h.FilterID,
		Filter:   *filter,
		Date:     h.Date,
	}
}

func UserSearchHistoryDomainToEntity(h *history.UserSearchHistory) *entities.UserSearchHistory {
	user := UserDomainToEntity(&h.User)
	filter := FilterDomainToEntity(&h.Filter)
	return &entities.UserSearchHistory{
		Model:    entities.Model{ID: h.ID},
		UserName: h.UserName,
		User:     *user,
		FilterID: h.FilterID,
		Filter:   *filter,
		Date:     h.Date,
	}
}
