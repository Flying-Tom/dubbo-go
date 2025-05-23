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

package graceful_shutdown

import (
	"time"
)

import (
	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

func init() {
	hessian.RegisterPOJO(&common.URL{})
}

func parseDuration(timeout string, desc string, def time.Duration) time.Duration {
	res, err := time.ParseDuration(timeout)
	if err != nil {
		logger.Errorf("The %s configuration is invalid: %s, and we will use the default value: %s, err: %v",
			desc, timeout, def.String(), err)
		res = def
	}

	return res
}
