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
type BACnetUnconfirmedServiceRequestWhoHas struct {
	*BACnetUnconfirmedServiceRequest
	DeviceInstanceRangeLowLimit  *BACnetContextTagUnsignedInteger
	DeviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger
	ObjectIdentifier             *BACnetContextTagObjectIdentifier
	ObjectName                   *BACnetContextTagOctetString

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWhoHas interface {
	IBACnetUnconfirmedServiceRequest
	// GetDeviceInstanceRangeLowLimit returns DeviceInstanceRangeLowLimit (property field)
	GetDeviceInstanceRangeLowLimit() *BACnetContextTagUnsignedInteger
	// GetDeviceInstanceRangeHighLimit returns DeviceInstanceRangeHighLimit (property field)
	GetDeviceInstanceRangeHighLimit() *BACnetContextTagUnsignedInteger
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetObjectName returns ObjectName (property field)
	GetObjectName() *BACnetContextTagOctetString
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
func (m *BACnetUnconfirmedServiceRequestWhoHas) GetServiceChoice() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestWhoHas) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeLowLimit() *BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeLowLimit
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeHighLimit() *BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeHighLimit
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetObjectName() *BACnetContextTagOctetString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHas factory function for BACnetUnconfirmedServiceRequestWhoHas
func NewBACnetUnconfirmedServiceRequestWhoHas(deviceInstanceRangeLowLimit *BACnetContextTagUnsignedInteger, deviceInstanceRangeHighLimit *BACnetContextTagUnsignedInteger, objectIdentifier *BACnetContextTagObjectIdentifier, objectName *BACnetContextTagOctetString, len uint16) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceRangeLowLimit:     deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit:    deviceInstanceRangeHighLimit,
		ObjectIdentifier:                objectIdentifier,
		ObjectName:                      objectName,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestWhoHas(structType interface{}) *BACnetUnconfirmedServiceRequestWhoHas {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHas); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHas); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHas"
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (deviceInstanceRangeLowLimit)
	if m.DeviceInstanceRangeLowLimit != nil {
		lengthInBits += (*m.DeviceInstanceRangeLowLimit).GetLengthInBits()
	}

	// Optional Field (deviceInstanceRangeHighLimit)
	if m.DeviceInstanceRangeHighLimit != nil {
		lengthInBits += (*m.DeviceInstanceRangeHighLimit).GetLengthInBits()
	}

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += (*m.ObjectIdentifier).GetLengthInBits()
	}

	// Optional Field (objectName)
	if m.ObjectName != nil {
		lengthInBits += (*m.ObjectName).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoHasParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHas"); pullErr != nil {
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

	// Optional Field (objectIdentifier) (Can be skipped, if a given expression evaluates to false)
	var objectIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'objectIdentifier' field")
		default:
			objectIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (objectName) (Can be skipped, if a given expression evaluates to false)
	var objectName *BACnetContextTagOctetString = nil
	if bool((objectIdentifier) == (nil)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_OCTET_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'objectName' field")
		default:
			objectName = CastBACnetContextTagOctetString(_val)
			if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHas"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceRangeLowLimit:     CastBACnetContextTagUnsignedInteger(deviceInstanceRangeLowLimit),
		DeviceInstanceRangeHighLimit:    CastBACnetContextTagUnsignedInteger(deviceInstanceRangeHighLimit),
		ObjectIdentifier:                CastBACnetContextTagObjectIdentifier(objectIdentifier),
		ObjectName:                      CastBACnetContextTagOctetString(objectName),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHas"); pushErr != nil {
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

		// Optional Field (objectIdentifier) (Can be skipped, if the value is null)
		var objectIdentifier *BACnetContextTagObjectIdentifier = nil
		if m.ObjectIdentifier != nil {
			if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
				return pushErr
			}
			objectIdentifier = m.ObjectIdentifier
			_objectIdentifierErr := objectIdentifier.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
				return popErr
			}
			if _objectIdentifierErr != nil {
				return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
			}
		}

		// Optional Field (objectName) (Can be skipped, if the value is null)
		var objectName *BACnetContextTagOctetString = nil
		if m.ObjectName != nil {
			if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
				return pushErr
			}
			objectName = m.ObjectName
			_objectNameErr := objectName.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
				return popErr
			}
			if _objectNameErr != nil {
				return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHas"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
