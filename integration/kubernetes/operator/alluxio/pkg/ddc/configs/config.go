package configs

import (
	dataset "github.com/Alluxio/alluxio/api/v1alpha1"
)

type Config struct {
	Name string

	Namespace string
	// The runtime type, now only is alluxio
	Runtime string

	// The name of the ConfigMap which contains the Value config files.
	TemplateFile string

	// Runtime options which are appened to configMaps.
	Options map[string]string

	DatasetConfig
}

type DatasetConfig struct {
	Mounts []dataset.Mount `json:"mounts,omitempty"`

	// Specifies how to prefetch dataset.
	PrefetchStrategy dataset.PrefetchStrategy `json:"prefetchStrategy,omitempty"`

	// replicas is the min replicas of dataset in the cluster
	Replicas *int32 `json:"replicas,omitempty"`

	LowWaterMarkRatio float64 `json:"lowWaterMarkRatio,omitempty"`

	HighWaterMarkRatio float64 `json:"highWaterMarkRatio,omitempty"`

	// NodeAffinity defines constraints that limit what nodes this dataset can be cached to.
	// This field influences the scheduling of pods that use the cached dataset.
	NodeAffinity *dataset.CacheableNodeAffinity `json:"nodeAffinity,omitempty"`
}
