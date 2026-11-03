package util

import (
	"be-library/app/model"
	"fmt"
	"strconv"
	"time"
)

type Member interface {
	CreateMembership() string
}
type member struct {
	repository Repository
}

func NewMember(repository Repository) *member {
	return &member{repository}
}

func getYearMonth() (string, string) {
	now := time.Now()
	year := now.Format("06")
	month := now.Format("01")
	return year, month
}

func (m *member) CreateMembership() string {
	year, month := getYearMonth()
	yearNumber, _ := strconv.Atoi(year)
	monthNumber, _ := strconv.Atoi(month)
	member := m.repository.GetByYearsMonths(uint8(yearNumber), uint8(monthNumber))
	if member.ID == "" {
		if err := m.repository.CreateMembership(model.Membership{
			Year:   uint8(yearNumber),
			Month:  uint8(monthNumber),
			Number: 1,
		}); err == nil {
			return fmt.Sprintf("A%s%s%04d", year, month, 1)
		}
	} else {
		newNumber := member.Number + 1
		if err := m.repository.UpdateMembership(uint8(yearNumber), uint8(monthNumber), newNumber); err == nil {
			return fmt.Sprintf("A%s%s%04d", year, month, newNumber)
		}
	}
	return ""
}
