package storage

import (
	"1.Redis/internal/models"
	"time"
)

func FetchDataFromDb(id string) (*models.Data, error) {
	time.Sleep(1 * time.Second)

	data := &models.Data{
		ID:   id,
		Info: "some infor",
	}

	return data, nil
}
