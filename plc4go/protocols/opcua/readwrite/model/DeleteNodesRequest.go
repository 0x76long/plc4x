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

// DeleteNodesRequest is the corresponding interface of DeleteNodesRequest
type DeleteNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToDelete returns NoOfNodesToDelete (property field)
	GetNoOfNodesToDelete() int32
	// GetNodesToDelete returns NodesToDelete (property field)
	GetNodesToDelete() []ExtensionObjectDefinition
}

// DeleteNodesRequestExactly can be used when we want exactly this type and not a type which fulfills DeleteNodesRequest.
// This is useful for switch cases.
type DeleteNodesRequestExactly interface {
	DeleteNodesRequest
	isDeleteNodesRequest() bool
}

// _DeleteNodesRequest is the data-structure of this message
type _DeleteNodesRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader     ExtensionObjectDefinition
	NoOfNodesToDelete int32
	NodesToDelete     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesRequest) GetIdentifier() string {
	return "500"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteNodesRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteNodesRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteNodesRequest) GetNoOfNodesToDelete() int32 {
	return m.NoOfNodesToDelete
}

func (m *_DeleteNodesRequest) GetNodesToDelete() []ExtensionObjectDefinition {
	return m.NodesToDelete
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteNodesRequest factory function for _DeleteNodesRequest
func NewDeleteNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToDelete int32, nodesToDelete []ExtensionObjectDefinition) *_DeleteNodesRequest {
	_result := &_DeleteNodesRequest{
		RequestHeader:              requestHeader,
		NoOfNodesToDelete:          noOfNodesToDelete,
		NodesToDelete:              nodesToDelete,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteNodesRequest(structType any) DeleteNodesRequest {
	if casted, ok := structType.(DeleteNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteNodesRequest) GetTypeName() string {
	return "DeleteNodesRequest"
}

func (m *_DeleteNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToDelete)
	lengthInBits += 32

	// Array field
	if len(m.NodesToDelete) > 0 {
		for _curItem, element := range m.NodesToDelete {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToDelete), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteNodesRequestParse(ctx context.Context, theBytes []byte, identifier string) (DeleteNodesRequest, error) {
	return DeleteNodesRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteNodesRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteNodesRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of DeleteNodesRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfNodesToDelete)
	_noOfNodesToDelete, _noOfNodesToDeleteErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfNodesToDelete", 32)
	if _noOfNodesToDeleteErr != nil {
		return nil, errors.Wrap(_noOfNodesToDeleteErr, "Error parsing 'noOfNodesToDelete' field of DeleteNodesRequest")
	}
	noOfNodesToDelete := _noOfNodesToDelete

	nodesToDelete, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToDelete", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("384"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfNodesToDelete))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToDelete' field"))
	}

	if closeErr := readBuffer.CloseContext("DeleteNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteNodesRequest")
	}

	// Create a partially initialized instance
	_child := &_DeleteNodesRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfNodesToDelete:          noOfNodesToDelete,
		NodesToDelete:              nodesToDelete,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteNodesRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (noOfNodesToDelete)
		noOfNodesToDelete := int32(m.GetNoOfNodesToDelete())
		_noOfNodesToDeleteErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfNodesToDelete", 32, int32((noOfNodesToDelete)))
		if _noOfNodesToDeleteErr != nil {
			return errors.Wrap(_noOfNodesToDeleteErr, "Error serializing 'noOfNodesToDelete' field")
		}

		// Array Field (nodesToDelete)
		if pushErr := writeBuffer.PushContext("nodesToDelete", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodesToDelete")
		}
		for _curItem, _element := range m.GetNodesToDelete() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodesToDelete()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodesToDelete' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodesToDelete", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodesToDelete")
		}

		if popErr := writeBuffer.PopContext("DeleteNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteNodesRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteNodesRequest) isDeleteNodesRequest() bool {
	return true
}

func (m *_DeleteNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
