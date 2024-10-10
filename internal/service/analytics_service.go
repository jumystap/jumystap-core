package service

import (
	"time"

	"github.com/jumystap/jumystap-core/internal/model"
	"github.com/jumystap/jumystap-core/internal/repository"
)

type AnalyticsService struct {
    repo *repository.AnalyticsRepository
}

func NewAnalyticsService(repo *repository.AnalyticsRepository) *AnalyticsService {
    return &AnalyticsService{repo: repo}
}

func (s *AnalyticsService) GetAnalytics(startDate string, endDate string) (*model.Analytics, error) {
    if (startDate == "" ) {
        startDate = "2024-08-01"
    }

    if (endDate == "" ) {
        endDate = time.Now().Format("2006-01-02")
    }

    analytics := &model.Analytics{}

    userCount, err := s.repo.GetCountOfUsers(startDate, endDate)
    if err != nil {
        return analytics, err
    }

    graduateCount, err := s.repo.GetCountOfGraduates(startDate, endDate)
    if err != nil {
        return nil, err
    }

    noneGraduateCount, err := s.repo.GetCountOfNoneGraduates(startDate, endDate)
    if err != nil {
        return nil, err
    }

    companiesCount, err := s.repo.GetCountOfCompanies(startDate, endDate)
    if err != nil {
        return nil, err
    }

    announcementsCount, err := s.repo.GetCountOfAnnouncements(startDate, endDate)
    if err != nil {
        return nil, err
    }

    responsesCount, err := s.repo.GetCountOfResponses(startDate, endDate)
    if err != nil {
        return nil, err
    }

    employeesRespondedCount, err := s.repo.GetCountOfEmployeesResponded(startDate, endDate)
    if err != nil {
        return nil, err
    }

    companiesRespondedCount, err := s.repo.GetCountOfCompaniesResponded(startDate, endDate)
    if err != nil {
        return nil, err
    }

    analytics.UserCount = userCount
    analytics.GraduateCount = graduateCount
    analytics.NoneGraduateCount = noneGraduateCount
    analytics.CompaniesCount = companiesCount
    analytics.AnnouncementsCount = announcementsCount
    analytics.EmployeesRespondedCount = employeesRespondedCount
    analytics.CompaniesRespondedCount = companiesRespondedCount
    analytics.ResponsesCount = responsesCount

    return analytics, nil
}
