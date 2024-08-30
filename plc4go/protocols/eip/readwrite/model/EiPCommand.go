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

// EiPCommand is an enum
type EiPCommand uint16

type IEiPCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	EiPCommand_RegisterSession   EiPCommand = 0x0065
	EiPCommand_UnregisterSession EiPCommand = 0x0066
	EiPCommand_SendRRData        EiPCommand = 0x006F
)

var EiPCommandValues []EiPCommand

func init() {
	_ = errors.New
	EiPCommandValues = []EiPCommand{
		EiPCommand_RegisterSession,
		EiPCommand_UnregisterSession,
		EiPCommand_SendRRData,
	}
}

func EiPCommandByValue(value uint16) (enum EiPCommand, ok bool) {
	switch value {
	case 0x0065:
		return EiPCommand_RegisterSession, true
	case 0x0066:
		return EiPCommand_UnregisterSession, true
	case 0x006F:
		return EiPCommand_SendRRData, true
	}
	return 0, false
}

func EiPCommandByName(value string) (enum EiPCommand, ok bool) {
	switch value {
	case "RegisterSession":
		return EiPCommand_RegisterSession, true
	case "UnregisterSession":
		return EiPCommand_UnregisterSession, true
	case "SendRRData":
		return EiPCommand_SendRRData, true
	}
	return 0, false
}

func EiPCommandKnows(value uint16) bool {
	for _, typeValue := range EiPCommandValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastEiPCommand(structType any) EiPCommand {
	castFunc := func(typ any) EiPCommand {
		if sEiPCommand, ok := typ.(EiPCommand); ok {
			return sEiPCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m EiPCommand) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m EiPCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EiPCommandParse(ctx context.Context, theBytes []byte) (EiPCommand, error) {
	return EiPCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EiPCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EiPCommand, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("EiPCommand", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading EiPCommand")
	}
	if enum, ok := EiPCommandByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for EiPCommand")
		return EiPCommand(val), nil
	} else {
		return enum, nil
	}
}

func (e EiPCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e EiPCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("EiPCommand", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e EiPCommand) PLC4XEnumName() string {
	switch e {
	case EiPCommand_RegisterSession:
		return "RegisterSession"
	case EiPCommand_UnregisterSession:
		return "UnregisterSession"
	case EiPCommand_SendRRData:
		return "SendRRData"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e EiPCommand) String() string {
	return e.PLC4XEnumName()
}
