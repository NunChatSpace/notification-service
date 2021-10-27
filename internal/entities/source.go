package entities

import "gorm.io/gorm"

func (d Db) Source() SourceInterface {
	return sourceDb(d)
}

type sourceDb struct {
	gorm *gorm.DB
}

type SourceInterface interface {
	Add(info SourceModel) (SourceModel, error)
	Get(uid string) (SourceModel, error)
	GetByIds(uids []string) ([]SourceModel, error)
}

type SourceModel struct {
	Model
	Name           string `json:"name"`
	ServerAddresss string `json:"server_address"`
	DatabaseName   string `json:"database_name"`
	TblName        string `json:"table_name"`
	DataKey        string `json:"data_key"`
}

func (SourceModel) TableName() string {
	return "source"
}

func (sm sourceDb) Add(info SourceModel) (SourceModel, error) {
	tx := sm.gorm.Create(&info)

	return info, tx.Error
}

func (sm sourceDb) Get(uid string) (SourceModel, error) {
	source := SourceModel{}
	tx := sm.gorm.First(&source, "id = ?", uid)

	return source, tx.Error
}

func (sm sourceDb) GetByIds(uids []string) ([]SourceModel, error) {
	source := []SourceModel{}
	tx := sm.gorm.Where("owner_id IN ?", uids).Find(&source)

	return source, tx.Error
}
