package main

import "gorm.io/gorm"

//return &domain.User{UserID: 12243}, nil
type User struct {
	gorm.Model
}
