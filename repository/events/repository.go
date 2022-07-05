package events

type EventRepository interface {
	UpdateEventCounter() error
}
