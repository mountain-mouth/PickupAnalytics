```mermaid
erDiagram
    users {
        int id PK
        text name
        datetime created_at
        datetime updated_at
        int age
        int experience_years
        bool is_published
    }

    area {
        int id PK
        text name

    }
```
