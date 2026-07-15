package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets table
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes.)
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded conncetion pool to execute the statement.
	// The first parameter is the SQL statement, followed by the values for the placeholder
	// parameters: title, content and expiry in that order. This method returns a
	// sq.Result type, which contains some basic information about what happened when the
	// statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() on the result to get the ID of our newly
	// inserted row in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Write the SQL statement we want to execute. Again, I've split it over two lines
	// for readability.
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use the QueryRow() method on the connection pool to execute our SQL statement,
	// passing in the untrusted id variable as the value for the placeholder parameter.
	// THis returns a pointer to a sql.Row value which holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	// Initialize a new zeroed Snippet struct.
	var s Snippet

	// Use row.Scan() to copy the values from each field in the sql.Row to the corresponding
	// field in the Snippet struct. Notice that the arguments to row.Scan are *pointers* to the
	// place you want to copy the data into, and the number of arguments must be exactly the same
	// as the number of columns returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a sql.ErrNoRows error.
		// We use the errors.Is() function to check for that error specifically, and return
		// our own ErrNoRecord error instead.
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	// In practice, you can shorten the code slightly by leveraging the fact that errors from
	// DB.QueryRow() are deferred until Scan() is called. It makes no functional difference, but
	// if you want it's perfectly OK to rewrite the code to look something like this.
	//err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	//if errors.Is(err, sql.ErrNoRows) {
	//	return Snippet{}, ErrNoRecord
	//} else {
	//	return Snippet{}, err
	//}

	// If everything went OK, then return the filled Snippet struct.
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Use the Query() method on the connection pool to execute our SQL statement.
	// This returns a sql.Rows resultset containing the result of our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is always properly closed
	// before the Latest() method returns. This defer statement should come *after* you
	// check for an error from the Query() method. Otherwise, if Query() returns an error,
	// you'll get a panic trying to close a nil resultset.
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	var snippets []Snippet

	// Use rows.Next to iterate through the rows in the resultset. This prepares the first
	// (and then each subsequent) row to be acted on by the rows.Scan() method. If iteration
	// over all the rows completes then the resultset automatically closes itself and frees
	// up the underlying database connection.
	for rows.Next() {
		// Create a new zero value Snippet struct.
		var s Snippet
		// Use rows.Scan() to copy the values from each field in the row to the new Snippet
		// struct that we created. Again, the arguments to row.Scan() must be pointers to the
		// place you want to copy the data into, and the number of arguments must be exactly
		// the same as the number of columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		snippets = append(snippets, s)
	}

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any error that
	// was encountered during the iteration. It's important to call this don't assume the
	// iteration completed successfully over the entire result set.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return snippets, nil
}
