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

// BACnetConstructedDataFaultParameters is the data-structure of this message
type BACnetConstructedDataFaultParameters struct {
	*BACnetConstructedData
	FaultParameters *BACnetFaultParameter

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataFaultParameters is the corresponding interface of BACnetConstructedDataFaultParameters
type IBACnetConstructedDataFaultParameters interface {
	IBACnetConstructedData
	// GetFaultParameters returns FaultParameters (property field)
	GetFaultParameters() *BACnetFaultParameter
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

func (m *BACnetConstructedDataFaultParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFaultParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFaultParameters) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFaultParameters) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFaultParameters) GetFaultParameters() *BACnetFaultParameter {
	return m.FaultParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFaultParameters factory function for BACnetConstructedDataFaultParameters
func NewBACnetConstructedDataFaultParameters(faultParameters *BACnetFaultParameter, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataFaultParameters {
	_result := &BACnetConstructedDataFaultParameters{
		FaultParameters:       faultParameters,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFaultParameters(structType interface{}) *BACnetConstructedDataFaultParameters {
	if casted, ok := structType.(BACnetConstructedDataFaultParameters); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultParameters); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFaultParameters(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFaultParameters(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFaultParameters) GetTypeName() string {
	return "BACnetConstructedDataFaultParameters"
}

func (m *BACnetConstructedDataFaultParameters) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFaultParameters) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultParameters)
	lengthInBits += m.FaultParameters.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataFaultParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFaultParametersParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataFaultParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultParameters"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultParameters)
	if pullErr := readBuffer.PullContext("faultParameters"); pullErr != nil {
		return nil, pullErr
	}
	_faultParameters, _faultParametersErr := BACnetFaultParameterParse(readBuffer)
	if _faultParametersErr != nil {
		return nil, errors.Wrap(_faultParametersErr, "Error parsing 'faultParameters' field")
	}
	faultParameters := CastBACnetFaultParameter(_faultParameters)
	if closeErr := readBuffer.CloseContext("faultParameters"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultParameters"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFaultParameters{
		FaultParameters:       CastBACnetFaultParameter(faultParameters),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFaultParameters) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultParameters"); pushErr != nil {
			return pushErr
		}

		// Simple Field (faultParameters)
		if pushErr := writeBuffer.PushContext("faultParameters"); pushErr != nil {
			return pushErr
		}
		_faultParametersErr := m.FaultParameters.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("faultParameters"); popErr != nil {
			return popErr
		}
		if _faultParametersErr != nil {
			return errors.Wrap(_faultParametersErr, "Error serializing 'faultParameters' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultParameters"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFaultParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
