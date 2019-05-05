# mmgo
mgo package package.

# Usage
```golang
// NewCtx - (dbName ,collectionName )
var Coll1 = mmgo.NewCtx("test", "coll1")

func main(){
    Coll1.Insert(&map[string]interface{}{
		"name": "alice",
		"age":  12,
		"info": "996",
    })
    Coll1.Insert(&map[string]interface{}{
		"name": "bob",
		"age":  20,
		"info": "996",
    })
    
    var collection = []map[string]interface{}
    err := testColl.FindAll(nil, nil, &collection)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("[LOG] res = %v\n", collection)
}
```

# API USAGE
- [NewCtx(name, coll string) *MgoDB](#NewCtx)
- [SetName(name string)](#SetName_SetCollection)
- [IsEmpty() bool](#IsEmpty)
- [Count(query interface{}) (int, error)](#Count)
- [SetName(name string) ](#SetName)
- [SetCollection(name string) ](#SetCollection)
- [connect() (*mgo.Session, *mgo.Collection) ](#connect)
- [getDb() (*mgo.Session, *mgo.Database) ](#getDb)
- [IsEmpty() bool ](#IsEmpty)
- [Count(query interface{}) (int, error) ](#Count)
- [Insert(docs ...interface{}) error ](#Insert)
- [FindOne(query, selector, result interface{}) error ](#FindOne)
- [FindAll(query, selector, result interface{}) error ](#FindAll)
- [FindPage(page, limit int, query, selector, result interface{}) error ](#FindPage)
- [FindIter(query interface{}) *mgo.Iter ](#FindIter)
- [Update(selector, update interface{}) error ](#Update)
- [Upsert(selector, update interface{}) error ](#Upsert)
- [UpdateAll(selector, update interface{}) error ](#UpdateAll)
- [Remove(selector interface{}) error ](#Remove)
- [RemoveAll(selector interface{}) error ](#RemoveAll)
- [RemoveRepeat(selector interface{}) ](#RemoveRepeat)
- [RemoveRepeatByKey(key string) ](#RemoveRepeatByKey)
- [BulkInsert(docs ...interface{}) (*mgo.BulkResult, error) ](#BulkInsert)
- [BulkRemove(selector ...interface{}) (*mgo.BulkResult, error) ](#BulkRemove)
- [BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error) ](#BulkRemoveAll)
- [BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error) ](#BulkUpdate)
- [BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error) ](#BulkUpdateAll)
- [BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) ](#BulkUpsert)
- [PipeAll(pipeline, result interface{}, allowDiskUse bool) error ](#PipeAll)
- [PipeOne(pipeline, result interface{}, allowDiskUse bool) error ](#PipeOne)
- [PipeIter(pipeline interface{}, allowDiskUse bool) *mgo.Iter ](#PipeIter)
- [Explain(pipeline, result interface{}) error ](#Explain)
- [GridFSCreate(prefix, name string) (*mgo.GridFile, error) ](#GridFSCreate)
- [GridFSFindOne(prefix string, query, result interface{}) error ](#GridFSFindOne)
- [GridFSFindAll(prefix string, query, result interface{}) error ](#GridFSFindAll)
- [GridFSOpen(prefix, name string) (*mgo.GridFile, error) ](#GridFSOpen)
- [GridFSRemove(prefix, name string) error ](#GridFSRemove)



#### NewCtx
> NewCtx(name, coll string) *MgoDB

```golang
var Coll = mmgo.NewCtx("test", "coll1")
```

#### SetName_SetCollection
> SetName(name string) <br>
> SetCollection(name string)

```golang
var Coll = mmgo.NewCtx("test", "coll")
Coll.SetName("test1")
// Coll.SetCollection("coll")
```

#### IsEmpty
> (m *MgoDB) IsEmpty() bool

```golang
var Coll = mmgo.NewCtx("test", "coll")
if Coll.IsEmpty() {
    fmt.Println("coll is empty")
}else{
    fmt.Println("coll is not empty")
}
```

### Count 
> Count(query interface{}) (int, error)  
```golang
```

### Insert 
> Insert(docs ...interface{}) error  
```golang
```

### FindOne 
> FindOne(query, selector, result interface{}) error  
```golang
```

### FindAll 
> FindAll(query, selector, result interface{}) error  
```golang
```

### FindPage 
> FindPage(page, limit int, query, selector, result interface{}) error  
```golang
```

### FindIter 
> FindIter(query interface{}) *mgo.Iter  
```golang
```

### Update 
> Update(selector, update interface{}) error  
```golang
```

### Upsert 
> Upsert(selector, update interface{}) error  
```golang
```

### UpdateAll 
> UpdateAll(selector, update interface{}) error  
```golang
```

### Remove 
> Remove(selector interface{}) error  
```golang
```

### RemoveAll 
> RemoveAll(selector interface{}) error  
```golang
```

### RemoveRepeat 
> RemoveRepeat(selector interface{})  
```golang
```

### RemoveRepeatByKey 
> RemoveRepeatByKey(key string)  
```golang
```

### BulkInsert 
> BulkInsert(docs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkRemove 
> BulkRemove(selector ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkRemoveAll 
> BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpdate 
> BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpdateAll 
> BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpsert 
> BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### PipeAll 
> PipeAll(pipeline, result interface{}, allowDiskUse bool) error  
```golang
```

### PipeOne 
> PipeOne(pipeline, result interface{}, allowDiskUse bool) error  
```golang
```

### PipeIter 
> PipeIter(pipeline interface{}, allowDiskUse bool) *mgo.Iter  
```golang
```

### Explain 
> Explain(pipeline, result interface{}) error  
```golang
```

### GridFSCreate 
> GridFSCreate(prefix, name string) (*mgo.GridFile, error)  
```golang
```

### GridFSFindOne 
> GridFSFindOne(prefix string, query, result interface{}) error  
```golang
```

### GridFSFindAll 
> GridFSFindAll(prefix string, query, result interface{}) error  
```golang
```

### GridFSOpen 
> GridFSOpen(prefix, name string) (*mgo.GridFile, error)  
```golang
```

### GridFSRemove 
> GridFSRemove(prefix, name string) error  
```golang
```




# LICENSE
GPL-3.0