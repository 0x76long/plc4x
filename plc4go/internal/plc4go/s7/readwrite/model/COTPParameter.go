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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type COTPParameter struct {

	// Arguments.
	Rest  uint8
	Child ICOTPParameterChild
}

// The corresponding interface
type ICOTPParameter interface {
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICOTPParameterParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICOTPParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICOTPParameterChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *COTPParameter)
	GetParent() *COTPParameter

	GetTypeName() string
	ICOTPParameter
}

// NewCOTPParameter factory function for COTPParameter
func NewCOTPParameter(rest uint8) *COTPParameter {
	return &COTPParameter{Rest: rest}
}

func CastCOTPParameter(structType interface{}) *COTPParameter {
	if casted, ok := structType.(COTPParameter); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return casted
	}
	return nil
}

func (m *COTPParameter) GetTypeName() string {
	return "COTPParameter"
}

func (m *COTPParameter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPParameter) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *COTPParameter) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	// Implicit Field (parameterLength)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPParameter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterParse(readBuffer utils.ReadBuffer, rest uint8) (*COTPParameter, error) {
	if pullErr := readBuffer.PullContext("COTPParameter"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType, _parameterTypeErr := readBuffer.ReadUint8("parameterType", 8)
	if _parameterTypeErr != nil {
		return nil, errors.Wrap(_parameterTypeErr, "Error parsing 'parameterType' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	parameterLength, _parameterLengthErr := readBuffer.ReadUint8("parameterLength", 8)
	_ = parameterLength
	if _parameterLengthErr != nil {
		return nil, errors.Wrap(_parameterLengthErr, "Error parsing 'parameterLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *COTPParameter
	var typeSwitchError error
	switch {
	case parameterType == 0xC0: // COTPParameterTpduSize
		_parent, typeSwitchError = COTPParameterTpduSizeParse(readBuffer, rest)
	case parameterType == 0xC1: // COTPParameterCallingTsap
		_parent, typeSwitchError = COTPParameterCallingTsapParse(readBuffer, rest)
	case parameterType == 0xC2: // COTPParameterCalledTsap
		_parent, typeSwitchError = COTPParameterCalledTsapParse(readBuffer, rest)
	case parameterType == 0xC3: // COTPParameterChecksum
		_parent, typeSwitchError = COTPParameterChecksumParse(readBuffer, rest)
	case parameterType == 0xE0: // COTPParameterDisconnectAdditionalInformation
		_parent, typeSwitchError = COTPParameterDisconnectAdditionalInformationParse(readBuffer, rest)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("COTPParameter"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *COTPParameter) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *COTPParameter) SerializeParent(writeBuffer utils.WriteBuffer, child ICOTPParameter, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("COTPParameter"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType := uint8(child.GetParameterType())
	_parameterTypeErr := writeBuffer.WriteUint8("parameterType", 8, (parameterType))

	if _parameterTypeErr != nil {
		return errors.Wrap(_parameterTypeErr, "Error serializing 'parameterType' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength := uint8(uint8(uint8(m.GetLengthInBytes())) - uint8(uint8(2)))
	_parameterLengthErr := writeBuffer.WriteUint8("parameterLength", 8, (parameterLength))
	if _parameterLengthErr != nil {
		return errors.Wrap(_parameterLengthErr, "Error serializing 'parameterLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("COTPParameter"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *COTPParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
