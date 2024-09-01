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

// S7ParameterModeTransition is the corresponding interface of S7ParameterModeTransition
type S7ParameterModeTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCurrentMode returns CurrentMode (property field)
	GetCurrentMode() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
}

// S7ParameterModeTransitionExactly can be used when we want exactly this type and not a type which fulfills S7ParameterModeTransition.
// This is useful for switch cases.
type S7ParameterModeTransitionExactly interface {
	S7ParameterModeTransition
	isS7ParameterModeTransition() bool
}

// _S7ParameterModeTransition is the data-structure of this message
type _S7ParameterModeTransition struct {
	*_S7Parameter
	Method           uint8
	CpuFunctionType  uint8
	CpuFunctionGroup uint8
	CurrentMode      uint8
	SequenceNumber   uint8
	// Reserved Fields
	reservedField0 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterModeTransition) GetParameterType() uint8 {
	return 0x01
}

func (m *_S7ParameterModeTransition) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterModeTransition) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterModeTransition) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterModeTransition) GetMethod() uint8 {
	return m.Method
}

func (m *_S7ParameterModeTransition) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *_S7ParameterModeTransition) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *_S7ParameterModeTransition) GetCurrentMode() uint8 {
	return m.CurrentMode
}

func (m *_S7ParameterModeTransition) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterModeTransition factory function for _S7ParameterModeTransition
func NewS7ParameterModeTransition(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, currentMode uint8, sequenceNumber uint8) *_S7ParameterModeTransition {
	_result := &_S7ParameterModeTransition{
		Method:           method,
		CpuFunctionType:  cpuFunctionType,
		CpuFunctionGroup: cpuFunctionGroup,
		CurrentMode:      currentMode,
		SequenceNumber:   sequenceNumber,
		_S7Parameter:     NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterModeTransition(structType any) S7ParameterModeTransition {
	if casted, ok := structType.(S7ParameterModeTransition); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterModeTransition); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterModeTransition) GetTypeName() string {
	return "S7ParameterModeTransition"
}

func (m *_S7ParameterModeTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (currentMode)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterModeTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterModeTransitionParse(ctx context.Context, theBytes []byte, messageType uint8) (S7ParameterModeTransition, error) {
	return S7ParameterModeTransitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterModeTransitionParseWithBufferProducer(messageType uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7ParameterModeTransition, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7ParameterModeTransition, error) {
		return S7ParameterModeTransitionParseWithBuffer(ctx, readBuffer, messageType)
	}
}

func S7ParameterModeTransitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterModeTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterModeTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterModeTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0010))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	itemLength, err := ReadImplicitField[uint8](ctx, "itemLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	method, err := ReadSimpleField(ctx, "method", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'method' field"))
	}

	cpuFunctionType, err := ReadSimpleField(ctx, "cpuFunctionType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionType' field"))
	}

	cpuFunctionGroup, err := ReadSimpleField(ctx, "cpuFunctionGroup", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionGroup' field"))
	}

	currentMode, err := ReadSimpleField(ctx, "currentMode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentMode' field"))
	}

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}

	if closeErr := readBuffer.CloseContext("S7ParameterModeTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterModeTransition")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterModeTransition{
		_S7Parameter:     &_S7Parameter{},
		Method:           method,
		CpuFunctionType:  cpuFunctionType,
		CpuFunctionGroup: cpuFunctionGroup,
		CurrentMode:      currentMode,
		SequenceNumber:   sequenceNumber,
		reservedField0:   reservedField0,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterModeTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterModeTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterModeTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterModeTransition")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0010), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		// Simple Field (method)
		method := uint8(m.GetMethod())
		_methodErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("method", 8, uint8((method)))
		if _methodErr != nil {
			return errors.Wrap(_methodErr, "Error serializing 'method' field")
		}

		// Simple Field (cpuFunctionType)
		cpuFunctionType := uint8(m.GetCpuFunctionType())
		_cpuFunctionTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("cpuFunctionType", 4, uint8((cpuFunctionType)))
		if _cpuFunctionTypeErr != nil {
			return errors.Wrap(_cpuFunctionTypeErr, "Error serializing 'cpuFunctionType' field")
		}

		// Simple Field (cpuFunctionGroup)
		cpuFunctionGroup := uint8(m.GetCpuFunctionGroup())
		_cpuFunctionGroupErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("cpuFunctionGroup", 4, uint8((cpuFunctionGroup)))
		if _cpuFunctionGroupErr != nil {
			return errors.Wrap(_cpuFunctionGroupErr, "Error serializing 'cpuFunctionGroup' field")
		}

		// Simple Field (currentMode)
		currentMode := uint8(m.GetCurrentMode())
		_currentModeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("currentMode", 8, uint8((currentMode)))
		if _currentModeErr != nil {
			return errors.Wrap(_currentModeErr, "Error serializing 'currentMode' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.GetSequenceNumber())
		_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sequenceNumber", 8, uint8((sequenceNumber)))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterModeTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterModeTransition")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterModeTransition) isS7ParameterModeTransition() bool {
	return true
}

func (m *_S7ParameterModeTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
