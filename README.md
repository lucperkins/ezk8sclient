# ezk8sclient

The easiest Go client for the Kubernetes API. ezk8sclient enables you to very quickly create a "default" client for the Kubernetes API:

```go
import "github.com/lucperkins/ezk8sclient"

client, err := NewClient()
```

Ordinarily you'd have to do this:

```go
import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

flags := genericclioptions.NewConfigFlags()
restConfig, err := flags.ToRESTConfig()
if err != nil {
        return nil, err
}

client, err := kubernetes.NewForConfig(restConfig)
if err != nil {
        return nil, err
}
```

Just a quick time saver for y'all.
