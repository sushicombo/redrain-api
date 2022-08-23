package mysql

func (ed *EventDB) UpdateEventCounter() error {
	_, err := ed.db.DB.Exec(abdul)

	if err != nil {
		return err
	}

	return nil
}
