package dpfm_api_output_formatter

import (
	"data-platform-api-new-orders-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderDate,
			&pm.OrderType,
			&pm.OrderStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderDocReferenceStatus,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderBlockStatus,
			&pm.HeaderDeliveryBlockStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			OrderID:                          data.OrderID,
			OrderDate:                        data.OrderDate,
			OrderType:                        data.OrderType,
			OrderStatus:                      data.OrderStatus,
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			ContractType:                     data.ContractType,
			OrderValidityStartDate:           data.OrderValidityStartDate,
			OrderValidityEndDate:             data.OrderValidityEndDate,
			InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
			TotalNetAmount:                   data.TotalNetAmount,
			TotalTaxAmount:                   data.TotalTaxAmount,
			TotalGrossAmount:                 data.TotalGrossAmount,
			HeaderDeliveryStatus:             data.HeaderDeliveryStatus,
			HeaderBillingStatus:              data.HeaderBillingStatus,
			HeaderDocReferenceStatus:         data.HeaderDocReferenceStatus,
			TransactionCurrency:              data.TransactionCurrency,
			PricingDate:                      data.PricingDate,
			PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
			RequestedDeliveryDate:            data.RequestedDeliveryDate,
			RequestedDeliveryTime:            data.RequestedDeliveryTime,
			HeaderCompleteDeliveryIsDefined:  data.HeaderCompleteDeliveryIsDefined,
			Incoterms:                        data.Incoterms,
			PaymentTerms:                     data.PaymentTerms,
			PaymentMethod:                    data.PaymentMethod,
			Contract:	                      data.Contract,
			ContractItem:	                  data.ContractItem,
			Project:	                      data.Project,
			WBSElement:	                      data.WBSElement,
			ProductionVersion:                data.ProductionVersion,
			ProductionVersionItem:	          data.ProductionVersionItem,
			ProductionOrder:                  data.ProductionOrder,
			ProductionOrderItem:	          data.ProductionOrderItem,
			Operations:		                  data.Operations,
			OperationsItem:			          data.OperationsItem,
			OperationID:		              data.OperationID,
			ReferenceDocument:                data.ReferenceDocument,
			ReferenceDocumentItem:            data.ReferenceDocumentItem,
			AccountAssignmentGroup:           data.AccountAssignmentGroup,
			AccountingExchangeRate:           data.AccountingExchangeRate,
			InvoiceDocumentDate:              data.InvoiceDocumentDate,
			IsExportImport:                   data.IsExportImport,
			HeaderText:                       data.HeaderText,
			HeaderBlockStatus:                data.HeaderBlockStatus,
			HeaderDeliveryBlockStatus:        data.HeaderDeliveryBlockStatus,
			HeaderBillingBlockStatus:         data.HeaderBillingBlockStatus,
			ExternalReferenceDocument:		  data.ExternalReferenceDocument,
			CertificateAuthorityChain:		  data.CertificateAuthorityChain,
			UsageControlChain:		  		  data.UsageControlChain,
			CreationDate:                     data.CreationDate,
			CreationTime:                     data.CreationTime,
			LastChangeDate:                   data.LastChangeDate,
			LastChangeTime:                   data.LastChangeTime,
			IsCancelled:                      data.IsCancelled,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
			&pm.EmailAddress,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
			EmailAddress:            data.EmailAddress,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

		data := pm
		address = append(address, Address{
			OrderID:     data.OrderID,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderStatus,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.ProductSpecification,
			&pm.MarkingOfMaterial,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityStartTime,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.StockConfirmationPlantBatchValidityEndTime,
			&pm.ServicesRenderingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.QuantityPerPackage,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ProductWeightUnit,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.InvoiceDocumentDate,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityStartTime,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.ProductionPlantBatchValidityEndTime,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.InspectionPlan,
			&pm.InspectionLot,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Contract,
			&pm.ContractItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.ItemBillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.Equipment,
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, nil
		}

		data := pm
		item = append(item, Item{
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderItemCategory:                             data.OrderItemCategory,
			OrderStatus:                      			   data.OrderStatus,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			Buyer:                            			   data.Buyer,
			Seller:                           			   data.Seller,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverFromPlant:                              data.DeliverFromPlant,
			OrderItemText:                                 data.OrderItemText,
			OrderItemTextByBuyer:                          data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                         data.OrderItemTextBySeller,
			Product:                                       data.Product,
			SizeOrDimensionText:                           data.SizeOrDimensionText,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			ProductSpecification:                          data.ProductSpecification,
			MarkingOfMaterial:                             data.MarkingOfMaterial,
			BaseUnit:                                      data.BaseUnit,
			DeliveryUnit:                                  data.DeliveryUnit,
			ProductionVersion:                             data.ProductionVersion,
			ProductionVersionItem:                         data.ProductionVersionItem,
			BillOfMaterial:                                data.BillOfMaterial,
			BillOfMaterialItem:                            data.BillOfMaterialItem,
			ProductionOrder:                  			   data.ProductionOrder,
			ProductionOrderItem:	          			   data.ProductionOrderItem,
			Operations:		                  			   data.Operations,
			OperationsItem:			          			   data.OperationsItem,
			OperationID:		              			   data.OperationID,
			PricingDate:                                   data.PricingDate,
			PriceDetnExchangeRate:                         data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                         data.RequestedDeliveryDate,
			RequestedDeliveryTime:                         data.RequestedDeliveryTime,
			DeliverToPlantTimeZone:                        data.DeliverToPlantTimeZone,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityStartTime:          data.DeliverToPlantBatchValidityStartTime,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverToPlantBatchValidityEndTime:            data.DeliverToPlantBatchValidityEndTime,
			DeliverFromPlantTimeZone:                      data.DeliverFromPlantTimeZone,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityStartTime:        data.DeliverFromPlantBatchValidityStartTime,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliverFromPlantBatchValidityEndTime:          data.DeliverFromPlantBatchValidityEndTime,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                data.StockConfirmationPlantTimeZone,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityStartTime:  data.StockConfirmationPlantBatchValidityStartTime,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			StockConfirmationPlantBatchValidityEndTime:    data.StockConfirmationPlantBatchValidityEndTime,
			ServicesRenderingDate:                         data.ServicesRenderingDate,
			OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:                   data.OrderQuantityInDeliveryUnit,
			QuantityPerPackage:                            data.QuantityPerPackage,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
			ProductWeightUnit:                             data.ProductWeightUnit,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			InternalCapacityQuantity:                      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:                  data.InternalCapacityQuantityUnit,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			InvoiceDocumentDate:                           data.InvoiceDocumentDate,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityStartTime:         data.ProductionPlantBatchValidityStartTime,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			ProductionPlantBatchValidityEndTime:           data.ProductionPlantBatchValidityEndTime,
			InspectionPlantBusinessPartner:                data.InspectionPlantBusinessPartner,
			InspectionPlant:                               data.InspectionPlant,
			InspectionPlan:                                data.InspectionPlan,
			InspectionLot:                           	   data.InspectionLot,
			Incoterms:                                     data.Incoterms,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			DueCalculationBaseDate:                        data.DueCalculationBaseDate,
			PaymentDueDate:                                data.PaymentDueDate,
			NetPaymentDays:                                data.NetPaymentDays,
			PaymentMethod:                                 data.PaymentMethod,
			Contract:	                      			   data.Contract,
			ContractItem:	                  			   data.ContractItem,
			Project:                                       data.Project,
			WBSElement:                                    data.WBSElement,
			AccountingExchangeRate:                        data.AccountingExchangeRate,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                            data.ItemDeliveryStatus,
			IssuingStatus:                                 data.IssuingStatus,
			ReceivingStatus:                               data.ReceivingStatus,
			ItemBillingStatus:                             data.ItemBillingStatus,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
			Equipment:                                     data.Equipment,
			FreightAgreement:                              data.FreightAgreement,
			FreightAgreementItem:                          data.FreightAgreementItem,
			ItemBlockStatus:                               data.ItemBlockStatus,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
			ExternalReferenceDocument:		  			   data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:	  			   data.ExternalReferenceDocumentItem,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
	}

	return &item, nil
}

func ConvertToItemPricingElement(rows *sql.Rows) (*[]ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPricingElement, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
			CreationDate:               data.CreationDate,
			CreationTime:               data.CreationTime,
			LastChangeDate:             data.LastChangeDate,
			LastChangeTime:             data.LastChangeTime,
			IsCancelled:                data.IsCancelled,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemPricingElement, nil
	}

	return &itemPricingElement, nil
}

func ConvertToItemScheduleLine(rows *sql.Rows) (*[]ItemScheduleLine, error) {
	defer rows.Close()
	itemScheduleLine := make([]ItemScheduleLine, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemScheduleLine{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ScheduleLine,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Product,
			&pm.StockConfirmationBussinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.ConfirmedDeliveryDate,
			&pm.ConfirmedDeliveryTime,
			&pm.ScheduleLineOrderQuantityInBaseUnit,
			&pm.OriginalOrderQuantityInBaseUnit,
			&pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			&pm.DeliveredQuantityInBaseUnit,
			&pm.UndeliveredQuantityInBaseUnit,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.PlusMinusFlag,
			&pm.ItemScheduleLineDeliveryBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.ExternalReferenceDocumentItemScheduleLine,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemScheduleLine, err
		}

		data := pm
		itemScheduleLine = append(itemScheduleLine, ItemScheduleLine{
			OrderID:                                         data.OrderID,
			OrderItem:                                       data.OrderItem,
			ScheduleLine:                                    data.ScheduleLine,
			SupplyChainRelationshipID:                       data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID:         data.SupplyChainRelationshipStockConfPlantID,
			Product:                                         data.Product,
			StockConfirmationBussinessPartner:               data.StockConfirmationBussinessPartner,
			StockConfirmationPlant:                          data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                  data.StockConfirmationPlantTimeZone,
			StockConfirmationPlantBatch:                     data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:    data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:      data.StockConfirmationPlantBatchValidityEndDate,
			RequestedDeliveryDate:                           data.RequestedDeliveryDate,
			RequestedDeliveryTime:                           data.RequestedDeliveryTime,
			ConfirmedDeliveryDate:                           data.ConfirmedDeliveryDate,
			ConfirmedDeliveryTime:                           data.ConfirmedDeliveryTime,
			ScheduleLineOrderQuantityInBaseUnit:             data.ScheduleLineOrderQuantityInBaseUnit,
			OriginalOrderQuantityInBaseUnit:                 data.OriginalOrderQuantityInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			DeliveredQuantityInBaseUnit:                     data.DeliveredQuantityInBaseUnit,
			UndeliveredQuantityInBaseUnit:                   data.UndeliveredQuantityInBaseUnit,
			OpenConfirmedQuantityInBaseUnit:                 data.OpenConfirmedQuantityInBaseUnit,
			StockIsFullyConfirmed:                           data.StockIsFullyConfirmed,
			PlusMinusFlag:                                   data.PlusMinusFlag,
			ItemScheduleLineDeliveryBlockStatus:             data.ItemScheduleLineDeliveryBlockStatus,
			ExternalReferenceDocument:		  			   	 data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:	  			   	 data.ExternalReferenceDocumentItem,
			ExternalReferenceDocumentItemScheduleLine:	   	 data.ExternalReferenceDocumentItemScheduleLine,
			CreationDate:                                    data.CreationDate,
			CreationTime:                                    data.CreationTime,
			LastChangeDate:                                  data.LastChangeDate,
			LastChangeTime:                                  data.LastChangeTime,
			IsCancelled:                                     data.IsCancelled,
			IsMarkedForDeletion:                             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemScheduleLine, nil
	}

	return &itemScheduleLine, nil
}
