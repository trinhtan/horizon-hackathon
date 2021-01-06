package contract

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	SolcCmdText   = "[SOLC_CMD]"
	DirectoryText = "[DIRECTORY]"
	ContractText  = "[CONTRACT]"
)

var (
	// BaseABIBINGenCmdEthereum is the base command for contract compilation to ABI and BIN
	BaseABIBINGenCmdEthereum = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./ethereum-contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/ebrelayer/contract/generated/ethereum/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmdEthereum is the base command for contract binding generation
	BaseBindingGenCmdEthereum = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/ebrelayer/contract/generated/ethereum/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/ebrelayer/contract/generated/ethereum/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/ebrelayer/contract/generated/ethereum/bindings/%s/%s.go", ContractText, ContractText)},
		"")
	// BaseABIBINGenCmdHarmony is the base command for contract compilation to ABI and BIN
	BaseABIBINGenCmdHarmony = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./harmony-contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/ebrelayer/contract/generated/harmony/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmdHarmony is the base command for contract binding generation
	BaseBindingGenCmdHarmony = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/ebrelayer/contract/generated/harmony/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/ebrelayer/contract/generated/harmony/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/ebrelayer/contract/generated/harmony/bindings/%s/%s.go", ContractText, ContractText)},
		"")
)

// CompileEthereumContracts compiles contracts to BIN and ABI files
func CompileEthereumContracts(contracts BridgeContracts) error {
	for _, contract := range contracts {
		// Construct generic BIN/ABI generation cmd with contract's directory path and name
		baseDirectory := ""
		if contract.String() == BridgeBank.String() {
			baseDirectory = contract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmdEthereum, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, contract.String(), -1)

		// Segment BIN and ABI generation cmds
		contractBINGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "bin", -1)
		err := execCmd(contractBINGenCmd)
		if err != nil {
			return err
		}

		contractABIGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "abi", -1)
		err = execCmd(contractABIGenCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// CompileHarmonyContracts compiles contracts to BIN and ABI files
func CompileHarmonyContracts(contracts BridgeContracts) error {
	for _, contract := range contracts {
		// Construct generic BIN/ABI generation cmd with contract's directory path and name
		baseDirectory := ""
		if contract.String() == BridgeBank.String() {
			baseDirectory = contract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmdHarmony, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, contract.String(), -1)

		// Segment BIN and ABI generation cmds
		contractBINGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "bin", -1)
		err := execCmd(contractBINGenCmd)
		if err != nil {
			return err
		}

		contractABIGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "abi", -1)
		err = execCmd(contractABIGenCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateEthereumBindings generates bindings for each ethereum contract
func GenerateEthereumBindings(contracts BridgeContracts) error {
	for _, contract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmdEthereum, ContractText, contract.String(), -1)
		err := execCmd(genBindingCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateHarmonyBindings generates bindings for each ethereum contract
func GenerateHarmonyBindings(contracts BridgeContracts) error {
	for _, contract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmdHarmony, ContractText, contract.String(), -1)
		err := execCmd(genBindingCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// execCmd executes a bash cmd
func execCmd(cmd string) error {

	mainCmd := exec.Command("sh", "-c", cmd)

	var out bytes.Buffer
	var stderr bytes.Buffer

	mainCmd.Stdout = &out
	mainCmd.Stderr = &stderr

	err := mainCmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	fmt.Println("Result: Successfully!")

	return nil
}
