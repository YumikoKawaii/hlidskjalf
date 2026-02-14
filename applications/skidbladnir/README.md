<div align="center">

# Skidbladnir

### *The Folding Ship — Always Finds Favorable Wind*

---

**Egress Sidecar Proxy**

</div>

---

## About

> *"Skidbladnir is the best of ships, crafted with cunning..."*

**Skidbladnir** is a transparent egress proxy that rides alongside application containers as a sidecar. An init container sets up iptables rules to intercept all outbound TCP traffic and redirect it through the proxy.

It speaks both HTTP/1.1 and h2c, matching the protocol of whatever passes through.

---

## How It Works

The init container creates iptables rules on the OUTPUT chain:

- Traffic from UID 1337 (Skidbladnir itself) passes through untouched
- Loopback traffic is excluded
- Everything else is redirected to port 15001

Skidbladnir reads the `Host` header and forwards the request to its original destination. Currently a transparent passthrough — no retries, timeouts, or mTLS.

---

## Configuration

| Variable | Default | Description |
|---|---|---|
| `SKIDBLADNIR_SERVER__HTTP` | `0.0.0.0:15001` | Listen address |

---

<div align="center">

*~ Small enough to fold, swift enough to carry ~*

</div>
