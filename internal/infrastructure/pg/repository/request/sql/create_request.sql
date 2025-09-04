INSERT INTO requests(user_id,
                     type,
                     payload,
                     status,
                     created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING ID;