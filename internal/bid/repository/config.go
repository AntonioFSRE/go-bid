package repository

var (
	getBidQuery    = `SELECT ttl, price, setAt FROM bid WHERE id = $1`
	createBidQuery = `INSERT INTO bid (id, author_id, ttl, price, setAt) 
									VALUES ($1, $2, $3, $4, now()) RETURNING *`
	updateBidQuery = `UPDATE bid 
									SET price = COALESCE(NULLIF($1, ''), price), 
										user_id = COALESCE(NULLIF($2, ''), user_id), 
									WHERE id = $3 IF PRICE < $1`

)
