/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MonitoringTestSpec defines the desired state of MonitoringTest
type MonitoringTestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Email string `json:"email,omitempty"`
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:validation:ExclusiveMaximum=false

	// Foo is an example field of MonitoringTest. Edit monitoringtest_types.go to remove/update
	// Size defines the number of Memcached instances
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Size int32 `json:"size,omitempty"`

	// Port defines the port that will be used to init the container with the image
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ContainerPort int32 `json:"containerPort,omitempty"`
}

// MonitoringTestStatus defines the observed state of MonitoringTest
type MonitoringTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions store the status conditions of the Memcached instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MonitoringTest is the Schema for the monitoringtests API
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=testMonitoring
// +kubebuilder:printcolumn:name="Domain Name",type=string,JSONPath=`.spec.domain`,priority=0,description="Domain name of the Cumulocity IoT Edge"
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.status.version`,priority=0,description="Deployed version of the Cumulocity IoT Edge"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.state`,priority=0,description="Current state of the deployment"
// +kubebuilder:printcolumn:name="Last Deployed",type="date",format="date",JSONPath=`.status.lastDeployedTime`,priority=0,description="Time since the Cumulocity IoT Edge was last successfully deployed or updated"
// +kubebuilder:printcolumn:name="Deployed Generation",type="integer",JSONPath=`.status.deployedGeneration`,priority=0,description="Generation of the CumulocityIoTEdge resource which is deployed"
// +kubebuilder:printcolumn:name="Latest Generation",type="integer",JSONPath=`.metadata.generation`,priority=0,description="Latest generation of the CumulocityIoTEdge resource"
// +kubebuilder:printcolumn:name="Deploying Generation",type="integer",JSONPath=`.status.deployingGeneration`,priority=0,description="Generation of the CumulocityIoTEdge resource which is being deployed"
// +kubebuilder:printcolumn:name="Warnings",type="string",JSONPath=`.status.warnings`,priority=1,description="Warning messages generated while validating the CumulocityIoTEdge resource"
type MonitoringTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringTestSpec   `json:"spec,omitempty"`
	Status MonitoringTestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MonitoringTestList contains a list of MonitoringTest
type MonitoringTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringTest{}, &MonitoringTestList{})
}
