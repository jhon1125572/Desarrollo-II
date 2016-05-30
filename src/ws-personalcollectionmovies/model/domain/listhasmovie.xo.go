// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// ListHasMovie represents a row from public.list_has_movie.
type ListHasMovie struct {
	Username string // username
	IDList   string // id_list
	ImdbID   string // imdb_id
	Erased   bool   // erased

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ListHasMovie exists in the database.
func (lhm *ListHasMovie) Exists() bool {
	return lhm._exists
}

// Deleted provides information if the ListHasMovie has been deleted from the database.
func (lhm *ListHasMovie) Deleted() bool {
	return lhm._deleted
}

// Insert inserts the ListHasMovie to the database.
func (lhm *ListHasMovie) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if lhm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.list_has_movie (` +
		`username, id_list, erased` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING imdb_id`

	// run query
	XOLog(sqlstr, lhm.Username, lhm.IDList, lhm.Erased)
	err = db.QueryRow(sqlstr, lhm.Username, lhm.IDList, lhm.Erased).Scan(&lhm.ImdbID)
	if err != nil {
		return err
	}

	// set existence
	lhm._exists = true

	return nil
}

// Update updates the ListHasMovie in the database.
func (lhm *ListHasMovie) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !lhm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if lhm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.list_has_movie SET (` +
		`username, id_list, erased` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE imdb_id = $4`

	// run query
	XOLog(sqlstr, lhm.Username, lhm.IDList, lhm.Erased, lhm.ImdbID)
	_, err = db.Exec(sqlstr, lhm.Username, lhm.IDList, lhm.Erased, lhm.ImdbID)
	return err
}

// Save saves the ListHasMovie to the database.
func (lhm *ListHasMovie) Save(db XODB) error {
	if lhm.Exists() {
		return lhm.Update(db)
	}

	return lhm.Insert(db)
}

// Upsert performs an upsert for ListHasMovie.
//
// NOTE: PostgreSQL 9.5+ only
func (lhm *ListHasMovie) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if lhm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.list_has_movie (` +
		`username, id_list, imdb_id, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (imdb_id) DO UPDATE SET (` +
		`username, id_list, imdb_id, erased` +
		`) = (` +
		`EXCLUDED.username, EXCLUDED.id_list, EXCLUDED.imdb_id, EXCLUDED.erased` +
		`)`

	// run query
	XOLog(sqlstr, lhm.Username, lhm.IDList, lhm.ImdbID, lhm.Erased)
	_, err = db.Exec(sqlstr, lhm.Username, lhm.IDList, lhm.ImdbID, lhm.Erased)
	if err != nil {
		return err
	}

	// set existence
	lhm._exists = true

	return nil
}

// Delete deletes the ListHasMovie from the database.
func (lhm *ListHasMovie) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !lhm._exists {
		return nil
	}

	// if deleted, bail
	if lhm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.list_has_movie WHERE imdb_id = $1`

	// run query
	XOLog(sqlstr, lhm.ImdbID)
	_, err = db.Exec(sqlstr, lhm.ImdbID)
	if err != nil {
		return err
	}

	// set deleted
	lhm._deleted = true

	return nil
}

// List returns the List associated with the ListHasMovie's Username (username).
//
// Generated from foreign key 'list_list_has_movie_fk'.
func (lhm *ListHasMovie) List(db XODB) (*List, error) {
	return ListByIDList(db, lhm.Username)
}

// MovieMapping returns the MovieMapping associated with the ListHasMovie's ImdbID (imdb_id).
//
// Generated from foreign key 'movie_list_has_movie_fk'.
func (lhm *ListHasMovie) MovieMapping(db XODB) (*MovieMapping, error) {
	return MovieMappingByImdbID(db, lhm.ImdbID)
}

// ListHasMovieByUsernameIDListImdbID retrieves a row from 'public.list_has_movie' as a ListHasMovie.
//
// Generated from index 'pk_list_has_movie'.
func ListHasMovieByUsernameIDListImdbID(db XODB, username string, idList string, imdbID string) (*ListHasMovie, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, id_list, imdb_id, erased ` +
		`FROM public.list_has_movie ` +
		`WHERE username = $1 AND id_list = $2 AND imdb_id = $3`

	// run query
	XOLog(sqlstr, username, idList, imdbID)
	lhm := ListHasMovie{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username, idList, imdbID).Scan(&lhm.Username, &lhm.IDList, &lhm.ImdbID, &lhm.Erased)
	if err != nil {
		return nil, err
	}

	return &lhm, nil
}
