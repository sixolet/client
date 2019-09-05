package generic

import (
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
	duckv1alpha1 "knative.dev/pkg/apis/duck/v1alpha1"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

func (r V1alpha1ConfigurationSpec) GetTemplate() RevisionTemplateSpec {
	if r.DeprecatedRevisionTemplate != nil {
		return V1alpha1RevisionTemplateSpec{r.DeprecatedRevisionTemplate}
	}
	return V1alpha1RevisionTemplateSpec{r.Template}
}

func (r V1alpha1ServiceSpec) GetTemplate() RevisionTemplateSpec {
	if r.DeprecatedRunLatest != nil {
		return V1alpha1ConfigurationSpec{&r.DeprecatedRunLatest.Configuration}.GetTemplate()
	} else if r.DeprecatedPinned != nil {
		return V1alpha1ConfigurationSpec{&r.DeprecatedPinned.Configuration}.GetTemplate()
	} else if r.DeprecatedRelease != nil {
		return V1alpha1ConfigurationSpec{&r.DeprecatedRelease.Configuration}.GetTemplate()
	}
	return V1alpha1RevisionTemplateSpec{r.Template}
}

func (r V1alpha1RevisionSpec) GetContainers() []corev1.Container {
	if r.DeprecatedContainer != nil {
		return []corev1.Container{*r.DeprecatedContainer}
	}
	return r.Containers
}

func (r V1alpha1RevisionSpec) SetContainers(o []corev1.Container) {
	if r.DeprecatedContainer != nil {
		r.DeprecatedContainer = &o[0]
	}
	r.Containers = o
}

func (s V1alpha1TrafficTargetSlice) Upsert(
	tag string, revisionname string, configurationname string, latestrevision *bool, percent int, url *apis.URL) TrafficTarget {
	ins := v1alpha1.TrafficTarget{}
	ins.Tag = tag
	ins.RevisionName = revisionname
	ins.ConfigurationName = configurationname
	ins.LatestRevision = latestrevision
	ins.Percent = percent
	ins.URL = url
	idx := s.Index(tag, revisionname, latestrevision)
	if idx >= 0 {
		s.Elts[idx] = ins
	} else {
		idx = len(s.Elts)
		s.Elts = append(s.Elts, ins)
	}
	return V1alpha1TrafficTarget{&s.Elts[idx]}
}

func (r V1alpha1ServiceStatus) GetAddress() *duckv1beta1.Addressable {
	a := r.Address
	if a == nil {
		return nil
	}
	return &r.Address.Addressable
}

func (r V1alpha1ServiceStatus) SetAddress(o *duckv1beta1.Addressable) {
	if o == nil {
		r.Address = nil
	}
	// a copy is the best we can do.
	r.Address = &duckv1alpha1.Addressable{Addressable: *o}
}

func (r V1alpha1RouteStatus) GetAddress() *duckv1beta1.Addressable {
	a := r.Address
	if a == nil {
		return nil
	}
	return &r.Address.Addressable
}

func (r V1alpha1RouteStatus) SetAddress(o *duckv1beta1.Addressable) {
	if o == nil {
		r.Address = nil
	}
	// a copy is the best we can do.
	r.Address = &duckv1alpha1.Addressable{Addressable: *o}
}
