# mmgo
> opp in mgo package.

mgo package package.

# Usage
```golang
// NewCtx - (dbName ,collectionName )
var Coll1 = mmgo.NewCtx("test", "coll1")
var Coll2 = mmgo.NewCtx("article", "p1")

func main(){
    Coll1.Insert(bson.M{
		"name": "alice",
		"age":  12,
		"info": "996",
    })
    Coll1.Insert(bson.M{
		"name": "bob",
		"age":  20,
		"info": "996",
    })
    
    var collection = []bson.M
    err := Coll2.FindAll(nil, nil, &collection)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("[LOG] res = %v\n", collection)
}
```

# API_USAGE
- [NewCtx(name, coll string) *MgoDB](#NewCtx)
- [SetName(name string)](#SetName_SetCollection)
- [IsEmpty() bool](#IsEmpty)
- [SetName(name string) ](#SetName)
- [SetCollection(name string) ](#SetCollection)
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



### NewCtx
[top](#API_USAGE)

> NewCtx(name, coll string) *MgoDB

```golang
var Coll = mmgo.NewCtx("test", "coll1")
```

### SetName_SetCollection
[top](#API_USAGE)

> SetName(name string) <br>
> SetCollection(name string)

```golang
var Coll = mmgo.NewCtx("test", "coll")
Coll.SetName("test1")
// Coll.SetCollection("coll")
```

### IsEmpty
[top](#API_USAGE)

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
[top](#API_USAGE)

> Count(query interface{}) (int, error)  
```golang
```

### Insert 
[top](#API_USAGE)

> Insert(docs ...interface{}) error  
```golang
```

### FindOne 
[top](#API_USAGE)

> FindOne(query, selector, result interface{}) error  
```golang
```

### FindAll 
[top](#API_USAGE)

> FindAll(query, selector, result interface{}) error  
```golang
```

### FindPage 
[top](#API_USAGE)

> FindPage(page, limit int, query, selector, result interface{}) error  
```golang
```

### FindIter 
[top](#API_USAGE)

> FindIter(query interface{}) *mgo.Iter  
```golang
```

### Update 
[top](#API_USAGE)

> Update(selector, update interface{}) error  
```golang
```

### Upsert 
[top](#API_USAGE)

> Upsert(selector, update interface{}) error  
```golang
```

### UpdateAll 
[top](#API_USAGE)

> UpdateAll(selector, update interface{}) error  
```golang
```

### Remove 
[top](#API_USAGE)

> Remove(selector interface{}) error  
```golang
```

### RemoveAll 
[top](#API_USAGE)

> RemoveAll(selector interface{}) error  
```golang
```

### RemoveRepeat 
[top](#API_USAGE)

> RemoveRepeat(selector interface{})  
```golang
```

### RemoveRepeatByKey 
[top](#API_USAGE)

> RemoveRepeatByKey(key string)  
```golang
```

### BulkInsert 
[top](#API_USAGE)

> BulkInsert(docs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkRemove 
[top](#API_USAGE)

> BulkRemove(selector ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkRemoveAll 
[top](#API_USAGE)

> BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpdate 
[top](#API_USAGE)

> BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpdateAll 
[top](#API_USAGE)

> BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### BulkUpsert 
[top](#API_USAGE)

> BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error)  
```golang
```

### PipeAll 
[top](#API_USAGE)

> PipeAll(pipeline, result interface{}, allowDiskUse bool) error  
```golang
```

### PipeOne 
[top](#API_USAGE)

> PipeOne(pipeline, result interface{}, allowDiskUse bool) error  
```golang
```

### PipeIter 
[top](#API_USAGE)

> PipeIter(pipeline interface{}, allowDiskUse bool) *mgo.Iter  
```golang
```

### Explain 
[top](#API_USAGE)

> Explain(pipeline, result interface{}) error  
```golang
```

### GridFSCreate 
[top](#API_USAGE)

> GridFSCreate(prefix, name string) (*mgo.GridFile, error)  
```golang
```

### GridFSFindOne 
[top](#API_USAGE)

> GridFSFindOne(prefix string, query, result interface{}) error  
```golang
```

### GridFSFindAll 
[top](#API_USAGE)

> GridFSFindAll(prefix string, query, result interface{}) error  
```golang
```

### GridFSOpen 
[top](#API_USAGE)

> GridFSOpen(prefix, name string) (*mgo.GridFile, error)  
```golang
```

### GridFSRemove 
[top](#API_USAGE)

> GridFSRemove(prefix, name string) error  
```golang
```



# LICENSE
GPL-3.0