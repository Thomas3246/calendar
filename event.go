package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("INVALID TITLE")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string { //Get-method for titile
	return e.title
}
