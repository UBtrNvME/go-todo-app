package handler

import (
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call service
	id, err := h.services.TodoListInterface.Create(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []domain.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	// Call service
	lists, err := h.services.TodoListInterface.GetAll(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
func (h *Handler) getListById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}
	// Get id of the list
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	// Call service
	list, err := h.services.TodoListInterface.GetById(userID, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
