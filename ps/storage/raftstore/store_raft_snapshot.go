// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package raftstore

import (
	"github.com/tiglabs/raft/proto"
	"github.com/vearch/vearch/util/log"
)

// Snapshot implements the raft interface.
func (s *Store) Snapshot() (proto.Snapshot, error) {
	return s.GetEngine().NewSnapshot()
}

// ApplySnapshot implements the raft interface.
func (s *Store) ApplySnapshot(peers []proto.Peer, iter proto.SnapIterator) error {
	err := s.GetEngine().ApplySnapshot(peers, iter)
	if err == nil {
		log.Debug("store info is [%+v]", s)
		log.Debug("store status is [%+v]", s.Status())
	}
	return err
}
