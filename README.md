# kubeaks

**Name:** Kubeaks

**Description:**
Kubeaks is a tools to help you switch between AKS cluster, even though the cluster just created.

## **Tools commands:**
* [COMMAND] $ kubeaks init
  * will appear an interactive terminal to fill up the question example subscription, rg, clustername
the information later will be saved in ~/.kubeaks.conf
* [COMMAND] $ kubeaks switch/s <CLUSTER_NAME>
  * this will switch to aks cluster <clustername>
  * When to swich a cluster, the cluster infomation will be pickup from ~/.kubeaks.conf
  * <img width="299" alt="image" src="https://github.com/aoshfan/kubeaks/assets/42878756/183fcf47-5751-4f18-a91e-505f3dc07a20">

