package events

type Events struct {
	ID      int64  `json:"id" db:"id"`
	Event   string `json:"event" db:"event"`
	Counter int64  `json:"counter" db:"counter"`
}

type EventsReq struct {
	ID int64 `json:"id" db:"id"`
}

type EventsResp struct {
	Event string `json:"event" db:"event"`
}
