package model

type User struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
    Gender      string  `json:"gender"`
    DateOfBirth string  `json:"date_of_birth"`
    Email       string  `json:"email"`
    Phone       string  `json:"phone"`
    Password    string  `json:"password"`
}

type LoginRequest struct {
    Email       string  `json:"email"`
    Password    string  `json:"password"`
}
