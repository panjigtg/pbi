package entity

type Transaction struct {

}

func (Transaction) TableName() string {
	return "trx"
}