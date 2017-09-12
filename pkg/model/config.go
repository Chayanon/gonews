package model

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session

const database = "gonews"

func Init(mongoURL string) error {
	var err error
	mongoSession, err = mgo.Dial(mongoURL)
	//return err
	if err != nil {
		return err
	}
	return nil
}
