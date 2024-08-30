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

// RegisterNodesResponse is the corresponding interface of RegisterNodesResponse
type RegisterNodesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfRegisteredNodeIds returns NoOfRegisteredNodeIds (property field)
	GetNoOfRegisteredNodeIds() int32
	// GetRegisteredNodeIds returns RegisteredNodeIds (property field)
	GetRegisteredNodeIds() []NodeId
}

// RegisterNodesResponseExactly can be used when we want exactly this type and not a type which fulfills RegisterNodesResponse.
// This is useful for switch cases.
type RegisterNodesResponseExactly interface {
	RegisterNodesResponse
	isRegisterNodesResponse() bool
}

// _RegisterNodesResponse is the data-structure of this message
type _RegisterNodesResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader        ExtensionObjectDefinition
	NoOfRegisteredNodeIds int32
	RegisteredNodeIds     []NodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterNodesResponse) GetIdentifier() string {
	return "563"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterNodesResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RegisterNodesResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterNodesResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_RegisterNodesResponse) GetNoOfRegisteredNodeIds() int32 {
	return m.NoOfRegisteredNodeIds
}

func (m *_RegisterNodesResponse) GetRegisteredNodeIds() []NodeId {
	return m.RegisteredNodeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRegisterNodesResponse factory function for _RegisterNodesResponse
func NewRegisterNodesResponse(responseHeader ExtensionObjectDefinition, noOfRegisteredNodeIds int32, registeredNodeIds []NodeId) *_RegisterNodesResponse {
	_result := &_RegisterNodesResponse{
		ResponseHeader:             responseHeader,
		NoOfRegisteredNodeIds:      noOfRegisteredNodeIds,
		RegisteredNodeIds:          registeredNodeIds,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRegisterNodesResponse(structType any) RegisterNodesResponse {
	if casted, ok := structType.(RegisterNodesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterNodesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterNodesResponse) GetTypeName() string {
	return "RegisterNodesResponse"
}

func (m *_RegisterNodesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfRegisteredNodeIds)
	lengthInBits += 32

	// Array field
	if len(m.RegisteredNodeIds) > 0 {
		for _curItem, element := range m.RegisteredNodeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RegisteredNodeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterNodesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RegisterNodesResponseParse(ctx context.Context, theBytes []byte, identifier string) (RegisterNodesResponse, error) {
	return RegisterNodesResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RegisterNodesResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RegisterNodesResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RegisterNodesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterNodesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of RegisterNodesResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfRegisteredNodeIds)
	_noOfRegisteredNodeIds, _noOfRegisteredNodeIdsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfRegisteredNodeIds", 32)
	if _noOfRegisteredNodeIdsErr != nil {
		return nil, errors.Wrap(_noOfRegisteredNodeIdsErr, "Error parsing 'noOfRegisteredNodeIds' field of RegisterNodesResponse")
	}
	noOfRegisteredNodeIds := _noOfRegisteredNodeIds

	registeredNodeIds, err := ReadCountArrayField[NodeId](ctx, "registeredNodeIds", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer), uint64(noOfRegisteredNodeIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'registeredNodeIds' field"))
	}

	if closeErr := readBuffer.CloseContext("RegisterNodesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterNodesResponse")
	}

	// Create a partially initialized instance
	_child := &_RegisterNodesResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfRegisteredNodeIds:      noOfRegisteredNodeIds,
		RegisteredNodeIds:          registeredNodeIds,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RegisterNodesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterNodesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterNodesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterNodesResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfRegisteredNodeIds)
		noOfRegisteredNodeIds := int32(m.GetNoOfRegisteredNodeIds())
		_noOfRegisteredNodeIdsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfRegisteredNodeIds", 32, int32((noOfRegisteredNodeIds)))
		if _noOfRegisteredNodeIdsErr != nil {
			return errors.Wrap(_noOfRegisteredNodeIdsErr, "Error serializing 'noOfRegisteredNodeIds' field")
		}

		// Array Field (registeredNodeIds)
		if pushErr := writeBuffer.PushContext("registeredNodeIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for registeredNodeIds")
		}
		for _curItem, _element := range m.GetRegisteredNodeIds() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetRegisteredNodeIds()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'registeredNodeIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("registeredNodeIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for registeredNodeIds")
		}

		if popErr := writeBuffer.PopContext("RegisterNodesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterNodesResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterNodesResponse) isRegisterNodesResponse() bool {
	return true
}

func (m *_RegisterNodesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
