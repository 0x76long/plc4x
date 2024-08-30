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

package readwrite

import (
	"context"

	"github.com/apache/plc4x/plc4go/protocols/opcua/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type OpcuaParserHelper struct {
}

func (m OpcuaParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (any, error) {
	switch typeName {
	case "LocaleId":
		return model.LocaleIdParseWithBuffer(context.Background(), io)
	case "ImageGIF":
		return model.ImageGIFParseWithBuffer(context.Background(), io)
	case "EncodedTicket":
		return model.EncodedTicketParseWithBuffer(context.Background(), io)
	case "OpenChannelMessage":
		response, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.OpenChannelMessageParseWithBuffer(context.Background(), io, response)
	case "ImageJPG":
		return model.ImageJPGParseWithBuffer(context.Background(), io)
	case "PascalByteString":
		return model.PascalByteStringParseWithBuffer(context.Background(), io)
	case "DiagnosticInfo":
		return model.DiagnosticInfoParseWithBuffer(context.Background(), io)
	case "PascalString":
		return model.PascalStringParseWithBuffer(context.Background(), io)
	case "TwoByteNodeId":
		return model.TwoByteNodeIdParseWithBuffer(context.Background(), io)
	case "OpcuaAPU":
		response, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.OpcuaAPUParseWithBuffer(context.Background(), io, response)
	case "Index":
		return model.IndexParseWithBuffer(context.Background(), io)
	case "StatusCode":
		return model.StatusCodeParseWithBuffer(context.Background(), io)
	case "NormalizedString":
		return model.NormalizedStringParseWithBuffer(context.Background(), io)
	case "QualifiedName":
		return model.QualifiedNameParseWithBuffer(context.Background(), io)
	case "NumericNodeId":
		return model.NumericNodeIdParseWithBuffer(context.Background(), io)
	case "FourByteNodeId":
		return model.FourByteNodeIdParseWithBuffer(context.Background(), io)
	case "AudioDataType":
		return model.AudioDataTypeParseWithBuffer(context.Background(), io)
	case "SecurityHeader":
		return model.SecurityHeaderParseWithBuffer(context.Background(), io)
	case "UserIdentityTokenDefinition":
		identifier, err := utils.StrToString(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.UserIdentityTokenDefinitionParseWithBuffer(context.Background(), io, identifier)
	case "ContinuationPoint":
		return model.ContinuationPointParseWithBuffer(context.Background(), io)
	case "Variant":
		return model.VariantParseWithBuffer(context.Background(), io)
	case "Payload":
		extensible, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		byteCount, err := utils.StrToUint32(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.PayloadParseWithBuffer(context.Background(), io, extensible, byteCount)
	case "ExtensionObjectEncodingMask":
		return model.ExtensionObjectEncodingMaskParseWithBuffer(context.Background(), io)
	case "DurationString":
		return model.DurationStringParseWithBuffer(context.Background(), io)
	case "Structure":
		return model.StructureParseWithBuffer(context.Background(), io)
	case "OpcuaConstants":
		return model.OpcuaConstantsParseWithBuffer(context.Background(), io)
	case "ExtensionHeader":
		return model.ExtensionHeaderParseWithBuffer(context.Background(), io)
	case "UtcTime":
		return model.UtcTimeParseWithBuffer(context.Background(), io)
	case "MessagePDU":
		response, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.MessagePDUParseWithBuffer(context.Background(), io, response)
	case "Counter":
		return model.CounterParseWithBuffer(context.Background(), io)
	case "SequenceHeader":
		return model.SequenceHeaderParseWithBuffer(context.Background(), io)
	case "NodeId":
		return model.NodeIdParseWithBuffer(context.Background(), io)
	case "RsaEncryptedSecret":
		return model.RsaEncryptedSecretParseWithBuffer(context.Background(), io)
	case "ExtensionObject":
		includeEncodingMask, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.ExtensionObjectParseWithBuffer(context.Background(), io, includeEncodingMask)
	case "LocalizedText":
		return model.LocalizedTextParseWithBuffer(context.Background(), io)
	case "IntegerId":
		return model.IntegerIdParseWithBuffer(context.Background(), io)
	case "ByteStringArray":
		return model.ByteStringArrayParseWithBuffer(context.Background(), io)
	case "Handle":
		return model.HandleParseWithBuffer(context.Background(), io)
	case "ImagePNG":
		return model.ImagePNGParseWithBuffer(context.Background(), io)
	case "XmlElement":
		return model.XmlElementParseWithBuffer(context.Background(), io)
	case "SessionAuthenticationToken":
		return model.SessionAuthenticationTokenParseWithBuffer(context.Background(), io)
	case "DataValue":
		return model.DataValueParseWithBuffer(context.Background(), io)
	case "GuidNodeId":
		return model.GuidNodeIdParseWithBuffer(context.Background(), io)
	case "GuidValue":
		return model.GuidValueParseWithBuffer(context.Background(), io)
	case "TrimmedString":
		return model.TrimmedStringParseWithBuffer(context.Background(), io)
	case "ApplicationInstanceCertificate":
		return model.ApplicationInstanceCertificateParseWithBuffer(context.Background(), io)
	case "BitFieldMaskDataType":
		return model.BitFieldMaskDataTypeParseWithBuffer(context.Background(), io)
	case "ImageBMP":
		return model.ImageBMPParseWithBuffer(context.Background(), io)
	case "ExtensionObjectDefinition":
		identifier, err := utils.StrToString(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.ExtensionObjectDefinitionParseWithBuffer(context.Background(), io, identifier)
	case "ExpandedNodeId":
		return model.ExpandedNodeIdParseWithBuffer(context.Background(), io)
	case "OpcuaProtocolLimits":
		return model.OpcuaProtocolLimitsParseWithBuffer(context.Background(), io)
	case "NumericRange":
		return model.NumericRangeParseWithBuffer(context.Background(), io)
	case "SemanticVersionString":
		return model.SemanticVersionStringParseWithBuffer(context.Background(), io)
	case "ByteStringNodeId":
		return model.ByteStringNodeIdParseWithBuffer(context.Background(), io)
	case "TimeString":
		return model.TimeStringParseWithBuffer(context.Background(), io)
	case "EccEncryptedSecret":
		return model.EccEncryptedSecretParseWithBuffer(context.Background(), io)
	case "StringNodeId":
		return model.StringNodeIdParseWithBuffer(context.Background(), io)
	case "VersionTime":
		return model.VersionTimeParseWithBuffer(context.Background(), io)
	case "UriString":
		return model.UriStringParseWithBuffer(context.Background(), io)
	case "DecimalString":
		return model.DecimalStringParseWithBuffer(context.Background(), io)
	case "NodeIdTypeDefinition":
		return model.NodeIdTypeDefinitionParseWithBuffer(context.Background(), io)
	case "DateString":
		return model.DateStringParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
