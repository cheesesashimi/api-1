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
type MachineOSBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec describes the configuration of the machine OS.
	// +kubebuilder:validation:Required
	Spec MachineOSBuildSpec `json:"spec"`

	// status describes the last observed state of this machine OS.
	// +optional
	Status MachineOSBuildStatus `json:"status"`
}

// Determines what MachineConfig inputs should be consumed by the build. Only
// one of these fields should be populated.
type BuildConfig struct {
	// Reference to a rendered MachineConfig. Will not automatically rebuild.
	RenderedMachineConfig *corev1.ObjectReference `json:"renderedMachineConfig,omitempty"`

	// Reference to a MachineConfigPool. Will automatically rebuild if the pool
	// picks up a new rendered MachineConfig.
	MachineConfigPool *corev1.ObjectReference `json:"machineConfigPool,omitempty"`

	// Note: Since the corev1.ObjectReference type suggests (but does not fully
	// require) a GroupVersionKind, we could theoretically get away with a single
	// field here. We'd need to validate that the Kind is either a MachineConfig
	// or a MachineConfigPool.
}

type MachineOSBuildSpec struct {
	// Specifies the Containerfile that a given Machine OS should be built with.
	Containerfile string `json:"containerfile"`

	// How one specifies what MachineConfig or MachineConfigPool to build with.
	BuildConfig `json:",inline"`

	// The base OS image which will be used for the Machine OS build. If not
	// present, this value will be silently populated with the clusters' current OS image.
	// +optional
	BaseOSImage string `json:"baseOSImage,omitempty"`
}

type MachineOSBuildStatus struct {
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

	// MachineOSBuildHistory is a reference to a successfully produced MachineOSImage.
	BuildHistory []MachineOSBuildHistory `json:"machineOSHistory"`

	// Contains references to the build pods and temporal objects such as
	// ConfigMaps and Secrets that are created / consumed by the build process.
	// Will be cleared after the build completes successfully.
	BuildObjects []corev1.ObjectReference `json:"buildObjects"`
}

// Represents when a given Machine OS was built.
type MachineOSBuildHistory struct {
	// The time when the build was performed.
	Built metav1.Time `json:"built"`
	// The image pullspec (in digested format) for the built OS image.
	ImagePullspec string `json:"imagePullspec"`
	// Reference to the produced MachineOSImage.
	corev1.ObjectReference `json:",inline"`
}
