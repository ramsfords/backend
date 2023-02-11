/* eslint-disable */
import type { freightClass } from "./freight_class";
import type { packageType } from "./package";

export const protobufPackage = "v1";

export enum dimensionUom {
  DIMENSION_NONE = 0,
  INCH = 1,
  CM = 2,
  UNRECOGNIZED = -1,
}

export enum weightUom {
  WEIGHT_NONE = 0,
  LB = 1,
  KG = 2,
  UNRECOGNIZED = -1,
}

export enum commodityServices {
  /** SERVICES_NONE - @gotags: dynamodbav:"NONE,omitempty" */
  SERVICES_NONE = 0,
  /** PROTECT_FROM_FREEZE - @gotags: dynamodbav:"protect_from_freeze,omitempty" */
  PROTECT_FROM_FREEZE = 1,
  /** SORT_AND_SEGREGATE - @gotags: dynamodbav:"sort_and_segregate,omitempty" */
  SORT_AND_SEGREGATE = 2,
  /** GUARANTEED - @gotags: dynamodbav:"guaranteed,omitempty" */
  GUARANTEED = 3,
  /** HAZARDOUS - @gotags: dynamodbav:"hazardous,omitempty" */
  HAZARDOUS = 4,
  /** STACKABLE - @gotags: dynamodbav:"stackabe,omitempty" */
  STACKABLE = 5,
  UNRECOGNIZED = -1,
}

export interface commodity {
  /** @gotags: dynamodbav:"density,omitempty" */
  density: number;
  /** @gotags: dynamodbav:"length,omitempty" */
  length: number;
  /** @gotags: dynamodbav:"width,omitempty" */
  width: number;
  /** @gotags: dynamodbav:"height,omitempty" */
  height: number;
  /** @gotags: dynamodbav:"weight,omitempty" */
  weight: number;
  /** @gotags: dynamodbav:"dimension_uom,omitempty" */
  dimensionUom: dimensionUom;
  /** @gotags: dynamodbav:"weight_uom,omitempty" */
  weightUom: weightUom;
  /** @gotags: dynamodbav:"dimension_display,omitempty" */
  dimensionDisplay: string;
  /** @gotags: dynamodbav:"package_type,omitempty" */
  packageType: packageType;
  /** @gotags: dynamodbav:"quantity,omitempty" */
  quantity: number;
  /** @gotags: dynamodbav:"freight_class,omitempty" */
  freightClass: freightClass;
  /** @gotags: dynamodbav:"stackable,omitempty" */
  stackable: boolean;
  /** @gotags: dynamodbav:"protect_from_freeze,omitempty" */
  protectFromFreeze: boolean;
  /** @gotags: dynamodbav:"sort_and_segregate,omitempty" */
  sortAndSegregate: boolean;
  /** @gotags: dynamodbav:"guaranteed,omitempty" */
  guaranteed: boolean;
  /** @gotags: dynamodbav:"hazardous,omitempty" */
  hazardous: boolean;
  /** @gotags: dynamodbav:"commodity_instructions,omitempty" */
  commodityInstructions: string;
  /** @gotags: dynamodbav:"commodity_services,omitempty" */
  commodityServices: commodityServices[];
  /** @gotags: dynamodbav:"index,omitempty" */
  index: number;
  /** @gotags: dynamodbav:"shipment_description,omitempty" */
  shipmentDescription: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
  /** @gotags: dynamodbav:"quote_id,omitempty" */
  quoteId: string;
}
