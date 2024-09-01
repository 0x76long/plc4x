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

// InterfaceOptions2 is the corresponding interface of InterfaceOptions2
type InterfaceOptions2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetBurden returns Burden (property field)
	GetBurden() bool
	// GetClockGen returns ClockGen (property field)
	GetClockGen() bool
}

// InterfaceOptions2Exactly can be used when we want exactly this type and not a type which fulfills InterfaceOptions2.
// This is useful for switch cases.
type InterfaceOptions2Exactly interface {
	InterfaceOptions2
	isInterfaceOptions2() bool
}

// _InterfaceOptions2 is the data-structure of this message
type _InterfaceOptions2 struct {
	Burden   bool
	ClockGen bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
	reservedField4 *bool
	reservedField5 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions2) GetBurden() bool {
	return m.Burden
}

func (m *_InterfaceOptions2) GetClockGen() bool {
	return m.ClockGen
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInterfaceOptions2 factory function for _InterfaceOptions2
func NewInterfaceOptions2(burden bool, clockGen bool) *_InterfaceOptions2 {
	return &_InterfaceOptions2{Burden: burden, ClockGen: clockGen}
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions2(structType any) InterfaceOptions2 {
	if casted, ok := structType.(InterfaceOptions2); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions2); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions2) GetTypeName() string {
	return "InterfaceOptions2"
}

func (m *_InterfaceOptions2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (burden)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (clockGen)
	lengthInBits += 1

	return lengthInBits
}

func (m *_InterfaceOptions2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceOptions2Parse(ctx context.Context, theBytes []byte) (InterfaceOptions2, error) {
	return InterfaceOptions2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions2ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions2, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions2, error) {
		return InterfaceOptions2ParseWithBuffer(ctx, readBuffer)
	}
}

func InterfaceOptions2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	burden, err := ReadSimpleField(ctx, "burden", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'burden' field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField3, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField4, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	reservedField5, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	clockGen, err := ReadSimpleField(ctx, "clockGen", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clockGen' field"))
	}

	if closeErr := readBuffer.CloseContext("InterfaceOptions2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions2")
	}

	// Create the instance
	return &_InterfaceOptions2{
		Burden:         burden,
		ClockGen:       clockGen,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
		reservedField2: reservedField2,
		reservedField3: reservedField3,
		reservedField4: reservedField4,
		reservedField5: reservedField5,
	}, nil
}

func (m *_InterfaceOptions2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("InterfaceOptions2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	// Simple Field (burden)
	burden := bool(m.GetBurden())
	_burdenErr := /*TODO: migrate me*/ writeBuffer.WriteBit("burden", (burden))
	if _burdenErr != nil {
		return errors.Wrap(_burdenErr, "Error serializing 'burden' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 4")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 5")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 6")
	}

	// Simple Field (clockGen)
	clockGen := bool(m.GetClockGen())
	_clockGenErr := /*TODO: migrate me*/ writeBuffer.WriteBit("clockGen", (clockGen))
	if _clockGenErr != nil {
		return errors.Wrap(_clockGenErr, "Error serializing 'clockGen' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions2")
	}
	return nil
}

func (m *_InterfaceOptions2) isInterfaceOptions2() bool {
	return true
}

func (m *_InterfaceOptions2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
