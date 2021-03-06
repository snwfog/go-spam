// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: external/services/grpc/definition/spam.proto

package definition

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HealthResponse_Status int32

const (
	HealthResponse_UNKNOWN     HealthResponse_Status = 0
	HealthResponse_SERVING     HealthResponse_Status = 1
	HealthResponse_NOT_SERVING HealthResponse_Status = 2
)

var HealthResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthResponse_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthResponse_Status) String() string {
	return proto.EnumName(HealthResponse_Status_name, int32(x))
}

func (HealthResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7b9b4e504ff007c, []int{1, 0}
}

type HealthRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b9b4e504ff007c, []int{0}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return m.Size()
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (m *HealthRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthResponse struct {
	Status HealthResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=com.stackadapt.HealthResponse_Status" json:"status,omitempty"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b9b4e504ff007c, []int{1}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return m.Size()
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() HealthResponse_Status {
	if m != nil {
		return m.Status
	}
	return HealthResponse_UNKNOWN
}

type Request struct {
	Key   []byte `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b9b4e504ff007c, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Response struct {
	Count uint64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b9b4e504ff007c, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("com.stackadapt.HealthResponse_Status", HealthResponse_Status_name, HealthResponse_Status_value)
	proto.RegisterType((*HealthRequest)(nil), "com.stackadapt.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "com.stackadapt.HealthResponse")
	proto.RegisterType((*Request)(nil), "com.stackadapt.Request")
	proto.RegisterType((*Response)(nil), "com.stackadapt.Response")
}

func init() {
	proto.RegisterFile("external/services/grpc/definition/spam.proto", fileDescriptor_b7b9b4e504ff007c)
}

var fileDescriptor_b7b9b4e504ff007c = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x33, 0xbd, 0xbd, 0xe9, 0xbd, 0xa7, 0x5a, 0xcb, 0x20, 0x18, 0x2a, 0x84, 0x12, 0x14,
	0x2a, 0x48, 0x82, 0xed, 0x52, 0xdc, 0x58, 0xa4, 0x4a, 0x21, 0x85, 0xc4, 0x3f, 0xe0, 0x46, 0xa6,
	0xe9, 0x69, 0x0d, 0x36, 0x99, 0x98, 0x99, 0x88, 0xdd, 0xfb, 0x00, 0x3e, 0x96, 0xcb, 0x2e, 0x5d,
	0x4a, 0xfb, 0x22, 0xd2, 0xa4, 0xa9, 0x56, 0x77, 0xf3, 0x3b, 0xe7, 0x9b, 0xf9, 0xce, 0xcc, 0xc0,
	0x21, 0x3e, 0x4b, 0x8c, 0x43, 0x36, 0xb6, 0x04, 0xc6, 0x4f, 0xbe, 0x87, 0xc2, 0x1a, 0xc5, 0x91,
	0x67, 0x0d, 0x70, 0xe8, 0x87, 0xbe, 0xf4, 0x79, 0x68, 0x89, 0x88, 0x05, 0x66, 0x14, 0x73, 0xc9,
	0x69, 0xc5, 0xe3, 0x81, 0x29, 0x24, 0xf3, 0x1e, 0xd8, 0x80, 0x45, 0xb2, 0xb6, 0x3b, 0xe2, 0x7c,
	0x34, 0x46, 0x2b, 0xed, 0xf6, 0x93, 0xa1, 0x85, 0x41, 0x24, 0x27, 0x19, 0x6c, 0x1c, 0xc0, 0xe6,
	0x39, 0xb2, 0xb1, 0xbc, 0x77, 0xf0, 0x31, 0x41, 0x21, 0xa9, 0x06, 0xa5, 0xa5, 0x44, 0x23, 0x75,
	0xd2, 0xf8, 0xef, 0xe4, 0xd1, 0x78, 0x21, 0x50, 0xc9, 0x59, 0x11, 0xf1, 0x50, 0x20, 0x3d, 0x01,
	0x55, 0x48, 0x26, 0x13, 0x91, 0xb2, 0x95, 0xe6, 0xbe, 0xb9, 0xee, 0x36, 0xd7, 0x79, 0xd3, 0x4d,
	0x61, 0x67, 0xb9, 0xc9, 0x68, 0x81, 0x9a, 0x55, 0x68, 0x19, 0x4a, 0x57, 0x76, 0xd7, 0xee, 0xdd,
	0xd8, 0x55, 0x65, 0x11, 0xdc, 0x33, 0xe7, 0xfa, 0xc2, 0xee, 0x54, 0x09, 0xdd, 0x82, 0xb2, 0xdd,
	0xbb, 0xbc, 0xcb, 0x0b, 0x05, 0xe3, 0x08, 0x4a, 0xf9, 0xac, 0x55, 0xf8, 0xd3, 0xc5, 0x49, 0xea,
	0xde, 0x70, 0x16, 0x4b, 0xba, 0x0d, 0x7f, 0xdb, 0x3c, 0x09, 0xa5, 0x56, 0xa8, 0x93, 0x46, 0xd1,
	0xc9, 0x82, 0x51, 0x87, 0x7f, 0xab, 0x91, 0x57, 0x04, 0xf9, 0x46, 0x34, 0xdb, 0x50, 0x74, 0x23,
	0x16, 0xd0, 0x63, 0x50, 0x5d, 0x94, 0x1d, 0x94, 0x74, 0xe7, 0xe7, 0x55, 0x96, 0xd2, 0x9a, 0xf6,
	0xbb, 0x91, 0x1d, 0x7d, 0xba, 0xf7, 0x36, 0xd3, 0xc9, 0x74, 0xa6, 0x93, 0x8f, 0x99, 0x4e, 0x5e,
	0xe7, 0xba, 0x32, 0x9d, 0xeb, 0xca, 0xfb, 0x5c, 0x57, 0x6e, 0xe1, 0xeb, 0xa7, 0xfa, 0x6a, 0xfa,
	0xf0, 0xad, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0xa7, 0x25, 0x21, 0xd5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpamClient is the client API for Spam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpamClient interface {
	SetGet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type spamClient struct {
	cc *grpc.ClientConn
}

func NewSpamClient(cc *grpc.ClientConn) SpamClient {
	return &spamClient{cc}
}

func (c *spamClient) SetGet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.stackadapt.Spam/SetGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpamServer is the server API for Spam service.
type SpamServer interface {
	SetGet(context.Context, *Request) (*Response, error)
}

// UnimplementedSpamServer can be embedded to have forward compatible implementations.
type UnimplementedSpamServer struct {
}

func (*UnimplementedSpamServer) SetGet(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGet not implemented")
}

func RegisterSpamServer(s *grpc.Server, srv SpamServer) {
	s.RegisterService(&_Spam_serviceDesc, srv)
}

func _Spam_SetGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpamServer).SetGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.stackadapt.Spam/SetGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpamServer).SetGet(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spam_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.stackadapt.Spam",
	HandlerType: (*SpamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGet",
			Handler:    _Spam_SetGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/services/grpc/definition/spam.proto",
}

func (m *HealthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarintSpam(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HealthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSpam(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintSpam(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSpam(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintSpam(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpam(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HealthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovSpam(uint64(l))
	}
	return n
}

func (m *HealthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovSpam(uint64(m.Status))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSpam(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovSpam(uint64(m.Count))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovSpam(uint64(m.Count))
	}
	return n
}

func sovSpam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpam(x uint64) (n int) {
	return sovSpam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSpam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= HealthResponse_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpam
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSpam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpam
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSpam
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSpam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpam = fmt.Errorf("proto: unexpected end of group")
)
