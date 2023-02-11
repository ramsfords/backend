/* eslint-disable */

export const protobufPackage = "user";

export interface ForgotPasswordData {
  /** @gotags: dynamodbav:"email" */
  email: string;
}

export interface ResetPasswordToken {
  /** @gotags: dynamodbav:"issued_to" */
  issuedTo: string;
  /** @gotags: dynamodbav:"issued_on" */
  issuedOn: string;
  /** @gotags: dynamodbav:"expires_on" */
  expiresOn: string;
  /** @gotags: dynamodbav:"token" */
  token: string;
}

export interface ResetPasswordData {
  /** @gotags: dynamodbav:"token" */
  token: string;
  /** @gotags: dynamodbav:"new_password" */
  newPassword: string;
  /** @gotags: dynamodbav:"confirm_password" */
  confirmPassword: string;
  /** @gotags: dynamodbav:"user_id" */
  userId: string;
}
