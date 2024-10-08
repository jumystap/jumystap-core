package repository

import "database/sql"

type AnalyticsRepository struct {
    db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) *AnalyticsRepository {
    return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) GetCountOfUsers () (uint64, error) {
    const query = "SELECT COUNT(*) FROM users"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfGraduates () (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id=2 AND is_graduate = 1"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfNoneGraduates () (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id=2 AND is_graduate = 0"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfCompanies () (uint64, error) {
    const query = "SELECT COUNT(*) FROM users WHERE role_id !=2"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfAnnouncements () (uint64, error) {
    const query = "SELECT COUNT(*) FROM announcements"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfResponses () (uint64, error) {
    const query = "SELECT COUNT(*) FROM responses"

    var count uint64

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *AnalyticsRepository) GetCountOfEmployeesResponded () (uint64, error) {
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

func (r *AnalyticsRepository) GetCountOfCompaniesResponded () (uint64, error) {
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
