/* eslint-disable */

export const protobufPackage = "v1";

export enum shipmentStatusOption {
  None = 0,
  quote_generated = 1,
  booked = 2,
  pickup_driver_assigned = 3,
  pickup_driver_dispatched = 4,
  piked_up = 5,
  arrived_at_shipping_warehouse = 6,
  arrived_at_transit_warehouse = 7,
  arrived_at_delivery_warehouse = 8,
  document_needed = 9,
  shippers_information_needed = 10,
  receiver_infomration_needed = 11,
  arrived_at_custom = 12,
  released_from_custom = 13,
  on_hold = 14,
  delivery_driver_assigned = 15,
  delivery_driver_despatched = 16,
  arrived_at_delivery_location = 17,
  delivery_refused_by_receiver = 18,
  delivery_rescheduled_by_receiver = 19,
  unable_to_deliver = 20,
  contact_to_schedule_delivery = 21,
  delivery_retried = 22,
  delivered = 23,
  wrong_delivery_address = 24,
  delivery_rescheduled_by_shipper = 25,
  delivery_rescheduled_by_carrier = 26,
  UNRECOGNIZED = -1,
}

export interface shipmentStatus {
  /** @gotags: dynamodbav:"shipment_id,omitempty" */
  shipmentId: string;
  /** @gotags: dynamodbav:"reference_number,omitempty" */
  referenceNumber: string;
  /** @gotags: dynamodbav:"bol_number,omitempty" */
  bolNumber: string;
  /** @gotags: dynamodbav:"date,omitempty" */
  date: string;
  /** @gotags: dynamodbav:"comment,omitempty" */
  comment: string;
  /** @gotags: dynamodbav:"status,omitempty" */
  status: shipmentStatusOption;
  /** @gotags: dynamodbav:"note,omitempty" */
  note: string;
  /** @gotags: dynamodbav:"current_location,omitempty" */
  currentLocation: string;
  /** @gotags: dynamodbav:"intervention_needed,omitempty" */
  interventionNeeded: boolean;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
}
