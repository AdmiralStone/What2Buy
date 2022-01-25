package models

type Item struct {
	ItemID    int     `json:"itemId"`
	ItemName  string  `json:"itemName"`
	ItemPrice float64 `json:"itemPrice"`
	ItemVotes int     `json:"itemVotes"`
}
