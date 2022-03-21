package items

type Item struct {
	Id                string      `json:"id"`
	Seller            string      `json:"seller"`
	Tittle            string      `json:"tittle"`
	Description       Description `json:"description"`
	Picture           []Picture   `json:"picture"`
	Video             string      `json:"video"`
	Price             float64     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
