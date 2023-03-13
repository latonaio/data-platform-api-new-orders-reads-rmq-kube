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
			&pm.CreationDate,
			&pm.LastChangeDate,
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
			CreationDate:                     data.CreationDate,
			LastChangeDate:                   data.LastChangeDate,
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

func ConvertToHeaderDoc(rows *sql.Rows) (*[]HeaderDoc, error) {
	defer rows.Close()
	headerDoc := make([]HeaderDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeaderDoc{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headerDoc, err
		}

		data := pm
		headerDoc = append(headerDoc, HeaderDoc{
			OrderID:                  data.OrderID,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headerDoc, nil
	}

	return &headerDoc, nil
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
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityStartTime,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverToPlantBatchValidityEndTime,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityStartTime,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliverFromPlantBatchValidityEndTime,
			&pm.DeliveryUnit,
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
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
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
			&pm.Project,
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
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
			&pm.ItemBillingBlockStatus,
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
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			OrderItemText:                                 data.OrderItemText,
			OrderItemTextByBuyer:                          data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                         data.OrderItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			PricingDate:                                   data.PricingDate,
			PriceDetnExchangeRate:                         data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                         data.RequestedDeliveryDate,
			RequestedDeliveryTime:                         data.RequestedDeliveryTime,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverToPlantTimeZone:                        data.DeliverToPlantTimeZone,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityStartTime:          data.DeliverToPlantBatchValidityStartTime,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverToPlantBatchValidityEndTime:            data.DeliverToPlantBatchValidityEndTime,
			DeliverFromPlant:                              data.DeliverFromPlant,
			DeliverFromPlantTimeZone:                      data.DeliverFromPlantTimeZone,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityStartTime:        data.DeliverFromPlantBatchValidityStartTime,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliverFromPlantBatchValidityEndTime:          data.DeliverFromPlantBatchValidityEndTime,
			DeliveryUnit:                                  data.DeliveryUnit,
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
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
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
			Project:                                       data.Project,
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
			ItemBlockStatus:                               data.ItemBlockStatus,
			ItemDeliveryBlockStatus:                       data.ItemDeliveryBlockStatus,
			ItemBillingBlockStatus:                        data.ItemBillingBlockStatus,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
			InspectionPlan:                                data.InspectionPlan,
			InspectionPlant:                               data.InspectionPlant,
			InspectionOrder:                               data.InspectionOrder,
		})
	}

	return &item, nil
}

func ConvertToItems(rows *sql.Rows) (*[]Item, error) {
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
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.Product,
			&pm.NetAmount,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.RequestedDeliveryDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			OrderID:                     data.OrderID,
			OrderItem:                   data.OrderItem,
			OrderItemCategory:           data.OrderItemCategory,
			OrderItemText:               data.OrderItemText,
			OrderItemTextByBuyer:        data.OrderItemTextByBuyer,
			OrderItemTextBySeller:       data.OrderItemTextBySeller,
			OrderQuantityInBaseUnit:     data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit: data.OrderQuantityInDeliveryUnit,
			BaseUnit:                    data.BaseUnit,
			DeliveryUnit:                data.DeliveryUnit,
			Product:                     data.Product,
			NetAmount:                   data.NetAmount,
			DeliverToParty:              data.DeliverToParty,
			DeliverFromParty:            data.DeliverFromParty,
			RequestedDeliveryDate:       data.RequestedDeliveryDate,
			IsCancelled:                 data.IsCancelled,
			IsMarkedForDeletion:         data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
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
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.PricingProcedureCounter,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPricingElement, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
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
			&pm.ScheduleLineOrderQuantity,
			&pm.OriginalOrderQuantityInBaseUnit,
			&pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			&pm.ConfirmedOrderQuantityByPDTAvailCheck,
			&pm.DeliveredQuantityInBaseUnit,
			&pm.UndeliveredQuantityInBaseUnit,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.PlusMinusFlag,
			&pm.ItemScheduleLineDeliveryBlockStatus,
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
			ScheduleLineOrderQuantity:                       data.ScheduleLineOrderQuantity,
			OriginalOrderQuantityInBaseUnit:                 data.OriginalOrderQuantityInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheck:           data.ConfirmedOrderQuantityByPDTAvailCheck,
			DeliveredQuantityInBaseUnit:                     data.DeliveredQuantityInBaseUnit,
			UndeliveredQuantityInBaseUnit:                   data.UndeliveredQuantityInBaseUnit,
			OpenConfirmedQuantityInBaseUnit:                 data.OpenConfirmedQuantityInBaseUnit,
			StockIsFullyConfirmed:                           data.StockIsFullyConfirmed,
			PlusMinusFlag:                                   data.PlusMinusFlag,
			ItemScheduleLineDeliveryBlockStatus:             data.ItemScheduleLineDeliveryBlockStatus,
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

func ConvertToHeadersBySeller(rows *sql.Rows) (*[]HeadersBySeller, error) {
	defer rows.Close()
	headersBySeller := make([]HeadersBySeller, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeadersBySeller{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderDeliveryStatus,
			&pm.DeliverToBusinessPartnerFullName,
			&pm.SellerBusinessPartnerFullName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headersBySeller, err
		}

		data := pm
		headersBySeller = append(headersBySeller, HeadersBySeller{
			OrderID:                          data.OrderID,
			HeaderDeliveryStatus:             data.HeaderDeliveryStatus,
			DeliverToBusinessPartnerFullName: data.DeliverToBusinessPartnerFullName,
			SellerBusinessPartnerFullName:    data.SellerBusinessPartnerFullName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headersBySeller, nil
	}

	return &headersBySeller, nil
}

func ConvertToHeadersByBuyer(rows *sql.Rows) (*[]HeadersByBuyer, error) {
	defer rows.Close()
	headersByBuyer := make([]HeadersByBuyer, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeadersByBuyer{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderDeliveryStatus,
			&pm.DeliverToBusinessPartnerFullName,
			&pm.BuyerBusinessPartnerFullName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headersByBuyer, err
		}

		data := pm
		headersByBuyer = append(headersByBuyer, HeadersByBuyer{
			OrderID:                          data.OrderID,
			HeaderDeliveryStatus:             data.HeaderDeliveryStatus,
			DeliverToBusinessPartnerFullName: data.DeliverToBusinessPartnerFullName,
			BuyerBusinessPartnerFullName:     data.BuyerBusinessPartnerFullName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headersByBuyer, nil
	}

	return &headersByBuyer, nil
}
