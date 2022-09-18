package models

type User struct {
	// db 表示的是数据库里面的tag
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
