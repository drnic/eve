# Go Patch Web UI

This repo is an initial experiment in generating a Web UI to allow lay people to modify a YAML-based deployment (e.g. Kubernetes, BOSH deploy, or BOSH env). They make changes, press "Apply", and eventually the deployment is modified.

The intermediate stage after user changes in the UI is a BOSH operator file ([go-patch](https://github.com/cppforlife/go-patch)) which is subsequently applied to a `bosh deploy base.yml -o ui.yml` or `bosh create-env base.yml -o ui.yml` command which is triggered by a CI system. 

## Scenario

Consider a Web UI for a Concourse CI deployment to allow scaling of workers:

```
Workers/Linux:
  Instances: 5
  Instance Type: m4.xlarge

Workers/Windows:
  Instances: 1
  Instance Type: m4.xlarge

<Apply Changes>
```

Internally, the "Apply Changes" button would convert the form values into a `go-patch` file that is specific to the Concourse CI deployment (assuming it is a BOSH deployment; but perhaps could be a YAML manifest for a Kubernetes deployment)

```yaml
- type: replace
  path: /instance_groups/name=worker/instances
  value: ((worker-linux-instances))

- type: replace
  path: /instance_groups/name=worker/vm_type
  value: ((worker-linux-vm-type))

- type: replace
  path: /instance_groups/name=windows-worker/instances
  value: ((worker-windows-instances))

- type: replace
  path: /instance_groups/name=windows-worker/vm_type
  value: ((worker-windows-vm-type))
```

This file would be saved to a Git repository, rather than a database. In turn, it is a assumed that CI system will be watching the Git repository, see the new Git commit, and re-deploy the target system.
