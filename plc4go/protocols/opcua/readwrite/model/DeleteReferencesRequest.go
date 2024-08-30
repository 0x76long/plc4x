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

// DeleteReferencesRequest is the corresponding interface of DeleteReferencesRequest
type DeleteReferencesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfReferencesToDelete returns NoOfReferencesToDelete (property field)
	GetNoOfReferencesToDelete() int32
	// GetReferencesToDelete returns ReferencesToDelete (property field)
	GetReferencesToDelete() []ExtensionObjectDefinition
}

// DeleteReferencesRequestExactly can be used when we want exactly this type and not a type which fulfills DeleteReferencesRequest.
// This is useful for switch cases.
type DeleteReferencesRequestExactly interface {
	DeleteReferencesRequest
	isDeleteReferencesRequest() bool
}

// _DeleteReferencesRequest is the data-structure of this message
type _DeleteReferencesRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader          ExtensionObjectDefinition
	NoOfReferencesToDelete int32
	ReferencesToDelete     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesRequest) GetIdentifier() string {
	return "506"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteReferencesRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteReferencesRequest) GetNoOfReferencesToDelete() int32 {
	return m.NoOfReferencesToDelete
}

func (m *_DeleteReferencesRequest) GetReferencesToDelete() []ExtensionObjectDefinition {
	return m.ReferencesToDelete
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteReferencesRequest factory function for _DeleteReferencesRequest
func NewDeleteReferencesRequest(requestHeader ExtensionObjectDefinition, noOfReferencesToDelete int32, referencesToDelete []ExtensionObjectDefinition) *_DeleteReferencesRequest {
	_result := &_DeleteReferencesRequest{
		RequestHeader:              requestHeader,
		NoOfReferencesToDelete:     noOfReferencesToDelete,
		ReferencesToDelete:         referencesToDelete,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteReferencesRequest(structType any) DeleteReferencesRequest {
	if casted, ok := structType.(DeleteReferencesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesRequest) GetTypeName() string {
	return "DeleteReferencesRequest"
}

func (m *_DeleteReferencesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfReferencesToDelete)
	lengthInBits += 32

	// Array field
	if len(m.ReferencesToDelete) > 0 {
		for _curItem, element := range m.ReferencesToDelete {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencesToDelete), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteReferencesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteReferencesRequestParse(ctx context.Context, theBytes []byte, identifier string) (DeleteReferencesRequest, error) {
	return DeleteReferencesRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteReferencesRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteReferencesRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteReferencesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of DeleteReferencesRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfReferencesToDelete)
	_noOfReferencesToDelete, _noOfReferencesToDeleteErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfReferencesToDelete", 32)
	if _noOfReferencesToDeleteErr != nil {
		return nil, errors.Wrap(_noOfReferencesToDeleteErr, "Error parsing 'noOfReferencesToDelete' field of DeleteReferencesRequest")
	}
	noOfReferencesToDelete := _noOfReferencesToDelete

	referencesToDelete, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "referencesToDelete", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("387"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfReferencesToDelete))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencesToDelete' field"))
	}

	if closeErr := readBuffer.CloseContext("DeleteReferencesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesRequest")
	}

	// Create a partially initialized instance
	_child := &_DeleteReferencesRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfReferencesToDelete:     noOfReferencesToDelete,
		ReferencesToDelete:         referencesToDelete,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteReferencesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesRequest")
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

		// Simple Field (noOfReferencesToDelete)
		noOfReferencesToDelete := int32(m.GetNoOfReferencesToDelete())
		_noOfReferencesToDeleteErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfReferencesToDelete", 32, int32((noOfReferencesToDelete)))
		if _noOfReferencesToDeleteErr != nil {
			return errors.Wrap(_noOfReferencesToDeleteErr, "Error serializing 'noOfReferencesToDelete' field")
		}

		// Array Field (referencesToDelete)
		if pushErr := writeBuffer.PushContext("referencesToDelete", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referencesToDelete")
		}
		for _curItem, _element := range m.GetReferencesToDelete() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReferencesToDelete()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'referencesToDelete' field")
			}
		}
		if popErr := writeBuffer.PopContext("referencesToDelete", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referencesToDelete")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesRequest) isDeleteReferencesRequest() bool {
	return true
}

func (m *_DeleteReferencesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
