package tzgo

import (
	"fmt"
)

// Head -
func (t *TezosNode) Head() (Block, error) {
	uri := "/chains/main/blocks/head"

	var ret Block
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	if t.isOutOfSync(ret.Header) {
		return ret, fmt.Errorf("Out of sync")
	}

	return ret, nil
}

// Header -
func (t *TezosNode) Header() (Header, error) {
	uri := "/chains/main/blocks/head/header"

	var ret Header
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// ActiveDelegatesWithRolls -
func (t *TezosNode) ActiveDelegatesWithRolls() ([]string, error) {
	uri := "/chains/main/blocks/head/context/raw/json/active_delegates_with_rolls"

	var ret []string
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// StakingBalance -
func (t *TezosNode) StakingBalance(address string) (string, error) {
	uri := fmt.Sprintf("/chains/main/blocks/head/context/delegates/%s/staking_balance", address)

	var ret string
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// NetworkPoints -
func (t *TezosNode) NetworkPoints() ([]NetworkPointWithURI, error) {
	uri := "/network/points"

	var ret []NetworkPointWithURI
	if err := t.get(uri, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// PendingOperations -
func (t *TezosNode) PendingOperations() (Mempool, error) {
	uri := "/chains/main/mempool/pending_operations"

	var ret Mempool
	if err := t.get(uri, &ret); err != nil {
		return Mempool{}, err
	}

	return ret, nil
}

// Operations -
func (t *TezosNode) Operations(level int) ([][]Operations, error) {
	uri := fmt.Sprintf("/chains/main/blocks/%d/operations", level)

	var ret [][]Operations
	if err := t.get(uri, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// ChainID -
func (t *TezosNode) ChainID() (string, error) {
	uri := "/chains/main/chain_id"

	var ret string
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// GetCurrentCycle - returns current cycle
func (t *TezosNode) GetCurrentCycle() (int, error) {
	uri := "/chains/main/blocks/head/header"

	var ret Header
	if err := t.get(uri, &ret); err != nil {
		return 0, err
	}

	if t.isOutOfSync(ret) {
		return 0, fmt.Errorf("Out of sync")
	}

	return (ret.Level + 1) / 4096, nil
}

// GetShortBlock -
func (t *TezosNode) GetShortBlock(block int) (ShortBlock, error) {
	uri := fmt.Sprintf("/chains/main/blocks/%d", block)

	var ret ShortBlock
	if err := t.get(uri, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}
