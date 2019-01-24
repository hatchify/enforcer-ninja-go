package sdk

const (
	// ViewDay represents a view range of the current day
	ViewDay View = "day"
	// ViewWeek represents a view range of the last seven days
	ViewWeek View = "week"
	// ViewMonth represents a view range of the last thirty days
	ViewMonth View = "month"
)

// View represents the view range for reporting data
type View string

// Validate will validate the selected View
func (v View) Validate() (err error) {
	switch v {
	case ViewDay:
	case ViewWeek:
	case ViewMonth:

	default:
		return ErrInvalidView
	}

	return
}
