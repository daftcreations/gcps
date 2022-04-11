package gcp

import (
	"testing"

	"github.com/daftcreations/gcps/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestContainsProfile(t *testing.T) {
	type test struct {
		input   []types.List
		profile string
		want    bool
	}

	tt := []test{
		{
			input: []types.List{
				{
					Name:     "abc",
					IsActive: true,
				},
				{
					Name:     "default",
					IsActive: false,
				},
			},
			profile: "test",
			want:    false,
		},
		{
			input: []types.List{
				{
					Name:     "abc",
					IsActive: true,
				},
				{
					Name:     "default",
					IsActive: false,
				},
			},
			profile: "default",
			want:    true,
		},
		{
			input: []types.List{
				{
					Name:     "abc",
					IsActive: true,
				},
				{
					Name:     "default",
					IsActive: false,
				},
			},
			profile: "abc",
			want:    false,
		},
	}

	for _, td := range tt {
		have := ContainsProfile(td.input, td.profile)
		require.Equal(t, td.want, have, "for %+v want: %s but have: %s", td.input, td.want, have)
	}
}
