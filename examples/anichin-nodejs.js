// Anichin API (Dracin API) — Node.js Example
// Membutuhkan Node.js 18+ (built-in fetch)
//
// Jalankan: node examples/anichin-nodejs.js

const ANICHIN_BASE = "https://api.anichin.bio";
const ANICHIN_TOKEN = "TRIAL-ANICHIN-2026";

const headers = {
  "X-API-Key": ANICHIN_TOKEN,
  "User-Agent": "Mozilla/5.0 (anichin-nodejs-example)",
};

async function anichinGet(source, path, params = {}) {
  const url = new URL(`${ANICHIN_BASE}/${source}/${path}`);
  Object.entries(params).forEach(([k, v]) => url.searchParams.set(k, v));

  const res = await fetch(url, { headers });
  if (!res.ok) throw new Error(`Anichin API error: ${res.status} ${res.statusText}`);
  return res.json();
}

(async () => {
  const trending = await anichinGet("dramabox", "trending");
  console.log(`Anichin trending (DramaBox): ${trending.items.length} drama`);
  console.log(`#1 → ${trending.items[0].title}`);

  const firstId = trending.items[0].id;
  const ep = await anichinGet("dramabox", "episode", { id: firstId, ep: 1 });
  console.log(`Episode 1 video: ${ep.videoUrl}`);
  console.log(`Quality tersedia:`, ep.qualityList.map(q => q.label).join(", "));
})();
