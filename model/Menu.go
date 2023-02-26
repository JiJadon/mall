package model

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Path        string `json:"path" form:"path"`
	Icon        string `json:"icon" form:"icon"`
	Description string `json:"description" form:"description"`
	Pid         int    `json:"pid" form:"pid"`
	PagePath    string `json:"pagePath" form:"page_path"`
	Sort        int    `json:"sort" form:"sort"`
}
