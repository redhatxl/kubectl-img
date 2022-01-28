

[![ci](https://github.com/redhatxl/kubeimg/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/redhatxl/kubeimg/actions/workflows/ci.yml)

​	



![Kubectl image Logo](docs/logo/logo.png)

`kubectl image` is a kubectl plugin that allows you to show kubernetes resource image.

[TOC]




## Installing

### Pre-built binaries

See the [release](https://github.com/iovisor/kubectl-trace/releases) page for the full list of pre-built assets.

The commands here show `amd64` versions, `386` versions are available in the releases page.

**Linux**

```bash
curl -L -o kubectl-img.tar.gz https://github.com/iovisor/kubectl-trace/releases/download/v0.1.0-rc.1/kubectl-trace_0.1.0-rc.1_linux_amd64.tar.gz
tar -xvf kubectl-trace.tar.gz
mv kubectl-trace /usr/local/bin/kubectl-trace
```

**OSX**

```bash
curl -L -o kubectl-trace.tar.gz https://github.com/iovisor/kubectl-trace/releases/download/v0.1.0-rc.1/kubectl-trace_0.1.0-rc.1_darwin_amd64.tar.gz
tar -xvf kubectl-trace.tar.gz
mv kubectl-trace /usr/local/bin/kubectl-trace
```


**Windows**

In PowerShell v5+
```powershell
$url = "https://github.com/iovisor/kubectl-trace/releases/download/v0.1.0-rc.1/kubectl-trace_0.1.0-rc.1_windows_amd64.zip"
$output = "$PSScriptRoot\kubectl-trace.zip"

Invoke-WebRequest -Uri $url -OutFile $output
Expand-Archive "$PSScriptRoot\kubectl-trace.zip" -DestinationPath "$PSScriptRoot\kubectl-trace"
```

### Source

Using go modules, you can build kubectl-trace at any git tag:

```
GO111MODULE=on go get github.com/iovisor/kubectl-trace/cmd/kubectl-trace@latest
```

This will download and compile `kubectl-trace` so that you can use it as a kubectl plugin with `kubectl trace`, note that you will need to be on a recent version of go which supports go modules.

To keep track of the ref you used to build, you can add an ldflag at build time to set this to match the ref provided to go modules:

```
> GO111MODULE=on go get -ldflags='-X github.com/iovisor/kubectl-trace/pkg/version.gitCommit=v0.1.2' github.com/iovisor/kubectl-trace/cmd/kubectl-trace@v0.1.2
> $GOHOME/bin/kubectl-trace version
git commit: v0.1.2
build date: 2021-08-10 12:38:37.921341766 -0400 EDT m=+0.034327432
```

**Note:** It is recommended you build tagged revisions only if you are looking for stability. Building branches such as `master` or the `latest` tag may result in a more unstable build which has received less QA than a tagged release.

### Packages

You can't find the package for your distro of choice?
You are very welcome and encouraged to create it and then [open an issue](https://github.com/iovisor/kubectl-trace/issues/new) to inform us for review.

#### Arch - AUR

The official [PKGBUILD](https://aur.archlinux.org/cgit/aur.git/tree/PKGBUILD?h=kubectl-trace-git) is on AUR.

If you use `yay` to manage AUR packages you can do:

```
yay -S kubectl-trace-git
```

## Architecture

See [architecture.md](/docs/architecture.md)

## Usage

You don't need to setup anything on your cluster before using it, please don't use it already
on a production system, just because this isn't yet 100% ready.

```shell
 xuel@kaliarchmacbookpro  ~  kubectl img image -h
show k8s resource image

Usage:
  kubeimg image [flags]

Flags:
  -b, --cronjobs       show cronjobs image
  -e, --daemonsets     show daemonsets image
  -d, --deployments    show deployments image
  -h, --help           help for image
  -o, --jobs           show jobs image
  -j, --json           show json format
  -f, --statefulsets   show statefulsets image

Global Flags:
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default cache directory (default "/Users/xuel/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
```

### View Deployments Images

```shell
# View the images of all deployments of the entire kubernetes cluster
$ kubectl img image --deployments
# View the images of all deployments of the entire kubernetes cluster
$ kubectl img image --deployments -n default
```

![](https://kaliarch-bucket-1251990360.cos.ap-beijing.myqcloud.com/blog_img/20220128112944.png)

### View all resource

```shell
# view all resource 
$ kubectl img image -bedof
```

![](https://kaliarch-bucket-1251990360.cos.ap-beijing.myqcloud.com/blog_img/20220128114642.png)

### View format

Table display is used by default

```shell
# Table display is used by default
$ kubectl img image --deployments -n default -j
```

![](https://kaliarch-bucket-1251990360.cos.ap-beijing.myqcloud.com/blog_img/20220128113907.png)