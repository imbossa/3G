// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Translation struct {
	Source      string `json:"source"       gorm:"size:255"  example:"auto"`
	Destination string `json:"destination"  gorm:"size:255"  example:"en"`
	Original    string `json:"original"     gorm:"type:text" example:"текст для перевода"`
	Translation string `json:"translation"  gorm:"type:text" example:"text for translation"`
}
