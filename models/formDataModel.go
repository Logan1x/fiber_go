package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormData struct {
	gorm.Model
	ProjectID       uuid.UUID // Foreign key
	FormSubject     string
	FormMessage     string
	FormSentByName  *string // optional field
	FormSentByEmail *string // optional field

}
