package formulas

import (
	"database/sql"
	"errors"
	"formulink-backend/internal/dto"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"net/http"
)

type FormulaLikesHandler struct {
	db *sql.DB
}

func NewFormulaLikesHandler(db *sql.DB) *FormulaLikesHandler {
	return &FormulaLikesHandler{db: db}
}

func (fh *FormulaLikesHandler) HandleLike(c echo.Context) error {
	var req dto.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := `SELECT 1 FROM formula_likes WHERE user_id = $1 AND formula_id = $2`
	row := fh.db.QueryRow(query, req.UserId, req.FormulaId)

	var exists int
	err := row.Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		query = `INSERT INTO formula_likes (user_id, formula_id) VALUES ($1, $2)`
	} else if err == nil {
		query = `DELETE FROM formula_likes WHERE user_id = $1 AND formula_id = $2`
	} else {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if _, err = fh.db.Exec(query, req.UserId, req.FormulaId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (fh *FormulaLikesHandler) GetStatus(c echo.Context) error {
	var req dto.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := `SELECT EXISTS (SELECT 1 FROM formula_likes WHERE user_id = $1 AND formula_id = $2)`
	row := fh.db.QueryRow(query, req.UserId, req.FormulaId)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if exists {
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (fh *FormulaLikesHandler) GetAllLikedFormulas(c echo.Context) error {
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var formulasIDs []uuid.UUID
	query := `SELECT * from formula_likes where user_id = $1`
	rows, err := fh.db.Query(query, id)
	if err != nil {
		logger.Lg().Logf(0, "err: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		var f model.FormulaLike
		if err = rows.Scan(&f.UserID, &f.FormulaID, &f.CreatedAt); err != nil {
			logger.Lg().Logf(0, "can't parse db data into strnig | err: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		formulasIDs = append(formulasIDs, f.FormulaID)
	}

	var formulas []model.Formula

	query = `SELECT * from formulas where id = $1`
	for _, fId := range formulasIDs {
		var formula model.Formula

		row := fh.db.QueryRow(query, fId)
		if err = row.Scan(
			&formula.Id,
			&formula.SectionId,
			&formula.Name,
			&formula.Description,
			&formula.Expression,
			pq.Array(&formula.Parameters),
			&formula.Difficulty,
		); err != nil {
			logger.Lg().Logf(0, "err: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		formulas = append(formulas, formula)
	}
	return c.JSON(http.StatusOK, formulas)
}
