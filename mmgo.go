package mmgo

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// mgo object
// 储存数据库访问的上下文
type MgoDB struct {
	CurrentDBname       string
	CurrentDBcollection string

	inited bool
}

// mgo.Session
// 每次连接都会克隆这个对象，代表新的DB连接
var mgoSession *mgo.Session

// 全局上下文，本来想写成可以默认对某个数据库进行访问
// 逻辑过于复杂和冗余，删去了
// var gloabal_ctx = NewEmptyCtx()

// create empty mgodb context
// 几乎无用，理由同上
// func NewEmptyCtx() *MgoDB {
// 	ctx := &MgoDB{
// 		CurrentDBcollection: "",
// 		CurrentDBname:       "",
// 	}
// 	ctx.init()
// 	return ctx
// }

// create new context
// 参数分别是数据库key值，和集合名称
func NewCtx(name, coll string) *MgoDB {
	ctx := &MgoDB{
		CurrentDBcollection: coll,
		CurrentDBname:       name,
		inited:              false,
	}
	if ctx.init() == nil {
		ctx.inited = true
	}
	return ctx
}

// init context
// 保证连接有效，产生error表示无法正常连接数据库
func (m *MgoDB) init() error {
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
		log.Printf("Create Session: %s\n", err)
	}
	m.inited = true
	return err
}

// change db context db name
// 修改数据库名称，除特殊不建议使用
func (m *MgoDB) SetName(name string) {
	m.CurrentDBname = name
}

// change db context collection name
// 修改数据库名称，除特殊不建议使用
func (m *MgoDB) SetCollection(name string) {
	m.CurrentDBcollection = name
}

// all of first call
// 所有调用过程均会经过这里
func (m *MgoDB) connect() (*mgo.Session, *mgo.Collection, error) {
	db := m.CurrentDBname
	collection := m.CurrentDBcollection
	//
	if !m.inited {
		if err := m.init(); err != nil {
			return nil, nil, err
		}
	}
	ms := mgoSession.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c, nil
}

// get *mgo.Database, Provide a call to some special function
func (m *MgoDB) getDb() (*mgo.Session, *mgo.Database, error) {
	db := m.CurrentDBname
	//
	if !m.inited {
		if err := m.init(); err != nil {
			return nil, nil, err
		}
	}
	ms := mgoSession.Copy()
	return ms, ms.DB(db), nil
}

// return mgo context context state
// 判断能否与数据库连接
func (m *MgoDB) IsLive() bool {
	if m.inited {
		return true
	}
	if m.init() == nil {
		return true
	}
	return false
}

// return collection ture/false rmpty
// 判断是否在操作一个空集合
func (m *MgoDB) IsEmpty() bool {
	ms, c, err := m.connect()
	if err != nil {
		return true
	}
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

// 根据查询返回个数
func (m *MgoDB) Count(query interface{}) (int, error) {
	ms, c, err := m.connect()
	if err != nil {
		return 0, err
	}
	defer ms.Close()
	return c.Find(query).Count()
}

// 插入操作，可一次插入多个文档
func (m *MgoDB) Insert(docs ...interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Insert(docs...)
}

func (m *MgoDB) FindOne(query, selector, result interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func (m *MgoDB) FindAll(query, selector, result interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func (m *MgoDB) FindPage(page, limit int, query, selector, result interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func (m *MgoDB) FindPageReverse(page, limit int, query, selector, result interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Find(query).Sort("-_id").Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func (m *MgoDB) FindIter(query interface{}) (*mgo.Iter, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()

	return c.Find(query).Iter(), nil
}

func (m *MgoDB) Update(selector, update interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Update(selector, update)
}

func (m *MgoDB) Upsert(selector, update interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	_, err = c.Upsert(selector, update)
	return err
}

func (m *MgoDB) UpdateAll(selector, update interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	_, err = c.UpdateAll(selector, update)
	return err
}

func (m *MgoDB) Remove(selector interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	return c.Remove(selector)
}

func (m *MgoDB) RemoveAll(selector interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()

	_, err = c.RemoveAll(selector)
	return err
}

func (m *MgoDB) RemoveRepeat(selector interface{}) {
	var resarr []interface{}
	err := m.FindAll(selector, nil, &resarr)
	if err != nil || len(resarr) == 1 {
		return
	}
	// for i, _ := range resarr {
	// 	if i == len(resarr)-1 {
	// 		break
	// 	}
	// 	err = Remove(selector)
	// }
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
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func (m *MgoDB) BulkRemove(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func (m *MgoDB) BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func (m *MgoDB) BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func (m *MgoDB) PipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
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
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func (m *MgoDB) PipeIter(pipeline interface{}, allowDiskUse bool) (*mgo.Iter, error) {
	ms, c, err := m.connect()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}

	return pipe.Iter(), nil

}

func (m *MgoDB) Explain(pipeline, result interface{}) error {
	ms, c, err := m.connect()
	if err != nil {
		return err
	}
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}
func (m *MgoDB) GridFSCreate(prefix, name string) (*mgo.GridFile, error) {
	ms, d, err := m.getDb()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Create(name)
}

func (m *MgoDB) GridFSFindOne(prefix string, query, result interface{}) error {
	ms, d, err := m.getDb()
	if err != nil {
		return err
	}
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).One(result)
}

func (m *MgoDB) GridFSFindAll(prefix string, query, result interface{}) error {
	ms, d, err := m.getDb()
	if err != nil {
		return err
	}
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).All(result)
}

func (m *MgoDB) GridFSOpen(prefix, name string) (*mgo.GridFile, error) {
	ms, d, err := m.getDb()
	if err != nil {
		return nil, err
	}
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Open(name)
}

func (m *MgoDB) GridFSRemove(prefix, name string) error {
	ms, d, err := m.getDb()
	if err != nil {
		return err
	}
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Remove(name)
}
