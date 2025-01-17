// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// List represents a row from public.list.
type List struct {
	Username    string // username
	IDList      string // id_list
	Name        string // name
	Description string // description
	Erased      bool   // erased

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the List exists in the database.
func (l *List) Exists() bool {
	return l._exists
}

// Deleted provides information if the List has been deleted from the database.
func (l *List) Deleted() bool {
	return l._deleted
}

// Insert inserts the List to the database.
func (l *List) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if l._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.list (` +
		`username, name, description, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id_list`

	// run query
	XOLog(sqlstr, l.Username, l.Name, l.Description, l.Erased)
	err = db.QueryRow(sqlstr, l.Username, l.Name, l.Description, l.Erased).Scan(&l.IDList)
	if err != nil {
		return err
	}

	// set existence
	l._exists = true

	return nil
}

// Update updates the List in the database.
func (l *List) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if l._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.list SET (` +
		`username, name, description, erased` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id_list = $5`

	// run query
	XOLog(sqlstr, l.Username, l.Name, l.Description, l.Erased, l.IDList)
	_, err = db.Exec(sqlstr, l.Username, l.Name, l.Description, l.Erased, l.IDList)
	return err
}

// Save saves the List to the database.
func (l *List) Save(db XODB) error {
	if l.Exists() {
		return l.Update(db)
	}

	return l.Insert(db)
}

// Upsert performs an upsert for List.
//
// NOTE: PostgreSQL 9.5+ only
func (l *List) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if l._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.list (` +
		`username, id_list, name, description, erased` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id_list) DO UPDATE SET (` +
		`username, id_list, name, description, erased` +
		`) = (` +
		`EXCLUDED.username, EXCLUDED.id_list, EXCLUDED.name, EXCLUDED.description, EXCLUDED.erased` +
		`)`

	// run query
	XOLog(sqlstr, l.Username, l.IDList, l.Name, l.Description, l.Erased)
	_, err = db.Exec(sqlstr, l.Username, l.IDList, l.Name, l.Description, l.Erased)
	if err != nil {
		return err
	}

	// set existence
	l._exists = true

	return nil
}

// Delete deletes the List from the database.
func (l *List) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return nil
	}

	// if deleted, bail
	if l._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.list WHERE id_list = $1`

	// run query
	XOLog(sqlstr, l.IDList)
	_, err = db.Exec(sqlstr, l.IDList)
	if err != nil {
		return err
	}

	// set deleted
	l._deleted = true

	return nil
}

// Useraccount returns the Useraccount associated with the List's Username (username).
//
// Generated from foreign key 'useraccount_list_fk'.
func (l *List) Useraccount(db XODB) (*Useraccount, error) {
	return UseraccountByUsername(db, l.Username)
}

// ListByUsernameIDList retrieves a row from 'public.list' as a List.
//
// Generated from index 'pk_list'.
func ListByUsernameIDList(db XODB, username string, idList string) (*List, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, id_list, name, description, erased ` +
		`FROM public.list ` +
		`WHERE username = $1 AND id_list = $2`

	// run query
	XOLog(sqlstr, username, idList)
	l := List{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username, idList).Scan(&l.Username, &l.IDList, &l.Name, &l.Description, &l.Erased)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
