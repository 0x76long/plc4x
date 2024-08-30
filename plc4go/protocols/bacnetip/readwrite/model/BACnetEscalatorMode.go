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

// BACnetEscalatorMode is an enum
type BACnetEscalatorMode uint16

type IBACnetEscalatorMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetEscalatorMode_UNKNOWN                  BACnetEscalatorMode = 0
	BACnetEscalatorMode_STOP                     BACnetEscalatorMode = 1
	BACnetEscalatorMode_UP                       BACnetEscalatorMode = 2
	BACnetEscalatorMode_DOWN                     BACnetEscalatorMode = 3
	BACnetEscalatorMode_INSPECTION               BACnetEscalatorMode = 4
	BACnetEscalatorMode_OUT_OF_SERVICE           BACnetEscalatorMode = 5
	BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE BACnetEscalatorMode = 0xFFFF
)

var BACnetEscalatorModeValues []BACnetEscalatorMode

func init() {
	_ = errors.New
	BACnetEscalatorModeValues = []BACnetEscalatorMode{
		BACnetEscalatorMode_UNKNOWN,
		BACnetEscalatorMode_STOP,
		BACnetEscalatorMode_UP,
		BACnetEscalatorMode_DOWN,
		BACnetEscalatorMode_INSPECTION,
		BACnetEscalatorMode_OUT_OF_SERVICE,
		BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEscalatorModeByValue(value uint16) (enum BACnetEscalatorMode, ok bool) {
	switch value {
	case 0:
		return BACnetEscalatorMode_UNKNOWN, true
	case 0xFFFF:
		return BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetEscalatorMode_STOP, true
	case 2:
		return BACnetEscalatorMode_UP, true
	case 3:
		return BACnetEscalatorMode_DOWN, true
	case 4:
		return BACnetEscalatorMode_INSPECTION, true
	case 5:
		return BACnetEscalatorMode_OUT_OF_SERVICE, true
	}
	return 0, false
}

func BACnetEscalatorModeByName(value string) (enum BACnetEscalatorMode, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetEscalatorMode_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE, true
	case "STOP":
		return BACnetEscalatorMode_STOP, true
	case "UP":
		return BACnetEscalatorMode_UP, true
	case "DOWN":
		return BACnetEscalatorMode_DOWN, true
	case "INSPECTION":
		return BACnetEscalatorMode_INSPECTION, true
	case "OUT_OF_SERVICE":
		return BACnetEscalatorMode_OUT_OF_SERVICE, true
	}
	return 0, false
}

func BACnetEscalatorModeKnows(value uint16) bool {
	for _, typeValue := range BACnetEscalatorModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEscalatorMode(structType any) BACnetEscalatorMode {
	castFunc := func(typ any) BACnetEscalatorMode {
		if sBACnetEscalatorMode, ok := typ.(BACnetEscalatorMode); ok {
			return sBACnetEscalatorMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEscalatorMode) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetEscalatorMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEscalatorModeParse(ctx context.Context, theBytes []byte) (BACnetEscalatorMode, error) {
	return BACnetEscalatorModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEscalatorModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEscalatorMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("BACnetEscalatorMode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetEscalatorMode")
	}
	if enum, ok := BACnetEscalatorModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetEscalatorMode")
		return BACnetEscalatorMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetEscalatorMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetEscalatorMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("BACnetEscalatorMode", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEscalatorMode) PLC4XEnumName() string {
	switch e {
	case BACnetEscalatorMode_UNKNOWN:
		return "UNKNOWN"
	case BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEscalatorMode_STOP:
		return "STOP"
	case BACnetEscalatorMode_UP:
		return "UP"
	case BACnetEscalatorMode_DOWN:
		return "DOWN"
	case BACnetEscalatorMode_INSPECTION:
		return "INSPECTION"
	case BACnetEscalatorMode_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetEscalatorMode) String() string {
	return e.PLC4XEnumName()
}
