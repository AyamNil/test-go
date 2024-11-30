# API Documentation: Comments Endpoint

## Overview
This API allows users to fetch, create, and delete comments. Below is the detailed information on how to interact with the endpoints.

---

## Endpoints

### 1. **Fetch Comments**
- **Endpoint:** `/fetch-comments`
- **Method:** `GET`
- **Description:** Retrieve a list of comments.

---

### 2. **Create Comment**
- **Endpoint:** `/comments`
- **Method:** `POST`
- **Description:** Add a new comment.
- **Request Body Example:**
  ```json
  {
    "userId": 1,
    "id": 1,
    "title": "TEST Post",
    "body": "This is Post API testing"
  },
