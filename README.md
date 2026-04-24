# InsightaLabs API

A profile search engine with structured filtering and NLP query parsing.

## Endpoints

GET /api/profiles  
GET /api/profiles/search?q=

## Features

- Filtering (gender, age, country)
- Pagination (page, limit)
- Sorting (age, created_at, gender_probability)
- Natural language parsing

## NLP Examples

- "females above 30"
- "adult males from kenya"

## Error Format

{ "status": "error", "message": "Unable to interpret query" }