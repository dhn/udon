# udon

```

  88   88 8888b.   dP"Yb  88b 88 
  88   88  8I  Yb dP   Yb 88Yb88 
  Y8   8P  8I  dY Yb   dP 88 Y88 
   YbodP' 8888Y"   YbodP  88  Y8                                                   


                     made with <3
```

**udon**: A simple tool that helps to find domains based on the Google Analytics ID.

<p align="left">
<a href="https://goreportcard.com/report/github.com/dhn/udon/"><img src="https://goreportcard.com/badge/github.com/dhn/udon"></a>
<a href="https://github.com/dhn/udon/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
</p>

# Building

* Download & install Go: https://golang.org/doc/install

```
go install github.com/dhn/udon@latest
```

# Usage

```
Usage of udon:
  -json
        Print results as JSON
  -s string
        UA ID to find domains for
  -silent
        Show only domains in output
  -version
        Show version of udon
```

# Running udon

**udon** can be used to find domains/subdomains for a given Google Analytics ID. In this example for "UA-33427076":

```json
➜  $ udon -silent -json -s UA-33427076 | jq -c
{"domain":"soester-anzeiger.de","source":"site-overview"}
{"domain":"wa.de","source":"site-overview"}
{"domain":"suederlaender-tageblatt.de","source":"site-overview"}
{"domain":"trauer.nrw","source":"site-overview"}
{"domain":"come-on.de","source":"site-overview"}
{"domain":"immobilien.wa.de","source":"spyonweb"}
{"domain":"nrw-jobs.de","source":"spyonweb"}
{"domain":"auto.wa.de","source":"hackertarget"}
{"domain":"come-on.de","source":"hackertarget"}
{"domain":"immobilien.wa.de","source":"hackertarget"}
{"domain":"soester-anzeiger.de","source":"hackertarget"}
{"domain":"trauer.nrw","source":"hackertarget"}
{"domain":"wa-mediengruppe.de","source":"hackertarget"}
{"domain":"wa.de","source":"hackertarget"}
{"domain":"web.archive.org","source":"hackertarget"}
{"domain":"www.come-on.de","source":"hackertarget"}
{"domain":"www.soester-anzeiger.de","source":"hackertarget"}
{"domain":"www.wa.de","source":"hackertarget"}
{"domain":"wa.de","source":"osint.sh"}
{"domain":"soester-anzeiger.de","source":"osint.sh"}
{"domain":"come-on.de","source":"osint.sh"}
{"domain":"dispspotegcred.tk","source":"osint.sh"}
{"domain":"conmamanan.tk","source":"osint.sh"}
```

To find UA id's following simple `nuclei` [1] template can be used:

```yaml
id: google-analytics-id

info:
  name: Google Analytics ID Discovery
  author: dhn
  severity: info
  tags: tech

requests:
  - method: GET

    path:
      - "{{BaseURL}}"

    extractors:
      - type: regex
        name: ua-id
        part: body
        regex:
          - "UA-[0-9]+(-[0-9]+)"

    redirects: true
    max-redirects: 5
    stop-at-first-match: true
    matchers-condition: and
    matchers:
      - type: regex
        part: body
        regex:
          - "UA-[0-9]+(-[0-9]+)"

      - type: status
        status:
          - 200
```

## Disclaimer

**udon** leverages multiple open APIs, it is developed for individuals to help them for research or internal work. If you wish to incorporate this tool into a commercial offering or purposes, you must agree to the Terms of the leveraged services:

- Hackertarget - https://hackertarget.com
- OSINT - https://osint.sh
- Site-Overview - http://site-overview.com
- SpyOnWeb - https://spyonweb.com

---
You expressly understand and agree that **udon** (creators and contributors) shall not be liable for any damages or losses resulting from your use of this tool or third-party products that use it.

Creators aren't in charge of any and have/has no responsibility for any kind of:

- Unlawful or illegal use of the tool.
- Legal or Law infringement (acted in any country, state, municipality, place) by third parties and users.
- Act against ethical and / or human moral, ethic, and peoples and cultures of the world.
- Malicious act, capable of causing damage to third parties, promoted or distributed by third parties or the user through this tool.

This disclaimer was shameless stolen from: https://github.com/projectdiscovery/subfinder/blob/master/DISCLAIMER.md

## Reference

- [1] https://github.com/projectdiscovery/nuclei