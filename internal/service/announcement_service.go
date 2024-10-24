package service

import (
	"strconv"

	"github.com/jumystap/jumystap-core/internal/model"
	"github.com/jumystap/jumystap-core/internal/repository"
)

type AnnouncementService struct {
    repository *repository.AnnouncementRepository
}

func NewAnnouncementService (repository *repository.AnnouncementRepository) *AnnouncementService {
    return &AnnouncementService{repository: repository}
}

func (s *AnnouncementService) GetAllAnnouncements(page string) ([]*model.Announcement, error) {
    var offset int
    if page == "" {
        offset = 0
    }else {
        offset, _ = strconv.Atoi(page)
        offset = (offset - 1)*10
    }
    announcements, err := s.repository.GetAllAnnouncements(offset)
    if err != nil {
        return nil, err
    }

    return announcements, nil
}
