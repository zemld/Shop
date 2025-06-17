package models

type ItemResponse struct {
	Item    Item   `json:"item"`
	Message string `json:"message"`
}

type ItemBoughtResponse struct {
	Item    Item   `json:"item_in_store"`
	Message string `json:"message"`
	Bought  int    `json:"bought"`
}

type ItemDeliveredResponse struct {
	Item      Item   `json:"item_in_store"`
	Message   string `json:"message"`
	Delivered int    `json:"delivered"`
}

type ItemWithNewPriceResponse struct {
	Item     Item    `json:"item"`
	Message  string  `json:"message"`
	OldPrice float64 `json:"old_price"`
}
