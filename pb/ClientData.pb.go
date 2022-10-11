// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: ClientData.proto

package pb

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

type ClientOnScreenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldUpdate bool   `protobuf:"varint,1,opt,name=ShouldUpdate,proto3" json:"ShouldUpdate,omitempty"`
	OnScreenText string `protobuf:"bytes,2,opt,name=OnScreenText,proto3" json:"OnScreenText,omitempty"`
}

func (x *ClientOnScreenData) Reset() {
	*x = ClientOnScreenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientOnScreenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientOnScreenData) ProtoMessage() {}

func (x *ClientOnScreenData) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientOnScreenData.ProtoReflect.Descriptor instead.
func (*ClientOnScreenData) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{0}
}

func (x *ClientOnScreenData) GetShouldUpdate() bool {
	if x != nil {
		return x.ShouldUpdate
	}
	return false
}

func (x *ClientOnScreenData) GetOnScreenText() string {
	if x != nil {
		return x.OnScreenText
	}
	return ""
}

type ClientExecData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldExec bool   `protobuf:"varint,1,opt,name=ShouldExec,proto3" json:"ShouldExec,omitempty"`
	Command    string `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (x *ClientExecData) Reset() {
	*x = ClientExecData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientExecData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientExecData) ProtoMessage() {}

func (x *ClientExecData) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientExecData.ProtoReflect.Descriptor instead.
func (*ClientExecData) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{1}
}

func (x *ClientExecData) GetShouldExec() bool {
	if x != nil {
		return x.ShouldExec
	}
	return false
}

func (x *ClientExecData) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type ClientExecOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *ClientDataRequest `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Output []byte             `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ClientExecOutput) Reset() {
	*x = ClientExecOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientExecOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientExecOutput) ProtoMessage() {}

func (x *ClientExecOutput) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientExecOutput.ProtoReflect.Descriptor instead.
func (*ClientExecOutput) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{2}
}

func (x *ClientExecOutput) GetId() *ClientDataRequest {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ClientExecOutput) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{3}
}

type FloodData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldFlood bool `protobuf:"varint,1,opt,name=ShouldFlood,proto3" json:"ShouldFlood,omitempty"`
}

func (x *FloodData) Reset() {
	*x = FloodData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloodData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloodData) ProtoMessage() {}

func (x *FloodData) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloodData.ProtoReflect.Descriptor instead.
func (*FloodData) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{4}
}

func (x *FloodData) GetShouldFlood() bool {
	if x != nil {
		return x.ShouldFlood
	}
	return false
}

type ClientDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientDataRequest) Reset() {
	*x = ClientDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDataRequest) ProtoMessage() {}

func (x *ClientDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDataRequest.ProtoReflect.Descriptor instead.
func (*ClientDataRequest) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{5}
}

func (x *ClientDataRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientDataOnScreenTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnScreen *ClientOnScreenData `protobuf:"bytes,1,opt,name=on_screen,json=onScreen,proto3" json:"on_screen,omitempty"`
}

func (x *ClientDataOnScreenTextResponse) Reset() {
	*x = ClientDataOnScreenTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDataOnScreenTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDataOnScreenTextResponse) ProtoMessage() {}

func (x *ClientDataOnScreenTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ClientData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDataOnScreenTextResponse.ProtoReflect.Descriptor instead.
func (*ClientDataOnScreenTextResponse) Descriptor() ([]byte, []int) {
	return file_ClientData_proto_rawDescGZIP(), []int{6}
}

func (x *ClientDataOnScreenTextResponse) GetOnScreen() *ClientOnScreenData {
	if x != nil {
		return x.OnScreen
	}
	return nil
}

var File_ClientData_proto protoreflect.FileDescriptor

var file_ClientData_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5c, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4f, 0x6e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x6e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x6e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x54, 0x65, 0x78, 0x74, 0x22, 0x4a, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x78,
	0x65, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x45, 0x78, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x53, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x45, 0x78, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x51, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65, 0x63, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x09, 0x46,
	0x6c, 0x6f, 0x6f, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x53,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x46, 0x6c, 0x6f, 0x6f, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x1e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x6e, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x09, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6f, 0x6e, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x32, 0xf2, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x12, 0x54, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x4f, 0x6e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x32, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x64, 0x12,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x6c, 0x6f, 0x6f, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientData_proto_rawDescOnce sync.Once
	file_ClientData_proto_rawDescData = file_ClientData_proto_rawDesc
)

func file_ClientData_proto_rawDescGZIP() []byte {
	file_ClientData_proto_rawDescOnce.Do(func() {
		file_ClientData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientData_proto_rawDescData)
	})
	return file_ClientData_proto_rawDescData
}

var file_ClientData_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ClientData_proto_goTypes = []interface{}{
	(*ClientOnScreenData)(nil),             // 0: pb.ClientOnScreenData
	(*ClientExecData)(nil),                 // 1: pb.ClientExecData
	(*ClientExecOutput)(nil),               // 2: pb.ClientExecOutput
	(*Void)(nil),                           // 3: pb.Void
	(*FloodData)(nil),                      // 4: pb.FloodData
	(*ClientDataRequest)(nil),              // 5: pb.ClientDataRequest
	(*ClientDataOnScreenTextResponse)(nil), // 6: pb.ClientDataOnScreenTextResponse
}
var file_ClientData_proto_depIdxs = []int32{
	5, // 0: pb.ClientExecOutput.id:type_name -> pb.ClientDataRequest
	0, // 1: pb.ClientDataOnScreenTextResponse.on_screen:type_name -> pb.ClientOnScreenData
	5, // 2: pb.Consumer.SubscribeOnScreenText:input_type -> pb.ClientDataRequest
	5, // 3: pb.Consumer.GetCommand:input_type -> pb.ClientDataRequest
	2, // 4: pb.Consumer.SetCommandOutput:input_type -> pb.ClientExecOutput
	3, // 5: pb.Consumer.GetFlood:input_type -> pb.Void
	6, // 6: pb.Consumer.SubscribeOnScreenText:output_type -> pb.ClientDataOnScreenTextResponse
	1, // 7: pb.Consumer.GetCommand:output_type -> pb.ClientExecData
	3, // 8: pb.Consumer.SetCommandOutput:output_type -> pb.Void
	4, // 9: pb.Consumer.GetFlood:output_type -> pb.FloodData
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ClientData_proto_init() }
func file_ClientData_proto_init() {
	if File_ClientData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientOnScreenData); i {
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
		file_ClientData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientExecData); i {
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
		file_ClientData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientExecOutput); i {
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
		file_ClientData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_ClientData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloodData); i {
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
		file_ClientData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDataRequest); i {
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
		file_ClientData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDataOnScreenTextResponse); i {
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
			RawDescriptor: file_ClientData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ClientData_proto_goTypes,
		DependencyIndexes: file_ClientData_proto_depIdxs,
		MessageInfos:      file_ClientData_proto_msgTypes,
	}.Build()
	File_ClientData_proto = out.File
	file_ClientData_proto_rawDesc = nil
	file_ClientData_proto_goTypes = nil
	file_ClientData_proto_depIdxs = nil
}
