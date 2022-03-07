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
type BACnetUnconfirmedServiceRequestWhoIs struct {
	*BACnetUnconfirmedServiceRequest
	DeviceInstanceRangeLowLimit  *BACnetContextTagUnsignedInteger
	DeviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWhoIs interface {
	IBACnetUnconfirmedServiceRequest
	// GetDeviceInstanceRangeLowLimit returns DeviceInstanceRangeLowLimit (property field)
	GetDeviceInstanceRangeLowLimit() *BACnetContextTagUnsignedInteger
	// GetDeviceInstanceRangeHighLimit returns DeviceInstanceRangeHighLimit (property field)
	GetDeviceInstanceRangeHighLimit() *BACnetContextTagUnsignedInteger
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
func (m *BACnetUnconfirmedServiceRequestWhoIs) GetServiceChoice() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestWhoIs) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetUnconfirmedServiceRequestWhoIs) GetDeviceInstanceRangeLowLimit() *BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeLowLimit
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetDeviceInstanceRangeHighLimit() *BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoIs factory function for BACnetUnconfirmedServiceRequestWhoIs
func NewBACnetUnconfirmedServiceRequestWhoIs(deviceInstanceRangeLowLimit *BACnetContextTagUnsignedInteger, deviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger, len uint16) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWhoIs{
		DeviceInstanceRangeLowLimit:     deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit:    deviceInstanceRangeHighLimit,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestWhoIs(structType interface{}) *BACnetUnconfirmedServiceRequestWhoIs {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoIs); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoIs); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWhoIs(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWhoIs(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoIs"
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (deviceInstanceRangeLowLimit)
	if m.DeviceInstanceRangeLowLimit != nil {
		lengthInBits += (*m.DeviceInstanceRangeLowLimit).GetLengthInBits()
	}

	// Optional Field (deviceInstanceRangeHighLimit)
	if m.DeviceInstanceRangeHighLimit != nil {
		lengthInBits += (*m.DeviceInstanceRangeHighLimit).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoIsParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoIs"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (deviceInstanceRangeLowLimit) (Can be skipped, if a given expression evaluates to false)
	var deviceInstanceRangeLowLimit *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("deviceInstanceRangeLowLimit"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeLowLimit' field")
		default:
			deviceInstanceRangeLowLimit = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("deviceInstanceRangeLowLimit"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (deviceInstanceRangeHighLimit) (Can be skipped, if a given expression evaluates to false)
	var deviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger = nil
	if bool((deviceInstanceRangeLowLimit) != (nil)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("deviceInstanceRangeHighLimit"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeHighLimit' field")
		default:
			deviceInstanceRangeHighLimit = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("deviceInstanceRangeHighLimit"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoIs"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWhoIs{
		DeviceInstanceRangeLowLimit:     CastBACnetContextTagUnsignedInteger(deviceInstanceRangeLowLimit),
		DeviceInstanceRangeHighLimit:    CastBACnetContextTagUnsignedInteger(deviceInstanceRangeHighLimit),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoIs"); pushErr != nil {
			return pushErr
		}

		// Optional Field (deviceInstanceRangeLowLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeLowLimit *BACnetContextTagUnsignedInteger = nil
		if m.DeviceInstanceRangeLowLimit != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeLowLimit"); pushErr != nil {
				return pushErr
			}
			deviceInstanceRangeLowLimit = m.DeviceInstanceRangeLowLimit
			_deviceInstanceRangeLowLimitErr := deviceInstanceRangeLowLimit.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeLowLimit"); popErr != nil {
				return popErr
			}
			if _deviceInstanceRangeLowLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeLowLimitErr, "Error serializing 'deviceInstanceRangeLowLimit' field")
			}
		}

		// Optional Field (deviceInstanceRangeHighLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger = nil
		if m.DeviceInstanceRangeHighLimit != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeHighLimit"); pushErr != nil {
				return pushErr
			}
			deviceInstanceRangeHighLimit = m.DeviceInstanceRangeHighLimit
			_deviceInstanceRangeHighLimitErr := deviceInstanceRangeHighLimit.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeHighLimit"); popErr != nil {
				return popErr
			}
			if _deviceInstanceRangeHighLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeHighLimitErr, "Error serializing 'deviceInstanceRangeHighLimit' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoIs"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWhoIs) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
