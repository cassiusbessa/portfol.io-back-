package entities

import "time"

type Project struct {
	ID          string
	Name        string
	Description string
	Link        *string
	Image       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

func (p Project) validateName(name string) error {
	err := emptyStringValidator(name, "name")
	if err != nil {
		return err
	}
	err = invalidFieldLenghValidator(name, "name", 3, 32)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) validateDescription(description string) error {
	err := emptyStringValidator(description, "description")
	if err != nil {
		return err
	}
	err = invalidFieldLenghValidator(description, "description", 3, 255)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) validator() error {
	err := p.validateName(p.Name)
	if err != nil {
		return err
	}
	err = p.validateDescription(p.Description)
	if err != nil {
		return err
	}
	return nil
}

func NewProject(project Project) (*Project, error) {
	err := project.validator()
	if err != nil {
		return nil, err
	}
	return &Project{
		Name:        project.Name,
		Description: project.Description,
		Link:        project.Link,
		Image:       project.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
