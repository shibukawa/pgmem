package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_buffercache_mark_dirty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v65 int64
	_ = v65
	var v74 int64
	_ = v74
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+30)) = uint16(v2)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_get_call_result_type(m, l0, v2, v8+int32(-4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L44
	}
L3:
	;
	return int64(0)
L4:
	;
	if v18 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = F_superuser(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L41
	}
L8:
	;
	if v24 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty[0]))
	if v31 < v14 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v35 = m.G0
	v37 = v35 - int32(32)
	m.G0 = v37
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty[1]))
	F_ResourceOwnerEnlarge(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty[2]))
	v49 = v46 + v14*int32(56)
	v51 = v49 - int32(32)
	v52 = int64(4194304)
	v54 = base.AtomicRmwOr64(m, v51, int32(0), v52)
	if v54&v52 != int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v65 = v54
	goto L17
L15:
	;
	goto L16
L16:
	;
	v145 = F_MarkDirtyUnpinnedBufferInternal(m, v14, v49-int32(56), v8+int32(-35))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L38
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = int32(_a_F_pg_buffercache_mark_dirty_0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = int32(_a_F_pg_buffercache_mark_dirty_1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(_a_F_pg_buffercache_mark_dirty_2)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = int32(0)
	v74 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v74
	if v65&int64(4194304) != v74 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	goto L22
L20:
	;
	goto L21
L21:
	;
	v109 = int32(_a_F_pg_buffercache_mark_dirty_3)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty[3]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(8))+8))
	if v112 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_perform_spin_delay(m, v37+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v91 = int64(0)
	v94 = base.AtomicRmwCmpxchg64(m, v51, int32(0), v91, v91)
	if v94&int64(4194304) != v91 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v129 = int64(4194304)
	v131 = base.AtomicRmwOr64(m, v51, int32(0), v129)
	if v131&v129 != int64(0) {
		v65 = v131
		goto L17
	} else {
		goto L37
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_mark_dirty[3])) = v127
	goto L27
L29:
	;
	if int32(999) < v110 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v110 < int32(11) {
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v117 = int32(900)
	if v117 <= v110 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = v117
	goto L35
L34:
	;
	v120 = v110
	goto L35
L35:
	;
	v127 = v120 + int32(100)
	goto L28
L36:
	;
	v127 = v110 - int32(1)
	goto L28
L37:
	;
	goto L18
L38:
	;
	m.G0 = v37 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(v145)
	v152 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v159 = F_heap_form_tuple(m, v154, v8+int32(-32), v8+int32(-34))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v162 = F_HeapTupleHeaderGetDatum(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	m.G0 = v10 - int32(-64)
	return v162
L41:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_mark_dirty_4), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_5), int32(827), int32(_a_F_pg_buffercache_mark_dirty_6))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_pg_buffercache_mark_dirty_6)
	F_errmsg(m, int32(_a_F_pg_buffercache_mark_dirty_7), v8+int32(-48))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_5), int32(691), int32(_a_F_pg_buffercache_mark_dirty_8))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_mark_dirty_9), v10)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_mark_dirty_5), int32(832), int32(_a_F_pg_buffercache_mark_dirty_6))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
