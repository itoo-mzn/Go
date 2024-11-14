package testsample

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	m.Run()
}

func setup() {
	fmt.Println("テスト全体を実行する前")
}

func teardown() {
	fmt.Println("テスト全体が実行された後")
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常系",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "正常系2",
			args: args{
				a: 2,
				b: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		testStart()
		defer testEnd()

		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Add(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})

		// got := Add(tt.args.a, tt.args.b)
		// if got != tt.want {
		// 	t.Errorf("want %d, got %d", tt.want, got)
		// }
	}

}

func Add(a, b int) int {
	return a + b
}

func testStart() {
	fmt.Println("1つのテストが実行される前")
	log.Println("Fooのログ1")
}

func testEnd() {
	fmt.Println("1つのテストが実行された後")
}
