# Запросы приложения

Этот пакет содержит обработчики запросов (Query Handlers) согласно паттерну CQRS.

## Document Queries
- **FetchDocumentByIDHandler**: Получает метаданные документа по ID и генерирует URL для скачивания из Minio.
- **FetchDocumentListHandler**: Получает список всех документов.
