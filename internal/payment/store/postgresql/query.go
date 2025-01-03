package postgresql

const (
	queryCreatePayment = `
		INSERT INTO
			payment
			(
				id,
				user_id,
				proof_payment_url,
				date,
				status,
				create_time
			)
		VALUES
			(
				:id,
				:user_id,
				:proof_payment_url,
				:date,
				:status,
				:create_time
			)
		RETURNING
			id
	`

	queryGetPayment = `
		SELECT
			p.id,
			p.user_id,
			u.fullname as user_name,
			u.type as user_type,
			u.quota as user_quota,
			p.proof_payment_url,
			p.date,
			p.status,
			p.create_time,
			p.update_time
		FROM
			payment p
		LEFT JOIN
			user_info u
		ON
			u.id = p.user_id
		%s
	`

	queryUpdatePayment = `
		UPDATE
			payment
		SET
			user_id = :user_id,
			proof_payment_url = :proof_payment_url,
			date = :date,
			status = :status,
			update_time = :update_time
		WHERE
			id = :id
	`
)
