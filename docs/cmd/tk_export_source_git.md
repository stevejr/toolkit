## tk export source git

Export git sources in YAML format

### Synopsis

Export git sources in YAML format

```
tk export source git [name] [flags]
```

### Examples

```
  # Export all git sources
  export source git --all > sources.yaml

  # Export a git source including the SSH keys or basic auth credentials
  export source git my-private-repo --with-credentials > source.yaml

```

### Options

```
  -h, --help   help for git
```

### Options inherited from parent commands

```
      --all                  select all resources
      --components strings   list of components, accepts comma-separated values (default [source-controller,kustomize-controller])
      --kubeconfig string    path to the kubeconfig file (default "~/.kube/config")
      --namespace string     the namespace scope for this operation (default "gitops-system")
      --timeout duration     timeout for this operation (default 5m0s)
      --verbose              print generated objects
      --with-credentials     include credential secrets
```

### SEE ALSO

* [tk export source](tk_export_source.md)	 - Export source commands

