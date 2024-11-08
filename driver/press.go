package driver

import "fmt"

// KeyEvent sends a key event with the specified keycode
// Parameters:
//   - keyCode: The Android key code to send
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) KeyEvent(keyCode KeyCode) bool {
	if output, err := d.Run("input", "keyevent", fmt.Sprintf("%d", keyCode)); err != nil || output != "" {
		return false
	}
	return true
}

// Home simulates pressing the home button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Home() bool {
	return d.KeyEvent(KEYCODE_HOME)
}

// Back simulates pressing the back button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Back() bool {
	return d.KeyEvent(KEYCODE_BACK)
}

// Enter simulates pressing the enter key
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Enter() bool {
	return d.KeyEvent(KEYCODE_ENTER)
}

// Search simulates pressing the search button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Search() bool {
	return d.KeyEvent(KEYCODE_SEARCH)
}

// Menu simulates pressing the menu button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Menu() bool {
	return d.KeyEvent(KEYCODE_MENU)
}

// VolumeUp simulates pressing the volume up button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) VolumeUp() bool {
	return d.KeyEvent(KEYCODE_VOLUME_UP)
}

// VolumeDown simulates pressing the volume down button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) VolumeDown() bool {
	return d.KeyEvent(KEYCODE_VOLUME_DOWN)
}

// Power simulates pressing the power button
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) Power() bool {
	return d.KeyEvent(KEYCODE_POWER)
}
