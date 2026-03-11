package user

import (
	"database/sql"
)

type User struct {
	email          string
	hashedPassword []byte
}

func (s *Server) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT email, password FROM employees WHERE email = $1
		UNION ALL
		SELECT email, password FROM clients WHERE email = $1
		LIMIT 1
	`

	var user User

	err := s.database.QueryRow(query, email).Scan(&user.email, &user.hashedPassword)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateRefreshToken(email string) {
}
