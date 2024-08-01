package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("INVALID YEAR")
	}
	d.year = year
	return nil
}
func (d *Date) Year() int { //Get-метод для года
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("INVALID MONTH")
	}
	d.month = month
	return nil
}
func (d *Date) Month() int { //Get-метод для месяца
	return d.month
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("INVALID DAY")
	}
	d.day = day
	return nil
}
func (d *Date) Day() int {
	return d.day
}
