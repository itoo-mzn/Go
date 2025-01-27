package mail

import (
	"net/mail"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *mail.Address
		wantErr bool
	}{
		{
			name:  "正常系　表示名とメアド",
			input: "Alice <alice@example.com>",
			want: &mail.Address{
				Name:    "Alice",
				Address: "alice@example.com",
			},
			wantErr: false,
		},
		{
			name:  "正常系　メアドのみ",
			input: "bob@example.com",
			want: &mail.Address{
				Name:    "",
				Address: "bob@example.com",
			},
			wantErr: false,
		},
		{
			name:  "正常系　メアドを囲う箇所以外に<>を含む",
			input: "<Alice> <alice@example.com>",
			want: &mail.Address{
				Name:    "<Alice>",
				Address: "alice@example.com",
			},
			wantErr: false,
		},
		{
			name:  "準異常系　表示名に使用できない文字が含まれている（:）",
			input: "Alice : <alice@example.com>",
			want: &mail.Address{
				Name:    "Alice :",
				Address: "alice@example.com",
			},
			wantErr: false,
		},
		{
			name:  "準異常系　表示名に使用できない文字が含まれている（;）",
			input: "Alice ; <alice@example.com>",
			want: &mail.Address{
				Name:    "Alice ;",
				Address: "alice@example.com",
			},
			wantErr: false,
		},
		{
			name:  "準異常系　表示名もメールアドレス",
			input: "bob@example.com <bob@example.com>",
			want: &mail.Address{
				Name:    "bob@example.com",
				Address: "bob@example.com",
			},
			wantErr: false,
		},
		{
			name:    "準異常系　複数の送信者（複数は想定していない）",
			input:   "Alice <alice@example.com>, Bob <bob@example.com>",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "準異常系　group定義どおり（RFCには準拠している）",
			input:   "A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系　<>が含まれていない（angle-addr定義に違反）",
			input:   "Invalid Email Format",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系　メールアドレスが含まれてない",
			input:   "Invalid Email Format",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系　入力が空",
			input:   "",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.want.Name {
				t.Errorf("Parse() got name = %v, want %v", got.Name, tt.want.Name)
			}
			if got.Address != tt.want.Address {
				t.Errorf("Parse() got address = %v, want %v", got.Address, tt.want.Address)
			}
		})
	}
}
