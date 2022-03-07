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
type BACnetUnconfirmedServiceRequestUnconfirmedEventNotification struct {
	*BACnetUnconfirmedServiceRequest

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedEventNotification interface {
	IBACnetUnconfirmedServiceRequest
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
func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetServiceChoice() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

// NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification factory function for BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
func NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(len uint16) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
