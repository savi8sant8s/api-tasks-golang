package validation

const REGEX_NAME string = "^[a-zA-ZáéíóúÁÉÍÓÚÃÕãõç ]{3,34}$"
const REGEX_EMAIL string = "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"
const REGEX_PASSWORD string = "^.{8,32}$"

const REGEX_TASK_TITLE string = "^.{3,20}$"
const REGEX_TASK_MESSAGE string = "^.{3,100}$"
