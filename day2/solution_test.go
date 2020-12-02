package day2

import "testing"

func TestGetValidPasswordsNumber(t *testing.T) {
	type args struct {
		passwords     []Password
		validatorFunc ValidatorFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should get valid passwords number for old company",
			args: args{
				passwords: []Password{
					{
						pp: PasswordPolicy{
							minLength:    1,
							maxLength:    3,
							requiredChar: "a",
						},
						value: "abcde",
					},
					{
						pp: PasswordPolicy{
							minLength:    1,
							maxLength:    3,
							requiredChar: "b",
						},
						value: "cdefg",
					},
					{
						pp: PasswordPolicy{
							minLength:    2,
							maxLength:    9,
							requiredChar: "c",
						},
						value: "ccccccccc",
					},
				},
				validatorFunc: ValidateOldCompanyPassword,
			},
			want: 2,
		},
		{
			name: "Should get valid passwords number for current company",
			args: args{
				passwords: []Password{
					{
						pp: PasswordPolicy{
							minLength:    1,
							maxLength:    3,
							requiredChar: "a",
						},
						value: "abcde",
					},
					{
						pp: PasswordPolicy{
							minLength:    1,
							maxLength:    3,
							requiredChar: "b",
						},
						value: "cdefg",
					},
					{
						pp: PasswordPolicy{
							minLength:    2,
							maxLength:    9,
							requiredChar: "c",
						},
						value: "ccccccccc",
					},
				},
				validatorFunc: ValidateCurrenctCompanyPassword,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValidPasswordsNumber(tt.args.passwords, tt.args.validatorFunc); got != tt.want {
				t.Errorf("GetValidPasswordsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
