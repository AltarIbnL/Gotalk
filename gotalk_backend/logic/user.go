package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

// 存放业务逻辑代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	var exist bool
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}
	if exist {
		//用户已存在
		return errors.New("用户已存在")
	}
	//mysql.QueryUserByName()
	// 生成UID
	userID := snowflake.GenID()

	// 构造user实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	//mysql.SignUp()
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
