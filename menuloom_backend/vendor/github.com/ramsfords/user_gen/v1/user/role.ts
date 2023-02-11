/* eslint-disable */

export const protobufPackage = "user";

export enum Role {
  SystemOwner = 0,
  SystemManager = 1,
  SystemStaff = 2,
  SystemDeveloper = 3,
  SystemWarehouseManager = 4,
  SystemWarehouseStaff = 5,
  BusinessOwner = 6,
  BusinessManager = 7,
  BusinessStaff = 8,
  BusinessWarehouseManager = 9,
  BusinessWarehouseStaff = 10,
  ThreePLOwner = 11,
  ThreePLStaff = 12,
  UNRECOGNIZED = -1,
}

export interface UpdateUserRoleData {
  /** @gotags: dynamodbav:"token" */
  token: string;
  /** @gotags: dynamodbav:"new_role" */
  newRole: Role[];
}
