
package macropoint

import (
	"time"

	"code.uber.internal/freight/uf-supply-locations/mapper"

	"code.uber.internal/freight/uf-supply-locations/entity"
	"code.uber.internal/freight/uf-supply-locations/model"
	"code.uber.internal/freight/uf-supply-locations/utils"
	"code.uber.internal/glue/basefactory"
	"github.com/google/uuid"
)

var (
	testJobUUID                     = uuid.New().String()
	testCarrierUUID                 = uuid.New().String()
	testTruckNumber                 = uuid.New().String()
	testPickupBusinessFacilityUUID  = uuid.New().String()
	testDropoffBusinessFacilityUUID = uuid.New().String()
	testPickupStartTime             = time.Now()
	testPickupEndTime               = time.Now()
	testDropOffStartTime            = time.Now()
	testDropOffEndTime              = time.Now()
	testCreatedAt                   = time.Now()
	testUpdatedAt                   = time.Now()

	testDotNumber = uuid.New().String()
)

func getMacropointOrderData() model.MacropointOrder {
	return model.MacropointOrder{
		JobUUID:                     testJobUUID,
		LoadUUID:                    &testJobUUID,
		TruckNumber:                 &testTruckNumber,
		PickupBusinessFacilityUUID:  &testPickupBusinessFacilityUUID,
		DropoffBusinessFacilityUUID: &testDropoffBusinessFacilityUUID,
		PickupStartTime:             &testPickupStartTime,
		PickupEndTime:               &testPickupEndTime,
		DropoffStartTime:            &testDropOffStartTime,
		DropoffEndTime:              &testDropOffEndTime,
		CreatedAt:                   testCreatedAt,
		UpdatedAt:                   testUpdatedAt,
	}
}

func getPickupAddrData() entity.Address {
	postalcode := "02130"
	address := "Ms Alice Smith Apartment 1c 213 Derrick Street"
	city := "Boston"
	state := "MA"
	country := "USA"
	return entity.Address{
		PostalCode: &postalcode,
		Address:    &address,
		City:       &city,
		State:      &state,
		Country:    &country,
	}
}

func getDropoffAddrData() entity.Address {
	postalcode := "74136"
	address := "4616 Philli Lane"
	city := "Tulsa"
	state := "Oklahoma"
	country := "USA"
	return entity.Address{
		PostalCode: &postalcode,
		Address:    &address,
		City:       &city,
		State:      &state,
		Country:    &country,
	}
}

func getCreateOrderResponse() model.MacropointCreateOrderResponse {
	return model.MacropointCreateOrderResponse{
		OrderID:           basefactory.String(),
		Status:            "Success",
		TrackingRequestID: basefactory.String(),
	}
}

func getChangeOrderResponse() model.MacropointChangeOrderResponse {
	return model.MacropointChangeOrderResponse{
		OrderID:           basefactory.String(),
		Status:            "Success",
		TrackingRequestID: basefactory.String(),
	}
}

func getMacroPointOrderRequestTestData() model.MacropointOrderRequest {
	pickAddr := getPickupAddrData()
	dropoffAddr := getDropoffAddrData()
	orderData := getMacropointOrderData()
	pickupStartTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.PickupStartTime))
	pickupEndTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.PickupEndTime))
	dropoffStartTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.DropoffStartTime))
	dropoffEndTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.DropoffEndTime))

	startTrackTime := orderData.PickupStartTime.Add(model.StartTimeOffset)
	trackingDuration := orderData.DropoffEndTime.Sub(*orderData.PickupStartTime) + model.TrackDurationOffset
	req := model.MacropointOrderRequest{
		TrackStartDateTime: mapper.DateTimeToMacropointTimestamp(startTrackTime),
		Notifications: model.Notifications{
			Notification: model.Notification{
				PartnerMPID:            model.PartnerMPID,
				IDNumber:               testJobUUID,
				TrackDurationInHours:   int(trackingDuration.Hours()),
				TrackIntervalInMinutes: model.TrackIntervalInMinutes,
			},
		},
		Carrier: model.CarrierData{
			CarrierID:   testDotNumber,
			CarrierName: testDotNumber,
		},
		Vehicle: model.VehicleData{
			VehicleType: model.VehicleTypeDryVan,
		},
		TripSheet: model.TripSheetData{
			Stops: model.StopsData{
				Stop: []model.StopData{
					{
						StopName:      model.STOP1,
						StopType:      model.PickupStopType,
						StopID:        model.ONE,
						Address:       mapper.AddressToStopAddress(pickAddr),
						StartDateTime: pickupStartTimeStr,
						EndDateTime:   pickupEndTimeStr,
					},
					{
						StopName:      model.STOP2,
						StopType:      model.DropoffStopType,
						StopID:        model.TWO,
						Address:       mapper.AddressToStopAddress(dropoffAddr),
						StartDateTime: dropoffStartTimeStr,
						EndDateTime:   dropoffEndTimeStr,
					},
				},
			},
		},
		TrackVia: model.TrackViaData{Number: model.Number{
			Type:        "VehicleID",
			AssetNumber: testTruckNumber,
		}},
	}

	return req
}

func getMacroPointChangeOrderRequestTestData() model.MacropointOrderRequest {
	pickAddr := getPickupAddrData()
	dropoffAddr := getDropoffAddrData()
	orderData := getMacropointOrderData()
	pickupStartTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.PickupStartTime))
	pickupEndTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.PickupEndTime))
	dropoffStartTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.DropoffStartTime))
	dropoffEndTimeStr := mapper.DateTimeToMacropointTimestamp(utils.TimeFromPointer(orderData.DropoffEndTime))
	startTrackTime := orderData.PickupStartTime.Add(model.StartTimeOffset)
	trackingDuration := orderData.DropoffEndTime.Sub(*orderData.PickupStartTime) + model.TrackDurationOffset
	req := model.MacropointOrderRequest{
		TrackStartDateTime: mapper.DateTimeToMacropointTimestamp(startTrackTime),
		Notifications: model.Notifications{
			Notification: model.Notification{
				PartnerMPID:            model.PartnerMPID,
				IDNumber:               testJobUUID,
				TrackDurationInHours:   int(trackingDuration.Hours()),
				TrackIntervalInMinutes: model.TrackIntervalInMinutes,
			},
		},
		Carrier: model.CarrierData{
			CarrierID:   testDotNumber,
			CarrierName: testDotNumber,
		},
		Vehicle: model.VehicleData{
			VehicleType: model.VehicleTypeDryVan,
		},
		TripSheet: model.TripSheetData{
			Stops: model.StopsData{
				Stop: []model.StopData{
					{
						StopName:      model.STOP1,
						StopType:      model.PickupStopType,
						StopID:        model.ONE,
						Address:       mapper.AddressToStopAddress(pickAddr),
						StartDateTime: pickupStartTimeStr,
						EndDateTime:   pickupEndTimeStr,
					},
					{
						StopName:      model.STOP2,
						StopType:      model.DropoffStopType,
						StopID:        model.TWO,
						Address:       mapper.AddressToStopAddress(dropoffAddr),
						StartDateTime: dropoffStartTimeStr,
						EndDateTime:   dropoffEndTimeStr,
					},
				},
			},
		},
		TrackVia: model.TrackViaData{Number: model.Number{
			Type:        "VehicleID",
			AssetNumber: testTruckNumber,
		}},
	}

	return req
}
