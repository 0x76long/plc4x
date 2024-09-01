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

// CALReplyLong is the corresponding interface of CALReplyLong
type CALReplyLong interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALReply
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetSerialInterfaceAddress returns SerialInterfaceAddress (property field)
	GetSerialInterfaceAddress() SerialInterfaceAddress
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() ReplyNetwork
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
}

// CALReplyLongExactly can be used when we want exactly this type and not a type which fulfills CALReplyLong.
// This is useful for switch cases.
type CALReplyLongExactly interface {
	CALReplyLong
	isCALReplyLong() bool
}

// _CALReplyLong is the data-structure of this message
type _CALReplyLong struct {
	*_CALReply
	TerminatingByte        uint32
	UnitAddress            UnitAddress
	BridgeAddress          BridgeAddress
	SerialInterfaceAddress SerialInterfaceAddress
	ReservedByte           *byte
	ReplyNetwork           ReplyNetwork
	// Reserved Fields
	reservedField0 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALReplyLong) InitializeParent(parent CALReply, calType byte, calData CALData) {
	m.CalType = calType
	m.CalData = calData
}

func (m *_CALReplyLong) GetParent() CALReply {
	return m._CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALReplyLong) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *_CALReplyLong) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

func (m *_CALReplyLong) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CALReplyLong) GetSerialInterfaceAddress() SerialInterfaceAddress {
	return m.SerialInterfaceAddress
}

func (m *_CALReplyLong) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *_CALReplyLong) GetReplyNetwork() ReplyNetwork {
	return m.ReplyNetwork
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALReplyLong) GetIsUnitAddress() bool {
	ctx := context.Background()
	_ = ctx
	unitAddress := m.UnitAddress
	_ = unitAddress
	bridgeAddress := m.BridgeAddress
	_ = bridgeAddress
	reservedByte := m.ReservedByte
	_ = reservedByte
	replyNetwork := m.ReplyNetwork
	_ = replyNetwork
	return bool(bool((m.GetTerminatingByte() & 0xff) == (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReplyLong factory function for _CALReplyLong
func NewCALReplyLong(terminatingByte uint32, unitAddress UnitAddress, bridgeAddress BridgeAddress, serialInterfaceAddress SerialInterfaceAddress, reservedByte *byte, replyNetwork ReplyNetwork, calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReplyLong {
	_result := &_CALReplyLong{
		TerminatingByte:        terminatingByte,
		UnitAddress:            unitAddress,
		BridgeAddress:          bridgeAddress,
		SerialInterfaceAddress: serialInterfaceAddress,
		ReservedByte:           reservedByte,
		ReplyNetwork:           replyNetwork,
		_CALReply:              NewCALReply(calType, calData, cBusOptions, requestContext),
	}
	_result._CALReply._CALReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALReplyLong(structType any) CALReplyLong {
	if casted, ok := structType.(CALReplyLong); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyLong); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyLong) GetTypeName() string {
	return "CALReplyLong"
}

func (m *_CALReplyLong) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += m.UnitAddress.GetLengthInBits(ctx)
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)
	}

	// Simple field (serialInterfaceAddress)
	lengthInBits += m.SerialInterfaceAddress.GetLengthInBits(ctx)

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += m.ReplyNetwork.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CALReplyLong) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALReplyLongParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (CALReplyLong, error) {
	return CALReplyLongParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func CALReplyLongParseWithBufferProducer(cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (CALReplyLong, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CALReplyLong, error) {
		return CALReplyLongParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	}
}

func CALReplyLongParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (CALReplyLong, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyLong"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyLong")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x86))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	terminatingByte, err := ReadPeekField[uint32](ctx, "terminatingByte", ReadUnsignedInt(readBuffer, uint8(24)), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminatingByte' field"))
	}

	isUnitAddress, err := ReadVirtualField[bool](ctx, "isUnitAddress", (*bool)(nil), bool((terminatingByte&0xff) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnitAddress' field"))
	}
	_ = isUnitAddress

	_unitAddress, err := ReadOptionalField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	var unitAddress UnitAddress
	if _unitAddress != nil {
		unitAddress = *_unitAddress
	}

	_bridgeAddress, err := ReadOptionalField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	var bridgeAddress BridgeAddress
	if _bridgeAddress != nil {
		bridgeAddress = *_bridgeAddress
	}

	serialInterfaceAddress, err := ReadSimpleField[SerialInterfaceAddress](ctx, "serialInterfaceAddress", ReadComplex[SerialInterfaceAddress](SerialInterfaceAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serialInterfaceAddress' field"))
	}

	reservedByte, err := ReadOptionalField[byte](ctx, "reservedByte", ReadByte(readBuffer, 8), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reservedByte' field"))
	}

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "wrong reservedByte"})
	}

	_replyNetwork, err := ReadOptionalField[ReplyNetwork](ctx, "replyNetwork", ReadComplex[ReplyNetwork](ReplyNetworkParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'replyNetwork' field"))
	}
	var replyNetwork ReplyNetwork
	if _replyNetwork != nil {
		replyNetwork = *_replyNetwork
	}

	if closeErr := readBuffer.CloseContext("CALReplyLong"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyLong")
	}

	// Create a partially initialized instance
	_child := &_CALReplyLong{
		_CALReply: &_CALReply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		TerminatingByte:        terminatingByte,
		UnitAddress:            unitAddress,
		BridgeAddress:          bridgeAddress,
		SerialInterfaceAddress: serialInterfaceAddress,
		ReservedByte:           reservedByte,
		ReplyNetwork:           replyNetwork,
		reservedField0:         reservedField0,
	}
	_child._CALReply._CALReplyChildRequirements = _child
	return _child, nil
}

func (m *_CALReplyLong) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALReplyLong) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyLong"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyLong")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x86), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}
		// Virtual field
		isUnitAddress := m.GetIsUnitAddress()
		_ = isUnitAddress
		if _isUnitAddressErr := writeBuffer.WriteVirtual(ctx, "isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		if err := WriteOptionalField[UnitAddress](ctx, "unitAddress", GetRef(m.GetUnitAddress()), WriteComplex[UnitAddress](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if err := WriteOptionalField[BridgeAddress](ctx, "bridgeAddress", GetRef(m.GetBridgeAddress()), WriteComplex[BridgeAddress](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
		}

		// Simple Field (serialInterfaceAddress)
		if pushErr := writeBuffer.PushContext("serialInterfaceAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serialInterfaceAddress")
		}
		_serialInterfaceAddressErr := writeBuffer.WriteSerializable(ctx, m.GetSerialInterfaceAddress())
		if popErr := writeBuffer.PopContext("serialInterfaceAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serialInterfaceAddress")
		}
		if _serialInterfaceAddressErr != nil {
			return errors.Wrap(_serialInterfaceAddressErr, "Error serializing 'serialInterfaceAddress' field")
		}

		if err := WriteOptionalField[byte](ctx, "reservedByte", m.GetReservedByte(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'reservedByte' field")
		}

		if err := WriteOptionalField[ReplyNetwork](ctx, "replyNetwork", GetRef(m.GetReplyNetwork()), WriteComplex[ReplyNetwork](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'replyNetwork' field")
		}

		if popErr := writeBuffer.PopContext("CALReplyLong"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyLong")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALReplyLong) isCALReplyLong() bool {
	return true
}

func (m *_CALReplyLong) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
