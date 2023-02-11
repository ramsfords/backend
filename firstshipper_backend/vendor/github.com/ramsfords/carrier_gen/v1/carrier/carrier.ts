/* eslint-disable */

export const protobufPackage = "v1";

export interface carrier {
  /** @gotags: dynamodbav:"name,omitempty" */
  name: string;
  /** @gotags: dynamodbav:"auth_url,omitempty" */
  authUrl: string;
  /** @gotags: dynamodbav:"rate_url,omitempty" */
  rateUrl: string;
  /** @gotags: dynamodbav:"add_address_url,omitempty" */
  addAddressUrl: string;
  /** @gotags: dynamodbav:"get_address_url,omitempty" */
  getAddressUrl: string;
  /** @gotags: dynamodbav:"quote_history_url,omitempty" */
  quoteHistoryUrl: string;
}
