package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_disable_timeout(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var __phi78 int32
	_ = __phi78
	var v79 int32
	_ = v79
	var __phi79 int32
	_ = __phi79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v141 int32
	_ = v141
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, _consts[529])) = v2
	v14 = l0 * int32(40)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[531]))))
	if v17 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v22 <= int32(0) {
		v51 = int32(-1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[532]))) = uint8(v116)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v116 < v119 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[1245])))
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)) = uint8(v71)
	v74 = v26 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v74 < v76 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v26 = v2
	goto L8
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v26 < v44 {
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v31 = v26 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[1245])))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v51 = int32(-1)
	goto L5
L10:
	;
	v38 = v26 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v38 < v40 {
		v26 = v38
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v51 = v26
	goto L5
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51
	v58 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v58 - int32(1)
	F_errmsg_internal(m, int32(488652), v8)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(516127), int32(143), int32(28781))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	__phi78 = v74
	__phi79 = v26
	v78 = __phi78
	v79 = __phi79
	goto L20
L18:
	;
	goto L19
L19:
	;
	v103 = int32(4561016)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	*(*int32)(unsafe.Add(mBase, _consts[530])) = v105 - int32(1)
	goto L3
L20:
	;
	v83 = int32(2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(v83)%32))+uint32(_consts[1245])))
	*(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(v83)%32))+uint32(_consts[1245]))) = v91
	v94 = v78 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[530]))
	if v94 < v96 {
		__phi78 = v94
		__phi79 = v78
		v78 = __phi78
		v79 = __phi79
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	goto L21
L23:
	;
	v125 = m.G0
	v126 = int32(16)
	v127 = v125 - v126
	m.G0 = v127
	F___gettimeofday(m, v127)
	mBase = m.M
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	v131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v127)+8)))
	m.G0 = v127 + v126
	goto L26
L24:
	;
	goto L25
L25:
	;
	m.G0 = v8 + int32(16)
	return
L26:
	;
	F_schedule_alarm(m, v131+v130*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
