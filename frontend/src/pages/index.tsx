import {Book} from "@/types/book";
import {useEffect, useState} from "react";
import {addBook, deleteBook, getBooks} from "@/data/book";

export default function Home() {
    const [books, setBooks] = useState<Book[]>([])
    const [loading, setLoading] = useState<boolean>(true)
    const [showModal, setShowModal] = useState<boolean>(false)
    const [errorMessage, setErrorMessage] = useState<string>('')
    const [bookToDelete, setBookToDelete] = useState<Book | null>(null)

    const [form, setForm] = useState({
        title: '',
        author: '',
        year: '',
    })

    const handleAddBook = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            await addBook({
                id: '',
                title: form.title,
                author: form.author,
                year: parseInt(form.year),
            })
            await loadBooks()
            setForm({title: '', author: '', year: ''})
            setShowModal(false)
        } catch (err) {
            console.error('Failed to add book:', err)
        }
    }

    const confirmDelete = async () => {
        if (!bookToDelete) return
        try {
            await deleteBook(bookToDelete.id)
            await loadBooks()
            setBookToDelete(null)
        } catch (err) {
            setErrorMessage('Failed to delete book')
        }
    }

    const loadBooks = async () => {
        try {
            setLoading(true)
            setErrorMessage('')
            const data = await getBooks()
            setBooks(data)
        } catch (err: any) {
            setErrorMessage(err.message || 'Something went wrong')
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        loadBooks()
    }, [])

    return (
        <main className="pt-24 px-8">
            <nav
                className="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
                <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                    <a href="https://rzndwb.xyz/" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
                            📚 BookLib
                        </span>
                    </a>
                    <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
                        <button type="button" onClick={() => setShowModal(true)}
                                className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                            Add Book
                        </button>
                        <button data-collapse-toggle="navbar-sticky" type="button"
                                className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                                aria-controls="navbar-sticky" aria-expanded="false">
                            <span className="sr-only">Open main menu</span>
                            <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                                 viewBox="0 0 17 14">
                                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round"
                                      strokeWidth="2" d="M1 1h15M1 7h15M1 13h15"/>
                            </svg>
                        </button>
                    </div>
                    <div className="items-center justify-between hidden w-full md:flex md:w-auto md:order-1"
                         id="navbar-sticky">
                    </div>
                </div>
            </nav>

            <div className="flex overflow-x-auto space-x-4 pb-4">
                {loading ? (
                    <p className="text-center text-gray-600">Loading books...</p>
                ) : (
                    <div className="flex flex-wrap gap-4">
                        {books.map((book) => (
                            <div
                                key={book.id}
                                className="min-w-[250px] max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow-sm dark:bg-gray-800 dark:border-gray-700"
                            >
                                <a href="#">
                                    <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                                        {book.title}
                                    </h5>
                                </a>
                                <p className="mb-3 font-normal text-gray-600 dark:text-gray-400">
                                    by <b>{book.author}</b>, {book.year}
                                </p>

                                <div className="flex flex-wrap gap-2">
                                    <button
                                        onClick={() => setBookToDelete(book)}
                                        className="inline-flex items-center px-3 py-2 text-sm font-medium text-white bg-red-700 rounded-lg hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-800"
                                    >
                                        Delete
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            fill="none"
                                            viewBox="0 0 24 24"
                                            strokeWidth={1.5}
                                            stroke="currentColor"
                                            className="w-4 h-4 ml-2"
                                        >
                                            <path
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                d="M6 18L18 6M6 6l12 12"
                                            />
                                        </svg>
                                    </button>

                                    <a
                                        href="#"
                                        className="inline-flex items-center px-3 py-2 text-sm font-medium text-white bg-yellow-700 rounded-lg hover:bg-yellow-800 focus:ring-4 focus:outline-none focus:ring-yellow-300 dark:bg-yellow-600 dark:hover:bg-yellow-700 dark:focus:ring-yellow-800"
                                    >
                                        Edit
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            fill="none"
                                            viewBox="0 0 24 24"
                                            strokeWidth={1.5}
                                            stroke="currentColor"
                                            className="w-4 h-4 ml-2"
                                        >
                                            <path
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                d="M15.232 5.232l3.536 3.536M16.5 3.75a2.121 2.121 0 113 3L7.5 19.5H3v-4.5L16.5 3.75z"
                                            />
                                        </svg>
                                    </a>

                                    <a
                                        href="#"
                                        className="inline-flex items-center px-3 py-2 text-sm font-medium text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                    >
                                        Detail
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            fill="none"
                                            viewBox="0 0 14 10"
                                            className="w-4 h-4 ml-2"
                                        >
                                            <path
                                                stroke="currentColor"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                strokeWidth="2"
                                                d="M1 5h12m0 0L9 1m4 4L9 9"
                                            />
                                        </svg>
                                    </a>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>

            {showModal && (
                <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
                    <div className="bg-white rounded-lg shadow-xl w-full max-w-md p-6 relative">
                        <h3 className="text-xl font-semibold mb-4">Add New Book</h3>
                        <form onSubmit={handleAddBook} className="space-y-4">
                            <div>
                                <label className="block text-sm font-medium mb-1">Title</label>
                                <input
                                    type="text"
                                    required
                                    value={form.title}
                                    onChange={(e) => setForm({...form, title: e.target.value})}
                                    className="w-full border border-gray-300 p-2 rounded"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium mb-1">Author</label>
                                <input
                                    type="text"
                                    required
                                    value={form.author}
                                    onChange={(e) => setForm({...form, author: e.target.value})}
                                    className="w-full border border-gray-300 p-2 rounded"
                                />
                            </div>
                            <div>
                                <label className="block text-sm font-medium mb-1">Year</label>
                                <input
                                    type="number"
                                    min="1400"
                                    max="2100"
                                    required
                                    value={form.year}
                                    onChange={(e) =>
                                        setForm({...form, year: e.target.value.replace(/\D/, '')})
                                    }
                                    className="w-full border border-gray-300 p-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                                />
                            </div>
                            <div className="flex justify-end space-x-3 pt-2">
                                <button
                                    type="button"
                                    onClick={() => setShowModal(false)}
                                    className="px-4 py-2 rounded border text-gray-700 hover:bg-gray-100"
                                >
                                    Cancel
                                </button>
                                <button
                                    type="submit"
                                    className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                                >
                                    Save
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
            {bookToDelete && (
                <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
                    <div className="bg-white p-6 rounded-lg shadow-md max-w-sm w-full">
                        <h2 className="text-lg font-semibold text-gray-800 mb-2">Confirm Delete</h2>
                        <p className="text-gray-700 mb-4">
                            Are you sure you want to delete <b>{bookToDelete.title}</b>?
                        </p>
                        <div className="flex justify-end space-x-3">
                            <button
                                className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300"
                                onClick={() => setBookToDelete(null)}
                            >
                                Cancel
                            </button>
                            <button
                                className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
                                onClick={confirmDelete}
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                </div>
            )}
            {errorMessage !== '' && (
                <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
                    <div className="bg-white p-6 rounded-lg shadow-md max-w-sm w-full">
                        <h2 className="text-lg font-semibold text-red-600 mb-2">Error</h2>
                        <p className="text-gray-700 mb-4">{errorMessage}</p>
                        <div className="flex justify-end space-x-3">
                            <button
                                className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300"
                                onClick={() => setErrorMessage('')}
                            >
                                Close
                            </button>
                            <button
                                className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                                onClick={loadBooks}
                            >
                                Retry
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </main>
    )
}
