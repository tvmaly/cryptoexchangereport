package cointracker

type AddedCoinsStorage map[string]string

func NewAddedCoinsStorage() *AddedCoinsStorage {
	return &AddedCoinsStorage{}
}

func (acs *AddedCoinsStorage) HasCoins() bool {
	if len(*acs) > 0 {
		return true
	} else {
		return false
	}
}

func (acs *AddedCoinsStorage) Add(exchange, coin string) {
	(*acs)[exchange] = coin
}
