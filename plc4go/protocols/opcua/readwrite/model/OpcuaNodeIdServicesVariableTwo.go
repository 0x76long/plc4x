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

// OpcuaNodeIdServicesVariableTwo is an enum
type OpcuaNodeIdServicesVariableTwo int32

type IOpcuaNodeIdServicesVariableTwo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_FalseState              OpcuaNodeIdServicesVariableTwo = 2374
	OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_TrueState               OpcuaNodeIdServicesVariableTwo = 2375
	OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_Id                      OpcuaNodeIdServicesVariableTwo = 8996
	OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TransitionTime          OpcuaNodeIdServicesVariableTwo = 9000
	OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_EffectiveTransitionTime OpcuaNodeIdServicesVariableTwo = 9001
	OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TrueState               OpcuaNodeIdServicesVariableTwo = 11110
	OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_FalseState              OpcuaNodeIdServicesVariableTwo = 11111
)

var OpcuaNodeIdServicesVariableTwoValues []OpcuaNodeIdServicesVariableTwo

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTwoValues = []OpcuaNodeIdServicesVariableTwo{
		OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_FalseState,
		OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_TrueState,
		OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_Id,
		OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TransitionTime,
		OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TrueState,
		OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_FalseState,
	}
}

func OpcuaNodeIdServicesVariableTwoByValue(value int32) (enum OpcuaNodeIdServicesVariableTwo, ok bool) {
	switch value {
	case 11110:
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TrueState, true
	case 11111:
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_FalseState, true
	case 2374:
		return OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_FalseState, true
	case 2375:
		return OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_TrueState, true
	case 8996:
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_Id, true
	case 9000:
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TransitionTime, true
	case 9001:
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_EffectiveTransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTwoByName(value string) (enum OpcuaNodeIdServicesVariableTwo, ok bool) {
	switch value {
	case "TwoStateVariableType_TrueState":
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TrueState, true
	case "TwoStateVariableType_FalseState":
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_FalseState, true
	case "TwoStateDiscreteType_FalseState":
		return OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_FalseState, true
	case "TwoStateDiscreteType_TrueState":
		return OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_TrueState, true
	case "TwoStateVariableType_Id":
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_Id, true
	case "TwoStateVariableType_TransitionTime":
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TransitionTime, true
	case "TwoStateVariableType_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_EffectiveTransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTwoKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTwoValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTwo(structType any) OpcuaNodeIdServicesVariableTwo {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTwo {
		if sOpcuaNodeIdServicesVariableTwo, ok := typ.(OpcuaNodeIdServicesVariableTwo); ok {
			return sOpcuaNodeIdServicesVariableTwo
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTwo) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTwo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTwoParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTwo, error) {
	return OpcuaNodeIdServicesVariableTwoParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTwoParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTwo, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTwo", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTwo")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTwoByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTwo")
		return OpcuaNodeIdServicesVariableTwo(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTwo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTwo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTwo", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTwo) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TrueState:
		return "TwoStateVariableType_TrueState"
	case OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_FalseState:
		return "TwoStateVariableType_FalseState"
	case OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_FalseState:
		return "TwoStateDiscreteType_FalseState"
	case OpcuaNodeIdServicesVariableTwo_TwoStateDiscreteType_TrueState:
		return "TwoStateDiscreteType_TrueState"
	case OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_Id:
		return "TwoStateVariableType_Id"
	case OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_TransitionTime:
		return "TwoStateVariableType_TransitionTime"
	case OpcuaNodeIdServicesVariableTwo_TwoStateVariableType_EffectiveTransitionTime:
		return "TwoStateVariableType_EffectiveTransitionTime"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTwo) String() string {
	return e.PLC4XEnumName()
}
