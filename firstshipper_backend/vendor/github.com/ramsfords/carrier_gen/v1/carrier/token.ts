/* eslint-disable */

export const protobufPackage = "v1";

export interface refreshToken {
  /** @gotags: dynamodbav:"refresh_token,omitempty" */
  refreshToken: string;
}

export interface token {
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
}
