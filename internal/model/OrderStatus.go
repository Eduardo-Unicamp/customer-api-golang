package model

type OrderStatus int

const (
	Unspecified OrderStatus = iota
	PENDING
	PAID
	CANCELED
)

func (os OrderStatus) String() string {
	strings := [...]string{"Unespecifed", "PENDING", "PAID", "CANCELLED"}
	if os < 0 || int(os) >= len(strings) {

		return "Unspecified"
	}
	return strings[os]
}
