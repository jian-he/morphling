/*
Copyright 2021 The Alibaba Authors.

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

package util

import (
	"fmt"

	morphlingv1alpha1 "github.com/alibaba/morphling/api/v1alpha1"
	"github.com/alibaba/morphling/pkg/controllers/consts"
)

func GetServiceDeploymentName(t *morphlingv1alpha1.Trial) string {
	return t.Name + "-" + "deployment"
}

func GetServiceName(t *morphlingv1alpha1.Trial) string {
	return t.Name + "-" + "service"
}

func GetStressTestJobName(t *morphlingv1alpha1.Trial) string {
	return t.Name + "-" + "client-job"
}

func GetServiceEndpoint(t *morphlingv1alpha1.Trial) string {
	return fmt.Sprintf("%s:%d",
		GetServiceName(t),
		consts.DefaultServicePort)
}

func GetDBStorageEndpoint() string {
	return fmt.Sprintf("%s.%s:%d",
		consts.DefaultMorphlingDBManagerServiceName,
		consts.DefaultMorphlingNamespace,
		consts.DefaultMorphlingDBManagerServicePort)
}
