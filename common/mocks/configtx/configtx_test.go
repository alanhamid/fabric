/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configtx

import (
	"testing"

	configtxapi "github.com/hyperledger/fabric/common/configtx/api"
)

func TestConfigtxTransactionalInterface(t *testing.T) {
	_ = configtxapi.Transactional(&Transactional{})
}

func TestConfigtxPolicyHandlerInterface(t *testing.T) {
	_ = configtxapi.PolicyHandler(&PolicyHandler{})
}

func TestConfigtxInitializerInterface(t *testing.T) {
	_ = configtxapi.Initializer(&Initializer{})
}

func TestConfigtxManagerInterface(t *testing.T) {
	_ = configtxapi.Manager(&Manager{})
}

func TestConfigtxResourcesInterface(t *testing.T) {
	_ = configtxapi.Resources(&Resources{})
}

func TestConfigtxHandlerInterface(t *testing.T) {
	_ = configtxapi.Handler(&Handler{})
}