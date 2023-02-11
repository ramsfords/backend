/* eslint-disable */

export const protobufPackage = "v1";

export interface reference {
  /** @gotags: dynamodbav:"vendor_bol,omitempty" */
  vendorBol: string;
  /** @gotags: dynamodbav:"bol,omitempty" */
  bol: string;
  /** @gotags: dynamodbav:"vendor_reference_id,omitempty" */
  vendorReferenceId: string;
  /** @gotags: dynamodbav:"pickup_no,omitempty" */
  pickupNo: string;
  /** @gotags: dynamodbav:"invoice_no,omitempty" */
  invoiceNo: string;
  /** @gotags: dynamodbav:"vendor_invoice_no,omitempty" */
  vendorInvoiceNo: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
}
