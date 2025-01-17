// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// UserCommentMovie represents a row from public.user_comment_movie.
type UserCommentMovie struct {
	Username    string // username
	ID          string // id
	Dateandtime int64  // dateandtime
	Comment     string // comment
	Erased      bool   // erased

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserCommentMovie exists in the database.
func (ucm *UserCommentMovie) Exists() bool {
	return ucm._exists
}

// Deleted provides information if the UserCommentMovie has been deleted from the database.
func (ucm *UserCommentMovie) Deleted() bool {
	return ucm._deleted
}

// Insert inserts the UserCommentMovie to the database.
func (ucm *UserCommentMovie) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ucm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.user_comment_movie (` +
		`username, id, comment, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING dateandtime`

	// run query
	XOLog(sqlstr, ucm.Username, ucm.ID, ucm.Comment, ucm.Erased)
	err = db.QueryRow(sqlstr, ucm.Username, ucm.ID, ucm.Comment, ucm.Erased).Scan(&ucm.Dateandtime)
	if err != nil {
		return err
	}

	// set existence
	ucm._exists = true

	return nil
}

// Update updates the UserCommentMovie in the database.
func (ucm *UserCommentMovie) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ucm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ucm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.user_comment_movie SET (` +
		`username, id, comment, erased` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE dateandtime = $5`

	// run query
	XOLog(sqlstr, ucm.Username, ucm.ID, ucm.Comment, ucm.Erased, ucm.Dateandtime)
	_, err = db.Exec(sqlstr, ucm.Username, ucm.ID, ucm.Comment, ucm.Erased, ucm.Dateandtime)
	return err
}

// Save saves the UserCommentMovie to the database.
func (ucm *UserCommentMovie) Save(db XODB) error {
	if ucm.Exists() {
		return ucm.Update(db)
	}

	return ucm.Insert(db)
}

// Upsert performs an upsert for UserCommentMovie.
//
// NOTE: PostgreSQL 9.5+ only
func (ucm *UserCommentMovie) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ucm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.user_comment_movie (` +
		`username, id, dateandtime, comment, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (dateandtime) DO UPDATE SET (` +
		`username, id, dateandtime, comment, erased` +
		`) = (` +
		`EXCLUDED.username, EXCLUDED.id, EXCLUDED.dateandtime, EXCLUDED.comment, EXCLUDED.erased` +
		`)`

	// run query
	XOLog(sqlstr, ucm.Username, ucm.ID, ucm.Dateandtime, ucm.Comment, ucm.Erased)
	_, err = db.Exec(sqlstr, ucm.Username, ucm.ID, ucm.Dateandtime, ucm.Comment, ucm.Erased)
	if err != nil {
		return err
	}

	// set existence
	ucm._exists = true

	return nil
}

// Delete deletes the UserCommentMovie from the database.
func (ucm *UserCommentMovie) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ucm._exists {
		return nil
	}

	// if deleted, bail
	if ucm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.user_comment_movie WHERE dateandtime = $1`

	// run query
	XOLog(sqlstr, ucm.Dateandtime)
	_, err = db.Exec(sqlstr, ucm.Dateandtime)
	if err != nil {
		return err
	}

	// set deleted
	ucm._deleted = true

	return nil
}

// MovieMapping returns the MovieMapping associated with the UserCommentMovie's ID (id).
//
// Generated from foreign key 'movie_user_comment_movie_fk'.
func (ucm *UserCommentMovie) MovieMapping(db XODB) (*MovieMapping, error) {
	return MovieMappingByID(db, ucm.ID)
}

// Useraccount returns the Useraccount associated with the UserCommentMovie's Username (username).
//
// Generated from foreign key 'useraccount_user_comment_movie_fk'.
func (ucm *UserCommentMovie) Useraccount(db XODB) (*Useraccount, error) {
	return UseraccountByUsername(db, ucm.Username)
}

// UserCommentMovieByUsernameIDDateandtime retrieves a row from 'public.user_comment_movie' as a UserCommentMovie.
//
// Generated from index 'pk_user_comment_movie'.
func UserCommentMovieByUsernameIDDateandtime(db XODB, username string, id string, dateandtime int64) (*UserCommentMovie, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, id, dateandtime, comment, erased ` +
		`FROM public.user_comment_movie ` +
		`WHERE username = $1 AND id = $2 AND dateandtime = $3`

	// run query
	XOLog(sqlstr, username, id, dateandtime)
	ucm := UserCommentMovie{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username, id, dateandtime).Scan(&ucm.Username, &ucm.ID, &ucm.Dateandtime, &ucm.Comment, &ucm.Erased)
	if err != nil {
		return nil, err
	}

	return &ucm, nil
}
