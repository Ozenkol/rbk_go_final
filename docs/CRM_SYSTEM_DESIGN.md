# CRM System Design

## Architecture
The system follows Clean Architecture and DDD principles with a CQRS pattern.

### Layers:
- **Domain**: Core entities and business logic.
- **Application**: CQRS Command and Query handlers.
- **Infrastructure**: GORM-based persistence and external adapters.
- **Delivery**: HTTP API using the Gin framework.

## Domain Model
- **Client**: Represents a customer (Person or Company).
- **Deal**: Manages the sales pipeline stages.
- **Proposal**: Commercial offers sent to clients.
- **Task**: To-do items and reminders for employees.
- **Communication**: History of interactions (Calls, Emails, WhatsApp).
- **Contract**: Legal agreements with clients.
- **Invoice**: Billing and payment tracking.
- **Product**: Catalog of goods and services.

## Multi-tenancy
- Each resource is linked to a `CompanyID` and `UserID` for isolation.

## External Integrations
- **Object Storage**: Pre-signed URLs for document management.
- **Messaging**: Interfaces for WhatsApp and SMS delivery.
- **Telephony**: Call logging capabilities.
