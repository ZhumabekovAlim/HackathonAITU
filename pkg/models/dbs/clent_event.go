package dbs

import (
	"awesomeProject/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
)

type ClientEventModel struct {
	DB *sql.DB
}

func (m *ClientEventModel) Insert(event *models.Client_Events) error {
	stmt := `
        INSERT INTO astana.client_event
        (client_id, event_id) 
        VALUES (?, ?);`

	_, err := m.DB.Exec(stmt, event.ClientId, event.EventId)
	if err != nil {
		return err
	}

	return nil
}

func (m *ClientEventModel) GetClientEventByClientId(client_id string) ([]byte, error) {
	stmt := `SELECT * FROM astana.client_event WHERE client_id = ?`

	clientEventRow := m.DB.QueryRow(stmt, client_id)

	s := &models.Client_Events{}

	err := clientEventRow.Scan(&s.Id, &s.ClientId, &s.EventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedClientEvent, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return convertedClientEvent, nil
}
