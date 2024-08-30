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

// Constant values.
const ModbusConstants_MODBUSTCPDEFAULTPORT uint16 = uint16(502)

// ModbusConstants is the corresponding interface of ModbusConstants
type ModbusConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ModbusConstantsExactly can be used when we want exactly this type and not a type which fulfills ModbusConstants.
// This is useful for switch cases.
type ModbusConstantsExactly interface {
	ModbusConstants
	isModbusConstants() bool
}

// _ModbusConstants is the data-structure of this message
type _ModbusConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusConstants) GetModbusTcpDefaultPort() uint16 {
	return ModbusConstants_MODBUSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusConstants factory function for _ModbusConstants
func NewModbusConstants() *_ModbusConstants {
	return &_ModbusConstants{}
}

// Deprecated: use the interface for direct cast
func CastModbusConstants(structType any) ModbusConstants {
	if casted, ok := structType.(ModbusConstants); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusConstants); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusConstants) GetTypeName() string {
	return "ModbusConstants"
}

func (m *_ModbusConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (modbusTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusConstantsParse(ctx context.Context, theBytes []byte) (ModbusConstants, error) {
	return ModbusConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	modbusTcpDefaultPort, err := readConstField[uint16](ctx, "modbusTcpDefaultPort", ReadUnsignedShort(readBuffer, 16), ModbusConstants.MODBUSTCPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modbusTcpDefaultPort' field"))
	}
	_ = modbusTcpDefaultPort

	if closeErr := readBuffer.CloseContext("ModbusConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusConstants")
	}

	// Create the instance
	return &_ModbusConstants{}, nil
}

func (m *_ModbusConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusConstants")
	}

	// Const Field (modbusTcpDefaultPort)
	_modbusTcpDefaultPortErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint16("modbusTcpDefaultPort", 16, uint16(502))
	if _modbusTcpDefaultPortErr != nil {
		return errors.Wrap(_modbusTcpDefaultPortErr, "Error serializing 'modbusTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("ModbusConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusConstants")
	}
	return nil
}

func (m *_ModbusConstants) isModbusConstants() bool {
	return true
}

func (m *_ModbusConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
