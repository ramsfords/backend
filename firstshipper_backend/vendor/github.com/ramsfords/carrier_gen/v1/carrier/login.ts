/* eslint-disable */

export const protobufPackage = "v1";

export interface login {
  /** @gotags: dynamodbav:"email,omitempty" */
  email: string;
  /** @gotags: dynamodbav:"password,omitempty" */
  password: string;
}
