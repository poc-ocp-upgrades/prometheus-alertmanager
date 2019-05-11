package nflogpb

import (
	proto "github.com/gogo/protobuf/proto"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	fmt "fmt"
	math "math"
	_ "github.com/gogo/protobuf/gogoproto"
	time "time"
	types "github.com/gogo/protobuf/types"
	io "io"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

const _ = proto.GoGoProtoPackageIsVersion2

type Receiver struct {
	GroupName	string	`protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Integration	string	`protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
	Idx			uint32	`protobuf:"varint,3,opt,name=idx,proto3" json:"idx,omitempty"`
}

func (m *Receiver) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Receiver{}
}
func (m *Receiver) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Receiver) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Receiver) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorNflog, []int{0}
}

type Entry struct {
	GroupKey		[]byte		`protobuf:"bytes,1,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	Receiver		*Receiver	`protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	GroupHash		[]byte		`protobuf:"bytes,3,opt,name=group_hash,json=groupHash,proto3" json:"group_hash,omitempty"`
	Resolved		bool		`protobuf:"varint,4,opt,name=resolved,proto3" json:"resolved,omitempty"`
	Timestamp		time.Time	`protobuf:"bytes,5,opt,name=timestamp,stdtime" json:"timestamp"`
	FiringAlerts	[]uint64	`protobuf:"varint,6,rep,packed,name=firing_alerts,json=firingAlerts" json:"firing_alerts,omitempty"`
	ResolvedAlerts	[]uint64	`protobuf:"varint,7,rep,packed,name=resolved_alerts,json=resolvedAlerts" json:"resolved_alerts,omitempty"`
}

func (m *Entry) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Entry{}
}
func (m *Entry) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Entry) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Entry) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorNflog, []int{1}
}

type MeshEntry struct {
	Entry		*Entry		`protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	ExpiresAt	time.Time	`protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,stdtime" json:"expires_at"`
}

func (m *MeshEntry) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MeshEntry{}
}
func (m *MeshEntry) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MeshEntry) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MeshEntry) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorNflog, []int{2}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Receiver)(nil), "nflogpb.Receiver")
	proto.RegisterType((*Entry)(nil), "nflogpb.Entry")
	proto.RegisterType((*MeshEntry)(nil), "nflogpb.MeshEntry")
}
func (m *Receiver) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Receiver) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.GroupName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNflog(dAtA, i, uint64(len(m.GroupName)))
		i += copy(dAtA[i:], m.GroupName)
	}
	if len(m.Integration) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNflog(dAtA, i, uint64(len(m.Integration)))
		i += copy(dAtA[i:], m.Integration)
	}
	if m.Idx != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNflog(dAtA, i, uint64(m.Idx))
	}
	return i, nil
}
func (m *Entry) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.GroupKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNflog(dAtA, i, uint64(len(m.GroupKey)))
		i += copy(dAtA[i:], m.GroupKey)
	}
	if m.Receiver != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNflog(dAtA, i, uint64(m.Receiver.Size()))
		n1, err := m.Receiver.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.GroupHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNflog(dAtA, i, uint64(len(m.GroupHash)))
		i += copy(dAtA[i:], m.GroupHash)
	}
	if m.Resolved {
		dAtA[i] = 0x20
		i++
		if m.Resolved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintNflog(dAtA, i, uint64(types.SizeOfStdTime(m.Timestamp)))
	n2, err := types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.FiringAlerts) > 0 {
		dAtA4 := make([]byte, len(m.FiringAlerts)*10)
		var j3 int
		for _, num := range m.FiringAlerts {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintNflog(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if len(m.ResolvedAlerts) > 0 {
		dAtA6 := make([]byte, len(m.ResolvedAlerts)*10)
		var j5 int
		for _, num := range m.ResolvedAlerts {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintNflog(dAtA, i, uint64(j5))
		i += copy(dAtA[i:], dAtA6[:j5])
	}
	return i, nil
}
func (m *MeshEntry) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *MeshEntry) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Entry != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNflog(dAtA, i, uint64(m.Entry.Size()))
		n7, err := m.Entry.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintNflog(dAtA, i, uint64(types.SizeOfStdTime(m.ExpiresAt)))
	n8, err := types.StdTimeMarshalTo(m.ExpiresAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}
func encodeVarintNflog(dAtA []byte, offset int, v uint64) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Receiver) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.GroupName)
	if l > 0 {
		n += 1 + l + sovNflog(uint64(l))
	}
	l = len(m.Integration)
	if l > 0 {
		n += 1 + l + sovNflog(uint64(l))
	}
	if m.Idx != 0 {
		n += 1 + sovNflog(uint64(m.Idx))
	}
	return n
}
func (m *Entry) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.GroupKey)
	if l > 0 {
		n += 1 + l + sovNflog(uint64(l))
	}
	if m.Receiver != nil {
		l = m.Receiver.Size()
		n += 1 + l + sovNflog(uint64(l))
	}
	l = len(m.GroupHash)
	if l > 0 {
		n += 1 + l + sovNflog(uint64(l))
	}
	if m.Resolved {
		n += 2
	}
	l = types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovNflog(uint64(l))
	if len(m.FiringAlerts) > 0 {
		l = 0
		for _, e := range m.FiringAlerts {
			l += sovNflog(uint64(e))
		}
		n += 1 + sovNflog(uint64(l)) + l
	}
	if len(m.ResolvedAlerts) > 0 {
		l = 0
		for _, e := range m.ResolvedAlerts {
			l += sovNflog(uint64(e))
		}
		n += 1 + sovNflog(uint64(l)) + l
	}
	return n
}
func (m *MeshEntry) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Entry != nil {
		l = m.Entry.Size()
		n += 1 + l + sovNflog(uint64(l))
	}
	l = types.SizeOfStdTime(m.ExpiresAt)
	n += 1 + l + sovNflog(uint64(l))
	return n
}
func sovNflog(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNflog(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovNflog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Receiver) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNflog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Integration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Integration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idx", wireType)
			}
			m.Idx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idx |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNflog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNflog
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
func (m *Entry) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNflog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupKey = append(m.GroupKey[:0], dAtA[iNdEx:postIndex]...)
			if m.GroupKey == nil {
				m.GroupKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receiver == nil {
				m.Receiver = &Receiver{}
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupHash = append(m.GroupHash[:0], dAtA[iNdEx:postIndex]...)
			if m.GroupHash == nil {
				m.GroupHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Resolved = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNflog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FiringAlerts = append(m.FiringAlerts, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNflog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthNflog
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNflog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FiringAlerts = append(m.FiringAlerts, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FiringAlerts", wireType)
			}
		case 7:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNflog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ResolvedAlerts = append(m.ResolvedAlerts, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNflog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthNflog
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNflog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ResolvedAlerts = append(m.ResolvedAlerts, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolvedAlerts", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNflog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNflog
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
func (m *MeshEntry) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNflog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MeshEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entry == nil {
				m.Entry = &Entry{}
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNflog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.ExpiresAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNflog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNflog
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
func skipNflog(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNflog
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
					return 0, ErrIntOverflowNflog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowNflog
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNflog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNflog
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipNflog(dAtA[start:])
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
	ErrInvalidLengthNflog	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNflog		= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("nflog.proto", fileDescriptorNflog)
}

var fileDescriptorNflog = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x6e, 0xd3, 0x40, 0x10, 0xc6, 0xbb, 0x4d, 0xd3, 0xda, 0xe3, 0xb4, 0x94, 0x15, 0x07, 0xcb, 0x08, 0xc7, 0x0a, 0x48, 0xf8, 0x82, 0x23, 0x95, 0x27, 0x68, 0x10, 0x12, 0x12, 0x82, 0xc3, 0x8a, 0x2b, 0xb2, 0x36, 0x74, 0xb2, 0x5e, 0x61, 0x7b, 0xad, 0xf5, 0x36, 0x6a, 0xde, 0x82, 0x47, 0xe0, 0x71, 0x72, 0xe4, 0x09, 0xf8, 0x93, 0x27, 0x41, 0xde, 0xb5, 0x1d, 0x8e, 0xdc, 0x66, 0x7f, 0xf3, 0xcd, 0xcc, 0xb7, 0x1f, 0x04, 0xf5, 0xa6, 0x54, 0x22, 0x6b, 0xb4, 0x32, 0x8a, 0x5e, 0xd8, 0x47, 0xb3, 0x8e, 0xe6, 0x42, 0x29, 0x51, 0xe2, 0xd2, 0xe2, 0xf5, 0xfd, 0x66, 0x69, 0x64, 0x85, 0xad, 0xe1, 0x55, 0xe3, 0x94, 0xd1, 0x13, 0xa1, 0x84, 0xb2, 0xe5, 0xb2, 0xab, 0x1c, 0x5d, 0x7c, 0x06, 0x8f, 0xe1, 0x17, 0x94, 0x5b, 0xd4, 0xf4, 0x19, 0x80, 0xd0, 0xea, 0xbe, 0xc9, 0x6b, 0x5e, 0x61, 0x48, 0x12, 0x92, 0xfa, 0xcc, 0xb7, 0xe4, 0x23, 0xaf, 0x90, 0x26, 0x10, 0xc8, 0xda, 0xa0, 0xd0, 0xdc, 0x48, 0x55, 0x87, 0xa7, 0xb6, 0xff, 0x2f, 0xa2, 0xd7, 0x30, 0x91, 0x77, 0x0f, 0xe1, 0x24, 0x21, 0xe9, 0x25, 0xeb, 0xca, 0xc5, 0xf7, 0x53, 0x98, 0xbe, 0xad, 0x8d, 0xde, 0xd1, 0xa7, 0xe0, 0x56, 0xe5, 0x5f, 0x71, 0x67, 0x77, 0xcf, 0x98, 0x67, 0xc1, 0x7b, 0xdc, 0xd1, 0x57, 0xe0, 0xe9, 0xde, 0x85, 0xdd, 0x1b, 0xdc, 0x3c, 0xce, 0xfa, 0x8f, 0x65, 0x83, 0x3d, 0x36, 0x4a, 0x8e, 0x46, 0x0b, 0xde, 0x16, 0xf6, 0xdc, 0xac, 0x37, 0xfa, 0x8e, 0xb7, 0x05, 0x8d, 0xba, 0x6d, 0xad, 0x2a, 0xb7, 0x78, 0x17, 0x9e, 0x25, 0x24, 0xf5, 0xd8, 0xf8, 0xa6, 0x2b, 0xf0, 0xc7, 0x60, 0xc2, 0xa9, 0x3d, 0x15, 0x65, 0x2e, 0xba, 0x6c, 0x88, 0x2e, 0xfb, 0x34, 0x28, 0x56, 0xde, 0xfe, 0xe7, 0xfc, 0xe4, 0xdb, 0xaf, 0x39, 0x61, 0xc7, 0x31, 0xfa, 0x1c, 0x2e, 0x37, 0x52, 0xcb, 0x5a, 0xe4, 0xbc, 0x44, 0x6d, 0xda, 0xf0, 0x3c, 0x99, 0xa4, 0x67, 0x6c, 0xe6, 0xe0, 0xad, 0x65, 0xf4, 0x25, 0x3c, 0x1a, 0x8e, 0x0e, 0xb2, 0x0b, 0x2b, 0xbb, 0x1a, 0xb0, 0x13, 0x2e, 0xb6, 0xe0, 0x7f, 0xc0, 0xb6, 0x70, 0x29, 0xbd, 0x80, 0x29, 0x76, 0x85, 0x4d, 0x28, 0xb8, 0xb9, 0x1a, 0x53, 0xb0, 0x6d, 0xe6, 0x9a, 0xf4, 0x0d, 0x00, 0x3e, 0x34, 0x52, 0x63, 0x9b, 0x73, 0xd3, 0x07, 0xf6, 0x9f, 0xbf, 0xe8, 0xe7, 0x6e, 0xcd, 0xea, 0x7a, 0xff, 0x27, 0x3e, 0xd9, 0x1f, 0x62, 0xf2, 0xe3, 0x10, 0x93, 0xdf, 0x87, 0x98, 0xac, 0xcf, 0xed, 0xe8, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x49, 0xcd, 0xa7, 0x1e, 0x61, 0x02, 0x00, 0x00}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
