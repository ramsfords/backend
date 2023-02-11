/* eslint-disable */
import type { role } from "./role";

export const protobufPackage = "v1";

export interface contact {
  /** @gotags: dynamodbav:"first_name,omitempty" */
  firstName: string;
  /** @gotags: dynamodbav:"last_name,omitempty" */
  lastName: string;
  /** @gotags: dynamodbav:"phone_number,omitempty" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"phone_number_extension,omitempty" */
  phoneNumberExtension: string;
  /** @gotags: dynamodbav:"email_address,omitempty" */
  emailAddress: string;
  /** @gotags: dynamodbav:"phone_number_display,omitempty" */
  phoneNumberDisplay: string;
  /** @gotags: dynamodbav:"roles,omitempty" */
  roles: role[];
  /** @gotags: dynamodbav:"preffered_contact_method,omitempty" */
  prefferedContactMethod: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
}
