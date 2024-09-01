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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CipConnectionManagerResponse is the corresponding interface of CipConnectionManagerResponse
type CipConnectionManagerResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetOtConnectionId returns OtConnectionId (property field)
	GetOtConnectionId() uint32
	// GetToConnectionId returns ToConnectionId (property field)
	GetToConnectionId() uint32
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetOtApi returns OtApi (property field)
	GetOtApi() uint32
	// GetToApi returns ToApi (property field)
	GetToApi() uint32
}

// CipConnectionManagerResponseExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerResponse.
// This is useful for switch cases.
type CipConnectionManagerResponseExactly interface {
	CipConnectionManagerResponse
	isCipConnectionManagerResponse() bool
}

// _CipConnectionManagerResponse is the data-structure of this message
type _CipConnectionManagerResponse struct {
	*_CipService
	OtConnectionId         uint32
	ToConnectionId         uint32
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	OtApi                  uint32
	ToApi                  uint32
	// Reserved Fields
	reservedField0 *uint32
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerResponse) GetService() uint8 {
	return 0x5B
}

func (m *_CipConnectionManagerResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectionManagerResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerResponse) InitializeParent(parent CipService) {}

func (m *_CipConnectionManagerResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerResponse) GetOtConnectionId() uint32 {
	return m.OtConnectionId
}

func (m *_CipConnectionManagerResponse) GetToConnectionId() uint32 {
	return m.ToConnectionId
}

func (m *_CipConnectionManagerResponse) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerResponse) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOtApi() uint32 {
	return m.OtApi
}

func (m *_CipConnectionManagerResponse) GetToApi() uint32 {
	return m.ToApi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerResponse factory function for _CipConnectionManagerResponse
func NewCipConnectionManagerResponse(otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, otApi uint32, toApi uint32, serviceLen uint16) *_CipConnectionManagerResponse {
	_result := &_CipConnectionManagerResponse{
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		OtApi:                  otApi,
		ToApi:                  toApi,
		_CipService:            NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerResponse(structType any) CipConnectionManagerResponse {
	if casted, ok := structType.(CipConnectionManagerResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerResponse) GetTypeName() string {
	return "CipConnectionManagerResponse"
}

func (m *_CipConnectionManagerResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 24

	// Simple field (otConnectionId)
	lengthInBits += 32

	// Simple field (toConnectionId)
	lengthInBits += 32

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (otApi)
	lengthInBits += 32

	// Simple field (toApi)
	lengthInBits += 32

	// Implicit Field (replySize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipConnectionManagerResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectionManagerResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipConnectionManagerResponse, error) {
	return CipConnectionManagerResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectionManagerResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CipConnectionManagerResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CipConnectionManagerResponse, error) {
		return CipConnectionManagerResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func CipConnectionManagerResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectionManagerResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(24)), uint32(0x000000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	otConnectionId, err := ReadSimpleField(ctx, "otConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otConnectionId' field"))
	}

	toConnectionId, err := ReadSimpleField(ctx, "toConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toConnectionId' field"))
	}

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}

	otApi, err := ReadSimpleField(ctx, "otApi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otApi' field"))
	}

	toApi, err := ReadSimpleField(ctx, "toApi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toApi' field"))
	}

	replySize, err := ReadImplicitField[uint8](ctx, "replySize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'replySize' field"))
	}
	_ = replySize

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	if closeErr := readBuffer.CloseContext("CipConnectionManagerResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerResponse")
	}

	// Create a partially initialized instance
	_child := &_CipConnectionManagerResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		OtApi:                  otApi,
		ToApi:                  toApi,
		reservedField0:         reservedField0,
		reservedField1:         reservedField1,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectionManagerResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerResponse")
		}

		if err := WriteReservedField[uint32](ctx, "reserved", uint32(0x000000), WriteUnsignedInt(writeBuffer, 24)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		// Simple Field (otConnectionId)
		otConnectionId := uint32(m.GetOtConnectionId())
		_otConnectionIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("otConnectionId", 32, uint32((otConnectionId)))
		if _otConnectionIdErr != nil {
			return errors.Wrap(_otConnectionIdErr, "Error serializing 'otConnectionId' field")
		}

		// Simple Field (toConnectionId)
		toConnectionId := uint32(m.GetToConnectionId())
		_toConnectionIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("toConnectionId", 32, uint32((toConnectionId)))
		if _toConnectionIdErr != nil {
			return errors.Wrap(_toConnectionIdErr, "Error serializing 'toConnectionId' field")
		}

		// Simple Field (connectionSerialNumber)
		connectionSerialNumber := uint16(m.GetConnectionSerialNumber())
		_connectionSerialNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("connectionSerialNumber", 16, uint16((connectionSerialNumber)))
		if _connectionSerialNumberErr != nil {
			return errors.Wrap(_connectionSerialNumberErr, "Error serializing 'connectionSerialNumber' field")
		}

		// Simple Field (originatorVendorId)
		originatorVendorId := uint16(m.GetOriginatorVendorId())
		_originatorVendorIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("originatorVendorId", 16, uint16((originatorVendorId)))
		if _originatorVendorIdErr != nil {
			return errors.Wrap(_originatorVendorIdErr, "Error serializing 'originatorVendorId' field")
		}

		// Simple Field (originatorSerialNumber)
		originatorSerialNumber := uint32(m.GetOriginatorSerialNumber())
		_originatorSerialNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("originatorSerialNumber", 32, uint32((originatorSerialNumber)))
		if _originatorSerialNumberErr != nil {
			return errors.Wrap(_originatorSerialNumberErr, "Error serializing 'originatorSerialNumber' field")
		}

		// Simple Field (otApi)
		otApi := uint32(m.GetOtApi())
		_otApiErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("otApi", 32, uint32((otApi)))
		if _otApiErr != nil {
			return errors.Wrap(_otApiErr, "Error serializing 'otApi' field")
		}

		// Simple Field (toApi)
		toApi := uint32(m.GetToApi())
		_toApiErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("toApi", 32, uint32((toApi)))
		if _toApiErr != nil {
			return errors.Wrap(_toApiErr, "Error serializing 'toApi' field")
		}
		replySize := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(30)))
		if err := WriteImplicitField(ctx, "replySize", replySize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'replySize' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerResponse) isCipConnectionManagerResponse() bool {
	return true
}

func (m *_CipConnectionManagerResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
