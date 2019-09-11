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


// Code generated by go run ./cmd/tools --api v1beta1 --interface --out generated.go --my-package knative.dev/client/pkg/serving/generic --interface-package knative.dev/client/pkg/serving/generic. DO NOT EDIT.
package generic

import(
	servingv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis/duck/v1beta1"
	"knative.dev/pkg/apis"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)


type Configuration interface {
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetClusterName() string
	SetClusterName(o string)
	GetCreationTimestamp() *v1.Time
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetKind() string
	SetKind(o string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetName() string
	SetName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetSpec() ConfigurationSpec
	GetStatus() ConfigurationStatus
	GetUID() types.UID
	SetUID(o types.UID)
}

type ConfigurationSpec interface {
	GetTemplate() RevisionTemplateSpec
}

type ConfigurationStatus interface {
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetLatestCreatedRevisionName() string
	SetLatestCreatedRevisionName(o string)
	GetLatestReadyRevisionName() string
	SetLatestReadyRevisionName(o string)
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
}

type Revision interface {
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetClusterName() string
	SetClusterName(o string)
	GetCreationTimestamp() *v1.Time
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetKind() string
	SetKind(o string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetName() string
	SetName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetSpec() RevisionSpec
	GetStatus() RevisionStatus
	GetUID() types.UID
	SetUID(o types.UID)
}

type RevisionSpec interface {
	GetContainerConcurrency() servingv1beta1.RevisionContainerConcurrencyType
	SetContainerConcurrency(o servingv1beta1.RevisionContainerConcurrencyType)
	GetContainers() []corev1.Container
	SetContainers(o []corev1.Container)
	GetServiceAccountName() string
	SetServiceAccountName(o string)
	GetTimeoutSeconds() *int64
	SetTimeoutSeconds(o *int64)
	GetVolumes() []corev1.Volume
	SetVolumes(o []corev1.Volume)
}

type RevisionStatus interface {
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetImageDigest() string
	SetImageDigest(o string)
	GetLogURL() string
	SetLogURL(o string)
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetServiceName() string
	SetServiceName(o string)
}

type RevisionTemplateSpec interface {
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetName() string
	SetName(o string)
	GetSpec() RevisionSpec
}

type Route interface {
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetClusterName() string
	SetClusterName(o string)
	GetCreationTimestamp() *v1.Time
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetKind() string
	SetKind(o string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetName() string
	SetName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetSpec() RouteSpec
	GetStatus() RouteStatus
	GetUID() types.UID
	SetUID(o types.UID)
}

type RouteSpec interface {
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type RouteStatus interface {
	GetAddress() *v1beta1.Addressable
	SetAddress(o *v1beta1.Addressable)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
}

type Service interface {
	GetAPIVersion() string
	SetAPIVersion(o string)
	GetAnnotations() map[string]string
	SetAnnotations(o map[string]string)
	GetClusterName() string
	SetClusterName(o string)
	GetCreationTimestamp() *v1.Time
	GetDeletionGracePeriodSeconds() *int64
	SetDeletionGracePeriodSeconds(o *int64)
	GetDeletionTimestamp() *v1.Time
	SetDeletionTimestamp(o *v1.Time)
	GetFinalizers() []string
	SetFinalizers(o []string)
	GetGenerateName() string
	SetGenerateName(o string)
	GetGeneration() int64
	SetGeneration(o int64)
	GetInitializers() *v1.Initializers
	SetInitializers(o *v1.Initializers)
	GetKind() string
	SetKind(o string)
	GetLabels() map[string]string
	SetLabels(o map[string]string)
	GetName() string
	SetName(o string)
	GetNamespace() string
	SetNamespace(o string)
	GetOwnerReferences() []v1.OwnerReference
	SetOwnerReferences(o []v1.OwnerReference)
	GetResourceVersion() string
	SetResourceVersion(o string)
	GetSelfLink() string
	SetSelfLink(o string)
	GetSpec() ServiceSpec
	GetStatus() ServiceStatus
	GetUID() types.UID
	SetUID(o types.UID)
}

type ServiceSpec interface {
	GetTemplate() RevisionTemplateSpec
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type ServiceStatus interface {
	GetAddress() *v1beta1.Addressable
	SetAddress(o *v1beta1.Addressable)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetLatestCreatedRevisionName() string
	SetLatestCreatedRevisionName(o string)
	GetLatestReadyRevisionName() string
	SetLatestReadyRevisionName(o string)
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
}

type TrafficTarget interface {
	GetConfigurationName() string
	SetConfigurationName(o string)
	GetLatestRevision() *bool
	SetLatestRevision(o *bool)
	GetPercent() int
	SetPercent(o int)
	GetRevisionName() string
	SetRevisionName(o string)
	GetTag() string
	SetTag(o string)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
}

type TrafficTargetSlice interface {
	Iter() chan TrafficTarget
	Index(latestrevision *bool, revisionname string, tag string) int
	Get(i int) TrafficTarget
	Find(latestrevision *bool, revisionname string, tag string)(TrafficTarget, bool)
	Filter(predicate func (e TrafficTarget) bool) TrafficTargetSlice
	Upsert(configurationname string, latestrevision *bool, percent int, revisionname string, tag string, url *apis.URL) TrafficTarget
}
