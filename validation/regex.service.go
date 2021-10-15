package validation

const (
	REGEX_NAME = "^[a-zA-ZáéíóúÁÉÍÓÚÃÕãõç ]{3,34}$"
	REGEX_EMAIL = "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"
	REGEX_PASSWORD = "^.{8,32}$"
	REGEX_TASK_TITLE = "^.{3,20}$"
	REGEX_TASK_MESSAGE = "^.{3,100}$"
)
