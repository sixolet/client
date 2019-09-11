// Copyright © 2019 The Knative Authors
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

// Code generated. DO NOT EDIT.
// knative.dev/client/tools/generic-apigen --api v1beta1 --interface --out generated.go --my-package knative.dev/client/pkg/serving/generic --interface-package knative.dev/client/pkg/serving/generic

package generic

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/pkg/apis"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

type Configuration interface {
	GetKind() string
	SetKind(o string)
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetName() string
	SetName(o string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetUID() types.UID
	SetUID(o types.UID)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetCreationTimestamp() *v1.Time
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetClusterName() string
	SetClusterName(o string)
	GetSpec() ConfigurationSpec
	GetStatus() ConfigurationStatus
}

type ConfigurationSpec interface {
	GetTemplate() RevisionTemplateSpec
}

type ConfigurationStatus interface {
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetLatestReadyRevisionName() string
	SetLatestReadyRevisionName(o string)
	GetLatestCreatedRevisionName() string
	SetLatestCreatedRevisionName(o string)
}

type Revision interface {
	GetKind() string
	SetKind(o string)
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetName() string
	SetName(o string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetUID() types.UID
	SetUID(o types.UID)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetCreationTimestamp() *v1.Time
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetClusterName() string
	SetClusterName(o string)
	GetSpec() RevisionSpec
	GetStatus() RevisionStatus
}

type RevisionSpec interface {
	GetVolumes() []corev1.Volume
	SetVolumes(o []corev1.Volume)
	GetContainers() []corev1.Container
	SetContainers(o []corev1.Container)
	GetServiceAccountName() string
	SetServiceAccountName(o string)
	GetContainerConcurrency() v1beta1.RevisionContainerConcurrencyType
	SetContainerConcurrency(o v1beta1.RevisionContainerConcurrencyType)
	GetTimeoutSeconds() *int64
	SetTimeoutSeconds(o *int64)
}

type RevisionStatus interface {
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetServiceName() string
	SetServiceName(o string)
	GetLogURL() string
	SetLogURL(o string)
	GetImageDigest() string
	SetImageDigest(o string)
}

type RevisionTemplateSpec interface {
	GetName() string
	SetName(o string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetSpec() RevisionSpec
}

type Route interface {
	GetKind() string
	SetKind(o string)
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetName() string
	SetName(o string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetUID() types.UID
	SetUID(o types.UID)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetCreationTimestamp() *v1.Time
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetClusterName() string
	SetClusterName(o string)
	GetSpec() RouteSpec
	GetStatus() RouteStatus
}

type RouteSpec interface {
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type RouteStatus interface {
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
	GetAddress() *duckv1beta1.Addressable
	SetAddress(o *duckv1beta1.Addressable)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type Service interface {
	GetKind() string
	SetKind(o string)
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetName() string
	SetName(o string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetUID() types.UID
	SetUID(o types.UID)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetCreationTimestamp() *v1.Time
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetClusterName() string
	SetClusterName(o string)
	GetSpec() ServiceSpec
	GetStatus() ServiceStatus
}

type ServiceSpec interface {
	GetTemplate() RevisionTemplateSpec
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type ServiceStatus interface {
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetLatestReadyRevisionName() string
	SetLatestReadyRevisionName(o string)
	GetLatestCreatedRevisionName() string
	SetLatestCreatedRevisionName(o string)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
	GetAddress() *duckv1beta1.Addressable
	SetAddress(o *duckv1beta1.Addressable)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type TrafficTarget interface {
	GetTag() string
	SetTag(o string)
	GetRevisionName() string
	SetRevisionName(o string)
	GetConfigurationName() string
	SetConfigurationName(o string)
	GetLatestRevision() *bool
	SetLatestRevision(o *bool)
	GetPercent() int
	SetPercent(o int)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
}

type TrafficTargetSlice interface {
	Iter() chan TrafficTarget
	Index(tag string, revisionname string, latestrevision *bool) int
	Get(i int) TrafficTarget
	Find(tag string, revisionname string, latestrevision *bool) (TrafficTarget, bool)
	Filter(predicate func(e TrafficTarget) bool) TrafficTargetSlice
	Upsert(tag string, revisionname string, configurationname string, latestrevision *bool, percent int, url *apis.URL) TrafficTarget
}
