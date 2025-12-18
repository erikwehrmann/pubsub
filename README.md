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
