package structs

import "time"

type Project struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Budget       float64   `json:"budget"`
	DepartmentID int       `json:"department"`
}
