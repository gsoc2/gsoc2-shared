module stitcher

go 1.15

//replace github.com/gsoc2/gsoc2-shared => ../

require (
	cloud.google.com/go/storage v1.30.1
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/docker/docker v20.10.5+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/gsoc2/gsoc2-shared v0.0.0-20231021014948-a187b7551fdc
	github.com/morikuni/aec v1.0.0 // indirect
	google.golang.org/api v0.128.0
	gopkg.in/yaml.v2 v2.4.0
)
