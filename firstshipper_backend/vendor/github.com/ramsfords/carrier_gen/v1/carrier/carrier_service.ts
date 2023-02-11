/* eslint-disable */

export const protobufPackage = "v1";

export interface Ping {
  /** @gotags: dynamodbav:"hi,omitempty" */
  hi: string;
}

export interface id {
  /** @gotags: dynamodbav:"id,omitempty" */
  id: string;
  /** @gotags: dynamodbav:"business_id,omitempty" */
  businessId: string;
}

export interface ids {
  /** @gotags: dynamodbav:"ids,omitempty" */
  ids: id[];
}
