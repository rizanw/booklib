import {Book} from "@/types/book";
import {apiFetch} from "@/utils/api";

const API_URL = `${process.env.NEXT_PUBLIC_BOOKLIB_API_BASE_URL}/api/v1/books`;

interface ApiResponse<T> {
    data: T;
    error: string;
    status: string;
}

export async function getBooks(): Promise<Book[]> {
    const res = await apiFetch(API_URL, {method: 'GET'});
    const json: ApiResponse<Book[]> = res.data;

    if (json.status !== 'success') {
        throw new Error(`API error: ${json.error}`);
    }

    return json.data || [];
}

export async function getBookById(id: string): Promise<Book> {
    const res = await apiFetch(`${API_URL}/${id}`, {method: 'GET'});
    const json: ApiResponse<Book> = res.data;
    return json.data;
}

export async function addBook(book: Book) {
    const res = await apiFetch(API_URL, {
        method: 'POST',
        data: book, // Axios auto-serializes JSON
    });
    const json: ApiResponse<Book> = res.data;
    if (json.status !== 'success') {
        throw new Error(`API error: ${json.error}`);
    }
    return json.data;
}

export async function updateBook(book: Book) {
    const res = await apiFetch(`${API_URL}/${book.id}`, {
        method: 'PUT',
        data: book,
    });
    if (res.status >= 400) throw new Error('Failed to update book');
}

export async function deleteBook(id: string): Promise<void> {
    const res = await apiFetch(`${API_URL}/${id}`, {method: 'DELETE'});
    if (res.status >= 400) throw new Error('Failed to delete book');
}
