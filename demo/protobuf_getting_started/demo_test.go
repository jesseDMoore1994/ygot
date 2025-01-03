//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"

	ocpb "github.com/openconfig/ygot/demo/protobuf_getting_started/ribproto/openconfig"
	ocepb "github.com/openconfig/ygot/demo/protobuf_getting_started/ribproto/openconfig/enums"
)

func TestProtoGenerate(t *testing.T) {
	tests := []struct {
		name          string
		inTestFunc    func(*ipv4Prefix) *ocpb.Device
		inPrefix      *ipv4Prefix
		wantTextProto string
	}{{
		name:       "simple route entry test",
		inTestFunc: buildRouteProto,
		inPrefix: &ipv4Prefix{
			atomicAggregate: true,
			localPref:       100,
			med:             10,
			nextHop:         "10.0.1.1",
			origin:          ocepb.OpenconfigBgpTypesBgpOriginAttrType_OPENCONFIGBGPTYPESBGPORIGINATTRTYPE_EGP,
			originatorID:    "192.0.2.42",
			prefix:          "192.0.2.0/24",
			protocolOrigin:  ocepb.OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_BGP,
		},
		wantTextProto: "route_entry.txtpb",
	}}

	for _, tt := range tests {
		got := tt.inTestFunc(tt.inPrefix)

		want := &ocpb.Device{}

		wantStr, err := os.ReadFile(filepath.Join("testdata", tt.wantTextProto))
		if err != nil {
			t.Errorf("%s: os.ReadFile(testdata/%s): could not read file, got: %v, want: nil", tt.name, tt.wantTextProto, err)
		}

		if err := prototext.Unmarshal(wantStr, want); err != nil {
			t.Errorf("%s: prototext.Unmarshal(file: %s): could not unmarshal test proto, got: %v, want: nil", tt.name, tt.wantTextProto, err)
		}

		if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
			t.Errorf("%s: %T: did not get expected return proto, diff(-got,+want):\n%s", tt.name, tt.inTestFunc, diff)
		}
	}
}
