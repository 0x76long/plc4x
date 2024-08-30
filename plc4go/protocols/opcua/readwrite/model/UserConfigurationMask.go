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

// UserConfigurationMask is an enum
type UserConfigurationMask uint32

type IUserConfigurationMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	UserConfigurationMask_userConfigurationMaskNone               UserConfigurationMask = 0
	UserConfigurationMask_userConfigurationMaskNoDelete           UserConfigurationMask = 1
	UserConfigurationMask_userConfigurationMaskDisabled           UserConfigurationMask = 2
	UserConfigurationMask_userConfigurationMaskNoChangeByUser     UserConfigurationMask = 4
	UserConfigurationMask_userConfigurationMaskMustChangePassword UserConfigurationMask = 8
)

var UserConfigurationMaskValues []UserConfigurationMask

func init() {
	_ = errors.New
	UserConfigurationMaskValues = []UserConfigurationMask{
		UserConfigurationMask_userConfigurationMaskNone,
		UserConfigurationMask_userConfigurationMaskNoDelete,
		UserConfigurationMask_userConfigurationMaskDisabled,
		UserConfigurationMask_userConfigurationMaskNoChangeByUser,
		UserConfigurationMask_userConfigurationMaskMustChangePassword,
	}
}

func UserConfigurationMaskByValue(value uint32) (enum UserConfigurationMask, ok bool) {
	switch value {
	case 0:
		return UserConfigurationMask_userConfigurationMaskNone, true
	case 1:
		return UserConfigurationMask_userConfigurationMaskNoDelete, true
	case 2:
		return UserConfigurationMask_userConfigurationMaskDisabled, true
	case 4:
		return UserConfigurationMask_userConfigurationMaskNoChangeByUser, true
	case 8:
		return UserConfigurationMask_userConfigurationMaskMustChangePassword, true
	}
	return 0, false
}

func UserConfigurationMaskByName(value string) (enum UserConfigurationMask, ok bool) {
	switch value {
	case "userConfigurationMaskNone":
		return UserConfigurationMask_userConfigurationMaskNone, true
	case "userConfigurationMaskNoDelete":
		return UserConfigurationMask_userConfigurationMaskNoDelete, true
	case "userConfigurationMaskDisabled":
		return UserConfigurationMask_userConfigurationMaskDisabled, true
	case "userConfigurationMaskNoChangeByUser":
		return UserConfigurationMask_userConfigurationMaskNoChangeByUser, true
	case "userConfigurationMaskMustChangePassword":
		return UserConfigurationMask_userConfigurationMaskMustChangePassword, true
	}
	return 0, false
}

func UserConfigurationMaskKnows(value uint32) bool {
	for _, typeValue := range UserConfigurationMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastUserConfigurationMask(structType any) UserConfigurationMask {
	castFunc := func(typ any) UserConfigurationMask {
		if sUserConfigurationMask, ok := typ.(UserConfigurationMask); ok {
			return sUserConfigurationMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m UserConfigurationMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m UserConfigurationMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UserConfigurationMaskParse(ctx context.Context, theBytes []byte) (UserConfigurationMask, error) {
	return UserConfigurationMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UserConfigurationMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UserConfigurationMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("UserConfigurationMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading UserConfigurationMask")
	}
	if enum, ok := UserConfigurationMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for UserConfigurationMask")
		return UserConfigurationMask(val), nil
	} else {
		return enum, nil
	}
}

func (e UserConfigurationMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e UserConfigurationMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("UserConfigurationMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e UserConfigurationMask) PLC4XEnumName() string {
	switch e {
	case UserConfigurationMask_userConfigurationMaskNone:
		return "userConfigurationMaskNone"
	case UserConfigurationMask_userConfigurationMaskNoDelete:
		return "userConfigurationMaskNoDelete"
	case UserConfigurationMask_userConfigurationMaskDisabled:
		return "userConfigurationMaskDisabled"
	case UserConfigurationMask_userConfigurationMaskNoChangeByUser:
		return "userConfigurationMaskNoChangeByUser"
	case UserConfigurationMask_userConfigurationMaskMustChangePassword:
		return "userConfigurationMaskMustChangePassword"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e UserConfigurationMask) String() string {
	return e.PLC4XEnumName()
}
