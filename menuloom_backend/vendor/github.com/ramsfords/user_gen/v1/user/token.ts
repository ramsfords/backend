/* eslint-disable */

export const protobufPackage = "user";

export interface RefreshTokenData {
  /**
   * format: JWT
   * @gotags: dynamodbav:"refresh_token"
   */
  refreshToken: string;
}

export interface RefreshToken {
  /**
   * format: JWT
   * @gotags: dynamodbav:"access_token"
   */
  accessToken: string;
  /**
   * format: JWT
   * @gotags: dynamodbav:"refresh_token"
   */
  refreshToken: string;
  /** @gotags: dynamodbav:"token" */
  token: string;
}
