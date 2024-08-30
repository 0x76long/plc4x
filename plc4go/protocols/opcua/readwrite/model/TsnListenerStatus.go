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

// TsnListenerStatus is an enum
type TsnListenerStatus uint32

type ITsnListenerStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TsnListenerStatus_tsnListenerStatusNone          TsnListenerStatus = 0
	TsnListenerStatus_tsnListenerStatusReady         TsnListenerStatus = 1
	TsnListenerStatus_tsnListenerStatusPartialFailed TsnListenerStatus = 2
	TsnListenerStatus_tsnListenerStatusFailed        TsnListenerStatus = 3
)

var TsnListenerStatusValues []TsnListenerStatus

func init() {
	_ = errors.New
	TsnListenerStatusValues = []TsnListenerStatus{
		TsnListenerStatus_tsnListenerStatusNone,
		TsnListenerStatus_tsnListenerStatusReady,
		TsnListenerStatus_tsnListenerStatusPartialFailed,
		TsnListenerStatus_tsnListenerStatusFailed,
	}
}

func TsnListenerStatusByValue(value uint32) (enum TsnListenerStatus, ok bool) {
	switch value {
	case 0:
		return TsnListenerStatus_tsnListenerStatusNone, true
	case 1:
		return TsnListenerStatus_tsnListenerStatusReady, true
	case 2:
		return TsnListenerStatus_tsnListenerStatusPartialFailed, true
	case 3:
		return TsnListenerStatus_tsnListenerStatusFailed, true
	}
	return 0, false
}

func TsnListenerStatusByName(value string) (enum TsnListenerStatus, ok bool) {
	switch value {
	case "tsnListenerStatusNone":
		return TsnListenerStatus_tsnListenerStatusNone, true
	case "tsnListenerStatusReady":
		return TsnListenerStatus_tsnListenerStatusReady, true
	case "tsnListenerStatusPartialFailed":
		return TsnListenerStatus_tsnListenerStatusPartialFailed, true
	case "tsnListenerStatusFailed":
		return TsnListenerStatus_tsnListenerStatusFailed, true
	}
	return 0, false
}

func TsnListenerStatusKnows(value uint32) bool {
	for _, typeValue := range TsnListenerStatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTsnListenerStatus(structType any) TsnListenerStatus {
	castFunc := func(typ any) TsnListenerStatus {
		if sTsnListenerStatus, ok := typ.(TsnListenerStatus); ok {
			return sTsnListenerStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m TsnListenerStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TsnListenerStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TsnListenerStatusParse(ctx context.Context, theBytes []byte) (TsnListenerStatus, error) {
	return TsnListenerStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TsnListenerStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TsnListenerStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("TsnListenerStatus", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TsnListenerStatus")
	}
	if enum, ok := TsnListenerStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TsnListenerStatus")
		return TsnListenerStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e TsnListenerStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TsnListenerStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("TsnListenerStatus", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TsnListenerStatus) PLC4XEnumName() string {
	switch e {
	case TsnListenerStatus_tsnListenerStatusNone:
		return "tsnListenerStatusNone"
	case TsnListenerStatus_tsnListenerStatusReady:
		return "tsnListenerStatusReady"
	case TsnListenerStatus_tsnListenerStatusPartialFailed:
		return "tsnListenerStatusPartialFailed"
	case TsnListenerStatus_tsnListenerStatusFailed:
		return "tsnListenerStatusFailed"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e TsnListenerStatus) String() string {
	return e.PLC4XEnumName()
}
