package postgres

import (
	"database/sql"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
)

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) usecases.TagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) FindAllTags() ([]entities.Tag, error) {
	var tags []entities.Tag

	rows, err := r.db.Query(`
	SELECT id, name
	FROM tags;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag entities.Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
