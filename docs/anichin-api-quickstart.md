# Anichin API Quickstart — Dracin API dalam 5 Menit

> Panduan tercepat memanggil **Anichin API** (Dracin API / Anichin Official API) untuk pertama kali. Cocok untuk developer yang ingin langsung menarik data short drama dari 15 sumber: DramaBox, ReelShort, ShortMax, NetShort, GoodShort, DramaWave, FlickReels, FreeReels, StardustTV, iDrama, DramaNova, StarShort, DramaBite, Melolo, MoboReels.

## 1. Dapatkan Token (Gratis)

Token trial Anichin sudah disediakan publik:

```
TRIAL-ANICHIN-2026
```

- **Berlaku**: 1 hari
- **Rate limit**: 50 req/menit
- Untuk akses produksi tanpa limit → [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot)

## 2. Panggil Endpoint Pertama

```bash
curl -s "https://api.anichin.bio/dramabox/trending" \
  -H "X-API-Key: TRIAL-ANICHIN-2026" \
  -H "User-Agent: Mozilla/5.0"
```

✅ Selesai. Anda baru saja memanggil **Anichin API** dan mendapat daftar drama trending dari DramaBox.

## 3. Pola Universal

Skema URL Anichin API sama untuk semua sumber:

```
GET https://api.anichin.bio/{source}/{path}?{params}
```

Ganti `{source}` saja untuk berpindah platform — `dramabox` → `reelshort` → `melolo`.

## 4. Resep Umum

### Cari drama
```
GET /dramabox/search?query=ceo
```

### Detail + daftar episode
```
GET /dramabox/detail?id=42000008521
```

### URL video (multi-quality)
```
GET /dramabox/episode?id=42000008521&ep=1
```

Response berisi `videoUrl` (default) + `qualityList[]` (1080p / 720p / 540p).

## 5. Selanjutnya

- Baca [README.md](../README.md) untuk referensi lengkap semua endpoint Anichin.
- Lihat [contoh kode](../examples/) untuk Node.js, Python, Go, PHP.
- Coba [WebSocket API](../README.md#websocket-api) untuk koneksi long-lived.

---

**Anichin** · **Dracin API** · [api.anichin.bio](https://api.anichin.bio)
