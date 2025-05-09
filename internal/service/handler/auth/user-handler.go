package handler

import (
	"database/sql"
	"errors"
	"formulink-backend/internal/dto"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetByTelegramID(telegramID int) (*model.User, error) {
	var user model.User
	row := s.db.QueryRow("SELECT * FROM users WHERE telegramid = $1", telegramID)
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(req dto.CreateUserRequest) (*model.User, error) {
	id := uuid.New()
	_, err := s.db.Exec(`INSERT INTO users (id, telegramid, username) VALUES ($1, $2, $3)`,
		id, req.TelegramId, req.Username)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         id,
		TelegramID: req.TelegramId,
		Username:   req.Username,
	}, nil
}

func (s *UserService) SetNeedOnboardingFalse(tgId int) error {
	query := `UPDATE users 
			SET need_onboarding = true
			WHERE telegramid = $1
		`
	if _, err := s.db.Exec(query, tgId); err != nil {
		logger.Lg().Logf(0, "can't update need_onboarding to false | err: %v", err)
		return err
	}
	return nil
}
