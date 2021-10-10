package dao

import (
	"savi8sant8s/api/entity"
)

type SessionDao struct {}

func (this *SessionDao) Create(session entity.Session) entity.Session {
	db.Select("UserID", "Token", "Expired").Create(&session)
	return session
}

func (this *SessionDao) Close(token string)bool {
	db.Model(&entity.Session{}).Where("expired = ?", false).Where("token = ?", token).Update("expired", true)
	return true
}

func (this *SessionDao) Valid(token string) bool {
	count := int64(0) 
	db.Take(&entity.Session{}).Where("token = ? AND expired = ?", token, false).Count(&count)
	return count > 0
}

func (this *SessionDao) UserID(token string) uint {
	session := new(entity.Session)
	db.Select("user_id").Where("token = ?", token).Find(&session)
	return session.UserID
}