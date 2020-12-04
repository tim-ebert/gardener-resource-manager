// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"

	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/gardener/gardener-resource-manager/cmd/gardener-resource-manager/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		<-signals.SetupSignalHandler()
	}()

	if err := app.NewResourceManagerCommand().ExecuteContext(ctx); err != nil {
		if log := runtimelog.Log; log.Enabled() {
			log.Error(err, "error running gardener-resource-manager")
		} else {
			fmt.Printf("error running gardener-resource-manager: %v", err)
		}
		os.Exit(1)
	}
}
