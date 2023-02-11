/* eslint-disable */

export const protobufPackage = "user";

export interface Ok {
  /** @gotags: dynamodbav:"success" */
  success: boolean;
  /** @gotags: dynamodbav:"status_code" */
  statusCode: number;
  /** @gotags: dynamodbav:"message" */
  message: string;
}
