# 03 · RBAC

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/03-rbac/setup.sh
```

This creates namespace `ci` with the Deployment `builder` and the Secret `registry-token`.

## Task

1. Create a ServiceAccount `deployer` in `ci`.
2. Create a Role `deployer` in `ci` that allows `get`, `list`, `create`, `update` and `patch` on Deployments and StatefulSets, and nothing else. Bind it to the ServiceAccount.
3. Create a ClusterRole `node-viewer` that allows `get` and `list` on nodes. Bind it to the ServiceAccount with a ClusterRoleBinding `deployer-node-viewer`.
4. Write a kubeconfig to `/opt/cka/deployer.kubeconfig` that authenticates as `deployer` with a token valid for 24h.
5. Using only that kubeconfig, scale `ci/builder` to 2 replicas.

## Done when

- `kubectl auth can-i create deployments -n ci --as=system:serviceaccount:ci:deployer` → `yes`
- `kubectl auth can-i get secrets -n ci --as=system:serviceaccount:ci:deployer` → `no`
- `kubectl auth can-i create deployments -n default --as=system:serviceaccount:ci:deployer` → `no`
- `kubectl auth can-i list nodes --as=system:serviceaccount:ci:deployer` → `yes`
- `kubectl --kubeconfig /opt/cka/deployer.kubeconfig -n ci get deploy builder` shows 2/2.

## Solution

Put your manifests and commands in `solution/`.
