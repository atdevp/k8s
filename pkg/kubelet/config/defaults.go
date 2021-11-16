package config

// default config
const (
	DefaultKubeletPodsDirName               = "pods"
	DefaultKubeletVolumesDirName            = "volumes"
	DefaultKubeletVolumeSubPathsDirName     = "volume-subpaths"
	DefaultKubeletVolumeDevicesDirName      = "volumeDevices"
	DefaultKubeletPluginDirName             = "plugins"
	DefaultKubeletPluginRegistrationDirName = "plugins_registry"
	DefaultKubeletContainersDirName         = "containers"
	DefaultKubeletPluginContainersDirName   = "plugin-containers"
	DefaultKubeletPodResourcesDirName       = "pod-resources"
	KubeletPluginsDirSELinuxLabel           = "system_u:object_r:container_file_t:s0"
)
