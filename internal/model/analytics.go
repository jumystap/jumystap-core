package model

type Analytics struct {
    UserCount                   uint64  `json:"user_count"`
    GraduateCount               uint64  `json:"graduate_count"`
    NoneGraduateCount           uint64  `json:"none_graduate_count"`
    CompaniesCount              uint64  `json:"companies_count"`
    AnnouncementsCount          uint64  `json:"announcements_count"`
    ResponsesCount              uint64  `json:"responses_count"`
    EmployeesRespondedCount     uint64  `json:"employees_responded_count"`
    CompaniesRespondedCount     uint64  `json:"companies_responded_count"`
}
