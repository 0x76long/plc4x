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
type CALData struct {
	CommandTypeContainer CALCommandTypeContainer
	Child                ICALDataChild
}

// The corresponding interface
type ICALData interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() CALCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() CALCommandType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICALDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICALData, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICALDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer)
	GetParent() *CALData

	GetTypeName() string
	ICALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *CALData) GetCommandTypeContainer() CALCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *CALData) GetCommandType() CALCommandType {
	return m.GetCommandTypeContainer().CommandType()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALData factory function for CALData
func NewCALData(commandTypeContainer CALCommandTypeContainer) *CALData {
	return &CALData{CommandTypeContainer: commandTypeContainer}
}

func CastCALData(structType interface{}) *CALData {
	if casted, ok := structType.(CALData); ok {
		return &casted
	}
	if casted, ok := structType.(*CALData); ok {
		return casted
	}
	return nil
}

func (m *CALData) GetTypeName() string {
	return "CALData"
}

func (m *CALData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALData) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CALData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *CALData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataParse(readBuffer utils.ReadBuffer) (*CALData, error) {
	if pullErr := readBuffer.PullContext("CALData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, pullErr
	}
	_commandTypeContainer, _commandTypeContainerErr := CALCommandTypeContainerParse(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := CALCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *CALData
	var typeSwitchError error
	switch {
	case commandType == CALCommandType_RESET: // CALDataRequestReset
		_parent, typeSwitchError = CALDataRequestResetParse(readBuffer)
	case commandType == CALCommandType_RECALL: // CALDataRequestRecall
		_parent, typeSwitchError = CALDataRequestRecallParse(readBuffer)
	case commandType == CALCommandType_IDENTIFY: // CALDataRequestIdentify
		_parent, typeSwitchError = CALDataRequestIdentifyParse(readBuffer)
	case commandType == CALCommandType_GET_STATUS: // CALDataRequestGetStatus
		_parent, typeSwitchError = CALDataRequestGetStatusParse(readBuffer)
	case commandType == CALCommandType_REPLY: // CALDataReplyReply
		_parent, typeSwitchError = CALDataReplyReplyParse(readBuffer, commandTypeContainer)
	case commandType == CALCommandType_ACKNOWLEDGE: // CALDataReplyAcknowledge
		_parent, typeSwitchError = CALDataReplyAcknowledgeParse(readBuffer)
	case commandType == CALCommandType_STATUS: // CALDataReplyStatus
		_parent, typeSwitchError = CALDataReplyStatusParse(readBuffer, commandTypeContainer)
	case commandType == CALCommandType_STATUS_EXTENDED: // CALDataReplyStatusExtended
		_parent, typeSwitchError = CALDataReplyStatusExtendedParse(readBuffer, commandTypeContainer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("CALData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, commandTypeContainer)
	return _parent, nil
}

func (m *CALData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CALData) SerializeParent(writeBuffer utils.WriteBuffer, child ICALData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("CALData"); pushErr != nil {
		return pushErr
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return pushErr
	}
	_commandTypeContainerErr := m.CommandTypeContainer.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return popErr
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual("commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CALData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CALData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
