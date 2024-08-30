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

// CBusHeader is the corresponding interface of CBusHeader
type CBusHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPriorityClass returns PriorityClass (property field)
	GetPriorityClass() PriorityClass
	// GetDp returns Dp (property field)
	GetDp() bool
	// GetRc returns Rc (property field)
	GetRc() uint8
	// GetDestinationAddressType returns DestinationAddressType (property field)
	GetDestinationAddressType() DestinationAddressType
}

// CBusHeaderExactly can be used when we want exactly this type and not a type which fulfills CBusHeader.
// This is useful for switch cases.
type CBusHeaderExactly interface {
	CBusHeader
	isCBusHeader() bool
}

// _CBusHeader is the data-structure of this message
type _CBusHeader struct {
	PriorityClass          PriorityClass
	Dp                     bool
	Rc                     uint8
	DestinationAddressType DestinationAddressType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusHeader) GetPriorityClass() PriorityClass {
	return m.PriorityClass
}

func (m *_CBusHeader) GetDp() bool {
	return m.Dp
}

func (m *_CBusHeader) GetRc() uint8 {
	return m.Rc
}

func (m *_CBusHeader) GetDestinationAddressType() DestinationAddressType {
	return m.DestinationAddressType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusHeader factory function for _CBusHeader
func NewCBusHeader(priorityClass PriorityClass, dp bool, rc uint8, destinationAddressType DestinationAddressType) *_CBusHeader {
	return &_CBusHeader{PriorityClass: priorityClass, Dp: dp, Rc: rc, DestinationAddressType: destinationAddressType}
}

// Deprecated: use the interface for direct cast
func CastCBusHeader(structType any) CBusHeader {
	if casted, ok := structType.(CBusHeader); ok {
		return casted
	}
	if casted, ok := structType.(*CBusHeader); ok {
		return *casted
	}
	return nil
}

func (m *_CBusHeader) GetTypeName() string {
	return "CBusHeader"
}

func (m *_CBusHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (priorityClass)
	lengthInBits += 2

	// Simple field (dp)
	lengthInBits += 1

	// Simple field (rc)
	lengthInBits += 2

	// Simple field (destinationAddressType)
	lengthInBits += 3

	return lengthInBits
}

func (m *_CBusHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusHeaderParse(ctx context.Context, theBytes []byte) (CBusHeader, error) {
	return CBusHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CBusHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CBusHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CBusHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityClass)
	if pullErr := readBuffer.PullContext("priorityClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityClass")
	}
	_priorityClass, _priorityClassErr := PriorityClassParseWithBuffer(ctx, readBuffer)
	if _priorityClassErr != nil {
		return nil, errors.Wrap(_priorityClassErr, "Error parsing 'priorityClass' field of CBusHeader")
	}
	priorityClass := _priorityClass
	if closeErr := readBuffer.CloseContext("priorityClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityClass")
	}

	// Simple Field (dp)
	_dp, _dpErr := /*TODO: migrate me*/ readBuffer.ReadBit("dp")
	if _dpErr != nil {
		return nil, errors.Wrap(_dpErr, "Error parsing 'dp' field of CBusHeader")
	}
	dp := _dp

	// Simple Field (rc)
	_rc, _rcErr := /*TODO: migrate me*/ readBuffer.ReadUint8("rc", 2)
	if _rcErr != nil {
		return nil, errors.Wrap(_rcErr, "Error parsing 'rc' field of CBusHeader")
	}
	rc := _rc

	// Simple Field (destinationAddressType)
	if pullErr := readBuffer.PullContext("destinationAddressType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destinationAddressType")
	}
	_destinationAddressType, _destinationAddressTypeErr := DestinationAddressTypeParseWithBuffer(ctx, readBuffer)
	if _destinationAddressTypeErr != nil {
		return nil, errors.Wrap(_destinationAddressTypeErr, "Error parsing 'destinationAddressType' field of CBusHeader")
	}
	destinationAddressType := _destinationAddressType
	if closeErr := readBuffer.CloseContext("destinationAddressType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destinationAddressType")
	}

	if closeErr := readBuffer.CloseContext("CBusHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusHeader")
	}

	// Create the instance
	return &_CBusHeader{
		PriorityClass:          priorityClass,
		Dp:                     dp,
		Rc:                     rc,
		DestinationAddressType: destinationAddressType,
	}, nil
}

func (m *_CBusHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusHeader")
	}

	// Simple Field (priorityClass)
	if pushErr := writeBuffer.PushContext("priorityClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for priorityClass")
	}
	_priorityClassErr := writeBuffer.WriteSerializable(ctx, m.GetPriorityClass())
	if popErr := writeBuffer.PopContext("priorityClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priorityClass")
	}
	if _priorityClassErr != nil {
		return errors.Wrap(_priorityClassErr, "Error serializing 'priorityClass' field")
	}

	// Simple Field (dp)
	dp := bool(m.GetDp())
	_dpErr := /*TODO: migrate me*/ writeBuffer.WriteBit("dp", (dp))
	if _dpErr != nil {
		return errors.Wrap(_dpErr, "Error serializing 'dp' field")
	}

	// Simple Field (rc)
	rc := uint8(m.GetRc())
	_rcErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("rc", 2, uint8((rc)))
	if _rcErr != nil {
		return errors.Wrap(_rcErr, "Error serializing 'rc' field")
	}

	// Simple Field (destinationAddressType)
	if pushErr := writeBuffer.PushContext("destinationAddressType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for destinationAddressType")
	}
	_destinationAddressTypeErr := writeBuffer.WriteSerializable(ctx, m.GetDestinationAddressType())
	if popErr := writeBuffer.PopContext("destinationAddressType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for destinationAddressType")
	}
	if _destinationAddressTypeErr != nil {
		return errors.Wrap(_destinationAddressTypeErr, "Error serializing 'destinationAddressType' field")
	}

	if popErr := writeBuffer.PopContext("CBusHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusHeader")
	}
	return nil
}

func (m *_CBusHeader) isCBusHeader() bool {
	return true
}

func (m *_CBusHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
