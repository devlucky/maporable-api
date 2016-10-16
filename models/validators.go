package models

import "time"

func validateDate(d string) (error) {
	_, err := time.Parse(time.RFC3339, d)
	return err
}