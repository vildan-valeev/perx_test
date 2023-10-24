package domain_test

import (
	"reflect"
	"testing"

	"github.com/vildan-valeev/perx_test/internal/domain"
	"github.com/vildan-valeev/perx_test/internal/transport/dto"
)

func TestStatusString(t *testing.T) {
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

func TestToDTO(t *testing.T) {
	type args struct {
		items domain.Items
	}
	tests := []struct {
		name string
		args args
		want dto.ItemsDTO
	}{
		{
			name: "test 1 - empty",
			args: args{
				items: domain.Items{},
			},
			want: dto.ItemsDTO{},
		},
		{
			name: "test 2 - default fields",
			args: args{
				items: domain.Items{
					0: {},
				},
			},
			want: dto.ItemsDTO{
				{
					Status:      domain.StatusUnknown.String(),
					ReceiptTime: -62135596800000,
					StartTime:   -62135596800000,
					EndTime:     -62135596800000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.args.items.ToDTO()) {
				t.Fail()
			}
		})
	}
}
