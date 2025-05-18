package main

import (
	"fmt"
	"github.com/STUD-IT-team/kinoshki-RFE/internal/storage"
)

func main() {
	cfg := storage.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "event_user",
		Password: "event_pass",
		DBName:   "event_db",
	}
	var participantStorage storage.ParticipantStorage
	myStorage, err := storage.NewStorage(cfg, participantStorage)
	if err != nil {
		fmt.Printf("Failed to initialize storage: %v", err)
	}

	defer func(myStorage *storage.Storage) {
		err := myStorage.CloseStorage()
		if err != nil {
			fmt.Printf("Failed to close storage: %v", err)
		}
	}(myStorage)

	newPart := storage.Participant{
		FullName:   "Иванов Иван Иванович",
		StudyGroup: "ИУ8-24",
		Phone:      "79161234567",
		Telegram:   "ivanov_ii",
	}

	id, err := myStorage.CreateParticipant(newPart)
	if err != nil {
		fmt.Printf("Create error: %v", err)
	} else {
		fmt.Printf("Created participant with ID: %d", id)
	}

	part, err := myStorage.GetParticipant(id)
	if err != nil {
		fmt.Printf("Get error: %v", err)
	} else {
		fmt.Printf("Participant: %v\n", part)
	}

	part.StudyGroup = "ИУ8-25"
	if err := myStorage.UpdateParticipant(part); err != nil {
		fmt.Printf("Update error: %v", err)
	} else {
		fmt.Println("Participant updated successfully")
	}
	
	if err := myStorage.DeleteParticipant(id); err != nil {
		fmt.Printf("Delete error: %v", err)
	} else {
		fmt.Println("Participant deleted successfully")
	}
}
