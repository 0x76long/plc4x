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

// OpcuaNodeIdServicesVariableImage is an enum
type OpcuaNodeIdServicesVariableImage int32

type IOpcuaNodeIdServicesVariableImage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableImage_ImageItemType_XAxisDefinition OpcuaNodeIdServicesVariableImage = 12055
	OpcuaNodeIdServicesVariableImage_ImageItemType_YAxisDefinition OpcuaNodeIdServicesVariableImage = 12056
)

var OpcuaNodeIdServicesVariableImageValues []OpcuaNodeIdServicesVariableImage

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableImageValues = []OpcuaNodeIdServicesVariableImage{
		OpcuaNodeIdServicesVariableImage_ImageItemType_XAxisDefinition,
		OpcuaNodeIdServicesVariableImage_ImageItemType_YAxisDefinition,
	}
}

func OpcuaNodeIdServicesVariableImageByValue(value int32) (enum OpcuaNodeIdServicesVariableImage, ok bool) {
	switch value {
	case 12055:
		return OpcuaNodeIdServicesVariableImage_ImageItemType_XAxisDefinition, true
	case 12056:
		return OpcuaNodeIdServicesVariableImage_ImageItemType_YAxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableImageByName(value string) (enum OpcuaNodeIdServicesVariableImage, ok bool) {
	switch value {
	case "ImageItemType_XAxisDefinition":
		return OpcuaNodeIdServicesVariableImage_ImageItemType_XAxisDefinition, true
	case "ImageItemType_YAxisDefinition":
		return OpcuaNodeIdServicesVariableImage_ImageItemType_YAxisDefinition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableImageKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableImageValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableImage(structType any) OpcuaNodeIdServicesVariableImage {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableImage {
		if sOpcuaNodeIdServicesVariableImage, ok := typ.(OpcuaNodeIdServicesVariableImage); ok {
			return sOpcuaNodeIdServicesVariableImage
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableImage) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableImage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableImageParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableImage, error) {
	return OpcuaNodeIdServicesVariableImageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableImageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableImage, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableImage", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableImage")
	}
	if enum, ok := OpcuaNodeIdServicesVariableImageByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableImage")
		return OpcuaNodeIdServicesVariableImage(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableImage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableImage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableImage", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableImage) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableImage_ImageItemType_XAxisDefinition:
		return "ImageItemType_XAxisDefinition"
	case OpcuaNodeIdServicesVariableImage_ImageItemType_YAxisDefinition:
		return "ImageItemType_YAxisDefinition"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableImage) String() string {
	return e.PLC4XEnumName()
}
