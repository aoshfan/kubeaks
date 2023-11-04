[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=bugs)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=coverage)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=aoshfan_kubeaks&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=aoshfan_kubeaks)  

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
