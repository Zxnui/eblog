package model

type MenuDetail struct {
	Name string	`json:"name"`
	Id int	`json:"id"`
	ParentId int `json:"parentId"`
}

type Menu struct {
	Menu_list []MenuDetail `json:"menu_list"`
	MaxId int `json:"maxId"`
}
