package dao

import (
	"savi8sant8s/api/entity"
)

type SessionDao struct {}

func (sd *SessionDao) Create(session entity.Session) entity.Session {
	db.Select("UserID", "Token", "Expired").Create(&session)
	return session
}

func (sd *SessionDao) Close(token string) entity.Session {
	var session = new(entity.Session)
	db.Model(&session).Where("expired = ?", false).Update("expired", true)
	return *session
}