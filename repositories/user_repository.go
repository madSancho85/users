package repositories

import (
	"database/sql"
	"time"

	"users/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (id, first_name, last_name, age, recording_date) VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.FirstName, user.LastName, user.Age, user.RecordingDate)
	return err
}

func (r *UserRepository) GetUsers() ([]*models.User, error) {
	rows, err := r.db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.RecordingDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUsersByDateAndAge(dateFrom, dateTo time.Time, ageFrom, ageTo int) ([]*models.User, error) {
	rows, err := r.db.Query(`SELECT * FROM users WHERE recording_date >= $1 AND recording_date <= $2 AND age >= $3 AND age <= $4`,
		dateFrom, dateTo, ageFrom, ageTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.RecordingDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
