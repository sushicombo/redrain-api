package monggodb

type EventDMonggoDB struct {
	connectionMonggo string
}

func InitEventRepoMonggo(conn string) EventDMonggoDB {
	return EventDMonggoDB{
		connectionMonggo: conn,
	}
}
