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

// NetworkGroupDataType is the corresponding interface of NetworkGroupDataType
type NetworkGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetServerUri returns ServerUri (property field)
	GetServerUri() PascalString
	// GetNoOfNetworkPaths returns NoOfNetworkPaths (property field)
	GetNoOfNetworkPaths() int32
	// GetNetworkPaths returns NetworkPaths (property field)
	GetNetworkPaths() []ExtensionObjectDefinition
}

// NetworkGroupDataTypeExactly can be used when we want exactly this type and not a type which fulfills NetworkGroupDataType.
// This is useful for switch cases.
type NetworkGroupDataTypeExactly interface {
	NetworkGroupDataType
	isNetworkGroupDataType() bool
}

// _NetworkGroupDataType is the data-structure of this message
type _NetworkGroupDataType struct {
	*_ExtensionObjectDefinition
	ServerUri        PascalString
	NoOfNetworkPaths int32
	NetworkPaths     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NetworkGroupDataType) GetIdentifier() string {
	return "11946"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NetworkGroupDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_NetworkGroupDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkGroupDataType) GetServerUri() PascalString {
	return m.ServerUri
}

func (m *_NetworkGroupDataType) GetNoOfNetworkPaths() int32 {
	return m.NoOfNetworkPaths
}

func (m *_NetworkGroupDataType) GetNetworkPaths() []ExtensionObjectDefinition {
	return m.NetworkPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkGroupDataType factory function for _NetworkGroupDataType
func NewNetworkGroupDataType(serverUri PascalString, noOfNetworkPaths int32, networkPaths []ExtensionObjectDefinition) *_NetworkGroupDataType {
	_result := &_NetworkGroupDataType{
		ServerUri:                  serverUri,
		NoOfNetworkPaths:           noOfNetworkPaths,
		NetworkPaths:               networkPaths,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNetworkGroupDataType(structType any) NetworkGroupDataType {
	if casted, ok := structType.(NetworkGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkGroupDataType) GetTypeName() string {
	return "NetworkGroupDataType"
}

func (m *_NetworkGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (serverUri)
	lengthInBits += m.ServerUri.GetLengthInBits(ctx)

	// Simple field (noOfNetworkPaths)
	lengthInBits += 32

	// Array field
	if len(m.NetworkPaths) > 0 {
		for _curItem, element := range m.NetworkPaths {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NetworkPaths), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NetworkGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NetworkGroupDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (NetworkGroupDataType, error) {
	return NetworkGroupDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NetworkGroupDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NetworkGroupDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NetworkGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (serverUri)
	if pullErr := readBuffer.PullContext("serverUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverUri")
	}
	_serverUri, _serverUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _serverUriErr != nil {
		return nil, errors.Wrap(_serverUriErr, "Error parsing 'serverUri' field of NetworkGroupDataType")
	}
	serverUri := _serverUri.(PascalString)
	if closeErr := readBuffer.CloseContext("serverUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverUri")
	}

	// Simple Field (noOfNetworkPaths)
	_noOfNetworkPaths, _noOfNetworkPathsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfNetworkPaths", 32)
	if _noOfNetworkPathsErr != nil {
		return nil, errors.Wrap(_noOfNetworkPathsErr, "Error parsing 'noOfNetworkPaths' field of NetworkGroupDataType")
	}
	noOfNetworkPaths := _noOfNetworkPaths

	networkPaths, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "networkPaths", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("11945"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfNetworkPaths))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkPaths' field"))
	}

	if closeErr := readBuffer.CloseContext("NetworkGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkGroupDataType")
	}

	// Create a partially initialized instance
	_child := &_NetworkGroupDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ServerUri:                  serverUri,
		NoOfNetworkPaths:           noOfNetworkPaths,
		NetworkPaths:               networkPaths,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NetworkGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NetworkGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NetworkGroupDataType")
		}

		// Simple Field (serverUri)
		if pushErr := writeBuffer.PushContext("serverUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverUri")
		}
		_serverUriErr := writeBuffer.WriteSerializable(ctx, m.GetServerUri())
		if popErr := writeBuffer.PopContext("serverUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverUri")
		}
		if _serverUriErr != nil {
			return errors.Wrap(_serverUriErr, "Error serializing 'serverUri' field")
		}

		// Simple Field (noOfNetworkPaths)
		noOfNetworkPaths := int32(m.GetNoOfNetworkPaths())
		_noOfNetworkPathsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfNetworkPaths", 32, int32((noOfNetworkPaths)))
		if _noOfNetworkPathsErr != nil {
			return errors.Wrap(_noOfNetworkPathsErr, "Error serializing 'noOfNetworkPaths' field")
		}

		// Array Field (networkPaths)
		if pushErr := writeBuffer.PushContext("networkPaths", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkPaths")
		}
		for _curItem, _element := range m.GetNetworkPaths() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNetworkPaths()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'networkPaths' field")
			}
		}
		if popErr := writeBuffer.PopContext("networkPaths", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkPaths")
		}

		if popErr := writeBuffer.PopContext("NetworkGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NetworkGroupDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NetworkGroupDataType) isNetworkGroupDataType() bool {
	return true
}

func (m *_NetworkGroupDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
