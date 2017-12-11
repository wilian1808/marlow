package models

type MultiAuto struct {
	table  bool   `marlow:"tableName=multi_auto&dialect=postgres&primaryKey=id"`
	ID     uint   `marlow:"column=id&autoIncrement=true"`
	Status string `marlow:"column=status&autoIncrement=true"`
	Name   string `marlow:"column=name"`
}
