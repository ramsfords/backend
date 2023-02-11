/* eslint-disable */
import type { emailTemplate } from "./email_template";

export const protobufPackage = "v1";

export interface sendEmail {
  /** @gotags: dynamodbav:"email_subject,omitempty" */
  emailSubject: string;
  /** @gotags: dynamodbav:"receiver_email,omitempty" */
  receiverEmailAddress: string;
  /** @gotags: dynamodbav:"receiver_name,omitempty" */
  receiverName: string;
  /** @gotags: dynamodbav:"email_template,omitempty" */
  emailPurpose: emailTemplate;
  /** @gotags: dynamodbav:"token,omitempty" */
  token: string;
  /** @gotags: dynamodbav:"success,omitempty" */
  success: boolean;
}
