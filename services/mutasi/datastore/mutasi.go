package datastore

import "github.com/Diasisme/asssesment-march-ihsan.git/models"

func (f *DatabaseData) Mutasi(request models.Mutasi) error {
	return f.DB.Create(&request).Error
}
