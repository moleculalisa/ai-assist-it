package models

import "time"

 //User представляет пользователя системы
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

//CreateUserRequest содержит данные для создания нового пользователя
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,max=100"`
}
//UpdateUserRequest содержит данныке для обновления пользователя
type UpdateUserRequest struct {
	Email *string `json:"email" validate:"omitempty,email"`
	Name  *string `json:"name" validate:"omitempty,max=100"`
}

//ErrorResponse представляет ошибку API
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

//SuccessResponse представляет успешный ответ API
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}
