package grpcForRedis

import (
	context "context"
	fmt "fmt"
	"log"
	"time"

	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:10086"
)

var c RedisOpClient

func RPCGetAuth(uid string) (name, passwd string) {
	r, err := c.GetAuth(context.Background(), &Tuid{Uid: uid})
	if err != nil {
		fmt.Println("error in getAuth: ", err)
		return "", ""
	}
	return r.Name, r.Passwd
}

func RPCSetAuthInfo(name, passwd string, expire time.Duration) string {
	uid, err := c.SetAuthInfo(context.Background(), &ItemInfo{Name: name, Passwd: passwd, Duration: int64(expire)})
	if err != nil {
		fmt.Println("error in setAuthInfo: ", err)
		return ""
	}
	return uid.Uid
}

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c = NewRedisOpClient(conn)

}
