/*
Copyright 2022 The Crossplane Authors.

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

package v1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:verbs=create;update,path=/validate-apiextensions-crossplane-io-v1-compositeresourcedefinition,mutating=false,failurePolicy=fail,groups=apiextensions.crossplane.io,resources=compositeresourcedefinitions,versions=v1,name=compositeresourcedefinitions.apiextensions.crossplane.io,sideEffects=None,admissionReviewVersions=v1

// SetupWebhookWithManager sets up the Composition webhook with the provided manager and CustomValidator.
func (c *CompositeResourceDefinition) SetupWebhookWithManager(mgr ctrl.Manager, validator admission.CustomValidator) error {
	// Needed to inject validator in order to avoid dependency cycles.
	return ctrl.NewWebhookManagedBy(mgr).
		WithValidator(validator).
		For(c).
		Complete()
}
