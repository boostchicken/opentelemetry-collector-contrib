// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jaegerthrifthttpexporter

import (
	"errors"
)

const (
	// Jaeger Tags
	ocTimeEventUnknownType    = "oc.timeevent.unknown.type"
	opencensusLanguage        = "opencensus.language"
	opencensusExporterVersion = "opencensus.exporterversion"
	opencensusCoreLibVersion  = "opencensus.corelibversion"
	opencensusResourceType    = "opencensus.resourcetype"
)

var (
	errZeroTraceID    = errors.New("OC span has an all zeros trace ID")
	errInvalidTraceID = errors.New("OC span has a non 16 byte array trace ID")
	errZeroSpanID     = errors.New("OC span has an all zeros span ID")
	errInvalidSpanID  = errors.New("OC span has a non 16 byte array span ID")
)
