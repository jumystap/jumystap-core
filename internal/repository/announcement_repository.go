package repository

import (
	"database/sql"
	"fmt"

	"github.com/jumystap/jumystap-core/internal/model"
)

type AnnouncementRepository struct {
    db *sql.DB
}

func NewAnnouncementRepository (db *sql.DB) *AnnouncementRepository {
    return &AnnouncementRepository{db: db}
}

func (r *AnnouncementRepository) GetAllAnnouncements(offset int) ([]*model.Announcement, error) {
    announcements := []*model.Announcement{}

    rows, err := r.db.Query(`
            SELECT id, title, description, cost, cost_min, cost_max, city, work_time, work_hours, salary_type, education, experience
			FROM announcements
            LIMIT 10 OFFSET ?
    `, offset)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	defer rows.Close()
	for rows.Next() {
		user, err := scanRowIntoAnnouncement(rows)
		if err != nil {
			return nil, fmt.Errorf("%s", err)
		}
		announcements = append(announcements, user)
	}

    return announcements, nil
}

func scanRowIntoAnnouncement(rows *sql.Rows) (*model.Announcement, error) {
	announcement := new(model.Announcement)

	err := rows.Scan(
		&announcement.Id,
		&announcement.Title,
		&announcement.Description,
		&announcement.Cost,
		&announcement.CostMin,
		&announcement.CostMax,
		&announcement.City,
		&announcement.WorkTime,
		&announcement.WorkHours,
		&announcement.SalaryType,
		&announcement.Education,
		&announcement.Experience,
	)
	if err != nil {
		return nil, err
	}

	return announcement, err
}
