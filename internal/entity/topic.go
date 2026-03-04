package entity

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`   // Can be text, URL, or file path
	Type      string `json:"type"`      // e.g., "video", "pdf", "text"
	SectionID uint   `json:"section_id"`
}
