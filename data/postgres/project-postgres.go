package postgres

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregates"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) usecases.ProjectRepository {
	return &ProjectRepository{db}
}

func (r *ProjectRepository) CreateProject(project *entities.Project, userId string, tagsId []int) error {
	var idValues string
	for i, id := range tagsId {
		if i == 0 {
			idValues = strconv.Itoa(id)
		} else {
			idValues += ", " + strconv.Itoa(id)
		}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec(`
		INSERT INTO projects (id, name, description, image, created_at, updated_at, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`, project.ID, project.Name, project.Description, project.Image, project.CreatedAt, project.UpdatedAt, userId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO project_tags (project_id, tag_id)
	SELECT $1, id
	FROM tags
	WHERE id IN (`+idValues+`);
`, project.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectRepository) FindAllProjects() ([]aggregates.Project, error) {
	var projects []aggregates.Project
	rows, err := r.db.Query(`
		SELECT 
			p.id, p.name, p.description, p.image, p.created_at, p.updated_at, p.delete_at, 
			u.id, u.full_name, u.email, u.image,
			ARRAY_AGG(t.name) AS tag_names
		FROM projects p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN project_tags pt ON p.id = pt.project_id
		LEFT JOIN tags t ON pt.tag_id = t.id
		WHERE p.delete_at IS NULL
		GROUP BY p.id, u.id
		ORDER BY p.created_at DESC;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fullProject aggregates.Project
		var nullUserImage, nullProjectImage sql.NullString
		var tagNames string
		var project entities.Project
		var user entities.User
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &nullProjectImage, &project.CreatedAt, &project.UpdatedAt, &project.DeleteAt,
			&user.ID, &user.FullName, &user.Email, &nullUserImage, &tagNames,
		)
		if err != nil {
			return nil, err
		}
		project.Image = &nullProjectImage.String
		user.Image = &nullUserImage.String
		tagNamesArray := strings.Split(tagNames[1:len(tagNames)-1], ", ")
		tagNamesArray = strings.Split(tagNamesArray[0], ",")
		fullProject = aggregates.NewProject(project, user, tagNamesArray)
		projects = append(projects, fullProject)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return projects, nil

}

func (r *ProjectRepository) FindProjectsByUserId(userId string) ([]aggregates.Project, error) {
	var projects []aggregates.Project
	rows, err := r.db.Query(`
		SELECT
			p.id, p.name, p.description, p.image, p.created_at, p.updated_at, p.delete_at,
			u.id, u.full_name, u.email, u.image,
			ARRAY_AGG(t.name) AS tag_names
		FROM projects p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN project_tags pt ON p.id = pt.project_id
		LEFT JOIN tags t ON pt.tag_id = t.id
		WHERE p.user_id = $1 AND p.delete_at IS NULL
		GROUP BY p.id, u.id
		ORDER BY p.created_at DESC;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nullUserImage, nullProjectImage sql.NullString
		var fullProject aggregates.Project
		var project entities.Project
		var user entities.User
		var tagNames string
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &nullProjectImage, &project.CreatedAt, &project.UpdatedAt, &project.DeleteAt,
			&user.ID, &user.FullName, &user.Email, &nullUserImage, &tagNames,
		)
		if err != nil {
			return nil, err
		}
		project.Image = &nullProjectImage.String
		user.Image = &nullUserImage.String
		tagNamesArray := strings.Split(tagNames[1:len(tagNames)-1], ", ")
		tagNamesArray = strings.Split(tagNamesArray[0], ",")
		fullProject = aggregates.NewProject(project, user, tagNamesArray)
		projects = append(projects, fullProject)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) FindProjectByNameAndUserId(name, userId string) (*entities.Project, error) {
	var project entities.Project
	var image sql.NullString

	err := r.db.QueryRow(`
		SELECT id, name, description, image, created_at, updated_at, delete_at
		FROM projects
		WHERE name = $1 AND user_id = $2 AND delete_at IS NULL;
	`, name, userId).Scan(
		&project.ID, &project.Name, &project.Description, &image, &project.CreatedAt, &project.UpdatedAt, &project.DeleteAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	project.Image = &image.String
	return &project, nil
}

func (r *ProjectRepository) UpdateProject(project *entities.Project, tagsId []int) (*aggregates.Project, error) {
	var nullUserImage, nullProjectImage sql.NullString
	var fullProject aggregates.Project
	var user entities.User

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.QueryRow(`
		UPDATE projects AS p
		SET name = $1, description = $2, image = $3, updated_at = $4
		FROM users AS u
		WHERE p.id = $5 AND p.user_id = u.id
		RETURNING p.id, p.name, p.description, p.image, p.created_at, p.updated_at, p.delete_at, u.full_name, u.email, u.image;
	`, project.Name, project.Description, project.Image, project.UpdatedAt, project.ID).Scan(
		&project.ID, &project.Name, &project.Description, &nullProjectImage, &project.CreatedAt, &project.UpdatedAt, &project.DeleteAt,
		&user.FullName, &user.Email, &nullUserImage,
	)
	if err != nil {
		return nil, err
	}
	project.Image = &nullProjectImage.String
	user.Image = &nullUserImage.String

	_, err = tx.Exec(`
		DELETE FROM project_tags
		WHERE project_id = $1;
	`, project.ID)
	if err != nil {
		return nil, err
	}

	if len(tagsId) > 0 {
		var idValues string
		for i, id := range tagsId {
			if i == 0 {
				idValues = strconv.Itoa(id)
			} else {
				idValues += ", " + strconv.Itoa(id)
			}
		}

		_, err = tx.Exec(`
			INSERT INTO project_tags (project_id, tag_id)
			SELECT $1, id
			FROM tags
			WHERE id IN (`+idValues+`);
		`, project.ID)
		if err != nil {
			return nil, err
		}
	}

	fullProject = aggregates.NewProject(*project, user, nil)
	return &fullProject, nil
}

func (r *ProjectRepository) DeleteProject(projectId, userId string) (bool, error) {
	_, err := r.db.Exec(`
		UPDATE projects
		SET delete_at = NOW()
		WHERE id = $1 AND user_id = $2;
	`, projectId, userId)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
