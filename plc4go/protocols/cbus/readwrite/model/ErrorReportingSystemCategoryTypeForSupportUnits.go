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

// ErrorReportingSystemCategoryTypeForSupportUnits is an enum
type ErrorReportingSystemCategoryTypeForSupportUnits uint8

type IErrorReportingSystemCategoryTypeForSupportUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ErrorReportingSystemCategoryTypeForSupportUnits_POWER_SUPPLIES ErrorReportingSystemCategoryTypeForSupportUnits = 0x0
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_1     ErrorReportingSystemCategoryTypeForSupportUnits = 0x1
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_2     ErrorReportingSystemCategoryTypeForSupportUnits = 0x2
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_3     ErrorReportingSystemCategoryTypeForSupportUnits = 0x3
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_4     ErrorReportingSystemCategoryTypeForSupportUnits = 0x4
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_5     ErrorReportingSystemCategoryTypeForSupportUnits = 0x5
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_6     ErrorReportingSystemCategoryTypeForSupportUnits = 0x6
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_7     ErrorReportingSystemCategoryTypeForSupportUnits = 0x7
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_8     ErrorReportingSystemCategoryTypeForSupportUnits = 0x8
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_9     ErrorReportingSystemCategoryTypeForSupportUnits = 0x9
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_10    ErrorReportingSystemCategoryTypeForSupportUnits = 0xA
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_11    ErrorReportingSystemCategoryTypeForSupportUnits = 0xB
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_12    ErrorReportingSystemCategoryTypeForSupportUnits = 0xC
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_13    ErrorReportingSystemCategoryTypeForSupportUnits = 0xD
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_14    ErrorReportingSystemCategoryTypeForSupportUnits = 0xE
	ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_15    ErrorReportingSystemCategoryTypeForSupportUnits = 0xF
)

var ErrorReportingSystemCategoryTypeForSupportUnitsValues []ErrorReportingSystemCategoryTypeForSupportUnits

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryTypeForSupportUnitsValues = []ErrorReportingSystemCategoryTypeForSupportUnits{
		ErrorReportingSystemCategoryTypeForSupportUnits_POWER_SUPPLIES,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_1,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_2,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_3,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_4,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_5,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_6,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_7,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_8,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_9,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_10,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_11,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_12,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_13,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_14,
		ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_15,
	}
}

func ErrorReportingSystemCategoryTypeForSupportUnitsByValue(value uint8) (enum ErrorReportingSystemCategoryTypeForSupportUnits, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryTypeForSupportUnits_POWER_SUPPLIES, true
	case 0x1:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_1, true
	case 0x2:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_3, true
	case 0x4:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_4, true
	case 0x5:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_5, true
	case 0x6:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_6, true
	case 0x7:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_7, true
	case 0x8:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_8, true
	case 0x9:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_9, true
	case 0xA:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_10, true
	case 0xB:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_11, true
	case 0xC:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_12, true
	case 0xD:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_13, true
	case 0xE:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_14, true
	case 0xF:
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForSupportUnitsByName(value string) (enum ErrorReportingSystemCategoryTypeForSupportUnits, ok bool) {
	switch value {
	case "POWER_SUPPLIES":
		return ErrorReportingSystemCategoryTypeForSupportUnits_POWER_SUPPLIES, true
	case "RESERVED_1":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_3, true
	case "RESERVED_4":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_4, true
	case "RESERVED_5":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_5, true
	case "RESERVED_6":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_6, true
	case "RESERVED_7":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_7, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_8, true
	case "RESERVED_9":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_9, true
	case "RESERVED_10":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_10, true
	case "RESERVED_11":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_11, true
	case "RESERVED_12":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_12, true
	case "RESERVED_13":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_13, true
	case "RESERVED_14":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_14, true
	case "RESERVED_15":
		return ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForSupportUnitsKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryTypeForSupportUnitsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryTypeForSupportUnits(structType any) ErrorReportingSystemCategoryTypeForSupportUnits {
	castFunc := func(typ any) ErrorReportingSystemCategoryTypeForSupportUnits {
		if sErrorReportingSystemCategoryTypeForSupportUnits, ok := typ.(ErrorReportingSystemCategoryTypeForSupportUnits); ok {
			return sErrorReportingSystemCategoryTypeForSupportUnits
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryTypeForSupportUnits) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryTypeForSupportUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeForSupportUnitsParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryTypeForSupportUnits, error) {
	return ErrorReportingSystemCategoryTypeForSupportUnitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryTypeForSupportUnitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryTypeForSupportUnits, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("ErrorReportingSystemCategoryTypeForSupportUnits", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryTypeForSupportUnits")
	}
	if enum, ok := ErrorReportingSystemCategoryTypeForSupportUnitsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ErrorReportingSystemCategoryTypeForSupportUnits")
		return ErrorReportingSystemCategoryTypeForSupportUnits(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryTypeForSupportUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryTypeForSupportUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("ErrorReportingSystemCategoryTypeForSupportUnits", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryTypeForSupportUnits) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryTypeForSupportUnits_POWER_SUPPLIES:
		return "POWER_SUPPLIES"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_4:
		return "RESERVED_4"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_5:
		return "RESERVED_5"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_6:
		return "RESERVED_6"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_7:
		return "RESERVED_7"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_9:
		return "RESERVED_9"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_10:
		return "RESERVED_10"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_11:
		return "RESERVED_11"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_12:
		return "RESERVED_12"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_13:
		return "RESERVED_13"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_14:
		return "RESERVED_14"
	case ErrorReportingSystemCategoryTypeForSupportUnits_RESERVED_15:
		return "RESERVED_15"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ErrorReportingSystemCategoryTypeForSupportUnits) String() string {
	return e.PLC4XEnumName()
}
