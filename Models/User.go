package models

import "time"

type User struct {
    Id uint `json:"id" gorm:"not null;primaryKey"`
    Name string `json:"name" gorm:"size:255"`
    Password string `json:"password" gorm:"size:255"`
    Created_at time.Time `json:"created_at"`
    Updated_at time.Time `json:"updated_at"`
}
