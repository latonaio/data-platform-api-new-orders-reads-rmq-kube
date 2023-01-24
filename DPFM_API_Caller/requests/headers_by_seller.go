package requests

type HeadersBySeller struct {
	OrderID                          int     `json:"OrderID"`
	HeaderDeliveryStatus             *string `json:"HeaderDeliveryStatus"`
	DeliverToBusinessPartnerFullName *string `json:"DeliverToBusinessPartnerFullName"`
	SellerBusinessPartnerFullName    *string `json:"SellerBusinessPartnerFullName"`
}
