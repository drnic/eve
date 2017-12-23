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

## Mapping File

This application will be responsible for (optionally) generating the web HTML, consuming the HTML form, generating the `go-patch` file, `git commit` to a local Git repository and performing a `git push`.

Therefore we need a standard for describing the HTML form and its mapping to a resulting `go-patch` file. The `go-patch` file will be specific to the resulting YAML files upon which it will be merged/applied.

For example, the following operation is specific to a Concourse CI BOSH deployment manifest:

```yaml
- type: replace
  path: /instance_groups/name=worker/instances
  value: ((worker-linux-instances))
```

The same UI for a Concourse deployment that is backed by Kubernetes will require a different `go-patch` operator file since the Kubernetes YAML deployment file has a different schema.

### Pivotal Product Templates

One consideration for a mapping file could be Pivotal's Product Template metadata files ([spec](https://docs.pivotal.io/tiledev/2-0/product-template-reference.html)).

They are made up of:

* [Forms](https://docs.pivotal.io/tiledev/2-0/product-template-reference.html#form-properties) - collections of UI elements shown on different tabs (and thus each are saved together as a 'form').
* [Property Blueprints](https://docs.pivotal.io/tiledev/2-0/product-template-reference.html#property-blueprints) - the mapping of form UI to data types and snippets of YAML for the resulting `bosh deploy`
* [Configurable Properties](https://docs.pivotal.io/tiledev/2-0/product-template-reference.html#configurable-props) - form element types (string, integer, etc) that include pre-defined validations for the web UI
* [Job Types](https://docs.pivotal.io/tiledev/2-0/product-template-reference.html#job-types) - the different VMs that make up the resulting running system

In our problem space, the Job Types are replaced by the `go-patch` operator file. We are assuming that Job Types (called Instance Groups in BOSH for example), are defined elsewhere. Our responsibility is to create an Operator file that will successfully merge with that base YAML file.

Pivotal's implementation of Product Templates, commercially known as Ops Manager and called `tempest` within the Ops Manager VM, is written as a Ruby on Rails web app and is proprietary to Pivotal.

### BYO HTML + JavaScript validations

There is this thing called HTML, with bells and whistles like JavaScript and CSS, that is pretty handy for describing forms that appear in web browsers. Since supporting BYO HTML forms is in scope, perhaps we just start with it and worry about fancy generation of Forms + Property Blueprints + Configurable Properties later (or never). Let the author of the Web UI provide their own HTML + Javascript (for validations).

In this scenario, what exactly is `go-patch-web-ui`? Perhaps its not actually the Web UI. Perhaps its just a CLI that converts the incoming HTML form into an Operator file performs the `git` commands. Perhaps it is packaged as a Cloud Foundry buildpack to make the CLI available to the wrapper web app; or manually packaged by the wrapper web app in a Docker image.