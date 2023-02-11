/* eslint-disable */
import type { Role } from "./role";

export const protobufPackage = "user";

export interface SignUpDatas {
  /** @gotags: dynamodbav:"first_name" */
  firstName: string;
  /** @gotags: dynamodbav:"middle_name" */
  middleName: string;
  /** @gotags: dynamodbav:"last_name" */
  lastName: string;
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"password" */
  password: string;
  /** @gotags: dynamodbav:"phone_number" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"role" */
  role: Role[];
}
