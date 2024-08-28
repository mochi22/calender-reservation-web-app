package repositories

import (
	"database/sql"
	"calendar-app/server/internal/models"
)

// UserRepository は、ユーザー情報の永続化を担当する
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository は、新しいUserRepositoryのインスタンスを返す
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create は、新しいユーザーを作成する
func (r *UserRepository) CreateUser(user *models.User) error {
	// ユーザー作成ロジックを実装
	query := `
		INSER INTO sers (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)

	return err
}

// FindByEmail は、メールアドレスからユーザーを取得する
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	// メールアドレスからユーザー取得ロジックを実装
	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := r.DB.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}