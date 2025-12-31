Work in progress...

# Go Event-Driven API with Pub/Sub

A simple Go HTTP API demonstrating an **event-driven architecture** using a custom in-memory **Pub/Sub system**.  
Events are published by the API and processed asynchronously by background subscribers, with **retry logic** and **context-aware shutdown**.


## Features

- HTTP API endpoint (`POST /orders`) that creates an order event
- Custom Pub/Sub bus using Go channels for fan-out
- Background subscribers (workers) process events independently
- Retry logic with exponential backoff for resilient subscribers
- Graceful shutdown using `context.Context`


## Running API

```bash
docker compose up --build
```

API will be available at http://localhost:8080

## Using the API

```bash
curl -X POST http://localhost:8080/orders
```

## Expected Output
```
[LOGGER] Event received: events.OrderCreated{OrderID:"3e710a5c-8829-41f8-8c50-2569d0f11a23", Amount:99.99, CreatedAt:time.Date(2025, time.December, 31, 19, 10, 48, 941699966, time.Local)}
[EMAIL] Sending email for event {3e710a5c-8829-41f8-8c50-2569d0f11a23 99.99 2025-12-31 19:10:48.941699966 +0000 UTC m=+2.871329170}
[EMAIL] Email sent
```
