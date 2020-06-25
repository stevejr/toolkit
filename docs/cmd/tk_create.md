## tk create

Create resources

### Synopsis

Create resources

### Options

```
      --export              export in YAML format to stdout
  -h, --help                help for create
      --interval duration   source sync interval (default 1m0s)
```

### Options inherited from parent commands

```
      --components strings   list of components, accepts comma-separated values (default [source-controller,kustomize-controller])
      --kubeconfig string    path to the kubeconfig file (default "~/.kube/config")
      --namespace string     the namespace scope for this operation (default "gitops-system")
      --timeout duration     timeout for this operation (default 5m0s)
      --verbose              print generated objects
```

### SEE ALSO

* [tk](tk.md)	 - Command line utility for assembling Kubernetes CD pipelines
* [tk create kustomization](tk_create_kustomization.md)	 - Create or update a Kustomization resource
* [tk create source](tk_create_source.md)	 - Create sources

