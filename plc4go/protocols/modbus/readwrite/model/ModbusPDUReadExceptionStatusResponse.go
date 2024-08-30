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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadExceptionStatusResponse is the corresponding interface of ModbusPDUReadExceptionStatusResponse
type ModbusPDUReadExceptionStatusResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() uint8
}

// ModbusPDUReadExceptionStatusResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadExceptionStatusResponse.
// This is useful for switch cases.
type ModbusPDUReadExceptionStatusResponseExactly interface {
	ModbusPDUReadExceptionStatusResponse
	isModbusPDUReadExceptionStatusResponse() bool
}

// _ModbusPDUReadExceptionStatusResponse is the data-structure of this message
type _ModbusPDUReadExceptionStatusResponse struct {
	*_ModbusPDU
	Value uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadExceptionStatusResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadExceptionStatusResponse) GetValue() uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadExceptionStatusResponse factory function for _ModbusPDUReadExceptionStatusResponse
func NewModbusPDUReadExceptionStatusResponse(value uint8) *_ModbusPDUReadExceptionStatusResponse {
	_result := &_ModbusPDUReadExceptionStatusResponse{
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadExceptionStatusResponse(structType any) ModbusPDUReadExceptionStatusResponse {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusResponse"
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusPDUReadExceptionStatusResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadExceptionStatusResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadExceptionStatusResponse, error) {
	return ModbusPDUReadExceptionStatusResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadExceptionStatusResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadExceptionStatusResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadExceptionStatusResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := /*TODO: migrate me*/ readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ModbusPDUReadExceptionStatusResponse")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadExceptionStatusResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadExceptionStatusResponse{
		_ModbusPDU: &_ModbusPDU{},
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadExceptionStatusResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadExceptionStatusResponse")
		}

		// Simple Field (value)
		value := uint8(m.GetValue())
		_valueErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("value", 8, uint8((value)))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadExceptionStatusResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadExceptionStatusResponse) isModbusPDUReadExceptionStatusResponse() bool {
	return true
}

func (m *_ModbusPDUReadExceptionStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
