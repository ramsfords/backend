/* eslint-disable */

export const protobufPackage = "restaurant";

export enum RestaurantRole {
  ADMIN = 0,
  USER = 1,
  STAFF = 2,
  CASHIER = 3,
  MANAGER = 4,
  SYSTEM_ADMIN = 5,
  SYSTEM_STAFF = 6,
  SYSTEM_MANAGER = 7,
  UNRECOGNIZED = -1,
}

export interface restaurantAddress {
  /** @gotags: dynamodbav:"street1" */
  street1: string;
  /** @gotags: dynamodbav:"city" */
  city: string;
  /** @gotags: dynamodbav:"state" */
  state: string;
  /** @gotags: dynamodbav:"country" */
  country: string;
  /** @gotags: dynamodbav:"zipCode" */
  zipCode: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
}

export interface openWindow {
  /** @gotags: dynamodbav:"from" */
  from: string;
  /** @gotags: dynamodbav:"to" */
  to: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
}

export interface hours {
  /** @gotags: dynamodbav:"openHours" */
  openHours: openWindow[];
}

export interface CreateRestaurantData {
  /** @gotags: dynamodbav:"id" */
  id: string;
  /** @gotags: dynamodbav:"name" */
  name: string;
  /** @gotags: dynamodbav:"restaurantName" */
  restaurantName: string;
  /** @gotags: dynamodbav:"restaurantWebUrl" */
  restaurantWebUrl: string;
  /** @gotags: dynamodbav:"restaurantS3DevUrl" */
  restaurantS3DevUrl: string;
  /** @gotags: dynamodbav:"restaurantS3StaticProdUrl" */
  restaurantS3StaticProdUrl: string;
  /** @gotags: dynamodbav:"address" */
  address:
    | restaurantAddress
    | undefined;
  /** @gotags: dynamodbav:"phoneNumber" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"ownerId" */
  ownerId: string;
  /** @gotags: dynamodbav:"openHours" */
  openHours: { [key: string]: hours };
  /** @gotags: dynamodbav:"created" */
  created: string;
  /** @gotags: dynamodbav:"updated" */
  updated: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
  /** @gotags: dynamodbav:"pk" */
  pk: string;
  /** @gotags: dynamodbav:"sk" */
  sk: string;
  /** @gotags: dynamodbav:"googleMapViewportUrl" */
  googleMapViewportUrl: string;
  /** @gotags: dynamodbav:"s3StaticImagesUrl" */
  s3StaticImagesUrl: string;
  /** @gotags: dynamodbav:"authEmail" */
  authEmail: string;
  /** @gotags: dynamodbav:"userName" */
  userName: string;
  /** @gotags: dynamodbav:"origin" */
  origin: string;
  /** @gotags: dynamodbav:"role" */
  role: RestaurantRole[];
  /** @gotags: dynamodbav:"tokenKey" */
  tokenKey: string;
  /** @gotags: dynamodbav:"restaurantIds" */
  restaurantIds: string;
}

export interface CreateRestaurantData_OpenHoursEntry {
  key: string;
  value: hours | undefined;
}

export interface restaurantResponse {
  /** @gotags: dynamodbav:"message" */
  message: string;
  /** @gotags: dynamodbav:"id" */
  id: string;
}
