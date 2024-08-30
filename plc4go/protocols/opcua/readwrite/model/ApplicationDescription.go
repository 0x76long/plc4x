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

// ApplicationDescription is the corresponding interface of ApplicationDescription
type ApplicationDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetApplicationUri returns ApplicationUri (property field)
	GetApplicationUri() PascalString
	// GetProductUri returns ProductUri (property field)
	GetProductUri() PascalString
	// GetApplicationName returns ApplicationName (property field)
	GetApplicationName() LocalizedText
	// GetApplicationType returns ApplicationType (property field)
	GetApplicationType() ApplicationType
	// GetGatewayServerUri returns GatewayServerUri (property field)
	GetGatewayServerUri() PascalString
	// GetDiscoveryProfileUri returns DiscoveryProfileUri (property field)
	GetDiscoveryProfileUri() PascalString
	// GetNoOfDiscoveryUrls returns NoOfDiscoveryUrls (property field)
	GetNoOfDiscoveryUrls() int32
	// GetDiscoveryUrls returns DiscoveryUrls (property field)
	GetDiscoveryUrls() []PascalString
}

// ApplicationDescriptionExactly can be used when we want exactly this type and not a type which fulfills ApplicationDescription.
// This is useful for switch cases.
type ApplicationDescriptionExactly interface {
	ApplicationDescription
	isApplicationDescription() bool
}

// _ApplicationDescription is the data-structure of this message
type _ApplicationDescription struct {
	*_ExtensionObjectDefinition
	ApplicationUri      PascalString
	ProductUri          PascalString
	ApplicationName     LocalizedText
	ApplicationType     ApplicationType
	GatewayServerUri    PascalString
	DiscoveryProfileUri PascalString
	NoOfDiscoveryUrls   int32
	DiscoveryUrls       []PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApplicationDescription) GetIdentifier() string {
	return "310"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApplicationDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ApplicationDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApplicationDescription) GetApplicationUri() PascalString {
	return m.ApplicationUri
}

func (m *_ApplicationDescription) GetProductUri() PascalString {
	return m.ProductUri
}

func (m *_ApplicationDescription) GetApplicationName() LocalizedText {
	return m.ApplicationName
}

func (m *_ApplicationDescription) GetApplicationType() ApplicationType {
	return m.ApplicationType
}

func (m *_ApplicationDescription) GetGatewayServerUri() PascalString {
	return m.GatewayServerUri
}

func (m *_ApplicationDescription) GetDiscoveryProfileUri() PascalString {
	return m.DiscoveryProfileUri
}

func (m *_ApplicationDescription) GetNoOfDiscoveryUrls() int32 {
	return m.NoOfDiscoveryUrls
}

func (m *_ApplicationDescription) GetDiscoveryUrls() []PascalString {
	return m.DiscoveryUrls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApplicationDescription factory function for _ApplicationDescription
func NewApplicationDescription(applicationUri PascalString, productUri PascalString, applicationName LocalizedText, applicationType ApplicationType, gatewayServerUri PascalString, discoveryProfileUri PascalString, noOfDiscoveryUrls int32, discoveryUrls []PascalString) *_ApplicationDescription {
	_result := &_ApplicationDescription{
		ApplicationUri:             applicationUri,
		ProductUri:                 productUri,
		ApplicationName:            applicationName,
		ApplicationType:            applicationType,
		GatewayServerUri:           gatewayServerUri,
		DiscoveryProfileUri:        discoveryProfileUri,
		NoOfDiscoveryUrls:          noOfDiscoveryUrls,
		DiscoveryUrls:              discoveryUrls,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApplicationDescription(structType any) ApplicationDescription {
	if casted, ok := structType.(ApplicationDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ApplicationDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ApplicationDescription) GetTypeName() string {
	return "ApplicationDescription"
}

func (m *_ApplicationDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (applicationUri)
	lengthInBits += m.ApplicationUri.GetLengthInBits(ctx)

	// Simple field (productUri)
	lengthInBits += m.ProductUri.GetLengthInBits(ctx)

	// Simple field (applicationName)
	lengthInBits += m.ApplicationName.GetLengthInBits(ctx)

	// Simple field (applicationType)
	lengthInBits += 32

	// Simple field (gatewayServerUri)
	lengthInBits += m.GatewayServerUri.GetLengthInBits(ctx)

	// Simple field (discoveryProfileUri)
	lengthInBits += m.DiscoveryProfileUri.GetLengthInBits(ctx)

	// Simple field (noOfDiscoveryUrls)
	lengthInBits += 32

	// Array field
	if len(m.DiscoveryUrls) > 0 {
		for _curItem, element := range m.DiscoveryUrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiscoveryUrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ApplicationDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApplicationDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (ApplicationDescription, error) {
	return ApplicationDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ApplicationDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ApplicationDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApplicationDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApplicationDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (applicationUri)
	if pullErr := readBuffer.PullContext("applicationUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for applicationUri")
	}
	_applicationUri, _applicationUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _applicationUriErr != nil {
		return nil, errors.Wrap(_applicationUriErr, "Error parsing 'applicationUri' field of ApplicationDescription")
	}
	applicationUri := _applicationUri.(PascalString)
	if closeErr := readBuffer.CloseContext("applicationUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for applicationUri")
	}

	// Simple Field (productUri)
	if pullErr := readBuffer.PullContext("productUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for productUri")
	}
	_productUri, _productUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _productUriErr != nil {
		return nil, errors.Wrap(_productUriErr, "Error parsing 'productUri' field of ApplicationDescription")
	}
	productUri := _productUri.(PascalString)
	if closeErr := readBuffer.CloseContext("productUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for productUri")
	}

	// Simple Field (applicationName)
	if pullErr := readBuffer.PullContext("applicationName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for applicationName")
	}
	_applicationName, _applicationNameErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _applicationNameErr != nil {
		return nil, errors.Wrap(_applicationNameErr, "Error parsing 'applicationName' field of ApplicationDescription")
	}
	applicationName := _applicationName.(LocalizedText)
	if closeErr := readBuffer.CloseContext("applicationName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for applicationName")
	}

	// Simple Field (applicationType)
	if pullErr := readBuffer.PullContext("applicationType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for applicationType")
	}
	_applicationType, _applicationTypeErr := ApplicationTypeParseWithBuffer(ctx, readBuffer)
	if _applicationTypeErr != nil {
		return nil, errors.Wrap(_applicationTypeErr, "Error parsing 'applicationType' field of ApplicationDescription")
	}
	applicationType := _applicationType
	if closeErr := readBuffer.CloseContext("applicationType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for applicationType")
	}

	// Simple Field (gatewayServerUri)
	if pullErr := readBuffer.PullContext("gatewayServerUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for gatewayServerUri")
	}
	_gatewayServerUri, _gatewayServerUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _gatewayServerUriErr != nil {
		return nil, errors.Wrap(_gatewayServerUriErr, "Error parsing 'gatewayServerUri' field of ApplicationDescription")
	}
	gatewayServerUri := _gatewayServerUri.(PascalString)
	if closeErr := readBuffer.CloseContext("gatewayServerUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for gatewayServerUri")
	}

	// Simple Field (discoveryProfileUri)
	if pullErr := readBuffer.PullContext("discoveryProfileUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for discoveryProfileUri")
	}
	_discoveryProfileUri, _discoveryProfileUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _discoveryProfileUriErr != nil {
		return nil, errors.Wrap(_discoveryProfileUriErr, "Error parsing 'discoveryProfileUri' field of ApplicationDescription")
	}
	discoveryProfileUri := _discoveryProfileUri.(PascalString)
	if closeErr := readBuffer.CloseContext("discoveryProfileUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for discoveryProfileUri")
	}

	// Simple Field (noOfDiscoveryUrls)
	_noOfDiscoveryUrls, _noOfDiscoveryUrlsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfDiscoveryUrls", 32)
	if _noOfDiscoveryUrlsErr != nil {
		return nil, errors.Wrap(_noOfDiscoveryUrlsErr, "Error parsing 'noOfDiscoveryUrls' field of ApplicationDescription")
	}
	noOfDiscoveryUrls := _noOfDiscoveryUrls

	discoveryUrls, err := ReadCountArrayField[PascalString](ctx, "discoveryUrls", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfDiscoveryUrls))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryUrls' field"))
	}

	if closeErr := readBuffer.CloseContext("ApplicationDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApplicationDescription")
	}

	// Create a partially initialized instance
	_child := &_ApplicationDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ApplicationUri:             applicationUri,
		ProductUri:                 productUri,
		ApplicationName:            applicationName,
		ApplicationType:            applicationType,
		GatewayServerUri:           gatewayServerUri,
		DiscoveryProfileUri:        discoveryProfileUri,
		NoOfDiscoveryUrls:          noOfDiscoveryUrls,
		DiscoveryUrls:              discoveryUrls,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ApplicationDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApplicationDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApplicationDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApplicationDescription")
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

		// Simple Field (productUri)
		if pushErr := writeBuffer.PushContext("productUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for productUri")
		}
		_productUriErr := writeBuffer.WriteSerializable(ctx, m.GetProductUri())
		if popErr := writeBuffer.PopContext("productUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for productUri")
		}
		if _productUriErr != nil {
			return errors.Wrap(_productUriErr, "Error serializing 'productUri' field")
		}

		// Simple Field (applicationName)
		if pushErr := writeBuffer.PushContext("applicationName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for applicationName")
		}
		_applicationNameErr := writeBuffer.WriteSerializable(ctx, m.GetApplicationName())
		if popErr := writeBuffer.PopContext("applicationName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for applicationName")
		}
		if _applicationNameErr != nil {
			return errors.Wrap(_applicationNameErr, "Error serializing 'applicationName' field")
		}

		// Simple Field (applicationType)
		if pushErr := writeBuffer.PushContext("applicationType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for applicationType")
		}
		_applicationTypeErr := writeBuffer.WriteSerializable(ctx, m.GetApplicationType())
		if popErr := writeBuffer.PopContext("applicationType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for applicationType")
		}
		if _applicationTypeErr != nil {
			return errors.Wrap(_applicationTypeErr, "Error serializing 'applicationType' field")
		}

		// Simple Field (gatewayServerUri)
		if pushErr := writeBuffer.PushContext("gatewayServerUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for gatewayServerUri")
		}
		_gatewayServerUriErr := writeBuffer.WriteSerializable(ctx, m.GetGatewayServerUri())
		if popErr := writeBuffer.PopContext("gatewayServerUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for gatewayServerUri")
		}
		if _gatewayServerUriErr != nil {
			return errors.Wrap(_gatewayServerUriErr, "Error serializing 'gatewayServerUri' field")
		}

		// Simple Field (discoveryProfileUri)
		if pushErr := writeBuffer.PushContext("discoveryProfileUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for discoveryProfileUri")
		}
		_discoveryProfileUriErr := writeBuffer.WriteSerializable(ctx, m.GetDiscoveryProfileUri())
		if popErr := writeBuffer.PopContext("discoveryProfileUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for discoveryProfileUri")
		}
		if _discoveryProfileUriErr != nil {
			return errors.Wrap(_discoveryProfileUriErr, "Error serializing 'discoveryProfileUri' field")
		}

		// Simple Field (noOfDiscoveryUrls)
		noOfDiscoveryUrls := int32(m.GetNoOfDiscoveryUrls())
		_noOfDiscoveryUrlsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDiscoveryUrls", 32, int32((noOfDiscoveryUrls)))
		if _noOfDiscoveryUrlsErr != nil {
			return errors.Wrap(_noOfDiscoveryUrlsErr, "Error serializing 'noOfDiscoveryUrls' field")
		}

		// Array Field (discoveryUrls)
		if pushErr := writeBuffer.PushContext("discoveryUrls", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for discoveryUrls")
		}
		for _curItem, _element := range m.GetDiscoveryUrls() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiscoveryUrls()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'discoveryUrls' field")
			}
		}
		if popErr := writeBuffer.PopContext("discoveryUrls", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for discoveryUrls")
		}

		if popErr := writeBuffer.PopContext("ApplicationDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApplicationDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApplicationDescription) isApplicationDescription() bool {
	return true
}

func (m *_ApplicationDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
