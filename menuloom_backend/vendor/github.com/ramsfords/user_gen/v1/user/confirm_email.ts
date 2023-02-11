/* eslint-disable */

export const protobufPackage = "user";

/**
 * ClientId:         &data.ClientId,
 * Username:         &data.UserName,
 * ConfirmationCode: &data.ConfirmationCode,
 */
export interface ConfirmEmailData {
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"user_name" */
  userName: string;
  /** @gotags: dynamodbav:"confirmation_code" */
  confirmationCode: string;
  /** @gotags: dynamodbav:"token" */
  token: string;
}
