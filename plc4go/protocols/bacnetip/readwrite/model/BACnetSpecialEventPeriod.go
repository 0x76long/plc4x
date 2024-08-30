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

// BACnetSpecialEventPeriod is the corresponding interface of BACnetSpecialEventPeriod
type BACnetSpecialEventPeriod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetSpecialEventPeriodExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEventPeriod.
// This is useful for switch cases.
type BACnetSpecialEventPeriodExactly interface {
	BACnetSpecialEventPeriod
	isBACnetSpecialEventPeriod() bool
}

// _BACnetSpecialEventPeriod is the data-structure of this message
type _BACnetSpecialEventPeriod struct {
	_BACnetSpecialEventPeriodChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetSpecialEventPeriodChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetSpecialEventPeriodParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetSpecialEventPeriod, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetSpecialEventPeriodChild interface {
	utils.Serializable
	InitializeParent(parent BACnetSpecialEventPeriod, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetSpecialEventPeriod

	GetTypeName() string
	BACnetSpecialEventPeriod
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventPeriod) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetSpecialEventPeriod) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriod factory function for _BACnetSpecialEventPeriod
func NewBACnetSpecialEventPeriod(peekedTagHeader BACnetTagHeader) *_BACnetSpecialEventPeriod {
	return &_BACnetSpecialEventPeriod{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventPeriod(structType any) BACnetSpecialEventPeriod {
	if casted, ok := structType.(BACnetSpecialEventPeriod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventPeriod) GetTypeName() string {
	return "BACnetSpecialEventPeriod"
}

func (m *_BACnetSpecialEventPeriod) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetSpecialEventPeriod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventPeriodParse(ctx context.Context, theBytes []byte) (BACnetSpecialEventPeriod, error) {
	return BACnetSpecialEventPeriodParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSpecialEventPeriodParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventPeriod, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "Validation failed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetSpecialEventPeriodChildSerializeRequirement interface {
		BACnetSpecialEventPeriod
		InitializeParent(BACnetSpecialEventPeriod, BACnetTagHeader)
		GetParent() BACnetSpecialEventPeriod
	}
	var _childTemp any
	var _child BACnetSpecialEventPeriodChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetSpecialEventPeriodCalendarEntry
		_childTemp, typeSwitchError = BACnetSpecialEventPeriodCalendarEntryParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetSpecialEventPeriodCalendarReference
		_childTemp, typeSwitchError = BACnetSpecialEventPeriodCalendarReferenceParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetSpecialEventPeriod")
	}
	_child = _childTemp.(BACnetSpecialEventPeriodChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriod")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetSpecialEventPeriod) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetSpecialEventPeriod, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriod"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriod")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriod"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriod")
	}
	return nil
}

func (m *_BACnetSpecialEventPeriod) isBACnetSpecialEventPeriod() bool {
	return true
}

func (m *_BACnetSpecialEventPeriod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
