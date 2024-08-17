package service

import (
	"context"

	"gorm.io/gorm"
)

func (s *service) Delete(ctx context.Context, id uint) error {
	rowsAffected, err := s.taskRepo.Delete(ctx, id)
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
