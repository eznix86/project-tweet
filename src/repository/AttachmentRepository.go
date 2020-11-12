package repository

import (
	"errors"

	"github.com/theArtechnology/project-tweet/src/entity"
)

// AttachmentRepository is a Repository (meaning Storage) for all my attachments
type AttachmentRepository struct {
	AttachementList []entity.Attachment // in memory AttachementList
}

// Find is
func (r *AttachmentRepository) Find(id int) (*entity.Attachment, error) {
	for _, attachment := range r.AttachementList {
		if attachment.ID == id {
			return &attachment, nil
		}
	}
	return nil, errors.New("attachment not found")
}

// Add is
func (r *AttachmentRepository) Add(entity entity.Attachment) error {
	r.AttachementList = append(r.AttachementList, entity)
	return nil
}

// Remove is
func (r *AttachmentRepository) Remove(id int) error {
	for index, attachment := range r.AttachementList {
		if attachment.ID == id {
			r.AttachementList = append(r.AttachementList[:index], r.AttachementList[index:]...)
			return nil
		}
	}

	return errors.New("attachment not found")
}

func (r *AttachmentRepository) getAll(entity entity.Attachment) ([]entity.Attachment, error) {
	return r.AttachementList, nil
}
