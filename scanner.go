package scanner

import (
	"bytes"

	"github.com/karalabe/hid"
)

const (
	END_OF_CHUND = 40

	// HID_REPORT_BYTE_SIGNIFICANCE
	MODIFIER   = 0
	RESERVED   = 1
	KEY_CODE_1 = 2
	KEY_CODE_2 = 3
	KEY_CODE_3 = 4
	KEY_CODE_4 = 5
	KEY_CODE_5 = 6
	KEY_CODE_6 = 7

	// MODIFIER_BITS
	LEFT_CTRL   byte = 0x1
	LEFT_SHIFT  byte = 0x2
	LEFT_ALT    byte = 0x3
	LEFT_GUI    byte = 0x4
	RIGHT_CTRL  byte = 0x5
	RIGHT_SHIFT byte = 0x6
	RIGHT_ALT   byte = 0x7
	RIGHT_GUI   byte = 0x8
)

func Start(vendorId uint16, productId uint16, onScan func(string), onError func(error)) {
	devices := hid.Enumerate(vendorId, productId)
	devInfo := devices[0]
	dev, err := devInfo.Open()
	if err != nil {
		panic(err)
	}
	defer dev.Close()

	var buffer bytes.Buffer
	for {
		chunk := make([]byte, 64)

		_, err := dev.Read(chunk)
		if err != nil {
			onError(err)
			continue
		}

		keyCode1 := chunk[KEY_CODE_1]

		if keyCode1 != END_OF_CHUND {
			keyCode, ok := keyCodes[keyCode1]
			if ok {
				modifierByte := chunk[MODIFIER]
				isShiftModified := modifierByte&LEFT_SHIFT > 0 || modifierByte&RIGHT_SHIFT > 0

				if isShiftModified && keyCode.shift != "" {
					buffer.WriteString(keyCode.shift)
				} else {
					buffer.WriteString(keyCode.unmodified)
				}

			}
		} else {
			onScan(buffer.String())
			buffer.Reset()
		}
	}
}
