// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/note/note.proto

package note

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	mi := &file_api_note_note_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	mi := &file_api_note_note_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNoteRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *UpdateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type NoteIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoteIDRequest) Reset() {
	*x = NoteIDRequest{}
	mi := &file_api_note_note_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteIDRequest) ProtoMessage() {}

func (x *NoteIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteIDRequest.ProtoReflect.Descriptor instead.
func (*NoteIDRequest) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{2}
}

func (x *NoteIDRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type UserIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserIDRequest) Reset() {
	*x = UserIDRequest{}
	mi := &file_api_note_note_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDRequest) ProtoMessage() {}

func (x *UserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDRequest.ProtoReflect.Descriptor instead.
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{3}
}

func (x *UserIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type NoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoteResponse) Reset() {
	*x = NoteResponse{}
	mi := &file_api_note_note_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteResponse) ProtoMessage() {}

func (x *NoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteResponse.ProtoReflect.Descriptor instead.
func (*NoteResponse) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{4}
}

func (x *NoteResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NoteResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NoteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NoteResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_api_note_note_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_note_note_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_note_note_proto_rawDescGZIP(), []int{5}
}

var File_api_note_note_proto protoreflect.FileDescriptor

const file_api_note_note_proto_rawDesc = "" +
	"\n" +
	"\x13api/note/note.proto\x12\x04note\"\\\n" +
	"\x11CreateNoteRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"\\\n" +
	"\x11UpdateNoteRequest\x12\x17\n" +
	"\anote_id\x18\x01 \x01(\x03R\x06noteId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"(\n" +
	"\rNoteIDRequest\x12\x17\n" +
	"\anote_id\x18\x01 \x01(\x03R\x06noteId\"(\n" +
	"\rUserIDRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"g\n" +
	"\fNoteResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x03R\x06userId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\"\a\n" +
	"\x05Empty2\xea\x01\n" +
	"\vNoteService\x129\n" +
	"\n" +
	"CreateNote\x12\x17.note.CreateNoteRequest\x1a\x12.note.NoteResponse\x125\n" +
	"\bGetNotes\x12\x13.note.UserIDRequest\x1a\x12.note.NoteResponse0\x01\x129\n" +
	"\n" +
	"UpdateNote\x12\x17.note.UpdateNoteRequest\x1a\x12.note.NoteResponse\x12.\n" +
	"\n" +
	"DeleteNote\x12\x13.note.NoteIDRequest\x1a\v.note.EmptyB+Z)github.com/bekhuli/notes-service/api/noteb\x06proto3"

var (
	file_api_note_note_proto_rawDescOnce sync.Once
	file_api_note_note_proto_rawDescData []byte
)

func file_api_note_note_proto_rawDescGZIP() []byte {
	file_api_note_note_proto_rawDescOnce.Do(func() {
		file_api_note_note_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_note_note_proto_rawDesc), len(file_api_note_note_proto_rawDesc)))
	})
	return file_api_note_note_proto_rawDescData
}

var file_api_note_note_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_note_note_proto_goTypes = []any{
	(*CreateNoteRequest)(nil), // 0: note.CreateNoteRequest
	(*UpdateNoteRequest)(nil), // 1: note.UpdateNoteRequest
	(*NoteIDRequest)(nil),     // 2: note.NoteIDRequest
	(*UserIDRequest)(nil),     // 3: note.UserIDRequest
	(*NoteResponse)(nil),      // 4: note.NoteResponse
	(*Empty)(nil),             // 5: note.Empty
}
var file_api_note_note_proto_depIdxs = []int32{
	0, // 0: note.NoteService.CreateNote:input_type -> note.CreateNoteRequest
	3, // 1: note.NoteService.GetNotes:input_type -> note.UserIDRequest
	1, // 2: note.NoteService.UpdateNote:input_type -> note.UpdateNoteRequest
	2, // 3: note.NoteService.DeleteNote:input_type -> note.NoteIDRequest
	4, // 4: note.NoteService.CreateNote:output_type -> note.NoteResponse
	4, // 5: note.NoteService.GetNotes:output_type -> note.NoteResponse
	4, // 6: note.NoteService.UpdateNote:output_type -> note.NoteResponse
	5, // 7: note.NoteService.DeleteNote:output_type -> note.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_note_note_proto_init() }
func file_api_note_note_proto_init() {
	if File_api_note_note_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_note_note_proto_rawDesc), len(file_api_note_note_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_note_note_proto_goTypes,
		DependencyIndexes: file_api_note_note_proto_depIdxs,
		MessageInfos:      file_api_note_note_proto_msgTypes,
	}.Build()
	File_api_note_note_proto = out.File
	file_api_note_note_proto_goTypes = nil
	file_api_note_note_proto_depIdxs = nil
}
