package model

import "database/sql"

type Announcement struct {
    Id          string          `json:"id"`
    Title       string          `json:"title"`
    Description sql.NullString  `json:"description"`
    Cost        interface{}     `json:"cost"`
    CostMin     interface{}     `json:"cost_min"`
    CostMax     interface{}     `json:"cost_max"`
    SalaryType  string          `json:"salary_type"`
    Experience  string          `json:"experience"`
    Education   string          `json:"education"`
    WorkTime    sql.NullString  `json:"work_time"`
    WorkHours   sql.NullString  `json:"work_hours"`
    City        string          `json:"city"`
}
