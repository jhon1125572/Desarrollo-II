// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// UserNextviewsMovie represents a row from public.user_nextviews_movies.
type UserNextviewsMovie struct {
	Username    string // username
	ID          string // id
	Dateandtime string  // dateandtime int64
	Erased      bool   // erased

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserNextviewsMovie exists in the database.
func (unm *UserNextviewsMovie) Exists() bool {
	return unm._exists
}

// Deleted provides information if the UserNextviewsMovie has been deleted from the database.
func (unm *UserNextviewsMovie) Deleted() bool {
	return unm._deleted
}

// Insert inserts the UserNextviewsMovie to the database.
func (unm *UserNextviewsMovie) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if unm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.user_nextviews_movies (` +
		`username, id, dateandtime, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, unm.Username, unm.ID, unm.Dateandtime, unm.Erased)
	err = db.QueryRow(sqlstr, unm.Username, unm.ID, unm.Dateandtime, unm.Erased).Scan(&unm.ID)
	if err != nil {
		return err
	}

	// set existence
	unm._exists = true

	return nil
}

// Update updates the UserNextviewsMovie in the database.
func (unm *UserNextviewsMovie) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !unm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if unm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.user_nextviews_movies SET (` +
		`username, dateandtime, erased` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, unm.Username, unm.Dateandtime, unm.Erased, unm.ID)
	_, err = db.Exec(sqlstr, unm.Username, unm.Dateandtime, unm.Erased, unm.ID)
	return err
}

// Save saves the UserNextviewsMovie to the database.
func (unm *UserNextviewsMovie) Save(db XODB) error {
	if unm.Exists() {
		return unm.Update(db)
	}

	return unm.Insert(db)
}

// Upsert performs an upsert for UserNextviewsMovie.
//
// NOTE: PostgreSQL 9.5+ only
func (unm *UserNextviewsMovie) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if unm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.user_nextviews_movies (` +
		`username, id, dateandtime, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`username, id, dateandtime, erased` +
		`) = (` +
		`EXCLUDED.username, EXCLUDED.id, EXCLUDED.dateandtime, EXCLUDED.erased` +
		`)`

	// run query
	XOLog(sqlstr, unm.Username, unm.ID, unm.Dateandtime, unm.Erased)
	_, err = db.Exec(sqlstr, unm.Username, unm.ID, unm.Dateandtime, unm.Erased)
	if err != nil {
		return err
	}

	// set existence
	unm._exists = true

	return nil
}

// Delete deletes the UserNextviewsMovie from the database.
func (unm *UserNextviewsMovie) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !unm._exists {
		return nil
	}

	// if deleted, bail
	if unm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.user_nextviews_movies WHERE id = $1`

	// run query
	XOLog(sqlstr, unm.ID)
	_, err = db.Exec(sqlstr, unm.ID)
	if err != nil {
		return err
	}

	// set deleted
	unm._deleted = true

	return nil
}

// MovieMapping returns the MovieMapping associated with the UserNextviewsMovie's ID (id).
//
// Generated from foreign key 'movie_user_nextviews_movies_fk'.
func (unm *UserNextviewsMovie) MovieMapping(db XODB) (*MovieMapping, error) {
	return MovieMappingByID(db, unm.ID)
}

// Useraccount returns the Useraccount associated with the UserNextviewsMovie's Username (username).
//
// Generated from foreign key 'useraccount_user_nextviews_movies_fk'.
func (unm *UserNextviewsMovie) Useraccount(db XODB) (*Useraccount, error) {
	return UseraccountByUsername(db, unm.Username)
}

// UserNextviewsMovieByUsernameID retrieves a row from 'public.user_nextviews_movies' as a UserNextviewsMovie.
//
// Generated from index 'pk_user_next_views_movie'.
func UserNextviewsMovieByUsernameID(db XODB, username string, id string) (*UserNextviewsMovie, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, id, dateandtime, erased ` +
		`FROM public.user_nextviews_movies ` +
		`WHERE username = $1 AND id = $2`

	// run query
	XOLog(sqlstr, username, id)
	unm := UserNextviewsMovie{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username, id).Scan(&unm.Username, &unm.ID, &unm.Dateandtime, &unm.Erased)
	if err != nil {
		return nil, err
	}

	return &unm, nil
}
