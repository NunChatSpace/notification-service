package entities

import "gorm.io/gorm"

func (d Db) Expression() ExpressionInterface {
	return expressionDb(d)
}

type expressionDb struct {
	gorm *gorm.DB
}

type ExpressionInterface interface {
	Add(info ExpressionModel) (ExpressionModel, error)
	Get(uid string) (ExpressionModel, error)
	GetAll(owner_uid string) ([]ExpressionModel, error)
	SetIsIncondition(flag bool, id string) error
}

type ExpressionModel struct {
	Model
	Name                string `json:"name"`
	SourceID            string `json:"source_id"`
	Condition           string `json:"condition"`
	DestinationID       string `json:"destination_id"`
	Interval            string `json:"interval"`
	IsIncondition       string `json:"is_incondition"`
	MessageInCondition  string `json:"message_in_condition"`
	MessageOutCondition string `json:"message_out_condition"`
}

func (ExpressionModel) TableName() string {
	return "expression"
}

func (em expressionDb) Add(info ExpressionModel) (ExpressionModel, error) {
	tx := em.gorm.Create(&info)

	return info, tx.Error
}

func (em expressionDb) Get(uid string) (ExpressionModel, error) {
	expression := ExpressionModel{}
	tx := em.gorm.First(&expression, "id = ?", uid)

	return expression, tx.Error
}

func (em expressionDb) GetAll(owner_uid string) ([]ExpressionModel, error) {
	expression := []ExpressionModel{}
	tx := em.gorm.Where("owner_id = ?", owner_uid).Find(&expression)

	return expression, tx.Error
}

func (em expressionDb) SetIsIncondition(flag bool, id string) error {
	tx := em.gorm.Model(&ExpressionModel{}).Where("id = ?", id).Update("is_incondition", flag)
	return tx.Error
}
