package digest_auth_client

import "testing"

func TestWwwAuthenticate(t *testing.T) {
	testCases := []struct {
		name  string
		value string
	}{
		{
			name:  "algorithm without quotes akamai",
			value: `Digest realm="2028110@p-ep2028110.i.akamaientrypoint.net", qop="auth", nonce="c2b968b7d4ca137d8756a11a5bccd5d2", opaque="12f916ce37cd09d63aaf902e66bb2328", algorithm=MD5`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := newWwwAuthenticate(tc.value)
			if a.Algorithm != "MD5" {
				t.Errorf("algo got: %q, want: %q", a.Algorithm, "MD5")
			}
		})
	}
}
