package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type TagUseCase struct {
	tagRepository TagRepository
}

func NewTagUseCase(tagRepository TagRepository) TagUseCase {
	return TagUseCase{
		tagRepository: tagRepository,
	}
}

func (t TagUseCase) FindAllTags() ([]entities.Tag, error) {
	tags, err := t.tagRepository.FindAllTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
