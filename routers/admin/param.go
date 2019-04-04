package admin

type LoginReq struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

type RegisterReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Roles    []string `json:"roles"`
	Introduce string `json:"introduce"`
}