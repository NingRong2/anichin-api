# Anichin API — Official Documentation (Dracin API)

<p align="center">
  <img src="assets/banner.svg" alt="Anichin API - Official Documentation - Dracin API banner" width="100%">
</p>

> **Anichin Official API** — REST + WebSocket aggregator untuk **15 sumber drama pendek** (short drama) dalam satu endpoint terunifikasi. Mendukung *trending*, *for you*, *search*, *detail* + daftar episode, *episode video URL* multi-quality, plus subtitle multi-bahasa.
>
> 📚 **Quickstart**: [docs/anichin-api-quickstart.md](docs/anichin-api-quickstart.md) · **Reference**: [docs/dracin-api-reference.md](docs/dracin-api-reference.md) · **Contoh kode**: [examples/](examples/)

[![Status](https://img.shields.io/badge/status-online-brightgreen)](https://api.anichin.bio)
[![License](https://img.shields.io/badge/trial-FREE%201%20HARI-orange)](https://api.anichin.bio)
[![Rate Limit](https://img.shields.io/badge/rate%20limit-50%20req%2Fmin-blue)](#rate-limit)
[![Telegram](https://img.shields.io/badge/Telegram-%40Anichin__Premium__Bot-26A5E4)](https://t.me/Anichin_Premium_Bot)

---

## Daftar Isi

- [Tentang](#tentang)
- [Base URL](#base-url)
- [Autentikasi](#autentikasi)
- [Rate Limit](#rate-limit)
- [Daftar Sumber (Sources)](#daftar-sumber-sources)
- [Format URL Endpoint](#format-url-endpoint)
- [Endpoint Inti (Base Endpoints)](#endpoint-inti-base-endpoints)
  - [GET `/{source}/trending`](#1-get-sourcetrending)
  - [GET `/{source}/foryou`](#2-get-sourceforyou)
  - [GET `/{source}/search`](#3-get-sourcesearch)
  - [GET `/{source}/detail`](#4-get-sourcedetail)
  - [GET `/{source}/episode`](#5-get-sourceepisode)
- [Endpoint Tambahan (Extra Endpoints)](#endpoint-tambahan-extra-endpoints)
- [WebSocket API](#websocket-api)
- [Format Respons & Error](#format-respons--error)
- [Contoh Penggunaan (Code Snippets)](#contoh-penggunaan-code-snippets)
  - [cURL](#curl)
  - [JavaScript / Node.js](#javascript--nodejs)
  - [Python](#python)
  - [Go](#go)
  - [PHP](#php)
- [Telegram Bot Integration](#telegram-bot-integration)
- [FAQ](#faq)
- [Catatan & Disclaimer](#catatan--disclaimer)

---

## Tentang

**Anichin API** (juga dikenal sebagai **Dracin API**) adalah REST API yang menyatukan **15 sumber short drama** populer ke dalam satu skema endpoint yang konsisten. Dengan satu *base path* yang sama, Anda bisa mengakses *trending*, *recommended*, *search*, detail drama, daftar episode, hingga URL video multi-quality (1080p / 720p / 540p) dari banyak platform sekaligus.

**Fitur utama:**

- 15 sumber drama pendek terunifikasi
- Skema endpoint identik antar-sumber → ganti `{source}` saja
- Multi-quality video (1080p / 720p / 540p)
- Subtitle multi-bahasa (sampai 23 bahasa, tergantung sumber)
- Mendukung **REST** dan **WebSocket** (`/ws`)
- Free Trial Token tersedia
- Dokumentasi interaktif live di [api.anichin.bio](https://api.anichin.bio)

---

## Base URL

```
https://api.anichin.bio
```

Semua endpoint berbasis HTTPS dan mengembalikan `application/json`.

CORS aktif untuk semua origin:

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, OPTIONS
Access-Control-Allow-Headers: Content-Type, X-API-Key
```

---

## Autentikasi

Setiap request **wajib** menyertakan API key pada header `X-API-Key`.

| Header        | Wajib | Nilai                                                                  |
| ------------- | :---: | ---------------------------------------------------------------------- |
| `X-API-Key`   | ✅    | Trial token: `TRIAL-ANICHIN-2026` (gratis 1 hari, rate-limited)        |
| `User-Agent`  | ✅    | Disarankan UA browser, mis. `Mozilla/5.0` — request tanpa UA bisa di-block sebagai `forbidden` |

### Trial Token

```
TRIAL-ANICHIN-2026
```

> Token trial gratis berlaku **1 hari**, dengan rate limit **50 req/menit**. Untuk akses tanpa batas / produksi, hubungi via [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot).

Contoh:

```bash
curl -s "https://api.anichin.bio/dramabox/trending" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" \
  -H "User-Agent: Mozilla/5.0"
```

---

## Rate Limit

| Limit               | Nilai           |
| ------------------- | --------------- |
| Maksimum per menit  | **50 req/min**  |
| Window              | Sliding 60 detik |

Saat rate limit terlewati, response berisi field `error` + `reset` (detik tersisa sebelum quota di-reset). Contoh dari channel WebSocket:

```json
{ "type": "ratelimit", "remaining": 0, "limit": 50, "reset": 38 }
```

---

## Daftar Sumber (Sources)

| `{source}`     | Nama Platform | Contoh ID Drama         | Endpoint Tambahan                              |
| -------------- | ------------- | ----------------------- | ---------------------------------------------- |
| `dramabox`     | DramaBox      | `42000007806`           | `hotrank`, `recommended`                       |
| `reelshort`    | ReelShort     | `699d1eefa3a7262cff05534b` | `homepage`                                  |
| `shortmax`     | ShortMax      | `18854`                 | `recommended`, `homepage`                      |
| `netshort`     | NetShort      | `2034157133506805762`   | `recommended`                                  |
| `goodshort`    | GoodShort     | `31001188126`           | `homepage`                                     |
| `dramawave`    | DramaWave     | `LeMYdgoXZM`            | `recommended`                                  |
| `flickreels`   | FlickReels    | `5672`                  | `hotrank`                                      |
| `freereels`    | FreeReels     | `51bAUXzvfP`            | `homepage`, `hotrank`                          |
| `stardusttv`   | StardustTV    | `146`                   | —                                              |
| `idrama`       | iDrama        | `160000641712`          | —                                              |
| `dramanova`    | DramaNova     | `102062`                | `hot`, `new`, `hot+`, `drama18`, `romance`     |
| `starshort`    | StarShort     | `j0NM`                  | `hotrank`, `latest`                            |
| `dramabite`    | DramaBite     | `15384`                 | `hotrank`, `latest`                            |
| `melolo`       | Melolo        | `7522723499182394385`   | `latest`                                       |
| `moboreels`    | MoboReels     | `41896322`              | —                                              |

> Setiap sumber otomatis punya 5 endpoint inti (lihat di bawah). Endpoint tambahan adalah ekstra spesifik per sumber.

---

## Format URL Endpoint

Skema URL semua endpoint:

```
GET https://api.anichin.bio/{source}/{path}?{query_params}
```

- `{source}` — kunci sumber dari [tabel di atas](#daftar-sumber-sources), mis. `dramabox`.
- `{path}` — nama endpoint, mis. `trending`, `search`, `detail`, `episode`.
- `{query_params}` — parameter query, mis. `?id=42000008521&ep=1`.

Contoh:

```
https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1
https://api.anichin.bio/shortmax/search?query=love
https://api.anichin.bio/dramanova/hot?page=2
```

---

## Endpoint Inti (Base Endpoints)

Lima endpoint berikut **tersedia untuk semua sumber** dengan kontrak yang sama.

### 1. `GET /{source}/trending`

Daftar drama yang sedang trending di sumber tersebut.

**Query Parameters:** *(tidak ada)*

**Contoh request:**

```bash
curl -s "https://api.anichin.bio/dramabox/trending" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" \
  -H "User-Agent: Mozilla/5.0"
```

**Contoh response (200):**

```json
{
  "code": 200,
  "hasMore": false,
  "items": [
    {
      "id": "42000008521",
      "dramaId": "42000008521",
      "title": "Di Balik Ruang Rahasia CEO",
      "description": "Saat karyawan magang, Mila, ...",
      "synopsis": "Saat karyawan magang, Mila, ...",
      "cover": "https://hwztchapter.dramaboxdb.com/.../42000008521.jpg",
      "posterImg": "https://hwztchapter.dramaboxdb.com/.../42000008521.jpg",
      "episodes": 74,
      "totalEpisodes": 74,
      "isCompleted": "1",
      "defaultLanguage": "in",
      "categoryNames": [],
      "publishedAt": "",
      "viewCount": 0,
      "likeCount": 0,
      "favoriteCount": 0
    }
  ]
}
```

---

### 2. `GET /{source}/foryou`

Rekomendasi personal/feed berbasis halaman.

**Query Parameters:**

| Param  | Tipe   | Wajib | Default | Deskripsi                  |
| ------ | ------ | :---: | ------- | -------------------------- |
| `page` | int    |       | `1`     | Nomor halaman (paginated). |

**Contoh:**

```bash
curl -s "https://api.anichin.bio/dramabox/foryou?page=1" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"
```

Struktur response identik dengan `/trending` (`items[]` + `hasMore`).

---

### 3. `GET /{source}/search`

Cari drama berdasarkan kata kunci.

**Query Parameters:**

| Param   | Tipe   | Wajib | Default | Deskripsi                        |
| ------- | ------ | :---: | ------- | -------------------------------- |
| `query` | string | ✅    | `love`  | Kata kunci pencarian (URL-encoded). |

**Contoh:**

```bash
curl -s "https://api.anichin.bio/shortmax/search?query=ceo" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"
```

Response: list drama yang cocok (`items[]`).

---

### 4. `GET /{source}/detail`

Detail drama lengkap + daftar episode.

**Query Parameters:**

| Param | Tipe   | Wajib | Deskripsi                                    |
| ----- | ------ | :---: | -------------------------------------------- |
| `id`  | string | ✅    | ID drama (lihat kolom *Contoh ID* di tabel sumber). |

**Contoh:**

```bash
curl -s "https://api.anichin.bio/dramabox/detail?id=42000008521" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"
```

**Contoh response (200):**

```json
{
  "code": 200,
  "data": {
    "id": "42000008521",
    "dramaId": "42000008521",
    "title": "Di Balik Ruang Rahasia CEO",
    "description": "...",
    "cover": "https://.../42000008521.jpg",
    "defaultLanguage": "in",
    "episodes": [
      { "episodeNumber": 1, "number": 1, "title": "Episode 1", "locked": false, "videoUrl": "" },
      { "episodeNumber": 2, "number": 2, "title": "Episode 2", "locked": false, "videoUrl": "" },
      { "episodeNumber": 6, "number": 6, "title": "Episode 6", "locked": true, "qualityList": [], "subtitles": [], "videoUrl": "" }
    ]
  }
}
```

> Episode dengan `"locked": true` membutuhkan tier berbayar / unlock di sumber asalnya.

---

### 5. `GET /{source}/episode`

Ambil URL video episode + daftar quality dan subtitle.

**Query Parameters:**

| Param | Tipe   | Wajib | Default | Deskripsi                                |
| ----- | ------ | :---: | ------- | ---------------------------------------- |
| `id`  | string | ✅    | —       | ID drama.                                |
| `ep`  | int    | ✅    | `1`     | Nomor episode (mulai dari 1).            |

**Contoh:**

```bash
curl -s "https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"
```

**Contoh response (200):**

```json
{
  "code": 200,
  "msg": "SUCCESS",
  "episodeNumber": 1,
  "number": 1,
  "locked": false,
  "videoUrl": "https://thwztvideo.dramaboxdb.com/.../700545423.720p.narrowv3.mp4",
  "qualityList": [
    { "label": "1080p", "url": "https://.../700545423.1080p.nav2.mp4" },
    { "label": "720p",  "url": "https://.../700545423.720p.narrowv3.mp4", "isDefault": true },
    { "label": "540p",  "url": "https://.../700545423.540p.narrowv2.mp4" }
  ]
}
```

> Beberapa sumber juga mengembalikan field `subtitles[]` (multi-bahasa) di dalam response episode.

---

## Endpoint Tambahan (Extra Endpoints)

Sebagai pelengkap, beberapa sumber memiliki endpoint khusus.

| Path           | Sumber yang Mendukung                                                                 | Params              | Deskripsi                  |
| -------------- | ------------------------------------------------------------------------------------- | ------------------- | -------------------------- |
| `hotrank`      | `dramabox`, `flickreels`, `freereels`, `starshort`, `dramabite`                       | —                   | Ranking drama terpopuler   |
| `recommended`  | `dramabox`, `shortmax`, `netshort`, `dramawave`                                       | —                   | Drama yang direkomendasikan |
| `homepage`     | `shortmax`, `goodshort`, `freereels`, `reelshort`                                     | —                   | Feed halaman utama         |
| `latest`       | `starshort`, `dramabite`, `melolo`                                                    | `page=1`            | Rilis terbaru (paginated)  |
| `hot`          | `dramanova`                                                                           | `page=1`            | Hot / Popular              |
| `new`          | `dramanova`                                                                           | `page=1`            | Rilis terbaru              |
| `hot+`         | `dramanova`                                                                           | —                   | Hot+ 18+ (mature)          |
| `drama18`      | `dramanova`                                                                           | —                   | Drama 18+                  |
| `romance`      | `dramanova`                                                                           | `page=1`            | Kategori romance           |

Contoh:

```bash
curl -s "https://api.anichin.bio/dramanova/hot?page=2" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"

curl -s "https://api.anichin.bio/dramabox/hotrank" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" -H "User-Agent: Mozilla/5.0"
```

---

## WebSocket API

Selain REST, semua endpoint juga bisa dipanggil via **WebSocket** untuk koneksi *long-lived* + monitoring rate limit real-time.

**URL:**

```
wss://api.anichin.bio/ws
```

### 1. Autentikasi

Setelah `onopen`, kirim payload auth:

```json
{ "type": "auth", "token": "TRIAL-ANICHIN-2026" }
```

Server membalas:

```json
{ "type": "auth", "message": "ok" }
```

### 2. Eksekusi Endpoint

```json
{
  "id": "req_1",
  "action": "execute",
  "source": "dramabox",
  "path": "episode",
  "params": { "id": "42000008521", "ep": "1" }
}
```

Server membalas dengan `id` yang sama:

```json
{
  "id": "req_1",
  "status": 200,
  "ms": 384,
  "data": { "code": 200, "videoUrl": "...", "qualityList": [ ... ] }
}
```

### 3. Notifikasi Rate Limit

Server akan push event `ratelimit` berkala:

```json
{ "type": "ratelimit", "limit": 50, "remaining": 12, "reset": 22 }
```

### Contoh JavaScript

```js
const ws = new WebSocket("wss://api.anichin.bio/ws");

ws.onopen = () => {
  ws.send(JSON.stringify({ type: "auth", token: "TRIAL-ANICHIN-2026" }));
};

ws.onmessage = (e) => {
  const msg = JSON.parse(e.data);
  if (msg.type === "ratelimit") {
    console.log(`Sisa quota: ${msg.remaining}/${msg.limit}, reset ${msg.reset}s`);
    return;
  }
  console.log("Response:", msg);
};

function execute(source, path, params) {
  ws.send(JSON.stringify({
    id: "req_" + Date.now(),
    action: "execute",
    source, path, params
  }));
}

// Setelah ws siap:
// execute("dramabox", "trending", {});
// execute("dramabox", "episode", { id: "42000008521", ep: "1" });
```

---

## Format Respons & Error

### Sukses

| Field        | Tipe         | Deskripsi                                        |
| ------------ | ------------ | ------------------------------------------------ |
| `code`       | int          | `200` jika sukses.                               |
| `msg`        | string       | `"SUCCESS"` (di sebagian endpoint).              |
| `items`      | array<obj>   | Daftar drama (untuk endpoint listing).           |
| `data`       | object       | Objek tunggal (untuk `detail`).                  |
| `hasMore`    | bool         | Indikator pagination.                            |
| `videoUrl`   | string (URL) | URL video default (untuk `episode`).             |
| `qualityList`| array<obj>   | List quality `{label, url, isDefault?}`.         |
| `subtitles`  | array<obj>   | List subtitle `{lang, url}` (jika ada).          |

### Error

Format umum:

```json
{ "error": "forbidden" }
```

| HTTP | `error` value         | Penyebab                                          |
| ---- | --------------------- | ------------------------------------------------- |
| 200  | —                     | OK                                                |
| 400  | `bad_request`         | Param `id` / `query` / `ep` hilang atau invalid.  |
| 401  | `unauthorized`        | Token tidak ada / tidak valid.                    |
| 403  | `forbidden`           | UA di-block, token expired, atau IP di-blacklist. |
| 404  | `not_found`           | `{source}` atau `{path}` tidak dikenal.           |
| 429  | `rate_limited`        | Lewat 50 req/menit. Cek field `reset`.            |
| 5xx  | `internal_error`      | Error di server / sumber upstream.                |

Saat 429 / WebSocket rate limit:

```json
{ "error": "rate_limited", "reset": 38 }
```

---

## Contoh Penggunaan (Code Snippets)

Semua contoh berikut memanggil:

```
GET https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1
```

### cURL

```bash
curl -s "https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" \
  -H "User-Agent: Mozilla/5.0"
```

### JavaScript / Node.js

```js
// Node.js 18+ (built-in fetch)
const url = "https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1";

const res = await fetch(url, {
  headers: {
    "X-API-Key": "TRIAL-ANICHIN-2026",
    "User-Agent": "Mozilla/5.0"
  }
});

if (!res.ok) throw new Error(`API error: ${res.status} ${res.statusText}`);

const data = await res.json();
console.log(JSON.stringify(data, null, 2));
```

### Python

```python
# pip install requests
import requests

url = "https://api.anichin.bio/dramabox/episode"
headers = {
    "X-API-Key": "TRIAL-ANICHIN-2026",
    "User-Agent": "Mozilla/5.0",
}
params = {"id": "42000008521", "ep": "1"}

r = requests.get(url, headers=headers, params=params, timeout=30)
r.raise_for_status()
print(r.json())
```

### Go

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func main() {
    url := "https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1"

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("X-API-Key", "TRIAL-ANICHIN-2026")
    req.Header.Set("User-Agent", "Mozilla/5.0")

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var data map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&data)
    fmt.Printf("%+v\n", data)
}
```

### PHP

```php
<?php
$url = "https://api.anichin.bio/dramabox/episode?id=42000008521&ep=1";

$ch = curl_init();
curl_setopt_array($ch, [
    CURLOPT_URL            => $url,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_FOLLOWLOCATION => true,
    CURLOPT_TIMEOUT        => 30,
    CURLOPT_HTTPHEADER     => [
        "X-API-Key: TRIAL-ANICHIN-2026",
        "User-Agent: Mozilla/5.0",
    ],
]);

$response = curl_exec($ch);
$status   = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($status !== 200) die("API error: HTTP $status");
print_r(json_decode($response, true));
```

---

## Telegram Bot Integration

Mau langsung nonton tanpa coding? Buka bot resmi:

> **[@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot)**

Bot ini menyediakan akses streaming langsung dari semua 15 sumber drama, plus fitur premium (token tanpa rate limit, akses unlock, dsb.).

---

## FAQ

**Q: Apakah API ini gratis selamanya?**
Trial token `TRIAL-ANICHIN-2026` gratis 1 hari dengan rate limit 50 req/min. Untuk produksi, hubungi via Telegram bot di atas.

**Q: Kenapa response saya `{"error":"forbidden"}` padahal token benar?**
Sertakan header `User-Agent` (mis. `Mozilla/5.0`). Request tanpa UA biasanya di-block.

**Q: Apakah saya bisa memanggil banyak sumber dengan kode yang sama?**
Ya. Skema endpoint identik — cukup ganti `{source}`. Contoh: `/dramabox/trending` ↔ `/shortmax/trending` ↔ `/melolo/trending`.

**Q: Bagaimana cara dapat ID drama untuk endpoint `detail` / `episode`?**
Ambil dari field `id` (atau `dramaId`) pada response `/trending`, `/foryou`, `/search`, atau endpoint listing lainnya.

**Q: Apakah video URL bisa di-stream langsung?**
Ya, `videoUrl` dan `qualityList[].url` adalah URL `.mp4` direct-play (HTTP range supported). Episode `locked: true` perlu unlock di sumber asalnya.

**Q: Berapa lama URL video valid?**
URL bersifat *signed* dan dapat expired. Refresh dengan memanggil ulang `/{source}/episode` saat dibutuhkan.

**Q: Ada subtitle?**
Ya, di sebagian sumber. Cek field `subtitles[]` pada response `/episode` (objek `{lang, url}` per bahasa).

---

## Catatan & Disclaimer

- **Anichin API** adalah aggregator/proxy ke sumber drama pendek pihak ketiga. Hak cipta seluruh konten video, gambar, dan metadata tetap milik penyedia masing-masing (DramaBox, ReelShort, ShortMax, dst.).
- API ini ditujukan untuk **edukasi, riset, dan pengembangan aplikasi pribadi**. Penyalahgunaan (scraping massal, redistribusi komersial konten berhak cipta, dsb.) bukan tanggung jawab Anichin.
- Skema endpoint, sumber, dan field response dapat berubah sewaktu-waktu. Selalu cek dokumentasi interaktif terbaru di [api.anichin.bio](https://api.anichin.bio).
- Untuk laporan bug, request fitur, atau upgrade ke tier non-trial → [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot).

---

<p align="center">
  Dibuat dengan ❤ untuk komunitas drama Indonesia · <a href="https://api.anichin.bio">api.anichin.bio</a>
</p>
