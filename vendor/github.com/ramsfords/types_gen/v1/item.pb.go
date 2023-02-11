// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: item.proto

package v1

import (
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

type ServingTime int32

const (
	ServingTime_ALLDAY    ServingTime = 0
	ServingTime_MORNING   ServingTime = 1
	ServingTime_LUNCH     ServingTime = 2
	ServingTime_DINNER    ServingTime = 3
	ServingTime_BREAKFAST ServingTime = 4
	ServingTime_BRUNCH    ServingTime = 5
)

// Enum value maps for ServingTime.
var (
	ServingTime_name = map[int32]string{
		0: "ALLDAY",
		1: "MORNING",
		2: "LUNCH",
		3: "DINNER",
		4: "BREAKFAST",
		5: "BRUNCH",
	}
	ServingTime_value = map[string]int32{
		"ALLDAY":    0,
		"MORNING":   1,
		"LUNCH":     2,
		"DINNER":    3,
		"BREAKFAST": 4,
		"BRUNCH":    5,
	}
)

func (x ServingTime) Enum() *ServingTime {
	p := new(ServingTime)
	*p = x
	return p
}

func (x ServingTime) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServingTime) Descriptor() protoreflect.EnumDescriptor {
	return file_item_proto_enumTypes[0].Descriptor()
}

func (ServingTime) Type() protoreflect.EnumType {
	return &file_item_proto_enumTypes[0]
}

func (x ServingTime) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServingTime.Descriptor instead.
func (ServingTime) EnumDescriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"assetId"
	AssetId string `protobuf:"bytes,1,opt,name=assetId,proto3" json:"assetId,omitempty" dynamodbav:"assetId"`
	// @gotags: dynamodbav:"publicId"
	PublicId string `protobuf:"bytes,2,opt,name=publicId,proto3" json:"publicId,omitempty" dynamodbav:"publicId"`
	// @gotags: dynamodbav:"url"
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty" dynamodbav:"url"`
	// @gotags: dynamodbav:"secureUrl"
	SecureUrl string `protobuf:"bytes,4,opt,name=secureUrl,proto3" json:"secureUrl,omitempty" dynamodbav:"secureUrl"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Image) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetSecureUrl() string {
	if x != nil {
		return x.SecureUrl
	}
	return ""
}

type Categories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"categories"
	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty" dynamodbav:"categories"`
}

func (x *Categories) Reset() {
	*x = Categories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Categories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Categories) ProtoMessage() {}

func (x *Categories) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Categories.ProtoReflect.Descriptor instead.
func (*Categories) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *Categories) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id"`
	// @gotags: dynamodbav:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" dynamodbav:"name"`
	// @gotags: dynamodbav:"localName"
	LocalName string `protobuf:"bytes,3,opt,name=localName,proto3" json:"localName,omitempty" dynamodbav:"localName"`
	// @gotags: dynamodbav:"description"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" dynamodbav:"description"`
	// @gotags: dynamodbav:"images"
	Images []*Image `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty" dynamodbav:"images"`
	// @gotags: dynamodbav:"servingTime"
	ServingTime ServingTime `protobuf:"varint,6,opt,name=servingTime,proto3,enum=v1.ServingTime" json:"servingTime,omitempty" dynamodbav:"servingTime"`
	// @gotags: dynamodbav:"restaurantId"
	RestaurantId string `protobuf:"bytes,7,opt,name=restaurantId,proto3" json:"restaurantId,omitempty" dynamodbav:"restaurantId"`
	// @gotags: dynamodbav:"type"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type"`
	// @gotags: dynamodbav:"pk"
	Pk string `protobuf:"bytes,9,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk"`
	// @gotags: dynamodbav:"sk"
	Sk string `protobuf:"bytes,10,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk"`
	// @gotags: dynamodbav:"rank"
	Rank int32 `protobuf:"varint,11,opt,name=rank,proto3" json:"rank,omitempty" dynamodbav:"rank"`
	// @gotags: dynamodbav:"items"
	Items []*Item `protobuf:"bytes,12,rep,name=items,proto3" json:"items,omitempty" dynamodbav:"items"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Category) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Category) GetServingTime() ServingTime {
	if x != nil {
		return x.ServingTime
	}
	return ServingTime_ALLDAY
}

func (x *Category) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Category) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Category) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *Category) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *Category) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Category) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"items"
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" dynamodbav:"items"`
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{3}
}

func (x *Items) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id"`
	// @gotags: dynamodbav:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" dynamodbav:"name"`
	// @gotags: dynamodbav:"localName"
	LocalName string `protobuf:"bytes,3,opt,name=localName,proto3" json:"localName,omitempty" dynamodbav:"localName"`
	// @gotags: dynamodbav:"description"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" dynamodbav:"description"`
	// @gotags: dynamodbav:"price"
	Price string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty" dynamodbav:"price"`
	// @gotags: dynamodbav:"images"
	Images []*Image `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty" dynamodbav:"images"`
	// @gotags: dynamodbav:"spiceLevel"
	SpiceLevel string `protobuf:"bytes,7,opt,name=spiceLevel,proto3" json:"spiceLevel,omitempty" dynamodbav:"spiceLevel"`
	// @gotags: dynamodbav:"isAvailable"
	IsAvailable bool `protobuf:"varint,11,opt,name=isAvailable,proto3" json:"isAvailable,omitempty" dynamodbav:"isAvailable"`
	// @gotags: dynamodbav:"cookingTime"
	CookingTime string `protobuf:"bytes,12,opt,name=cookingTime,proto3" json:"cookingTime,omitempty" dynamodbav:"cookingTime"`
	// @gotags: dynamodbav:"reviews"
	Reviews []string `protobuf:"bytes,13,rep,name=reviews,proto3" json:"reviews,omitempty" dynamodbav:"reviews"`
	// @gotags: dynamodbav:"restaurantId"
	RestaurantId string `protobuf:"bytes,14,opt,name=restaurantId,proto3" json:"restaurantId,omitempty" dynamodbav:"restaurantId"`
	// @gotags: dynamodbav:"type"
	Type string `protobuf:"bytes,15,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type"`
	// @gotags: dynamodbav:"pk"
	Pk string `protobuf:"bytes,16,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk"`
	// @gotags: dynamodbav:"sk"
	Sk string `protobuf:"bytes,17,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk"`
	// @gotags: dynamodbav:"categories"
	// it holds names of categories it belongs to
	Categories []string `protobuf:"bytes,18,rep,name=categories,proto3" json:"categories,omitempty" dynamodbav:"categories"`
	// @gotags: dynamodbav:"tags"
	Tags []string `protobuf:"bytes,19,rep,name=tags,proto3" json:"tags,omitempty" dynamodbav:"tags"`
	// @gotags: dynamodbav:"isVeg"
	IsVeg bool `protobuf:"varint,20,opt,name=isVeg,proto3" json:"isVeg,omitempty" dynamodbav:"isVeg"`
	// @gotags: dynamodbav:"isVegan"
	IsVegan bool `protobuf:"varint,21,opt,name=isVegan,proto3" json:"isVegan,omitempty" dynamodbav:"isVegan"`
	// @gotags: dynamodbav:"isGlutenFree"
	IsGlutenFree bool `protobuf:"varint,22,opt,name=isGlutenFree,proto3" json:"isGlutenFree,omitempty" dynamodbav:"isGlutenFree"`
	// @gotags: dynamodbav:"isDairyFree"
	IsDairyFree bool `protobuf:"varint,23,opt,name=isDairyFree,proto3" json:"isDairyFree,omitempty" dynamodbav:"isDairyFree"`
	// @gotags: dynamodbav:"isNutFree"
	IsNutFree bool `protobuf:"varint,24,opt,name=isNutFree,proto3" json:"isNutFree,omitempty" dynamodbav:"isNutFree"`
	// @gotags: dynamodbav:"isEggFree"
	IsEggFree bool `protobuf:"varint,25,opt,name=isEggFree,proto3" json:"isEggFree,omitempty" dynamodbav:"isEggFree"`
	// @gotags: dynamodbav:"isSoyFree"
	IsSoyFree bool `protobuf:"varint,26,opt,name=isSoyFree,proto3" json:"isSoyFree,omitempty" dynamodbav:"isSoyFree"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Item) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Item) GetSpiceLevel() string {
	if x != nil {
		return x.SpiceLevel
	}
	return ""
}

func (x *Item) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

func (x *Item) GetCookingTime() string {
	if x != nil {
		return x.CookingTime
	}
	return ""
}

func (x *Item) GetReviews() []string {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *Item) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Item) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Item) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *Item) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *Item) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Item) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Item) GetIsVeg() bool {
	if x != nil {
		return x.IsVeg
	}
	return false
}

func (x *Item) GetIsVegan() bool {
	if x != nil {
		return x.IsVegan
	}
	return false
}

func (x *Item) GetIsGlutenFree() bool {
	if x != nil {
		return x.IsGlutenFree
	}
	return false
}

func (x *Item) GetIsDairyFree() bool {
	if x != nil {
		return x.IsDairyFree
	}
	return false
}

func (x *Item) GetIsNutFree() bool {
	if x != nil {
		return x.IsNutFree
	}
	return false
}

func (x *Item) GetIsEggFree() bool {
	if x != nil {
		return x.IsEggFree
	}
	return false
}

func (x *Item) GetIsSoyFree() bool {
	if x != nil {
		return x.IsSoyFree
	}
	return false
}

type ItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"success"
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" dynamodbav:"success"`
	// @gotags: dynamodbav:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" dynamodbav:"message"`
}

func (x *ItemResponse) Reset() {
	*x = ItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemResponse) ProtoMessage() {}

func (x *ItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemResponse.ProtoReflect.Descriptor instead.
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{5}
}

func (x *ItemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x6d, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x22,
	0x3a, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0xd0, 0x02, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x31, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6b,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1e,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x27,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xfd, 0x04, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x70, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x70, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x70, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x73, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x56, 0x65, 0x67,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x56, 0x65, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x56, 0x65, 0x67, 0x61, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x56, 0x65, 0x67, 0x61, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x47, 0x6c, 0x75,
	0x74, 0x65, 0x6e, 0x46, 0x72, 0x65, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x47, 0x6c, 0x75, 0x74, 0x65, 0x6e, 0x46, 0x72, 0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x73, 0x44, 0x61, 0x69, 0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x44, 0x61, 0x69, 0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x4e, 0x75, 0x74, 0x46, 0x72, 0x65, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x4e, 0x75, 0x74, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x45, 0x67, 0x67, 0x46, 0x72, 0x65, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x45, 0x67, 0x67, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53,
	0x6f, 0x79, 0x46, 0x72, 0x65, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x53, 0x6f, 0x79, 0x46, 0x72, 0x65, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x58, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c,
	0x4c, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x55, 0x4e, 0x43, 0x48, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52,
	0x45, 0x41, 0x4b, 0x46, 0x41, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x55,
	0x4e, 0x43, 0x48, 0x10, 0x05, 0x32, 0xd3, 0x01, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5e, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31,
	0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_item_proto_goTypes = []interface{}{
	(ServingTime)(0),     // 0: v1.servingTime
	(*Image)(nil),        // 1: v1.image
	(*Categories)(nil),   // 2: v1.categories
	(*Category)(nil),     // 3: v1.category
	(*Items)(nil),        // 4: v1.items
	(*Item)(nil),         // 5: v1.item
	(*ItemResponse)(nil), // 6: v1.itemResponse
}
var file_item_proto_depIdxs = []int32{
	3,  // 0: v1.categories.categories:type_name -> v1.category
	1,  // 1: v1.category.images:type_name -> v1.image
	0,  // 2: v1.category.servingTime:type_name -> v1.servingTime
	5,  // 3: v1.category.items:type_name -> v1.item
	5,  // 4: v1.items.items:type_name -> v1.item
	1,  // 5: v1.item.images:type_name -> v1.image
	5,  // 6: v1.itemService.createItem:input_type -> v1.item
	4,  // 7: v1.itemService.createItems:input_type -> v1.items
	3,  // 8: v1.itemService.createCategory:input_type -> v1.category
	2,  // 9: v1.itemService.createCategories:input_type -> v1.categories
	6,  // 10: v1.itemService.createItem:output_type -> v1.itemResponse
	6,  // 11: v1.itemService.createItems:output_type -> v1.itemResponse
	6,  // 12: v1.itemService.createCategory:output_type -> v1.itemResponse
	6,  // 13: v1.itemService.createCategories:output_type -> v1.itemResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Categories); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
		file_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemResponse); i {
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
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		EnumInfos:         file_item_proto_enumTypes,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}