package Models

import "time"

type Post struct {
    Id uint `gorm:"primaryKey"`
    Name string `gorm:"size:255"`
    Password string `gorm:"size:255"`
    Created_at time.Time
    Updated_at time.Time
}
