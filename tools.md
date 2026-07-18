These are all good design questions. Here's how I'd approach each one.

## 1. Xsolla for payments

Yes, this is actually one of the reasons your project being inspired by Xsolla is useful.

Xsolla provides a **sandbox environment** where you can simulate successful and failed payments using test payment methods and cards. No real money is charged, and you still receive receipts and webhooks, making it ideal for developing and testing your payment flow. ([Xsolla Documentation][1])

However, there's an important limitation:

* **Before** signing a licensing agreement, your project can only use **sandbox**.
* **After** going live, sandbox is still available, but after the first real payment, access becomes more restricted (typically limited to authorized project users). ([Xsolla Documentation][1])

For a portfolio project, I'd implement the full flow:

```text
React
    │
    ▼
Go API
    │
Create Pending Payment
    │
Redirect to Xsolla Sandbox
    │
Xsolla webhook
    │
Update payment
    │
Create reservation
```

This demonstrates:

* external API integration
* webhook handling
* idempotency
* asynchronous workflows
* payment state machine

Those are excellent backend topics for interviews.

---

## 2. Can PostgreSQL do "movie about sharks" → Jaws?

### Short answer

**Not by itself.**

There are different levels of search.

### Level 1 — Full Text Search (`tsvector`)

Can answer:

> shark attack

if your movie overview contains:

> "A giant shark terrorizes a seaside town..."

Yes.

### Level 2 — `pg_trgm`

Can answer:

```
Jwas
```

↓

```
Jaws
```

It fixes typos.

### Level 3 — `pgvector`

Now we're talking about semantic similarity.

Suppose you embed movie descriptions.

User searches:

```
movie about sharks
```

The model converts this into a vector.

The movie overview

```
A great white shark attacks beachgoers...
```

is converted into another vector.

The vectors are close together.

This returns **Jaws**, even if the query doesn't exactly match the text.

That's semantic search.

---

### Do you need Elasticsearch?

No.

Nowadays many teams build semantic search with:

```
PostgreSQL
+
pgvector
+
Full Text Search
```

instead of

```
PostgreSQL
+
Elasticsearch
```

A common production architecture is:

```
tsvector
        +
pg_trgm
        +
pgvector
```

This gives:

* keyword search
* typo tolerance
* semantic search

without introducing another database. ([PostgreSQL][2])

For 1,000 movies, this is more than sufficient.

---

## 3. What about Goose and other Go tools?

This is where Go 1.24+ is really nice.

Instead of asking every developer to install tools globally:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

you can pin them in your module.

For example:

```bash
go get -tool github.com/pressly/goose/v3/cmd/goose@latest
go get -tool github.com/gofiber/cli/fiber@latest
go get -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

Then everyone can simply run:

```bash
go tool goose
go tool fiber
go tool golangci-lint
```

using the exact versions recorded in your `go.mod`. Go 1.24 introduced `tool` directives specifically for managing executable dependencies this way. ([Go][3])

I'd use this for all your project tooling:

* Goose
* Fiber CLI
* golangci-lint
* swag
* mockgen (if you add it)
* sqlc (if you ever adopt it)

That way:

* no global installs
* reproducible versions
* easier onboarding for anyone cloning your repository

---

## What I would build if I were making this project

Your stack is becoming very close to something I'd expect from a strong backend portfolio:

* **Go + Fiber**
* **PostgreSQL** (source of truth)
* **Redis** (seat holds with TTL)
* **RabbitMQ** (notifications, payment-related events after the booking transaction)
* **TMDB Seeder**
* **Xsolla Sandbox** (payment integration)
* **Go tool directives** (versioned project tools)
* **Goose** (database migrations)
* **Docker Compose**
* **Swagger/OpenAPI**
* **JWT + RBAC**
* **PostgreSQL Full-Text Search + `pg_trgm`**, with **`pgvector`** if you want to experiment with semantic search

I would skip Elasticsearch for this project unless your goal is specifically to learn Elasticsearch. For a catalog of 500–1,000 movies, PostgreSQL plus its search extensions can provide excellent functionality with much less operational complexity.

[1]: https://developers.xsolla.com/dev-resources/testing/general-info/?utm_source=chatgpt.com "Xsolla Documentation – Sandbox and production environments"
[2]: https://www.postgresql.org/docs/9.5/datatype-textsearch.html?utm_source=chatgpt.com "PostgreSQL: Documentation: 9.5: Text Search Types"
[3]: https://go.dev/doc/go1.24?file=go1.24&utm_source=chatgpt.com "Go 1.24 Release Notes - The Go Programming Language"
