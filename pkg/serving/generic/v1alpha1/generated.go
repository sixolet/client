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
// knative.dev/client/tools/generic-apigen --api v1alpha1 --out generated.go --my-package knative.dev/client/pkg/serving/generic/v1alpha1 --interface-package knative.dev/client/pkg/serving/generic

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/client/pkg/serving/generic"
	"knative.dev/pkg/apis"
	servingv1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

type V1alpha1Configuration struct {
	*servingv1alpha1.Configuration
}

func (r V1alpha1Configuration) GetKind() string {
	return r.Kind
}
func (r V1alpha1Configuration) SetKind(o string) {
	r.Kind = o
}
func (r V1alpha1Configuration) GetAPIVersion() string {
	return r.APIVersion
}
func (r V1alpha1Configuration) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r V1alpha1Configuration) GetName() string {
	return r.Name
}
func (r V1alpha1Configuration) SetName(o string) {
	r.Name = o
}
func (r V1alpha1Configuration) GetGenerateName() string {
	return r.GenerateName
}
func (r V1alpha1Configuration) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r V1alpha1Configuration) GetNamespace() string {
	return r.Namespace
}
func (r V1alpha1Configuration) SetNamespace(o string) {
	r.Namespace = o
}
func (r V1alpha1Configuration) GetSelfLink() string {
	return r.SelfLink
}
func (r V1alpha1Configuration) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r V1alpha1Configuration) GetUID() types.UID {
	return r.UID
}
func (r V1alpha1Configuration) SetUID(o types.UID) {
	r.UID = o
}
func (r V1alpha1Configuration) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r V1alpha1Configuration) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r V1alpha1Configuration) GetGeneration() int64 {
	return r.Generation
}
func (r V1alpha1Configuration) SetGeneration(o int64) {
	r.Generation = o
}
func (r V1alpha1Configuration) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r V1alpha1Configuration) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r V1alpha1Configuration) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r V1alpha1Configuration) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r V1alpha1Configuration) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r V1alpha1Configuration) GetLabels() map[string]string {
	return r.Labels
}
func (r V1alpha1Configuration) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r V1alpha1Configuration) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r V1alpha1Configuration) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r V1alpha1Configuration) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r V1alpha1Configuration) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r V1alpha1Configuration) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r V1alpha1Configuration) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r V1alpha1Configuration) GetFinalizers() []string {
	return r.Finalizers
}
func (r V1alpha1Configuration) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r V1alpha1Configuration) GetClusterName() string {
	return r.ClusterName
}
func (r V1alpha1Configuration) SetClusterName(o string) {
	r.ClusterName = o
}
func (r V1alpha1Configuration) GetSpec() generic.ConfigurationSpec {
	return V1alpha1ConfigurationSpec{&r.Spec}
}
func (r V1alpha1Configuration) GetStatus() generic.ConfigurationStatus {
	return V1alpha1ConfigurationStatus{&r.Status}
}

type V1alpha1ConfigurationSpec struct {
	*servingv1alpha1.ConfigurationSpec
}

type V1alpha1ConfigurationStatus struct {
	*servingv1alpha1.ConfigurationStatus
}

func (r V1alpha1ConfigurationStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r V1alpha1ConfigurationStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r V1alpha1ConfigurationStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r V1alpha1ConfigurationStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r V1alpha1ConfigurationStatus) GetLatestReadyRevisionName() string {
	return r.LatestReadyRevisionName
}
func (r V1alpha1ConfigurationStatus) SetLatestReadyRevisionName(o string) {
	r.LatestReadyRevisionName = o
}
func (r V1alpha1ConfigurationStatus) GetLatestCreatedRevisionName() string {
	return r.LatestCreatedRevisionName
}
func (r V1alpha1ConfigurationStatus) SetLatestCreatedRevisionName(o string) {
	r.LatestCreatedRevisionName = o
}

type V1alpha1ManualType struct {
	*servingv1alpha1.ManualType
}

type V1alpha1PinnedType struct {
	*servingv1alpha1.PinnedType
}

func (r V1alpha1PinnedType) GetRevisionName() string {
	return r.RevisionName
}
func (r V1alpha1PinnedType) SetRevisionName(o string) {
	r.RevisionName = o
}
func (r V1alpha1PinnedType) GetConfiguration() generic.ConfigurationSpec {
	return V1alpha1ConfigurationSpec{&r.Configuration}
}

type V1alpha1ReleaseType struct {
	*servingv1alpha1.ReleaseType
}

func (r V1alpha1ReleaseType) GetRevisions() []string {
	return r.Revisions
}
func (r V1alpha1ReleaseType) SetRevisions(o []string) {
	r.Revisions = o
}
func (r V1alpha1ReleaseType) GetRolloutPercent() int {
	return r.RolloutPercent
}
func (r V1alpha1ReleaseType) SetRolloutPercent(o int) {
	r.RolloutPercent = o
}
func (r V1alpha1ReleaseType) GetConfiguration() generic.ConfigurationSpec {
	return V1alpha1ConfigurationSpec{&r.Configuration}
}

type V1alpha1Revision struct {
	*servingv1alpha1.Revision
}

func (r V1alpha1Revision) GetKind() string {
	return r.Kind
}
func (r V1alpha1Revision) SetKind(o string) {
	r.Kind = o
}
func (r V1alpha1Revision) GetAPIVersion() string {
	return r.APIVersion
}
func (r V1alpha1Revision) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r V1alpha1Revision) GetName() string {
	return r.Name
}
func (r V1alpha1Revision) SetName(o string) {
	r.Name = o
}
func (r V1alpha1Revision) GetGenerateName() string {
	return r.GenerateName
}
func (r V1alpha1Revision) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r V1alpha1Revision) GetNamespace() string {
	return r.Namespace
}
func (r V1alpha1Revision) SetNamespace(o string) {
	r.Namespace = o
}
func (r V1alpha1Revision) GetSelfLink() string {
	return r.SelfLink
}
func (r V1alpha1Revision) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r V1alpha1Revision) GetUID() types.UID {
	return r.UID
}
func (r V1alpha1Revision) SetUID(o types.UID) {
	r.UID = o
}
func (r V1alpha1Revision) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r V1alpha1Revision) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r V1alpha1Revision) GetGeneration() int64 {
	return r.Generation
}
func (r V1alpha1Revision) SetGeneration(o int64) {
	r.Generation = o
}
func (r V1alpha1Revision) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r V1alpha1Revision) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r V1alpha1Revision) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r V1alpha1Revision) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r V1alpha1Revision) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r V1alpha1Revision) GetLabels() map[string]string {
	return r.Labels
}
func (r V1alpha1Revision) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r V1alpha1Revision) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r V1alpha1Revision) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r V1alpha1Revision) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r V1alpha1Revision) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r V1alpha1Revision) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r V1alpha1Revision) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r V1alpha1Revision) GetFinalizers() []string {
	return r.Finalizers
}
func (r V1alpha1Revision) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r V1alpha1Revision) GetClusterName() string {
	return r.ClusterName
}
func (r V1alpha1Revision) SetClusterName(o string) {
	r.ClusterName = o
}
func (r V1alpha1Revision) GetSpec() generic.RevisionSpec {
	return V1alpha1RevisionSpec{&r.Spec}
}
func (r V1alpha1Revision) GetStatus() generic.RevisionStatus {
	return V1alpha1RevisionStatus{&r.Status}
}

type V1alpha1RevisionSpec struct {
	*servingv1alpha1.RevisionSpec
}

func (r V1alpha1RevisionSpec) GetVolumes() []corev1.Volume {
	return r.Volumes
}
func (r V1alpha1RevisionSpec) SetVolumes(o []corev1.Volume) {
	r.Volumes = o
}
func (r V1alpha1RevisionSpec) GetServiceAccountName() string {
	return r.ServiceAccountName
}
func (r V1alpha1RevisionSpec) SetServiceAccountName(o string) {
	r.ServiceAccountName = o
}
func (r V1alpha1RevisionSpec) GetContainerConcurrency() v1beta1.RevisionContainerConcurrencyType {
	return r.ContainerConcurrency
}
func (r V1alpha1RevisionSpec) SetContainerConcurrency(o v1beta1.RevisionContainerConcurrencyType) {
	r.ContainerConcurrency = o
}
func (r V1alpha1RevisionSpec) GetTimeoutSeconds() *int64 {
	return r.TimeoutSeconds
}
func (r V1alpha1RevisionSpec) SetTimeoutSeconds(o *int64) {
	r.TimeoutSeconds = o
}

type V1alpha1RevisionStatus struct {
	*servingv1alpha1.RevisionStatus
}

func (r V1alpha1RevisionStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r V1alpha1RevisionStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r V1alpha1RevisionStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r V1alpha1RevisionStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r V1alpha1RevisionStatus) GetServiceName() string {
	return r.ServiceName
}
func (r V1alpha1RevisionStatus) SetServiceName(o string) {
	r.ServiceName = o
}
func (r V1alpha1RevisionStatus) GetLogURL() string {
	return r.LogURL
}
func (r V1alpha1RevisionStatus) SetLogURL(o string) {
	r.LogURL = o
}
func (r V1alpha1RevisionStatus) GetImageDigest() string {
	return r.ImageDigest
}
func (r V1alpha1RevisionStatus) SetImageDigest(o string) {
	r.ImageDigest = o
}

type V1alpha1RevisionTemplateSpec struct {
	*servingv1alpha1.RevisionTemplateSpec
}

func (r V1alpha1RevisionTemplateSpec) GetName() string {
	return r.Name
}
func (r V1alpha1RevisionTemplateSpec) SetName(o string) {
	r.Name = o
}
func (r V1alpha1RevisionTemplateSpec) GetLabels() map[string]string {
	return r.Labels
}
func (r V1alpha1RevisionTemplateSpec) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r V1alpha1RevisionTemplateSpec) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r V1alpha1RevisionTemplateSpec) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r V1alpha1RevisionTemplateSpec) GetSpec() generic.RevisionSpec {
	return V1alpha1RevisionSpec{&r.Spec}
}

type V1alpha1Route struct {
	*servingv1alpha1.Route
}

func (r V1alpha1Route) GetKind() string {
	return r.Kind
}
func (r V1alpha1Route) SetKind(o string) {
	r.Kind = o
}
func (r V1alpha1Route) GetAPIVersion() string {
	return r.APIVersion
}
func (r V1alpha1Route) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r V1alpha1Route) GetName() string {
	return r.Name
}
func (r V1alpha1Route) SetName(o string) {
	r.Name = o
}
func (r V1alpha1Route) GetGenerateName() string {
	return r.GenerateName
}
func (r V1alpha1Route) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r V1alpha1Route) GetNamespace() string {
	return r.Namespace
}
func (r V1alpha1Route) SetNamespace(o string) {
	r.Namespace = o
}
func (r V1alpha1Route) GetSelfLink() string {
	return r.SelfLink
}
func (r V1alpha1Route) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r V1alpha1Route) GetUID() types.UID {
	return r.UID
}
func (r V1alpha1Route) SetUID(o types.UID) {
	r.UID = o
}
func (r V1alpha1Route) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r V1alpha1Route) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r V1alpha1Route) GetGeneration() int64 {
	return r.Generation
}
func (r V1alpha1Route) SetGeneration(o int64) {
	r.Generation = o
}
func (r V1alpha1Route) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r V1alpha1Route) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r V1alpha1Route) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r V1alpha1Route) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r V1alpha1Route) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r V1alpha1Route) GetLabels() map[string]string {
	return r.Labels
}
func (r V1alpha1Route) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r V1alpha1Route) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r V1alpha1Route) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r V1alpha1Route) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r V1alpha1Route) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r V1alpha1Route) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r V1alpha1Route) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r V1alpha1Route) GetFinalizers() []string {
	return r.Finalizers
}
func (r V1alpha1Route) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r V1alpha1Route) GetClusterName() string {
	return r.ClusterName
}
func (r V1alpha1Route) SetClusterName(o string) {
	r.ClusterName = o
}
func (r V1alpha1Route) GetSpec() generic.RouteSpec {
	return V1alpha1RouteSpec{&r.Spec}
}
func (r V1alpha1Route) GetStatus() generic.RouteStatus {
	return V1alpha1RouteStatus{&r.Status}
}

type V1alpha1RouteSpec struct {
	*servingv1alpha1.RouteSpec
}

func (r V1alpha1RouteSpec) GetTraffic() generic.TrafficTargetSlice {
	return V1alpha1TrafficTargetSlice{r.Traffic}
}
func (r V1alpha1RouteSpec) SetTraffic(o generic.TrafficTargetSlice) {
	r.Traffic = o.(V1alpha1TrafficTargetSlice).Elts
}

type V1alpha1RouteStatus struct {
	*servingv1alpha1.RouteStatus
}

func (r V1alpha1RouteStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r V1alpha1RouteStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r V1alpha1RouteStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r V1alpha1RouteStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r V1alpha1RouteStatus) GetURL() *apis.URL {
	return r.URL
}
func (r V1alpha1RouteStatus) SetURL(o *apis.URL) {
	r.URL = o
}
func (r V1alpha1RouteStatus) GetTraffic() generic.TrafficTargetSlice {
	return V1alpha1TrafficTargetSlice{r.Traffic}
}
func (r V1alpha1RouteStatus) SetTraffic(o generic.TrafficTargetSlice) {
	r.Traffic = o.(V1alpha1TrafficTargetSlice).Elts
}

type V1alpha1RunLatestType struct {
	*servingv1alpha1.RunLatestType
}

func (r V1alpha1RunLatestType) GetConfiguration() generic.ConfigurationSpec {
	return V1alpha1ConfigurationSpec{&r.Configuration}
}

type V1alpha1Service struct {
	*servingv1alpha1.Service
}

func (r V1alpha1Service) GetKind() string {
	return r.Kind
}
func (r V1alpha1Service) SetKind(o string) {
	r.Kind = o
}
func (r V1alpha1Service) GetAPIVersion() string {
	return r.APIVersion
}
func (r V1alpha1Service) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r V1alpha1Service) GetName() string {
	return r.Name
}
func (r V1alpha1Service) SetName(o string) {
	r.Name = o
}
func (r V1alpha1Service) GetGenerateName() string {
	return r.GenerateName
}
func (r V1alpha1Service) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r V1alpha1Service) GetNamespace() string {
	return r.Namespace
}
func (r V1alpha1Service) SetNamespace(o string) {
	r.Namespace = o
}
func (r V1alpha1Service) GetSelfLink() string {
	return r.SelfLink
}
func (r V1alpha1Service) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r V1alpha1Service) GetUID() types.UID {
	return r.UID
}
func (r V1alpha1Service) SetUID(o types.UID) {
	r.UID = o
}
func (r V1alpha1Service) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r V1alpha1Service) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r V1alpha1Service) GetGeneration() int64 {
	return r.Generation
}
func (r V1alpha1Service) SetGeneration(o int64) {
	r.Generation = o
}
func (r V1alpha1Service) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r V1alpha1Service) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r V1alpha1Service) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r V1alpha1Service) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r V1alpha1Service) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r V1alpha1Service) GetLabels() map[string]string {
	return r.Labels
}
func (r V1alpha1Service) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r V1alpha1Service) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r V1alpha1Service) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r V1alpha1Service) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r V1alpha1Service) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r V1alpha1Service) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r V1alpha1Service) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r V1alpha1Service) GetFinalizers() []string {
	return r.Finalizers
}
func (r V1alpha1Service) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r V1alpha1Service) GetClusterName() string {
	return r.ClusterName
}
func (r V1alpha1Service) SetClusterName(o string) {
	r.ClusterName = o
}
func (r V1alpha1Service) GetSpec() generic.ServiceSpec {
	return V1alpha1ServiceSpec{&r.Spec}
}
func (r V1alpha1Service) GetStatus() generic.ServiceStatus {
	return V1alpha1ServiceStatus{&r.Status}
}

type V1alpha1ServiceSpec struct {
	*servingv1alpha1.ServiceSpec
}

func (r V1alpha1ServiceSpec) GetTraffic() generic.TrafficTargetSlice {
	return V1alpha1TrafficTargetSlice{r.Traffic}
}
func (r V1alpha1ServiceSpec) SetTraffic(o generic.TrafficTargetSlice) {
	r.Traffic = o.(V1alpha1TrafficTargetSlice).Elts
}

type V1alpha1ServiceStatus struct {
	*servingv1alpha1.ServiceStatus
}

func (r V1alpha1ServiceStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r V1alpha1ServiceStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r V1alpha1ServiceStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r V1alpha1ServiceStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r V1alpha1ServiceStatus) GetURL() *apis.URL {
	return r.URL
}
func (r V1alpha1ServiceStatus) SetURL(o *apis.URL) {
	r.URL = o
}
func (r V1alpha1ServiceStatus) GetTraffic() generic.TrafficTargetSlice {
	return V1alpha1TrafficTargetSlice{r.Traffic}
}
func (r V1alpha1ServiceStatus) SetTraffic(o generic.TrafficTargetSlice) {
	r.Traffic = o.(V1alpha1TrafficTargetSlice).Elts
}
func (r V1alpha1ServiceStatus) GetLatestReadyRevisionName() string {
	return r.LatestReadyRevisionName
}
func (r V1alpha1ServiceStatus) SetLatestReadyRevisionName(o string) {
	r.LatestReadyRevisionName = o
}
func (r V1alpha1ServiceStatus) GetLatestCreatedRevisionName() string {
	return r.LatestCreatedRevisionName
}
func (r V1alpha1ServiceStatus) SetLatestCreatedRevisionName(o string) {
	r.LatestCreatedRevisionName = o
}

type V1alpha1TrafficTarget struct {
	*servingv1alpha1.TrafficTarget
}

func (r V1alpha1TrafficTarget) GetTag() string {
	return r.Tag
}
func (r V1alpha1TrafficTarget) SetTag(o string) {
	r.Tag = o
}
func (r V1alpha1TrafficTarget) GetRevisionName() string {
	return r.RevisionName
}
func (r V1alpha1TrafficTarget) SetRevisionName(o string) {
	r.RevisionName = o
}
func (r V1alpha1TrafficTarget) GetConfigurationName() string {
	return r.ConfigurationName
}
func (r V1alpha1TrafficTarget) SetConfigurationName(o string) {
	r.ConfigurationName = o
}
func (r V1alpha1TrafficTarget) GetLatestRevision() *bool {
	return r.LatestRevision
}
func (r V1alpha1TrafficTarget) SetLatestRevision(o *bool) {
	r.LatestRevision = o
}
func (r V1alpha1TrafficTarget) GetPercent() int {
	return r.Percent
}
func (r V1alpha1TrafficTarget) SetPercent(o int) {
	r.Percent = o
}
func (r V1alpha1TrafficTarget) GetURL() *apis.URL {
	return r.URL
}
func (r V1alpha1TrafficTarget) SetURL(o *apis.URL) {
	r.URL = o
}

type V1alpha1TrafficTargetSlice struct {
	Elts []servingv1alpha1.TrafficTarget
}

func (s V1alpha1TrafficTargetSlice) Iter() chan generic.TrafficTarget {
	ret := make(chan generic.TrafficTarget, len(s.Elts))
	for _, elt := range s.Elts {
		ret <- V1alpha1TrafficTarget{&elt}
	}
	close(ret)
	return ret
}
func (s V1alpha1TrafficTargetSlice) Index(tag string, revisionname string, latestrevision *bool) int {
	for i, elt := range s.Elts {
		if elt.Tag != tag {
			continue
		}
		if elt.RevisionName != revisionname {
			continue
		}
		var v bool
		if elt.LatestRevision != nil {
			v = *elt.LatestRevision
		} else if latestrevision != nil {
			continue
		}
		if v != *latestrevision {
			continue
		}
		return i
	}
	return -1
}
func (s V1alpha1TrafficTargetSlice) Get(i int) generic.TrafficTarget {
	return V1alpha1TrafficTarget{&s.Elts[i]}
}
func (s V1alpha1TrafficTargetSlice) Find(tag string, revisionname string, latestrevision *bool) (generic.TrafficTarget, bool) {
	i := s.Index(tag, revisionname, latestrevision)
	if i < 0 {
		return V1alpha1TrafficTarget{nil}, false
	}
	return s.Get(i), true
}
func (s V1alpha1TrafficTargetSlice) Filter(predicate func(e generic.TrafficTarget) bool) generic.TrafficTargetSlice {
	ret := []servingv1alpha1.TrafficTarget{}
	for _, elt := range s.Elts {
		if predicate(V1alpha1TrafficTarget{&elt}) {
			ret = append(ret, elt)
		}
	}
	return V1alpha1TrafficTargetSlice{ret}
}
