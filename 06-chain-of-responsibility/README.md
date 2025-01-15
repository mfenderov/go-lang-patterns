# Chain of Responsibility

The **Chain of Responsibility(sometimes called Middleware)** is a design pattern used to build a chain of processing
steps that a request goes through before reaching its final destination.
Each chain link is responsible for a specific task, such as logging,
authentication, or modifying the request or response. This pattern is highly modular, flexible, and commonly used in web
frameworks and servers.

---

## Why Use the Chain of Responsibility Pattern?

1. **Separation of Concerns**: Each chain link handles a specific responsibility (e.g., logging, validation,
   authentication) without mixing logic.
2. **Modularity**: Chain links are independent and reusable.
3. **Scalability**: You can easily add, remove, or reorder chain links without affecting other links.
4. **Composability**: Chain links can be combined to form a complex processing pipeline.
5. **Maintainability**: Simplifies the codebase by breaking down responsibilities into smaller, manageable pieces.

---

## When to Use Chain of Responsibility?

- **Web Servers**: Handling HTTP requests and responses.
- **API Gateways**: Applying security, logging, and rate-limiting policies.
- **Event Pipelines**: Processing events through a series of stages (e.g., data enrichment, validation).
- **Message Brokers**: Intercepting and processing messages between producers and consumers.
