package event

import (
	"github.com/sushicombo/redrain-api/repository/events"
)

type Event struct {
	repo events.EventRepository
}

func InitEventUsecase(repo events.EventRepository) Event {
	return Event{
		repo: repo,
	}
}
