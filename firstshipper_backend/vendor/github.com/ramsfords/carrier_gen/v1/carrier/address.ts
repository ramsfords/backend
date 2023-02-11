/* eslint-disable */

export const protobufPackage = "v1";

export interface address {
  /** @gotags: dynamodbav:"address_line_1,omitempty" */
  addressLine1: string;
  /** @gotags: dynamodbav:"address_line_2,omitempty" */
  addressLine2: string;
  /** @gotags: dynamodbav:"street_name,omitempty" */
  streetName: string;
  /** @gotags: dynamodbav:"city,omitempty" */
  city: string;
  /** @gotags: dynamodbav:"county,omitempty" */
  county: string;
  /** @gotags: dynamodbav:"zip_code,omitempty" */
  zipCode: string;
  /** @gotags: dynamodbav:"state,omitempty" */
  state: string;
  /** @gotags: dynamodbav:"state_code,omitempty" */
  stateCode: string;
  /** @gotags: dynamodbav:"extended_zip_code,omitempty" */
  extendedZipCode: string;
  /** @gotags: dynamodbav:"country,omitempty" */
  country: string;
  /** @gotags: dynamodbav:"country_code,omitempty" */
  countryCode: string;
  /** @gotags: dynamodbav:"free_form_address,omitempty" */
  freeFormAddress: string;
  /** @gotags: dynamodbav:"local_name,omitempty" */
  localName: string;
  /** @gotags: dynamodbav:"lat,omitempty" */
  lat: number;
  /** @gotags: dynamodbav:"long,omitempty" */
  long: number;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
}
