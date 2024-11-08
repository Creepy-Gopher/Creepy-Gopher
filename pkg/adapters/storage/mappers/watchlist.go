package mappers

import (
	"creepy/internal/watchlist"
	"creepy/pkg/adapters/storage/entities"
)

func WatchListEntityToDomain(w *entities.WatchList) *watchlist.WatchList {
	user := UserEntityToDomain(&w.User)
	filter := FilterEntityToDomain(&w.Filter)
	return &watchlist.WatchList{
		ID:         w.ID,
		UserID:     w.UserID,
		User:       *user,
		FilterID:   w.FilterID,
		Filter:     *filter,
		FilterName: w.FilterName,
	}
}

func WatchListDomainToEntity(w *watchlist.WatchList) *entities.WatchList {
	user := UserDomainToEntity(&w.User)
	filter := FilterDomainToEntity(&w.Filter)
	return &entities.WatchList{
		Model:      entities.Model{ID: w.ID},
		UserID:     w.UserID,
		User:       *user,
		FilterID:   w.FilterID,
		Filter:     *filter,
		FilterName: w.FilterName,
	}
}
