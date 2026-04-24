 # 📘 Insighta Labs — Intelligence Query Engine (Stage 2)

## Overview

This project is a backend intelligence query engine built for **Insighta Labs**, designed to allow clients to query demographic profiles using:

* Structured filtering (gender, age, country, etc.)
* Sorting (age, created_at, gender_probability)
* Pagination (page/limit)
* Natural language queries (rule-based parsing, no AI/LLMs)

The system is built using Go and PostgreSQL, exposing RESTful endpoints for querying a profiles dataset.

---

# 📦 API Endpoints

## 1. Get Profiles

### `GET /api/profiles`

Supports:

* Filtering
* Sorting
* Pagination

### Query Parameters

#### Filters

* `gender` → male | female
* `age_group` → child | teenager | adult | senior
* `country_id` → ISO country code (e.g. NG, KE)
* `min_age` → integer
* `max_age` → integer
* `min_gender_probability` → float
* `min_country_probability` → float

#### Sorting

* `sort_by` → age | created_at | gender_probability
* `order` → asc | desc

#### Pagination

* `page` (default: 1)
* `limit` (default: 10, max: 50)

---

### Example Request

```
/api/profiles?gender=male&country_id=NG&min_age=25&sort_by=age&order=desc&page=1&limit=10
```

---

## 2. Natural Language Search

### `GET /api/profiles/search?q=...`

Parses human-readable queries into structured filters.

---

### Supported Query Patterns

| Natural Language Input             | Parsed Filters                                |
| ---------------------------------- | --------------------------------------------- |
| young males                        | gender=male + age 16–24                       |
| females above 30                   | gender=female + min_age=30                    |
| people from angola                 | country_id=AO                                 |
| adult males from kenya             | gender=male + age_group=adult + country_id=KE |
| male and female teenagers above 17 | age_group=teenager + min_age=17               |

---

# 🧠 Natural Language Parsing Approach

## Strategy

This system uses **rule-based keyword matching only** (NO AI / NO LLM).

Parsing is done in sequential stages:

---

## 1. Gender Detection

Keywords:

* "male", "males" → gender = male
* "female", "females" → gender = female

Special:

* If both "male and female" → gender filter is removed (treated as all)

---

## 2. Age Mapping Rules

### Explicit rules:

* "above X" → min_age = X
* "below X" → max_age = X

### Relative terms:

* "young" → age 16–24
* "teenager" → age_group = teenager
* "adult" → age_group = adult
* "senior" → age_group = senior

---

## 3. Country Parsing

Pattern:

* "from {country}" → maps to ISO country code

Example:

* Kenya → KE
* Nigeria → NG
* Angola → AO

---

## 4. Combined Parsing Logic

Filters are merged using AND logic:

Example:

> "adult males from kenya"

Results in:

* gender = male
* age_group = adult
* country_id = KE

---

# ⚠️ Query Validation Rules

All invalid inputs return:

```json
{ "status": "error", "message": "Invalid query parameters" }
```

### Invalid cases:

* unsupported `sort_by`
* invalid `order`
* malformed query structure

---

### Uninterpretable Natural Language Query

If query cannot be parsed:

```json
{ "status": "error", "message": "Unable to interpret query" }
```

---

# 📄 Pagination Behavior

* Pagination is applied after filtering and sorting
* OFFSET is calculated as:

```
offset = (page - 1) * limit
```

* Ensures:

  * no overlap between pages
  * stable deterministic ordering
  * consistent dataset slicing

---

# ⚙️ System Architecture

```
Handler → Service → Parser → Query Builder → Repository → PostgreSQL
```

### Responsibilities

* **Handler**: HTTP request/response
* **Service**: business logic orchestration
* **Parser**: natural language → structured filters
* **Query Builder**: SQL generation
* **Repository**: DB execution layer

---

# 🧪 Data Seeding

* Seeded from `scripts/seed.json`
* Uses `ON CONFLICT (name) DO NOTHING` to prevent duplicates
* UUID v7 generated for all records
* Age groups automatically derived during insertion

---

# ⚠️ Limitations

This system does NOT support:

### 1. Advanced NLP

* No machine learning
* No semantic understanding
* Strict keyword-based parsing only

---

### 2. Ambiguous Queries

Examples not supported:

* “youngish people”
* “wealthy young adults”
* “mostly females under 40”

---

### 3. Complex Boolean Logic

Not supported:

* OR conditions across filters (except implicit gender dual case)
* nested logic like parentheses expressions

---

### 4. Geographic inference

* Only direct country name → ISO mapping supported
* No city-level or regional inference

---

# 🚀 Performance Considerations

* Indexed filtering on:

  * gender
  * country_id
  * age
* Pagination ensures no full table scan is returned to clients
* Query builder avoids unnecessary joins

---

# 🌐 CORS

All responses include:

```
Access-Control-Allow-Origin: *
```

Required for grading access.

---

# 📊 Summary

This system enables:

* Fast demographic filtering
* Deterministic pagination
* Safe SQL query generation
* Rule-based natural language interpretation

---

# ✅ Submission Notes

* All endpoints tested with production dataset
* Seed script is idempotent
* Pagination is stable and non-overlapping
* Natural language parser is deterministic and rule-based

