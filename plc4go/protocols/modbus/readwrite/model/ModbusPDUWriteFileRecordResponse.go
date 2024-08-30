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

// ModbusPDUWriteFileRecordResponse is the corresponding interface of ModbusPDUWriteFileRecordResponse
type ModbusPDUWriteFileRecordResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetItems returns Items (property field)
	GetItems() []ModbusPDUWriteFileRecordResponseItem
}

// ModbusPDUWriteFileRecordResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteFileRecordResponse.
// This is useful for switch cases.
type ModbusPDUWriteFileRecordResponseExactly interface {
	ModbusPDUWriteFileRecordResponse
	isModbusPDUWriteFileRecordResponse() bool
}

// _ModbusPDUWriteFileRecordResponse is the data-structure of this message
type _ModbusPDUWriteFileRecordResponse struct {
	*_ModbusPDU
	Items []ModbusPDUWriteFileRecordResponseItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteFileRecordResponse) GetFunctionFlag() uint8 {
	return 0x15
}

func (m *_ModbusPDUWriteFileRecordResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUWriteFileRecordResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponse) GetItems() []ModbusPDUWriteFileRecordResponseItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteFileRecordResponse factory function for _ModbusPDUWriteFileRecordResponse
func NewModbusPDUWriteFileRecordResponse(items []ModbusPDUWriteFileRecordResponseItem) *_ModbusPDUWriteFileRecordResponse {
	_result := &_ModbusPDUWriteFileRecordResponse{
		Items:      items,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteFileRecordResponse(structType any) ModbusPDUWriteFileRecordResponse {
	if casted, ok := structType.(ModbusPDUWriteFileRecordResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponse) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponse"
}

func (m *_ModbusPDUWriteFileRecordResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteFileRecordResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteFileRecordResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUWriteFileRecordResponse, error) {
	return ModbusPDUWriteFileRecordResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUWriteFileRecordResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteFileRecordResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUWriteFileRecordResponse")
	}

	items, err := ReadLengthArrayField[ModbusPDUWriteFileRecordResponseItem](ctx, "items", ReadComplex[ModbusPDUWriteFileRecordResponseItem](ModbusPDUWriteFileRecordResponseItemParseWithBuffer, readBuffer), int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteFileRecordResponse{
		_ModbusPDU: &_ModbusPDU{},
		Items:      items,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteFileRecordResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteFileRecordResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	itemsArraySizeInBytes := func(items []ModbusPDUWriteFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes(ctx))
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordResponse")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		_byteCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("byteCount", 8, uint8((byteCount)))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteFileRecordResponse) isModbusPDUWriteFileRecordResponse() bool {
	return true
}

func (m *_ModbusPDUWriteFileRecordResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
