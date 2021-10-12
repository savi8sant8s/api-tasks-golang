package utils

const (
	BLANK = ""

	EMAIL_ALREADY_REGISTERED = "Já existe um usuário cadastrado com esse email."
	EMAIL_NOT_REGISTERED = "Não existe usuário cadastrado com esse email."
	USER_REGISTERED = "Usuário cadastrado com sucesso."
	INVALID_CREDENTIALS = "As credenciais informadas estão incorretas."
	LOGIN_SUCCESS = "Login realizado com sucesso."
	LOGOUT_SUCCESS = "Logout realizado com sucesso"

	ERROR_NAME = "O nome deve possuir de 3 a 34 letras."
	ERROR_EMAIL = "O email deve seguir o padrão fulano@email.com."
	ERROR_PASSWORD = "A senha deve possuir de 8 a 32 caracteres."
	ERROR_INCORRECT_AUTH_HEADER = "Token não especificado corretamente."
	ERROR_INVALID_TOKEN = "Token inválido ou sessão expirada."

	ERROR_TASK_TITLE = "O título da tarefa deve possuir de 3 a 20 caracteres."
	ERROR_TASK_MESSAGE = "A mensagem da tarefa deve possuir de 3 a 100 caracteres."

	CREATE_TASK_SUCCESS = "Tarefa criada com sucesso."
	ALL_TASKS_SUCCESS = "Todas as tarefas retornadas com sucesso."
	DELETE_TASK_SUCCESS = "Tarefa deletada com sucesso."
	UPDATE_TASK_SUCCESS = "Tarefa alterada com sucesso."
	NOT_FOUND_TASK_SUCCESS = "Tarefa não encontrada com esse ID."
)