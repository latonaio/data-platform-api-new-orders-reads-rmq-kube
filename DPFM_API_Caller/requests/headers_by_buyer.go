package requests

type HeadersByBuyer struct {
	OrderID                          int     `json:"OrderID"`
	HeaderDeliveryStatus             *string `json:"HeaderDeliveryStatus"`
	DeliverToBusinessPartnerFullName *string `json:"DeliverToBusinessPartnerFullName"`
	BuyerBusinessPartnerFullName     *string `json:"BuyerBusinessPartnerFullName"`
}
