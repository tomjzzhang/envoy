/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import "unsafe"

type HttpCAPI interface {
	HttpContinue(r unsafe.Pointer, status uint64)
	HttpSendLocalReply(r unsafe.Pointer, responseCode int, bodyText string, headers map[string]string, grpcStatus int64, details string)

	// Send a specialized reply that indicates that the filter has failed on the go side. Internally this is used for
	// when unhandled panics are detected.
	HttpSendPanicReply(r unsafe.Pointer, details string)
	// experience api, memory unsafe
	HttpGetHeader(r unsafe.Pointer, key *string, value *string)
	HttpCopyHeaders(r unsafe.Pointer, num uint64, bytes uint64) map[string][]string
	HttpSetHeader(r unsafe.Pointer, key *string, value *string, add bool)
	HttpRemoveHeader(r unsafe.Pointer, key *string)

	HttpGetBuffer(r unsafe.Pointer, bufferPtr uint64, value *string, length uint64)
	HttpSetBufferHelper(r unsafe.Pointer, bufferPtr uint64, value string, action BufferAction)

	HttpCopyTrailers(r unsafe.Pointer, num uint64, bytes uint64) map[string][]string
	HttpSetTrailer(r unsafe.Pointer, key *string, value *string, add bool)
	HttpRemoveTrailer(r unsafe.Pointer, key *string)

	HttpGetStringValue(r unsafe.Pointer, id int) (string, bool)
	HttpGetIntegerValue(r unsafe.Pointer, id int) (uint64, bool)

	// TODO: HttpGetDynamicMetadata(r unsafe.Pointer, filterName string) map[string]interface{}
	HttpSetDynamicMetadata(r unsafe.Pointer, filterName string, key string, value interface{})

	HttpLog(level LogType, message string)
	HttpLogLevel() LogType

	HttpFinalize(r unsafe.Pointer, reason int)

	HttpSetStringFilterState(r unsafe.Pointer, key string, value string, stateType StateType, lifeSpan LifeSpan, streamSharing StreamSharing)
	HttpGetStringFilterState(r unsafe.Pointer, key string) string
}

type NetworkCAPI interface {
	// DownstreamWrite writes buffer data into downstream connection.
	DownstreamWrite(f unsafe.Pointer, bufferPtr unsafe.Pointer, bufferLen int, endStream int)
	// DownstreamClose closes the downstream connection
	DownstreamClose(f unsafe.Pointer, closeType int)
	// DownstreamFinalize cleans up the resource of downstream connection, should be called only by runtime.SetFinalizer
	DownstreamFinalize(f unsafe.Pointer, reason int)
	// DownstreamInfo gets the downstream connection info of infoType
	DownstreamInfo(f unsafe.Pointer, infoType int) string

	// UpstreamConnect creates an envoy upstream connection to address
	UpstreamConnect(libraryID string, addr string) unsafe.Pointer
	// UpstreamWrite writes buffer data into upstream connection.
	UpstreamWrite(f unsafe.Pointer, bufferPtr unsafe.Pointer, bufferLen int, endStream int)
	// UpstreamClose closes the upstream connection
	UpstreamClose(f unsafe.Pointer, closeType int)
	// UpstreamFinalize cleans up the resource of upstream connection, should be called only by runtime.SetFinalizer
	UpstreamFinalize(f unsafe.Pointer, reason int)
	// UpstreamInfo gets the upstream connection info of infoType
	UpstreamInfo(f unsafe.Pointer, infoType int) string
}
