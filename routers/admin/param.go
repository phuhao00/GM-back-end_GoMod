package admin

type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	User_sex  int 	 `json:"user_sex"`
	Nick_name string `json:"nick_name"`
}

type UpdateReq struct {
	Name      string `json:"name"`
	ColumnName  string
	ColumnVal  string
}
type GetUserHavePlayGamesReq struct {
	Name      string `json:"name"`
}
