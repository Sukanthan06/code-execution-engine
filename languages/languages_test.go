package languages

import "testing"

func TestGetLanguage(t *testing.T) {
	tests := []struct {
		name          string
		expectedName  string
		expectedExist bool
	}{
		{"python", "Python", true},
		{"javascript", "JavaScript", true},
		{"c", "C", true},
		{"cpp", "C++", true},
		{"invalid_lang", "", false},
	}

	for _, tt := range tests {
		lang, exists := GetLanguage(tt.name)
		if exists != tt.expectedExist {
			t.Errorf("GetLanguage(%q) exists = %v, expected %v", tt.name, exists, tt.expectedExist)
		}
		if exists && lang.Name != tt.expectedName {
			t.Errorf("GetLanguage(%q) Name = %q, expected %q", tt.name, lang.Name, tt.expectedName)
		}
	}
}

func TestIsSupported(t *testing.T) {
	if !IsSupported("python") {
		t.Errorf("expected python to be supported")
	}
	if IsSupported("unsupported") {
		t.Errorf("expected unsupported not to be supported")
	}
}
