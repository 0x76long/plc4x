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

// BACnetConfirmedServiceRequestAtomicReadFileStream is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileStream
type BACnetConfirmedServiceRequestAtomicReadFileStream interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetRequestOctetCount returns RequestOctetCount (property field)
	GetRequestOctetCount() BACnetApplicationTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestAtomicReadFileStream is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFileStream()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	CreateBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
}

// _BACnetConfirmedServiceRequestAtomicReadFileStream is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileStream struct {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
	FileStartPosition BACnetApplicationTagSignedInteger
	RequestOctetCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestAtomicReadFileStream = (*_BACnetConfirmedServiceRequestAtomicReadFileStream)(nil)
var _ BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordRequirements = (*_BACnetConfirmedServiceRequestAtomicReadFileStream)(nil)

// NewBACnetConfirmedServiceRequestAtomicReadFileStream factory function for _BACnetConfirmedServiceRequestAtomicReadFileStream
func NewBACnetConfirmedServiceRequestAtomicReadFileStream(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag, fileStartPosition BACnetApplicationTagSignedInteger, requestOctetCount BACnetApplicationTagUnsignedInteger) *_BACnetConfirmedServiceRequestAtomicReadFileStream {
	if fileStartPosition == nil {
		panic("fileStartPosition of type BACnetApplicationTagSignedInteger for BACnetConfirmedServiceRequestAtomicReadFileStream must not be nil")
	}
	if requestOctetCount == nil {
		panic("requestOctetCount of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestAtomicReadFileStream must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileStream{
		BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartPosition: fileStartPosition,
		RequestOctetCount: requestOctetCount,
	}
	_result.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder is a builder for BACnetConfirmedServiceRequestAtomicReadFileStream
type BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fileStartPosition BACnetApplicationTagSignedInteger, requestOctetCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	// WithFileStartPosition adds FileStartPosition (property field)
	WithFileStartPosition(BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	// WithFileStartPositionBuilder adds FileStartPosition (property field) which is build by the builder
	WithFileStartPositionBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	// WithRequestOctetCount adds RequestOctetCount (property field)
	WithRequestOctetCount(BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	// WithRequestOctetCountBuilder adds RequestOctetCount (property field) which is build by the builder
	WithRequestOctetCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// Build builds the BACnetConfirmedServiceRequestAtomicReadFileStream or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAtomicReadFileStream, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAtomicReadFileStream
}

// NewBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder() creates a BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
func NewBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	return &_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder{_BACnetConfirmedServiceRequestAtomicReadFileStream: new(_BACnetConfirmedServiceRequestAtomicReadFileStream)}
}

type _BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder struct {
	*_BACnetConfirmedServiceRequestAtomicReadFileStream

	parentBuilder *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) = (*_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) setParent(contract BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract) {
	b.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = contract
	contract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)._SubType = b._BACnetConfirmedServiceRequestAtomicReadFileStream
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) WithMandatoryFields(fileStartPosition BACnetApplicationTagSignedInteger, requestOctetCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	return b.WithFileStartPosition(fileStartPosition).WithRequestOctetCount(requestOctetCount)
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) WithFileStartPosition(fileStartPosition BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	b.FileStartPosition = fileStartPosition
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) WithFileStartPositionBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	builder := builderSupplier(b.FileStartPosition.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.FileStartPosition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) WithRequestOctetCount(requestOctetCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	b.RequestOctetCount = requestOctetCount
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) WithRequestOctetCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	builder := builderSupplier(b.RequestOctetCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RequestOctetCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) Build() (BACnetConfirmedServiceRequestAtomicReadFileStream, error) {
	if b.FileStartPosition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fileStartPosition' not set"))
	}
	if b.RequestOctetCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestOctetCount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestAtomicReadFileStream.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) MustBuild() BACnetConfirmedServiceRequestAtomicReadFileStream {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) Done() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder().(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) buildForBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord() (BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder().(*_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder
func (b *_BACnetConfirmedServiceRequestAtomicReadFileStream) CreateBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder() BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestAtomicReadFileStreamBuilder()
	}
	return &_BACnetConfirmedServiceRequestAtomicReadFileStreamBuilder{_BACnetConfirmedServiceRequestAtomicReadFileStream: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract {
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetRequestOctetCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestOctetCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileStream(structType any) BACnetConfirmedServiceRequestAtomicReadFileStream {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileStream"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	// Simple field (requestOctetCount)
	lengthInBits += m.RequestOctetCount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) (__bACnetConfirmedServiceRequestAtomicReadFileStream BACnetConfirmedServiceRequestAtomicReadFileStream, err error) {
	m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	requestOctetCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestOctetCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestOctetCount' field"))
	}
	m.RequestOctetCount = requestOctetCount

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestOctetCount", m.GetRequestOctetCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestOctetCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) IsBACnetConfirmedServiceRequestAtomicReadFileStream() {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) deepCopy() *_BACnetConfirmedServiceRequestAtomicReadFileStream {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicReadFileStreamCopy := &_BACnetConfirmedServiceRequestAtomicReadFileStream{
		m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.FileStartPosition),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RequestOctetCount),
	}
	_BACnetConfirmedServiceRequestAtomicReadFileStreamCopy.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)._SubType = m
	return _BACnetConfirmedServiceRequestAtomicReadFileStreamCopy
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) String() string {
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
