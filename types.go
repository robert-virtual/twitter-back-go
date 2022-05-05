package main

// los nombres de campos deben ser en mayuscula para q sean
// accedibles de lo contrario seran campos privados

type post struct {
	Id        *string  `json:"id"`        //db
	Content   string   `json:"content"`   //db
	Images    *[]image `json:"images"`    //lista de imagenes
	UserId    string   `json:"userId"`    //db
	CreatedAt *string  `json:"cretaedAt"` //db
	User      *user    `json:"user"`
}
type user struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Password *string `json:"password"` // es tipo *string es un  puntero a un string en este caso lo uso solo para poder dejar este campo en null
}
type image struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	PostId string `json:"postId"`
}
