package handler

import (
	"database/sql"
	"errors"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"net/http"
	"strconv"
)

type FormulaHandler struct {
	db *sql.DB
}

func NewFormulaHandler(_db *sql.DB) *FormulaHandler {
	return &FormulaHandler{
		db: _db,
	}
}

func (fh *FormulaHandler) GetFormulasBySectionId(c echo.Context) error {
	var formulas []model.Formula
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Lg().Fatalf("can't parse id | err: id is missing")
		return c.JSON(http.StatusBadRequest, "id is missing")
	}

	query := `SELECT id, section_id, name,  description, expression, parameters, difficulty from formulas WHERE section_id = $1`
	rows, err := fh.db.Query(query, id)
	if err != nil {
		return c.JSON(http.StatusOK, "there no formulas by this section ID")
	}

	for rows.Next() {
		var formula model.Formula
		if err := rows.Scan(
			&formula.Id,
			&formula.SectionId,
			&formula.Name,
			&formula.Description,
			&formula.Expression,
			pq.Array(&formula.Parameters),
			&formula.Difficulty,
		); err != nil {
			logger.Lg().Fatalf("can't parse db data to model.Formula | err: %v", err)
			return c.JSON(http.StatusInternalServerError, "can't parse db data to model.Formula")
		}
		formulas = append(formulas, formula)
	}

	return c.JSON(http.StatusOK, formulas)
}

func (fh *FormulaHandler) GetFormulaById(c echo.Context) error {
	var formula model.Formula
	idStr := c.Param("id")
	if idStr == "" {
		logger.Lg().Debug("can't parse id | err: id is missing or wrong type")
		return c.JSON(http.StatusBadRequest, "id is missing | err: id is missing or can't or wrong type")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Lg().Debug("can't parse id | err: %v", err)
		return c.JSON(http.StatusBadRequest, "id is invalid format")
	}

	query := `SELECT id, section_id, name, description, expression, parameters, difficulty FROM formulas WHERE id = $1`
	row := fh.db.QueryRow(query, id)
	err = row.Scan(
		&formula.Id,
		&formula.SectionId,
		&formula.Name,
		&formula.Description,
		&formula.Expression,
		pq.Array(&formula.Parameters),
		&formula.Difficulty,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusNotFound, "formula doesn't exist")
	}

	return c.JSON(http.StatusOK, formula)
}
