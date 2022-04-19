package main

import (
	"fmt"
	"time"
)

type VaccineScheduler struct {
	AppointmentMap map[time.Time]map[string]map[time.Time]ProviderAppointment
	PatientMap     map[string]Patient
}

type Patient struct {
	PatientID string
	Aggregate ProviderAgg
}

type ProviderAgg struct {
	ProviderID  string
	Appointment ProviderAppointment
}

type Provider struct {
	ProviderID string
}

type ProviderAppointment struct {
	PatientID       string
	AppointmentTime time.Time
}

// ScheduleAppointment reserves a patient appointment with this provider and
// appointment time
//
// Returns:
//     An error if this appointment does not exist or this patient already has
//     an appointment
func (s *VaccineScheduler) ScheduleAppointment(patientID string, providerID string, appointmentTime time.Time) error {
	return nil
}

// CancelAppointment cancels an existing appointment for this patient. If this
// patient has no appointment, do nothing.
func (s *VaccineScheduler) CancelAppointment(patientID string) {
}

// GetPatientAppointment gets this patient's appointment information
//
// Returns:
//     This patient's appointment time and provider ID, or
//     empty values if this patient has no scheduled appointment
func (s *VaccineScheduler) GetPatientAppointment(patientID string) (time.Time, string) {
	return time.Time{}, ""
}

// GetAvailableAppointments gets open appointments on this day
// for patients to browse (day contains a calendar date)
//
// Returns:
//    A mapping of appointment time to list of provider IDs indicating which
//    providers have available appointments for each appointment time on this
//    day
func (s *VaccineScheduler) GetAvailableAppointments(day time.Time) map[time.Time][]string {
	return map[time.Time][]string{}
}

// AddAppointment makes a new appointment with this provider available
//
// Returns:
//     An error if this provider already has an appointment at this time
func (s *VaccineScheduler) AddAppointment(providerID string, appointmentTime time.Time) error {
	//Extract day
	year, month, day := appointmentTime.Date()
	roundDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	firstMap, exists := s.AppointmentMap[roundDate]
	if !exists {
		var map2 = make(map[time.Time]ProviderAppointment)
		map2[appointmentTime] = ProviderAppointment{
			PatientID:       "",
			AppointmentTime: appointmentTime,
		}
		s.AppointmentMap[roundDate] = map[string]map[time.Time]ProviderAppointment{
			providerID: map2,
		}
		return nil
	}
	//Check for the Provider id
	secondMap, exists := firstMap[providerID]
	if !exists {
		var map1 = make(map[time.Time]ProviderAppointment)
		map1[appointmentTime] = ProviderAppointment{
			PatientID:       "",
			AppointmentTime: appointmentTime,
		}
		firstMap[providerID] = secondMap
		return nil
	} else {
		appointment, exists := secondMap[appointmentTime]
		if exists {
			return fmt.Errorf("conflict")
		} else {
			secondMap[appointmentTime] = ProviderAppointment{
				PatientID:       "",
				AppointmentTime: appointmentTime,
			}
			secondMap[appointmentTime] = appointment
		}
	}
	return nil
}

// RemoveAppointment removes an available appointment for a provider at this
// time. If this provider has no appointment at this time, do nothing.
func (s *VaccineScheduler) RemoveAppointment(providerID string, appointmentTime time.Time) {
}

// GetProviderSchedule gets this provider's patient schedule for this day
//
// Returns:
//     A list of ProviderAppointment structs (containing appointment times and
//     patient IDs), sorted by appointment time, which represents the patient
//     schedule for this provider on this day
func (s *VaccineScheduler) GetProviderSchedule(providerID string, day time.Time) []ProviderAppointment {
	return nil
}

func main() {
	scheduler := VaccineScheduler{
		AppointmentMap: map[time.Time]map[string]map[time.Time]ProviderAppointment{},
		PatientMap:     map[string]Patient{},
	}
	err := scheduler.AddAppointment("123", time.Date(2022, 04, 19, 5, 30, 0, 0, time.UTC))
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = scheduler.AddAppointment("123", time.Date(2022, 04, 19, 5, 30, 0, 0, time.UTC))
	fmt.Printf(err.Error())
}
