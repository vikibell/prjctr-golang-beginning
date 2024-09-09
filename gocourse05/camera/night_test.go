package camera

import (
	"testing"
)

func TestNightCameraRetrieveData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		cameraData := NewNightCameraData(1, "Коза", "Побігла вліво")
		c := NewNightCamera(1, "Canon", []Data{cameraData})

		data := c.RetrieveData()
		if len(data) != 1 {
			t.Fatalf("RetrieveData() should return 1 item.")
		}

		got := data[0]
		if got != cameraData {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, cameraData)
		}
	})
}

func TestNightCameraProcessData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{NewNightCameraData(1, "Коза", "Побігла вліво")})

		pd, err := c.ProcessData()
		if err != nil {
			t.Fatalf("processData() should not have error: %s", err)
		}

		got := pd.AnimalMovement
		want := "Коза, Побігла вліво; "
		if got != want {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{})

		_, err := c.ProcessData()
		if err == nil {
			t.Errorf("ProcessData() should return error but got <nil>")
		}
	})
}
