package user_stats

import (
	"database/sql"
	"formulink-backend/internal/dto"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
	"time"
)

type UserStatsRepository struct {
	db *sql.DB
}

func NewUserStatsRepository(db *sql.DB) *UserStatsRepository {
	return &UserStatsRepository{db: db}
}

func (usr *UserStatsRepository) createNewRecord(req dto.NewRecordRequest) (dto.NewRecordResponse, error) {
	id := uuid.New()
	query := `INSERT into user_stats (id, user_id, "right", fail, completed_at) values ($1, $2, $3, $4, $5)`
	if _, err := usr.db.Exec(query, id, req.UserId, req.Right, req.Fail, time.Now()); err != nil {
		logger.Lg().Logf(0, "can't insert data | err: %v", err)
		return dto.NewRecordResponse{}, err
	}
	return dto.NewRecordResponse{
		Id:          id,
		CompletedAt: time.Now(),
	}, nil

}

func (usr *UserStatsRepository) getUserStats(userId uuid.UUID) *[]model.UserStats {
	var stats []model.UserStats

	query := `SELECT * from user_stats WHERE user_id = $1`
	rows, err := usr.db.Query(query, userId)
	if err != nil {
		logger.Lg().Logf(0, "ISE | err: %v", err)
		return nil
	}

	for rows.Next() {
		var s model.UserStats
		if err = rows.Scan(
			&s.Id,
			&s.UserId,
			&s.Right,
			&s.Fail,
			&s.CompletedAt,
		); err == nil {
			stats = append(stats, s)
		} else {
			logger.Lg().Logf(0, "ISE | err: %v", err)
			return nil
		}
	}
	return &stats
}
