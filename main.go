package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Appointment struct {
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
}

var (
	appointmentFile = "appointment.json" // File to store appointment
	mutex           sync.RWMutex
)

func loadAppointment() (Appointment, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	var appointment Appointment
	file, err := os.Open(appointmentFile)
	if err != nil {
		return appointment, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&appointment); err != nil {
		return appointment, err
	}

	return appointment, nil
}

func saveAppointment(appointment Appointment) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.Create(appointmentFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(appointment); err != nil {
		return err
	}

	return nil
}

func SetAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var newAppointment Appointment

	if err := json.NewDecoder(r.Body).Decode(&newAppointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	if err := saveAppointment(newAppointment); err != nil {
		http.Error(w, "Error saving appointment", http.StatusInternalServerError)
		log.Printf("Error saving appointment: %v", err)
		return
	}

	log.Printf("New appointment set: %v", newAppointment)

	w.WriteHeader(http.StatusCreated)
}

func CheckAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	appointment, err := loadAppointment()
	if err != nil {
		http.Error(w, "Error loading appointment", http.StatusInternalServerError)
		log.Printf("Error loading appointment: %v", err)
		return
	}

	currentTime := time.Now()
	endTime := appointment.Time.Add(appointment.Duration)
	log.Printf("Appointment Time: %v, Current Time: %v, End Time: %v", appointment.Time, currentTime, endTime)

	if currentTime.After(endTime) {
		log.Println("Appointment completed")
		fmt.Fprintln(w, "Appointment completed")
		return
	}

	log.Println("Appointment not completed")
	fmt.Fprintln(w, "Appointment not completed")
}

func main() {
	http.HandleFunc("/api/set-appointment", SetAppointmentHandler)
	http.HandleFunc("/api/check-appointment", CheckAppointmentHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
