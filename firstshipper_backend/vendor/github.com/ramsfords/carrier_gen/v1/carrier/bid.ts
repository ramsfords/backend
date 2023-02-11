/* eslint-disable */
import type { amount } from "./amount";

export const protobufPackage = "v1";

export interface bid {
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
  /** @gotags: dynamodbav:"bid_id,omitempty" */
  bidId: string;
  /** @gotags: dynamodbav:"carrier,omitempty" */
  carrier: string;
  /** @gotags: dynamodbav:"amount,omitempty" */
  amount:
    | amount
    | undefined;
  /** @gotags: dynamodbav:"transit_time,omitempty" */
  transitTime: string;
  /** @gotags: dynamodbav:"guranteed,omitempty" */
  guranteed: boolean;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"company_large_logo_url,omitempty" */
  companyLargeLogoUrl: string;
  /** @gotags: dynamodbav:"company_small_logo_url,omitempty" */
  companySmallLogoUrl: string;
  /** @gotags: dynamodbav:"delivery_date,omitempty" */
  deliveryDate: string;
  /** @gotags: dynamodbav:"vendor_quote_id,omitempty" */
  vendorQuoteId: string;
  /** @gotags: dynamodbav:"capacity_provider_quote_id,omitempty" */
  capacityProviderQuoteId: string;
  /** @gotags: dynamodbav:"vendor_name,omitempty" */
  vendorName: string;
  /** @gotags: dynamodbav:"carrier_name,omitempty" */
  carrierName: string;
  /** @gotags: dynamodbav:"carrier_code,omitempty" */
  carrierCode: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
  /** @gotags: dynamodbav:"carrier_quote_id,omitempty" */
  carrierQuoteId: string;
}

export interface bids {
  /** @gotags: dynamodbav:"bids,omitempty" */
  bids: bid[];
}
