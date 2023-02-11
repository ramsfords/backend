/* eslint-disable */

export const protobufPackage = "v1";

export interface userHome {
  /** @gotags: dynamodbav:"user_id,omitempty" */
  userId: string;
  /** @gotags: dynamodbav:"user_email,omitempty" */
  userEmail: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
}
