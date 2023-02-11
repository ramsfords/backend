/* eslint-disable */

export const protobufPackage = "v1";

export interface confirmEmail {
  /** @gotags: dynamodbav:"email,omitempty" */
  email: string;
  /** @gotags: dynamodbav:"user_name,omitempty" */
  userName: string;
  /** @gotags: dynamodbav:"confirmation_code,omitempty" */
  confirmationCode: string;
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
}
