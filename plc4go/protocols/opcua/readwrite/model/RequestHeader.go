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

// RequestHeader is the corresponding interface of RequestHeader
type RequestHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAuthenticationToken returns AuthenticationToken (property field)
	GetAuthenticationToken() NodeId
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
	// GetReturnDiagnostics returns ReturnDiagnostics (property field)
	GetReturnDiagnostics() uint32
	// GetAuditEntryId returns AuditEntryId (property field)
	GetAuditEntryId() PascalString
	// GetTimeoutHint returns TimeoutHint (property field)
	GetTimeoutHint() uint32
	// GetAdditionalHeader returns AdditionalHeader (property field)
	GetAdditionalHeader() ExtensionObject
}

// RequestHeaderExactly can be used when we want exactly this type and not a type which fulfills RequestHeader.
// This is useful for switch cases.
type RequestHeaderExactly interface {
	RequestHeader
	isRequestHeader() bool
}

// _RequestHeader is the data-structure of this message
type _RequestHeader struct {
	*_ExtensionObjectDefinition
	AuthenticationToken NodeId
	Timestamp           int64
	RequestHandle       uint32
	ReturnDiagnostics   uint32
	AuditEntryId        PascalString
	TimeoutHint         uint32
	AdditionalHeader    ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RequestHeader) GetIdentifier() string {
	return "391"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestHeader) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RequestHeader) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestHeader) GetAuthenticationToken() NodeId {
	return m.AuthenticationToken
}

func (m *_RequestHeader) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_RequestHeader) GetRequestHandle() uint32 {
	return m.RequestHandle
}

func (m *_RequestHeader) GetReturnDiagnostics() uint32 {
	return m.ReturnDiagnostics
}

func (m *_RequestHeader) GetAuditEntryId() PascalString {
	return m.AuditEntryId
}

func (m *_RequestHeader) GetTimeoutHint() uint32 {
	return m.TimeoutHint
}

func (m *_RequestHeader) GetAdditionalHeader() ExtensionObject {
	return m.AdditionalHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestHeader factory function for _RequestHeader
func NewRequestHeader(authenticationToken NodeId, timestamp int64, requestHandle uint32, returnDiagnostics uint32, auditEntryId PascalString, timeoutHint uint32, additionalHeader ExtensionObject) *_RequestHeader {
	_result := &_RequestHeader{
		AuthenticationToken:        authenticationToken,
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ReturnDiagnostics:          returnDiagnostics,
		AuditEntryId:               auditEntryId,
		TimeoutHint:                timeoutHint,
		AdditionalHeader:           additionalHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestHeader(structType any) RequestHeader {
	if casted, ok := structType.(RequestHeader); ok {
		return casted
	}
	if casted, ok := structType.(*RequestHeader); ok {
		return *casted
	}
	return nil
}

func (m *_RequestHeader) GetTypeName() string {
	return "RequestHeader"
}

func (m *_RequestHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (authenticationToken)
	lengthInBits += m.AuthenticationToken.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (requestHandle)
	lengthInBits += 32

	// Simple field (returnDiagnostics)
	lengthInBits += 32

	// Simple field (auditEntryId)
	lengthInBits += m.AuditEntryId.GetLengthInBits(ctx)

	// Simple field (timeoutHint)
	lengthInBits += 32

	// Simple field (additionalHeader)
	lengthInBits += m.AdditionalHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RequestHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestHeaderParse(ctx context.Context, theBytes []byte, identifier string) (RequestHeader, error) {
	return RequestHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RequestHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RequestHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RequestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (authenticationToken)
	if pullErr := readBuffer.PullContext("authenticationToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationToken")
	}
	_authenticationToken, _authenticationTokenErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _authenticationTokenErr != nil {
		return nil, errors.Wrap(_authenticationTokenErr, "Error parsing 'authenticationToken' field of RequestHeader")
	}
	authenticationToken := _authenticationToken.(NodeId)
	if closeErr := readBuffer.CloseContext("authenticationToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationToken")
	}

	// Simple Field (timestamp)
	_timestamp, _timestampErr := /*TODO: migrate me*/ readBuffer.ReadInt64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of RequestHeader")
	}
	timestamp := _timestamp

	// Simple Field (requestHandle)
	_requestHandle, _requestHandleErr := /*TODO: migrate me*/ readBuffer.ReadUint32("requestHandle", 32)
	if _requestHandleErr != nil {
		return nil, errors.Wrap(_requestHandleErr, "Error parsing 'requestHandle' field of RequestHeader")
	}
	requestHandle := _requestHandle

	// Simple Field (returnDiagnostics)
	_returnDiagnostics, _returnDiagnosticsErr := /*TODO: migrate me*/ readBuffer.ReadUint32("returnDiagnostics", 32)
	if _returnDiagnosticsErr != nil {
		return nil, errors.Wrap(_returnDiagnosticsErr, "Error parsing 'returnDiagnostics' field of RequestHeader")
	}
	returnDiagnostics := _returnDiagnostics

	// Simple Field (auditEntryId)
	if pullErr := readBuffer.PullContext("auditEntryId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for auditEntryId")
	}
	_auditEntryId, _auditEntryIdErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _auditEntryIdErr != nil {
		return nil, errors.Wrap(_auditEntryIdErr, "Error parsing 'auditEntryId' field of RequestHeader")
	}
	auditEntryId := _auditEntryId.(PascalString)
	if closeErr := readBuffer.CloseContext("auditEntryId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for auditEntryId")
	}

	// Simple Field (timeoutHint)
	_timeoutHint, _timeoutHintErr := /*TODO: migrate me*/ readBuffer.ReadUint32("timeoutHint", 32)
	if _timeoutHintErr != nil {
		return nil, errors.Wrap(_timeoutHintErr, "Error parsing 'timeoutHint' field of RequestHeader")
	}
	timeoutHint := _timeoutHint

	// Simple Field (additionalHeader)
	if pullErr := readBuffer.PullContext("additionalHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalHeader")
	}
	_additionalHeader, _additionalHeaderErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _additionalHeaderErr != nil {
		return nil, errors.Wrap(_additionalHeaderErr, "Error parsing 'additionalHeader' field of RequestHeader")
	}
	additionalHeader := _additionalHeader.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("additionalHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalHeader")
	}

	if closeErr := readBuffer.CloseContext("RequestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestHeader")
	}

	// Create a partially initialized instance
	_child := &_RequestHeader{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		AuthenticationToken:        authenticationToken,
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ReturnDiagnostics:          returnDiagnostics,
		AuditEntryId:               auditEntryId,
		TimeoutHint:                timeoutHint,
		AdditionalHeader:           additionalHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RequestHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestHeader")
		}

		// Simple Field (authenticationToken)
		if pushErr := writeBuffer.PushContext("authenticationToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authenticationToken")
		}
		_authenticationTokenErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationToken())
		if popErr := writeBuffer.PopContext("authenticationToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authenticationToken")
		}
		if _authenticationTokenErr != nil {
			return errors.Wrap(_authenticationTokenErr, "Error serializing 'authenticationToken' field")
		}

		// Simple Field (timestamp)
		timestamp := int64(m.GetTimestamp())
		_timestampErr := /*TODO: migrate me*/ writeBuffer.WriteInt64("timestamp", 64, int64((timestamp)))
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (requestHandle)
		requestHandle := uint32(m.GetRequestHandle())
		_requestHandleErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("requestHandle", 32, uint32((requestHandle)))
		if _requestHandleErr != nil {
			return errors.Wrap(_requestHandleErr, "Error serializing 'requestHandle' field")
		}

		// Simple Field (returnDiagnostics)
		returnDiagnostics := uint32(m.GetReturnDiagnostics())
		_returnDiagnosticsErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("returnDiagnostics", 32, uint32((returnDiagnostics)))
		if _returnDiagnosticsErr != nil {
			return errors.Wrap(_returnDiagnosticsErr, "Error serializing 'returnDiagnostics' field")
		}

		// Simple Field (auditEntryId)
		if pushErr := writeBuffer.PushContext("auditEntryId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for auditEntryId")
		}
		_auditEntryIdErr := writeBuffer.WriteSerializable(ctx, m.GetAuditEntryId())
		if popErr := writeBuffer.PopContext("auditEntryId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for auditEntryId")
		}
		if _auditEntryIdErr != nil {
			return errors.Wrap(_auditEntryIdErr, "Error serializing 'auditEntryId' field")
		}

		// Simple Field (timeoutHint)
		timeoutHint := uint32(m.GetTimeoutHint())
		_timeoutHintErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("timeoutHint", 32, uint32((timeoutHint)))
		if _timeoutHintErr != nil {
			return errors.Wrap(_timeoutHintErr, "Error serializing 'timeoutHint' field")
		}

		// Simple Field (additionalHeader)
		if pushErr := writeBuffer.PushContext("additionalHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalHeader")
		}
		_additionalHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetAdditionalHeader())
		if popErr := writeBuffer.PopContext("additionalHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalHeader")
		}
		if _additionalHeaderErr != nil {
			return errors.Wrap(_additionalHeaderErr, "Error serializing 'additionalHeader' field")
		}

		if popErr := writeBuffer.PopContext("RequestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestHeader")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestHeader) isRequestHeader() bool {
	return true
}

func (m *_RequestHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
