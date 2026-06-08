# Architectural Changes

This document outlines the architectural changes made to improve the security and functionality of the API.

## 1. Authentication and User Identification
- **Removed `user_id` from Request Bodies**: To prevent unauthorized users from performing actions on behalf of others, the `user_id` field has been removed from all API request bodies.
- **Token-based User Identification**: The `user_id` is now extracted from the JWT authentication token.
- **HTTP Handler Extraction**: HTTP handlers extract the token from the `Authorization` header, use a utility function in the `AuthService` to retrieve the `user_id`, and set it on the domain entities before passing them to the Command Handlers.

## 2. Document API Enhancements
- **Minio Integration**: The Document API now integrates with Minio (Blob Storage).
- **URL Return**: When a document metadata is saved in PostgreSQL, the API returns the download URL of the file stored in Minio. This is handled by the `CreateDocumentHandler` using the `ObjectStorageProvider`.

## 3. Clean Architecture and CQRS
- The project continues to adhere to Clean Architecture, DDD, and CQRS principles.
- Business logic remains intact in the domain layer.
- Clear separation between Commands (state-changing operations) and Queries (data-retrieval operations).
