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
type BACnetContextTagBitString struct {
	*BACnetContextTag
	Payload *BACnetTagPayloadBitString

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
}

// The corresponding interface
type IBACnetContextTagBitString interface {
	IBACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadBitString
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
func (m *BACnetContextTagBitString) GetDataType() BACnetDataType {
	return BACnetDataType_BIT_STRING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagBitString) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagBitString) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetContextTagBitString) GetPayload() *BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagBitString factory function for BACnetContextTagBitString
func NewBACnetContextTagBitString(payload *BACnetTagPayloadBitString, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagBitString{
		Payload:          payload,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagBitString(structType interface{}) *BACnetContextTagBitString {
	if casted, ok := structType.(BACnetContextTagBitString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagBitString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagBitString(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagBitString(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagBitString) GetTypeName() string {
	return "BACnetContextTagBitString"
}

func (m *BACnetContextTagBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetContextTagBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagBitStringParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, header *BACnetTagHeader) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagBitString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadBitString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagBitString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagBitString{
		Payload:          CastBACnetTagPayloadBitString(payload),
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBitString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBitString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
