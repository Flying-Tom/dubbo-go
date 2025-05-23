/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rest

import (
	"net/url"
	"strings"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"dubbo.apache.org/dubbo-go/v3/registry/servicediscovery/synthesizer"
)

func init() {
	synthesizer.AddSynthesizer(NewRestSubscribedURLsSynthesizer())
}

// SubscribedURLsSynthesizer implementation for rest protocol
type RestSubscribedURLsSynthesizer struct{}

func (r RestSubscribedURLsSynthesizer) Support(subscribedURL *common.URL) bool {
	return subscribedURL.Protocol == "rest"
}

func (r RestSubscribedURLsSynthesizer) Synthesize(subscribedURL *common.URL, serviceInstances []registry.ServiceInstance) []*common.URL {
	urls := make([]*common.URL, len(serviceInstances))
	for i, s := range serviceInstances {
		splitHost := strings.Split(s.GetHost(), ":")
		u := common.NewURLWithOptions(common.WithProtocol(subscribedURL.Protocol), common.WithIp(splitHost[0]),
			common.WithPort(splitHost[1]), common.WithPath(subscribedURL.GetParam(constant.InterfaceKey, subscribedURL.Path)),
			common.WithParams(url.Values{}),
			common.WithParamsValue(constant.SideKey, constant.ProviderProtocol),
			common.WithParamsValue(constant.ApplicationKey, s.GetServiceName()),
			common.WithParamsValue(constant.RegistryKey, "true"),
		)
		urls[i] = u
	}
	return urls
}

func NewRestSubscribedURLsSynthesizer() RestSubscribedURLsSynthesizer {
	return RestSubscribedURLsSynthesizer{}
}
