// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1/storage.proto

package storage

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for `CreateReadSession`.
type CreateReadSessionRequest struct {
	// Required. The request project that owns the session, in the form of
	// `projects/{project_id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Session to be created.
	ReadSession *ReadSession `protobuf:"bytes,2,opt,name=read_session,json=readSession,proto3" json:"read_session,omitempty"`
	// Max initial number of streams. If unset or zero, the server will
	// provide a value of streams so as to produce reasonable throughput. Must be
	// non-negative. The number of streams may be lower than the requested number,
	// depending on the amount parallelism that is reasonable for the table. Error
	// will be returned if the max count is greater than the current system
	// max limit of 1,000.
	//
	// Streams must be read starting from offset 0.
	MaxStreamCount       int32    `protobuf:"varint,3,opt,name=max_stream_count,json=maxStreamCount,proto3" json:"max_stream_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReadSessionRequest) Reset()         { *m = CreateReadSessionRequest{} }
func (m *CreateReadSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateReadSessionRequest) ProtoMessage()    {}
func (*CreateReadSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{0}
}

func (m *CreateReadSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReadSessionRequest.Unmarshal(m, b)
}
func (m *CreateReadSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReadSessionRequest.Marshal(b, m, deterministic)
}
func (m *CreateReadSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReadSessionRequest.Merge(m, src)
}
func (m *CreateReadSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateReadSessionRequest.Size(m)
}
func (m *CreateReadSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReadSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReadSessionRequest proto.InternalMessageInfo

func (m *CreateReadSessionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateReadSessionRequest) GetReadSession() *ReadSession {
	if m != nil {
		return m.ReadSession
	}
	return nil
}

func (m *CreateReadSessionRequest) GetMaxStreamCount() int32 {
	if m != nil {
		return m.MaxStreamCount
	}
	return 0
}

// Request message for `ReadRows`.
type ReadRowsRequest struct {
	// Required. Stream to read rows from.
	ReadStream string `protobuf:"bytes,1,opt,name=read_stream,json=readStream,proto3" json:"read_stream,omitempty"`
	// The offset requested must be less than the last row read from Read.
	// Requesting a larger offset is undefined. If not specified, start reading
	// from offset zero.
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRowsRequest) Reset()         { *m = ReadRowsRequest{} }
func (m *ReadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRowsRequest) ProtoMessage()    {}
func (*ReadRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{1}
}

func (m *ReadRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRowsRequest.Unmarshal(m, b)
}
func (m *ReadRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRowsRequest.Marshal(b, m, deterministic)
}
func (m *ReadRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRowsRequest.Merge(m, src)
}
func (m *ReadRowsRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRowsRequest.Size(m)
}
func (m *ReadRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRowsRequest proto.InternalMessageInfo

func (m *ReadRowsRequest) GetReadStream() string {
	if m != nil {
		return m.ReadStream
	}
	return ""
}

func (m *ReadRowsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Information on if the current connection is being throttled.
type ThrottleState struct {
	// How much this connection is being throttled. Zero means no throttling,
	// 100 means fully throttled.
	ThrottlePercent      int32    `protobuf:"varint,1,opt,name=throttle_percent,json=throttlePercent,proto3" json:"throttle_percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThrottleState) Reset()         { *m = ThrottleState{} }
func (m *ThrottleState) String() string { return proto.CompactTextString(m) }
func (*ThrottleState) ProtoMessage()    {}
func (*ThrottleState) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{2}
}

func (m *ThrottleState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThrottleState.Unmarshal(m, b)
}
func (m *ThrottleState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThrottleState.Marshal(b, m, deterministic)
}
func (m *ThrottleState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThrottleState.Merge(m, src)
}
func (m *ThrottleState) XXX_Size() int {
	return xxx_messageInfo_ThrottleState.Size(m)
}
func (m *ThrottleState) XXX_DiscardUnknown() {
	xxx_messageInfo_ThrottleState.DiscardUnknown(m)
}

var xxx_messageInfo_ThrottleState proto.InternalMessageInfo

func (m *ThrottleState) GetThrottlePercent() int32 {
	if m != nil {
		return m.ThrottlePercent
	}
	return 0
}

// Estimated stream statistics for a given Stream.
type StreamStats struct {
	// Represents the progress of the current stream.
	Progress             *StreamStats_Progress `protobuf:"bytes,2,opt,name=progress,proto3" json:"progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StreamStats) Reset()         { *m = StreamStats{} }
func (m *StreamStats) String() string { return proto.CompactTextString(m) }
func (*StreamStats) ProtoMessage()    {}
func (*StreamStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{3}
}

func (m *StreamStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamStats.Unmarshal(m, b)
}
func (m *StreamStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamStats.Marshal(b, m, deterministic)
}
func (m *StreamStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamStats.Merge(m, src)
}
func (m *StreamStats) XXX_Size() int {
	return xxx_messageInfo_StreamStats.Size(m)
}
func (m *StreamStats) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamStats.DiscardUnknown(m)
}

var xxx_messageInfo_StreamStats proto.InternalMessageInfo

func (m *StreamStats) GetProgress() *StreamStats_Progress {
	if m != nil {
		return m.Progress
	}
	return nil
}

type StreamStats_Progress struct {
	// The fraction of rows assigned to the stream that have been processed by
	// the server so far, not including the rows in the current response
	// message.
	//
	// This value, along with `at_response_end`, can be used to interpolate
	// the progress made as the rows in the message are being processed using
	// the following formula: `at_response_start + (at_response_end -
	// at_response_start) * rows_processed_from_response / rows_in_response`.
	//
	// Note that if a filter is provided, the `at_response_end` value of the
	// previous response may not necessarily be equal to the
	// `at_response_start` value of the current response.
	AtResponseStart float64 `protobuf:"fixed64,1,opt,name=at_response_start,json=atResponseStart,proto3" json:"at_response_start,omitempty"`
	// Similar to `at_response_start`, except that this value includes the
	// rows in the current response.
	AtResponseEnd        float64  `protobuf:"fixed64,2,opt,name=at_response_end,json=atResponseEnd,proto3" json:"at_response_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamStats_Progress) Reset()         { *m = StreamStats_Progress{} }
func (m *StreamStats_Progress) String() string { return proto.CompactTextString(m) }
func (*StreamStats_Progress) ProtoMessage()    {}
func (*StreamStats_Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{3, 0}
}

func (m *StreamStats_Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamStats_Progress.Unmarshal(m, b)
}
func (m *StreamStats_Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamStats_Progress.Marshal(b, m, deterministic)
}
func (m *StreamStats_Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamStats_Progress.Merge(m, src)
}
func (m *StreamStats_Progress) XXX_Size() int {
	return xxx_messageInfo_StreamStats_Progress.Size(m)
}
func (m *StreamStats_Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamStats_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_StreamStats_Progress proto.InternalMessageInfo

func (m *StreamStats_Progress) GetAtResponseStart() float64 {
	if m != nil {
		return m.AtResponseStart
	}
	return 0
}

func (m *StreamStats_Progress) GetAtResponseEnd() float64 {
	if m != nil {
		return m.AtResponseEnd
	}
	return 0
}

// Response from calling `ReadRows` may include row data, progress and
// throttling information.
type ReadRowsResponse struct {
	// Row data is returned in format specified during session creation.
	//
	// Types that are valid to be assigned to Rows:
	//	*ReadRowsResponse_AvroRows
	//	*ReadRowsResponse_ArrowRecordBatch
	Rows isReadRowsResponse_Rows `protobuf_oneof:"rows"`
	// Number of serialized rows in the rows block.
	RowCount int64 `protobuf:"varint,6,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	// Statistics for the stream.
	Stats *StreamStats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	// Throttling state. If unset, the latest response still describes
	// the current throttling status.
	ThrottleState        *ThrottleState `protobuf:"bytes,5,opt,name=throttle_state,json=throttleState,proto3" json:"throttle_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadRowsResponse) Reset()         { *m = ReadRowsResponse{} }
func (m *ReadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*ReadRowsResponse) ProtoMessage()    {}
func (*ReadRowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{4}
}

func (m *ReadRowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRowsResponse.Unmarshal(m, b)
}
func (m *ReadRowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRowsResponse.Marshal(b, m, deterministic)
}
func (m *ReadRowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRowsResponse.Merge(m, src)
}
func (m *ReadRowsResponse) XXX_Size() int {
	return xxx_messageInfo_ReadRowsResponse.Size(m)
}
func (m *ReadRowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRowsResponse proto.InternalMessageInfo

type isReadRowsResponse_Rows interface {
	isReadRowsResponse_Rows()
}

type ReadRowsResponse_AvroRows struct {
	AvroRows *AvroRows `protobuf:"bytes,3,opt,name=avro_rows,json=avroRows,proto3,oneof"`
}

type ReadRowsResponse_ArrowRecordBatch struct {
	ArrowRecordBatch *ArrowRecordBatch `protobuf:"bytes,4,opt,name=arrow_record_batch,json=arrowRecordBatch,proto3,oneof"`
}

func (*ReadRowsResponse_AvroRows) isReadRowsResponse_Rows() {}

func (*ReadRowsResponse_ArrowRecordBatch) isReadRowsResponse_Rows() {}

func (m *ReadRowsResponse) GetRows() isReadRowsResponse_Rows {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ReadRowsResponse) GetAvroRows() *AvroRows {
	if x, ok := m.GetRows().(*ReadRowsResponse_AvroRows); ok {
		return x.AvroRows
	}
	return nil
}

func (m *ReadRowsResponse) GetArrowRecordBatch() *ArrowRecordBatch {
	if x, ok := m.GetRows().(*ReadRowsResponse_ArrowRecordBatch); ok {
		return x.ArrowRecordBatch
	}
	return nil
}

func (m *ReadRowsResponse) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *ReadRowsResponse) GetStats() *StreamStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *ReadRowsResponse) GetThrottleState() *ThrottleState {
	if m != nil {
		return m.ThrottleState
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadRowsResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadRowsResponse_AvroRows)(nil),
		(*ReadRowsResponse_ArrowRecordBatch)(nil),
	}
}

// Request message for `SplitReadStream`.
type SplitReadStreamRequest struct {
	// Required. Name of the stream to split.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A value in the range (0.0, 1.0) that specifies the fractional point at
	// which the original stream should be split. The actual split point is
	// evaluated on pre-filtered rows, so if a filter is provided, then there is
	// no guarantee that the division of the rows between the new child streams
	// will be proportional to this fractional value. Additionally, because the
	// server-side unit for assigning data is collections of rows, this fraction
	// will always map to a data storage boundary on the server side.
	Fraction             float64  `protobuf:"fixed64,2,opt,name=fraction,proto3" json:"fraction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SplitReadStreamRequest) Reset()         { *m = SplitReadStreamRequest{} }
func (m *SplitReadStreamRequest) String() string { return proto.CompactTextString(m) }
func (*SplitReadStreamRequest) ProtoMessage()    {}
func (*SplitReadStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{5}
}

func (m *SplitReadStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitReadStreamRequest.Unmarshal(m, b)
}
func (m *SplitReadStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitReadStreamRequest.Marshal(b, m, deterministic)
}
func (m *SplitReadStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitReadStreamRequest.Merge(m, src)
}
func (m *SplitReadStreamRequest) XXX_Size() int {
	return xxx_messageInfo_SplitReadStreamRequest.Size(m)
}
func (m *SplitReadStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitReadStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SplitReadStreamRequest proto.InternalMessageInfo

func (m *SplitReadStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SplitReadStreamRequest) GetFraction() float64 {
	if m != nil {
		return m.Fraction
	}
	return 0
}

// Response message for `SplitReadStream`.
type SplitReadStreamResponse struct {
	// Primary stream, which contains the beginning portion of
	// |original_stream|. An empty value indicates that the original stream can no
	// longer be split.
	PrimaryStream *ReadStream `protobuf:"bytes,1,opt,name=primary_stream,json=primaryStream,proto3" json:"primary_stream,omitempty"`
	// Remainder stream, which contains the tail of |original_stream|. An empty
	// value indicates that the original stream can no longer be split.
	RemainderStream      *ReadStream `protobuf:"bytes,2,opt,name=remainder_stream,json=remainderStream,proto3" json:"remainder_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SplitReadStreamResponse) Reset()         { *m = SplitReadStreamResponse{} }
func (m *SplitReadStreamResponse) String() string { return proto.CompactTextString(m) }
func (*SplitReadStreamResponse) ProtoMessage()    {}
func (*SplitReadStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bf7871be7b3410, []int{6}
}

func (m *SplitReadStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitReadStreamResponse.Unmarshal(m, b)
}
func (m *SplitReadStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitReadStreamResponse.Marshal(b, m, deterministic)
}
func (m *SplitReadStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitReadStreamResponse.Merge(m, src)
}
func (m *SplitReadStreamResponse) XXX_Size() int {
	return xxx_messageInfo_SplitReadStreamResponse.Size(m)
}
func (m *SplitReadStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitReadStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SplitReadStreamResponse proto.InternalMessageInfo

func (m *SplitReadStreamResponse) GetPrimaryStream() *ReadStream {
	if m != nil {
		return m.PrimaryStream
	}
	return nil
}

func (m *SplitReadStreamResponse) GetRemainderStream() *ReadStream {
	if m != nil {
		return m.RemainderStream
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateReadSessionRequest)(nil), "google.cloud.bigquery.storage.v1.CreateReadSessionRequest")
	proto.RegisterType((*ReadRowsRequest)(nil), "google.cloud.bigquery.storage.v1.ReadRowsRequest")
	proto.RegisterType((*ThrottleState)(nil), "google.cloud.bigquery.storage.v1.ThrottleState")
	proto.RegisterType((*StreamStats)(nil), "google.cloud.bigquery.storage.v1.StreamStats")
	proto.RegisterType((*StreamStats_Progress)(nil), "google.cloud.bigquery.storage.v1.StreamStats.Progress")
	proto.RegisterType((*ReadRowsResponse)(nil), "google.cloud.bigquery.storage.v1.ReadRowsResponse")
	proto.RegisterType((*SplitReadStreamRequest)(nil), "google.cloud.bigquery.storage.v1.SplitReadStreamRequest")
	proto.RegisterType((*SplitReadStreamResponse)(nil), "google.cloud.bigquery.storage.v1.SplitReadStreamResponse")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1/storage.proto", fileDescriptor_79bf7871be7b3410)
}

var fileDescriptor_79bf7871be7b3410 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xef, 0x38, 0x89, 0xe5, 0x8c, 0x93, 0xd8, 0x9d, 0x43, 0x6b, 0x0c, 0x82, 0x68, 0x55, 0xa1,
	0xd4, 0x71, 0x76, 0x89, 0x2b, 0x95, 0x12, 0x28, 0xc8, 0x8e, 0x2a, 0x8a, 0xb8, 0x98, 0x31, 0x7f,
	0xa4, 0x1e, 0x58, 0x8d, 0xd7, 0xe3, 0xf5, 0xa2, 0xdd, 0x9d, 0xed, 0xcc, 0xd8, 0x6e, 0x55, 0x55,
	0x48, 0x70, 0xe2, 0xcc, 0xf7, 0xe0, 0x3b, 0x70, 0x29, 0x12, 0x17, 0x24, 0xb8, 0xf5, 0xd4, 0x03,
	0x07, 0xf8, 0x0a, 0x9c, 0xd0, 0xfc, 0x59, 0x67, 0x71, 0x49, 0xed, 0xe4, 0x36, 0xfb, 0xe6, 0xf7,
	0x7e, 0xf3, 0x7e, 0x6f, 0xde, 0x7b, 0x3b, 0xd0, 0x0d, 0x19, 0x0b, 0x63, 0xea, 0x05, 0x31, 0x9b,
	0x8e, 0xbc, 0x61, 0x14, 0x3e, 0x9c, 0x52, 0xfe, 0xd8, 0x13, 0x92, 0x71, 0x12, 0x52, 0x6f, 0x76,
	0x9c, 0x2f, 0xdd, 0x8c, 0x33, 0xc9, 0xd0, 0xbe, 0xc1, 0xbb, 0x1a, 0xef, 0xe6, 0x78, 0x37, 0x07,
	0xcd, 0x8e, 0x9b, 0x6f, 0x58, 0x46, 0x92, 0x45, 0x1e, 0x49, 0x53, 0x26, 0x89, 0x8c, 0x58, 0x2a,
	0x8c, 0x7f, 0xf3, 0x7a, 0x61, 0x37, 0x88, 0x23, 0x9a, 0x4a, 0xbb, 0xf1, 0x56, 0x61, 0x63, 0x1c,
	0xd1, 0x78, 0xe4, 0x0f, 0xe9, 0x84, 0xcc, 0x22, 0xc6, 0x2d, 0xe0, 0xb5, 0x02, 0x80, 0x53, 0xc1,
	0xa6, 0x3c, 0xb0, 0x41, 0x35, 0xdb, 0x2b, 0x45, 0x10, 0xce, 0xd9, 0xdc, 0xa2, 0x0f, 0x57, 0xa3,
	0x67, 0x9c, 0x59, 0xf0, 0xd1, 0x1a, 0xf9, 0xe1, 0x94, 0x24, 0x06, 0xee, 0xfc, 0x05, 0x60, 0xe3,
	0x94, 0x53, 0x22, 0x29, 0xa6, 0x64, 0x34, 0xa0, 0x42, 0x44, 0x2c, 0xc5, 0xf4, 0xe1, 0x94, 0x0a,
	0x89, 0x3e, 0x85, 0xe5, 0x8c, 0x70, 0x9a, 0xca, 0x06, 0xd8, 0x07, 0x07, 0xdb, 0xbd, 0x5b, 0x2f,
	0xba, 0xa5, 0x7f, 0xba, 0x47, 0xf0, 0x50, 0x93, 0xe7, 0x9a, 0x12, 0x92, 0x92, 0x90, 0x72, 0x7b,
	0x31, 0x24, 0x8b, 0x84, 0x1b, 0xb0, 0xc4, 0xeb, 0x73, 0xf6, 0x0d, 0x0d, 0x24, 0xb6, 0x14, 0xe8,
	0x0b, 0xb8, 0xc3, 0x29, 0x19, 0xf9, 0xc2, 0x9c, 0xd1, 0x28, 0xed, 0x83, 0x83, 0x6a, 0xe7, 0xc8,
	0x5d, 0x75, 0x3f, 0x6e, 0x21, 0xb0, 0xde, 0xc6, 0x8b, 0x6e, 0x09, 0x57, 0xf9, 0x99, 0x05, 0x1d,
	0xc0, 0x7a, 0x42, 0x1e, 0xf9, 0x46, 0x94, 0x1f, 0xb0, 0x69, 0x2a, 0x1b, 0x1b, 0xfb, 0xe0, 0x60,
	0x0b, 0xef, 0x25, 0xe4, 0xd1, 0x40, 0x9b, 0x4f, 0x95, 0xd5, 0x79, 0x0a, 0x6b, 0x8a, 0x0a, 0xb3,
	0xb9, 0xc8, 0x05, 0x62, 0x58, 0x35, 0x31, 0x69, 0x98, 0x55, 0x79, 0xac, 0x55, 0x1e, 0xc2, 0x9b,
	0x79, 0x30, 0x79, 0x2c, 0x4b, 0x02, 0x75, 0x5c, 0xda, 0x11, 0x43, 0xbe, 0x58, 0xa3, 0x6b, 0xb0,
	0xcc, 0xc6, 0x63, 0x41, 0xa5, 0x56, 0xb8, 0x81, 0xed, 0x97, 0x73, 0x02, 0x77, 0x3f, 0x9f, 0x70,
	0x26, 0x65, 0x4c, 0x07, 0x92, 0x48, 0x8a, 0x6e, 0xc2, 0xba, 0xb4, 0x06, 0x3f, 0xa3, 0x3c, 0xc8,
	0xf3, 0xbc, 0x85, 0x6b, 0xb9, 0xbd, 0x6f, 0xcc, 0xce, 0xcf, 0x00, 0x56, 0x0d, 0xbd, 0x72, 0x15,
	0x08, 0xc3, 0x4a, 0xc6, 0x59, 0xc8, 0xa9, 0x10, 0x36, 0x8f, 0xb7, 0x57, 0xe7, 0xb1, 0x40, 0xe0,
	0xf6, 0xad, 0x37, 0x5e, 0xf0, 0x34, 0xbf, 0x86, 0x95, 0xdc, 0x8a, 0x5a, 0xf0, 0x2a, 0x91, 0x3e,
	0xa7, 0x22, 0x63, 0xa9, 0xa0, 0xbe, 0x90, 0x84, 0x9b, 0xd8, 0x00, 0xae, 0x11, 0x89, 0xad, 0x7d,
	0xa0, 0xcc, 0xe8, 0x6d, 0x58, 0x2b, 0x62, 0x69, 0x3a, 0xd2, 0x21, 0x01, 0xbc, 0x7b, 0x86, 0xbc,
	0x97, 0x8e, 0x9c, 0xef, 0x37, 0x60, 0xfd, 0x2c, 0xff, 0xc6, 0x8e, 0x3e, 0x81, 0xdb, 0xaa, 0x76,
	0x7d, 0xce, 0xe6, 0x42, 0x5f, 0x5b, 0xb5, 0xd3, 0x5a, 0xad, 0xa4, 0x3b, 0xe3, 0x4c, 0xd1, 0xdc,
	0xbf, 0x82, 0x2b, 0xc4, 0xae, 0xd1, 0x10, 0x22, 0xdd, 0x34, 0x3e, 0xa7, 0x01, 0xe3, 0x23, 0x7f,
	0x48, 0x64, 0x30, 0x69, 0x6c, 0x6a, 0xce, 0xce, 0x1a, 0x9c, 0xca, 0x17, 0x6b, 0xd7, 0x9e, 0xf2,
	0xbc, 0x7f, 0x05, 0xd7, 0xc9, 0x92, 0x0d, 0xbd, 0x0e, 0xb7, 0xd5, 0x09, 0xa6, 0xca, 0xca, 0xfa,
	0x7a, 0x2b, 0x9c, 0xcd, 0x75, 0x7d, 0xa1, 0x53, 0xb8, 0x25, 0x54, 0x72, 0xd7, 0xaf, 0xec, 0xc2,
	0x8d, 0x60, 0xe3, 0x8b, 0xbe, 0x84, 0x7b, 0x8b, 0xa2, 0x50, 0x16, 0xda, 0xd8, 0xd2, 0x6c, 0xde,
	0x6a, 0xb6, 0xff, 0x54, 0x17, 0xde, 0x95, 0xc5, 0xcf, 0x5e, 0x19, 0x6e, 0xaa, 0x1c, 0x3b, 0x4f,
	0xe0, 0xb5, 0x41, 0x16, 0x47, 0xb2, 0x50, 0xbc, 0xb6, 0x17, 0xee, 0xc1, 0xcd, 0x94, 0x24, 0xf4,
	0xf2, 0x4d, 0xa0, 0xdd, 0x51, 0x13, 0x56, 0xc6, 0x9c, 0x04, 0x32, 0x6f, 0x71, 0x80, 0x17, 0xdf,
	0xce, 0x2f, 0x00, 0x5e, 0x7f, 0xe9, 0x74, 0x5b, 0x09, 0x03, 0xb8, 0x97, 0xf1, 0x28, 0x21, 0xfc,
	0x71, 0xb1, 0x1b, 0xab, 0x9d, 0xf6, 0x9a, 0x03, 0xc2, 0xb0, 0xed, 0x5a, 0x0e, 0xdb, 0x8b, 0x5f,
	0xc1, 0x3a, 0xa7, 0x09, 0x89, 0xd2, 0x11, 0xe5, 0x39, 0x6d, 0xe9, 0x12, 0xb4, 0xb5, 0x05, 0x8b,
	0x31, 0x74, 0x9e, 0x95, 0xe1, 0x4e, 0x2f, 0x0a, 0x3f, 0x53, 0x3e, 0x0a, 0x87, 0xfe, 0x06, 0xf0,
	0xea, 0x4b, 0x73, 0x14, 0x9d, 0xac, 0x3e, 0xe5, 0xbc, 0xe1, 0xdb, 0xbc, 0xd8, 0x64, 0x74, 0xa2,
	0xe7, 0xdd, 0x1b, 0x66, 0xd2, 0xb6, 0x8b, 0x63, 0xb6, 0xbd, 0x3c, 0x1c, 0xbf, 0xfb, 0xe3, 0xcf,
	0x1f, 0x4b, 0x1f, 0x38, 0xef, 0xaa, 0x3f, 0xc1, 0x93, 0x22, 0xd2, 0x95, 0x64, 0x18, 0xd3, 0xbb,
	0x99, 0x19, 0xdb, 0xc2, 0x6b, 0x79, 0x23, 0x22, 0x89, 0xa0, 0x7a, 0xa9, 0xf7, 0x84, 0xd7, 0x7a,
	0x7a, 0x02, 0x5a, 0xe8, 0x37, 0x00, 0x2b, 0x79, 0x27, 0xa3, 0xe3, 0xf5, 0xc2, 0x2c, 0x4c, 0xdd,
	0x66, 0xe7, 0x22, 0x2e, 0xa6, 0x3c, 0x9c, 0x07, 0xcf, 0xbb, 0xa8, 0x30, 0xaa, 0xdb, 0x66, 0xa8,
	0x6a, 0x31, 0x1f, 0xa1, 0xbb, 0x05, 0x31, 0x7a, 0xbb, 0xa8, 0x22, 0x66, 0x81, 0xf9, 0xa1, 0x7b,
	0x2d, 0xcf, 0x4a, 0xd5, 0x4b, 0x0d, 0x54, 0x92, 0xde, 0x01, 0xe8, 0x19, 0x80, 0xb5, 0xa5, 0xc2,
	0x44, 0x77, 0xd6, 0xe8, 0xdf, 0xff, 0xed, 0xa4, 0xe6, 0x7b, 0x97, 0xf0, 0xb4, 0x32, 0x3f, 0xd4,
	0x8a, 0xee, 0xa0, 0xdb, 0x5a, 0x91, 0x6a, 0xa8, 0x0b, 0x49, 0x69, 0xfe, 0x04, 0x7e, 0xed, 0xbe,
	0xf9, 0xea, 0x96, 0xfd, 0xbd, 0xfb, 0x03, 0x98, 0x48, 0x99, 0x89, 0x13, 0xcf, 0x9b, 0xcf, 0xe7,
	0xcb, 0x0d, 0x4d, 0xa6, 0x72, 0xb2, 0x78, 0x36, 0xb4, 0xd7, 0x05, 0xba, 0x2a, 0xf9, 0x2c, 0x8d,
	0x57, 0x7b, 0xe8, 0x9c, 0x1c, 0x65, 0x31, 0x91, 0x63, 0xc6, 0x93, 0xde, 0xb7, 0xf0, 0x46, 0xc0,
	0x92, 0x95, 0x09, 0xeb, 0xed, 0x0c, 0xcc, 0xba, 0xaf, 0x5e, 0x2d, 0x7d, 0xf0, 0xe0, 0x63, 0xeb,
	0x11, 0xb2, 0x98, 0xa4, 0xa1, 0xcb, 0x78, 0xe8, 0x85, 0x34, 0xd5, 0x6f, 0x1a, 0xef, 0xec, 0xd8,
	0xf3, 0x5f, 0x41, 0xef, 0xdb, 0xe5, 0xb0, 0xac, 0x7d, 0x6e, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x29, 0xec, 0xd8, 0xe0, 0x58, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BigQueryReadClient is the client API for BigQueryRead service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BigQueryReadClient interface {
	// Creates a new read session. A read session divides the contents of a
	// BigQuery table into one or more streams, which can then be used to read
	// data from the table. The read session also specifies properties of the
	// data to be read, such as a list of columns or a push-down filter describing
	// the rows to be returned.
	//
	// A particular row can be read by at most one stream. When the caller has
	// reached the end of each stream in the session, then all the data in the
	// table has been read.
	//
	// Data is assigned to each stream such that roughly the same number of
	// rows can be read from each stream. Because the server-side unit for
	// assigning data is collections of rows, the API does not guarantee that
	// each stream will return the same number or rows. Additionally, the
	// limits are enforced based on the number of pre-filtered rows, so some
	// filters can lead to lopsided assignments.
	//
	// Read sessions automatically expire 24 hours after they are created and do
	// not require manual clean-up by the caller.
	CreateReadSession(ctx context.Context, in *CreateReadSessionRequest, opts ...grpc.CallOption) (*ReadSession, error)
	// Reads rows from the stream in the format prescribed by the ReadSession.
	// Each response contains one or more table rows, up to a maximum of 100 MiB
	// per response; read requests which attempt to read individual rows larger
	// than 100 MiB will fail.
	//
	// Each request also returns a set of stream statistics reflecting the current
	// state of the stream.
	ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigQueryRead_ReadRowsClient, error)
	// Splits a given `ReadStream` into two `ReadStream` objects. These
	// `ReadStream` objects are referred to as the primary and the residual
	// streams of the split. The original `ReadStream` can still be read from in
	// the same manner as before. Both of the returned `ReadStream` objects can
	// also be read from, and the rows returned by both child streams will be
	// the same as the rows read from the original stream.
	//
	// Moreover, the two child streams will be allocated back-to-back in the
	// original `ReadStream`. Concretely, it is guaranteed that for streams
	// original, primary, and residual, that original[0-j] = primary[0-j] and
	// original[j-n] = residual[0-m] once the streams have been read to
	// completion.
	SplitReadStream(ctx context.Context, in *SplitReadStreamRequest, opts ...grpc.CallOption) (*SplitReadStreamResponse, error)
}

type bigQueryReadClient struct {
	cc grpc.ClientConnInterface
}

func NewBigQueryReadClient(cc grpc.ClientConnInterface) BigQueryReadClient {
	return &bigQueryReadClient{cc}
}

func (c *bigQueryReadClient) CreateReadSession(ctx context.Context, in *CreateReadSessionRequest, opts ...grpc.CallOption) (*ReadSession, error) {
	out := new(ReadSession)
	err := c.cc.Invoke(ctx, "/google.cloud.bigquery.storage.v1.BigQueryRead/CreateReadSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigQueryReadClient) ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigQueryRead_ReadRowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BigQueryRead_serviceDesc.Streams[0], "/google.cloud.bigquery.storage.v1.BigQueryRead/ReadRows", opts...)
	if err != nil {
		return nil, err
	}
	x := &bigQueryReadReadRowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BigQueryRead_ReadRowsClient interface {
	Recv() (*ReadRowsResponse, error)
	grpc.ClientStream
}

type bigQueryReadReadRowsClient struct {
	grpc.ClientStream
}

func (x *bigQueryReadReadRowsClient) Recv() (*ReadRowsResponse, error) {
	m := new(ReadRowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bigQueryReadClient) SplitReadStream(ctx context.Context, in *SplitReadStreamRequest, opts ...grpc.CallOption) (*SplitReadStreamResponse, error) {
	out := new(SplitReadStreamResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.bigquery.storage.v1.BigQueryRead/SplitReadStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigQueryReadServer is the server API for BigQueryRead service.
type BigQueryReadServer interface {
	// Creates a new read session. A read session divides the contents of a
	// BigQuery table into one or more streams, which can then be used to read
	// data from the table. The read session also specifies properties of the
	// data to be read, such as a list of columns or a push-down filter describing
	// the rows to be returned.
	//
	// A particular row can be read by at most one stream. When the caller has
	// reached the end of each stream in the session, then all the data in the
	// table has been read.
	//
	// Data is assigned to each stream such that roughly the same number of
	// rows can be read from each stream. Because the server-side unit for
	// assigning data is collections of rows, the API does not guarantee that
	// each stream will return the same number or rows. Additionally, the
	// limits are enforced based on the number of pre-filtered rows, so some
	// filters can lead to lopsided assignments.
	//
	// Read sessions automatically expire 24 hours after they are created and do
	// not require manual clean-up by the caller.
	CreateReadSession(context.Context, *CreateReadSessionRequest) (*ReadSession, error)
	// Reads rows from the stream in the format prescribed by the ReadSession.
	// Each response contains one or more table rows, up to a maximum of 100 MiB
	// per response; read requests which attempt to read individual rows larger
	// than 100 MiB will fail.
	//
	// Each request also returns a set of stream statistics reflecting the current
	// state of the stream.
	ReadRows(*ReadRowsRequest, BigQueryRead_ReadRowsServer) error
	// Splits a given `ReadStream` into two `ReadStream` objects. These
	// `ReadStream` objects are referred to as the primary and the residual
	// streams of the split. The original `ReadStream` can still be read from in
	// the same manner as before. Both of the returned `ReadStream` objects can
	// also be read from, and the rows returned by both child streams will be
	// the same as the rows read from the original stream.
	//
	// Moreover, the two child streams will be allocated back-to-back in the
	// original `ReadStream`. Concretely, it is guaranteed that for streams
	// original, primary, and residual, that original[0-j] = primary[0-j] and
	// original[j-n] = residual[0-m] once the streams have been read to
	// completion.
	SplitReadStream(context.Context, *SplitReadStreamRequest) (*SplitReadStreamResponse, error)
}

// UnimplementedBigQueryReadServer can be embedded to have forward compatible implementations.
type UnimplementedBigQueryReadServer struct {
}

func (*UnimplementedBigQueryReadServer) CreateReadSession(ctx context.Context, req *CreateReadSessionRequest) (*ReadSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadSession not implemented")
}
func (*UnimplementedBigQueryReadServer) ReadRows(req *ReadRowsRequest, srv BigQueryRead_ReadRowsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadRows not implemented")
}
func (*UnimplementedBigQueryReadServer) SplitReadStream(ctx context.Context, req *SplitReadStreamRequest) (*SplitReadStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitReadStream not implemented")
}

func RegisterBigQueryReadServer(s *grpc.Server, srv BigQueryReadServer) {
	s.RegisterService(&_BigQueryRead_serviceDesc, srv)
}

func _BigQueryRead_CreateReadSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigQueryReadServer).CreateReadSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.bigquery.storage.v1.BigQueryRead/CreateReadSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigQueryReadServer).CreateReadSession(ctx, req.(*CreateReadSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigQueryRead_ReadRows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRowsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BigQueryReadServer).ReadRows(m, &bigQueryReadReadRowsServer{stream})
}

type BigQueryRead_ReadRowsServer interface {
	Send(*ReadRowsResponse) error
	grpc.ServerStream
}

type bigQueryReadReadRowsServer struct {
	grpc.ServerStream
}

func (x *bigQueryReadReadRowsServer) Send(m *ReadRowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BigQueryRead_SplitReadStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitReadStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigQueryReadServer).SplitReadStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.bigquery.storage.v1.BigQueryRead/SplitReadStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigQueryReadServer).SplitReadStream(ctx, req.(*SplitReadStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigQueryRead_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.bigquery.storage.v1.BigQueryRead",
	HandlerType: (*BigQueryReadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReadSession",
			Handler:    _BigQueryRead_CreateReadSession_Handler,
		},
		{
			MethodName: "SplitReadStream",
			Handler:    _BigQueryRead_SplitReadStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRows",
			Handler:       _BigQueryRead_ReadRows_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/cloud/bigquery/storage/v1/storage.proto",
}