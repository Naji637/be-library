package util

import (
	"be-library/app/model"
	"fmt"
	"strconv"
	"time"
)

type Barcode interface {
	CreateBarcode() string
}
type barcode struct {
	repository Repository
}

func NewBarcode(repository Repository) *barcode {
	return &barcode{repository}
}
func getYearMonth() (string, string) {
	now := time.Now()
	year := now.Format("06")
	month := now.Format("01")
	return year, month
}
func (b *barcode) CreateBarcode() string {
	year, month := getYearMonth()
	yearNumber, _ := strconv.Atoi(year)
	monthNumber, _ := strconv.Atoi(month)
	barcode := b.repository.GetByYearsMonths(uint8(yearNumber), uint8(monthNumber))
	if barcode.ID == "" {
		if err := b.repository.CreateBarcode(model.Barcode{
			Year:   uint8(yearNumber),
			Month:  uint8(monthNumber),
			Number: 1,
		}); err == nil {
			return fmt.Sprintf("B%s%s%04d", year, month, 1)
		}
	} else {
		newNumber := barcode.Number + 1
		if err := b.repository.UpdateBarcode(uint8(yearNumber), uint8(monthNumber), newNumber); err == nil {
			return fmt.Sprintf("B%s%s%04d", year, month, newNumber)
		}
	}
	return ""

}
