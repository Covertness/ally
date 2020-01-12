// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressFactoryABI is the input ABI used to generate the binding from.
const AddressFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"contractAddresses\",\"type\":\"address[]\"}],\"name\":\"batchCollect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"createReceivers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receiversMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tracker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"sendFundsFromReceiverTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AddressFactoryBin is the compiled bytecode used for deploying new contracts.
var AddressFactoryBin = "0x6080604052600062000016620000f560201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350620000d4620000c8620000f560201b60201c565b620000fd60201b60201c565b6000600260006101000a81548160ff02191690831515021790555062000322565b600033905090565b620001188160016200015e60201b620014391790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b6200017082826200024260201b60201c565b15620001e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620002cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018062002bc56022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b61289380620003326000396000f3fe60806040523480156200001157600080fd5b5060043610620001005760003560e01c80638456cb591162000099578063dbb00a7b116200006f578063dbb00a7b1462000302578063f2fde38b1462000497578063f9cabac814620004de578063ff5422f0146200050f5762000100565b80638456cb5914620002865780638da5cb5b14620002925780638f32d59b14620002de5762000100565b80635c975abb11620000db5780635c975abb14620002035780636ef8d66d1462000227578063715018a6146200023357806382dc1ec4146200023f5762000100565b80633f4ba83a1462000105578063457f4fa2146200011157806346fbf68e14620001a4575b600080fd5b6200010f62000580565b005b6200018a600480360360808110156200012957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050620006f7565b604051808215151515815260200191505060405180910390f35b620001e960048036036020811015620001bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050620008ac565b604051808215151515815260200191505060405180910390f35b6200020d620008cb565b604051808215151515815260200191505060405180910390f35b62000231620008e2565b005b6200023d620008f8565b005b62000284600480360360208110156200025757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505062000a34565b005b6200029062000aad565b005b6200029c62000c25565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b620002e862000c4e565b604051808215151515815260200191505060405180910390f35b62000495600480360360808110156200031a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156200037857600080fd5b8201836020820111156200038b57600080fd5b80359060200191846020830284011164010000000083111715620003ae57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156200040f57600080fd5b8201836020820111156200042257600080fd5b803590602001918460208302840111640100000000831117156200044557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505062000cae565b005b620004dc60048036036020811015620004af57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505062000ed7565b005b6200050d60048036036020811015620004f657600080fd5b810190808035906020019092919050505062000f62565b005b6200053e600480360360208110156200052757600080fd5b81019080803590602001909291905050506200105f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b620005946200058e62001092565b620008ac565b620005eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806200279d6030913960400191505060405180910390fd5b600260009054906101000a900460ff166200066e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa620006b462001092565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b60006200070362000c4e565b62000776576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166389d4e6cd8585856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b1580156200086557600080fd5b505af11580156200087a573d6000803e3d6000fd5b505050506040513d60208110156200089157600080fd5b81019080805190602001909291905050509050949350505050565b6000620008c48260016200109a90919063ffffffff16565b9050919050565b6000600260009054906101000a900460ff16905090565b620008f6620008f062001092565b6200117a565b565b6200090262000c4e565b62000975576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b62000a4862000a4262001092565b620008ac565b62000a9f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806200279d6030913960400191505060405180910390fd5b62000aaa81620011d6565b50565b62000ac162000abb62001092565b620008ac565b62000b18576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806200279d6030913960400191505060405180910390fd5b600260009054906101000a900460ff161562000b9c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25862000be262001092565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662000c9262001092565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b62000cb862000c4e565b62000d2b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b815181511462000d3a57600080fd5b60008090505b815181101562000ed05781818151811062000d5757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166389d4e6cd8685848151811062000d8857fe5b6020026020010151876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b15801562000e2e57600080fd5b505af115801562000e43573d6000803e3d6000fd5b505050506040513d602081101562000e5a57600080fd5b810190808051906020019092919050505062000ec2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180620028366029913960400191505060405180910390fd5b808060010191505062000d40565b5050505050565b62000ee162000c4e565b62000f54576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b62000f5f8162001232565b50565b62000f6c62000c4e565b62000fdf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60405162000fed9062001517565b604051809103906000f0801580156200100a573d6000803e3d6000fd5b506003600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562001123576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620028146022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b620011908160016200137890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b620011ec8160016200143990919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620012ba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180620027cd6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6200138482826200109a565b620013db576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180620027f36021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6200144582826200109a565b15620014b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61127780620015268339019056fe608060405260006100146100ee60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100ce6100c36100ee60201b60201c565b6100f660201b60201c565b6000600260006101000a81548160ff021916908315150217905550610313565b600033905090565b61010e81600161015460201b610d8a1790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b610164828261023560201b60201c565b156101d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156102bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806112556022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b610f33806103226000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806382dc1ec41161007157806382dc1ec41461014a5780638456cb591461018e57806389d4e6cd146101985780638da5cb5b1461021e5780638f32d59b14610268578063f2fde38b1461028a576100a9565b80633f4ba83a146100ae57806346fbf68e146100b85780635c975abb146101145780636ef8d66d14610136578063715018a614610140575b600080fd5b6100b66102ce565b005b6100fa600480360360208110156100ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061043c565b604051808215151515815260200191505060405180910390f35b61011c610459565b604051808215151515815260200191505060405180910390f35b61013e610470565b005b610148610482565b005b61018c6004803603602081101561016057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105bb565b005b61019661062c565b005b610204600480360360608110156101ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061079b565b604051808215151515815260200191505060405180910390f35b6102266108e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61027061090b565b604051808215151515815260200191505060405180910390f35b6102cc600480360360208110156102a057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610969565b005b6102de6102d96109ef565b61043c565b610333576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180610e666030913960400191505060405180910390fd5b600260009054906101000a900460ff166103b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6103f96109ef565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b60006104528260016109f790919063ffffffff16565b9050919050565b6000600260009054906101000a900460ff16905090565b61048061047b6109ef565b610ad5565b565b61048a61090b565b6104fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6105cb6105c66109ef565b61043c565b610620576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180610e666030913960400191505060405180910390fd5b61062981610b2f565b50565b61063c6106376109ef565b61043c565b610691576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180610e666030913960400191505060405180910390fd5b600260009054906101000a900460ff1615610714576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586107586109ef565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b60006107a561090b565b610817576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561089e57600080fd5b505af11580156108b2573d6000803e3d6000fd5b505050506040513d60208110156108c857600080fd5b810190808051906020019092919050505090509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661094d6109ef565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b61097161090b565b6109e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6109ec81610b89565b50565b600033905090565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610edd6022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b610ae9816001610ccd90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b610b43816001610d8a90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610e966026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610cd782826109f7565b610d2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180610ebc6021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b610d9482826109f7565b15610e07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505056fe506175736572526f6c653a2063616c6c657220646f6573206e6f742068617665207468652050617573657220726f6c654f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373a265627a7a723158206b1f27b3245c7bc6b72e450cb48d521aee3f401525aa9456415419f4354da31164736f6c634300050d0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373506175736572526f6c653a2063616c6c657220646f6573206e6f742068617665207468652050617573657220726f6c654f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65526f6c65733a206163636f756e7420697320746865207a65726f20616464726573736261746368436f6c6c65637427732063616c6c20746f2073656e6446756e6473546f206661696c6564a265627a7a723158205cfa468942941d139d08728360eab37c4183b325aefe1351619df279e454a22664736f6c634300050d0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployAddressFactory deploys a new Ethereum contract, binding an instance of AddressFactory to it.
func DeployAddressFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressFactory{AddressFactoryCaller: AddressFactoryCaller{contract: contract}, AddressFactoryTransactor: AddressFactoryTransactor{contract: contract}, AddressFactoryFilterer: AddressFactoryFilterer{contract: contract}}, nil
}

// AddressFactory is an auto generated Go binding around an Ethereum contract.
type AddressFactory struct {
	AddressFactoryCaller     // Read-only binding to the contract
	AddressFactoryTransactor // Write-only binding to the contract
	AddressFactoryFilterer   // Log filterer for contract events
}

// AddressFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressFactorySession struct {
	Contract     *AddressFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressFactoryCallerSession struct {
	Contract *AddressFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddressFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressFactoryTransactorSession struct {
	Contract     *AddressFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddressFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressFactoryRaw struct {
	Contract *AddressFactory // Generic contract binding to access the raw methods on
}

// AddressFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressFactoryCallerRaw struct {
	Contract *AddressFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AddressFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressFactoryTransactorRaw struct {
	Contract *AddressFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressFactory creates a new instance of AddressFactory, bound to a specific deployed contract.
func NewAddressFactory(address common.Address, backend bind.ContractBackend) (*AddressFactory, error) {
	contract, err := bindAddressFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressFactory{AddressFactoryCaller: AddressFactoryCaller{contract: contract}, AddressFactoryTransactor: AddressFactoryTransactor{contract: contract}, AddressFactoryFilterer: AddressFactoryFilterer{contract: contract}}, nil
}

// NewAddressFactoryCaller creates a new read-only instance of AddressFactory, bound to a specific deployed contract.
func NewAddressFactoryCaller(address common.Address, caller bind.ContractCaller) (*AddressFactoryCaller, error) {
	contract, err := bindAddressFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryCaller{contract: contract}, nil
}

// NewAddressFactoryTransactor creates a new write-only instance of AddressFactory, bound to a specific deployed contract.
func NewAddressFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressFactoryTransactor, error) {
	contract, err := bindAddressFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryTransactor{contract: contract}, nil
}

// NewAddressFactoryFilterer creates a new log filterer instance of AddressFactory, bound to a specific deployed contract.
func NewAddressFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFactoryFilterer, error) {
	contract, err := bindAddressFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryFilterer{contract: contract}, nil
}

// bindAddressFactory binds a generic wrapper to an already deployed contract.
func bindAddressFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressFactory *AddressFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressFactory.Contract.AddressFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressFactory *AddressFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.Contract.AddressFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressFactory *AddressFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressFactory.Contract.AddressFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressFactory *AddressFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressFactory *AddressFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressFactory *AddressFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressFactory.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AddressFactory *AddressFactoryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AddressFactory.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AddressFactory *AddressFactorySession) IsOwner() (bool, error) {
	return _AddressFactory.Contract.IsOwner(&_AddressFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AddressFactory *AddressFactoryCallerSession) IsOwner() (bool, error) {
	return _AddressFactory.Contract.IsOwner(&_AddressFactory.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AddressFactory *AddressFactoryCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AddressFactory.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AddressFactory *AddressFactorySession) IsPauser(account common.Address) (bool, error) {
	return _AddressFactory.Contract.IsPauser(&_AddressFactory.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AddressFactory *AddressFactoryCallerSession) IsPauser(account common.Address) (bool, error) {
	return _AddressFactory.Contract.IsPauser(&_AddressFactory.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressFactory *AddressFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AddressFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressFactory *AddressFactorySession) Owner() (common.Address, error) {
	return _AddressFactory.Contract.Owner(&_AddressFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressFactory *AddressFactoryCallerSession) Owner() (common.Address, error) {
	return _AddressFactory.Contract.Owner(&_AddressFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AddressFactory *AddressFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AddressFactory.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AddressFactory *AddressFactorySession) Paused() (bool, error) {
	return _AddressFactory.Contract.Paused(&_AddressFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AddressFactory *AddressFactoryCallerSession) Paused() (bool, error) {
	return _AddressFactory.Contract.Paused(&_AddressFactory.CallOpts)
}

// ReceiversMap is a free data retrieval call binding the contract method 0xff5422f0.
//
// Solidity: function receiversMap(uint256 ) constant returns(address)
func (_AddressFactory *AddressFactoryCaller) ReceiversMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AddressFactory.contract.Call(opts, out, "receiversMap", arg0)
	return *ret0, err
}

// ReceiversMap is a free data retrieval call binding the contract method 0xff5422f0.
//
// Solidity: function receiversMap(uint256 ) constant returns(address)
func (_AddressFactory *AddressFactorySession) ReceiversMap(arg0 *big.Int) (common.Address, error) {
	return _AddressFactory.Contract.ReceiversMap(&_AddressFactory.CallOpts, arg0)
}

// ReceiversMap is a free data retrieval call binding the contract method 0xff5422f0.
//
// Solidity: function receiversMap(uint256 ) constant returns(address)
func (_AddressFactory *AddressFactoryCallerSession) ReceiversMap(arg0 *big.Int) (common.Address, error) {
	return _AddressFactory.Contract.ReceiversMap(&_AddressFactory.CallOpts, arg0)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AddressFactory *AddressFactoryTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AddressFactory *AddressFactorySession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.AddPauser(&_AddressFactory.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AddressFactory *AddressFactoryTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.AddPauser(&_AddressFactory.TransactOpts, account)
}

// BatchCollect is a paid mutator transaction binding the contract method 0xdbb00a7b.
//
// Solidity: function batchCollect(address tracker, address receiver, uint256[] amounts, address[] contractAddresses) returns()
func (_AddressFactory *AddressFactoryTransactor) BatchCollect(opts *bind.TransactOpts, tracker common.Address, receiver common.Address, amounts []*big.Int, contractAddresses []common.Address) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "batchCollect", tracker, receiver, amounts, contractAddresses)
}

// BatchCollect is a paid mutator transaction binding the contract method 0xdbb00a7b.
//
// Solidity: function batchCollect(address tracker, address receiver, uint256[] amounts, address[] contractAddresses) returns()
func (_AddressFactory *AddressFactorySession) BatchCollect(tracker common.Address, receiver common.Address, amounts []*big.Int, contractAddresses []common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.BatchCollect(&_AddressFactory.TransactOpts, tracker, receiver, amounts, contractAddresses)
}

// BatchCollect is a paid mutator transaction binding the contract method 0xdbb00a7b.
//
// Solidity: function batchCollect(address tracker, address receiver, uint256[] amounts, address[] contractAddresses) returns()
func (_AddressFactory *AddressFactoryTransactorSession) BatchCollect(tracker common.Address, receiver common.Address, amounts []*big.Int, contractAddresses []common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.BatchCollect(&_AddressFactory.TransactOpts, tracker, receiver, amounts, contractAddresses)
}

// CreateReceivers is a paid mutator transaction binding the contract method 0xf9cabac8.
//
// Solidity: function createReceivers(uint256 ID) returns()
func (_AddressFactory *AddressFactoryTransactor) CreateReceivers(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "createReceivers", ID)
}

// CreateReceivers is a paid mutator transaction binding the contract method 0xf9cabac8.
//
// Solidity: function createReceivers(uint256 ID) returns()
func (_AddressFactory *AddressFactorySession) CreateReceivers(ID *big.Int) (*types.Transaction, error) {
	return _AddressFactory.Contract.CreateReceivers(&_AddressFactory.TransactOpts, ID)
}

// CreateReceivers is a paid mutator transaction binding the contract method 0xf9cabac8.
//
// Solidity: function createReceivers(uint256 ID) returns()
func (_AddressFactory *AddressFactoryTransactorSession) CreateReceivers(ID *big.Int) (*types.Transaction, error) {
	return _AddressFactory.Contract.CreateReceivers(&_AddressFactory.TransactOpts, ID)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AddressFactory *AddressFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AddressFactory *AddressFactorySession) Pause() (*types.Transaction, error) {
	return _AddressFactory.Contract.Pause(&_AddressFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AddressFactory *AddressFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _AddressFactory.Contract.Pause(&_AddressFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressFactory *AddressFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressFactory *AddressFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressFactory.Contract.RenounceOwnership(&_AddressFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressFactory *AddressFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressFactory.Contract.RenounceOwnership(&_AddressFactory.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AddressFactory *AddressFactoryTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AddressFactory *AddressFactorySession) RenouncePauser() (*types.Transaction, error) {
	return _AddressFactory.Contract.RenouncePauser(&_AddressFactory.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AddressFactory *AddressFactoryTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _AddressFactory.Contract.RenouncePauser(&_AddressFactory.TransactOpts)
}

// SendFundsFromReceiverTo is a paid mutator transaction binding the contract method 0x457f4fa2.
//
// Solidity: function sendFundsFromReceiverTo(uint256 ID, address tracker, uint256 amount, address receiver) returns(bool)
func (_AddressFactory *AddressFactoryTransactor) SendFundsFromReceiverTo(opts *bind.TransactOpts, ID *big.Int, tracker common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "sendFundsFromReceiverTo", ID, tracker, amount, receiver)
}

// SendFundsFromReceiverTo is a paid mutator transaction binding the contract method 0x457f4fa2.
//
// Solidity: function sendFundsFromReceiverTo(uint256 ID, address tracker, uint256 amount, address receiver) returns(bool)
func (_AddressFactory *AddressFactorySession) SendFundsFromReceiverTo(ID *big.Int, tracker common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.SendFundsFromReceiverTo(&_AddressFactory.TransactOpts, ID, tracker, amount, receiver)
}

// SendFundsFromReceiverTo is a paid mutator transaction binding the contract method 0x457f4fa2.
//
// Solidity: function sendFundsFromReceiverTo(uint256 ID, address tracker, uint256 amount, address receiver) returns(bool)
func (_AddressFactory *AddressFactoryTransactorSession) SendFundsFromReceiverTo(ID *big.Int, tracker common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.SendFundsFromReceiverTo(&_AddressFactory.TransactOpts, ID, tracker, amount, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressFactory *AddressFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressFactory *AddressFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.TransferOwnership(&_AddressFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressFactory *AddressFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressFactory.Contract.TransferOwnership(&_AddressFactory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AddressFactory *AddressFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AddressFactory *AddressFactorySession) Unpause() (*types.Transaction, error) {
	return _AddressFactory.Contract.Unpause(&_AddressFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AddressFactory *AddressFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _AddressFactory.Contract.Unpause(&_AddressFactory.TransactOpts)
}

// AddressFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressFactory contract.
type AddressFactoryOwnershipTransferredIterator struct {
	Event *AddressFactoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AddressFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressFactoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddressFactoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AddressFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the AddressFactory contract.
type AddressFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressFactory *AddressFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryOwnershipTransferredIterator{contract: _AddressFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressFactory *AddressFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressFactoryOwnershipTransferred)
				if err := _AddressFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressFactory *AddressFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*AddressFactoryOwnershipTransferred, error) {
	event := new(AddressFactoryOwnershipTransferred)
	if err := _AddressFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AddressFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AddressFactory contract.
type AddressFactoryPausedIterator struct {
	Event *AddressFactoryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AddressFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressFactoryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddressFactoryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AddressFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressFactoryPaused represents a Paused event raised by the AddressFactory contract.
type AddressFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AddressFactory *AddressFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*AddressFactoryPausedIterator, error) {

	logs, sub, err := _AddressFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AddressFactoryPausedIterator{contract: _AddressFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AddressFactory *AddressFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AddressFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _AddressFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressFactoryPaused)
				if err := _AddressFactory.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AddressFactory *AddressFactoryFilterer) ParsePaused(log types.Log) (*AddressFactoryPaused, error) {
	event := new(AddressFactoryPaused)
	if err := _AddressFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AddressFactoryPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the AddressFactory contract.
type AddressFactoryPauserAddedIterator struct {
	Event *AddressFactoryPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AddressFactoryPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressFactoryPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddressFactoryPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AddressFactoryPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressFactoryPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressFactoryPauserAdded represents a PauserAdded event raised by the AddressFactory contract.
type AddressFactoryPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*AddressFactoryPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AddressFactory.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryPauserAddedIterator{contract: _AddressFactory.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *AddressFactoryPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AddressFactory.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressFactoryPauserAdded)
				if err := _AddressFactory.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) ParsePauserAdded(log types.Log) (*AddressFactoryPauserAdded, error) {
	event := new(AddressFactoryPauserAdded)
	if err := _AddressFactory.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AddressFactoryPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the AddressFactory contract.
type AddressFactoryPauserRemovedIterator struct {
	Event *AddressFactoryPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AddressFactoryPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressFactoryPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddressFactoryPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AddressFactoryPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressFactoryPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressFactoryPauserRemoved represents a PauserRemoved event raised by the AddressFactory contract.
type AddressFactoryPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*AddressFactoryPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AddressFactory.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AddressFactoryPauserRemovedIterator{contract: _AddressFactory.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *AddressFactoryPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AddressFactory.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressFactoryPauserRemoved)
				if err := _AddressFactory.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AddressFactory *AddressFactoryFilterer) ParsePauserRemoved(log types.Log) (*AddressFactoryPauserRemoved, error) {
	event := new(AddressFactoryPauserRemoved)
	if err := _AddressFactory.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AddressFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AddressFactory contract.
type AddressFactoryUnpausedIterator struct {
	Event *AddressFactoryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AddressFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressFactoryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddressFactoryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AddressFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressFactoryUnpaused represents a Unpaused event raised by the AddressFactory contract.
type AddressFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AddressFactory *AddressFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AddressFactoryUnpausedIterator, error) {

	logs, sub, err := _AddressFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AddressFactoryUnpausedIterator{contract: _AddressFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AddressFactory *AddressFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AddressFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _AddressFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressFactoryUnpaused)
				if err := _AddressFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AddressFactory *AddressFactoryFilterer) ParseUnpaused(log types.Log) (*AddressFactoryUnpaused, error) {
	event := new(AddressFactoryUnpaused)
	if err := _AddressFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
