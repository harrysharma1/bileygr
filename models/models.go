package models

type Users struct {
	id           string `json:"id", database:"id"`
	username     string `json:"username", database:"username"`
	email        string `json:"email", database:"email"`
	password     string `json:"password", database:"password"`
	profileImage string `json:"profile_image", database:"profile_image"`
	createdAt    string `json:"created_at", database:"created_at"`
}
