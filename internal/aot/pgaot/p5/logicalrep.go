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
	v41 = base.I32_div_s(v39, int32(32))
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
	F_errfinish(m, int32(_a_F_logicalrep_rel_mark_updatable_1), int32(330), int32(_a_F_logicalrep_rel_mark_updatable_2))
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
func F_logicalrep_worker_onexit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[1]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	m.T0[v11].(func(*base.Module, int32))(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 == int32(2) {
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
	v19 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[3]))
	if v21 == v19 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	v156 = F_LWLockAcquire(m, v152+int32(_a_F_logicalrep_worker_onexit_0), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L40
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	v58 = F_LWLockAcquire(m, v54+int32(_a_F_logicalrep_worker_onexit_0), int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = v19
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v27<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	F_shm_mq_detach(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v44 = v27 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v44 < v45 {
		v27 = v44
		goto L12
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(0)
	goto L16
L18:
	;
	goto L13
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[5]))
	if v61 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	F_LWLockRelease(m, v140+int32(_a_F_logicalrep_worker_onexit_0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L39
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v70 = int32(0)
	v71 = v61
	v72 = v68
	v74 = int32(0)
	goto L22
L22:
	;
	v78 = v72 + v70*int32(112)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+32)))
	if v79 != int32(1) {
		v95 = v71
		v96 = v72
		v98 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v98 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v100 = v70 + int32(1)
	if v100 < v95 {
		v70 = v100
		v71 = v95
		v72 = v96
		v74 = v98
		goto L22
	} else {
		goto L29
	}
L25:
	;
	v83 = v78 + int32(16)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	if v84 != v66 {
		v95 = v71
		v96 = v72
		v98 = v74
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v86 == int32(0) {
		v95 = v71
		v96 = v72
		v98 = v74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v89 = F_lappend(m, v74, v83)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[5]))
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v95 = v92
	v96 = v94
	v98 = v89
	goto L24
L29:
	;
	goto L23
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v104 <= int32(0) {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(0)
	v109 = v104
	goto L32
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v108<<(uint(int32(2))%32))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+16)))
	if v119 == int32(0) {
		v129 = v109
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L20
L34:
	;
	v131 = v108 + int32(1)
	if v131 < v129 {
		v108 = v131
		v109 = v129
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v122 != int32(3) {
		v129 = v109
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_logicalrep_worker_stop_internal(m, v118, int32(15))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v129 = v128
	goto L34
L38:
	;
	goto L33
L39:
	;
	goto L8
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v160 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v159)+20)) = v160
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+16)) = uint8(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v159)+28)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v159)+36)) = v162
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+68)) = uint8(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+64)) = int32(-1)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[4]))
	F_LWLockRelease(m, v175+int32(_a_F_logicalrep_worker_onexit_0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[2]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+60))
	if v182 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_FileSetDeleteAll(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[7])))
	if v186 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v189 = int32(1)
	F_LockReleaseAll(m, v189, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_onexit[6]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v195 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v197 = F_kill(m, v195, int32(10))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	return
L53:
	;
	goto L52
}
