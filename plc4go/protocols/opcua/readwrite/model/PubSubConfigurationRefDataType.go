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

// PubSubConfigurationRefDataType is the corresponding interface of PubSubConfigurationRefDataType
type PubSubConfigurationRefDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetConfigurationMask returns ConfigurationMask (property field)
	GetConfigurationMask() PubSubConfigurationRefMask
	// GetElementIndex returns ElementIndex (property field)
	GetElementIndex() uint16
	// GetConnectionIndex returns ConnectionIndex (property field)
	GetConnectionIndex() uint16
	// GetGroupIndex returns GroupIndex (property field)
	GetGroupIndex() uint16
}

// PubSubConfigurationRefDataTypeExactly can be used when we want exactly this type and not a type which fulfills PubSubConfigurationRefDataType.
// This is useful for switch cases.
type PubSubConfigurationRefDataTypeExactly interface {
	PubSubConfigurationRefDataType
	isPubSubConfigurationRefDataType() bool
}

// _PubSubConfigurationRefDataType is the data-structure of this message
type _PubSubConfigurationRefDataType struct {
	*_ExtensionObjectDefinition
	ConfigurationMask PubSubConfigurationRefMask
	ElementIndex      uint16
	ConnectionIndex   uint16
	GroupIndex        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfigurationRefDataType) GetIdentifier() string {
	return "25521"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfigurationRefDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PubSubConfigurationRefDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfigurationRefDataType) GetConfigurationMask() PubSubConfigurationRefMask {
	return m.ConfigurationMask
}

func (m *_PubSubConfigurationRefDataType) GetElementIndex() uint16 {
	return m.ElementIndex
}

func (m *_PubSubConfigurationRefDataType) GetConnectionIndex() uint16 {
	return m.ConnectionIndex
}

func (m *_PubSubConfigurationRefDataType) GetGroupIndex() uint16 {
	return m.GroupIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPubSubConfigurationRefDataType factory function for _PubSubConfigurationRefDataType
func NewPubSubConfigurationRefDataType(configurationMask PubSubConfigurationRefMask, elementIndex uint16, connectionIndex uint16, groupIndex uint16) *_PubSubConfigurationRefDataType {
	_result := &_PubSubConfigurationRefDataType{
		ConfigurationMask:          configurationMask,
		ElementIndex:               elementIndex,
		ConnectionIndex:            connectionIndex,
		GroupIndex:                 groupIndex,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPubSubConfigurationRefDataType(structType any) PubSubConfigurationRefDataType {
	if casted, ok := structType.(PubSubConfigurationRefDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfigurationRefDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfigurationRefDataType) GetTypeName() string {
	return "PubSubConfigurationRefDataType"
}

func (m *_PubSubConfigurationRefDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (configurationMask)
	lengthInBits += 32

	// Simple field (elementIndex)
	lengthInBits += 16

	// Simple field (connectionIndex)
	lengthInBits += 16

	// Simple field (groupIndex)
	lengthInBits += 16

	return lengthInBits
}

func (m *_PubSubConfigurationRefDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubConfigurationRefDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PubSubConfigurationRefDataType, error) {
	return PubSubConfigurationRefDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PubSubConfigurationRefDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PubSubConfigurationRefDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PubSubConfigurationRefDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfigurationRefDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (configurationMask)
	if pullErr := readBuffer.PullContext("configurationMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for configurationMask")
	}
	_configurationMask, _configurationMaskErr := PubSubConfigurationRefMaskParseWithBuffer(ctx, readBuffer)
	if _configurationMaskErr != nil {
		return nil, errors.Wrap(_configurationMaskErr, "Error parsing 'configurationMask' field of PubSubConfigurationRefDataType")
	}
	configurationMask := _configurationMask
	if closeErr := readBuffer.CloseContext("configurationMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for configurationMask")
	}

	// Simple Field (elementIndex)
	_elementIndex, _elementIndexErr := /*TODO: migrate me*/ readBuffer.ReadUint16("elementIndex", 16)
	if _elementIndexErr != nil {
		return nil, errors.Wrap(_elementIndexErr, "Error parsing 'elementIndex' field of PubSubConfigurationRefDataType")
	}
	elementIndex := _elementIndex

	// Simple Field (connectionIndex)
	_connectionIndex, _connectionIndexErr := /*TODO: migrate me*/ readBuffer.ReadUint16("connectionIndex", 16)
	if _connectionIndexErr != nil {
		return nil, errors.Wrap(_connectionIndexErr, "Error parsing 'connectionIndex' field of PubSubConfigurationRefDataType")
	}
	connectionIndex := _connectionIndex

	// Simple Field (groupIndex)
	_groupIndex, _groupIndexErr := /*TODO: migrate me*/ readBuffer.ReadUint16("groupIndex", 16)
	if _groupIndexErr != nil {
		return nil, errors.Wrap(_groupIndexErr, "Error parsing 'groupIndex' field of PubSubConfigurationRefDataType")
	}
	groupIndex := _groupIndex

	if closeErr := readBuffer.CloseContext("PubSubConfigurationRefDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfigurationRefDataType")
	}

	// Create a partially initialized instance
	_child := &_PubSubConfigurationRefDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ConfigurationMask:          configurationMask,
		ElementIndex:               elementIndex,
		ConnectionIndex:            connectionIndex,
		GroupIndex:                 groupIndex,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PubSubConfigurationRefDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfigurationRefDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfigurationRefDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfigurationRefDataType")
		}

		// Simple Field (configurationMask)
		if pushErr := writeBuffer.PushContext("configurationMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for configurationMask")
		}
		_configurationMaskErr := writeBuffer.WriteSerializable(ctx, m.GetConfigurationMask())
		if popErr := writeBuffer.PopContext("configurationMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for configurationMask")
		}
		if _configurationMaskErr != nil {
			return errors.Wrap(_configurationMaskErr, "Error serializing 'configurationMask' field")
		}

		// Simple Field (elementIndex)
		elementIndex := uint16(m.GetElementIndex())
		_elementIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("elementIndex", 16, uint16((elementIndex)))
		if _elementIndexErr != nil {
			return errors.Wrap(_elementIndexErr, "Error serializing 'elementIndex' field")
		}

		// Simple Field (connectionIndex)
		connectionIndex := uint16(m.GetConnectionIndex())
		_connectionIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("connectionIndex", 16, uint16((connectionIndex)))
		if _connectionIndexErr != nil {
			return errors.Wrap(_connectionIndexErr, "Error serializing 'connectionIndex' field")
		}

		// Simple Field (groupIndex)
		groupIndex := uint16(m.GetGroupIndex())
		_groupIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("groupIndex", 16, uint16((groupIndex)))
		if _groupIndexErr != nil {
			return errors.Wrap(_groupIndexErr, "Error serializing 'groupIndex' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfigurationRefDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfigurationRefDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfigurationRefDataType) isPubSubConfigurationRefDataType() bool {
	return true
}

func (m *_PubSubConfigurationRefDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
