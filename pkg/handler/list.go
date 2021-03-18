package handler

import (
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (h *Handler) getAllLists(c *gin.Context) {

}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
