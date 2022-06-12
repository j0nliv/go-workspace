package models

import "time"



type Product struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Description string 		`json:"description"`
	CreatedOn 	time.Time 	`json:"createdOn"`
	ChangedOn 	time.Time 	`json:"changedOn"`
}