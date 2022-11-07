package cmd

import (
	"testing"
)

func Test_getFraction(t *testing.T) {
	type args struct {
		dec     float64
		divider string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1/2",
			args: args{
				dec:     0.5,
				divider: "/",
			},
			want: "1/2",
		},
		{
			name: "1/3",
			args: args{
				dec:     0.33333,
				divider: "/",
			},
			want: "1/3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFraction(tt.args.dec, tt.args.divider); got != tt.want {
				t.Errorf("getFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFraction(t *testing.T) {
	type args struct {
		odd string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "1/2",
			args: args{
				odd: "1/2",
			},
			want:    1.5,
			wantErr: false,
		},
		{
			name: "9/4",
			args: args{
				odd: "9/4",
			},
			want:    3.25,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFraction(tt.args.odd)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFraction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
