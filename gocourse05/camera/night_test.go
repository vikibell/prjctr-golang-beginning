package camera

import (
	"testing"
)

func TestNightCameraRetrieveData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{NewNightCameraData(1, "Коза", "Побігла вліво")})

		data, err := c.RetrieveData()
		if err != nil {
			t.Fatalf("RetrieveData() should not have error: %s", err)
		}

		if len(data) == 0 {
			t.Errorf("RetrieveData() should not return empty data.")
		}

		for _, d := range data {
			if d.Animal != "Коза" || d.Movement != "Побігла вліво" {
				t.Errorf("RetrieveData() = %+v", d)
			}
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{})

		_, err := c.RetrieveData()
		if err == nil {
			t.Errorf("RetrieveData() should return error but got <nil>")
		}
	})
}

func TestNightCameraProcessData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{NewNightCameraData(1, "Коза", "Побігла вліво")})

		pd, err := c.ProcessData()
		if err != nil {
			t.Errorf("processData() should not have error: %s", err)
		}

		if pd.AnimalMovement != "Коза, Побігла вліво; " {
			t.Errorf("ProcessData() = %+v", pd)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewNightCamera(1, "Canon", []Data{})

		_, err := c.RetrieveData()
		if err == nil {
			t.Errorf("RetrieveData() should return error but got <nil>")
		}
	})
}
