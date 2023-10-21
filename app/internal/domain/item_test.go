package domain_test

import (
	"perx/internal/domain"
	"testing"
)

func TestStatusSting(t *testing.T) {
	type args struct {
		item domain.Status
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Unknown",
			args: args{
				item: domain.StatusUnknown,
			},
			want: "Unknown",
		},
		{
			name: "processed",
			args: args{
				item: domain.StatusProcessed,
			},
			want: "В процессе",
		},
		{
			name: "in queue",
			args: args{
				item: domain.StatusInQueue,
			},
			want: "В очереди",
		},
		{
			name: "done",
			args: args{
				item: domain.StatusDone,
			},
			want: "Завершен",
		},
		{
			name: "error",
			args: args{
				item: domain.StatusError,
			},
			want: "Ошибка",
		},
		{
			name: "empty",
			args: args{
				item: 3151,
			},
			want: "",
		},
		{
			name: "empty",
			args: args{
				item: -154,
			},
			want: "",
		},
		{
			name: "Unknown",
			args: args{
				item: 0,
			},
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.args.item.String() {
				t.Fail()
			}

		})
	}
}
