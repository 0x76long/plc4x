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

// BACnetUnconfirmedServiceRequestIHave is the corresponding interface of BACnetUnconfirmedServiceRequestIHave
type BACnetUnconfirmedServiceRequestIHave interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetApplicationTagCharacterString
}

// BACnetUnconfirmedServiceRequestIHaveExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestIHave.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestIHaveExactly interface {
	BACnetUnconfirmedServiceRequestIHave
	isBACnetUnconfirmedServiceRequestIHave() bool
}

// _BACnetUnconfirmedServiceRequestIHave is the data-structure of this message
type _BACnetUnconfirmedServiceRequestIHave struct {
	*_BACnetUnconfirmedServiceRequest
	DeviceIdentifier BACnetApplicationTagObjectIdentifier
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
	ObjectName       BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_I_HAVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIHave) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetObjectName() BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestIHave factory function for _BACnetUnconfirmedServiceRequestIHave
func NewBACnetUnconfirmedServiceRequestIHave(deviceIdentifier BACnetApplicationTagObjectIdentifier, objectIdentifier BACnetApplicationTagObjectIdentifier, objectName BACnetApplicationTagCharacterString, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestIHave {
	_result := &_BACnetUnconfirmedServiceRequestIHave{
		DeviceIdentifier:                 deviceIdentifier,
		ObjectIdentifier:                 objectIdentifier,
		ObjectName:                       objectName,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestIHave(structType any) BACnetUnconfirmedServiceRequestIHave {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestIHave); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestIHave); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIHave"
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestIHave) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestIHaveParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestIHave, error) {
	return BACnetUnconfirmedServiceRequestIHaveParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestIHaveParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestIHave, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIHave"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestIHave")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field of BACnetUnconfirmedServiceRequestIHave")
	}
	deviceIdentifier := _deviceIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetUnconfirmedServiceRequestIHave")
	}
	objectIdentifier := _objectIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (objectName)
	if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectName")
	}
	_objectName, _objectNameErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _objectNameErr != nil {
		return nil, errors.Wrap(_objectNameErr, "Error parsing 'objectName' field of BACnetUnconfirmedServiceRequestIHave")
	}
	objectName := _objectName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectName")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIHave"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestIHave")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestIHave{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceIdentifier: deviceIdentifier,
		ObjectIdentifier: objectIdentifier,
		ObjectName:       objectName,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestIHave) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIHave"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestIHave")
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		_deviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetDeviceIdentifier())
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (objectName)
		if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectName")
		}
		_objectNameErr := writeBuffer.WriteSerializable(ctx, m.GetObjectName())
		if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectName")
		}
		if _objectNameErr != nil {
			return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIHave"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestIHave")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestIHave) isBACnetUnconfirmedServiceRequestIHave() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestIHave) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
