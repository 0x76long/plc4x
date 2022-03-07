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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7Message_PROTOCOLID uint8 = 0x32

// The data-structure of this message
type S7Message struct {
	TpduReference uint16
	Parameter     *S7Parameter
	Payload       *S7Payload
	Child         IS7MessageChild
}

// The corresponding interface
type IS7Message interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetTpduReference returns TpduReference (property field)
	GetTpduReference() uint16
	// GetParameter returns Parameter (property field)
	GetParameter() *S7Parameter
	// GetPayload returns Payload (property field)
	GetPayload() *S7Payload
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7MessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Message, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7MessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Message, tpduReference uint16, parameter *S7Parameter, payload *S7Payload)
	GetParent() *S7Message

	GetTypeName() string
	IS7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *S7Message) GetTpduReference() uint16 {
	return m.TpduReference
}

func (m *S7Message) GetParameter() *S7Parameter {
	return m.Parameter
}

func (m *S7Message) GetPayload() *S7Payload {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7Message factory function for S7Message
func NewS7Message(tpduReference uint16, parameter *S7Parameter, payload *S7Payload) *S7Message {
	return &S7Message{TpduReference: tpduReference, Parameter: parameter, Payload: payload}
}

func CastS7Message(structType interface{}) *S7Message {
	if casted, ok := structType.(S7Message); ok {
		return &casted
	}
	if casted, ok := structType.(*S7Message); ok {
		return casted
	}
	return nil
}

func (m *S7Message) GetTypeName() string {
	return "S7Message"
}

func (m *S7Message) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7Message) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7Message) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (tpduReference)
	lengthInBits += 16

	// Implicit Field (parameterLength)
	lengthInBits += 16

	// Implicit Field (payloadLength)
	lengthInBits += 16

	// Optional Field (parameter)
	if m.Parameter != nil {
		lengthInBits += (*m.Parameter).GetLengthInBits()
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += (*m.Payload).GetLengthInBits()
	}

	return lengthInBits
}

func (m *S7Message) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7MessageParse(readBuffer utils.ReadBuffer) (*S7Message, error) {
	if pullErr := readBuffer.PullContext("S7Message"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (protocolId)
	protocolId, _protocolIdErr := readBuffer.ReadUint8("protocolId", 8)
	if _protocolIdErr != nil {
		return nil, errors.Wrap(_protocolIdErr, "Error parsing 'protocolId' field")
	}
	if protocolId != S7Message_PROTOCOLID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7Message_PROTOCOLID) + " but got " + fmt.Sprintf("%d", protocolId))
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (tpduReference)
	_tpduReference, _tpduReferenceErr := readBuffer.ReadUint16("tpduReference", 16)
	if _tpduReferenceErr != nil {
		return nil, errors.Wrap(_tpduReferenceErr, "Error parsing 'tpduReference' field")
	}
	tpduReference := _tpduReference

	// Implicit Field (parameterLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	parameterLength, _parameterLengthErr := readBuffer.ReadUint16("parameterLength", 16)
	_ = parameterLength
	if _parameterLengthErr != nil {
		return nil, errors.Wrap(_parameterLengthErr, "Error parsing 'parameterLength' field")
	}

	// Implicit Field (payloadLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	payloadLength, _payloadLengthErr := readBuffer.ReadUint16("payloadLength", 16)
	_ = payloadLength
	if _payloadLengthErr != nil {
		return nil, errors.Wrap(_payloadLengthErr, "Error parsing 'payloadLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7Message
	var typeSwitchError error
	switch {
	case messageType == 0x01: // S7MessageRequest
		_parent, typeSwitchError = S7MessageRequestParse(readBuffer)
	case messageType == 0x02: // S7MessageResponse
		_parent, typeSwitchError = S7MessageResponseParse(readBuffer)
	case messageType == 0x03: // S7MessageResponseData
		_parent, typeSwitchError = S7MessageResponseDataParse(readBuffer)
	case messageType == 0x07: // S7MessageUserData
		_parent, typeSwitchError = S7MessageUserDataParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Optional Field (parameter) (Can be skipped, if a given expression evaluates to false)
	var parameter *S7Parameter = nil
	if bool((parameterLength) > (0)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("parameter"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := S7ParameterParse(readBuffer, messageType)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'parameter' field")
		default:
			parameter = CastS7Parameter(_val)
			if closeErr := readBuffer.CloseContext("parameter"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	var payload *S7Payload = nil
	if bool((payloadLength) > (0)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := S7PayloadParse(readBuffer, messageType, (parameter))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'payload' field")
		default:
			payload = CastS7Payload(_val)
			if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("S7Message"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, tpduReference, parameter, payload)
	return _parent, nil
}

func (m *S7Message) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Message) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Message, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Message"); pushErr != nil {
		return pushErr
	}

	// Const Field (protocolId)
	_protocolIdErr := writeBuffer.WriteUint8("protocolId", 8, 0x32)
	if _protocolIdErr != nil {
		return errors.Wrap(_protocolIdErr, "Error serializing 'protocolId' field")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (tpduReference)
	tpduReference := uint16(m.TpduReference)
	_tpduReferenceErr := writeBuffer.WriteUint16("tpduReference", 16, (tpduReference))
	if _tpduReferenceErr != nil {
		return errors.Wrap(_tpduReferenceErr, "Error serializing 'tpduReference' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength := uint16(utils.InlineIf(bool((m.GetParameter()) != (nil)), func() interface{} { return uint16((*m.GetParameter()).GetLengthInBytes()) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	_parameterLengthErr := writeBuffer.WriteUint16("parameterLength", 16, (parameterLength))
	if _parameterLengthErr != nil {
		return errors.Wrap(_parameterLengthErr, "Error serializing 'parameterLength' field")
	}

	// Implicit Field (payloadLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	payloadLength := uint16(utils.InlineIf(bool((m.GetPayload()) != (nil)), func() interface{} { return uint16((*m.GetPayload()).GetLengthInBytes()) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	_payloadLengthErr := writeBuffer.WriteUint16("payloadLength", 16, (payloadLength))
	if _payloadLengthErr != nil {
		return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (parameter) (Can be skipped, if the value is null)
	var parameter *S7Parameter = nil
	if m.Parameter != nil {
		if pushErr := writeBuffer.PushContext("parameter"); pushErr != nil {
			return pushErr
		}
		parameter = m.Parameter
		_parameterErr := parameter.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("parameter"); popErr != nil {
			return popErr
		}
		if _parameterErr != nil {
			return errors.Wrap(_parameterErr, "Error serializing 'parameter' field")
		}
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload *S7Payload = nil
	if m.Payload != nil {
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		payload = m.Payload
		_payloadErr := payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := writeBuffer.PopContext("S7Message"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7Message) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
