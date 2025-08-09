import {Book} from "@/types/book";

const API_URL = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/v1/books`

interface ApiResponse<T> {
    data: T
    error: string
    status: string
}

export async function getBooks(): Promise<Book[]> {
    const res = await fetch(API_URL)
    const json: ApiResponse<Book[]> = await res.json()

    if (json.status !== 'success') {
        throw new Error(`API error: ${json.error}`)
    }

    return json.data || []
}

export async function getBookById(id: string): Promise<Book> {
    const res = await fetch(`${API_URL}/${id}`)
    const json: ApiResponse<Book> = await res.json()
    return json.data
}

export async function addBook(book: Book) {
    const res = await fetch(`${API_URL}`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(book),
    })
    const json: ApiResponse<Book> = await res.json()
    if (json.status !== 'success') {
        throw new Error(`API error: ${json.error}`)
    }
    return json.data
}

export async function updateBook(book: Book) {
    const res = await fetch(`${API_URL}/${book.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(book),
    })
    if (!res.ok) throw new Error('Failed to update book')
}

export async function deleteBook(id: string): Promise<void> {
    const res = await fetch(`${API_URL}/${id}`, {method: 'DELETE'})
    if (!res.ok) throw new Error('Failed to delete book')
}
