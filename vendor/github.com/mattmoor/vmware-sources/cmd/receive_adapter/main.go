/*
Copyright 2020 The Knative Authors

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

package main

import (
	"context"

	"github.com/mattmoor/vmware-sources/pkg/vsphere"
	"k8s.io/client-go/kubernetes"
	"knative.dev/eventing/pkg/adapter/v2"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
)

func main() {
	ctx := signals.NewContext()
	kc := kubernetes.NewForConfigOrDie(sharedmain.ParseAndGetConfigOrDie())
	ctx = context.WithValue(ctx, kubeclient.Key{}, kc)
	adapter.MainWithContext(ctx, "vspheresource", vsphere.NewEnvConfig, vsphere.NewAdapter)
}
