package mmgo

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type MgoDB struct {
	CurrentDBname       string
	CurrentDBcollection string
}

var mgoSession *mgo.Session

// var gloabal_ctx = NewEmptyCtx()

func NewEmptyCtx() *MgoDB {
	ctx := &MgoDB{
		CurrentDBcollection: "",
		CurrentDBname:       "",
	}
	ctx.init()
	return ctx
}

func NewCtx(name, coll string) *MgoDB {
	ctx := &MgoDB{
		CurrentDBcollection: name,
		CurrentDBname:       coll,
	}
	ctx.init()
	return ctx
}

func (m *MgoDB) init() {
	var err error
	if mgoAuthUser != "" && mgoAuthPassword != "" {
		dialInfo := &mgo.DialInfo{
			Addrs:     []string{mgoHost},
			Timeout:   mgoTimeout,
			Source:    mgoAuthdb,
			Username:  mgoAuthUser,
			Password:  mgoAuthPassword,
			PoolLimit: PoolLimit,
		}
		mgoSession, err = mgo.DialWithInfo(dialInfo)
	} else {
		mgoSession, err = mgo.Dial(mgoHost)
	}
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
}

func (m *MgoDB) SetName(name string) {
	m.CurrentDBname = name
}

func (m *MgoDB) SetCollection(name string) {
	m.CurrentDBcollection = name
}

func (m *MgoDB) connect() (*mgo.Session, *mgo.Collection) {
	db := m.CurrentDBname
	collection := m.CurrentDBcollection
	//
	ms := mgoSession.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func (m *MgoDB) getDb() (*mgo.Session, *mgo.Database) {
	db := m.CurrentDBname
	//
	ms := mgoSession.Copy()
	return ms, ms.DB(db)
}

func (m *MgoDB) IsEmpty() bool {
	ms, c := m.connect()
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func (m *MgoDB) Count(query interface{}) (int, error) {
	ms, c := m.connect()
	defer ms.Close()
	return c.Find(query).Count()
}

func (m *MgoDB) Insert(docs ...interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Insert(docs...)
}

func (m *MgoDB) FindOne(query, selector, result interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func (m *MgoDB) FindAll(query, selector, result interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func (m *MgoDB) FindPage(page, limit int, query, selector, result interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func (m *MgoDB) FindIter(query interface{}) *mgo.Iter {
	ms, c := m.connect()
	defer ms.Close()

	return c.Find(query).Iter()
}

func (m *MgoDB) Update(selector, update interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Update(selector, update)
}

func (m *MgoDB) Upsert(selector, update interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func (m *MgoDB) UpdateAll(selector, update interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func (m *MgoDB) Remove(selector interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	return c.Remove(selector)
}

func (m *MgoDB) RemoveAll(selector interface{}) error {
	ms, c := m.connect()
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

func (m *MgoDB) RemoveRepeat(selector interface{}) {
	var resarr []interface{}
	err := m.FindAll(selector, nil, &resarr)
	if err != nil || len(resarr) == 1 {
		return
	}
	m.RemoveAll(selector)
	m.Insert(resarr[0])
}

func (m *MgoDB) RemoveRepeatByKey(key string) {
	var resmapArr []map[string]interface{}
	err := m.FindAll(nil, nil, &resmapArr)
	if err != nil || len(resmapArr) == 1 {
		return
	}
	var taskArr []interface{}
	for _, resmap := range resmapArr {
		for k, v := range resmap {
			if k == key {
				taskArr = append(taskArr, v)
			}
		}
	}
	for _, v := range taskArr {
		m.RemoveRepeat(bson.M{key: v})
	}
}

//insert one or multi documents
func (m *MgoDB) BulkInsert(docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func (m *MgoDB) BulkRemove(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func (m *MgoDB) BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := m.connect()
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func (m *MgoDB) PipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := m.connect()
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func (m *MgoDB) PipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := m.connect()
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func (m *MgoDB) PipeIter(pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	ms, c := m.connect()
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}

	return pipe.Iter()

}

func (m *MgoDB) Explain(pipeline, result interface{}) error {
	ms, c := m.connect()
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}
func (m *MgoDB) GridFSCreate(prefix, name string) (*mgo.GridFile, error) {
	ms, d := m.getDb()
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Create(name)
}

func (m *MgoDB) GridFSFindOne(prefix string, query, result interface{}) error {
	ms, d := m.getDb()
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).One(result)
}

func (m *MgoDB) GridFSFindAll(prefix string, query, result interface{}) error {
	ms, d := m.getDb()
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).All(result)
}

func (m *MgoDB) GridFSOpen(prefix, name string) (*mgo.GridFile, error) {
	ms, d := m.getDb()
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Open(name)
}

func (m *MgoDB) GridFSRemove(prefix, name string) error {
	ms, d := m.getDb()
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Remove(name)
}
