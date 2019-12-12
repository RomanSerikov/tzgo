package tzgo

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON -
func (n *NetworkPointWithURI) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.URI, &n.NetworkPoint}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in NetworkPointWithURI: %d != %d", g, e)
	}
	return nil
}

// UnmarshalJSON -
func (n *FailedOperation) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Hash, &n.PendingOperation}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in PendingOperation: %d != %d", g, e)
	}

	return nil
}
