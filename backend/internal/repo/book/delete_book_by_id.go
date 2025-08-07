package book

import "context"

func (r *repo) DeleteBook(ctx context.Context, id string) error {
	query := `DELETE FROM books WHERE id = $1`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}
