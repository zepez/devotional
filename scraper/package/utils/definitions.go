package utils

import "time"

type Devotional struct {
	Id          string    `json:"id"`
	Source      string    `json:"source"`
	Name        string    `json:"name"`
	Html        string    `json:"html"`
	Plain       string    `json:"plain_text"`
	Image       string    `json:"image"`
	Target_Date string    `json:"target_date"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
