// Prompt Model for GORM / Database
package models

import "gorm.io/gorm"

type Response struct {
	gorm.Model
	ResponseText string `gorm:"size:1400" json:"responseText"`
	Category     string `json:"category"`
	PromptID     uint   `json:"promptId"`
}
