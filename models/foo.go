package models

import "github.com/attapon-th/null"

const (
	fooTableName string = "foo"
	fooSchema    string = "public"
	fooTable     string = fooSchema + "." + fooTableName
)

// Foo Demo model
type Foo struct {
	ModelID
	FooName  null.String     `json:"foo_name" gorm:"type:column:foo_name;varchar(100)"`
	FooDate  null.DateString `json:"foo_date" gorm:"column:foo_date;type:date"`
	FooInt   null.Int        `json:"foo_int" gorm:"column:foo_int;type:int"`
	FooFloat null.Float      `json:"foo_float" gorm:"column:foo_float;type:float"`
	FooTime  null.Time       `json:"foo_time" gorm:"column:foo_time;timestamp"`
	IsActive bool            `json:"is_active" gorm:"column:is_active;default:true"`
	ModelRecordVersion
}

// TableName get table name
func (m Foo) TableName() string {
	return fooTable
}
