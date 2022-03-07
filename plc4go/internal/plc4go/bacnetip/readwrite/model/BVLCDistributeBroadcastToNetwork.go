/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BVLCDistributeBroadcastToNetwork struct {
	*BVLC
	Npdu *NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

// The corresponding interface
type IBVLCDistributeBroadcastToNetwork interface {
	IBVLC
	// GetNpdu returns Npdu (property field)
	GetNpdu() *NPDU
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *BVLCDistributeBroadcastToNetwork) GetBvlcFunction() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BVLCDistributeBroadcastToNetwork) InitializeParent(parent *BVLC) {}

func (m *BVLCDistributeBroadcastToNetwork) GetParent() *BVLC {
	return m.BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BVLCDistributeBroadcastToNetwork) GetNpdu() *NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCDistributeBroadcastToNetwork factory function for BVLCDistributeBroadcastToNetwork
func NewBVLCDistributeBroadcastToNetwork(npdu *NPDU, bvlcPayloadLength uint16) *BVLC {
	child := &BVLCDistributeBroadcastToNetwork{
		Npdu: npdu,
		BVLC: NewBVLC(),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCDistributeBroadcastToNetwork(structType interface{}) *BVLCDistributeBroadcastToNetwork {
	if casted, ok := structType.(BVLCDistributeBroadcastToNetwork); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCDistributeBroadcastToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCDistributeBroadcastToNetwork(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCDistributeBroadcastToNetwork(casted.Child)
	}
	return nil
}

func (m *BVLCDistributeBroadcastToNetwork) GetTypeName() string {
	return "BVLCDistributeBroadcastToNetwork"
}

func (m *BVLCDistributeBroadcastToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCDistributeBroadcastToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits()

	return lengthInBits
}

func (m *BVLCDistributeBroadcastToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCDistributeBroadcastToNetworkParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCDistributeBroadcastToNetwork"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, pullErr
	}
	_npdu, _npduErr := NPDUParse(readBuffer, uint16(bvlcPayloadLength))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}
	npdu := CastNPDU(_npdu)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCDistributeBroadcastToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCDistributeBroadcastToNetwork{
		Npdu: CastNPDU(npdu),
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCDistributeBroadcastToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDistributeBroadcastToNetwork"); pushErr != nil {
			return pushErr
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return pushErr
		}
		_npduErr := m.Npdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return popErr
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCDistributeBroadcastToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCDistributeBroadcastToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
