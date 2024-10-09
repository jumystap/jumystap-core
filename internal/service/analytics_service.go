package service

import (
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
    analytics := &model.Analytics{}

    userCount, err := s.repo.GetCountOfUsers()
    if err != nil {
        return analytics, err
    }

    graduateCount, err := s.repo.GetCountOfGraduates()
    if err != nil {
        return nil, err
    }

    noneGraduateCount, err := s.repo.GetCountOfNoneGraduates()
    if err != nil {
        return nil, err
    }

    companiesCount, err := s.repo.GetCountOfCompanies()
    if err != nil {
        return nil, err
    }

    announcementsCount, err := s.repo.GetCountOfAnnouncements()
    if err != nil {
        return nil, err
    }

    responsesCount, err := s.repo.GetCountOfResponses()
    if err != nil {
        return nil, err
    }

    employeesRespondedCount, err := s.repo.GetCountOfEmployeesResponded()
    if err != nil {
        return nil, err
    }

    companiesRespondedCount, err := s.repo.GetCountOfCompaniesResponded()
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
