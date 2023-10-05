# k8s-operator-monitoring

- health Prometheus endpoint
- Health of different components + metrics
- Operator can depend on this Endpoint to monitor the status and try to correct when necessary.

## Objectives

 1. Explore how to expose an health endpoint in the Operator
 2. Define the metrics which this endpoint checks and relays
 3. Operator itself will keep monitoring this endpoint
     - if any component is unhealthy, the Operator reacts with operations like (restart pod, restart the deployment etc..) ???

## Steps to do quick start

### Create project

```shell
$ export GO111MODULE=on
$ mkdir -p $HOME/projects/monitor-operator
$ cd $HOME/projects/monitor-operator
# we'll use a domain of monitor.com
# so all API groups will be <group>.monitor.com
$ operator-sdk init --domain=monitor.com --repo=github.com/nonus25/monitor-operator
```

```shell
❯ operator-sdk init --domain=monitor.com --repo=github.com/nonus25/monitor-operator
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.14.1
Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
```

### Create a new API and Controller

Create a new Custom Resource Definition(CRD) API with group cache version v1alpha1 and Kind monitor. When prompted, enter yes y for creating both the resource and controller.

```shell
❯ operator-sdk create api --group=cache --version=v1alpha1 --kind=Monitor
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/monitor_types.go
controllers/monitor_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
mkdir -p /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```

After modifying the `api/v1alpha1/*_types.go` file always run the following command to update the generated code for that resource type:

```shell
❯ make generate
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
```

```shell
❯ make manifests
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```

### Implement the Controller

For this example replace the generated controller file `controllers/monitor_controller.go` with the example `monitor_controller.go` implementation.

The example controller executes the following reconciliation logic for each monitor CR:

Create a monitor Deployment if it doesn’t exist
Ensure that the Deployment size is the same as specified by the monitor CR spec
Update the monitor CR status using the status writer with the names of the monitor pods

### Install CRD

```shell
❯ make generate
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
❯ make manifests
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```

```shell
❯ make install
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize || { curl -Ss "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash -s -- 3.8.7 /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin; }
{Version:kustomize/v3.8.7 GitCommit:ad092cc7a91c07fdf63a2e4b7f13fa588a39af4f BuildDate:2020-11-11T23:14:14Z GoOs:linux GoArch:amd64}
kustomize installed to /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/monitors.cache.monitor.com created
```

### Run locally outside the cluster

To run the operator locally execute the following command:

```shell
make run ENABLE_WEBHOOKS=false
```

and kept this running in a separate window

### Run as a Deployment inside the cluster

```shell
export user=nonus25

make docker-build IMG=docker.io/$user/monitor-operator:v0.0.1
```

```shell
❯ make docker-build IMG=docker.io/$user/monitor-operator:v0.0.1
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/setup-envtest || GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
KUBEBUILDER_ASSETS="/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/k8s/1.26.0-linux-amd64" go test ./... -coverprofile cover.out
?       github.com/nonus25/monitor-operator     [no test files]
?       github.com/nonus25/monitor-operator/api/v1alpha1        [no test files]
ok      github.com/nonus25/monitor-operator/controllers 0.013s  coverage: 0.0% of statements
docker build -t docker.io/nonus25/monitor-operator:v0.0.1 .
[+] Building 18.6s (18/18) FINISHED                                                                                                         docker:default
 => [internal] load .dockerignore                                                                                                           0.0s
 => => transferring context: 171B                                                                                                           0.0s
 => [internal] load build definition from Dockerfile                                                                                        0.0s
 => => transferring dockerfile: 1.29kB                                                                                                      0.0s
 => [internal] load metadata for gcr.io/distroless/static:nonroot                                                                           0.5s
 => [internal] load metadata for docker.io/library/golang:1.19                                                                              1.1s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                               0.0s
 => CACHED [stage-1 1/3] FROM gcr.io/distroless/static:nonroot@sha256:92d40eea0b5307a94f2ebee3e94095e704015fb41e35fc1fcbd1d151cc282222      0.0s
 => [builder 1/9] FROM docker.io/library/golang:1.19@sha256:3025bf670b8363ec9f1b4c4f27348e6d9b7fec607c47e401e40df816853e743a                0.0s
 => [internal] load build context                                                                                                           0.0s
 => => transferring context: 96.48kB                                                                                                        0.0s
 => CACHED [builder 2/9] WORKDIR /workspace                                                                                                 0.0s
 => [builder 3/9] COPY go.mod go.mod                                                                                                        0.1s
 => [builder 4/9] COPY go.sum go.sum                                                                                                        0.0s
 => [builder 5/9] RUN go mod download                                                                                                       3.7s
 => [builder 6/9] COPY main.go main.go                                                                                                      0.0s
 => [builder 7/9] COPY api/ api/                                                                                                            0.1s
 => [builder 8/9] COPY controllers/ controllers/                                                                                            0.0s
 => [builder 9/9] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go                                                  13.1s
 => [stage-1 2/3] COPY --from=builder /workspace/manager .                                                                                  0.1s
 => exporting to image                                                                                                                      0.1s
 => => exporting layers                                                                                                                     0.1s
 => => writing image sha256:7f50c409ca19d51bb80a1de88c50958021881c890e0c29f0a3af22666c4a69d2                                                0.0s
 => => naming to docker.io/nonus25/monitor-operator:v0.0.1 
 ```

Push the image to a repository:

```shell
make docker-push IMG=docker.io/$user/monitor-operator:v0.0.1
```

```shell
❯ make docker-push IMG=docker.io/nonus25/monitor-operator:v0.0.1
docker push docker.io/nonus25/monitor-operator:v0.0.1
The push refers to repository [docker.io/nonus25/monitor-operator]
2f606ade1e8e: Pushed 
4cb10dd2545b: Pushed 
d2d7ec0f6756: Pushed 
1a73b54f556b: Pushed 
e624a5370eca: Layer already exists 
d52f02c6501c: Pushed 
ff5700ec5418: Pushed 
7bea6b893187: Pushed 
6fbdf253bbc2: Pushed 
e023e0e48e6e: Pushed 
v0.0.1: digest: sha256:3466f0f98db1ebb0f76dd4dd29933130c48fab1d11a0e2c9de7751231a239e04 size: 2402
```

### Deploy the operator

For this example we will run the operator in the `monitor` namespace which can be specified for all resources in `config/default/kustomization.yaml`:

```shell
cd config/default/ && kustomize edit set namespace "monitor" && cd ../..
```

Run the following to deploy the operator. This will also install the RBAC manifests from config/rbac.

```shell
make deploy IMG=docker.io/$user/monitor-operator:v0.0.1
```

```shell
❯ make deploy IMG=docker.io/nonus25/monitor-operator:v0.0.1
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize || { curl -Ss "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash -s -- 3.8.7 /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin; }
cd config/manager && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize edit set image controller=docker.io/nonus25/monitor-operator:v0.0.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize build config/default | kubectl apply -f -
namespace/monitor created
customresourcedefinition.apiextensions.k8s.io/monitors.cache.monitor.com unchanged
serviceaccount/k8s-operator-monitoring-controller-manager created
role.rbac.authorization.k8s.io/k8s-operator-monitoring-leader-election-role created
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-manager-role created
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-metrics-reader created
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-proxy-role created
rolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-leader-election-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-manager-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-proxy-rolebinding created
service/k8s-operator-monitoring-controller-manager-metrics-service created
deployment.apps/k8s-operator-monitoring-controller-manager created
```

Verify that the monitor-operator is up and running:

```shell
❯ kubectl get deployment -n monitor
Alias tip: kgd -n monitor
NAME                                         READY   UP-TO-DATE   AVAILABLE   AGE
k8s-operator-monitoring-controller-manager   1/1     1            1           107s
```

### Create a Monitor CR

Update the sample Monitor CR manifest at `config/samples/cache_v1alpha1_monitor.yaml` and define the `spec` as the following:

```yaml
apiVersion: cache.monitor.com/v1alpha1
kind: Monitor
metadata:
  labels:
    app.kubernetes.io/name: monitor
    app.kubernetes.io/instance: monitor-sample
    app.kubernetes.io/part-of: k8s-operator-monitoring
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: k8s-operator-monitoring
  name: monitor-sample
spec:
  # TODO(user): Add fields here
```

```yaml
apiVersion: cache.monitor.com/v1alpha1
kind: Monitor
metadata:
  labels:
    app.kubernetes.io/name: monitor
    app.kubernetes.io/instance: monitor-sample
    app.kubernetes.io/part-of: k8s-operator-monitoring
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: k8s-operator-monitoring
  name: monitor-sample
  namespace: monitor
spec:
  size: 3
```

Create image the one we will be using for this deployment, for simplicity lets use same application with `Dockerfile` modifications

```Dockerfile
FROM golang:1.19 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o monitor main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/monitor .
USER 65532:65532

ENTRYPOINT ["/monitor"]
```

then repeat same process for generating the image:

```shell
export user=nonus25

make docker-build IMG=docker.io/$user/monitor-operator:v0.0.2
make docker-push IMG=docker.io/$user/monitor-operator:v0.0.2
```

There where we running `make run` we need to export this env variable `MONITOR_IMAGE`

```shell
export MONITOR_IMAGE="$user/monitor-operator:v0.0.2"
```

and then start again `make run`

```shell
❯ vim test/sample/manifest/monitor-sample.yaml
❯ kubectl apply -f test/sample/manifest/monitor-sample.yaml
Alias tip: kaf test/sample/manifest/monitor-sample.yaml
monitor.cache.monitor.com/monitor-sample created
```

and we should see this kind of output

```shell
NAME                                                          READY   STATUS    RESTARTS        AGE
k8s-operator-monitoring-controller-manager-75485bc78c-bzh4m   2/2     Running   0               2m13s
monitor-sample-5d9786b9fd-rfd8j                               1/1     Running   7 (6m24s ago)   26m
monitor-sample-5d9786b9fd-xxdxz                               1/1     Running   7 (6m20s ago)   26m
monitor-sample-5d9786b9fd-4gltt                               1/1     Running   7 (6m2s ago)    26m
```

## Metrics

By default, controller-runtime builds a global prometheus registry and publishes a collection of performance metrics for each controller. [default merics](https://book.kubebuilder.io/reference/metrics-reference)

Make sure cluster role for monitoring been installed `<namePrefix>-metrics-reader`, `namePrefix` we can find in `config/default/kustomization.yaml`

```shell
❯ k get clusterrole | grep metrics-reader
system:aggregated-metrics-reader                                       2023-07-05T07:07:33Z
k8s-operator-monitoring-metrics-reader                                 2023-10-03T09:33:55Z
```

we can see role is installed `k8s-operator-monitoring-metrics-reader`

### Exporting Metrics for Prometheus

1. Install Prometheus and Prometheus Operator
2. Uncomment the line `- ../prometheus` in the `config/default/kustomization.yaml`. It creates the `ServiceMonitor` resource which enables exporting the metrics.

To add needed CRD's for metrics we need to install Prometheus

```shell
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus-operator prometheus-community/prometheus-operator
```

Or we can try to install `bundle.yaml` from `https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/master/bundle.yaml`

```shell
❯ kaf https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml
customresourcedefinition.apiextensions.k8s.io/alertmanagerconfigs.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/alertmanagers.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/podmonitors.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/probes.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/prometheusrules.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/scrapeconfigs.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/servicemonitors.monitoring.coreos.com created
customresourcedefinition.apiextensions.k8s.io/thanosrulers.monitoring.coreos.com created
clusterrolebinding.rbac.authorization.k8s.io/prometheus-operator unchanged
clusterrole.rbac.authorization.k8s.io/prometheus-operator unchanged
deployment.apps/prometheus-operator unchanged
serviceaccount/prometheus-operator unchanged
service/prometheus-operator unchanged
Error from server (Invalid): error when creating "https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml": CustomResourceDefinition.apiextensions.k8s.io "prometheusagents.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
Error from server (Invalid): error when creating "https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml": CustomResourceDefinition.apiextensions.k8s.io "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
```

Not sure about those errors yet

```shell
Error from server (Invalid): error when creating "https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml": CustomResourceDefinition.apiextensions.k8s.io "prometheusagents.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
Error from server (Invalid): error when creating "https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml": CustomResourceDefinition.apiextensions.k8s.io "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
```

and then this works

```shell
❯ make deploy IMG=docker.io/nonus25/monitor-operator:v0.0.4
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
test -s /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize || { curl -Ss "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash -s -- 3.8.7 /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin; }
cd config/manager && /home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize edit set image controller=docker.io/nonus25/monitor-operator:v0.0.4
/home/auc/go/src/github.com/nonus25/k8s-operator-monitoring/bin/kustomize build config/default | kubectl apply -f -
namespace/monitor unchanged
customresourcedefinition.apiextensions.k8s.io/monitors.cache.monitor.com unchanged
serviceaccount/k8s-operator-monitoring-controller-manager unchanged
role.rbac.authorization.k8s.io/k8s-operator-monitoring-leader-election-role unchanged
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-manager-role configured
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-metrics-reader unchanged
clusterrole.rbac.authorization.k8s.io/k8s-operator-monitoring-proxy-role unchanged
rolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-leader-election-rolebinding unchanged
clusterrolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-manager-rolebinding unchanged
clusterrolebinding.rbac.authorization.k8s.io/k8s-operator-monitoring-proxy-rolebinding unchanged
service/k8s-operator-monitoring-controller-manager-metrics-service unchanged
deployment.apps/k8s-operator-monitoring-controller-manager unchanged
servicemonitor.monitoring.coreos.com/k8s-operator-monitoring-controller-manager-metrics-monitor created
```

```shell
❯ k get servicemonitors.monitoring.coreos.com
NAME                                                         AGE
k8s-operator-monitoring-controller-manager-metrics-monitor   5m54s
```