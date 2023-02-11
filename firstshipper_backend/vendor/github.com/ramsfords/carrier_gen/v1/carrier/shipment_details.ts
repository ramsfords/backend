/* eslint-disable */

export const protobufPackage = "v1";

export interface shipmentDetails {
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
  /** @gotags: dynamodbav:"requester_id,omitempty" */
  requesterId: string;
  /** @gotags: dynamodbav:"mode,omitempty" */
  mode: string;
  /** @gotags: dynamodbav:"liable_party_id,omitempty" */
  liablePartyId: string;
  /** @gotags: dynamodbav:"pickup_date,omitempty" */
  pickupDate: string;
  /** @gotags: dynamodbav:"display_date,omitempty" */
  displayDate: string;
  /** @gotags: dynamodbav:"delivery_date,omitempty" */
  deliveryDate: string;
  /** @gotags: dynamodbav:"vendor_bids,omitempty" */
  totalItems: number;
  /** @gotags: dynamodbav:"total_weight,omitempty" */
  totalWeight: number;
  /** @gotags: dynamodbav:"valid_until,omitempty" */
  validUntil: string;
  /** @gotags: dynamodbav:"edit_mode,omitempty" */
  editMode: boolean;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
}
