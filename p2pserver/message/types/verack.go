/*
 * Copyright (C) 2019 The TesraSupernet Authors
 * This file is part of The TesraSupernet library.
 *
 * The TesraSupernet is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The TesraSupernet is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The TesraSupernet.  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"io"

	comm "github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/p2pserver/common"
)

type VerACK struct {
	//TODO remove this legecy field when upgrade network layer protocal
	isConsensus bool
}

//Serialize message payload
func (this *VerACK) Serialization(sink *comm.ZeroCopySink) {
	sink.WriteBool(this.isConsensus)
}

func (this *VerACK) CmdType() string {
	return common.VERACK_TYPE
}

//Deserialize message payload
func (this *VerACK) Deserialization(source *comm.ZeroCopySource) error {
	var irregular, eof bool
	this.isConsensus, irregular, eof = source.NextBool()
	if eof {
		return io.ErrUnexpectedEOF
	}
	if irregular {
		return comm.ErrIrregularData
	}

	return nil
}
