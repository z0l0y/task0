package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"task0/go/dao"
	"task0/go/entity"
)

func CreateUser(user *entity.User) (err error) {
	// 使用bcrypt，对用户的密码进行加密
	passwordBytes := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

func GetUser(user *entity.User) (err error) {
	err = AuthenticateUser(user.Name, user.Password)
	if err != nil {
		return err
	}
	// 注意这里我们不能用#{}，这是Java中占位符的写法，在gorm中是不正确的，在gorm中应为？
	errQuery := dao.SqlSession.Where("name = ?", user.Name).First(&user).Error
	if errQuery != nil {
		if gorm.IsRecordNotFoundError(errQuery) {
			// 没有找到匹配的记录
			return errors.New("未找到匹配的用户") // 或者返回相应的错误信息
		}
		return errQuery // 其他查询错误
	}
	return
}

func GetAllUser() ([]*entity.User, error) {
	var userList []*entity.User
	err := dao.SqlSession.Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func DeleteUserByUsername(name string) (err error) {
	err = dao.SqlSession.Where("name = ?", name).Delete(&entity.User{}).Error
	return
}

func GetUserByName(username string) (user *entity.User, err error) {
	user = &entity.User{} // 将变量赋值为指向 User 类型的指针
	if err = dao.SqlSession.Where("name = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User) (err error) {
	// 使用bcrypt，对用户的密码进行加密
	passwordBytes := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = dao.SqlSession.Save(user).Error
	return
}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(username, password string) error {
	// 从数据库中获取用户名对应的加密密码
	var user entity.User
	err := dao.SqlSession.Where("name = ?", username).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 没有找到对应的用户记录
			return errors.New("用户名不存在")
		}
		// 其他数据库查询错误
		return err
	}
	// 验证密码的正确性
	err = VerifyPassword(user.Password, password)
	if err != nil {
		// 密码不匹配
		return errors.New("密码不正确")
	}

	return nil
}
