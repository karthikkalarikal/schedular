package domain

import "time"

type Employee struct {
	Name       string `json:"name" validate:"required"`
	Department string `json:"department" validate:"required"`
	Role       string `json:"role" validate:"required"`

	// Availability	[]Availability
}

type EmployeeResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
}

type Availability struct {
	Day       time.Weekday `json:"day"`
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
}

type Response struct {
	Code int
	Data interface{}
	Err  string
}
