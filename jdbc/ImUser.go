package jdbc

type ImUser struct {
	UserId int8 `db:"user_id"`
	UserName string `db:"user_name"`
	NickName string	`db:"nick_name"`
}
