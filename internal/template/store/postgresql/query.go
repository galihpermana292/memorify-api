package postgresql

const (
	queryCreateTemplate = `
		INSERT INTO
			template
			(
				id,
				name,
				label,
				thumbnail_uri,
				create_time
			)
		VALUES
			(
				:id,
				:name,
				:label,
				:thumbnail_uri,
				:create_time
			)
		RETURNING
			id
	`

	queryGetTemplate = `
		SELECT
			t.id,
			t.name,
			t.label,
			t.thumbnail_uri,
			t.create_time,
			t.update_time
		FROM
			template t
		%s
	`

	queryUpdateTemplate = `
		UPDATE
			template
		SET
			name = :name,
			label = :label,
			thumbnail_uri = :thumbnail_uri,
			update_time = :update_time
		WHERE
			id = :id 
	`

	queryDeleteTemplate = `
		DELETE FROM
			template
		WHERE
			id = :id
	`
)
