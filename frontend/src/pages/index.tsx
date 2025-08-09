import Link from 'next/link'
import {Book} from "@/types/book";
import {useEffect, useState} from "react";
import {getBooks} from "@/data/book";

export default function Home() {
    const [books, setBooks] = useState<Book[]>([])
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const load = async () => {
            try {
                const data = await getBooks()
                setBooks(data)
            } catch (err) {
                console.error(err)
            } finally {
                setLoading(false)
            }
        }
        load()
    }, [])

    return (
        <main className="p-6">
            <nav
                className="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
                <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                    <a href="https://rzndwb.xyz/" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
                            📚 BookLib
                        </span>
                    </a>
                    <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
                        <button type="button" data-modal-target="default-modal" data-modal-toggle="default-modal"
                                className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                            Add Book
                        </button>
                        <button data-collapse-toggle="navbar-sticky" type="button"
                                className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                                aria-controls="navbar-sticky" aria-expanded="false">
                            <span className="sr-only">Open main menu</span>
                            <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                                 viewBox="0 0 17 14">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                                      stroke-width="2" d="M1 1h15M1 7h15M1 13h15"/>
                            </svg>
                        </button>
                    </div>
                    <div className="items-center justify-between hidden w-full md:flex md:w-auto md:order-1"
                         id="navbar-sticky">
                    </div>
                </div>
            </nav>

            <ul className="space-y-4">
                {loading ? (
                    <p className="text-center text-gray-600">Loading books...</p>
                ) : (
                    <ul className="space-y-4 max-w-3xl mx-auto mt-6">
                        {books.map((book) => (
                            <li key={book.id}
                                className="bg-white border p-4 rounded shadow flex justify-between items-start">
                                <div>
                                    <h2 className="text-xl font-semibold">{book.title}</h2>
                                    <p className="text-gray-600">by {book.author}</p>
                                    <p className="text-sm text-gray-500 mt-1">{book.year}</p>
                                </div>
                                <div className="space-x-3 text-sm">
                                    <Link href={`/${book.id}`} className="text-blue-600 hover:underline">View</Link>
                                    <Link href={`/${book.id}/edit`}
                                          className="text-yellow-500 hover:underline">Edit</Link>
                                    <button
                                        onClick={() => function () {
                                            // todo: delete book
                                        }}
                                        className="text-red-500 hover:underline"
                                    >
                                        Delete
                                    </button>
                                </div>
                            </li>
                        ))}
                    </ul>
                )}
            </ul>

        </main>
    )
}
