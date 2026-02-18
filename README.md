<div align="center">

```
    ██╗  ██╗██╗     ██╗██████╗ ███████╗██╗  ██╗     ██╗ █████╗ ██╗     ███████╗
    ██║  ██║██║     ██║██╔══██╗██╔════╝██║ ██╔╝     ██║██╔══██╗██║     ██╔════╝
    ███████║██║     ██║██║  ██║███████╗█████╔╝      ██║███████║██║     █████╗
    ██╔══██║██║     ██║██║  ██║╚════██║██╔═██╗ ██   ██║██╔══██║██║     ██╔══╝
    ██║  ██║███████╗██║██████╔╝███████║██║  ██╗╚█████╔╝██║  ██║███████╗██║
    ╚═╝  ╚═╝╚══════╝╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚════╝ ╚═╝  ╚═╝╚══════╝╚═╝
```

### *Odin's Throne — The All-Seeing Seat*

---

**A Kubernetes Laboratory Monorepo**

[![Author](https://img.shields.io/badge/Author-Yumiko%20Sturluson-ff69b4?style=for-the-badge)](https://github.com/yumikokawaii)
[![License](https://img.shields.io/badge/License-Private-9370DB?style=for-the-badge)]()
[![K8s](https://img.shields.io/badge/Kubernetes-k3d-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)]()
[![Coffee](https://img.shields.io/badge/Powered%20by-Coffee-brown?style=for-the-badge)]()

</div>

---

## About

> *"From Hlidskjalf, Odin gazes upon all the nine worlds..."*

Welcome to **Hlidskjalf**!

A monorepo containing multiple mock services and Kubernetes workloads designed for experimentation, learning, and development purposes. Powered by **k3d**, this project provides a lightweight yet powerful local Kubernetes environment.

Think of it as your own little world to orchestrate containers!

---

### Bifrost — The Rainbow Bridge

API gateway that routes HTTP/gRPC requests to backend services by URL path prefix. Discovers services dynamically via Kubernetes API annotations (`bifrost.io/port`, `bifrost.io/prefixes`) and load-balances across pod endpoints using round-robin. Supports HTTP/2 (h2c).

### Skidbladnir — The Ship That Always Finds Favorable Wind

Egress-only network sidecar proxy deployed alongside every service. Intercepts all outbound TCP traffic via iptables OUTPUT chain redirect (port 15001, UID 1337 exempt).

**Client-side load balancing:** HTTP/2 clients multiplex all requests over a single long-lived connection, which kube-proxy pins to one backend pod — making replica scaling useless. Skidbladnir solves this by resolving service hostnames to pod IPs directly via the Kubernetes Endpoints API and round-robining each request across all available pods. Services are discovered lazily on first request, and endpoint watches keep the pod list in sync as pods scale up, down, or restart.

Also provides inbound rate limiting via a distributed fair-share protocol between sidecar instances.

---

## Author

<div align="center">

**~ Yumiko Sturluson ~**

*Software Engineer*

*コードよ、わがまほうとなれ*

</div>

---

## Getting Started

Interested in exploring the nine worlds?

Reach out for more information:

**Contact:** [yumiko.stl@gmail.com](mailto:yumiko.stl@gmail.com)

---

<div align="center">

*~ Built with mass amounts of coffee ~*

</div>

