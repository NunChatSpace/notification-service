package entities

import "gorm.io/gorm"

func (d Db) Destination() DestinationInterface {
	return destinationDb(d)
}

type destinationDb struct {
	gorm *gorm.DB
}

type DestinationInterface interface {
	Add(info DestinationModel) (DestinationModel, error)
	Get(uid string) (DestinationModel, error)
	GetByIds(uids []string) ([]DestinationModel, error)
}

type DestinationModel struct {
	Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func (DestinationModel) TableName() string {
	return "destination"
}

func (dm destinationDb) Add(info DestinationModel) (DestinationModel, error) {
	tx := dm.gorm.Create(&info)

	return info, tx.Error
}

func (dm destinationDb) Get(uid string) (DestinationModel, error) {
	dest := DestinationModel{}
	dm.gorm.First(&dest, "id = ?", uid)

	return dest, nil
}

func (dm destinationDb) GetByIds(uids []string) ([]DestinationModel, error) {
	dests := []DestinationModel{}
	tx := dm.gorm.Where("owner_id IN ?", uids).Find(&dests)

	return dests, tx.Error
}
