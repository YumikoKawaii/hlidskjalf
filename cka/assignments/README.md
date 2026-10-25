# CKA assignments

Each assignment has two parts:

1. `setup.sh`: run from the Mac, in the repo root. It drives the Lima VMs through `limactl shell` and puts the cluster into the starting state for the task. Some assignments plant faults, others only prepare resources.
2. `solution/`: your work goes here (commands, manifests, notes on the root cause).

Prerequisites: all three VMs running, `bootstrap.sh` run on every node, `control-plane.sh` run on `cp1`, workers joined. Answer files go under `/opt/cka/` on `cp1` (setup creates it).

```bash
./cka/assignments/01-etcd-backup-restore/setup.sh
limactl shell cp1
```

| #  | Assignment | Domain | Faults |
|----|------------|--------|--------|
| 01 | [etcd backup & restore](01-etcd-backup-restore/) | Cluster Architecture | none |
| 02 | [Cluster upgrade](02-cluster-upgrade/) | Cluster Architecture | none |
| 03 | [RBAC](03-rbac/) | Cluster Architecture | none |
| 04 | [Helm](04-helm/) | Cluster Architecture | none |
| 05 | [Kustomize](05-kustomize/) | Cluster Architecture | none |
| 06 | [CRDs & operators](06-crds-operators/) | Cluster Architecture | none |
| 07 | [Rollout & rollback](07-rollout-rollback/) | Workloads & Scheduling | none |
| 08 | [HPA & metrics-server](08-hpa-metrics/) | Workloads & Scheduling | none |
| 09 | [Scheduling](09-scheduling/) | Workloads & Scheduling | none |
| 10 | [Native sidecar](10-native-sidecar/) | Workloads & Scheduling | none |
| 11 | [ResourceQuota & LimitRange](11-quota-limitrange/) | Workloads & Scheduling | none |
| 12 | [Services & DNS](12-services-dns/) | Services & Networking | none |
| 13 | [NetworkPolicy](13-network-policy/) | Services & Networking | none |
| 14 | [Gateway API](14-gateway-api/) | Services & Networking | none |
| 15 | [CoreDNS](15-coredns/) | Services & Networking | none |
| 16 | [PV & PVC](16-pv-pvc/) | Storage | none |
| 17 | [StorageClass](17-storageclass/) | Storage | none |
| 18 | [Broken kubelet](18-broken-kubelet/) | Troubleshooting | yes |
| 19 | [Broken scheduler](19-broken-scheduler/) | Troubleshooting | yes |
| 20 | [Broken Service](20-broken-service/) | Troubleshooting | yes |
| 21 | [CrashLoopBackOff](21-crashloop/) | Troubleshooting | yes |
| 22 | [Resource usage](22-resource-usage/) | Troubleshooting | none |
| 23 | [Node network](23-node-network/) | Troubleshooting | yes |
