// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configtx.proto
	common/configuration.proto
	common/ledger.proto
	common/msp_principal.proto
	common/policies.proto

It has these top-level messages:
	LastConfig
	Metadata
	MetadataSignature
	Header
	ChannelHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigEnvelope
	ConfigGroupSchema
	ConfigValueSchema
	ConfigPolicySchema
	Config
	ConfigUpdateEnvelope
	ConfigUpdate
	ConfigGroup
	ConfigValue
	ConfigPolicy
	ConfigSignature
	HashingAlgorithm
	BlockDataHashingStructure
	OrdererAddresses
	BlockchainInfo
	MSPPrincipal
	OrganizationUnit
	MSPRole
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE              HeaderType = 0
	HeaderType_CONFIG               HeaderType = 1
	HeaderType_CONFIG_UPDATE        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION HeaderType = 3
	HeaderType_ORDERER_TRANSACTION  HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO    HeaderType = 5
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",
	2: "CONFIG_UPDATE",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":              0,
	"CONFIG":               1,
	"CONFIG_UPDATE":        2,
	"ENDORSER_TRANSACTION": 3,
	"ORDERER_TRANSACTION":  4,
	"DELIVER_SEEK_INFO":    5,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// This enum enlists indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIG         BlockMetadataIndex = 1
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
	BlockMetadataIndex_ORDERER             BlockMetadataIndex = 3
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIG",
	2: "TRANSACTIONS_FILTER",
	3: "ORDERER",
}
var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIG":         1,
	"TRANSACTIONS_FILTER": 2,
	"ORDERER":             3,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}
func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// LastConfig is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfig struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *LastConfig) Reset()                    { *m = LastConfig{} }
func (m *LastConfig) String() string            { return proto.CompactTextString(m) }
func (*LastConfig) ProtoMessage()               {}
func (*LastConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value      []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MetadataSignature) Reset()                    { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string            { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()               {}
func (*MetadataSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Header struct {
	ChannelHeader   *ChannelHeader   `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader" json:"channel_header,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader" json:"signature_header,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Header) GetChannelHeader() *ChannelHeader {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ChannelHeader) Reset()                    { *m = ChannelHeader{} }
func (m *ChannelHeader) String() string            { return proto.CompactTextString(m) }
func (*ChannelHeader) ProtoMessage()               {}
func (*ChannelHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChannelHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BlockHeader is the element of the block which forms the block chain
// The block header is hashed using the configured chain hashing algorithm
// over the ASN.1 encoding of the BlockHeader
type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*LastConfig)(nil), "common.LastConfig")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChannelHeader)(nil), "common.ChannelHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 886 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x6e, 0xe3, 0xc4,
	0x1b, 0xad, 0xe3, 0xfc, 0x69, 0xbe, 0x34, 0xed, 0x74, 0xb2, 0xfd, 0xd5, 0xbf, 0xc2, 0x6a, 0x2b,
	0x23, 0x50, 0x69, 0x45, 0x22, 0xca, 0x0d, 0x48, 0xdc, 0x38, 0xf1, 0xa4, 0x6b, 0x6d, 0xd6, 0x5e,
	0xc6, 0xce, 0x22, 0x58, 0x24, 0x6b, 0x92, 0x4c, 0x13, 0x8b, 0xc4, 0x8e, 0x6c, 0xa7, 0x6a, 0x6f,
	0xb9, 0xe4, 0x02, 0x21, 0xc1, 0x2d, 0x2f, 0xc0, 0x93, 0xf0, 0x16, 0xbc, 0x04, 0x12, 0xb7, 0x68,
	0x3c, 0xb6, 0x37, 0xe9, 0x22, 0x71, 0x15, 0x9f, 0xf3, 0x9d, 0x99, 0xef, 0xcc, 0xf9, 0xc6, 0x31,
	0x74, 0xa6, 0xd1, 0x6a, 0x15, 0x85, 0x3d, 0xf9, 0xd3, 0x5d, 0xc7, 0x51, 0x1a, 0xe1, 0xba, 0x44,
	0x67, 0xcf, 0xe6, 0x51, 0x34, 0x5f, 0xf2, 0x5e, 0xc6, 0x4e, 0x36, 0xb7, 0xbd, 0x34, 0x58, 0xf1,
	0x24, 0x65, 0xab, 0xb5, 0x14, 0xea, 0x3a, 0xc0, 0x88, 0x25, 0xe9, 0x20, 0x0a, 0x6f, 0x83, 0x39,
	0x7e, 0x02, 0xb5, 0x20, 0x9c, 0xf1, 0x7b, 0x4d, 0x39, 0x57, 0x2e, 0xaa, 0x54, 0x02, 0xfd, 0x0d,
	0xec, 0xbf, 0xe4, 0x29, 0x9b, 0xb1, 0x94, 0x09, 0xc5, 0x1d, 0x5b, 0x6e, 0x78, 0xa6, 0x38, 0xa0,
	0x12, 0xe0, 0x2f, 0x00, 0x92, 0x60, 0x1e, 0xb2, 0x74, 0x13, 0xf3, 0x44, 0xab, 0x9c, 0xab, 0x17,
	0xad, 0xeb, 0xff, 0x77, 0x73, 0x47, 0xc5, 0x5a, 0xb7, 0x50, 0xd0, 0x2d, 0xb1, 0xfe, 0x1d, 0x1c,
	0xbf, 0x23, 0xc0, 0x1f, 0x03, 0x2a, 0x25, 0xfe, 0x82, 0xb3, 0x19, 0x8f, 0xf3, 0x86, 0x47, 0x25,
	0xff, 0x3c, 0xa3, 0xf1, 0xfb, 0xd0, 0x2c, 0x29, 0xad, 0x92, 0x69, 0xde, 0x12, 0xfa, 0x8f, 0x0a,
	0xd4, 0x73, 0xe1, 0x97, 0x70, 0x38, 0x5d, 0xb0, 0x30, 0xe4, 0xcb, 0xed, 0x1d, 0x5b, 0xd7, 0x27,
	0x85, 0xcf, 0x81, 0xac, 0x4a, 0x39, 0x6d, 0x4f, 0xb7, 0x21, 0xee, 0xff, 0x8b, 0xa3, 0x4a, 0xb6,
	0xfe, 0xb4, 0x58, 0xef, 0xee, 0x3a, 0x7b, 0xc7, 0xaa, 0xfe, 0xa7, 0x02, 0xed, 0x9d, 0x26, 0x18,
	0x43, 0x35, 0x7d, 0x58, 0xcb, 0x30, 0x6b, 0x34, 0x7b, 0xc6, 0x1a, 0x34, 0xee, 0x78, 0x9c, 0x04,
	0x51, 0x98, 0x35, 0xa8, 0xd1, 0x02, 0xe2, 0xcf, 0xa1, 0x59, 0x8e, 0x4f, 0x53, 0xb3, 0xe6, 0x67,
	0x5d, 0x39, 0xe0, 0x6e, 0x31, 0xe0, 0xae, 0x57, 0x28, 0xe8, 0x5b, 0x31, 0x7e, 0x0a, 0x50, 0x9c,
	0x3d, 0x98, 0x69, 0xd5, 0x73, 0xe5, 0xa2, 0x49, 0x9b, 0x39, 0x63, 0xcd, 0x70, 0x07, 0x6a, 0xe9,
	0xbd, 0xa8, 0xd4, 0xb2, 0x4a, 0x35, 0xbd, 0xb7, 0x66, 0x62, 0xd2, 0x7c, 0x1d, 0x4d, 0x17, 0x5a,
	0x5d, 0xde, 0x85, 0x0c, 0x88, 0xb8, 0xf9, 0x7d, 0xca, 0xc3, 0xcc, 0x5f, 0x43, 0xc6, 0x5d, 0x12,
	0xba, 0x01, 0x47, 0x8f, 0x52, 0x10, 0xc7, 0x99, 0xc6, 0x9c, 0xa5, 0x51, 0x31, 0xc1, 0x02, 0x8a,
	0x06, 0x61, 0x14, 0x4e, 0x8b, 0xa9, 0x49, 0xa0, 0x13, 0x68, 0xbc, 0x62, 0x0f, 0xcb, 0x88, 0xcd,
	0xf0, 0x47, 0x50, 0xdf, 0x99, 0xd4, 0x61, 0x91, 0x74, 0x1e, 0x70, 0x5e, 0x15, 0x29, 0x8a, 0xeb,
	0x93, 0xef, 0x93, 0x3d, 0xeb, 0x7d, 0xd8, 0x27, 0xe1, 0x1d, 0x5f, 0x46, 0x32, 0xd1, 0xb5, 0xdc,
	0xb2, 0xb0, 0x90, 0xc3, 0xff, 0xb8, 0x3c, 0x3f, 0x29, 0x50, 0xeb, 0x2f, 0xa3, 0xe9, 0xf7, 0xf8,
	0xea, 0x91, 0x93, 0x4e, 0xe1, 0x24, 0x2b, 0x3f, 0xb2, 0xf3, 0xe1, 0x96, 0x9d, 0xd6, 0xf5, 0xf1,
	0x8e, 0xd4, 0x64, 0x29, 0x93, 0x0e, 0xf1, 0xa7, 0xb0, 0xbf, 0xca, 0x2f, 0x7e, 0x3e, 0xcc, 0x93,
	0x1d, 0x69, 0xf1, 0x56, 0xd0, 0x52, 0xa6, 0xcf, 0xa1, 0xb5, 0xd5, 0x10, 0xff, 0x0f, 0xea, 0xe1,
	0x66, 0x35, 0xc9, 0x5d, 0x55, 0x69, 0x8e, 0xf0, 0x07, 0xd0, 0x5e, 0xc7, 0xfc, 0x2e, 0x88, 0x36,
	0x89, 0xbf, 0x60, 0xc9, 0x22, 0x3f, 0xd9, 0x41, 0x41, 0x3e, 0x67, 0xc9, 0x02, 0xbf, 0x07, 0x4d,
	0xb1, 0xa7, 0x14, 0xa8, 0x99, 0x60, 0x5f, 0x10, 0xa2, 0xa8, 0x3f, 0x83, 0x66, 0x69, 0xb7, 0x8c,
	0x57, 0x39, 0x57, 0xcb, 0x78, 0xaf, 0xa0, 0xbd, 0x63, 0x12, 0x9f, 0x6d, 0x9d, 0x46, 0x0a, 0x4b,
	0x7c, 0xf9, 0xbb, 0x02, 0x75, 0x37, 0x65, 0xe9, 0x26, 0xc1, 0x2d, 0x68, 0x8c, 0xed, 0x17, 0xb6,
	0xf3, 0xb5, 0x8d, 0xf6, 0xf0, 0x01, 0x34, 0xdc, 0xf1, 0x60, 0x40, 0x5c, 0x17, 0xfd, 0xa1, 0x60,
	0x04, 0xad, 0xbe, 0x61, 0xfa, 0x94, 0x7c, 0x35, 0x26, 0xae, 0x87, 0x7e, 0x56, 0xf1, 0x21, 0x34,
	0x87, 0x0e, 0xed, 0x5b, 0xa6, 0x49, 0x6c, 0xf4, 0x4b, 0x86, 0x6d, 0xc7, 0xf3, 0x87, 0xce, 0xd8,
	0x36, 0xd1, 0xaf, 0x2a, 0x7e, 0x0a, 0x5a, 0xae, 0xf6, 0x89, 0xed, 0x59, 0xde, 0x37, 0xbe, 0xe7,
	0x38, 0xfe, 0xc8, 0xa0, 0x37, 0x04, 0xfd, 0xa6, 0xe2, 0x33, 0x38, 0xb1, 0x6c, 0x8f, 0x50, 0xdb,
	0x18, 0xf9, 0x2e, 0xa1, 0xaf, 0x09, 0xf5, 0x09, 0xa5, 0x0e, 0x45, 0x7f, 0xa9, 0x58, 0x83, 0x8e,
	0xa0, 0xac, 0x01, 0xf1, 0xc7, 0xb6, 0xf1, 0xda, 0xb0, 0x46, 0x46, 0x7f, 0x44, 0xd0, 0xdf, 0xea,
	0xe5, 0x0f, 0x0a, 0x80, 0xcc, 0xd7, 0x13, 0x6f, 0x63, 0x0b, 0x1a, 0x2f, 0x89, 0xeb, 0x1a, 0x37,
	0x04, 0xed, 0x61, 0x80, 0xfa, 0xc0, 0xb1, 0x87, 0xd6, 0x0d, 0x52, 0xf0, 0x31, 0xb4, 0xe5, 0xb3,
	0x3f, 0x7e, 0x65, 0x1a, 0x1e, 0x41, 0x15, 0xac, 0xc1, 0x13, 0x62, 0x9b, 0x0e, 0x75, 0x09, 0xf5,
	0x3d, 0x6a, 0xd8, 0xae, 0x31, 0xf0, 0x2c, 0xc7, 0x46, 0x2a, 0x3e, 0x85, 0x8e, 0x43, 0x4d, 0x42,
	0x1f, 0x15, 0xaa, 0xf8, 0x04, 0x8e, 0x4d, 0x32, 0xb2, 0x84, 0x37, 0x97, 0x90, 0x17, 0xbe, 0x65,
	0x0f, 0x1d, 0x54, 0xbb, 0x7c, 0x03, 0x78, 0x27, 0x5e, 0x4b, 0xfc, 0x0f, 0xe3, 0x43, 0x00, 0xd7,
	0xba, 0xb1, 0x0d, 0x6f, 0x4c, 0x89, 0x8b, 0xf6, 0xf0, 0x11, 0xb4, 0x46, 0x86, 0xeb, 0xf9, 0xa5,
	0xa7, 0x53, 0xe8, 0x6c, 0x6d, 0xef, 0xfa, 0x43, 0x6b, 0xe4, 0x11, 0x8a, 0x2a, 0xe2, 0x14, 0x79,
	0x7f, 0xa4, 0xf6, 0x3f, 0xf9, 0xf6, 0x6a, 0x1e, 0xa4, 0x8b, 0xcd, 0x44, 0x5c, 0xb7, 0xde, 0xe2,
	0x61, 0xcd, 0xe3, 0x25, 0x9f, 0xcd, 0x79, 0xdc, 0xbb, 0x65, 0x93, 0x38, 0x98, 0xca, 0x8f, 0x45,
	0x92, 0x7f, 0x50, 0x26, 0xf5, 0x0c, 0x7e, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xa9,
	0x91, 0x13, 0x68, 0x06, 0x00, 0x00,
}