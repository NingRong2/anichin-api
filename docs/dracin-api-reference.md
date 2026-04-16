# Dracin API Reference — Anichin Official

> Referensi lengkap **Dracin API** (Anichin Official API). Halaman ini me-redirect ke dokumentasi lengkap di [README.md](../README.md). **Dracin API** dan **Anichin API** adalah API yang sama — Dracin adalah nama alternatif dari aggregator short drama resmi Anichin.

## Quick Reference

| Hal           | Nilai                                          |
| ------------- | ---------------------------------------------- |
| Base URL      | `https://api.anichin.bio`                      |
| Auth Header   | `X-API-Key: TRIAL-ANICHIN-2026`                |
| Rate Limit    | 50 req/min                                     |
| WebSocket     | `wss://api.anichin.bio/ws`                     |
| Sumber Drama  | 15 (DramaBox, ReelShort, ShortMax, dst.)       |

## Endpoint per Sumber

```
GET /{source}/trending
GET /{source}/foryou?page=1
GET /{source}/search?query={keyword}
GET /{source}/detail?id={dramaId}
GET /{source}/episode?id={dramaId}&ep={n}
```

Endpoint tambahan (per sumber): `hotrank`, `recommended`, `homepage`, `latest`, `hot`, `new`, `hot+`, `drama18`, `romance`.

## Lihat Juga

- [README.md](../README.md) — referensi lengkap **Anichin API / Dracin API**
- [anichin-api-quickstart.md](anichin-api-quickstart.md) — quickstart 5 menit
- [examples/](../examples/) — contoh kode multi-bahasa
