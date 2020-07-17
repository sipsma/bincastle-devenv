module github.com/sipsma/bincastle-devenv

go 1.14

// TODO need to figure out how to remove need for these replaces https://github.com/moby/buildkit/pull/1425#issuecomment-609220074

replace github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe

replace github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305

replace github.com/godbus/dbus => github.com/godbus/dbus v0.0.0-20181101234600-2ff6f7ffd60f

// TODO if I don't have these lines, "go build" downgrades runc to rc8 in this go.mod file everytime (why???!!!)
replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.0-rc10

replace github.com/opencontainers/runtime-spec => github.com/opencontainers/runtime-spec v0.1.2-0.20190207185410-29686dbc5559

replace github.com/containerd/containerd => github.com/containerd/containerd v1.3.1-0.20200512144102-f13ba8f2f2fd

replace github.com/docker/docker => github.com/docker/docker v1.4.2-0.20200227233006-38f52c9fec82

require (
	github.com/moby/buildkit v0.7.1-0.20200717034530-97ca82f4c422
	github.com/opencontainers/go-digest v1.0.0
	github.com/sipsma/bincastle v0.0.0-20200717182751-a5051970c7c2
	github.com/sipsma/bincastle-distro v0.0.0-20200717200940-04970ec04316
)
