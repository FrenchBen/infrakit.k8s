# InfraKit.K8S

[![CircleCI](https://circleci.com/gh/docker/infrakit.aws.svg?style=shield&circle-token=e74dcf8c25027948307a7618041e1d1997ded50a)](https://circleci.com/gh/docker/infrakit.k8s)

[InfraKit](https://github.com/docker/infrakit) plugin for creating and managing resources in Kubernetes.

## Flavor plugin

An InfraKit instance plugin is provided, which creates a Kubernetes cluster.

### Building and running

To build the Kubernetes Flavor plugin, run `make binaries`.  The plugin binary will be located at
`./build/infrakit-flavor-kubernetes`.

At a minimum, the plugin requires the SSL folder to use for storing instance SSL certs. This folder must exist prior to using the plugin. 

```console
$ mkdir ./ssl
$ build/infrakit-flavor-kubernetes --ssl-dir ./ssl
INFO[0000] Starting plugin
INFO[0000] Listening on: unix:///run/infrakit/plugins/flavor-kubernetes.sock
INFO[0000] listener protocol= unix addr= /run/infrakit/plugins/flavor-kubernetes.sock err= <nil>
```

### Example

To continue with an example, we will use the [default](https://github.com/docker/infrakit/tree/master/cmd/group) Group
plugin:
```console
$ build/infrakit-group-default
INFO[0000] Starting discovery
INFO[0000] Starting plugin
INFO[0000] Starting
INFO[0000] Listening on: unix:///run/infrakit/plugins/group.sock
INFO[0000] listener protocol= unix addr= /run/infrakit/plugins/group.sock err= <nil>
```

and the [Vagrant](https://github.com/docker/infrakit/tree/master/example/instance/vagrant) Instance plugin with the specified Kubernetes template:
```console
$ build/infrakit-instance-vagrant --dir ./tutorial/ --template ./templates/kubernetes.tmpl 
INFO[0000] Starting plugin
INFO[0000] Listening on: unix:///run/infrakit/plugins/instance-vagrant.sock
INFO[0000] listener protocol= unix addr= /run/infrakit/plugins/instance-vagrant.sock err= <nil>
```

We will use a basic configuration that creates a single cluster with 1 etcd, 1 controller, 1 worker, from the examples directory.


Finally, instruct the Group plugin to start watching the group for:
etcd
```console
$ build/infrakit group watch example/etcd.json
watching k8s-etcd
```
controller
```console
$ build/infrakit group watch example/controller.json
watching k8s-controller
```
worker
```console
$ build/infrakit group watch example/worker.json
watching k8s-worker
```

In the console running the Group plugin, we will see input like the following:
```
INFO[1208] Watching group ''
INFO[1219] Adding 1 instances to group to reach desired 1
INFO[1219] Created instance i-ba0412a2 with tags map[infrakit.config_sha:dUBtWGmkptbGg29ecBgv1VJYzys= infrakit.group:aws-example]
```

Additionally, the CLI will report the newly-created instance:
```console
$ build/infrakit group inspect aws-example
ID                             	LOGICAL                        	TAGS
i-ba0412a2                     	172.31.41.13                   	Name=infrakit-example,infrakit.config_sha=dUBtWGmkptbGg29ecBgv1VJYzys=,infrakit.group=aws-example
```

Retrieve the IP address of the host from the AWS console, and use SSH to verify that our shell code ran:

```console
$ ssh ubuntu@55.55.55.55 cat /hello
Hello, World!
```

### Plugin properties

The plugin expects properties in the following format:
```json
{
  "Tags": {
  },
  "RunInstancesInput": {
  }
}
```

The `Tags` property is a string-string mapping of EC2 instance tags to include on all instances that are created.
`RunInstancesInput` follows the structure of the type by the same name in the
[AWS go SDK](http://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#RunInstancesInput).


#### AWS API Credentials

The plugin can use API credentials from several sources.
- config file:
  see [AWS docs](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-config-files)
- EC2 instance metadata:
  see [AWS docs](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html)

Additional credentials sources are supported, but are not generally recommended as they are less secure:
- command line arguments: `--session-token`, or  `--access-key-id` and `--secret-access-key`
- environment variables:
  see [AWS docs](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-environment)


## Reporting security issues

The maintainers take security seriously. If you discover a security issue,
please bring it to their attention right away!

Please **DO NOT** file a public issue, instead send your report privately to
[security@docker.com](mailto:security@docker.com).

Security reports are greatly appreciated and we will publicly thank you for it.
We also like to send gifts—if you're into Docker schwag, make sure to let
us know. We currently do not offer a paid security bounty program, but are not
ruling it out in the future.


## Copyright and license

Copyright © 2016 Docker, Inc. All rights reserved, except as follows. Code
is released under the Apache 2.0 license. The README.md file, and files in the
"docs" folder are licensed under the Creative Commons Attribution 4.0
International License under the terms and conditions set forth in the file
"LICENSE.docs". You may obtain a duplicate copy of the same license, titled
CC-BY-SA-4.0, at http://creativecommons.org/licenses/by/4.0/.
