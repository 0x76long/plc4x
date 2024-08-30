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

// BACnetConstructedDataReasonForDisable is the corresponding interface of BACnetConstructedDataReasonForDisable
type BACnetConstructedDataReasonForDisable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReasonForDisable returns ReasonForDisable (property field)
	GetReasonForDisable() []BACnetAccessCredentialDisableReasonTagged
}

// BACnetConstructedDataReasonForDisableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataReasonForDisable.
// This is useful for switch cases.
type BACnetConstructedDataReasonForDisableExactly interface {
	BACnetConstructedDataReasonForDisable
	isBACnetConstructedDataReasonForDisable() bool
}

// _BACnetConstructedDataReasonForDisable is the data-structure of this message
type _BACnetConstructedDataReasonForDisable struct {
	*_BACnetConstructedData
	ReasonForDisable []BACnetAccessCredentialDisableReasonTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReasonForDisable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReasonForDisable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REASON_FOR_DISABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReasonForDisable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataReasonForDisable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForDisable) GetReasonForDisable() []BACnetAccessCredentialDisableReasonTagged {
	return m.ReasonForDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReasonForDisable factory function for _BACnetConstructedDataReasonForDisable
func NewBACnetConstructedDataReasonForDisable(reasonForDisable []BACnetAccessCredentialDisableReasonTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReasonForDisable {
	_result := &_BACnetConstructedDataReasonForDisable{
		ReasonForDisable:       reasonForDisable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReasonForDisable(structType any) BACnetConstructedDataReasonForDisable {
	if casted, ok := structType.(BACnetConstructedDataReasonForDisable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReasonForDisable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReasonForDisable) GetTypeName() string {
	return "BACnetConstructedDataReasonForDisable"
}

func (m *_BACnetConstructedDataReasonForDisable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.ReasonForDisable) > 0 {
		for _, element := range m.ReasonForDisable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataReasonForDisable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataReasonForDisableParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReasonForDisable, error) {
	return BACnetConstructedDataReasonForDisableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataReasonForDisableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReasonForDisable, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReasonForDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReasonForDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reasonForDisable, err := ReadTerminatedArrayField[BACnetAccessCredentialDisableReasonTagged](ctx, "reasonForDisable", ReadComplex[BACnetAccessCredentialDisableReasonTagged](func(ctx context.Context, buffer utils.ReadBuffer) (BACnetAccessCredentialDisableReasonTagged, error) {
		v, err := BACnetAccessCredentialDisableReasonTaggedParseWithBuffer(ctx, readBuffer, (uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS))
		if err != nil {
			return nil, err
		}
		return v.(BACnetAccessCredentialDisableReasonTagged), nil
	}, readBuffer), func() bool { return IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber) })
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reasonForDisable' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReasonForDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReasonForDisable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataReasonForDisable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ReasonForDisable: reasonForDisable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataReasonForDisable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReasonForDisable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReasonForDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReasonForDisable")
		}

		// Array Field (reasonForDisable)
		if pushErr := writeBuffer.PushContext("reasonForDisable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reasonForDisable")
		}
		for _curItem, _element := range m.GetReasonForDisable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReasonForDisable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'reasonForDisable' field")
			}
		}
		if popErr := writeBuffer.PopContext("reasonForDisable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reasonForDisable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReasonForDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReasonForDisable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReasonForDisable) isBACnetConstructedDataReasonForDisable() bool {
	return true
}

func (m *_BACnetConstructedDataReasonForDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
