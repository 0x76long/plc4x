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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SetPublishingModeRequest is the corresponding interface of SetPublishingModeRequest
type SetPublishingModeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetNoOfSubscriptionIds returns NoOfSubscriptionIds (property field)
	GetNoOfSubscriptionIds() int32
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
}

// SetPublishingModeRequestExactly can be used when we want exactly this type and not a type which fulfills SetPublishingModeRequest.
// This is useful for switch cases.
type SetPublishingModeRequestExactly interface {
	SetPublishingModeRequest
	isSetPublishingModeRequest() bool
}

// _SetPublishingModeRequest is the data-structure of this message
type _SetPublishingModeRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader       ExtensionObjectDefinition
	PublishingEnabled   bool
	NoOfSubscriptionIds int32
	SubscriptionIds     []uint32
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetPublishingModeRequest) GetIdentifier() string {
	return "799"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetPublishingModeRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SetPublishingModeRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetPublishingModeRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_SetPublishingModeRequest) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_SetPublishingModeRequest) GetNoOfSubscriptionIds() int32 {
	return m.NoOfSubscriptionIds
}

func (m *_SetPublishingModeRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSetPublishingModeRequest factory function for _SetPublishingModeRequest
func NewSetPublishingModeRequest(requestHeader ExtensionObjectDefinition, publishingEnabled bool, noOfSubscriptionIds int32, subscriptionIds []uint32) *_SetPublishingModeRequest {
	_result := &_SetPublishingModeRequest{
		RequestHeader:              requestHeader,
		PublishingEnabled:          publishingEnabled,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetPublishingModeRequest(structType any) SetPublishingModeRequest {
	if casted, ok := structType.(SetPublishingModeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetPublishingModeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetPublishingModeRequest) GetTypeName() string {
	return "SetPublishingModeRequest"
}

func (m *_SetPublishingModeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	return lengthInBits
}

func (m *_SetPublishingModeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetPublishingModeRequestParse(ctx context.Context, theBytes []byte, identifier string) (SetPublishingModeRequest, error) {
	return SetPublishingModeRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SetPublishingModeRequestParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (SetPublishingModeRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SetPublishingModeRequest, error) {
		return SetPublishingModeRequestParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func SetPublishingModeRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SetPublishingModeRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetPublishingModeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetPublishingModeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	publishingEnabled, err := ReadSimpleField(ctx, "publishingEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingEnabled' field"))
	}

	noOfSubscriptionIds, err := ReadSimpleField(ctx, "noOfSubscriptionIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSubscriptionIds' field"))
	}

	subscriptionIds, err := ReadCountArrayField[uint32](ctx, "subscriptionIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfSubscriptionIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionIds' field"))
	}

	if closeErr := readBuffer.CloseContext("SetPublishingModeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetPublishingModeRequest")
	}

	// Create a partially initialized instance
	_child := &_SetPublishingModeRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		PublishingEnabled:          publishingEnabled,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SetPublishingModeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetPublishingModeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetPublishingModeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetPublishingModeRequest")
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

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		// Simple Field (publishingEnabled)
		publishingEnabled := bool(m.GetPublishingEnabled())
		_publishingEnabledErr := /*TODO: migrate me*/ writeBuffer.WriteBit("publishingEnabled", (publishingEnabled))
		if _publishingEnabledErr != nil {
			return errors.Wrap(_publishingEnabledErr, "Error serializing 'publishingEnabled' field")
		}

		// Simple Field (noOfSubscriptionIds)
		noOfSubscriptionIds := int32(m.GetNoOfSubscriptionIds())
		_noOfSubscriptionIdsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfSubscriptionIds", 32, int32((noOfSubscriptionIds)))
		if _noOfSubscriptionIdsErr != nil {
			return errors.Wrap(_noOfSubscriptionIdsErr, "Error serializing 'noOfSubscriptionIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "subscriptionIds", m.GetSubscriptionIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionIds' field")
		}

		if popErr := writeBuffer.PopContext("SetPublishingModeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetPublishingModeRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetPublishingModeRequest) isSetPublishingModeRequest() bool {
	return true
}

func (m *_SetPublishingModeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
