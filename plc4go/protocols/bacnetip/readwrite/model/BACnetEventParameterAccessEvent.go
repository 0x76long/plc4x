/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventParameterAccessEvent is the data-structure of this message
type BACnetEventParameterAccessEvent struct {
	*BACnetEventParameter
	OpeningTag               *BACnetOpeningTag
	ListOfAccessEvents       *BACnetEventParameterAccessEventListOfAccessEvents
	AccessEventTimeReference *BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag               *BACnetClosingTag
}

// IBACnetEventParameterAccessEvent is the corresponding interface of BACnetEventParameterAccessEvent
type IBACnetEventParameterAccessEvent interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() *BACnetEventParameterAccessEventListOfAccessEvents
	// GetAccessEventTimeReference returns AccessEventTimeReference (property field)
	GetAccessEventTimeReference() *BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetEventParameterAccessEvent) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterAccessEvent) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterAccessEvent) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterAccessEvent) GetListOfAccessEvents() *BACnetEventParameterAccessEventListOfAccessEvents {
	return m.ListOfAccessEvents
}

func (m *BACnetEventParameterAccessEvent) GetAccessEventTimeReference() *BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.AccessEventTimeReference
}

func (m *BACnetEventParameterAccessEvent) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterAccessEvent factory function for BACnetEventParameterAccessEvent
func NewBACnetEventParameterAccessEvent(openingTag *BACnetOpeningTag, listOfAccessEvents *BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference *BACnetDeviceObjectPropertyReferenceEnclosed, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterAccessEvent {
	_result := &BACnetEventParameterAccessEvent{
		OpeningTag:               openingTag,
		ListOfAccessEvents:       listOfAccessEvents,
		AccessEventTimeReference: accessEventTimeReference,
		ClosingTag:               closingTag,
		BACnetEventParameter:     NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterAccessEvent(structType interface{}) *BACnetEventParameterAccessEvent {
	if casted, ok := structType.(BACnetEventParameterAccessEvent); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterAccessEvent(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterAccessEvent(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterAccessEvent) GetTypeName() string {
	return "BACnetEventParameterAccessEvent"
}

func (m *BACnetEventParameterAccessEvent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterAccessEvent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (listOfAccessEvents)
	lengthInBits += m.ListOfAccessEvents.GetLengthInBits()

	// Simple field (accessEventTimeReference)
	lengthInBits += m.AccessEventTimeReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventParameterAccessEvent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterAccessEventParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEvent"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(13)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfAccessEvents)
	if pullErr := readBuffer.PullContext("listOfAccessEvents"); pullErr != nil {
		return nil, pullErr
	}
	_listOfAccessEvents, _listOfAccessEventsErr := BACnetEventParameterAccessEventListOfAccessEventsParse(readBuffer, uint8(uint8(0)))
	if _listOfAccessEventsErr != nil {
		return nil, errors.Wrap(_listOfAccessEventsErr, "Error parsing 'listOfAccessEvents' field")
	}
	listOfAccessEvents := CastBACnetEventParameterAccessEventListOfAccessEvents(_listOfAccessEvents)
	if closeErr := readBuffer.CloseContext("listOfAccessEvents"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (accessEventTimeReference)
	if pullErr := readBuffer.PullContext("accessEventTimeReference"); pullErr != nil {
		return nil, pullErr
	}
	_accessEventTimeReference, _accessEventTimeReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _accessEventTimeReferenceErr != nil {
		return nil, errors.Wrap(_accessEventTimeReferenceErr, "Error parsing 'accessEventTimeReference' field")
	}
	accessEventTimeReference := CastBACnetDeviceObjectPropertyReferenceEnclosed(_accessEventTimeReference)
	if closeErr := readBuffer.CloseContext("accessEventTimeReference"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(13)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEvent"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterAccessEvent{
		OpeningTag:               CastBACnetOpeningTag(openingTag),
		ListOfAccessEvents:       CastBACnetEventParameterAccessEventListOfAccessEvents(listOfAccessEvents),
		AccessEventTimeReference: CastBACnetDeviceObjectPropertyReferenceEnclosed(accessEventTimeReference),
		ClosingTag:               CastBACnetClosingTag(closingTag),
		BACnetEventParameter:     &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterAccessEvent) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEvent"); pushErr != nil {
			return pushErr
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return pushErr
		}
		_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return popErr
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (listOfAccessEvents)
		if pushErr := writeBuffer.PushContext("listOfAccessEvents"); pushErr != nil {
			return pushErr
		}
		_listOfAccessEventsErr := m.ListOfAccessEvents.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfAccessEvents"); popErr != nil {
			return popErr
		}
		if _listOfAccessEventsErr != nil {
			return errors.Wrap(_listOfAccessEventsErr, "Error serializing 'listOfAccessEvents' field")
		}

		// Simple Field (accessEventTimeReference)
		if pushErr := writeBuffer.PushContext("accessEventTimeReference"); pushErr != nil {
			return pushErr
		}
		_accessEventTimeReferenceErr := m.AccessEventTimeReference.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accessEventTimeReference"); popErr != nil {
			return popErr
		}
		if _accessEventTimeReferenceErr != nil {
			return errors.Wrap(_accessEventTimeReferenceErr, "Error serializing 'accessEventTimeReference' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return pushErr
		}
		_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return popErr
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEvent"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
