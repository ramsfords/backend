/* eslint-disable */
import type { PhoneNumberData } from "./phone_number";
import type { Role } from "./role";

export const protobufPackage = "user";

export interface UserData {
  /** @gotags: dynamodbav:"type" */
  Type: string;
  /** @gotags: dynamodbav:"cognito_id" */
  cognitoId: string;
  /** @gotags: dynamodbav:"user_name" */
  userName: string;
  /** @gotags: dynamodbav:"user_id" */
  userId: string;
  /** @gotags: dynamodbav:"first_name" */
  firstName: string;
  /** @gotags: dynamodbav:"middle_name" */
  middleName: string;
  /** @gotags: dynamodbav:"last_name" */
  lastName: string;
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"hashed_password" */
  hashedPassword: string;
  /** @gotags: dynamodbav:"avatar_url" */
  avatarUrl: string;
  /** @gotags: dynamodbav:"roles" */
  roles: Role[];
  /** @gotags: dynamodbav:"new_password_required" */
  newPasswordRequired: boolean;
  /** @gotags: dynamodbav:"password_changed_at" */
  passwordChangedAt: string;
  /** @gotags: dynamodbav:"created_on" */
  createdOn: string;
  /** @gotags: dynamodbav:"updated_on" */
  updatedOn: string;
  /** @gotags: dynamodbav:"deleted_on" */
  deletedOn: string;
  /** @gotags: dynamodbav:"phone_numbers" */
  phoneNumbers: PhoneNumberData[];
  /** @gotags: dynamodbav:"email_verified" */
  emailVerified: boolean;
  /** @gotags: dynamodbav:"reset_password_token" */
  resetPasswordToken: string[];
  /** @gotags: dynamodbav:"sessions" */
  sessions: string[];
  /** @gotags: dynamodbav:"sk" */
  pk: string;
  /** @gotags: dynamodbav:"pk" */
  sk: string;
  /** @gotags: dynamodbav:"business_ids" */
  BusinessIds: string[];
  /** @gotags: dynamodbav:"unsuscribed_to_marketing_email" */
  unsuscribedToMarketingEmail: boolean;
}

export interface MeData {
  /** @gotags: dynamodbav:"token" */
  token: string[];
}

export interface SignUpData {
  username: string;
  password: string;
  confirmPassword: string;
  email: string;
  name: string;
  origin: string;
  role: Role[];
  emailVisibility: boolean;
  restaurantIds: string;
  type: string;
}

export interface userResponse {
  message: string;
  id: string;
}

export interface Login {
  email: string;
  password: string;
}

export interface ResetPassword {
  /** @gotags: dynamodbav:"token" */
  token: string;
  /** @gotags: dynamodbav:"password" */
  password: string;
  /** @gotags: dynamodbav:"confirmPassword" */
  confirmPassword: string;
}

export interface NonAuthUser {
  /** @gotags: dynamodbav:"email" */
  email: string;
  /** @gotags: dynamodbav:"name" */
  name: string;
  /** @gotags: dynamodbav:"origin" */
  origin: string;
  /** @gotags: dynamodbav:"role" */
  role: Role[];
  /** @gotags: dynamodbav:"tokenKey" */
  tokenKey: string;
  /** @gotags: dynamodbav:"restaurantIds" */
  restaurantIds: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
}
