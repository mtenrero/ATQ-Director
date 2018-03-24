package dockerInterpreter

import "testing"

func TestVolumeBinding(t *testing.T) {
	volume, err := CreateVolumeWithData("TEST", "/tmp")

	if err != nil {
		t.Errorf("Error Creating Binding Mount: %s", err)
	}

	t.Log(volume)
}

func TestVolumeRemove(t *testing.T) {
	volume, err := CreateVolumeWithData("TESTREMOVE", "/tmp")
	if err != nil {
		t.Errorf("Error Creating Binding Mount: %s", err)
	}

	errRemove := RemoveVolume(volume.Name, true)

	if err != nil {
		t.Errorf("Error Removing Volume : %s", errRemove)
	}
}
