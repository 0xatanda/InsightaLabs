# 🚀 Overview

This project is a **demographic intelligence API system** built for Insighta Labs.
It enables clients (marketing, product, and analytics teams) to:

* Filter large profile datasets
* Sort and paginate results
* Query using natural language (rule-based parsing)
* Retrieve structured demographic insights

Built using **Go + PostgreSQL** without external web frameworks.

---

## 🧱 Tech Stack

* Go (net/http)
* PostgreSQL
* SQL Query Builder (custom)
* Rule-based Natural Language Parser
* EC2 Deployment (Linux Ubuntu)

---

## 📂 Project Structure

```text
cmd/
  api/          → main API server
  seed/         → database seeding CLI

internal/
  config/       → DB connection
  handler/      → HTTP handlers
  service/      → business logic
  repository/   → DB queries
  query/        → SQL builder
  parser/       → natural language parser
  utils/        → helpers (UUID, JSON, etc)

scripts/
  seed.json     → dataset (2026 profiles)
```

---

## 🧠 Features

### ✅ 1. Profiles API

`GET /api/profiles`

Supports:

* Filtering (gender, age_group, country_id, age range)
* Sorting (age, created_at, gender_probability)
* Pagination (page, limit)

Example:

```bash
/api/profiles?gender=male&country_id=NG&sort_by=age&order=desc&page=1&limit=10
```

---

### ✅ 2. Natural Language Search Engine

`GET /api/profiles/search?q=...`

Converts plain English into structured filters using rule-based parsing.

#### Examples

| Query                    | Interpretation                             |
| ------------------------ | ------------------------------------------ |
| young males from nigeria | gender=male + age 16–24 + country=NG       |
| females above 30         | gender=female + min_age=30                 |
| adult males from kenya   | gender=male + age_group=adult + country=KE |
| teenage users above 17   | age_group=teenager + min_age=17            |

---

### ⚠️ Parsing Rules

* “young” → age 16–24
* “male / female” → gender filter
* country names → ISO codes
* “above X” → min_age
* age groups: child / teenager / adult / senior

---

## 🧠 Limitations

This system uses **rule-based parsing only**:

* No AI / LLM usage
* No fuzzy matching
* No complex boolean logic (AND/OR)
* Limited slang interpretation
* Strict keyword-based mapping only

If query cannot be parsed:

```json
{
  "status": "error",
  "message": "Unable to interpret query"
}
```

---

## 🗄️ Database Schema

Table: `profiles`

| Field               | Type             |
| ------------------- | ---------------- |
| id                  | UUID v7          |
| name                | VARCHAR (unique) |
| gender              | VARCHAR          |
| gender_probability  | FLOAT            |
| age                 | INT              |
| age_group           | VARCHAR          |
| country_id          | VARCHAR(2)       |
| country_name        | VARCHAR          |
| country_probability | FLOAT            |
| created_at          | TIMESTAMP (UTC)  |

---

## 🌱 Data Seeding

To seed the database:

```bash
go run ./cmd/seed
```

* Reads `scripts/seed.json`
* Inserts 2026 profiles
* Uses `ON CONFLICT (name) DO NOTHING` to avoid duplicates

---

## 🌐 Live Deployment

Base URL:

```
http://54.89.73.180
```

---

## 📡 API Endpoints

### Profiles

```bash
GET /api/profiles
```

### Search

```bash
GET /api/profiles/search?q=young males from nigeria
```

---

## ⚙️ Error Handling

Standard response format:

```json
{
  "status": "error",
  "message": "description"
}
```

### HTTP Codes

* 400 → Bad Request
* 404 → Not Found
* 422 → Unprocessable Entity
* 500 → Server Error

---

## ⚡ Performance Notes

* Indexed queries on filters
* Pagination prevents full-table scans
* Lightweight SQL builder for dynamic queries


## 🏁 Status

✔ Filtering system
✔ Sorting system
✔ Pagination
✔ Natural language query engine
✔ Production deployment
✔ Seeded dataset (2026 profiles)

---
