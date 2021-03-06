package alluxio

type Alluxio struct {
	Image           string            `yaml:"image"`
	ImageTag        string            `yaml:"imageTag"`
	ImagePullPolicy string            `yaml:"imagePullPolicy"`
	NodeSelector    map[string]string `yaml:"nodeSelector,omitempty"`
	ReplicaCount    int               `yaml:"replicaCount,omitempty"`

	Properties map[string]string `yaml:"properties,omitempty"`

	Master Master `yaml:"master,omitempty"`

	Woker Woker `yaml:"worker,omitempty"`

	Fuse Fuse `yaml:"fuse,omitempty"`

	Tieredstore Tieredstore `yaml:"tieredstore,omitempty"`
}

type Woker struct {
	JvmOptions   string            `yaml:"jvmOptions,omitempty"`
	NodeSelector map[string]string `yaml:"nodeSelector,omitempty"`
}

type Master struct {
	JvmOptions   string            `yaml:"jvmOptions,omitempty"`
	NodeSelector map[string]string `yaml:"nodeSelector,omitempty"`
}

type Fuse struct {
	Image              string            `yaml:"image,omitempty"`
	NodeSelector       map[string]string `yaml:"nodeSelector,omitempty"`
	ImageTag           string            `yaml:"imageTag,omitempty"`
	ImagePullPolicy    string            `yaml:"imagePullPolicy,omitempty"`
	Env                map[string]string `yaml:"env,omitempty"`
	JvmOptions         string            `yaml:"jvmOptions,omitempty"`
	ShortCircuitPolicy string            `yaml:"shortCircuitPolicy,omitempty"`
	Args               []string          `yaml:"args,omitempty"`
}

type Tieredstore struct {
	Levels []Level `yaml:"levels,omitempty"`
}

type Level struct {
	Alias string  `yaml:"alias,omitempty"`
	Level int     `yaml:"level"`
	Type  string  `yaml:"type,omitempty"`
	Path  string  `yaml:"path,omitempty"`
	Quota string  `yaml:"quota,omitempty"`
	High  float64 `yaml:"high,omitempty"`
	Low   float64 `yaml:"low,omitempty"`
}

type cacheStates struct {
	cacheCapacity string
	// cacheable        string
	lowWaterMark     string
	highWaterMark    string
	cached           string
	cachedPercentage string
	nonCacheable     string
}
