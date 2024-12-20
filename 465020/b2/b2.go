package main  
import (  
    "log"  
    "sync"  
    "time"  
)  
type UserDataMap struct {  
    data   sync.Map  
    version int64  
    logger *log.Logger  
    mu     sync.RWMutex  
}  
func NewUserDataMap() *UserDataMap {  
    return &UserDataMap{  
        data:   sync.Map{},  
        version: 0,  
        logger: log.New(log.New(os.Stderr, "UserDataMap: ", log.LstdFlags), "", 0),  
    }  
}  
func (udm *UserDataMap) Get(key string) (interface{}, bool) {  
    udm.mu.RLock()  
    defer udm.mu.RUnlock()  
    value, ok := udm.data.Load(key)  
    udm.logger.Printf("Fetching data for key: %s", key)  
    return value, ok  
}  
func (udm *UserDataMap) Set(key string, value interface{}) {  
    udm.mu.Lock()  
    defer udm.mu.Unlock()  
    udm.data.Store(key, value)  
    atomic.AddInt64(&udm.version, 1)  
    udm.logger.Printf("Setting data for key: %s, value: %v", key, value)  
}  
func (udm *UserDataMap) Delete(key string) {  
    udm.mu.Lock()  
    defer udm.mu.Unlock()  
    udm.data.Delete(key)  
    atomic.AddInt64(&udm.version, 1)  
    udm.logger.Printf("Deleting data for key: %s", key)  
}  
func (udm *UserDataMap) GetVersion() int64 {  
    return atomic.LoadInt64(&udm.version)  
}  
func (udm *UserDataMap) Snapshot() map[string]interface{} {  
    snapshot := make(map[string]interface{})  
    udm.mu.RLock()  
    defer udm.mu.RUnlock()  
    udm.data.Range(func(key, value interface{}) bool {  
        snapshot[key.(string)] = value  
        return true  
    })  
    return snapshot  
}  
func main() {  
    udm := NewUserDataMap()  
    go func() {  
        for {  
            udm.Set("user1", map[string]string{"name": "Alice", "email": "alice@example.com"})  
            time.Sleep(2 * time.Second)  
        }  
    }()  
    go func() {  
        for {  
            udm.Get("user1")  
            time.Sleep(1 * time.Second)  
        }  
    }()  
    time.Sleep(10 * time.Second)  
    snapshot := udm.Snapshot()  
    log.Println("Snapshot of current state:", snapshot)  
}  