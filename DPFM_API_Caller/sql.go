package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-new-orders-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-new-orders-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var itemScheduleLine *[]dpfm_api_output_formatter.ItemScheduleLine
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner
	var headerDoc *[]dpfm_api_output_formatter.HeaderDoc
	var headersBySeller *[]dpfm_api_output_formatter.HeadersBySeller
	var headersByBuyer *[]dpfm_api_output_formatter.HeadersByBuyer
	// var items *[]dpfm_api_output_formatter.Items

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		case "ItemPricingElements":
			func() {
				itemPricingElement = c.ItemPricingElements(mtx, input, output, errs, log)
			}()
		case "ItemScheduleLine":
			func() {
				itemScheduleLine = c.ItemScheduleLine(mtx, input, output, errs, log)
			}()
		case "ItemScheduleLines":
			func() {
				itemScheduleLine = c.ItemScheduleLines(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "HeaderDoc":
			func() {
				headerDoc = c.HeaderDoc(mtx, input, output, errs, log)
			}()
		case "HeadersBySeller":
			func() {
				headersBySeller = c.HeadersBySeller(mtx, input, output, errs, log)
			}()
		case "HeadersByBuyer":
			func() {
				headersByBuyer = c.HeadersByBuyer(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		ItemScheduleLine:   itemScheduleLine,
		Address:            address,
		Partner:            partner,
		HeaderDoc:          headerDoc,
		HeadersBySeller:    headersBySeller,
		HeadersByBuyer:     headersByBuyer,
		// Items:              items,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	orderID := input.Header.OrderID
	rows, err := c.db.Query(
		`SELECT
		header.OrderID,header.OrderDate,header.OrderType,header.SupplyChainRelationshipID,header.SupplyChainRelationshipBillingID,header.
		SupplyChainRelationshipPaymentID,header.Buyer,header.Seller,header.BillToParty,header.BillFromParty,header.BillToCountry,header.
		BillFromCountry,header.Payer,header.Payee,header.CreationDate,header.LastChangeDate,header.ContractType,header.OrderValidityStartDate,header.
		OrderValidityEndDate,header.InvoicePeriodStartDate,header.InvoicePeriodEndDate,header.TotalNetAmount,header.TotalTaxAmount,header.
		TotalGrossAmount,header.HeaderDeliveryStatus,header.HeaderBillingStatus,header.HeaderDocReferenceStatus,header.
		TransactionCurrency,header.PricingDate,header.PriceDetnExchangeRate,header.RequestedDeliveryDate,header.RequestedDeliveryTime,header.HeaderCompleteDeliveryIsDefined,header.
		Incoterms,terms.PaymentTermsName,method.PaymentMethodName,header.ReferenceDocument,header.ReferenceDocumentItem,header.AccountAssignmentGroup,header.
		AccountingExchangeRate,header.InvoiceDocumentDate,header.IsExportImport,header.HeaderText,header.HeaderBlockStatus,header.
		HeaderDeliveryBlockStatus,header.HeaderBillingBlockStatus,header.IsCancelled,header.IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data AS header
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_text_data AS terms
		ON header.PaymentTerms = terms.PaymentTerms
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_text_data AS method
		ON header.PaymentMethod = method.PaymentMethod
		WHERE OrderID = ?;`, orderID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.HeaderCompleteDeliveryIsDefined != nil {
		where = fmt.Sprintf("%s\nAND HeaderCompleteDeliveryIsDefined = %v", where, *input.Header.HeaderCompleteDeliveryIsDefined)
	}
	if input.Header.HeaderDeliveryBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryBlockStatus = %v", where, *input.Header.HeaderDeliveryBlockStatus)
	}
	if input.Header.HeaderDeliveryStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderDeliveryStatus = '%s'", where, *input.Header.HeaderDeliveryStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	idWhere := ""
	if input.Header.Buyer != nil && input.Header.Seller != nil {
		idWhere = fmt.Sprintf("\nAND ( Buyer = %d OR Seller = %d ) ", *input.Header.Buyer, *input.Header.Seller)
	} else if input.Header.Buyer != nil {
		idWhere = fmt.Sprintf("\nAND Buyer = %d ", *input.Header.Buyer)
	} else if input.Header.Seller != nil {
		idWhere = fmt.Sprintf("\nAND Seller = %d ", *input.Header.Seller)
	}

	rows, err := c.db.Query(
		`SELECT
		header.OrderID,header.OrderDate,header.OrderType,header.SupplyChainRelationshipID,header.SupplyChainRelationshipBillingID,header.
		SupplyChainRelationshipPaymentID,header.Buyer,header.Seller,header.BillToParty,header.BillFromParty,header.BillToCountry,header.
		BillFromCountry,header.Payer,header.Payee,header.CreationDate,header.LastChangeDate,header.ContractType,header.OrderValidityStartDate,header.
		OrderValidityEndDate,header.InvoicePeriodStartDate,header.InvoicePeriodEndDate,header.TotalNetAmount,header.TotalTaxAmount,header.
		TotalGrossAmount,header.HeaderDeliveryStatus,header.HeaderBillingStatus,header.HeaderDocReferenceStatus,header.
		TransactionCurrency,header.PricingDate,header.PriceDetnExchangeRate,header.RequestedDeliveryDate,header.RequestedDeliveryDate,header.HeaderCompleteDeliveryIsDefined,header.
		Incoterms,terms.PaymentTermsName,method.PaymentMethodName,header.ReferenceDocument,header.ReferenceDocumentItem,header.AccountAssignmentGroup,header.
		AccountingExchangeRate,header.InvoiceDocumentDate,header.IsExportImport,header.HeaderText,header.HeaderBlockStatus,header.
		HeaderDeliveryBlockStatus,header.HeaderBillingBlockStatus,header.IsCancelled,header.IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data AS header
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_text_data AS terms
		ON header.PaymentTerms = terms.PaymentTerms
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_text_data AS method
		ON header.PaymentMethod = method.PaymentMethod
		` + where + idWhere + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, orderID, v.OrderItem)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT 
		OrderID, OrderItem, OrderItemCategory, SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID, SupplyChainRelationshipStockConfPlantID, SupplyChainRelationshipProductionPlantID, 
		OrderItemText, OrderItemTextByBuyer, OrderItemTextBySeller, Product, ProductStandardID, ProductGroup, BaseUnit, PricingDate, 
		PriceDetnExchangeRate, RequestedDeliveryDate, RequestedDeliveryTime, DeliverToParty, DeliverFromParty, CreationDate, CreationTime, 
		LastChangeDate, LastChangeTime, DeliverToPlant, DeliverToPlantTimeZone, DeliverToPlantStorageLocation, ProductIsBatchManagedInDeliverToPlant, 
		BatchMgmtPolicyInDeliverToPlant, DeliverToPlantBatch, DeliverToPlantBatchValidityStartDate, DeliverToPlantBatchValidityStartTime,
		DeliverToPlantBatchValidityEndDate, DeliverToPlantBatchValidityEndTime, DeliverFromPlant, DeliverFromPlantTimeZone, 
		DeliverFromPlantStorageLocation, ProductIsBatchManagedInDeliverFromPlant, BatchMgmtPolicyInDeliverFromPlant, 
		DeliverFromPlantBatch, DeliverFromPlantBatchValidityStartDate, DeliverFromPlantBatchValidityStartTime, DeliverFromPlantBatchValidityEndDate,
		DeliverFromPlantBatchValidityEndTime, DeliveryUnit, StockConfirmationBusinessPartner, StockConfirmationPlant, 
		StockConfirmationPlantTimeZone, ProductIsBatchManagedInStockConfirmationPlant, BatchMgmtPolicyInStockConfirmationPlant, 
		StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate, StockConfirmationPlantBatchValidityStartTime, 
		StockConfirmationPlantBatchValidityEndDate, StockConfirmationPlantBatchValidityEndTime, ServicesRenderingDate, 
		OrderQuantityInBaseUnit, OrderQuantityInDeliveryUnit, StockConfirmationPolicy, StockConfirmationStatus, 
		ConfirmedOrderQuantityInBaseUnit, ItemWeightUnit, ProductGrossWeight, ItemGrossWeight, ProductNetWeight, ItemNetWeight,
		InternalCapacityQuantity, InternalCapacityQuantityUnit, NetAmount, TaxAmount, GrossAmount, InvoiceDocumentDate,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantTimeZone, ProductionPlantStorageLocation, 
		ProductIsBatchManagedInProductionPlant, BatchMgmtPolicyInProductionPlant, ProductionPlantBatch, ProductionPlantBatchValidityStartDate, 
		ProductionPlantBatchValidityStartTime, ProductionPlantBatchValidityEndDate, ProductionPlantBatchValidityEndTime,
		Incoterms, TransactionTaxClassification, ProductTaxClassificationBillToCountry, ProductTaxClassificationBillFromCountry, 
		DefinedTaxClassification, AccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, DueCalculationBaseDate,
		PaymentDueDate, NetPaymentDays, PaymentMethod, Project, AccountingExchangeRate, ReferenceDocument, ReferenceDocumentItem,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, IssuingStatus, ReceivingStatus, ItemBillingStatus, TaxCode, TaxRate, 
		CountryOfOrigin, CountryOfOriginLanguage, ItemBlockStatus, ItemDeliveryBlockStatus, ItemBillingBlockStatus, IsCancelled,
		IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}
	where := fmt.Sprintf("WHERE item.OrderID IS NOT NULL\nAND header.OrderID = %d", input.Header.OrderID)
	where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	if input.Header.Buyer != nil {
		where = fmt.Sprintf("%s\nAND header.Buyer = %d", where, *input.Header.Buyer)
	}
	if input.Header.Seller != nil {
		where = fmt.Sprintf("%s\nAND header.Seller = %d", where, *input.Header.Seller)
	}
	if item != nil {
		if item.ItemCompleteDeliveryIsDefined != nil {
			where = fmt.Sprintf("%s\nAND item.ItemCompleteDeliveryIsDefined = %v", where, *item.ItemCompleteDeliveryIsDefined)
		}
		if item.ItemDeliveryStatus != nil {
			where = fmt.Sprintf("%s\nAND item.ItemDeliveryStatus = '%v'", where, *item.ItemDeliveryStatus)
		}
		if item.ItemDeliveryBlockStatus != nil {
			where = fmt.Sprintf("%s\nAND item.ItemDeliveryBlockStatus = %v", where, *item.ItemDeliveryBlockStatus)
		}
		if item.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND item.IsCancelled = %v", where, *item.IsCancelled)
		}
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT 
			item.OrderID,
			item.OrderItem,
			item.OrderItemCategory,
			item.OrderItemText,
			item.OrderItemTextByBuyer,
			item.OrderItemTextBySeller,
			item.OrderQuantityInBaseUnit,
			item.OrderQuantityInDeliveryUnit,
			item.BaseUnit,
			item.DeliveryUnit,
			item.Product,
			item.NetAmount,
			item.DeliverToParty,
			item.DeliverFromParty,
			item.RequestedDeliveryDate,
			item.IsCancelled,
			item.IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data as header
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data as item
		ON header.OrderID = item.OrderID ` + where + ` ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItems(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, orderID, v.OrderItem)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElements(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemPricingElement {
	where := fmt.Sprintf("WHERE OrderID = %d", input.Header.OrderID)
	if input.Header.Buyer != nil {
		where = fmt.Sprintf("%s\nAND Buyer = %d", where, *input.Header.Buyer)
	}
	if input.Header.Seller != nil {
		where = fmt.Sprintf("%s\nAND Seller = %d", where, *input.Header.Seller)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		` + where + ` ;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
func (c *DPFMAPICaller) ItemScheduleLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemScheduleLine {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemScheduleLine := v.ItemScheduleLine
		for _, w := range itemScheduleLine {
			args = append(args, orderID, v.OrderItem, w.ScheduleLine)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem, ScheduleLine) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemScheduleLine(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemScheduleLines(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemScheduleLine {
	orderID := input.Header.OrderID
	item := input.Header.Item[0]

	where := fmt.Sprintf("WHERE (OrderID, OrderItem) IN ( (%d, %d) ) ", orderID, item.OrderItem)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemScheduleLine(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	orderID := input.Header.OrderID
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, orderID, v.AddressID)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_address_data
		WHERE (OrderID, AddressID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	orderID := input.Header.OrderID
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, orderID, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data
		WHERE (OrderID, PartnerFunction, BusinessPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderDoc(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeaderDoc {
	var args []interface{}
	orderID := input.Header.OrderID
	headerDoc := input.Header.HeaderDoc

	cnt := 0
	for _, v := range headerDoc {
		args = append(args, orderID, v.DocType, v.DocVersionID, v.DocID)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_doc_data
		WHERE (OrderID, DocType, DocVersionID, DocID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeaderDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySeller(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeadersBySeller {
	seller := input.Header.Seller

	rows, err := c.db.Query(
		`SELECT OrdersHeader.OrderID,
		OrdersHeader.HeaderDeliveryStatus,
		OrdersHeaderPartnerDeliverTo.BusinessPartnerFullName as DeliverToBusinessPartnerFullName,
		OrdersHeaderPartnerSeller.BusinessPartnerFullName as SellerBusinessPartnerFullName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data as OrdersHeader
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data as OrdersHeaderPartnerDeliverTo
		ON OrdersHeader.OrderID = OrdersHeaderPartnerDeliverTo.OrderID AND OrdersHeaderPartnerDeliverTo.PartnerFunction = "DELIVER_TO"
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data as OrdersHeaderPartnerSeller
		ON OrdersHeader.OrderID = OrdersHeaderPartnerSeller.OrderID AND OrdersHeaderPartnerSeller.PartnerFunction = "SELLER"
		WHERE (OrdersHeader.Seller) = (?);`, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeadersBySeller(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBuyer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeadersByBuyer {
	buyer := input.Header.Buyer

	rows, err := c.db.Query(
		`SELECT OrdersHeader.OrderID,
		OrdersHeader.HeaderDeliveryStatus,
		OrdersHeaderPartnerDeliverTo.BusinessPartnerFullName as DeliverToBusinessPartnerFullName,
		OrdersHeaderPartnerBuyer.BusinessPartnerFullName as BuyerBusinessPartnerFullName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data as OrdersHeader
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data as OrdersHeaderPartnerDeliverTo
		ON OrdersHeader.OrderID = OrdersHeaderPartnerDeliverTo.OrderID AND OrdersHeaderPartnerDeliverTo.PartnerFunction = "DELIVER_TO"
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data as OrdersHeaderPartnerBuyer
		ON OrdersHeader.OrderID = OrdersHeaderPartnerBuyer.OrderID AND OrdersHeaderPartnerBuyer.PartnerFunction = "BUYER"
		WHERE (OrdersHeader.Buyer) = (?);`, buyer,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeadersByBuyer(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

// func (c *DPFMAPICaller) Items(
// 	mtx *sync.Mutex,
// 	input *dpfm_api_input_reader.SDC,
// 	output *dpfm_api_output_formatter.SDC,
// 	errs *[]error,
// 	log *logger.Logger,
// ) *[]dpfm_api_output_formatter.Items {
// 	orderID := input.Header.OrderID

// 	rows, err := c.db.Query(
// 		`SELECT OrderItems.OrderID,
// 		OrderItems.OrderItem,
// 		OrderItems.OrderItemText,
// 		OrderItems.OrderQuantityInDeliveryUnit,
// 		OrderItems.NetAmount,
// 		OrderItems.Product,
// 		ProductDescription.ProductDescription,
// 		OrdersItemPricingElement.ConditionRateValue,
// 		OrdersItemScheduleLine.ConfirmedDeliveryDate
// 		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data as OrderItems
// 		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data as ProductDescription
// 		ON OrderItems.Product = ProductDescription.Product
// 		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data as OrdersItemPricingElement
// 		ON OrderItems.OrderID = OrdersItemPricingElement.OrderID AND OrderItems.OrderItem = OrdersItemPricingElement.OrderItem
// 		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data as OrdersItemScheduleLine
// 		ON OrderItems.OrderID = OrdersItemScheduleLine.OrderID AND OrderItems.OrderItem = OrdersItemScheduleLine.OrderItem
// 		WHERE (OrderItems.OrderID) = (?);`, orderID,
// 	)
// 	if err != nil {
// 		*errs = append(*errs, err)
// 		return nil
// 	}
// defer rows.Close()

// 	data, err := dpfm_api_output_formatter.ConvertToItems(rows)
// 	if err != nil {
// 		*errs = append(*errs, err)
// 		return nil
// 	}

// 	return data
// }
