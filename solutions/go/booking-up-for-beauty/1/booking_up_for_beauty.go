package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t,err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil{
        return time.Time{}
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    appointmentTime, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
    return appointmentTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    appointmentTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
    hour := appointmentTime.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    appointmentTime, err := time.Parse("1/2/2006 15:04:05", date)
    if err != nil {
        return ""
    }
    return "You have an appointment on " + appointmentTime.Format("Monday, January 2, 2006") + ", at " + appointmentTime.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
