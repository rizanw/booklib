import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {Book} from "@/types/book";
import {getBookById} from "@/data/book";

export default function BookDetailPage() {
    const router = useRouter();
    const {id} = router.query;

    const [book, setBook] = useState<Book | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!id) return;

        async function fetchBook() {
            setLoading(true);
            setError(null);

            try {
                const fetchedBook = await getBookById(id as string);
                if (!fetchedBook) {
                    setError("Book not found.");
                } else {
                    setBook(fetchedBook);
                }
            } catch {
                setError("Failed to load book details.");
            } finally {
                setLoading(false);
            }
        }

        fetchBook();
    }, [id]);

    if (loading) {
        return (
            <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
                <div className="bg-white p-6 rounded shadow-md dark:bg-gray-800 dark:text-white">
                    <li className="flex items-center">
                        <div role="status">
                            <svg
                                aria-hidden="true"
                                className="w-4 h-4 me-2 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
                                viewBox="0 0 100 101"
                                fill="none"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                <path
                                    d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                    fill="currentColor"
                                />
                                <path
                                    d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                    fill="currentFill"
                                />
                            </svg>
                            <span className="sr-only">Loading...</span>
                        </div>
                        Loading book details...
                    </li>
                </div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="min-h-screen flex items-center justify-center text-red-600 text-lg">
                {error}
            </div>
        );
    }

    if (!book) {
        return null; // or a fallback UI if needed
    }

    return (
        <main className="min-h-screen px-8 pt-24">
            <nav
                className="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
                <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                    <a href="https://rzndwb.xyz/" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
                            📚 BookLib
                        </span>
                    </a>
                    <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
                    </div>
                    <div className="items-center justify-between hidden w-full md:flex md:w-auto md:order-1"
                         id="navbar-sticky">
                    </div>
                </div>
            </nav>

            <div className="max-w-2xl mx-auto bg-white p-6 rounded-lg shadow-md dark:bg-gray-800">
                <h1 className="text-3xl font-bold text-gray-900 dark:text-white mb-4">{book.title}</h1>
                <p className="text-lg text-gray-700 dark:text-gray-300 mb-2">
                    <strong>Author:</strong> {book.author}
                </p>
                <p className="text-lg text-gray-700 dark:text-gray-300 mb-2">
                    <strong>Published Year:</strong> {book.year}
                </p>
                <a
                    href="/"
                    className="inline-block mt-4 px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700"
                >
                    ← Back to Home
                </a>
            </div>
        </main>
    );
}
