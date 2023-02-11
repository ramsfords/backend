/* eslint-disable */

export const protobufPackage = "v1";

export interface forgotPassword {
  /** @gotags: dynamodbav:"email" */
  email: string;
}

export interface resetPasswordToken {
  /** @gotags: dynamodbav:"issued_to,omitempty" */
  issuedTo: string;
  /** @gotags: dynamodbav:"issued_on,omitempty" */
  issuedOn: string;
  /** @gotags: dynamodbav:"expires_on,omitempty" */
  expiresOn: string;
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
}

export interface resetPassword {
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
  /** @gotags: dynamodbav:"new_password,omitempty" */
  newPassword: string;
  /** @gotags: dynamodbav:"confirm_password,omitempty" */
  confirmPassword: string;
  /** @gotags: dynamodbav:"email,omitempty" */
  email: string;
}
