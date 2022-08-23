package monggodb

import "fmt"

func (ed *EventDMonggoDB) UpdateEventCounter() error {
	fmt.Println("updated events")

	return nil
}
