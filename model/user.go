package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	StudentId string `json:"studentId" gorm:"unique"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	MBTI      string `json:"mbti"`
}
