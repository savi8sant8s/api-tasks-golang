package validation

import (
	"regexp"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
)

func ValidTask(task entity.Task) (bool, string) {
	matchTaskTitle, _ := regexp.MatchString(REGEX_TASK_TITLE, task.Title)
	if !matchTaskTitle {
		return false, utils.ERROR_TASK_TITLE
	}
	matchTaskMessage, _ := regexp.MatchString(REGEX_TASK_MESSAGE, task.Message)
	if !matchTaskMessage {
		return false, utils.ERROR_TASK_MESSAGE
	}
	return true, utils.BLANK
}
