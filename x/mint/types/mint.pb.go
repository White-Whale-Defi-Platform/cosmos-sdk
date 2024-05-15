// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/mint/v1beta1/mint.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	// Deprecated: Use of inflation in Minter is deprecated as we no longer use this in minting.
	Inflation cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"` // Deprecated: Do not use.
	// current annual expected provisions
	// Deprecated: Use of annual_provisions in Minter is deprecated as we no longer use this in minting.
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annual_provisions"` // Deprecated: Do not use.
	// epoch_provisions represent rewards for the current epoch.
	EpochProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=epoch_provisions,json=epochProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"epoch_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params defines the parameters for the x/mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_max"`
	// minimum inflation rate
	InflationMin cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"goal_bonded"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// maximum supply for the token
	MaxSupply cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=max_supply,json=maxSupply,proto3,customtype=cosmossdk.io/math.Int" json:"max_supply"`
	// epoch_identifier mint epoch identifier e.g. (day, week).
	EpochIdentifier string `protobuf:"bytes,8,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// reduction_period_in_epochs the number of epochs it takes
	// to reduce the rewards.
	ReductionPeriodInEpochs int64 `protobuf:"varint,9,opt,name=reduction_period_in_epochs,json=reductionPeriodInEpochs,proto3" json:"reduction_period_in_epochs,omitempty"`
	// reduction_factor is the reduction multiplier to execute
	// at the end of each period set by reduction_period_in_epochs.
	ReductionFactor cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"reduction_factor"`
	// minting_rewards_distribution_start_epoch start epoch to distribute minting
	// rewards
	MintingRewardsDistributionStartEpoch int64 `protobuf:"varint,11,opt,name=minting_rewards_distribution_start_epoch,json=mintingRewardsDistributionStartEpoch,proto3" json:"minting_rewards_distribution_start_epoch,omitempty"`
	// genesis_epoch_provisions epoch provisions from the first epoch.
	GenesisEpochProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"genesis_epoch_provisions"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *Params) GetReductionPeriodInEpochs() int64 {
	if m != nil {
		return m.ReductionPeriodInEpochs
	}
	return 0
}

func (m *Params) GetMintingRewardsDistributionStartEpoch() int64 {
	if m != nil {
		return m.MintingRewardsDistributionStartEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "cosmos.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "cosmos.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/mint/v1beta1/mint.proto", fileDescriptor_2df116d183c1e223) }

var fileDescriptor_2df116d183c1e223 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x4e, 0x14, 0x31,
	0x1c, 0xc7, 0x77, 0x00, 0x57, 0xb6, 0x40, 0x80, 0x02, 0x5a, 0x30, 0x0c, 0x1b, 0x62, 0xcc, 0x06,
	0xc3, 0xae, 0x48, 0xd4, 0xc4, 0x8b, 0x09, 0x2e, 0x26, 0x6b, 0x24, 0x59, 0x86, 0x83, 0x51, 0x0f,
	0xb5, 0x3b, 0x53, 0x86, 0xca, 0x4e, 0x3b, 0xb6, 0x5d, 0xdc, 0x7d, 0x05, 0x4f, 0xc6, 0x37, 0xf0,
	0xe6, 0x91, 0x03, 0x0f, 0xc1, 0x91, 0x78, 0x32, 0x1c, 0x88, 0x81, 0x03, 0xaf, 0x61, 0xa6, 0x1d,
	0x67, 0x41, 0x22, 0x51, 0xd6, 0xcb, 0x64, 0xe6, 0xf7, 0xe7, 0xf3, 0xed, 0xb7, 0xd3, 0x5f, 0x81,
	0xeb, 0x0b, 0x15, 0x09, 0x55, 0x89, 0x18, 0xd7, 0x95, 0x9d, 0xa5, 0x06, 0xd5, 0x64, 0xc9, 0x7c,
	0x94, 0x63, 0x29, 0xb4, 0x80, 0x13, 0x36, 0x5f, 0x36, 0xa1, 0x34, 0x3f, 0x33, 0x19, 0x8a, 0x50,
	0x98, 0x7c, 0x25, 0x79, 0xb3, 0xa5, 0x33, 0xd3, 0xb6, 0x14, 0xdb, 0x44, 0xda, 0x67, 0x53, 0xe3,
	0x24, 0x62, 0x5c, 0x54, 0xcc, 0xd3, 0x86, 0xe6, 0xbf, 0xf4, 0x81, 0xfc, 0x1a, 0xe3, 0x9a, 0x4a,
	0xb8, 0x0e, 0x0a, 0x8c, 0x6f, 0x36, 0x89, 0x66, 0x82, 0x23, 0xa7, 0xe8, 0x94, 0x0a, 0x2b, 0xcb,
	0xfb, 0x47, 0x73, 0xb9, 0xc3, 0xa3, 0xb9, 0x5b, 0x16, 0xa3, 0x82, 0xed, 0x32, 0x13, 0x95, 0x88,
	0xe8, 0xad, 0xf2, 0x0b, 0x1a, 0x12, 0xbf, 0x53, 0xa5, 0xfe, 0xb7, 0xbd, 0x45, 0x90, 0xaa, 0x54,
	0xa9, 0x8f, 0x1c, 0xaf, 0x4b, 0x81, 0x6f, 0xc1, 0x38, 0xe1, 0xbc, 0x45, 0x9a, 0xc9, 0x6a, 0x76,
	0x98, 0x62, 0x82, 0x2b, 0xd4, 0x77, 0x75, 0xf4, 0x98, 0xa5, 0xd5, 0x33, 0x18, 0x6c, 0x80, 0x31,
	0x1a, 0x0b, 0x7f, 0xeb, 0xac, 0x40, 0xbf, 0x11, 0x78, 0xf4, 0x17, 0x02, 0x87, 0x7b, 0x8b, 0x63,
	0x36, 0xbd, 0xa8, 0x82, 0xed, 0xe2, 0xce, 0xbd, 0xf2, 0x83, 0xfb, 0xde, 0xa8, 0x01, 0x76, 0x35,
	0xe6, 0x3f, 0x0f, 0x82, 0x7c, 0x9d, 0x48, 0x12, 0x29, 0x38, 0x0b, 0x40, 0xf2, 0x0b, 0x70, 0x40,
	0xb9, 0x88, 0xec, 0x26, 0x79, 0x85, 0x24, 0x52, 0x4d, 0x02, 0xf0, 0x1d, 0x98, 0xca, 0xcc, 0x63,
	0x49, 0x34, 0xc5, 0xfe, 0x16, 0xe1, 0x21, 0x4d, 0x3d, 0x3f, 0xfc, 0x67, 0xcf, 0x5f, 0x4f, 0x77,
	0x17, 0x1c, 0x6f, 0x22, 0x83, 0x7a, 0x44, 0xd3, 0xa7, 0x06, 0x09, 0xdf, 0x80, 0x91, 0xae, 0x56,
	0x44, 0xda, 0xa9, 0xed, 0xab, 0x6a, 0x0c, 0x67, 0xb0, 0x35, 0xd2, 0xfe, 0x0d, 0xce, 0x38, 0x1a,
	0xf8, 0x5f, 0x70, 0xc6, 0xe1, 0x4b, 0x30, 0x14, 0x0a, 0xd2, 0xc4, 0x0d, 0xc1, 0x03, 0x1a, 0xa0,
	0x6b, 0x3d, 0xa1, 0x41, 0x82, 0x5a, 0x31, 0x24, 0x78, 0x07, 0x8c, 0x36, 0x9a, 0xc2, 0xdf, 0x56,
	0x38, 0xa6, 0x12, 0x77, 0x28, 0x91, 0x28, 0x5f, 0x74, 0x4a, 0x03, 0xde, 0x88, 0x0d, 0xd7, 0xa9,
	0x7c, 0x45, 0x89, 0x84, 0xcf, 0x01, 0x88, 0x48, 0x1b, 0xab, 0x56, 0x1c, 0x37, 0x3b, 0xe8, 0xba,
	0xd1, 0xbf, 0x9b, 0xea, 0x4f, 0x5d, 0xd4, 0xaf, 0x71, 0x7d, 0x46, 0xb9, 0xc6, 0xb5, 0x57, 0x88,
	0x48, 0x7b, 0xc3, 0x74, 0xc3, 0x27, 0xbf, 0x0e, 0x20, 0x0b, 0x28, 0xd7, 0x6c, 0x93, 0x51, 0x89,
	0x06, 0x0d, 0x71, 0xf2, 0x92, 0xd3, 0x55, 0xcb, 0x8a, 0xe1, 0x3a, 0x98, 0x91, 0x34, 0x68, 0xf9,
	0x66, 0xab, 0x63, 0x2a, 0x99, 0x08, 0x30, 0xe3, 0xd8, 0x94, 0x29, 0x54, 0x28, 0x3a, 0xa5, 0xfe,
	0x3f, 0xa0, 0x6e, 0x66, 0x7d, 0x75, 0xd3, 0x56, 0xe3, 0xab, 0xa6, 0x29, 0x19, 0x8a, 0x2e, 0x72,
	0x93, 0xf8, 0x5a, 0x48, 0x04, 0x7a, 0x1c, 0x8a, 0x0c, 0xf8, 0xcc, 0xf0, 0x60, 0x13, 0x94, 0x92,
	0x73, 0xcf, 0x78, 0x88, 0x25, 0xfd, 0x40, 0x64, 0xa0, 0x70, 0xc0, 0x94, 0x96, 0xac, 0xd1, 0x32,
	0xb2, 0x4a, 0x13, 0xa9, 0xad, 0x0b, 0x34, 0x74, 0x89, 0x89, 0xdb, 0x29, 0xc5, 0xb3, 0x90, 0xea,
	0x19, 0xc6, 0x46, 0x82, 0x30, 0x96, 0xe0, 0x7b, 0x80, 0x42, 0xca, 0xa9, 0x62, 0x0a, 0x5f, 0x18,
	0xf7, 0xe1, 0xde, 0x9c, 0xdd, 0x48, 0xc1, 0xab, 0xe7, 0xa7, 0xfe, 0xf1, 0xec, 0xc7, 0xd3, 0xdd,
	0x05, 0xd4, 0x2d, 0xaf, 0xb4, 0xed, 0xed, 0x6c, 0x6f, 0x82, 0x95, 0xe5, 0xfd, 0x63, 0xd7, 0x39,
	0x38, 0x76, 0x9d, 0x1f, 0xc7, 0xae, 0xf3, 0xe9, 0xc4, 0xcd, 0x1d, 0x9c, 0xb8, 0xb9, 0xef, 0x27,
	0x6e, 0xee, 0xf5, 0xf4, 0xb9, 0x15, 0xa4, 0x5d, 0xba, 0x13, 0x53, 0xd5, 0xc8, 0x9b, 0x4b, 0x77,
	0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x56, 0x80, 0x73, 0xef, 0x05, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EpochProvisions.Size()
		i -= size
		if _, err := m.EpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.GenesisEpochProvisions.Size()
		i -= size
		if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.MintingRewardsDistributionStartEpoch != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MintingRewardsDistributionStartEpoch))
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.ReductionPeriodInEpochs != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ReductionPeriodInEpochs))
		i--
		dAtA[i] = 0x48
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintMint(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.MaxSupply.Size()
		i -= size
		if _, err := m.MaxSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.EpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	l = m.MaxSupply.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.ReductionPeriodInEpochs != 0 {
		n += 1 + sovMint(uint64(m.ReductionPeriodInEpochs))
	}
	l = m.ReductionFactor.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.MintingRewardsDistributionStartEpoch != 0 {
		n += 1 + sovMint(uint64(m.MintingRewardsDistributionStartEpoch))
	}
	l = m.GenesisEpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInEpochs", wireType)
			}
			m.ReductionPeriodInEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingRewardsDistributionStartEpoch", wireType)
			}
			m.MintingRewardsDistributionStartEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintingRewardsDistributionStartEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
