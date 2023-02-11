/* eslint-disable */
import type { role } from "./role";

export const protobufPackage = "v1";

export interface addStaff {
  token: string;
  roles: role[];
  newStaffEmail: string;
  password: string;
  firstName: string;
  lastName: string;
  phoneNumber: string;
  businessId: string;
}
