# kubeaks init


This command will check if clustername is exists in output.yaml, if exists kubeaks will login into Azure and get AKS credentials and output it into kubeconfig.yaml after that it will convert the kubeconfig.yaml using `kubelogin`.

```mermaid
flowchart TD
    A[kube cluster switch --clustername 'name'] -->
    B(read output.yaml) --> 
    C{found clustername}
    C --> C1[Exit]
    C --> C2[Login Azure]
    C2 --> d(Create kubeconfig.yaml)
    d --> e(Kubelogin -l azure)
```
## Todo:
* todo: kubeconfig path take from output.yaml
* todo: try using Stdin, Stdout, Stderr :D
* todo: what condition kubeaks will request credentials
* todo: if name already exists in kubeconfig, should not make a duplication