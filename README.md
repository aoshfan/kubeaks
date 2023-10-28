# kubeaks

**Name:** Kubeaks

**Description:**
Kubeaks is a tools to help you switch between AKS cluster, even though the cluster just created.

**How cluster switch happen**
1. it will read from `~/.kubeconfig`

## **Tools commands:**
* [COMMAND] $ kubeaks init
  * will appear an interactive terminal to fill up the question example subscription, rg, clustername
the information later will be saved in _~/.kubeaks.yaml_
* [COMMAND] $ kubeaks switch/s <CLUSTER_NAME>
  * this will switch to aks cluster <clustername>
  * When to swich a cluster, the cluster infomation will be pickup from _~/.kubeaks.yaml_
  * _~/.kubeaks.yaml_
    ```
    name: blue
      azure:
        subscription: 12345678-1111-1234-1234-1234567890ab
    	resourceGroup: rg-1
    	clusterName: aks-blue
      kubeconfig:
        name: rg-aks-dev-blue
    
    name: green
      azure:
        subscription: 12345678-2222-1234-1234-1234567890ab
    	   resourceGroup: rg-2
    	   clusterName: aks-green
      kubeconfig:
        name: rg-aks-dev-green
    
    name: red
      azure:
        subscription: 12345678-3333-1234-1234-1234567890ab
    	   resourceGroup: rg-1
    	   clusterName: aks-red
      kubeconfig:
        name: rg-aks-dev-red
    ```

