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

// NodeId is the corresponding interface of NodeId
type NodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeIdTypeDefinition
	// GetId returns Id (virtual field)
	GetId() string
}

// NodeIdExactly can be used when we want exactly this type and not a type which fulfills NodeId.
// This is useful for switch cases.
type NodeIdExactly interface {
	NodeId
	isNodeId() bool
}

// _NodeId is the data-structure of this message
type _NodeId struct {
	NodeId NodeIdTypeDefinition
	// Reserved Fields
	reservedField0 *int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeId) GetNodeId() NodeIdTypeDefinition {
	return m.NodeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeId) GetId() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetNodeId().GetIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeId factory function for _NodeId
func NewNodeId(nodeId NodeIdTypeDefinition) *_NodeId {
	return &_NodeId{NodeId: nodeId}
}

// Deprecated: use the interface for direct cast
func CastNodeId(structType any) NodeId {
	if casted, ok := structType.(NodeId); ok {
		return casted
	}
	if casted, ok := structType.(*NodeId); ok {
		return *casted
	}
	return nil
}

func (m *_NodeId) GetTypeName() string {
	return "NodeId"
}

func (m *_NodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdParse(ctx context.Context, theBytes []byte) (NodeId, error) {
	return NodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *int8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NodeId")
		}
		if reserved != int8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": int8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdTypeDefinitionParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of NodeId")
	}
	nodeId := _nodeId.(NodeIdTypeDefinition)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Virtual field
	_id := nodeId.GetIdentifier()
	id := fmt.Sprintf("%v", _id)
	_ = id

	if closeErr := readBuffer.CloseContext("NodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeId")
	}

	// Create the instance
	return &_NodeId{
		NodeId:         nodeId,
		reservedField0: reservedField0,
	}, nil
}

func (m *_NodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NodeId")
	}

	// Reserved Field (reserved)
	{
		var reserved int8 = int8(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": int8(0x00),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := /*TODO: migrate me*/ writeBuffer.WriteInt8("reserved", 2, int8(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (nodeId)
	if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for nodeId")
	}
	_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
	if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for nodeId")
	}
	if _nodeIdErr != nil {
		return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
	}
	// Virtual field
	id := m.GetId()
	_ = id
	if _idErr := writeBuffer.WriteVirtual(ctx, "id", m.GetId()); _idErr != nil {
		return errors.Wrap(_idErr, "Error serializing 'id' field")
	}

	if popErr := writeBuffer.PopContext("NodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NodeId")
	}
	return nil
}

func (m *_NodeId) isNodeId() bool {
	return true
}

func (m *_NodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
