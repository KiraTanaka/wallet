SELECT id,
	user_id,
	balance
FROM wallets
WHERE id = $1