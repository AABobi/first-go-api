package routes

import (
	"example/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	//Wysyłałem response w json format -> gin dostarcza nam tamą mapę
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Get event error"})
		return
	}
	fmt.Println(event)
	//context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	//Token z headeara pobeiramy?
	var event models.Event
	//Przychodzi call z frontendu, jakiś json post i ShouldBindJSON zapisujemy wartośc do event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	//Pobieramy z contextu userId i przerabiamy je na int64
	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to updated event"})
		return
	}

	var updatedEvent models.Event
	//Zapisuje pod adresem updatedEvent dane z jsona, który przyszły na ten endpoit (np; z frontu)
	err = context.ShouldBindJSON(&updatedEvent)
	fmt.Println("updateEvent", updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not request"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not request2"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Removed"})
}
