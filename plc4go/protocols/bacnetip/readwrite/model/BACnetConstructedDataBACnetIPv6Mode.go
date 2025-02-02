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

// BACnetConstructedDataBACnetIPv6Mode is the corresponding interface of BACnetConstructedDataBACnetIPv6Mode
type BACnetConstructedDataBACnetIPv6Mode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBacnetIpv6Mode returns BacnetIpv6Mode (property field)
	GetBacnetIpv6Mode() BACnetIPModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetIPModeTagged
	// IsBACnetConstructedDataBACnetIPv6Mode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPv6Mode()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPv6ModeBuilder
	CreateBACnetConstructedDataBACnetIPv6ModeBuilder() BACnetConstructedDataBACnetIPv6ModeBuilder
}

// _BACnetConstructedDataBACnetIPv6Mode is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6Mode struct {
	BACnetConstructedDataContract
	BacnetIpv6Mode BACnetIPModeTagged
}

var _ BACnetConstructedDataBACnetIPv6Mode = (*_BACnetConstructedDataBACnetIPv6Mode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPv6Mode)(nil)

// NewBACnetConstructedDataBACnetIPv6Mode factory function for _BACnetConstructedDataBACnetIPv6Mode
func NewBACnetConstructedDataBACnetIPv6Mode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, bacnetIpv6Mode BACnetIPModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6Mode {
	if bacnetIpv6Mode == nil {
		panic("bacnetIpv6Mode of type BACnetIPModeTagged for BACnetConstructedDataBACnetIPv6Mode must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPv6Mode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BacnetIpv6Mode:                bacnetIpv6Mode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPv6ModeBuilder is a builder for BACnetConstructedDataBACnetIPv6Mode
type BACnetConstructedDataBACnetIPv6ModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bacnetIpv6Mode BACnetIPModeTagged) BACnetConstructedDataBACnetIPv6ModeBuilder
	// WithBacnetIpv6Mode adds BacnetIpv6Mode (property field)
	WithBacnetIpv6Mode(BACnetIPModeTagged) BACnetConstructedDataBACnetIPv6ModeBuilder
	// WithBacnetIpv6ModeBuilder adds BacnetIpv6Mode (property field) which is build by the builder
	WithBacnetIpv6ModeBuilder(func(BACnetIPModeTaggedBuilder) BACnetIPModeTaggedBuilder) BACnetConstructedDataBACnetIPv6ModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBACnetIPv6Mode or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPv6Mode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPv6Mode
}

// NewBACnetConstructedDataBACnetIPv6ModeBuilder() creates a BACnetConstructedDataBACnetIPv6ModeBuilder
func NewBACnetConstructedDataBACnetIPv6ModeBuilder() BACnetConstructedDataBACnetIPv6ModeBuilder {
	return &_BACnetConstructedDataBACnetIPv6ModeBuilder{_BACnetConstructedDataBACnetIPv6Mode: new(_BACnetConstructedDataBACnetIPv6Mode)}
}

type _BACnetConstructedDataBACnetIPv6ModeBuilder struct {
	*_BACnetConstructedDataBACnetIPv6Mode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPv6ModeBuilder) = (*_BACnetConstructedDataBACnetIPv6ModeBuilder)(nil)

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBACnetIPv6Mode
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) WithMandatoryFields(bacnetIpv6Mode BACnetIPModeTagged) BACnetConstructedDataBACnetIPv6ModeBuilder {
	return b.WithBacnetIpv6Mode(bacnetIpv6Mode)
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) WithBacnetIpv6Mode(bacnetIpv6Mode BACnetIPModeTagged) BACnetConstructedDataBACnetIPv6ModeBuilder {
	b.BacnetIpv6Mode = bacnetIpv6Mode
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) WithBacnetIpv6ModeBuilder(builderSupplier func(BACnetIPModeTaggedBuilder) BACnetIPModeTaggedBuilder) BACnetConstructedDataBACnetIPv6ModeBuilder {
	builder := builderSupplier(b.BacnetIpv6Mode.CreateBACnetIPModeTaggedBuilder())
	var err error
	b.BacnetIpv6Mode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetIPModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) Build() (BACnetConstructedDataBACnetIPv6Mode, error) {
	if b.BacnetIpv6Mode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bacnetIpv6Mode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBACnetIPv6Mode.deepCopy(), nil
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) MustBuild() BACnetConstructedDataBACnetIPv6Mode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBACnetIPv6ModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBACnetIPv6ModeBuilder().(*_BACnetConstructedDataBACnetIPv6ModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBACnetIPv6ModeBuilder creates a BACnetConstructedDataBACnetIPv6ModeBuilder
func (b *_BACnetConstructedDataBACnetIPv6Mode) CreateBACnetConstructedDataBACnetIPv6ModeBuilder() BACnetConstructedDataBACnetIPv6ModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataBACnetIPv6ModeBuilder()
	}
	return &_BACnetConstructedDataBACnetIPv6ModeBuilder{_BACnetConstructedDataBACnetIPv6Mode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetBacnetIpv6Mode() BACnetIPModeTagged {
	return m.BacnetIpv6Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetActualValue() BACnetIPModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetIPModeTagged(m.GetBacnetIpv6Mode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6Mode(structType any) BACnetConstructedDataBACnetIPv6Mode {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6Mode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6Mode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6Mode"
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (bacnetIpv6Mode)
	lengthInBits += m.BacnetIpv6Mode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPv6Mode BACnetConstructedDataBACnetIPv6Mode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6Mode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6Mode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetIpv6Mode, err := ReadSimpleField[BACnetIPModeTagged](ctx, "bacnetIpv6Mode", ReadComplex[BACnetIPModeTagged](BACnetIPModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetIpv6Mode' field"))
	}
	m.BacnetIpv6Mode = bacnetIpv6Mode

	actualValue, err := ReadVirtualField[BACnetIPModeTagged](ctx, "actualValue", (*BACnetIPModeTagged)(nil), bacnetIpv6Mode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6Mode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6Mode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6Mode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6Mode")
		}

		if err := WriteSimpleField[BACnetIPModeTagged](ctx, "bacnetIpv6Mode", m.GetBacnetIpv6Mode(), WriteComplex[BACnetIPModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bacnetIpv6Mode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6Mode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6Mode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) IsBACnetConstructedDataBACnetIPv6Mode() {}

func (m *_BACnetConstructedDataBACnetIPv6Mode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) deepCopy() *_BACnetConstructedDataBACnetIPv6Mode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPv6ModeCopy := &_BACnetConstructedDataBACnetIPv6Mode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetIPModeTagged](m.BacnetIpv6Mode),
	}
	_BACnetConstructedDataBACnetIPv6ModeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPv6ModeCopy
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) String() string {
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
