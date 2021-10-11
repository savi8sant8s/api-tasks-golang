package data

import "savi8sant8s/api/entity"

type MessageTasks struct {
	Message
	Tasks []entity.Task `json:"tasks"`
}