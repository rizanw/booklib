import {Book} from "@/types/book";

const API_URL = 'http://0.0.0.0:8080/api/v1/books'

export async function getBooks(): Promise<Book[]> {
    const res = await fetch(API_URL)
    if (!res.ok) throw new Error('Failed to fetch books')
    return res.json()
}

export async function getBookById(id: string): Promise<Book> {
    const res = await fetch(`${API_URL}/${id}`)
    if (!res.ok) throw new Error('Failed to fetch book')
    return res.json()
}

export async function updateBook(book: Book): Promise<Book> {
    const res = await fetch(API_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(book),
    })
    if (!res.ok) throw new Error('Failed to add book')
    return res.json()
}

export async function deleteBook(id: number): Promise<void> {
    const res = await fetch(`${API_URL}/${id}`, {method: 'DELETE'})
    if (!res.ok) throw new Error('Failed to delete book')
}
