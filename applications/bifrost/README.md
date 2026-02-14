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

Bifrost discovers backend services via annotations on Kubernetes Service objects. It watches for changes in real time, resolves requests by longest prefix match, and load-balances across pod IPs with round-robin.

---

## Configuration

| Variable | Default | Description |
|---|---|---|
| `SERVER__HTTP` | `0.0.0.0:10080` | Listen address |
| `NAMESPACE` | `hlidskjalf` | Kubernetes namespace to watch |

---

<div align="center">

*~ The bridge that connects all services ~*

</div>
