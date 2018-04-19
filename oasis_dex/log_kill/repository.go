package log_kill

import (
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type Datastore interface {
	Remove(model LogKillModel) error
}

type Repository struct {
	DB *postgres.DB
}

func (d Repository) Remove(model LogKillModel) error {
	_, err := d.DB.Exec(`DELETE FROM oasis.offer WHERE id = $1`, model.ID)
	if err != nil {
		return err
	}
	return nil
}
