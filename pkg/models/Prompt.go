// Prompt Model for GORM / Database
package models

import "gorm.io/gorm"

type Prompt struct {
	gorm.Model
	PromptText string     `gorm:"unique" json:"promptText"`
	Category   string     `json:"category"`
	Responses  []Response `json:"responses"`
}
