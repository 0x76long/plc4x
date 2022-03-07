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
type BACnetConfirmedServiceACKReadRange struct {
	*BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKReadRange interface {
	IBACnetConfirmedServiceACK
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
func (m *BACnetConfirmedServiceACKReadRange) GetServiceChoice() uint8 {
	return 0x1A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceACKReadRange) InitializeParent(parent *BACnetConfirmedServiceACK) {}

func (m *BACnetConfirmedServiceACKReadRange) GetParent() *BACnetConfirmedServiceACK {
	return m.BACnetConfirmedServiceACK
}

// NewBACnetConfirmedServiceACKReadRange factory function for BACnetConfirmedServiceACKReadRange
func NewBACnetConfirmedServiceACKReadRange() *BACnetConfirmedServiceACK {
	child := &BACnetConfirmedServiceACKReadRange{
		BACnetConfirmedServiceACK: NewBACnetConfirmedServiceACK(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceACK
}

func CastBACnetConfirmedServiceACKReadRange(structType interface{}) *BACnetConfirmedServiceACKReadRange {
	if casted, ok := structType.(BACnetConfirmedServiceACKReadRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceACKReadRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceACK); ok {
		return CastBACnetConfirmedServiceACKReadRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceACK); ok {
		return CastBACnetConfirmedServiceACKReadRange(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceACKReadRange) GetTypeName() string {
	return "BACnetConfirmedServiceACKReadRange"
}

func (m *BACnetConfirmedServiceACKReadRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACKReadRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKReadRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceACKReadRangeParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACKReadRange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACKReadRange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKReadRange{
		BACnetConfirmedServiceACK: &BACnetConfirmedServiceACK{},
	}
	_child.BACnetConfirmedServiceACK.Child = _child
	return _child.BACnetConfirmedServiceACK, nil
}

func (m *BACnetConfirmedServiceACKReadRange) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACKReadRange"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACKReadRange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceACKReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
