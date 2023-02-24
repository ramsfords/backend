package books

import (
	"fmt"
	"strings"

	"github.com/ramsfords/backend/foundations/zoho"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Address struct {
	Attention string `json:"attention,omitempty"`
	Address   string `json:"address,omitempty"`
	Street2   string `json:"street2,omitempty"`
	StateCode string `json:"state_code,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
	Fax       string `json:"fax,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type Tag struct {
	TagID       string `json:"tag_id,omitempty"`
	TagOptionID string `json:"tag_option_id,omitempty"`
}
type DefaultTemplates struct {
	InvoiceTemplateID                             string `json:"invoice_template_id,omitempty"`
	EstimateTemplateID                            string `json:"estimate_template_id,omitempty"`
	CreditnoteTemplateID                          string `json:"creditnote_template_id,omitempty"`
	PurchaseorderTemplateID                       string `json:"purchaseorder_template_id,omitempty"`
	SalesorderTemplateID                          string `json:"salesorder_template_id,omitempty"`
	RetainerinvoiceTemplateID                     string `json:"retainerinvoice_template_id,omitempty"`
	PaymentthankyouTemplateID                     string `json:"paymentthankyou_template_id,omitempty"`
	RetainerinvoicePaymentthankyouTemplateID      string `json:"retainerinvoice_paymentthankyou_template_id,omitempty"`
	InvoiceEmailTemplateID                        string `json:"invoice_email_template_id,omitempty"`
	EstimateEmailTemplateID                       string `json:"estimate_email_template_id,omitempty"`
	CreditnoteEmailTemplateID                     string `json:"creditnote_email_template_id,omitempty"`
	PurchaseorderEmailTemplateID                  string `json:"purchaseorder_email_template_id,omitempty"`
	SalesorderEmailTemplateID                     string `json:"salesorder_email_template_id,omitempty"`
	RetainerinvoiceEmailTemplateID                string `json:"retainerinvoice_email_template_id,omitempty"`
	PaymentthankyouEmailTemplateID                string `json:"paymentthankyou_email_template_id,omitempty"`
	RetainerinvoicePaymentthankyouEmailTemplateID string `json:"retainerinvoice_paymentthankyou_email_template_id,omitempty"`
}
type CustomField struct {
	Index int    `json:"index,omitempty"`
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}
type Contact struct {
	ContactID                        string           `json:"contact_id,omitempty"`
	ContactName                      string           `json:"contact_name,omitempty"`
	CompanyName                      string           `json:"company_name,omitempty"`
	HasTransaction                   bool             `json:"has_transaction,omitempty"`
	ContactType                      string           `json:"contact_type,omitempty"`
	CustomerSubType                  string           `json:"customer_sub_type,omitempty"`
	CreditLimit                      string           `json:"credit_limit,omitempty"`
	IsPortalEnabled                  bool             `json:"is_portal_enabled,omitempty"`
	LanguageCode                     string           `json:"language_code,omitempty"`
	IsTaxable                        bool             `json:"is_taxable,omitempty"`
	TaxID                            string           `json:"tax_id,omitempty"`
	TaxName                          string           `json:"tax_name,omitempty"`
	TaxPercentage                    string           `json:"tax_percentage,omitempty"`
	TaxAuthorityID                   string           `json:"tax_authority_id,omitempty"`
	TaxExemptionID                   string           `json:"tax_exemption_id,omitempty"`
	TaxAuthorityName                 string           `json:"tax_authority_name,omitempty"`
	TaxExemptionCode                 string           `json:"tax_exemption_code,omitempty"`
	PlaceOfContact                   string           `json:"place_of_contact,omitempty"`
	GstNo                            string           `json:"gst_no,omitempty"`
	VatTreatment                     string           `json:"vat_treatment,omitempty"`
	TaxTreatment                     string           `json:"tax_treatment,omitempty"`
	GstTreatment                     string           `json:"gst_treatment,omitempty"`
	IsLinkedWithZohocrm              bool             `json:"is_linked_with_zohocrm,omitempty"`
	Website                          string           `json:"website,omitempty"`
	OwnerID                          string           `json:"owner_id,omitempty"`
	PrimaryContactID                 string           `json:"primary_contact_id,omitempty"`
	PaymentTerms                     int              `json:"payment_terms,omitempty"`
	PaymentTermsLabel                string           `json:"payment_terms_label,omitempty"`
	CurrencyID                       string           `json:"currency_id,omitempty"`
	CurrencyCode                     string           `json:"currency_code,omitempty"`
	CurrencySymbol                   string           `json:"currency_symbol,omitempty"`
	OpeningBalanceAmount             float32          `json:"opening_balance_amount,omitempty"`
	ExchangeRate                     string           `json:"exchange_rate,omitempty"`
	OutstandingReceivableAmount      float32          `json:"outstanding_receivable_amount,omitempty"`
	OutstandingReceivableAmountBcy   float32          `json:"outstanding_receivable_amount_bcy,omitempty"`
	UnusedCreditsReceivableAmount    float32          `json:"unused_credits_receivable_amount,omitempty"`
	UnusedCreditsReceivableAmountBcy float32          `json:"unused_credits_receivable_amount_bcy,omitempty"`
	Status                           string           `json:"status,omitempty"`
	PaymentReminderEnabled           bool             `json:"payment_reminder_enabled,omitempty"`
	CustomFields                     []CustomField    `json:"custom_fields,omitempty"`
	BillingAddress                   Address          `json:"billing_address,omitempty"`
	ShippingAddress                  Address          `json:"shipping_address,omitempty"`
	Facebook                         string           `json:"facebook,omitempty"`
	Twitter                          string           `json:"twitter,omitempty"`
	ContactPersons                   []ContactPerson  `json:"contact_persons,omitempty"`
	DefaultTemplates                 DefaultTemplates `json:"default_templates,omitempty"`
	Notes                            string           `json:"notes,omitempty"`
	CreatedTime                      string           `json:"created_time,omitempty"`
	LastModifiedTime                 string           `json:"last_modified_time,omitempty"`
}
type ContactPerson struct {
	ContactPersonID  string `json:"contact_person_id,omitempty"`
	Salutation       string `json:"salutation,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
	Designation      string `json:"designation,omitempty"`
	Department       string `json:"department,omitempty"`
	Skype            string `json:"skype,omitempty"`
	IsPrimaryContact bool   `json:"is_primary_contact,omitempty"`
	EnablePortal     bool   `json:"enable_portal,omitempty"`
}
type ContactResponse struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Contact Contact `json:"contact,omitempty"`
}

func (api API) CreateContact(contact Contact) (Contact, error) {
	var responseData ContactResponse
	endPoint := &zoho.Endpoint{
		Name: "books",
		URL:  "https://books.zoho.com/api/v3/contacts",
		URLParameters: map[string]zoho.Parameter{
			"organization_id": zoho.Parameter(api.id),
		},
		Method:       zoho.HTTPPost,
		ResponseData: &responseData,
		RequestBody:  contact,
		BodyFormat:   zoho.JSON,
	}
	err := api.Zoho.HTTPRequest(endPoint)
	if err == nil {
		contactData, ok := endPoint.ResponseData.(*ContactResponse)
		if ok {
			return contactData.Contact, nil
		}
		if !ok {
			return Contact{}, err
		}

	} else {
		return Contact{}, err
	}
	return Contact{}, err
}
func NewContact(business *v1.Business) Contact {
	return Contact{
		ContactName:       business.AdminUser.Name,
		CompanyName:       business.BusinessName,
		ContactType:       "customer",
		CustomerSubType:   "business",
		PaymentTerms:      15,
		PaymentTermsLabel: "Net 15",
		BillingAddress: Address{
			Attention: business.AdminUser.Name,
			Address:   business.Address.AddressLine1,
			Street2:   business.Address.AddressLine2,
			StateCode: business.Address.StateCode,
			City:      business.Address.City,
			State:     business.Address.State,
			Zip:       business.Address.ZipCode,
			Country:   "U.S.A",
			Phone:     business.PhoneNumber.PhoneNumber,
		},
		ShippingAddress: Address{
			Attention: business.AdminUser.Name,
			Address:   business.Address.AddressLine1,
			Street2:   business.Address.AddressLine2,
			StateCode: business.Address.StateCode,
			City:      business.Address.City,
			State:     business.Address.State,
			Zip:       business.Address.ZipCode,
			Country:   "U.S.A",
			Phone:     business.PhoneNumber.PhoneNumber,
		},
		ContactPersons: []ContactPerson{
			{
				FirstName:        strings.Split(business.AdminUser.Name, " ")[0],
				LastName:         strings.Split(business.AdminUser.Name, " ")[1],
				Email:            business.AdminUser.Email,
				Phone:            business.PhoneNumber.PhoneNumber,
				Mobile:           business.PhoneNumber.PhoneNumber,
				IsPrimaryContact: true,
			},
		},
	}
}
func InsertProspects(reciever *v1.Location, api API) error {
	propect := Contact{
		ContactName:       reciever.Contact.Name,
		CompanyName:       reciever.CompanyName,
		ContactType:       "customer",
		CustomerSubType:   "business",
		PaymentTerms:      15,
		PaymentTermsLabel: "Net 15",
		BillingAddress: Address{
			Attention: reciever.Contact.Name,
			Address:   reciever.Address.AddressLine1,
			Street2:   reciever.Address.AddressLine2,
			StateCode: reciever.Address.StateCode,
			City:      reciever.Address.City,
			State:     reciever.Address.State,
			Zip:       reciever.Address.ZipCode,
			Country:   "U.S.A",
			Phone:     reciever.Contact.PhoneNumber,
		},
		ShippingAddress: Address{
			Attention: reciever.Contact.Name,
			Address:   reciever.Address.AddressLine1,
			Street2:   reciever.Address.AddressLine2,
			StateCode: reciever.Address.StateCode,
			City:      reciever.Address.City,
			State:     reciever.Address.State,
			Zip:       reciever.Address.ZipCode,
			Country:   "U.S.A",
			Phone:     reciever.Contact.PhoneNumber,
		},
		ContactPersons: []ContactPerson{
			{
				FirstName:        strings.Split(reciever.Contact.Name, " ")[0],
				LastName:         strings.Split(reciever.Contact.Name, " ")[1],
				Email:            reciever.Contact.EmailAddress,
				Phone:            reciever.Contact.PhoneNumber,
				Mobile:           reciever.Contact.PhoneNumber,
				IsPrimaryContact: true,
			},
		},
	}
	_, err := api.CreateContact(propect)
	if err != nil {
		fmt.Println("could not insert prospect")
	}
	return err
}
