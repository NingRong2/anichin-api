# Changelog — Anichin API Documentation

Semua perubahan penting pada dokumentasi **Anichin API** (Dracin API) dicatat di file ini.

Format mengacu pada [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) dan repository ini menggunakan [Semantic Versioning](https://semver.org/).

---

## [Unreleased]

### Added
- Banner SVG resmi Anichin API di `assets/banner.svg`
- Sitemap untuk GitHub Pages (`docs/sitemap.xml`)
- Contoh kode multi-bahasa di folder `examples/`

---

## [1.0.0] — 2026-04-16

### Added
- Dokumentasi lengkap **Anichin API / Dracin API** untuk 15 sumber short drama:
  DramaBox, ReelShort, ShortMax, NetShort, GoodShort, DramaWave, FlickReels,
  FreeReels, StardustTV, iDrama, DramaNova, StarShort, DramaBite, Melolo, MoboReels.
- 5 endpoint inti per sumber: `trending`, `foryou`, `search`, `detail`, `episode`.
- Endpoint tambahan per sumber: `hotrank`, `recommended`, `homepage`, `latest`,
  `hot`, `new`, `hot+`, `drama18`, `romance`.
- Dokumentasi WebSocket `wss://api.anichin.bio/ws` (auth, execute, ratelimit).
- Code snippets: cURL, Node.js, Python, Go, PHP.
- Trial token `TRIAL-ANICHIN-2026` dengan rate limit 50 req/min.
- File pendukung repo: `LICENSE` (MIT), `CONTRIBUTING.md`, `SECURITY.md`,
  `.github/FUNDING.yml`, banner & sitemap.

### SEO
- Topics & description repository GitHub dioptimalkan untuk keyword
  *Anichin*, *Anichin API*, *Anichin Official*, *Dracin*, *Dracin API*.
- Schema.org `WebAPI` & `TechArticle` di GitHub Pages.

---

[Unreleased]: ../../compare/v1.0.0...HEAD
[1.0.0]: ../../releases/tag/v1.0.0
