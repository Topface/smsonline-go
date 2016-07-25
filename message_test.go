package smsonline

import "testing"

var tableSignTest = []struct {
	user   string
	secret string
	out    string
}{
	{"test", "secret", "6b57dabc49d963b8a9dbd48d60779698"},
	{"test1", "secret2", "e0ca61370304b800d30c573e8a011833"},
	{"123123", "3534t5g", "72431e042fe0fe6f49c98e3b2c028935"},
}

func Test_setCharset(t *testing.T) {
	const Charset = "temp"
	m := message{}
	m.setCharset(Charset)

	if m.charset != Charset {
		t.Fatalf("Incorect charset, want %s, got %s", Charset, m.charset)
	}
}

func Test_setDelay(t *testing.T) {
	const Delay = 100
	m := message{}
	m.setDelay(Delay)

	if m.delay != Delay {
		t.Fatalf("Incorect delay, want %d, got %d", Delay, m.delay)
	}
}

func Test_setDelay2(t *testing.T) {
	const Delay = maxDelay + 1
	m := message{}
	m.setDelay(Delay)

	if m.delay != maxDelay {
		t.Fatalf("Incorect delay, want %d, got %d", maxDelay, m.delay)
	}
}

func Test_setAck(t *testing.T) {
	const Ack = false
	m := message{}
	m.setAck(Ack)

	if m.reportType != typeNoReport {
		t.Fatalf("Incorect report type, want %d, got %d", typeNoReport, m.reportType)
	}
}

func Test_setAck2(t *testing.T) {
	const Ack = true
	m := message{}
	m.setAck(Ack)

	if m.reportType != typeReport {
		t.Fatalf("Incorect report type, want %d, got %d", typeReport, m.reportType)
	}
}

func Test_setBinaryType(t *testing.T) {
	const Binary = false
	m := message{}
	m.setBinaryType(Binary)

	if m.textType != typeText {
		t.Fatalf("Incorect text type, want %d, got %d", typeText, m.textType)
	}
}

func Test_setBinaryType2(t *testing.T) {
	const Binary = true
	m := message{}
	m.setBinaryType(Binary)

	if m.textType != typeBinary {
		t.Fatalf("Incorect text type, want %d, got %d", typeBinary, m.textType)
	}
}

func Test_getSign(t *testing.T) {
	m := makeSms("from", "text", "to")
	for _, testData := range tableSignTest {
		sign := m.getSign(testData.user, testData.secret)
		if sign != testData.out {
			t.Errorf("Incorect sign, want %s, got %s", testData.out, sign)
		}
	}

}

func Test_getMessageData(t *testing.T) {
	const From = "from"
	const Text = "text"
	const To = "to"
	m := makeSms(From, Text, To)
	data := m.getMessageData("user", "secret")

	if data.Get("txt") != Text {
		t.Errorf("Incorect txt data, want %s, got %s", Text, data.Get("txt"))
	}

	if data.Get("from") != From {
		t.Errorf("Incorect from data, want %s, got %s", From, data.Get("from"))
	}

	if data.Get("phones[]") != To {
		t.Errorf("Incorect phones[] data, want %s, got %s", To, data.Get("phones[]"))
	}
}
