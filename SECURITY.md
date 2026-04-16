# Security Policy — Anichin API

## Versi yang Didukung

| Versi          | Status       |
| -------------- | ------------ |
| Anichin API v1 | ✅ Supported |
| Trial Token    | ✅ Supported |

## Melaporkan Kerentanan

Jika Anda menemukan kerentanan keamanan pada **Anichin API** (Dracin API), jangan buka issue publik. Sebagai gantinya:

1. **Telegram (preferred)**: kirim DM ke [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot) dengan prefix `[SECURITY]`.
2. Sertakan:
   - Deskripsi singkat kerentanan
   - Endpoint Anichin yang terdampak (mis. `/dramabox/episode`)
   - Langkah reproduksi (PoC)
   - Dampak potensial
   - Rekomendasi mitigasi (opsional)

## SLA Respon

| Severity | Respon Awal | Patch Target |
| -------- | ----------- | ------------ |
| Critical | < 24 jam    | < 7 hari     |
| High     | < 48 jam    | < 14 hari    |
| Medium   | < 7 hari    | < 30 hari    |
| Low      | < 14 hari   | best-effort  |

## Scope

**In-scope:**
- Endpoint `api.anichin.bio/*`
- WebSocket `wss://api.anichin.bio/ws`
- Token leak / privilege escalation
- Rate limit bypass
- Injection (NoSQL, command, dsb.)

**Out-of-scope:**
- Konten upstream (DramaBox, ReelShort, dll.) — lapor langsung ke penyedia masing-masing
- Telegram bot (`@Anichin_Premium_Bot`) — gunakan kanal Telegram
- DoS / brute-force tanpa bukti dampak nyata

## Disclosure

Anichin mengikuti **Coordinated Disclosure**: laporan akan di-patch dahulu sebelum dipublikasikan. Reporter akan disebut di [CHANGELOG.md](CHANGELOG.md) (opt-in).
