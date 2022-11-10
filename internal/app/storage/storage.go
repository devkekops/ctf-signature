package storage

import (
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            string  `json:"id"`
	Date          string  `json:"date"`
	FromAccountID int     `json:"from_account_id"`
	ToAccountID   int     `json:"to_account_id"`
	Sum           float64 `json:"sum"`
	Message       string  `json:"message"`
}

type PaymentRepository interface {
	GetPayments() ([]Payment, error)
	GetPayment() ([]Payment, error)
}

type PaymentRepo struct {
	arr []Payment
}

func NewPaymentRepo() *PaymentRepo {
	var arr []Payment

	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 1; i < 10000; i++ {
		b := make([]rune, 50)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		p := Payment{
			ID:            uuid.New().String(),
			Date:          time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0).String(),
			FromAccountID: rand.Intn(1000),
			ToAccountID:   rand.Intn(1000),
			Sum:           math.Round(rand.Float64()*100)/100 + float64(rand.Intn(5000000)),
			Message:       string(b),
		}

		arr = append(arr, p)
	}

	lastPayment := Payment{
		ID:            uuid.New().String(),
		Date:          time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0).String(),
		FromAccountID: rand.Intn(1000),
		ToAccountID:   rand.Intn(1000),
		Sum:           math.Round(rand.Float64()*100)/100 + float64(rand.Intn(5000000)),
		Message:       "here is your flag: the_security_of_the_md5_is_severely_compromised",
	}

	arr = append(arr, lastPayment)

	return &PaymentRepo{
		arr: arr,
	}
}

func (r *PaymentRepo) GetPayments() ([]Payment, error) {

	return r.arr, nil
}

func (r *PaymentRepo) GetPayment() ([]Payment, error) {

	return r.arr, nil
}
