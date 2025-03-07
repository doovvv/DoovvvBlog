package model

import (
	"doovvvblog/utils"
	"doovvvblog/utils/errmsg"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;" json:"username" validate:"required,min=4,max=12" `
	Password string `gorm:"type:varchar(100);not null;" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=1,lte=2"`
}

func (u *User) TableName() string {
	return "user"
}

// 新增用户
func CreateUser(u *User) int {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errmsg.ERROR
	}
	u.Password = hashedPassword
	err = DB.Create(&u).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查用户名是否存在
func CheckUser(username string) (code int) {
	var users User
	DB.Select("id").Where("username = ?", username).Limit(1).First(&users)
	fmt.Println(username)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 查询所有用户
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := DB.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 删除用户
func DeleteUser(id int) (code int) {
	var user User
	err := DB.Where("id =?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
func EditUser(id uint, u *User) (code int) {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = u.Username
	maps["role"] = u.Role
	err := DB.Model(&user).Where("id =?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登录
func CheckLogin(username string, password string) int {
	var user User
	err := DB.Where("username =?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
