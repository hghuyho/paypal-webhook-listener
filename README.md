# Webhook endpoint for gRPC-Gateway
Implementation for Webhook Endpoint for gRPC-Gateway (Eg.PayPal listener)
### Setup local development
####  How to generate code
    ```bash
    make proto
    ```
####  How to run
    ```bash
    make server
    ```
...
Expose localhost:8080 via ngrok to listen PayPal webhook.