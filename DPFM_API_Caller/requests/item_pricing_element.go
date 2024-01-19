package requests

type ItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionType              string   `json:"ConditionType"`
	PricingDate                string   `json:"PricingDate"`
	ConditionRateValue         float32  `json:"ConditionRateValue"`
	ConditionRateValueUnit     int      `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     int      `json:"ConditionScaleQuantity"`
	ConditionCurrency          string   `json:"ConditionCurrency"`
	ConditionQuantity          float32  `json:"ConditionQuantity"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            float32  `json:"ConditionAmount"`
	TransactionCurrency        string   `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
	CreationDate               string   `json:"CreationDate"`
	CreationTime               string   `json:"CreationTime"`
	LastChangeDate             string   `json:"LastChangeDate"`
	LastChangeTime             string   `json:"LastChangeTime"`
	IsCancelled                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
}
