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
	session := new(entity.Session)
	db.Model(&session).Where("expired = ?", false).Update("expired", true)
	return *session
}

func (sd *SessionDao) Expired(token string) bool {
	valid := db.Where("token = ?", token).Where("expired = ?", true).Take(&entity.Session{})
	return valid.RowsAffected > 0
}

func (sd *SessionDao) UserID(token string) uint {
	session := new(entity.Session)
	db.Select("user_id").Where("token = ?", token).Find(&session)
	return session.UserID
}