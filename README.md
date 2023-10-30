# kubeaks

**Name:** Kubeaks

**Description:**
Kubeaks is a tools to help you switch between AKS cluster, even though the cluster just created.

## **How cluster switch happen**

1. Firstly, it will use `~/.kubeconfig` to login into the cluster.
2. If error, `kubeaks` will try to authenticate using Azure SDK, and write the login into `~/.kubeconfig`.

## **Tools commands:**

* [COMMAND] $ kubeaks init
  * will appear an interactive terminal to fill up the question example subscription, rg, clustername
the information later will be saved in _~/.kubeaks.yaml_
* [COMMAND] $ kubeaks switch/s ***CLUSTER_NAME***
  * this will switch to aks cluster ***CLUSTER_NAME***
  * When to swich a cluster, the cluster infomation will be pickup from _~/.kubeaks.yaml_
  * _~/.kubeaks.yaml_

    ```yaml
    - name: blue
        azure:
          subscription: 12345678-1111-1234-1234-1234567890ab
          resourceGroup: rg-1
          clusterName: aks-blue
        kubeconfig:
          name: rg-aks-dev-blue
    
    - name: green
        azure:
          subscription: 12345678-2222-1234-1234-1234567890ab
          resourceGroup: rg-2
          clusterName: aks-green
        kubeconfig:
          name: rg-aks-dev-green
    
    - name: red
        azure:
          subscription: 12345678-3333-1234-1234-1234567890ab
          resourceGroup: rg-1
          clusterName: aks-red
        kubeconfig:
          name: rg-aks-dev-red
    ```

## Manual way of switch AKS cluster

```bash
# 1. set subscription
$ az account set --subscription 12345678-1234-1234-1234-1234567890ab

# 2. Get AKS cluster credentials by specifing RG & AKS name
$ az aks get-credentials --resource-group dev --name blue

# 3. After getting the credentials, we use kubelogin
$ kubelogin convert-kubeconfig -l azurecli

# 4. verify
$ kubectl get namespaces
```
