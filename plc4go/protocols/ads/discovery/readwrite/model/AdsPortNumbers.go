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

// AdsPortNumbers is an enum
type AdsPortNumbers uint16

type IAdsPortNumbers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AdsPortNumbers_LOGGER               AdsPortNumbers = 100
	AdsPortNumbers_EVENT_LOGGER         AdsPortNumbers = 110
	AdsPortNumbers_IO                   AdsPortNumbers = 300
	AdsPortNumbers_ADDITIONAL_TASK_1    AdsPortNumbers = 301
	AdsPortNumbers_ADDITIONAL_TASK_2    AdsPortNumbers = 302
	AdsPortNumbers_NC                   AdsPortNumbers = 500
	AdsPortNumbers_PLC_RUNTIME_SYSTEM_1 AdsPortNumbers = 801
	AdsPortNumbers_PLC_RUNTIME_SYSTEM_2 AdsPortNumbers = 811
	AdsPortNumbers_PLC_RUNTIME_SYSTEM_3 AdsPortNumbers = 821
	AdsPortNumbers_PLC_RUNTIME_SYSTEM_4 AdsPortNumbers = 831
	AdsPortNumbers_CAM_SWITCH           AdsPortNumbers = 900
	AdsPortNumbers_SYSTEM_SERVICE       AdsPortNumbers = 10000
	AdsPortNumbers_SCOPE                AdsPortNumbers = 14000
)

var AdsPortNumbersValues []AdsPortNumbers

func init() {
	_ = errors.New
	AdsPortNumbersValues = []AdsPortNumbers{
		AdsPortNumbers_LOGGER,
		AdsPortNumbers_EVENT_LOGGER,
		AdsPortNumbers_IO,
		AdsPortNumbers_ADDITIONAL_TASK_1,
		AdsPortNumbers_ADDITIONAL_TASK_2,
		AdsPortNumbers_NC,
		AdsPortNumbers_PLC_RUNTIME_SYSTEM_1,
		AdsPortNumbers_PLC_RUNTIME_SYSTEM_2,
		AdsPortNumbers_PLC_RUNTIME_SYSTEM_3,
		AdsPortNumbers_PLC_RUNTIME_SYSTEM_4,
		AdsPortNumbers_CAM_SWITCH,
		AdsPortNumbers_SYSTEM_SERVICE,
		AdsPortNumbers_SCOPE,
	}
}

func AdsPortNumbersByValue(value uint16) (enum AdsPortNumbers, ok bool) {
	switch value {
	case 100:
		return AdsPortNumbers_LOGGER, true
	case 10000:
		return AdsPortNumbers_SYSTEM_SERVICE, true
	case 110:
		return AdsPortNumbers_EVENT_LOGGER, true
	case 14000:
		return AdsPortNumbers_SCOPE, true
	case 300:
		return AdsPortNumbers_IO, true
	case 301:
		return AdsPortNumbers_ADDITIONAL_TASK_1, true
	case 302:
		return AdsPortNumbers_ADDITIONAL_TASK_2, true
	case 500:
		return AdsPortNumbers_NC, true
	case 801:
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_1, true
	case 811:
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_2, true
	case 821:
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_3, true
	case 831:
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_4, true
	case 900:
		return AdsPortNumbers_CAM_SWITCH, true
	}
	return 0, false
}

func AdsPortNumbersByName(value string) (enum AdsPortNumbers, ok bool) {
	switch value {
	case "LOGGER":
		return AdsPortNumbers_LOGGER, true
	case "SYSTEM_SERVICE":
		return AdsPortNumbers_SYSTEM_SERVICE, true
	case "EVENT_LOGGER":
		return AdsPortNumbers_EVENT_LOGGER, true
	case "SCOPE":
		return AdsPortNumbers_SCOPE, true
	case "IO":
		return AdsPortNumbers_IO, true
	case "ADDITIONAL_TASK_1":
		return AdsPortNumbers_ADDITIONAL_TASK_1, true
	case "ADDITIONAL_TASK_2":
		return AdsPortNumbers_ADDITIONAL_TASK_2, true
	case "NC":
		return AdsPortNumbers_NC, true
	case "PLC_RUNTIME_SYSTEM_1":
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_1, true
	case "PLC_RUNTIME_SYSTEM_2":
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_2, true
	case "PLC_RUNTIME_SYSTEM_3":
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_3, true
	case "PLC_RUNTIME_SYSTEM_4":
		return AdsPortNumbers_PLC_RUNTIME_SYSTEM_4, true
	case "CAM_SWITCH":
		return AdsPortNumbers_CAM_SWITCH, true
	}
	return 0, false
}

func AdsPortNumbersKnows(value uint16) bool {
	for _, typeValue := range AdsPortNumbersValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAdsPortNumbers(structType any) AdsPortNumbers {
	castFunc := func(typ any) AdsPortNumbers {
		if sAdsPortNumbers, ok := typ.(AdsPortNumbers); ok {
			return sAdsPortNumbers
		}
		return 0
	}
	return castFunc(structType)
}

func (m AdsPortNumbers) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m AdsPortNumbers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsPortNumbersParse(ctx context.Context, theBytes []byte) (AdsPortNumbers, error) {
	return AdsPortNumbersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsPortNumbersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsPortNumbers, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("AdsPortNumbers", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AdsPortNumbers")
	}
	if enum, ok := AdsPortNumbersByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AdsPortNumbers")
		return AdsPortNumbers(val), nil
	} else {
		return enum, nil
	}
}

func (e AdsPortNumbers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AdsPortNumbers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("AdsPortNumbers", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AdsPortNumbers) PLC4XEnumName() string {
	switch e {
	case AdsPortNumbers_LOGGER:
		return "LOGGER"
	case AdsPortNumbers_SYSTEM_SERVICE:
		return "SYSTEM_SERVICE"
	case AdsPortNumbers_EVENT_LOGGER:
		return "EVENT_LOGGER"
	case AdsPortNumbers_SCOPE:
		return "SCOPE"
	case AdsPortNumbers_IO:
		return "IO"
	case AdsPortNumbers_ADDITIONAL_TASK_1:
		return "ADDITIONAL_TASK_1"
	case AdsPortNumbers_ADDITIONAL_TASK_2:
		return "ADDITIONAL_TASK_2"
	case AdsPortNumbers_NC:
		return "NC"
	case AdsPortNumbers_PLC_RUNTIME_SYSTEM_1:
		return "PLC_RUNTIME_SYSTEM_1"
	case AdsPortNumbers_PLC_RUNTIME_SYSTEM_2:
		return "PLC_RUNTIME_SYSTEM_2"
	case AdsPortNumbers_PLC_RUNTIME_SYSTEM_3:
		return "PLC_RUNTIME_SYSTEM_3"
	case AdsPortNumbers_PLC_RUNTIME_SYSTEM_4:
		return "PLC_RUNTIME_SYSTEM_4"
	case AdsPortNumbers_CAM_SWITCH:
		return "CAM_SWITCH"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e AdsPortNumbers) String() string {
	return e.PLC4XEnumName()
}
