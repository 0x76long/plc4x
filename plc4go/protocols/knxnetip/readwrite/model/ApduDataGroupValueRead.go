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

// ApduDataGroupValueRead is the corresponding interface of ApduDataGroupValueRead
type ApduDataGroupValueRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataGroupValueReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataGroupValueRead.
// This is useful for switch cases.
type ApduDataGroupValueReadExactly interface {
	ApduDataGroupValueRead
	isApduDataGroupValueRead() bool
}

// _ApduDataGroupValueRead is the data-structure of this message
type _ApduDataGroupValueRead struct {
	*_ApduData
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataGroupValueRead) GetApciType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataGroupValueRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataGroupValueRead) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataGroupValueRead factory function for _ApduDataGroupValueRead
func NewApduDataGroupValueRead(dataLength uint8) *_ApduDataGroupValueRead {
	_result := &_ApduDataGroupValueRead{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataGroupValueRead(structType any) ApduDataGroupValueRead {
	if casted, ok := structType.(ApduDataGroupValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataGroupValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataGroupValueRead) GetTypeName() string {
	return "ApduDataGroupValueRead"
}

func (m *_ApduDataGroupValueRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 6

	return lengthInBits
}

func (m *_ApduDataGroupValueRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataGroupValueReadParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduDataGroupValueRead, error) {
	return ApduDataGroupValueReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataGroupValueReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataGroupValueRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataGroupValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataGroupValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 6)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ApduDataGroupValueRead")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataGroupValueRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataGroupValueRead{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		reservedField0: reservedField0,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataGroupValueRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataGroupValueRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataGroupValueRead")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 6, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataGroupValueRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataGroupValueRead) isApduDataGroupValueRead() bool {
	return true
}

func (m *_ApduDataGroupValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
