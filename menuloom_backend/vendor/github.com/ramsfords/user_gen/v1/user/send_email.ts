/* eslint-disable */

export const protobufPackage = "user";

export interface SendEmailData {
  /** @gotags: dynamodbav:"email_subject" */
  emailSubject: string;
  /** @gotags: dynamodbav:"receiver_email" */
  receiverEmailAddress: string;
  /** @gotags: dynamodbav:"receiver_name" */
  receiverName: string;
  /** @gotags: dynamodbav:"email_purpose" */
  emailPurpose: string;
  /** @gotags: dynamodbav:"token" */
  token: string;
  /** @gotags: dynamodbav:"success" */
  success: boolean;
}
