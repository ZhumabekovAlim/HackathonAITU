package dbs

import (
	"awesomeProject/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
)

type ClientSigthModel struct {
	DB *sql.DB
}

func (m *ClientSigthModel) Insert(sight *models.Client_Sights) error {
	stmt := `
        INSERT INTO astana.client_sight
        (client_id, sight_id) 
        VALUES (?, ?);`

	_, err := m.DB.Exec(stmt, sight)
	if err != nil {
		return err
	}

	return nil
}

func (m *ClientSigthModel) GetClientSightByClientId(client_id string) ([]byte, error) {
	stmt := `SELECT * FROM astana.client_sight WHERE client_sight.client_id = ?`

	clientSightRow := m.DB.QueryRow(stmt, client_id)

	s := &models.Client_Sights{}

	err := clientSightRow.Scan(&s.Id, &s.ClientId, &s.SightId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedClientSight, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return convertedClientSight, nil
}
