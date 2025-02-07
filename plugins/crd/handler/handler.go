// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"github.com/contiv/vpp/plugins/crd/handler/customnetwork"
	"github.com/contiv/vpp/plugins/crd/handler/nodeconfig"
	"github.com/contiv/vpp/plugins/crd/handler/servicefunctionchain"
	"github.com/contiv/vpp/plugins/crd/handler/telemetry"
)

// Handler is implemented by any handler.
// The Handle method is used to process an event
type Handler interface {
	//
	Init() error
	//
	ObjectCreated(obj interface{})
	//
	ObjectDeleted(obj interface{})
	//
	ObjectUpdated(oldObj, newObj interface{})
}

// Map maps each event handler function to a name for easily lookup
var Map = map[string]interface{}{
	"default":              &Default{},
	"telemetry":            &telemetry.Handler{},
	"nodeConfig":           &nodeconfig.Handler{},
	"customNetwork":        &customnetwork.Handler{},
	"serviceFunctionChain": &servicefunctionchain.Handler{},
}

// Default handler implements Handler interface
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init() error {
	return nil
}

// ObjectCreated is called when a CRD object is created
func (d *Default) ObjectCreated(obj interface{}) {
}

// ObjectDeleted is called when a CRD object is deleted
func (d *Default) ObjectDeleted(obj interface{}) {
}

// ObjectUpdated is called when a CRD object is updated
func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {
}
