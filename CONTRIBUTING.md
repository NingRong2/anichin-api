# Berkontribusi ke Anichin API Documentation

Terima kasih sudah tertarik berkontribusi pada **Anichin API** (juga dikenal sebagai **Dracin API** / **Anichin Official API**). Repository ini adalah dokumentasi terbuka untuk REST + WebSocket API resmi Anichin yang menyatukan 15 sumber drama pendek (DramaBox, ReelShort, ShortMax, NetShort, dst.).

## Cara Berkontribusi

### 1. Lapor Bug / Request Fitur
- Buka [Issue baru](../../issues/new) dengan label `bug` atau `enhancement`.
- Sertakan: endpoint Anichin yang dipanggil, request payload, response aktual, dan response yang diharapkan.
- Untuk masalah token / akses premium, hubungi [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot).

### 2. Tambah Contoh Kode
Contoh kode ditempatkan di folder `examples/` dengan format nama:

```
examples/anichin-{bahasa}.{ext}
examples/dracin-{bahasa}.{ext}
```

Contoh: `anichin-rust.rs`, `anichin-kotlin.kt`, `dracin-swift.swift`.

Setiap contoh harus:
- Menggunakan trial token `TRIAL-ANICHIN-2026`
- Menyertakan header `User-Agent: Mozilla/5.0`
- Memanggil minimal 1 endpoint (`/trending`, `/search`, `/detail`, atau `/episode`)
- Berkomentar dalam Bahasa Indonesia ATAU Bahasa Inggris

### 3. Perbaiki Dokumentasi
- Edit `README.md` atau file di `docs/`.
- Pertahankan struktur heading (H2/H3) agar TOC tidak rusak.
- Sertakan keyword **Anichin**, **Dracin**, atau **Anichin Official** secara natural — bukan keyword stuffing.

### 4. Workflow Pull Request
1. Fork repo
2. Buat branch: `git checkout -b feat/anichin-{deskripsi}`
3. Commit dengan format konvensional:
   - `docs: tambah contoh anichin-rust`
   - `fix: perbaiki rate limit example`
   - `feat: dokumentasi endpoint dramanova/romance`
4. Push & buka PR ke `main`

## Code Style

| Bahasa     | Style                          |
| ---------- | ------------------------------ |
| JavaScript | Standard / Prettier (2 spasi)  |
| Python     | PEP 8 (4 spasi, type hints)    |
| Go         | `gofmt`                        |
| PHP        | PSR-12                         |
| Markdown   | CommonMark + GFM               |

## Code of Conduct

Bersikap respek. Konten dewasa (`hot+`, `drama18`) hanya boleh disebut dalam konteks dokumentasi teknis — bukan promosi.

## Lisensi Kontribusi

Dengan berkontribusi, Anda menyetujui kontribusi Anda dirilis di bawah [MIT License](LICENSE) yang sama dengan repository Anichin API ini.

---

Pertanyaan? → [@Anichin_Premium_Bot](https://t.me/Anichin_Premium_Bot)
