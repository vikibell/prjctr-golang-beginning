package camera

import (
	"testing"
)

func TestDayCameraRetrieveData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewDayCamera(1, "Canon", []Data{NewDayCameraData(1, "Коза", "Побігла вліво")})

		data, err := c.retrieveData()
		if err != nil {
			t.Errorf("retrieveData() should not have error: %s", err)
		}

		if len(data) == 0 {
			t.Errorf("retrieveData() should not return empty data.")
		}

		for _, d := range data {
			if d.Animal != "Коза" || d.Movement != "Побігла вліво" {
				t.Errorf("retrieveData() = %+v", d)
			}
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewDayCamera(1, "Canon", []Data{})
		_, err := c.retrieveData()
		if err == nil {
			t.Errorf("retrieveData() should return error but got <nil>")
		}
	})
}

func TestDayCameraProcessData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewDayCamera(1, "Canon", []Data{NewDayCameraData(1, "Коза", "Побігла вліво")})

		pd, err := c.ProcessData()
		if err != nil {
			t.Errorf("processData() should not have error: %s", err)
		}

		if pd.AnimalMovement != "Коза, Побігла вліво; " {
			t.Errorf("ProcessData() = %+v", pd)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewDayCamera(1, "Canon", []Data{})

		_, err := c.retrieveData()
		if err == nil {
			t.Errorf("retrieveData() should return error but got <nil>")
		}
	})
}
