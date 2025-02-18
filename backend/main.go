package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderType string    `json:"order_type" binding:"required"` // "buy" 또는 "sell"
	UserID    int64     `json:"user_id"`
	StockID   string    `json:"stock_id" binding:"required"`
	Price     int64     `json:"price" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"` // "pending" or "matched"
}

var db *gorm.DB
var redis_mutex *redsync.Mutex

var mutex sync.Mutex

func postOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// lock 환경 구성하여 동시성 문제 해결
	// mutex.Lock()
	// defer mutex.Unlock()

	// distributed lock
	for err := redis_mutex.Lock(); err != nil; err = redis_mutex.Lock() {
	}
	defer redis_mutex.Unlock()

	var count int64
	db.Where("status = ?", "matched").Find(&[]Order{}).Count(&count)
	if count < 50 {
		newOrder.Status = "matched"
	} else {
		newOrder.Status = "pending"
	}
	db.Create(&newOrder)

	c.JSON(http.StatusCreated, newOrder)
}

func initPostgresDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Order{})
	return db
}

func initRedis() *redsync.Redsync {
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     os.Getenv("REDIS_IP"),
		Password: os.Getenv("REDIS_PW"),
	})
	pool := goredis.NewPool(client)

	rs := redsync.New(pool)

	return rs
}

func main() {
	db = initPostgresDB()

	redis_mutex = initRedis().NewMutex("order-lock")

	r := gin.Default()
	r.POST("/order", postOrder)
	r.Run("0.0.0.0:8080")
}
