package store

import (
	"log"

	"github.com/devinmcgloin/morph/src/api/generator"
	"github.com/devinmcgloin/morph/src/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddSrc takes a source struct and adds it to the db in the proper image
func (ds *MgoStore) AddSrc(imageID bson.ObjectId, src model.ImgSource) error {
	session := ds.session.Copy()
	defer session.Close()

	c := session.DB(dbname).C("images")

	err := c.UpdateId(imageID, bson.M{"$push": bson.M{"sources": src}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// AddUser takes the given user and adds it to the Database.
func (ds *MgoStore) AddUser(user model.User) error {
	session := ds.session.Copy()
	defer session.Close()

	c := session.DB(dbname).C("users")

	err := c.Insert(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ds *MgoStore) AddUserImage(user bson.ObjectId, images ...bson.ObjectId) error {
	session := ds.session.Copy()
	defer session.Close()

	c := session.DB(dbname).C("users")

	var imgID []bson.ObjectId

	for img := range images {
		change := mgo.Change{
			Update:    bson.M{"$push": bson.M{"images": img}},
			ReturnNew: true,
		}

		_, err := c.FindId(user).Apply(change, &imgID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (ds *MgoStore) AddImg(image model.Image) error {
	session := ds.session.Copy()
	defer session.Close()

	log.Println(image)
	c := session.DB(dbname).C("images")

	err := c.Insert(image)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// UpdateImage matches the DB image with the one passed in.
func (ds *MgoStore) UpdateImage(image model.Image) error {
	session := ds.session.Copy()
	defer session.Close()

	c := session.DB(dbname).C("images")

	err := c.UpdateId(image.ID, image)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ds *MgoStore) GetNewImageShortCode() string {
	randTitle := generator.RandString(12)
	if ds.ExistsImageShortCode(randTitle) {
		randTitle = generator.RandString(12)
	}
	return randTitle
}

func (ds *MgoStore) GetNewAlbumShortCode() string {
	randTitle := generator.RandString(12)
	if ds.ExistsAlbumShortCode(randTitle) {
		randTitle = generator.RandString(12)
	}
	return randTitle
}