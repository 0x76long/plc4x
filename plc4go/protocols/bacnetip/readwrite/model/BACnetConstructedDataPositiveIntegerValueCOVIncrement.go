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

// BACnetConstructedDataPositiveIntegerValueCOVIncrement is the corresponding interface of BACnetConstructedDataPositiveIntegerValueCOVIncrement
type BACnetConstructedDataPositiveIntegerValueCOVIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPositiveIntegerValueCOVIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPositiveIntegerValueCOVIncrement()
	// CreateBuilder creates a BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
	CreateBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder() BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
}

// _BACnetConstructedDataPositiveIntegerValueCOVIncrement is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueCOVIncrement struct {
	BACnetConstructedDataContract
	CovIncrement BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPositiveIntegerValueCOVIncrement = (*_BACnetConstructedDataPositiveIntegerValueCOVIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPositiveIntegerValueCOVIncrement)(nil)

// NewBACnetConstructedDataPositiveIntegerValueCOVIncrement factory function for _BACnetConstructedDataPositiveIntegerValueCOVIncrement
func NewBACnetConstructedDataPositiveIntegerValueCOVIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covIncrement BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueCOVIncrement {
	if covIncrement == nil {
		panic("covIncrement of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPositiveIntegerValueCOVIncrement must not be nil")
	}
	_result := &_BACnetConstructedDataPositiveIntegerValueCOVIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovIncrement:                  covIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder is a builder for BACnetConstructedDataPositiveIntegerValueCOVIncrement
type BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(covIncrement BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
	// WithCovIncrement adds CovIncrement (property field)
	WithCovIncrement(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
	// WithCovIncrementBuilder adds CovIncrement (property field) which is build by the builder
	WithCovIncrementBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPositiveIntegerValueCOVIncrement or returns an error if something is wrong
	Build() (BACnetConstructedDataPositiveIntegerValueCOVIncrement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPositiveIntegerValueCOVIncrement
}

// NewBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder() creates a BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
func NewBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder() BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder {
	return &_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder{_BACnetConstructedDataPositiveIntegerValueCOVIncrement: new(_BACnetConstructedDataPositiveIntegerValueCOVIncrement)}
}

type _BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder struct {
	*_BACnetConstructedDataPositiveIntegerValueCOVIncrement

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) = (*_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder)(nil)

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataPositiveIntegerValueCOVIncrement
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) WithMandatoryFields(covIncrement BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder {
	return b.WithCovIncrement(covIncrement)
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) WithCovIncrement(covIncrement BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder {
	b.CovIncrement = covIncrement
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) WithCovIncrementBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder {
	builder := builderSupplier(b.CovIncrement.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.CovIncrement, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) Build() (BACnetConstructedDataPositiveIntegerValueCOVIncrement, error) {
	if b.CovIncrement == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'covIncrement' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPositiveIntegerValueCOVIncrement.deepCopy(), nil
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) MustBuild() BACnetConstructedDataPositiveIntegerValueCOVIncrement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder().(*_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder creates a BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder
func (b *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) CreateBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder() BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder {
	if b == nil {
		return NewBACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder()
	}
	return &_BACnetConstructedDataPositiveIntegerValueCOVIncrementBuilder{_BACnetConstructedDataPositiveIntegerValueCOVIncrement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetCovIncrement() BACnetApplicationTagUnsignedInteger {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueCOVIncrement(structType any) BACnetConstructedDataPositiveIntegerValueCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueCOVIncrement"
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPositiveIntegerValueCOVIncrement BACnetConstructedDataPositiveIntegerValueCOVIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covIncrement, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covIncrement", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	m.CovIncrement = covIncrement

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueCOVIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueCOVIncrement")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covIncrement", m.GetCovIncrement(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueCOVIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) IsBACnetConstructedDataPositiveIntegerValueCOVIncrement() {
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) deepCopy() *_BACnetConstructedDataPositiveIntegerValueCOVIncrement {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPositiveIntegerValueCOVIncrementCopy := &_BACnetConstructedDataPositiveIntegerValueCOVIncrement{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.CovIncrement),
	}
	_BACnetConstructedDataPositiveIntegerValueCOVIncrementCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPositiveIntegerValueCOVIncrementCopy
}

func (m *_BACnetConstructedDataPositiveIntegerValueCOVIncrement) String() string {
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
