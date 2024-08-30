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

// OpcuaNodeIdServicesVariableRead is an enum
type OpcuaNodeIdServicesVariableRead int32

type IOpcuaNodeIdServicesVariableRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableRead_ReadMethodType_InputArguments  OpcuaNodeIdServicesVariableRead = 11744
	OpcuaNodeIdServicesVariableRead_ReadMethodType_OutputArguments OpcuaNodeIdServicesVariableRead = 11745
)

var OpcuaNodeIdServicesVariableReadValues []OpcuaNodeIdServicesVariableRead

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableReadValues = []OpcuaNodeIdServicesVariableRead{
		OpcuaNodeIdServicesVariableRead_ReadMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRead_ReadMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableReadByValue(value int32) (enum OpcuaNodeIdServicesVariableRead, ok bool) {
	switch value {
	case 11744:
		return OpcuaNodeIdServicesVariableRead_ReadMethodType_InputArguments, true
	case 11745:
		return OpcuaNodeIdServicesVariableRead_ReadMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableReadByName(value string) (enum OpcuaNodeIdServicesVariableRead, ok bool) {
	switch value {
	case "ReadMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRead_ReadMethodType_InputArguments, true
	case "ReadMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableRead_ReadMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableReadKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableReadValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableRead(structType any) OpcuaNodeIdServicesVariableRead {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRead {
		if sOpcuaNodeIdServicesVariableRead, ok := typ.(OpcuaNodeIdServicesVariableRead); ok {
			return sOpcuaNodeIdServicesVariableRead
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRead) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableReadParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRead, error) {
	return OpcuaNodeIdServicesVariableReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRead, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRead", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRead")
	}
	if enum, ok := OpcuaNodeIdServicesVariableReadByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRead")
		return OpcuaNodeIdServicesVariableRead(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRead", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRead) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRead_ReadMethodType_InputArguments:
		return "ReadMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRead_ReadMethodType_OutputArguments:
		return "ReadMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRead) String() string {
	return e.PLC4XEnumName()
}
