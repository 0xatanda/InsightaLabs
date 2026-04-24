# InsightaLabs Stage 2 – Intelligence Query Engine

## Overview
This system provides advanced filtering, sorting, pagination, and natural language query parsing for demographic profiles.

---

## Filtering System

Supported filters:
- gender
- age_group
- country_id
- min_age / max_age
- min_gender_probability
- min_country_probability

All filters are combined using AND logic.

---

## Sorting

Allowed fields:
- age
- created_at
- gender_probability

Order:
- asc
- desc

Invalid sort parameters are rejected with:
{
  "status": "error",
  "message": "Invalid query parameters"
}

---

## Pagination

Query params:
- page (default 1)
- limit (default 10, max 50)

Implemented using SQL LIMIT and OFFSET:
OFFSET = (page - 1) * limit

---

## Natural Language Parsing

Rule-based parser (no AI used).

Mappings:

- "young" → age 16–24
- "females" → gender=female
- "males" → gender=male
- "adult" → age_group=adult
- "teenagers" → age_group=teenager
- "above X" → min_age=X
- "from X" → country_id mapping (ISO code lookup)

Examples:
- "females above 30" → gender=female + min_age=30
- "adult males from kenya" → gender=male + age_group=adult + country_id=KE

Uninterpretable queries return:
{
  "status": "error",
  "message": "Unable to interpret query"
}

---

## Limitations

- No fuzzy matching for country names
- No synonym expansion beyond fixed rules
- Cannot handle complex nested logic (e.g. OR conditions)
- NLP is rule-based only (no LLMs used)