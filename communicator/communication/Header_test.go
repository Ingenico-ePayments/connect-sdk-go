package communication

import "testing"

func TestGetDispositionFilename(t *testing.T) {
	testData := map[string]string{}

	testData["attachment; filename=testfile"] = "testfile"
	testData["attachment; filename=\"testfile\""] = "testfile"
	testData["attachment; filename=\"testfile"] = "\"testfile"
	testData["attachment; filename=testfile\""] = "testfile\""
	testData["attachment; filename='testfile'"] = "testfile"
	testData["attachment; filename='testfile"] = "'testfile"
	testData["attachment; filename=testfile'"] = "testfile'"

	testData["filename=testfile"] = "testfile"
	testData["filename=\"testfile\""] = "testfile"
	testData["filename=\"testfile"] = "\"testfile"
	testData["filename=testfile\""] = "testfile\""
	testData["filename='testfile'"] = "testfile"
	testData["filename='testfile"] = "'testfile"
	testData["filename=testfile'"] = "testfile'"

	testData["attachment; filename=testfile; x=y"] = "testfile"
	testData["attachment; filename=\"testfile\"; x=y"] = "testfile"
	testData["attachment; filename=\"testfile; x=y"] = "\"testfile"
	testData["attachment; filename=testfile\"; x=y"] = "testfile\""
	testData["attachment; filename='testfile'; x=y"] = "testfile"
	testData["attachment; filename='testfile; x=y"] = "'testfile"
	testData["attachment; filename=testfile'; x=y"] = "testfile'"

	testData["attachment"] = ""

	testData["filename=\""] = "\""
	testData["filename='"] = "'"

	for value, expected := range testData {
		header, _ := NewHeader("Content-Disposition", value)
		headers := append(Headers{}, *header)
		if expected != headers.GetDispositionFilename() {
			t.Fatalf("TestGetDispositionFilename : %s should yield %s", value, expected)
		}
	}
}