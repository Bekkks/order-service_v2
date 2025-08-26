package main

import (
	"context"
	"encoding/json"
	"log"
	"crudl/internal/domain"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Опционально: установка посева для воспроизводимости
	// gofakeit.Seed(time.Now().UnixNano())

	for range 10 {
		data := CreateRandomOrder()
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal("ошибка при маршалинге данных:", err)
		}

		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "my-order",
			Balancer: &kafka.LeastBytes{},
		})
		defer writer.Close()

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("order" + strconv.Itoa(gofakeit.Number(0, 999))),
				Value: jsonData,
			},
		)
		if err != nil {
			log.Fatal("ошибка при записи сообщения:", err)
		}
	}

	log.Println("Сообщение успешно отправлено")
}

func CreateRandomOrder() domain.Order {
	ordUID := uuid.New()
	itemsCount := gofakeit.Number(1, 3) // 1-3 случайных предмета
	items := make([]domain.Item, itemsCount)

	for i := 0; i < itemsCount; i++ {
		items[i] = domain.Item{
			OrderUID:    ordUID.String(),
			ChrtID:      gofakeit.Int64() % 1_000_001, // Ограничение до 0..1_000_000
			TrackNumber: "TN" + strconv.Itoa(gofakeit.Number(0, 999_999_999)),
			Price:       gofakeit.Number(10, 500),
			Rid:         gofakeit.LetterN(8),
			Name:        "Item " + gofakeit.Word(),
			Sale:        gofakeit.Number(0, 50),
			Size:        gofakeit.RandomString([]string{"S", "M", "L", "XL"}),
			TotalPrice:  gofakeit.Number(10, 500),
			NmID:        gofakeit.Int64() % 1_000_000_001, // Ограничение до 0..1_000_000_000
			Brand:       "Brand " + gofakeit.Company(),
			Status:      gofakeit.Number(0, 4),
		}
	}

	return domain.Order{
		OrderUID:          ordUID.String(),
		TrackNumber:       "TN" + strconv.Itoa(gofakeit.Number(0, 999_999_999)),
		Entry:             "entry" + gofakeit.LetterN(3),
		Locale:            "en-US", // Оставлено как есть
		InternalSignature: gofakeit.LetterN(10),
		CustomerID:        "cust" + strconv.Itoa(gofakeit.Number(0, 999)),
		DeliveryService:   "service" + gofakeit.Word(),
		ShardKey:          "shard" + strconv.Itoa(gofakeit.Number(0, 9)),
		SmID:              gofakeit.Number(0, 9),
		DateCreated:       gofakeit.DateRange(time.Now().Add(-1000*time.Hour), time.Now()),
		OofShard:          "oof" + gofakeit.LetterN(3),
		Delivery: domain.Delivery{
			OrderUID: ordUID.String(),
			Name:     gofakeit.Name(),
			Phone:    gofakeit.Phone(),
			Zip:      gofakeit.Zip(),
			City:     gofakeit.City(),
			Address:  gofakeit.Street(),
			Region:   gofakeit.State(),
			Email:    gofakeit.Email(),
		},
		Payment: domain.Payment{
			OrderUID:     ordUID.String(),
			Transaction:  "txn" + strconv.Itoa(gofakeit.Number(0, 999_999)),
			RequestID:    "req" + strconv.Itoa(gofakeit.Number(0, 999_999)),
			Currency:     "USD", // Оставлено как есть
			Provider:     "provider" + gofakeit.Word(),
			Amount:       gofakeit.Number(100, 1000),
			PaymentDt:    gofakeit.Date().Unix(),
			Bank:         "bank" + gofakeit.Company(),
			DeliveryCost: gofakeit.Number(10, 50),
			GoodsTotal:   gofakeit.Number(50, 900),
			CustomFee:    gofakeit.Number(0, 100),
		},
		Items: items,
	}
}
