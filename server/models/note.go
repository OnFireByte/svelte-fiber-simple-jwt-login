package models

type NoteStatus string

const (
	NoteStatusActive  NoteStatus = "active"
	NoteStatusSuccess NoteStatus = "success"
)

type Note struct {
	OwnModel
	Uuid     string     `json:"uuid" gorm:"unique;not null;index"`
	Content  string     `json:"content"`
	Owner_id uint       `json:"-"`
	Owner    User       `json:"-" gorm:"foreignkey:Owner_id"`
	Status   NoteStatus `json:"status"`
}