// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/metadata.proto

/*
Package metadata is a generated protocol buffer package.

It is generated from these files:
	metadata/metadata.proto

It has these top-level messages:
	HashingCosts
	WrappedKeyData
	ProtectorData
	EncryptionOptions
	WrappedPolicyKey
	PolicyData
	Config
*/
package metadata

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

// Specifies the method in which an outside secret is obtained for a Protector
type SourceType int32

const (
	SourceType_default           SourceType = 0
	SourceType_pam_passphrase    SourceType = 1
	SourceType_custom_passphrase SourceType = 2
	SourceType_raw_key           SourceType = 3
)

var SourceType_name = map[int32]string{
	0: "default",
	1: "pam_passphrase",
	2: "custom_passphrase",
	3: "raw_key",
}
var SourceType_value = map[string]int32{
	"default":           0,
	"pam_passphrase":    1,
	"custom_passphrase": 2,
	"raw_key":           3,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}
func (SourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Type of encryption; should match declarations of unix.FS_ENCRYPTION_MODE
type EncryptionOptions_Mode int32

const (
	EncryptionOptions_default     EncryptionOptions_Mode = 0
	EncryptionOptions_AES_256_XTS EncryptionOptions_Mode = 1
	EncryptionOptions_AES_256_GCM EncryptionOptions_Mode = 2
	EncryptionOptions_AES_256_CBC EncryptionOptions_Mode = 3
	EncryptionOptions_AES_256_CTS EncryptionOptions_Mode = 4
	EncryptionOptions_AES_128_CBC EncryptionOptions_Mode = 5
	EncryptionOptions_AES_128_CTS EncryptionOptions_Mode = 6
)

var EncryptionOptions_Mode_name = map[int32]string{
	0: "default",
	1: "AES_256_XTS",
	2: "AES_256_GCM",
	3: "AES_256_CBC",
	4: "AES_256_CTS",
	5: "AES_128_CBC",
	6: "AES_128_CTS",
}
var EncryptionOptions_Mode_value = map[string]int32{
	"default":     0,
	"AES_256_XTS": 1,
	"AES_256_GCM": 2,
	"AES_256_CBC": 3,
	"AES_256_CTS": 4,
	"AES_128_CBC": 5,
	"AES_128_CTS": 6,
}

func (x EncryptionOptions_Mode) String() string {
	return proto.EnumName(EncryptionOptions_Mode_name, int32(x))
}
func (EncryptionOptions_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// Cost parameters to be used in our hashing functions.
type HashingCosts struct {
	Time        int64 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Memory      int64 `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	Parallelism int64 `protobuf:"varint,4,opt,name=parallelism" json:"parallelism,omitempty"`
}

func (m *HashingCosts) Reset()                    { *m = HashingCosts{} }
func (m *HashingCosts) String() string            { return proto.CompactTextString(m) }
func (*HashingCosts) ProtoMessage()               {}
func (*HashingCosts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HashingCosts) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HashingCosts) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *HashingCosts) GetParallelism() int64 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

// This structure is used for our authenticated wrapping/unwrapping of keys.
type WrappedKeyData struct {
	IV           []byte `protobuf:"bytes,1,opt,name=IV,proto3" json:"IV,omitempty"`
	EncryptedKey []byte `protobuf:"bytes,2,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Hmac         []byte `protobuf:"bytes,3,opt,name=hmac,proto3" json:"hmac,omitempty"`
}

func (m *WrappedKeyData) Reset()                    { *m = WrappedKeyData{} }
func (m *WrappedKeyData) String() string            { return proto.CompactTextString(m) }
func (*WrappedKeyData) ProtoMessage()               {}
func (*WrappedKeyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WrappedKeyData) GetIV() []byte {
	if m != nil {
		return m.IV
	}
	return nil
}

func (m *WrappedKeyData) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

func (m *WrappedKeyData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

// The associated data for each protector
type ProtectorData struct {
	ProtectorDescriptor string     `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor" json:"protector_descriptor,omitempty"`
	Source              SourceType `protobuf:"varint,2,opt,name=source,enum=metadata.SourceType" json:"source,omitempty"`
	// These are only used by some of the protector types
	Name       string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Costs      *HashingCosts   `protobuf:"bytes,4,opt,name=costs" json:"costs,omitempty"`
	Salt       []byte          `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	Uid        int64           `protobuf:"varint,6,opt,name=uid" json:"uid,omitempty"`
	WrappedKey *WrappedKeyData `protobuf:"bytes,7,opt,name=wrapped_key,json=wrappedKey" json:"wrapped_key,omitempty"`
}

func (m *ProtectorData) Reset()                    { *m = ProtectorData{} }
func (m *ProtectorData) String() string            { return proto.CompactTextString(m) }
func (*ProtectorData) ProtoMessage()               {}
func (*ProtectorData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProtectorData) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *ProtectorData) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_default
}

func (m *ProtectorData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtectorData) GetCosts() *HashingCosts {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *ProtectorData) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ProtectorData) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ProtectorData) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// Encryption policy specifics, corresponds to the fscrypt_policy struct
type EncryptionOptions struct {
	Padding   int64                  `protobuf:"varint,1,opt,name=padding" json:"padding,omitempty"`
	Contents  EncryptionOptions_Mode `protobuf:"varint,2,opt,name=contents,enum=metadata.EncryptionOptions_Mode" json:"contents,omitempty"`
	Filenames EncryptionOptions_Mode `protobuf:"varint,3,opt,name=filenames,enum=metadata.EncryptionOptions_Mode" json:"filenames,omitempty"`
}

func (m *EncryptionOptions) Reset()                    { *m = EncryptionOptions{} }
func (m *EncryptionOptions) String() string            { return proto.CompactTextString(m) }
func (*EncryptionOptions) ProtoMessage()               {}
func (*EncryptionOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EncryptionOptions) GetPadding() int64 {
	if m != nil {
		return m.Padding
	}
	return 0
}

func (m *EncryptionOptions) GetContents() EncryptionOptions_Mode {
	if m != nil {
		return m.Contents
	}
	return EncryptionOptions_default
}

func (m *EncryptionOptions) GetFilenames() EncryptionOptions_Mode {
	if m != nil {
		return m.Filenames
	}
	return EncryptionOptions_default
}

type WrappedPolicyKey struct {
	ProtectorDescriptor string          `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor" json:"protector_descriptor,omitempty"`
	WrappedKey          *WrappedKeyData `protobuf:"bytes,2,opt,name=wrapped_key,json=wrappedKey" json:"wrapped_key,omitempty"`
}

func (m *WrappedPolicyKey) Reset()                    { *m = WrappedPolicyKey{} }
func (m *WrappedPolicyKey) String() string            { return proto.CompactTextString(m) }
func (*WrappedPolicyKey) ProtoMessage()               {}
func (*WrappedPolicyKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WrappedPolicyKey) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *WrappedPolicyKey) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// The associated data for each policy
type PolicyData struct {
	KeyDescriptor     string              `protobuf:"bytes,1,opt,name=key_descriptor,json=keyDescriptor" json:"key_descriptor,omitempty"`
	Options           *EncryptionOptions  `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	WrappedPolicyKeys []*WrappedPolicyKey `protobuf:"bytes,3,rep,name=wrapped_policy_keys,json=wrappedPolicyKeys" json:"wrapped_policy_keys,omitempty"`
}

func (m *PolicyData) Reset()                    { *m = PolicyData{} }
func (m *PolicyData) String() string            { return proto.CompactTextString(m) }
func (*PolicyData) ProtoMessage()               {}
func (*PolicyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PolicyData) GetKeyDescriptor() string {
	if m != nil {
		return m.KeyDescriptor
	}
	return ""
}

func (m *PolicyData) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *PolicyData) GetWrappedPolicyKeys() []*WrappedPolicyKey {
	if m != nil {
		return m.WrappedPolicyKeys
	}
	return nil
}

// Data stored in the config file
type Config struct {
	Source        SourceType         `protobuf:"varint,1,opt,name=source,enum=metadata.SourceType" json:"source,omitempty"`
	HashCosts     *HashingCosts      `protobuf:"bytes,2,opt,name=hash_costs,json=hashCosts" json:"hash_costs,omitempty"`
	Compatibility string             `protobuf:"bytes,3,opt,name=compatibility" json:"compatibility,omitempty"`
	Options       *EncryptionOptions `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Config) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_default
}

func (m *Config) GetHashCosts() *HashingCosts {
	if m != nil {
		return m.HashCosts
	}
	return nil
}

func (m *Config) GetCompatibility() string {
	if m != nil {
		return m.Compatibility
	}
	return ""
}

func (m *Config) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*HashingCosts)(nil), "metadata.HashingCosts")
	proto.RegisterType((*WrappedKeyData)(nil), "metadata.WrappedKeyData")
	proto.RegisterType((*ProtectorData)(nil), "metadata.ProtectorData")
	proto.RegisterType((*EncryptionOptions)(nil), "metadata.EncryptionOptions")
	proto.RegisterType((*WrappedPolicyKey)(nil), "metadata.WrappedPolicyKey")
	proto.RegisterType((*PolicyData)(nil), "metadata.PolicyData")
	proto.RegisterType((*Config)(nil), "metadata.Config")
	proto.RegisterEnum("metadata.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("metadata.EncryptionOptions_Mode", EncryptionOptions_Mode_name, EncryptionOptions_Mode_value)
}

func init() { proto.RegisterFile("metadata/metadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xd2, 0xb5, 0xeb, 0xe9, 0x1f, 0x32, 0x6f, 0x8c, 0x08, 0x6e, 0xaa, 0x00, 0xd2,
	0x84, 0xa6, 0xa1, 0x15, 0x0d, 0x81, 0x84, 0x90, 0xa0, 0x9b, 0x60, 0x4c, 0x13, 0xc3, 0xad, 0x06,
	0x48, 0x48, 0x95, 0x97, 0x78, 0xab, 0xb5, 0x24, 0xb6, 0x62, 0x57, 0x55, 0xee, 0x78, 0x07, 0xde,
	0x81, 0x47, 0xe0, 0x21, 0x78, 0x2a, 0x64, 0xa7, 0x49, 0xd3, 0x4d, 0x9a, 0x06, 0x37, 0xd1, 0xf1,
	0x67, 0xfb, 0x7c, 0xc7, 0x3f, 0xfb, 0x04, 0xee, 0xc7, 0x54, 0x91, 0x90, 0x28, 0xf2, 0xac, 0x08,
	0x76, 0x44, 0xca, 0x15, 0x47, 0xab, 0xc5, 0xd8, 0xff, 0x0e, 0xed, 0x0f, 0x44, 0x4e, 0x58, 0x72,
	0x31, 0xe0, 0x52, 0x49, 0x84, 0xa0, 0xa6, 0x58, 0x4c, 0x3d, 0xbb, 0x67, 0x6d, 0x39, 0xd8, 0xc4,
	0x68, 0x13, 0xea, 0x31, 0x8d, 0x79, 0x9a, 0x79, 0x8e, 0x51, 0xe7, 0x23, 0xd4, 0x83, 0x96, 0x20,
	0x29, 0x89, 0x22, 0x1a, 0x31, 0x19, 0x7b, 0x35, 0x33, 0x59, 0x95, 0xfc, 0x6f, 0xd0, 0xfd, 0x92,
	0x12, 0x21, 0x68, 0x78, 0x44, 0xb3, 0x7d, 0xa2, 0x08, 0xea, 0x82, 0x7d, 0x78, 0xea, 0x59, 0x3d,
	0x6b, 0xab, 0x8d, 0xed, 0xc3, 0x53, 0xf4, 0x08, 0x3a, 0x34, 0x09, 0xd2, 0x4c, 0x28, 0x1a, 0x8e,
	0x2f, 0x69, 0x66, 0x8c, 0xdb, 0xb8, 0x5d, 0x8a, 0x47, 0x34, 0xd3, 0x45, 0x4d, 0x62, 0x12, 0x18,
	0xfb, 0x36, 0x36, 0xb1, 0xff, 0xd3, 0x86, 0xce, 0x49, 0xca, 0x15, 0x0d, 0x14, 0x4f, 0x4d, 0xea,
	0x5d, 0xd8, 0x10, 0x85, 0x30, 0x0e, 0xa9, 0x0c, 0x52, 0x26, 0x14, 0x4f, 0x8d, 0x59, 0x13, 0xaf,
	0x97, 0x73, 0xfb, 0xe5, 0x14, 0xda, 0x86, 0xba, 0xe4, 0xd3, 0x34, 0xc8, 0xcf, 0xdb, 0xed, 0x6f,
	0xec, 0x94, 0xa0, 0x86, 0x46, 0x1f, 0x65, 0x82, 0xe2, 0xf9, 0x1a, 0x5d, 0x46, 0x42, 0x62, 0x6a,
	0xca, 0x68, 0x62, 0x13, 0xa3, 0x6d, 0x58, 0x09, 0x34, 0x38, 0x73, 0xfa, 0x56, 0x7f, 0x73, 0x91,
	0xa0, 0x8a, 0x15, 0xe7, 0x8b, 0x74, 0x06, 0x49, 0x22, 0xe5, 0xad, 0xe4, 0x07, 0xd1, 0x31, 0x72,
	0xc1, 0x99, 0xb2, 0xd0, 0xab, 0x1b, 0x7a, 0x3a, 0x44, 0xaf, 0xa0, 0x35, 0xcb, 0xa9, 0x19, 0x22,
	0x0d, 0x93, 0xd9, 0x5b, 0x64, 0x5e, 0x46, 0x8a, 0x61, 0x56, 0x8e, 0xfd, 0x5f, 0x36, 0xac, 0x1d,
	0xe4, 0xe8, 0x18, 0x4f, 0x3e, 0x99, 0xaf, 0x44, 0x1e, 0x34, 0x04, 0x09, 0x43, 0x96, 0x5c, 0x18,
	0x18, 0x0e, 0x2e, 0x86, 0xe8, 0x35, 0xac, 0x06, 0x3c, 0x51, 0x34, 0x51, 0x72, 0x8e, 0xa0, 0xb7,
	0xf0, 0xb9, 0x96, 0x68, 0xe7, 0x98, 0x87, 0x14, 0x97, 0x3b, 0xd0, 0x1b, 0x68, 0x9e, 0xb3, 0x88,
	0x6a, 0x10, 0xd2, 0x50, 0xb9, 0xcd, 0xf6, 0xc5, 0x16, 0x3f, 0x83, 0x9a, 0x96, 0x50, 0x0b, 0x1a,
	0x21, 0x3d, 0x27, 0xd3, 0x48, 0xb9, 0x77, 0xd0, 0x5d, 0x68, 0xbd, 0x3d, 0x18, 0x8e, 0xfb, 0x7b,
	0x2f, 0xc6, 0x5f, 0x47, 0x43, 0xd7, 0xaa, 0x0a, 0xef, 0x07, 0xc7, 0xae, 0x5d, 0x15, 0x06, 0xef,
	0x06, 0xae, 0xb3, 0x24, 0x8c, 0x86, 0x6e, 0xad, 0x10, 0x76, 0xfb, 0x2f, 0xcd, 0x8a, 0x95, 0x25,
	0x61, 0x34, 0x74, 0xeb, 0xfe, 0x0f, 0x0b, 0xdc, 0x39, 0xc7, 0x13, 0x1e, 0xb1, 0x20, 0xd3, 0xef,
	0xec, 0x3f, 0x5e, 0xd0, 0x95, 0xbb, 0xb2, 0xff, 0xe1, 0xae, 0x7e, 0x5b, 0x00, 0xb9, 0xb7, 0x79,
	0xbe, 0x4f, 0xa0, 0x7b, 0x49, 0xb3, 0xeb, 0xb6, 0x9d, 0x4b, 0x9a, 0x55, 0x0c, 0xf7, 0xa0, 0xc1,
	0x73, 0x9c, 0x73, 0xb3, 0x87, 0x37, 0x10, 0xc7, 0xc5, 0x5a, 0xf4, 0x11, 0xd6, 0x8b, 0x3a, 0x85,
	0xf1, 0xd4, 0xe5, 0xea, 0x4b, 0x73, 0xb6, 0x5a, 0xfd, 0x07, 0xd7, 0xea, 0x2d, 0x99, 0xe0, 0xb5,
	0xd9, 0x15, 0x45, 0xfa, 0x7f, 0x2c, 0xa8, 0x0f, 0x78, 0x72, 0xce, 0x2e, 0x2a, 0x0d, 0x64, 0xdd,
	0xa2, 0x81, 0xf6, 0x00, 0x26, 0x44, 0x4e, 0xc6, 0x79, 0xc7, 0xd8, 0x37, 0x76, 0x4c, 0x53, 0xaf,
	0xcc, 0xff, 0x49, 0x8f, 0xa1, 0x13, 0xf0, 0x58, 0x10, 0xc5, 0xce, 0x58, 0xc4, 0x54, 0x36, 0x6f,
	0xc0, 0x65, 0xb1, 0x0a, 0xa6, 0x76, 0x7b, 0x30, 0x4f, 0x3f, 0x03, 0x2c, 0x2a, 0x5d, 0x7e, 0x89,
	0x08, 0xba, 0x82, 0xc4, 0x63, 0x41, 0xa4, 0x14, 0x93, 0x94, 0x48, 0xea, 0x5a, 0xe8, 0x1e, 0xac,
	0x05, 0x53, 0xa9, 0xf8, 0x92, 0x6c, 0xeb, 0x7d, 0x29, 0x99, 0x69, 0xa6, 0xae, 0x73, 0x56, 0x37,
	0x3f, 0xd9, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc6, 0x2e, 0x75, 0x7f, 0x05, 0x00,
	0x00,
}
