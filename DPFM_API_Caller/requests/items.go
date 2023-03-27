package requests

type Items struct {
	OrderID                     int      `json:"OrderID"`
	OrderItem                   int      `json:"OrderItem"`
	OrderItemCategory           *string  `json:"OrderItemCategory"`
	OrderItemText               *string  `json:"OrderItemText"`
	OrderItemTextByBuyer        *string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller       *string  `json:"OrderItemTextBySeller"`
	OrderQuantityInBaseUnit     *float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit"`
	BaseUnit                    *string  `json:"BaseUnit"`
	DeliveryUnit                *string  `json:"DeliveryUnit"`
	Product                     *string  `json:"Product"`
	NetAmount                   *float32 `json:"NetAmount"`
	DeliverToParty              *int     `json:"DeliverToParty"`
	DeliverFromParty            *int     `json:"DeliverFromParty"`
	RequestedDeliveryDate       *string  `json:"RequestedDeliveryDate"`
	IsCancelled                 *bool    `json:"IsCancelled"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
}
