package models

import "github.com/google/uuid"

type Project struct {
	BaseModel
	UserID      uuid.UUID // Foreign key
	ProjectName string
	ProjectDesc *string
	Forms       []FormData // Association with FormData
}
