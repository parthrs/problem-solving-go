package misc

import (
	"testing"

	"github.com/rs/zerolog"
)

type Input struct {
	Command string
	Inputs  []int
}

func TestBanking(t *testing.T) {
	inputs := []Input{
		{"CREATE_ACCOUNT", []int{1, 100}},  // New acc 100
		{"CREATE_ACCOUNT", []int{2, 101}},  // New acc 101
		{"CREATE_ACCOUNT", []int{3, 102}},  // New acc 102
		{"CREATE_ACCOUNT", []int{4, 102}},  // Invalid: acc already created
		{"DEPOSIT", []int{5, 99, 100}},     // Invalid: acc N/A
		{"DEPOSIT", []int{6, 100, 100}},    // 100 balance: 100
		{"DEPOSIT", []int{7, 101, 100}},    // 101 balance: 100
		{"DEPOSIT", []int{8, 102, 100}},    // 102 balance: 100
		{"PAY", []int{9, 101, 1000}},       // Invalid: acc balance not enough
		{"PAY", []int{10, 101, 50}},        // 101 balance: 50
		{"DEPOSIT", []int{11, 101, 500}},   // 101 balance: 550
		{"DEPOSIT", []int{12, 102, 100}},   // 102 balance: 200
		{"TOP_ACTIVITY", []int{13, 2}},     // [100(100), 102(200)]
		{"TOP_ACTIVITY", []int{14, 3}},     // [100(100), 102(200), 101(650)]
		{"TOP_ACTIVITY", []int{15, 4}},     // [100(100), 102(200), 101(650)]
		{"CREATE_ACCOUNT", []int{16, 104}}, // New acc 104
		{"DEPOSIT", []int{17, 104, 600}},   // 104 balance: 600
		{"PAY", []int{18, 104, 50}},        // 104 balance: 550
		{"TOP_ACTIVITY", []int{19, 4}},     // [100(100), 102(200), 104(650), 101(650)]
		{"CREATE_ACCOUNT", []int{20, 105}}, // New acc 105
		{"DEPOSIT", []int{21, 105, 100}},   // 105 balance: 100
		{"DEPOSIT", []int{22, 105, 100}},   // 105 balance: 200
		{"DEPOSIT", []int{23, 105, 100}},   // 105 balance: 300
		{"DEPOSIT", []int{24, 105, 350}},   // 105 balance: 650
		{"TOP_ACTIVITY", []int{25, 10}},    // [100(100), 102(200), 105(650), 104(650), 101(650)]
	}
	expecteds := []string{
		"true",
		"true",
		"true",
		"false",
		"",
		"100",
		"100",
		"100",
		"",
		"50",
		"550",
		"200",
		"[100(100), 102(200)]",
		"[100(100), 102(200), 101(650)]",
		"[100(100), 102(200), 101(650)]",
		"true",
		"600",
		"550",
		"[100(100), 102(200), 104(650), 101(650)]",
		"true",
		"100",
		"200",
		"300",
		"650",
		"[100(100), 102(200), 105(650), 104(650), 101(650)]",
	}
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	myBank := NewBank()
	result := []string{}
	for _, input := range inputs {
		result = append(result, myBank.OperationParser(input.Command, input.Inputs))
	}

	for i := range inputs {
		if result[i] != expecteds[i] {
			t.Errorf("Incorrect output for %v, Expected: %v but got %v", inputs[i], expecteds[i], result[i])
		}
	}
}

func TestBankingTransfers(t *testing.T) {
	inputs := []Input{
		{"CREATE_ACCOUNT", []int{1, 1}},                 // "true" New acc 1
		{"CREATE_ACCOUNT", []int{2, 2}},                 // "true" New acc 2
		{"DEPOSIT", []int{3, 1, 2000}},                  // "2000" Acc 1 balance: 2000
		{"DEPOSIT", []int{4, 2, 3000}},                  // "3000" Acc 2 balance: 3000
		{"TRANSFER", []int{5, 1, 2, 5000}},              // "false" Invalid: acc 1 lacks balance
		{"TRANSFER", []int{16, 1, 2, 1000}},             // "transfer1" Tx submitted
		{"ACCEPT_TRANSFER", []int{20, 1, 1}},            // "false" Invalid: Tx not meant for 1
		{"ACCEPT_TRANSFER", []int{21, 11, 1}},           // "false" Invalid: No such acc as 11
		{"ACCEPT_TRANSFER", []int{22, 1, 2}},            // "false" Invalid: No such tx request
		{"ACCEPT_TRANSFER", []int{25, 2, 1}},            // "true" Tx accepted Acc 1 balance: 1000, Acc 2 balance: 4000
		{"ACCEPT_TRANSFER", []int{30, 2, 1}},            // "false" Tx already accepted
		{"TRANSFER", []int{40, 1, 2, 1000}},             // "transfer2" Tx submitted
		{"ACCEPT_TRANSFER", []int{45 + 86400000, 2, 2}}, // "false" Tx expired
		{"TRANSFER", []int{50 + 86400000, 1, 1, 1000}},  // "" Source and Target are same
	}

	expecteds := []string{
		"true",
		"true",
		"2000",
		"3000",
		"",
		"transfer1",
		"false",
		"false",
		"false",
		"true",
		"false",
		"transfer2",
		"false",
		"",
	}
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	myBank := NewBank()
	result := []string{}
	for _, input := range inputs {
		result = append(result, myBank.OperationParser(input.Command, input.Inputs))
	}

	for i := range inputs {
		if result[i] != expecteds[i] {
			t.Errorf("Incorrect output for %v, Expected: %v but got %v", inputs[i], expecteds[i], result[i])
		}
	}
}
