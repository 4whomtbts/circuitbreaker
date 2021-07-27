package config

type CcbConfig struct {
	CircuitBreakerLevel string    `yaml:"circuit_breaker_level"`
	WatchInterval       int       `yaml:"watch_interval"`
	CpuConfig           CpuConfig `yaml:"cpu"`
	GpuConfig           GpuConfig `yaml:"gpu"`
	NodeExporters       []string  `yaml:"nodeExporters"`
	DcgmExporters       []string  `yaml:"dcgmExporters"`
	Emails              []string  `yaml:"emails"`
	EmailSender         string    `yaml:"emailSender"`
	EmailSenderPwd      string    `yaml:"emailSenderPwd"`
	SshUser             string    `yaml:"sshUser"`
	SshPwd              string    `yaml:"sshPwd"`
	ExcludedImages      []string  `yaml:"excludedImages"`
}

type CpuConfig struct {
	TriggerPoint int `yaml:"triggerPoint"`
	TolerableNumber int `yaml:"tolerableNumber"`
}

type GpuConfig struct {
	TriggerPoint int `yaml:"triggerPoint"`
	TolerableNumber int `yaml:"tolerableNumber"`

}