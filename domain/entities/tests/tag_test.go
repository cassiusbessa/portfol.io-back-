package entities_test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

func TestNewTag(t *testing.T) {
	tag := entities.Tag{
		Name: "Valid Tag",
	}
	testCases := []struct {
		description string
		tag         entities.Tag
		wantErr     bool
	}{
		{
			description: "should return nil when tag is valid",
			tag:         tag,
			wantErr:     false,
		},
		{
			description: "should return error when tag name is empty",
			tag: entities.Tag{
				Name: "",
			},
			wantErr: true,
		},
		{
			description: "should return error when tag name is too short",
			tag: entities.Tag{
				Name: "a",
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			_, err := entities.NewTag(tC.tag)
			if err != nil && !tC.wantErr {
				t.Errorf("expected nil, got %s", err)
			}
			if err == nil && tC.wantErr {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
