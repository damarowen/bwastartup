package payment

import "github.com/google/uuid"

//untuk mencegah cycle import, maka di buat type struct Transaction
type Transaction struct {
	Id uuid.UUID
	Amount int
}