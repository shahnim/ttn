// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto
// DO NOT EDIT!

/*
	Package networkserver is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto

	It has these top-level messages:
		DevicesRequest
		DevicesResponse
		StatusRequest
		Status
*/
package networkserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/TheThingsNetwork/ttn/api"
import lorawan "github.com/TheThingsNetwork/ttn/api/protocol/lorawan"
import broker "github.com/TheThingsNetwork/ttn/api/broker"
import handler "github.com/TheThingsNetwork/ttn/api/handler"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevicesRequest struct {
	DevAddr *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
	FCnt    uint32                                              `protobuf:"varint,2,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *DevicesRequest) Reset()                    { *m = DevicesRequest{} }
func (m *DevicesRequest) String() string            { return proto.CompactTextString(m) }
func (*DevicesRequest) ProtoMessage()               {}
func (*DevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{0} }

type DevicesResponse struct {
	Results []*lorawan.Device `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *DevicesResponse) Reset()                    { *m = DevicesResponse{} }
func (m *DevicesResponse) String() string            { return proto.CompactTextString(m) }
func (*DevicesResponse) ProtoMessage()               {}
func (*DevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{1} }

func (m *DevicesResponse) GetResults() []*lorawan.Device {
	if m != nil {
		return m.Results
	}
	return nil
}

// message StatusRequest is used to request the status of this NetworkServer
type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{2} }

// message Status is the response to the StatusRequest
type Status struct {
	System            *api.SystemStats    `protobuf:"bytes,1,opt,name=system" json:"system,omitempty"`
	Component         *api.ComponentStats `protobuf:"bytes,2,opt,name=component" json:"component,omitempty"`
	Uplink            *api.Rates          `protobuf:"bytes,11,opt,name=uplink" json:"uplink,omitempty"`
	Downlink          *api.Rates          `protobuf:"bytes,12,opt,name=downlink" json:"downlink,omitempty"`
	Activations       *api.Rates          `protobuf:"bytes,13,opt,name=activations" json:"activations,omitempty"`
	DevicesPerAddress *api.Percentiles    `protobuf:"bytes,21,opt,name=devices_per_address,json=devicesPerAddress" json:"devices_per_address,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{3} }

func (m *Status) GetSystem() *api.SystemStats {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *Status) GetComponent() *api.ComponentStats {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *Status) GetUplink() *api.Rates {
	if m != nil {
		return m.Uplink
	}
	return nil
}

func (m *Status) GetDownlink() *api.Rates {
	if m != nil {
		return m.Downlink
	}
	return nil
}

func (m *Status) GetActivations() *api.Rates {
	if m != nil {
		return m.Activations
	}
	return nil
}

func (m *Status) GetDevicesPerAddress() *api.Percentiles {
	if m != nil {
		return m.DevicesPerAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*DevicesRequest)(nil), "networkserver.DevicesRequest")
	proto.RegisterType((*DevicesResponse)(nil), "networkserver.DevicesResponse")
	proto.RegisterType((*StatusRequest)(nil), "networkserver.StatusRequest")
	proto.RegisterType((*Status)(nil), "networkserver.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for NetworkServer service

type NetworkServerClient interface {
	// Broker requests devices with DevAddr for MIC check
	GetDevices(ctx context.Context, in *DevicesRequest, opts ...grpc.CallOption) (*DevicesResponse, error)
	// Broker requests device activation "template" from Network Server
	PrepareActivation(ctx context.Context, in *broker.DeduplicatedDeviceActivationRequest, opts ...grpc.CallOption) (*broker.DeduplicatedDeviceActivationRequest, error)
	// Broker confirms device activation
	Activate(ctx context.Context, in *handler.DeviceActivationResponse, opts ...grpc.CallOption) (*handler.DeviceActivationResponse, error)
	// Broker informs Network Server about Uplink
	Uplink(ctx context.Context, in *broker.DeduplicatedUplinkMessage, opts ...grpc.CallOption) (*broker.DeduplicatedUplinkMessage, error)
	// Broker informs Network Server about Downlink, NetworkServer should sign, add MIC, ...
	Downlink(ctx context.Context, in *broker.DownlinkMessage, opts ...grpc.CallOption) (*broker.DownlinkMessage, error)
}

type networkServerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerClient(cc *grpc.ClientConn) NetworkServerClient {
	return &networkServerClient{cc}
}

func (c *networkServerClient) GetDevices(ctx context.Context, in *DevicesRequest, opts ...grpc.CallOption) (*DevicesResponse, error) {
	out := new(DevicesResponse)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/GetDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) PrepareActivation(ctx context.Context, in *broker.DeduplicatedDeviceActivationRequest, opts ...grpc.CallOption) (*broker.DeduplicatedDeviceActivationRequest, error) {
	out := new(broker.DeduplicatedDeviceActivationRequest)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/PrepareActivation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Activate(ctx context.Context, in *handler.DeviceActivationResponse, opts ...grpc.CallOption) (*handler.DeviceActivationResponse, error) {
	out := new(handler.DeviceActivationResponse)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Uplink(ctx context.Context, in *broker.DeduplicatedUplinkMessage, opts ...grpc.CallOption) (*broker.DeduplicatedUplinkMessage, error) {
	out := new(broker.DeduplicatedUplinkMessage)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Uplink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Downlink(ctx context.Context, in *broker.DownlinkMessage, opts ...grpc.CallOption) (*broker.DownlinkMessage, error) {
	out := new(broker.DownlinkMessage)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Downlink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServer service

type NetworkServerServer interface {
	// Broker requests devices with DevAddr for MIC check
	GetDevices(context.Context, *DevicesRequest) (*DevicesResponse, error)
	// Broker requests device activation "template" from Network Server
	PrepareActivation(context.Context, *broker.DeduplicatedDeviceActivationRequest) (*broker.DeduplicatedDeviceActivationRequest, error)
	// Broker confirms device activation
	Activate(context.Context, *handler.DeviceActivationResponse) (*handler.DeviceActivationResponse, error)
	// Broker informs Network Server about Uplink
	Uplink(context.Context, *broker.DeduplicatedUplinkMessage) (*broker.DeduplicatedUplinkMessage, error)
	// Broker informs Network Server about Downlink, NetworkServer should sign, add MIC, ...
	Downlink(context.Context, *broker.DownlinkMessage) (*broker.DownlinkMessage, error)
}

func RegisterNetworkServerServer(s *grpc.Server, srv NetworkServerServer) {
	s.RegisterService(&_NetworkServer_serviceDesc, srv)
}

func _NetworkServer_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).GetDevices(ctx, req.(*DevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_PrepareActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DeduplicatedDeviceActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).PrepareActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/PrepareActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).PrepareActivation(ctx, req.(*broker.DeduplicatedDeviceActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(handler.DeviceActivationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Activate(ctx, req.(*handler.DeviceActivationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Uplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DeduplicatedUplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Uplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Uplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Uplink(ctx, req.(*broker.DeduplicatedUplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Downlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Downlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Downlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Downlink(ctx, req.(*broker.DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkserver.NetworkServer",
	HandlerType: (*NetworkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDevices",
			Handler:    _NetworkServer_GetDevices_Handler,
		},
		{
			MethodName: "PrepareActivation",
			Handler:    _NetworkServer_PrepareActivation_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _NetworkServer_Activate_Handler,
		},
		{
			MethodName: "Uplink",
			Handler:    _NetworkServer_Uplink_Handler,
		},
		{
			MethodName: "Downlink",
			Handler:    _NetworkServer_Downlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorNetworkserver,
}

// Client API for NetworkServerManager service

type NetworkServerManagerClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Status, error)
}

type networkServerManagerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerManagerClient(cc *grpc.ClientConn) NetworkServerManagerClient {
	return &networkServerManagerClient{cc}
}

func (c *networkServerManagerClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServerManager/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServerManager service

type NetworkServerManagerServer interface {
	GetStatus(context.Context, *StatusRequest) (*Status, error)
}

func RegisterNetworkServerManagerServer(s *grpc.Server, srv NetworkServerManagerServer) {
	s.RegisterService(&_NetworkServerManager_serviceDesc, srv)
}

func _NetworkServerManager_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerManagerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServerManager/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerManagerServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkserver.NetworkServerManager",
	HandlerType: (*NetworkServerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _NetworkServerManager_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorNetworkserver,
}

func (m *DevicesRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DevicesRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DevAddr != nil {
		data[i] = 0xa
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.DevAddr.Size()))
		n1, err := m.DevAddr.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FCnt != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *DevicesResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DevicesResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0xa
			i++
			i = encodeVarintNetworkserver(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StatusRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Status) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Status) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.System != nil {
		data[i] = 0xa
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.System.Size()))
		n2, err := m.System.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Component != nil {
		data[i] = 0x12
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.Component.Size()))
		n3, err := m.Component.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Uplink != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.Uplink.Size()))
		n4, err := m.Uplink.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Downlink != nil {
		data[i] = 0x62
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.Downlink.Size()))
		n5, err := m.Downlink.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Activations != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.Activations.Size()))
		n6, err := m.Activations.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.DevicesPerAddress != nil {
		data[i] = 0xaa
		i++
		data[i] = 0x1
		i++
		i = encodeVarintNetworkserver(data, i, uint64(m.DevicesPerAddress.Size()))
		n7, err := m.DevicesPerAddress.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeFixed64Networkserver(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Networkserver(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNetworkserver(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DevicesRequest) Size() (n int) {
	var l int
	_ = l
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovNetworkserver(uint64(m.FCnt))
	}
	return n
}

func (m *DevicesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovNetworkserver(uint64(l))
		}
	}
	return n
}

func (m *StatusRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Status) Size() (n int) {
	var l int
	_ = l
	if m.System != nil {
		l = m.System.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Component != nil {
		l = m.Component.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Uplink != nil {
		l = m.Uplink.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Downlink != nil {
		l = m.Downlink.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Activations != nil {
		l = m.Activations.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.DevicesPerAddress != nil {
		l = m.DevicesPerAddress.Size()
		n += 2 + l + sovNetworkserver(uint64(l))
	}
	return n
}

func sovNetworkserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkserver(x uint64) (n int) {
	return sovNetworkserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicesRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *DevicesResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &lorawan.Device{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *StatusRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *Status) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field System", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.System == nil {
				m.System = &api.SystemStats{}
			}
			if err := m.System.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Component == nil {
				m.Component = &api.ComponentStats{}
			}
			if err := m.Component.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uplink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Uplink == nil {
				m.Uplink = &api.Rates{}
			}
			if err := m.Uplink.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downlink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Downlink == nil {
				m.Downlink = &api.Rates{}
			}
			if err := m.Downlink.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Activations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Activations == nil {
				m.Activations = &api.Rates{}
			}
			if err := m.Activations.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevicesPerAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DevicesPerAddress == nil {
				m.DevicesPerAddress = &api.Percentiles{}
			}
			if err := m.DevicesPerAddress.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func skipNetworkserver(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNetworkserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetworkserver(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetworkserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto", fileDescriptorNetworkserver)
}

var fileDescriptorNetworkserver = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x41, 0x4f, 0x13, 0x41,
	0x14, 0xc7, 0x5d, 0xd0, 0x52, 0x5e, 0xa9, 0xc0, 0x20, 0xb1, 0x69, 0xb4, 0x42, 0x0f, 0xa6, 0x46,
	0xdd, 0x8d, 0x35, 0xd1, 0x98, 0x70, 0xa0, 0x50, 0xc3, 0xc1, 0x40, 0xea, 0x82, 0x89, 0x37, 0x32,
	0xdd, 0x7d, 0x6c, 0x37, 0x6c, 0x67, 0xd6, 0x99, 0xd9, 0x22, 0xdf, 0xc4, 0xab, 0x67, 0xbf, 0x88,
	0x47, 0xcf, 0x1e, 0x8c, 0xc1, 0x2f, 0x62, 0x3a, 0x33, 0x5b, 0x58, 0x81, 0x50, 0x4f, 0xbb, 0xef,
	0xfd, 0x7f, 0xef, 0xcd, 0xeb, 0xfe, 0xdf, 0x14, 0xde, 0x46, 0xb1, 0x1a, 0x64, 0x7d, 0x37, 0xe0,
	0x43, 0xef, 0x60, 0x80, 0x07, 0x83, 0x98, 0x45, 0x72, 0x0f, 0xd5, 0x09, 0x17, 0xc7, 0x9e, 0x52,
	0xcc, 0xa3, 0x69, 0xec, 0x31, 0x13, 0x4b, 0x14, 0x23, 0x14, 0xc5, 0xc8, 0x4d, 0x05, 0x57, 0x9c,
	0x54, 0x0b, 0xc9, 0xfa, 0xf3, 0x0b, 0x5d, 0x23, 0x1e, 0x71, 0x4f, 0x53, 0xfd, 0xec, 0x48, 0x47,
	0x3a, 0xd0, 0x6f, 0xa6, 0xba, 0x80, 0x5f, 0x3b, 0x04, 0x4d, 0x63, 0x8b, 0x77, 0xa6, 0xc1, 0x35,
	0x1a, 0xf0, 0xc4, 0x4b, 0xb8, 0xa0, 0x27, 0x94, 0x79, 0x21, 0x8e, 0xe2, 0x00, 0x6d, 0x8b, 0xd7,
	0xd3, 0xb4, 0xe8, 0x0b, 0x7e, 0x8c, 0xc2, 0x3e, 0x6c, 0xe1, 0x9b, 0x69, 0x0a, 0x07, 0x94, 0x85,
	0x09, 0x8a, 0xfc, 0x69, 0x4a, 0x9b, 0x9f, 0xe1, 0x6e, 0x57, 0xcf, 0x20, 0x7d, 0xfc, 0x94, 0xa1,
	0x54, 0xe4, 0x3d, 0x94, 0x43, 0x1c, 0x1d, 0xd2, 0x30, 0x14, 0x35, 0x67, 0xcd, 0x69, 0x2d, 0x6c,
	0xbd, 0xfa, 0xf9, 0xeb, 0x51, 0xfb, 0xa6, 0x23, 0x02, 0x2e, 0xd0, 0x53, 0xa7, 0x29, 0x4a, 0xb7,
	0x8b, 0xa3, 0x4e, 0x18, 0x0a, 0x7f, 0x2e, 0x34, 0x2f, 0x64, 0x05, 0xee, 0x1c, 0x1d, 0x06, 0x4c,
	0xd5, 0x66, 0xd6, 0x9c, 0x56, 0xd5, 0xbf, 0x7d, 0xb4, 0xcd, 0x54, 0x73, 0x03, 0x16, 0x27, 0x27,
	0xcb, 0x94, 0x33, 0x89, 0xe4, 0x09, 0xcc, 0x09, 0x94, 0x59, 0xa2, 0x64, 0xcd, 0x59, 0x9b, 0x6d,
	0x55, 0xda, 0x8b, 0xae, 0xfd, 0x50, 0xae, 0x41, 0xfd, 0x5c, 0x6f, 0x2e, 0x42, 0x75, 0x5f, 0x51,
	0x95, 0xe5, 0x63, 0x37, 0xbf, 0xce, 0x40, 0xc9, 0x64, 0x48, 0x0b, 0x4a, 0xf2, 0x54, 0x2a, 0x1c,
	0xea, 0xf9, 0x2b, 0xed, 0x25, 0x77, 0x6c, 0xd3, 0xbe, 0x4e, 0x8d, 0x11, 0xe9, 0x5b, 0x9d, 0xbc,
	0x80, 0xf9, 0x80, 0x0f, 0x53, 0xce, 0xd0, 0x0e, 0x57, 0x69, 0xaf, 0x68, 0x78, 0x3b, 0xcf, 0x1a,
	0xfe, 0x9c, 0x22, 0x4d, 0x28, 0x65, 0x69, 0x12, 0xb3, 0xe3, 0x5a, 0x45, 0xf3, 0xa0, 0x79, 0x9f,
	0x2a, 0x94, 0xbe, 0x55, 0xc8, 0x63, 0x28, 0x87, 0xfc, 0x84, 0x69, 0x6a, 0xe1, 0x12, 0x35, 0xd1,
	0xc8, 0x33, 0xa8, 0xd0, 0x40, 0xc5, 0x23, 0xaa, 0x62, 0xce, 0x64, 0xad, 0x7a, 0x09, 0xbd, 0x28,
	0x93, 0x4d, 0x58, 0x31, 0xeb, 0x22, 0x0f, 0x53, 0x14, 0xda, 0x20, 0x94, 0xb2, 0xb6, 0x7a, 0xe1,
	0x37, 0xf6, 0x50, 0x04, 0xc8, 0x54, 0x9c, 0xa0, 0xf4, 0x97, 0x2d, 0xdc, 0x43, 0xd1, 0x31, 0x68,
	0xfb, 0xdb, 0x2c, 0x54, 0xad, 0x67, 0xfb, 0xfa, 0x4e, 0x90, 0x77, 0x00, 0x3b, 0xa8, 0xac, 0x0f,
	0xe4, 0xa1, 0x5b, 0xbc, 0x46, 0xc5, 0xcd, 0xa8, 0x37, 0xae, 0x93, 0xad, 0x7d, 0x43, 0x58, 0xee,
	0x09, 0x4c, 0xa9, 0xc0, 0xce, 0x64, 0x6c, 0xf2, 0xd4, 0xb5, 0xab, 0xda, 0xc5, 0x70, 0xfc, 0x79,
	0x02, 0xaa, 0x30, 0x34, 0x95, 0xe7, 0x54, 0x7e, 0xc2, 0xff, 0xc0, 0xa4, 0x07, 0x65, 0x9b, 0x44,
	0xb2, 0xee, 0xe6, 0x6b, 0x7d, 0x99, 0x36, 0xd3, 0xd5, 0x6f, 0x46, 0xc8, 0x1e, 0x94, 0x3e, 0x18,
	0x07, 0xd7, 0xaf, 0x1a, 0xc4, 0x68, 0xbb, 0x28, 0x25, 0x8d, 0xc6, 0xfd, 0x6e, 0x42, 0xc8, 0x06,
	0x94, 0xbb, 0xb9, 0xd7, 0xf7, 0x27, 0xb8, 0xcd, 0xe4, 0x7d, 0xae, 0x13, 0xda, 0x1f, 0xe1, 0x5e,
	0xc1, 0xac, 0x5d, 0xca, 0x68, 0x84, 0x82, 0x6c, 0xc2, 0xfc, 0x0e, 0x2a, 0xbb, 0xeb, 0x0f, 0xfe,
	0xf1, 0xa4, 0x70, 0x29, 0xea, 0xab, 0x57, 0xaa, 0x5b, 0x4b, 0xdf, 0xcf, 0x1a, 0xce, 0x8f, 0xb3,
	0x86, 0xf3, 0xfb, 0xac, 0xe1, 0x7c, 0xf9, 0xd3, 0xb8, 0xd5, 0x2f, 0xe9, 0x7f, 0x83, 0x97, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x66, 0xb2, 0x91, 0x81, 0x7a, 0x05, 0x00, 0x00,
}
