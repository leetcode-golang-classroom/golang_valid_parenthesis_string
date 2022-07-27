package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "(*))"
	for idx := 0; idx < b.N; idx++ {
		checkValidString(s)
	}
}
func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s = \"()\"",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "s = \"(*)\"",
			args: args{s: "(*)"},
			want: true,
		},
		{
			name: "s = \"(*))\"",
			args: args{s: "(*))"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
