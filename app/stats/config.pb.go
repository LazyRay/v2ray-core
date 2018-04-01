package stats

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/app/stats/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x2f,
	0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x81, 0x29, 0x2b, 0x4a, 0xd5, 0x4b, 0x2c, 0x28, 0xd0, 0x03, 0x2b, 0x51, 0xe2,
	0xe0, 0x62, 0x73, 0x06, 0xab, 0x72, 0xb2, 0xe2, 0x92, 0x48, 0xce, 0xcf, 0xd5, 0xc3, 0xa6, 0x2a,
	0x80, 0x31, 0x8a, 0x15, 0xcc, 0x58, 0xc5, 0x24, 0x12, 0x66, 0x14, 0x94, 0x58, 0xa9, 0xe7, 0x0c,
	0x92, 0x77, 0x2c, 0x28, 0xd0, 0x0b, 0x06, 0x09, 0x27, 0xb1, 0x81, 0xad, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x24, 0xc6, 0x41, 0x8b, 0x00, 0x00, 0x00,
}