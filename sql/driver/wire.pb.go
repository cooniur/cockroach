// Code generated by protoc-gen-gogo.
// source: cockroach/sql/driver/wire.proto
// DO NOT EDIT!

/*
	Package driver is a generated protocol buffer package.

	It is generated from these files:
		cockroach/sql/driver/wire.proto

	It has these top-level messages:
		RequestHeader
		ResponseHeader
		Datum
		Result
		Request
		Response
*/
package driver

import proto "github.com/gogo/protobuf/proto"
import math "math"
import cockroach_proto3 "github.com/cockroachdb/cockroach/proto"
import cockroach_proto2 "github.com/cockroachdb/cockroach/proto"

// discarding unused import gogoproto "gogoproto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// RequestHeader is supplied with every Request.
type RequestHeader struct {
	// User is the originating user.
	User string `protobuf:"bytes,5,opt,name=user" json:"user"`
	// Session settings that were returned in the last response that
	// contained them, being reflected back to the server.
	Session []byte `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// The transaction state returned in the previous response being
	// reflected back.
	Txn []byte `protobuf:"bytes,2,opt,name=txn" json:"txn,omitempty"`
	// CmdID is optionally specified for request idempotence
	// (i.e. replay protection).
	CmdID cockroach_proto3.ClientCmdID `protobuf:"bytes,3,opt,name=cmd_id" json:"cmd_id"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}

func (m *RequestHeader) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *RequestHeader) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *RequestHeader) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *RequestHeader) GetCmdID() cockroach_proto3.ClientCmdID {
	if m != nil {
		return m.CmdID
	}
	return cockroach_proto3.ClientCmdID{}
}

// ResponseHeader is returned with every Response.
type ResponseHeader struct {
	// Error is non-nil if an error occurred.
	Error *cockroach_proto2.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// Setting that should be reflected back in all subsequent requests.
	// When not set, future requests should continue to use existing settings.
	Session []byte `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	// Transaction message returned in a response; not to be interpreted by
	// the recipient and reflected in a subsequent request. When not set,
	// the subsequent request should not contain a transaction object.
	Txn []byte `protobuf:"bytes,3,opt,name=txn" json:"txn,omitempty"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}

func (m *ResponseHeader) GetError() *cockroach_proto2.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ResponseHeader) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ResponseHeader) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

type Datum struct {
	BoolVal   *bool    `protobuf:"varint,1,opt,name=bool_val" json:"bool_val,omitempty"`
	IntVal    *int64   `protobuf:"varint,2,opt,name=int_val" json:"int_val,omitempty"`
	FloatVal  *float64 `protobuf:"fixed64,3,opt,name=float_val" json:"float_val,omitempty"`
	BytesVal  []byte   `protobuf:"bytes,4,opt,name=bytes_val" json:"bytes_val,omitempty"`
	StringVal *string  `protobuf:"bytes,5,opt,name=string_val" json:"string_val,omitempty"`
}

func (m *Datum) Reset()      { *m = Datum{} }
func (*Datum) ProtoMessage() {}

func (m *Datum) GetBoolVal() bool {
	if m != nil && m.BoolVal != nil {
		return *m.BoolVal
	}
	return false
}

func (m *Datum) GetIntVal() int64 {
	if m != nil && m.IntVal != nil {
		return *m.IntVal
	}
	return 0
}

func (m *Datum) GetFloatVal() float64 {
	if m != nil && m.FloatVal != nil {
		return *m.FloatVal
	}
	return 0
}

func (m *Datum) GetBytesVal() []byte {
	if m != nil {
		return m.BytesVal
	}
	return nil
}

func (m *Datum) GetStringVal() string {
	if m != nil && m.StringVal != nil {
		return *m.StringVal
	}
	return ""
}

// A Result is a collection of rows.
type Result struct {
	// The names of the columns returned in the result set in the order specified
	// in the SQL statement. The number of columns will equal the number of
	// values in each Row.
	Columns []string `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	// The rows in the result set.
	Rows []Result_Row `protobuf:"bytes,2,rep,name=rows" json:"rows"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

func (m *Result) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Result) GetRows() []Result_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// A Row is a collection of values representing a row in a result.
type Result_Row struct {
	Values []Datum `protobuf:"bytes,1,rep,name=values" json:"values"`
}

func (m *Result_Row) Reset()         { *m = Result_Row{} }
func (m *Result_Row) String() string { return proto.CompactTextString(m) }
func (*Result_Row) ProtoMessage()    {}

func (m *Result_Row) GetValues() []Datum {
	if m != nil {
		return m.Values
	}
	return nil
}

// An SQL request to cockroach. A transaction can consist of multiple
// requests.
type Request struct {
	// Request header.
	RequestHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	// SQL statement(s) to be serially executed by the server. Multiple
	// statements are passed as a single string separated by semicolons.
	Sql string `protobuf:"bytes,2,opt,name=sql" json:"sql"`
	// Parameters referred to in the above SQL statement(s) using "?".
	Params []Datum `protobuf:"bytes,3,rep,name=params" json:"params"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *Request) GetParams() []Datum {
	if m != nil {
		return m.Params
	}
	return nil
}

type Response struct {
	ResponseHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	// The list of results. There is one result object per SQL statement in the
	// request.
	Results []Result `protobuf:"bytes,2,rep,name=results" json:"results"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetResults() []Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *RequestHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RequestHeader) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Session != nil {
		data[i] = 0xa
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	if m.Txn != nil {
		data[i] = 0x12
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Txn)))
		i += copy(data[i:], m.Txn)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintWire(data, i, uint64(m.CmdID.Size()))
	n1, err := m.CmdID.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x2a
	i++
	i = encodeVarintWire(data, i, uint64(len(m.User)))
	i += copy(data[i:], m.User)
	return i, nil
}

func (m *ResponseHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ResponseHeader) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintWire(data, i, uint64(m.Error.Size()))
		n2, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Session != nil {
		data[i] = 0x12
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	if m.Txn != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintWire(data, i, uint64(len(m.Txn)))
		i += copy(data[i:], m.Txn)
	}
	return i, nil
}

func (m *Datum) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Datum) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BoolVal != nil {
		data[i] = 0x8
		i++
		if *m.BoolVal {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.IntVal != nil {
		data[i] = 0x10
		i++
		i = encodeVarintWire(data, i, uint64(*m.IntVal))
	}
	if m.FloatVal != nil {
		data[i] = 0x19
		i++
		i = encodeFixed64Wire(data, i, uint64(math.Float64bits(*m.FloatVal)))
	}
	if m.BytesVal != nil {
		data[i] = 0x22
		i++
		i = encodeVarintWire(data, i, uint64(len(m.BytesVal)))
		i += copy(data[i:], m.BytesVal)
	}
	if m.StringVal != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintWire(data, i, uint64(len(*m.StringVal)))
		i += copy(data[i:], *m.StringVal)
	}
	return i, nil
}

func (m *Result) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Result) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Rows) > 0 {
		for _, msg := range m.Rows {
			data[i] = 0x12
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Result_Row) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Result_Row) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintWire(data, i, uint64(m.RequestHeader.Size()))
	n3, err := m.RequestHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x12
	i++
	i = encodeVarintWire(data, i, uint64(len(m.Sql)))
	i += copy(data[i:], m.Sql)
	if len(m.Params) > 0 {
		for _, msg := range m.Params {
			data[i] = 0x1a
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintWire(data, i, uint64(m.ResponseHeader.Size()))
	n4, err := m.ResponseHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0x12
			i++
			i = encodeVarintWire(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Wire(data []byte, offset int, v uint64) int {
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
func encodeFixed32Wire(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWire(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RequestHeader) Size() (n int) {
	var l int
	_ = l
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovWire(uint64(l))
	}
	if m.Txn != nil {
		l = len(m.Txn)
		n += 1 + l + sovWire(uint64(l))
	}
	l = m.CmdID.Size()
	n += 1 + l + sovWire(uint64(l))
	l = len(m.User)
	n += 1 + l + sovWire(uint64(l))
	return n
}

func (m *ResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovWire(uint64(l))
	}
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovWire(uint64(l))
	}
	if m.Txn != nil {
		l = len(m.Txn)
		n += 1 + l + sovWire(uint64(l))
	}
	return n
}

func (m *Datum) Size() (n int) {
	var l int
	_ = l
	if m.BoolVal != nil {
		n += 2
	}
	if m.IntVal != nil {
		n += 1 + sovWire(uint64(*m.IntVal))
	}
	if m.FloatVal != nil {
		n += 9
	}
	if m.BytesVal != nil {
		l = len(m.BytesVal)
		n += 1 + l + sovWire(uint64(l))
	}
	if m.StringVal != nil {
		l = len(*m.StringVal)
		n += 1 + l + sovWire(uint64(l))
	}
	return n
}

func (m *Result) Size() (n int) {
	var l int
	_ = l
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			l = len(s)
			n += 1 + l + sovWire(uint64(l))
		}
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Result_Row) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Request) Size() (n int) {
	var l int
	_ = l
	l = m.RequestHeader.Size()
	n += 1 + l + sovWire(uint64(l))
	l = len(m.Sql)
	n += 1 + l + sovWire(uint64(l))
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	l = m.ResponseHeader.Size()
	n += 1 + l + sovWire(uint64(l))
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func sovWire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Datum) GetValue() interface{} {
	if this.BoolVal != nil {
		return this.BoolVal
	}
	if this.IntVal != nil {
		return this.IntVal
	}
	if this.FloatVal != nil {
		return this.FloatVal
	}
	if this.BytesVal != nil {
		return this.BytesVal
	}
	if this.StringVal != nil {
		return this.StringVal
	}
	return nil
}

func (this *Datum) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *bool:
		this.BoolVal = vt
	case *int64:
		this.IntVal = vt
	case *float64:
		this.FloatVal = vt
	case []byte:
		this.BytesVal = vt
	case *string:
		this.StringVal = vt
	default:
		return false
	}
	return true
}
func (m *RequestHeader) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txn = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CmdID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CmdID.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ResponseHeader) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &cockroach_proto2.Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txn = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Datum) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolVal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.BoolVal = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IntVal = &v
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatVal", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.FloatVal = &v2
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesVal = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.StringVal = &s
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Result) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, Result_Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Result_Row) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, Datum{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequestHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sql = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, Datum{})
			if err := m.Params[len(m.Params)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ResponseHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
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
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipWire(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
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
			for {
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
				return 0, ErrInvalidLengthWire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
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
				next, err := skipWire(data[start:])
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
	ErrInvalidLengthWire = fmt.Errorf("proto: negative length found during unmarshaling")
)
