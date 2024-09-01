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

// CBusPointToPointToMultiPointCommandStatus is the corresponding interface of CBusPointToPointToMultiPointCommandStatus
type CBusPointToPointToMultiPointCommandStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToPointToMultiPointCommand
	// GetStatusRequest returns StatusRequest (property field)
	GetStatusRequest() StatusRequest
}

// CBusPointToPointToMultiPointCommandStatusExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointToMultiPointCommandStatus.
// This is useful for switch cases.
type CBusPointToPointToMultiPointCommandStatusExactly interface {
	CBusPointToPointToMultiPointCommandStatus
	isCBusPointToPointToMultiPointCommandStatus() bool
}

// _CBusPointToPointToMultiPointCommandStatus is the data-structure of this message
type _CBusPointToPointToMultiPointCommandStatus struct {
	*_CBusPointToPointToMultiPointCommand
	StatusRequest StatusRequest
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

func (m *_CBusPointToPointToMultiPointCommandStatus) InitializeParent(parent CBusPointToPointToMultiPointCommand, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte) {
	m.BridgeAddress = bridgeAddress
	m.NetworkRoute = networkRoute
	m.PeekedApplication = peekedApplication
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetParent() CBusPointToPointToMultiPointCommand {
	return m._CBusPointToPointToMultiPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultiPointCommandStatus) GetStatusRequest() StatusRequest {
	return m.StatusRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointToMultiPointCommandStatus factory function for _CBusPointToPointToMultiPointCommandStatus
func NewCBusPointToPointToMultiPointCommandStatus(statusRequest StatusRequest, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, cBusOptions CBusOptions) *_CBusPointToPointToMultiPointCommandStatus {
	_result := &_CBusPointToPointToMultiPointCommandStatus{
		StatusRequest:                        statusRequest,
		_CBusPointToPointToMultiPointCommand: NewCBusPointToPointToMultiPointCommand(bridgeAddress, networkRoute, peekedApplication, cBusOptions),
	}
	_result._CBusPointToPointToMultiPointCommand._CBusPointToPointToMultiPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultiPointCommandStatus(structType any) CBusPointToPointToMultiPointCommandStatus {
	if casted, ok := structType.(CBusPointToPointToMultiPointCommandStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultiPointCommandStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetTypeName() string {
	return "CBusPointToPointToMultiPointCommandStatus"
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (statusRequest)
	lengthInBits += m.StatusRequest.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToPointToMultiPointCommandStatusParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommandStatus, error) {
	return CBusPointToPointToMultiPointCommandStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointToMultiPointCommandStatusParseWithBufferProducer(cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusPointToPointToMultiPointCommandStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusPointToPointToMultiPointCommandStatus, error) {
		return CBusPointToPointToMultiPointCommandStatusParseWithBuffer(ctx, readBuffer, cBusOptions)
	}
}

func CBusPointToPointToMultiPointCommandStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommandStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultiPointCommandStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultiPointCommandStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0xFF))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	statusRequest, err := ReadSimpleField[StatusRequest](ctx, "statusRequest", ReadComplex[StatusRequest](StatusRequestParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusRequest' field"))
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultiPointCommandStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultiPointCommandStatus")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToPointToMultiPointCommandStatus{
		_CBusPointToPointToMultiPointCommand: &_CBusPointToPointToMultiPointCommand{
			CBusOptions: cBusOptions,
		},
		StatusRequest:  statusRequest,
		reservedField0: reservedField0,
	}
	_child._CBusPointToPointToMultiPointCommand._CBusPointToPointToMultiPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointToMultiPointCommandStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultiPointCommandStatus")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0xFF), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		// Simple Field (statusRequest)
		if pushErr := writeBuffer.PushContext("statusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusRequest")
		}
		_statusRequestErr := writeBuffer.WriteSerializable(ctx, m.GetStatusRequest())
		if popErr := writeBuffer.PopContext("statusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusRequest")
		}
		if _statusRequestErr != nil {
			return errors.Wrap(_statusRequestErr, "Error serializing 'statusRequest' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointToMultiPointCommandStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultiPointCommandStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointToMultiPointCommandStatus) isCBusPointToPointToMultiPointCommandStatus() bool {
	return true
}

func (m *_CBusPointToPointToMultiPointCommandStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
