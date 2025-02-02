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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtDomainAddressWrite is the corresponding interface of ApduDataExtDomainAddressWrite
type ApduDataExtDomainAddressWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtDomainAddressWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtDomainAddressWrite()
	// CreateBuilder creates a ApduDataExtDomainAddressWriteBuilder
	CreateApduDataExtDomainAddressWriteBuilder() ApduDataExtDomainAddressWriteBuilder
}

// _ApduDataExtDomainAddressWrite is the data-structure of this message
type _ApduDataExtDomainAddressWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtDomainAddressWrite = (*_ApduDataExtDomainAddressWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtDomainAddressWrite)(nil)

// NewApduDataExtDomainAddressWrite factory function for _ApduDataExtDomainAddressWrite
func NewApduDataExtDomainAddressWrite(length uint8) *_ApduDataExtDomainAddressWrite {
	_result := &_ApduDataExtDomainAddressWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtDomainAddressWriteBuilder is a builder for ApduDataExtDomainAddressWrite
type ApduDataExtDomainAddressWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtDomainAddressWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtDomainAddressWrite or returns an error if something is wrong
	Build() (ApduDataExtDomainAddressWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtDomainAddressWrite
}

// NewApduDataExtDomainAddressWriteBuilder() creates a ApduDataExtDomainAddressWriteBuilder
func NewApduDataExtDomainAddressWriteBuilder() ApduDataExtDomainAddressWriteBuilder {
	return &_ApduDataExtDomainAddressWriteBuilder{_ApduDataExtDomainAddressWrite: new(_ApduDataExtDomainAddressWrite)}
}

type _ApduDataExtDomainAddressWriteBuilder struct {
	*_ApduDataExtDomainAddressWrite

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtDomainAddressWriteBuilder) = (*_ApduDataExtDomainAddressWriteBuilder)(nil)

func (b *_ApduDataExtDomainAddressWriteBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtDomainAddressWrite
}

func (b *_ApduDataExtDomainAddressWriteBuilder) WithMandatoryFields() ApduDataExtDomainAddressWriteBuilder {
	return b
}

func (b *_ApduDataExtDomainAddressWriteBuilder) Build() (ApduDataExtDomainAddressWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtDomainAddressWrite.deepCopy(), nil
}

func (b *_ApduDataExtDomainAddressWriteBuilder) MustBuild() ApduDataExtDomainAddressWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtDomainAddressWriteBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtDomainAddressWriteBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtDomainAddressWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtDomainAddressWriteBuilder().(*_ApduDataExtDomainAddressWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtDomainAddressWriteBuilder creates a ApduDataExtDomainAddressWriteBuilder
func (b *_ApduDataExtDomainAddressWrite) CreateApduDataExtDomainAddressWriteBuilder() ApduDataExtDomainAddressWriteBuilder {
	if b == nil {
		return NewApduDataExtDomainAddressWriteBuilder()
	}
	return &_ApduDataExtDomainAddressWriteBuilder{_ApduDataExtDomainAddressWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressWrite) GetExtApciType() uint8 {
	return 0x20
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressWrite(structType any) ApduDataExtDomainAddressWrite {
	if casted, ok := structType.(ApduDataExtDomainAddressWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressWrite) GetTypeName() string {
	return "ApduDataExtDomainAddressWrite"
}

func (m *_ApduDataExtDomainAddressWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtDomainAddressWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtDomainAddressWrite ApduDataExtDomainAddressWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressWrite")
	}

	return m, nil
}

func (m *_ApduDataExtDomainAddressWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressWrite) IsApduDataExtDomainAddressWrite() {}

func (m *_ApduDataExtDomainAddressWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtDomainAddressWrite) deepCopy() *_ApduDataExtDomainAddressWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtDomainAddressWriteCopy := &_ApduDataExtDomainAddressWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtDomainAddressWriteCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtDomainAddressWriteCopy
}

func (m *_ApduDataExtDomainAddressWrite) String() string {
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
