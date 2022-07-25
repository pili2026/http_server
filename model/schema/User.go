package schema

type User struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}
