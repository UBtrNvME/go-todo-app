package repository

import (
	"fmt"
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r TodoListPostgres) Create(userID int, list domain.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r TodoListPostgres) GetAll(userID int) ([]domain.TodoList, error) {
	var lists []domain.TodoList
	getListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, getListsQuery, userID)

	return lists, err
}

func (r TodoListPostgres) GetById(userID, listId int) (domain.TodoList, error) {
	var list domain.TodoList
	getListQuery := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id 
				WHERE ul.user_id = $1 AND ul.user_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, getListQuery, userID, listId)

	return list, err
}

func (r TodoListPostgres) Update(userID, listId int, data interface{}) error {
	panic("implement me")
}

func (r TodoListPostgres) Delete(userID, listId int) error {
	panic("implement me")
}
