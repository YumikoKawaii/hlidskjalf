# Hlidskjalf

*Odin's Throne — The All-Seeing Seat*

Kubernetes certification labs.

## CKA

Three-node kubeadm lab on Lima VMs (Apple Silicon).

```bash
./cka/up.sh
limactl shell cp1
```

| Node    | IP            | CPU | RAM  |
|---------|---------------|-----|------|
| cp1     | 192.168.104.1 | 2   | 4 GB |
| worker1 | 192.168.104.3 | 2   | 3 GB |
| worker2 | 192.168.104.4 | 2   | 3 GB |

Practice assignments: [`cka/assignments/`](cka/assignments/).

Stop:

```bash
for n in cp1 worker1 worker2; do limactl stop $n; done
```
