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

// AdsDeleteDeviceNotificationRequest is the corresponding interface of AdsDeleteDeviceNotificationRequest
type AdsDeleteDeviceNotificationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
}

// AdsDeleteDeviceNotificationRequestExactly can be used when we want exactly this type and not a type which fulfills AdsDeleteDeviceNotificationRequest.
// This is useful for switch cases.
type AdsDeleteDeviceNotificationRequestExactly interface {
	AdsDeleteDeviceNotificationRequest
	isAdsDeleteDeviceNotificationRequest() bool
}

// _AdsDeleteDeviceNotificationRequest is the data-structure of this message
type _AdsDeleteDeviceNotificationRequest struct {
	*_AmsPacket
	NotificationHandle uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsDeleteDeviceNotificationRequest) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDeleteDeviceNotificationRequest factory function for _AdsDeleteDeviceNotificationRequest
func NewAdsDeleteDeviceNotificationRequest(notificationHandle uint32, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsDeleteDeviceNotificationRequest {
	_result := &_AdsDeleteDeviceNotificationRequest{
		NotificationHandle: notificationHandle,
		_AmsPacket:         NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationRequest(structType any) AdsDeleteDeviceNotificationRequest {
	if casted, ok := structType.(AdsDeleteDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationRequest) GetTypeName() string {
	return "AdsDeleteDeviceNotificationRequest"
}

func (m *_AdsDeleteDeviceNotificationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDeleteDeviceNotificationRequestParse(ctx context.Context, theBytes []byte) (AdsDeleteDeviceNotificationRequest, error) {
	return AdsDeleteDeviceNotificationRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDeleteDeviceNotificationRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDeleteDeviceNotificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := /*TODO: migrate me*/ readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field of AdsDeleteDeviceNotificationRequest")
	}
	notificationHandle := _notificationHandle

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsDeleteDeviceNotificationRequest{
		_AmsPacket:         &_AmsPacket{},
		NotificationHandle: notificationHandle,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsDeleteDeviceNotificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeleteDeviceNotificationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationRequest")
		}

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.GetNotificationHandle())
		_notificationHandleErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("notificationHandle", 32, uint32((notificationHandle)))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationRequest) isAdsDeleteDeviceNotificationRequest() bool {
	return true
}

func (m *_AdsDeleteDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
