package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllFoods() ([]*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			id, data_bank_id, food_name, food_description, coalesce(food_image, ''),
			created_at, updated_at
		from
			foods
		order by
			food_name
		`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	//Ensure rows are closed once they are no longer needed. apply defer
	defer rows.Close()

	var foods []*models.Food

	for rows.Next() {
		var food models.Food
		err := rows.Scan(
			&food.ID,
			&food.DataBankID,
			&food.FoodName,
			&food.FoodDescription,
			&food.FoodImage,
			&food.CreatedAT,
			&food.UpdatedAT,
		)
		if err != nil {
			return nil, err
		}
		foods = append(foods, &food)
	}

	return foods, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
				created_at, updated_at 
				from users where email = $1`


	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
				created_at, updated_at 
				from users where id = $1`


	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}