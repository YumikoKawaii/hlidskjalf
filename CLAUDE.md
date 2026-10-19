# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Communication
- Be terse. No preamble, no summary at the end.
- Don't explain what you're about to do, just do it.
- No "I'll now...", "Let me...", "Here's what I did..."
- Only output what was asked for.

## Shell behavior
- Never run background shell processes or jobs
- Never use & to background commands
- Never start monitoring loops or watchers
- Run commands synchronously, wait for output, done

## Project Overview

**Hlidskjalf** is a Kubernetes certification lab (CKA first, CKS later), named after Odin's throne in Norse mythology.

## CKA lab (`cka/`)

Three Lima VMs (vz, Ubuntu 24.04 arm64) on the `lima:user-v2` network:

| Node    | IP            | CPU | RAM  | Disk  |
|---------|---------------|-----|------|-------|
| cp1     | 192.168.104.1 | 2   | 4 GB | 20 GB |
| worker1 | 192.168.104.3 | 2   | 3 GB | 20 GB |
| worker2 | 192.168.104.4 | 2   | 3 GB | 20 GB |

- `cka/up.sh [node...]` creates/starts VMs and runs `cka/node-setup.sh` inside each (default: all three).
- `K8S_MINOR` (default `1.36`) selects the pkgs.k8s.io repo; one minor behind stable to leave room for `kubeadm upgrade` practice.
- Nodes get containerd (SystemdCgroup), kubeadm/kubelet/kubectl (held), crictl, etcdctl. Cluster bootstrap (`kubeadm init`/`join`, CNI) is done by hand as exam practice.
- `limactl start|stop` take one instance per call.
