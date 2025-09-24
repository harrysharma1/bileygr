package models

type Users struct {
	Id           string `json:"id", database:"id"`
	Username     string `json:"username", database:"username"`
	Email        string `json:"email", database:"email"`
	Password     string `json:"-", database:"password"`
	ProfileImage string `json:"profile_image", database:"profile_image"`
	CreatedAt    string `json:"created_at", database:"created_at"`
}
