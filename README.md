# InsightaLabs API

## Base URL
http://54.89.73.180

## Endpoints

### GET /api/profiles
Supports:
- Filtering
- Sorting
- Pagination

### GET /api/profiles/search?q=
Natural language query parsing.

---

## Filtering

- gender
- age_group
- country_id
- min_age / max_age

---

## Sorting

Allowed fields:
- age
- created_at
- gender_probability

---

## Pagination

- page (default: 1)
- limit (default: 10, max: 50)

---

## NLP Examples

- "young males" → male + age 16–24
- "females above 30" → female + min_age 30
- "adult males from kenya" → male + adult + KE

---

## Errors

All errors follow:

{
  "status": "error",
  "message": "<message>"
}