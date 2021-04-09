package models

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserModel struct {
	Id    int64  `gorm:"type:int"`
	Name  string `gorm:"type:text"`
	Email string `gorm:"type:text"`
}
