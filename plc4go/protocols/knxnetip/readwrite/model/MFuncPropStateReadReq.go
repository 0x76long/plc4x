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

// MFuncPropStateReadReq is the corresponding interface of MFuncPropStateReadReq
type MFuncPropStateReadReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMFuncPropStateReadReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMFuncPropStateReadReq()
	// CreateBuilder creates a MFuncPropStateReadReqBuilder
	CreateMFuncPropStateReadReqBuilder() MFuncPropStateReadReqBuilder
}

// _MFuncPropStateReadReq is the data-structure of this message
type _MFuncPropStateReadReq struct {
	CEMIContract
}

var _ MFuncPropStateReadReq = (*_MFuncPropStateReadReq)(nil)
var _ CEMIRequirements = (*_MFuncPropStateReadReq)(nil)

// NewMFuncPropStateReadReq factory function for _MFuncPropStateReadReq
func NewMFuncPropStateReadReq(size uint16) *_MFuncPropStateReadReq {
	_result := &_MFuncPropStateReadReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MFuncPropStateReadReqBuilder is a builder for MFuncPropStateReadReq
type MFuncPropStateReadReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MFuncPropStateReadReqBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the MFuncPropStateReadReq or returns an error if something is wrong
	Build() (MFuncPropStateReadReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MFuncPropStateReadReq
}

// NewMFuncPropStateReadReqBuilder() creates a MFuncPropStateReadReqBuilder
func NewMFuncPropStateReadReqBuilder() MFuncPropStateReadReqBuilder {
	return &_MFuncPropStateReadReqBuilder{_MFuncPropStateReadReq: new(_MFuncPropStateReadReq)}
}

type _MFuncPropStateReadReqBuilder struct {
	*_MFuncPropStateReadReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MFuncPropStateReadReqBuilder) = (*_MFuncPropStateReadReqBuilder)(nil)

func (b *_MFuncPropStateReadReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._MFuncPropStateReadReq
}

func (b *_MFuncPropStateReadReqBuilder) WithMandatoryFields() MFuncPropStateReadReqBuilder {
	return b
}

func (b *_MFuncPropStateReadReqBuilder) Build() (MFuncPropStateReadReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MFuncPropStateReadReq.deepCopy(), nil
}

func (b *_MFuncPropStateReadReqBuilder) MustBuild() MFuncPropStateReadReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MFuncPropStateReadReqBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_MFuncPropStateReadReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MFuncPropStateReadReqBuilder) DeepCopy() any {
	_copy := b.CreateMFuncPropStateReadReqBuilder().(*_MFuncPropStateReadReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMFuncPropStateReadReqBuilder creates a MFuncPropStateReadReqBuilder
func (b *_MFuncPropStateReadReq) CreateMFuncPropStateReadReqBuilder() MFuncPropStateReadReqBuilder {
	if b == nil {
		return NewMFuncPropStateReadReqBuilder()
	}
	return &_MFuncPropStateReadReqBuilder{_MFuncPropStateReadReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropStateReadReq) GetMessageCode() uint8 {
	return 0xF9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropStateReadReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMFuncPropStateReadReq(structType any) MFuncPropStateReadReq {
	if casted, ok := structType.(MFuncPropStateReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropStateReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropStateReadReq) GetTypeName() string {
	return "MFuncPropStateReadReq"
}

func (m *_MFuncPropStateReadReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropStateReadReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MFuncPropStateReadReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mFuncPropStateReadReq MFuncPropStateReadReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropStateReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropStateReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropStateReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropStateReadReq")
	}

	return m, nil
}

func (m *_MFuncPropStateReadReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropStateReadReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropStateReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropStateReadReq")
		}

		if popErr := writeBuffer.PopContext("MFuncPropStateReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropStateReadReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropStateReadReq) IsMFuncPropStateReadReq() {}

func (m *_MFuncPropStateReadReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MFuncPropStateReadReq) deepCopy() *_MFuncPropStateReadReq {
	if m == nil {
		return nil
	}
	_MFuncPropStateReadReqCopy := &_MFuncPropStateReadReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	_MFuncPropStateReadReqCopy.CEMIContract.(*_CEMI)._SubType = m
	return _MFuncPropStateReadReqCopy
}

func (m *_MFuncPropStateReadReq) String() string {
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
