package datastore

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/sirupsen/logrus"
)

func (f *DatabaseData) MigrationDB() (err error) {

	err = f.DB.AutoMigrate(&models.Mutasi{})
	if err != nil {
		remark := "Proses migrasi tabel DB gagal dilakukan"
		f.log.Error(logrus.Fields{
			"error": err,
		}, nil, remark)
		return err
	}
	return
}
