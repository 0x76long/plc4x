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

// BACnetConfirmedServiceChoice is an enum
type BACnetConfirmedServiceChoice uint8

type IBACnetConfirmedServiceChoice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM                   BACnetConfirmedServiceChoice = 0x00
	BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION          BACnetConfirmedServiceChoice = 0x01
	BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE BACnetConfirmedServiceChoice = 0x1F
	BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION        BACnetConfirmedServiceChoice = 0x02
	BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY                   BACnetConfirmedServiceChoice = 0x03
	BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY              BACnetConfirmedServiceChoice = 0x04
	BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION               BACnetConfirmedServiceChoice = 0x1D
	BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION               BACnetConfirmedServiceChoice = 0x1B
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV                       BACnetConfirmedServiceChoice = 0x05
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY              BACnetConfirmedServiceChoice = 0x1C
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE     BACnetConfirmedServiceChoice = 0x1E
	BACnetConfirmedServiceChoice_ATOMIC_READ_FILE                    BACnetConfirmedServiceChoice = 0x06
	BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE                   BACnetConfirmedServiceChoice = 0x07
	BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT                    BACnetConfirmedServiceChoice = 0x08
	BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT                 BACnetConfirmedServiceChoice = 0x09
	BACnetConfirmedServiceChoice_CREATE_OBJECT                       BACnetConfirmedServiceChoice = 0x0A
	BACnetConfirmedServiceChoice_DELETE_OBJECT                       BACnetConfirmedServiceChoice = 0x0B
	BACnetConfirmedServiceChoice_READ_PROPERTY                       BACnetConfirmedServiceChoice = 0x0C
	BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE              BACnetConfirmedServiceChoice = 0x0E
	BACnetConfirmedServiceChoice_READ_RANGE                          BACnetConfirmedServiceChoice = 0x1A
	BACnetConfirmedServiceChoice_WRITE_PROPERTY                      BACnetConfirmedServiceChoice = 0x0F
	BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE             BACnetConfirmedServiceChoice = 0x10
	BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL        BACnetConfirmedServiceChoice = 0x11
	BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER          BACnetConfirmedServiceChoice = 0x12
	BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE              BACnetConfirmedServiceChoice = 0x13
	BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE                 BACnetConfirmedServiceChoice = 0x14
	BACnetConfirmedServiceChoice_VT_OPEN                             BACnetConfirmedServiceChoice = 0x15
	BACnetConfirmedServiceChoice_VT_CLOSE                            BACnetConfirmedServiceChoice = 0x16
	BACnetConfirmedServiceChoice_VT_DATA                             BACnetConfirmedServiceChoice = 0x17
	BACnetConfirmedServiceChoice_AUTHENTICATE                        BACnetConfirmedServiceChoice = 0x18
	BACnetConfirmedServiceChoice_REQUEST_KEY                         BACnetConfirmedServiceChoice = 0x19
	BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL           BACnetConfirmedServiceChoice = 0x0D
)

var BACnetConfirmedServiceChoiceValues []BACnetConfirmedServiceChoice

func init() {
	_ = errors.New
	BACnetConfirmedServiceChoiceValues = []BACnetConfirmedServiceChoice{
		BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM,
		BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION,
		BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE,
		BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION,
		BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY,
		BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY,
		BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION,
		BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_ATOMIC_READ_FILE,
		BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE,
		BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT,
		BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT,
		BACnetConfirmedServiceChoice_CREATE_OBJECT,
		BACnetConfirmedServiceChoice_DELETE_OBJECT,
		BACnetConfirmedServiceChoice_READ_PROPERTY,
		BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_READ_RANGE,
		BACnetConfirmedServiceChoice_WRITE_PROPERTY,
		BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL,
		BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER,
		BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE,
		BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE,
		BACnetConfirmedServiceChoice_VT_OPEN,
		BACnetConfirmedServiceChoice_VT_CLOSE,
		BACnetConfirmedServiceChoice_VT_DATA,
		BACnetConfirmedServiceChoice_AUTHENTICATE,
		BACnetConfirmedServiceChoice_REQUEST_KEY,
		BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL,
	}
}

func BACnetConfirmedServiceChoiceByValue(value uint8) (enum BACnetConfirmedServiceChoice, ok bool) {
	switch value {
	case 0x00:
		return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM, true
	case 0x01:
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION, true
	case 0x02:
		return BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION, true
	case 0x03:
		return BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY, true
	case 0x04:
		return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY, true
	case 0x05:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV, true
	case 0x06:
		return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE, true
	case 0x07:
		return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE, true
	case 0x08:
		return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT, true
	case 0x09:
		return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT, true
	case 0x0A:
		return BACnetConfirmedServiceChoice_CREATE_OBJECT, true
	case 0x0B:
		return BACnetConfirmedServiceChoice_DELETE_OBJECT, true
	case 0x0C:
		return BACnetConfirmedServiceChoice_READ_PROPERTY, true
	case 0x0D:
		return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL, true
	case 0x0E:
		return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE, true
	case 0x0F:
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY, true
	case 0x10:
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE, true
	case 0x11:
		return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL, true
	case 0x12:
		return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER, true
	case 0x13:
		return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE, true
	case 0x14:
		return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE, true
	case 0x15:
		return BACnetConfirmedServiceChoice_VT_OPEN, true
	case 0x16:
		return BACnetConfirmedServiceChoice_VT_CLOSE, true
	case 0x17:
		return BACnetConfirmedServiceChoice_VT_DATA, true
	case 0x18:
		return BACnetConfirmedServiceChoice_AUTHENTICATE, true
	case 0x19:
		return BACnetConfirmedServiceChoice_REQUEST_KEY, true
	case 0x1A:
		return BACnetConfirmedServiceChoice_READ_RANGE, true
	case 0x1B:
		return BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION, true
	case 0x1C:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY, true
	case 0x1D:
		return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION, true
	case 0x1E:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE, true
	case 0x1F:
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	}
	return 0, false
}

func BACnetConfirmedServiceChoiceByName(value string) (enum BACnetConfirmedServiceChoice, ok bool) {
	switch value {
	case "ACKNOWLEDGE_ALARM":
		return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM, true
	case "CONFIRMED_COV_NOTIFICATION":
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION, true
	case "CONFIRMED_EVENT_NOTIFICATION":
		return BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION, true
	case "GET_ALARM_SUMMARY":
		return BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY, true
	case "GET_ENROLLMENT_SUMMARY":
		return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY, true
	case "SUBSCRIBE_COV":
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV, true
	case "ATOMIC_READ_FILE":
		return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE, true
	case "ATOMIC_WRITE_FILE":
		return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE, true
	case "ADD_LIST_ELEMENT":
		return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT, true
	case "REMOVE_LIST_ELEMENT":
		return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT, true
	case "CREATE_OBJECT":
		return BACnetConfirmedServiceChoice_CREATE_OBJECT, true
	case "DELETE_OBJECT":
		return BACnetConfirmedServiceChoice_DELETE_OBJECT, true
	case "READ_PROPERTY":
		return BACnetConfirmedServiceChoice_READ_PROPERTY, true
	case "READ_PROPERTY_CONDITIONAL":
		return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL, true
	case "READ_PROPERTY_MULTIPLE":
		return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE, true
	case "WRITE_PROPERTY":
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY, true
	case "WRITE_PROPERTY_MULTIPLE":
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE, true
	case "DEVICE_COMMUNICATION_CONTROL":
		return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL, true
	case "CONFIRMED_PRIVATE_TRANSFER":
		return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER, true
	case "CONFIRMED_TEXT_MESSAGE":
		return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE, true
	case "REINITIALIZE_DEVICE":
		return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE, true
	case "VT_OPEN":
		return BACnetConfirmedServiceChoice_VT_OPEN, true
	case "VT_CLOSE":
		return BACnetConfirmedServiceChoice_VT_CLOSE, true
	case "VT_DATA":
		return BACnetConfirmedServiceChoice_VT_DATA, true
	case "AUTHENTICATE":
		return BACnetConfirmedServiceChoice_AUTHENTICATE, true
	case "REQUEST_KEY":
		return BACnetConfirmedServiceChoice_REQUEST_KEY, true
	case "READ_RANGE":
		return BACnetConfirmedServiceChoice_READ_RANGE, true
	case "LIFE_SAFETY_OPERATION":
		return BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION, true
	case "SUBSCRIBE_COV_PROPERTY":
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY, true
	case "GET_EVENT_INFORMATION":
		return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION, true
	case "SUBSCRIBE_COV_PROPERTY_MULTIPLE":
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE, true
	case "CONFIRMED_COV_NOTIFICATION_MULTIPLE":
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	}
	return 0, false
}

func BACnetConfirmedServiceChoiceKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceChoiceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceChoice(structType any) BACnetConfirmedServiceChoice {
	castFunc := func(typ any) BACnetConfirmedServiceChoice {
		if sBACnetConfirmedServiceChoice, ok := typ.(BACnetConfirmedServiceChoice); ok {
			return sBACnetConfirmedServiceChoice
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceChoice) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetConfirmedServiceChoice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceChoiceParse(ctx context.Context, theBytes []byte) (BACnetConfirmedServiceChoice, error) {
	return BACnetConfirmedServiceChoiceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceChoiceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceChoice, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetConfirmedServiceChoice", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetConfirmedServiceChoice")
	}
	if enum, ok := BACnetConfirmedServiceChoiceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetConfirmedServiceChoice")
		return BACnetConfirmedServiceChoice(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetConfirmedServiceChoice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetConfirmedServiceChoice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetConfirmedServiceChoice", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceChoice) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM:
		return "ACKNOWLEDGE_ALARM"
	case BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION:
		return "CONFIRMED_COV_NOTIFICATION"
	case BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION:
		return "CONFIRMED_EVENT_NOTIFICATION"
	case BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY:
		return "GET_ALARM_SUMMARY"
	case BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY:
		return "GET_ENROLLMENT_SUMMARY"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV:
		return "SUBSCRIBE_COV"
	case BACnetConfirmedServiceChoice_ATOMIC_READ_FILE:
		return "ATOMIC_READ_FILE"
	case BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE:
		return "ATOMIC_WRITE_FILE"
	case BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT:
		return "ADD_LIST_ELEMENT"
	case BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT:
		return "REMOVE_LIST_ELEMENT"
	case BACnetConfirmedServiceChoice_CREATE_OBJECT:
		return "CREATE_OBJECT"
	case BACnetConfirmedServiceChoice_DELETE_OBJECT:
		return "DELETE_OBJECT"
	case BACnetConfirmedServiceChoice_READ_PROPERTY:
		return "READ_PROPERTY"
	case BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL:
		return "READ_PROPERTY_CONDITIONAL"
	case BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE:
		return "READ_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_WRITE_PROPERTY:
		return "WRITE_PROPERTY"
	case BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE:
		return "WRITE_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL:
		return "DEVICE_COMMUNICATION_CONTROL"
	case BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER:
		return "CONFIRMED_PRIVATE_TRANSFER"
	case BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE:
		return "CONFIRMED_TEXT_MESSAGE"
	case BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE:
		return "REINITIALIZE_DEVICE"
	case BACnetConfirmedServiceChoice_VT_OPEN:
		return "VT_OPEN"
	case BACnetConfirmedServiceChoice_VT_CLOSE:
		return "VT_CLOSE"
	case BACnetConfirmedServiceChoice_VT_DATA:
		return "VT_DATA"
	case BACnetConfirmedServiceChoice_AUTHENTICATE:
		return "AUTHENTICATE"
	case BACnetConfirmedServiceChoice_REQUEST_KEY:
		return "REQUEST_KEY"
	case BACnetConfirmedServiceChoice_READ_RANGE:
		return "READ_RANGE"
	case BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION:
		return "LIFE_SAFETY_OPERATION"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY:
		return "SUBSCRIBE_COV_PROPERTY"
	case BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION:
		return "GET_EVENT_INFORMATION"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE:
		return "SUBSCRIBE_COV_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE:
		return "CONFIRMED_COV_NOTIFICATION_MULTIPLE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetConfirmedServiceChoice) String() string {
	return e.PLC4XEnumName()
}
