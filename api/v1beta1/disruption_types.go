/*

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

package v1beta1

import (
	"strconv"
	"strings"

	chaostypes "github.com/DataDog/chaos-fi-controller/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DisruptionSpec defines the desired state of Disruption
type DisruptionSpec struct {
	// +optional
	// +kubebuilder:validation:Minimum=0
	Count    *int       `json:"count"`    // number of pods to target
	Selector labels.Set `json:"selector"` // label selector
	// +optional
	NetworkFailure *NetworkFailureSpec `json:"networkFailure,omitempty"`
	// +optional
	NetworkLatency *NetworkLatencySpec `json:"networkLatency,omitempty"`
	// +optional
	NodeFailure *NodeFailureSpec `json:"nodeFailure,omitempty"`
}

// NetworkFailureSpec represents a network failure injection
type NetworkFailureSpec struct {
	Hosts       []string `json:"hosts,omitempty"`
	Port        int      `json:"port"`
	Probability int      `json:"probability"`
	Protocol    string   `json:"protocol"`
}

// GenerateArgs generates injection or cleanup pod arguments for the given spec
func (s *NetworkFailureSpec) GenerateArgs(mode chaostypes.PodMode, uid types.UID, containerID string) []string {
	args := []string{}

	switch mode {
	case chaostypes.PodModeInject:
		args = []string{
			"network-failure",
			"inject",
			"--uid",
			string(uid),
			"--container-id",
			containerID,
			"--port",
			strconv.Itoa(s.Port),
			"--protocol",
			s.Protocol,
			"--probability",
			strconv.Itoa(s.Probability),
			"--hosts",
		}
		args = append(args, strings.Split(strings.Join(s.Hosts, " --hosts "), " ")...)
	case chaostypes.PodModeClean:
		args = []string{
			"network-failure",
			"clean",
			"--uid",
			string(uid),
			"--container-id",
			containerID,
		}
	}

	return args
}

// NetworkLatencySpec represents a network latency injection
type NetworkLatencySpec struct {
	Delay uint     `json:"delay"`
	Hosts []string `json:"hosts"`
}

// GenerateArgs generates injection or cleanup pod arguments for the given spec
func (s *NetworkLatencySpec) GenerateArgs(mode chaostypes.PodMode, uid types.UID, containerID string) []string {
	args := []string{}

	switch mode {
	case chaostypes.PodModeInject:
		args = []string{
			"network-latency",
			"inject",
			"--uid",
			string(uid),
			"--container-id",
			containerID,
			"--delay",
			strconv.Itoa(int(s.Delay)),
			"--hosts",
		}
		args = append(args, strings.Split(strings.Join(s.Hosts, " --hosts "), " ")...)
	case chaostypes.PodModeClean:
		args = []string{
			"network-latency",
			"clean",
			"--uid",
			string(uid),
			"--container-id",
			containerID,
			"--hosts",
		}
		args = append(args, strings.Split(strings.Join(s.Hosts, " --hosts "), " ")...)
	}

	return args
}

// NodeFailureSpec represents a node failure injection
type NodeFailureSpec struct {
	Shutdown bool `json:"shutdown,omitempty"`
}

// GenerateArgs generates injection or cleanup pod arguments for the given spec
func (s *NodeFailureSpec) GenerateArgs(mode chaostypes.PodMode, uid types.UID, containerID string) []string {
	args := []string{}

	switch mode {
	case chaostypes.PodModeInject:
		args = []string{
			"node-failure",
			"inject",
			"--uid",
			string(uid),
		}
		if s.Shutdown {
			args = append(args, "--shutdown")
		}
	}

	return args
}

// DisruptionStatus defines the observed state of Disruption
type DisruptionStatus struct {
	// +optional
	IsFinalizing bool `json:"isFinalizing"`
	// +optional
	IsInjected bool `json:"isInjected"`
	// +optional
	TargetPods []string `json:"targetPods,omitempty"`
}

// +kubebuilder:object:root=true

// Disruption is the Schema for the disruptions API
// +kubebuilder:resource:shortName=dis
type Disruption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DisruptionSpec   `json:"spec,omitempty"`
	Status DisruptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DisruptionList contains a list of Disruption
type DisruptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disruption `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Disruption{}, &DisruptionList{})
}