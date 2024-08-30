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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AdsDiscovery_HEADER uint32 = 0x71146603

// AdsDiscovery is the corresponding interface of AdsDiscovery
type AdsDiscovery interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetRequestId returns RequestId (property field)
	GetRequestId() uint32
	// GetOperation returns Operation (property field)
	GetOperation() Operation
	// GetAmsNetId returns AmsNetId (property field)
	GetAmsNetId() AmsNetId
	// GetPortNumber returns PortNumber (property field)
	GetPortNumber() AdsPortNumbers
	// GetBlocks returns Blocks (property field)
	GetBlocks() []AdsDiscoveryBlock
}

// AdsDiscoveryExactly can be used when we want exactly this type and not a type which fulfills AdsDiscovery.
// This is useful for switch cases.
type AdsDiscoveryExactly interface {
	AdsDiscovery
	isAdsDiscovery() bool
}

// _AdsDiscovery is the data-structure of this message
type _AdsDiscovery struct {
	RequestId  uint32
	Operation  Operation
	AmsNetId   AmsNetId
	PortNumber AdsPortNumbers
	Blocks     []AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscovery) GetRequestId() uint32 {
	return m.RequestId
}

func (m *_AdsDiscovery) GetOperation() Operation {
	return m.Operation
}

func (m *_AdsDiscovery) GetAmsNetId() AmsNetId {
	return m.AmsNetId
}

func (m *_AdsDiscovery) GetPortNumber() AdsPortNumbers {
	return m.PortNumber
}

func (m *_AdsDiscovery) GetBlocks() []AdsDiscoveryBlock {
	return m.Blocks
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscovery) GetHeader() uint32 {
	return AdsDiscovery_HEADER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscovery factory function for _AdsDiscovery
func NewAdsDiscovery(requestId uint32, operation Operation, amsNetId AmsNetId, portNumber AdsPortNumbers, blocks []AdsDiscoveryBlock) *_AdsDiscovery {
	return &_AdsDiscovery{RequestId: requestId, Operation: operation, AmsNetId: amsNetId, PortNumber: portNumber, Blocks: blocks}
}

// Deprecated: use the interface for direct cast
func CastAdsDiscovery(structType any) AdsDiscovery {
	if casted, ok := structType.(AdsDiscovery); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscovery); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscovery) GetTypeName() string {
	return "AdsDiscovery"
}

func (m *_AdsDiscovery) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (header)
	lengthInBits += 32

	// Simple field (requestId)
	lengthInBits += 32

	// Simple field (operation)
	lengthInBits += 32

	// Simple field (amsNetId)
	lengthInBits += m.AmsNetId.GetLengthInBits(ctx)

	// Simple field (portNumber)
	lengthInBits += 16

	// Implicit Field (numBlocks)
	lengthInBits += 32

	// Array field
	if len(m.Blocks) > 0 {
		for _curItem, element := range m.Blocks {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Blocks), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdsDiscovery) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryParse(ctx context.Context, theBytes []byte) (AdsDiscovery, error) {
	return AdsDiscoveryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsDiscoveryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscovery, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsDiscovery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscovery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := readConstField[uint32](ctx, "header", ReadUnsignedInt(readBuffer, 32), AdsDiscovery.HEADER, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	_ = header

	// Simple Field (requestId)
	_requestId, _requestIdErr := /*TODO: migrate me*/ readBuffer.ReadUint32("requestId", 32)
	if _requestIdErr != nil {
		return nil, errors.Wrap(_requestIdErr, "Error parsing 'requestId' field of AdsDiscovery")
	}
	requestId := _requestId

	// Simple Field (operation)
	if pullErr := readBuffer.PullContext("operation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for operation")
	}
	_operation, _operationErr := OperationParseWithBuffer(ctx, readBuffer)
	if _operationErr != nil {
		return nil, errors.Wrap(_operationErr, "Error parsing 'operation' field of AdsDiscovery")
	}
	operation := _operation
	if closeErr := readBuffer.CloseContext("operation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for operation")
	}

	// Simple Field (amsNetId)
	if pullErr := readBuffer.PullContext("amsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for amsNetId")
	}
	_amsNetId, _amsNetIdErr := AmsNetIdParseWithBuffer(ctx, readBuffer)
	if _amsNetIdErr != nil {
		return nil, errors.Wrap(_amsNetIdErr, "Error parsing 'amsNetId' field of AdsDiscovery")
	}
	amsNetId := _amsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("amsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for amsNetId")
	}

	// Simple Field (portNumber)
	if pullErr := readBuffer.PullContext("portNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for portNumber")
	}
	_portNumber, _portNumberErr := AdsPortNumbersParseWithBuffer(ctx, readBuffer)
	if _portNumberErr != nil {
		return nil, errors.Wrap(_portNumberErr, "Error parsing 'portNumber' field of AdsDiscovery")
	}
	portNumber := _portNumber
	if closeErr := readBuffer.CloseContext("portNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for portNumber")
	}

	// Implicit Field (numBlocks) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numBlocks, _numBlocksErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("numBlocks", 32)
	_ = numBlocks
	if _numBlocksErr != nil {
		return nil, errors.Wrap(_numBlocksErr, "Error parsing 'numBlocks' field of AdsDiscovery")
	}

	blocks, err := ReadCountArrayField[AdsDiscoveryBlock](ctx, "blocks", ReadComplex[AdsDiscoveryBlock](AdsDiscoveryBlockParseWithBuffer, readBuffer), uint64(numBlocks), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'blocks' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsDiscovery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscovery")
	}

	// Create the instance
	return &_AdsDiscovery{
		RequestId:  requestId,
		Operation:  operation,
		AmsNetId:   amsNetId,
		PortNumber: portNumber,
		Blocks:     blocks,
	}, nil
}

func (m *_AdsDiscovery) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscovery) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDiscovery"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDiscovery")
	}

	// Const Field (header)
	_headerErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint32("header", 32, uint32(0x71146603))
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Simple Field (requestId)
	requestId := uint32(m.GetRequestId())
	_requestIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("requestId", 32, uint32((requestId)))
	if _requestIdErr != nil {
		return errors.Wrap(_requestIdErr, "Error serializing 'requestId' field")
	}

	// Simple Field (operation)
	if pushErr := writeBuffer.PushContext("operation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for operation")
	}
	_operationErr := writeBuffer.WriteSerializable(ctx, m.GetOperation())
	if popErr := writeBuffer.PopContext("operation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for operation")
	}
	if _operationErr != nil {
		return errors.Wrap(_operationErr, "Error serializing 'operation' field")
	}

	// Simple Field (amsNetId)
	if pushErr := writeBuffer.PushContext("amsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for amsNetId")
	}
	_amsNetIdErr := writeBuffer.WriteSerializable(ctx, m.GetAmsNetId())
	if popErr := writeBuffer.PopContext("amsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for amsNetId")
	}
	if _amsNetIdErr != nil {
		return errors.Wrap(_amsNetIdErr, "Error serializing 'amsNetId' field")
	}

	// Simple Field (portNumber)
	if pushErr := writeBuffer.PushContext("portNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for portNumber")
	}
	_portNumberErr := writeBuffer.WriteSerializable(ctx, m.GetPortNumber())
	if popErr := writeBuffer.PopContext("portNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for portNumber")
	}
	if _portNumberErr != nil {
		return errors.Wrap(_portNumberErr, "Error serializing 'portNumber' field")
	}

	// Implicit Field (numBlocks) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numBlocks := uint32(uint32(len(m.GetBlocks())))
	_numBlocksErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("numBlocks", 32, uint32((numBlocks)))
	if _numBlocksErr != nil {
		return errors.Wrap(_numBlocksErr, "Error serializing 'numBlocks' field")
	}

	// Array Field (blocks)
	if pushErr := writeBuffer.PushContext("blocks", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for blocks")
	}
	for _curItem, _element := range m.GetBlocks() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetBlocks()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'blocks' field")
		}
	}
	if popErr := writeBuffer.PopContext("blocks", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for blocks")
	}

	if popErr := writeBuffer.PopContext("AdsDiscovery"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDiscovery")
	}
	return nil
}

func (m *_AdsDiscovery) isAdsDiscovery() bool {
	return true
}

func (m *_AdsDiscovery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
