# Sharingan - Fees API 🚀

---

#### Fees API in encore (encore.dev) that uses a temporal (temporal.io) workflow started at the beginning of a fee period, and allows for the progressive accrual of fees. At the end of the billing period, the total invoice and bill summation are available.

---

## 🔥 Run locally
1. Clone repo:
```shell
git clone git@github.com:dsha256/pave-coding-challenge-david-shainidze.git
```

2. Install [Temporal](https://learn.temporal.io/getting_started/go/), and run:
```shell
temporal server start-dev --dynamic-config-value frontend.enableUpdateWorkflowExecution=true
```

3. Install [Encore](https://encore.dev/docs/install), navigate to the root folder of the project and run:
```shell
go mod tidy
```
```shell
encore run
```

3. Install Encore(https://encore.dev/docs/install), navigate to the root folder of the project and run:
```shell
encore run
```

* ### Then you get two endpoints:
  - API - http://127.0.0.1:4000
  - Development Dashboard http://localhost:9400/sharingan-ynp2

---

# 👀 Want to give it a shot?
Go to the development dashboard (http://localhost:9400/sharingan-ynp2) and the Encore's `API Explorer` wil show you the way 😊💀🤞

## > [ API Explorer](https://encore.dev/docs/observability/dev-dash)
<div>
  <a href="https://github.com/dsha256/pave-coding-challenge-david-shainidze">
    <img src="/.github/images/encore_api_explorer.png" alt="API Explorer" />
  </a>
</div>