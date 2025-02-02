/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataIPv6DefaultGateway is the corresponding interface of BACnetConstructedDataIPv6DefaultGateway
type BACnetConstructedDataIPv6DefaultGateway interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6DefaultGateway returns Ipv6DefaultGateway (property field)
	GetIpv6DefaultGateway() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataIPv6DefaultGateway is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6DefaultGateway()
	// CreateBuilder creates a BACnetConstructedDataIPv6DefaultGatewayBuilder
	CreateBACnetConstructedDataIPv6DefaultGatewayBuilder() BACnetConstructedDataIPv6DefaultGatewayBuilder
}

// _BACnetConstructedDataIPv6DefaultGateway is the data-structure of this message
type _BACnetConstructedDataIPv6DefaultGateway struct {
	BACnetConstructedDataContract
	Ipv6DefaultGateway BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPv6DefaultGateway = (*_BACnetConstructedDataIPv6DefaultGateway)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6DefaultGateway)(nil)

// NewBACnetConstructedDataIPv6DefaultGateway factory function for _BACnetConstructedDataIPv6DefaultGateway
func NewBACnetConstructedDataIPv6DefaultGateway(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6DefaultGateway BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DefaultGateway {
	if ipv6DefaultGateway == nil {
		panic("ipv6DefaultGateway of type BACnetApplicationTagOctetString for BACnetConstructedDataIPv6DefaultGateway must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6DefaultGateway{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6DefaultGateway:            ipv6DefaultGateway,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6DefaultGatewayBuilder is a builder for BACnetConstructedDataIPv6DefaultGateway
type BACnetConstructedDataIPv6DefaultGatewayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6DefaultGateway BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DefaultGatewayBuilder
	// WithIpv6DefaultGateway adds Ipv6DefaultGateway (property field)
	WithIpv6DefaultGateway(BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DefaultGatewayBuilder
	// WithIpv6DefaultGatewayBuilder adds Ipv6DefaultGateway (property field) which is build by the builder
	WithIpv6DefaultGatewayBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6DefaultGatewayBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPv6DefaultGateway or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6DefaultGateway, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6DefaultGateway
}

// NewBACnetConstructedDataIPv6DefaultGatewayBuilder() creates a BACnetConstructedDataIPv6DefaultGatewayBuilder
func NewBACnetConstructedDataIPv6DefaultGatewayBuilder() BACnetConstructedDataIPv6DefaultGatewayBuilder {
	return &_BACnetConstructedDataIPv6DefaultGatewayBuilder{_BACnetConstructedDataIPv6DefaultGateway: new(_BACnetConstructedDataIPv6DefaultGateway)}
}

type _BACnetConstructedDataIPv6DefaultGatewayBuilder struct {
	*_BACnetConstructedDataIPv6DefaultGateway

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6DefaultGatewayBuilder) = (*_BACnetConstructedDataIPv6DefaultGatewayBuilder)(nil)

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPv6DefaultGateway
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) WithMandatoryFields(ipv6DefaultGateway BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DefaultGatewayBuilder {
	return b.WithIpv6DefaultGateway(ipv6DefaultGateway)
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) WithIpv6DefaultGateway(ipv6DefaultGateway BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DefaultGatewayBuilder {
	b.Ipv6DefaultGateway = ipv6DefaultGateway
	return b
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) WithIpv6DefaultGatewayBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6DefaultGatewayBuilder {
	builder := builderSupplier(b.Ipv6DefaultGateway.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.Ipv6DefaultGateway, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) Build() (BACnetConstructedDataIPv6DefaultGateway, error) {
	if b.Ipv6DefaultGateway == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6DefaultGateway' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6DefaultGateway.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) MustBuild() BACnetConstructedDataIPv6DefaultGateway {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6DefaultGatewayBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6DefaultGatewayBuilder().(*_BACnetConstructedDataIPv6DefaultGatewayBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6DefaultGatewayBuilder creates a BACnetConstructedDataIPv6DefaultGatewayBuilder
func (b *_BACnetConstructedDataIPv6DefaultGateway) CreateBACnetConstructedDataIPv6DefaultGatewayBuilder() BACnetConstructedDataIPv6DefaultGatewayBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6DefaultGatewayBuilder()
	}
	return &_BACnetConstructedDataIPv6DefaultGatewayBuilder{_BACnetConstructedDataIPv6DefaultGateway: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DEFAULT_GATEWAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetIpv6DefaultGateway() BACnetApplicationTagOctetString {
	return m.Ipv6DefaultGateway
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6DefaultGateway())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DefaultGateway(structType any) BACnetConstructedDataIPv6DefaultGateway {
	if casted, ok := structType.(BACnetConstructedDataIPv6DefaultGateway); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DefaultGateway); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetTypeName() string {
	return "BACnetConstructedDataIPv6DefaultGateway"
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6DefaultGateway)
	lengthInBits += m.Ipv6DefaultGateway.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6DefaultGateway BACnetConstructedDataIPv6DefaultGateway, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DefaultGateway"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DefaultGateway")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6DefaultGateway, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6DefaultGateway", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6DefaultGateway' field"))
	}
	m.Ipv6DefaultGateway = ipv6DefaultGateway

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipv6DefaultGateway)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DefaultGateway"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DefaultGateway")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DefaultGateway"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DefaultGateway")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6DefaultGateway", m.GetIpv6DefaultGateway(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6DefaultGateway' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DefaultGateway"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DefaultGateway")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) IsBACnetConstructedDataIPv6DefaultGateway() {}

func (m *_BACnetConstructedDataIPv6DefaultGateway) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) deepCopy() *_BACnetConstructedDataIPv6DefaultGateway {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6DefaultGatewayCopy := &_BACnetConstructedDataIPv6DefaultGateway{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.Ipv6DefaultGateway),
	}
	_BACnetConstructedDataIPv6DefaultGatewayCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6DefaultGatewayCopy
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
