package generic

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck/v1beta1"
	servingv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"
)

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
	GetAddress() *v1beta1.Addressable
	SetAddress(o *v1beta1.Addressable)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type Servingv1beta1ServiceStatus struct {
	*servingv1beta1.ServiceStatus
}

func (r Servingv1beta1ServiceStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r Servingv1beta1ServiceStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r Servingv1beta1ServiceStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r Servingv1beta1ServiceStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r Servingv1beta1ServiceStatus) GetLatestReadyRevisionName() string {
	return r.LatestReadyRevisionName
}
func (r Servingv1beta1ServiceStatus) SetLatestReadyRevisionName(o string) {
	r.LatestReadyRevisionName = o
}
func (r Servingv1beta1ServiceStatus) GetLatestCreatedRevisionName() string {
	return r.LatestCreatedRevisionName
}
func (r Servingv1beta1ServiceStatus) SetLatestCreatedRevisionName(o string) {
	r.LatestCreatedRevisionName = o
}
func (r Servingv1beta1ServiceStatus) GetURL() *apis.URL {
	return r.URL
}
func (r Servingv1beta1ServiceStatus) SetURL(o *apis.URL) {
	r.URL = o
}
func (r Servingv1beta1ServiceStatus) GetAddress() *v1beta1.Addressable {
	return r.Address
}
func (r Servingv1beta1ServiceStatus) SetAddress(o *v1beta1.Addressable) {
	r.Address = o
}
func (r Servingv1beta1ServiceStatus) GetTraffic() TrafficTargetSlice {
	return Servingv1beta1TrafficTargetSlice{&r.Traffic}
}
func (r Servingv1beta1ServiceStatus) SetTraffic(o TrafficTargetSlice) {
	r.Traffic = o
}

type ConfigurationSpec interface {
	GetTemplate() RevisionTemplateSpec
}

type Servingv1beta1ConfigurationSpec struct {
	*servingv1beta1.ConfigurationSpec
}

func (r Servingv1beta1ConfigurationSpec) GetTemplate() RevisionTemplateSpec {
	return Servingv1beta1RevisionTemplateSpec{&r.Template}
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

type Servingv1beta1Service struct {
	*servingv1beta1.Service
}

func (r Servingv1beta1Service) GetKind() string {
	return r.Kind
}
func (r Servingv1beta1Service) SetKind(o string) {
	r.Kind = o
}
func (r Servingv1beta1Service) GetAPIVersion() string {
	return r.APIVersion
}
func (r Servingv1beta1Service) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r Servingv1beta1Service) GetName() string {
	return r.Name
}
func (r Servingv1beta1Service) SetName(o string) {
	r.Name = o
}
func (r Servingv1beta1Service) GetGenerateName() string {
	return r.GenerateName
}
func (r Servingv1beta1Service) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r Servingv1beta1Service) GetNamespace() string {
	return r.Namespace
}
func (r Servingv1beta1Service) SetNamespace(o string) {
	r.Namespace = o
}
func (r Servingv1beta1Service) GetSelfLink() string {
	return r.SelfLink
}
func (r Servingv1beta1Service) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r Servingv1beta1Service) GetUID() types.UID {
	return r.UID
}
func (r Servingv1beta1Service) SetUID(o types.UID) {
	r.UID = o
}
func (r Servingv1beta1Service) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r Servingv1beta1Service) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r Servingv1beta1Service) GetGeneration() int64 {
	return r.Generation
}
func (r Servingv1beta1Service) SetGeneration(o int64) {
	r.Generation = o
}
func (r Servingv1beta1Service) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r Servingv1beta1Service) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r Servingv1beta1Service) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r Servingv1beta1Service) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r Servingv1beta1Service) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r Servingv1beta1Service) GetLabels() map[string]string {
	return r.Labels
}
func (r Servingv1beta1Service) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r Servingv1beta1Service) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r Servingv1beta1Service) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r Servingv1beta1Service) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r Servingv1beta1Service) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r Servingv1beta1Service) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r Servingv1beta1Service) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r Servingv1beta1Service) GetFinalizers() []string {
	return r.Finalizers
}
func (r Servingv1beta1Service) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r Servingv1beta1Service) GetClusterName() string {
	return r.ClusterName
}
func (r Servingv1beta1Service) SetClusterName(o string) {
	r.ClusterName = o
}
func (r Servingv1beta1Service) GetSpec() ServiceSpec {
	return Servingv1beta1ServiceSpec{&r.Spec}
}
func (r Servingv1beta1Service) GetStatus() ServiceStatus {
	return Servingv1beta1ServiceStatus{&r.Status}
}

type RevisionSpec interface {
	GetVolumes() []corev1.Volume
	SetVolumes(o []corev1.Volume)
	GetContainers() []corev1.Container
	SetContainers(o []corev1.Container)
	GetServiceAccountName() string
	SetServiceAccountName(o string)
	GetContainerConcurrency() servingv1beta1.RevisionContainerConcurrencyType
	SetContainerConcurrency(o servingv1beta1.RevisionContainerConcurrencyType)
	GetTimeoutSeconds() *int64
	SetTimeoutSeconds(o *int64)
}

type Servingv1beta1RevisionSpec struct {
	*servingv1beta1.RevisionSpec
}

func (r Servingv1beta1RevisionSpec) GetVolumes() []corev1.Volume {
	return r.Volumes
}
func (r Servingv1beta1RevisionSpec) SetVolumes(o []corev1.Volume) {
	r.Volumes = o
}
func (r Servingv1beta1RevisionSpec) GetContainers() []corev1.Container {
	return r.Containers
}
func (r Servingv1beta1RevisionSpec) SetContainers(o []corev1.Container) {
	r.Containers = o
}
func (r Servingv1beta1RevisionSpec) GetServiceAccountName() string {
	return r.ServiceAccountName
}
func (r Servingv1beta1RevisionSpec) SetServiceAccountName(o string) {
	r.ServiceAccountName = o
}
func (r Servingv1beta1RevisionSpec) GetContainerConcurrency() servingv1beta1.RevisionContainerConcurrencyType {
	return r.ContainerConcurrency
}
func (r Servingv1beta1RevisionSpec) SetContainerConcurrency(o servingv1beta1.RevisionContainerConcurrencyType) {
	r.ContainerConcurrency = o
}
func (r Servingv1beta1RevisionSpec) GetTimeoutSeconds() *int64 {
	return r.TimeoutSeconds
}
func (r Servingv1beta1RevisionSpec) SetTimeoutSeconds(o *int64) {
	r.TimeoutSeconds = o
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

type Servingv1beta1ConfigurationStatus struct {
	*servingv1beta1.ConfigurationStatus
}

func (r Servingv1beta1ConfigurationStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r Servingv1beta1ConfigurationStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r Servingv1beta1ConfigurationStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r Servingv1beta1ConfigurationStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r Servingv1beta1ConfigurationStatus) GetLatestReadyRevisionName() string {
	return r.LatestReadyRevisionName
}
func (r Servingv1beta1ConfigurationStatus) SetLatestReadyRevisionName(o string) {
	r.LatestReadyRevisionName = o
}
func (r Servingv1beta1ConfigurationStatus) GetLatestCreatedRevisionName() string {
	return r.LatestCreatedRevisionName
}
func (r Servingv1beta1ConfigurationStatus) SetLatestCreatedRevisionName(o string) {
	r.LatestCreatedRevisionName = o
}

type RouteStatus interface {
	GetObservedGeneration() int64
	SetObservedGeneration(o int64)
	GetConditions() []apis.Condition
	SetConditions(o []apis.Condition)
	GetURL() *apis.URL
	SetURL(o *apis.URL)
	GetAddress() *v1beta1.Addressable
	SetAddress(o *v1beta1.Addressable)
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type Servingv1beta1RouteStatus struct {
	*servingv1beta1.RouteStatus
}

func (r Servingv1beta1RouteStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r Servingv1beta1RouteStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r Servingv1beta1RouteStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r Servingv1beta1RouteStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r Servingv1beta1RouteStatus) GetURL() *apis.URL {
	return r.URL
}
func (r Servingv1beta1RouteStatus) SetURL(o *apis.URL) {
	r.URL = o
}
func (r Servingv1beta1RouteStatus) GetAddress() *v1beta1.Addressable {
	return r.Address
}
func (r Servingv1beta1RouteStatus) SetAddress(o *v1beta1.Addressable) {
	r.Address = o
}
func (r Servingv1beta1RouteStatus) GetTraffic() TrafficTargetSlice {
	return Servingv1beta1TrafficTargetSlice{&r.Traffic}
}
func (r Servingv1beta1RouteStatus) SetTraffic(o TrafficTargetSlice) {
	r.Traffic = o
}

type TrafficTargetSlice interface {
	Iter() chan TrafficTarget
	Index(tag string, revisionname string, latestrevision bool) int
	Get(i int) TrafficTarget
	Find(tag string, revisionname string, latestrevision bool) (TrafficTarget, bool)
	Remove(tag string, revisionname string, latestrevision bool)
	Upsert(tag string, revisionname string, configurationname string, latestrevision *bool, percent int, url *apis.URL)
	MatchesTag(tag string) []TrafficTarget
	MatchesRevisionName(revisionname string) []TrafficTarget
	MatchesLatestRevision(latestrevision bool) []TrafficTarget
}

type Servingv1beta1TrafficTargetSlice struct {
	Elts []servingv1beta1.TrafficTarget
}

func (s Servingv1beta1TrafficTargetSlice) Iter() chan TrafficTarget {
	ret := make(chan TrafficTarget, len(s.Elts))
	for _, elt := range s.Elts {
		ret <- Servingv1beta1TrafficTarget{&elt}
	}
	close(ret)
	return ret
}
func (s Servingv1beta1TrafficTargetSlice) Index(tag string, revisionname string, latestrevision bool) int {
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
		}
		if v != latestrevision {
			continue
		}
		return i
	}
	return -1
}
func (s Servingv1beta1TrafficTargetSlice) Get(i int) TrafficTarget {
	return Servingv1beta1TrafficTarget{&s.Elts[i]}
}
func (s Servingv1beta1TrafficTargetSlice) Find(tag string, revisionname string, latestrevision bool) (TrafficTarget, bool) {
	i := s.Index(tag, revisionname, latestrevision)
	if i < 0 {
		return Servingv1beta1TrafficTarget{nil}, false
	}
	return Get(i), true
}
func (s Servingv1beta1TrafficTargetSlice) Remove(tag string, revisionname string, latestrevision bool) {
}
func (s Servingv1beta1TrafficTargetSlice) Upsert(tag string, revisionname string, configurationname string, latestrevision *bool, percent int, url *apis.URL) {
}
func (s Servingv1beta1TrafficTargetSlice) MatchesTag(tag string) []TrafficTarget {
}
func (s Servingv1beta1TrafficTargetSlice) MatchesRevisionName(revisionname string) []TrafficTarget {
}
func (s Servingv1beta1TrafficTargetSlice) MatchesLatestRevision(latestrevision bool) []TrafficTarget {
}

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

type Servingv1beta1Configuration struct {
	*servingv1beta1.Configuration
}

func (r Servingv1beta1Configuration) GetKind() string {
	return r.Kind
}
func (r Servingv1beta1Configuration) SetKind(o string) {
	r.Kind = o
}
func (r Servingv1beta1Configuration) GetAPIVersion() string {
	return r.APIVersion
}
func (r Servingv1beta1Configuration) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r Servingv1beta1Configuration) GetName() string {
	return r.Name
}
func (r Servingv1beta1Configuration) SetName(o string) {
	r.Name = o
}
func (r Servingv1beta1Configuration) GetGenerateName() string {
	return r.GenerateName
}
func (r Servingv1beta1Configuration) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r Servingv1beta1Configuration) GetNamespace() string {
	return r.Namespace
}
func (r Servingv1beta1Configuration) SetNamespace(o string) {
	r.Namespace = o
}
func (r Servingv1beta1Configuration) GetSelfLink() string {
	return r.SelfLink
}
func (r Servingv1beta1Configuration) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r Servingv1beta1Configuration) GetUID() types.UID {
	return r.UID
}
func (r Servingv1beta1Configuration) SetUID(o types.UID) {
	r.UID = o
}
func (r Servingv1beta1Configuration) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r Servingv1beta1Configuration) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r Servingv1beta1Configuration) GetGeneration() int64 {
	return r.Generation
}
func (r Servingv1beta1Configuration) SetGeneration(o int64) {
	r.Generation = o
}
func (r Servingv1beta1Configuration) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r Servingv1beta1Configuration) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r Servingv1beta1Configuration) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r Servingv1beta1Configuration) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r Servingv1beta1Configuration) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r Servingv1beta1Configuration) GetLabels() map[string]string {
	return r.Labels
}
func (r Servingv1beta1Configuration) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r Servingv1beta1Configuration) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r Servingv1beta1Configuration) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r Servingv1beta1Configuration) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r Servingv1beta1Configuration) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r Servingv1beta1Configuration) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r Servingv1beta1Configuration) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r Servingv1beta1Configuration) GetFinalizers() []string {
	return r.Finalizers
}
func (r Servingv1beta1Configuration) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r Servingv1beta1Configuration) GetClusterName() string {
	return r.ClusterName
}
func (r Servingv1beta1Configuration) SetClusterName(o string) {
	r.ClusterName = o
}
func (r Servingv1beta1Configuration) GetSpec() ConfigurationSpec {
	return Servingv1beta1ConfigurationSpec{&r.Spec}
}
func (r Servingv1beta1Configuration) GetStatus() ConfigurationStatus {
	return Servingv1beta1ConfigurationStatus{&r.Status}
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

type Servingv1beta1TrafficTarget struct {
	*servingv1beta1.TrafficTarget
}

func (r Servingv1beta1TrafficTarget) GetTag() string {
	return r.Tag
}
func (r Servingv1beta1TrafficTarget) SetTag(o string) {
	r.Tag = o
}
func (r Servingv1beta1TrafficTarget) GetRevisionName() string {
	return r.RevisionName
}
func (r Servingv1beta1TrafficTarget) SetRevisionName(o string) {
	r.RevisionName = o
}
func (r Servingv1beta1TrafficTarget) GetConfigurationName() string {
	return r.ConfigurationName
}
func (r Servingv1beta1TrafficTarget) SetConfigurationName(o string) {
	r.ConfigurationName = o
}
func (r Servingv1beta1TrafficTarget) GetLatestRevision() *bool {
	return r.LatestRevision
}
func (r Servingv1beta1TrafficTarget) SetLatestRevision(o *bool) {
	r.LatestRevision = o
}
func (r Servingv1beta1TrafficTarget) GetPercent() int {
	return r.Percent
}
func (r Servingv1beta1TrafficTarget) SetPercent(o int) {
	r.Percent = o
}
func (r Servingv1beta1TrafficTarget) GetURL() *apis.URL {
	return r.URL
}
func (r Servingv1beta1TrafficTarget) SetURL(o *apis.URL) {
	r.URL = o
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

type Servingv1beta1Route struct {
	*servingv1beta1.Route
}

func (r Servingv1beta1Route) GetKind() string {
	return r.Kind
}
func (r Servingv1beta1Route) SetKind(o string) {
	r.Kind = o
}
func (r Servingv1beta1Route) GetAPIVersion() string {
	return r.APIVersion
}
func (r Servingv1beta1Route) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r Servingv1beta1Route) GetName() string {
	return r.Name
}
func (r Servingv1beta1Route) SetName(o string) {
	r.Name = o
}
func (r Servingv1beta1Route) GetGenerateName() string {
	return r.GenerateName
}
func (r Servingv1beta1Route) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r Servingv1beta1Route) GetNamespace() string {
	return r.Namespace
}
func (r Servingv1beta1Route) SetNamespace(o string) {
	r.Namespace = o
}
func (r Servingv1beta1Route) GetSelfLink() string {
	return r.SelfLink
}
func (r Servingv1beta1Route) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r Servingv1beta1Route) GetUID() types.UID {
	return r.UID
}
func (r Servingv1beta1Route) SetUID(o types.UID) {
	r.UID = o
}
func (r Servingv1beta1Route) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r Servingv1beta1Route) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r Servingv1beta1Route) GetGeneration() int64 {
	return r.Generation
}
func (r Servingv1beta1Route) SetGeneration(o int64) {
	r.Generation = o
}
func (r Servingv1beta1Route) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r Servingv1beta1Route) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r Servingv1beta1Route) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r Servingv1beta1Route) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r Servingv1beta1Route) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r Servingv1beta1Route) GetLabels() map[string]string {
	return r.Labels
}
func (r Servingv1beta1Route) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r Servingv1beta1Route) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r Servingv1beta1Route) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r Servingv1beta1Route) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r Servingv1beta1Route) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r Servingv1beta1Route) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r Servingv1beta1Route) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r Servingv1beta1Route) GetFinalizers() []string {
	return r.Finalizers
}
func (r Servingv1beta1Route) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r Servingv1beta1Route) GetClusterName() string {
	return r.ClusterName
}
func (r Servingv1beta1Route) SetClusterName(o string) {
	r.ClusterName = o
}
func (r Servingv1beta1Route) GetSpec() RouteSpec {
	return Servingv1beta1RouteSpec{&r.Spec}
}
func (r Servingv1beta1Route) GetStatus() RouteStatus {
	return Servingv1beta1RouteStatus{&r.Status}
}

type RouteSpec interface {
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type Servingv1beta1RouteSpec struct {
	*servingv1beta1.RouteSpec
}

func (r Servingv1beta1RouteSpec) GetTraffic() TrafficTargetSlice {
	return Servingv1beta1TrafficTargetSlice{&r.Traffic}
}
func (r Servingv1beta1RouteSpec) SetTraffic(o TrafficTargetSlice) {
	r.Traffic = o
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

type Servingv1beta1Revision struct {
	*servingv1beta1.Revision
}

func (r Servingv1beta1Revision) GetKind() string {
	return r.Kind
}
func (r Servingv1beta1Revision) SetKind(o string) {
	r.Kind = o
}
func (r Servingv1beta1Revision) GetAPIVersion() string {
	return r.APIVersion
}
func (r Servingv1beta1Revision) SetAPIVersion(o string) {
	r.APIVersion = o
}
func (r Servingv1beta1Revision) GetName() string {
	return r.Name
}
func (r Servingv1beta1Revision) SetName(o string) {
	r.Name = o
}
func (r Servingv1beta1Revision) GetGenerateName() string {
	return r.GenerateName
}
func (r Servingv1beta1Revision) SetGenerateName(o string) {
	r.GenerateName = o
}
func (r Servingv1beta1Revision) GetNamespace() string {
	return r.Namespace
}
func (r Servingv1beta1Revision) SetNamespace(o string) {
	r.Namespace = o
}
func (r Servingv1beta1Revision) GetSelfLink() string {
	return r.SelfLink
}
func (r Servingv1beta1Revision) SetSelfLink(o string) {
	r.SelfLink = o
}
func (r Servingv1beta1Revision) GetUID() types.UID {
	return r.UID
}
func (r Servingv1beta1Revision) SetUID(o types.UID) {
	r.UID = o
}
func (r Servingv1beta1Revision) GetResourceVersion() string {
	return r.ResourceVersion
}
func (r Servingv1beta1Revision) SetResourceVersion(o string) {
	r.ResourceVersion = o
}
func (r Servingv1beta1Revision) GetGeneration() int64 {
	return r.Generation
}
func (r Servingv1beta1Revision) SetGeneration(o int64) {
	r.Generation = o
}
func (r Servingv1beta1Revision) GetCreationTimestamp() *v1.Time {
	return &r.CreationTimestamp
}
func (r Servingv1beta1Revision) GetDeletionTimestamp() *v1.Time {
	return r.DeletionTimestamp
}
func (r Servingv1beta1Revision) SetDeletionTimestamp(o *v1.Time) {
	r.DeletionTimestamp = o
}
func (r Servingv1beta1Revision) GetDeletionGracePeriodSeconds() *int64 {
	return r.DeletionGracePeriodSeconds
}
func (r Servingv1beta1Revision) SetDeletionGracePeriodSeconds(o *int64) {
	r.DeletionGracePeriodSeconds = o
}
func (r Servingv1beta1Revision) GetLabels() map[string]string {
	return r.Labels
}
func (r Servingv1beta1Revision) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r Servingv1beta1Revision) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r Servingv1beta1Revision) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r Servingv1beta1Revision) GetOwnerReferences() []v1.OwnerReference {
	return r.OwnerReferences
}
func (r Servingv1beta1Revision) SetOwnerReferences(o []v1.OwnerReference) {
	r.OwnerReferences = o
}
func (r Servingv1beta1Revision) GetInitializers() *v1.Initializers {
	return r.Initializers
}
func (r Servingv1beta1Revision) SetInitializers(o *v1.Initializers) {
	r.Initializers = o
}
func (r Servingv1beta1Revision) GetFinalizers() []string {
	return r.Finalizers
}
func (r Servingv1beta1Revision) SetFinalizers(o []string) {
	r.Finalizers = o
}
func (r Servingv1beta1Revision) GetClusterName() string {
	return r.ClusterName
}
func (r Servingv1beta1Revision) SetClusterName(o string) {
	r.ClusterName = o
}
func (r Servingv1beta1Revision) GetSpec() RevisionSpec {
	return Servingv1beta1RevisionSpec{&r.Spec}
}
func (r Servingv1beta1Revision) GetStatus() RevisionStatus {
	return Servingv1beta1RevisionStatus{&r.Status}
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

type Servingv1beta1RevisionStatus struct {
	*servingv1beta1.RevisionStatus
}

func (r Servingv1beta1RevisionStatus) GetObservedGeneration() int64 {
	return r.ObservedGeneration
}
func (r Servingv1beta1RevisionStatus) SetObservedGeneration(o int64) {
	r.ObservedGeneration = o
}
func (r Servingv1beta1RevisionStatus) GetConditions() []apis.Condition {
	return r.Conditions
}
func (r Servingv1beta1RevisionStatus) SetConditions(o []apis.Condition) {
	r.Conditions = o
}
func (r Servingv1beta1RevisionStatus) GetServiceName() string {
	return r.ServiceName
}
func (r Servingv1beta1RevisionStatus) SetServiceName(o string) {
	r.ServiceName = o
}
func (r Servingv1beta1RevisionStatus) GetLogURL() string {
	return r.LogURL
}
func (r Servingv1beta1RevisionStatus) SetLogURL(o string) {
	r.LogURL = o
}
func (r Servingv1beta1RevisionStatus) GetImageDigest() string {
	return r.ImageDigest
}
func (r Servingv1beta1RevisionStatus) SetImageDigest(o string) {
	r.ImageDigest = o
}

type ServiceSpec interface {
	GetTemplate() RevisionTemplateSpec
	GetTraffic() TrafficTargetSlice
	SetTraffic(o TrafficTargetSlice)
}

type Servingv1beta1ServiceSpec struct {
	*servingv1beta1.ServiceSpec
}

func (r Servingv1beta1ServiceSpec) GetTemplate() RevisionTemplateSpec {
	return Servingv1beta1RevisionTemplateSpec{&r.Template}
}
func (r Servingv1beta1ServiceSpec) GetTraffic() TrafficTargetSlice {
	return Servingv1beta1TrafficTargetSlice{&r.Traffic}
}
func (r Servingv1beta1ServiceSpec) SetTraffic(o TrafficTargetSlice) {
	r.Traffic = o
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

type Servingv1beta1RevisionTemplateSpec struct {
	*servingv1beta1.RevisionTemplateSpec
}

func (r Servingv1beta1RevisionTemplateSpec) GetName() string {
	return r.Name
}
func (r Servingv1beta1RevisionTemplateSpec) SetName(o string) {
	r.Name = o
}
func (r Servingv1beta1RevisionTemplateSpec) GetLabels() map[string]string {
	return r.Labels
}
func (r Servingv1beta1RevisionTemplateSpec) SetLabels(o map[string]string) {
	r.Labels = o
}
func (r Servingv1beta1RevisionTemplateSpec) GetAnnotations() map[string]string {
	return r.Annotations
}
func (r Servingv1beta1RevisionTemplateSpec) SetAnnotations(o map[string]string) {
	r.Annotations = o
}
func (r Servingv1beta1RevisionTemplateSpec) GetSpec() RevisionSpec {
	return Servingv1beta1RevisionSpec{&r.Spec}
}
