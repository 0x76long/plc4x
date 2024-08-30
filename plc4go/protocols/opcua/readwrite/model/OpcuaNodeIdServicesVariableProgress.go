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

// OpcuaNodeIdServicesVariableProgress is an enum
type OpcuaNodeIdServicesVariableProgress int32

type IOpcuaNodeIdServicesVariableProgress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableProgress_ProgressEventType_Context  OpcuaNodeIdServicesVariableProgress = 12502
	OpcuaNodeIdServicesVariableProgress_ProgressEventType_Progress OpcuaNodeIdServicesVariableProgress = 12503
)

var OpcuaNodeIdServicesVariableProgressValues []OpcuaNodeIdServicesVariableProgress

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableProgressValues = []OpcuaNodeIdServicesVariableProgress{
		OpcuaNodeIdServicesVariableProgress_ProgressEventType_Context,
		OpcuaNodeIdServicesVariableProgress_ProgressEventType_Progress,
	}
}

func OpcuaNodeIdServicesVariableProgressByValue(value int32) (enum OpcuaNodeIdServicesVariableProgress, ok bool) {
	switch value {
	case 12502:
		return OpcuaNodeIdServicesVariableProgress_ProgressEventType_Context, true
	case 12503:
		return OpcuaNodeIdServicesVariableProgress_ProgressEventType_Progress, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableProgressByName(value string) (enum OpcuaNodeIdServicesVariableProgress, ok bool) {
	switch value {
	case "ProgressEventType_Context":
		return OpcuaNodeIdServicesVariableProgress_ProgressEventType_Context, true
	case "ProgressEventType_Progress":
		return OpcuaNodeIdServicesVariableProgress_ProgressEventType_Progress, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableProgressKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableProgressValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableProgress(structType any) OpcuaNodeIdServicesVariableProgress {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableProgress {
		if sOpcuaNodeIdServicesVariableProgress, ok := typ.(OpcuaNodeIdServicesVariableProgress); ok {
			return sOpcuaNodeIdServicesVariableProgress
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableProgress) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableProgress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableProgressParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableProgress, error) {
	return OpcuaNodeIdServicesVariableProgressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableProgressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableProgress, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableProgress", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableProgress")
	}
	if enum, ok := OpcuaNodeIdServicesVariableProgressByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableProgress")
		return OpcuaNodeIdServicesVariableProgress(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableProgress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableProgress", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableProgress) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableProgress_ProgressEventType_Context:
		return "ProgressEventType_Context"
	case OpcuaNodeIdServicesVariableProgress_ProgressEventType_Progress:
		return "ProgressEventType_Progress"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableProgress) String() string {
	return e.PLC4XEnumName()
}
