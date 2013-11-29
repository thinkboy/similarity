package similarity

import (
	"testing"
)

func TestSimilarity(t *testing.T) {
	str1 := "这剧不错啊,我特别想看，尤其是那个,什么情况啊这是，这是怎么了。我想看啊，行不行了。"
	str2 := "啊这剧不错,wo特别想看，尤其是那个，o test -test.be23让4人的功夫 位她4亲通过企鹅过"

	sim := Similarity(str1, str2)
	t.Logf("similarity : %d%%", sim)
}

func Benchmark_Similarity(b *testing.B) {
	str1 := "这剧不错啊,我特别想看，尤其是那个,什么情况啊这是，这是怎么了。我想看啊，行不行了。"
	str2 := "啊这剧不错,wo特别想看，尤其是那个，o test -test.be23让4人的功夫 位她4亲通过企鹅过"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Similarity(str1, str2)
	}
}
