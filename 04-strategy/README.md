# Strategy Pattern

## What is the Strategy Pattern?

The **Strategy Pattern** is a behavioral design pattern that allows a behavior to be selected at runtime. It defines a
family of algorithms, encapsulates each one, and makes them interchangeable. This enables the client code to choose or
switch between different strategies without altering the code that uses them.

---

## Why Do We Need the Strategy Pattern?

The Strategy Pattern is useful when you need to define multiple variations of an algorithm or behavior and switch
between them dynamically at runtime. By encapsulating each algorithm in its own, this pattern promotes **open/closed
principle**, meaning the system is open for extension but closed for modification.

### Key Benefits:

- **Flexibility:** Swap algorithms at runtime without modifying the client code.
- **Reusability:** Algorithms are encapsulated, making them reusable across different contexts.
- **Separation of Concerns:** Decouples the behavior from the context in which it is used.

---

## Common Use Cases

- **Sorting Algorithms:** Selecting different sorting techniques, like quicksort or mergesort, depending on data size or
  structure.
- **Payment Systems:** Supporting multiple payment methods (e.g., credit card, PayPal, crypto) with interchangeable
  strategies.
- **AI in Games:** Enabling characters to switch between behaviors (e.g., aggressive, defensive) based on the game
  state.
- **Compression Utilities:** Choosing different compression algorithms, such as ZIP, RAR, or GZIP, depending on the use
  case.