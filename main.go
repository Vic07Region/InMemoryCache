package main

import (
    "fmt"
    "sync"
    "hash/fnv"
   ) 
   

type Cache interface {
    Set(k, v string)
    Get(k string) (v string, ok bool)
}

type memoryInCache struct{
    memory map[string]string
    mutex sync.RWMutex
}

type Sharding struct {
    Shards []*memoryInCache
    
}

func hashkey(key string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(key))
    return h.Sum32()
}

func (s *Sharding) GetShard(key string) *memoryInCache {
    shridx := hashkey(key) % uint32(len(s.Shards))
    return s.Shards[shridx]
}

func (s *Sharding) Set(key, value string){
    shard := s.GetShard(key)
    shard.mutex.Lock()
    shard.memory[key]=value
    shard.mutex.Unlock()
}

func (s *Sharding) Get(key string) (value string, ok bool){
    shard := s.GetShard(key)
    shard.mutex.RLock()
    value, ok = shard.memory[key]
    shard.mutex.RUnlock()
    return value, ok
}

func CreateMemory(numshards int) Cache{
    s := &Sharding{
        Shards: make([]*memoryInCache, numshards),
    }
    
    for i:= range s.Shards{
        s.Shards[i] = &memoryInCache{
            memory: make(map[string]string)}
    }
    return s
} 

func main() {
    micache := CreateMemory(5)
    micache.Set("test","my first cache")
    fmt.Println(micache.Get("test"))
}
