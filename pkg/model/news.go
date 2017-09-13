package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//News type
type News struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

// var (
// 	newsStorage []News
// 	mutexNews   sync.RWMutex
// )

// func generateID() string {
// 	buf := make([]byte, 16)
// 	rand.Read(buf)
// 	return base64.StdEncoding.EncodeToString(buf)
// }

//CreateNews inserts news into database
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	// news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil
	// mutexNews.Lock()
	// defer mutexNews.Unlock()
	// newsStorage = append(newsStorage, news)
}

//ListNews list all news from database
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil
	// mutexNews.RLock()
	// defer mutexNews.RUnlock()
	// r := make([]*News, len(newsStorage))
	// for i := range newsStorage {
	// 	n := newsStorage[i]
	// 	r[i] = &n
	// }
	// return r
}

//GetNews retrives news from database
func GetNews(id string) (*News, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
	// mutexNews.RLock()
	// defer mutexNews.RUnlock()
	// for _, news := range newsStorage {
	// 	if news.ID == id {
	// 		n := news
	// 		return &n
	// 	}
	// }
	// return nil
}

//DeleteNews delete a news from database
func DeleteNews(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
	// mutexNews.Lock()
	// defer mutexNews.Unlock()
	// for i, news := range newsStorage {
	// 	if news.ID == id {
	// 		newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
	// 		return
	// 	}
	// }
}

//UpdateNews update news
func UpdateNews(news *News) error {
	if news.ID == "" {
		return fmt.Errorf("required id to update")
	}
	news.UpdatedAt = time.Now()
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").UpdateId(news.ID, news)
	if err != nil {
		return err
	}
	return nil
}
