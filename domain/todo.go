package domain

type TodoList struct {
	ID          int    `json: "-"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

type UsersList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json: "-"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}

type ListsItems struct {
	ID     int
	ListID int
	ItemID int
}
