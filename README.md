# Eve

Create your own lovel web UIs for YAML-based deployments, such as BOSH, Kubernetes, and Cloud Foundry.

```bash
$ eve convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge'
```

Prints to stdout:

```yaml
- type: replace
  path: /instance_groups/name=worker/instances
  value: "5"
- type: replace
  path: /instance_groups/name=worker/vm_type
  value: m4.xlarge
```

Or to write to a file:

```bash
$ eve convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge' \
  --target tmp/bosh-scaling-operator.yml
```


This repo is an initial experiment in generating a Web UI to allow lay people to modify a YAML-based deployment (e.g. Kubernetes, BOSH deploy, or BOSH env). They make changes, press "Apply", and eventually the deployment is modified.

The intermediate stage after user changes in the UI is a BOSH operator file ([go-patch](https://github.com/cppforlife/go-patch)) which is subsequently applied to a `bosh deploy base.yml -o ui.yml` or `bosh create-env base.yml -o ui.yml` command which is triggered by a CI system. 

By creating/changing a file in a Git repository, the operators of the system have an audit log of changes in the Git history. "Bob changed the instance count to 5", "Sue changed the instance count to 1".

By using the the `eve` system, we allow decoupling of a bespoke Web UI from the backend system that performs the deployment.

The decoupled nature of the wrapper Web UI and the backend deployment system (BOSH, Kubernetes, Cloud Foundry) may make it difficult to provide "state of deployment" feedback to the Web UI user.

* [CI pipeline](https://ci-ohio.starkandwayne.com/teams/cfcommunity/pipelines/eve)
* [Issues](https://github.com/starkandwayne/eve/issues)

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
  value: 5

- type: replace
  path: /instance_groups/name=worker/vm_type
  value: m4.xlarge

- type: replace
  path: /instance_groups/name=windows-worker/instances
  value: 1

- type: replace
  path: /instance_groups/name=windows-worker/vm_type
  value: m4.xlarge
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

In this scenario, what exactly is `eve`? Perhaps its not actually the Web UI. Perhaps its just a CLI that converts the incoming HTML form into an Operator file performs the `git` commands. Perhaps it is packaged as a Cloud Foundry buildpack to make the CLI available to the wrapper web app; or manually packaged by the wrapper web app in a Docker image.

## KISS

We would also just focus on an the HTML form design and its mapping to a generated Operator file as the primary problem being solved. Web UI is solved by dozens of web frameworks. Triggering of deployments to BOSH/Kubernetes/Cloud Foundry is solved by CI systems watching Git repos. This tool would solve the UI -> Operator file problem.

Perhaps this tool could even ignore that Git layer. But my guess is bespoke Web apps would just rewrite the same `git pull; git commit -a -m "user change"; git push` code over and over. So it might as well be implemented by the tool. Perhaps make it optional so end users can choose a different system for delivering the Operator file to their backend deployment system.

## HTML form examples

```html
<form action="/update-deployment" method="POST">
  <fieldset>
    <legend>Workers/Linux:</legend>
    <label>Instances:</label> <input type="text" name="workers-linux-instances" value="5">
    <label>Instance Type:</label> <input type="text" name="workers-linux-instance-type" value="m4.xlarge">
  </fieldset>
  <input type="submit" value="Apply Changes">
</form>
```

The wrapper web app will receive the form POST to `/update-deployment` and contain form fields `workers-linux-instances` and `workers-linux-instance-type`. It will then pass these to the `eve` CLI to generate the Operator file (given a mapping file):

```bash
eve \
  --mapping path/to/mapping.yml \
  --inputs '{"workers-linux-instances": 5, "workers-linux-instance-type": "m4.xlarge"}' \
  --target path/to/ui.yml
```

Or perhaps

```bash
go run main.go convert \
  --mapping fixtures/bosh-scaling/mapping.yml \
  --inputs 'workers-linux-instances:5' \
  --inputs 'workers-linux-instance-type:m4.xlarge'
```

This will create a file at `path/to/ui.yml` similar to the following (dependent upon the mapping in `path/to/mapping.yml`):

```yaml
- type: replace
  path: /instance_groups/name=worker/instances
  value: 5

- type: replace
  path: /instance_groups/name=worker/vm_type
  value: m4.xlarge
```

## Plan

Build a nice Web UI app as a wrapper for an existing deployment (or multiple deployments). Figure out where `eve` would be hooked in to post changes.

Figure out how `eve` is hooked in to fetch current values.

Get `eve` to do the thing its supposed to do to trigger a deployment.