package google

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalenderService(t *testing.T) {
	svc := NewCalendarService()
	err := svc.AddToCalendar()
	require.NoError(t, err)
}
