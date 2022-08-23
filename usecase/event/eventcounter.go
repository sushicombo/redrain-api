package event

func (e Event) UpdateEventCounter() error {
	err := e.repo.UpdateEventCounter()

	return err
}
