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

// BACnetConstructedDataProportionalConstant is the corresponding interface of BACnetConstructedDataProportionalConstant
type BACnetConstructedDataProportionalConstant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProportionalConstant returns ProportionalConstant (property field)
	GetProportionalConstant() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataProportionalConstant is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProportionalConstant()
	// CreateBuilder creates a BACnetConstructedDataProportionalConstantBuilder
	CreateBACnetConstructedDataProportionalConstantBuilder() BACnetConstructedDataProportionalConstantBuilder
}

// _BACnetConstructedDataProportionalConstant is the data-structure of this message
type _BACnetConstructedDataProportionalConstant struct {
	BACnetConstructedDataContract
	ProportionalConstant BACnetApplicationTagReal
}

var _ BACnetConstructedDataProportionalConstant = (*_BACnetConstructedDataProportionalConstant)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProportionalConstant)(nil)

// NewBACnetConstructedDataProportionalConstant factory function for _BACnetConstructedDataProportionalConstant
func NewBACnetConstructedDataProportionalConstant(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, proportionalConstant BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProportionalConstant {
	if proportionalConstant == nil {
		panic("proportionalConstant of type BACnetApplicationTagReal for BACnetConstructedDataProportionalConstant must not be nil")
	}
	_result := &_BACnetConstructedDataProportionalConstant{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProportionalConstant:          proportionalConstant,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProportionalConstantBuilder is a builder for BACnetConstructedDataProportionalConstant
type BACnetConstructedDataProportionalConstantBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(proportionalConstant BACnetApplicationTagReal) BACnetConstructedDataProportionalConstantBuilder
	// WithProportionalConstant adds ProportionalConstant (property field)
	WithProportionalConstant(BACnetApplicationTagReal) BACnetConstructedDataProportionalConstantBuilder
	// WithProportionalConstantBuilder adds ProportionalConstant (property field) which is build by the builder
	WithProportionalConstantBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataProportionalConstantBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataProportionalConstant or returns an error if something is wrong
	Build() (BACnetConstructedDataProportionalConstant, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProportionalConstant
}

// NewBACnetConstructedDataProportionalConstantBuilder() creates a BACnetConstructedDataProportionalConstantBuilder
func NewBACnetConstructedDataProportionalConstantBuilder() BACnetConstructedDataProportionalConstantBuilder {
	return &_BACnetConstructedDataProportionalConstantBuilder{_BACnetConstructedDataProportionalConstant: new(_BACnetConstructedDataProportionalConstant)}
}

type _BACnetConstructedDataProportionalConstantBuilder struct {
	*_BACnetConstructedDataProportionalConstant

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataProportionalConstantBuilder) = (*_BACnetConstructedDataProportionalConstantBuilder)(nil)

func (b *_BACnetConstructedDataProportionalConstantBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataProportionalConstant
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) WithMandatoryFields(proportionalConstant BACnetApplicationTagReal) BACnetConstructedDataProportionalConstantBuilder {
	return b.WithProportionalConstant(proportionalConstant)
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) WithProportionalConstant(proportionalConstant BACnetApplicationTagReal) BACnetConstructedDataProportionalConstantBuilder {
	b.ProportionalConstant = proportionalConstant
	return b
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) WithProportionalConstantBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataProportionalConstantBuilder {
	builder := builderSupplier(b.ProportionalConstant.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.ProportionalConstant, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) Build() (BACnetConstructedDataProportionalConstant, error) {
	if b.ProportionalConstant == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'proportionalConstant' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataProportionalConstant.deepCopy(), nil
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) MustBuild() BACnetConstructedDataProportionalConstant {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataProportionalConstantBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataProportionalConstantBuilder().(*_BACnetConstructedDataProportionalConstantBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataProportionalConstantBuilder creates a BACnetConstructedDataProportionalConstantBuilder
func (b *_BACnetConstructedDataProportionalConstant) CreateBACnetConstructedDataProportionalConstantBuilder() BACnetConstructedDataProportionalConstantBuilder {
	if b == nil {
		return NewBACnetConstructedDataProportionalConstantBuilder()
	}
	return &_BACnetConstructedDataProportionalConstantBuilder{_BACnetConstructedDataProportionalConstant: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProportionalConstant) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROPORTIONAL_CONSTANT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetProportionalConstant() BACnetApplicationTagReal {
	return m.ProportionalConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProportionalConstant) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetProportionalConstant())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProportionalConstant(structType any) BACnetConstructedDataProportionalConstant {
	if casted, ok := structType.(BACnetConstructedDataProportionalConstant); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProportionalConstant); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProportionalConstant) GetTypeName() string {
	return "BACnetConstructedDataProportionalConstant"
}

func (m *_BACnetConstructedDataProportionalConstant) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (proportionalConstant)
	lengthInBits += m.ProportionalConstant.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProportionalConstant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProportionalConstant) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProportionalConstant BACnetConstructedDataProportionalConstant, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProportionalConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProportionalConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	proportionalConstant, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "proportionalConstant", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proportionalConstant' field"))
	}
	m.ProportionalConstant = proportionalConstant

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), proportionalConstant)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProportionalConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProportionalConstant")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProportionalConstant) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProportionalConstant) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProportionalConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProportionalConstant")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "proportionalConstant", m.GetProportionalConstant(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'proportionalConstant' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProportionalConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProportionalConstant")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProportionalConstant) IsBACnetConstructedDataProportionalConstant() {}

func (m *_BACnetConstructedDataProportionalConstant) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProportionalConstant) deepCopy() *_BACnetConstructedDataProportionalConstant {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProportionalConstantCopy := &_BACnetConstructedDataProportionalConstant{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.ProportionalConstant),
	}
	_BACnetConstructedDataProportionalConstantCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProportionalConstantCopy
}

func (m *_BACnetConstructedDataProportionalConstant) String() string {
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
