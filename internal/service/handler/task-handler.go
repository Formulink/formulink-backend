package handler

import (
	"database/sql"
	"formulink-backend/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type TaskHandler struct {
	db    *sql.DB
	redis *redis.Client
}

func NewTaskHandler(_db *sql.DB, _redis *redis.Client) *TaskHandler {
	return &TaskHandler{
		db:    _db,
		redis: _redis,
	}
}

func (th *TaskHandler) GetTasksByFormulaId(c echo.Context) error {
	var tasks []model.Task
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := `SELECT * from tasks where formula_id = $1`
	rows, err := th.db.Query(query, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}

	for rows.Next() {
		var task model.Task
		if err = rows.Scan(
			&task.Id,
			&task.FormulaId,
			&task.Difficulty,
			&task.TaskText,
			&task.Result,
		); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		tasks = append(tasks, task)
	}
	return c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) GetTaskById(c echo.Context) error {
	var task model.Task
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := `SELECT * from tasks where id = $1`
	row := th.db.QueryRow(query, id)
	if row == nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}
	err = row.Scan(
		&task.Id,
		&task.FormulaId,
		&task.Difficulty,
		&task.TaskText,
		&task.Result,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, task)
}
