/* eslint-disable */
import type { bid } from "./bid";
import type { commodity } from "./commodity";
import type { location } from "./location";
import type { shipmentDetails } from "./shipment_details";

export const protobufPackage = "v1";

export enum freightChargeTerm {
  CHARGE_NONE = 0,
  THIRD_PARTY = 1,
  INBOUND_COLLECT = 2,
  PREPAID = 3,
  UNRECOGNIZED = -1,
}

export interface bookingResponse {
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
  /** @gotags: dynamodbav:"booking_success,omitempty" */
  bookingSuccess: boolean;
  /** @gotags: dynamodbav:"rapid_booking_confirmation,omitempty" */
  DispatchResponse:
    | DispatchResponse
    | undefined;
  /** @gotags: dynamodbav:"booking_info,omitempty" */
  bookingInfo:
    | bookingInfo
    | undefined;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
  /** @gotags: dynamodbav:"svg_data,omitempty" */
  svgData: string;
  /** @gotags: dynamodbav:"bol_url,omitempty" */
  bolUrl: string;
}

export interface bookingInfo {
  firstShipperBolNumber: string;
  freightTerm: freightChargeTerm;
  specialInstructions: string;
  carrierName: string;
  carrierProNumber: string;
  carrierLogoUrl: string;
  carrierDispatchContactNumber: string;
  carrierDispatchEmail: string;
  carrierBolNumber: string;
  carrierReference: string;
}

export interface bookings {
  /** @gotags: dynamodbav:"bookings,omitempty" */
  bookingResponses: bookingResponse[];
}

export interface DispatchResponse {
  /** @gotags: dynamodbav:"shipmentId,omitempty" */
  ShipmentID: number;
  /** @gotags: dynamodbav:"securityKey,omitempty" */
  SecurityKey: string;
  /** @gotags: dynamodbav:"pickupNumber,omitempty" */
  PickupNumber: string;
  /** @gotags: dynamodbav:"carrierName,omitempty" */
  CarrierName: string;
  /** @gotags: dynamodbav:"carrierPhone,omitempty" */
  CarrierPhone: string;
  /** @gotags: dynamodbav:"carrierPRONumber,omitempty" */
  CarrierPRONumber: string;
  /** @gotags: dynamodbav:"handlingUnitTotal,omitempty" */
  HandlingUnitTotal: number;
  /** @gotags: dynamodbav:"isShipmentEdit,omitempty" */
  IsShipmentEdit: boolean;
  /** @gotags: dynamodbav:"isShipmentManual,omitempty" */
  IsShipmentManual: boolean;
  /** @gotags: dynamodbav:"serviceType,omitempty" */
  ServiceType: number;
  /** @gotags: dynamodbav:"isTrackingEmailSend,omitempty" */
  IsTrackingEmailSend: boolean;
  /** @gotags: dynamodbav:"isTrackingAPIEnabled,omitempty" */
  IsTrackingAPIEnabled: boolean;
  /** @gotags: dynamodbav:"customerBOLNumber,omitempty" */
  CustomerBOLNumber: string;
  /** @gotags: dynamodbav:"shipperEmail,omitempty" */
  ShipperEmail: string;
  /** @gotags: dynamodbav:"consigneeEmail,omitempty" */
  ConsigneeEmail: string;
  /** @gotags: dynamodbav:"result,omitempty" */
  Result: Result | undefined;
}

export interface Result {
  /** @gotags: dynamodbav:"capacityProviderBolUrl,omitempty" */
  CapacityProviderBolUrl: string;
  /** @gotags: dynamodbav:"shipmentIdentifier,omitempty" */
  ShipmentIdentifier: string;
  /** @gotags: dynamodbav:"pickupNote,omitempty" */
  PickupNote: string;
  /** @gotags: dynamodbav:"pickupDateTime,omitempty" */
  PickupDateTime: string;
  /** @gotags: dynamodbav:"errors,omitempty" */
  Errors: string[];
  /** @gotags: dynamodbav:"infoMessages,omitempty" */
  InfoMessages: string[];
  /** @gotags: dynamodbav:"type,omitempty" */
  Type: string;
}
