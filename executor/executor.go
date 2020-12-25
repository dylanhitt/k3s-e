package executor

import (
	"bytes"
	"text/template"
)

type Executor interface {
	Execute() error
	GetJoinToken() string
}

type InitOpts struct {
	SystemdType string
	EnvFile     string
	BinDir      string
	K3sCMD      string
}

var systemdTmpl = `
[Unit]
Description=Lightweight Kubernetes
Documentation=https://k3s.io
Wants=network-online.target
After=network-online.target
[Install]
WantedBy=multi-user.target
[Service]
Type={{ .SystemdType }}
EnvironmentFile={{ .EnvFile }}
KillMode=process
Delegate=yes
# Having non-zero Limit*s causes performance problems due to accounting overhead
# in the kernel. We recommend using cgroups to do container-local accounting.
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity
TasksMax=infinity
TimeoutStartSec=0
Restart=always
RestartSec=5s
ExecStartPre=-/sbin/modprobe br_netfilter
ExecStartPre=-/sbin/modprobe overlay
ExecStart={{ .BinDir }}/k3s \\
    {{ .K3sCMD }}
`

func createInitConfig(sc InitOpts) string {
	t := template.Must(template.New("systemd").Parse(systemdTmpl))

	var tpl bytes.Buffer
	err := t.ExecuteTemplate(&tpl, "systemd", sc)
	if err != nil {
		panic(err)
	}

	return tpl.String()
}
