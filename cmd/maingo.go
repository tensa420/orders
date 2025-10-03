package main

import (
	"context"
	"net/http"

	"github.com/go-jose/go-jose/v4/json"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

import in "order/pkg/pb/inventory/inventory"
import pay "order/pkg/pb/payment/payment"

type Order struct {
	order_uuid       string   `json:"order_uuid"`
	user_uuid        string   `json:"user_uuid"`
	part_uuids       []string `json:"part_uuids"`
	total_price      float64  `json:"total_price"`
	transaction_uuid *string  `json:"transaction_uuid"`
	payment_method   *string  `json:"payment_method"`
	status           string   `json:"status"`
}

type PaymentMethod int

const (
	UNKNOWN PaymentMethod = iota
	CARD
	SBP
	CREDIT_CARD
	INVESTOR_MONEY
)

var InventoryAddress = ":50052"
var PaymentAddress = ":50051"

func HandleCreateOrder(w http.ResponseWriter, r *http.Request, ctx context.Context, client in.InventoryServiceClient) (string, float64, error) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	req := &in.ListPartsRequest{
		Filter: &in.PartsFilter{
			Names: order.part_uuids,
		},
	}

	resp, err := client.ListParts(ctx, req)
	if err != nil {
		status.Errorf(codes.InvalidArgument, "Something went wrong")
	}
	var total_price float64
	for _, part := range resp.GetParts() {
		if part == nil {
			return "", 0, status.Errorf(codes.NotFound, "one of your parts is unavailible")
		}
		total_price += part.Price
	}
	order.order_uuid = uuid.New().String()
	order.status = "PENDING_PAYMENT"

}
func main() {

}
