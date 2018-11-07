// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/v2/table.proto

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf4 "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Possible timestamp granularities to use when keeping multiple versions
// of data in a table.
type Table_TimestampGranularity int32

const (
	// The user did not specify a granularity. Should not be returned.
	// When specified during table creation, MILLIS will be used.
	Table_TIMESTAMP_GRANULARITY_UNSPECIFIED Table_TimestampGranularity = 0
	// The table keeps data versioned at a granularity of 1ms.
	Table_MILLIS Table_TimestampGranularity = 1
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "TIMESTAMP_GRANULARITY_UNSPECIFIED",
	1: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"TIMESTAMP_GRANULARITY_UNSPECIFIED": 0,
	"MILLIS": 1,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

// Defines a view over a table's fields.
type Table_View int32

const (
	// Uses the default view for each method as documented in its request.
	Table_VIEW_UNSPECIFIED Table_View = 0
	// Only populates `name`.
	Table_NAME_ONLY Table_View = 1
	// Only populates `name` and fields related to the table's schema.
	Table_SCHEMA_VIEW Table_View = 2
	// Populates all fields.
	Table_FULL Table_View = 4
)

var Table_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "NAME_ONLY",
	2: "SCHEMA_VIEW",
	4: "FULL",
}
var Table_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"NAME_ONLY":        1,
	"SCHEMA_VIEW":      2,
	"FULL":             4,
}

func (x Table_View) String() string {
	return proto.EnumName(Table_View_name, int32(x))
}
func (Table_View) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// (`OutputOnly`)
	// The unique name of the table. Values are of the form
	// `projects/<project>/instances/<instance>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	// Views: `NAME_ONLY`, `SCHEMA_VIEW`, `FULL`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (`CreationOnly`)
	// The column families configured for this table, mapped by column family ID.
	// Views: `SCHEMA_VIEW`, `FULL`
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// (`CreationOnly`)
	// The granularity (e.g. `MILLIS`, `MICROS`) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// If unspecified at creation time, the value will be set to `MILLIS`.
	// Views: `SCHEMA_VIEW`, `FULL`
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,enum=google.bigtable.admin.v2.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

func (m *Table) GetGranularity() Table_TimestampGranularity {
	if m != nil {
		return m.Granularity
	}
	return Table_TIMESTAMP_GRANULARITY_UNSPECIFIED
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,1,opt,name=gc_rule,json=gcRule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()                    { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string            { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()               {}
func (*ColumnFamily) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Garbage collection rules.
	//
	// Types that are valid to be assigned to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (m *GcRule) Reset()                    { *m = GcRule{} }
func (m *GcRule) String() string            { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()               {}
func (*GcRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isGcRule_Rule interface {
	isGcRule_Rule()
}

type GcRule_MaxNumVersions struct {
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,oneof"`
}
type GcRule_MaxAge struct {
	MaxAge *google_protobuf4.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,oneof"`
}
type GcRule_Intersection_ struct {
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,oneof"`
}
type GcRule_Union_ struct {
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}
func (*GcRule_MaxAge) isGcRule_Rule()         {}
func (*GcRule_Intersection_) isGcRule_Rule()  {}
func (*GcRule_Union_) isGcRule_Rule()         {}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *GcRule) GetMaxNumVersions() int32 {
	if x, ok := m.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (m *GcRule) GetMaxAge() *google_protobuf4.Duration {
	if x, ok := m.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := m.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if x, ok := m.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GcRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GcRule_OneofMarshaler, _GcRule_OneofUnmarshaler, _GcRule_OneofSizer, []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
}

func _GcRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MaxAge); err != nil {
			return err
		}
	case *GcRule_Intersection_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Intersection); err != nil {
			return err
		}
	case *GcRule_Union_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Union); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GcRule.Rule has unexpected type %T", x)
	}
	return nil
}

func _GcRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GcRule)
	switch tag {
	case 1: // rule.max_num_versions
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &GcRule_MaxNumVersions{int32(x)}
		return true, err
	case 2: // rule.max_age
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf4.Duration)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_MaxAge{msg}
		return true, err
	case 3: // rule.intersection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Intersection)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Intersection_{msg}
		return true, err
	case 4: // rule.union
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Union)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Union_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GcRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		s := proto.Size(x.MaxAge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Intersection_:
		s := proto.Size(x.Intersection)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Union_:
		s := proto.Size(x.Union)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()                    { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()               {}
func (*GcRule_Intersection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()                    { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()               {}
func (*GcRule_Union) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 1} }

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Table)(nil), "google.bigtable.admin.v2.Table")
	proto.RegisterType((*ColumnFamily)(nil), "google.bigtable.admin.v2.ColumnFamily")
	proto.RegisterType((*GcRule)(nil), "google.bigtable.admin.v2.GcRule")
	proto.RegisterType((*GcRule_Intersection)(nil), "google.bigtable.admin.v2.GcRule.Intersection")
	proto.RegisterType((*GcRule_Union)(nil), "google.bigtable.admin.v2.GcRule.Union")
	proto.RegisterEnum("google.bigtable.admin.v2.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Table_View", Table_View_name, Table_View_value)
}

func init() { proto.RegisterFile("google/bigtable/admin/v2/table.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6e, 0xda, 0x3c,
	0x18, 0xc7, 0x09, 0x09, 0xf4, 0xed, 0x43, 0xdf, 0x36, 0xf2, 0x7a, 0xc0, 0x50, 0xd5, 0x31, 0xb4,
	0x4d, 0x68, 0xd2, 0x12, 0x29, 0xad, 0xa6, 0x7d, 0x4f, 0x81, 0x06, 0x88, 0x04, 0x0c, 0x85, 0x8f,
	0xa9, 0xd3, 0xa4, 0xc8, 0x50, 0x37, 0xb2, 0x16, 0x3b, 0x28, 0x24, 0xac, 0xdc, 0xc3, 0x0e, 0x76,
	0x1d, 0xbb, 0x99, 0xdd, 0xd2, 0x14, 0x27, 0x68, 0xb4, 0x2b, 0xea, 0xb4, 0x23, 0x6c, 0x3f, 0xff,
	0xff, 0xef, 0xf9, 0xb0, 0x09, 0x3c, 0xf2, 0x82, 0xc0, 0xf3, 0x89, 0x3e, 0xa5, 0x5e, 0x84, 0xa7,
	0x3e, 0xd1, 0xf1, 0x05, 0xa3, 0x5c, 0x5f, 0x1a, 0xba, 0xd8, 0x6a, 0xf3, 0x30, 0x88, 0x02, 0x54,
	0x4e, 0x55, 0xda, 0x5a, 0xa5, 0x09, 0x95, 0xb6, 0x34, 0x2a, 0x47, 0x99, 0x1f, 0xcf, 0xa9, 0x8e,
	0x39, 0x0f, 0x22, 0x1c, 0xd1, 0x80, 0x2f, 0x52, 0x5f, 0xe5, 0x38, 0x8b, 0x8a, 0xdd, 0x34, 0xbe,
	0xd4, 0x2f, 0xe2, 0x50, 0x08, 0xb2, 0xf8, 0x83, 0x9b, 0xf1, 0x88, 0x32, 0xb2, 0x88, 0x30, 0x9b,
	0xa7, 0x82, 0xda, 0x4f, 0x19, 0x0a, 0xa3, 0x24, 0x23, 0x42, 0xa0, 0x70, 0xcc, 0x48, 0x59, 0xaa,
	0x4a, 0xf5, 0x5d, 0x47, 0xac, 0xd1, 0x67, 0x38, 0x98, 0x05, 0x7e, 0xcc, 0xb8, 0x7b, 0x89, 0x19,
	0xf5, 0x29, 0x59, 0x94, 0xe5, 0xaa, 0x5c, 0x2f, 0x19, 0x27, 0xda, 0xb6, 0x82, 0x35, 0x41, 0xd3,
	0x9a, 0xc2, 0xd6, 0xca, 0x5c, 0x16, 0x8f, 0xc2, 0x95, 0xb3, 0x3f, 0xbb, 0x76, 0x88, 0x26, 0x50,
	0xf2, 0x42, 0xcc, 0x63, 0x1f, 0x87, 0x34, 0x5a, 0x95, 0x95, 0xaa, 0x54, 0xdf, 0x37, 0x4e, 0xef,
	0x22, 0x8f, 0xd6, 0x1d, 0xb4, 0x7f, 0x7b, 0x9d, 0x4d, 0x50, 0x85, 0xc2, 0xbd, 0x5b, 0xd2, 0x23,
	0x15, 0xe4, 0x2f, 0x64, 0x95, 0xf5, 0x97, 0x2c, 0xd1, 0x1b, 0x28, 0x2c, 0xb1, 0x1f, 0x93, 0x72,
	0xbe, 0x2a, 0xd5, 0x4b, 0xc6, 0x93, 0xed, 0xa9, 0x37, 0x78, 0x2b, 0x27, 0x35, 0xbd, 0xca, 0xbf,
	0x90, 0x6a, 0x36, 0x1c, 0xde, 0x56, 0x0f, 0x7a, 0x0c, 0x0f, 0x47, 0x76, 0xcf, 0x1a, 0x8e, 0xcc,
	0xde, 0xc0, 0x6d, 0x3b, 0x66, 0x7f, 0xdc, 0x35, 0x1d, 0x7b, 0x74, 0xee, 0x8e, 0xfb, 0xc3, 0x81,
	0xd5, 0xb4, 0x5b, 0xb6, 0x75, 0xa6, 0xe6, 0x10, 0x40, 0xb1, 0x67, 0x77, 0xbb, 0xf6, 0x50, 0x95,
	0x6a, 0x2d, 0x50, 0x26, 0x94, 0x7c, 0x45, 0x87, 0xa0, 0x4e, 0x6c, 0xeb, 0xe3, 0x0d, 0xe5, 0xff,
	0xb0, 0xdb, 0x37, 0x7b, 0x96, 0xfb, 0xa1, 0xdf, 0x3d, 0x57, 0x25, 0x74, 0x00, 0xa5, 0x61, 0xb3,
	0x63, 0xf5, 0x4c, 0x37, 0xd1, 0xaa, 0x79, 0xf4, 0x1f, 0x28, 0xad, 0x71, 0xb7, 0xab, 0x2a, 0x35,
	0x1b, 0xf6, 0x36, 0xab, 0x45, 0x2f, 0x61, 0xc7, 0x9b, 0xb9, 0x61, 0xec, 0xa7, 0x57, 0x5b, 0x32,
	0xaa, 0xdb, 0xdb, 0x6c, 0xcf, 0x9c, 0xd8, 0x27, 0x4e, 0xd1, 0x13, 0xbf, 0xb5, 0xef, 0x32, 0x14,
	0xd3, 0x23, 0xf4, 0x14, 0x54, 0x86, 0xaf, 0x5c, 0x1e, 0x33, 0x77, 0x49, 0xc2, 0x45, 0xf2, 0x04,
	0x05, 0xae, 0xd0, 0xc9, 0x39, 0xfb, 0x0c, 0x5f, 0xf5, 0x63, 0x36, 0xc9, 0xce, 0xd1, 0x29, 0xec,
	0x24, 0x5a, 0xec, 0xad, 0x07, 0x7b, 0x7f, 0x9d, 0x71, 0xfd, 0x0c, 0xb5, 0xb3, 0xec, 0x99, 0x76,
	0x72, 0x4e, 0x91, 0xe1, 0x2b, 0xd3, 0x23, 0x68, 0x08, 0x7b, 0x94, 0x47, 0x24, 0x5c, 0x90, 0x59,
	0x12, 0x29, 0xcb, 0xc2, 0xfa, 0xec, 0xae, 0x62, 0x35, 0x7b, 0xc3, 0xd4, 0xc9, 0x39, 0xd7, 0x20,
	0xe8, 0x1d, 0x14, 0x62, 0x9e, 0xd0, 0x94, 0xbb, 0x6e, 0x38, 0xa3, 0x8d, 0x79, 0x8a, 0x49, 0x6d,
	0x95, 0x16, 0xec, 0x6d, 0xf2, 0xd1, 0x73, 0x28, 0x24, 0x93, 0x4c, 0x7a, 0x97, 0xff, 0x6a, 0x94,
	0xa9, 0xbc, 0xf2, 0x1e, 0x0a, 0x82, 0xfc, 0xaf, 0x80, 0x46, 0x11, 0x94, 0x64, 0xd1, 0xf8, 0x26,
	0xc1, 0xd1, 0x2c, 0x60, 0x5b, 0x6d, 0x0d, 0x10, 0xff, 0x92, 0x41, 0x32, 0xe8, 0x81, 0xf4, 0xe9,
	0x6d, 0xa6, 0xf3, 0x02, 0x1f, 0x73, 0x4f, 0x0b, 0x42, 0x4f, 0xf7, 0x08, 0x17, 0xd7, 0xa0, 0xa7,
	0x21, 0x3c, 0xa7, 0x8b, 0x3f, 0x3f, 0x4e, 0xaf, 0xc5, 0xe2, 0x47, 0xfe, 0xb8, 0x9d, 0xfa, 0x9b,
	0x7e, 0x10, 0x5f, 0x68, 0x8d, 0x75, 0x36, 0x53, 0x64, 0x9b, 0x18, 0xd3, 0xa2, 0x40, 0x9d, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xf3, 0x89, 0x9d, 0xe6, 0x04, 0x00, 0x00,
}