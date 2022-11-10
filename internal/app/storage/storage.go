package storage

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            string    `json:"id"`
	Date          time.Time `json:"date"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Sum           float64   `json:"sum"`
	Message       string    `json:"message"`
}

type PaymentRepository interface {
	GetPayments(int64) ([]Payment, error)
	GetPayment(string) (Payment, error)
}

type PaymentRepo struct {
	idToPaymentMap map[string]Payment
}

func NewPaymentRepo() *PaymentRepo {
	idToPaymentMap := make(map[string]Payment)

	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	now := time.Now().Unix()

	for i := 1; i < 10000; i++ {
		b := make([]rune, 50)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		p := Payment{
			ID:            uuid.New().String(),
			Date:          time.Unix(now, 0),
			FromAccountID: rand.Intn(1000),
			ToAccountID:   rand.Intn(1000),
			Sum:           math.Round(rand.Float64()*100)/100 + float64(rand.Intn(1000000)),
			Message:       string(b),
		}

		idToPaymentMap[p.ID] = p

		now = now - rand.Int63n(172800)
	}

	lastPayment := Payment{
		ID:            uuid.New().String(),
		Date:          time.Unix(now, 0),
		FromAccountID: rand.Intn(50000),
		ToAccountID:   rand.Intn(50000),
		Sum:           math.Round(rand.Float64()*100)/100 + float64(rand.Intn(1000000)),
		Message:       "here is your flag: the_security_of_the_md5_is_severely_compromised",
	}

	idToPaymentMap[lastPayment.ID] = lastPayment

	return &PaymentRepo{
		idToPaymentMap: idToPaymentMap,
	}
}

type timeSlice []Payment

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].Date.After(p[j].Date)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (r *PaymentRepo) GetPayments(offset int64) ([]Payment, error) {
	startTime := time.Unix(time.Now().Unix()-offset, 0)

	var payments timeSlice
	for _, v := range r.idToPaymentMap {
		if v.Date.After(startTime) {
			payments = append(payments, v)
		}
	}

	sort.Sort(payments)

	return payments, nil
}

func (r *PaymentRepo) GetPayment(id string) (Payment, error) {
	var p Payment
	p, exist := r.idToPaymentMap[id]
	if !exist {
		return p, fmt.Errorf("not found payment %s", id)
	}

	return p, nil
}
