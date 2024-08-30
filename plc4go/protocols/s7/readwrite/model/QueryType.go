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

// QueryType is an enum
type QueryType uint8

type IQueryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	QueryType_BYALARMTYPE QueryType = 0x01
	QueryType_ALARM_8     QueryType = 0x02
	QueryType_ALARM_S     QueryType = 0x04
	QueryType_ALARM_8P    QueryType = 0x09
)

var QueryTypeValues []QueryType

func init() {
	_ = errors.New
	QueryTypeValues = []QueryType{
		QueryType_BYALARMTYPE,
		QueryType_ALARM_8,
		QueryType_ALARM_S,
		QueryType_ALARM_8P,
	}
}

func QueryTypeByValue(value uint8) (enum QueryType, ok bool) {
	switch value {
	case 0x01:
		return QueryType_BYALARMTYPE, true
	case 0x02:
		return QueryType_ALARM_8, true
	case 0x04:
		return QueryType_ALARM_S, true
	case 0x09:
		return QueryType_ALARM_8P, true
	}
	return 0, false
}

func QueryTypeByName(value string) (enum QueryType, ok bool) {
	switch value {
	case "BYALARMTYPE":
		return QueryType_BYALARMTYPE, true
	case "ALARM_8":
		return QueryType_ALARM_8, true
	case "ALARM_S":
		return QueryType_ALARM_S, true
	case "ALARM_8P":
		return QueryType_ALARM_8P, true
	}
	return 0, false
}

func QueryTypeKnows(value uint8) bool {
	for _, typeValue := range QueryTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastQueryType(structType any) QueryType {
	castFunc := func(typ any) QueryType {
		if sQueryType, ok := typ.(QueryType); ok {
			return sQueryType
		}
		return 0
	}
	return castFunc(structType)
}

func (m QueryType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m QueryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryTypeParse(ctx context.Context, theBytes []byte) (QueryType, error) {
	return QueryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func QueryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (QueryType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("QueryType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading QueryType")
	}
	if enum, ok := QueryTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for QueryType")
		return QueryType(val), nil
	} else {
		return enum, nil
	}
}

func (e QueryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e QueryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("QueryType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e QueryType) PLC4XEnumName() string {
	switch e {
	case QueryType_BYALARMTYPE:
		return "BYALARMTYPE"
	case QueryType_ALARM_8:
		return "ALARM_8"
	case QueryType_ALARM_S:
		return "ALARM_S"
	case QueryType_ALARM_8P:
		return "ALARM_8P"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e QueryType) String() string {
	return e.PLC4XEnumName()
}
