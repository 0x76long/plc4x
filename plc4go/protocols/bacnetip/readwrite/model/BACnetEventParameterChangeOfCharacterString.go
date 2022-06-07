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

// BACnetEventParameterChangeOfCharacterString is the data-structure of this message
type BACnetEventParameterChangeOfCharacterString struct {
	*BACnetEventParameter
	OpeningTag        *BACnetOpeningTag
	TimeDelay         *BACnetContextTagUnsignedInteger
	ListOfAlarmValues *BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	ClosingTag        *BACnetClosingTag
}

// IBACnetEventParameterChangeOfCharacterString is the corresponding interface of BACnetEventParameterChangeOfCharacterString
type IBACnetEventParameterChangeOfCharacterString interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() *BACnetContextTagUnsignedInteger
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() *BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
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

func (m *BACnetEventParameterChangeOfCharacterString) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterChangeOfCharacterString) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterChangeOfCharacterString) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterChangeOfCharacterString) GetTimeDelay() *BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *BACnetEventParameterChangeOfCharacterString) GetListOfAlarmValues() *BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *BACnetEventParameterChangeOfCharacterString) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfCharacterString factory function for BACnetEventParameterChangeOfCharacterString
func NewBACnetEventParameterChangeOfCharacterString(openingTag *BACnetOpeningTag, timeDelay *BACnetContextTagUnsignedInteger, listOfAlarmValues *BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterChangeOfCharacterString {
	_result := &BACnetEventParameterChangeOfCharacterString{
		OpeningTag:           openingTag,
		TimeDelay:            timeDelay,
		ListOfAlarmValues:    listOfAlarmValues,
		ClosingTag:           closingTag,
		BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterChangeOfCharacterString(structType interface{}) *BACnetEventParameterChangeOfCharacterString {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterChangeOfCharacterString(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterChangeOfCharacterString(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterChangeOfCharacterString) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterString"
}

func (m *BACnetEventParameterChangeOfCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterChangeOfCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventParameterChangeOfCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfCharacterStringParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterChangeOfCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(17)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, pullErr
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := CastBACnetContextTagUnsignedInteger(_timeDelay)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfAlarmValues)
	if pullErr := readBuffer.PullContext("listOfAlarmValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfAlarmValues, _listOfAlarmValuesErr := BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParse(readBuffer, uint8(uint8(1)))
	if _listOfAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfAlarmValuesErr, "Error parsing 'listOfAlarmValues' field")
	}
	listOfAlarmValues := CastBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(_listOfAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfAlarmValues"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(17)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterChangeOfCharacterString{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		TimeDelay:            CastBACnetContextTagUnsignedInteger(timeDelay),
		ListOfAlarmValues:    CastBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(listOfAlarmValues),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetEventParameter: &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterChangeOfCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterString"); pushErr != nil {
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return pushErr
		}
		_timeDelayErr := m.TimeDelay.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return popErr
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (listOfAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfAlarmValues"); pushErr != nil {
			return pushErr
		}
		_listOfAlarmValuesErr := m.ListOfAlarmValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfAlarmValues"); popErr != nil {
			return popErr
		}
		if _listOfAlarmValuesErr != nil {
			return errors.Wrap(_listOfAlarmValuesErr, "Error serializing 'listOfAlarmValues' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterChangeOfCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
