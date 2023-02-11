/* eslint-disable */

export const protobufPackage = "v1";

export interface business {
  /** @gotags: dynamodbav:"type" */
  type: string;
  /** @gotags: dynamodbav:"business_name" */
  businessName: string;
  /** @gotags: dynamodbav:"business_id" */
  businessId: string;
  /** @gotags: dynamodbav:"business_email" */
  businessEmail: string;
  /** @gotags: dynamodbav:"accounting_email" */
  accountingEmail: string;
  /** @gotags: dynamodbav:"customer_service_email" */
  customerServiceEmail: string;
  /** @gotags: dynamodbav:"high_priority_email" */
  highPriorityEmail: string;
  /** @gotags: dynamodbav:"avatar_url" */
  avatarUrl: string;
  /** @gotags: dynamodbav:"admin_email" */
  adminEmail: string;
  /** @gotags: dynamodbav:"created_on" */
  createAt: string;
  /** @gotags: dynamodbav:"updated_on" */
  updatedAt: string;
  /** @gotags: dynamodbav:"deleted_on" */
  deletedAt: string;
  /** @gotags: dynamodbav:"phone_number" */
  phoneNumber: string;
  /** @gotags: dynamodbav:"needs_address_update" */
  needsAddressUpdate: boolean;
  /** @gotags: dynamodbav:"needs_default_pickup_address_update" */
  needsDefaultPickupAddressUpdate: boolean;
  /** @gotags: dynamodbav:"needs_default_pickup_address_update" */
  needsDefaultDeliveryAddressUpdate: boolean;
}

export interface businesses {
  /** @gotags: dynamodbav:"businesses" */
  businesses: business[];
}
