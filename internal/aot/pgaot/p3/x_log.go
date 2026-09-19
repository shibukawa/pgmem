package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetXLogInsertRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(0), int32(1))
	if v8 != 0 {
		F_s_lock(m, v5, int32(_a_F_GetXLogInsertRecPtr_0), int32(_a_F_GetXLogInsertRecPtr_1), int32(_a_F_GetXLogInsertRecPtr_2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5))), uint32(v17))
			v21 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[1])))
			v22 = base.I64_div_u_s(v16, v21)
			v24 = v16 - v22*v21
			if base.Ui64(v24) <= base.Ui64(int64(8151)) {
				v42 = v24 + int64(40)
			} else {
				v30 = v24 - int64(8152)
				v31 = int64(8168)
				v32 = base.I64_div_u_s(v30, v31)
				v42 = v30 - v32*v31 + v32<<(uint(int64(13))%64) + int64(8216)
			}
			v44 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[2])))
			return v22*v44 + v42&int64(4294967295)
		}
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v17 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5))), uint32(v17))
		v21 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[1])))
		v22 = base.I64_div_u_s(v16, v21)
		v24 = v16 - v22*v21
		if base.Ui64(v24) <= base.Ui64(int64(8151)) {
			v42 = v24 + int64(40)
		} else {
			v30 = v24 - int64(8152)
			v31 = int64(8168)
			v32 = base.I64_div_u_s(v30, v31)
			v42 = v30 - v32*v31 + v32<<(uint(int64(13))%64) + int64(8216)
		}
		v44 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[2])))
		return v22*v44 + v42&int64(4294967295)
	}
}
func F_GetXLogWriteRecPtr(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	v1 = int64(0)
	v3 = int32(_a_F_GetXLogWriteRecPtr_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[0]))
	v8 = base.AtomicRmwCmpxchg64(m, v4, int32(280), v1, v1)
	*(*int64)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[1])) = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[0]))
	v16 = base.AtomicRmwCmpxchg64(m, v12, int32(272), v1, v1)
	*(*int64)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[2])) = v16
	return v16
}
func F_XLogArchiveNotify(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(_a_F_XLogArchiveNotify_0)
	v13 = v7 + int32(48)
	v18 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_XLogArchiveNotify_1), v7+int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = F_AllocateFile(m, v13, int32(_a_F_XLogArchiveNotify_2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v7 + int32(1072)
	return
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v42 = F_FreeFile(m, v21)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if v27 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	F_errmsg(m, int32(_a_F_XLogArchiveNotify_3), v7)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_XLogArchiveNotify_4), int32(457), int32(_a_F_XLogArchiveNotify_5))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L3
L13:
	;
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v65 = F_strlen(m, l0)
	mBase = m.M
	if v65 != int32(16) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(48)
	F_errmsg(m, int32(_a_F_XLogArchiveNotify_6), v7+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_XLogArchiveNotify_4), int32(465), int32(_a_F_XLogArchiveNotify_5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[0])))
	if v189 != int32(1) {
		goto L3
	} else {
		goto L52
	}
L23:
	;
	v68 = int32(_a_F_XLogArchiveNotify_7)
	v72 = m.G0
	v74 = v72 - int32(32)
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v74)+24)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v74)+16)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[1])))
	if v83 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v151 != int32(8) {
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v151 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[2])))
	if v87 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v91 = l0
	goto L31
L29:
	;
	goto L30
L30:
	;
	v101 = v68
	v102 = v83
	goto L34
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v97 == v83 {
		v91 = v91 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v151 = v91 - l0
	goto L24
L33:
	;
	goto L32
L34:
	;
	v109 = v74 + int32(base.Ui32(v102)>>(uint(int32(3))%32))&int32(28)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110 | v111<<(uint(v102)%32)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v115 != 0 {
		v101 = v101 + v111
		v102 = v115
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v118 == int32(0) {
		v141 = l0
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v151 = v141 - l0
	goto L24
L38:
	;
	v122 = l0
	v123 = v118
	goto L39
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(base.Ui32(v123)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v131)>>(uint(v123)%32))&int32(1) == int32(0) {
		v141 = v122
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v141 = v139
	goto L37
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	v139 = v122 + int32(1)
	if v137 != 0 {
		v122 = v139
		v123 = v137
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v155 = l0 + int32(8)
	v156 = int32(_a_F_XLogArchiveNotify_8)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[3])))
	if base.B2i32(v159 == int32(0))|base.B2i32(v159 != v162) != 0 {
		v180 = v159
		v181 = v162
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v180-v181 != 0 {
		goto L22
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v165 = v155
	v166 = v156
	goto L47
L47:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v170 == int32(0) {
		v180 = v170
		v181 = v169
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v180 = v170
	v181 = v169
	goto L45
L49:
	;
	v173 = int32(1)
	if v170 == v169 {
		v165 = v165 + v173
		v166 = v166 + v173
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[4]))
	v187 = base.AtomicRmwXchg32(m, v184, int32(4), int32(1))
	goto L22
L52:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[4]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v194 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[5]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v204 = v199 + v194*int32(640) + int32(20)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v205 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	goto L3
L56:
	;
	goto L55
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v208 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	if v211 == int32(0) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[6]))
	if v215 == v211 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v217 = m.G0
	v219 = v217 - int32(16)
	m.G0 = v219
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[7]))
	if v222 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v245 = F_pgmem_kill(m, v211, int32(23))
	mBase = m.M
	goto L57
L64:
	;
	m.G0 = v219 + int32(16)
	goto L56
L65:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+15)) = uint8(v225)
	goto L66
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[8]))
	v233 = F_write(m, v229, v219+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v233 {
		goto L64
	} else {
		goto L68
	}
L67:
	;
	goto L64
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[9]))
	if v237 == int32(27) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
}
func F_XLogDropRelation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v10
	F_forget_invalid_pages(m, v6, l1, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_XLogFileInitInternal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int64
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v360 int64
	_ = v360
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int64
	_ = v427
	var v431 int32
	_ = v431
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v452 int64
	_ = v452
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v480 int64
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v525 int64
	_ = v525
	var v526 int64
	_ = v526
	var v530 int64
	_ = v530
	var v580 int32
	_ = v580
	var v581 int64
	_ = v581
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v597 int64
	_ = v597
	var v601 int32
	_ = v601
	var v617 int32
	_ = v617
	var v618 int64
	_ = v618
	var v622 int64
	_ = v622
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v643 int64
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	v10 = m.G0
	v12 = v10 - int32(1168)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = l1
	v17 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0])))
	v18 = base.I64_div_u_s(int64(4294967296), v17)
	v19 = base.I64_div_u_s(l0, v18)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+116)) = uint32(v19)
	v22 = l0 - v18*v19
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+120)) = uint32(v22)
	v28 = F_pg_snprintf(m, l3, int32(1024), int32(_a_F_XLogFileInitInternal_0), v12+int32(112))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v32)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[1]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[2]))
	if v41 != int32(14) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v44 = int32(_a_F_XLogFileInitInternal_1)
	goto L5
L4:
	;
	v44 = v32
	goto L5
L5:
	;
	v45 = v35 << (uint(int32(13)) % 32) & v44
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[3])))
	if v47 != int32(1) {
		v69 = v45
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v73 = F_BasicOpenFile(m, l3, v69|int32(_a_F_XLogFileInitInternal_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[4]))
	switch v51 {
	case 0, 1, 3:
		v69 = v45
		goto L6
	case 2:
		goto L8
	case 4:
		goto L10
	default:
		goto L9
	}
L8:
	;
	v69 = v45 | int32(_a_F_XLogFileInitInternal_3)
	goto L6
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v69 = v45 | int32(_a_F_XLogFileInitInternal_4)
	goto L6
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
	F_errmsg_internal(m, int32(_a_F_XLogFileInitInternal_5), v12)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(_a_F_XLogFileInitInternal_7), int32(_a_F_XLogFileInitInternal_8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L167
	}
L15:
	;
	v747 = int32(_a_F_XLogFileInitInternal_9)
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	v749 = F_close(m, v112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = v748
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L163
	}
L16:
	;
	v725 = v12 + int32(144)
	v726 = F_unlink(m, v725)
	mBase = m.M
	v727 = F_close(m, v112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = v329
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L159
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L155
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L151
	}
L19:
	;
	if v73 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	if v78 != int32(44) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v682 = v73
	goto L22
L22:
	;
	m.G0 = v12 + int32(1168)
	return v682
L23:
	;
	v83 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_XLogFileInitInternal_10), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v94 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v94
	v97 = v12 + int32(144)
	v102 = F_pg_snprintf(m, v97, int32(1024), int32(_a_F_XLogFileInitInternal_11), v12+int32(80))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3227), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v104 = F_unlink(m, v97)
	mBase = m.M
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[1]))
	if v108&int32(4) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = int32(_a_F_XLogFileInitInternal_13)
	goto L33
L32:
	;
	v111 = int32(194)
	goto L33
L33:
	;
	v112 = F_BasicOpenFile(m, v97, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v112 < int32(0) {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v116 = int32(0)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[6])))
	v121 = m.G0
	v123 = v121 - int32(16)
	m.G0 = v123
	if v118 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(167772234)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[8])))
	if v141 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	F___clock_gettime(m, int32(1), v123)
	mBase = m.M
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v123)+8)))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	v132 = v127 + v128*int64(1000000000)
	goto L39
L38:
	;
	v132 = int64(0)
	goto L39
L39:
	;
	m.G0 = v123 + int32(16)
	goto L36
L40:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = int32(0)
	v334 = int32(2)
	v337 = int32(1)
	v339 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0])))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[8])))
	if v342 != 0 {
		goto L88
	} else {
		goto L89
	}
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0]))
	v156 = m.G0
	v158 = v156 - int32(1024)
	m.G0 = v158
	v161 = v145
	v162 = int64(0)
	v170 = int32(0)
	goto L45
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = int32(0)
	v316 = int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0]))
	v322 = F_pwrite(m, v112, int32(_a_F_XLogFileInitInternal_14), v316, base.I64_extend_i32_s(v318-v316))
	mBase = m.M
	if v322 == v316 {
		v329 = v116
		goto L40
	} else {
		goto L84
	}
L44:
	;
	if int32(0) <= v296 {
		v329 = v116
		goto L40
	} else {
		goto L83
	}
L45:
	;
	v172 = int32(0)
	if v161 == v172 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	m.G0 = v158 + int32(1024)
	goto L44
L47:
	;
	goto L46
L48:
	;
	v296 = v170
	goto L47
L49:
	;
	goto L50
L50:
	;
	v176 = v161
	v178 = v172
	goto L51
L51:
	;
	v189 = v158 + v178<<(uint(int32(3))%32)
	v190 = int32(_a_F_XLogFileInitInternal_15)
	if base.Ui32(v190) <= base.Ui32(v176) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v204 = m.G0
	v206 = v204 - int32(1024)
	m.G0 = v206
	if v198 <= int32(128) {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	goto L52
L54:
	;
	v193 = v190
	goto L56
L55:
	;
	v193 = v176
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(_a_F_XLogFileInitInternal_16)
	v198 = v178 + int32(1)
	v199 = v176 - v193
	if base.Ui32(int32(126)) < base.Ui32(v178) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	if v199 != 0 {
		v176 = v199
		v178 = v198
		goto L51
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	m.G0 = v206 + int32(1024)
	if int32(0) <= v280 {
		v161 = v199
		v162 = v162 + base.I64_extend_i32_u(v280)
		v170 = v170 + v280
		goto L45
	} else {
		goto L82
	}
L60:
	;
	v213 = v158
	v214 = v198
	v217 = int32(0)
	v221 = v162
	goto L63
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = int32(28)
	v280 = int32(-1)
	goto L59
L63:
	;
	if v214 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v280 = v232
	goto L59
L65:
	;
	if v228 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v226 = F_pwrite(m, v112, v224, v225, v221)
	mBase = m.M
	v228 = v226
	goto L65
L67:
	;
	goto L68
L68:
	;
	v227 = F_pwritev(m, v112, v213, v214, v221)
	mBase = m.M
	v228 = v227
	goto L65
L69:
	;
	v280 = int32(-1)
	goto L59
L70:
	;
	goto L71
L71:
	;
	v232 = v228 + v217
	v238 = v213
	v239 = v214
	v241 = v228
	goto L72
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v247) <= base.Ui32(v241) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v238 == v206 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v253 = v239 - int32(1)
	if v253 != 0 {
		v238 = v238 + int32(8)
		v239 = v253
		v241 = v241 - v247
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v280 = v232
	goto L59
L78:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v261 + v241
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v264 - v241
	if int32(0) < v239 {
		v213 = v206
		v214 = v239
		v217 = v232
		v221 = v221 + base.I64_extend_i32_u(v228)
		goto L63
	} else {
		goto L81
	}
L79:
	;
	v256 = v239 << (uint(int32(3)) % 32)
	if v256 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	base.MemoryCopy(m, v206, v238, v256)
	goto L78
L81:
	;
	goto L64
L82:
	;
	v296 = v280
	goto L47
L83:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	v329 = v311
	goto L40
L84:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	if v326 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v328 = v326
	goto L87
L86:
	;
	v328 = int32(51)
	goto L87
L87:
	;
	v329 = v328
	goto L40
L88:
	;
	v343 = v339
	goto L90
L89:
	;
	v343 = int64(1)
	goto L90
L90:
	;
	v347 = m.G0
	v349 = v347 - int32(16)
	m.G0 = v349
	if v132 != int64(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v329 != 0 {
		goto L16
	} else {
		goto L108
	}
L92:
	;
	F___clock_gettime(m, int32(1), v349)
	mBase = m.M
	v355 = int64(*(*int32)(unsafe.Add(mBase, uint32(v349)+8)))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
	v360 = v355 + (v356*int64(1000000000) - v132)
	goto L95
L93:
	;
	goto L94
L94:
	;
	v447 = int32(824)
	v448 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[9])) = v448 + base.I64_extend_i32_u(v337)
	v452 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[10])) = v452 + v343
	F_pgstat_count_backend_io_op(m, v334, v334, int32(7), v337, v343)
	mBase = m.M
	v457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v457)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[12])) = uint8(v457)
	m.G0 = v349 + int32(16)
	goto L91
L95:
	;
	v410 = int32(824)
	v411 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[13])) = v411 + v360
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[2]))
	v422 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v415))|base.B2i32(int32(1)<<(uint(v415)%32)&int32(_a_F_XLogFileInitInternal_17) == v422) == v422 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v427 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[14])) = v427 + v360
	v431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v431)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[15])) = uint8(v431)
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L94
L108:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[6])))
	v469 = m.G0
	v471 = v469 - int32(16)
	m.G0 = v471
	if v466 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = int32(167772233)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[3])))
	if v490 != int32(1) {
		v504 = int32(0)
		goto L114
	} else {
		goto L115
	}
L110:
	;
	F___clock_gettime(m, int32(1), v471)
	mBase = m.M
	v475 = int64(*(*int32)(unsafe.Add(mBase, uint32(v471)+8)))
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v471)))
	v480 = v475 + v476*int64(1000000000)
	goto L112
L111:
	;
	v480 = int64(0)
	goto L112
L112:
	;
	m.G0 = v471 + int32(16)
	goto L109
L113:
	;
	if v504 != 0 {
		goto L15
	} else {
		goto L120
	}
L114:
	;
	goto L113
L115:
	;
	goto L116
L116:
	;
	v495 = F_fsync(m, v112)
	mBase = m.M
	if v495 != int32(-1) {
		v504 = v495
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v504 = int32(-1)
	goto L114
L118:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	if v499 == int32(27) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v506))) = int32(0)
	v509 = int32(2)
	v511 = int32(1)
	v513 = int64(0)
	v517 = m.G0
	v519 = v517 - int32(16)
	m.G0 = v519
	if v480 != v513 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v635 = F_close(m, v112)
	mBase = m.M
	if v635 != 0 {
		goto L14
	} else {
		goto L138
	}
L122:
	;
	F___clock_gettime(m, int32(1), v519)
	mBase = m.M
	v525 = int64(*(*int32)(unsafe.Add(mBase, uint32(v519)+8)))
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v519)))
	v530 = v525 + (v526*int64(1000000000) - v480)
	goto L125
L123:
	;
	goto L124
L124:
	;
	v617 = int32(776)
	v618 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[16])) = v618 + base.I64_extend_i32_u(v511)
	v622 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[17])) = v622 + v513
	F_pgstat_count_backend_io_op(m, v509, v509, v511, v511, v513)
	mBase = m.M
	v627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v627)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[12])) = uint8(v627)
	m.G0 = v519 + int32(16)
	goto L121
L125:
	;
	v580 = int32(776)
	v581 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[18])) = v581 + v530
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[2]))
	v592 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v585))|base.B2i32(int32(1)<<(uint(v585)%32)&int32(_a_F_XLogFileInitInternal_17) == v592) == v592 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v597 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[19]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[19])) = v597 + v530
	v601 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v601)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[15])) = uint8(v601)
	goto L137
L136:
	;
	goto L137
L137:
	;
	goto L124
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = l0
	v643 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[20])))
	v645 = F_InstallXLogFileSegment(m, v12+int32(136), v12+int32(144), int32(1), l0+v643, l1)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	v682 = int32(-1)
	goto L22
L140:
	;
	F_errmsg_internal(m, v668, int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L149
	}
L141:
	;
	if v645 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v647 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v647)
	v651 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v659 = F_unlink(m, v12+int32(144))
	mBase = m.M
	v662 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	if v651 == int32(0) {
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v668 = int32(_a_F_XLogFileInitInternal_18)
	v669 = int32(3349)
	goto L140
L147:
	;
	if v662 == int32(0) {
		goto L139
	} else {
		goto L148
	}
L148:
	;
	v668 = int32(_a_F_XLogFileInitInternal_19)
	v669 = int32(3359)
	goto L140
L149:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), v669, int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L139
L151:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l3
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_20), v12+int32(96))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3216), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_21), v12+int32(16))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3241), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v725
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_22), v12-int32(-64))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3302), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_23), v12+int32(48))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3316), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_24), v12+int32(32))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3326), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogInsert(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v14 int32
	_ = v14
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v418 int64
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 int64
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int64
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int64
	_ = v529
	var v530 int64
	_ = v530
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v610 int64
	_ = v610
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int64
	_ = v625
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int64
	_ = v634
	var v637 int32
	_ = v637
	var v638 int64
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int64
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v713 int32
	_ = v713
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v840 int64
	_ = v840
	var v841 int64
	_ = v841
	var v845 int64
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v888 int64
	_ = v888
	var v889 int64
	_ = v889
	var v896 int64
	_ = v896
	var v898 int32
	_ = v898
	var v903 int64
	_ = v903
	var v904 int64
	_ = v904
	var v906 int64
	_ = v906
	var v912 int64
	_ = v912
	var v913 int64
	_ = v913
	var v914 int64
	_ = v914
	var v924 int64
	_ = v924
	var v926 int64
	_ = v926
	var v930 int64
	_ = v930
	var v933 int64
	_ = v933
	var v934 int64
	_ = v934
	var v936 int64
	_ = v936
	var v939 int64
	_ = v939
	var v944 int64
	_ = v944
	var v946 int64
	_ = v946
	var v947 int64
	_ = v947
	var v948 int64
	_ = v948
	var v950 int64
	_ = v950
	var v955 int64
	_ = v955
	var v964 int64
	_ = v964
	var v966 int64
	_ = v966
	var v970 int64
	_ = v970
	var v974 int64
	_ = v974
	var v975 int64
	_ = v975
	var v977 int64
	_ = v977
	var v983 int64
	_ = v983
	var v984 int64
	_ = v984
	var v985 int64
	_ = v985
	var v995 int64
	_ = v995
	var v997 int64
	_ = v997
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int64
	_ = v1021
	var v1023 int64
	_ = v1023
	var v1024 int64
	_ = v1024
	var v1026 int64
	_ = v1026
	var v1029 int64
	_ = v1029
	var v1034 int64
	_ = v1034
	var v1036 int64
	_ = v1036
	var v1037 int64
	_ = v1037
	var v1038 int64
	_ = v1038
	var v1040 int64
	_ = v1040
	var v1045 int64
	_ = v1045
	var v1054 int64
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int64
	_ = v1057
	var v1058 int64
	_ = v1058
	var v1061 int64
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int64
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1074 int64
	_ = v1074
	var v1075 int64
	_ = v1075
	var v1081 int64
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1093 int64
	_ = v1093
	var v1098 int64
	_ = v1098
	var v1100 int64
	_ = v1100
	var v1103 int64
	_ = v1103
	var v1108 int64
	_ = v1108
	var v1110 int64
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1112 int64
	_ = v1112
	var v1114 int64
	_ = v1114
	var v1119 int64
	_ = v1119
	var v1128 int64
	_ = v1128
	var v1132 int64
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1140 int64
	_ = v1140
	var v1142 int64
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int64
	_ = v1146
	var v1151 int64
	_ = v1151
	var v1169 int64
	_ = v1169
	var v1176 int64
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1186 int64
	_ = v1186
	var v1192 int64
	_ = v1192
	var v1193 int64
	_ = v1193
	var v1194 int64
	_ = v1194
	var v1204 int64
	_ = v1204
	var v1206 int64
	_ = v1206
	var v1222 int64
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1250 int64
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1258 int64
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1264 int64
	_ = v1264
	var v1265 int64
	_ = v1265
	var v1267 int64
	_ = v1267
	var v1273 int64
	_ = v1273
	var v1274 int64
	_ = v1274
	var v1275 int64
	_ = v1275
	var v1285 int64
	_ = v1285
	var v1287 int64
	_ = v1287
	var v1291 int64
	_ = v1291
	var v1293 int64
	_ = v1293
	var v1295 int64
	_ = v1295
	var v1298 int64
	_ = v1298
	var v1303 int64
	_ = v1303
	var v1305 int64
	_ = v1305
	var v1306 int64
	_ = v1306
	var v1307 int64
	_ = v1307
	var v1309 int64
	_ = v1309
	var v1314 int64
	_ = v1314
	var v1323 int64
	_ = v1323
	var v1327 int64
	_ = v1327
	var v1329 int64
	_ = v1329
	var v1331 int64
	_ = v1331
	var v1337 int64
	_ = v1337
	var v1338 int64
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1349 int64
	_ = v1349
	var v1356 int64
	_ = v1356
	var v1357 int64
	_ = v1357
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1381 int64
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int64
	_ = v1392
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int64
	_ = v1433
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1471 int64
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1499 int64
	_ = v1499
	var v1500 int64
	_ = v1500
	var v1502 int64
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int64
	_ = v1511
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1532 int32
	_ = v1532
	var v1549 int32
	_ = v1549
	var v1552 int64
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1565 int64
	_ = v1565
	var v1569 int64
	_ = v1569
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1614 int64
	_ = v1614
	var v1622 int64
	_ = v1622
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1690 int32
	_ = v1690
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int64
	_ = v1729
	var v1730 int64
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1747 int32
	_ = v1747
	var v1748 int64
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1751 int64
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1758 int64
	_ = v1758
	var v1761 int64
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1769 int64
	_ = v1769
	var v1773 int64
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int64
	_ = v1776
	var v1780 int64
	_ = v1780
	var v1787 int32
	_ = v1787
	var v1796 int64
	_ = v1796
	var v1798 int64
	_ = v1798
	var v1804 int64
	_ = v1804
	var v1807 int64
	_ = v1807
	var v1811 int64
	_ = v1811
	var v1813 int64
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1819 int64
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1824 int64
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1830 int64
	_ = v1830
	var v1840 int64
	_ = v1840
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1925 int32
	_ = v1925
	var v1931 int32
	_ = v1931
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1989 int32
	_ = v1989
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2101 int64
	_ = v2101
	var v2107 int32
	_ = v2107
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2136 int64
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2163 int32
	_ = v2163
	var v2169 int32
	_ = v2169
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2227 int32
	_ = v2227
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2301 int64
	_ = v2301
	var v2344 int32
	_ = v2344
	v1 = l0
	v14 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(_a_F_XLogInsert_0)
	m.G0 = v41
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[0])))
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[2])) = int32(_a_F_XLogInsert_1)
	v2344 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3])) = v2344
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[4])) = v2344
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])) = uint8(v2344)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[0])) = uint8(v2344)
	m.G0 = v41 + int32(_a_F_XLogInsert_0)
	return v2301
L2:
	;
	v2136 = int64(40)
	v2138 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v2138 <= int32(0) {
		v2301 = v2136
		goto L1
	} else {
		goto L381
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L72
	} else {
		goto L377
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L72
	} else {
		goto L373
	}
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L72
	} else {
		goto L370
	}
L6:
	;
	if l1&int32(12) != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L72
	} else {
		goto L367
	}
L9:
	;
	if v1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[6]))
	if v48 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[7]))
	v80 = v14
	v81 = v14
	v83 = v14
	goto L14
L13:
	;
	goto L12
L14:
	;
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v41+int32(56)))) = v96
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(55)))) = uint8(v99)
	goto L16
L15:
	;
	v1897 = int32(0)
	v1899 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v1899 <= v1897 {
		v2301 = v1840
		goto L1
	} else {
		goto L356
	}
L16:
	;
	v101 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[11])) = v104
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[12])) = v101
	v110 = v104 + int32(24)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[13]))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v1))))
	v117 = v114<<(uint(int32(1))%32) | l1
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v119 <= v101 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])))
	if v566&int32(1) == int32(0) {
		v582 = v540
		goto L120
	} else {
		goto L121
	}
L18:
	;
	v122 = int64(0)
	v529 = v122
	v530 = v122
	v540 = v110
	v544 = int32(_a_F_XLogInsert_2)
	v554 = v80
	v555 = v81
	v557 = v83
	v558 = v101
	goto L17
L19:
	;
	goto L20
L20:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v41)+56))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	v131 = int64(0)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+55)))
	v136 = int32(0)
	v140 = v131
	v141 = v131
	v151 = v110
	v154 = v136
	v155 = int32(_a_F_XLogInsert_2)
	v158 = v119
	v159 = v129
	v162 = v136
	v165 = v80
	v166 = v81
	v168 = v83
	v169 = v101
	goto L21
L21:
	;
	v178 = v159 + v162*int32(_a_F_XLogInsert_3)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v179 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v529 = v504
	v530 = v505
	v540 = v507
	v544 = v510
	v554 = v517
	v555 = v518
	v557 = v519
	v558 = v520
	goto L17
L23:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v182&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v504 = v140
	v505 = v141
	v507 = v151
	v509 = v154
	v510 = v155
	v513 = v158
	v514 = v159
	v517 = v165
	v518 = v166
	v519 = v168
	v520 = v169
	goto L25
L25:
	;
	v525 = v162 + int32(1)
	if v525 < v513 {
		v140 = v504
		v141 = v505
		v151 = v507
		v154 = v509
		v155 = v510
		v158 = v513
		v159 = v514
		v162 = v525
		v165 = v517
		v166 = v518
		v168 = v519
		v169 = v520
		goto L21
	} else {
		goto L119
	}
L26:
	;
	v208 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v178)+28))
	if v218 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v204 = v141
	v206 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v186 = int32(0)
	if base.B2i32(v132&int32(1) == v186)|v182&int32(2) != 0 {
		v204 = v141
		v206 = v186
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	v193 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v192))))
	v196 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v192)+4)))
	v197 = v193<<(uint(int64(32))%64) | v196
	if base.Ui64(v197) <= base.Ui64(v125) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v204 = v141
	v206 = int32(1)
	goto L26
L32:
	;
	goto L33
L33:
	;
	if base.Ui64(v141-int64(1)) < base.Ui64(v197) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v203 = v141
	goto L36
L35:
	;
	v203 = v197
	goto L36
L36:
	;
	v204 = v203
	v206 = v186
	goto L26
L37:
	;
	v219 = v206 ^ int32(1) | int32(base.Ui32(v182&int32(16))>>(uint(int32(4))%32))
	goto L39
L38:
	;
	v219 = v208
	goto L39
L39:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+16)))
	v223 = int32(6)
	if v182&v223 == v223 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v227 = v220 | int32(64)
	goto L42
L41:
	;
	v227 = v220
	goto L42
L42:
	;
	v230 = base.B2i32(v117&int32(2) != int32(0)) | v206
	if v230 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v432 = int32(0)
	if v219 == v432 {
		goto L103
	} else {
		goto L104
	}
L44:
	;
	v418 = v140
	v420 = v155
	v423 = v227
	v425 = int32(0)
	v426 = v208
	v427 = v165
	v428 = v166
	v429 = v168
	v430 = v169
	goto L43
L45:
	;
	goto L46
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	v235 = int32(0)
	if v182&int32(8) == v235 {
		v255 = v208
		v256 = v235
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[15]))
	if v258 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L48:
	;
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+12)))
	if base.Ui32(v240) < base.Ui32(int32(24)) {
		v255 = v208
		v256 = v235
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+14)))
	v249 = base.B2i32(base.Ui32(v240) < base.Ui32(v243)) & base.B2i32(base.Ui32(v243) < base.Ui32(int32(_a_F_XLogInsert_4)))
	if v249 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v250 = v243 - v240
	goto L52
L51:
	;
	v250 = int32(0)
	goto L52
L52:
	;
	if v249 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v252 = v240
	goto L55
L54:
	;
	v252 = int32(0)
	goto L55
L55:
	;
	v255 = v250
	v256 = v252
	goto L47
L56:
	;
	v335 = v178 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v335
	v338 = v255 & int32(_a_F_XLogInsert_5)
	v340 = base.B2i32(v338 != int32(0))
	if v206 != 0 {
		goto L82
	} else {
		goto L83
	}
L57:
	;
	v263 = int32(0)
	v330 = v255 & int32(_a_F_XLogInsert_5)
	v331 = v263
	v333 = v263
	goto L56
L58:
	;
	goto L59
L59:
	;
	v265 = int32(0)
	v268 = v255 & int32(_a_F_XLogInsert_5)
	if v268 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v256 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v283 = v234
	v285 = v265
	goto L62
L62:
	;
	switch v258 - int32(1) {
	case 0:
		goto L69
	case 1:
		goto L71
	case 2:
		goto L70
	default:
		v330 = v268
		v331 = int32(0)
		v333 = v265
		goto L56
	}
L63:
	;
	base.MemoryCopy(m, v41-int32(-64), v234, v256)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v273 = v268 + v256
	v274 = int32(_a_F_XLogInsert_6) - v273
	if v274 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	base.MemoryCopy(m, v41-int32(-64)+v256, v273+v234, v274)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v283 = v41 - int32(-64)
	v285 = int32(2)
	goto L62
L69:
	;
	v318 = int32(_a_F_XLogInsert_6) - v268
	v321 = F_pglz_compress(m, v283, v318, v178-int32(-64), v52)
	mBase = m.M
	v322 = int32(0)
	v327 = base.B2i32(v321+v285 < v318) & base.B2i32(v322 <= v321)
	if v327 != 0 {
		goto L79
	} else {
		goto L80
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L72
	} else {
		goto L76
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	return int64(0)
L73:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_7), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(984), int32(_a_F_XLogInsert_9))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_10), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(995), int32(_a_F_XLogInsert_9))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v328 = v321
	goto L81
L80:
	;
	v328 = v322
	goto L81
L81:
	;
	v330 = v268
	v331 = v327
	v333 = v328
	goto L56
L82:
	;
	v343 = v340 | int32(2)
	goto L84
L83:
	;
	v343 = v340
	goto L84
L84:
	;
	if v331 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178+v404))) = v407
	v418 = v140 + base.I64_extend_i32_u(v406)&int64(65535)
	v420 = v403
	v423 = v227 | int32(16)
	v425 = v331
	v426 = v255
	v427 = v256
	v428 = v405
	v429 = v406
	v430 = v169 + int32(1)
	goto L43
L86:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[15]))
	switch v345 - int32(1) {
	case 0:
		goto L90
	case 1:
		goto L92
	case 2:
		goto L91
	default:
		v376 = v343
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v338 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+44)) = v178 - int32(-64)
	v403 = v335
	v404 = int32(48)
	v405 = v376
	v406 = v333
	v407 = v333 & int32(_a_F_XLogInsert_5)
	goto L85
L90:
	;
	v376 = v343 | int32(4)
	goto L89
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L72
	} else {
		goto L96
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L72
	} else {
		goto L93
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_7), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L72
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(739), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L72
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_10), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L72
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(747), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L72
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+44)) = v234
	v387 = int32(_a_F_XLogInsert_6)
	v403 = v335
	v404 = int32(48)
	v405 = v343
	v406 = v387
	v407 = v387
	goto L85
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+48)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v178)+44)) = v234
	v392 = v178 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+40)) = v392
	v394 = v330 + v256
	*(*int32)(unsafe.Add(mBase, uint32(v178)+56)) = v234 + v394
	v397 = int32(_a_F_XLogInsert_6)
	v403 = v392
	v404 = int32(60)
	v405 = v343
	v406 = v397 - v255
	v407 = v397 - v394
	goto L85
L102:
	;
	if v154 == int32(0) {
		v462 = v432
		v463 = v445
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v444 = v418
	v445 = v423
	v446 = int32(0)
	v447 = v420
	goto L102
L104:
	;
	goto L105
L105:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v178)+28))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v178)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v437
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v178)+36))
	v444 = v418 + base.I64_extend_i32_u(v436)
	v445 = v423 | int32(32)
	v446 = v436
	v447 = v443
	goto L102
L106:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+2)) = uint16(v446)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)) = uint8(v463)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v162)
	if v230 == int32(0) {
		v486 = v151 + int32(4)
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	if v450 != v451 {
		v462 = v432
		v463 = v445
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if v453 != v454 {
		v462 = v432
		v463 = v445
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v460 = base.B2i32(v458 == v459)
	if v458 == v459 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v461 = v445 | int32(-128)
	goto L112
L111:
	;
	v461 = v445
	goto L112
L112:
	;
	v462 = v460
	v463 = v461
	goto L106
L113:
	;
	if v462 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)) = uint8(v428)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+6)) = uint16(v427)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+4)) = uint16(v429)
	if base.B2i32(v426&int32(_a_F_XLogInsert_5) == int32(0))|(v425^int32(1)) != 0 {
		v486 = v151 + int32(9)
		goto L113
	} else {
		goto L115
	}
L115:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+9)) = uint16(v426)
	v486 = v151 + int32(11)
	goto L113
L116:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v486)+8)) = v489
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v178)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v486))) = v491
	v495 = v486 + int32(12)
	goto L118
L117:
	;
	v495 = v486
	goto L118
L118:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v178)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v496
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	v504 = v444
	v505 = v204
	v507 = v495 + int32(4)
	v509 = v178
	v510 = v447
	v513 = v501
	v514 = v503
	v517 = v427
	v518 = v428
	v519 = v429
	v520 = v430
	goto L25
L119:
	;
	goto L22
L120:
	;
	v583 = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+78)))
	if v586 != 0 {
		v600 = v583
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v572 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XLogInsert[17])))
	if v572 == int32(0) {
		v582 = v540
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v575 = int32(253)
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v575)
	v578 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XLogInsert[17])))
	*(*uint16)(unsafe.Add(mBase, uint32(v540)+1)) = uint16(v578)
	v582 = v540 + int32(3)
	goto L120
L123:
	;
	if v600 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[18]))
	if v588 < int32(2) {
		v600 = v583
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v585)+20))
	if v591 != int32(2) {
		v600 = v583
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v585)+28))
	if v594 < int32(2) {
		v600 = v583
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	v600 = base.B2i32(v597 != int32(0))
	goto L123
L128:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+1)) = v602
	v604 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v582))) = uint8(v604)
	v608 = v582 + int32(5)
	goto L130
L129:
	;
	v608 = v582
	goto L130
L130:
	;
	v610 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	if v610 != int64(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if base.Ui64(int64(256)) <= base.Ui64(v610) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v638 = v529
	v639 = v608
	v640 = v544
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = int32(0)
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[10]))
	v646 = v639 - v645
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[20])) = v646
	v649 = v638 + base.I64_extend_i32_u(v646)
	v651 = int32(24)
	v655 = m.Env.Pgmem_crc32c(m, int32(-1), v645+v651, v646-v651)
	mBase = m.M
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[12]))
	if v657 != 0 {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v631
	v634 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[2]))
	v638 = v634 + v529
	v639 = v629
	v640 = v637
	goto L133
L135:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v610) {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v622 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v608))) = uint8(v622)
	v625 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+1)) = uint8(v625)
	v629 = v608 + int32(2)
	goto L134
L138:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v608)+1)) = uint32(v610)
	v618 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v608))) = uint8(v618)
	v629 = v608 + int32(5)
	goto L134
L139:
	;
	v671 = v655
	v672 = v657
	goto L142
L140:
	;
	v713 = v655
	goto L141
L141:
	;
	if base.Ui64(int64(1069547521)) <= base.Ui64(v649) {
		goto L3
	} else {
		goto L145
	}
L142:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v672)+8))
	v698 = m.Env.Pgmem_crc32c(m, v671, v696, v697)
	mBase = m.M
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	if v699 != 0 {
		v671 = v698
		v672 = v699
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v713 = v698
	goto L141
L144:
	;
	goto L143
L145:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	goto L146
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+17)) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+16)) = uint8(v117)
	*(*uint32)(unsafe.Add(mBase, uint32(v104))) = uint32(v649)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v104)+20)) = v713
	*(*int64)(unsafe.Add(mBase, uint32(v104)+8)) = int64(0)
	v751 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])))
	v752 = int32(0)
	v755 = m.G0
	v757 = v755 - int32(16)
	m.G0 = v757
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[11]))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+17)))
	if v762 != 0 {
		v775 = v752
		v776 = int32(1)
		goto L148
	} else {
		goto L149
	}
L147:
	;
	if v1840 == int64(0) {
		v80 = v554
		v81 = v555
		v83 = v557
		goto L14
	} else {
		goto L355
	}
L148:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])))
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	if v782 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L149:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+16)))
	v765 = v763 & int32(240)
	if v765 == int32(64) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v775 = int32(1)
	v776 = int32(0)
	goto L148
L151:
	;
	goto L152
L152:
	;
	if v765 == int32(224) {
		v775 = v752
		v776 = int32(0)
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v775 = v752
	v776 = int32(1)
	goto L148
L154:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L72
	} else {
		goto L351
	}
L155:
	;
	m.G0 = v757 + int32(16)
	goto L147
L156:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L72
	} else {
		goto L320
	}
L157:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v761)+20))
	v1374 = m.Env.Pgmem_crc32c(m, v1372, v761, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v761)+20)) = v1374 ^ int32(-1)
	v1381 = v1357 & int64(8191)
	if v1381 == int64(0) {
		goto L277
	} else {
		goto L278
	}
L158:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	v1241 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v1244 = base.AtomicRmwXchg32(m, v1241, int32(0), int32(1))
	if v1244 != 0 {
		goto L257
	} else {
		goto L258
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L72
	} else {
		goto L254
	}
L160:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v778)+308))
	v802 = int32(_a_F_XLogInsert_12)
	v804 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24])) = v804 + int32(1)
	if v776 != 0 {
		goto L169
	} else {
		goto L170
	}
L161:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[25])))
	if v786 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	if v782 == int32(0) {
		goto L159
	} else {
		goto L168
	}
L164:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v778)+316))
	v792 = base.B2i32(v790 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[25])) = uint8(v792)
	if v790 != int32(2) {
		goto L159
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23])) = int32(1)
	goto L160
L167:
	;
	goto L166
L168:
	;
	goto L160
L169:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[26]))
	if v809 == int32(-1) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L72
	} else {
		goto L189
	}
L172:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[27]))
	v816 = base.I32_rem_s(v814, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[26])) = v816
	v818 = v816
	goto L174
L173:
	;
	v818 = v809
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[28])) = v818
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[29]))
	v827 = F_LWLockAcquire(m, v822+v818<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L72
	} else {
		goto L175
	}
L175:
	;
	if v827 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v831 = int32(_a_F_XLogInsert_13)
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[26]))
	v837 = base.I32_rem_s(v833+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[26])) = v837
	goto L178
L177:
	;
	goto L178
L178:
	;
	v840 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8]))
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v778)+152))
	if v840 != v841 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8])) = v841
	v845 = v841
	goto L181
L180:
	;
	v845 = v840
	goto L181
L181:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778)+160)))
	if v846 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	if v780&int32(1)&base.B2i32(base.Ui64(v845) <= base.Ui64(v530-int64(1))) != 0 {
		goto L158
	} else {
		goto L187
	}
L183:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v778)+164))
	v852 = base.B2i32(int32(0) < v850)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])) = uint8(v852)
	if int32(0) < v850 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v855 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])) = uint8(v855)
	goto L182
L186:
	;
	goto L158
L187:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L72
	} else {
		goto L188
	}
L188:
	;
	v866 = int32(_a_F_XLogInsert_12)
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24])) = v868 - int32(1)
	v1840 = int64(0)
	goto L155
L189:
	;
	if v775 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v882 = base.AtomicRmwXchg32(m, v879, int32(0), int32(1))
	if v882 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L192
L192:
	;
	v1006 = int32(0)
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v1011 = base.AtomicRmwXchg32(m, v1008, v1006, int32(1))
	if v1011 != 0 {
		goto L215
	} else {
		goto L216
	}
L193:
	;
	F_s_lock(m, v879, int32(_a_F_XLogInsert_14), int32(1134), int32(_a_F_XLogInsert_15))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L72
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v879)+16))
	v889 = *(*int64)(unsafe.Add(mBase, uint32(v879)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v879)+16)) = v889
	v896 = v889 + base.I64_extend_i32_s((v877+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v879)+8)) = v896
	v898 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v879))), uint32(v898))
	v903 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30])))
	v904 = base.I64_div_u_s(v889, v903)
	v906 = v889 - v904*v903
	if base.Ui64(v906) <= base.Ui64(int64(8151)) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	goto L195
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v757)+8)) = v930
	v933 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30])))
	v934 = base.I64_div_u_s(v896, v933)
	v936 = v896 - v934*v933
	if base.Ui64(v936) <= base.Ui64(int64(8151)) {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	v926 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v930 = v904*v926 + v924&int64(4294967295)
	goto L197
L199:
	;
	v924 = v906 + int64(40)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v912 = v906 - int64(8152)
	v913 = int64(8168)
	v914 = base.I64_div_u_s(v912, v913)
	v924 = v912 - v914*v913 + v914<<(uint(int64(13))%64) + int64(8216)
	goto L198
L202:
	;
	v966 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v970 = v934*v966 + v964&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v970
	v974 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30])))
	v975 = base.I64_div_u_s(v888, v974)
	v977 = v888 - v975*v974
	if base.Ui64(v977) <= base.Ui64(int64(8151)) {
		goto L212
	} else {
		goto L213
	}
L203:
	;
	v939 = int64(0)
	if v936 == v939 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v946 = v936 - int64(8152)
	v947 = int64(8168)
	v948 = base.I64_div_u_s(v946, v947)
	v950 = v948 << (uint(int64(13)) % 64)
	v955 = v946 - v948*v947
	if v955 == int64(0) {
		v964 = v950 - int64(-8192)
		goto L202
	} else {
		goto L209
	}
L206:
	;
	v944 = v939
	goto L208
L207:
	;
	v944 = v936 + int64(40)
	goto L208
L208:
	;
	v964 = v944
	goto L202
L209:
	;
	v964 = v955 + v950 + int64(8216)
	goto L202
L210:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v761)+8)) = v975*v997 + v995&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v778)+152)) = v930
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8])) = v930
	v1356 = v970
	v1357 = v930
	goto L157
L211:
	;
	v997 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	goto L210
L212:
	;
	v995 = v977 + int64(40)
	goto L211
L213:
	;
	goto L214
L214:
	;
	v983 = v977 - int64(8152)
	v984 = int64(8168)
	v985 = base.I64_div_u_s(v983, v984)
	v995 = v983 - v985*v984 + v985<<(uint(int64(13))%64) + int64(8216)
	goto L211
L215:
	;
	F_s_lock(m, v1008, int32(_a_F_XLogInsert_14), int32(1183), int32(_a_F_XLogInsert_16))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L72
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v1017 = int32(8)
	v1018 = v757 + v1017
	v1021 = *(*int64)(unsafe.Add(mBase, uint32(v1008)+8))
	v1023 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30])))
	v1024 = base.I64_div_u_s(v1021, v1023)
	v1026 = v1021 - v1024*v1023
	if base.Ui64(v1026) <= base.Ui64(int64(8151)) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L217
L219:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31]))
	v1057 = base.I64_extend_i32_s(v1056)
	v1058 = v1024 * v1057
	v1061 = v1058 + v1054&int64(4294967295)
	v1063 = v1056 - int32(1)
	v1064 = base.I64_extend_i32_s(v1063)
	v1065 = v1061 & v1064
	if v1065 == int64(0) {
		goto L228
	} else {
		goto L229
	}
L220:
	;
	v1029 = int64(0)
	if v1026 == v1029 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	v1036 = v1026 - int64(8152)
	v1037 = int64(8168)
	v1038 = base.I64_div_u_s(v1036, v1037)
	v1040 = v1038 << (uint(int64(13)) % 64)
	v1045 = v1036 - v1038*v1037
	if v1045 == int64(0) {
		v1054 = v1040 - int64(-8192)
		goto L219
	} else {
		goto L226
	}
L223:
	;
	v1034 = v1029
	goto L225
L224:
	;
	v1034 = v1026 + int64(40)
	goto L225
L225:
	;
	v1054 = v1034
	goto L219
L226:
	;
	v1054 = v1045 + v1040 + int64(8216)
	goto L219
L227:
	;
	if v1065 == int64(0) {
		v1690 = v1006
		goto L156
	} else {
		goto L253
	}
L228:
	;
	v1068 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1008))), uint32(v1068))
	*(*int64)(unsafe.Add(mBase, uint32(v1018))) = v1061
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v1061
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1074 = v1021 + int64(24)
	v1075 = *(*int64)(unsafe.Add(mBase, uint32(v1008)+16))
	if base.Ui64(v1026) <= base.Ui64(int64(8151)) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1093 = v1026 + int64(40)
	goto L233
L232:
	;
	v1081 = v1026 - int64(8152)
	v1082 = int64(8168)
	v1083 = base.I64_div_u_s(v1081, v1082)
	v1093 = v1081 - v1083*v1082 + v1083<<(uint(int64(13))%64) + int64(8216)
	goto L233
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1018))) = v1093&int64(4294967295) + v1058
	v1098 = base.I64_div_u_s(v1074, v1023)
	v1100 = v1074 - v1098*v1023
	if base.Ui64(v1100) <= base.Ui64(int64(8151)) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1132 = v1057*v1098 + v1128&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v1132
	v1135 = v1063 & base.I32_wrap_i64(v1132)
	if v1135 == int32(0) {
		v1176 = v1074
		goto L242
	} else {
		goto L243
	}
L235:
	;
	v1103 = int64(0)
	if v1100 == v1103 {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	goto L237
L237:
	;
	v1110 = v1100 - int64(8152)
	v1111 = int64(8168)
	v1112 = base.I64_div_u_s(v1110, v1111)
	v1114 = v1112 << (uint(int64(13)) % 64)
	v1119 = v1110 - v1112*v1111
	if v1119 == int64(0) {
		v1128 = v1114 - int64(-8192)
		goto L234
	} else {
		goto L241
	}
L238:
	;
	v1108 = v1103
	goto L240
L239:
	;
	v1108 = v1100 + int64(40)
	goto L240
L240:
	;
	v1128 = v1108
	goto L234
L241:
	;
	v1128 = v1119 + v1114 + int64(8216)
	goto L234
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1008)+16)) = v1021
	*(*int64)(unsafe.Add(mBase, uint32(v1008)+8)) = v1176
	v1181 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1008))), uint32(v1181))
	v1184 = base.I64_div_u_s(v1075, v1023)
	v1186 = v1075 - v1023*v1184
	if base.Ui64(v1186) <= base.Ui64(int64(8151)) {
		goto L250
	} else {
		goto L251
	}
L243:
	;
	v1140 = v1132 + base.I64_extend_i32_u(v1056-v1135)
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v1140
	v1142 = base.I64_div_u_s(v1140, v1057)
	v1145 = base.I32_wrap_i64(v1140) & int32(_a_F_XLogInsert_17)
	v1146 = v1140 & v1064
	if v1146&int64(35184372080640) == int64(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1151 = v1142 * v1023
	if v1145 == int32(0) {
		v1176 = v1151
		goto L242
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1169 = v1142*v1023 + (int64(base.Ui64(v1146)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v1145 == int32(0) {
		v1176 = v1169
		goto L242
	} else {
		goto L248
	}
L247:
	;
	v1176 = v1151 + base.I64_extend_i32_u(v1145-int32(40))
	goto L242
L248:
	;
	v1176 = v1169 + base.I64_extend_i32_u(v1145-int32(24))
	goto L242
L249:
	;
	v1206 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	*(*int64)(unsafe.Add(mBase, uint32(v761+v1017))) = v1184*v1206 + v1204&int64(4294967295)
	goto L227
L250:
	;
	v1204 = v1186 + int64(40)
	goto L249
L251:
	;
	goto L252
L252:
	;
	v1192 = v1186 - int64(8152)
	v1193 = int64(8168)
	v1194 = base.I64_div_u_s(v1192, v1193)
	v1204 = v1192 - v1194*v1193 + v1194<<(uint(int64(13))%64) + int64(8216)
	goto L249
L253:
	;
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(v757)+8))
	v1356 = v1222
	v1357 = v1223
	goto L157
L254:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_18), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L72
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_14), int32(779), int32(_a_F_XLogInsert_19))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L72
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_s_lock(m, v1241, int32(_a_F_XLogInsert_14), int32(1134), int32(_a_F_XLogInsert_15))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L72
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1250 = *(*int64)(unsafe.Add(mBase, uint32(v1241)+16))
	v1251 = *(*int64)(unsafe.Add(mBase, uint32(v1241)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1241)+16)) = v1251
	v1258 = v1251 + base.I64_extend_i32_s((v1239+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v1241)+8)) = v1258
	v1260 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1241))), uint32(v1260))
	v1264 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30])))
	v1265 = base.I64_div_u_s(v1251, v1264)
	v1267 = v1251 - v1265*v1264
	if base.Ui64(v1267) <= base.Ui64(int64(8151)) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L259
L261:
	;
	v1287 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v1291 = v1265*v1287 + v1285&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v757)+8)) = v1291
	v1293 = base.I64_div_u_s(v1258, v1264)
	v1295 = v1258 - v1293*v1264
	if base.Ui64(v1295) <= base.Ui64(int64(8151)) {
		goto L266
	} else {
		goto L267
	}
L262:
	;
	v1285 = v1267 + int64(40)
	goto L261
L263:
	;
	goto L264
L264:
	;
	v1273 = v1267 - int64(8152)
	v1274 = int64(8168)
	v1275 = base.I64_div_u_s(v1273, v1274)
	v1285 = v1273 - v1275*v1274 + v1275<<(uint(int64(13))%64) + int64(8216)
	goto L261
L265:
	;
	v1327 = v1293*v1287 + v1323&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v1327
	v1329 = base.I64_div_u_s(v1250, v1264)
	v1331 = v1250 - v1329*v1264
	if base.Ui64(v1331) <= base.Ui64(int64(8151)) {
		goto L274
	} else {
		goto L275
	}
L266:
	;
	v1298 = int64(0)
	if v1295 == v1298 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	v1305 = v1295 - int64(8152)
	v1306 = int64(8168)
	v1307 = base.I64_div_u_s(v1305, v1306)
	v1309 = v1307 << (uint(int64(13)) % 64)
	v1314 = v1305 - v1307*v1306
	if v1314 == int64(0) {
		v1323 = v1309 - int64(-8192)
		goto L265
	} else {
		goto L272
	}
L269:
	;
	v1303 = v1298
	goto L271
L270:
	;
	v1303 = v1295 + int64(40)
	goto L271
L271:
	;
	v1323 = v1303
	goto L265
L272:
	;
	v1323 = v1314 + v1309 + int64(8216)
	goto L265
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v761)+8)) = v1329*v1287 + v1349&int64(4294967295)
	v1356 = v1327
	v1357 = v1291
	goto L157
L274:
	;
	v1349 = v1331 + int64(40)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1337 = v1331 - int64(8152)
	v1338 = int64(8168)
	v1339 = base.I64_div_u_s(v1337, v1338)
	v1349 = v1337 - v1339*v1338 + v1339<<(uint(int64(13))%64) + int64(8216)
	goto L273
L277:
	;
	v1386 = int32(0)
	goto L279
L278:
	;
	v1386 = int32(_a_F_XLogInsert_6) - base.I32_wrap_i64(v1381)
	goto L279
L279:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	v1388 = F_GetXLogBuffer(m, v1357, v801)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L72
	} else {
		goto L280
	}
L280:
	;
	v1392 = v1357
	v1403 = v1388
	v1406 = v1386
	v1410 = int32(_a_F_XLogInsert_2)
	v1413 = v752
	goto L281
L281:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+4))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+8))
	if v1406 < v1429 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v775 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L283:
	;
	v1433 = v1392
	v1444 = v1403
	v1447 = v1406
	v1448 = v1429
	v1449 = v1428
	v1454 = v1413
	goto L286
L284:
	;
	v1511 = v1392
	v1522 = v1403
	v1525 = v1406
	v1526 = v1429
	v1527 = v1428
	v1532 = v1413
	goto L285
L285:
	;
	if v1526 != 0 {
		goto L302
	} else {
		goto L303
	}
L286:
	;
	if v1447 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1511 = v1500
	v1522 = v1492
	v1525 = v1507
	v1526 = v1494
	v1527 = v1493
	v1532 = v1474
	goto L285
L288:
	;
	base.MemoryCopy(m, v1444, v1449, v1447)
	goto L290
L289:
	;
	goto L290
L290:
	;
	v1471 = v1433 + base.I64_extend_i32_s(v1447)
	v1472 = F_GetXLogBuffer(m, v1471, v801)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L72
	} else {
		goto L291
	}
L291:
	;
	v1474 = v1447 + v1454
	*(*int32)(unsafe.Add(mBase, uint32(v1472)+16)) = v1387 - v1474
	v1477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1472)+2)))
	v1478 = int32(1)
	v1479 = v1477 | v1478
	*(*uint16)(unsafe.Add(mBase, uint32(v1472)+2)) = uint16(v1479)
	v1484 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31]))
	v1490 = base.B2i32(v1471&base.I64_extend_i32_s(v1484-v1478) == int64(0))
	if v1471&base.I64_extend_i32_s(v1484-v1478) == int64(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1491 = int32(40)
	goto L294
L293:
	;
	v1491 = int32(24)
	goto L294
L294:
	;
	v1492 = v1472 + v1491
	v1493 = v1447 + v1449
	v1494 = v1448 - v1447
	if v1471&base.I64_extend_i32_s(v1484-v1478) == int64(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1499 = int64(40)
	goto L297
L296:
	;
	v1499 = int64(24)
	goto L297
L297:
	;
	v1500 = v1499 + v1471
	v1502 = v1500 & int64(8191)
	if v1502 == int64(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1507 = int32(0)
	goto L300
L299:
	;
	v1507 = int32(_a_F_XLogInsert_6) - base.I32_wrap_i64(v1502)
	goto L300
L300:
	;
	if v1507 < v1494 {
		v1433 = v1500
		v1444 = v1492
		v1447 = v1507
		v1448 = v1494
		v1449 = v1493
		v1454 = v1474
		goto L286
	} else {
		goto L301
	}
L301:
	;
	goto L287
L302:
	;
	base.MemoryCopy(m, v1522, v1527, v1526)
	goto L304
L303:
	;
	goto L304
L304:
	;
	v1549 = v1525 - v1526
	v1552 = v1511 + base.I64_extend_i32_s(v1526)
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1410)))
	if v1553 != 0 {
		v1392 = v1552
		v1403 = v1522 + v1526
		v1406 = v1549
		v1410 = v1553
		v1413 = v1526 + v1532
		goto L281
	} else {
		goto L305
	}
L305:
	;
	goto L282
L306:
	;
	if v1622 != v1356 {
		goto L154
	} else {
		goto L315
	}
L307:
	;
	v1622 = (v1552 + int64(7)) & int64(-8)
	goto L306
L308:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31]))
	if v1552&base.I64_extend_i32_s(v1557-int32(1)) == int64(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1565 = v1552 + base.I64_extend_i32_s(v1549)
	if base.Ui64(v1356) <= base.Ui64(v1565) {
		v1622 = v1565
		goto L306
	} else {
		goto L310
	}
L310:
	;
	v1569 = v1565
	goto L311
L311:
	;
	v1605 = F_GetXLogBuffer(m, v1569, v801)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L72
	} else {
		goto L313
	}
L312:
	;
	v1622 = v1614
	goto L306
L313:
	;
	v1607 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1605)+16)) = v1607
	*(*int64)(unsafe.Add(mBase, uint32(v1605)+8)) = v1607
	*(*int64)(unsafe.Add(mBase, uint32(v1605))) = v1607
	v1614 = v1569 - int64(-8192)
	if base.Ui64(v1614) < base.Ui64(v1356) {
		v1569 = v1614
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v1659 = int32(1)
	if v751&int32(2) != 0 {
		v1690 = v1659
		goto L156
	} else {
		goto L316
	}
L316:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[29]))
	v1666 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[28]))
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	if v1668 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1669 = int32(0)
	goto L319
L318:
	;
	v1669 = v1666
	goto L319
L319:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1663+v1669<<(uint(int32(7))%32))+24)) = v1357
	v1690 = v1659
	goto L156
L320:
	;
	v1714 = int32(_a_F_XLogInsert_12)
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24])) = v1716 - int32(1)
	v1721 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1721)))
	if v1722 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1723 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1721)+70)) = uint8(v1723)
	goto L323
L322:
	;
	goto L323
L323:
	;
	if v600 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v1727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1726)+78)) = uint8(v1727)
	goto L326
L325:
	;
	goto L326
L326:
	;
	v1729 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	v1730 = *(*int64)(unsafe.Add(mBase, uint32(v757)+8))
	if base.Ui64(int64(8192)) <= base.Ui64(v1729^v1730) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v1738 = base.AtomicRmwXchg32(m, v1735, int32(440), int32(1))
	if v1738 != 0 {
		goto L330
	} else {
		goto L331
	}
L328:
	;
	goto L329
L329:
	;
	if v775 != 0 {
		goto L339
	} else {
		goto L340
	}
L330:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	F_s_lock(m, v1740+int32(440), int32(_a_F_XLogInsert_14), int32(968), int32(_a_F_XLogInsert_19))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L72
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1748 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	v1750 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v1751 = *(*int64)(unsafe.Add(mBase, uint32(v1750)+184))
	if base.Ui64(v1751) < base.Ui64(v1748) {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L332
L334:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1750)+184)) = v1748
	goto L336
L335:
	;
	goto L336
L336:
	;
	v1754 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1750)+440)), uint32(v1754))
	v1758 = int64(0)
	v1761 = base.AtomicRmwCmpxchg64(m, v1750, int32(280), v1758, v1758)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[33])) = v1761
	v1765 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	v1769 = base.AtomicRmwCmpxchg64(m, v1765, int32(272), v1758, v1758)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[34])) = v1769
	goto L329
L337:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[35])) = v1773
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[36])) = v1776
	v1840 = v1773
	goto L155
L338:
	;
	v1813 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v761))))
	v1815 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[37])) = uint8(v1815)
	v1817 = int32(_a_F_XLogInsert_20)
	v1819 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[38])) = v1813 + v1819
	v1822 = int32(_a_F_XLogInsert_21)
	v1824 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[39]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[39])) = v1824 + int64(1)
	v1828 = int32(_a_F_XLogInsert_22)
	v1830 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[40]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[40])) = v1830 + base.I64_extend_i32_s(v558)
	v1840 = v1811
	goto L155
L339:
	;
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	F_XLogFlush(m, v1773)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L72
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1804 = *(*int64)(unsafe.Add(mBase, uint32(v757)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[36])) = v1804
	v1807 = *(*int64)(unsafe.Add(mBase, uint32(v757)))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[35])) = v1807
	if v1690 == int32(0) {
		v1840 = v1807
		goto L155
	} else {
		goto L350
	}
L342:
	;
	v1776 = *(*int64)(unsafe.Add(mBase, uint32(v757)+8))
	if v1690 == int32(0) {
		goto L337
	} else {
		goto L343
	}
L343:
	;
	v1780 = v1776 + int64(24)
	if base.Ui64(int64(8192)) <= base.Ui64(v1780^v1776) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31]))
	if v1780&base.I64_extend_i32_s(v1787-int32(1)^int32(_a_F_XLogInsert_17)) == int64(0) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v1798 = v1780
	goto L346
L346:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[35])) = v1798
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[36])) = v1776
	v1811 = v1798
	goto L338
L347:
	;
	v1796 = int64(64)
	goto L349
L348:
	;
	v1796 = int64(48)
	goto L349
L349:
	;
	v1798 = v1796 + v1776
	goto L346
L350:
	;
	v1811 = v1807
	goto L338
L351:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L72
	} else {
		goto L352
	}
L352:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_23), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L72
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_14), int32(1367), int32(_a_F_XLogInsert_24))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L72
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	goto L15
L356:
	;
	v1903 = v1899 & int32(7)
	v1905 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	if base.Ui32(int32(8)) <= base.Ui32(v1899) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1925 = v1897
	v1931 = int32(0)
	goto L360
L358:
	;
	v1989 = v1897
	goto L359
L359:
	;
	v2027 = int32(0)
	v2028 = v1989
	goto L364
L360:
	;
	v1951 = v1905 + v1925*int32(_a_F_XLogInsert_3)
	v1952 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[41]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[42]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[43]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[44]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[45]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[46]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951)+uint32(_c_F_XLogInsert[47]))) = uint8(v1952)
	*(*uint8)(unsafe.Add(mBase, uint32(v1951))) = uint8(v1952)
	v1968 = int32(8)
	v1969 = v1925 + v1968
	v1971 = v1931 + v1968
	if v1971 != v1899&int32(2147483640) {
		v1925 = v1969
		v1931 = v1971
		goto L360
	} else {
		goto L362
	}
L361:
	;
	if v1903 == int32(0) {
		v2301 = v1840
		goto L1
	} else {
		goto L363
	}
L362:
	;
	goto L361
L363:
	;
	v1989 = v1969
	goto L359
L364:
	;
	v2055 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1905+v2028*int32(_a_F_XLogInsert_3)))) = uint8(v2055)
	v2057 = int32(1)
	v2060 = v2027 + v2057
	if v2060 != v1903 {
		v2027 = v2060
		v2028 = v2028 + v2057
		goto L364
	} else {
		goto L366
	}
L365:
	;
	v2301 = v1840
	goto L1
L366:
	;
	goto L365
L367:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_25), int32(0))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L72
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(480), int32(_a_F_XLogInsert_26))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L72
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = l1
	F_errmsg_internal(m, int32(_a_F_XLogInsert_27), v41+int32(48))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L72
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(489), int32(_a_F_XLogInsert_26))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L72
	} else {
		goto L372
	}
L372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L373:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_28), int32(0))
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L72
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = int32(-1)
	v2101 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v2101
	F_errdetail_internal(m, int32(_a_F_XLogInsert_29), v41+int32(32))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L72
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(874), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L72
	} else {
		goto L376
	}
L376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L377:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_30), int32(0))
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L72
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v117 & int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(1069547520)
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v649
	F_errdetail_internal(m, int32(_a_F_XLogInsert_31), v41)
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L72
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(919), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L72
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	v2142 = v2138 & int32(7)
	v2144 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	if base.Ui32(int32(8)) <= base.Ui32(v2138) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v2163 = v14
	v2169 = v14
	goto L385
L383:
	;
	v2227 = v14
	goto L384
L384:
	;
	v2264 = v14
	v2265 = v2227
	goto L389
L385:
	;
	v2189 = v2144 + v2163*int32(_a_F_XLogInsert_3)
	v2190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[41]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[42]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[43]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[44]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[45]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[46]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189)+uint32(_c_F_XLogInsert[47]))) = uint8(v2190)
	*(*uint8)(unsafe.Add(mBase, uint32(v2189))) = uint8(v2190)
	v2206 = int32(8)
	v2207 = v2163 + v2206
	v2209 = v2169 + v2206
	if v2209 != v2138&int32(2147483640) {
		v2163 = v2207
		v2169 = v2209
		goto L385
	} else {
		goto L387
	}
L386:
	;
	if v2142 == int32(0) {
		v2301 = v2136
		goto L1
	} else {
		goto L388
	}
L387:
	;
	goto L386
L388:
	;
	v2227 = v2207
	goto L384
L389:
	;
	v2292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2144+v2265*int32(_a_F_XLogInsert_3)))) = uint8(v2292)
	v2294 = int32(1)
	v2297 = v2264 + v2294
	if v2297 != v2142 {
		v2264 = v2297
		v2265 = v2265 + v2294
		goto L389
	} else {
		goto L391
	}
L390:
	;
	v2301 = v2136
	goto L1
L391:
	;
	goto L390
}
func F_XLogPageRead(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v348 int64
	_ = v348
	var v357 int64
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v515 int64
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int64
	_ = v592
	var v596 int64
	_ = v596
	var v597 int64
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v678 int32
	_ = v678
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v684 int64
	_ = v684
	var v687 int64
	_ = v687
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int64
	_ = v734
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int64
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int64
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int64
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1121 int32
	_ = v1121
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1149 int64
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int64
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int64
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1178 int64
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1218 int64
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1245 int32
	_ = v1245
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int64
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int64
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1302 int64
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1341 int32
	_ = v1341
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1373 int64
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1404 int64
	_ = v1404
	var v1410 int32
	_ = v1410
	var v1425 int32
	_ = v1425
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1461 int64
	_ = v1461
	var v1462 int64
	_ = v1462
	var v1466 int64
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int64
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1492 int64
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1504 int64
	_ = v1504
	var v1505 int64
	_ = v1505
	var v1509 int64
	_ = v1509
	var v1559 int32
	_ = v1559
	var v1560 int64
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1576 int64
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1596 int32
	_ = v1596
	var v1597 int64
	_ = v1597
	var v1601 int64
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1615 int32
	_ = v1615
	var v1618 int64
	_ = v1618
	var v1621 int64
	_ = v1621
	var v1622 int64
	_ = v1622
	var v1623 int64
	_ = v1623
	var v1626 int64
	_ = v1626
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1644 int32
	_ = v1644
	var v1649 int64
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1684 int64
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1697 int32
	_ = v1697
	var v1704 int32
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1728 int32
	_ = v1728
	var v1729 int64
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1741 int64
	_ = v1741
	var v1742 int64
	_ = v1742
	var v1746 int64
	_ = v1746
	var v1796 int32
	_ = v1796
	var v1797 int64
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1808 int32
	_ = v1808
	var v1813 int64
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1833 int32
	_ = v1833
	var v1834 int64
	_ = v1834
	var v1838 int64
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1859 int64
	_ = v1859
	var v1860 int64
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1876 int64
	_ = v1876
	var v1878 int64
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1959 int32
	_ = v1959
	v4 = l3
	v6 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(1216)
	m.G0 = v34
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v40 = base.I32_wrap_i64(l1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if v44 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v115 = (v37 - int32(1)) & v40
	v118 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v112))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2])) = v118
	v120 = int64(32)
	v122 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(v120) % 64)))
	v124 = l1 + base.I64_extend_i32_s(l2)
	if v113 != 0 {
		v143 = v6
		v144 = int32(0)
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v112 = v37
	v113 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v49 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v51 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v37))
	if v49 == v51 {
		v112 = v37
		v113 = v6
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v53 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[3])))
	if v55 != v53 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v101 = int32(_a_F_XLogPageRead_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	v103 = F_close(m, v102)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v112 = v111
	v113 = v53
	goto L1
L7:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[5])))
	if v59&int32(1) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[6]))
	v70 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[7]))
	v72 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v73 = base.I64_div_u_s(v70, v72)
	goto L9
L9:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v65-int32(1))+v73) <= base.Ui64(v49)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v78 = F_GetRedoRecPtr(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[6]))
	v90 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[7]))
	v92 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v93 = base.I64_div_u_s(v90, v92)
	goto L13
L13:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v85-int32(1))+v93) <= base.Ui64(v83)) == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_RequestCheckpoint(m, int32(128))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	m.G0 = v34 + int32(1216)
	return v1959
L17:
	;
	v150 = v144
	v155 = v143
	v167 = v6
	goto L27
L18:
	;
	v130 = int32(_a_F_XLogPageRead_1)
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v132 == int32(3) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v143 = v130
	v144 = int32(3)
	goto L17
L20:
	;
	v136 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(v124) <= base.Ui64(v136) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v143 = v130
	v144 = int32(2)
	goto L17
L23:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v138 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v1959 = int32(-2)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v143 = v130
	v144 = int32(1)
	goto L17
L27:
	;
	switch v150 {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L30
	default:
		goto L32
	}
L28:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1940 {
		goto L438
	} else {
		goto L439
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9])) = v115
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v155
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[11])))
	v1455 = m.G0
	v1457 = v1455 - int32(16)
	m.G0 = v1457
	if v1452 != 0 {
		goto L352
	} else {
		goto L353
	}
L31:
	;
	v150 = int32(2)
	v155 = v1425
	goto L27
L32:
	;
	v1404 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(int64(8191)) < base.Ui64(v1404^l1) {
		v1425 = int32(_a_F_XLogPageRead_1)
		goto L31
	} else {
		goto L350
	}
L33:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+5)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[12])))
	if v184 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	v150 = int32(1)
	v167 = v176
	goto L27
L35:
	;
	v203 = int32(1)
	v204 = v167 & v203
	v215 = v202
	v229 = int32(0)
	goto L45
L36:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	if v188 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v199 = int32(2)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13])) = v199
	v202 = v199
	goto L35
L39:
	;
	if v188 != int32(3) {
		v202 = v188
		goto L35
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v196)
	v199 = int32(1)
	goto L38
L42:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v192&int32(1) != 0 {
		v202 = v188
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L11
	} else {
		goto L347
	}
L45:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])))
	if v243 != 0 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L11
	} else {
		goto L344
	}
L47:
	;
	v473 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v473)
	if base.Ui32(int32(2)) <= base.Ui32(v469-int32(1)) {
		goto L122
	} else {
		goto L123
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13])) = v433
	if v215 == v433 {
		v468 = v431
		v469 = v215
		goto L47
	} else {
		goto L105
	}
L49:
	;
	v431 = v428
	v433 = int32(1)
	goto L48
L50:
	;
	if v204 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v419 = int32(0)
	if v215 != int32(2) {
		v468 = v419
		v469 = v215
		goto L47
	} else {
		goto L103
	}
L53:
	;
	v1959 = int32(-2)
	goto L16
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v215-int32(1)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v280 = F_WalRcvStreaming(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L72
	}
L57:
	;
	if v215 == int32(3) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v267 != int32(1) {
		goto L29
	} else {
		goto L64
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v256
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_2), v34)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3765), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v270 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	if v270 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v275 == int32(0) {
		goto L29
	} else {
		goto L70
	}
L69:
	;
	goto L29
L70:
	;
	v431 = int32(1)
	v433 = int32(3)
	goto L48
L71:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[16]))
	if v302 != int32(1) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v280 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L11
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	v289 = F_LWLockAcquire(m, v285+int32(1152), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L11
	} else {
		goto L77
	}
L76:
	;
	goto L71
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[18]))
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v292)+320)) = uint8(v293)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	F_LWLockRelease(m, v296+int32(1152))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L71
L79:
	;
	v313 = m.G0
	v314 = int32(16)
	v315 = v313 - v314
	m.G0 = v315
	F_gettimeofday(m, v315)
	mBase = m.M
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v315)+8)))
	m.G0 = v315 + v314
	v327 = v319 + v318*int64(1000000) - int64(946684800000000)
	goto L83
L80:
	;
	v305 = F_rescanLatestTimeLine(m, v179, v178)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	if v305 == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v428 = int32(0)
	goto L49
L83:
	;
	v329 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19]))
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[20]))
	goto L84
L84:
	;
	if base.B2i32(base.I64_extend_i32_s(v331)*int64(1000) <= v327-v329) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[20]))
	v342 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19]))
	if v327 <= v342 {
		v360 = int32(0)
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v415 = v327
	goto L87
L87:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19])) = v415
	v428 = int32(0)
	goto L49
L88:
	;
	v364 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v348 = v327 - v342
	if base.B2i32(int64(0) < v342)^base.B2i32(v348 < v327)|base.B2i32(int64(2147483646000) < v348) != 0 {
		v360 = int32(2147483647)
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v357 = base.I64_div_s(v348+int64(999), int64(1000))
	v360 = base.I32_wrap_i64(v357)
	goto L89
L92:
	;
	if v364 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+180)) = base.I32_wrap_i64(v124)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+176)) = base.I32_wrap_i64(int64(base.Ui64(v124) >> (uint(v120) % 64)))
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_5), v34+int32(176))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L11
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_KnownAssignedTransactionIdsIdleMaintenance(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L98
	}
L96:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3744), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v386 = F_WaitLatch(m, v381+int32(4), int32(41), v340-v360, int32(150994948))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v389+int32(4)))) = int32(0)
	goto L100
L100:
	;
	v397 = m.G0
	v398 = int32(16)
	v399 = v397 - v398
	m.G0 = v399
	F_gettimeofday(m, v399)
	mBase = m.M
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v399)))
	v403 = int64(*(*int32)(unsafe.Add(mBase, uint32(v399)+8)))
	m.G0 = v399 + v398
	goto L101
L101:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v415 = v403 + v402*int64(1000000) - int64(946684800000000)
	goto L87
L103:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[12])))
	if v423&int32(1) == int32(0) {
		v468 = v419
		v469 = v215
		goto L47
	} else {
		goto L104
	}
L104:
	;
	v428 = v419
	goto L49
L105:
	;
	v438 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	if v438 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v215<<(uint(int32(2))%32))+uint32(_c_F_XLogPageRead[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v442
	v447 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])))
	if v447 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v468 = v431
	v469 = v467
	goto L47
L110:
	;
	v448 = int32(_a_F_XLogPageRead_6)
	goto L112
L111:
	;
	v448 = int32(_a_F_XLogPageRead_7)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+168)) = v448
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451<<(uint(int32(2))%32))+uint32(_c_F_XLogPageRead[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+164)) = v454
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_8), v34+int32(160))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3782), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	goto L109
L115:
	;
	goto L46
L116:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+80))
	if v1355 != 0 {
		goto L339
	} else {
		goto L340
	}
L117:
	;
	v1267 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L11
	} else {
		goto L324
	}
L118:
	;
	v150 = int32(3)
	v155 = v1245
	goto L27
L119:
	;
	v1135 = F_WalRcvStreaming(m)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L11
	} else {
		goto L300
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23])) = v733
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	v742 = F_LWLockAcquire(m, v738+int32(1152), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L11
	} else {
		goto L191
	}
L121:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	v725 = F_tliOfPointInHistory(m, v4, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L11
	} else {
		goto L186
	}
L122:
	;
	if v469 != int32(3) {
		goto L44
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v517 {
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[25])))
	if (v468|(v482^int32(-1)))&int32(1) != 0 {
		v497 = v468
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v499 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[25])) = uint8(v499)
	if v497 == v499 {
		goto L119
	} else {
		goto L131
	}
L127:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	v490 = int32(1)
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[16]))
	if v492 != v490 {
		v497 = v490
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v495 = F_rescanLatestTimeLine(m, v179, v178)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v497 = v490
	goto L126
L131:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[26]))
	if v504 == int32(0) {
		goto L119
	} else {
		goto L132
	}
L132:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	if v507 == int32(0) {
		goto L119
	} else {
		goto L133
	}
L133:
	;
	if v180&v203 == int32(0) {
		goto L121
	} else {
		goto L134
	}
L134:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[27]))
	v515 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[28]))
	v733 = v513
	v734 = v515
	goto L120
L135:
	;
	v520 = F_close(m, v517)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = int32(-1)
	goto L137
L136:
	;
	goto L137
L137:
	;
	if v181&v203 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23])) = int32(0)
	goto L140
L139:
	;
	goto L140
L140:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v530 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	if v532 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v678
	v682 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v683 = base.I64_div_u_s(int64(4294967296), v682)
	v684 = base.I64_div_u_s(v530, v683)
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+52)) = uint32(v684)
	v687 = v530 - v683*v684
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+56)) = uint32(v687)
	v690 = v34 + int32(192)
	v695 = F_pg_snprintf(m, v690, int32(1024), int32(_a_F_XLogPageRead_9), v34+int32(48))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L11
	} else {
		goto L178
	}
L142:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	v537 = F_readTimeLineHistory(m, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L145
	}
L143:
	;
	v541 = v532
	goto L144
L144:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v542 <= int32(0) {
		goto L141
	} else {
		goto L147
	}
L145:
	;
	if v537 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v541 = v537
	goto L144
L147:
	;
	if v528 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v548 = v528
	goto L150
L149:
	;
	v548 = int32(0)
	goto L150
L150:
	;
	v557 = int32(0)
	goto L151
L151:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v583+v557<<(uint(int32(2))%32))))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	if base.Ui32(v588) < base.Ui32(v590) {
		goto L141
	} else {
		goto L153
	}
L152:
	;
	goto L141
L153:
	;
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v587)+8))
	if v592 != int64(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v643 = v557 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v643 < v644 {
		v557 = v643
		goto L151
	} else {
		goto L177
	}
L155:
	;
	v596 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v597 = base.I64_div_u_s(v592, v596)
	if base.Ui64(v530) < base.Ui64(v597) {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if base.Ui32(int32(1)) < base.Ui32(v548) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L157
L159:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	if v629 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L160:
	;
	if v548&int32(1) != 0 {
		goto L154
	} else {
		goto L170
	}
L161:
	;
	v601 = int32(1)
	v603 = F_XLogFileRead(m, v530, v588, v601, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	if v603 == int32(-1) {
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v609 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L11
	} else {
		goto L164
	}
L164:
	;
	if v609 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_10), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L11
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v627 = v603
	goto L159
L168:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(_a_F_XLogPageRead_11), int32(_a_F_XLogPageRead_12))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L11
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v623 = F_XLogFileRead(m, v530, v588, int32(2), int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L171
	}
L171:
	;
	if v623 == int32(-1) {
		goto L154
	} else {
		goto L172
	}
L172:
	;
	v627 = v623
	goto L159
L173:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24])) = v541
	goto L175
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v627
	v636 = int32(_a_F_XLogPageRead_1)
	v638 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v638 == int32(3) {
		v1245 = v636
		goto L118
	} else {
		goto L176
	}
L176:
	;
	v1425 = v636
	goto L31
L177:
	;
	goto L152
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30])) = int32(44)
	v702 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L11
	} else {
		goto L179
	}
L179:
	;
	if v702 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L11
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = int32(-1)
	v721 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v721)
	v1341 = v229
	goto L116
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v690
	F_errmsg(m, int32(_a_F_XLogPageRead_13), v34+int32(32))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L11
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(_a_F_XLogPageRead_14), int32(_a_F_XLogPageRead_12))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	if base.Ui32(v725) < base.Ui32(v728) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v731 = v728
	goto L189
L188:
	;
	v731 = int32(0)
	goto L189
L189:
	;
	if v731 != 0 {
		goto L115
	} else {
		goto L190
	}
L190:
	;
	v733 = v725
	v734 = v124
	goto L120
L191:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[18]))
	v746 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v745)+320)) = uint8(v746)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	F_LWLockRelease(m, v749+int32(1152))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[26]))
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[31]))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[32])))
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[33]))
	v762 = F_time(m)
	mBase = m.M
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v767 = base.AtomicRmwXchg32(m, v761, int32(1456), int32(1))
	if v767 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	F_s_lock(m, v761+int32(1456), int32(_a_F_XLogPageRead_15), int32(263), int32(_a_F_XLogPageRead_16))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L11
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v778 = v761 + int32(104)
	if v755 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L195
L197:
	;
	if v757 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L198:
	;
	goto L204
L199:
	;
	goto L200
L200:
	;
	v898 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v778))) = uint8(v898)
	goto L197
L201:
	;
	goto L197
L202:
	;
	v895 = F_strlen(m, v884)
	mBase = m.M
	goto L201
L204:
	;
	goto L205
L205:
	;
	v785 = int32(1023)
	if (v778^v755)&int32(3) != 0 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v888 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v885))) = uint8(v888)
	goto L202
L207:
	;
	v869 = v864
	v870 = v865
	v871 = v866
	goto L228
L208:
	;
	if v859 == int32(0) {
		v884 = v857
		v885 = v858
		goto L206
	} else {
		goto L227
	}
L209:
	;
	v857 = v755
	v858 = v778
	v859 = v785
	goto L208
L210:
	;
	goto L211
L211:
	;
	v789 = int32(0)
	if base.B2i32(v755&int32(3) == v789)|int32(0) == v789 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	if v825 == int32(0) {
		v884 = v822
		v885 = v823
		goto L206
	} else {
		goto L221
	}
L213:
	;
	v801 = v755
	v802 = v778
	v803 = v785
	goto L216
L214:
	;
	goto L215
L215:
	;
	v822 = v755
	v823 = v778
	v824 = v785
	v825 = int32(1)
	goto L212
L216:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	*(*uint8)(unsafe.Add(mBase, uint32(v802))) = uint8(v805)
	if v805 == int32(0) {
		v864 = v801
		v865 = v802
		v866 = v803
		goto L207
	} else {
		goto L218
	}
L217:
	;
	v822 = v816
	v823 = v810
	v824 = v812
	v825 = v814
	goto L212
L218:
	;
	v809 = int32(1)
	v810 = v802 + v809
	v812 = v803 - v809
	v813 = int32(0)
	v814 = base.B2i32(v812 != v813)
	v816 = v801 + v809
	if v816&int32(3) == v813 {
		v822 = v816
		v823 = v810
		v824 = v812
		v825 = v814
		goto L212
	} else {
		goto L219
	}
L219:
	;
	if v812 != 0 {
		v801 = v816
		v802 = v810
		v803 = v812
		goto L216
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822))))
	if base.B2i32(v828 == int32(0))|base.B2i32(base.Ui32(v824) < base.Ui32(int32(4))) != 0 {
		v857 = v822
		v858 = v823
		v859 = v824
		goto L208
	} else {
		goto L222
	}
L222:
	;
	v835 = v822
	v836 = v823
	v837 = v824
	goto L223
L223:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	v843 = int32(-2139062144)
	if (int32(16843008)-v840|v840)&v843 != v843 {
		v864 = v835
		v865 = v836
		v866 = v837
		goto L207
	} else {
		goto L225
	}
L224:
	;
	v857 = v851
	v858 = v849
	v859 = v853
	goto L208
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v836))) = v840
	v848 = int32(4)
	v849 = v836 + v848
	v851 = v835 + v848
	v853 = v837 - v848
	if base.Ui32(int32(3)) < base.Ui32(v853) {
		v835 = v851
		v836 = v849
		v837 = v853
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v864 = v857
	v865 = v858
	v866 = v859
	goto L207
L228:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869))))
	*(*uint8)(unsafe.Add(mBase, uint32(v870))) = uint8(v873)
	if v873 == int32(0) {
		v884 = v869
		v885 = v870
		goto L206
	} else {
		goto L230
	}
L229:
	;
	v884 = v880
	v885 = v878
	goto L206
L230:
	;
	v877 = int32(1)
	v878 = v870 + v877
	v880 = v869 + v877
	v882 = v871 - v877
	if v882 != 0 {
		v869 = v880
		v870 = v878
		v871 = v882
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1031 = v734 & base.I64_extend_i32_s(int32(0)-v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v761)+1452)) = uint8(v1030)
	*(*int64)(unsafe.Add(mBase, uint32(v761)+24)) = v762
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v761)+8))
	if v1036 != 0 {
		goto L267
	} else {
		goto L268
	}
L233:
	;
	v1028 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v761)+1388)) = uint8(v1028)
	v1030 = v759
	goto L232
L234:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757))))
	if v903 == int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v907 = v761 + int32(1388)
	goto L239
L236:
	;
	v1030 = int32(0)
	goto L232
L237:
	;
	v1024 = F_strlen(m, v1013)
	mBase = m.M
	goto L236
L239:
	;
	goto L240
L240:
	;
	v914 = int32(63)
	if (v907^v757)&int32(3) != 0 {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v1017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1014))) = uint8(v1017)
	goto L237
L242:
	;
	v998 = v993
	v999 = v994
	v1000 = v995
	goto L263
L243:
	;
	if v988 == int32(0) {
		v1013 = v986
		v1014 = v987
		goto L241
	} else {
		goto L262
	}
L244:
	;
	v986 = v757
	v987 = v907
	v988 = v914
	goto L243
L245:
	;
	goto L246
L246:
	;
	v918 = int32(0)
	if base.B2i32(v757&int32(3) == v918)|int32(0) == v918 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	if v954 == int32(0) {
		v1013 = v951
		v1014 = v952
		goto L241
	} else {
		goto L256
	}
L248:
	;
	v930 = v757
	v931 = v907
	v932 = v914
	goto L251
L249:
	;
	goto L250
L250:
	;
	v951 = v757
	v952 = v907
	v953 = v914
	v954 = int32(1)
	goto L247
L251:
	;
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930))))
	*(*uint8)(unsafe.Add(mBase, uint32(v931))) = uint8(v934)
	if v934 == int32(0) {
		v993 = v930
		v994 = v931
		v995 = v932
		goto L242
	} else {
		goto L253
	}
L252:
	;
	v951 = v945
	v952 = v939
	v953 = v941
	v954 = v943
	goto L247
L253:
	;
	v938 = int32(1)
	v939 = v931 + v938
	v941 = v932 - v938
	v942 = int32(0)
	v943 = base.B2i32(v941 != v942)
	v945 = v930 + v938
	if v945&int32(3) == v942 {
		v951 = v945
		v952 = v939
		v953 = v941
		v954 = v943
		goto L247
	} else {
		goto L254
	}
L254:
	;
	if v941 != 0 {
		v930 = v945
		v931 = v939
		v932 = v941
		goto L251
	} else {
		goto L255
	}
L255:
	;
	goto L252
L256:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951))))
	if base.B2i32(v957 == int32(0))|base.B2i32(base.Ui32(v953) < base.Ui32(int32(4))) != 0 {
		v986 = v951
		v987 = v952
		v988 = v953
		goto L243
	} else {
		goto L257
	}
L257:
	;
	v964 = v951
	v965 = v952
	v966 = v953
	goto L258
L258:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	v972 = int32(-2139062144)
	if (int32(16843008)-v969|v969)&v972 != v972 {
		v993 = v964
		v994 = v965
		v995 = v966
		goto L242
	} else {
		goto L260
	}
L259:
	;
	v986 = v980
	v987 = v978
	v988 = v982
	goto L243
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v965))) = v969
	v977 = int32(4)
	v978 = v965 + v977
	v980 = v964 + v977
	v982 = v966 - v977
	if base.Ui32(int32(3)) < base.Ui32(v982) {
		v964 = v980
		v965 = v978
		v966 = v982
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v993 = v986
	v994 = v987
	v995 = v988
	goto L242
L263:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1002)
	if v1002 == int32(0) {
		v1013 = v998
		v1014 = v999
		goto L241
	} else {
		goto L265
	}
L264:
	;
	v1013 = v1009
	v1014 = v1007
	goto L241
L265:
	;
	v1006 = int32(1)
	v1007 = v999 + v1006
	v1009 = v998 + v1006
	v1011 = v1000 - v1006
	if v1011 != 0 {
		v998 = v1009
		v999 = v1007
		v1000 = v1011
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	v1037 = int32(4)
	goto L269
L268:
	;
	v1037 = int32(1)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+8)) = v1037
	v1039 = *(*int64)(unsafe.Add(mBase, uint32(v761)+32))
	if v1039 != int64(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v761)+32)) = v1031
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	v1050 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v761)+1456)), uint32(v1050))
	if v1036 == v1050 {
		goto L276
	} else {
		goto L277
	}
L271:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v761)+56))
	if v1042 == v733 {
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v761)+64)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v761)+56)) = v733
	*(*int64)(unsafe.Add(mBase, uint32(v761)+48)) = v1031
	goto L270
L274:
	;
	goto L273
L275:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8])) = int64(0)
	goto L119
L276:
	;
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[5])))
	if v1057 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	goto L278
L278:
	;
	if v1049 != int32(-1) {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	goto L275
L280:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v1061+int32(28)))) = int32(1)
	v1068 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[35]))
	v1070 = F_pgmem_kill(m, v1068, int32(10))
	mBase = m.M
	goto L282
L281:
	;
	goto L282
L282:
	;
	goto L279
L283:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[36]))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)))
	v1080 = v1075 + v1049*int32(640) + int32(20)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)))
	if v1081 != 0 {
		goto L287
	} else {
		goto L288
	}
L284:
	;
	goto L285
L285:
	;
	goto L275
L286:
	;
	goto L285
L287:
	;
	goto L286
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = int32(1)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+4))
	if v1084 == int32(0) {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+12))
	if v1087 == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[37]))
	if v1091 == v1087 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1093 = m.G0
	v1095 = v1093 - int32(16)
	m.G0 = v1095
	v1098 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[38]))
	if v1098 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	goto L293
L293:
	;
	v1121 = F_pgmem_kill(m, v1087, int32(23))
	mBase = m.M
	goto L287
L294:
	;
	m.G0 = v1095 + int32(16)
	goto L286
L295:
	;
	v1101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1095)+15)) = uint8(v1101)
	goto L296
L296:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[39]))
	v1109 = F_write(m, v1105, v1095+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v1109 {
		goto L294
	} else {
		goto L298
	}
L297:
	;
	goto L294
L298:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30]))
	if v1113 == int32(27) {
		goto L296
	} else {
		goto L299
	}
L299:
	;
	goto L297
L300:
	;
	if v1135 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1140)
	v1341 = v229
	goto L116
L302:
	;
	goto L303
L303:
	;
	v1143 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(v124) < base.Ui64(v1143) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v1230 = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[40])) = v1230
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1230
	v1245 = v155
	goto L118
L305:
	;
	if v204 == int32(0) {
		goto L117
	} else {
		goto L323
	}
L306:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1203 {
		goto L304
	} else {
		goto L317
	}
L307:
	;
	v1149 = F_GetWalRcvFlushRecPtr(m, v34+int32(192), int32(_a_F_XLogPageRead_17))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L11
	} else {
		goto L308
	}
L308:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8])) = v1149
	if base.Ui64(v1149) <= base.Ui64(v124) {
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[41]))
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	if v1154 != v1156 {
		goto L305
	} else {
		goto L310
	}
L310:
	;
	v1158 = *(*int64)(unsafe.Add(mBase, uint32(v34)+192))
	if base.Ui64(v124) < base.Ui64(v1158) {
		goto L306
	} else {
		goto L311
	}
L311:
	;
	v1164 = m.G0
	v1165 = int32(16)
	v1166 = v1164 - v1165
	m.G0 = v1166
	F_gettimeofday(m, v1166)
	mBase = m.M
	v1169 = *(*int64)(unsafe.Add(mBase, uint32(v1166)))
	v1170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1166)+8)))
	m.G0 = v1166 + v1165
	v1178 = v1170 + v1169*int64(1000000) - int64(946684800000000)
	goto L312
L312:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[42])) = v1178
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1184 = base.AtomicRmwXchg32(m, v1181, int32(96), int32(1))
	if v1184 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	F_s_lock(m, v1186+int32(96), int32(_a_F_XLogPageRead_3), int32(_a_F_XLogPageRead_18), int32(_a_F_XLogPageRead_19))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L11
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	*(*int64)(unsafe.Add(mBase, uint32(v1195)+72)) = v1178
	v1197 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1195)+96)), uint32(v1197))
	goto L306
L316:
	;
	goto L315
L317:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	if v1207 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	v1213 = F_readTimeLineHistory(m, v1212)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L11
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1218 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v1220 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[41]))
	v1223 = F_XLogFileRead(m, v1218, v1220, int32(3), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L11
	} else {
		goto L322
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24])) = v1213
	goto L320
L322:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1223
	v1341 = v229
	goto L116
L323:
	;
	v1959 = int32(-2)
	goto L16
L324:
	;
	if v1267 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1270)
	v1341 = v229
	goto L116
L326:
	;
	goto L327
L327:
	;
	if v229 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_WalRcvForceReply(m)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L11
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	F_KnownAssignedTransactionIdsIdleMaintenance(m)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L11
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[43]))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+124))
	if v1285 != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1313 = F_WaitLatch(m, v1307+int32(4), int32(33), int32(-1), int32(83886090))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L11
	} else {
		goto L337
	}
L334:
	;
	v1286 = *(*int64)(unsafe.Add(mBase, uint32(v1285)+16))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+120))
	v1288 = *(*int64)(unsafe.Add(mBase, uint32(v1287)+16))
	v1291 = base.I32_wrap_i64(v1286 - v1288)
	goto L336
L335:
	;
	v1291 = int32(0)
	goto L336
L336:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+112))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+16))
	v1295 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[44]))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+64)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+56)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+60)) = v1296 + v1293
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	v1302 = *(*int64)(unsafe.Add(mBase, uint32(v1301)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1279)+16)) = v1302 - int64(-8192)
	goto L333
L337:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v1316+int32(4)))) = int32(0)
	goto L338
L338:
	;
	v1341 = int32(1)
	goto L116
L339:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L11
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L11
	} else {
		goto L343
	}
L342:
	;
	goto L341
L343:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v215 = v1362
	v229 = v1341
	goto L45
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+152)) = v725
	v1369 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+156)) = v1369
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+148)) = uint32(v4)
	v1373 = int64(base.Ui64(v4) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+144)) = uint32(v1373)
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_20), v34+int32(144))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L11
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3891), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L11
	} else {
		goto L346
	}
L346:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L347:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1390
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_2), v34+int32(16))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L11
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(4033), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L11
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v1425 = base.I32_wrap_i64(v1404)&(v1410-int32(1)) - v115
	goto L31
L351:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = int32(167772235)
	v1475 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	v1476 = int32(_a_F_XLogPageRead_1)
	v1478 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_XLogPageRead[9])))
	v1479 = F_pread(m, v1475, l4, v1476, v1478)
	mBase = m.M
	if v1479 != v1476 {
		goto L358
	} else {
		goto L359
	}
L352:
	;
	F___clock_gettime(m, int32(1), v1457)
	mBase = m.M
	v1461 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1457)+8)))
	v1462 = *(*int64)(unsafe.Add(mBase, uint32(v1457)))
	v1466 = v1461 + v1462*int64(1000000000)
	goto L354
L353:
	;
	v1466 = int64(0)
	goto L354
L354:
	;
	m.G0 = v1457 + int32(16)
	goto L351
L355:
	;
	v150 = int32(0)
	goto L27
L356:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10]))
	v1959 = v1937
	goto L16
L357:
	;
	v1914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v1914 != 0 {
		goto L431
	} else {
		goto L432
	}
L358:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30]))
	v1485 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v1485))) = int32(0)
	v1491 = int32(1)
	v1492 = base.I64_extend_i32_s(v1479)
	v1496 = m.G0
	v1498 = v1496 - int32(16)
	m.G0 = v1498
	if v1466 != int64(0) {
		goto L362
	} else {
		goto L363
	}
L359:
	;
	goto L360
L360:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v1722))) = int32(0)
	v1728 = int32(1)
	v1729 = int64(8192)
	v1733 = m.G0
	v1735 = v1733 - int32(16)
	m.G0 = v1735
	if v1466 != int64(0) {
		goto L401
	} else {
		goto L402
	}
L361:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v1615
	v1618 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v1621 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v1622 = base.I64_div_u_s(int64(4294967296), v1621)
	v1623 = base.I64_div_u_s(v1618, v1622)
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+132)) = uint32(v1623)
	v1626 = v1618 - v1623*v1622
	*(*uint32)(unsafe.Add(mBase, uint32(v34)+136)) = uint32(v1626)
	v1634 = F_pg_snprintf(m, v34+int32(192), int32(64), int32(_a_F_XLogPageRead_21), v34+int32(128))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L11
	} else {
		goto L378
	}
L362:
	;
	F___clock_gettime(m, int32(1), v1498)
	mBase = m.M
	v1504 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1498)+8)))
	v1505 = *(*int64)(unsafe.Add(mBase, uint32(v1498)))
	v1509 = v1504 + (v1505*int64(1000000000) - v1466)
	goto L365
L363:
	;
	goto L364
L364:
	;
	v1596 = int32(880)
	v1597 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46])) = v1597 + base.I64_extend_i32_u(v1491)
	v1601 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47])) = v1601 + v1492
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1491, v1492)
	mBase = m.M
	v1606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1606)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[49])) = uint8(v1606)
	m.G0 = v1498 + int32(16)
	goto L361
L365:
	;
	v1559 = int32(880)
	v1560 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50])) = v1560 + v1509
	v1564 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[51]))
	v1571 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1564))|base.B2i32(int32(1)<<(uint(v1564)%32)&int32(_a_F_XLogPageRead_22) == v1571) == v1571 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1576 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52])) = v1576 + v1509
	v1580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1580)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[53])) = uint8(v1580)
	goto L377
L376:
	;
	goto L377
L377:
	;
	goto L364
L378:
	;
	if v1479 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), v1717, int32(_a_F_XLogPageRead_23))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L11
	} else {
		goto L399
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30])) = v1483
	if v42 != int32(15) {
		v1654 = v42
		goto L383
	} else {
		goto L384
	}
L381:
	;
	goto L382
L382:
	;
	if v42 != int32(15) {
		v1689 = v42
		goto L391
	} else {
		goto L392
	}
L383:
	;
	v1656 = F_errstart(m, v1654, int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L11
	} else {
		goto L387
	}
L384:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1644 != int32(2) {
		v1654 = v42
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1649 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v124 == v1649 {
		v1654 = int32(14)
		goto L383
	} else {
		goto L386
	}
L386:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v124
	v1654 = int32(15)
	goto L383
L387:
	;
	if v1656 == int32(0) {
		goto L357
	} else {
		goto L388
	}
L388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+84)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = v40
	v1665 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+92)) = v1665
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v34 + int32(192)
	F_errmsg(m, int32(_a_F_XLogPageRead_24), v34+int32(80))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L11
	} else {
		goto L390
	}
L390:
	;
	v1717 = int32(3445)
	goto L379
L391:
	;
	v1691 = F_errstart(m, v1689, int32(0))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L11
	} else {
		goto L395
	}
L392:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1679 != int32(2) {
		v1689 = v42
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v1684 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v124 == v1684 {
		v1689 = int32(14)
		goto L391
	} else {
		goto L394
	}
L394:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v124
	v1689 = int32(15)
	goto L391
L395:
	;
	if v1691 == int32(0) {
		goto L357
	} else {
		goto L396
	}
L396:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L11
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v1479
	*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = int32(_a_F_XLogPageRead_1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+100)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v34)+104)) = v40
	v1704 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+108)) = v1704
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v34 + int32(192)
	F_errmsg(m, int32(_a_F_XLogPageRead_25), v34+int32(96))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L11
	} else {
		goto L398
	}
L398:
	;
	v1717 = int32(3452)
	goto L379
L399:
	;
	goto L357
L400:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1184)) = v1852
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v1855 != int32(1) {
		goto L356
	} else {
		goto L417
	}
L401:
	;
	F___clock_gettime(m, int32(1), v1735)
	mBase = m.M
	v1741 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1735)+8)))
	v1742 = *(*int64)(unsafe.Add(mBase, uint32(v1735)))
	v1746 = v1741 + (v1742*int64(1000000000) - v1466)
	goto L404
L402:
	;
	goto L403
L403:
	;
	v1833 = int32(880)
	v1834 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46])) = v1834 + base.I64_extend_i32_u(v1728)
	v1838 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47])) = v1838 + v1729
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1728, v1729)
	mBase = m.M
	v1843 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1843)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[49])) = uint8(v1843)
	m.G0 = v1735 + int32(16)
	goto L400
L404:
	;
	v1796 = int32(880)
	v1797 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50])) = v1797 + v1746
	v1801 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[51]))
	v1808 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1801))|base.B2i32(int32(1)<<(uint(v1801)%32)&int32(_a_F_XLogPageRead_22) == v1808) == v1808 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1813 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52])) = v1813 + v1746
	v1817 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1817)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[53])) = uint8(v1817)
	goto L416
L415:
	;
	goto L416
L416:
	;
	goto L403
L417:
	;
	v1859 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v1860 = base.I64_rem_u_s(l1, v1859)
	if v1860 != int64(0) {
		goto L356
	} else {
		goto L418
	}
L418:
	;
	v1863 = F_XLogReaderValidatePageHeader(m, l0, l1, l4)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L11
	} else {
		goto L419
	}
L419:
	;
	if v1863 != 0 {
		goto L356
	} else {
		goto L420
	}
L420:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865))))
	if v1866 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1905 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1904))) = uint8(v1905)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v1905)
	goto L357
L422:
	;
	if v42 != int32(15) {
		v1883 = v42
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1886 = F_errstart(m, v1883, int32(0))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L11
	} else {
		goto L427
	}
L424:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1872 != int32(2) {
		v1883 = v42
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v1876 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1878 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v1876 == v1878 {
		v1883 = int32(14)
		goto L423
	} else {
		goto L426
	}
L426:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v1876
	v1883 = int32(15)
	goto L423
L427:
	;
	if v1886 == int32(0) {
		goto L421
	} else {
		goto L428
	}
L428:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v1890
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_26), v34-int32(-64))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L11
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3508), int32(_a_F_XLogPageRead_23))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L11
	} else {
		goto L430
	}
L430:
	;
	goto L421
L431:
	;
	v1959 = int32(-2)
	goto L16
L432:
	;
	goto L433
L433:
	;
	v1917 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1917)
	v1920 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1920 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1923 = F_close(m, v1920)
	mBase = m.M
	goto L436
L435:
	;
	goto L436
L436:
	;
	v1924 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1924
	v1929 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v1929
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1929
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v1935 != 0 {
		goto L355
	} else {
		goto L437
	}
L437:
	;
	v1959 = v1924
	goto L16
L438:
	;
	v1943 = F_close(m, v1940)
	mBase = m.M
	goto L440
L439:
	;
	goto L440
L440:
	;
	v1944 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1944
	v1949 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v1949
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1949
	v1959 = v1944
	goto L16
}
func F_XLogPrefetchResetStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	F_gettimeofday(m, v8)
	mBase = m.M
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	m.G0 = v8 + v7
	v22 = base.AtomicRmwXchg64(m, v2, int32(0), v12+v11*int64(1000000)-int64(946684800000000))
	v23 = int32(_a_F_XLogPrefetchResetStats_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v25 = int64(0)
	v27 = base.AtomicRmwXchg64(m, v24, int32(8), v25)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v32 = base.AtomicRmwXchg64(m, v29, int32(16), v25)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v37 = base.AtomicRmwXchg64(m, v34, int32(24), v25)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v42 = base.AtomicRmwXchg64(m, v39, int32(32), v25)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v47 = base.AtomicRmwXchg64(m, v44, int32(40), v25)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPrefetchResetStats[0]))
	v52 = base.AtomicRmwXchg64(m, v49, int32(48), v25)
	return
}
func F_XLogPrefetcherBeginRead(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v6 - int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_XLogBeginRead(m, v10, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_XLogPrefetcherGetReader(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_XLogReadAhead(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int64
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int64
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v287 int64
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v390 int64
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v410 int64
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int64
	_ = v456
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v494 int64
	_ = v494
	var v495 int64
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int64
	_ = v574
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v718 int32
	_ = v718
	var v748 int32
	_ = v748
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v784 int32
	_ = v784
	var v786 int64
	_ = v786
	var v790 int64
	_ = v790
	var v796 int32
	_ = v796
	var v823 int64
	_ = v823
	var v827 int64
	_ = v827
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int64
	_ = v867
	var v870 int64
	_ = v870
	var v876 int32
	_ = v876
	var v879 int64
	_ = v879
	var v883 int64
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v953 int64
	_ = v953
	var v961 int64
	_ = v961
	var v967 int32
	_ = v967
	var v971 int64
	_ = v971
	var v978 int64
	_ = v978
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v992 int64
	_ = v992
	var v997 int64
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int64
	_ = v1009
	var v1013 int64
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1034 int32
	_ = v1034
	var v1036 int64
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1046 int64
	_ = v1046
	var v1049 int64
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int64
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1250 int32
	_ = v1250
	var v1296 int64
	_ = v1296
	var v1299 int64
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1329 int32
	_ = v1329
	var v1357 int32
	_ = v1357
	var v1361 int64
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1403 int64
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	v2 = l1
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(_a_F_XLogReadAhead_0)
	m.G0 = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v30 != 0 {
		v1451 = v3
		v1455 = v28
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1455 + int32(_a_F_XLogReadAhead_0)
	return v1451
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v32)
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v34
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v43 = v39 & int64(-8192)
	v44 = int32(_a_F_XLogReadAhead_1)
	v45 = base.I32_wrap_i64(v39)
	v47 = v45 & int32(_a_F_XLogReadAhead_2)
	if base.Ui32(v44) <= base.Ui32(v47) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v50 = v44
	goto L5
L4:
	;
	v50 = v47
	goto L5
L5:
	;
	v53 = F_ReadPageInternal(m, l0, v43, v50+int32(24))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v53 == int32(-2) {
		v1451 = v3
		v1455 = v28
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v61 = v53
	v64 = v47
	v65 = v3
	v69 = v45
	v82 = v39
	v83 = v43
	goto L12
L9:
	;
	if v1414 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L10:
	;
	v1404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1379)+1256)) = uint8(v1404)
	*(*int64)(unsafe.Add(mBase, uint32(v1379)+56)) = v1403
	*(*int64)(unsafe.Add(mBase, uint32(v1379)+48)) = v110
	v1408 = v1379
	v1414 = v1385
	v1416 = v1387
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v495
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+17)))
	if v497 != 0 {
		goto L145
	} else {
		goto L146
	}
L12:
	;
	if v61 < int32(0) {
		v1408 = l0
		v1414 = v65
		v1416 = v28
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v430 = int32(_a_F_XLogReadAhead_3)
	v431 = v109 + v130
	if base.Ui32(v430) <= base.Ui32(v431) {
		goto L135
	} else {
		goto L136
	}
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+2)))
	if v89&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v92 = int32(40)
	goto L17
L16:
	;
	v92 = int32(24)
	goto L17
L17:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v113 = int32(0)
	if base.B2i32(v89&int32(1) == v113)|base.B2i32(v92 != v109) == v113 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v109 = v92
	v110 = v82 + base.I64_extend_i32_u(v92)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v92) <= base.Ui32(v64) {
		v109 = v64
		v110 = v82
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v69
	v102 = int64(base.Ui64(v82) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+112)) = uint32(v102)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_4), v28+int32(112))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v1408 = l0
	v1414 = v65
	v1416 = v28
	goto L9
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)) = uint32(v110)
	v121 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28))) = uint32(v121)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_5), v28)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v126 = base.I32_wrap_i64(v110)
	v128 = v126 & int32(_a_F_XLogReadAhead_2)
	v129 = v88 + v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if base.Ui32(v109) <= base.Ui32(int32(_a_F_XLogReadAhead_1)) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v1408 = l0
	v1414 = v65
	v1416 = v28
	goto L9
L28:
	;
	v155 = v130 + int32(2037)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v156 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v136 = F_ValidXLogRecordHeader(m, l0, v110, v133, v129, base.B2i32(v41 == int64(0)))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(23)) < base.Ui32(v130) {
		goto L28
	} else {
		goto L34
	}
L32:
	;
	if v136 == int32(0) {
		v1408 = l0
		v1414 = v65
		v1416 = v28
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v126
	v147 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+96)) = uint32(v147)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_6), v28+int32(96))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v1408 = l0
	v1414 = v65
	v1416 = v28
	goto L9
L36:
	;
	v202 = int32(_a_F_XLogReadAhead_3) - v128
	v203 = base.B2i32(base.Ui32(v130) <= base.Ui32(v202))
	if v203 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v196 = int32(0)
	if v2 != 0 {
		v1451 = v196
		v1455 = v28
		goto L1
	} else {
		goto L52
	}
L38:
	;
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+4)) = uint8(v188)
	v198 = v186
	v200 = base.B2i32(v186 == v188)
	goto L36
L39:
	;
	if base.Ui32(v172-v171) <= base.Ui32(v155) {
		goto L37
	} else {
		goto L51
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v155) <= base.Ui32(v177-v175+v174) {
		v186 = v175
		goto L38
	} else {
		goto L49
	}
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v159 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v171) < base.Ui32(v172) {
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v163 = v159
	goto L46
L45:
	;
	v160 = int32(_a_F_XLogReadAhead_7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v160
	v163 = v160
	goto L46
L46:
	;
	v164 = F_palloc(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v164
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v169)
	v174 = v164
	v175 = v164
	v176 = v164
	goto L40
L48:
	;
	v174 = v156
	v175 = v171
	v176 = v172
	goto L40
L49:
	;
	if base.Ui32(v155) < base.Ui32(v176-v174) {
		v186 = v174
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	v186 = v171
	goto L38
L52:
	;
	v198 = v196
	v200 = int32(1)
	goto L36
L53:
	;
	if v202 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	goto L13
L56:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	base.MemoryCopy(m, v206, v207+v128, v202)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v212 = int32(_a_F_XLogReadAhead_8)
	v213 = int32(-8192)
	v216 = v130&v213 - v213
	if base.Ui32(v216) <= base.Ui32(v212) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v219 = v212
	goto L61
L60:
	;
	v219 = v216
	goto L61
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v226 = v202
	v234 = base.B2i32(base.Ui32(v109) < base.Ui32(int32(_a_F_XLogReadAhead_9)))
	v236 = v222 + v202
	v248 = v83
	goto L63
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v251
	v416 = int32(_a_F_XLogReadAhead_1)
	v417 = base.I32_wrap_i64(v251)
	v419 = v417 & int32(_a_F_XLogReadAhead_2)
	if base.Ui32(v416) <= base.Ui32(v419) {
		goto L130
	} else {
		goto L131
	}
L63:
	;
	v249 = int32(0)
	v251 = v248 - int64(-8192)
	v253 = F_ReadPageInternal(m, l0, v251, int32(24))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L6
	} else {
		goto L65
	}
L64:
	;
	v374 = int32(-1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v376 = int32(24)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v381 = m.Env.Pgmem_crc32c(m, v374, v375+v376, v378-v376)
	mBase = m.M
	v383 = m.Env.Pgmem_crc32c(m, v381, v375, int32(20))
	mBase = m.M
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v375)+20))
	if v383^v384 != v374 {
		goto L123
	} else {
		goto L124
	}
L65:
	;
	if v253 == int32(-2) {
		v1451 = v249
		v1455 = v28
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if v253 < int32(0) {
		v1379 = l0
		v1385 = v198
		v1387 = v28
		v1403 = v251
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+2)))
	if v260&int32(8) != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v260&int32(1) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v126
	v269 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+16)) = uint32(v269)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_10), v28+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v130 == v226+v276 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v1379 = l0
	v1385 = v198
	v1387 = v28
	v1403 = v251
	goto L10
L73:
	;
	v280 = v276
	goto L75
L74:
	;
	v280 = int32(0)
	goto L75
L75:
	;
	if v280 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v126
	v287 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28-int32(-64)))) = uint32(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = base.I64_extend_i32_u(v130) - base.I64_extend_i32_u(v226)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_11), v28+int32(48))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v299 = int32(_a_F_XLogReadAhead_3)
	v300 = v130 + int32(24) - v226
	if base.Ui32(v299) <= base.Ui32(v300) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v1379 = l0
	v1385 = v198
	v1387 = v28
	v1403 = v251
	goto L10
L80:
	;
	v303 = v299
	goto L82
L81:
	;
	v303 = v300
	goto L82
L82:
	;
	v304 = F_ReadPageInternal(m, l0, v251, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	if v304 == int32(-2) {
		v1451 = v249
		v1455 = v28
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v304 < int32(0) {
		v1379 = l0
		v1385 = v198
		v1387 = v28
		v1403 = v251
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+2)))
	if v312&int32(2) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v315 = int32(40)
	goto L88
L87:
	;
	v315 = int32(24)
	goto L88
L88:
	;
	if base.Ui32(v304) < base.Ui32(v315) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v317 = F_ReadPageInternal(m, l0, v251, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L92
	}
L90:
	;
	v319 = v304
	goto L91
L91:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v323 = int32(_a_F_XLogReadAhead_3) - v315
	if base.Ui32(v321) < base.Ui32(v323) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v319 = v317
	goto L91
L93:
	;
	v325 = v321
	goto L95
L94:
	;
	v325 = v323
	goto L95
L95:
	;
	v326 = v325 + v315
	if base.Ui32(v319) < base.Ui32(v326) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v328 = F_ReadPageInternal(m, l0, v251, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v325 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	base.MemoryCopy(m, v236, v315+v320, v325)
	goto L102
L101:
	;
	goto L102
L102:
	;
	if v234&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v336 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v340 = F_ValidXLogRecordHeader(m, l0, v110, v336, v337, base.B2i32(v41 == int64(0)))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v344 = v226 + v325
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1248))
	if base.Ui32(v130) <= base.Ui32(v345) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	if v340 == int32(0) {
		v1379 = l0
		v1385 = v198
		v1387 = v28
		v1403 = v251
		goto L10
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v371 = v325 + v236
	goto L110
L109:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v349 = int32(0)
	v350 = base.B2i32(v344 == v349)
	if v350 == v349 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if base.Ui32(v344) < base.Ui32(v130) {
		v226 = v344
		v234 = int32(1)
		v236 = v371
		v248 = v251
		goto L63
	} else {
		goto L122
	}
L111:
	;
	base.MemoryCopy(m, v28+int32(128), v348, v344)
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v348 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_pfree(m, v348)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v358 = F_palloc(m, v219)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L6
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1248)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1244)) = v358
	if v350 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	base.MemoryCopy(m, v358, v28+int32(128), v344)
	goto L121
L120:
	;
	goto L121
L121:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v371 = v367 + v344
	goto L110
L122:
	;
	goto L64
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v126
	v390 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+32)) = uint32(v390)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_12), v28+int32(32))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v110
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v398&int32(2) != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v1379 = l0
	v1385 = v198
	v1387 = v28
	v1403 = v251
	goto L10
L127:
	;
	v410 = int64(40)
	goto L129
L128:
	;
	v410 = int64(24)
	goto L129
L129:
	;
	v473 = v375
	v494 = v251
	v495 = base.I64_extend_i32_u((v400+int32(7))&int32(-8)) + (v410 + v251)
	goto L11
L130:
	;
	v422 = v416
	goto L132
L131:
	;
	v422 = v419
	goto L132
L132:
	;
	v425 = F_ReadPageInternal(m, l0, v251, v422+int32(24))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	if v425 != int32(-2) {
		v61 = v425
		v64 = v419
		v65 = v198
		v69 = v417
		v82 = v251
		v83 = v251
		goto L12
	} else {
		goto L134
	}
L134:
	;
	v1451 = v249
	v1455 = v28
	goto L1
L135:
	;
	v434 = v430
	goto L137
L136:
	;
	v434 = v431
	goto L137
L137:
	;
	v435 = F_ReadPageInternal(m, l0, v83, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	if v435 == int32(-2) {
		v1451 = int32(0)
		v1455 = v28
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v435 < int32(0) {
		v1408 = l0
		v1414 = v198
		v1416 = v28
		goto L9
	} else {
		goto L140
	}
L140:
	;
	v441 = int32(-1)
	v442 = int32(24)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v447 = m.Env.Pgmem_crc32c(m, v441, v129+v442, v444-v442)
	mBase = m.M
	v449 = m.Env.Pgmem_crc32c(m, v447, v129, int32(20))
	mBase = m.M
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v129)+20))
	if v449^v450 != v441 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v126
	v456 = int64(base.Ui64(v110) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v28)+80)) = uint32(v456)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_12), v28+int32(80))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v110
	v473 = v129
	v494 = v83
	v495 = v110 + base.I64_extend_i32_u((v130+int32(7))&int32(-8))
	goto L11
L144:
	;
	v1408 = l0
	v1414 = v198
	v1416 = v28
	goto L9
L145:
	;
	if v200 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+16)))
	if v498&int32(240) != int32(64) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = (v495 + base.I64_extend_i32_s(v503-int32(1))) & base.I64_extend_i32_s(int32(0)-v503)
	goto L145
L148:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v514 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L149:
	;
	v564 = v198
	goto L150
L150:
	;
	v568 = int32(0)
	v570 = m.G0
	v572 = v570 - int32(176)
	m.G0 = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v473)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v564)+48)) = v574
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v473)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v564)+40)) = v576
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v473)))
	*(*int64)(unsafe.Add(mBase, uint32(v564)+32)) = v578
	*(*int64)(unsafe.Add(mBase, uint32(v564)+16)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v564)+68)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+60)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+56)) = uint16(v568)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+8)) = v568
	v590 = v564 + int32(76)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v593 = v591 - int32(24)
	if v593 == v568 {
		v1250 = v590
		goto L172
	} else {
		goto L173
	}
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v558)+4)) = uint8(v556)
	v564 = v558
	goto L150
L152:
	;
	v553 = F_palloc(m, v155)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L168
	}
L153:
	;
	if base.Ui32(v155) < base.Ui32(v530-v529) {
		v556 = int32(0)
		v558 = v529
		goto L151
	} else {
		goto L167
	}
L154:
	;
	v537 = int32(0)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v155) <= base.Ui32(v538-v535+v536) {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v517 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v529) < base.Ui32(v530) {
		goto L153
	} else {
		goto L162
	}
L158:
	;
	v521 = v517
	goto L160
L159:
	;
	v518 = int32(_a_F_XLogReadAhead_7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v518
	v521 = v518
	goto L160
L160:
	;
	v522 = F_palloc(m, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v522
	v527 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v527)
	v533 = v522
	v535 = v522
	v536 = v522
	goto L154
L162:
	;
	v533 = v530
	v535 = v529
	v536 = v514
	goto L154
L163:
	;
	v556 = v537
	v558 = v535
	goto L151
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v533-v536) <= base.Ui32(v155) {
		goto L152
	} else {
		goto L166
	}
L166:
	;
	v556 = v537
	v558 = v536
	goto L151
L167:
	;
	goto L152
L168:
	;
	v556 = int32(1)
	v558 = v553
	goto L151
L169:
	;
	m.G0 = v572 + int32(176)
	if v1357 != 0 {
		goto L283
	} else {
		goto L284
	}
L170:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)))) = v1329
	v1357 = int32(0)
	goto L169
L171:
	;
	v1296 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+4)) = uint32(v1296)
	v1299 = int64(base.Ui64(v1296) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572))) = uint32(v1299)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_13), v572)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L6
	} else {
		goto L282
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = (v1250 - v564 + int32(7)) & int32(-8)
	v1357 = int32(1)
	goto L169
L173:
	;
	v600 = int32(-1)
	v604 = v473 + int32(24)
	v606 = v593
	v609 = v568
	v611 = v568
	goto L175
L174:
	;
	if v1106 != v1111 {
		goto L171
	} else {
		goto L259
	}
L175:
	;
	v625 = v606 - int32(1)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
	switch v626 - int32(252) {
	case 0:
		goto L179
	case 1:
		goto L180
	case 2:
		goto L181
	case 3:
		goto L182
	default:
		goto L178
	}
L176:
	;
	v1100 = v1072
	v1102 = int32(0)
	v1104 = v1096
	v1106 = v1078
	v1111 = v1083
	goto L174
L177:
	;
	if base.Ui32(v1083) < base.Ui32(v1078) {
		v600 = v1072
		v604 = v1096
		v606 = v1078
		v609 = v1081
		v611 = v1083
		goto L175
	} else {
		goto L258
	}
L178:
	;
	if base.Ui32(v626) <= base.Ui32(int32(32)) {
		goto L188
	} else {
		goto L189
	}
L179:
	;
	if base.Ui32(v606) < base.Ui32(int32(5)) {
		goto L171
	} else {
		goto L186
	}
L180:
	;
	if base.Ui32(v606) < base.Ui32(int32(3)) {
		goto L171
	} else {
		goto L185
	}
L181:
	;
	if base.Ui32(v606) < base.Ui32(int32(5)) {
		goto L171
	} else {
		goto L184
	}
L182:
	;
	if v625 == int32(0) {
		goto L171
	} else {
		goto L183
	}
L183:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v564)+68)) = v631
	v633 = int32(2)
	v1100 = v600
	v1102 = v631
	v1104 = v604 + v633
	v1106 = v606 - v633
	v1111 = v631 + v611
	goto L174
L184:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v604)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v564)+68)) = v640
	v642 = int32(5)
	v1100 = v600
	v1102 = v640
	v1104 = v604 + v642
	v1106 = v606 - v642
	v1111 = v640 + v611
	goto L174
L185:
	;
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+1)))
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+56)) = uint16(v649)
	v651 = int32(3)
	v1072 = v600
	v1078 = v606 - v651
	v1081 = v609
	v1083 = v611
	v1096 = v604 + v651
	goto L177
L186:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v604)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v564)+60)) = v657
	v659 = int32(5)
	v1072 = v600
	v1078 = v606 - v659
	v1081 = v609
	v1083 = v611
	v1096 = v604 + v659
	goto L177
L187:
	;
	if v626 <= v600 {
		goto L203
	} else {
		goto L204
	}
L188:
	;
	v666 = v600 + int32(1)
	if v626 <= v666 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v786 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+168)) = uint32(v786)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+160)) = v626
	v790 = int64(base.Ui64(v786) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+164)) = uint32(v790)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_14), v572+int32(160))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L6
	} else {
		goto L202
	}
L191:
	;
	v673 = (v600 ^ int32(-1) + v626) & int32(7)
	if v673 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v677 = int32(0)
	v683 = v666
	goto L195
L193:
	;
	v718 = v666
	goto L194
L194:
	;
	if base.Ui32(v626-v600-int32(2)) <= base.Ui32(int32(6)) {
		goto L187
	} else {
		goto L198
	}
L195:
	;
	v702 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v590+v683*int32(52)))) = uint8(v702)
	v704 = int32(1)
	v705 = v683 + v704
	v707 = v677 + v704
	if v707 != v673 {
		v677 = v707
		v683 = v705
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v718 = v705
	goto L194
L197:
	;
	goto L196
L198:
	;
	v748 = v718
	goto L199
L199:
	;
	v766 = v590 + v748*int32(52)
	v767 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+364)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+312)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+260)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+208)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+156)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+104)) = uint8(v767)
	*(*uint8)(unsafe.Add(mBase, uint32(v766)+52)) = uint8(v767)
	v784 = v748 + int32(8)
	if v784 != v626 {
		v748 = v784
		goto L199
	} else {
		goto L201
	}
L200:
	;
	goto L187
L201:
	;
	goto L200
L202:
	;
	goto L170
L203:
	;
	v823 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+152)) = uint32(v823)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+144)) = v626
	v827 = int64(base.Ui64(v823) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+148)) = uint32(v827)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_15), v572+int32(144))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L6
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564)+72)) = v626
	v837 = v590 + v626*int32(52)
	v838 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+30)) = uint8(v838)
	v840 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v837))) = uint8(v840)
	if v625 == v838 {
		goto L171
	} else {
		goto L207
	}
L206:
	;
	goto L170
L207:
	;
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+28)) = uint8(v844)
	*(*int32)(unsafe.Add(mBase, uint32(v837)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v837)+16)) = v844 & int32(15)
	v853 = int32(1)
	v854 = int32(base.Ui32(v844)>>(uint(int32(5))%32)) & v853
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+43)) = uint8(v854)
	v859 = int32(base.Ui32(v844)>>(uint(int32(4))%32)) & v853
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+29)) = uint8(v859)
	v862 = v606 & int32(-2)
	if v862 == int32(2) {
		goto L171
	} else {
		goto L208
	}
L208:
	;
	v865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+48)) = uint16(v865)
	if v854 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v890 = int32(4)
	v891 = v606 - v890
	v893 = v604 + v890
	v894 = v611 + v865
	if v859 == int32(0) {
		v1020 = v891
		v1022 = v893
		v1023 = v894
		goto L217
	} else {
		goto L218
	}
L210:
	;
	if v865 != 0 {
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	if v865 == int32(0) {
		goto L209
	} else {
		goto L215
	}
L213:
	;
	v867 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+20)) = uint32(v867)
	v870 = int64(base.Ui64(v867) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+16)) = uint32(v870)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_16), v572+int32(16))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L6
	} else {
		goto L214
	}
L214:
	;
	goto L170
L215:
	;
	v879 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+136)) = uint32(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+128)) = v865
	v883 = int64(base.Ui64(v879) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+132)) = uint32(v883)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_17), v572+int32(128))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	goto L170
L217:
	;
	if int32(0) <= base.I32_extend8_s(v844) {
		goto L249
	} else {
		goto L250
	}
L218:
	;
	if base.Ui32(v891) < base.Ui32(int32(2)) {
		goto L171
	} else {
		goto L219
	}
L219:
	;
	v899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v893))))
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+40)) = uint16(v899)
	if v862 == int32(6) {
		goto L171
	} else {
		goto L220
	}
L220:
	;
	v903 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+36)) = uint16(v903)
	if v606 == int32(8) {
		goto L171
	} else {
		goto L221
	}
L221:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+42)) = uint8(v907)
	v909 = int32(1)
	v912 = int32(base.Ui32(v907)>>(uint(v909)%32)) & v909
	*(*uint8)(unsafe.Add(mBase, uint32(v837)+30)) = uint8(v912)
	v914 = int32(9)
	v915 = v606 - v914
	v919 = v907 & int32(28)
	if v919 != 0 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v907&int32(1) != 0 {
		goto L232
	} else {
		goto L233
	}
L223:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+38)) = uint16(v933)
	v935 = v915
	v936 = v604 + v914
	v937 = v933
	goto L222
L224:
	;
	if v907&int32(1) != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	goto L226
L226:
	;
	v933 = int32(_a_F_XLogReadAhead_3) - v899
	goto L223
L227:
	;
	if base.Ui32(v915) < base.Ui32(int32(2)) {
		goto L171
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v933 = int32(0)
	goto L223
L230:
	;
	v924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+9)))
	*(*uint16)(unsafe.Add(mBase, uint32(v837)+38)) = uint16(v924)
	v926 = int32(11)
	v935 = v606 - v926
	v936 = v604 + v926
	v937 = v924
	goto L222
L231:
	;
	if v907&int32(29)|v1004 != 0 {
		v1020 = v935
		v1022 = v936
		v1023 = v894 + v899
		goto L217
	} else {
		goto L246
	}
L232:
	;
	v940 = int32(0)
	if base.B2i32(v903 == v940)|base.B2i32(v937&int32(_a_F_XLogReadAhead_18) == v940) == v940 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	if (v903|v937)&int32(_a_F_XLogReadAhead_18) != 0 {
		goto L240
	} else {
		goto L241
	}
L235:
	;
	if v899 != int32(_a_F_XLogReadAhead_3) {
		v1004 = int32(0)
		goto L231
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v953 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+112)) = uint32(v953)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+104)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v572)+100)) = v937 & int32(_a_F_XLogReadAhead_18)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+96)) = v903
	v961 = int64(base.Ui64(v953) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+108)) = uint32(v961)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_19), v572+int32(96))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L6
	} else {
		goto L239
	}
L238:
	;
	goto L237
L239:
	;
	goto L170
L240:
	;
	v971 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+92)) = uint32(v971)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+84)) = v937 & int32(_a_F_XLogReadAhead_18)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v903
	v978 = int64(base.Ui64(v971) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+88)) = uint32(v978)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_20), v572+int32(80))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L6
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v985 = int32(_a_F_XLogReadAhead_3)
	if base.B2i32(v919 == int32(0))|base.B2i32(v899 != v985) != 0 {
		v1004 = base.B2i32(v899 == v985)
		goto L231
	} else {
		goto L244
	}
L243:
	;
	goto L170
L244:
	;
	v992 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+40)) = uint32(v992)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+32)) = int32(_a_F_XLogReadAhead_3)
	v997 = int64(base.Ui64(v992) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+36)) = uint32(v997)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_21), v572+int32(32))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L6
	} else {
		goto L245
	}
L245:
	;
	goto L170
L246:
	;
	v1009 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+72)) = uint32(v1009)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v865
	v1013 = int64(base.Ui64(v1009) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+68)) = uint32(v1013)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_22), v572-int32(-64))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L6
	} else {
		goto L247
	}
L247:
	;
	goto L170
L248:
	;
	if base.Ui32(v1060) < base.Ui32(int32(4)) {
		goto L171
	} else {
		goto L257
	}
L249:
	;
	if base.Ui32(v1020) < base.Ui32(int32(12)) {
		goto L171
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	if v609 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v837)+12)) = v1034
	v1036 = *(*int64)(unsafe.Add(mBase, uint32(v1022)))
	*(*int64)(unsafe.Add(mBase, uint32(v837)+4)) = v1036
	v1038 = int32(12)
	v1060 = v1020 - v1038
	v1061 = v1022 + v1038
	v1062 = v837 + int32(4)
	goto L248
L253:
	;
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+52)) = uint32(v1046)
	v1049 = int64(base.Ui64(v1046) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v572)+48)) = uint32(v1049)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_23), v572+int32(48))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L6
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v609)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v837)+12)) = v1056
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v609)))
	*(*int64)(unsafe.Add(mBase, uint32(v837)+4)) = v1058
	v1060 = v1020
	v1061 = v1022
	v1062 = v609
	goto L248
L256:
	;
	goto L170
L257:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1061)))
	*(*int32)(unsafe.Add(mBase, uint32(v837)+20)) = v1065
	v1067 = int32(4)
	v1072 = v626
	v1078 = v1060 - v1067
	v1081 = v1062
	v1083 = v1023
	v1096 = v1061 + v1067
	goto L177
L258:
	;
	goto L176
L259:
	;
	v1126 = v564 + int32(76)
	v1127 = int32(52)
	v1131 = v1126 + v1100*v1127 + v1127
	v1132 = int32(0)
	if v1132 <= v1100 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1139 = int32(0)
	v1141 = v1104
	v1145 = v1132
	v1147 = v1131
	goto L263
L261:
	;
	v1207 = v1102
	v1209 = v1104
	v1215 = v1131
	goto L262
L262:
	;
	if v1207 == int32(0) {
		v1250 = v1215
		goto L172
	} else {
		goto L278
	}
L263:
	;
	v1163 = v1126 + v1139*int32(52)
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	if v1164 != int32(1) {
		v1195 = v1141
		v1196 = v1147
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v564)+68))
	v1207 = v1203
	v1209 = v1195
	v1215 = v1196
	goto L262
L265:
	;
	v1198 = v1145 + int32(1)
	v1200 = v1198 & int32(255)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v564)+72))
	if v1200 <= v1201 {
		v1139 = v1200
		v1141 = v1195
		v1145 = v1198
		v1147 = v1196
		goto L263
	} else {
		goto L277
	}
L266:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+29)))
	if v1167 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+32)) = v1147
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163)+40)))
	if v1171 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1177 = v1141
	v1178 = v1147
	goto L269
L269:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+43)))
	if v1179 != int32(1) {
		v1195 = v1177
		v1196 = v1178
		goto L265
	} else {
		goto L273
	}
L270:
	;
	base.MemoryCopy(m, v1147, v1141, v1171)
	goto L272
L271:
	;
	goto L272
L272:
	;
	v1173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163)+40)))
	v1177 = v1173 + v1141
	v1178 = v1147 + v1173
	goto L269
L273:
	;
	v1185 = (v1178 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+44)) = v1185
	v1187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163)+48)))
	if v1187 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	base.MemoryCopy(m, v1185, v1177, v1187)
	goto L276
L275:
	;
	goto L276
L276:
	;
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163)+48)))
	v1195 = v1189 + v1177
	v1196 = v1185 + v1189
	goto L265
L277:
	;
	goto L264
L278:
	;
	v1234 = (v1215 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+64)) = v1234
	if v1207 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	base.MemoryCopy(m, v1234, v1209, v1207)
	goto L281
L280:
	;
	goto L281
L281:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v564)+68))
	v1250 = v1234 + v1237
	goto L172
L282:
	;
	goto L170
L283:
	;
	v1361 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v564)+24)) = v1361
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+4)))
	if v1363 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	if base.Ui32(v130) <= base.Ui32(v202) {
		v1408 = l0
		v1414 = v564
		v1416 = v28
		goto L9
	} else {
		goto L298
	}
L286:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v1366 != v564 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	goto L288
L288:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v1374 != 0 {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v1369 = v1368
	goto L291
L290:
	;
	v1369 = v1366
	goto L291
L291:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1369 + v1370
	goto L288
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1374)+8)) = v564
	goto L294
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v564
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1377 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1451 = v564
	v1455 = v28
	goto L1
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v564
	v1451 = v564
	v1455 = v28
	goto L1
L298:
	;
	v1379 = l0
	v1385 = v564
	v1387 = v28
	v1403 = v494
	goto L10
L299:
	;
	v1440 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1408)+1192)) = v1440
	*(*int64)(unsafe.Add(mBase, uint32(v1408)+1176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1408)+132)) = v1440
	v1451 = v1440
	v1455 = v1416
	goto L1
L300:
	;
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414)+4)))
	if v1435 != int32(1) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	F_pfree(m, v1414)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L6
	} else {
		goto L302
	}
L302:
	;
	goto L299
}
func F_XLogReadBufferForRedo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_XLogReadBufferForRedoExtended(m, l0, l1, v4, v4, l2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_XLogReaderFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1168))
	if v3 != int32(-1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		m.T0[v6].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v9 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
				F_pfree(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
					if v20 != 0 {
						F_pfree(m, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
				if v12 != int32(1) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
					F_pfree(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
						if v20 != 0 {
							F_pfree(m, v20)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					F_pfree(m, v9)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
						F_pfree(m, v17)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
							if v20 != 0 {
								F_pfree(m, v20)
								mBase = m.M
								v22 = m.ExcPending
								if v22 != 0 {
									return
								} else {
									v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									F_pfree(m, v23)
									mBase = m.M
									v25 = m.ExcPending
									if v25 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v27 = m.ExcPending
										if v27 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		if v9 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
			F_pfree(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
				if v20 != 0 {
					F_pfree(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					F_pfree(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
			if v12 != int32(1) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
				F_pfree(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
					if v20 != 0 {
						F_pfree(m, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						F_pfree(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_pfree(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
					F_pfree(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
						if v20 != 0 {
							F_pfree(m, v20)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								F_pfree(m, v23)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							F_pfree(m, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogReaderValidatePageHeader(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int64
	_ = v148
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int64
	_ = v222
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	v13 = m.G0
	v15 = v13 - int32(320)
	m.G0 = v15
	v17 = base.I32_wrap_i64(l1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v21 = v17 & (v18 - int32(1))
	v22 = base.I64_extend_i32_s(v18)
	v23 = base.I64_div_u_s(l1, v22)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if v24 != int32(_a_F_XLogReaderValidatePageHeader_0) {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v27
		v30 = base.I64_div_u_s(int64(4294967296), v22)
		v31 = base.I64_div_u_s(v23, v30)
		*(*uint32)(unsafe.Add(mBase, uint32(v15)+244)) = uint32(v31)
		v34 = v23 - v30*v31
		*(*uint32)(unsafe.Add(mBase, uint32(v15)+248)) = uint32(v34)
		v37 = v15 + int32(256)
		v42 = F_pg_snprintf(m, v37, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(240))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v17
			v50 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+216)) = uint32(v50)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v46
			*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v37
			F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_2), v15+int32(208))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v237 = int32(0)
				m.G0 = v15 + int32(320)
				return v237
			}
		}
	} else {
		v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
		if base.Ui32(int32(16)) <= base.Ui32(v60) {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v63
			v66 = base.I64_div_u_s(int64(4294967296), v22)
			v67 = base.I64_div_u_s(v23, v66)
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v67)
			v70 = v23 - v66*v67
			*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v70)
			v73 = v15 + int32(256)
			v78 = F_pg_snprintf(m, v73, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(32))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v17
				v84 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v84)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v73
				F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_3), v15)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v237 = int32(0)
					m.G0 = v15 + int32(320)
					return v237
				}
			}
		} else {
			if v60&int32(2) != 0 {
				v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				if v94 == int64(0) {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
					if v18 != v108 {
						v110 = int32(0)
						F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_4), v110)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							v237 = v110
							m.G0 = v15 + int32(320)
							return v237
						}
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
						if v115 == int32(_a_F_XLogReaderValidatePageHeader_5) {
							v159 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							if l1 != v159 {
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v161
								v164 = base.I64_div_u_s(int64(4294967296), v22)
								v165 = base.I64_div_u_s(v23, v164)
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+180)) = uint32(v165)
								v168 = v23 - v164*v165
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+184)) = uint32(v168)
								v171 = v15 + int32(256)
								v176 = F_pg_snprintf(m, v171, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(176))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int32(0)
								} else {
									v178 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v21
									*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v17
									v181 = int64(32)
									v182 = int64(base.Ui64(l1) >> (uint(v181) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+156)) = uint32(v182)
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+148)) = uint32(v178)
									v186 = int64(base.Ui64(v178) >> (uint(v181) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+144)) = uint32(v186)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v171
									F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_6), v15+int32(144))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int32(0)
									} else {
										v237 = int32(0)
										m.G0 = v15 + int32(320)
										return v237
									}
								}
							} else {
								v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
								if base.Ui64(l1) <= base.Ui64(v195) {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
									v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
									v237 = int32(1)
									m.G0 = v15 + int32(320)
									return v237
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
									if base.Ui32(v198) <= base.Ui32(v197) {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
										v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
										v237 = int32(1)
										m.G0 = v15 + int32(320)
										return v237
									} else {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v200
										v203 = base.I64_div_u_s(int64(4294967296), v22)
										v204 = base.I64_div_u_s(v23, v203)
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v204)
										v207 = v23 - v203*v204
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v207)
										v210 = v15 + int32(256)
										v215 = F_pg_snprintf(m, v210, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(128))
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return int32(0)
										} else {
											v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
											*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
											v222 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v222)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v218
											*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v217
											*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v210
											F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_7), v15+int32(96))
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return int32(0)
											} else {
												v237 = int32(0)
												m.G0 = v15 + int32(320)
												return v237
											}
										}
									}
								}
							}
						} else {
							v118 = int32(0)
							F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_8), v118)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int32(0)
							} else {
								v237 = v118
								m.G0 = v15 + int32(320)
								return v237
							}
						}
					}
				} else {
					v97 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					if v97 == v94 {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
						if v18 != v108 {
							v110 = int32(0)
							F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_4), v110)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v237 = v110
								m.G0 = v15 + int32(320)
								return v237
							}
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
							if v115 == int32(_a_F_XLogReaderValidatePageHeader_5) {
								v159 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								if l1 != v159 {
									v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v161
									v164 = base.I64_div_u_s(int64(4294967296), v22)
									v165 = base.I64_div_u_s(v23, v164)
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+180)) = uint32(v165)
									v168 = v23 - v164*v165
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+184)) = uint32(v168)
									v171 = v15 + int32(256)
									v176 = F_pg_snprintf(m, v171, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(176))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int32(0)
									} else {
										v178 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v21
										*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v17
										v181 = int64(32)
										v182 = int64(base.Ui64(l1) >> (uint(v181) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+156)) = uint32(v182)
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+148)) = uint32(v178)
										v186 = int64(base.Ui64(v178) >> (uint(v181) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v15)+144)) = uint32(v186)
										*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v171
										F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_6), v15+int32(144))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int32(0)
										} else {
											v237 = int32(0)
											m.G0 = v15 + int32(320)
											return v237
										}
									}
								} else {
									v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
									if base.Ui64(l1) <= base.Ui64(v195) {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
										v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
										v237 = int32(1)
										m.G0 = v15 + int32(320)
										return v237
									} else {
										v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
										if base.Ui32(v198) <= base.Ui32(v197) {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
											v237 = int32(1)
											m.G0 = v15 + int32(320)
											return v237
										} else {
											v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
											*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v200
											v203 = base.I64_div_u_s(int64(4294967296), v22)
											v204 = base.I64_div_u_s(v23, v203)
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v204)
											v207 = v23 - v203*v204
											*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v207)
											v210 = v15 + int32(256)
											v215 = F_pg_snprintf(m, v210, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(128))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
												v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
												*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
												*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
												v222 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v222)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v218
												*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v217
												*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v210
												F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_7), v15+int32(96))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return int32(0)
												} else {
													v237 = int32(0)
													m.G0 = v15 + int32(320)
													return v237
												}
											}
										}
									}
								}
							} else {
								v118 = int32(0)
								F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_8), v118)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v237 = v118
									m.G0 = v15 + int32(320)
									return v237
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v15)+200)) = v94
						*(*int64)(unsafe.Add(mBase, uint32(v15)+192)) = v97
						F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_9), v15+int32(192))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							v237 = int32(0)
							m.G0 = v15 + int32(320)
							return v237
						}
					}
				}
			} else {
				if v21 != 0 {
					v159 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					if l1 != v159 {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v161
						v164 = base.I64_div_u_s(int64(4294967296), v22)
						v165 = base.I64_div_u_s(v23, v164)
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+180)) = uint32(v165)
						v168 = v23 - v164*v165
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+184)) = uint32(v168)
						v171 = v15 + int32(256)
						v176 = F_pg_snprintf(m, v171, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(176))
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							v178 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v21
							*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v17
							v181 = int64(32)
							v182 = int64(base.Ui64(l1) >> (uint(v181) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+156)) = uint32(v182)
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+148)) = uint32(v178)
							v186 = int64(base.Ui64(v178) >> (uint(v181) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+144)) = uint32(v186)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v171
							F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_6), v15+int32(144))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
								v237 = int32(0)
								m.G0 = v15 + int32(320)
								return v237
							}
						}
					} else {
						v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1200))
						if base.Ui64(l1) <= base.Ui64(v195) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
							v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
							v237 = int32(1)
							m.G0 = v15 + int32(320)
							return v237
						} else {
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
							if base.Ui32(v198) <= base.Ui32(v197) {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1200)) = l1
								v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1208)) = v234
								v237 = int32(1)
								m.G0 = v15 + int32(320)
								return v237
							} else {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v200
								v203 = base.I64_div_u_s(int64(4294967296), v22)
								v204 = base.I64_div_u_s(v23, v203)
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+132)) = uint32(v204)
								v207 = v23 - v203*v204
								*(*uint32)(unsafe.Add(mBase, uint32(v15)+136)) = uint32(v207)
								v210 = v15 + int32(256)
								v215 = F_pg_snprintf(m, v210, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(128))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return int32(0)
								} else {
									v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1208))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v21
									*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v17
									v222 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v15)+108)) = uint32(v222)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v218
									*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v217
									*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v210
									F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_7), v15+int32(96))
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int32(0)
									} else {
										v237 = int32(0)
										m.G0 = v15 + int32(320)
										return v237
									}
								}
							}
						}
					}
				} else {
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v123
					v126 = base.I64_div_u_s(int64(4294967296), v22)
					v127 = base.I64_div_u_s(v23, v126)
					*(*uint32)(unsafe.Add(mBase, uint32(v15)+84)) = uint32(v127)
					v130 = v23 - v126*v127
					*(*uint32)(unsafe.Add(mBase, uint32(v15)+88)) = uint32(v130)
					v133 = v15 + int32(256)
					v138 = F_pg_snprintf(m, v133, int32(64), int32(_a_F_XLogReaderValidatePageHeader_1), v15+int32(80))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
						v141 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v141
						*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v17
						v148 = int64(base.Ui64(l1) >> (uint(int64(32)) % 64))
						*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v148)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v133
						F_report_invalid_record(m, l0, int32(_a_F_XLogReaderValidatePageHeader_3), v15+int32(48))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							v237 = v141
							m.G0 = v15 + int32(320)
							return v237
						}
					}
				}
			}
		}
	}
}
func F_XLogSendLogical(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v1)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v16 = F_XLogReadRecord(m, v13, v6+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		if v18 == int32(0) {
			if v16 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				F_LogicalDecodingProcessRecord(m, v22, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
					*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2])) = v30
					v34 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3]))
					if v34 != int64(0) {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+40))
						if base.Ui64(v40) < base.Ui64(v34) {
							v75 = v34
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
							if base.Ui64(v75) <= base.Ui64(v79) {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v85 == int32(0) {
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
							if v99 != 0 {
								F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[7])))
							if v44 == int32(1) {
								v48 = F_GetXLogReplayRecPtr(m, int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v73 = v48
									*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
									v75 = v73
									v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
									v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
									if base.Ui64(v75) <= base.Ui64(v79) {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v89 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									} else {
										v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
										if v85 == int32(0) {
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
											if v89 == int32(0) {
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
											}
										}
									}
									v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
									v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
									if v99 != 0 {
										F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
											*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
											v110 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
											m.G0 = v6 + int32(16)
											return
										}
									} else {
										v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
										v110 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
										m.G0 = v6 + int32(16)
										return
									}
								}
							} else {
								v52 = int32(_a_F_XLogSendLogical_2)
								v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
								v54 = int64(0)
								v57 = base.AtomicRmwCmpxchg64(m, v53, int32(280), v54, v54)
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
								v65 = base.AtomicRmwCmpxchg64(m, v61, int32(272), v54, v54)
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v65
								v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
								v73 = v72
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
								v75 = v73
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
								if base.Ui64(v75) <= base.Ui64(v79) {
									v82 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v85 == int32(0) {
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v89 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
								if v99 != 0 {
									F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
										v110 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[7])))
						if v44 == int32(1) {
							v48 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v73 = v48
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
								v75 = v73
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
								if base.Ui64(v75) <= base.Ui64(v79) {
									v82 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v85 == int32(0) {
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v89 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
								if v99 != 0 {
									F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
										v110 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v52 = int32(_a_F_XLogSendLogical_2)
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v54 = int64(0)
							v57 = base.AtomicRmwCmpxchg64(m, v53, int32(280), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v65 = base.AtomicRmwCmpxchg64(m, v61, int32(272), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v65
							v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
							v73 = v72
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
							v75 = v73
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
							if base.Ui64(v75) <= base.Ui64(v79) {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v85 == int32(0) {
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
							if v99 != 0 {
								F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v34 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3]))
				if v34 != int64(0) {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+40))
					if base.Ui64(v40) < base.Ui64(v34) {
						v75 = v34
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
						if base.Ui64(v75) <= base.Ui64(v79) {
							v82 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
							if v89 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
							}
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
							if v85 == int32(0) {
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							}
						}
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
						v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
						if v99 != 0 {
							F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
							*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
							v110 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[7])))
						if v44 == int32(1) {
							v48 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v73 = v48
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
								v75 = v73
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
								if base.Ui64(v75) <= base.Ui64(v79) {
									v82 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v85 == int32(0) {
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v89 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
								if v99 != 0 {
									F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
										v110 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v52 = int32(_a_F_XLogSendLogical_2)
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v54 = int64(0)
							v57 = base.AtomicRmwCmpxchg64(m, v53, int32(280), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v65 = base.AtomicRmwCmpxchg64(m, v61, int32(272), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v65
							v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
							v73 = v72
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
							v75 = v73
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
							if base.Ui64(v75) <= base.Ui64(v79) {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v85 == int32(0) {
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
							if v99 != 0 {
								F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[7])))
					if v44 == int32(1) {
						v48 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v73 = v48
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
							v75 = v73
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
							if base.Ui64(v75) <= base.Ui64(v79) {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v85 == int32(0) {
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v89 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
							if v99 != 0 {
								F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
									v110 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						}
					} else {
						v52 = int32(_a_F_XLogSendLogical_2)
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
						v54 = int64(0)
						v57 = base.AtomicRmwCmpxchg64(m, v53, int32(280), v54, v54)
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
						v65 = base.AtomicRmwCmpxchg64(m, v61, int32(272), v54, v54)
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v65
						v72 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
						v73 = v72
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v73
						v75 = v73
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+40))
						if base.Ui64(v75) <= base.Ui64(v79) {
							v82 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v82)
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
							if v89 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
							}
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
							if v85 == int32(0) {
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v89 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							}
						}
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
						v99 = base.AtomicRmwXchg32(m, v96, int32(76), int32(1))
						if v99 != 0 {
							F_s_lock(m, v96+int32(76), int32(_a_F_XLogSendLogical_0), int32(3496), int32(_a_F_XLogSendLogical_1))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
								v110 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v108 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
							*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
							v110 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v96)+76)), uint32(v110))
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return
			} else {
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v120
				F_errmsg_internal(m, int32(_a_F_XLogSendLogical_3), v6)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XLogSendLogical_0), int32(3445), int32(_a_F_XLogSendLogical_1))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_XLogSendPhysical(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v115 int64
	_ = v115
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int64
	_ = v269
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v278 int64
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v336 int64
	_ = v336
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v388 int64
	_ = v388
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v439 int64
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v455 int64
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int64
	_ = v490
	var v491 int64
	_ = v491
	var v494 int64
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v510 int64
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int64
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int64
	_ = v535
	var v538 int64
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int64
	_ = v554
	var v557 int64
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v605 int64
	_ = v605
	var v606 int64
	_ = v606
	var v610 int64
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int64
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int64
	_ = v645
	var v646 int64
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int64
	_ = v710
	var v711 int64
	_ = v711
	var v719 int64
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int64
	_ = v729
	var v731 int64
	_ = v731
	var v733 int64
	_ = v733
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v740 int64
	_ = v740
	var v742 int64
	_ = v742
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int64
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v798 int64
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v809 int64
	_ = v809
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[0]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[1])))
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[2]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 == int32(4) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = base.AtomicRmwXchg32(m, v26, int32(76), int32(1))
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_s_lock(m, v26+int32(76), int32(_a_F_XLogSendPhysical_0), int32(3869), int32(_a_F_XLogSendPhysical_1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(4)
	v42 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+76)), uint32(v42))
	goto L1
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	m.G0 = v19 + int32(128)
	return
L10:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v49)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])))
	if v52 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v148 = m.G0
	v149 = int32(16)
	v150 = v148 - v149
	m.G0 = v150
	F_gettimeofday(m, v150)
	mBase = m.M
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	v154 = int64(*(*int32)(unsafe.Add(mBase, uint32(v150)+8)))
	m.G0 = v150 + v149
	goto L45
L14:
	;
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	v144 = v56
	goto L13
L15:
	;
	goto L16
L16:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])))
	if v58 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v124 = F_readTimeLineHistory(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L42
	}
L18:
	;
	v64 = F_GetWalRcvFlushRecPtr(m, int32(0), v19+int32(88))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v102 = int32(_a_F_XLogSendPhysical_2)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v104 = int64(0)
	v107 = base.AtomicRmwCmpxchg64(m, v103, int32(280), v104, v104)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[8])) = v107
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v115 = base.AtomicRmwCmpxchg64(m, v111, int32(272), v104, v104)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[9])) = v115
	goto L40
L21:
	;
	v68 = F_GetXLogReplayRecPtr(m, v19+int32(32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])))
	if v74 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v84 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+316))
	v82 = base.B2i32(v80 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])) = uint8(v82)
	v84 = v82
	goto L26
L25:
	;
	v84 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+308))
	goto L30
L28:
	;
	goto L29
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[11]))
	if v94 != v71 {
		v123 = v71
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])) = uint8(v91)
	v123 = v89
	goto L17
L31:
	;
	if base.Ui64(v68) < base.Ui64(v64) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v64
	goto L34
L33:
	;
	v97 = v68
	goto L34
L34:
	;
	if v71 == v70 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v99 = v97
	goto L37
L36:
	;
	v99 = v68
	goto L37
L37:
	;
	v144 = v99
	goto L13
L38:
	;
	v144 = v122
	goto L13
L40:
	;
	goto L41
L41:
	;
	v122 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[8]))
	goto L38
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[11]))
	v130 = F_tliSwitchPoint(m, v128, v124, int32(_a_F_XLogSendPhysical_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5])) = v130
	F_list_free_deep(m, v124)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])) = uint8(v136)
	v139 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	v144 = v139
	goto L13
L45:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[12])))
	if v164 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v231 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])))
	if v233 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[14]))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	if v169 == v144 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v144
	v173 = v168 + int32(8)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[15])))
	v178 = base.I32_rem_s(v174+int32(1), int32(_a_F_XLogSendPhysical_4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[16])))
	if v178 == v179 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v183 = v173 + v178<<(uint(int32(4))%32)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[17]))) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[18]))) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[16]))) = int32(-1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[19])))
	if v191 == v178 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v195 = v173 + v178<<(uint(int32(4))%32)
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[20]))) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[21]))) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[19]))) = int32(-1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[22])))
	if v203 == v178 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v207 = v173 + v178<<(uint(int32(4))%32)
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[23]))) = v208
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[24]))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[22]))) = int32(-1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v215 = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v173+v174<<(uint(v215)%32)))) = v144
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[15])))
	*(*int64)(unsafe.Add(mBase, uint32(v168+v219<<(uint(v215)%32))+16)) = v154 + v153*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_XLogSendPhysical[15]))) = v178
	goto L46
L58:
	;
	if base.Ui64(v144) <= base.Ui64(v231) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	v237 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	if base.Ui64(v231) < base.Ui64(v237) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+1168))
	if int32(0) <= v241 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+1168))
	v245 = F_close(m, v244)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v240)+1168)) = int32(-1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v249 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[26]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	m.T0[v253].(func(*base.Module, int32, int32, int32))(m, int32(99), v249, v249)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v257)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[1])) = uint8(v257)
	v264 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	if v264 == int32(0) {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v269 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+4)) = uint32(v269)
	v272 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+12)) = uint32(v272)
	v274 = int64(32)
	v275 = int64(base.Ui64(v269) >> (uint(v274) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19))) = uint32(v275)
	v278 = int64(base.Ui64(v272) >> (uint(v274) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+8)) = uint32(v278)
	F_errmsg_internal(m, int32(_a_F_XLogSendPhysical_5), v19)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_XLogSendPhysical_0), int32(3270), int32(_a_F_XLogSendPhysical_6))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	goto L9
L70:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v290)
	goto L9
L71:
	;
	goto L72
L72:
	;
	v294 = v231 + int64(131072)
	v295 = base.B2i32(base.Ui64(v144) <= base.Ui64(v294))
	v298 = v295 & (v233 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v298)
	v300 = int32(_a_F_XLogSendPhysical_7)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v302)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[28])) = v302
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v302
	goto L73
L73:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogSendPhysical_7), int32(1))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v312 = int32(_a_F_XLogSendPhysical_8)
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v314 = int32(_a_F_XLogSendPhysical_7)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v317 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v313+v315))) = uint8(v317)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v313 + int32(1)
	F_enlargeStringInfo(m, v314, int32(8))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v327 = int32(_a_F_XLogSendPhysical_8)
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v329 = int32(_a_F_XLogSendPhysical_7)
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v332 = int64(56)
	v334 = int64(65280)
	v336 = int64(40)
	v339 = int64(16711680)
	v341 = int64(24)
	v343 = int64(4278190080)
	v345 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v328+v330))) = v231<<(uint(v332)%64) | v231&v334<<(uint(v336)%64) | (v231&v339<<(uint(v341)%64) | v231&v343<<(uint(v345)%64)) | (int64(base.Ui64(v231)>>(uint(v345)%64))&v343 | int64(base.Ui64(v231)>>(uint(v341)%64))&v339 | (int64(base.Ui64(v231)>>(uint(v336)%64))&v334 | int64(base.Ui64(v231)>>(uint(v332)%64))))
	v369 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v328 + v369
	F_enlargeStringInfo(m, v329, v369)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v376 = int32(_a_F_XLogSendPhysical_8)
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v378 = int32(_a_F_XLogSendPhysical_7)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v381 = int64(56)
	v383 = int64(65280)
	v385 = int64(40)
	v388 = int64(16711680)
	v390 = int64(24)
	v392 = int64(4278190080)
	v394 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v377+v379))) = v144<<(uint(v381)%64) | v144&v383<<(uint(v385)%64) | (v144&v388<<(uint(v390)%64) | v144&v392<<(uint(v394)%64)) | (int64(base.Ui64(v144)>>(uint(v394)%64))&v392 | int64(base.Ui64(v144)>>(uint(v390)%64))&v388 | (int64(base.Ui64(v144)>>(uint(v385)%64))&v383 | int64(base.Ui64(v144)>>(uint(v381)%64))))
	v418 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v377 + v418
	F_enlargeStringInfo(m, v378, v418)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v425 = int32(_a_F_XLogSendPhysical_8)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v427 = int32(_a_F_XLogSendPhysical_7)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v426+v428))) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v426 + int32(8)
	if base.Ui64(v144) <= base.Ui64(v294) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v439 = v144
	goto L80
L79:
	;
	v439 = v294 & int64(-8192)
	goto L80
L80:
	;
	v441 = base.I32_wrap_i64(v439 - v231)
	F_enlargeStringInfo(m, v427, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v448 = v441
	v455 = v231
	goto L82
L82:
	;
	v460 = int32(_a_F_XLogSendPhysical_8)
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v465 = v462 + v464
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+1184))
	v469 = m.G0
	v471 = v469 - int32(16)
	m.G0 = v471
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])))
	if v476 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v684 = int32(_a_F_XLogSendPhysical_8)
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v687 = v686 + v626
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v687
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v692 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v690+v687))) = uint8(v692)
	v694 = int32(_a_F_XLogSendPhysical_9)
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v692)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[31])) = v692
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32])) = v692
	goto L124
L84:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v622 = v596 + v621
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v622
	v625 = v455 + base.I64_extend_i32_u(v596)
	v626 = v448 - v596
	if v626 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L7
	} else {
		goto L106
	}
L86:
	;
	m.G0 = v471 + int32(16)
	goto L84
L87:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v474)+316))
	v482 = base.B2i32(v480 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])) = uint8(v482)
	if v480 != int32(2) {
		v596 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v474)+308))
	if v468 != v487 {
		v596 = int32(0)
		goto L86
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v490 = v455 + base.I64_extend_i32_u(v448)
	v491 = int64(0)
	v494 = base.AtomicRmwCmpxchg64(m, v474, int32(264), v491, v491)
	if base.Ui64(v494) < base.Ui64(v490) {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	if v448 == int32(0) {
		v563 = v465
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v596 = v563 - v465
	goto L86
L94:
	;
	v498 = v465
	v499 = v448
	v510 = v455
	goto L95
L95:
	;
	v517 = base.I32_wrap_i64(v510) & int32(_a_F_XLogSendPhysical_10)
	v518 = int32(_a_F_XLogSendPhysical_4) - v517
	v520 = v510 + base.I64_extend_i32_u(v518)
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+300))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+304))
	v530 = base.I64_rem_u_s(int64(base.Ui64(v510)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v526+int32(1)))
	v531 = base.I32_wrap_i64(v530)
	v533 = v531 << (uint(int32(3)) % 32)
	v535 = int64(0)
	v538 = base.AtomicRmwCmpxchg64(m, v523+v533, int32(0), v535, v535)
	if v520 != v538 {
		v563 = v498
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v563 = v559
	goto L93
L97:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+296))
	if base.Ui32(v499) < base.Ui32(v518) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v544 = v499
	goto L100
L99:
	;
	v544 = v518
	goto L100
L100:
	;
	if v544 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v498, v542+v531<<(uint(int32(13))%32)+v517, v544)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+300))
	v554 = int64(0)
	v557 = base.AtomicRmwCmpxchg64(m, v552+v533, int32(0), v554, v554)
	if v520 != v557 {
		v563 = v498
		goto L93
	} else {
		goto L104
	}
L104:
	;
	v559 = v498 + v544
	v562 = v499 - v544
	if v562 != 0 {
		v498 = v559
		v499 = v562
		v510 = v510 + base.I64_extend_i32_u(v544)
		goto L95
	} else {
		goto L105
	}
L105:
	;
	goto L96
L106:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v471)+12)) = uint32(v494)
	v605 = int64(32)
	v606 = int64(base.Ui64(v494) >> (uint(v605) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v471)+8)) = uint32(v606)
	*(*uint32)(unsafe.Add(mBase, uint32(v471)+4)) = uint32(v490)
	v610 = int64(base.Ui64(v490) >> (uint(v605) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v471))) = uint32(v610)
	F_errmsg(m, int32(_a_F_XLogSendPhysical_11), v471)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_XLogSendPhysical_12), int32(1773), int32(_a_F_XLogSendPhysical_13))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v645 = int64(*(*int32)(unsafe.Add(mBase, uint32(v644)+1160)))
	v646 = base.I64_div_u_s(v625, v645)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+1184))
	F_CheckXLogRemoved(m, v646, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L7
	} else {
		goto L114
	}
L110:
	;
	v630 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v630)+1184))
	v636 = v19 + int32(88)
	v637 = F_WALRead(m, v630, v632+v622, v625, v626, v634, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	if v637 != 0 {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	F_WALReadRaiseError(m, v636)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])))
	if v651 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	goto L83
L116:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[2]))
	v658 = base.AtomicRmwXchg32(m, v655, int32(76), int32(1))
	if v658 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_s_lock(m, v655+int32(76), int32(_a_F_XLogSendPhysical_0), int32(3367), int32(_a_F_XLogSendPhysical_6))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+16)))
	v667 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v655)+16)) = uint8(v667)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v655)+76)), uint32(v667))
	if v666 != int32(1) {
		goto L115
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)+1168))
	if v676 < int32(0) {
		goto L115
	} else {
		goto L122
	}
L122:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v675)+1168))
	v680 = F_close(m, v679)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v675)+1168)) = int32(-1)
	goto L123
L123:
	;
	v448 = v626
	v455 = v625
	goto L82
L124:
	;
	v705 = m.G0
	v706 = int32(16)
	v707 = v705 - v706
	m.G0 = v707
	F_gettimeofday(m, v707)
	mBase = m.M
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v707)))
	v711 = int64(*(*int32)(unsafe.Add(mBase, uint32(v707)+8)))
	m.G0 = v707 + v706
	v719 = v711 + v710*int64(1000000) - int64(946684800000000)
	goto L125
L125:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogSendPhysical_9), int32(8))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v724 = int32(_a_F_XLogSendPhysical_14)
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32]))
	v726 = int32(_a_F_XLogSendPhysical_9)
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	v729 = int64(56)
	v731 = int64(65280)
	v733 = int64(40)
	v736 = int64(16711680)
	v738 = int64(24)
	v740 = int64(4278190080)
	v742 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v725+v727))) = v719<<(uint(v729)%64) | v719&v731<<(uint(v733)%64) | (v719&v736<<(uint(v738)%64) | v719&v740<<(uint(v742)%64)) | (int64(base.Ui64(v719)>>(uint(v742)%64))&v740 | int64(base.Ui64(v719)>>(uint(v738)%64))&v736 | (int64(base.Ui64(v719)>>(uint(v733)%64))&v731 | int64(base.Ui64(v719)>>(uint(v729)%64))))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32])) = v725 + int32(8)
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	*(*int64)(unsafe.Add(mBase, uint32(v770)+17)) = v773
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[26]))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+20))
	m.T0[v780].(func(*base.Module, int32, int32, int32))(m, int32(100), v770, v777)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13])) = v439
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[2]))
	v789 = base.AtomicRmwXchg32(m, v786, int32(76), int32(1))
	if v789 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_s_lock(m, v786+int32(76), int32(_a_F_XLogSendPhysical_0), int32(3399), int32(_a_F_XLogSendPhysical_6))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L7
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v798 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v786)+8)) = v798
	v800 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v786)+76)), uint32(v800))
	v804 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[33])))
	if v804 != int32(1) {
		goto L9
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+20)) = uint32(v798)
	v809 = int64(base.Ui64(v798) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+16)) = uint32(v809)
	v812 = v19 + int32(32)
	v817 = F_pg_snprintf(m, v812, int32(50), int32(_a_F_XLogSendPhysical_15), v19+int32(16))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v819 = F_strlen(m, v812)
	mBase = m.M
	goto L9
}
func F_XLogSetAsyncXactLSN(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(440), int32(1))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	F_s_lock(m, v10+int32(440), int32(_a_F_XLogSetAsyncXactLSN_0), int32(2618), int32(_a_F_XLogSetAsyncXactLSN_1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+216))
	if base.Ui64(l0) <= base.Ui64(v20) {
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
	v22 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v19)+440)), uint32(v22))
	return
L7:
	;
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+216)) = l0
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+321)))
	v27 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v19)+440)), uint32(v27))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+60))
	if v60 == int32(-1) {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v31 = int64(0)
	v34 = base.AtomicRmwCmpxchg64(m, v19, int32(280), v31, v31)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[2])) = v34
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v42 = base.AtomicRmwCmpxchg64(m, v38, int32(272), v31, v31)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[3])) = v42
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[4]))
	if v45 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v48 = int64(13)
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[2]))
	if base.I32_wrap_i64(int64(base.Ui64(l0)>>(uint(v48)%64))-int64(base.Ui64(v51)>>(uint(v48)%64))) < v45 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = v63 + v60*int32(640) + int32(20)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L9
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v72 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v75 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[5]))
	if v79 == v75 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = m.G0
	v83 = v81 - int32(16)
	m.G0 = v83
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[6]))
	if v86 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v109 = F_pgmem_kill(m, v75, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v83 + int32(16)
	goto L15
L24:
	;
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+15)) = uint8(v89)
	goto L25
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[7]))
	v97 = F_write(m, v93, v83+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v97 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[8]))
	if v101 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
}
func F_XLogWalRcvSendReply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0]))
		if v14 <= int32(0) {
			m.G0 = v9 + int32(32)
			return
		} else {
			v20 = m.G0
			v21 = int32(16)
			v22 = v20 - v21
			m.G0 = v22
			F_gettimeofday(m, v22)
			mBase = m.M
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
			v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+8)))
			m.G0 = v22 + v21
			v34 = v26 + v25*int64(1000000) - int64(946684800000000)
			v36 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[1]))
			if l0 != 0 {
				v38 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
				v50 = v38
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
				v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
				if v58 <= int64(0) {
					v64 = int64(9223372036854775807)
				} else {
					v64 = v58*int64(1000000) + v34
				}
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
				v67 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v69 = int32(_a_F_XLogWalRcvSendReply_0)
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
					*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
					*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
					F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = int32(_a_F_XLogWalRcvSendReply_0)
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
						v83 = int32(_a_F_XLogWalRcvSendReply_1)
						v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						v86 = int32(114)
						*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
						v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
						v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
						F_enlargeStringInfo(m, v81, int32(8))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(_a_F_XLogWalRcvSendReply_0)
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v102 = int32(_a_F_XLogWalRcvSendReply_1)
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v105 = int64(56)
							v107 = int64(65280)
							v109 = int64(40)
							v112 = int64(16711680)
							v114 = int64(24)
							v116 = int64(4278190080)
							v118 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v144 = int32(8)
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
							v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
							F_enlargeStringInfo(m, v100, v144)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								v153 = int32(_a_F_XLogWalRcvSendReply_0)
								v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v155 = int32(_a_F_XLogWalRcvSendReply_1)
								v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v158 = int64(56)
								v160 = int64(65280)
								v162 = int64(40)
								v165 = int64(16711680)
								v167 = int64(24)
								v169 = int64(4278190080)
								v171 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
								v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v197 = int32(8)
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
								F_enlargeStringInfo(m, v153, v197)
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v206 = int32(_a_F_XLogWalRcvSendReply_1)
									v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v209 = int64(56)
									v211 = int64(65280)
									v213 = int64(40)
									v216 = int64(16711680)
									v218 = int64(24)
									v220 = int64(4278190080)
									v222 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
									v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
									v254 = m.G0
									v255 = int32(16)
									v256 = v254 - v255
									m.G0 = v256
									F_gettimeofday(m, v256)
									mBase = m.M
									v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
									v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
									m.G0 = v256 + v255
									v268 = v260 + v259*int64(1000000) - int64(946684800000000)
									F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return
									} else {
										v273 = int32(_a_F_XLogWalRcvSendReply_0)
										v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v275 = int32(_a_F_XLogWalRcvSendReply_1)
										v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v278 = int64(56)
										v280 = int64(65280)
										v282 = int64(40)
										v285 = int64(16711680)
										v287 = int64(24)
										v289 = int64(4278190080)
										v291 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
										v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
										F_enlargeStringInfo(m, v273, int32(1))
										mBase = m.M
										v323 = m.ExcPending
										if v323 != 0 {
											return
										} else {
											v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v326 = int32(_a_F_XLogWalRcvSendReply_1)
											v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
											v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
											v338 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v339 = m.ExcPending
											if v339 != 0 {
												return
											} else {
												if v338 != 0 {
													v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
													if v2 != 0 {
														v346 = int32(_a_F_XLogWalRcvSendReply_2)
													} else {
														v346 = int32(_a_F_XLogWalRcvSendReply_3)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
													v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
													v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
													v354 = int64(32)
													v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
													v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
													F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
														mBase = m.M
														v367 = m.ExcPending
														if v367 != 0 {
															return
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
															v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
															v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
															v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												} else {
													v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
													v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
													v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
													v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
													m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														m.G0 = v9 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
				v42 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
				if v42 != v36 {
					v50 = v40
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
					v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
					if v58 <= int64(0) {
						v64 = int64(9223372036854775807)
					} else {
						v64 = v58*int64(1000000) + v34
					}
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
					v67 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = int32(_a_F_XLogWalRcvSendReply_0)
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
						v71 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
						F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = int32(_a_F_XLogWalRcvSendReply_0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v83 = int32(_a_F_XLogWalRcvSendReply_1)
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v86 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
							v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
							F_enlargeStringInfo(m, v81, int32(8))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(_a_F_XLogWalRcvSendReply_0)
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v102 = int32(_a_F_XLogWalRcvSendReply_1)
								v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v105 = int64(56)
								v107 = int64(65280)
								v109 = int64(40)
								v112 = int64(16711680)
								v114 = int64(24)
								v116 = int64(4278190080)
								v118 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v144 = int32(8)
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
								v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
								F_enlargeStringInfo(m, v100, v144)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(_a_F_XLogWalRcvSendReply_0)
									v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v155 = int32(_a_F_XLogWalRcvSendReply_1)
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v158 = int64(56)
									v160 = int64(65280)
									v162 = int64(40)
									v165 = int64(16711680)
									v167 = int64(24)
									v169 = int64(4278190080)
									v171 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
									v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v197 = int32(8)
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
									F_enlargeStringInfo(m, v153, v197)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v206 = int32(_a_F_XLogWalRcvSendReply_1)
										v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v209 = int64(56)
										v211 = int64(65280)
										v213 = int64(40)
										v216 = int64(16711680)
										v218 = int64(24)
										v220 = int64(4278190080)
										v222 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
										v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
										v254 = m.G0
										v255 = int32(16)
										v256 = v254 - v255
										m.G0 = v256
										F_gettimeofday(m, v256)
										mBase = m.M
										v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
										v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
										m.G0 = v256 + v255
										v268 = v260 + v259*int64(1000000) - int64(946684800000000)
										F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return
										} else {
											v273 = int32(_a_F_XLogWalRcvSendReply_0)
											v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v275 = int32(_a_F_XLogWalRcvSendReply_1)
											v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v278 = int64(56)
											v280 = int64(65280)
											v282 = int64(40)
											v285 = int64(16711680)
											v287 = int64(24)
											v289 = int64(4278190080)
											v291 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
											v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
											F_enlargeStringInfo(m, v273, int32(1))
											mBase = m.M
											v323 = m.ExcPending
											if v323 != 0 {
												return
											} else {
												v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v326 = int32(_a_F_XLogWalRcvSendReply_1)
												v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
												v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
												v338 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v339 = m.ExcPending
												if v339 != 0 {
													return
												} else {
													if v338 != 0 {
														v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
														if v2 != 0 {
															v346 = int32(_a_F_XLogWalRcvSendReply_2)
														} else {
															v346 = int32(_a_F_XLogWalRcvSendReply_3)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
														v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
														v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
														v354 = int64(32)
														v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
														v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
														F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
														mBase = m.M
														v362 = m.ExcPending
														if v362 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
															mBase = m.M
															v367 = m.ExcPending
															if v367 != 0 {
																return
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
														v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
														v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
														v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v45 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
					if v45 != v40 {
						v50 = v40
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
						v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
						if v58 <= int64(0) {
							v64 = int64(9223372036854775807)
						} else {
							v64 = v58*int64(1000000) + v34
						}
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
						v67 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = int32(_a_F_XLogWalRcvSendReply_0)
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v71 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
							F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = int32(_a_F_XLogWalRcvSendReply_0)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v83 = int32(_a_F_XLogWalRcvSendReply_1)
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v86 = int32(114)
								*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
								v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
								v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
								F_enlargeStringInfo(m, v81, int32(8))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(_a_F_XLogWalRcvSendReply_0)
									v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v102 = int32(_a_F_XLogWalRcvSendReply_1)
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v105 = int64(56)
									v107 = int64(65280)
									v109 = int64(40)
									v112 = int64(16711680)
									v114 = int64(24)
									v116 = int64(4278190080)
									v118 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
									v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v144 = int32(8)
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
									v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
									F_enlargeStringInfo(m, v100, v144)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a_F_XLogWalRcvSendReply_0)
										v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v155 = int32(_a_F_XLogWalRcvSendReply_1)
										v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v158 = int64(56)
										v160 = int64(65280)
										v162 = int64(40)
										v165 = int64(16711680)
										v167 = int64(24)
										v169 = int64(4278190080)
										v171 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
										v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v197 = int32(8)
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
										F_enlargeStringInfo(m, v153, v197)
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v206 = int32(_a_F_XLogWalRcvSendReply_1)
											v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v209 = int64(56)
											v211 = int64(65280)
											v213 = int64(40)
											v216 = int64(16711680)
											v218 = int64(24)
											v220 = int64(4278190080)
											v222 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
											v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
											v254 = m.G0
											v255 = int32(16)
											v256 = v254 - v255
											m.G0 = v256
											F_gettimeofday(m, v256)
											mBase = m.M
											v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
											v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
											m.G0 = v256 + v255
											v268 = v260 + v259*int64(1000000) - int64(946684800000000)
											F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = int32(_a_F_XLogWalRcvSendReply_0)
												v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v275 = int32(_a_F_XLogWalRcvSendReply_1)
												v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												v278 = int64(56)
												v280 = int64(65280)
												v282 = int64(40)
												v285 = int64(16711680)
												v287 = int64(24)
												v289 = int64(4278190080)
												v291 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
												v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
												F_enlargeStringInfo(m, v273, int32(1))
												mBase = m.M
												v323 = m.ExcPending
												if v323 != 0 {
													return
												} else {
													v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
													v326 = int32(_a_F_XLogWalRcvSendReply_1)
													v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
													v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
													v338 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v339 = m.ExcPending
													if v339 != 0 {
														return
													} else {
														if v338 != 0 {
															v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
															if v2 != 0 {
																v346 = int32(_a_F_XLogWalRcvSendReply_2)
															} else {
																v346 = int32(_a_F_XLogWalRcvSendReply_3)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
															v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
															v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
															v354 = int64(32)
															v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
															v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
															F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
															mBase = m.M
															v362 = m.ExcPending
															if v362 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
																mBase = m.M
																v367 = m.ExcPending
																if v367 != 0 {
																	return
																} else {
																	v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																	v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																	v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																	v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																	m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																	mBase = m.M
																	v380 = m.ExcPending
																	if v380 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
															v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
															v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
															v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v48 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5]))
						if v34 < v48 {
							m.G0 = v9 + int32(32)
							return
						} else {
							v50 = v40
							*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
							*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
							v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
							if v58 <= int64(0) {
								v64 = int64(9223372036854775807)
							} else {
								v64 = v58*int64(1000000) + v34
							}
							*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
							v67 = F_GetXLogReplayRecPtr(m, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = int32(_a_F_XLogWalRcvSendReply_0)
								v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v71 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
								F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									v81 = int32(_a_F_XLogWalRcvSendReply_0)
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v83 = int32(_a_F_XLogWalRcvSendReply_1)
									v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v86 = int32(114)
									*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
									v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
									v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
									F_enlargeStringInfo(m, v81, int32(8))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = int32(_a_F_XLogWalRcvSendReply_0)
										v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v102 = int32(_a_F_XLogWalRcvSendReply_1)
										v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v105 = int64(56)
										v107 = int64(65280)
										v109 = int64(40)
										v112 = int64(16711680)
										v114 = int64(24)
										v116 = int64(4278190080)
										v118 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
										v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v144 = int32(8)
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
										v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
										F_enlargeStringInfo(m, v100, v144)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											v153 = int32(_a_F_XLogWalRcvSendReply_0)
											v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v155 = int32(_a_F_XLogWalRcvSendReply_1)
											v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v158 = int64(56)
											v160 = int64(65280)
											v162 = int64(40)
											v165 = int64(16711680)
											v167 = int64(24)
											v169 = int64(4278190080)
											v171 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
											v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v197 = int32(8)
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
											F_enlargeStringInfo(m, v153, v197)
											mBase = m.M
											v203 = m.ExcPending
											if v203 != 0 {
												return
											} else {
												v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v206 = int32(_a_F_XLogWalRcvSendReply_1)
												v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												v209 = int64(56)
												v211 = int64(65280)
												v213 = int64(40)
												v216 = int64(16711680)
												v218 = int64(24)
												v220 = int64(4278190080)
												v222 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
												v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
												v254 = m.G0
												v255 = int32(16)
												v256 = v254 - v255
												m.G0 = v256
												F_gettimeofday(m, v256)
												mBase = m.M
												v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
												v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
												m.G0 = v256 + v255
												v268 = v260 + v259*int64(1000000) - int64(946684800000000)
												F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
												mBase = m.M
												v272 = m.ExcPending
												if v272 != 0 {
													return
												} else {
													v273 = int32(_a_F_XLogWalRcvSendReply_0)
													v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
													v275 = int32(_a_F_XLogWalRcvSendReply_1)
													v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													v278 = int64(56)
													v280 = int64(65280)
													v282 = int64(40)
													v285 = int64(16711680)
													v287 = int64(24)
													v289 = int64(4278190080)
													v291 = int64(8)
													*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
													v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
													F_enlargeStringInfo(m, v273, int32(1))
													mBase = m.M
													v323 = m.ExcPending
													if v323 != 0 {
														return
													} else {
														v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
														v326 = int32(_a_F_XLogWalRcvSendReply_1)
														v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
														*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
														v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
														*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
														v338 = F_errstart(m, int32(13), int32(0))
														mBase = m.M
														v339 = m.ExcPending
														if v339 != 0 {
															return
														} else {
															if v338 != 0 {
																v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
																if v2 != 0 {
																	v346 = int32(_a_F_XLogWalRcvSendReply_2)
																} else {
																	v346 = int32(_a_F_XLogWalRcvSendReply_3)
																}
																*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
																v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
																v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
																v354 = int64(32)
																v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
																v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
																*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
																F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
																mBase = m.M
																v362 = m.ExcPending
																if v362 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
																	mBase = m.M
																	v367 = m.ExcPending
																	if v367 != 0 {
																		return
																	} else {
																		v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																		v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																		v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																		v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																		v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																		m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																		mBase = m.M
																		v380 = m.ExcPending
																		if v380 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	}
																}
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v20 = m.G0
		v21 = int32(16)
		v22 = v20 - v21
		m.G0 = v22
		F_gettimeofday(m, v22)
		mBase = m.M
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+8)))
		m.G0 = v22 + v21
		v34 = v26 + v25*int64(1000000) - int64(946684800000000)
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[1]))
		if l0 != 0 {
			v38 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
			v50 = v38
			*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
			*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
			v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
			if v58 <= int64(0) {
				v64 = int64(9223372036854775807)
			} else {
				v64 = v58*int64(1000000) + v34
			}
			*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
			v67 = F_GetXLogReplayRecPtr(m, int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v69 = int32(_a_F_XLogWalRcvSendReply_0)
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
				v71 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
				*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
				*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
				F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v81 = int32(_a_F_XLogWalRcvSendReply_0)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
					v83 = int32(_a_F_XLogWalRcvSendReply_1)
					v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
					v86 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
					v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
					*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
					v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
					F_enlargeStringInfo(m, v81, int32(8))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v100 = int32(_a_F_XLogWalRcvSendReply_0)
						v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
						v102 = int32(_a_F_XLogWalRcvSendReply_1)
						v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						v105 = int64(56)
						v107 = int64(65280)
						v109 = int64(40)
						v112 = int64(16711680)
						v114 = int64(24)
						v116 = int64(4278190080)
						v118 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
						v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						v144 = int32(8)
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
						v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
						F_enlargeStringInfo(m, v100, v144)
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							v153 = int32(_a_F_XLogWalRcvSendReply_0)
							v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v155 = int32(_a_F_XLogWalRcvSendReply_1)
							v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v158 = int64(56)
							v160 = int64(65280)
							v162 = int64(40)
							v165 = int64(16711680)
							v167 = int64(24)
							v169 = int64(4278190080)
							v171 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
							v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v197 = int32(8)
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
							F_enlargeStringInfo(m, v153, v197)
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return
							} else {
								v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v206 = int32(_a_F_XLogWalRcvSendReply_1)
								v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v209 = int64(56)
								v211 = int64(65280)
								v213 = int64(40)
								v216 = int64(16711680)
								v218 = int64(24)
								v220 = int64(4278190080)
								v222 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
								v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
								v254 = m.G0
								v255 = int32(16)
								v256 = v254 - v255
								m.G0 = v256
								F_gettimeofday(m, v256)
								mBase = m.M
								v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
								v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
								m.G0 = v256 + v255
								v268 = v260 + v259*int64(1000000) - int64(946684800000000)
								F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return
								} else {
									v273 = int32(_a_F_XLogWalRcvSendReply_0)
									v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v275 = int32(_a_F_XLogWalRcvSendReply_1)
									v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v278 = int64(56)
									v280 = int64(65280)
									v282 = int64(40)
									v285 = int64(16711680)
									v287 = int64(24)
									v289 = int64(4278190080)
									v291 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
									v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
									F_enlargeStringInfo(m, v273, int32(1))
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
										return
									} else {
										v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v326 = int32(_a_F_XLogWalRcvSendReply_1)
										v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
										v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
										v338 = F_errstart(m, int32(13), int32(0))
										mBase = m.M
										v339 = m.ExcPending
										if v339 != 0 {
											return
										} else {
											if v338 != 0 {
												v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
												if v2 != 0 {
													v346 = int32(_a_F_XLogWalRcvSendReply_2)
												} else {
													v346 = int32(_a_F_XLogWalRcvSendReply_3)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
												v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
												v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
												v354 = int64(32)
												v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
												v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
												F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
												mBase = m.M
												v362 = m.ExcPending
												if v362 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
													mBase = m.M
													v367 = m.ExcPending
													if v367 != 0 {
														return
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
														v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
														v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
														v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											} else {
												v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
												v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
												v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
												m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
												mBase = m.M
												v380 = m.ExcPending
												if v380 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v40 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
			v42 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
			if v42 != v36 {
				v50 = v40
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
				v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
				if v58 <= int64(0) {
					v64 = int64(9223372036854775807)
				} else {
					v64 = v58*int64(1000000) + v34
				}
				*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
				v67 = F_GetXLogReplayRecPtr(m, int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v69 = int32(_a_F_XLogWalRcvSendReply_0)
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
					*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
					*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
					F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = int32(_a_F_XLogWalRcvSendReply_0)
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
						v83 = int32(_a_F_XLogWalRcvSendReply_1)
						v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						v86 = int32(114)
						*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
						v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
						v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
						F_enlargeStringInfo(m, v81, int32(8))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(_a_F_XLogWalRcvSendReply_0)
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v102 = int32(_a_F_XLogWalRcvSendReply_1)
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v105 = int64(56)
							v107 = int64(65280)
							v109 = int64(40)
							v112 = int64(16711680)
							v114 = int64(24)
							v116 = int64(4278190080)
							v118 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v144 = int32(8)
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
							v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
							F_enlargeStringInfo(m, v100, v144)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								v153 = int32(_a_F_XLogWalRcvSendReply_0)
								v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v155 = int32(_a_F_XLogWalRcvSendReply_1)
								v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v158 = int64(56)
								v160 = int64(65280)
								v162 = int64(40)
								v165 = int64(16711680)
								v167 = int64(24)
								v169 = int64(4278190080)
								v171 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
								v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v197 = int32(8)
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
								F_enlargeStringInfo(m, v153, v197)
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v206 = int32(_a_F_XLogWalRcvSendReply_1)
									v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v209 = int64(56)
									v211 = int64(65280)
									v213 = int64(40)
									v216 = int64(16711680)
									v218 = int64(24)
									v220 = int64(4278190080)
									v222 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
									v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
									v254 = m.G0
									v255 = int32(16)
									v256 = v254 - v255
									m.G0 = v256
									F_gettimeofday(m, v256)
									mBase = m.M
									v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
									v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
									m.G0 = v256 + v255
									v268 = v260 + v259*int64(1000000) - int64(946684800000000)
									F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return
									} else {
										v273 = int32(_a_F_XLogWalRcvSendReply_0)
										v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v275 = int32(_a_F_XLogWalRcvSendReply_1)
										v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v278 = int64(56)
										v280 = int64(65280)
										v282 = int64(40)
										v285 = int64(16711680)
										v287 = int64(24)
										v289 = int64(4278190080)
										v291 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
										v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
										F_enlargeStringInfo(m, v273, int32(1))
										mBase = m.M
										v323 = m.ExcPending
										if v323 != 0 {
											return
										} else {
											v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v326 = int32(_a_F_XLogWalRcvSendReply_1)
											v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
											v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
											v338 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v339 = m.ExcPending
											if v339 != 0 {
												return
											} else {
												if v338 != 0 {
													v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
													if v2 != 0 {
														v346 = int32(_a_F_XLogWalRcvSendReply_2)
													} else {
														v346 = int32(_a_F_XLogWalRcvSendReply_3)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
													v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
													v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
													v354 = int64(32)
													v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
													v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
													F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
														mBase = m.M
														v367 = m.ExcPending
														if v367 != 0 {
															return
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
															v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
															v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
															v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												} else {
													v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
													v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
													v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
													v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
													m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														m.G0 = v9 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v45 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
				if v45 != v40 {
					v50 = v40
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
					v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
					if v58 <= int64(0) {
						v64 = int64(9223372036854775807)
					} else {
						v64 = v58*int64(1000000) + v34
					}
					*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
					v67 = F_GetXLogReplayRecPtr(m, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = int32(_a_F_XLogWalRcvSendReply_0)
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
						v71 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
						*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
						F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = int32(_a_F_XLogWalRcvSendReply_0)
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v83 = int32(_a_F_XLogWalRcvSendReply_1)
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							v86 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
							v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
							F_enlargeStringInfo(m, v81, int32(8))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(_a_F_XLogWalRcvSendReply_0)
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v102 = int32(_a_F_XLogWalRcvSendReply_1)
								v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v105 = int64(56)
								v107 = int64(65280)
								v109 = int64(40)
								v112 = int64(16711680)
								v114 = int64(24)
								v116 = int64(4278190080)
								v118 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v144 = int32(8)
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
								v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
								F_enlargeStringInfo(m, v100, v144)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(_a_F_XLogWalRcvSendReply_0)
									v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v155 = int32(_a_F_XLogWalRcvSendReply_1)
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v158 = int64(56)
									v160 = int64(65280)
									v162 = int64(40)
									v165 = int64(16711680)
									v167 = int64(24)
									v169 = int64(4278190080)
									v171 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
									v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v197 = int32(8)
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
									F_enlargeStringInfo(m, v153, v197)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v206 = int32(_a_F_XLogWalRcvSendReply_1)
										v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v209 = int64(56)
										v211 = int64(65280)
										v213 = int64(40)
										v216 = int64(16711680)
										v218 = int64(24)
										v220 = int64(4278190080)
										v222 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
										v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
										v254 = m.G0
										v255 = int32(16)
										v256 = v254 - v255
										m.G0 = v256
										F_gettimeofday(m, v256)
										mBase = m.M
										v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
										v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
										m.G0 = v256 + v255
										v268 = v260 + v259*int64(1000000) - int64(946684800000000)
										F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return
										} else {
											v273 = int32(_a_F_XLogWalRcvSendReply_0)
											v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v275 = int32(_a_F_XLogWalRcvSendReply_1)
											v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v278 = int64(56)
											v280 = int64(65280)
											v282 = int64(40)
											v285 = int64(16711680)
											v287 = int64(24)
											v289 = int64(4278190080)
											v291 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
											v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
											F_enlargeStringInfo(m, v273, int32(1))
											mBase = m.M
											v323 = m.ExcPending
											if v323 != 0 {
												return
											} else {
												v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v326 = int32(_a_F_XLogWalRcvSendReply_1)
												v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
												v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
												v338 = F_errstart(m, int32(13), int32(0))
												mBase = m.M
												v339 = m.ExcPending
												if v339 != 0 {
													return
												} else {
													if v338 != 0 {
														v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
														if v2 != 0 {
															v346 = int32(_a_F_XLogWalRcvSendReply_2)
														} else {
															v346 = int32(_a_F_XLogWalRcvSendReply_3)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
														v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
														v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
														v354 = int64(32)
														v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
														v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
														F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
														mBase = m.M
														v362 = m.ExcPending
														if v362 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
															mBase = m.M
															v367 = m.ExcPending
															if v367 != 0 {
																return
															} else {
																v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																mBase = m.M
																v380 = m.ExcPending
																if v380 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													} else {
														v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
														v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
														v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
														v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
														v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
														m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
														mBase = m.M
														v380 = m.ExcPending
														if v380 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v48 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5]))
					if v34 < v48 {
						m.G0 = v9 + int32(32)
						return
					} else {
						v50 = v40
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v36
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4])) = v50
						v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
						if v58 <= int64(0) {
							v64 = int64(9223372036854775807)
						} else {
							v64 = v58*int64(1000000) + v34
						}
						*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v64
						v67 = F_GetXLogReplayRecPtr(m, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = int32(_a_F_XLogWalRcvSendReply_0)
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
							v71 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7])) = v71
							*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v71
							F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = int32(_a_F_XLogWalRcvSendReply_0)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
								v83 = int32(_a_F_XLogWalRcvSendReply_1)
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								v86 = int32(114)
								*(*uint8)(unsafe.Add(mBase, uint32(v82+v84))) = uint8(v86)
								v90 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
								*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v90 + int32(1)
								v95 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
								F_enlargeStringInfo(m, v81, int32(8))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(_a_F_XLogWalRcvSendReply_0)
									v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
									v102 = int32(_a_F_XLogWalRcvSendReply_1)
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v105 = int64(56)
									v107 = int64(65280)
									v109 = int64(40)
									v112 = int64(16711680)
									v114 = int64(24)
									v116 = int64(4278190080)
									v118 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v101+v103))) = v95<<(uint(v105)%64) | v95&v107<<(uint(v109)%64) | (v95&v112<<(uint(v114)%64) | v95&v116<<(uint(v118)%64)) | (int64(base.Ui64(v95)>>(uint(v118)%64))&v116 | int64(base.Ui64(v95)>>(uint(v114)%64))&v112 | (int64(base.Ui64(v95)>>(uint(v109)%64))&v107 | int64(base.Ui64(v95)>>(uint(v105)%64))))
									v143 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
									v144 = int32(8)
									*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v143 + v144
									v148 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
									F_enlargeStringInfo(m, v100, v144)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a_F_XLogWalRcvSendReply_0)
										v154 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
										v155 = int32(_a_F_XLogWalRcvSendReply_1)
										v156 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v158 = int64(56)
										v160 = int64(65280)
										v162 = int64(40)
										v165 = int64(16711680)
										v167 = int64(24)
										v169 = int64(4278190080)
										v171 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v154+v156))) = v148<<(uint(v158)%64) | v148&v160<<(uint(v162)%64) | (v148&v165<<(uint(v167)%64) | v148&v169<<(uint(v171)%64)) | (int64(base.Ui64(v148)>>(uint(v171)%64))&v169 | int64(base.Ui64(v148)>>(uint(v167)%64))&v165 | (int64(base.Ui64(v148)>>(uint(v162)%64))&v160 | int64(base.Ui64(v148)>>(uint(v158)%64))))
										v196 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
										v197 = int32(8)
										*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v196 + v197
										F_enlargeStringInfo(m, v153, v197)
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return
										} else {
											v205 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
											v206 = int32(_a_F_XLogWalRcvSendReply_1)
											v207 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											v209 = int64(56)
											v211 = int64(65280)
											v213 = int64(40)
											v216 = int64(16711680)
											v218 = int64(24)
											v220 = int64(4278190080)
											v222 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v205+v207))) = v67<<(uint(v209)%64) | v67&v211<<(uint(v213)%64) | (v67&v216<<(uint(v218)%64) | v67&v220<<(uint(v222)%64)) | (int64(base.Ui64(v67)>>(uint(v222)%64))&v220 | int64(base.Ui64(v67)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v67)>>(uint(v213)%64))&v211 | int64(base.Ui64(v67)>>(uint(v209)%64))))
											v247 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
											*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v247 + int32(8)
											v254 = m.G0
											v255 = int32(16)
											v256 = v254 - v255
											m.G0 = v256
											F_gettimeofday(m, v256)
											mBase = m.M
											v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
											v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
											m.G0 = v256 + v255
											v268 = v260 + v259*int64(1000000) - int64(946684800000000)
											F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = int32(_a_F_XLogWalRcvSendReply_0)
												v274 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
												v275 = int32(_a_F_XLogWalRcvSendReply_1)
												v276 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												v278 = int64(56)
												v280 = int64(65280)
												v282 = int64(40)
												v285 = int64(16711680)
												v287 = int64(24)
												v289 = int64(4278190080)
												v291 = int64(8)
												*(*int64)(unsafe.Add(mBase, uint32(v274+v276))) = v268<<(uint(v278)%64) | v268&v280<<(uint(v282)%64) | (v268&v285<<(uint(v287)%64) | v268&v289<<(uint(v291)%64)) | (int64(base.Ui64(v268)>>(uint(v291)%64))&v289 | int64(base.Ui64(v268)>>(uint(v287)%64))&v285 | (int64(base.Ui64(v268)>>(uint(v282)%64))&v280 | int64(base.Ui64(v268)>>(uint(v278)%64))))
												v316 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
												*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v316 + int32(8)
												F_enlargeStringInfo(m, v273, int32(1))
												mBase = m.M
												v323 = m.ExcPending
												if v323 != 0 {
													return
												} else {
													v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
													v326 = int32(_a_F_XLogWalRcvSendReply_1)
													v327 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													*(*uint8)(unsafe.Add(mBase, uint32(v325+v327))) = uint8(v2)
													v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
													*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v332 + int32(1)
													v338 = F_errstart(m, int32(13), int32(0))
													mBase = m.M
													v339 = m.ExcPending
													if v339 != 0 {
														return
													} else {
														if v338 != 0 {
															v341 = int64(base.Ui64(v67) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v341)
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v67)
															if v2 != 0 {
																v346 = int32(_a_F_XLogWalRcvSendReply_2)
															} else {
																v346 = int32(_a_F_XLogWalRcvSendReply_3)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v346
															v349 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v349)
															v352 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v352)
															v354 = int64(32)
															v355 = int64(base.Ui64(v349) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v355)
															v358 = int64(base.Ui64(v352) >> (uint(v354) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v358)
															F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v9)
															mBase = m.M
															v362 = m.ExcPending
															if v362 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1145), int32(_a_F_XLogWalRcvSendReply_6))
																mBase = m.M
																v367 = m.ExcPending
																if v367 != 0 {
																	return
																} else {
																	v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
																	v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
																	v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
																	v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
																	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
																	m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
																	mBase = m.M
																	v380 = m.ExcPending
																	if v380 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														} else {
															v371 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
															v373 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
															v375 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8]))
															v377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
															v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
															m.T0[v378].(func(*base.Module, int32, int32, int32))(m, v371, v373, v375)
															mBase = m.M
															v380 = m.ExcPending
															if v380 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v311 int64
	_ = v311
	var v315 int32
	_ = v315
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v336 int64
	_ = v336
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int64
	_ = v403
	var v406 int32
	_ = v406
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v423 int64
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int64
	_ = v441
	var v443 int32
	_ = v443
	var v446 int64
	_ = v446
	var v449 int32
	_ = v449
	var v453 int64
	_ = v453
	var v455 int32
	_ = v455
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v486 int64
	_ = v486
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v499 int64
	_ = v499
	var v500 int64
	_ = v500
	var v505 int32
	_ = v505
	var v511 int64
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v554 int64
	_ = v554
	var v555 int64
	_ = v555
	var v558 int64
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int64
	_ = v571
	var v575 int64
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int64
	_ = v584
	var v586 int32
	_ = v586
	var v587 int64
	_ = v587
	var v588 int32
	_ = v588
	var v596 int64
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v611 int64
	_ = v611
	var v613 int32
	_ = v613
	var v615 int64
	_ = v615
	var v617 int64
	_ = v617
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v644 int64
	_ = v644
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v651 int64
	_ = v651
	var v652 int64
	_ = v652
	var v655 int32
	_ = v655
	var v659 int64
	_ = v659
	var v661 int32
	_ = v661
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v672 int32
	_ = v672
	var v674 int64
	_ = v674
	var v676 int64
	_ = v676
	var v677 int64
	_ = v677
	var v681 int64
	_ = v681
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	v4 = int64(0)
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v22 = int32(_a_F_XLogWrite_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v27 = base.AtomicRmwCmpxchg64(m, v23, int32(280), v4, v4)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v35 = base.AtomicRmwCmpxchg64(m, v31, int32(272), v4, v4)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v35
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+304))
	v45 = base.I64_rem_u_s(int64(base.Ui64(v35)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v41+int32(1)))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v53 = v47
	v55 = v40
	v56 = v8
	v59 = base.I32_wrap_i64(v45)
	v60 = v8
	v61 = v8
	goto L2
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L13
	} else {
		goto L117
	}
L2:
	;
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v53) <= base.Ui64(v65) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v554 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v555 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v555) <= base.Ui64(v554) {
		goto L92
	} else {
		goto L93
	}
L4:
	;
	goto L3
L5:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+300))
	v71 = int64(0)
	v74 = base.AtomicRmwCmpxchg64(m, v67+v59<<(uint(int32(3))%32), int32(0), v71, v71)
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v74) <= base.Ui64(v76) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v74
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v86 = base.I64_div_u_s(v74-int64(1), base.I64_extend_i32_s(v84))
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	if v86 != v88 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v91 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v118 = v74
	v120 = v84
	goto L9
L9:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if v122 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v103 = v86
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v103
	v109 = F_XLogFileInit(m, v103, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v97 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v101 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v102 = base.I64_div_u_s(v97-int64(1), v101)
	v103 = v102
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v109
	F_ReserveExternalFD(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v117 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v118 = v117
	v120 = v115
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	v131 = base.I64_div_u_s(v118-int64(1), base.I64_extend_i32_s(v120))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v131
	v134 = F_XLogFileOpen(m, v131, l1)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v143 = v118
	v144 = v120
	goto L19
L19:
	;
	if v56 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v134
	F_ReserveExternalFD(m)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v143 = v142
	v144 = v140
	goto L19
L22:
	;
	v145 = v61
	goto L24
L23:
	;
	v145 = v59
	goto L24
L24:
	;
	v146 = base.B2i32(base.Ui64(v74) <= base.Ui64(v80))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v153 = v60
	goto L27
L26:
	;
	v153 = (base.I32_wrap_i64(v143) + int32(-8192)) & (v144 - int32(1))
	goto L27
L27:
	;
	v155 = v56 + int32(1)
	v157 = v155 << (uint(int32(13)) % 32)
	v160 = v146 & base.B2i32(base.Ui32(v144) <= base.Ui32(v153+v157))
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	if base.Ui64(v143) < base.Ui64(v80) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v146 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L29:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+304))
	if v160|base.B2i32(v59 == v164) == int32(0) {
		v511 = v80
		v514 = v155
		v518 = v153
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+296))
	v180 = v157
	v182 = v169 + v145<<(uint(int32(13))%32)
	v185 = v153
	goto L33
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7])) = int32(0)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[8])))
	v196 = m.G0
	v198 = v196 - int32(16)
	m.G0 = v198
	if v193 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v389 = int32(0)
	if v160 == v389 {
		v511 = v80
		v514 = v389
		v518 = v388
		goto L28
	} else {
		goto L67
	}
L35:
	;
	v211 = int32(_a_F_XLogWrite_1)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = int32(167772240)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v218 = F_pwrite(m, v216, v182, v180, base.I64_extend_i32_u(v185))
	mBase = m.M
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = int32(0)
	v226 = int32(1)
	v227 = base.I64_extend_i32_s(v218)
	v231 = m.G0
	v233 = v231 - int32(16)
	m.G0 = v233
	if v207 != int64(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F___clock_gettime(m, int32(1), v198)
	mBase = m.M
	v202 = int64(*(*int32)(unsafe.Add(mBase, uint32(v198)+8)))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	v207 = v202 + v203*int64(1000000000)
	goto L38
L37:
	;
	v207 = int64(0)
	goto L38
L38:
	;
	m.G0 = v198 + int32(16)
	goto L35
L39:
	;
	if v218 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	F___clock_gettime(m, int32(1), v233)
	mBase = m.M
	v239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v233)+8)))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v233)))
	v244 = v239 + (v240*int64(1000000000) - v207)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v331 = int32(888)
	v332 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[10])) = v332 + base.I64_extend_i32_u(v226)
	v336 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[11])) = v336 + v227
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v226, v227)
	mBase = m.M
	v341 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[12])) = uint8(v341)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[13])) = uint8(v341)
	m.G0 = v233 + int32(16)
	goto L39
L43:
	;
	v294 = int32(888)
	v295 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[14])) = v295 + v244
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[15]))
	v306 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v299))|base.B2i32(int32(1)<<(uint(v299)%32)&int32(_a_F_XLogWrite_2) == v306) == v306 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v311 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[16])) = v311 + v244
	v315 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[12])) = uint8(v315)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[17])) = uint8(v315)
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L42
L56:
	;
	if v385 != 0 {
		v180 = v385
		v182 = v387
		v185 = v388
		goto L33
	} else {
		goto L66
	}
L57:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7]))
	if v352 == int32(27) {
		v385 = v180
		v387 = v182
		v388 = v185
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v385 = v180 - v218
	v387 = v218 + v182
	v388 = v218 + v185
	goto L56
L60:
	;
	v356 = v19 + int32(32)
	v358 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	F_XLogFileName(m, v356, l1, v358, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7])) = v352
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v356
	F_errmsg(m, int32(_a_F_XLogWrite_3), v19)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_XLogWrite_4), int32(2455), int32(_a_F_XLogWrite_5))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L34
L67:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v395 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	F_issue_xlog_fsync(m, v393, v395, l1)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[18])) = uint8(v399)
	v403 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v403
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[19]))
	if int32(0) < v406 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v410 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v411 = m.G0
	v413 = v411 - int32(80)
	m.G0 = v413
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = l1
	v418 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v419 = base.I64_div_u_s(int64(4294967296), v418)
	v420 = base.I64_div_u_s(v410, v419)
	*(*uint32)(unsafe.Add(mBase, uint32(v413)+4)) = uint32(v420)
	v423 = v410 - v419*v420
	*(*uint32)(unsafe.Add(mBase, uint32(v413)+8)) = uint32(v423)
	v426 = v413 + int32(16)
	v429 = F_pg_snprintf(m, v426, int32(64), int32(_a_F_XLogWrite_6), v413)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L13
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v441 = F_time(m)
	mBase = m.M
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v443)+248)) = v441
	v446 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v443)+256)) = v446
	v449 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[20])))
	if v449 != int32(1) {
		v511 = v80
		v514 = v389
		v518 = v388
		goto L28
	} else {
		goto L74
	}
L72:
	;
	F_XLogArchiveNotify(m, v426)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	m.G0 = v413 + int32(80)
	goto L71
L74:
	;
	v453 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[21]))
	v460 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22]))
	v462 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v463 = base.I64_div_u_s(v460, v462)
	if base.Ui64(v453) < base.Ui64(base.I64_extend_i32_s(v455-int32(1))+v463) {
		v511 = v80
		v514 = v389
		v518 = v388
		goto L28
	} else {
		goto L75
	}
L75:
	;
	v468 = base.AtomicRmwXchg32(m, v443, int32(440), int32(1))
	v469 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v468 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	F_s_lock(m, v471+int32(440), int32(_a_F_XLogWrite_4), int32(_a_F_XLogWrite_7), int32(_a_F_XLogWrite_8))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L13
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v480)+200))
	v482 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v480)+440)), uint32(v482))
	v486 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22]))
	if base.Ui64(v486) < base.Ui64(v481) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22])) = v481
	v490 = v481
	goto L82
L81:
	;
	v490 = v486
	goto L82
L82:
	;
	v492 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[21]))
	v499 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v500 = base.I64_div_u_s(v490, v499)
	if base.Ui64(v492) < base.Ui64(base.I64_extend_i32_s(v494-int32(1))+v500) {
		v511 = v469
		v514 = v389
		v518 = v388
		goto L28
	} else {
		goto L83
	}
L83:
	;
	F_RequestCheckpoint(m, int32(128))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v511 = v469
	v514 = v389
	v518 = v388
	goto L28
L85:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v511
	goto L4
L86:
	;
	goto L87
L87:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+304))
	if v59 != v531 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v533 = v59 + int32(1)
	goto L90
L89:
	;
	v533 = int32(0)
	goto L90
L90:
	;
	if base.B2i32(l2 == int32(0))|v514 != 0 {
		v53 = v511
		v55 = v530
		v56 = v514
		v59 = v533
		v60 = v518
		v61 = v145
		goto L2
	} else {
		goto L91
	}
L91:
	;
	goto L4
L92:
	;
	v630 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v633 = base.AtomicRmwXchg32(m, v630, int32(440), int32(1))
	if v633 != 0 {
		goto L107
	} else {
		goto L108
	}
L93:
	;
	v558 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v558) <= base.Ui64(v554) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[23]))
	switch v561 - int32(2) {
	case 0, 2:
		v617 = v558
		goto L95
	default:
		goto L96
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v617
	v623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[18])) = uint8(v623)
	goto L92
L96:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v567 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v611 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	F_issue_xlog_fsync(m, v609, v611, l1)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L13
	} else {
		goto L106
	}
L98:
	;
	v571 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v575 = base.I64_div_u_s(v558-int64(1), base.I64_extend_i32_s(v565))
	if v571 == v575 {
		v609 = v567
		goto L97
	} else {
		goto L101
	}
L99:
	;
	v587 = v558
	v588 = v565
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	v596 = base.I64_div_u_s(v587-int64(1), base.I64_extend_i32_s(v588))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v596
	v599 = F_XLogFileOpen(m, v596, l1)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L13
	} else {
		goto L104
	}
L101:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v580 {
		v609 = v580
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v584 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v587 = v584
	v588 = v586
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v599
	F_ReserveExternalFD(m)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v609 = v605
	goto L97
L106:
	;
	v615 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v617 = v615
	goto L95
L107:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	F_s_lock(m, v635+int32(440), int32(_a_F_XLogWrite_4), int32(2568), int32(_a_F_XLogWrite_5))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v644 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v646)+184))
	if base.Ui64(v647) < base.Ui64(v644) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v646)+184)) = v644
	goto L113
L112:
	;
	goto L113
L113:
	;
	v651 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v646)+192))
	if base.Ui64(v652) < base.Ui64(v651) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v646)+192)) = v651
	goto L116
L115:
	;
	goto L116
L116:
	;
	v655 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v646)+440)), uint32(v655))
	v659 = base.AtomicRmwXchg64(m, v646, int32(272), v644)
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v663 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v665 = base.AtomicRmwXchg64(m, v661, int32(280), v663)
	m.G0 = v19 + int32(96)
	return
L117:
	;
	v674 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+20)) = uint32(v674)
	v676 = int64(32)
	v677 = int64(base.Ui64(v674) >> (uint(v676) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+16)) = uint32(v677)
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+28)) = uint32(v74)
	v681 = int64(base.Ui64(v74) >> (uint(v676) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+24)) = uint32(v681)
	F_errmsg_internal(m, int32(_a_F_XLogWrite_9), v19+int32(16))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_XLogWrite_4), int32(2354), int32(_a_F_XLogWrite_5))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L13
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
