/* eslint-disable */

export const protobufPackage = "v1";

export enum pickupLocationServices {
  /** PICKUP_LOCATION_NONE - @gotags: dynamodbav:"pickup_location_none,omitempty" */
  PICKUP_LOCATION_NONE = 0,
  /** PICKUP_LOCATION_WITH_DOCK - @gotags: dynamodbav:"pickup_location_with_dock,omitempty" */
  PICKUP_LOCATION_WITH_DOCK = 1,
  /** LIFTGATE_PICKUP - @gotags: dynamodbav:"liftgate_pickup,omitempty" */
  LIFTGATE_PICKUP = 2,
  /** PICKUP_APPOINTMENT - @gotags: dynamodbav:"pickup_appointment,omitempty" */
  PICKUP_APPOINTMENT = 3,
  /** INSIDE_PICKUP - @gotags: dynamodbav:"inside_pickup,omitempty" */
  INSIDE_PICKUP = 4,
  /** PICKUP_NOTIFICATION - @gotags: dynamodbav:"pickup_notification,omitempty" */
  PICKUP_NOTIFICATION = 5,
  /** SHIPPER_DELIVERY_NOTIFICATION - @gotags: dynamodbav:"shipper_delivery_notification,omitempty" */
  SHIPPER_DELIVERY_NOTIFICATION = 6,
  UNRECOGNIZED = -1,
}

export enum deliveryLocationServices {
  /** DELIVERY_LOCATION_NONE - @gotags: dynamodbav:"delivery_location_none,omitempty" */
  DELIVERY_LOCATION_NONE = 0,
  /** DELIVERY_LOCATION_WITH_DOCK - @gotags: dynamodbav:"delivery_location_with_dock,omitempty" */
  DELIVERY_LOCATION_WITH_DOCK = 1,
  /** LIFTGATE_DELIVERY - @gotags: dynamodbav:"liftgate_delivery,omitempty" */
  LIFTGATE_DELIVERY = 2,
  /** DELIVERY_APPOINTMENT - @gotags: dynamodbav:"delivery_appointment,omitempty" */
  DELIVERY_APPOINTMENT = 3,
  /** INSIDE_DELIVERY - @gotags: dynamodbav:"inside_delivery,omitempty" */
  INSIDE_DELIVERY = 4,
  /** RECEIVER_PICKUP_NOTIFICATION - @gotags: dynamodbav:"receiver_pickup_notification,omitempty" */
  RECEIVER_PICKUP_NOTIFICATION = 5,
  /** DELIVERY_NOTIFICATION - @gotags: dynamodbav:"delivery_notification,omitempty" */
  DELIVERY_NOTIFICATION = 6,
  UNRECOGNIZED = -1,
}
