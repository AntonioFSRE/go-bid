package repository

var (
	getUserQuery  = `SELECT * FROM users WHERE id = $1`

	findUserByNameQuery = `SELECT * FROM users WHERE name = $1`

)
