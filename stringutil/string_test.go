package stringutil

import (
	"reflect"
	"testing"
)

func TestStringRomove(t *testing.T) {
	type args struct {
		str   string
		array []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name: "test1",
			args: args{
				str:   "1",
				array: []string{"1", "2", "3"},
			},
			want:  []string{"2", "3"},
			want1: true,
		},
		{
			name: "test2",
			args: args{
				str:   "2",
				array: []string{"1", "2", "3"},
			},
			want:  []string{"1", "3"},
			want1: true,
		},
		{
			name: "test3",
			args: args{
				str:   "3",
				array: []string{"1", "2", "3"},
			},
			want:  []string{"1", "2"},
			want1: true,
		},
		{
			name: "test4",
			args: args{
				str:   "4",
				array: []string{"1", "2", "3"},
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StringRomove(tt.args.str, tt.args.array)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringRomove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringRomove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
