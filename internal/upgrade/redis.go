package upgrade

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func RedisGet() {
	ctx := context.Background()

	client := client()

	//err := client.Set(ctx, "mykey1", "hello_anybody", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//session := map[string]string{"name": "Marina", "surname": "Pinchuk", "company": "Adidas", "age": "25"}
	//for k, v := range session {
	//	err := client.HSet(ctx, "user2", k, v).Err()
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	//userSession := client.HGetAll(ctx, "user1").Val()
	//fmt.Println(userSession)
	//
	//val, err := client.Get(ctx, "mykey").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("mykey:", val)

	err := client.HDel(ctx, "user2", "name", "surname", "company", "age").Err()
	if err != nil {
		panic(err)
	}
}
