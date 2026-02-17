<div align="center">

# Bifrost

### *The Rainbow Bridge — Gateway to All Realms*

---

**API Gateway & Service Router**

</div>

---

## About

> *"Bifrost trembles as the sons of Muspell ride over it..."*

**Bifrost** is the API gateway of Hlidskjalf. It watches the Kubernetes realm for services, discovers them by their annotations, and routes requests to the right destination.

It speaks both HTTP/1.1 and HTTP/2 cleartext (h2c), choosing the right tongue for each traveler.

---

## How It Works

Bifrost discovers backend services via annotations on Kubernetes Service objects (headless services). It watches for changes in real time, resolves requests by longest prefix match, and load-balances across individual pod IPs with round-robin.

Supported annotations:
- `bifrost.io/port` — the port to forward traffic to
- `bifrost.io/prefixes` — JSON array of URL path prefixes to match

Bifrost selects the transport based on the incoming protocol: HTTP/1.1 requests use a pooled `http.Transport`, while HTTP/2 (gRPC) requests use `http2.Transport` with h2c. The HTTP/1.1 transport maintains a configurable pool of idle connections to reduce TCP handshake overhead.

---

## Configuration

| Variable | Default | Description |
|---|---|---|
| `SERVER__HTTP` | `0.0.0.0:10080` | Listen address |
| `NAMESPACE` | `hlidskjalf` | Kubernetes namespace to watch |
| `TRANSPORT__MAX_IDLE_CONNS` | `200` | Max idle connections (all hosts) |
| `TRANSPORT__MAX_IDLE_CONNS_PER_HOST` | `50` | Max idle connections per backend pod |

---

<div align="center">

*~ The bridge that connects all services ~*

</div>
