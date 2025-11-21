// https://dev.to/thesaltree/system-design-building-a-parking-lot-system-in-go-3n4h
package parkinglot

import (
	"fmt"
	"sync"
	"time"
)

var (
	parkingLotInstance *parkingLot
	once               sync.Once
)

type parkingLot struct {
	name  string
	floor []*parkingFloor
}

func GetParkingLotInstance() *parkingLot {
	once.Do(func() {
		parkingLotInstance = &parkingLot{}
	})
	return parkingLotInstance
}

type parkingFloor struct {
	floorId             int
	totalParking        int
	twoWheelerParking   int
	threeWheelerParking int
}

func (pl *parkingLot) AddFloor(floorId int) {
	pl.floor = append(pl.floor, NewParkingFloor(floorId))
}

func NewParkingFloor(floorId int) *parkingFloor {
	return &parkingFloor{}
}

type ParkingSpot struct {
	SpotId         int
	VehicleType    vehicles.VehicleType
	CurrentVehicle *vehicles.VehicleInterface
	lock           sync.Mutex
}

func (ps *ParkingSpot) ParkVehicle(vehicle vehicles.VehicleInterface) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	if vehicle.GetVehicleType != ps.VehicleType {
		return fmt.Errorf("vehicle type mismatch: expected %s got %v", ps.VehicleType, vehicle.GetVehicleType)
	}
	if ps.CurrentVehicle != nil {
		return fmt.Errorf("vehicle already present on the parking spot, no space for %v", vehicle.GetVehicleType)
	}
	ps.CurrentVehicle = &vehicle
	return nil
}

type ParkingTicket struct {
	Vehicle      vehicles.VehicleInterface
	EntryTime    time.Time
	ExitTime     time.Time
	Spot         *ParkingSpot
	TotalCharges float64
}

func NewParkingTicket(vehicle vehicle.VehicleInterface, spot *ParkingSpot) *ParkingTicket {
	return &ParkingTicket{Vehicle: vehicle, EntryTime: time.Now(), ExitTime: time.Time{}, Spot: spot, TotalCharges: 0.00}

}
