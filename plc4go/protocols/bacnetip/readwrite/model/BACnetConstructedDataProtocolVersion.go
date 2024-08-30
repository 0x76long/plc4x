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

// BACnetConstructedDataProtocolVersion is the corresponding interface of BACnetConstructedDataProtocolVersion
type BACnetConstructedDataProtocolVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProtocolVersion returns ProtocolVersion (property field)
	GetProtocolVersion() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataProtocolVersionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProtocolVersion.
// This is useful for switch cases.
type BACnetConstructedDataProtocolVersionExactly interface {
	BACnetConstructedDataProtocolVersion
	isBACnetConstructedDataProtocolVersion() bool
}

// _BACnetConstructedDataProtocolVersion is the data-structure of this message
type _BACnetConstructedDataProtocolVersion struct {
	*_BACnetConstructedData
	ProtocolVersion BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolVersion) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolVersion) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProtocolVersion) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetProtocolVersion() BACnetApplicationTagUnsignedInteger {
	return m.ProtocolVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolVersion) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetProtocolVersion())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolVersion factory function for _BACnetConstructedDataProtocolVersion
func NewBACnetConstructedDataProtocolVersion(protocolVersion BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolVersion {
	_result := &_BACnetConstructedDataProtocolVersion{
		ProtocolVersion:        protocolVersion,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolVersion(structType any) BACnetConstructedDataProtocolVersion {
	if casted, ok := structType.(BACnetConstructedDataProtocolVersion); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolVersion); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolVersion) GetTypeName() string {
	return "BACnetConstructedDataProtocolVersion"
}

func (m *_BACnetConstructedDataProtocolVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (protocolVersion)
	lengthInBits += m.ProtocolVersion.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataProtocolVersionParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolVersion, error) {
	return BACnetConstructedDataProtocolVersionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataProtocolVersionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolVersion, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolVersion)
	if pullErr := readBuffer.PullContext("protocolVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for protocolVersion")
	}
	_protocolVersion, _protocolVersionErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _protocolVersionErr != nil {
		return nil, errors.Wrap(_protocolVersionErr, "Error parsing 'protocolVersion' field of BACnetConstructedDataProtocolVersion")
	}
	protocolVersion := _protocolVersion.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("protocolVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for protocolVersion")
	}

	// Virtual field
	_actualValue := protocolVersion
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolVersion")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProtocolVersion{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProtocolVersion: protocolVersion,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProtocolVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolVersion")
		}

		// Simple Field (protocolVersion)
		if pushErr := writeBuffer.PushContext("protocolVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for protocolVersion")
		}
		_protocolVersionErr := writeBuffer.WriteSerializable(ctx, m.GetProtocolVersion())
		if popErr := writeBuffer.PopContext("protocolVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for protocolVersion")
		}
		if _protocolVersionErr != nil {
			return errors.Wrap(_protocolVersionErr, "Error serializing 'protocolVersion' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolVersion")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolVersion) isBACnetConstructedDataProtocolVersion() bool {
	return true
}

func (m *_BACnetConstructedDataProtocolVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
