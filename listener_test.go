package fstore

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ExampleStruct struct {
	IP          string `json:"ip"`
	HTTPVersion string `json:"http_version"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	TLS         struct {
		Ciphers    []string `json:"ciphers"`
		Extensions []struct {
			Name                     string   `json:"name"`
			Data                     string   `json:"data,omitempty"`
			SupportedGroups          []string `json:"supported_groups,omitempty"`
			MasterSecretData         string   `json:"master_secret_data,omitempty"`
			ExtendedMasterSecretData string   `json:"extended_master_secret_data,omitempty"`
			SignatureAlgorithms      []string `json:"signature_algorithms,omitempty"`
			StatusRequest            struct {
				CertificateStatusType   string `json:"certificate_status_type"`
				ResponderIDListLength   int    `json:"responder_id_list_length"`
				RequestExtensionsLength int    `json:"request_extensions_length"`
			} `json:"status_request,omitempty"`
			Versions   []string `json:"versions,omitempty"`
			Algorithms []string `json:"algorithms,omitempty"`
			ServerName string   `json:"server_name,omitempty"`
			SharedKeys []struct {
				TLSGREASE0Xfafa     string `json:"TLS_GREASE (0xfafa),omitempty"`
				X25519Kyber76825497 string `json:"X25519Kyber768 (25497),omitempty"`
				X2551929            string `json:"X25519 (29),omitempty"`
			} `json:"shared_keys,omitempty"`
			EllipticCurvesPointFormats []string `json:"elliptic_curves_point_formats,omitempty"`
			Protocols                  []string `json:"protocols,omitempty"`
			PSKKeyExchangeMode         string   `json:"PSK_Key_Exchange_Mode,omitempty"`
		} `json:"extensions"`
		TLSVersionRecord     string `json:"tls_version_record"`
		TLSVersionNegotiated string `json:"tls_version_negotiated"`
		Ja3                  string `json:"ja3"`
		Ja3Hash              string `json:"ja3_hash"`
		Peetprint            string `json:"peetprint"`
		PeetprintHash        string `json:"peetprint_hash"`
		ClientRandom         string `json:"client_random"`
		SessionID            string `json:"session_id"`
	} `json:"tls"`
	HTTP2 struct {
		AkamaiFingerprint     string `json:"akamai_fingerprint"`
		AkamaiFingerprintHash string `json:"akamai_fingerprint_hash"`
		SentFrames            []struct {
			FrameType string   `json:"frame_type"`
			Length    int      `json:"length"`
			Settings  []string `json:"settings,omitempty"`
			Increment int      `json:"increment,omitempty"`
			StreamID  int      `json:"stream_id,omitempty"`
			Headers   []string `json:"headers,omitempty"`
			Flags     []string `json:"flags,omitempty"`
			Priority  struct {
				Weight    int `json:"weight"`
				DependsOn int `json:"depends_on"`
				Exclusive int `json:"exclusive"`
			} `json:"priority,omitempty"`
		} `json:"sent_frames"`
	} `json:"http2"`
}

var fp = []byte(`{
  "ip": "1.1.1.1:59418",
  "http_version": "h2",
  "method": "GET",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
  "tls": {
    "ciphers": [
      "TLS_GREASE (0x8A8A)",
      "TLS_AES_128_GCM_SHA256",
      "TLS_AES_256_GCM_SHA384",
      "TLS_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
      "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
      "TLS_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_RSA_WITH_AES_128_CBC_SHA",
      "TLS_RSA_WITH_AES_256_CBC_SHA"
    ],
    "extensions": [
      {
        "name": "TLS_GREASE (0x9a9a)"
      },
      {
        "name": "signed_certificate_timestamp (18)"
      },
      {
        "name": "session_ticket (35)",
        "data": ""
      },
      {
        "name": "supported_groups (10)",
        "supported_groups": [
          "TLS_GREASE (0xfafa)",
          "X25519Kyber768 (25497)",
          "X25519 (29)",
          "P-256 (23)",
          "P-384 (24)"
        ]
      },
      {
        "name": "extended_master_secret (23)",
        "master_secret_data": "",
        "extended_master_secret_data": ""
      },
      {
        "name": "signature_algorithms (13)",
        "signature_algorithms": [
          "ecdsa_secp256r1_sha256",
          "rsa_pss_rsae_sha256",
          "rsa_pkcs1_sha256",
          "ecdsa_secp384r1_sha384",
          "rsa_pss_rsae_sha384",
          "rsa_pkcs1_sha384",
          "rsa_pss_rsae_sha512",
          "rsa_pkcs1_sha512"
        ]
      },
      {
        "name": "status_request (5)",
        "status_request": {
          "certificate_status_type": "OSCP (1)",
          "responder_id_list_length": 0,
          "request_extensions_length": 0
        }
      },
      {
        "name": "supported_versions (43)",
        "versions": [
          "TLS_GREASE (0xdada)",
          "TLS 1.3",
          "TLS 1.2"
        ]
      },
      {
        "name": "compress_certificate (27)",
        "algorithms": [
          "brotli (2)"
        ]
      },
      {
        "name": "server_name (0)",
        "server_name": "tls.peet.ws"
      },
      {
        "name": "extensionRenegotiationInfo (boringssl) (65281)",
        "data": "00"
      },
      {
        "name": "key_share (51)",
        "shared_keys": [
          {
            "TLS_GREASE (0xfafa)": "00"
          },
          {
            "X25519Kyber768 (25497)": "7b363296460eb1c02706b128953e6d75bc36d71d499de8623379ba831be44f39e31182fb5c724af745adf7012bbc0b8ef9ba6f66124f775b1b5346c2071cdd5a6d7cd6c3bff84291ab87fd735aae6403a3d109b5e18ca7e0b98833259c442a730ba31e6041a0f309fb4b5dcad1b6700a52c6c4421a0616adf43cd0639c4dd8b62f8945c9278cad96c60fb140aa0759e32a1365d06ada077188b33922e038eac0caea79cebc8c921106a33c6c8febc70613a3036b666a365a4bf83b0bdcd142387aa375a49e1352a4eb6656c9c53f5f6812789728067a6e355ccf02aa1b6cb16500024c8a35688c9170b074aeb5e0a70452ca83c94b428b75bec6245ad971d1c854ac2bc041f95293e45e0687ac7c152815c0ab34d79e569b35530736ebd90a6237746426be3a811be448243bdbb0f628c4022514bb04bd56d5113c10486843b0c574c414a176c6c62957032dda84c0dd9ac6c8b9251e2c0caf563dc606b726b135d5926357a3ca93327564276efe8970c648c62439aa2ff03a4f65a1cd8273d73b899967b705585703c47c3b7c80528c1eadb4254222c5bea95020fc114361bb7d65a1f768123e977bc51b82ee4bce2e54bee8b44ab1f929de1171ff950de2048280c8bffdd082e850861bb50fb580ad2b6276e84c529dc34402c761a53811ee206c68d66c7213a88eba6d39a51c3d8a0a05695a422931cc9a27bfbc3052694d29766b587862e8e0cbc126c22f4081cea9700ee989f8ab835d03b858962507f0aaa7734d56c2a867205c3c448e74a95eac7554d0f7c8853044ea73ca369aad310c3a4f06318837515cc5b90f29446dfc05fbac2641f52e6829c22b348c75107ee19a549140b3d7f32d898a1114f58ddc119adf642b31b2b0fdf13831eb4465b4c0246679111761e0aa8f1e135f91acb2eff9b37f6c459297be885b7f8e16c0d4364d33a9c51f9a8aa1b4a30127b97a96a6dc9257902a0d51a32fbb298ca7f0362d54c52959924a758b3110404fd14d28080711a7714a109ccb178dd68424651a75578499b21a7226c203a129281e3300ca9459909a0b5163bc8a39408cda65927365c4d6c7ed8a19ae37b7d7f02e900c02b9e7865cd621d681c4ca438a3d8863eef4cce9126ee95096f4d7833ffc8c895cca79c105d5d65bf26353448b40aeb53da0226fb9f063a66389ae63962f729da684403a67cc096545c7318af872aa713cc7a25260671a360bc8226f287c70e2a174e086fda954152671a996a49ee60d8609a2485618a1d7a9f7b6a2004dc16dd479f79121dbb3c53586aa8e075d8f76be89806b23e28ccc1cb285862bd625b9b0d053afe0b450b9c554589654853236f3bea59645234ccb289025f21603f082640420cfc7551afa653e94f80d45e8b037e8cda8d07074993f4f3068cbb21e464a2745c52dd226610080083d9467f3b831c379a40700147dd340d88b1875e801ea33770c47b6dea5795c3356593209073c6a9e5cc9696b227f013abc12642a91020dfb41dc761dac56368261bf5ddc6dfb96a19d43bce827b7ce41a90157afc7db49062cc034e88ce14b897ccb8c020c867e0550fc498659e84f64bc4772e43f5ee5bfc27bc5d245a66854745fbb460d61b8bc362bc876870bd2ce83009112a49e64739b8c0700b72acbd8a202862722e54051a4d37313b0c439b0f52eb97e77500d83e51ed3adcad746beec3863979b"
          },
          {
            "X25519 (29)": "e43687094584db53755b0a9abc84a7433579aa5598db91ae57950d158139f01b"
          }
        ]
      },
      {
        "name": "ec_point_formats (11)",
        "elliptic_curves_point_formats": [
          "0x00"
        ]
      },
      {
        "name": "application_layer_protocol_negotiation (16)",
        "protocols": [
          "h2",
          "http/1.1"
        ]
      },
      {
        "name": "application_settings (17513)",
        "protocols": [
          "h2"
        ]
      },
      {
        "name": "psk_key_exchange_modes (45)",
        "PSK_Key_Exchange_Mode": "PSK with (EC)DHE key establishment (psk_dhe_ke) (1)"
      },
      {
        "name": "TLS_GREASE (0x7a7a)"
      },
      {
        "name": "pre_shared_key (41)",
        "data": "007700714ba602a57be4dc4461eba9dbc6d54f5405897aeff7339546e52ee92bb71b55719247cac9b198e3af2eaeaa5ffbbc84c34404ea2ae9ba7a0cd839b770d8b537691d74416ca5d868bb61f367beb68c513818ca8d4db3e4c06b9362cc5327caffec6b1fc0323a5ef70d2e4536dbc638d5d8d61a26127900212001247a71878613f7d8ffc455d0b8a356c3a2ddd1c297a86d3f20e39d58d561f4"
      }
    ],
    "tls_version_record": "771",
    "tls_version_negotiated": "772",
    "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,18-35-10-23-13-5-43-27-0-65281-51-11-16-17513-45-41,25497-29-23-24,0",
    "ja3_hash": "9f6db5c1bf0924780372cca95bb7f8ba",
    "peetprint": "GREASE-772-771|2-1.1|GREASE-25497-29-23-24|1027-2052-1025-1283-2053-1281-2054-1537|1|2|GREASE-4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53|0-10-11-13-16-17513-18-23-27-35-41-43-45-5-51-65281-GREASE-GREASE",
    "peetprint_hash": "0cd24525a36f772d1f68f62ea0eb24ce",
    "client_random": "c62cb50847e207bb8c0f391ac9bec7fa47001aa7cc274f9cf1a5764ff1bc1827",
    "session_id": "eb6ddd9d4765afb1c5c780e2087efbce13f1f84b0da7e7b3107230d6216f752d"
  },
  "http2": {
    "akamai_fingerprint": "1:65536,2:0,4:6291456,6:262144|15663105|0|m,a,s,p",
    "akamai_fingerprint_hash": "90224459f8bf70b7d0a8797eb916dbc9",
    "sent_frames": [
      {
        "frame_type": "SETTINGS",
        "length": 24,
        "settings": [
          "HEADER_TABLE_SIZE = 65536",
          "ENABLE_PUSH = 0",
          "INITIAL_WINDOW_SIZE = 6291456",
          "MAX_HEADER_LIST_SIZE = 262144"
        ]
      },
      {
        "frame_type": "WINDOW_UPDATE",
        "length": 4,
        "increment": 15663105
      },
      {
        "frame_type": "HEADERS",
        "stream_id": 1,
        "length": 568,
        "headers": [
          ":method: GET",
          ":authority: tls.peet.ws",
          ":scheme: https",
          ":path: /api/all",
          "sec-ch-ua: \\\"Google Chrome\\\";v=\\\"117\\\", \\\"Not;A=Brand\\\";v=\\\"8\\\", \\\"Chromium\\\";v=\\\"117\\",
          "sec-ch-ua-mobile: ?0",
          "sec-ch-ua-platform: \\\"macOS\\",
          "upgrade-insecure-requests: 1",
          "dnt: 1",
          "user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
          "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
          "sec-fetch-site: same-origin",
          "sec-fetch-mode: navigate",
          "sec-fetch-user: ?1",
          "sec-fetch-dest: document",
          "referer: https://tls.peet.ws/",
          "accept-encoding: gzip, deflate, br",
          "accept-language: en-DE,en-US;q=0.9,en;q=0.8",
          "cookie: cf_clearance=MUoJKskype.U8euK.Lj3E9nLBa3Vfyc9eQNl7vpQKsk-1696544599-0-1-957868af.5df2091a.69ec06d1-160.2.1696544599"
        ],
        "flags": [
          "EndStream (0x1)",
          "EndHeaders (0x4)",
          "Priority (0x20)"
        ],
        "priority": {
          "weight": 256,
          "depends_on": 0,
          "exclusive": 1
        }
      }
    ]
  }
}`)
var fp2 = []byte(`{
	"appCodeName": "Mozilla",
	"appName": "Netscape",
	"appVersion": "5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	"cookieEnabled": true,
	"deviceMemory": 8,
	"doNotTrack": "1",
	"hardwareConcurrency": 8,
	"language": "en-DE",
	"languages": [
		"en-DE",
		"en-US",
		"en"
	],
	"maxTouchPoints": 0,
	"pdfViewerEnabled": true,
	"platform": "MacIntel",
	"product": "Gecko",
	"productSub": "20030107",
	"userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	"vendor": "Google Inc.",
	"vendorSub": "",
	"webdiver": false,
	"devicePixelRatio": 2,
	"innerWidth": 1420,
	"innerHeight": 770,
	"outerWidth": 855,
	"outerHeight": 855,
	"screen": {
		"availWidth": 1440,
		"availHeight": 875,
		"width": 1440,
		"height": 900,
		"colorDepth": 30,
		"pixelDepth": 30,
		"availLeft": 0,
		"availTop": 25,
		"orientation": {
			"angle": 0,
			"type": "landscape-primary"
		},
		"isExtended": false
	},
	"connection": {
		"effectiveType": "4g",
		"rtt": 100,
		"downlink": 1.95,
		"saveData": false
	},
	"plugins": [
		{
			"0": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"1": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"application/pdf": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"text/pdf": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"name": "PDF Viewer",
			"filename": "internal-pdf-viewer",
			"description": "Portable Document Format",
			"length": 2
		},
		{
			"0": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"1": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"application/pdf": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"text/pdf": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"name": "Chrome PDF Viewer",
			"filename": "internal-pdf-viewer",
			"description": "Portable Document Format",
			"length": 2
		},
		{
			"0": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"1": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"application/pdf": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"text/pdf": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"name": "Chromium PDF Viewer",
			"filename": "internal-pdf-viewer",
			"description": "Portable Document Format",
			"length": 2
		},
		{
			"0": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"1": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"application/pdf": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"text/pdf": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"name": "Microsoft Edge PDF Viewer",
			"filename": "internal-pdf-viewer",
			"description": "Portable Document Format",
			"length": 2
		},
		{
			"0": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"1": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"application/pdf": {
				"type": "application/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"text/pdf": {
				"type": "text/pdf",
				"suffixes": "pdf",
				"description": "Portable Document Format"
			},
			"name": "WebKit built-in PDF",
			"filename": "internal-pdf-viewer",
			"description": "Portable Document Format",
			"length": 2
		}
	],
	"userActivation": {
		"hasBeenActive": false,
		"isActive": false
	},
	"chrome.app": {
		"isInstalled": false,
		"InstallState": {
			"DISABLED": "disabled",
			"INSTALLED": "installed",
			"NOT_INSTALLED": "not_installed"
		},
		"RunningState": {
			"CANNOT_RUN": "cannot_run",
			"READY_TO_RUN": "ready_to_run",
			"RUNNING": "running"
		}
	},
	"wow64": false,
	"HighEntropyValues": {
		"architecture": "arm",
		"bitness": "64",
		"brands": [
			{
				"brand": "Google Chrome",
				"version": "117"
			},
			{
				"brand": "Not;A=Brand",
				"version": "8"
			},
			{
				"brand": "Chromium",
				"version": "117"
			}
		],
		"fullVersionList": [
			{
				"brand": "Google Chrome",
				"version": "117.0.5938.92"
			},
			{
				"brand": "Not;A=Brand",
				"version": "8.0.0.0"
			},
			{
				"brand": "Chromium",
				"version": "117.0.5938.92"
			}
		],
		"mobile": false,
		"model": "",
		"platform": "macOS",
		"platformVersion": "14.0.0",
		"uaFullVersion": "117.0.5938.92"
	},
	"darkmode": true,
	"availabeFonts": [
		"AVENIR",
		"Academy Engraved LET",
		"Al Bayan",
		"Al Nile",
		"Al Tarikh",
		"American Typewriter",
		"Andale Mono",
		"Apple Braille",
		"Apple Chancery",
		"Apple Color Emoji",
		"Apple SD Gothic Neo",
		"Apple Symbols",
		"AppleGothic",
		"AppleMyungjo",
		"Arial",
		"Arial Black",
		"Arial Hebrew",
		"Arial Hebrew Scholar",
		"Arial Narrow",
		"Arial Rounded MT Bold",
		"Arial Unicode MS",
		"Arima Madurai",
		"Avenir",
		"Avenir Black",
		"Avenir Black Oblique",
		"Avenir Book",
		"Avenir Heavy",
		"Avenir Light",
		"Avenir Medium",
		"Avenir Next",
		"Avenir Next Condensed Demi Bold",
		"Avenir Next Condensed Heavy",
		"Avenir Next Condensed Medium",
		"Avenir Next Condensed Ultra Light",
		"Avenir Next Demi Bold",
		"Avenir Next Heavy",
		"Avenir Next Medium",
		"Avenir Next Ultra Light",
		"Ayuthaya",
		"Baghdad",
		"Bai Jamjuree",
		"Baloo Bhaijaan",
		"Bangla MN",
		"Bangla Sangam MN",
		"Baskerville",
		"Beirut",
		"Big Caslon",
		"Bodoni 72",
		"Bodoni 72 Oldstyle",
		"Bodoni 72 Smallcaps",
		"Bodoni Ornaments",
		"Bradley Hand",
		"Brush Script MT",
		"Chakra Petch",
		"Chalkboard",
		"Chalkboard SE",
		"Chalkduster",
		"Charm",
		"Charmonman",
		"Charter",
		"Charter Black",
		"Cochin",
		"Comic Sans MS",
		"Copperplate",
		"Corsiva Hebrew",
		"Courier",
		"Courier New",
		"DIN Alternate",
		"DIN Condensed",
		"Damascus",
		"DecoType Naskh",
		"Devanagari MT",
		"Devanagari Sangam MN",
		"Didot",
		"Diwan Kufi",
		"Diwan Thuluth",
		"Euphemia UCAS",
		"Fahkwang",
		"Farah",
		"Farisi",
		"Futura",
		"GB18030 Bitmap",
		"Geeza Pro",
		"Geneva",
		"Georgia",
		"Gill Sans",
		"Gujarati MT",
		"Gujarati Sangam MN",
		"Gurmukhi MN",
		"Gurmukhi MT",
		"Gurmukhi Sangam MN",
		"Heiti SC",
		"Heiti TC",
		"Helvetica",
		"Helvetica Neue",
		"Herculanum",
		"Hiragino Kaku Gothic Pro W3",
		"Hiragino Kaku Gothic Pro W6",
		"Hiragino Kaku Gothic ProN",
		"Hiragino Kaku Gothic ProN W3",
		"Hiragino Kaku Gothic ProN W6",
		"Hiragino Kaku Gothic Std W8",
		"Hiragino Kaku Gothic StdN W8",
		"Hiragino Maru Gothic Pro W4",
		"Hiragino Maru Gothic ProN",
		"Hiragino Maru Gothic ProN W4",
		"Hiragino Mincho Pro W3",
		"Hiragino Mincho Pro W6",
		"Hiragino Mincho ProN",
		"Hiragino Mincho ProN W3",
		"Hiragino Mincho ProN W6",
		"Hiragino Sans",
		"Hiragino Sans CNS W3",
		"Hiragino Sans CNS W6",
		"Hiragino Sans GB",
		"Hiragino Sans GB W3",
		"Hiragino Sans GB W6",
		"Hiragino Sans W0",
		"Hiragino Sans W1",
		"Hiragino Sans W2",
		"Hiragino Sans W3",
		"Hiragino Sans W4",
		"Hiragino Sans W5",
		"Hiragino Sans W6",
		"Hiragino Sans W7",
		"Hiragino Sans W8",
		"Hiragino Sans W9",
		"Hoefler Text",
		"Hoefler Text Ornaments",
		"ITF Devanagari",
		"ITF Devanagari Marathi",
		"Impact",
		"InaiMathi",
		"Iowan Old Style Black",
		"K2D",
		"Kailasa",
		"Kannada MN",
		"Kannada Sangam MN",
		"Kavivanar",
		"Kefa",
		"Khmer MN",
		"Khmer Sangam MN",
		"Klee Demibold",
		"Klee Medium",
		"KoHo",
		"Kodchasan",
		"Kohinoor Bangla",
		"Kohinoor Devanagari",
		"Kohinoor Telugu",
		"Kokonor",
		"Krub",
		"Krungthep",
		"KufiStandardGK",
		"LUCIDA GRANDE",
		"Lao MN",
		"Lao Sangam MN",
		"LiHei Pro",
		"LiSong Pro",
		"Lucida Grande",
		"Luminari",
		"Malayalam MN",
		"Malayalam Sangam MN",
		"Mali",
		"Marion",
		"Marker Felt",
		"Menlo",
		"Microsoft Sans Serif",
		"Mishafi",
		"Mishafi Gold",
		"Modak",
		"Monaco",
		"Mshtakan",
		"Mukta",
		"Mukta Mahee",
		"Mukta Malar",
		"Mukta Vaani",
		"Muna",
		"Myanmar MN",
		"Myanmar Sangam MN",
		"Nadeem",
		"Nanum Brush Script",
		"Nanum Gothic",
		"Nanum Myeongjo",
		"Nanum Pen Script",
		"NanumGothic",
		"NanumGothic ExtraBold",
		"NanumMyeongjo",
		"New Peninim MT",
		"Niramit",
		"Noteworthy",
		"Noto Nastaliq Urdu",
		"Noto Sans Adlam",
		"Noto Sans Armenian",
		"Noto Sans Avestan",
		"Noto Sans Bamum",
		"Noto Sans Batak",
		"Noto Sans Brahmi",
		"Noto Sans Buginese",
		"Noto Sans Buhid",
		"Noto Sans Canadian Aboriginal",
		"Noto Sans Carian",
		"Noto Sans Chakma",
		"Noto Sans Cham",
		"Noto Sans Coptic",
		"Noto Sans Cuneiform",
		"Noto Sans Cypriot",
		"Noto Sans Egyptian Hieroglyphs",
		"Noto Sans Glagolitic",
		"Noto Sans Gothic",
		"Noto Sans Hanunoo",
		"Noto Sans Imperial Aramaic",
		"Noto Sans Inscriptional Pahlavi",
		"Noto Sans Inscriptional Parthian",
		"Noto Sans Javanese",
		"Noto Sans Kaithi",
		"Noto Sans Kannada",
		"Noto Sans Kayah Li",
		"Noto Sans Kharoshthi",
		"Noto Sans Lepcha",
		"Noto Sans Limbu",
		"Noto Sans Linear B",
		"Noto Sans Lisu",
		"Noto Sans Lycian",
		"Noto Sans Lydian",
		"Noto Sans Mandaic",
		"Noto Sans Meetei Mayek",
		"Noto Sans Mongolian",
		"Noto Sans Myanmar",
		"Noto Sans NKo",
		"Noto Sans New Tai Lue",
		"Noto Sans Ol Chiki",
		"Noto Sans Old Italic",
		"Noto Sans Old Persian",
		"Noto Sans Old South Arabian",
		"Noto Sans Old Turkic",
		"Noto Sans Oriya",
		"Noto Sans Osage",
		"Noto Sans Osmanya",
		"Noto Sans Phoenician",
		"Noto Sans Rejang",
		"Noto Sans Samaritan",
		"Noto Sans Saurashtra",
		"Noto Sans Sundanese",
		"Noto Sans Syloti Nagri",
		"Noto Sans Tagalog",
		"Noto Sans Tagbanwa",
		"Noto Sans Tai Le",
		"Noto Sans Tai Tham",
		"Noto Sans Tai Viet",
		"Noto Sans Thaana",
		"Noto Sans Tifinagh",
		"Noto Sans Ugaritic",
		"Noto Sans Vai",
		"Noto Sans Yi",
		"Noto Serif Kannada",
		"Noto Serif Myanmar",
		"OPTIMA",
		"OSAKA",
		"Optima",
		"Oriya MN",
		"Oriya Sangam MN",
		"Osaka",
		"Osaka-Mono",
		"PT Mono",
		"PT Sans",
		"PT Sans Caption",
		"PT Sans Narrow",
		"PT Serif",
		"PT Serif Caption",
		"Palatino",
		"Papyrus",
		"Party LET",
		"Phosphate",
		"PingFang HK",
		"PingFang SC",
		"PingFang TC",
		"Plantagenet Cherokee",
		"Raanana",
		"Rockwell",
		"STFangsong",
		"STHeiti",
		"STIX Two Math",
		"STIX Two Text",
		"STIXGeneral",
		"STIXIntegralsD",
		"STIXIntegralsSm",
		"STIXIntegralsUp",
		"STIXIntegralsUpD",
		"STIXIntegralsUpSm",
		"STIXNonUnicode",
		"STIXSizeFiveSym",
		"STIXSizeFourSym",
		"STIXSizeOneSym",
		"STIXSizeThreeSym",
		"STIXSizeTwoSym",
		"STIXVariants",
		"STKaiti",
		"STSong",
		"STXihei",
		"Sana",
		"Sarabun",
		"Sathu",
		"Savoye LET",
		"Seravek",
		"Seravek ExtraLight",
		"Seravek Light",
		"Seravek Medium",
		"Shree Devanagari 714",
		"SignPainter",
		"SignPainter-HouseScript",
		"Silom",
		"Sinhala MN",
		"Sinhala Sangam MN",
		"Skia",
		"Snell Roundhand",
		"Songti SC",
		"Songti TC",
		"Srisakdi",
		"Sukhumvit Set",
		"Symbol",
		"Tahoma",
		"Tamil MN",
		"Tamil Sangam MN",
		"Telugu MN",
		"Telugu Sangam MN",
		"Thonburi",
		"Times",
		"Times New Roman",
		"Toppan Bunkyu Gothic Regular",
		"Toppan Bunkyu Midashi Gothic Extrabold",
		"Toppan Bunkyu Midashi Mincho Extrabold",
		"Toppan Bunkyu Mincho Regular",
		"Trattatello",
		"Trebuchet MS",
		"Tsukushi A Round Gothic Bold",
		"Tsukushi A Round Gothic Regular",
		"Tsukushi B Round Gothic Bold",
		"Tsukushi B Round Gothic Regular",
		"Verdana",
		"Waseem",
		"Webdings",
		"Wingdings",
		"Wingdings 2",
		"Wingdings 3",
		"YuGothic Bold",
		"YuGothic Medium",
		"YuMincho +36p Kana Demibold",
		"YuMincho +36p Kana Extrabold",
		"YuMincho +36p Kana Medium",
		"YuMincho Demibold",
		"YuMincho Extrabold",
		"YuMincho Medium",
		"Zapf Dingbats",
		"Zapfino"
	],
	"stack_native": [
		11034,
		9195,
		40
	],
	"timing_native": 0.09999999403953552,
	"permissions": {
		"background-sync": "granted",
		"clipboard-read": "prompt",
		"clipboard-write": "granted",
		"geolocation": "prompt",
		"microphone": "prompt",
		"camera": "prompt",
		"notifications": "prompt",
		"payment-handler": "granted",
		"accelerometer": "granted",
		"gyroscope": "granted",
		"magnetometer": "granted",
		"storage-access": "granted",
		"persistent-storage": "prompt",
		"midi": "granted",
		"background-fetch": "granted",
		"screen-wake-lock": "granted",
		"display-capture": "prompt"
	},
	"navigator": [
		"vendorSub",
		"productSub",
		"vendor",
		"maxTouchPoints",
		"scheduling",
		"userActivation",
		"doNotTrack",
		"geolocation",
		"connection",
		"plugins",
		"mimeTypes",
		"pdfViewerEnabled",
		"webkitTemporaryStorage",
		"webkitPersistentStorage",
		"hardwareConcurrency",
		"cookieEnabled",
		"appCodeName",
		"appName",
		"appVersion",
		"platform",
		"product",
		"userAgent",
		"language",
		"languages",
		"onLine",
		"webdriver",
		"getGamepads",
		"javaEnabled",
		"sendBeacon",
		"vibrate",
		"bluetooth",
		"clipboard",
		"credentials",
		"keyboard",
		"managed",
		"mediaDevices",
		"storage",
		"serviceWorker",
		"virtualKeyboard",
		"wakeLock",
		"deviceMemory",
		"ink",
		"hid",
		"locks",
		"gpu",
		"mediaCapabilities",
		"mediaSession",
		"permissions",
		"presentation",
		"usb",
		"xr",
		"serial",
		"windowControlsOverlay",
		"userAgentData",
		"adAuctionComponents",
		"runAdAuction",
		"canLoadAdAuctionFencedFrame",
		"clearAppBadge",
		"getBattery",
		"getUserMedia",
		"requestMIDIAccess",
		"requestMediaKeySystemAccess",
		"setAppBadge",
		"webkitGetUserMedia",
		"deprecatedReplaceInURN",
		"deprecatedURNToURL",
		"getInstalledRelatedApps",
		"joinAdInterestGroup",
		"leaveAdInterestGroup",
		"updateAdInterestGroups",
		"registerProtocolHandler",
		"unregisterProtocolHandler"
	],
	"window": [
		"Object",
		"Function",
		"Array",
		"Number",
		"parseFloat",
		"parseInt",
		"Infinity",
		"NaN",
		"undefined",
		"Boolean",
		"String",
		"Symbol",
		"Date",
		"Promise",
		"RegExp",
		"Error",
		"AggregateError",
		"EvalError",
		"RangeError",
		"ReferenceError",
		"SyntaxError",
		"TypeError",
		"URIError",
		"globalThis",
		"JSON",
		"Math",
		"Intl",
		"ArrayBuffer",
		"Atomics",
		"Uint8Array",
		"Int8Array",
		"Uint16Array",
		"Int16Array",
		"Uint32Array",
		"Int32Array",
		"Float32Array",
		"Float64Array",
		"Uint8ClampedArray",
		"BigUint64Array",
		"BigInt64Array",
		"DataView",
		"Map",
		"BigInt",
		"Set",
		"WeakMap",
		"WeakSet",
		"Proxy",
		"Reflect",
		"FinalizationRegistry",
		"WeakRef",
		"decodeURI",
		"decodeURIComponent",
		"encodeURI",
		"encodeURIComponent",
		"escape",
		"unescape",
		"eval",
		"isFinite",
		"isNaN",
		"console",
		"Option",
		"Image",
		"Audio",
		"webkitURL",
		"webkitRTCPeerConnection",
		"webkitMediaStream",
		"WebKitMutationObserver",
		"WebKitCSSMatrix",
		"XSLTProcessor",
		"XPathResult",
		"XPathExpression",
		"XPathEvaluator",
		"XMLSerializer",
		"XMLHttpRequestUpload",
		"XMLHttpRequestEventTarget",
		"XMLHttpRequest",
		"XMLDocument",
		"WritableStreamDefaultWriter",
		"WritableStreamDefaultController",
		"WritableStream",
		"Worker",
		"Window",
		"WheelEvent",
		"WebSocket",
		"WebGLVertexArrayObject",
		"WebGLUniformLocation",
		"WebGLTransformFeedback",
		"WebGLTexture",
		"WebGLSync",
		"WebGLShaderPrecisionFormat",
		"WebGLShader",
		"WebGLSampler",
		"WebGLRenderingContext",
		"WebGLRenderbuffer",
		"WebGLQuery",
		"WebGLProgram",
		"WebGLFramebuffer",
		"WebGLContextEvent",
		"WebGLBuffer",
		"WebGLActiveInfo",
		"WebGL2RenderingContext",
		"WaveShaperNode",
		"VisualViewport",
		"VirtualKeyboardGeometryChangeEvent",
		"ValidityState",
		"VTTCue",
		"UserActivation",
		"URLSearchParams",
		"URLPattern",
		"URL",
		"UIEvent",
		"TrustedTypePolicyFactory",
		"TrustedTypePolicy",
		"TrustedScriptURL",
		"TrustedScript",
		"TrustedHTML",
		"TreeWalker",
		"TransitionEvent",
		"TransformStreamDefaultController",
		"TransformStream",
		"TrackEvent",
		"TouchList",
		"TouchEvent",
		"Touch",
		"TimeRanges",
		"TextTrackList",
		"TextTrackCueList",
		"TextTrackCue",
		"TextTrack",
		"TextMetrics",
		"TextEvent",
		"TextEncoderStream",
		"TextEncoder",
		"TextDecoderStream",
		"TextDecoder",
		"Text",
		"TaskSignal",
		"TaskPriorityChangeEvent",
		"TaskController",
		"TaskAttributionTiming",
		"SyncManager",
		"SubmitEvent",
		"StyleSheetList",
		"StyleSheet",
		"StylePropertyMapReadOnly",
		"StylePropertyMap",
		"StorageEvent",
		"Storage",
		"StereoPannerNode",
		"StaticRange",
		"SourceBufferList",
		"SourceBuffer",
		"ShadowRoot",
		"Selection",
		"SecurityPolicyViolationEvent",
		"ScriptProcessorNode",
		"ScreenOrientation",
		"Screen",
		"Scheduling",
		"Scheduler",
		"SVGViewElement",
		"SVGUseElement",
		"SVGUnitTypes",
		"SVGTransformList",
		"SVGTransform",
		"SVGTitleElement",
		"SVGTextPositioningElement",
		"SVGTextPathElement",
		"SVGTextElement",
		"SVGTextContentElement",
		"SVGTSpanElement",
		"SVGSymbolElement",
		"SVGSwitchElement",
		"SVGStyleElement",
		"SVGStringList",
		"SVGStopElement",
		"SVGSetElement",
		"SVGScriptElement",
		"SVGSVGElement",
		"SVGRectElement",
		"SVGRect",
		"SVGRadialGradientElement",
		"SVGPreserveAspectRatio",
		"SVGPolylineElement",
		"SVGPolygonElement",
		"SVGPointList",
		"SVGPoint",
		"SVGPatternElement",
		"SVGPathElement",
		"SVGNumberList",
		"SVGNumber",
		"SVGMetadataElement",
		"SVGMatrix",
		"SVGMaskElement",
		"SVGMarkerElement",
		"SVGMPathElement",
		"SVGLinearGradientElement",
		"SVGLineElement",
		"SVGLengthList",
		"SVGLength",
		"SVGImageElement",
		"SVGGraphicsElement",
		"SVGGradientElement",
		"SVGGeometryElement",
		"SVGGElement",
		"SVGForeignObjectElement",
		"SVGFilterElement",
		"SVGFETurbulenceElement",
		"SVGFETileElement",
		"SVGFESpotLightElement",
		"SVGFESpecularLightingElement",
		"SVGFEPointLightElement",
		"SVGFEOffsetElement",
		"SVGFEMorphologyElement",
		"SVGFEMergeNodeElement",
		"SVGFEMergeElement",
		"SVGFEImageElement",
		"SVGFEGaussianBlurElement",
		"SVGFEFuncRElement",
		"SVGFEFuncGElement",
		"SVGFEFuncBElement",
		"SVGFEFuncAElement",
		"SVGFEFloodElement",
		"SVGFEDropShadowElement",
		"SVGFEDistantLightElement",
		"SVGFEDisplacementMapElement",
		"SVGFEDiffuseLightingElement",
		"SVGFEConvolveMatrixElement",
		"SVGFECompositeElement",
		"SVGFEComponentTransferElement",
		"SVGFEColorMatrixElement",
		"SVGFEBlendElement",
		"SVGEllipseElement",
		"SVGElement",
		"SVGDescElement",
		"SVGDefsElement",
		"SVGComponentTransferFunctionElement",
		"SVGClipPathElement",
		"SVGCircleElement",
		"SVGAnimationElement",
		"SVGAnimatedTransformList",
		"SVGAnimatedString",
		"SVGAnimatedRect",
		"SVGAnimatedPreserveAspectRatio",
		"SVGAnimatedNumberList",
		"SVGAnimatedNumber",
		"SVGAnimatedLengthList",
		"SVGAnimatedLength",
		"SVGAnimatedInteger",
		"SVGAnimatedEnumeration",
		"SVGAnimatedBoolean",
		"SVGAnimatedAngle",
		"SVGAnimateTransformElement",
		"SVGAnimateMotionElement",
		"SVGAnimateElement",
		"SVGAngle",
		"SVGAElement",
		"Response",
		"ResizeObserverSize",
		"ResizeObserverEntry",
		"ResizeObserver",
		"Request",
		"ReportingObserver",
		"ReadableStreamDefaultReader",
		"ReadableStreamDefaultController",
		"ReadableStreamBYOBRequest",
		"ReadableStreamBYOBReader",
		"ReadableStream",
		"ReadableByteStreamController",
		"Range",
		"RadioNodeList",
		"RTCTrackEvent",
		"RTCStatsReport",
		"RTCSessionDescription",
		"RTCSctpTransport",
		"RTCRtpTransceiver",
		"RTCRtpSender",
		"RTCRtpReceiver",
		"RTCPeerConnectionIceEvent",
		"RTCPeerConnectionIceErrorEvent",
		"RTCPeerConnection",
		"RTCIceTransport",
		"RTCIceCandidate",
		"RTCErrorEvent",
		"RTCError",
		"RTCEncodedVideoFrame",
		"RTCEncodedAudioFrame",
		"RTCDtlsTransport",
		"RTCDataChannelEvent",
		"RTCDataChannel",
		"RTCDTMFToneChangeEvent",
		"RTCDTMFSender",
		"RTCCertificate",
		"PromiseRejectionEvent",
		"ProgressEvent",
		"Profiler",
		"ProcessingInstruction",
		"PopStateEvent",
		"PointerEvent",
		"PluginArray",
		"Plugin",
		"PictureInPictureWindow",
		"PictureInPictureEvent",
		"PeriodicWave",
		"PerformanceTiming",
		"PerformanceServerTiming",
		"PerformanceResourceTiming",
		"PerformancePaintTiming",
		"PerformanceObserverEntryList",
		"PerformanceObserver",
		"PerformanceNavigationTiming",
		"PerformanceNavigation",
		"PerformanceMeasure",
		"PerformanceMark",
		"PerformanceLongTaskTiming",
		"PerformanceEventTiming",
		"PerformanceEntry",
		"PerformanceElementTiming",
		"Performance",
		"Path2D",
		"PannerNode",
		"PageTransitionEvent",
		"OverconstrainedError",
		"OscillatorNode",
		"OffscreenCanvasRenderingContext2D",
		"OffscreenCanvas",
		"OfflineAudioContext",
		"OfflineAudioCompletionEvent",
		"NodeList",
		"NodeIterator",
		"NodeFilter",
		"Node",
		"NetworkInformation",
		"Navigator",
		"NavigationTransition",
		"NavigationHistoryEntry",
		"NavigationDestination",
		"NavigationCurrentEntryChangeEvent",
		"Navigation",
		"NavigateEvent",
		"NamedNodeMap",
		"MutationRecord",
		"MutationObserver",
		"MouseEvent",
		"MimeTypeArray",
		"MimeType",
		"MessagePort",
		"MessageEvent",
		"MessageChannel",
		"MediaStreamTrackProcessor",
		"MediaStreamTrackGenerator",
		"MediaStreamTrackEvent",
		"MediaStreamTrack",
		"MediaStreamEvent",
		"MediaStreamAudioSourceNode",
		"MediaStreamAudioDestinationNode",
		"MediaStream",
		"MediaSourceHandle",
		"MediaSource",
		"MediaRecorder",
		"MediaQueryListEvent",
		"MediaQueryList",
		"MediaList",
		"MediaError",
		"MediaEncryptedEvent",
		"MediaElementAudioSourceNode",
		"MediaCapabilities",
		"MathMLElement",
		"Location",
		"LayoutShiftAttribution",
		"LayoutShift",
		"LargestContentfulPaint",
		"KeyframeEffect",
		"KeyboardEvent",
		"IntersectionObserverEntry",
		"IntersectionObserver",
		"InputEvent",
		"InputDeviceInfo",
		"InputDeviceCapabilities",
		"ImageData",
		"ImageCapture",
		"ImageBitmapRenderingContext",
		"ImageBitmap",
		"IdleDeadline",
		"IIRFilterNode",
		"IDBVersionChangeEvent",
		"IDBTransaction",
		"IDBRequest",
		"IDBOpenDBRequest",
		"IDBObjectStore",
		"IDBKeyRange",
		"IDBIndex",
		"IDBFactory",
		"IDBDatabase",
		"IDBCursorWithValue",
		"IDBCursor",
		"History",
		"Headers",
		"HashChangeEvent",
		"HTMLVideoElement",
		"HTMLUnknownElement",
		"HTMLUListElement",
		"HTMLTrackElement",
		"HTMLTitleElement",
		"HTMLTimeElement",
		"HTMLTextAreaElement",
		"HTMLTemplateElement",
		"HTMLTableSectionElement",
		"HTMLTableRowElement",
		"HTMLTableElement",
		"HTMLTableColElement",
		"HTMLTableCellElement",
		"HTMLTableCaptionElement",
		"HTMLStyleElement",
		"HTMLSpanElement",
		"HTMLSourceElement",
		"HTMLSlotElement",
		"HTMLSelectElement",
		"HTMLScriptElement",
		"HTMLQuoteElement",
		"HTMLProgressElement",
		"HTMLPreElement",
		"HTMLPictureElement",
		"HTMLParamElement",
		"HTMLParagraphElement",
		"HTMLOutputElement",
		"HTMLOptionsCollection",
		"HTMLOptionElement",
		"HTMLOptGroupElement",
		"HTMLObjectElement",
		"HTMLOListElement",
		"HTMLModElement",
		"HTMLMeterElement",
		"HTMLMetaElement",
		"HTMLMenuElement",
		"HTMLMediaElement",
		"HTMLMarqueeElement",
		"HTMLMapElement",
		"HTMLLinkElement",
		"HTMLLegendElement",
		"HTMLLabelElement",
		"HTMLLIElement",
		"HTMLInputElement",
		"HTMLImageElement",
		"HTMLIFrameElement",
		"HTMLHtmlElement",
		"HTMLHeadingElement",
		"HTMLHeadElement",
		"HTMLHRElement",
		"HTMLFrameSetElement",
		"HTMLFrameElement",
		"HTMLFormElement",
		"HTMLFormControlsCollection",
		"HTMLFontElement",
		"HTMLFieldSetElement",
		"HTMLEmbedElement",
		"HTMLElement",
		"HTMLDocument",
		"HTMLDivElement",
		"HTMLDirectoryElement",
		"HTMLDialogElement",
		"HTMLDetailsElement",
		"HTMLDataListElement",
		"HTMLDataElement",
		"HTMLDListElement",
		"HTMLCollection",
		"HTMLCanvasElement",
		"HTMLButtonElement",
		"HTMLBodyElement",
		"HTMLBaseElement",
		"HTMLBRElement",
		"HTMLAudioElement",
		"HTMLAreaElement",
		"HTMLAnchorElement",
		"HTMLAllCollection",
		"GeolocationPositionError",
		"GeolocationPosition",
		"GeolocationCoordinates",
		"Geolocation",
		"GamepadHapticActuator",
		"GamepadEvent",
		"GamepadButton",
		"Gamepad",
		"GainNode",
		"FormDataEvent",
		"FormData",
		"FontFaceSetLoadEvent",
		"FontFace",
		"FocusEvent",
		"FileReader",
		"FileList",
		"File",
		"FeaturePolicy",
		"External",
		"EventTarget",
		"EventSource",
		"EventCounts",
		"Event",
		"ErrorEvent",
		"ElementInternals",
		"Element",
		"DynamicsCompressorNode",
		"DragEvent",
		"DocumentType",
		"DocumentFragment",
		"Document",
		"DelayNode",
		"DecompressionStream",
		"DataTransferItemList",
		"DataTransferItem",
		"DataTransfer",
		"DOMTokenList",
		"DOMStringMap",
		"DOMStringList",
		"DOMRectReadOnly",
		"DOMRectList",
		"DOMRect",
		"DOMQuad",
		"DOMPointReadOnly",
		"DOMPoint",
		"DOMParser",
		"DOMMatrixReadOnly",
		"DOMMatrix",
		"DOMImplementation",
		"DOMException",
		"DOMError",
		"CustomStateSet",
		"CustomEvent",
		"CustomElementRegistry",
		"Crypto",
		"CountQueuingStrategy",
		"ConvolverNode",
		"ConstantSourceNode",
		"CompressionStream",
		"CompositionEvent",
		"Comment",
		"CloseEvent",
		"ClipboardEvent",
		"CharacterData",
		"ChannelSplitterNode",
		"ChannelMergerNode",
		"CanvasRenderingContext2D",
		"CanvasPattern",
		"CanvasGradient",
		"CanvasCaptureMediaStreamTrack",
		"CSSVariableReferenceValue",
		"CSSUnparsedValue",
		"CSSUnitValue",
		"CSSTranslate",
		"CSSTransformValue",
		"CSSTransformComponent",
		"CSSSupportsRule",
		"CSSStyleValue",
		"CSSStyleSheet",
		"CSSStyleRule",
		"CSSStyleDeclaration",
		"CSSSkewY",
		"CSSSkewX",
		"CSSSkew",
		"CSSScale",
		"CSSRuleList",
		"CSSRule",
		"CSSRotate",
		"CSSPropertyRule",
		"CSSPositionValue",
		"CSSPerspective",
		"CSSPageRule",
		"CSSNumericValue",
		"CSSNumericArray",
		"CSSNamespaceRule",
		"CSSMediaRule",
		"CSSMatrixComponent",
		"CSSMathValue",
		"CSSMathSum",
		"CSSMathProduct",
		"CSSMathNegate",
		"CSSMathMin",
		"CSSMathMax",
		"CSSMathInvert",
		"CSSMathClamp",
		"CSSLayerStatementRule",
		"CSSLayerBlockRule",
		"CSSKeywordValue",
		"CSSKeyframesRule",
		"CSSKeyframeRule",
		"CSSImportRule",
		"CSSImageValue",
		"CSSGroupingRule",
		"CSSFontPaletteValuesRule",
		"CSSFontFaceRule",
		"CSSCounterStyleRule",
		"CSSContainerRule",
		"CSSConditionRule",
		"CSS",
		"CDATASection",
		"ByteLengthQueuingStrategy",
		"BroadcastChannel",
		"BlobEvent",
		"Blob",
		"BiquadFilterNode",
		"BeforeUnloadEvent",
		"BeforeInstallPromptEvent",
		"BaseAudioContext",
		"BarProp",
		"AudioWorkletNode",
		"AudioSinkInfo",
		"AudioScheduledSourceNode",
		"AudioProcessingEvent",
		"AudioParamMap",
		"AudioParam",
		"AudioNode",
		"AudioListener",
		"AudioDestinationNode",
		"AudioContext",
		"AudioBufferSourceNode",
		"AudioBuffer",
		"Attr",
		"AnimationEvent",
		"AnimationEffect",
		"Animation",
		"AnalyserNode",
		"AbstractRange",
		"AbortSignal",
		"AbortController",
		"window",
		"self",
		"document",
		"name",
		"location",
		"customElements",
		"history",
		"navigation",
		"locationbar",
		"menubar",
		"personalbar",
		"scrollbars",
		"statusbar",
		"toolbar",
		"status",
		"closed",
		"frames",
		"length",
		"top",
		"opener",
		"parent",
		"frameElement",
		"navigator",
		"origin",
		"external",
		"screen",
		"innerWidth",
		"innerHeight",
		"scrollX",
		"pageXOffset",
		"scrollY",
		"pageYOffset",
		"visualViewport",
		"screenX",
		"screenY",
		"outerWidth",
		"outerHeight",
		"devicePixelRatio",
		"event",
		"clientInformation",
		"offscreenBuffering",
		"screenLeft",
		"screenTop",
		"styleMedia",
		"onsearch",
		"isSecureContext",
		"trustedTypes",
		"performance",
		"onappinstalled",
		"onbeforeinstallprompt",
		"crypto",
		"indexedDB",
		"sessionStorage",
		"localStorage",
		"onbeforexrselect",
		"onabort",
		"onbeforeinput",
		"onblur",
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncontextlost",
		"oncontextmenu",
		"oncontextrestored",
		"oncuechange",
		"ondblclick",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"onformdata",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onload",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onreset",
		"onresize",
		"onscroll",
		"onsecuritypolicyviolation",
		"onseeked",
		"onseeking",
		"onselect",
		"onslotchange",
		"onstalled",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"onvolumechange",
		"onwaiting",
		"onwebkitanimationend",
		"onwebkitanimationiteration",
		"onwebkitanimationstart",
		"onwebkittransitionend",
		"onwheel",
		"onauxclick",
		"ongotpointercapture",
		"onlostpointercapture",
		"onpointerdown",
		"onpointermove",
		"onpointerrawupdate",
		"onpointerup",
		"onpointercancel",
		"onpointerover",
		"onpointerout",
		"onpointerenter",
		"onpointerleave",
		"onselectstart",
		"onselectionchange",
		"onanimationend",
		"onanimationiteration",
		"onanimationstart",
		"ontransitionrun",
		"ontransitionstart",
		"ontransitionend",
		"ontransitioncancel",
		"onafterprint",
		"onbeforeprint",
		"onbeforeunload",
		"onhashchange",
		"onlanguagechange",
		"onmessage",
		"onmessageerror",
		"onoffline",
		"ononline",
		"onpagehide",
		"onpageshow",
		"onpopstate",
		"onrejectionhandled",
		"onstorage",
		"onunhandledrejection",
		"onunload",
		"crossOriginIsolated",
		"scheduler",
		"alert",
		"atob",
		"blur",
		"btoa",
		"cancelAnimationFrame",
		"cancelIdleCallback",
		"captureEvents",
		"clearInterval",
		"clearTimeout",
		"close",
		"confirm",
		"createImageBitmap",
		"fetch",
		"find",
		"focus",
		"getComputedStyle",
		"getSelection",
		"matchMedia",
		"moveBy",
		"moveTo",
		"open",
		"postMessage",
		"print",
		"prompt",
		"queueMicrotask",
		"releaseEvents",
		"reportError",
		"requestAnimationFrame",
		"requestIdleCallback",
		"resizeBy",
		"resizeTo",
		"scroll",
		"scrollBy",
		"scrollTo",
		"setInterval",
		"setTimeout",
		"stop",
		"structuredClone",
		"webkitCancelAnimationFrame",
		"webkitRequestAnimationFrame",
		"chrome",
		"WebAssembly",
		"fence",
		"caches",
		"cookieStore",
		"ondevicemotion",
		"ondeviceorientation",
		"ondeviceorientationabsolute",
		"launchQueue",
		"sharedStorage",
		"documentPictureInPicture",
		"onbeforematch",
		"onbeforetoggle",
		"AbsoluteOrientationSensor",
		"Accelerometer",
		"AudioWorklet",
		"BatteryManager",
		"Cache",
		"CacheStorage",
		"Clipboard",
		"ClipboardItem",
		"CookieChangeEvent",
		"CookieStore",
		"CookieStoreManager",
		"Credential",
		"CredentialsContainer",
		"CryptoKey",
		"DeviceMotionEvent",
		"DeviceMotionEventAcceleration",
		"DeviceMotionEventRotationRate",
		"DeviceOrientationEvent",
		"FederatedCredential",
		"GPU",
		"GPUAdapter",
		"GPUAdapterInfo",
		"GPUBindGroup",
		"GPUBindGroupLayout",
		"GPUBuffer",
		"GPUBufferUsage",
		"GPUCanvasContext",
		"GPUColorWrite",
		"GPUCommandBuffer",
		"GPUCommandEncoder",
		"GPUCompilationInfo",
		"GPUCompilationMessage",
		"GPUComputePassEncoder",
		"GPUComputePipeline",
		"GPUDevice",
		"GPUDeviceLostInfo",
		"GPUError",
		"GPUExternalTexture",
		"GPUInternalError",
		"GPUMapMode",
		"GPUOutOfMemoryError",
		"GPUPipelineError",
		"GPUPipelineLayout",
		"GPUQuerySet",
		"GPUQueue",
		"GPURenderBundle",
		"GPURenderBundleEncoder",
		"GPURenderPassEncoder",
		"GPURenderPipeline",
		"GPUSampler",
		"GPUShaderModule",
		"GPUShaderStage",
		"GPUSupportedFeatures",
		"GPUSupportedLimits",
		"GPUTexture",
		"GPUTextureUsage",
		"GPUTextureView",
		"GPUUncapturedErrorEvent",
		"GPUValidationError",
		"GravitySensor",
		"Gyroscope",
		"Keyboard",
		"KeyboardLayoutMap",
		"LinearAccelerationSensor",
		"Lock",
		"LockManager",
		"MIDIAccess",
		"MIDIConnectionEvent",
		"MIDIInput",
		"MIDIInputMap",
		"MIDIMessageEvent",
		"MIDIOutput",
		"MIDIOutputMap",
		"MIDIPort",
		"MediaDeviceInfo",
		"MediaDevices",
		"MediaKeyMessageEvent",
		"MediaKeySession",
		"MediaKeyStatusMap",
		"MediaKeySystemAccess",
		"MediaKeys",
		"NavigationPreloadManager",
		"NavigatorManagedData",
		"OrientationSensor",
		"PasswordCredential",
		"RelativeOrientationSensor",
		"Sanitizer",
		"ScreenDetailed",
		"ScreenDetails",
		"Sensor",
		"SensorErrorEvent",
		"ServiceWorker",
		"ServiceWorkerContainer",
		"ServiceWorkerRegistration",
		"StorageManager",
		"SubtleCrypto",
		"VirtualKeyboard",
		"WGSLLanguageFeatures",
		"WebTransport",
		"WebTransportBidirectionalStream",
		"WebTransportDatagramDuplexStream",
		"WebTransportError",
		"Worklet",
		"XRDOMOverlayState",
		"XRLayer",
		"XRWebGLBinding",
		"AudioData",
		"EncodedAudioChunk",
		"EncodedVideoChunk",
		"ImageTrack",
		"ImageTrackList",
		"VideoColorSpace",
		"VideoFrame",
		"AudioDecoder",
		"AudioEncoder",
		"ImageDecoder",
		"VideoDecoder",
		"VideoEncoder",
		"AuthenticatorAssertionResponse",
		"AuthenticatorAttestationResponse",
		"AuthenticatorResponse",
		"PublicKeyCredential",
		"BarcodeDetector",
		"Bluetooth",
		"BluetoothCharacteristicProperties",
		"BluetoothDevice",
		"BluetoothRemoteGATTCharacteristic",
		"BluetoothRemoteGATTDescriptor",
		"BluetoothRemoteGATTServer",
		"BluetoothRemoteGATTService",
		"CaptureController",
		"DocumentPictureInPicture",
		"EyeDropper",
		"Fence",
		"FencedFrameConfig",
		"HTMLFencedFrameElement",
		"FileSystemDirectoryHandle",
		"FileSystemFileHandle",
		"FileSystemHandle",
		"FileSystemWritableFileStream",
		"FontData",
		"FragmentDirective",
		"HID",
		"HIDConnectionEvent",
		"HIDDevice",
		"HIDInputReportEvent",
		"IdentityCredential",
		"IdentityProvider",
		"IdleDetector",
		"LaunchParams",
		"LaunchQueue",
		"OTPCredential",
		"PaymentAddress",
		"PaymentRequest",
		"PaymentResponse",
		"PaymentMethodChangeEvent",
		"Presentation",
		"PresentationAvailability",
		"PresentationConnection",
		"PresentationConnectionAvailableEvent",
		"PresentationConnectionCloseEvent",
		"PresentationConnectionList",
		"PresentationReceiver",
		"PresentationRequest",
		"Serial",
		"SerialPort",
		"SharedStorage",
		"SharedStorageWorklet",
		"ToggleEvent",
		"USB",
		"USBAlternateInterface",
		"USBConfiguration",
		"USBConnectionEvent",
		"USBDevice",
		"USBEndpoint",
		"USBInTransferResult",
		"USBInterface",
		"USBIsochronousInTransferPacket",
		"USBIsochronousInTransferResult",
		"USBIsochronousOutTransferPacket",
		"USBIsochronousOutTransferResult",
		"USBOutTransferResult",
		"WakeLock",
		"WakeLockSentinel",
		"WindowControlsOverlay",
		"WindowControlsOverlayGeometryChangeEvent",
		"XRAnchor",
		"XRAnchorSet",
		"XRBoundedReferenceSpace",
		"XRCPUDepthInformation",
		"XRCamera",
		"XRDepthInformation",
		"XRFrame",
		"XRHitTestResult",
		"XRHitTestSource",
		"XRInputSource",
		"XRInputSourceArray",
		"XRInputSourceEvent",
		"XRInputSourcesChangeEvent",
		"XRLightEstimate",
		"XRLightProbe",
		"XRPose",
		"XRRay",
		"XRReferenceSpace",
		"XRReferenceSpaceEvent",
		"XRRenderState",
		"XRRigidTransform",
		"XRSession",
		"XRSessionEvent",
		"XRSpace",
		"XRSystem",
		"XRTransientInputHitTestResult",
		"XRTransientInputHitTestSource",
		"XRView",
		"XRViewerPose",
		"XRViewport",
		"XRWebGLDepthInformation",
		"XRWebGLLayer",
		"getScreenDetails",
		"openDatabase",
		"queryLocalFonts",
		"showDirectoryPicker",
		"showOpenFilePicker",
		"showSaveFilePicker",
		"originAgentCluster",
		"credentialless",
		"speechSynthesis",
		"oncontentvisibilityautostatechange",
		"onscrollend",
		"AnimationPlaybackEvent",
		"AnimationTimeline",
		"CSSAnimation",
		"CSSTransition",
		"DocumentTimeline",
		"BackgroundFetchManager",
		"BackgroundFetchRecord",
		"BackgroundFetchRegistration",
		"BluetoothUUID",
		"BrowserCaptureMediaStreamTrack",
		"CropTarget",
		"CSSStartingStyleRule",
		"ContentVisibilityAutoStateChangeEvent",
		"DelegatedInkTrailPresenter",
		"Ink",
		"DocumentPictureInPictureEvent",
		"Highlight",
		"HighlightRegistry",
		"MediaMetadata",
		"MediaSession",
		"MutationEvent",
		"NavigatorUAData",
		"Notification",
		"PaymentManager",
		"PaymentRequestUpdateEvent",
		"PeriodicSyncManager",
		"PermissionStatus",
		"Permissions",
		"PushManager",
		"PushSubscription",
		"PushSubscriptionOptions",
		"RemotePlayback",
		"ScrollTimeline",
		"ViewTimeline",
		"SharedWorker",
		"SpeechSynthesisErrorEvent",
		"SpeechSynthesisEvent",
		"SpeechSynthesisUtterance",
		"VideoPlaybackQuality",
		"ViewTransition",
		"VisibilityStateEntry",
		"webkitSpeechGrammar",
		"webkitSpeechGrammarList",
		"webkitSpeechRecognition",
		"webkitSpeechRecognitionError",
		"webkitSpeechRecognitionEvent",
		"webkitRequestFileSystem",
		"webkitResolveLocalFileSystemURL",
		"TEMPORARY",
		"PERSISTENT",
		"addEventListener",
		"dispatchEvent",
		"removeEventListener"
	],
	"document": [
		"location",
		"implementation",
		"URL",
		"documentURI",
		"compatMode",
		"characterSet",
		"charset",
		"inputEncoding",
		"contentType",
		"doctype",
		"documentElement",
		"xmlEncoding",
		"xmlVersion",
		"xmlStandalone",
		"domain",
		"referrer",
		"cookie",
		"lastModified",
		"readyState",
		"title",
		"dir",
		"body",
		"head",
		"images",
		"embeds",
		"plugins",
		"links",
		"forms",
		"scripts",
		"currentScript",
		"defaultView",
		"designMode",
		"onreadystatechange",
		"anchors",
		"applets",
		"fgColor",
		"linkColor",
		"vlinkColor",
		"alinkColor",
		"bgColor",
		"all",
		"scrollingElement",
		"onpointerlockchange",
		"onpointerlockerror",
		"hidden",
		"visibilityState",
		"wasDiscarded",
		"prerendering",
		"featurePolicy",
		"webkitVisibilityState",
		"webkitHidden",
		"onbeforecopy",
		"onbeforecut",
		"onbeforepaste",
		"onfreeze",
		"onprerenderingchange",
		"onresume",
		"onsearch",
		"onvisibilitychange",
		"fullscreenEnabled",
		"fullscreen",
		"onfullscreenchange",
		"onfullscreenerror",
		"webkitIsFullScreen",
		"webkitCurrentFullScreenElement",
		"webkitFullscreenEnabled",
		"webkitFullscreenElement",
		"onwebkitfullscreenchange",
		"onwebkitfullscreenerror",
		"rootElement",
		"pictureInPictureEnabled",
		"onbeforexrselect",
		"onabort",
		"onbeforeinput",
		"onblur",
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncontextlost",
		"oncontextmenu",
		"oncontextrestored",
		"oncuechange",
		"ondblclick",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"onformdata",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onload",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onreset",
		"onresize",
		"onscroll",
		"onsecuritypolicyviolation",
		"onseeked",
		"onseeking",
		"onselect",
		"onslotchange",
		"onstalled",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"onvolumechange",
		"onwaiting",
		"onwebkitanimationend",
		"onwebkitanimationiteration",
		"onwebkitanimationstart",
		"onwebkittransitionend",
		"onwheel",
		"onauxclick",
		"ongotpointercapture",
		"onlostpointercapture",
		"onpointerdown",
		"onpointermove",
		"onpointerrawupdate",
		"onpointerup",
		"onpointercancel",
		"onpointerover",
		"onpointerout",
		"onpointerenter",
		"onpointerleave",
		"onselectstart",
		"onselectionchange",
		"onanimationend",
		"onanimationiteration",
		"onanimationstart",
		"ontransitionrun",
		"ontransitionstart",
		"ontransitionend",
		"ontransitioncancel",
		"oncopy",
		"oncut",
		"onpaste",
		"children",
		"firstElementChild",
		"lastElementChild",
		"childElementCount",
		"activeElement",
		"styleSheets",
		"pointerLockElement",
		"fullscreenElement",
		"adoptedStyleSheets",
		"pictureInPictureElement",
		"fonts",
		"adoptNode",
		"append",
		"captureEvents",
		"caretRangeFromPoint",
		"clear",
		"close",
		"createAttribute",
		"createAttributeNS",
		"createCDATASection",
		"createComment",
		"createDocumentFragment",
		"createElement",
		"createElementNS",
		"createEvent",
		"createExpression",
		"createNSResolver",
		"createNodeIterator",
		"createProcessingInstruction",
		"createRange",
		"createTextNode",
		"createTreeWalker",
		"elementFromPoint",
		"elementsFromPoint",
		"evaluate",
		"execCommand",
		"exitFullscreen",
		"exitPictureInPicture",
		"exitPointerLock",
		"getElementById",
		"getElementsByClassName",
		"getElementsByName",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getSelection",
		"hasFocus",
		"importNode",
		"open",
		"prepend",
		"queryCommandEnabled",
		"queryCommandIndeterm",
		"queryCommandState",
		"queryCommandSupported",
		"queryCommandValue",
		"querySelector",
		"querySelectorAll",
		"releaseEvents",
		"replaceChildren",
		"webkitCancelFullScreen",
		"webkitExitFullscreen",
		"write",
		"writeln",
		"fragmentDirective",
		"onbeforematch",
		"onbeforetoggle",
		"browsingTopics",
		"hasPrivateToken",
		"hasRedemptionRecord",
		"timeline",
		"oncontentvisibilityautostatechange",
		"onscrollend",
		"getAnimations",
		"hasStorageAccess",
		"requestStorageAccess",
		"requestStorageAccessFor",
		"startViewTransition",
		"nodeType",
		"nodeName",
		"baseURI",
		"isConnected",
		"ownerDocument",
		"parentNode",
		"parentElement",
		"childNodes",
		"firstChild",
		"lastChild",
		"previousSibling",
		"nextSibling",
		"nodeValue",
		"textContent",
		"ELEMENT_NODE",
		"ATTRIBUTE_NODE",
		"TEXT_NODE",
		"CDATA_SECTION_NODE",
		"ENTITY_REFERENCE_NODE",
		"ENTITY_NODE",
		"PROCESSING_INSTRUCTION_NODE",
		"COMMENT_NODE",
		"DOCUMENT_NODE",
		"DOCUMENT_TYPE_NODE",
		"DOCUMENT_FRAGMENT_NODE",
		"NOTATION_NODE",
		"DOCUMENT_POSITION_DISCONNECTED",
		"DOCUMENT_POSITION_PRECEDING",
		"DOCUMENT_POSITION_FOLLOWING",
		"DOCUMENT_POSITION_CONTAINS",
		"DOCUMENT_POSITION_CONTAINED_BY",
		"DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC",
		"appendChild",
		"cloneNode",
		"compareDocumentPosition",
		"contains",
		"getRootNode",
		"hasChildNodes",
		"insertBefore",
		"isDefaultNamespace",
		"isEqualNode",
		"isSameNode",
		"lookupNamespaceURI",
		"lookupPrefix",
		"normalize",
		"removeChild",
		"replaceChild",
		"addEventListener",
		"dispatchEvent",
		"removeEventListener"
	],
	"documentElement": [
		"version",
		"title",
		"lang",
		"translate",
		"dir",
		"hidden",
		"accessKey",
		"draggable",
		"spellcheck",
		"autocapitalize",
		"contentEditable",
		"enterKeyHint",
		"isContentEditable",
		"inputMode",
		"virtualKeyboardPolicy",
		"offsetParent",
		"offsetTop",
		"offsetLeft",
		"offsetWidth",
		"offsetHeight",
		"innerText",
		"outerText",
		"onbeforexrselect",
		"onabort",
		"onbeforeinput",
		"onblur",
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncontextlost",
		"oncontextmenu",
		"oncontextrestored",
		"oncuechange",
		"ondblclick",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"onformdata",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onload",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onreset",
		"onresize",
		"onscroll",
		"onsecuritypolicyviolation",
		"onseeked",
		"onseeking",
		"onselect",
		"onslotchange",
		"onstalled",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"onvolumechange",
		"onwaiting",
		"onwebkitanimationend",
		"onwebkitanimationiteration",
		"onwebkitanimationstart",
		"onwebkittransitionend",
		"onwheel",
		"onauxclick",
		"ongotpointercapture",
		"onlostpointercapture",
		"onpointerdown",
		"onpointermove",
		"onpointerrawupdate",
		"onpointerup",
		"onpointercancel",
		"onpointerover",
		"onpointerout",
		"onpointerenter",
		"onpointerleave",
		"onselectstart",
		"onselectionchange",
		"onanimationend",
		"onanimationiteration",
		"onanimationstart",
		"ontransitionrun",
		"ontransitionstart",
		"ontransitionend",
		"ontransitioncancel",
		"oncopy",
		"oncut",
		"onpaste",
		"dataset",
		"nonce",
		"autofocus",
		"tabIndex",
		"style",
		"attributeStyleMap",
		"attachInternals",
		"blur",
		"click",
		"focus",
		"inert",
		"oncontentvisibilityautostatechange",
		"onscrollend",
		"popover",
		"onbeforetoggle",
		"onbeforematch",
		"hidePopover",
		"showPopover",
		"togglePopover",
		"namespaceURI",
		"prefix",
		"localName",
		"tagName",
		"id",
		"className",
		"classList",
		"slot",
		"attributes",
		"shadowRoot",
		"part",
		"assignedSlot",
		"innerHTML",
		"outerHTML",
		"scrollTop",
		"scrollLeft",
		"scrollWidth",
		"scrollHeight",
		"clientTop",
		"clientLeft",
		"clientWidth",
		"clientHeight",
		"onbeforecopy",
		"onbeforecut",
		"onbeforepaste",
		"onsearch",
		"elementTiming",
		"onfullscreenchange",
		"onfullscreenerror",
		"onwebkitfullscreenchange",
		"onwebkitfullscreenerror",
		"role",
		"ariaAtomic",
		"ariaAutoComplete",
		"ariaBusy",
		"ariaBrailleLabel",
		"ariaBrailleRoleDescription",
		"ariaChecked",
		"ariaColCount",
		"ariaColIndex",
		"ariaColSpan",
		"ariaCurrent",
		"ariaDescription",
		"ariaDisabled",
		"ariaExpanded",
		"ariaHasPopup",
		"ariaHidden",
		"ariaInvalid",
		"ariaKeyShortcuts",
		"ariaLabel",
		"ariaLevel",
		"ariaLive",
		"ariaModal",
		"ariaMultiLine",
		"ariaMultiSelectable",
		"ariaOrientation",
		"ariaPlaceholder",
		"ariaPosInSet",
		"ariaPressed",
		"ariaReadOnly",
		"ariaRelevant",
		"ariaRequired",
		"ariaRoleDescription",
		"ariaRowCount",
		"ariaRowIndex",
		"ariaRowSpan",
		"ariaSelected",
		"ariaSetSize",
		"ariaSort",
		"ariaValueMax",
		"ariaValueMin",
		"ariaValueNow",
		"ariaValueText",
		"children",
		"firstElementChild",
		"lastElementChild",
		"childElementCount",
		"previousElementSibling",
		"nextElementSibling",
		"after",
		"animate",
		"append",
		"attachShadow",
		"before",
		"closest",
		"computedStyleMap",
		"getAttribute",
		"getAttributeNS",
		"getAttributeNames",
		"getAttributeNode",
		"getAttributeNodeNS",
		"getBoundingClientRect",
		"getClientRects",
		"getElementsByClassName",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getInnerHTML",
		"hasAttribute",
		"hasAttributeNS",
		"hasAttributes",
		"hasPointerCapture",
		"insertAdjacentElement",
		"insertAdjacentHTML",
		"insertAdjacentText",
		"matches",
		"prepend",
		"querySelector",
		"querySelectorAll",
		"releasePointerCapture",
		"remove",
		"removeAttribute",
		"removeAttributeNS",
		"removeAttributeNode",
		"replaceChildren",
		"replaceWith",
		"requestFullscreen",
		"requestPointerLock",
		"scroll",
		"scrollBy",
		"scrollIntoView",
		"scrollIntoViewIfNeeded",
		"scrollTo",
		"setAttribute",
		"setAttributeNS",
		"setAttributeNode",
		"setAttributeNodeNS",
		"setPointerCapture",
		"toggleAttribute",
		"webkitMatchesSelector",
		"webkitRequestFullScreen",
		"webkitRequestFullscreen",
		"checkVisibility",
		"getAnimations",
		"setHTML",
		"nodeType",
		"nodeName",
		"baseURI",
		"isConnected",
		"ownerDocument",
		"parentNode",
		"parentElement",
		"childNodes",
		"firstChild",
		"lastChild",
		"previousSibling",
		"nextSibling",
		"nodeValue",
		"textContent",
		"ELEMENT_NODE",
		"ATTRIBUTE_NODE",
		"TEXT_NODE",
		"CDATA_SECTION_NODE",
		"ENTITY_REFERENCE_NODE",
		"ENTITY_NODE",
		"PROCESSING_INSTRUCTION_NODE",
		"COMMENT_NODE",
		"DOCUMENT_NODE",
		"DOCUMENT_TYPE_NODE",
		"DOCUMENT_FRAGMENT_NODE",
		"NOTATION_NODE",
		"DOCUMENT_POSITION_DISCONNECTED",
		"DOCUMENT_POSITION_PRECEDING",
		"DOCUMENT_POSITION_FOLLOWING",
		"DOCUMENT_POSITION_CONTAINS",
		"DOCUMENT_POSITION_CONTAINED_BY",
		"DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC",
		"appendChild",
		"cloneNode",
		"compareDocumentPosition",
		"contains",
		"getRootNode",
		"hasChildNodes",
		"insertBefore",
		"isDefaultNamespace",
		"isEqualNode",
		"isSameNode",
		"lookupNamespaceURI",
		"lookupPrefix",
		"normalize",
		"removeChild",
		"replaceChild",
		"addEventListener",
		"dispatchEvent",
		"removeEventListener"
	],
	"speechSynthesis": [
		{
			"voiceURI": "Anna",
			"name": "Anna",
			"lang": "de-DE",
			"localService": true,
			"default": true
		},
		{
			"voiceURI": "Aaron",
			"name": "Aaron",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Albert",
			"name": "Albert",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Alice",
			"name": "Alice",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Alva",
			"name": "Alva",
			"lang": "sv-SE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Amélie",
			"name": "Amélie",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Amira",
			"name": "Amira",
			"lang": "ms-MY",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Arthur",
			"name": "Arthur",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Bahh",
			"name": "Bahh",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Boing",
			"name": "Boing",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Carmit",
			"name": "Carmit",
			"lang": "he-IL",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Catherine",
			"name": "Catherine",
			"lang": "en-AU",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Cellos",
			"name": "Cellos",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Damayanti",
			"name": "Damayanti",
			"lang": "id-ID",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Daniel (Englisch (Vereinigtes Königreich))",
			"name": "Daniel (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Daniel (Französisch (Frankreich))",
			"name": "Daniel (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Daria",
			"name": "Daria",
			"lang": "bg-BG",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Deutsch (Deutschland))",
			"name": "Eddy (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Englisch (Vereinigtes Königreich))",
			"name": "Eddy (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Englisch (Vereinigte Staaten))",
			"name": "Eddy (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Spanisch (Spanien))",
			"name": "Eddy (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Spanisch (Mexiko))",
			"name": "Eddy (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Finnisch (Finnland))",
			"name": "Eddy (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Französisch (Kanada))",
			"name": "Eddy (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Französisch (Frankreich))",
			"name": "Eddy (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Italienisch (Italien))",
			"name": "Eddy (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Eddy (Portugiesisch (Brasilien))",
			"name": "Eddy (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Ellen",
			"name": "Ellen",
			"lang": "nl-BE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Deutsch (Deutschland))",
			"name": "Flo (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Englisch (Vereinigtes Königreich))",
			"name": "Flo (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Englisch (Vereinigte Staaten))",
			"name": "Flo (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Spanisch (Spanien))",
			"name": "Flo (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Spanisch (Mexiko))",
			"name": "Flo (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Finnisch (Finnland))",
			"name": "Flo (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Französisch (Kanada))",
			"name": "Flo (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Französisch (Frankreich))",
			"name": "Flo (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Italienisch (Italien))",
			"name": "Flo (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flo (Portugiesisch (Brasilien))",
			"name": "Flo (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Flüstern",
			"name": "Flüstern",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Fred",
			"name": "Fred",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Glocken",
			"name": "Glocken",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Gordon",
			"name": "Gordon",
			"lang": "en-AU",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Deutsch (Deutschland))",
			"name": "Grandma (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Englisch (Vereinigtes Königreich))",
			"name": "Grandma (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Englisch (Vereinigte Staaten))",
			"name": "Grandma (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Spanisch (Spanien))",
			"name": "Grandma (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Spanisch (Mexiko))",
			"name": "Grandma (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Finnisch (Finnland))",
			"name": "Grandma (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Französisch (Kanada))",
			"name": "Grandma (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Französisch (Frankreich))",
			"name": "Grandma (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Italienisch (Italien))",
			"name": "Grandma (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandma (Portugiesisch (Brasilien))",
			"name": "Grandma (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Deutsch (Deutschland))",
			"name": "Grandpa (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Englisch (Vereinigtes Königreich))",
			"name": "Grandpa (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Englisch (Vereinigte Staaten))",
			"name": "Grandpa (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Spanisch (Spanien))",
			"name": "Grandpa (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Spanisch (Mexiko))",
			"name": "Grandpa (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Finnisch (Finnland))",
			"name": "Grandpa (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Französisch (Kanada))",
			"name": "Grandpa (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Französisch (Frankreich))",
			"name": "Grandpa (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Italienisch (Italien))",
			"name": "Grandpa (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Grandpa (Portugiesisch (Brasilien))",
			"name": "Grandpa (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Gute Neuigkeiten",
			"name": "Gute Neuigkeiten",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Hattori",
			"name": "Hattori",
			"lang": "ja-JP",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Helena",
			"name": "Helena",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Ioana",
			"name": "Ioana",
			"lang": "ro-RO",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Jacques",
			"name": "Jacques",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Joana",
			"name": "Joana",
			"lang": "pt-PT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Junior",
			"name": "Junior",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Kanya",
			"name": "Kanya",
			"lang": "th-TH",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Karen",
			"name": "Karen",
			"lang": "en-AU",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Katrin",
			"name": "Katrin",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Kyoko",
			"name": "Kyoko",
			"lang": "ja-JP",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Lana",
			"name": "Lana",
			"lang": "hr-HR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Laura",
			"name": "Laura",
			"lang": "sk-SK",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Lekha",
			"name": "Lekha",
			"lang": "hi-IN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Lesya",
			"name": "Lesya",
			"lang": "uk-UA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Li-Mu",
			"name": "Li-Mu",
			"lang": "zh-CN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Linh",
			"name": "Linh",
			"lang": "vi-VN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Luciana",
			"name": "Luciana",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Majed",
			"name": "Majed",
			"lang": "ar-001",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Marie",
			"name": "Marie",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Martha",
			"name": "Martha",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Martin",
			"name": "Martin",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Meijia",
			"name": "Meijia",
			"lang": "zh-TW",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Melina",
			"name": "Melina",
			"lang": "el-GR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Milena",
			"name": "Milena",
			"lang": "ru-RU",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Moira",
			"name": "Moira",
			"lang": "en-IE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Monster",
			"name": "Monster",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Montse",
			"name": "Montse",
			"lang": "ca-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Mónica",
			"name": "Mónica",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Nicky",
			"name": "Nicky",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Nora",
			"name": "Nora",
			"lang": "nb-NO",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "O-Ren",
			"name": "O-Ren",
			"lang": "ja-JP",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Orgel",
			"name": "Orgel",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Paulina",
			"name": "Paulina",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Ralph",
			"name": "Ralph",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Deutsch (Deutschland))",
			"name": "Reed (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Englisch (Vereinigtes Königreich))",
			"name": "Reed (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Englisch (Vereinigte Staaten))",
			"name": "Reed (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Spanisch (Spanien))",
			"name": "Reed (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Spanisch (Mexiko))",
			"name": "Reed (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Finnisch (Finnland))",
			"name": "Reed (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Französisch (Kanada))",
			"name": "Reed (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Italienisch (Italien))",
			"name": "Reed (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Reed (Portugiesisch (Brasilien))",
			"name": "Reed (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rishi",
			"name": "Rishi",
			"lang": "en-IN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Deutsch (Deutschland))",
			"name": "Rocko (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Englisch (Vereinigtes Königreich))",
			"name": "Rocko (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Englisch (Vereinigte Staaten))",
			"name": "Rocko (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Spanisch (Spanien))",
			"name": "Rocko (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Spanisch (Mexiko))",
			"name": "Rocko (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Finnisch (Finnland))",
			"name": "Rocko (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Französisch (Kanada))",
			"name": "Rocko (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Französisch (Frankreich))",
			"name": "Rocko (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Italienisch (Italien))",
			"name": "Rocko (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Rocko (Portugiesisch (Brasilien))",
			"name": "Rocko (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Samantha",
			"name": "Samantha",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Deutsch (Deutschland))",
			"name": "Sandy (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Englisch (Vereinigtes Königreich))",
			"name": "Sandy (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Englisch (Vereinigte Staaten))",
			"name": "Sandy (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Spanisch (Spanien))",
			"name": "Sandy (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Spanisch (Mexiko))",
			"name": "Sandy (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Finnisch (Finnland))",
			"name": "Sandy (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Französisch (Kanada))",
			"name": "Sandy (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Französisch (Frankreich))",
			"name": "Sandy (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Italienisch (Italien))",
			"name": "Sandy (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sandy (Portugiesisch (Brasilien))",
			"name": "Sandy (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sara",
			"name": "Sara",
			"lang": "da-DK",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Satu",
			"name": "Satu",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Schlechte Neuigkeiten",
			"name": "Schlechte Neuigkeiten",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Seifenblasen",
			"name": "Seifenblasen",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Deutsch (Deutschland))",
			"name": "Shelley (Deutsch (Deutschland))",
			"lang": "de-DE",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Englisch (Vereinigtes Königreich))",
			"name": "Shelley (Englisch (Vereinigtes Königreich))",
			"lang": "en-GB",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Englisch (Vereinigte Staaten))",
			"name": "Shelley (Englisch (Vereinigte Staaten))",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Spanisch (Spanien))",
			"name": "Shelley (Spanisch (Spanien))",
			"lang": "es-ES",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Spanisch (Mexiko))",
			"name": "Shelley (Spanisch (Mexiko))",
			"lang": "es-MX",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Finnisch (Finnland))",
			"name": "Shelley (Finnisch (Finnland))",
			"lang": "fi-FI",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Französisch (Kanada))",
			"name": "Shelley (Französisch (Kanada))",
			"lang": "fr-CA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Französisch (Frankreich))",
			"name": "Shelley (Französisch (Frankreich))",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Italienisch (Italien))",
			"name": "Shelley (Italienisch (Italien))",
			"lang": "it-IT",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Shelley (Portugiesisch (Brasilien))",
			"name": "Shelley (Portugiesisch (Brasilien))",
			"lang": "pt-BR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Sinji",
			"name": "Sinji",
			"lang": "zh-HK",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Spaßvogel",
			"name": "Spaßvogel",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Superstar",
			"name": "Superstar",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Tessa",
			"name": "Tessa",
			"lang": "en-ZA",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Thomas",
			"name": "Thomas",
			"lang": "fr-FR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Tingting",
			"name": "Tingting",
			"lang": "zh-CN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Tünde",
			"name": "Tünde",
			"lang": "hu-HU",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Wobble",
			"name": "Wobble",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Xander",
			"name": "Xander",
			"lang": "nl-NL",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Yelda",
			"name": "Yelda",
			"lang": "tr-TR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Yu-shu",
			"name": "Yu-shu",
			"lang": "zh-CN",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Yuna",
			"name": "Yuna",
			"lang": "ko-KR",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Zarvox",
			"name": "Zarvox",
			"lang": "en-US",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Zosia",
			"name": "Zosia",
			"lang": "pl-PL",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Zuzana",
			"name": "Zuzana",
			"lang": "cs-CZ",
			"localService": true,
			"default": false
		},
		{
			"voiceURI": "Google Deutsch",
			"name": "Google Deutsch",
			"lang": "de-DE",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google US English",
			"name": "Google US English",
			"lang": "en-US",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google UK English Female",
			"name": "Google UK English Female",
			"lang": "en-GB",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google UK English Male",
			"name": "Google UK English Male",
			"lang": "en-GB",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google español",
			"name": "Google español",
			"lang": "es-ES",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google español de Estados Unidos",
			"name": "Google español de Estados Unidos",
			"lang": "es-US",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google français",
			"name": "Google français",
			"lang": "fr-FR",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google हिन्दी",
			"name": "Google हिन्दी",
			"lang": "hi-IN",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google Bahasa Indonesia",
			"name": "Google Bahasa Indonesia",
			"lang": "id-ID",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google italiano",
			"name": "Google italiano",
			"lang": "it-IT",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google 日本語",
			"name": "Google 日本語",
			"lang": "ja-JP",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google 한국의",
			"name": "Google 한국의",
			"lang": "ko-KR",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google Nederlands",
			"name": "Google Nederlands",
			"lang": "nl-NL",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google polski",
			"name": "Google polski",
			"lang": "pl-PL",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google português do Brasil",
			"name": "Google português do Brasil",
			"lang": "pt-BR",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google русский",
			"name": "Google русский",
			"lang": "ru-RU",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google 普通话（中国大陆）",
			"name": "Google 普通话（中国大陆）",
			"lang": "zh-CN",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google 粤語（香港）",
			"name": "Google 粤語（香港）",
			"lang": "zh-HK",
			"localService": false,
			"default": false
		},
		{
			"voiceURI": "Google 國語（臺灣）",
			"name": "Google 國語（臺灣）",
			"lang": "zh-TW",
			"localService": false,
			"default": false
		}
	],
	"css": {
		"0": "accent-color",
		"1": "align-content",
		"2": "align-items",
		"3": "align-self",
		"4": "alignment-baseline",
		"5": "animation-composition",
		"6": "animation-delay",
		"7": "animation-direction",
		"8": "animation-duration",
		"9": "animation-fill-mode",
		"10": "animation-iteration-count",
		"11": "animation-name",
		"12": "animation-play-state",
		"13": "animation-range-end",
		"14": "animation-range-start",
		"15": "animation-timeline",
		"16": "animation-timing-function",
		"17": "app-region",
		"18": "appearance",
		"19": "backdrop-filter",
		"20": "backface-visibility",
		"21": "background-attachment",
		"22": "background-blend-mode",
		"23": "background-clip",
		"24": "background-color",
		"25": "background-image",
		"26": "background-origin",
		"27": "background-position",
		"28": "background-repeat",
		"29": "background-size",
		"30": "baseline-shift",
		"31": "baseline-source",
		"32": "block-size",
		"33": "border-block-end-color",
		"34": "border-block-end-style",
		"35": "border-block-end-width",
		"36": "border-block-start-color",
		"37": "border-block-start-style",
		"38": "border-block-start-width",
		"39": "border-bottom-color",
		"40": "border-bottom-left-radius",
		"41": "border-bottom-right-radius",
		"42": "border-bottom-style",
		"43": "border-bottom-width",
		"44": "border-collapse",
		"45": "border-end-end-radius",
		"46": "border-end-start-radius",
		"47": "border-image-outset",
		"48": "border-image-repeat",
		"49": "border-image-slice",
		"50": "border-image-source",
		"51": "border-image-width",
		"52": "border-inline-end-color",
		"53": "border-inline-end-style",
		"54": "border-inline-end-width",
		"55": "border-inline-start-color",
		"56": "border-inline-start-style",
		"57": "border-inline-start-width",
		"58": "border-left-color",
		"59": "border-left-style",
		"60": "border-left-width",
		"61": "border-right-color",
		"62": "border-right-style",
		"63": "border-right-width",
		"64": "border-start-end-radius",
		"65": "border-start-start-radius",
		"66": "border-top-color",
		"67": "border-top-left-radius",
		"68": "border-top-right-radius",
		"69": "border-top-style",
		"70": "border-top-width",
		"71": "bottom",
		"72": "box-shadow",
		"73": "box-sizing",
		"74": "break-after",
		"75": "break-before",
		"76": "break-inside",
		"77": "buffered-rendering",
		"78": "caption-side",
		"79": "caret-color",
		"80": "clear",
		"81": "clip",
		"82": "clip-path",
		"83": "clip-rule",
		"84": "color",
		"85": "color-interpolation",
		"86": "color-interpolation-filters",
		"87": "color-rendering",
		"88": "column-count",
		"89": "column-gap",
		"90": "column-rule-color",
		"91": "column-rule-style",
		"92": "column-rule-width",
		"93": "column-span",
		"94": "column-width",
		"95": "contain-intrinsic-block-size",
		"96": "contain-intrinsic-height",
		"97": "contain-intrinsic-inline-size",
		"98": "contain-intrinsic-size",
		"99": "contain-intrinsic-width",
		"100": "container-name",
		"101": "container-type",
		"102": "content",
		"103": "cursor",
		"104": "cx",
		"105": "cy",
		"106": "d",
		"107": "direction",
		"108": "display",
		"109": "dominant-baseline",
		"110": "empty-cells",
		"111": "fill",
		"112": "fill-opacity",
		"113": "fill-rule",
		"114": "filter",
		"115": "flex-basis",
		"116": "flex-direction",
		"117": "flex-grow",
		"118": "flex-shrink",
		"119": "flex-wrap",
		"120": "float",
		"121": "flood-color",
		"122": "flood-opacity",
		"123": "font-family",
		"124": "font-kerning",
		"125": "font-optical-sizing",
		"126": "font-palette",
		"127": "font-size",
		"128": "font-stretch",
		"129": "font-style",
		"130": "font-synthesis-small-caps",
		"131": "font-synthesis-style",
		"132": "font-synthesis-weight",
		"133": "font-variant",
		"134": "font-variant-alternates",
		"135": "font-variant-caps",
		"136": "font-variant-east-asian",
		"137": "font-variant-ligatures",
		"138": "font-variant-numeric",
		"139": "font-variant-position",
		"140": "font-weight",
		"141": "grid-auto-columns",
		"142": "grid-auto-flow",
		"143": "grid-auto-rows",
		"144": "grid-column-end",
		"145": "grid-column-start",
		"146": "grid-row-end",
		"147": "grid-row-start",
		"148": "grid-template-areas",
		"149": "grid-template-columns",
		"150": "grid-template-rows",
		"151": "height",
		"152": "hyphenate-character",
		"153": "hyphenate-limit-chars",
		"154": "hyphens",
		"155": "image-orientation",
		"156": "image-rendering",
		"157": "initial-letter",
		"158": "inline-size",
		"159": "inset-block-end",
		"160": "inset-block-start",
		"161": "inset-inline-end",
		"162": "inset-inline-start",
		"163": "isolation",
		"164": "justify-content",
		"165": "justify-items",
		"166": "justify-self",
		"167": "left",
		"168": "letter-spacing",
		"169": "lighting-color",
		"170": "line-break",
		"171": "line-height",
		"172": "list-style-image",
		"173": "list-style-position",
		"174": "list-style-type",
		"175": "margin-block-end",
		"176": "margin-block-start",
		"177": "margin-bottom",
		"178": "margin-inline-end",
		"179": "margin-inline-start",
		"180": "margin-left",
		"181": "margin-right",
		"182": "margin-top",
		"183": "marker-end",
		"184": "marker-mid",
		"185": "marker-start",
		"186": "mask-type",
		"187": "math-depth",
		"188": "math-shift",
		"189": "math-style",
		"190": "max-block-size",
		"191": "max-height",
		"192": "max-inline-size",
		"193": "max-width",
		"194": "min-block-size",
		"195": "min-height",
		"196": "min-inline-size",
		"197": "min-width",
		"198": "mix-blend-mode",
		"199": "object-fit",
		"200": "object-position",
		"201": "object-view-box",
		"202": "offset-anchor",
		"203": "offset-distance",
		"204": "offset-path",
		"205": "offset-position",
		"206": "offset-rotate",
		"207": "opacity",
		"208": "order",
		"209": "orphans",
		"210": "outline-color",
		"211": "outline-offset",
		"212": "outline-style",
		"213": "outline-width",
		"214": "overflow-anchor",
		"215": "overflow-clip-margin",
		"216": "overflow-wrap",
		"217": "overflow-x",
		"218": "overflow-y",
		"219": "overlay",
		"220": "overscroll-behavior-block",
		"221": "overscroll-behavior-inline",
		"222": "padding-block-end",
		"223": "padding-block-start",
		"224": "padding-bottom",
		"225": "padding-inline-end",
		"226": "padding-inline-start",
		"227": "padding-left",
		"228": "padding-right",
		"229": "padding-top",
		"230": "paint-order",
		"231": "perspective",
		"232": "perspective-origin",
		"233": "pointer-events",
		"234": "position",
		"235": "r",
		"236": "resize",
		"237": "right",
		"238": "rotate",
		"239": "row-gap",
		"240": "ruby-position",
		"241": "rx",
		"242": "ry",
		"243": "scale",
		"244": "scroll-behavior",
		"245": "scroll-margin-block-end",
		"246": "scroll-margin-block-start",
		"247": "scroll-margin-inline-end",
		"248": "scroll-margin-inline-start",
		"249": "scroll-padding-block-end",
		"250": "scroll-padding-block-start",
		"251": "scroll-padding-inline-end",
		"252": "scroll-padding-inline-start",
		"253": "scroll-timeline-axis",
		"254": "scroll-timeline-name",
		"255": "scrollbar-gutter",
		"256": "shape-image-threshold",
		"257": "shape-margin",
		"258": "shape-outside",
		"259": "shape-rendering",
		"260": "speak",
		"261": "stop-color",
		"262": "stop-opacity",
		"263": "stroke",
		"264": "stroke-dasharray",
		"265": "stroke-dashoffset",
		"266": "stroke-linecap",
		"267": "stroke-linejoin",
		"268": "stroke-miterlimit",
		"269": "stroke-opacity",
		"270": "stroke-width",
		"271": "tab-size",
		"272": "table-layout",
		"273": "text-align",
		"274": "text-align-last",
		"275": "text-anchor",
		"276": "text-decoration",
		"277": "text-decoration-color",
		"278": "text-decoration-line",
		"279": "text-decoration-skip-ink",
		"280": "text-decoration-style",
		"281": "text-emphasis-color",
		"282": "text-emphasis-position",
		"283": "text-emphasis-style",
		"284": "text-indent",
		"285": "text-overflow",
		"286": "text-rendering",
		"287": "text-shadow",
		"288": "text-size-adjust",
		"289": "text-transform",
		"290": "text-underline-position",
		"291": "text-wrap",
		"292": "timeline-scope",
		"293": "top",
		"294": "touch-action",
		"295": "transform",
		"296": "transform-origin",
		"297": "transform-style",
		"298": "transition-behavior",
		"299": "transition-delay",
		"300": "transition-duration",
		"301": "transition-property",
		"302": "transition-timing-function",
		"303": "translate",
		"304": "unicode-bidi",
		"305": "user-select",
		"306": "vector-effect",
		"307": "vertical-align",
		"308": "view-timeline-axis",
		"309": "view-timeline-inset",
		"310": "view-timeline-name",
		"311": "view-transition-name",
		"312": "visibility",
		"313": "white-space-collapse",
		"314": "widows",
		"315": "width",
		"316": "will-change",
		"317": "word-break",
		"318": "word-spacing",
		"319": "writing-mode",
		"320": "x",
		"321": "y",
		"322": "z-index",
		"323": "zoom",
		"324": "-webkit-border-horizontal-spacing",
		"325": "-webkit-border-image",
		"326": "-webkit-border-vertical-spacing",
		"327": "-webkit-box-align",
		"328": "-webkit-box-decoration-break",
		"329": "-webkit-box-direction",
		"330": "-webkit-box-flex",
		"331": "-webkit-box-ordinal-group",
		"332": "-webkit-box-orient",
		"333": "-webkit-box-pack",
		"334": "-webkit-box-reflect",
		"335": "-webkit-font-smoothing",
		"336": "-webkit-line-break",
		"337": "-webkit-line-clamp",
		"338": "-webkit-locale",
		"339": "-webkit-mask-box-image",
		"340": "-webkit-mask-box-image-outset",
		"341": "-webkit-mask-box-image-repeat",
		"342": "-webkit-mask-box-image-slice",
		"343": "-webkit-mask-box-image-source",
		"344": "-webkit-mask-box-image-width",
		"345": "-webkit-mask-clip",
		"346": "-webkit-mask-composite",
		"347": "-webkit-mask-image",
		"348": "-webkit-mask-origin",
		"349": "-webkit-mask-position",
		"350": "-webkit-mask-repeat",
		"351": "-webkit-mask-size",
		"352": "-webkit-print-color-adjust",
		"353": "-webkit-rtl-ordering",
		"354": "-webkit-tap-highlight-color",
		"355": "-webkit-text-combine",
		"356": "-webkit-text-decorations-in-effect",
		"357": "-webkit-text-fill-color",
		"358": "-webkit-text-orientation",
		"359": "-webkit-text-security",
		"360": "-webkit-text-stroke-color",
		"361": "-webkit-text-stroke-width",
		"362": "-webkit-user-drag",
		"363": "-webkit-user-modify",
		"364": "-webkit-writing-mode",
		"accentColor": "auto",
		"additiveSymbols": "",
		"alignContent": "normal",
		"alignItems": "normal",
		"alignSelf": "auto",
		"alignmentBaseline": "auto",
		"all": "",
		"animation": "none 0s ease 0s 1 normal none running",
		"animationComposition": "replace",
		"animationDelay": "0s",
		"animationDirection": "normal",
		"animationDuration": "0s",
		"animationFillMode": "none",
		"animationIterationCount": "1",
		"animationName": "none",
		"animationPlayState": "running",
		"animationRange": "normal",
		"animationRangeEnd": "normal",
		"animationRangeStart": "normal",
		"animationTimeline": "auto",
		"animationTimingFunction": "ease",
		"appRegion": "none",
		"appearance": "none",
		"ascentOverride": "",
		"aspectRatio": "auto",
		"backdropFilter": "none",
		"backfaceVisibility": "visible",
		"background": "rgba(0, 0, 0, 0) none repeat scroll 0% 0% / auto padding-box border-box",
		"backgroundAttachment": "scroll",
		"backgroundBlendMode": "normal",
		"backgroundClip": "border-box",
		"backgroundColor": "rgba(0, 0, 0, 0)",
		"backgroundImage": "none",
		"backgroundOrigin": "padding-box",
		"backgroundPosition": "0% 0%",
		"backgroundPositionX": "0%",
		"backgroundPositionY": "0%",
		"backgroundRepeat": "repeat",
		"backgroundRepeatX": "repeat",
		"backgroundRepeatY": "repeat",
		"backgroundSize": "auto",
		"basePalette": "",
		"baselineShift": "0px",
		"baselineSource": "auto",
		"blockSize": "8px",
		"border": "0px none rgb(0, 0, 0)",
		"borderBlock": "0px none rgb(0, 0, 0)",
		"borderBlockColor": "rgb(0, 0, 0)",
		"borderBlockEnd": "0px none rgb(0, 0, 0)",
		"borderBlockEndColor": "rgb(0, 0, 0)",
		"borderBlockEndStyle": "none",
		"borderBlockEndWidth": "0px",
		"borderBlockStart": "0px none rgb(0, 0, 0)",
		"borderBlockStartColor": "rgb(0, 0, 0)",
		"borderBlockStartStyle": "none",
		"borderBlockStartWidth": "0px",
		"borderBlockStyle": "none",
		"borderBlockWidth": "0px",
		"borderBottom": "0px none rgb(0, 0, 0)",
		"borderBottomColor": "rgb(0, 0, 0)",
		"borderBottomLeftRadius": "0px",
		"borderBottomRightRadius": "0px",
		"borderBottomStyle": "none",
		"borderBottomWidth": "0px",
		"borderCollapse": "separate",
		"borderColor": "rgb(0, 0, 0)",
		"borderEndEndRadius": "0px",
		"borderEndStartRadius": "0px",
		"borderImage": "none",
		"borderImageOutset": "0",
		"borderImageRepeat": "stretch",
		"borderImageSlice": "100%",
		"borderImageSource": "none",
		"borderImageWidth": "1",
		"borderInline": "0px none rgb(0, 0, 0)",
		"borderInlineColor": "rgb(0, 0, 0)",
		"borderInlineEnd": "0px none rgb(0, 0, 0)",
		"borderInlineEndColor": "rgb(0, 0, 0)",
		"borderInlineEndStyle": "none",
		"borderInlineEndWidth": "0px",
		"borderInlineStart": "0px none rgb(0, 0, 0)",
		"borderInlineStartColor": "rgb(0, 0, 0)",
		"borderInlineStartStyle": "none",
		"borderInlineStartWidth": "0px",
		"borderInlineStyle": "none",
		"borderInlineWidth": "0px",
		"borderLeft": "0px none rgb(0, 0, 0)",
		"borderLeftColor": "rgb(0, 0, 0)",
		"borderLeftStyle": "none",
		"borderLeftWidth": "0px",
		"borderRadius": "0px",
		"borderRight": "0px none rgb(0, 0, 0)",
		"borderRightColor": "rgb(0, 0, 0)",
		"borderRightStyle": "none",
		"borderRightWidth": "0px",
		"borderSpacing": "0px 0px",
		"borderStartEndRadius": "0px",
		"borderStartStartRadius": "0px",
		"borderStyle": "none",
		"borderTop": "0px none rgb(0, 0, 0)",
		"borderTopColor": "rgb(0, 0, 0)",
		"borderTopLeftRadius": "0px",
		"borderTopRightRadius": "0px",
		"borderTopStyle": "none",
		"borderTopWidth": "0px",
		"borderWidth": "0px",
		"bottom": "auto",
		"boxShadow": "none",
		"boxSizing": "content-box",
		"breakAfter": "auto",
		"breakBefore": "auto",
		"breakInside": "auto",
		"bufferedRendering": "auto",
		"captionSide": "top",
		"caretColor": "rgb(0, 0, 0)",
		"clear": "none",
		"clip": "auto",
		"clipPath": "none",
		"clipRule": "nonzero",
		"color": "rgb(0, 0, 0)",
		"colorInterpolation": "srgb",
		"colorInterpolationFilters": "linearrgb",
		"colorRendering": "auto",
		"colorScheme": "normal",
		"columnCount": "auto",
		"columnFill": "balance",
		"columnGap": "normal",
		"columnRule": "0px none rgb(0, 0, 0)",
		"columnRuleColor": "rgb(0, 0, 0)",
		"columnRuleStyle": "none",
		"columnRuleWidth": "0px",
		"columnSpan": "none",
		"columnWidth": "auto",
		"columns": "auto auto",
		"contain": "none",
		"containIntrinsicBlockSize": "none",
		"containIntrinsicHeight": "none",
		"containIntrinsicInlineSize": "none",
		"containIntrinsicSize": "none",
		"containIntrinsicWidth": "none",
		"container": "none",
		"containerName": "none",
		"containerType": "normal",
		"content": "normal",
		"contentVisibility": "visible",
		"counterIncrement": "none",
		"counterReset": "none",
		"counterSet": "none",
		"cursor": "auto",
		"cx": "0px",
		"cy": "0px",
		"d": "none",
		"descentOverride": "",
		"direction": "ltr",
		"display": "block",
		"dominantBaseline": "auto",
		"emptyCells": "show",
		"fallback": "",
		"fill": "rgb(0, 0, 0)",
		"fillOpacity": "1",
		"fillRule": "nonzero",
		"filter": "none",
		"flex": "0 1 auto",
		"flexBasis": "auto",
		"flexDirection": "row",
		"flexFlow": "row nowrap",
		"flexGrow": "0",
		"flexShrink": "1",
		"flexWrap": "nowrap",
		"float": "none",
		"floodColor": "rgb(0, 0, 0)",
		"floodOpacity": "1",
		"font": "16px Times",
		"fontDisplay": "",
		"fontFamily": "Times",
		"fontFeatureSettings": "normal",
		"fontKerning": "auto",
		"fontOpticalSizing": "auto",
		"fontPalette": "normal",
		"fontSize": "16px",
		"fontStretch": "100%",
		"fontStyle": "normal",
		"fontSynthesis": "weight style small-caps",
		"fontSynthesisSmallCaps": "auto",
		"fontSynthesisStyle": "auto",
		"fontSynthesisWeight": "auto",
		"fontVariant": "normal",
		"fontVariantAlternates": "normal",
		"fontVariantCaps": "normal",
		"fontVariantEastAsian": "normal",
		"fontVariantLigatures": "normal",
		"fontVariantNumeric": "normal",
		"fontVariantPosition": "normal",
		"fontVariationSettings": "normal",
		"fontWeight": "400",
		"forcedColorAdjust": "auto",
		"gap": "normal",
		"grid": "none / none / none / row / auto / auto",
		"gridArea": "auto / auto / auto / auto",
		"gridAutoColumns": "auto",
		"gridAutoFlow": "row",
		"gridAutoRows": "auto",
		"gridColumn": "auto / auto",
		"gridColumnEnd": "auto",
		"gridColumnGap": "normal",
		"gridColumnStart": "auto",
		"gridGap": "normal normal",
		"gridRow": "auto / auto",
		"gridRowEnd": "auto",
		"gridRowGap": "normal",
		"gridRowStart": "auto",
		"gridTemplate": "none / none / none",
		"gridTemplateAreas": "none",
		"gridTemplateColumns": "none",
		"gridTemplateRows": "none",
		"height": "8px",
		"hyphenateCharacter": "auto",
		"hyphenateLimitChars": "auto",
		"hyphens": "manual",
		"imageOrientation": "from-image",
		"imageRendering": "auto",
		"inherits": "",
		"initialLetter": "normal",
		"initialValue": "",
		"inlineSize": "0px",
		"inset": "auto",
		"insetBlock": "auto",
		"insetBlockEnd": "auto",
		"insetBlockStart": "auto",
		"insetInline": "auto",
		"insetInlineEnd": "auto",
		"insetInlineStart": "auto",
		"isolation": "auto",
		"justifyContent": "normal",
		"justifyItems": "normal",
		"justifySelf": "auto",
		"left": "auto",
		"letterSpacing": "normal",
		"lightingColor": "rgb(255, 255, 255)",
		"lineBreak": "auto",
		"lineGapOverride": "",
		"lineHeight": "normal",
		"listStyle": "outside none disc",
		"listStyleImage": "none",
		"listStylePosition": "outside",
		"listStyleType": "disc",
		"margin": "0px",
		"marginBlock": "0px",
		"marginBlockEnd": "0px",
		"marginBlockStart": "0px",
		"marginBottom": "0px",
		"marginInline": "0px",
		"marginInlineEnd": "0px",
		"marginInlineStart": "0px",
		"marginLeft": "0px",
		"marginRight": "0px",
		"marginTop": "0px",
		"marker": "none",
		"markerEnd": "none",
		"markerMid": "none",
		"markerStart": "none",
		"mask": "none",
		"maskType": "luminance",
		"mathDepth": "0",
		"mathShift": "normal",
		"mathStyle": "normal",
		"maxBlockSize": "none",
		"maxHeight": "none",
		"maxInlineSize": "none",
		"maxWidth": "none",
		"minBlockSize": "0px",
		"minHeight": "0px",
		"minInlineSize": "0px",
		"minWidth": "0px",
		"mixBlendMode": "normal",
		"negative": "",
		"objectFit": "fill",
		"objectPosition": "50% 50%",
		"objectViewBox": "none",
		"offset": "none 0px auto 0deg",
		"offsetAnchor": "auto",
		"offsetDistance": "0px",
		"offsetPath": "none",
		"offsetPosition": "auto",
		"offsetRotate": "auto 0deg",
		"opacity": "1",
		"order": "0",
		"orphans": "2",
		"outline": "rgb(0, 0, 0) none 0px",
		"outlineColor": "rgb(0, 0, 0)",
		"outlineOffset": "0px",
		"outlineStyle": "none",
		"outlineWidth": "0px",
		"overflow": "visible",
		"overflowAnchor": "auto",
		"overflowClipMargin": "0px",
		"overflowWrap": "normal",
		"overflowX": "visible",
		"overflowY": "visible",
		"overlay": "none",
		"overrideColors": "",
		"overscrollBehavior": "auto",
		"overscrollBehaviorBlock": "auto",
		"overscrollBehaviorInline": "auto",
		"overscrollBehaviorX": "auto",
		"overscrollBehaviorY": "auto",
		"pad": "",
		"padding": "0px",
		"paddingBlock": "0px",
		"paddingBlockEnd": "0px",
		"paddingBlockStart": "0px",
		"paddingBottom": "0px",
		"paddingInline": "0px",
		"paddingInlineEnd": "0px",
		"paddingInlineStart": "0px",
		"paddingLeft": "0px",
		"paddingRight": "0px",
		"paddingTop": "0px",
		"page": "auto",
		"pageBreakAfter": "auto",
		"pageBreakBefore": "auto",
		"pageBreakInside": "auto",
		"pageOrientation": "",
		"paintOrder": "normal",
		"perspective": "none",
		"perspectiveOrigin": "0px 4px",
		"placeContent": "normal",
		"placeItems": "normal",
		"placeSelf": "auto",
		"pointerEvents": "auto",
		"position": "static",
		"prefix": "",
		"quotes": "auto",
		"r": "0px",
		"range": "",
		"resize": "none",
		"right": "auto",
		"rotate": "none",
		"rowGap": "normal",
		"rubyPosition": "over",
		"rx": "auto",
		"ry": "auto",
		"scale": "none",
		"scrollBehavior": "auto",
		"scrollMargin": "0px",
		"scrollMarginBlock": "0px",
		"scrollMarginBlockEnd": "0px",
		"scrollMarginBlockStart": "0px",
		"scrollMarginBottom": "0px",
		"scrollMarginInline": "0px",
		"scrollMarginInlineEnd": "0px",
		"scrollMarginInlineStart": "0px",
		"scrollMarginLeft": "0px",
		"scrollMarginRight": "0px",
		"scrollMarginTop": "0px",
		"scrollPadding": "auto",
		"scrollPaddingBlock": "auto",
		"scrollPaddingBlockEnd": "auto",
		"scrollPaddingBlockStart": "auto",
		"scrollPaddingBottom": "auto",
		"scrollPaddingInline": "auto",
		"scrollPaddingInlineEnd": "auto",
		"scrollPaddingInlineStart": "auto",
		"scrollPaddingLeft": "auto",
		"scrollPaddingRight": "auto",
		"scrollPaddingTop": "auto",
		"scrollSnapAlign": "none",
		"scrollSnapStop": "normal",
		"scrollSnapType": "none",
		"scrollTimeline": "none",
		"scrollTimelineAxis": "block",
		"scrollTimelineName": "none",
		"scrollbarGutter": "auto",
		"shapeImageThreshold": "0",
		"shapeMargin": "0px",
		"shapeOutside": "none",
		"shapeRendering": "auto",
		"size": "",
		"sizeAdjust": "",
		"speak": "normal",
		"speakAs": "",
		"src": "",
		"stopColor": "rgb(0, 0, 0)",
		"stopOpacity": "1",
		"stroke": "none",
		"strokeDasharray": "none",
		"strokeDashoffset": "0px",
		"strokeLinecap": "butt",
		"strokeLinejoin": "miter",
		"strokeMiterlimit": "4",
		"strokeOpacity": "1",
		"strokeWidth": "1px",
		"suffix": "",
		"symbols": "",
		"syntax": "",
		"system": "",
		"tabSize": "8",
		"tableLayout": "auto",
		"textAlign": "start",
		"textAlignLast": "auto",
		"textAnchor": "start",
		"textCombineUpright": "none",
		"textDecoration": "none solid rgb(0, 0, 0)",
		"textDecorationColor": "rgb(0, 0, 0)",
		"textDecorationLine": "none",
		"textDecorationSkipInk": "auto",
		"textDecorationStyle": "solid",
		"textDecorationThickness": "auto",
		"textEmphasis": "none rgb(0, 0, 0)",
		"textEmphasisColor": "rgb(0, 0, 0)",
		"textEmphasisPosition": "over",
		"textEmphasisStyle": "none",
		"textIndent": "0px",
		"textOrientation": "mixed",
		"textOverflow": "clip",
		"textRendering": "auto",
		"textShadow": "none",
		"textSizeAdjust": "auto",
		"textTransform": "none",
		"textUnderlineOffset": "auto",
		"textUnderlinePosition": "auto",
		"textWrap": "wrap",
		"timelineScope": "none",
		"top": "auto",
		"touchAction": "auto",
		"transform": "none",
		"transformBox": "view-box",
		"transformOrigin": "0px 4px",
		"transformStyle": "flat",
		"transition": "all 0s ease 0s",
		"transitionBehavior": "normal",
		"transitionDelay": "0s",
		"transitionDuration": "0s",
		"transitionProperty": "all",
		"transitionTimingFunction": "ease",
		"translate": "none",
		"unicodeBidi": "normal",
		"unicodeRange": "",
		"userSelect": "auto",
		"vectorEffect": "none",
		"verticalAlign": "baseline",
		"viewTimeline": "none",
		"viewTimelineAxis": "block",
		"viewTimelineInset": "0px",
		"viewTimelineName": "none",
		"viewTransitionName": "root",
		"visibility": "visible",
		"webkitAlignContent": "normal",
		"webkitAlignItems": "normal",
		"webkitAlignSelf": "auto",
		"webkitAnimation": "none 0s ease 0s 1 normal none running",
		"webkitAnimationDelay": "0s",
		"webkitAnimationDirection": "normal",
		"webkitAnimationDuration": "0s",
		"webkitAnimationFillMode": "none",
		"webkitAnimationIterationCount": "1",
		"webkitAnimationName": "none",
		"webkitAnimationPlayState": "running",
		"webkitAnimationTimingFunction": "ease",
		"webkitAppRegion": "none",
		"webkitAppearance": "none",
		"webkitBackfaceVisibility": "visible",
		"webkitBackgroundClip": "border-box",
		"webkitBackgroundOrigin": "padding-box",
		"webkitBackgroundSize": "auto",
		"webkitBorderAfter": "0px none rgb(0, 0, 0)",
		"webkitBorderAfterColor": "rgb(0, 0, 0)",
		"webkitBorderAfterStyle": "none",
		"webkitBorderAfterWidth": "0px",
		"webkitBorderBefore": "0px none rgb(0, 0, 0)",
		"webkitBorderBeforeColor": "rgb(0, 0, 0)",
		"webkitBorderBeforeStyle": "none",
		"webkitBorderBeforeWidth": "0px",
		"webkitBorderBottomLeftRadius": "0px",
		"webkitBorderBottomRightRadius": "0px",
		"webkitBorderEnd": "0px none rgb(0, 0, 0)",
		"webkitBorderEndColor": "rgb(0, 0, 0)",
		"webkitBorderEndStyle": "none",
		"webkitBorderEndWidth": "0px",
		"webkitBorderHorizontalSpacing": "0px",
		"webkitBorderImage": "none",
		"webkitBorderRadius": "0px",
		"webkitBorderStart": "0px none rgb(0, 0, 0)",
		"webkitBorderStartColor": "rgb(0, 0, 0)",
		"webkitBorderStartStyle": "none",
		"webkitBorderStartWidth": "0px",
		"webkitBorderTopLeftRadius": "0px",
		"webkitBorderTopRightRadius": "0px",
		"webkitBorderVerticalSpacing": "0px",
		"webkitBoxAlign": "stretch",
		"webkitBoxDecorationBreak": "slice",
		"webkitBoxDirection": "normal",
		"webkitBoxFlex": "0",
		"webkitBoxOrdinalGroup": "1",
		"webkitBoxOrient": "horizontal",
		"webkitBoxPack": "start",
		"webkitBoxReflect": "none",
		"webkitBoxShadow": "none",
		"webkitBoxSizing": "content-box",
		"webkitClipPath": "none",
		"webkitColumnBreakAfter": "auto",
		"webkitColumnBreakBefore": "auto",
		"webkitColumnBreakInside": "auto",
		"webkitColumnCount": "auto",
		"webkitColumnGap": "normal",
		"webkitColumnRule": "0px none rgb(0, 0, 0)",
		"webkitColumnRuleColor": "rgb(0, 0, 0)",
		"webkitColumnRuleStyle": "none",
		"webkitColumnRuleWidth": "0px",
		"webkitColumnSpan": "none",
		"webkitColumnWidth": "auto",
		"webkitColumns": "auto auto",
		"webkitFilter": "none",
		"webkitFlex": "0 1 auto",
		"webkitFlexBasis": "auto",
		"webkitFlexDirection": "row",
		"webkitFlexFlow": "row nowrap",
		"webkitFlexGrow": "0",
		"webkitFlexShrink": "1",
		"webkitFlexWrap": "nowrap",
		"webkitFontFeatureSettings": "normal",
		"webkitFontSmoothing": "auto",
		"webkitHyphenateCharacter": "auto",
		"webkitJustifyContent": "normal",
		"webkitLineBreak": "auto",
		"webkitLineClamp": "none",
		"webkitLocale": "auto",
		"webkitLogicalHeight": "8px",
		"webkitLogicalWidth": "0px",
		"webkitMarginAfter": "0px",
		"webkitMarginBefore": "0px",
		"webkitMarginEnd": "0px",
		"webkitMarginStart": "0px",
		"webkitMask": "",
		"webkitMaskBoxImage": "none",
		"webkitMaskBoxImageOutset": "0",
		"webkitMaskBoxImageRepeat": "stretch",
		"webkitMaskBoxImageSlice": "0 fill",
		"webkitMaskBoxImageSource": "none",
		"webkitMaskBoxImageWidth": "auto",
		"webkitMaskClip": "border-box",
		"webkitMaskComposite": "source-over",
		"webkitMaskImage": "none",
		"webkitMaskOrigin": "border-box",
		"webkitMaskPosition": "0% 0%",
		"webkitMaskPositionX": "0%",
		"webkitMaskPositionY": "0%",
		"webkitMaskRepeat": "repeat",
		"webkitMaskRepeatX": "",
		"webkitMaskRepeatY": "",
		"webkitMaskSize": "auto",
		"webkitMaxLogicalHeight": "none",
		"webkitMaxLogicalWidth": "none",
		"webkitMinLogicalHeight": "0px",
		"webkitMinLogicalWidth": "0px",
		"webkitOpacity": "1",
		"webkitOrder": "0",
		"webkitPaddingAfter": "0px",
		"webkitPaddingBefore": "0px",
		"webkitPaddingEnd": "0px",
		"webkitPaddingStart": "0px",
		"webkitPerspective": "none",
		"webkitPerspectiveOrigin": "0px 4px",
		"webkitPerspectiveOriginX": "",
		"webkitPerspectiveOriginY": "",
		"webkitPrintColorAdjust": "economy",
		"webkitRtlOrdering": "logical",
		"webkitRubyPosition": "before",
		"webkitShapeImageThreshold": "0",
		"webkitShapeMargin": "0px",
		"webkitShapeOutside": "none",
		"webkitTapHighlightColor": "rgba(0, 0, 0, 0.18)",
		"webkitTextCombine": "none",
		"webkitTextDecorationsInEffect": "none",
		"webkitTextEmphasis": "none rgb(0, 0, 0)",
		"webkitTextEmphasisColor": "rgb(0, 0, 0)",
		"webkitTextEmphasisPosition": "over",
		"webkitTextEmphasisStyle": "none",
		"webkitTextFillColor": "rgb(0, 0, 0)",
		"webkitTextOrientation": "vertical-right",
		"webkitTextSecurity": "none",
		"webkitTextSizeAdjust": "auto",
		"webkitTextStroke": "",
		"webkitTextStrokeColor": "rgb(0, 0, 0)",
		"webkitTextStrokeWidth": "0px",
		"webkitTransform": "none",
		"webkitTransformOrigin": "0px 4px",
		"webkitTransformOriginX": "",
		"webkitTransformOriginY": "",
		"webkitTransformOriginZ": "",
		"webkitTransformStyle": "flat",
		"webkitTransition": "all 0s ease 0s",
		"webkitTransitionDelay": "0s",
		"webkitTransitionDuration": "0s",
		"webkitTransitionProperty": "all",
		"webkitTransitionTimingFunction": "ease",
		"webkitUserDrag": "auto",
		"webkitUserModify": "read-only",
		"webkitUserSelect": "auto",
		"webkitWritingMode": "horizontal-tb",
		"whiteSpace": "normal",
		"whiteSpaceCollapse": "collapse",
		"widows": "2",
		"width": "0px",
		"willChange": "auto",
		"wordBreak": "normal",
		"wordSpacing": "0px",
		"wordWrap": "normal",
		"writingMode": "horizontal-tb",
		"x": "0px",
		"y": "0px",
		"zIndex": "auto",
		"zoom": "1",
		"cssText": "",
		"length": 365,
		"cssFloat": "none"
	},
	"keyboard": {
		"0": "Digit0",
		"1": "Digit1",
		"2": "Digit2",
		"3": "Digit3",
		"4": "Digit4",
		"5": "Digit5",
		"6": "Digit6",
		"7": "Digit7",
		"8": "Digit8",
		"9": "Digit9",
		"e": "KeyE",
		"d": "KeyD",
		"u": "KeyU",
		"ß": "Minus",
		"h": "KeyH",
		"y": "KeyZ",
		"´": "Equal",
		"p": "KeyP",
		"ö": "Semicolon",
		"+": "BracketRight",
		"-": "Slash",
		"ü": "BracketLeft",
		"l": "KeyL",
		"w": "KeyW",
		"s": "KeyS",
		"o": "KeyO",
		".": "Period",
		"v": "KeyV",
		"<": "IntlRo",
		"g": "KeyG",
		"j": "KeyJ",
		"q": "KeyQ",
		"t": "KeyT",
		"z": "KeyY",
		"ä": "Quote",
		"^": "IntlYen",
		"#": "Backslash",
		"k": "KeyK",
		"f": "KeyF",
		"i": "KeyI",
		"r": "KeyR",
		"x": "KeyX",
		"a": "KeyA",
		"m": "KeyM",
		"n": "KeyN",
		"b": "KeyB",
		"c": "KeyC",
		",": "Comma"
	},
	"audioTypes": {
		"audio/3gpp": "",
		"audio/aac": "probably",
		"audio/flac": "probably",
		"audio/mpeg": "probably",
		"audio/mp4; codecs=\"ac-3\"": "",
		"audio/mp4; codecs=\"ec-3\"": "",
		"audio/ogg; codecs=\"flac\"": "probably",
		"audio/ogg; codecs=\"vorbis\"": "probably",
		"audio/ogg; codecs=\"opus\"": "probably",
		"audio/wav; codecs=\"1\"": "probably",
		"audio/webm; codecs=\"vorbis\"": "probably",
		"audio/webm; codecs=\"opus\"": "probably",
		"audio/mp4; codecs=\"mp4a_40_2\"": ""
	},
	"videoTypes": {
		"video/mp4; codecs=\"flac\"": "probably",
		"video/ogg; codecs=\"theora\"": "probably",
		"video/ogg; codecs=\"opus\"": "probably",
		"video/webm; codecs=\"vp9, opus\"": "probably",
		"video/webm; codecs=\"vp8, vorbis\" ": "probably"
	},
	"audioContext": {
		"baseLatency": 0.005333333333333333,
		"outputLatency": 0,
		"sinkId": "",
		"destination": {
			"maxChannelCount": 2,
			"context": {
				"baseLatency": 0.005333333333333333,
				"outputLatency": 0,
				"sinkId": "",
				"destination": {
					"maxChannelCount": 2,
					"numberOfInputs": 1,
					"numberOfOutputs": 0,
					"channelCount": 2,
					"channelCountMode": "explicit",
					"channelInterpretation": "speakers"
				},
				"currentTime": 0,
				"sampleRate": 48000,
				"listener": {},
				"state": "running",
				"audioWorklet": {}
			},
			"numberOfInputs": 1,
			"numberOfOutputs": 0,
			"channelCount": 2,
			"channelCountMode": "explicit",
			"channelInterpretation": "speakers"
		},
		"currentTime": 0,
		"sampleRate": 48000,
		"listener": {
			"positionX": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"positionY": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"positionZ": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"forwardX": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"forwardY": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"forwardZ": {
				"value": -1,
				"automationRate": "a-rate",
				"defaultValue": -1,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"upX": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"upY": {
				"value": 1,
				"automationRate": "a-rate",
				"defaultValue": 1,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			},
			"upZ": {
				"value": 0,
				"automationRate": "a-rate",
				"defaultValue": 0,
				"minValue": -3.4028234663852886e+38,
				"maxValue": 3.4028234663852886e+38
			}
		},
		"state": "running",
		"audioWorklet": {}
	},
	"webrtc": {
		"video": {
			"codecs": [
				{
					"clockRate": 90000,
					"mimeType": "video/VP8"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/rtx"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/VP9",
					"sdpFmtpLine": "profile-id=0"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/VP9",
					"sdpFmtpLine": "profile-id=2"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/VP9",
					"sdpFmtpLine": "profile-id=1"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/VP9",
					"sdpFmtpLine": "profile-id=3"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=4d001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=4d001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=f4001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=f4001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/AV1"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/AV1",
					"sdpFmtpLine": "profile=1"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=64001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/H264",
					"sdpFmtpLine": "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=64001f"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/red"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/ulpfec"
				},
				{
					"clockRate": 90000,
					"mimeType": "video/flexfec-03",
					"sdpFmtpLine": "repair-window=10000000"
				}
			],
			"headerExtensions": [
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:toffset"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time"
				},
				{
					"direction": "sendrecv",
					"uri": "urn:3gpp:video-orientation"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/playout-delay"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/video-content-type"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/video-timing"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/color-space"
				},
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:sdes:mid"
				},
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id"
				},
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id"
				}
			]
		},
		"audio": {
			"codecs": [
				{
					"channels": 2,
					"clockRate": 48000,
					"mimeType": "audio/opus",
					"sdpFmtpLine": "minptime=10;useinbandfec=1"
				},
				{
					"channels": 2,
					"clockRate": 48000,
					"mimeType": "audio/red",
					"sdpFmtpLine": "111/111"
				},
				{
					"channels": 1,
					"clockRate": 8000,
					"mimeType": "audio/G722"
				},
				{
					"channels": 1,
					"clockRate": 8000,
					"mimeType": "audio/PCMU"
				},
				{
					"channels": 1,
					"clockRate": 8000,
					"mimeType": "audio/PCMA"
				},
				{
					"channels": 1,
					"clockRate": 8000,
					"mimeType": "audio/CN"
				},
				{
					"channels": 1,
					"clockRate": 48000,
					"mimeType": "audio/telephone-event"
				},
				{
					"channels": 1,
					"clockRate": 8000,
					"mimeType": "audio/telephone-event"
				}
			],
			"headerExtensions": [
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:ssrc-audio-level"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time"
				},
				{
					"direction": "sendrecv",
					"uri": "http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01"
				},
				{
					"direction": "sendrecv",
					"uri": "urn:ietf:params:rtp-hdrext:sdes:mid"
				}
			]
		}
	},
	"webgpu": {
		"features": {
			"size": 8
		},
		"limits": {
			"maxTextureDimension1D": 8192,
			"maxTextureDimension2D": 8192,
			"maxTextureDimension3D": 2048,
			"maxTextureArrayLayers": 256,
			"maxBindGroups": 4,
			"maxBindingsPerBindGroup": 1000,
			"maxDynamicUniformBuffersPerPipelineLayout": 10,
			"maxDynamicStorageBuffersPerPipelineLayout": 8,
			"maxSampledTexturesPerShaderStage": 16,
			"maxSamplersPerShaderStage": 16,
			"maxStorageBuffersPerShaderStage": 8,
			"maxStorageTexturesPerShaderStage": 8,
			"maxUniformBuffersPerShaderStage": 12,
			"maxUniformBufferBindingSize": 65536,
			"maxStorageBufferBindingSize": 4294967292,
			"minUniformBufferOffsetAlignment": 256,
			"minStorageBufferOffsetAlignment": 256,
			"maxVertexBuffers": 8,
			"maxBufferSize": 4294967296,
			"maxVertexAttributes": 16,
			"maxVertexBufferArrayStride": 2048,
			"maxInterStageShaderComponents": 60,
			"maxInterStageShaderVariables": 16,
			"maxColorAttachments": 8,
			"maxColorAttachmentBytesPerSample": 32,
			"maxComputeWorkgroupStorageSize": 32768,
			"maxComputeInvocationsPerWorkgroup": 1024,
			"maxComputeWorkgroupSizeX": 1024,
			"maxComputeWorkgroupSizeY": 1024,
			"maxComputeWorkgroupSizeZ": 64,
			"maxComputeWorkgroupsPerDimension": 65535
		},
		"isFallbackAdapter": false,
		"vendor": "apple",
		"architecture": "common-3",
		"device": "",
		"description": ""
	},
	"mediaDevices": [
		{
			"deviceId": "",
			"kind": "audioinput",
			"label": "",
			"groupId": ""
		},
		{
			"deviceId": "",
			"kind": "videoinput",
			"label": "",
			"groupId": ""
		},
		{
			"deviceId": "",
			"kind": "audiooutput",
			"label": "",
			"groupId": ""
		}
	],
	"is_bot": false,
	"status": "pass",
	"stack_worker": [
		6359,
		5299,
		39.99245283018868
	],
	"timing_worker": 0.09999999403953552,
	"gl": {
		"BLEND_EQUATION": [
			32774,
			32777
		],
		"BLEND_EQUATION_RGB": [
			32774,
			32777
		],
		"BLEND_EQUATION_ALPHA": [
			32774,
			34877
		],
		"BLEND_DST_RGB": [
			0,
			32968
		],
		"BLEND_SRC_RGB": [
			1,
			32969
		],
		"BLEND_DST_ALPHA": [
			0,
			32970
		],
		"BLEND_SRC_ALPHA": [
			1,
			32971
		],
		"BLEND_COLOR": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			32773
		],
		"CULL_FACE": [
			false,
			2884
		],
		"BLEND": [
			false,
			3042
		],
		"DITHER": [
			true,
			3024
		],
		"STENCIL_TEST": [
			false,
			2960
		],
		"DEPTH_TEST": [
			false,
			2929
		],
		"SCISSOR_TEST": [
			false,
			3089
		],
		"POLYGON_OFFSET_FILL": [
			false,
			32823
		],
		"SAMPLE_ALPHA_TO_COVERAGE": [
			false,
			32926
		],
		"SAMPLE_COVERAGE": [
			false,
			32928
		],
		"LINE_WIDTH": [
			1,
			2849
		],
		"ALIASED_POINT_SIZE_RANGE": [
			{
				"0": 1,
				"1": 64
			},
			33901
		],
		"ALIASED_LINE_WIDTH_RANGE": [
			{
				"0": 1,
				"1": 1
			},
			33902
		],
		"CULL_FACE_MODE": [
			1029,
			2885
		],
		"FRONT_FACE": [
			2305,
			2886
		],
		"DEPTH_RANGE": [
			{
				"0": 0,
				"1": 1
			},
			2928
		],
		"DEPTH_WRITEMASK": [
			true,
			2930
		],
		"DEPTH_CLEAR_VALUE": [
			1,
			2931
		],
		"DEPTH_FUNC": [
			513,
			2932
		],
		"STENCIL_CLEAR_VALUE": [
			0,
			2961
		],
		"STENCIL_FUNC": [
			519,
			2962
		],
		"STENCIL_FAIL": [
			7680,
			2964
		],
		"STENCIL_PASS_DEPTH_FAIL": [
			7680,
			2965
		],
		"STENCIL_PASS_DEPTH_PASS": [
			7680,
			2966
		],
		"STENCIL_REF": [
			0,
			2967
		],
		"STENCIL_VALUE_MASK": [
			2147483647,
			2963
		],
		"STENCIL_WRITEMASK": [
			2147483647,
			2968
		],
		"STENCIL_BACK_FUNC": [
			519,
			34816
		],
		"STENCIL_BACK_FAIL": [
			7680,
			34817
		],
		"STENCIL_BACK_PASS_DEPTH_FAIL": [
			7680,
			34818
		],
		"STENCIL_BACK_PASS_DEPTH_PASS": [
			7680,
			34819
		],
		"STENCIL_BACK_REF": [
			0,
			36003
		],
		"STENCIL_BACK_VALUE_MASK": [
			2147483647,
			36004
		],
		"STENCIL_BACK_WRITEMASK": [
			2147483647,
			36005
		],
		"VIEWPORT": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			2978
		],
		"SCISSOR_BOX": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			3088
		],
		"COLOR_CLEAR_VALUE": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			3106
		],
		"COLOR_WRITEMASK": [
			[
				true,
				true,
				true,
				true
			],
			3107
		],
		"UNPACK_ALIGNMENT": [
			4,
			3317
		],
		"PACK_ALIGNMENT": [
			4,
			3333
		],
		"MAX_TEXTURE_SIZE": [
			16384,
			3379
		],
		"MAX_VIEWPORT_DIMS": [
			{
				"0": 16384,
				"1": 16384
			},
			3386
		],
		"SUBPIXEL_BITS": [
			4,
			3408
		],
		"RED_BITS": [
			8,
			3410
		],
		"GREEN_BITS": [
			8,
			3411
		],
		"BLUE_BITS": [
			8,
			3412
		],
		"ALPHA_BITS": [
			8,
			3413
		],
		"DEPTH_BITS": [
			24,
			3414
		],
		"STENCIL_BITS": [
			0,
			3415
		],
		"POLYGON_OFFSET_UNITS": [
			0,
			10752
		],
		"POLYGON_OFFSET_FACTOR": [
			0,
			32824
		],
		"SAMPLE_BUFFERS": [
			1,
			32936
		],
		"SAMPLES": [
			4,
			32937
		],
		"SAMPLE_COVERAGE_VALUE": [
			1,
			32938
		],
		"SAMPLE_COVERAGE_INVERT": [
			false,
			32939
		],
		"COMPRESSED_TEXTURE_FORMATS": [
			{},
			34467
		],
		"GENERATE_MIPMAP_HINT": [
			4352,
			33170
		],
		"MAX_VERTEX_ATTRIBS": [
			16,
			34921
		],
		"MAX_VERTEX_UNIFORM_VECTORS": [
			1024,
			36347
		],
		"MAX_VARYING_VECTORS": [
			31,
			36348
		],
		"MAX_COMBINED_TEXTURE_IMAGE_UNITS": [
			32,
			35661
		],
		"MAX_VERTEX_TEXTURE_IMAGE_UNITS": [
			16,
			35660
		],
		"MAX_TEXTURE_IMAGE_UNITS": [
			16,
			34930
		],
		"MAX_FRAGMENT_UNIFORM_VECTORS": [
			1024,
			36349
		],
		"SHADING_LANGUAGE_VERSION": [
			"WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
			35724
		],
		"VENDOR": [
			"WebKit",
			7936
		],
		"RENDERER": [
			"WebKit WebGL",
			7937
		],
		"VERSION": [
			"WebGL 1.0 (OpenGL ES 2.0 Chromium)",
			7938
		],
		"MAX_CUBE_MAP_TEXTURE_SIZE": [
			16384,
			34076
		],
		"ACTIVE_TEXTURE": [
			33984,
			34016
		],
		"IMPLEMENTATION_COLOR_READ_TYPE": [
			5121,
			35738
		],
		"IMPLEMENTATION_COLOR_READ_FORMAT": [
			6408,
			35739
		],
		"MAX_RENDERBUFFER_SIZE": [
			16384,
			34024
		],
		"UNPACK_FLIP_Y_WEBGL": [
			false,
			37440
		],
		"UNPACK_PREMULTIPLY_ALPHA_WEBGL": [
			false,
			37441
		],
		"UNPACK_COLORSPACE_CONVERSION_WEBGL": [
			37444,
			37443
		],
		"UNMASKED_VENDOR_WEBGL": [
			"Google Inc. (Apple)",
			37445
		],
		"UNMASKED_RENDERER_WEBGL": [
			"ANGLE (Apple, Apple M2, OpenGL 4.1)",
			37446
		]
	},
	"gl2": {
		"BLEND_EQUATION": [
			32774,
			32777
		],
		"BLEND_EQUATION_RGB": [
			32774,
			32777
		],
		"BLEND_EQUATION_ALPHA": [
			32774,
			34877
		],
		"BLEND_DST_RGB": [
			0,
			32968
		],
		"BLEND_SRC_RGB": [
			1,
			32969
		],
		"BLEND_DST_ALPHA": [
			0,
			32970
		],
		"BLEND_SRC_ALPHA": [
			1,
			32971
		],
		"BLEND_COLOR": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			32773
		],
		"CULL_FACE": [
			false,
			2884
		],
		"BLEND": [
			false,
			3042
		],
		"DITHER": [
			true,
			3024
		],
		"STENCIL_TEST": [
			false,
			2960
		],
		"DEPTH_TEST": [
			false,
			2929
		],
		"SCISSOR_TEST": [
			false,
			3089
		],
		"POLYGON_OFFSET_FILL": [
			false,
			32823
		],
		"SAMPLE_ALPHA_TO_COVERAGE": [
			false,
			32926
		],
		"SAMPLE_COVERAGE": [
			false,
			32928
		],
		"LINE_WIDTH": [
			1,
			2849
		],
		"ALIASED_POINT_SIZE_RANGE": [
			{
				"0": 1,
				"1": 64
			},
			33901
		],
		"ALIASED_LINE_WIDTH_RANGE": [
			{
				"0": 1,
				"1": 1
			},
			33902
		],
		"CULL_FACE_MODE": [
			1029,
			2885
		],
		"FRONT_FACE": [
			2305,
			2886
		],
		"DEPTH_RANGE": [
			{
				"0": 0,
				"1": 1
			},
			2928
		],
		"DEPTH_WRITEMASK": [
			true,
			2930
		],
		"DEPTH_CLEAR_VALUE": [
			1,
			2931
		],
		"DEPTH_FUNC": [
			513,
			2932
		],
		"STENCIL_CLEAR_VALUE": [
			0,
			2961
		],
		"STENCIL_FUNC": [
			519,
			2962
		],
		"STENCIL_FAIL": [
			7680,
			2964
		],
		"STENCIL_PASS_DEPTH_FAIL": [
			7680,
			2965
		],
		"STENCIL_PASS_DEPTH_PASS": [
			7680,
			2966
		],
		"STENCIL_REF": [
			0,
			2967
		],
		"STENCIL_VALUE_MASK": [
			2147483647,
			2963
		],
		"STENCIL_WRITEMASK": [
			2147483647,
			2968
		],
		"STENCIL_BACK_FUNC": [
			519,
			34816
		],
		"STENCIL_BACK_FAIL": [
			7680,
			34817
		],
		"STENCIL_BACK_PASS_DEPTH_FAIL": [
			7680,
			34818
		],
		"STENCIL_BACK_PASS_DEPTH_PASS": [
			7680,
			34819
		],
		"STENCIL_BACK_REF": [
			0,
			36003
		],
		"STENCIL_BACK_VALUE_MASK": [
			2147483647,
			36004
		],
		"STENCIL_BACK_WRITEMASK": [
			2147483647,
			36005
		],
		"VIEWPORT": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			2978
		],
		"SCISSOR_BOX": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			3088
		],
		"COLOR_CLEAR_VALUE": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			3106
		],
		"COLOR_WRITEMASK": [
			[
				true,
				true,
				true,
				true
			],
			3107
		],
		"UNPACK_ALIGNMENT": [
			4,
			3317
		],
		"PACK_ALIGNMENT": [
			4,
			3333
		],
		"MAX_TEXTURE_SIZE": [
			16384,
			3379
		],
		"MAX_VIEWPORT_DIMS": [
			{
				"0": 16384,
				"1": 16384
			},
			3386
		],
		"SUBPIXEL_BITS": [
			4,
			3408
		],
		"RED_BITS": [
			8,
			3410
		],
		"GREEN_BITS": [
			8,
			3411
		],
		"BLUE_BITS": [
			8,
			3412
		],
		"ALPHA_BITS": [
			8,
			3413
		],
		"DEPTH_BITS": [
			24,
			3414
		],
		"STENCIL_BITS": [
			0,
			3415
		],
		"POLYGON_OFFSET_UNITS": [
			0,
			10752
		],
		"POLYGON_OFFSET_FACTOR": [
			0,
			32824
		],
		"SAMPLE_BUFFERS": [
			1,
			32936
		],
		"SAMPLES": [
			4,
			32937
		],
		"SAMPLE_COVERAGE_VALUE": [
			1,
			32938
		],
		"SAMPLE_COVERAGE_INVERT": [
			false,
			32939
		],
		"COMPRESSED_TEXTURE_FORMATS": [
			{},
			34467
		],
		"GENERATE_MIPMAP_HINT": [
			4352,
			33170
		],
		"MAX_VERTEX_ATTRIBS": [
			16,
			34921
		],
		"MAX_VERTEX_UNIFORM_VECTORS": [
			1024,
			36347
		],
		"MAX_VARYING_VECTORS": [
			31,
			36348
		],
		"MAX_COMBINED_TEXTURE_IMAGE_UNITS": [
			32,
			35661
		],
		"MAX_VERTEX_TEXTURE_IMAGE_UNITS": [
			16,
			35660
		],
		"MAX_TEXTURE_IMAGE_UNITS": [
			16,
			34930
		],
		"MAX_FRAGMENT_UNIFORM_VECTORS": [
			1024,
			36349
		],
		"SHADING_LANGUAGE_VERSION": [
			"WebGL GLSL ES 3.00 (OpenGL ES GLSL ES 3.0 Chromium)",
			35724
		],
		"VENDOR": [
			"WebKit",
			7936
		],
		"RENDERER": [
			"WebKit WebGL",
			7937
		],
		"VERSION": [
			"WebGL 2.0 (OpenGL ES 3.0 Chromium)",
			7938
		],
		"MAX_CUBE_MAP_TEXTURE_SIZE": [
			16384,
			34076
		],
		"ACTIVE_TEXTURE": [
			33984,
			34016
		],
		"IMPLEMENTATION_COLOR_READ_TYPE": [
			5121,
			35738
		],
		"IMPLEMENTATION_COLOR_READ_FORMAT": [
			6408,
			35739
		],
		"MAX_RENDERBUFFER_SIZE": [
			16384,
			34024
		],
		"UNPACK_FLIP_Y_WEBGL": [
			false,
			37440
		],
		"UNPACK_PREMULTIPLY_ALPHA_WEBGL": [
			false,
			37441
		],
		"UNPACK_COLORSPACE_CONVERSION_WEBGL": [
			37444,
			37443
		],
		"READ_BUFFER": [
			1029,
			3074
		],
		"UNPACK_ROW_LENGTH": [
			0,
			3314
		],
		"UNPACK_SKIP_ROWS": [
			0,
			3315
		],
		"UNPACK_SKIP_PIXELS": [
			0,
			3316
		],
		"PACK_ROW_LENGTH": [
			0,
			3330
		],
		"PACK_SKIP_ROWS": [
			0,
			3331
		],
		"PACK_SKIP_PIXELS": [
			0,
			3332
		],
		"UNPACK_SKIP_IMAGES": [
			0,
			32877
		],
		"UNPACK_IMAGE_HEIGHT": [
			0,
			32878
		],
		"MAX_3D_TEXTURE_SIZE": [
			2048,
			32883
		],
		"MAX_ELEMENTS_VERTICES": [
			1048575,
			33000
		],
		"MAX_ELEMENTS_INDICES": [
			150000,
			33001
		],
		"MAX_TEXTURE_LOD_BIAS": [
			16,
			34045
		],
		"MAX_DRAW_BUFFERS": [
			8,
			34852
		],
		"DRAW_BUFFER0": [
			1029,
			34853
		],
		"DRAW_BUFFER1": [
			1029,
			34854
		],
		"DRAW_BUFFER2": [
			1029,
			34855
		],
		"DRAW_BUFFER3": [
			1029,
			34856
		],
		"DRAW_BUFFER4": [
			1029,
			34857
		],
		"DRAW_BUFFER5": [
			1029,
			34858
		],
		"DRAW_BUFFER6": [
			1029,
			34859
		],
		"DRAW_BUFFER7": [
			1029,
			34860
		],
		"MAX_FRAGMENT_UNIFORM_COMPONENTS": [
			4096,
			35657
		],
		"MAX_VERTEX_UNIFORM_COMPONENTS": [
			4096,
			35658
		],
		"FRAGMENT_SHADER_DERIVATIVE_HINT": [
			4352,
			35723
		],
		"MAX_ARRAY_TEXTURE_LAYERS": [
			2048,
			35071
		],
		"MIN_PROGRAM_TEXEL_OFFSET": [
			-8,
			35076
		],
		"MAX_PROGRAM_TEXEL_OFFSET": [
			7,
			35077
		],
		"MAX_VARYING_COMPONENTS": [
			124,
			35659
		],
		"MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS": [
			4,
			35968
		],
		"RASTERIZER_DISCARD": [
			false,
			35977
		],
		"MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS": [
			64,
			35978
		],
		"MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS": [
			4,
			35979
		],
		"MAX_COLOR_ATTACHMENTS": [
			8,
			36063
		],
		"MAX_SAMPLES": [
			4,
			36183
		],
		"MAX_VERTEX_UNIFORM_BLOCKS": [
			16,
			35371
		],
		"MAX_FRAGMENT_UNIFORM_BLOCKS": [
			16,
			35373
		],
		"MAX_COMBINED_UNIFORM_BLOCKS": [
			32,
			35374
		],
		"MAX_UNIFORM_BUFFER_BINDINGS": [
			72,
			35375
		],
		"MAX_UNIFORM_BLOCK_SIZE": [
			65536,
			35376
		],
		"MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS": [
			266240,
			35377
		],
		"MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS": [
			266240,
			35379
		],
		"UNIFORM_BUFFER_OFFSET_ALIGNMENT": [
			256,
			35380
		],
		"MAX_VERTEX_OUTPUT_COMPONENTS": [
			64,
			37154
		],
		"MAX_FRAGMENT_INPUT_COMPONENTS": [
			128,
			37157
		],
		"MAX_SERVER_WAIT_TIMEOUT": [
			0,
			37137
		],
		"TRANSFORM_FEEDBACK_PAUSED": [
			false,
			36387
		],
		"TRANSFORM_FEEDBACK_ACTIVE": [
			false,
			36388
		],
		"MAX_ELEMENT_INDEX": [
			4294967295,
			36203
		],
		"MAX_CLIENT_WAIT_TIMEOUT_WEBGL": [
			0,
			37447
		],
		"UNMASKED_VENDOR_WEBGL": [
			"Google Inc. (Apple)",
			37445
		],
		"UNMASKED_RENDERER_WEBGL": [
			"ANGLE (Apple, Apple M2, OpenGL 4.1)",
			37446
		]
	},
	"gl_experimental": {
		"BLEND_EQUATION": [
			32774,
			32777
		],
		"BLEND_EQUATION_RGB": [
			32774,
			32777
		],
		"BLEND_EQUATION_ALPHA": [
			32774,
			34877
		],
		"BLEND_DST_RGB": [
			0,
			32968
		],
		"BLEND_SRC_RGB": [
			1,
			32969
		],
		"BLEND_DST_ALPHA": [
			0,
			32970
		],
		"BLEND_SRC_ALPHA": [
			1,
			32971
		],
		"BLEND_COLOR": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			32773
		],
		"CULL_FACE": [
			false,
			2884
		],
		"BLEND": [
			false,
			3042
		],
		"DITHER": [
			true,
			3024
		],
		"STENCIL_TEST": [
			false,
			2960
		],
		"DEPTH_TEST": [
			false,
			2929
		],
		"SCISSOR_TEST": [
			false,
			3089
		],
		"POLYGON_OFFSET_FILL": [
			false,
			32823
		],
		"SAMPLE_ALPHA_TO_COVERAGE": [
			false,
			32926
		],
		"SAMPLE_COVERAGE": [
			false,
			32928
		],
		"LINE_WIDTH": [
			1,
			2849
		],
		"ALIASED_POINT_SIZE_RANGE": [
			{
				"0": 1,
				"1": 64
			},
			33901
		],
		"ALIASED_LINE_WIDTH_RANGE": [
			{
				"0": 1,
				"1": 1
			},
			33902
		],
		"CULL_FACE_MODE": [
			1029,
			2885
		],
		"FRONT_FACE": [
			2305,
			2886
		],
		"DEPTH_RANGE": [
			{
				"0": 0,
				"1": 1
			},
			2928
		],
		"DEPTH_WRITEMASK": [
			true,
			2930
		],
		"DEPTH_CLEAR_VALUE": [
			1,
			2931
		],
		"DEPTH_FUNC": [
			513,
			2932
		],
		"STENCIL_CLEAR_VALUE": [
			0,
			2961
		],
		"STENCIL_FUNC": [
			519,
			2962
		],
		"STENCIL_FAIL": [
			7680,
			2964
		],
		"STENCIL_PASS_DEPTH_FAIL": [
			7680,
			2965
		],
		"STENCIL_PASS_DEPTH_PASS": [
			7680,
			2966
		],
		"STENCIL_REF": [
			0,
			2967
		],
		"STENCIL_VALUE_MASK": [
			2147483647,
			2963
		],
		"STENCIL_WRITEMASK": [
			2147483647,
			2968
		],
		"STENCIL_BACK_FUNC": [
			519,
			34816
		],
		"STENCIL_BACK_FAIL": [
			7680,
			34817
		],
		"STENCIL_BACK_PASS_DEPTH_FAIL": [
			7680,
			34818
		],
		"STENCIL_BACK_PASS_DEPTH_PASS": [
			7680,
			34819
		],
		"STENCIL_BACK_REF": [
			0,
			36003
		],
		"STENCIL_BACK_VALUE_MASK": [
			2147483647,
			36004
		],
		"STENCIL_BACK_WRITEMASK": [
			2147483647,
			36005
		],
		"VIEWPORT": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			2978
		],
		"SCISSOR_BOX": [
			{
				"0": 0,
				"1": 0,
				"2": 300,
				"3": 150
			},
			3088
		],
		"COLOR_CLEAR_VALUE": [
			{
				"0": 0,
				"1": 0,
				"2": 0,
				"3": 0
			},
			3106
		],
		"COLOR_WRITEMASK": [
			[
				true,
				true,
				true,
				true
			],
			3107
		],
		"UNPACK_ALIGNMENT": [
			4,
			3317
		],
		"PACK_ALIGNMENT": [
			4,
			3333
		],
		"MAX_TEXTURE_SIZE": [
			16384,
			3379
		],
		"MAX_VIEWPORT_DIMS": [
			{
				"0": 16384,
				"1": 16384
			},
			3386
		],
		"SUBPIXEL_BITS": [
			4,
			3408
		],
		"RED_BITS": [
			8,
			3410
		],
		"GREEN_BITS": [
			8,
			3411
		],
		"BLUE_BITS": [
			8,
			3412
		],
		"ALPHA_BITS": [
			8,
			3413
		],
		"DEPTH_BITS": [
			24,
			3414
		],
		"STENCIL_BITS": [
			0,
			3415
		],
		"POLYGON_OFFSET_UNITS": [
			0,
			10752
		],
		"POLYGON_OFFSET_FACTOR": [
			0,
			32824
		],
		"SAMPLE_BUFFERS": [
			1,
			32936
		],
		"SAMPLES": [
			4,
			32937
		],
		"SAMPLE_COVERAGE_VALUE": [
			1,
			32938
		],
		"SAMPLE_COVERAGE_INVERT": [
			false,
			32939
		],
		"COMPRESSED_TEXTURE_FORMATS": [
			{},
			34467
		],
		"GENERATE_MIPMAP_HINT": [
			4352,
			33170
		],
		"MAX_VERTEX_ATTRIBS": [
			16,
			34921
		],
		"MAX_VERTEX_UNIFORM_VECTORS": [
			1024,
			36347
		],
		"MAX_VARYING_VECTORS": [
			31,
			36348
		],
		"MAX_COMBINED_TEXTURE_IMAGE_UNITS": [
			32,
			35661
		],
		"MAX_VERTEX_TEXTURE_IMAGE_UNITS": [
			16,
			35660
		],
		"MAX_TEXTURE_IMAGE_UNITS": [
			16,
			34930
		],
		"MAX_FRAGMENT_UNIFORM_VECTORS": [
			1024,
			36349
		],
		"SHADING_LANGUAGE_VERSION": [
			"WebGL GLSL ES 3.00 (OpenGL ES GLSL ES 3.0 Chromium)",
			35724
		],
		"VENDOR": [
			"WebKit",
			7936
		],
		"RENDERER": [
			"WebKit WebGL",
			7937
		],
		"VERSION": [
			"WebGL 2.0 (OpenGL ES 3.0 Chromium)",
			7938
		],
		"MAX_CUBE_MAP_TEXTURE_SIZE": [
			16384,
			34076
		],
		"ACTIVE_TEXTURE": [
			33984,
			34016
		],
		"IMPLEMENTATION_COLOR_READ_TYPE": [
			5121,
			35738
		],
		"IMPLEMENTATION_COLOR_READ_FORMAT": [
			6408,
			35739
		],
		"MAX_RENDERBUFFER_SIZE": [
			16384,
			34024
		],
		"UNPACK_FLIP_Y_WEBGL": [
			false,
			37440
		],
		"UNPACK_PREMULTIPLY_ALPHA_WEBGL": [
			false,
			37441
		],
		"UNPACK_COLORSPACE_CONVERSION_WEBGL": [
			37444,
			37443
		],
		"READ_BUFFER": [
			1029,
			3074
		],
		"UNPACK_ROW_LENGTH": [
			0,
			3314
		],
		"UNPACK_SKIP_ROWS": [
			0,
			3315
		],
		"UNPACK_SKIP_PIXELS": [
			0,
			3316
		],
		"PACK_ROW_LENGTH": [
			0,
			3330
		],
		"PACK_SKIP_ROWS": [
			0,
			3331
		],
		"PACK_SKIP_PIXELS": [
			0,
			3332
		],
		"UNPACK_SKIP_IMAGES": [
			0,
			32877
		],
		"UNPACK_IMAGE_HEIGHT": [
			0,
			32878
		],
		"MAX_3D_TEXTURE_SIZE": [
			2048,
			32883
		],
		"MAX_ELEMENTS_VERTICES": [
			1048575,
			33000
		],
		"MAX_ELEMENTS_INDICES": [
			150000,
			33001
		],
		"MAX_TEXTURE_LOD_BIAS": [
			16,
			34045
		],
		"MAX_DRAW_BUFFERS": [
			8,
			34852
		],
		"DRAW_BUFFER0": [
			1029,
			34853
		],
		"DRAW_BUFFER1": [
			1029,
			34854
		],
		"DRAW_BUFFER2": [
			1029,
			34855
		],
		"DRAW_BUFFER3": [
			1029,
			34856
		],
		"DRAW_BUFFER4": [
			1029,
			34857
		],
		"DRAW_BUFFER5": [
			1029,
			34858
		],
		"DRAW_BUFFER6": [
			1029,
			34859
		],
		"DRAW_BUFFER7": [
			1029,
			34860
		],
		"MAX_FRAGMENT_UNIFORM_COMPONENTS": [
			4096,
			35657
		],
		"MAX_VERTEX_UNIFORM_COMPONENTS": [
			4096,
			35658
		],
		"FRAGMENT_SHADER_DERIVATIVE_HINT": [
			4352,
			35723
		],
		"MAX_ARRAY_TEXTURE_LAYERS": [
			2048,
			35071
		],
		"MIN_PROGRAM_TEXEL_OFFSET": [
			-8,
			35076
		],
		"MAX_PROGRAM_TEXEL_OFFSET": [
			7,
			35077
		],
		"MAX_VARYING_COMPONENTS": [
			124,
			35659
		],
		"MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS": [
			4,
			35968
		],
		"RASTERIZER_DISCARD": [
			false,
			35977
		],
		"MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS": [
			64,
			35978
		],
		"MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS": [
			4,
			35979
		],
		"MAX_COLOR_ATTACHMENTS": [
			8,
			36063
		],
		"MAX_SAMPLES": [
			4,
			36183
		],
		"MAX_VERTEX_UNIFORM_BLOCKS": [
			16,
			35371
		],
		"MAX_FRAGMENT_UNIFORM_BLOCKS": [
			16,
			35373
		],
		"MAX_COMBINED_UNIFORM_BLOCKS": [
			32,
			35374
		],
		"MAX_UNIFORM_BUFFER_BINDINGS": [
			72,
			35375
		],
		"MAX_UNIFORM_BLOCK_SIZE": [
			65536,
			35376
		],
		"MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS": [
			266240,
			35377
		],
		"MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS": [
			266240,
			35379
		],
		"UNIFORM_BUFFER_OFFSET_ALIGNMENT": [
			256,
			35380
		],
		"MAX_VERTEX_OUTPUT_COMPONENTS": [
			64,
			37154
		],
		"MAX_FRAGMENT_INPUT_COMPONENTS": [
			128,
			37157
		],
		"MAX_SERVER_WAIT_TIMEOUT": [
			0,
			37137
		],
		"TRANSFORM_FEEDBACK_PAUSED": [
			false,
			36387
		],
		"TRANSFORM_FEEDBACK_ACTIVE": [
			false,
			36388
		],
		"MAX_ELEMENT_INDEX": [
			4294967295,
			36203
		],
		"MAX_CLIENT_WAIT_TIMEOUT_WEBGL": [
			0,
			37447
		],
		"UNMASKED_VENDOR_WEBGL": [
			"Google Inc. (Apple)",
			37445
		],
		"UNMASKED_RENDERER_WEBGL": [
			"ANGLE (Apple, Apple M2, OpenGL 4.1)",
			37446
		]
	}
}`)

// func TestString(m *testing.T) {
// 	f := Listener()
// 	f.EnableDebug()
// 	f.Store(string(fp))
// }

func TestStruct(m *testing.T) {
	var data map[string]any

	err := json.Unmarshal(fp2, &data)
	if err != nil {
		panic(err)
	}
	f := Listener()
	// f.EnableDebug()

	f.Threshhold = 5
	f.UseKeyCompression = true

	res, _ := f.Store(data)
	smolData, _ := json.Marshal(res)
	valLookup, _ := json.Marshal(f.database.hashValues)
	keyLookup, _ := json.Marshal(f.database.hashKeys)
	fmt.Println("Data:", string(smolData))
	fmt.Print("\n==\n\n")
	fmt.Println("val lookup:", string(valLookup))
	fmt.Print("\n==\n\n")
	fmt.Println("key lookup:", string(keyLookup))
}
