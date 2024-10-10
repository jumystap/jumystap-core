package repository

import "database/sql"

type AnalyticsRepository struct {
    db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) *AnalyticsRepository {
    return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) GetCountOfUsers (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE DATE(created_at) >= ? AND DATE(created_at) <= ?"

    var count uint64

    err := r.db.QueryRow(query, startDate, endDate).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfGraduates (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id=2 AND is_graduate = 1"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfNoneGraduates (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id=2 AND is_graduate = 0"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfCompanies (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id !=2"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfAnnouncements (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM announcements"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfResponses (startDate string, endDate string) (uint64, error) {
    const query = "SELECT COUNT(*) FROM responses"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfEmployeesResponded (startDate string, endDate string) (uint64, error) {
    const query = ` SELECT COUNT(*) 
                    FROM (SELECT employee_id FROM responses GROUP BY employee_id) 
                    AS grouped_responses
                    `

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfCompaniesResponded (startDate string, endDate string) (uint64, error) {
    const query = ` SELECT COUNT(*) 
                    FROM (SELECT announcement_id FROM responses GROUP BY announcement_id) 
                    AS grouped_responses
                    `

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}
