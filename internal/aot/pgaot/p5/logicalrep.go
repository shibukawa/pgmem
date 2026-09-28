package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_rel_mark_updatable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_RelationGetIndexAttrBitmap(m, v12, int32(2))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v29 = int32(-1)
	goto L10
L2:
	;
	return
L3:
	;
	if v14 != 0 {
		v26 = v14
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_RelationGetIndexAttrBitmap(m, v16, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v18 != 0 {
		v26 = v18
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21 == int32(102) {
		v26 = v20
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v24)
	v26 = v20
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L30
	}
L9:
	;
	m.G0 = v8 + int32(16)
	return
L10:
	;
	if v26 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v106)
	goto L9
L12:
	;
	if v88 < int32(0) {
		goto L9
	} else {
		goto L23
	}
L13:
	;
	v88 = base.I32_ctz(v74) | v75<<(uint(int32(5))%32)
	goto L12
L14:
	;
	v88 = int32(-2)
	goto L12
L15:
	;
	v39 = v29 + int32(1)
	v41 = int32(base.Ui32(v39) >> (uint(int32(5)) % 32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v42 <= v41 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v45 = v26 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41<<(uint(int32(2))%32))))
	v52 = v49 & (int32(-1) << (uint(v39) % 32))
	if v52 != 0 {
		v74 = v52
		v75 = v41
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v54 = v41 + int32(1)
	if v54 == v42 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v57 = v54
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v45+v57<<(uint(int32(2))%32))))
	if v64 != 0 {
		v74 = v64
		v75 = v57
		goto L13
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v66 = v57 + int32(1)
	if v66 != v42 {
		v57 = v66
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if base.Ui32(v88) <= base.Ui32(int32(7)) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94+v88<<(uint(int32(1))%32)-int32(16)))))
	if int32(0) <= v100 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v104 = F_bms_is_member(m, v100, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L11
L28:
	;
	if v104 != 0 {
		v29 = v88
		goto L10
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v119
	F_errmsg(m, int32(_a_F_logicalrep_rel_mark_updatable_0), v8)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_mark_updatable_1), int32(342), int32(_a_F_logicalrep_rel_mark_updatable_2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_worker_onexit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[1]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	m.T0[v12].(func(*base.Module, int32))(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 == int32(3) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v20 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[3]))
	if v22 == v20 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	v166 = F_LWLockAcquire(m, v162+int32(_a_F_logicalrep_worker_onexit_0), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L43
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	v61 = F_LWLockAcquire(m, v57+int32(_a_F_logicalrep_worker_onexit_0), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L19
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v28 = v20
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v28<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	F_shm_mq_detach(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v46 = v28 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v46 < v47 {
		v28 = v46
		goto L12
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(0)
	goto L16
L18:
	;
	goto L13
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[5]))
	if v64 <= int32(0) {
		v143 = v3
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	F_LWLockRelease(m, v147+int32(_a_F_logicalrep_worker_onexit_0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L41
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v73 = int32(0)
	v75 = v64
	v76 = v71
	v77 = v3
	goto L22
L22:
	;
	v82 = v76 + v73<<(uint(int32(7))%32)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+32)))
	if v83 != int32(1) {
		v99 = v75
		v100 = v76
		v101 = v77
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v101 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v104 = v73 + int32(1)
	if v104 < v99 {
		v73 = v104
		v75 = v99
		v76 = v100
		v77 = v101
		goto L22
	} else {
		goto L29
	}
L25:
	;
	v87 = v82 + int32(16)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	if v88 != v69 {
		v99 = v75
		v100 = v76
		v101 = v77
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v90 == int32(0) {
		v99 = v75
		v100 = v76
		v101 = v77
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v93 = F_lappend(m, v77, v87)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[5]))
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v99 = v96
	v100 = v98
	v101 = v93
	goto L24
L29:
	;
	goto L23
L30:
	;
	v143 = int32(0)
	goto L20
L31:
	;
	goto L32
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v109 <= int32(0) {
		v143 = v101
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v113 = int32(0)
	v115 = v109
	goto L34
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v113<<(uint(int32(2))%32))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+16)))
	if v125 == int32(0) {
		v135 = v115
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v143 = v101
	goto L20
L36:
	;
	v137 = v113 + int32(1)
	if v137 < v135 {
		v113 = v137
		v115 = v135
		goto L34
	} else {
		goto L40
	}
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v128 != int32(4) {
		v135 = v115
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_logicalrep_worker_stop_internal(m, v124, int32(15))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v135 = v134
	goto L36
L40:
	;
	goto L35
L41:
	;
	F_list_free(m, v143)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L8
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v170 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v169)+20)) = v170
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+16)) = uint8(v172)
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v169)+28)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v169)+36)) = v172
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+68)) = uint8(v172)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+64)) = int32(-1)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	F_LWLockRelease(m, v185+int32(_a_F_logicalrep_worker_onexit_0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+60))
	if v192 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_FileSetDeleteAll(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[7])))
	if v196 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v199 = int32(1)
	F_LockReleaseAll(m, v199, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v205 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v207 = F_pgmem_kill(m, v205, int32(10))
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	return
}
