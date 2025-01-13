## Observer Pattern: Detailed Dive

The Observer Pattern defines a one-to-many dependency between objects so that when one object (the subject) changes
state, all its dependent objects (the observers) are notified and updated automatically.

## Why Use Observer?

Event-driven systems: Perfect for notifications or pub/sub models.
Decoupling: The subject doesn’t need to know about the observers’ implementation.
Scalability: Add or remove observers dynamically.

## Social Media Influencers Example 🌟

Scenario:  
Imagine a social media app where users (followers) get notified whenever an influencer posts something new. The
influencer is the subject, and the followers are the observers.

