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

// S7ParameterWriteVarResponse is the corresponding interface of S7ParameterWriteVarResponse
type S7ParameterWriteVarResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetNumItems returns NumItems (property field)
	GetNumItems() uint8
}

// S7ParameterWriteVarResponseExactly can be used when we want exactly this type and not a type which fulfills S7ParameterWriteVarResponse.
// This is useful for switch cases.
type S7ParameterWriteVarResponseExactly interface {
	S7ParameterWriteVarResponse
	isS7ParameterWriteVarResponse() bool
}

// _S7ParameterWriteVarResponse is the data-structure of this message
type _S7ParameterWriteVarResponse struct {
	*_S7Parameter
	NumItems uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterWriteVarResponse) GetParameterType() uint8 {
	return 0x05
}

func (m *_S7ParameterWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterWriteVarResponse) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterWriteVarResponse) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterWriteVarResponse) GetNumItems() uint8 {
	return m.NumItems
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterWriteVarResponse factory function for _S7ParameterWriteVarResponse
func NewS7ParameterWriteVarResponse(numItems uint8) *_S7ParameterWriteVarResponse {
	_result := &_S7ParameterWriteVarResponse{
		NumItems:     numItems,
		_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterWriteVarResponse(structType any) S7ParameterWriteVarResponse {
	if casted, ok := structType.(S7ParameterWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterWriteVarResponse) GetTypeName() string {
	return "S7ParameterWriteVarResponse"
}

func (m *_S7ParameterWriteVarResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterWriteVarResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterWriteVarResponseParse(ctx context.Context, theBytes []byte, messageType uint8) (S7ParameterWriteVarResponse, error) {
	return S7ParameterWriteVarResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterWriteVarResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterWriteVarResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterWriteVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numItems)
	_numItems, _numItemsErr := /*TODO: migrate me*/ readBuffer.ReadUint8("numItems", 8)
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field of S7ParameterWriteVarResponse")
	}
	numItems := _numItems

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterWriteVarResponse")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterWriteVarResponse{
		_S7Parameter: &_S7Parameter{},
		NumItems:     numItems,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterWriteVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterWriteVarResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterWriteVarResponse")
		}

		// Simple Field (numItems)
		numItems := uint8(m.GetNumItems())
		_numItemsErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numItems", 8, uint8((numItems)))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterWriteVarResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterWriteVarResponse) isS7ParameterWriteVarResponse() bool {
	return true
}

func (m *_S7ParameterWriteVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
