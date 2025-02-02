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

package _default

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

func TestNewDefaultBrowser(t *testing.T) {
	type args struct {
		defaultBrowserRequirements DefaultBrowserRequirements
	}
	tests := []struct {
		name string
		args args
		want DefaultBrowser
	}{
		{
			name: "just create it",
			want: &defaultBrowser{
				log: log.Logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultBrowser(tt.args.defaultBrowserRequirements); !assert.Equal(t, tt.want, got) {
				t.Errorf("NewDefaultBrowser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultBrowser_Browse(t *testing.T) {
	type fields struct {
		DefaultBrowserRequirements DefaultBrowserRequirements
	}
	type args struct {
		ctx           context.Context
		browseRequest apiModel.PlcBrowseRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantAsserter func(t *testing.T, results <-chan apiModel.PlcBrowseRequestResult) bool
	}{
		{
			name: "Browse empty",
			args: args{
				ctx:           context.Background(),
				browseRequest: spiModel.NewDefaultPlcBrowseRequest(nil, nil, nil),
			},
			wantAsserter: func(t *testing.T, results <-chan apiModel.PlcBrowseRequestResult) bool {
				timeout := time.NewTimer(2 * time.Second)
				defer utils.CleanupTimer(timeout)
				select {
				case result := <-results:
					assert.NotNil(t, result)
				case <-timeout.C:
					t.Error("timeout")
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &defaultBrowser{
				DefaultBrowserRequirements: tt.fields.DefaultBrowserRequirements,
			}
			assert.Truef(t, tt.wantAsserter(t, m.Browse(tt.args.ctx, tt.args.browseRequest)), "Browse(%v, %v)", tt.args.ctx, tt.args.browseRequest)
		})
	}
}

func Test_defaultBrowser_BrowseWithInterceptor(t *testing.T) {
	type fields struct {
		DefaultBrowserRequirements DefaultBrowserRequirements
	}
	type args struct {
		ctx           context.Context
		browseRequest apiModel.PlcBrowseRequest
		interceptor   func(result apiModel.PlcBrowseItem) bool
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		setup        func(t *testing.T, fields *fields, args *args)
		wantAsserter func(t *testing.T, results <-chan apiModel.PlcBrowseRequestResult) bool
	}{
		{
			name: "Browse empty",
			args: args{
				ctx:           context.Background(),
				browseRequest: spiModel.NewDefaultPlcBrowseRequest(nil, nil, nil),
			},
			wantAsserter: func(t *testing.T, results <-chan apiModel.PlcBrowseRequestResult) bool {
				timeout := time.NewTimer(2 * time.Second)
				defer utils.CleanupTimer(timeout)
				select {
				case result := <-results:
					assert.NotNil(t, result)
				case <-timeout.C:
					t.Error("timeout")
					return false
				}
				return true
			},
		},
		{
			name: "Browse something",
			args: args{
				ctx: func() context.Context {
					timeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
					t.Cleanup(cancelFunc)
					return timeout
				}(),
			},
			setup: func(t *testing.T, fields *fields, args *args) {
				requirements := NewMockDefaultBrowserRequirements(t)
				requirements.EXPECT().BrowseQuery(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(0, nil)
				fields.DefaultBrowserRequirements = requirements
				query := NewMockPlcQuery(t)
				args.browseRequest = spiModel.NewDefaultPlcBrowseRequest(
					map[string]apiModel.PlcQuery{
						"test": query,
					},
					[]string{"test"},
					nil,
				)
			},
			wantAsserter: func(t *testing.T, results <-chan apiModel.PlcBrowseRequestResult) bool {
				timeout := time.NewTimer(2 * time.Second)
				defer utils.CleanupTimer(timeout)
				select {
				case result := <-results:
					assert.NotNil(t, result)
				case <-timeout.C:
					t.Error("timeout")
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.fields, &tt.args)
			}
			m := &defaultBrowser{
				DefaultBrowserRequirements: tt.fields.DefaultBrowserRequirements,
			}
			assert.Truef(t, tt.wantAsserter(t, m.Browse(tt.args.ctx, tt.args.browseRequest)), "BrowseWithInterceptor(%v, %v, func())", tt.args.ctx, tt.args.browseRequest)
		})
	}
}
