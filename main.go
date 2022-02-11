package main

import (
	"awesomeProject/utilities"
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var TheatreRevenue float64

type seats struct {
	ASeats []string
	BSeats []string
	CSeats []string
}

// initData initialises the Movie Data
func initData() *map[string]seats {
	myshows := map[string]seats{}

	for i := 1; i < 4; i++ {
		var mySeats seats
		for i := 0; i < 10; i++ {
			mySeats.ASeats = append(mySeats.ASeats, "A"+strconv.Itoa(i))
			mySeats.BSeats = append(mySeats.BSeats, "B"+strconv.Itoa(i))
			mySeats.CSeats = append(mySeats.CSeats, "C"+strconv.Itoa(i))
		}
		index := strconv.Itoa(i)
		myshows["Show "+index] = mySeats
	}

	return &myshows
}

// ShowTickets Shows the tickets Available before booking
func ShowTickets(AvailableSeats *map[string]seats) {
	var bookingStatus bool
	fmt.Println(".............Available tickets are ...................")
	for Show, seats := range *AvailableSeats {
		fmt.Println(Show + " running at Audi " + Show[5:])
		fmt.Println(seats.ASeats)
		fmt.Println(seats.BSeats)
		fmt.Println(seats.CSeats)
		fmt.Println()
	}

	fmt.Println("Enter Show Number")
	var showName int
	_, err := fmt.Scanln(&showName)
	if status := utilities.CheckErrors(err); !status {
		return
	}
	ok := utilities.ContainsInt(showName, []int{1, 2, 3})
	if !ok {
		fmt.Println()
		fmt.Println("Invalid Show Name Try Again.....................")
		return
	}

	fmt.Println("Enter Seats separated by comma")
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}
	line = strings.ReplaceAll(line, " ", "")
	bookedSeats := strings.Split(line, ",")
	ok, err = ValidateSeats(bookedSeats)

	if !ok {
		fmt.Println()
		fmt.Println(err)
		return
	} else {
		bookingStatus = BookTickets(bookedSeats, AvailableSeats, showName)
		if !bookingStatus {
			fmt.Println("Please retry from available seats")
			return
		}
		_ = TotalUserCost(bookedSeats)
	}

}

// TotalUserCost Calculates the Movie ticket cost per user
func TotalUserCost(bookedSeats []string) float64 {
	var baseCost float64
	for _, v := range bookedSeats {
		res := strings.ToUpper(v[:1])
		if res == "A" {
			baseCost += 320
		} else if res == "B" {
			baseCost += 280
		} else {
			baseCost += 240
		}
	}
	TheatreRevenue += baseCost
	ServiceTax := math.Floor(baseCost*0.14*100) / 100
	CessTax := math.Floor(baseCost*0.005*100) / 100
	Total := baseCost + (ServiceTax + 2*CessTax)
	fmt.Println("Subtotal:", baseCost)
	fmt.Println("Service Tax @14%:", ServiceTax)
	fmt.Println("Swachh Bharat Cess @0.5%:", CessTax)
	fmt.Println("Krishi Kalyan Cess @0.5%:", CessTax)
	fmt.Println("Total:", Total)
	return Total
}

// BookTickets books tickets and updates seats in the data
// Returns true for successful transaction, and false for unsuccessful
func BookTickets(bookedSeats []string, AvailableSeats *map[string]seats, showNum int) bool {
	var seatAvailable bool
	UndoUpdate := map[string]int{} // To keep tract of booked tickets to revert later in case of failure
	seats := (*AvailableSeats)["Show "+strconv.Itoa(showNum)]
	for _, v := range bookedSeats {
		seatAvailable = false
		res := v[:1]
		if res == "A" {
			for index, seat := range seats.ASeats {
				if seat == v {
					UndoUpdate["A"] = index
					seats.ASeats[index] = " "
					seatAvailable = true
					break
				}
			}
		} else if res == "B" {
			for index, seat := range seats.BSeats {
				if seat == v {
					UndoUpdate["B"] = index
					seats.BSeats[index] = " "
					seatAvailable = true
					break
				}
			}
		} else {
			for index, seat := range seats.CSeats {
				if seat == v {
					UndoUpdate["C"] = index
					seats.CSeats[index] = " "
					seatAvailable = true
					break
				}
			}
		}
		if !seatAvailable {
			// If unavailable seat is encountered, revert the booked seats before that in the transaction
			// ["A1","A2","A3"], if A3 in unavailable revert A1 & A2 bookings
			for sequence, SeatIndex := range UndoUpdate {
				if sequence == "A" {
					(*AvailableSeats)["Show "+strconv.Itoa(showNum)].ASeats[SeatIndex] = "A" + strconv.Itoa(SeatIndex)
				} else if sequence == "B" {
					(*AvailableSeats)["Show "+strconv.Itoa(showNum)].BSeats[SeatIndex] = "B" + strconv.Itoa(SeatIndex)
				} else {
					(*AvailableSeats)["Show "+strconv.Itoa(showNum)].CSeats[SeatIndex] = "C" + strconv.Itoa(SeatIndex)
				}
			}
			return false
		}
	}

	return true
}

// ValidateSeats validates the booking request for
// Checking duplicate seat numbers
// Every seat number provided is unique
func ValidateSeats(bookedSeats []string) (bool, error) {

	var Unique []string
	var err error
	for k, v := range bookedSeats {
		r, _ := regexp.Compile("([A-Ca-c][0-9])")
		flag := r.MatchString(v)
		if !flag {
			err = errors.New("incorrect Seat number(s) try Again")
			return false, err
		}
		res := strings.ToUpper(v[:1]) + v[1:]
		bookedSeats[k] = res
		if ok := utilities.ContainsString(res, Unique); ok {
			err = errors.New("repeated Seat numbers Try Again")
			return false, err
		}
		Unique = append(Unique, res)

	}

	return true, err
}

// TotalRevenue Calculates total revenue for the Theatre
// It asks for password
func TotalRevenue(Password string) float64 {
	if Password != "admin123" {
		return 0
	}
	fmt.Println("..............Total Sales.....................")
	ServiceTax := math.Floor(TheatreRevenue*0.14*100) / 100
	CessTax := math.Floor(TheatreRevenue*0.005*100) / 100
	fmt.Println("Revenue:", TheatreRevenue)
	fmt.Println("Service Tax:", ServiceTax)
	fmt.Println("Swachh Bharat Cess:", CessTax)
	fmt.Println("Krishi Kalyan Cess:", CessTax)
	return TheatreRevenue
}

func main() {
	var choice string
	AvailableSeats := initData()

	for {
		fmt.Println(".........................................................................")
		fmt.Println("Type 1 to book tickets")
		fmt.Println("Type 2 to exit")
		fmt.Println("Type 3 to see total sales(admin)")
		fmt.Println("Enter Your choice: ")
		fmt.Println(".........................................................................")
		_, err := fmt.Scanln(&choice)
		if status := utilities.CheckErrors(err); !status {
			return
		}
		switch choice {

		case "1":
			ShowTickets(AvailableSeats)
		case "2":
			fmt.Println("Thank you. Please visit again")
			return
		case "3":
			var Password string
			fmt.Println("Admin feature Enter password")
			_, err := fmt.Scanln(&Password)
			if status := utilities.CheckErrors(err); !status {
				return
			}
			TotalRevenue(Password)
		default:
			fmt.Println("Invalid entry try again")
		}
	}

}
