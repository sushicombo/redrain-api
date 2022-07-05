package events

func (ed *EventDB) UpdateEventCounter() error {
	_, err := ed.db.DB.Exec(`UPDATE events SET counter = counter + 1 WHERE id = 1`)

	if err != nil {
		return err
	}

	return nil
}
