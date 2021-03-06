package piglatin

import (
        "bytes"
        "github.com/stretchr/testify/assert"
        "testing"
)

var translateTests = []struct {
	in  string
	out string
}{
	{
		in:  "dog",
		out: "ogday",
	},
	{
		in:  "cat",
		out: "atcay",
	},
	{
		in:  "hat",
		out: "athay",
	},
	{
		in:  "egg",
		out: "eggday",
	},
}

func TestTranslate(t *testing.T) {

	for i, test := range translateTests {

		actual := Translate(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)

	}

}

func TestTranslateMultiple(t *testing.T) {
        expected := []string{"atcay", "eggday"}
        actual := TranslateMultiple("cat", "egg")

        assert.Equal(t, expected, actual)
}

func TestWriter(t *testing.T) {
        var b bytes.Buffer
        pigWriter := NewWriter(&b)
        pigWriter.Write([]byte("hat egg"))

        assert.Equal(t, "athay eggday", string(b.Bytes()))
}
