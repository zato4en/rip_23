package ds

import (
	"time"
)

type Satellite struct {
	//Когда миграции создают GORM, для некоторых типов данных надо
	//добавить ограничения
	//JSON мы указываем для обращения к полям структуры извне через JSON запросы
	ID           uint      `json:"id" gorm:"primaryKey"`
	DateCreated  time.Time `json:"date_created"`
	DateFormed   time.Time `json:"date_formed"`
	DateAccepted time.Time `json:"date_accepted"`
	Status       string    `gorm:"type:varchar(255)" json:"status"`
	Satellite    string    `gorm:"type:varchar(255)" json:"satellite"`
	UserID       uint      `json:"-"`
	ModerID      uint      `json:"-"`
	User         Users     `gorm:"foreignKey:UserID" json:"-"`
}
