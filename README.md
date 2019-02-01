# Webhook Dispatcher

[![Build Status](https://drone-kr.webrelay.io/api/badges/webhookrelay/webhook-dispatcher/status.svg)](https://drone-kr.webrelay.io/webhookrelay/webhook-dispatcher)

Simple webhook dispatcher that can be used in the pipelines. Targeted at container workloads where native webhooks aren't that easy to get or configure (looking at you Google Cloud Builder).

## Usage

```bash
$ webhook-dispatcher --destination https://my.webhookrelay.com/v1/webhooks/544a6fe8-83fe-4361-a264-0fd486e1665d --body hello
{"status_code":200,"body":""}
```

Available options:

```bash
webhook-dispatcher --help
Usage of webhook-dispatcher:
  --basic-auth string
    	Optional basic authentication in a 'user:pass' format
  --body string
    	Webhook payload
  --destination string
    	Webhook destination (https://example.com/webhooks)
  --method string
    	Webhook method (defaults to 'POST')

```