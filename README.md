# HyperHive

[![GitHub](https://img.shields.io/badge/GitHub-Repository-blue)](https://github.com/Vajid-Hussain/HyperHive)
[![Live](https://img.shields.io/badge/Live-Demo-green)](https://hyperhive.vajid.tech/swagger/index.html)

## Overview

HyperHive is a high-performance API designed to provide scalable social media functionalities. It leverages modern technologies and best practices to deliver a robust platform capable of handling extensive user interactions and dynamic content updates.

## Features

- **Community-Focused Social Media APIs**: Built using the [Echo](https://echo.labstack.com/) Go framework, enabling essential social networking features.
- **Microservices Architecture**: Utilizes [gRPC](https://grpc.io/) for efficient, scalable, and language-agnostic inter-service communication.
- **Clean Code Principles**: Enhances code maintainability, readability, and long-term project flexibility. Optimizes performance with [Redis](https://redis.io/) caching.
- **Infrastructure as Code (IaC)**: Automates S3 bucket provisioning using [Terraform](https://www.terraform.io/).
- **Real-Time Communication**: Implements [Socket.IO](https://github.com/googollee/go-socket.io) for dynamic group and private chat functionalities.
- **Reliable Messaging**: Integrates [Kafka](https://kafka.apache.org/) as a robust message broker for reliable message delivery and fault tolerance.
- **Scalable Deployment**: Orchestrates dynamic application deployment with [Kubernetes](https://kubernetes.io/) for enhanced scalability.

## Installation

### Prerequisites

- Go 1.16+
- Docker
- Kubernetes
- Terraform
- Redis
- Kafka

### Steps

1. **Clone the repository:**

    ```bash
    git clone https://github.com/Vajid-Hussain/HyperHive.git
    cd HyperHive
    ```

2. **Setup environment variables:**

    Create a `.env` file in the root directory and add the necessary environment variables.

3. **Provision Infrastructure:**

    Use Terraform to provision required infrastructure components.

    ```bash
    terraform init
    terraform apply
    ```

4. **Run the application:**

    ```bash
    go run main.go
    ```

## Usage

### API Documentation

Access the live API documentation [here](https://hyperhive.vajid.tech/swagger/index.html).

### Real-Time Communication

The application supports real-time group and private chat functionalities using Socket.IO.

### Message Broker

Kafka is used as the message broker to ensure reliable message delivery and fault tolerance.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for review.

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a pull request

---

[GitHub Repository](https://github.com/Vajid-Hussain/HyperHive) | [Live Demo](https://hyperhive.vajid.tech/swagger/index.html)
