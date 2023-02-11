/* eslint-disable */

export const protobufPackage = "v1";

export interface status {
  /** @gotags: dynamodbav:"is_locked,omitempty" */
  isLocked: boolean;
  /** @gotags: dynamodbav:"since,omitempty" */
  since: string;
}
