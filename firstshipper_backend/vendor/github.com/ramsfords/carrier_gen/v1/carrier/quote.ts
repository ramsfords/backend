/* eslint-disable */
import type { bid } from "./bid";
import type { commodity } from "./commodity";
import type { location } from "./location";
import type { shipmentDetails } from "./shipment_details";

export const protobufPackage = "v1";

export interface quoteRequest {
  /** @gotags: dynamodbav:"shipment_details,omitempty" */
  shipmentDetails:
    | shipmentDetails
    | undefined;
  /** @gotags: dynamodbav:"commodities,omitempty" */
  commodities: commodity[];
  /** @gotags: dynamodbav:"pickup,omitempty" */
  pickup:
    | location
    | undefined;
  /** @gotags: dynamodbav:"delivery,omitempty" */
  delivery:
    | location
    | undefined;
  /** @gotags: dynamodbav:"bid,omitempty" */
  bid:
    | bid
    | undefined;
  /** @gotags: dynamodbav:"bids,omitempty" */
  bids: bid[];
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
  /** @gotags: dynamodbav:"booking_success,omitempty" */
  bookingSuccess: boolean;
  /** @gotags: dynamodbav:"special_instruction,omitempty" */
  specialInstruction: string;
}

export interface quotesResponse {
  /** @gotags: dynamodbav:"quote_request,omitempty" */
  quoteRequests: quoteRequest[];
}
