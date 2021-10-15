package data

import "savi8sant8s/gotasks/entity"

type MessageTasks struct {
	Message
	Tasks []entity.Task `json:"tasks"`
}
