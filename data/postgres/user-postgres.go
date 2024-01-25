package postgres

import (
	"database/sql"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	_, err := r.db.Exec(`
	INSERT INTO users (id, full_name, email, password, image, created_at, updated_at, delete_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`, user.ID, user.FullName, user.Email, user.Password, user.Image, user.CreatedAt, user.UpdatedAt, user.DeleteAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	var nullImage sql.NullString

	err := r.db.QueryRow(`
	SELECT id, full_name, email, password, image, created_at, updated_at, delete_at
	FROM users
	WHERE email = $1;
	`, email).Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &nullImage, &user.CreatedAt, &user.UpdatedAt, &user.DeleteAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user.Image = &nullImage.String
	return &user, nil
}
