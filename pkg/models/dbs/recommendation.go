package dbs

import (
	"awesomeProject/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
)

type Recommendations struct {
	DB *sql.DB
}

func (m *Recommendations) Insert(recommendation *models.Recommendation) error {
	stmt := `
        INSERT INTO astana.recommendation
        (cliend_id, sight_category_id, event_category_id) 
        VALUES (?, ?, ?);`

	_, err := m.DB.Exec(stmt, recommendation)
	if err != nil {
		return err
	}

	return nil
}

func (m *Recommendations) GetEventById(id string) ([]byte, error) {
	stmt := `SELECT id, cliend_id, sight_category_id, event_category_id FROM astana.recommendation WHERE id = ?`

	eventRow := m.DB.QueryRow(stmt, id)

	s := &models.Recommendation{}

	err := eventRow.Scan(&s.Id, &s.ClientId, &s.SightCategoryId, &s.EventCategoryId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedRec, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return convertedRec, nil
}

func (m *Recommendations) GetAllEvents() ([]byte, error) {
	stmt := `SELECT id, cliend_id, sight_category_id, event_category_id FROM astana.recommendation`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*models.Recommendation{}

	for rows.Next() {
		s := &models.Recommendation{}
		err = rows.Scan(&s.Id, &s.ClientId, &s.SightCategoryId, &s.EventCategoryId)
		if err != nil {
			return nil, err
		}
		events = append(events, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedEvents, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}
	return convertedEvents, nil
}

func (m *Recommendations) DeleteEventById(id int) error {
	stmt := `DELETE FROM astana.recommendation WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
