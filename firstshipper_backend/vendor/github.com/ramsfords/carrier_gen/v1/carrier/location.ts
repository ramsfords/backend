/* eslint-disable */
import type { address } from "./address";
import type { businessHours } from "./business_hours";
import type { contact } from "./contact";
import type { deliveryLocationServices, pickupLocationServices } from "./location_services";

export const protobufPackage = "v1";

export interface location {
  /** @gotags: dynamodbav:"id,omitempty" */
  id: string;
  /** @gotags: dynamodbav:"company_name,omitempty" */
  companyName: string;
  /** @gotags: dynamodbav:"address,omitempty" */
  address:
    | address
    | undefined;
  /** @gotags: dynamodbav:"contact,omitempty" */
  contact:
    | contact
    | undefined;
  /** @gotags: dynamodbav:"business_hours,omitempty", */
  businessHours:
    | businessHours
    | undefined;
  /** @gotags: dynamodbav:"shipper_pickup_ready_by,omitempty" */
  shipperPickupReadyBy: string;
  /** @gotags: dynamodbav:"shipper_instructions,omitempty" */
  shipperInstructions: string;
  /** @gotags: dynamodbav:"receiver_instructions,omitempty" */
  receiverInstructions: string;
  /** @gotags: dynamodbav:"delivery_location_services,omitempty" */
  deliveryLocationServices: deliveryLocationServices[];
  /** @gotags: dynamodbav:"pickup_location_services,omitempty" */
  pickupLocationServices: pickupLocationServices[];
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
}

export interface locations {
  /** @gotags: dynamodbav:"locations,omitempty" */
  Locations: location[];
}
