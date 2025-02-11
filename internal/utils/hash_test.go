package utils

import "testing"

func TestHashGenerate(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		algorithm string
		want      string
		wantErr   bool
	}{
		{
			name:      "MD5 success",
			text:      "test",
			algorithm: "md5",
			want:      "098f6bcd4621d373cade4e832627b4f6",
			wantErr:   false,
		},
		{
			name:      "SHA256 success",
			text:      "test",
			algorithm: "sha256",
			want:      "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			wantErr:   false,
		},
		{
			name:      "Empty string MD5",
			text:      "",
			algorithm: "md5",
			want:      "d41d8cd98f00b204e9800998ecf8427e",
			wantErr:   false,
		},
		{
			name:      "Empty string SHA256",
			text:      "",
			algorithm: "sha256",
			want:      "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			wantErr:   false,
		},
		{
			name:      "Invalid algorithm",
			text:      "test",
			algorithm: "invalid",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashGenerate(tt.text, tt.algorithm)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashGenerate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashGenerate() = %v, want %v", got, tt.want)
			}
		})
	}
}
