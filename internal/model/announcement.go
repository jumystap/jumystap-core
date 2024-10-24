package model

type Announcement struct {
    Id          string      `json:"id"`
    Title       string      `json:"title"`
    Description string      `json:"description"`
    Cost        interface{} `json:"cost"`
    CostMin     interface{} `json:"cost_min"`
    CostMax     interface{} `json:"cost_max"`
    SalaryType  string      `json:"salary_type"`
    Experience  string      `json:"experience"`
    Education   string      `json:"education"`
    WorkTime    string      `json:"work_time"`
    WorkHours   string      `json:"work_hours"`
    City        string      `json:"city"`
}
