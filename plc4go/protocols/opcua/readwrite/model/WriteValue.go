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

// WriteValue is the corresponding interface of WriteValue
type WriteValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetValue returns Value (property field)
	GetValue() DataValue
}

// WriteValueExactly can be used when we want exactly this type and not a type which fulfills WriteValue.
// This is useful for switch cases.
type WriteValueExactly interface {
	WriteValue
	isWriteValue() bool
}

// _WriteValue is the data-structure of this message
type _WriteValue struct {
	*_ExtensionObjectDefinition
	NodeId      NodeId
	AttributeId uint32
	IndexRange  PascalString
	Value       DataValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriteValue) GetIdentifier() string {
	return "670"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriteValue) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_WriteValue) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_WriteValue) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_WriteValue) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_WriteValue) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_WriteValue) GetValue() DataValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewWriteValue factory function for _WriteValue
func NewWriteValue(nodeId NodeId, attributeId uint32, indexRange PascalString, value DataValue) *_WriteValue {
	_result := &_WriteValue{
		NodeId:                     nodeId,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
		Value:                      value,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastWriteValue(structType any) WriteValue {
	if casted, ok := structType.(WriteValue); ok {
		return casted
	}
	if casted, ok := structType.(*WriteValue); ok {
		return *casted
	}
	return nil
}

func (m *_WriteValue) GetTypeName() string {
	return "WriteValue"
}

func (m *_WriteValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_WriteValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func WriteValueParse(ctx context.Context, theBytes []byte, identifier string) (WriteValue, error) {
	return WriteValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func WriteValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (WriteValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("WriteValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriteValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of WriteValue")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (attributeId)
	_attributeId, _attributeIdErr := /*TODO: migrate me*/ readBuffer.ReadUint32("attributeId", 32)
	if _attributeIdErr != nil {
		return nil, errors.Wrap(_attributeIdErr, "Error parsing 'attributeId' field of WriteValue")
	}
	attributeId := _attributeId

	// Simple Field (indexRange)
	if pullErr := readBuffer.PullContext("indexRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for indexRange")
	}
	_indexRange, _indexRangeErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _indexRangeErr != nil {
		return nil, errors.Wrap(_indexRangeErr, "Error parsing 'indexRange' field of WriteValue")
	}
	indexRange := _indexRange.(PascalString)
	if closeErr := readBuffer.CloseContext("indexRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for indexRange")
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := DataValueParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of WriteValue")
	}
	value := _value.(DataValue)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("WriteValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriteValue")
	}

	// Create a partially initialized instance
	_child := &_WriteValue{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
		Value:                      value,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_WriteValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriteValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriteValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriteValue")
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

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (indexRange)
		if pushErr := writeBuffer.PushContext("indexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for indexRange")
		}
		_indexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetIndexRange())
		if popErr := writeBuffer.PopContext("indexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for indexRange")
		}
		if _indexRangeErr != nil {
			return errors.Wrap(_indexRangeErr, "Error serializing 'indexRange' field")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("WriteValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriteValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriteValue) isWriteValue() bool {
	return true
}

func (m *_WriteValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
