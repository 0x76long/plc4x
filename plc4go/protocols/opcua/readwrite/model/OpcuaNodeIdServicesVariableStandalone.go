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

// OpcuaNodeIdServicesVariableStandalone is an enum
type OpcuaNodeIdServicesVariableStandalone int32

type IOpcuaNodeIdServicesVariableStandalone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_DataSetMetaData OpcuaNodeIdServicesVariableStandalone = 23830
	OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_IsConnected     OpcuaNodeIdServicesVariableStandalone = 23831
)

var OpcuaNodeIdServicesVariableStandaloneValues []OpcuaNodeIdServicesVariableStandalone

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableStandaloneValues = []OpcuaNodeIdServicesVariableStandalone{
		OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_DataSetMetaData,
		OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_IsConnected,
	}
}

func OpcuaNodeIdServicesVariableStandaloneByValue(value int32) (enum OpcuaNodeIdServicesVariableStandalone, ok bool) {
	switch value {
	case 23830:
		return OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_DataSetMetaData, true
	case 23831:
		return OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_IsConnected, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStandaloneByName(value string) (enum OpcuaNodeIdServicesVariableStandalone, ok bool) {
	switch value {
	case "StandaloneSubscribedDataSetType_DataSetMetaData":
		return OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_DataSetMetaData, true
	case "StandaloneSubscribedDataSetType_IsConnected":
		return OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_IsConnected, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStandaloneKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableStandaloneValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableStandalone(structType any) OpcuaNodeIdServicesVariableStandalone {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableStandalone {
		if sOpcuaNodeIdServicesVariableStandalone, ok := typ.(OpcuaNodeIdServicesVariableStandalone); ok {
			return sOpcuaNodeIdServicesVariableStandalone
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableStandalone) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableStandalone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableStandaloneParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableStandalone, error) {
	return OpcuaNodeIdServicesVariableStandaloneParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableStandaloneParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableStandalone, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableStandalone", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableStandalone")
	}
	if enum, ok := OpcuaNodeIdServicesVariableStandaloneByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableStandalone")
		return OpcuaNodeIdServicesVariableStandalone(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableStandalone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableStandalone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableStandalone", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableStandalone) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_DataSetMetaData:
		return "StandaloneSubscribedDataSetType_DataSetMetaData"
	case OpcuaNodeIdServicesVariableStandalone_StandaloneSubscribedDataSetType_IsConnected:
		return "StandaloneSubscribedDataSetType_IsConnected"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableStandalone) String() string {
	return e.PLC4XEnumName()
}
