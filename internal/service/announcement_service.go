package service

import (
	"github.com/jumystap/jumystap-core/internal/model"
	"github.com/jumystap/jumystap-core/internal/repository"
)

type AnnouncementService struct {
    repository *repository.AnnouncementRepository
}

func NewAnnouncementService (repository *repository.AnnouncementRepository) *AnnouncementService {
    return &AnnouncementService{repository: repository}
}

func (s *AnnouncementService) GetAllAnnouncements() ([]*model.Announcement, error) {
    announcements, err := s.repository.GetAllAnnouncements()
    if err != nil {
        return nil, err
    }

    return announcements, nil
}
