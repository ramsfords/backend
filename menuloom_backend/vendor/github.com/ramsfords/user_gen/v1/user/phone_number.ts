/* eslint-disable */

export const protobufPackage = "user";

export enum PhoneNumberType {
  Home = 0,
  Office = 1,
  CellPhone = 2,
  UNRECOGNIZED = -1,
}

export interface PhoneNumberData {
  /** @gotags: dynamodbav:"phone_number" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"phone_number_type" */
  type: PhoneNumberType;
}
