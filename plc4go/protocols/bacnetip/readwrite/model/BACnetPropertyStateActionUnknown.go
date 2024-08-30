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

// BACnetPropertyStateActionUnknown is the corresponding interface of BACnetPropertyStateActionUnknown
type BACnetPropertyStateActionUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetUnknownValue returns UnknownValue (property field)
	GetUnknownValue() BACnetContextTagUnknown
}

// BACnetPropertyStateActionUnknownExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStateActionUnknown.
// This is useful for switch cases.
type BACnetPropertyStateActionUnknownExactly interface {
	BACnetPropertyStateActionUnknown
	isBACnetPropertyStateActionUnknown() bool
}

// _BACnetPropertyStateActionUnknown is the data-structure of this message
type _BACnetPropertyStateActionUnknown struct {
	*_BACnetPropertyStates
	UnknownValue BACnetContextTagUnknown
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStateActionUnknown) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStateActionUnknown) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStateActionUnknown) GetUnknownValue() BACnetContextTagUnknown {
	return m.UnknownValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStateActionUnknown factory function for _BACnetPropertyStateActionUnknown
func NewBACnetPropertyStateActionUnknown(unknownValue BACnetContextTagUnknown, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStateActionUnknown {
	_result := &_BACnetPropertyStateActionUnknown{
		UnknownValue:          unknownValue,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStateActionUnknown(structType any) BACnetPropertyStateActionUnknown {
	if casted, ok := structType.(BACnetPropertyStateActionUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStateActionUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStateActionUnknown) GetTypeName() string {
	return "BACnetPropertyStateActionUnknown"
}

func (m *_BACnetPropertyStateActionUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unknownValue)
	lengthInBits += m.UnknownValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStateActionUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStateActionUnknownParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStateActionUnknown, error) {
	return BACnetPropertyStateActionUnknownParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStateActionUnknownParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStateActionUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStateActionUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStateActionUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unknownValue)
	if pullErr := readBuffer.PullContext("unknownValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unknownValue")
	}
	_unknownValue, _unknownValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_UNKNOWN))
	if _unknownValueErr != nil {
		return nil, errors.Wrap(_unknownValueErr, "Error parsing 'unknownValue' field of BACnetPropertyStateActionUnknown")
	}
	unknownValue := _unknownValue.(BACnetContextTagUnknown)
	if closeErr := readBuffer.CloseContext("unknownValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unknownValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStateActionUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStateActionUnknown")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStateActionUnknown{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		UnknownValue:          unknownValue,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStateActionUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStateActionUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStateActionUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStateActionUnknown")
		}

		// Simple Field (unknownValue)
		if pushErr := writeBuffer.PushContext("unknownValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unknownValue")
		}
		_unknownValueErr := writeBuffer.WriteSerializable(ctx, m.GetUnknownValue())
		if popErr := writeBuffer.PopContext("unknownValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unknownValue")
		}
		if _unknownValueErr != nil {
			return errors.Wrap(_unknownValueErr, "Error serializing 'unknownValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStateActionUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStateActionUnknown")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStateActionUnknown) isBACnetPropertyStateActionUnknown() bool {
	return true
}

func (m *_BACnetPropertyStateActionUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
