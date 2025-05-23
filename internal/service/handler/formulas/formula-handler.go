package formulas

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"formulink-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"time"
)

type FormulaHandler struct {
	db    *sql.DB
	redis *redis.Client
}

func NewFormulaHandler(_db *sql.DB, _redis *redis.Client) *FormulaHandler {
	return &FormulaHandler{
		db:    _db,
		redis: _redis,
	}
}

//func (fh *FormulaHandler) GetAllFormulas(c echo.Context) error {
//
//	ctx := context.TODO()
//	v, err := utils.Parse(fh.redis.Get(ctx, "fall").Bytes())
//	if err != nil {
//
//		var formulas []model.Formula
//		query := `SELECT * from formulas`
//		rows, err := fh.db.Query(query)
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, err)
//		}
//
//		for rows.Next() {
//			var formula model.Formula
//			err = rows.Scan(
//				&formula.Id,
//				&formula.SectionId,
//				&formula.Name,
//				&formula.Description,
//				&formula.Expression,
//				pq.Array(&formula.Parameters),
//				&formula.Difficulty,
//				&formula.VideoLink,
//				&formula.VideoName,
//			)
//			if err != nil {
//				logger.Lg().Logf(0, "can't parse data form db | err: %v", err)
//				return c.JSON(http.StatusInternalServerError, err)
//			}
//			formulas = append(formulas, formula)
//		}
//
//		go func() {
//			var bytes []byte
//			if bytes, err = json.Marshal(formulas); err != nil {
//				return
//			}
//			if err = fh.redis.Set(ctx, "fall", bytes, time.Hour).Err(); err != nil {
//			}
//		}()
//
//		return c.JSON(http.StatusOK, formulas)
//
//	}
//
//	return c.JSON(http.StatusOK, v)
//}

func (fh *FormulaHandler) GetAllFormulas(c echo.Context) error {

	var formulas []model.Formula

	query := `SELECT * FROM formulas`
	rows, err := fh.db.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var formula model.Formula
		err = rows.Scan(
			&formula.Id,
			&formula.SectionId,
			&formula.Name,
			&formula.Description,
			&formula.Expression,
			pq.Array(&formula.Parameters),
			&formula.Difficulty,
			&formula.VideoLink,
			&formula.VideoName,
		)
		if err != nil {
			logger.Lg().Logf(0, "can't parse data form db | err: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		formulas = append(formulas, formula)
	}

	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, formulas)
}

func (fh *FormulaHandler) GetFormulasBySectionId(c echo.Context) error {
	var formulas []model.Formula
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is missing")
	}

	query := `SELECT * from formulas WHERE section_id = $1`
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
			&formula.VideoLink,
			&formula.VideoName,
		); err != nil {
			logger.Lg().Logf(0, "can't parse data form db | err: %v", err)
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

	query := `SELECT * FROM formulas WHERE id = $1`
	row := fh.db.QueryRow(query, id)
	err = row.Scan(
		&formula.Id,
		&formula.SectionId,
		&formula.Name,
		&formula.Description,
		&formula.Expression,
		pq.Array(&formula.Parameters),
		&formula.Difficulty,
		&formula.VideoLink,
		&formula.VideoName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		logger.Lg().Logf(0, "can't parse data form db | err: %v", err)
		return c.JSON(http.StatusNotFound, "formula doesn't exist")
	}

	return c.JSON(http.StatusOK, formula)
}

func (fh *FormulaHandler) GetFormulaOfTheDay(c echo.Context) error {
	ctx := context.TODO()
	v, err := utils.Parse(fh.redis.Get(ctx, "fday").Bytes())
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = fh.setRandomFormula()
			if err != nil {
				logger.Lg().Logf(zapcore.InfoLevel, "error: %v", err)
				return c.JSON(http.StatusInternalServerError, ":(")
			}

			v, err = utils.Parse(fh.redis.Get(ctx, "fday").Bytes())
			if err != nil {
				logger.Lg().Logf(zapcore.InfoLevel, "error: %v", err)
				return c.JSON(http.StatusInternalServerError, "item hasn't been updated in redis")
			}
			return c.JSON(http.StatusOK, v)
		} else {
			logger.Lg().Logf(zapcore.InfoLevel, "error: %v", err)
			return c.JSON(http.StatusInternalServerError, ":(")
		}
	}

	logger.Lg().Logf(zapcore.InfoLevel, "item exists in redis %v", v)
	return c.JSON(http.StatusOK, v)
}

func (fh *FormulaHandler) setRandomFormula() error {
	var formula model.Formula

	query := `SELECT *
          FROM formulas 
          ORDER BY RANDOM() 
          LIMIT 1`

	row := fh.db.QueryRow(query)
	err := row.Scan(
		&formula.Id,
		&formula.SectionId,
		&formula.Name,
		&formula.Description,
		&formula.Expression,
		pq.Array(&formula.Parameters),
		&formula.Difficulty,
		&formula.VideoLink,
		&formula.VideoName,
	)
	if err != nil {
		logger.Lg().Logf(zapcore.InfoLevel, "problem is here (get rows)")
		return err
	}

	fjson, err := json.Marshal(formula)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	err = fh.redis.Set(ctx, "fday", fjson, time.Hour*24).Err()
	if err != nil {
		logger.Lg().Logf(zapcore.InfoLevel, "problem is here (set value) %v", err)
		return err
	}
	logger.Lg().Logf(zapcore.InfoLevel, "item succesfully updated in redis")
	return nil
}
