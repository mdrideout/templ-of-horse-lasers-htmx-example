package models

import "gorm.io/gorm"

type LeaderboardItemStore struct {
	gorm.Model
	Name  string `form:"name"`
	Score int    `form:"score"`
}
