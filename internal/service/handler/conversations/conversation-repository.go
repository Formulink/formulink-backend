package conversations

import (
	"database/sql"
	"formulink-backend/internal/dto"
	"formulink-backend/internal/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ConversationRepository struct {
	db *sql.DB
}

func NewConversationRepository(db *sql.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (cr *ConversationRepository) createNewConversation(userId uuid.UUID) (uuid.UUID, error) {
	query := `INSERT into conversations (id, user_id) VALUES ($1, $2)`
	id := uuid.New()
	if _, err := cr.db.Exec(query, id, userId); err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func (cr *ConversationRepository) getConversation(conversationId uuid.UUID) (model.Conversation, error) {
	var conversation model.Conversation

	query := `SELECT * from conversations where id = $1`
	row := cr.db.QueryRow(query, conversationId)
	if err := row.Scan(
		&conversation.Id,
		&conversation.UserId,
		pq.Array(&conversation.Messages),
		&conversation.CreatedAt,
	); err != nil {
		return model.Conversation{}, err
	}
	return conversation, nil
}

func (cr *ConversationRepository) getAllConversations(userId uuid.UUID) ([]model.Conversation, error) {
	var conversations []model.Conversation

	query := `SELECT * from conversations where user_id = $1`
	rows, err := cr.db.Query(query, userId)
	if err != nil {
		return []model.Conversation{}, err
	}

	for rows.Next() {
		var conversation model.Conversation
		if err = rows.Scan(
			&conversation.Id,
			&conversation.UserId,
			pq.Array(&conversation.Messages),
			&conversation.CreatedAt,
		); err != nil {
			return []model.Conversation{}, err
		}
		conversations = append(conversations, conversation)
	}
	return conversations, nil
}

func (cr *ConversationRepository) AddMessage(req dto.NewMessageDto) error {
	query := `
		UPDATE conversations
		SET messages = array_append(messages, $3)
		WHERE user_id = $1 and id = $2
	`
	if _, err := cr.db.Exec(query, req.UserId, req.ConversationId, req.Message); err != nil {
		return err
	}
	return nil
}

func (cr *ConversationRepository) deleteConversation(conversationId uuid.UUID) error {
	query := `DELETE from conversations where id = $1`
	if _, err := cr.db.Exec(query, conversationId); err != nil {
		return err
	}
	return nil
}
