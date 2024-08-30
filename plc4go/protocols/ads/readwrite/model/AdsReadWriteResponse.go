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

// AdsReadWriteResponse is the corresponding interface of AdsReadWriteResponse
type AdsReadWriteResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsReadWriteResponseExactly can be used when we want exactly this type and not a type which fulfills AdsReadWriteResponse.
// This is useful for switch cases.
type AdsReadWriteResponseExactly interface {
	AdsReadWriteResponse
	isAdsReadWriteResponse() bool
}

// _AdsReadWriteResponse is the data-structure of this message
type _AdsReadWriteResponse struct {
	*_AmsPacket
	Result ReturnCode
	Data   []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadWriteResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *_AdsReadWriteResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadWriteResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsReadWriteResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadWriteResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadWriteResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadWriteResponse factory function for _AdsReadWriteResponse
func NewAdsReadWriteResponse(result ReturnCode, data []byte, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsReadWriteResponse {
	_result := &_AdsReadWriteResponse{
		Result:     result,
		Data:       data,
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadWriteResponse(structType any) AdsReadWriteResponse {
	if casted, ok := structType.(AdsReadWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadWriteResponse) GetTypeName() string {
	return "AdsReadWriteResponse"
}

func (m *_AdsReadWriteResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadWriteResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsReadWriteResponseParse(ctx context.Context, theBytes []byte) (AdsReadWriteResponse, error) {
	return AdsReadWriteResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsReadWriteResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsReadWriteResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsReadWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParseWithBuffer(ctx, readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsReadWriteResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AdsReadWriteResponse")
	}

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsReadWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadWriteResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsReadWriteResponse{
		_AmsPacket: &_AmsPacket{},
		Result:     result,
		Data:       data,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadWriteResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadWriteResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadWriteResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(ctx, m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.GetData())))
		_lengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("length", 32, uint32((length)))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadWriteResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadWriteResponse) isAdsReadWriteResponse() bool {
	return true
}

func (m *_AdsReadWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
