## tk create source

Create source commands

### Synopsis

Create source commands

### Options

```
  -h, --help   help for source
```

### Options inherited from parent commands

```
      --components strings   list of components, accepts comma-separated values (default [source-controller,kustomize-controller])
      --export               export in yaml format to stdout
      --interval duration    source sync interval (default 1m0s)
      --kubeconfig string    path to the kubeconfig file (default "~/.kube/config")
      --namespace string     the namespace scope for this operation (default "gitops-system")
      --timeout duration     timeout for this operation (default 5m0s)
      --verbose              print generated objects
```

### SEE ALSO

* [tk create](tk_create.md)	 - Create commands
* [tk create source git](tk_create_source_git.md)	 - Create or update a git source

