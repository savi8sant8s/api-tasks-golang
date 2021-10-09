package data

type Register struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ErrorMessage struct {
	Erro bool `json:"erro"`
	Message string `json:"message"`
}

type Task struct {
	Title   string
	Message string
}