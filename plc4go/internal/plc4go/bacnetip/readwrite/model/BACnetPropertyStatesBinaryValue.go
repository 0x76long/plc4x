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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetPropertyStatesBinaryValue struct {
	*BACnetPropertyStates
	BinaryValue *BACnetBinaryPV

	// Arguments.
	TagNumber uint8
}

// The corresponding interface
type IBACnetPropertyStatesBinaryValue interface {
	IBACnetPropertyStates
	// GetBinaryValue returns BinaryValue (property field)
	GetBinaryValue() *BACnetBinaryPV
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
///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesBinaryValue) InitializeParent(parent *BACnetPropertyStates, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetPropertyStates.OpeningTag = openingTag
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
	m.BACnetPropertyStates.ClosingTag = closingTag
}

func (m *BACnetPropertyStatesBinaryValue) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetPropertyStatesBinaryValue) GetBinaryValue() *BACnetBinaryPV {
	return m.BinaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBinaryValue factory function for BACnetPropertyStatesBinaryValue
func NewBACnetPropertyStatesBinaryValue(binaryValue *BACnetBinaryPV, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetPropertyStates {
	child := &BACnetPropertyStatesBinaryValue{
		BinaryValue:          binaryValue,
		BACnetPropertyStates: NewBACnetPropertyStates(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	child.Child = child
	return child.BACnetPropertyStates
}

func CastBACnetPropertyStatesBinaryValue(structType interface{}) *BACnetPropertyStatesBinaryValue {
	if casted, ok := structType.(BACnetPropertyStatesBinaryValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBinaryValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesBinaryValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryValue"
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (binaryValue)
	if m.BinaryValue != nil {
		lengthInBits += (*m.BinaryValue).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyStatesBinaryValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBinaryValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetPropertyStates, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (binaryValue) (Can be skipped, if a given expression evaluates to false)
	var binaryValue *BACnetBinaryPV = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("binaryValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetBinaryPVParse(readBuffer, peekedTagNumber)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'binaryValue' field")
		default:
			binaryValue = CastBACnetBinaryPV(_val)
			if closeErr := readBuffer.CloseContext("binaryValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesBinaryValue{
		BinaryValue:          CastBACnetBinaryPV(binaryValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child.BACnetPropertyStates, nil
}

func (m *BACnetPropertyStatesBinaryValue) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryValue"); pushErr != nil {
			return pushErr
		}

		// Optional Field (binaryValue) (Can be skipped, if the value is null)
		var binaryValue *BACnetBinaryPV = nil
		if m.BinaryValue != nil {
			if pushErr := writeBuffer.PushContext("binaryValue"); pushErr != nil {
				return pushErr
			}
			binaryValue = m.BinaryValue
			_binaryValueErr := binaryValue.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("binaryValue"); popErr != nil {
				return popErr
			}
			if _binaryValueErr != nil {
				return errors.Wrap(_binaryValueErr, "Error serializing 'binaryValue' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesBinaryValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
