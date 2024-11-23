INSERT INTO wallet_operations
			(wallet_id,
			operation_type,
			amount)
VALUES     ($1,
			$2,
			$3)
RETURNING id