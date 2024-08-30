package camera

import (
	"testing"
)

func TestNightCameraRetrieveData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{NewNightCameraData(1, "Коза", "Побігла вліво")})

		data, _ := c.retrieveData()
		if len(*data) == 0 {
			t.Errorf("retrieveData() should not return empty data.")
		}

		for _, d := range *data {
			if d.Animal != "Коза" || d.Movement != "Побігла вліво" {
				t.Errorf("retrieveData() = %+v", d)
			}
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{})

		_, err := c.retrieveData()
		if err == nil {
			t.Errorf("retrieveData() should return error but got <nil>")
		}
	})
}

func TestNightCameraProcessData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{NewNightCameraData(1, "Коза", "Побігла вліво")})

		pd, _ := c.ProcessData()
		if pd.AnimalMovement != "Коза, Побігла вліво; " {
			t.Errorf("ProcessData() = %+v", pd)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{})

		_, err := c.retrieveData()
		if err == nil {
			t.Errorf("retrieveData() should return error but got <nil>")
		}
	})
}
