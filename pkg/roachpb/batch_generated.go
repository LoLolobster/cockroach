// Code generated by gen_batch.go; DO NOT EDIT

package roachpb

import (
	"bytes"
	"fmt"
	"strconv"
)

type reqCounts [35]int32

// getReqCounts returns the number of times each
// request type appears in the batch.
func (ba *BatchRequest) getReqCounts() reqCounts {
	var counts reqCounts
	for _, r := range ba.Requests {
		switch {
		case r.Get != nil:
			counts[0]++
		case r.Put != nil:
			counts[1]++
		case r.ConditionalPut != nil:
			counts[2]++
		case r.Increment != nil:
			counts[3]++
		case r.Delete != nil:
			counts[4]++
		case r.DeleteRange != nil:
			counts[5]++
		case r.Scan != nil:
			counts[6]++
		case r.BeginTransaction != nil:
			counts[7]++
		case r.EndTransaction != nil:
			counts[8]++
		case r.AdminSplit != nil:
			counts[9]++
		case r.AdminMerge != nil:
			counts[10]++
		case r.AdminTransferLease != nil:
			counts[11]++
		case r.AdminChangeReplicas != nil:
			counts[12]++
		case r.HeartbeatTxn != nil:
			counts[13]++
		case r.Gc != nil:
			counts[14]++
		case r.PushTxn != nil:
			counts[15]++
		case r.RangeLookup != nil:
			counts[16]++
		case r.ResolveIntent != nil:
			counts[17]++
		case r.ResolveIntentRange != nil:
			counts[18]++
		case r.Merge != nil:
			counts[19]++
		case r.TruncateLog != nil:
			counts[20]++
		case r.RequestLease != nil:
			counts[21]++
		case r.ReverseScan != nil:
			counts[22]++
		case r.ComputeChecksum != nil:
			counts[23]++
		case r.DeprecatedVerifyChecksum != nil:
			counts[24]++
		case r.CheckConsistency != nil:
			counts[25]++
		case r.Noop != nil:
			counts[26]++
		case r.InitPut != nil:
			counts[27]++
		case r.ChangeFrozen != nil:
			counts[28]++
		case r.TransferLease != nil:
			counts[29]++
		case r.LeaseInfo != nil:
			counts[30]++
		case r.WriteBatch != nil:
			counts[31]++
		case r.Export != nil:
			counts[32]++
		case r.Import != nil:
			counts[33]++
		case r.QueryTxn != nil:
			counts[34]++
		default:
			panic(fmt.Sprintf("unsupported request: %+v", r))
		}
	}
	return counts
}

var requestNames = []string{
	"Get",
	"Put",
	"CPut",
	"Inc",
	"Del",
	"DelRng",
	"Scan",
	"BeginTxn",
	"EndTxn",
	"AdmSplit",
	"AdmMerge",
	"AdmTransferLease",
	"AdmChangeReplicas",
	"HeartbeatTxn",
	"Gc",
	"PushTxn",
	"RngLookup",
	"ResolveIntent",
	"ResolveIntentRng",
	"Merge",
	"TruncLog",
	"RequestLease",
	"RevScan",
	"ComputeChksum",
	"DeprecatedVerifyChksum",
	"ChkConsistency",
	"Noop",
	"InitPut",
	"ChangeFrozen",
	"TransferLease",
	"LeaseInfo",
	"WriteBatch",
	"Export",
	"Import",
	"QueryTxn",
}

// Summary prints a short summary of the requests in a batch.
func (ba *BatchRequest) Summary() string {
	if len(ba.Requests) == 0 {
		return "empty batch"
	}
	counts := ba.getReqCounts()
	var buf struct {
		bytes.Buffer
		tmp [10]byte
	}
	for i, v := range counts {
		if v != 0 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}
			buf.Write(strconv.AppendInt(buf.tmp[:0], int64(v), 10))
			buf.WriteString(" ")
			buf.WriteString(requestNames[i])
		}
	}
	return buf.String()
}

// CreateReply creates replies for each of the contained requests, wrapped in a
// BatchResponse. The response objects are batch allocated to minimize
// allocation overhead.
func (ba *BatchRequest) CreateReply() *BatchResponse {
	br := &BatchResponse{}
	br.Responses = make([]ResponseUnion, len(ba.Requests))

	counts := ba.getReqCounts()

	var buf0 []GetResponse
	var buf1 []PutResponse
	var buf2 []ConditionalPutResponse
	var buf3 []IncrementResponse
	var buf4 []DeleteResponse
	var buf5 []DeleteRangeResponse
	var buf6 []ScanResponse
	var buf7 []BeginTransactionResponse
	var buf8 []EndTransactionResponse
	var buf9 []AdminSplitResponse
	var buf10 []AdminMergeResponse
	var buf11 []AdminTransferLeaseResponse
	var buf12 []AdminChangeReplicasResponse
	var buf13 []HeartbeatTxnResponse
	var buf14 []GCResponse
	var buf15 []PushTxnResponse
	var buf16 []RangeLookupResponse
	var buf17 []ResolveIntentResponse
	var buf18 []ResolveIntentRangeResponse
	var buf19 []MergeResponse
	var buf20 []TruncateLogResponse
	var buf21 []RequestLeaseResponse
	var buf22 []ReverseScanResponse
	var buf23 []ComputeChecksumResponse
	var buf24 []DeprecatedVerifyChecksumResponse
	var buf25 []CheckConsistencyResponse
	var buf26 []NoopResponse
	var buf27 []InitPutResponse
	var buf28 []ChangeFrozenResponse
	var buf29 []RequestLeaseResponse
	var buf30 []LeaseInfoResponse
	var buf31 []WriteBatchResponse
	var buf32 []ExportResponse
	var buf33 []ImportResponse
	var buf34 []QueryTxnResponse

	for i, r := range ba.Requests {
		switch {
		case r.Get != nil:
			if buf0 == nil {
				buf0 = make([]GetResponse, counts[0])
			}
			br.Responses[i].Get = &buf0[0]
			buf0 = buf0[1:]
		case r.Put != nil:
			if buf1 == nil {
				buf1 = make([]PutResponse, counts[1])
			}
			br.Responses[i].Put = &buf1[0]
			buf1 = buf1[1:]
		case r.ConditionalPut != nil:
			if buf2 == nil {
				buf2 = make([]ConditionalPutResponse, counts[2])
			}
			br.Responses[i].ConditionalPut = &buf2[0]
			buf2 = buf2[1:]
		case r.Increment != nil:
			if buf3 == nil {
				buf3 = make([]IncrementResponse, counts[3])
			}
			br.Responses[i].Increment = &buf3[0]
			buf3 = buf3[1:]
		case r.Delete != nil:
			if buf4 == nil {
				buf4 = make([]DeleteResponse, counts[4])
			}
			br.Responses[i].Delete = &buf4[0]
			buf4 = buf4[1:]
		case r.DeleteRange != nil:
			if buf5 == nil {
				buf5 = make([]DeleteRangeResponse, counts[5])
			}
			br.Responses[i].DeleteRange = &buf5[0]
			buf5 = buf5[1:]
		case r.Scan != nil:
			if buf6 == nil {
				buf6 = make([]ScanResponse, counts[6])
			}
			br.Responses[i].Scan = &buf6[0]
			buf6 = buf6[1:]
		case r.BeginTransaction != nil:
			if buf7 == nil {
				buf7 = make([]BeginTransactionResponse, counts[7])
			}
			br.Responses[i].BeginTransaction = &buf7[0]
			buf7 = buf7[1:]
		case r.EndTransaction != nil:
			if buf8 == nil {
				buf8 = make([]EndTransactionResponse, counts[8])
			}
			br.Responses[i].EndTransaction = &buf8[0]
			buf8 = buf8[1:]
		case r.AdminSplit != nil:
			if buf9 == nil {
				buf9 = make([]AdminSplitResponse, counts[9])
			}
			br.Responses[i].AdminSplit = &buf9[0]
			buf9 = buf9[1:]
		case r.AdminMerge != nil:
			if buf10 == nil {
				buf10 = make([]AdminMergeResponse, counts[10])
			}
			br.Responses[i].AdminMerge = &buf10[0]
			buf10 = buf10[1:]
		case r.AdminTransferLease != nil:
			if buf11 == nil {
				buf11 = make([]AdminTransferLeaseResponse, counts[11])
			}
			br.Responses[i].AdminTransferLease = &buf11[0]
			buf11 = buf11[1:]
		case r.AdminChangeReplicas != nil:
			if buf12 == nil {
				buf12 = make([]AdminChangeReplicasResponse, counts[12])
			}
			br.Responses[i].AdminChangeReplicas = &buf12[0]
			buf12 = buf12[1:]
		case r.HeartbeatTxn != nil:
			if buf13 == nil {
				buf13 = make([]HeartbeatTxnResponse, counts[13])
			}
			br.Responses[i].HeartbeatTxn = &buf13[0]
			buf13 = buf13[1:]
		case r.Gc != nil:
			if buf14 == nil {
				buf14 = make([]GCResponse, counts[14])
			}
			br.Responses[i].Gc = &buf14[0]
			buf14 = buf14[1:]
		case r.PushTxn != nil:
			if buf15 == nil {
				buf15 = make([]PushTxnResponse, counts[15])
			}
			br.Responses[i].PushTxn = &buf15[0]
			buf15 = buf15[1:]
		case r.RangeLookup != nil:
			if buf16 == nil {
				buf16 = make([]RangeLookupResponse, counts[16])
			}
			br.Responses[i].RangeLookup = &buf16[0]
			buf16 = buf16[1:]
		case r.ResolveIntent != nil:
			if buf17 == nil {
				buf17 = make([]ResolveIntentResponse, counts[17])
			}
			br.Responses[i].ResolveIntent = &buf17[0]
			buf17 = buf17[1:]
		case r.ResolveIntentRange != nil:
			if buf18 == nil {
				buf18 = make([]ResolveIntentRangeResponse, counts[18])
			}
			br.Responses[i].ResolveIntentRange = &buf18[0]
			buf18 = buf18[1:]
		case r.Merge != nil:
			if buf19 == nil {
				buf19 = make([]MergeResponse, counts[19])
			}
			br.Responses[i].Merge = &buf19[0]
			buf19 = buf19[1:]
		case r.TruncateLog != nil:
			if buf20 == nil {
				buf20 = make([]TruncateLogResponse, counts[20])
			}
			br.Responses[i].TruncateLog = &buf20[0]
			buf20 = buf20[1:]
		case r.RequestLease != nil:
			if buf21 == nil {
				buf21 = make([]RequestLeaseResponse, counts[21])
			}
			br.Responses[i].RequestLease = &buf21[0]
			buf21 = buf21[1:]
		case r.ReverseScan != nil:
			if buf22 == nil {
				buf22 = make([]ReverseScanResponse, counts[22])
			}
			br.Responses[i].ReverseScan = &buf22[0]
			buf22 = buf22[1:]
		case r.ComputeChecksum != nil:
			if buf23 == nil {
				buf23 = make([]ComputeChecksumResponse, counts[23])
			}
			br.Responses[i].ComputeChecksum = &buf23[0]
			buf23 = buf23[1:]
		case r.DeprecatedVerifyChecksum != nil:
			if buf24 == nil {
				buf24 = make([]DeprecatedVerifyChecksumResponse, counts[24])
			}
			br.Responses[i].DeprecatedVerifyChecksum = &buf24[0]
			buf24 = buf24[1:]
		case r.CheckConsistency != nil:
			if buf25 == nil {
				buf25 = make([]CheckConsistencyResponse, counts[25])
			}
			br.Responses[i].CheckConsistency = &buf25[0]
			buf25 = buf25[1:]
		case r.Noop != nil:
			if buf26 == nil {
				buf26 = make([]NoopResponse, counts[26])
			}
			br.Responses[i].Noop = &buf26[0]
			buf26 = buf26[1:]
		case r.InitPut != nil:
			if buf27 == nil {
				buf27 = make([]InitPutResponse, counts[27])
			}
			br.Responses[i].InitPut = &buf27[0]
			buf27 = buf27[1:]
		case r.ChangeFrozen != nil:
			if buf28 == nil {
				buf28 = make([]ChangeFrozenResponse, counts[28])
			}
			br.Responses[i].ChangeFrozen = &buf28[0]
			buf28 = buf28[1:]
		case r.TransferLease != nil:
			if buf29 == nil {
				buf29 = make([]RequestLeaseResponse, counts[29])
			}
			br.Responses[i].RequestLease = &buf29[0]
			buf29 = buf29[1:]
		case r.LeaseInfo != nil:
			if buf30 == nil {
				buf30 = make([]LeaseInfoResponse, counts[30])
			}
			br.Responses[i].LeaseInfo = &buf30[0]
			buf30 = buf30[1:]
		case r.WriteBatch != nil:
			if buf31 == nil {
				buf31 = make([]WriteBatchResponse, counts[31])
			}
			br.Responses[i].WriteBatch = &buf31[0]
			buf31 = buf31[1:]
		case r.Export != nil:
			if buf32 == nil {
				buf32 = make([]ExportResponse, counts[32])
			}
			br.Responses[i].Export = &buf32[0]
			buf32 = buf32[1:]
		case r.Import != nil:
			if buf33 == nil {
				buf33 = make([]ImportResponse, counts[33])
			}
			br.Responses[i].Import = &buf33[0]
			buf33 = buf33[1:]
		case r.QueryTxn != nil:
			if buf34 == nil {
				buf34 = make([]QueryTxnResponse, counts[34])
			}
			br.Responses[i].QueryTxn = &buf34[0]
			buf34 = buf34[1:]
		default:
			panic(fmt.Sprintf("unsupported request: %+v", r))
		}
	}
	return br
}
