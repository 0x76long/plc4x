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

// Status is an enum
type Status uint8

type IStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	Status_NO_ERROR                        Status = 0x00
	Status_PROTOCOL_TYPE_NOT_SUPPORTED     Status = 0x01
	Status_UNSUPPORTED_PROTOCOL_VERSION    Status = 0x02
	Status_OUT_OF_ORDER_SEQUENCE_NUMBER    Status = 0x04
	Status_INVALID_CONNECTION_ID           Status = 0x21
	Status_CONNECTION_TYPE_NOT_SUPPORTED   Status = 0x22
	Status_CONNECTION_OPTION_NOT_SUPPORTED Status = 0x23
	Status_NO_MORE_CONNECTIONS             Status = 0x24
	Status_NO_MORE_UNIQUE_CONNECTIONS      Status = 0x25
	Status_DATA_CONNECTION                 Status = 0x26
	Status_KNX_CONNECTION                  Status = 0x27
	Status_TUNNELLING_LAYER_NOT_SUPPORTED  Status = 0x29
)

var StatusValues []Status

func init() {
	_ = errors.New
	StatusValues = []Status{
		Status_NO_ERROR,
		Status_PROTOCOL_TYPE_NOT_SUPPORTED,
		Status_UNSUPPORTED_PROTOCOL_VERSION,
		Status_OUT_OF_ORDER_SEQUENCE_NUMBER,
		Status_INVALID_CONNECTION_ID,
		Status_CONNECTION_TYPE_NOT_SUPPORTED,
		Status_CONNECTION_OPTION_NOT_SUPPORTED,
		Status_NO_MORE_CONNECTIONS,
		Status_NO_MORE_UNIQUE_CONNECTIONS,
		Status_DATA_CONNECTION,
		Status_KNX_CONNECTION,
		Status_TUNNELLING_LAYER_NOT_SUPPORTED,
	}
}

func StatusByValue(value uint8) (enum Status, ok bool) {
	switch value {
	case 0x00:
		return Status_NO_ERROR, true
	case 0x01:
		return Status_PROTOCOL_TYPE_NOT_SUPPORTED, true
	case 0x02:
		return Status_UNSUPPORTED_PROTOCOL_VERSION, true
	case 0x04:
		return Status_OUT_OF_ORDER_SEQUENCE_NUMBER, true
	case 0x21:
		return Status_INVALID_CONNECTION_ID, true
	case 0x22:
		return Status_CONNECTION_TYPE_NOT_SUPPORTED, true
	case 0x23:
		return Status_CONNECTION_OPTION_NOT_SUPPORTED, true
	case 0x24:
		return Status_NO_MORE_CONNECTIONS, true
	case 0x25:
		return Status_NO_MORE_UNIQUE_CONNECTIONS, true
	case 0x26:
		return Status_DATA_CONNECTION, true
	case 0x27:
		return Status_KNX_CONNECTION, true
	case 0x29:
		return Status_TUNNELLING_LAYER_NOT_SUPPORTED, true
	}
	return 0, false
}

func StatusByName(value string) (enum Status, ok bool) {
	switch value {
	case "NO_ERROR":
		return Status_NO_ERROR, true
	case "PROTOCOL_TYPE_NOT_SUPPORTED":
		return Status_PROTOCOL_TYPE_NOT_SUPPORTED, true
	case "UNSUPPORTED_PROTOCOL_VERSION":
		return Status_UNSUPPORTED_PROTOCOL_VERSION, true
	case "OUT_OF_ORDER_SEQUENCE_NUMBER":
		return Status_OUT_OF_ORDER_SEQUENCE_NUMBER, true
	case "INVALID_CONNECTION_ID":
		return Status_INVALID_CONNECTION_ID, true
	case "CONNECTION_TYPE_NOT_SUPPORTED":
		return Status_CONNECTION_TYPE_NOT_SUPPORTED, true
	case "CONNECTION_OPTION_NOT_SUPPORTED":
		return Status_CONNECTION_OPTION_NOT_SUPPORTED, true
	case "NO_MORE_CONNECTIONS":
		return Status_NO_MORE_CONNECTIONS, true
	case "NO_MORE_UNIQUE_CONNECTIONS":
		return Status_NO_MORE_UNIQUE_CONNECTIONS, true
	case "DATA_CONNECTION":
		return Status_DATA_CONNECTION, true
	case "KNX_CONNECTION":
		return Status_KNX_CONNECTION, true
	case "TUNNELLING_LAYER_NOT_SUPPORTED":
		return Status_TUNNELLING_LAYER_NOT_SUPPORTED, true
	}
	return 0, false
}

func StatusKnows(value uint8) bool {
	for _, typeValue := range StatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastStatus(structType any) Status {
	castFunc := func(typ any) Status {
		if sStatus, ok := typ.(Status); ok {
			return sStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m Status) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m Status) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusParse(ctx context.Context, theBytes []byte) (Status, error) {
	return StatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Status, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("Status", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Status")
	}
	if enum, ok := StatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for Status")
		return Status(val), nil
	} else {
		return enum, nil
	}
}

func (e Status) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e Status) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("Status", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Status) PLC4XEnumName() string {
	switch e {
	case Status_NO_ERROR:
		return "NO_ERROR"
	case Status_PROTOCOL_TYPE_NOT_SUPPORTED:
		return "PROTOCOL_TYPE_NOT_SUPPORTED"
	case Status_UNSUPPORTED_PROTOCOL_VERSION:
		return "UNSUPPORTED_PROTOCOL_VERSION"
	case Status_OUT_OF_ORDER_SEQUENCE_NUMBER:
		return "OUT_OF_ORDER_SEQUENCE_NUMBER"
	case Status_INVALID_CONNECTION_ID:
		return "INVALID_CONNECTION_ID"
	case Status_CONNECTION_TYPE_NOT_SUPPORTED:
		return "CONNECTION_TYPE_NOT_SUPPORTED"
	case Status_CONNECTION_OPTION_NOT_SUPPORTED:
		return "CONNECTION_OPTION_NOT_SUPPORTED"
	case Status_NO_MORE_CONNECTIONS:
		return "NO_MORE_CONNECTIONS"
	case Status_NO_MORE_UNIQUE_CONNECTIONS:
		return "NO_MORE_UNIQUE_CONNECTIONS"
	case Status_DATA_CONNECTION:
		return "DATA_CONNECTION"
	case Status_KNX_CONNECTION:
		return "KNX_CONNECTION"
	case Status_TUNNELLING_LAYER_NOT_SUPPORTED:
		return "TUNNELLING_LAYER_NOT_SUPPORTED"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e Status) String() string {
	return e.PLC4XEnumName()
}
