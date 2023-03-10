package ai_parser

import (
	"github.com/ramsfords/backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

type ParserResponse struct {
	Amazon struct {
		Status        string `json:"status"`
		ExtractedData []struct {
			CustomerInformation struct {
				CustomerName              any `json:"customer_name"`
				CustomerAddress           any `json:"customer_address"`
				CustomerEmail             any `json:"customer_email"`
				CustomerID                any `json:"customer_id"`
				CustomerTaxID             any `json:"customer_tax_id"`
				CustomerMailingAddress    any `json:"customer_mailing_address"`
				CustomerBillingAddress    any `json:"customer_billing_address"`
				CustomerShippingAddress   any `json:"customer_shipping_address"`
				CustomerServiceAddress    any `json:"customer_service_address"`
				CustomerRemittanceAddress any `json:"customer_remittance_address"`
			} `json:"customer_information"`
			MerchantInformation struct {
				MerchantName    any `json:"merchant_name"`
				MerchantAddress any `json:"merchant_address"`
				MerchantPhone   any `json:"merchant_phone"`
				MerchantEmail   any `json:"merchant_email"`
				MerchantFax     any `json:"merchant_fax"`
				MerchantWebsite any `json:"merchant_website"`
				MerchantTaxID   any `json:"merchant_tax_id"`
				MerchantSiret   any `json:"merchant_siret"`
				MerchantSiren   any `json:"merchant_siren"`
			} `json:"merchant_information"`
			InvoiceNumber         any     `json:"invoice_number"`
			InvoiceTotal          float64 `json:"invoice_total"`
			InvoiceSubtotal       any     `json:"invoice_subtotal"`
			AmountDue             any     `json:"amount_due"`
			PreviousUnpaidBalance any     `json:"previous_unpaid_balance"`
			Discount              any     `json:"discount"`
			Taxes                 []struct {
				Value any `json:"value"`
				Rate  any `json:"rate"`
			} `json:"taxes"`
			PaymentTerm    any `json:"payment_term"`
			PurchaseOrder  any `json:"purchase_order"`
			Date           any `json:"date"`
			DueDate        any `json:"due_date"`
			ServiceDate    any `json:"service_date"`
			ServiceDueDate any `json:"service_due_date"`
			Locale         struct {
				Currency any `json:"currency"`
				Language any `json:"language"`
			} `json:"locale"`
			BankInformations struct {
				AccountNumber any `json:"account_number"`
				Iban          any `json:"iban"`
				Bsb           any `json:"bsb"`
				SortCode      any `json:"sort_code"`
				VatNumber     any `json:"vat_number"`
				RootingNumber any `json:"rooting_number"`
				Swift         any `json:"swift"`
			} `json:"bank_informations"`
			ItemLines []struct {
				Description string `json:"description"`
				Quantity    any    `json:"quantity"`
				Amount      any    `json:"amount"`
				UnitPrice   any    `json:"unit_price"`
				Discount    any    `json:"discount"`
				ProductCode any    `json:"product_code"`
				DateItem    any    `json:"date_item"`
				TaxItem     any    `json:"tax_item"`
				TaxRate     any    `json:"tax_rate"`
			} `json:"item_lines"`
		} `json:"extracted_data"`
	} `json:"amazon"`
	EdenAi struct {
		Status        string `json:"status"`
		ExtractedData []struct {
			CustomerInformation struct {
				CustomerName              any `json:"customer_name"`
				CustomerAddress           any `json:"customer_address"`
				CustomerEmail             any `json:"customer_email"`
				CustomerID                any `json:"customer_id"`
				CustomerTaxID             any `json:"customer_tax_id"`
				CustomerMailingAddress    any `json:"customer_mailing_address"`
				CustomerBillingAddress    any `json:"customer_billing_address"`
				CustomerShippingAddress   any `json:"customer_shipping_address"`
				CustomerServiceAddress    any `json:"customer_service_address"`
				CustomerRemittanceAddress any `json:"customer_remittance_address"`
			} `json:"customer_information"`
			MerchantInformation struct {
				MerchantName    any `json:"merchant_name"`
				MerchantAddress any `json:"merchant_address"`
				MerchantPhone   any `json:"merchant_phone"`
				MerchantEmail   any `json:"merchant_email"`
				MerchantFax     any `json:"merchant_fax"`
				MerchantWebsite any `json:"merchant_website"`
				MerchantTaxID   any `json:"merchant_tax_id"`
				MerchantSiret   any `json:"merchant_siret"`
				MerchantSiren   any `json:"merchant_siren"`
			} `json:"merchant_information"`
			InvoiceNumber         any     `json:"invoice_number"`
			InvoiceTotal          float64 `json:"invoice_total"`
			InvoiceSubtotal       any     `json:"invoice_subtotal"`
			AmountDue             any     `json:"amount_due"`
			PreviousUnpaidBalance any     `json:"previous_unpaid_balance"`
			Discount              any     `json:"discount"`
			Taxes                 []struct {
				Value any `json:"value"`
				Rate  any `json:"rate"`
			} `json:"taxes"`
			PaymentTerm    any `json:"payment_term"`
			PurchaseOrder  any `json:"purchase_order"`
			Date           any `json:"date"`
			DueDate        any `json:"due_date"`
			ServiceDate    any `json:"service_date"`
			ServiceDueDate any `json:"service_due_date"`
			Locale         struct {
				Currency any `json:"currency"`
				Language any `json:"language"`
			} `json:"locale"`
			BankInformations struct {
				AccountNumber any `json:"account_number"`
				Iban          any `json:"iban"`
				Bsb           any `json:"bsb"`
				SortCode      any `json:"sort_code"`
				VatNumber     any `json:"vat_number"`
				RootingNumber any `json:"rooting_number"`
				Swift         any `json:"swift"`
			} `json:"bank_informations"`
			ItemLines []struct {
				Description string `json:"description"`
				Quantity    any    `json:"quantity"`
				Amount      any    `json:"amount"`
				UnitPrice   any    `json:"unit_price"`
				Discount    any    `json:"discount"`
				ProductCode any    `json:"product_code"`
				DateItem    any    `json:"date_item"`
				TaxItem     any    `json:"tax_item"`
				TaxRate     any    `json:"tax_rate"`
			} `json:"item_lines"`
		} `json:"extracted_data"`
	} `json:"eden-ai"`
}
type ParseRequest struct {
	FileName           string             `json:"fileName"`
	InvoiceData        string             `json:"invoiceData"`
	BookingResponse    v1.BookingResponse `json:"bookingResponse"`
	AuthContext        model.Session      `json:"authContext"`
	AccountingPlatform string             `json:"accountingPlatform"`
	CustomerID         string             `json:"customerId"`
	PaymentTerm        string             `json:"paymentTerm"`
	PaymentMethod      string             `json:"paymentMethod"`
	IncludeShipping    bool               `json:"includeShipping"`
}
