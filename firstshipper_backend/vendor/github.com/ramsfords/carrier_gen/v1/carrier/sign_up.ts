/* eslint-disable */
import type { role } from "./role";

export const protobufPackage = "v1";

export interface signUp {
  /** @gotags: dynamodbav:"first_name,omitempty" */
  firstName: string;
  /** @gotags: dynamodbav:"middle_name,omitempty" */
  middleName: string;
  /** @gotags: dynamodbav:"last_name,omitempty" */
  lastName: string;
  /** @gotags: dynamodbav:"email,omitempty" */
  email: string;
  /** @gotags: dynamodbav:"password,omitempty" */
  password: string;
  /** @gotags: dynamodbav:"phone_number,omitempty" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"role,omitempty" */
  role: role[];
  /** @gotags: dynamodbav:"terms_aggrement,omitempty" */
  termsAggrement: boolean;
  /** @gotags: dynamodbav:"company_name,omitempty" */
  companyName: string;
}
