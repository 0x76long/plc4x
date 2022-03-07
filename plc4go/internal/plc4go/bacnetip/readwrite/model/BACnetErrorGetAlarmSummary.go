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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetErrorGetAlarmSummary struct {
	*BACnetError
}

// The corresponding interface
type IBACnetErrorGetAlarmSummary interface {
	IBACnetError
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
func (m *BACnetErrorGetAlarmSummary) GetServiceChoice() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetErrorGetAlarmSummary) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.BACnetError.ErrorClass = errorClass
	m.BACnetError.ErrorCode = errorCode
}

func (m *BACnetErrorGetAlarmSummary) GetParent() *BACnetError {
	return m.BACnetError
}

// NewBACnetErrorGetAlarmSummary factory function for BACnetErrorGetAlarmSummary
func NewBACnetErrorGetAlarmSummary(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetError {
	child := &BACnetErrorGetAlarmSummary{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	child.Child = child
	return child.BACnetError
}

func CastBACnetErrorGetAlarmSummary(structType interface{}) *BACnetErrorGetAlarmSummary {
	if casted, ok := structType.(BACnetErrorGetAlarmSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetErrorGetAlarmSummary); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastBACnetErrorGetAlarmSummary(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastBACnetErrorGetAlarmSummary(casted.Child)
	}
	return nil
}

func (m *BACnetErrorGetAlarmSummary) GetTypeName() string {
	return "BACnetErrorGetAlarmSummary"
}

func (m *BACnetErrorGetAlarmSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetErrorGetAlarmSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorGetAlarmSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorGetAlarmSummaryParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorGetAlarmSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetErrorGetAlarmSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorGetAlarmSummary{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child.BACnetError, nil
}

func (m *BACnetErrorGetAlarmSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorGetAlarmSummary"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorGetAlarmSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorGetAlarmSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
