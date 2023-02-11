/* eslint-disable */

export const protobufPackage = "v1";

/** @gotags: dynamodbav:"amount,omitempty" */
export interface amount {
  /** @gotags: dynamodbav:"full_amount,omitempty" */
  fullAmount: number;
  /** @gotags: dynamodbav:"discount_applied,omitempty" */
  discountApplied: number;
  /** @gotags: dynamodbav:"net_amount,omitempty" */
  netAmount: number;
}
