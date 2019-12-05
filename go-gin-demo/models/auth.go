package models

//type Auth struct {
//	ID       int    `gorm:"primary_key" json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

func CheckAuth(userId int) bool {
	var user User
	db.Select("user_id").Where("user_id = ?", userId).First(&user)
	if user.UserId > 0 {
		return true
	}
	return false
}
