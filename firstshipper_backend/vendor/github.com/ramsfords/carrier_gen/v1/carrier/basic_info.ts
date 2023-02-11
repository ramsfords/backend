/* eslint-disable */
import type { bookingResponse } from "./book";
import type { business } from "./business";
import type { location } from "./location";
import type { quoteRequest } from "./quote";
import type { user } from "./user";

export const protobufPackage = "v1";

export interface basicInfo {
  token: string;
  business: business | undefined;
  user: user | undefined;
  users: user[];
  valid: boolean;
  clientAddresses: location[];
  businessAddress: location | undefined;
  defaultPickupLocation: location | undefined;
  defaultDeliveryLocation: location | undefined;
  bookings: bookingResponse[];
  quotes: quoteRequest[];
}
