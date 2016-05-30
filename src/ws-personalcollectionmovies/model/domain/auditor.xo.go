// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

// Auditor represents a row from public.auditor.
type Auditor struct {
	Username string // username
}

// Root returns the Root associated with the Auditor's Username (username).
//
// Generated from foreign key 'root_auditor_fk'.
func (a *Auditor) Root(db XODB) (*Root, error) {
	return RootByUsername(db, a.Username)
}

// AuditorByUsername retrieves a row from 'public.auditor' as a Auditor.
//
// Generated from index 'pk_auditor'.
func AuditorByUsername(db XODB, username string) (*Auditor, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username ` +
		`FROM public.auditor ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	a := Auditor{}

	err = db.QueryRow(sqlstr, username).Scan(&a.Username)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
