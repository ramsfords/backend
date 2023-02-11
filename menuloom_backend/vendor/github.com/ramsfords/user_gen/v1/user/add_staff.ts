/* eslint-disable */
import type { Role } from "./role";

export const protobufPackage = "user";

export interface AddStaffData {
  /** @gotags: dynamodbav:"token" */
  token: string[];
  /** @gotags: dynamodbav:"roles" */
  roles: Role[];
  /** @gotags: dynamodbav:"new_staff_email" */
  newStaffEmail: string[];
}
