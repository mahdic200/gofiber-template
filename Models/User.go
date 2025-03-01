package Models

import "time"

type User struct {
    id uint
    name string
    password string
    created_at time.Time
    updated_at time.Time

}
