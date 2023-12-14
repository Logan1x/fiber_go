package models

import (
	"github.com/google/uuid"
)

type FormData struct {
	BaseModel
	ProjectID       uuid.UUID // Foreign key
	FormSubject     string
	FormMessage     string
	FormSentByName  *string // optional field
	FormSentByEmail *string // optional field

}
