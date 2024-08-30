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

// PerformUpdateType is an enum
type PerformUpdateType uint32

type IPerformUpdateType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PerformUpdateType_performUpdateTypeInsert  PerformUpdateType = 1
	PerformUpdateType_performUpdateTypeReplace PerformUpdateType = 2
	PerformUpdateType_performUpdateTypeUpdate  PerformUpdateType = 3
	PerformUpdateType_performUpdateTypeRemove  PerformUpdateType = 4
)

var PerformUpdateTypeValues []PerformUpdateType

func init() {
	_ = errors.New
	PerformUpdateTypeValues = []PerformUpdateType{
		PerformUpdateType_performUpdateTypeInsert,
		PerformUpdateType_performUpdateTypeReplace,
		PerformUpdateType_performUpdateTypeUpdate,
		PerformUpdateType_performUpdateTypeRemove,
	}
}

func PerformUpdateTypeByValue(value uint32) (enum PerformUpdateType, ok bool) {
	switch value {
	case 1:
		return PerformUpdateType_performUpdateTypeInsert, true
	case 2:
		return PerformUpdateType_performUpdateTypeReplace, true
	case 3:
		return PerformUpdateType_performUpdateTypeUpdate, true
	case 4:
		return PerformUpdateType_performUpdateTypeRemove, true
	}
	return 0, false
}

func PerformUpdateTypeByName(value string) (enum PerformUpdateType, ok bool) {
	switch value {
	case "performUpdateTypeInsert":
		return PerformUpdateType_performUpdateTypeInsert, true
	case "performUpdateTypeReplace":
		return PerformUpdateType_performUpdateTypeReplace, true
	case "performUpdateTypeUpdate":
		return PerformUpdateType_performUpdateTypeUpdate, true
	case "performUpdateTypeRemove":
		return PerformUpdateType_performUpdateTypeRemove, true
	}
	return 0, false
}

func PerformUpdateTypeKnows(value uint32) bool {
	for _, typeValue := range PerformUpdateTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPerformUpdateType(structType any) PerformUpdateType {
	castFunc := func(typ any) PerformUpdateType {
		if sPerformUpdateType, ok := typ.(PerformUpdateType); ok {
			return sPerformUpdateType
		}
		return 0
	}
	return castFunc(structType)
}

func (m PerformUpdateType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m PerformUpdateType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PerformUpdateTypeParse(ctx context.Context, theBytes []byte) (PerformUpdateType, error) {
	return PerformUpdateTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PerformUpdateTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PerformUpdateType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("PerformUpdateType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PerformUpdateType")
	}
	if enum, ok := PerformUpdateTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PerformUpdateType")
		return PerformUpdateType(val), nil
	} else {
		return enum, nil
	}
}

func (e PerformUpdateType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PerformUpdateType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("PerformUpdateType", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PerformUpdateType) PLC4XEnumName() string {
	switch e {
	case PerformUpdateType_performUpdateTypeInsert:
		return "performUpdateTypeInsert"
	case PerformUpdateType_performUpdateTypeReplace:
		return "performUpdateTypeReplace"
	case PerformUpdateType_performUpdateTypeUpdate:
		return "performUpdateTypeUpdate"
	case PerformUpdateType_performUpdateTypeRemove:
		return "performUpdateTypeRemove"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e PerformUpdateType) String() string {
	return e.PLC4XEnumName()
}
