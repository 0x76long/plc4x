/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReadExceptionStatusResponse struct {
	*ModbusPDU
	Value uint8
}

// The corresponding interface
type IModbusPDUReadExceptionStatusResponse interface {
	IModbusPDU
	// GetValue returns Value (property field)
	GetValue() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *ModbusPDUReadExceptionStatusResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusResponse) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *ModbusPDUReadExceptionStatusResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadExceptionStatusResponse) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadExceptionStatusResponse) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusPDUReadExceptionStatusResponse) GetValue() uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadExceptionStatusResponse factory function for ModbusPDUReadExceptionStatusResponse
func NewModbusPDUReadExceptionStatusResponse(value uint8) *ModbusPDU {
	child := &ModbusPDUReadExceptionStatusResponse{
		Value:     value,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadExceptionStatusResponse(structType interface{}) *ModbusPDUReadExceptionStatusResponse {
	if casted, ok := structType.(ModbusPDUReadExceptionStatusResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadExceptionStatusResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadExceptionStatusResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadExceptionStatusResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadExceptionStatusResponse) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusResponse"
}

func (m *ModbusPDUReadExceptionStatusResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadExceptionStatusResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *ModbusPDUReadExceptionStatusResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadExceptionStatusResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadExceptionStatusResponse{
		Value:     value,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadExceptionStatusResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := uint8(m.Value)
		_valueErr := writeBuffer.WriteUint8("value", 8, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadExceptionStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
