package todo

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

type UserList struct {
	Id     int
	UserId int
	ListID int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id      int
	ListsId int
	Item    int
}
