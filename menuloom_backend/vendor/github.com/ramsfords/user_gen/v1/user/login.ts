/* eslint-disable */

export const protobufPackage = "user";

export interface LoginUserData {
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"password" */
  password: string;
}
