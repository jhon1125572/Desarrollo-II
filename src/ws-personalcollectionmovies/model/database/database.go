package database

import (
    "github.com/knq/dburl"
    _ "github.com/lib/pq"
    "ws-personalcollectionmovies/log"
    "ws-personalcollectionmovies/model/domain"
    )
    
func OpenDataBase() (db domain.XODB) {
    // open database
	db, err := dburl.Open("pgsql://sotilstfjbrplf:pr-JQDyUOTCcm8FzQRmpm8a_g0@ec2-54-243-62-211.compute-1.amazonaws.com/d59ne6oopddc7q")
	if err != nil {
		log.Error(err)
	}
	return db
}