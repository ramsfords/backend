// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: bid.proto

package v1

import (
	_ "github.com/ramsfords/types_gen/v1/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bidId,omitempty"
	BidId string `protobuf:"bytes,1,opt,name=bidId,proto3" json:"bidId,omitempty" dynamodbav:"bidId,omitempty"`
	// @gotags: dynamodbav:"carrier,omitempty"
	Carrier string `protobuf:"bytes,2,opt,name=carrier,proto3" json:"carrier,omitempty" dynamodbav:"carrier,omitempty"`
	// @gotags: dynamodbav:"amount,omitempty"
	Amount *Amount `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty" dynamodbav:"amount,omitempty"`
	// @gotags: dynamodbav:"transitTime,omitempty"
	TransitTime string `protobuf:"bytes,5,opt,name=transitTime,proto3" json:"transitTime,omitempty" dynamodbav:"transitTime,omitempty"`
	// @gotags: dynamodbav:"guranteed,omitempty"
	Guranteed bool `protobuf:"varint,6,opt,name=guranteed,proto3" json:"guranteed,omitempty" dynamodbav:"guranteed,omitempty"`
	// @gotags: dynamodbav:"vendorLogo,omitempty"
	VendorLogo string `protobuf:"bytes,7,opt,name=vendorLogo,proto3" json:"vendorLogo,omitempty" dynamodbav:"vendorLogo,omitempty"`
	// @gotags: dynamodbav:"deliveryDate,omitempty"
	DeliveryDate string `protobuf:"bytes,10,opt,name=deliveryDate,proto3" json:"deliveryDate,omitempty" dynamodbav:"deliveryDate,omitempty"`
	// @gotags: dynamodbav:"vendorQuoteId,omitempty"
	VendorQuoteId string `protobuf:"bytes,11,opt,name=vendorQuoteId,proto3" json:"vendorQuoteId,omitempty" dynamodbav:"vendorQuoteId,omitempty"`
	// @gotags: dynamodbav:"capacityProviderQuoteId,omitempty"
	CapacityProviderQuoteId string `protobuf:"bytes,12,opt,name=capacityProviderQuoteId,proto3" json:"capacityProviderQuoteId,omitempty" dynamodbav:"capacityProviderQuoteId,omitempty"`
	// @gotags: dynamodbav:"vendorName,omitempty"
	VendorName string `protobuf:"bytes,13,opt,name=vendorName,proto3" json:"vendorName,omitempty" dynamodbav:"vendorName,omitempty"`
	// @gotags: dynamodbav:"carrierName,omitempty"
	CarrierName string `protobuf:"bytes,14,opt,name=carrierName,proto3" json:"carrierName,omitempty" dynamodbav:"carrierName,omitempty"`
	// @gotags: dynamodbav:"carrierCode,omitempty"
	CarrierCode string `protobuf:"bytes,15,opt,name=carrierCode,proto3" json:"carrierCode,omitempty" dynamodbav:"carrierCode,omitempty"`
	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,16,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type,omitempty"`
	// @gotags: dynamodbav:"carrierQuoteId,omitempty"
	CarrierQuoteId string `protobuf:"bytes,17,opt,name=carrierQuoteId,proto3" json:"carrierQuoteId,omitempty" dynamodbav:"carrierQuoteId,omitempty"`
	// @gotags: dynamodbav:"quoteId,omitempty"
	QuoteId string `protobuf:"bytes,18,opt,name=quoteId,proto3" json:"quoteId,omitempty" dynamodbav:"quoteId,omitempty"`
	// @gotags: dynamodbav:"carrierID,omitempty"
	CarrierID int64 `protobuf:"varint,19,opt,name=carrierID,proto3" json:"carrierID,omitempty" dynamodbav:"carrierID,omitempty"`
	// @gotags: dynamodbav:"destination,omitempty"
	Destination string `protobuf:"bytes,20,opt,name=destination,proto3" json:"destination,omitempty" dynamodbav:"destination,omitempty"`
	// @gotags: dynamodbav:"origin,omitempty"
	Origin string `protobuf:"bytes,21,opt,name=origin,proto3" json:"origin,omitempty" dynamodbav:"origin,omitempty"`
	// @gotags: dynamodbav:"opportunityId,omitempty"
	OpportunityId int64 `protobuf:"varint,22,opt,name=opportunityId,proto3" json:"opportunityId,omitempty" dynamodbav:"opportunityId,omitempty"`
	// @gotags: dynamodbav:"serviceLevelCode,omitempty"
	ServiceLevelCode string `protobuf:"bytes,23,opt,name=serviceLevelCode,proto3" json:"serviceLevelCode,omitempty" dynamodbav:"serviceLevelCode,omitempty"`
	// @gotags: dynamodbav:"serviceName,omitempty"
	ServiceName string `protobuf:"bytes,24,opt,name=serviceName,proto3" json:"serviceName,omitempty" dynamodbav:"serviceName,omitempty"`
	// @gotags: dynamodbav:"serviceType,omitempty"
	ServiceType int64 `protobuf:"varint,25,opt,name=serviceType,proto3" json:"serviceType,omitempty" dynamodbav:"serviceType,omitempty"`
	// @gotags: dynamodbav:"shipmentId,omitempty"
	ShipmentId string `protobuf:"bytes,26,opt,name=shipmentId,proto3" json:"shipmentId,omitempty" dynamodbav:"shipmentId,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetBidId() string {
	if x != nil {
		return x.BidId
	}
	return ""
}

func (x *Bid) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *Bid) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Bid) GetTransitTime() string {
	if x != nil {
		return x.TransitTime
	}
	return ""
}

func (x *Bid) GetGuranteed() bool {
	if x != nil {
		return x.Guranteed
	}
	return false
}

func (x *Bid) GetVendorLogo() string {
	if x != nil {
		return x.VendorLogo
	}
	return ""
}

func (x *Bid) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *Bid) GetVendorQuoteId() string {
	if x != nil {
		return x.VendorQuoteId
	}
	return ""
}

func (x *Bid) GetCapacityProviderQuoteId() string {
	if x != nil {
		return x.CapacityProviderQuoteId
	}
	return ""
}

func (x *Bid) GetVendorName() string {
	if x != nil {
		return x.VendorName
	}
	return ""
}

func (x *Bid) GetCarrierName() string {
	if x != nil {
		return x.CarrierName
	}
	return ""
}

func (x *Bid) GetCarrierCode() string {
	if x != nil {
		return x.CarrierCode
	}
	return ""
}

func (x *Bid) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Bid) GetCarrierQuoteId() string {
	if x != nil {
		return x.CarrierQuoteId
	}
	return ""
}

func (x *Bid) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *Bid) GetCarrierID() int64 {
	if x != nil {
		return x.CarrierID
	}
	return 0
}

func (x *Bid) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Bid) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Bid) GetOpportunityId() int64 {
	if x != nil {
		return x.OpportunityId
	}
	return 0
}

func (x *Bid) GetServiceLevelCode() string {
	if x != nil {
		return x.ServiceLevelCode
	}
	return ""
}

func (x *Bid) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Bid) GetServiceType() int64 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *Bid) GetShipmentId() string {
	if x != nil {
		return x.ShipmentId
	}
	return ""
}

type Bids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bids,omitempty"
	Bids []*Bid `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids,omitempty" dynamodbav:"bids,omitempty"`
}

func (x *Bids) Reset() {
	*x = Bids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bids) ProtoMessage() {}

func (x *Bids) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bids.ProtoReflect.Descriptor instead.
func (*Bids) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{1}
}

func (x *Bids) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_bid_proto protoreflect.FileDescriptor

var file_bid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x06, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69,
	0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x23, 0x0a,
	0x04, 0x62, 0x69, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x04, 0x62, 0x69,
	0x64, 0x73, 0x42, 0x5d, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x42, 0x69,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58,
	0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bid_proto_rawDescOnce sync.Once
	file_bid_proto_rawDescData = file_bid_proto_rawDesc
)

func file_bid_proto_rawDescGZIP() []byte {
	file_bid_proto_rawDescOnce.Do(func() {
		file_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bid_proto_rawDescData)
	})
	return file_bid_proto_rawDescData
}

var file_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bid_proto_goTypes = []interface{}{
	(*Bid)(nil),    // 0: v1.bid
	(*Bids)(nil),   // 1: v1.bids
	(*Amount)(nil), // 2: v1.amount
}
var file_bid_proto_depIdxs = []int32{
	2, // 0: v1.bid.amount:type_name -> v1.amount
	0, // 1: v1.bids.bids:type_name -> v1.bid
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bid_proto_init() }
func file_bid_proto_init() {
	if File_bid_proto != nil {
		return
	}
	file_amount_proto_init()
	file_carrier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bids); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bid_proto_goTypes,
		DependencyIndexes: file_bid_proto_depIdxs,
		MessageInfos:      file_bid_proto_msgTypes,
	}.Build()
	File_bid_proto = out.File
	file_bid_proto_rawDesc = nil
	file_bid_proto_goTypes = nil
	file_bid_proto_depIdxs = nil
}