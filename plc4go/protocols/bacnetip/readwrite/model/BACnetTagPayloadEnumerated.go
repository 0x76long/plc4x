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

// BACnetTagPayloadEnumerated is the corresponding interface of BACnetTagPayloadEnumerated
type BACnetTagPayloadEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetData returns Data (property field)
	GetData() []byte
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
}

// BACnetTagPayloadEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadEnumerated.
// This is useful for switch cases.
type BACnetTagPayloadEnumeratedExactly interface {
	BACnetTagPayloadEnumerated
	isBACnetTagPayloadEnumerated() bool
}

// _BACnetTagPayloadEnumerated is the data-structure of this message
type _BACnetTagPayloadEnumerated struct {
	Data []byte

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadEnumerated) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadEnumerated) GetActualValue() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(ParseVarUint(ctx, m.GetData()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadEnumerated factory function for _BACnetTagPayloadEnumerated
func NewBACnetTagPayloadEnumerated(data []byte, actualLength uint32) *_BACnetTagPayloadEnumerated {
	return &_BACnetTagPayloadEnumerated{Data: data, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadEnumerated(structType any) BACnetTagPayloadEnumerated {
	if casted, ok := structType.(BACnetTagPayloadEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadEnumerated) GetTypeName() string {
	return "BACnetTagPayloadEnumerated"
}

func (m *_BACnetTagPayloadEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadEnumeratedParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadEnumerated, error) {
	return BACnetTagPayloadEnumeratedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadEnumeratedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTagPayloadEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := readBuffer.ReadByteArray("data", int(actualLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	// Virtual field
	_actualValue := ParseVarUint(ctx, data)
	actualValue := uint32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadEnumerated")
	}

	// Create the instance
	return &_BACnetTagPayloadEnumerated{
		ActualLength: actualLength,
		Data:         data,
	}, nil
}

func (m *_BACnetTagPayloadEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadEnumerated"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadEnumerated")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ = actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadEnumerated"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadEnumerated")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadEnumerated) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadEnumerated) isBACnetTagPayloadEnumerated() bool {
	return true
}

func (m *_BACnetTagPayloadEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
