<?php
/**
 * Anichin API (Dracin API) - PHP Example
 * Membutuhkan ekstensi php-curl
 *
 * Jalankan: php examples/anichin-php.php
 */

const ANICHIN_BASE  = "https://api.anichin.bio";
const ANICHIN_TOKEN = "TRIAL-ANICHIN-2026";

function anichin_get(string $source, string $path, array $params = []): array {
    $url = ANICHIN_BASE . "/$source/$path";
    if ($params) $url .= "?" . http_build_query($params);

    $ch = curl_init();
    curl_setopt_array($ch, [
        CURLOPT_URL            => $url,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_TIMEOUT        => 30,
        CURLOPT_HTTPHEADER     => [
            "X-API-Key: " . ANICHIN_TOKEN,
            "User-Agent: Mozilla/5.0 (anichin-php-example)",
        ],
    ]);

    $body   = curl_exec($ch);
    $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    if ($status !== 200) throw new RuntimeException("Anichin API HTTP $status");
    return json_decode($body, true);
}

$trending = anichin_get("dramabox", "trending");
echo "Anichin trending (DramaBox): " . count($trending["items"]) . " drama\n";
echo "#1 -> " . $trending["items"][0]["title"] . "\n";

$id = $trending["items"][0]["id"];
$ep = anichin_get("dramabox", "episode", ["id" => $id, "ep" => 1]);
echo "Episode 1 video: " . $ep["videoUrl"] . "\n";
echo "Quality: " . implode(", ", array_column($ep["qualityList"], "label")) . "\n";
