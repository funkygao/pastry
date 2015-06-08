package pastry

import (
	"testing"
)

// Test that node versions are correctly updated
func TestNodeVersionUpdate(t *testing.T) {
	self_id, err := NodeIDFromBytes([]byte("this is a test Node for testing purposes only."))
	if err != nil {
		t.Fatalf(err.Error())
	}
	self := NewNode(self_id, "127.0.0.1", "127.0.0.1", "testing", 0)
	self.updateVersions(2, 3, 4)
	if self.routingTableVersion != 2 {
		t.Errorf("Routing table version was supposed to be %d, was %d instead.", 2, self.routingTableVersion)
	}
	if self.leafsetVersion != 3 {
		t.Errorf("Leafset version was supposed to be %d, was %d instead.", 3, self.leafsetVersion)
	}
	if self.neighborhoodSetVersion != 4 {
		t.Errorf("Neighborhood Set version was supposed to be %d, was %d instead.", 4, self.neighborhoodSetVersion)
	}
}

// Test that node versions are updated even when one version is lower
func TestNodeVersionUpdateMixed(t *testing.T) {
	self_id, err := NodeIDFromBytes([]byte("this is a test Node for testing purposes only."))
	if err != nil {
		t.Fatalf(err.Error())
	}
	self := NewNode(self_id, "127.0.0.1", "127.0.0.1", "testing", 0)
	self.updateVersions(2, 3, 4)
	if self.routingTableVersion != 2 {
		t.Errorf("Routing table version was supposed to be %d, was %d instead.", 2, self.routingTableVersion)
	}
	if self.leafsetVersion != 3 {
		t.Errorf("Leafset version was supposed to be %d, was %d instead.", 3, self.leafsetVersion)
	}
	if self.neighborhoodSetVersion != 4 {
		t.Errorf("Neighborhood Set version was supposed to be %d, was %d instead.", 4, self.neighborhoodSetVersion)
	}
	self.updateVersions(3, 3, 3)
	if self.routingTableVersion != 3 {
		t.Errorf("Routing table version was supposed to be %d, was %d instead.", 3, self.routingTableVersion)
	}
	if self.leafsetVersion != 3 {
		t.Errorf("Leafset version was supposed to be %d, was %d instead.", 3, self.leafsetVersion)
	}
	if self.neighborhoodSetVersion != 4 {
		t.Errorf("Neighborhood Set version was supposed to be %d, was %d instead.", 4, self.neighborhoodSetVersion)
	}
}
