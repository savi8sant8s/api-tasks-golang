package dao

import (
	"savi8sant8s/api/database"
	"savi8sant8s/api/entity"
)

type SessionDao struct {
	db database.Database
}

func (this *SessionDao) CreateSession(session entity.Session) {
	this.db.Instance().Select("UserID", "Token", "Expired").Create(&session)
}

func (this *SessionDao) CloseSession(token string) {
	this.db.Instance().Model(&entity.Session{}).Where("expired = ?", false).Where("token = ?", token).Update("expired", true)
}

func (this *SessionDao) CloseLastSessions(userId uint) {
	this.db.Instance().Model(&entity.Session{}).Where("expired = ?", false).Where("user_id = ?", userId).Update("expired", true)
}

func (this *SessionDao) ValidToken(token string) bool {
	count := int64(0) 
	this.db.Instance().Take(&entity.Session{}).Where("token = ? AND expired = ?", token, false).Count(&count)
	return count > 0
}

func (this *SessionDao) GetUserID(token string) uint {
	session := new(entity.Session)
	this.db.Instance().Select("user_id").Where("token = ?", token).Find(&session)
	return session.UserID
}