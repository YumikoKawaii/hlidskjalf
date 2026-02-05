<div align="center">

```
     ██████╗ ██████╗ ███████╗███████╗██████╗ ██╗   ██╗███████╗
    ██╔═══██╗██╔══██╗██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝
    ██║   ██║██████╔╝███████╗█████╗  ██████╔╝██║   ██║█████╗
    ██║   ██║██╔══██╗╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝
    ╚██████╔╝██████╔╝███████║███████╗██║  ██║ ╚████╔╝ ███████╗
     ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝
```

### *The Eye of Hlidskjalf*

---

**Observability Stack for the Nine Worlds**

[![Prometheus](https://img.shields.io/badge/Prometheus-v3.2.1-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)](https://prometheus.io/)
[![Grafana](https://img.shields.io/badge/Grafana-11.0.0-F46800?style=for-the-badge&logo=grafana&logoColor=white)](https://grafana.com/)
[![Loki](https://img.shields.io/badge/Loki-3.0.0-2C3239?style=for-the-badge&logo=grafana&logoColor=white)](https://grafana.com/oss/loki/)
[![Promtail](https://img.shields.io/badge/Promtail-3.0.0-2C3239?style=for-the-badge&logo=grafana&logoColor=white)](https://grafana.com/docs/loki/latest/clients/promtail/)
[![OpenTelemetry](https://img.shields.io/badge/OTel_Collector-0.120.0-7B5EA7?style=for-the-badge&logo=opentelemetry&logoColor=white)](https://opentelemetry.io/)

</div>

---

## About

> *"Odin sat upon Hlidskjalf and gazed out over all the worlds — he saw all things, the hidden and the revealed."*

**Observe** is the all-seeing eye of Hlidskjalf — a full observability platform deployed in the `observe` namespace that watches over every service in the realm.

It collects **metrics**, aggregates **logs**, and stands ready to trace every request as it travels between the worlds.

Nothing escapes its gaze.

---

## The Three Pillars

<div align="center">

| | Pillar | How |
|---|---|---|
| <img src="https://raw.githubusercontent.com/cncf/artwork/main/projects/prometheus/icon/color/prometheus-icon-color.svg" width="20"> | **Metrics** | Prometheus scrapes service metrics every 15 seconds |
| <img src="https://raw.githubusercontent.com/grafana/loki/main/docs/sources/logo.png" width="20"> | **Logs** | Promtail ships container logs to Loki in real time |
| <img src="https://raw.githubusercontent.com/cncf/artwork/main/projects/opentelemetry/icon/color/opentelemetry-icon-color.svg" width="20"> | **Traces** | OpenTelemetry Collector stands ready for distributed tracing |

</div>

Everything flows into **Grafana** — the single pane of glass where all knowledge converges.

---

## How It Sees

```
    ┌─────────────────────────────────────────────────┐
    │            hlidskjalf  (services)                │
    │                                                  │    
    └──────────┬──────────────────────┬────────────────┘
               │                      │
           stdout                 /metrics
               │                      │
               ▼                      ▼
         ┌──────────┐          ┌──────────┐
         │ Promtail │          │Prometheus│
         │(per node)│          │ (scrape) │
         └────┬─────┘          └────┬─────┘
              │                     │
              ▼                     │
         ┌──────────┐               │
         │   Loki   │               │
         └────┬─────┘               │
              │                     │
              └──────┐  ┌───────────┘
                     ▼  ▼
                ┌──────────┐
                │ Grafana  │  ← The All-Seeing Eye
                │  :3000   │
                └──────────┘
```

---

## At a Glance

<div align="center">

| Component | What it does |
|-----------|-------------|
| **Prometheus** | Pulls metrics from services, keeps 15 days of history |
| **Loki** | Stores and indexes logs, keeps 7 days of history |
| **Promtail** | Runs on every node, ships pod logs to Loki |
| **Grafana** | Queries everything, shows everything |
| **OTel Collector** | Receives OTLP traces & metrics, batches and exports |
| **Reloader** | Watches for config changes, restarts pods automatically |

</div>

---

<div align="center">

*~ きゅうつのせかいに、フリズスキャールヴからかくれるものはない ~*

</div>
