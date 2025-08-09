# 💻 Frontend - Book Library CRUD App

This is a React + Next.js + TypeScript application for managing a small library of books.

---

## 🛠️ Setup

### 1. Prerequisites

- Node.js ≥ 18
- npm

### 2. Install Dependencies

```bash
npm install
```

### 3. Environment Variables

Create a .env.local file in the project root with:

```
NEXT_PUBLIC_BOOKLIB_API_BASE_URL=http://localhost:8080
```

### 4. Run Development Server

```bash
npm run dev
```

## 🗂️ Project Structure

```
⏺ frontend/
  ├── public/              # Static assets served directly
  └── src/
      ├── data/           # Static data files and data utilities
      ├── pages/          # Next.js pages directory (file-based routing)
      │   └── books/      # Book-related pages
      ├── styles/         # CSS and styling files
      └── types/          # TypeScript type definitions

```
