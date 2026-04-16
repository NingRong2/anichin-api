"""
Anichin API (Dracin API) - Python Example
pip install requests

Jalankan: python examples/anichin-python.py
"""
import requests

ANICHIN_BASE = "https://api.anichin.bio"
ANICHIN_TOKEN = "TRIAL-ANICHIN-2026"

HEADERS = {
    "X-API-Key": ANICHIN_TOKEN,
    "User-Agent": "Mozilla/5.0 (anichin-python-example)",
}


def anichin_get(source: str, path: str, **params) -> dict:
    url = f"{ANICHIN_BASE}/{source}/{path}"
    r = requests.get(url, headers=HEADERS, params=params, timeout=30)
    r.raise_for_status()
    return r.json()


if __name__ == "__main__":
    trending = anichin_get("dramabox", "trending")
    print(f"Anichin trending (DramaBox): {len(trending['items'])} drama")
    print(f"#1 -> {trending['items'][0]['title']}")

    drama_id = trending["items"][0]["id"]
    ep = anichin_get("dramabox", "episode", id=drama_id, ep=1)
    print(f"Episode 1 video: {ep['videoUrl']}")
    print("Quality:", ", ".join(q["label"] for q in ep["qualityList"]))

    search = anichin_get("shortmax", "search", query="ceo")
    print(f"\nAnichin search ShortMax 'ceo': {len(search.get('items', []))} hasil")
