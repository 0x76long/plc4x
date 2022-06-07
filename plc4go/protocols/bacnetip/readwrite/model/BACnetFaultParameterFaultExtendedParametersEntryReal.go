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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetFaultParameterFaultExtendedParametersEntryReal is the data-structure of this message
type BACnetFaultParameterFaultExtendedParametersEntryReal struct {
	*BACnetFaultParameterFaultExtendedParametersEntry
	RealValue *BACnetApplicationTagReal
}

// IBACnetFaultParameterFaultExtendedParametersEntryReal is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryReal
type IBACnetFaultParameterFaultExtendedParametersEntryReal interface {
	IBACnetFaultParameterFaultExtendedParametersEntry
	// GetRealValue returns RealValue (property field)
	GetRealValue() *BACnetApplicationTagReal
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

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) InitializeParent(parent *BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameterFaultExtendedParametersEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetParent() *BACnetFaultParameterFaultExtendedParametersEntry {
	return m.BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetRealValue() *BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryReal factory function for BACnetFaultParameterFaultExtendedParametersEntryReal
func NewBACnetFaultParameterFaultExtendedParametersEntryReal(realValue *BACnetApplicationTagReal, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultExtendedParametersEntryReal {
	_result := &BACnetFaultParameterFaultExtendedParametersEntryReal{
		RealValue: realValue,
		BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultExtendedParametersEntryReal(structType interface{}) *BACnetFaultParameterFaultExtendedParametersEntryReal {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryReal(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryReal(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryReal"
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryRealParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultExtendedParametersEntryReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, pullErr
	}
	_realValue, _realValueErr := BACnetApplicationTagParse(readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field")
	}
	realValue := CastBACnetApplicationTagReal(_realValue)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultExtendedParametersEntryReal{
		RealValue: CastBACnetApplicationTagReal(realValue),
		BACnetFaultParameterFaultExtendedParametersEntry: &BACnetFaultParameterFaultExtendedParametersEntry{},
	}
	_child.BACnetFaultParameterFaultExtendedParametersEntry.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return pushErr
		}
		_realValueErr := m.RealValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return popErr
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryReal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
