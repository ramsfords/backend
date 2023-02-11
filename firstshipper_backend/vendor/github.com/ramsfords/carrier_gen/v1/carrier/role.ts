/* eslint-disable */

export const protobufPackage = "v1";

export enum role {
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

export interface updateUserRole {
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
  /** @gotags: dynamodbav:"new_role,omitempty" */
  newRole: role[];
  /** @gotags: dynamodbav:"staff_email,omitempty" */
  staffEmail: string;
}
