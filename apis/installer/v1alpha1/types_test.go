/*
Copyright AppsCode Inc. and Contributors.

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

package v1alpha1_test

import (
	"os"
	"testing"

	"go.openviz.dev/installer/apis/installer/v1alpha1"

	schemachecker "kmodules.xyz/schema-checker"
)

func TestDefaultValues(t *testing.T) {
	checker := schemachecker.New(os.DirFS("../../.."),
		schemachecker.TestCase{Obj: v1alpha1.GrafanaOperatorSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.KubeGrafanaDashboardsSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.MonitoringOperatorSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.TricksterSpec{}},
	)
	checker.TestAll(t)
}
