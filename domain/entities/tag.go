package entities

type Tag struct {
	ID   string
	Name string
}

func (t Tag) ValidateName(name string) error {
	err := emptyStringValidator(name, "name")
	if err != nil {
		return err
	}
	err = invalidFieldLenghValidator(name, "name", 2, 32)
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) validator() error {
	err := t.ValidateName(t.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewTag(tag Tag) (*Tag, error) {
	err := tag.validator()
	if err != nil {
		return nil, err
	}
	return &Tag{
		Name: tag.Name,
	}, nil
}
