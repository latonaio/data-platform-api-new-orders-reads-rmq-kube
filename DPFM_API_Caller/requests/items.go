package requests

type Items struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderItemText               *string  `json:"OrderItemText"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit"`
	NetAmount                   *float32 `json:"NetAmount"`
	Product                     *string  `json:"Product"`
	ProductDescription          *string  `json:"ProductDescription"`
	ConditionRateValue          *float32 `json:"ConditionRateValue"`
	ConfirmedDeliveryDate       *string  `json:"ConfirmedDeliveryDate"`
	ItemIsCancelled             bool     `json:"ItemIsCancelled"`
	ItemIsDeleted               bool     `json:"ItemIsDeleted"`
}
