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

// OpcuaIdentifierType is an enum
type OpcuaIdentifierType string

type IOpcuaIdentifierType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaIdentifierType_STRING_IDENTIFIER OpcuaIdentifierType = "s"
	OpcuaIdentifierType_NUMBER_IDENTIFIER OpcuaIdentifierType = "i"
	OpcuaIdentifierType_GUID_IDENTIFIER   OpcuaIdentifierType = "g"
	OpcuaIdentifierType_BINARY_IDENTIFIER OpcuaIdentifierType = "b"
)

var OpcuaIdentifierTypeValues []OpcuaIdentifierType

func init() {
	_ = errors.New
	OpcuaIdentifierTypeValues = []OpcuaIdentifierType{
		OpcuaIdentifierType_STRING_IDENTIFIER,
		OpcuaIdentifierType_NUMBER_IDENTIFIER,
		OpcuaIdentifierType_GUID_IDENTIFIER,
		OpcuaIdentifierType_BINARY_IDENTIFIER,
	}
}

func OpcuaIdentifierTypeByValue(value string) (enum OpcuaIdentifierType, ok bool) {
	switch value {
	case "b":
		return OpcuaIdentifierType_BINARY_IDENTIFIER, true
	case "g":
		return OpcuaIdentifierType_GUID_IDENTIFIER, true
	case "i":
		return OpcuaIdentifierType_NUMBER_IDENTIFIER, true
	case "s":
		return OpcuaIdentifierType_STRING_IDENTIFIER, true
	}
	return "", false
}

func OpcuaIdentifierTypeByName(value string) (enum OpcuaIdentifierType, ok bool) {
	switch value {
	case "BINARY_IDENTIFIER":
		return OpcuaIdentifierType_BINARY_IDENTIFIER, true
	case "GUID_IDENTIFIER":
		return OpcuaIdentifierType_GUID_IDENTIFIER, true
	case "NUMBER_IDENTIFIER":
		return OpcuaIdentifierType_NUMBER_IDENTIFIER, true
	case "STRING_IDENTIFIER":
		return OpcuaIdentifierType_STRING_IDENTIFIER, true
	}
	return "", false
}

func OpcuaIdentifierTypeKnows(value string) bool {
	for _, typeValue := range OpcuaIdentifierTypeValues {
		if string(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaIdentifierType(structType any) OpcuaIdentifierType {
	castFunc := func(typ any) OpcuaIdentifierType {
		if sOpcuaIdentifierType, ok := typ.(OpcuaIdentifierType); ok {
			return sOpcuaIdentifierType
		}
		return ""
	}
	return castFunc(structType)
}

func (m OpcuaIdentifierType) GetLengthInBits(ctx context.Context) uint16 {
	return 0
}

func (m OpcuaIdentifierType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaIdentifierTypeParse(ctx context.Context, theBytes []byte) (OpcuaIdentifierType, error) {
	return OpcuaIdentifierTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaIdentifierTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaIdentifierType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadString("OpcuaIdentifierType", uint32(8), utils.WithEncoding("UTF-8"))
	if err != nil {
		return "", errors.Wrap(err, "error reading OpcuaIdentifierType")
	}
	if enum, ok := OpcuaIdentifierTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaIdentifierType")
		return OpcuaIdentifierType(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaIdentifierType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaIdentifierType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteString("OpcuaIdentifierType", uint32(8), string(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()), utils.WithEncoding("UTF-8)"))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaIdentifierType) PLC4XEnumName() string {
	switch e {
	case OpcuaIdentifierType_BINARY_IDENTIFIER:
		return "BINARY_IDENTIFIER"
	case OpcuaIdentifierType_GUID_IDENTIFIER:
		return "GUID_IDENTIFIER"
	case OpcuaIdentifierType_NUMBER_IDENTIFIER:
		return "NUMBER_IDENTIFIER"
	case OpcuaIdentifierType_STRING_IDENTIFIER:
		return "STRING_IDENTIFIER"
	}
	return fmt.Sprintf("Unknown(%v)", string(e))
}

func (e OpcuaIdentifierType) String() string {
	return e.PLC4XEnumName()
}
