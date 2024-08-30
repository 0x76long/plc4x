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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// AssociatedQueryValueType is the corresponding interface of AssociatedQueryValueType
type AssociatedQueryValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetValueLength returns ValueLength (property field)
	GetValueLength() uint16
	// GetData returns Data (property field)
	GetData() []uint8
}

// AssociatedQueryValueTypeExactly can be used when we want exactly this type and not a type which fulfills AssociatedQueryValueType.
// This is useful for switch cases.
type AssociatedQueryValueTypeExactly interface {
	AssociatedQueryValueType
	isAssociatedQueryValueType() bool
}

// _AssociatedQueryValueType is the data-structure of this message
type _AssociatedQueryValueType struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	ValueLength   uint16
	Data          []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AssociatedQueryValueType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_AssociatedQueryValueType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_AssociatedQueryValueType) GetValueLength() uint16 {
	return m.ValueLength
}

func (m *_AssociatedQueryValueType) GetData() []uint8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAssociatedQueryValueType factory function for _AssociatedQueryValueType
func NewAssociatedQueryValueType(returnCode DataTransportErrorCode, transportSize DataTransportSize, valueLength uint16, data []uint8) *_AssociatedQueryValueType {
	return &_AssociatedQueryValueType{ReturnCode: returnCode, TransportSize: transportSize, ValueLength: valueLength, Data: data}
}

// Deprecated: use the interface for direct cast
func CastAssociatedQueryValueType(structType any) AssociatedQueryValueType {
	if casted, ok := structType.(AssociatedQueryValueType); ok {
		return casted
	}
	if casted, ok := structType.(*AssociatedQueryValueType); ok {
		return *casted
	}
	return nil
}

func (m *_AssociatedQueryValueType) GetTypeName() string {
	return "AssociatedQueryValueType"
}

func (m *_AssociatedQueryValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Simple field (valueLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AssociatedQueryValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AssociatedQueryValueTypeParse(ctx context.Context, theBytes []byte) (AssociatedQueryValueType, error) {
	return AssociatedQueryValueTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AssociatedQueryValueTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AssociatedQueryValueType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AssociatedQueryValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AssociatedQueryValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParseWithBuffer(ctx, readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of AssociatedQueryValueType")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	_transportSize, _transportSizeErr := DataTransportSizeParseWithBuffer(ctx, readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of AssociatedQueryValueType")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Simple Field (valueLength)
	_valueLength, _valueLengthErr := /*TODO: migrate me*/ readBuffer.ReadUint16("valueLength", 16)
	if _valueLengthErr != nil {
		return nil, errors.Wrap(_valueLengthErr, "Error parsing 'valueLength' field of AssociatedQueryValueType")
	}
	valueLength := _valueLength

	data, err := ReadCountArrayField[uint8](ctx, "data", ReadUnsignedByte(readBuffer, 8), uint64(valueLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("AssociatedQueryValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AssociatedQueryValueType")
	}

	// Create the instance
	return &_AssociatedQueryValueType{
		ReturnCode:    returnCode,
		TransportSize: transportSize,
		ValueLength:   valueLength,
		Data:          data,
	}, nil
}

func (m *_AssociatedQueryValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AssociatedQueryValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AssociatedQueryValueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AssociatedQueryValueType")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Simple Field (valueLength)
	valueLength := uint16(m.GetValueLength())
	_valueLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("valueLength", 16, uint16((valueLength)))
	if _valueLengthErr != nil {
		return errors.Wrap(_valueLengthErr, "Error serializing 'valueLength' field")
	}

	// Array Field (data)
	if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	for _curItem, _element := range m.GetData() {
		_ = _curItem
		_elementErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("", 8, uint8(_element))
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'data' field")
		}
	}
	if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}

	if popErr := writeBuffer.PopContext("AssociatedQueryValueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AssociatedQueryValueType")
	}
	return nil
}

func (m *_AssociatedQueryValueType) isAssociatedQueryValueType() bool {
	return true
}

func (m *_AssociatedQueryValueType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
