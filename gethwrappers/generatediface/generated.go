package generated

import "github.com/ethereum/go-ethereum/common"

// AbigenLog is an interface for abigen generated log topics
// Kept under package name 'generated' to minimize symbol churn for consumers.
type AbigenLog interface {
	Topic() common.Hash
}
