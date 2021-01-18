package v1

import (
	"context"
	"fmt"
	"github.com/ayubirz/rent-vs-buy-grpc-go/internal/protos/v1/breakeven"
	"github.com/ayubirz/rent-vs-buy-grpc-go/internal/services/rent_guess"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	"log"
)

type Server struct {
}

//CalcBreakeven takes a gRPC request and calculates the point that renting and buying will make equal sense
func (s *Server) CalcBreakeven(ctx context.Context, request *breakeven.BreakevenRequest) (*breakeven.BreakevenResponse, error) {
	log.Printf("Recieved request for CalcBreakeven: %v", request.GetHouseValue())
	
	inputRequest := request
	if request.HouseValue == 0 {
		inputRequest = defaultBreakevenRequest()
	}
	breakevenPoint := rent_guess.Guess(inputRequest)

	return &breakeven.BreakevenResponse{
		Response: fmt.Sprintf("Start renting if rent value is less than(Break-even point): %v", fmt.Sprintf("%.2f", breakevenPoint)),
	}, nil
}

//defaultBreakevenRequest returns a default gRPC request for the times that an empty request comes in
func defaultBreakevenRequest() *breakeven.BreakevenRequest {
	request := breakeven.BreakevenRequest{
		HouseValue:              750000,
		DownPayment:             10,
		PropertyTaxRate:         1,
		PropertyTransferTaxRate: 1,
		LivingYears:             30,
		MortgageInterestRate:    6,
		MortgageTerm:            30,
		MonthlyCommonCost:       250,
	}

	return &request
}
