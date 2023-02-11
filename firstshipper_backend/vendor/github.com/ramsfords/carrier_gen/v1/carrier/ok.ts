/* eslint-disable */

export const protobufPackage = "v1";

export interface ok {
  /** @gotags: dynamodbav:"success,omitempty" */
  success: boolean;
  /** @gotags: dynamodbav:"status_code,omitempty" */
  statusCode: number;
  /** @gotags: dynamodbav:"message,omitempty" */
  message: string;
}
