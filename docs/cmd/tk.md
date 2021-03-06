## tk

Command line utility for assembling Kubernetes CD pipelines

### Synopsis

Command line utility for assembling Kubernetes CD pipelines the GitOps way.

### Examples

```
  # Check prerequisites 
  tk check --pre

  # Install the latest version of the toolkit
  tk install --version=master

  # Create a source from a public Git repository
  tk create source git webapp-latest \
    --url=https://github.com/stefanprodan/podinfo \
    --branch=master \
    --interval=3m

  # List git sources and their status
  tk get sources git

  # Trigger a git sync
  tk sync source git webapp-latest

  # Export git sources in YAML format
  tk export source git --all > sources.yaml

  # Create a kustomization for deploying a series of microservices
  tk create kustomization webapp-dev \
    --source=webapp-latest \
    --path="./deploy/webapp/" \
    --prune=true \
    --interval=5m \
    --validate=client \
    --health-check="Deployment/backend.webapp" \
    --health-check="Deployment/frontend.webapp" \
    --health-check-timeout=2m

  # Trigger a git sync and apply changes if any
  tk sync kustomization webapp-dev --with-source

  # Suspend a kustomization reconciliation
  tk suspend kustomization webapp-dev

  # Export kustomizations in YAML format
  tk export kustomization --all > kustomizations.yaml

  # Resume a kustomization reconciliation
  tk resume kustomization webapp-dev

  # Delete a kustomization
  tk delete kustomization webapp-dev

  # Delete a git source
  tk delete source git webapp-latest

  # Uninstall the toolkit and delete CRDs
  tk uninstall --crds

```

### Options

```
      --components strings   list of components, accepts comma-separated values (default [source-controller,kustomize-controller])
  -h, --help                 help for tk
      --kubeconfig string    path to the kubeconfig file (default "~/.kube/config")
      --namespace string     the namespace scope for this operation (default "gitops-system")
      --timeout duration     timeout for this operation (default 5m0s)
      --verbose              print generated objects
```

### SEE ALSO

* [tk bootstrap](tk_bootstrap.md)	 - Bootstrap commands
* [tk check](tk_check.md)	 - Check requirements and installation
* [tk completion](tk_completion.md)	 - Generates bash completion scripts
* [tk create](tk_create.md)	 - Create commands
* [tk delete](tk_delete.md)	 - Delete commands
* [tk export](tk_export.md)	 - Export commands
* [tk get](tk_get.md)	 - Get commands
* [tk install](tk_install.md)	 - Install the toolkit components
* [tk resume](tk_resume.md)	 - Resume commands
* [tk suspend](tk_suspend.md)	 - Suspend commands
* [tk sync](tk_sync.md)	 - Synchronize commands
* [tk uninstall](tk_uninstall.md)	 - Uninstall the toolkit components

