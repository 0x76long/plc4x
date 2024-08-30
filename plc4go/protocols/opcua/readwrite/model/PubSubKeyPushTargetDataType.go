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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// PubSubKeyPushTargetDataType is the corresponding interface of PubSubKeyPushTargetDataType
type PubSubKeyPushTargetDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetApplicationUri returns ApplicationUri (property field)
	GetApplicationUri() PascalString
	// GetNoOfPushTargetFolder returns NoOfPushTargetFolder (property field)
	GetNoOfPushTargetFolder() int32
	// GetPushTargetFolder returns PushTargetFolder (property field)
	GetPushTargetFolder() []PascalString
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetUserTokenType returns UserTokenType (property field)
	GetUserTokenType() ExtensionObjectDefinition
	// GetRequestedKeyCount returns RequestedKeyCount (property field)
	GetRequestedKeyCount() uint16
	// GetRetryInterval returns RetryInterval (property field)
	GetRetryInterval() float64
	// GetNoOfPushTargetProperties returns NoOfPushTargetProperties (property field)
	GetNoOfPushTargetProperties() int32
	// GetPushTargetProperties returns PushTargetProperties (property field)
	GetPushTargetProperties() []ExtensionObjectDefinition
	// GetNoOfSecurityGroups returns NoOfSecurityGroups (property field)
	GetNoOfSecurityGroups() int32
	// GetSecurityGroups returns SecurityGroups (property field)
	GetSecurityGroups() []PascalString
}

// PubSubKeyPushTargetDataTypeExactly can be used when we want exactly this type and not a type which fulfills PubSubKeyPushTargetDataType.
// This is useful for switch cases.
type PubSubKeyPushTargetDataTypeExactly interface {
	PubSubKeyPushTargetDataType
	isPubSubKeyPushTargetDataType() bool
}

// _PubSubKeyPushTargetDataType is the data-structure of this message
type _PubSubKeyPushTargetDataType struct {
	*_ExtensionObjectDefinition
	ApplicationUri           PascalString
	NoOfPushTargetFolder     int32
	PushTargetFolder         []PascalString
	EndpointUrl              PascalString
	SecurityPolicyUri        PascalString
	UserTokenType            ExtensionObjectDefinition
	RequestedKeyCount        uint16
	RetryInterval            float64
	NoOfPushTargetProperties int32
	PushTargetProperties     []ExtensionObjectDefinition
	NoOfSecurityGroups       int32
	SecurityGroups           []PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubKeyPushTargetDataType) GetIdentifier() string {
	return "25272"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubKeyPushTargetDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PubSubKeyPushTargetDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubKeyPushTargetDataType) GetApplicationUri() PascalString {
	return m.ApplicationUri
}

func (m *_PubSubKeyPushTargetDataType) GetNoOfPushTargetFolder() int32 {
	return m.NoOfPushTargetFolder
}

func (m *_PubSubKeyPushTargetDataType) GetPushTargetFolder() []PascalString {
	return m.PushTargetFolder
}

func (m *_PubSubKeyPushTargetDataType) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_PubSubKeyPushTargetDataType) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_PubSubKeyPushTargetDataType) GetUserTokenType() ExtensionObjectDefinition {
	return m.UserTokenType
}

func (m *_PubSubKeyPushTargetDataType) GetRequestedKeyCount() uint16 {
	return m.RequestedKeyCount
}

func (m *_PubSubKeyPushTargetDataType) GetRetryInterval() float64 {
	return m.RetryInterval
}

func (m *_PubSubKeyPushTargetDataType) GetNoOfPushTargetProperties() int32 {
	return m.NoOfPushTargetProperties
}

func (m *_PubSubKeyPushTargetDataType) GetPushTargetProperties() []ExtensionObjectDefinition {
	return m.PushTargetProperties
}

func (m *_PubSubKeyPushTargetDataType) GetNoOfSecurityGroups() int32 {
	return m.NoOfSecurityGroups
}

func (m *_PubSubKeyPushTargetDataType) GetSecurityGroups() []PascalString {
	return m.SecurityGroups
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPubSubKeyPushTargetDataType factory function for _PubSubKeyPushTargetDataType
func NewPubSubKeyPushTargetDataType(applicationUri PascalString, noOfPushTargetFolder int32, pushTargetFolder []PascalString, endpointUrl PascalString, securityPolicyUri PascalString, userTokenType ExtensionObjectDefinition, requestedKeyCount uint16, retryInterval float64, noOfPushTargetProperties int32, pushTargetProperties []ExtensionObjectDefinition, noOfSecurityGroups int32, securityGroups []PascalString) *_PubSubKeyPushTargetDataType {
	_result := &_PubSubKeyPushTargetDataType{
		ApplicationUri:             applicationUri,
		NoOfPushTargetFolder:       noOfPushTargetFolder,
		PushTargetFolder:           pushTargetFolder,
		EndpointUrl:                endpointUrl,
		SecurityPolicyUri:          securityPolicyUri,
		UserTokenType:              userTokenType,
		RequestedKeyCount:          requestedKeyCount,
		RetryInterval:              retryInterval,
		NoOfPushTargetProperties:   noOfPushTargetProperties,
		PushTargetProperties:       pushTargetProperties,
		NoOfSecurityGroups:         noOfSecurityGroups,
		SecurityGroups:             securityGroups,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPubSubKeyPushTargetDataType(structType any) PubSubKeyPushTargetDataType {
	if casted, ok := structType.(PubSubKeyPushTargetDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubKeyPushTargetDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubKeyPushTargetDataType) GetTypeName() string {
	return "PubSubKeyPushTargetDataType"
}

func (m *_PubSubKeyPushTargetDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (applicationUri)
	lengthInBits += m.ApplicationUri.GetLengthInBits(ctx)

	// Simple field (noOfPushTargetFolder)
	lengthInBits += 32

	// Array field
	if len(m.PushTargetFolder) > 0 {
		for _curItem, element := range m.PushTargetFolder {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PushTargetFolder), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (userTokenType)
	lengthInBits += m.UserTokenType.GetLengthInBits(ctx)

	// Simple field (requestedKeyCount)
	lengthInBits += 16

	// Simple field (retryInterval)
	lengthInBits += 64

	// Simple field (noOfPushTargetProperties)
	lengthInBits += 32

	// Array field
	if len(m.PushTargetProperties) > 0 {
		for _curItem, element := range m.PushTargetProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PushTargetProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfSecurityGroups)
	lengthInBits += 32

	// Array field
	if len(m.SecurityGroups) > 0 {
		for _curItem, element := range m.SecurityGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PubSubKeyPushTargetDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubKeyPushTargetDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PubSubKeyPushTargetDataType, error) {
	return PubSubKeyPushTargetDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PubSubKeyPushTargetDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PubSubKeyPushTargetDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PubSubKeyPushTargetDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubKeyPushTargetDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (applicationUri)
	if pullErr := readBuffer.PullContext("applicationUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for applicationUri")
	}
	_applicationUri, _applicationUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _applicationUriErr != nil {
		return nil, errors.Wrap(_applicationUriErr, "Error parsing 'applicationUri' field of PubSubKeyPushTargetDataType")
	}
	applicationUri := _applicationUri.(PascalString)
	if closeErr := readBuffer.CloseContext("applicationUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for applicationUri")
	}

	// Simple Field (noOfPushTargetFolder)
	_noOfPushTargetFolder, _noOfPushTargetFolderErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfPushTargetFolder", 32)
	if _noOfPushTargetFolderErr != nil {
		return nil, errors.Wrap(_noOfPushTargetFolderErr, "Error parsing 'noOfPushTargetFolder' field of PubSubKeyPushTargetDataType")
	}
	noOfPushTargetFolder := _noOfPushTargetFolder

	pushTargetFolder, err := ReadCountArrayField[PascalString](ctx, "pushTargetFolder", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfPushTargetFolder))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pushTargetFolder' field"))
	}

	// Simple Field (endpointUrl)
	if pullErr := readBuffer.PullContext("endpointUrl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endpointUrl")
	}
	_endpointUrl, _endpointUrlErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _endpointUrlErr != nil {
		return nil, errors.Wrap(_endpointUrlErr, "Error parsing 'endpointUrl' field of PubSubKeyPushTargetDataType")
	}
	endpointUrl := _endpointUrl.(PascalString)
	if closeErr := readBuffer.CloseContext("endpointUrl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endpointUrl")
	}

	// Simple Field (securityPolicyUri)
	if pullErr := readBuffer.PullContext("securityPolicyUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPolicyUri")
	}
	_securityPolicyUri, _securityPolicyUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _securityPolicyUriErr != nil {
		return nil, errors.Wrap(_securityPolicyUriErr, "Error parsing 'securityPolicyUri' field of PubSubKeyPushTargetDataType")
	}
	securityPolicyUri := _securityPolicyUri.(PascalString)
	if closeErr := readBuffer.CloseContext("securityPolicyUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPolicyUri")
	}

	// Simple Field (userTokenType)
	if pullErr := readBuffer.PullContext("userTokenType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userTokenType")
	}
	_userTokenType, _userTokenTypeErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("306"))
	if _userTokenTypeErr != nil {
		return nil, errors.Wrap(_userTokenTypeErr, "Error parsing 'userTokenType' field of PubSubKeyPushTargetDataType")
	}
	userTokenType := _userTokenType.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("userTokenType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userTokenType")
	}

	// Simple Field (requestedKeyCount)
	_requestedKeyCount, _requestedKeyCountErr := /*TODO: migrate me*/ readBuffer.ReadUint16("requestedKeyCount", 16)
	if _requestedKeyCountErr != nil {
		return nil, errors.Wrap(_requestedKeyCountErr, "Error parsing 'requestedKeyCount' field of PubSubKeyPushTargetDataType")
	}
	requestedKeyCount := _requestedKeyCount

	// Simple Field (retryInterval)
	_retryInterval, _retryIntervalErr := /*TODO: migrate me*/ readBuffer.ReadFloat64("retryInterval", 64)
	if _retryIntervalErr != nil {
		return nil, errors.Wrap(_retryIntervalErr, "Error parsing 'retryInterval' field of PubSubKeyPushTargetDataType")
	}
	retryInterval := _retryInterval

	// Simple Field (noOfPushTargetProperties)
	_noOfPushTargetProperties, _noOfPushTargetPropertiesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfPushTargetProperties", 32)
	if _noOfPushTargetPropertiesErr != nil {
		return nil, errors.Wrap(_noOfPushTargetPropertiesErr, "Error parsing 'noOfPushTargetProperties' field of PubSubKeyPushTargetDataType")
	}
	noOfPushTargetProperties := _noOfPushTargetProperties

	pushTargetProperties, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "pushTargetProperties", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("14535"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfPushTargetProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pushTargetProperties' field"))
	}

	// Simple Field (noOfSecurityGroups)
	_noOfSecurityGroups, _noOfSecurityGroupsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfSecurityGroups", 32)
	if _noOfSecurityGroupsErr != nil {
		return nil, errors.Wrap(_noOfSecurityGroupsErr, "Error parsing 'noOfSecurityGroups' field of PubSubKeyPushTargetDataType")
	}
	noOfSecurityGroups := _noOfSecurityGroups

	securityGroups, err := ReadCountArrayField[PascalString](ctx, "securityGroups", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfSecurityGroups))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroups' field"))
	}

	if closeErr := readBuffer.CloseContext("PubSubKeyPushTargetDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubKeyPushTargetDataType")
	}

	// Create a partially initialized instance
	_child := &_PubSubKeyPushTargetDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ApplicationUri:             applicationUri,
		NoOfPushTargetFolder:       noOfPushTargetFolder,
		PushTargetFolder:           pushTargetFolder,
		EndpointUrl:                endpointUrl,
		SecurityPolicyUri:          securityPolicyUri,
		UserTokenType:              userTokenType,
		RequestedKeyCount:          requestedKeyCount,
		RetryInterval:              retryInterval,
		NoOfPushTargetProperties:   noOfPushTargetProperties,
		PushTargetProperties:       pushTargetProperties,
		NoOfSecurityGroups:         noOfSecurityGroups,
		SecurityGroups:             securityGroups,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PubSubKeyPushTargetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubKeyPushTargetDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubKeyPushTargetDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubKeyPushTargetDataType")
		}

		// Simple Field (applicationUri)
		if pushErr := writeBuffer.PushContext("applicationUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for applicationUri")
		}
		_applicationUriErr := writeBuffer.WriteSerializable(ctx, m.GetApplicationUri())
		if popErr := writeBuffer.PopContext("applicationUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for applicationUri")
		}
		if _applicationUriErr != nil {
			return errors.Wrap(_applicationUriErr, "Error serializing 'applicationUri' field")
		}

		// Simple Field (noOfPushTargetFolder)
		noOfPushTargetFolder := int32(m.GetNoOfPushTargetFolder())
		_noOfPushTargetFolderErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfPushTargetFolder", 32, int32((noOfPushTargetFolder)))
		if _noOfPushTargetFolderErr != nil {
			return errors.Wrap(_noOfPushTargetFolderErr, "Error serializing 'noOfPushTargetFolder' field")
		}

		// Array Field (pushTargetFolder)
		if pushErr := writeBuffer.PushContext("pushTargetFolder", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pushTargetFolder")
		}
		for _curItem, _element := range m.GetPushTargetFolder() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetPushTargetFolder()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'pushTargetFolder' field")
			}
		}
		if popErr := writeBuffer.PopContext("pushTargetFolder", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pushTargetFolder")
		}

		// Simple Field (endpointUrl)
		if pushErr := writeBuffer.PushContext("endpointUrl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpointUrl")
		}
		_endpointUrlErr := writeBuffer.WriteSerializable(ctx, m.GetEndpointUrl())
		if popErr := writeBuffer.PopContext("endpointUrl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpointUrl")
		}
		if _endpointUrlErr != nil {
			return errors.Wrap(_endpointUrlErr, "Error serializing 'endpointUrl' field")
		}

		// Simple Field (securityPolicyUri)
		if pushErr := writeBuffer.PushContext("securityPolicyUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPolicyUri")
		}
		_securityPolicyUriErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityPolicyUri())
		if popErr := writeBuffer.PopContext("securityPolicyUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPolicyUri")
		}
		if _securityPolicyUriErr != nil {
			return errors.Wrap(_securityPolicyUriErr, "Error serializing 'securityPolicyUri' field")
		}

		// Simple Field (userTokenType)
		if pushErr := writeBuffer.PushContext("userTokenType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userTokenType")
		}
		_userTokenTypeErr := writeBuffer.WriteSerializable(ctx, m.GetUserTokenType())
		if popErr := writeBuffer.PopContext("userTokenType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userTokenType")
		}
		if _userTokenTypeErr != nil {
			return errors.Wrap(_userTokenTypeErr, "Error serializing 'userTokenType' field")
		}

		// Simple Field (requestedKeyCount)
		requestedKeyCount := uint16(m.GetRequestedKeyCount())
		_requestedKeyCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("requestedKeyCount", 16, uint16((requestedKeyCount)))
		if _requestedKeyCountErr != nil {
			return errors.Wrap(_requestedKeyCountErr, "Error serializing 'requestedKeyCount' field")
		}

		// Simple Field (retryInterval)
		retryInterval := float64(m.GetRetryInterval())
		_retryIntervalErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("retryInterval", 64, (retryInterval))
		if _retryIntervalErr != nil {
			return errors.Wrap(_retryIntervalErr, "Error serializing 'retryInterval' field")
		}

		// Simple Field (noOfPushTargetProperties)
		noOfPushTargetProperties := int32(m.GetNoOfPushTargetProperties())
		_noOfPushTargetPropertiesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfPushTargetProperties", 32, int32((noOfPushTargetProperties)))
		if _noOfPushTargetPropertiesErr != nil {
			return errors.Wrap(_noOfPushTargetPropertiesErr, "Error serializing 'noOfPushTargetProperties' field")
		}

		// Array Field (pushTargetProperties)
		if pushErr := writeBuffer.PushContext("pushTargetProperties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pushTargetProperties")
		}
		for _curItem, _element := range m.GetPushTargetProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetPushTargetProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'pushTargetProperties' field")
			}
		}
		if popErr := writeBuffer.PopContext("pushTargetProperties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pushTargetProperties")
		}

		// Simple Field (noOfSecurityGroups)
		noOfSecurityGroups := int32(m.GetNoOfSecurityGroups())
		_noOfSecurityGroupsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfSecurityGroups", 32, int32((noOfSecurityGroups)))
		if _noOfSecurityGroupsErr != nil {
			return errors.Wrap(_noOfSecurityGroupsErr, "Error serializing 'noOfSecurityGroups' field")
		}

		// Array Field (securityGroups)
		if pushErr := writeBuffer.PushContext("securityGroups", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityGroups")
		}
		for _curItem, _element := range m.GetSecurityGroups() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSecurityGroups()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'securityGroups' field")
			}
		}
		if popErr := writeBuffer.PopContext("securityGroups", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityGroups")
		}

		if popErr := writeBuffer.PopContext("PubSubKeyPushTargetDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubKeyPushTargetDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubKeyPushTargetDataType) isPubSubKeyPushTargetDataType() bool {
	return true
}

func (m *_PubSubKeyPushTargetDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
