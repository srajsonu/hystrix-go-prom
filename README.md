# hystrix-go-prom

`hystrix-go-prom` is a custom collector for Hystrix metrics, designed to integrate with Prometheus for monitoring purposes.

## Table of Contents

- [About the Project](#about-the-project)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## About the Project

`hystrix-go-prom` provides a way to collect and expose Hystrix metrics to Prometheus. This allows for effective monitoring and alerting based on the health and performance of your Hystrix commands.

## Getting Started

To get a local copy up and running, follow these steps.

### Prerequisites

Ensure you have Go installed on your machine. You can download it from the official [Go website](https://golang.org/).

### Installation

1. **Initialize Go Modules:**  
   If your project is not already using Go modules, initialize it by running:
   ```bash
   go get github.com/srajsonu/hystrix-go-prom
