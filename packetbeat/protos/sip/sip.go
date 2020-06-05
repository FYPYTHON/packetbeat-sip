// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package sip

import (
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/monitoring"

	"github.com/elastic/beats/v7/packetbeat/protos"
)

var (
	debugf = logp.MakeDebug("sip")
)

// Packetbeats monitoring metrics
var (
	messageIgnored = monitoring.NewInt(nil, "sip.message_ignored")
)

const (
	transportTCP = 0
	transportUDP = 1
)

const (
	SIP_STATUS_RECEIVED         = 0
	SIP_STATUS_HEADER_RECEIVING = 1
	SIP_STATUS_BODY_RECEIVING   = 2
	SIP_STATUS_REJECTED         = 3
)

const (
	SIP_DETAIL_URI            = 1
	SIP_DETAIL_NAME_ADDR      = 2
	SIP_DETAIL_INT            = 3
	SIP_DETAIL_INT_METHOD     = 4
	SIP_DETAIL_INT_INT_METHOD = 5
	SIP_DETAIL_INT_STRING     = 6
	SIP_DETAIL_INT_INT        = 7
	SIP_DETAIL_INT_INT_STRING = 8
)

func init() {
	// Memo: Secound argment*New* is below New function.
	protos.Register("sip", New)
}

// New create a sip plugin
func New(
	testMode bool,
	results protos.Reporter,
	cfg *common.Config,
) (protos.Plugin, error) {
	p := &sipPlugin{}
	config := defaultConfig
	if !testMode {
		if err := cfg.Unpack(&config); err != nil {
			return nil, err
		}
	}

	if err := p.init(results, &config); err != nil {
		return nil, err
	}
	return p, nil
}

func getLastElementStrArray(array []common.NetString) common.NetString {
	return array[len(array)-1]
}

/**
 ******************************************************************
 * transport
 *******************************************************************
 **/

// Transport protocol.
// transport=0 tcp, transport=1, udp
type transport uint8

func (t transport) String() string {

	transportNames := []string{
		"tcp",
		"udp",
	}

	if int(t) >= len(transportNames) {
		return "impossible"
	}
	return transportNames[t]
}
