package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineOSBuild describes the OS image of a given MachineConfigPool.
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
// +kubebuilder:validation:XValidation:rule="self.metadata.name == self.spec.node.name",message="spec.node.name should match metadata.name"
type MachineOSImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec describes the configuration of the machine OS image.
	// +kubebuilder:validation:Required
	Spec MachineOSImageSpec `json:"spec"`

	// status describes the last observed state of this machine OS image.
	// +optional
	Status MachineOSImageStatus `json:"status"`
}

// My initial idea is that this wouldn't have a spec since a cluster admin or
// another controller would not be setting the spec on this.
type MachineOSImageSpec struct{}

type MachineOSImageInfo struct {
	// The base OS image pullspec that was used to build the OS image.
	BaseOSImage string `json:"baseOSImage"`
	// A reference to the rendered MachineConfig that was used to produce this OS image.
	RenderedMachineConfig corev1.ObjectReference `json:"renderedMachineConfig"`
	// The Containerfile that was used to produce this OS image.
	Containerfile string `json:"containerfile"`
	// The final image pullspec (in digested format) for the built OS image.
	ImagePullspec string `json:"finalImagePullspec"`
	// For when the image was built.
	Built metav1.Time `json:"built"`
}

type MachineOSImageStatus struct {
	// conditions represent the observations of a machine config node's current state.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// observedGeneration represents the generation observed by the controller.
	// This field is updated when the controller observes a change to the desiredConfig in the configVersion of the machine config node spec.
	// +kubebuilder:validation:Required
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	MachineOSImageInfo `json:",inline"`
}
