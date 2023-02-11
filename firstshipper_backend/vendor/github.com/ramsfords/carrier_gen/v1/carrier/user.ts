/* eslint-disable */
import type { role } from "./role";

export const protobufPackage = "v1";

export interface user {
  username: string;
  password: string;
  confirmPassword: string;
  email: string;
  name: string;
  origin: string;
  role: role[];
  emailVisibility: boolean;
  restaurantIds: string;
  type: string;
}

export interface LoginResponse {
  message: string;
  id: string;
}

export interface LoginUser {
  email: string;
  password: string;
}
