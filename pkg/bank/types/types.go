package types

// Money представляет собой денежную сумму в минимальных единицах(центы, копейки, дирамы и т.д.)
type Money int64

// Currenct predstavlyaet kod valuti
type Currency string

// Kodi valuti
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN predstavlyaer nomer karti
type PAN string

// Card predstavlyaet info o karte
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // We used Money
	MinBalance Money // We used Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}
