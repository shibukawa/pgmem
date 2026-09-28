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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(0), int32(1))
	if v8 != 0 {
		F_s_lock(m, v5, int32(_a_F_GetXLogInsertRecPtr_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v15 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5))), uint32(v15))
			v19 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[1])))
			v20 = base.I64_div_u_s(v14, v19)
			v22 = v14 - v20*v19
			if base.Ui64(v22) <= base.Ui64(int64(8151)) {
				v40 = v22 + int64(40)
			} else {
				v28 = v22 - int64(8152)
				v29 = int64(8168)
				v30 = base.I64_div_u_s(v28, v29)
				v40 = v28 - v30*v29 + v30<<(uint(int64(13))%64) + int64(8216)
			}
			v42 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[2])))
			return v20*v42 + v40&int64(4294967295)
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v15 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v5))), uint32(v15))
		v19 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[1])))
		v20 = base.I64_div_u_s(v14, v19)
		v22 = v14 - v20*v19
		if base.Ui64(v22) <= base.Ui64(int64(8151)) {
			v40 = v22 + int64(40)
		} else {
			v28 = v22 - int64(8152)
			v29 = int64(8168)
			v30 = base.I64_div_u_s(v28, v29)
			v40 = v28 - v30*v29 + v30<<(uint(int64(13))%64) + int64(8216)
		}
		v42 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetXLogInsertRecPtr[2])))
		return v20*v42 + v40&int64(4294967295)
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	v1 = int64(0)
	v3 = int32(_a_F_GetXLogWriteRecPtr_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[0]))
	v8 = base.AtomicRmwCmpxchg64(m, v4, int32(272), v1, v1)
	*(*int64)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[1])) = v8
	v10 = int32(0)
	v13 = base.AtomicRmwOr32(m, v10, int32(_a_F_GetXLogWriteRecPtr_1), v10)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[0]))
	v20 = base.AtomicRmwCmpxchg64(m, v16, int32(264), v1, v1)
	*(*int64)(unsafe.Add(mBase, _c_F_GetXLogWriteRecPtr[2])) = v20
	return v20
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
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
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
	F_errfinish(m, int32(_a_F_XLogArchiveNotify_4), int32(458), int32(_a_F_XLogArchiveNotify_5))
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
	F_errfinish(m, int32(_a_F_XLogArchiveNotify_4), int32(466), int32(_a_F_XLogArchiveNotify_5))
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
	v204 = v199 + v194*int32(768) + int32(316)
	v205 = int32(0)
	v208 = base.AtomicRmwOr32(m, v205, int32(_a_F_XLogArchiveNotify_9), v205)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v209 != 0 {
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
	v212 = int32(0)
	v215 = base.AtomicRmwOr32(m, v212, int32(_a_F_XLogArchiveNotify_9), v212)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v216 == v212 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	if v219 == int32(0) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[6]))
	if v223 == v219 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v225 = m.G0
	v227 = v225 - int32(16)
	m.G0 = v227
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[7]))
	if v230 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v253 = F_pgmem_kill(m, v219, int32(23))
	mBase = m.M
	goto L57
L64:
	;
	m.G0 = v227 + int32(16)
	goto L56
L65:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+15)) = uint8(v233)
	goto L66
L66:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[8]))
	v241 = F_write(m, v237, v227+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v241 {
		goto L64
	} else {
		goto L68
	}
L67:
	;
	goto L64
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_XLogArchiveNotify[9]))
	if v245 == int32(27) {
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int64
	_ = v407
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v428 int64
	_ = v428
	var v478 int32
	_ = v478
	var v479 int64
	_ = v479
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v495 int64
	_ = v495
	var v499 int32
	_ = v499
	var v515 int32
	_ = v515
	var v516 int64
	_ = v516
	var v520 int64
	_ = v520
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v548 int64
	_ = v548
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int64
	_ = v581
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int64
	_ = v593
	var v594 int64
	_ = v594
	var v598 int64
	_ = v598
	var v648 int32
	_ = v648
	var v649 int64
	_ = v649
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v665 int64
	_ = v665
	var v669 int32
	_ = v669
	var v685 int32
	_ = v685
	var v686 int64
	_ = v686
	var v690 int64
	_ = v690
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
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
		goto L16
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
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L170
	}
L15:
	;
	v756 = int32(_a_F_XLogFileInitInternal_9)
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	v758 = F_close(m, v112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = v757
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L166
	}
L16:
	;
	if v73 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	if v78 == int32(44) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v750 = v73
	goto L19
L19:
	;
	m.G0 = v12 + int32(1168)
	return v750
L20:
	;
	v402 = int32(2)
	v405 = int32(1)
	v407 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0])))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[6])))
	if v410 != 0 {
		goto L104
	} else {
		goto L105
	}
L21:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v398))) = int32(0)
	goto L20
L22:
	;
	v375 = v12 + int32(144)
	v376 = F_unlink(m, v375)
	mBase = m.M
	v377 = F_close(m, v112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = v373
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L100
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L96
	}
L24:
	;
	v83 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L92
	}
L27:
	;
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errmsg_internal(m, int32(_a_F_XLogFileInitInternal_10), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
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
		goto L33
	}
L31:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3290), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = F_unlink(m, v97)
	mBase = m.M
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[1]))
	if v108&int32(4) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v111 = int32(_a_F_XLogFileInitInternal_13)
	goto L36
L35:
	;
	v111 = int32(194)
	goto L36
L36:
	;
	v112 = F_BasicOpenFile(m, v97, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v112 < int32(0) {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[8])))
	v120 = m.G0
	v122 = v120 - int32(16)
	m.G0 = v122
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(167772236)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[6])))
	if v140 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	F___clock_gettime(m, int32(1), v122)
	mBase = m.M
	v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v122)+8)))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	v131 = v126 + v127*int64(1000000000)
	goto L42
L41:
	;
	v131 = int64(0)
	goto L42
L42:
	;
	m.G0 = v122 + int32(16)
	goto L39
L43:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = int32(0)
	if v331 != 0 {
		v373 = v331
		goto L22
	} else {
		goto L91
	}
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0]))
	v155 = m.G0
	v157 = v155 - int32(1024)
	m.G0 = v157
	v160 = v144
	v161 = int64(0)
	v169 = int32(0)
	goto L48
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = int32(0)
	v313 = int32(1)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[0]))
	v319 = F_pwrite(m, v112, int32(_a_F_XLogFileInitInternal_14), v313, base.I64_extend_i32_s(v315-v313))
	mBase = m.M
	if v319 == v313 {
		goto L21
	} else {
		goto L87
	}
L47:
	;
	if v295 < int32(0) {
		goto L43
	} else {
		goto L86
	}
L48:
	;
	v171 = int32(0)
	if v160 == v171 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	m.G0 = v157 + int32(1024)
	goto L47
L50:
	;
	goto L49
L51:
	;
	v295 = v169
	goto L50
L52:
	;
	goto L53
L53:
	;
	v175 = v160
	v177 = v171
	goto L54
L54:
	;
	v188 = v157 + v177<<(uint(int32(3))%32)
	v189 = int32(_a_F_XLogFileInitInternal_15)
	if base.Ui32(v189) <= base.Ui32(v175) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v203 = m.G0
	v205 = v203 - int32(1024)
	m.G0 = v205
	if v197 <= int32(128) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	goto L55
L57:
	;
	v192 = v189
	goto L59
L58:
	;
	v192 = v175
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(_a_F_XLogFileInitInternal_16)
	v197 = v177 + int32(1)
	v198 = v175 - v192
	if base.Ui32(int32(126)) < base.Ui32(v177) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v198 != 0 {
		v175 = v198
		v177 = v197
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	m.G0 = v205 + int32(1024)
	if int32(0) <= v279 {
		v160 = v198
		v161 = v161 + base.I64_extend_i32_u(v279)
		v169 = v169 + v279
		goto L48
	} else {
		goto L85
	}
L63:
	;
	v212 = v157
	v213 = v197
	v216 = int32(0)
	v220 = v161
	goto L66
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5])) = int32(28)
	v279 = int32(-1)
	goto L62
L66:
	;
	if v213 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v279 = v231
	goto L62
L68:
	;
	if v227 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v225 = F_pwrite(m, v112, v223, v224, v220)
	mBase = m.M
	v227 = v225
	goto L68
L70:
	;
	goto L71
L71:
	;
	v226 = F_pwritev(m, v112, v212, v213, v220)
	mBase = m.M
	v227 = v226
	goto L68
L72:
	;
	v279 = int32(-1)
	goto L62
L73:
	;
	goto L74
L74:
	;
	v231 = v227 + v216
	v237 = v212
	v238 = v213
	v240 = v227
	goto L75
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if base.Ui32(v246) <= base.Ui32(v240) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v237 == v205 {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	v252 = v238 - int32(1)
	if v252 != 0 {
		v237 = v237 + int32(8)
		v238 = v252
		v240 = v240 - v246
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	v279 = v231
	goto L62
L81:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v260 + v240
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v263 - v240
	if int32(0) < v238 {
		v212 = v205
		v213 = v238
		v216 = v231
		v220 = v220 + base.I64_extend_i32_u(v227)
		goto L66
	} else {
		goto L84
	}
L82:
	;
	v255 = v238 << (uint(int32(3)) % 32)
	if v255 == int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	base.MemoryCopy(m, v205, v237, v255)
	goto L81
L84:
	;
	goto L67
L85:
	;
	v295 = v279
	goto L50
L86:
	;
	goto L21
L87:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = int32(0)
	if v323 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v329 = v323
	goto L90
L89:
	;
	v329 = int32(51)
	goto L90
L90:
	;
	v373 = v329
	goto L22
L91:
	;
	goto L20
L92:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l3
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_17), v12+int32(96))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3279), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_18), v12+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3304), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v375
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_19), v12-int32(-64))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3357), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v411 = v407
	goto L106
L105:
	;
	v411 = int64(1)
	goto L106
L106:
	;
	v415 = m.G0
	v417 = v415 - int32(16)
	m.G0 = v417
	if v131 != int64(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[8])))
	v537 = m.G0
	v539 = v537 - int32(16)
	m.G0 = v539
	if v534 != 0 {
		goto L125
	} else {
		goto L126
	}
L108:
	;
	F___clock_gettime(m, int32(1), v417)
	mBase = m.M
	v423 = int64(*(*int32)(unsafe.Add(mBase, uint32(v417)+8)))
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v417)))
	v428 = v423 + (v424*int64(1000000000) - v131)
	goto L111
L109:
	;
	goto L110
L110:
	;
	v515 = int32(824)
	v516 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[9])) = v516 + base.I64_extend_i32_u(v405)
	v520 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[10])) = v520 + v411
	F_pgstat_count_backend_io_op(m, v402, v402, int32(7), v405, v411)
	mBase = m.M
	v525 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[12])) = uint8(v525)
	m.G0 = v417 + int32(16)
	goto L107
L111:
	;
	v478 = int32(824)
	v479 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[13])) = v479 + v428
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[2]))
	v490 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v483))|base.B2i32(int32(1)<<(uint(v483)%32)&int32(_a_F_XLogFileInitInternal_20) == v490) == v490 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v495 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[14])) = v495 + v428
	v499 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v499)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[15])) = uint8(v499)
	goto L123
L122:
	;
	goto L123
L123:
	;
	goto L110
L124:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = int32(167772235)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[3])))
	if v558 != int32(1) {
		v572 = int32(0)
		goto L129
	} else {
		goto L130
	}
L125:
	;
	F___clock_gettime(m, int32(1), v539)
	mBase = m.M
	v543 = int64(*(*int32)(unsafe.Add(mBase, uint32(v539)+8)))
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v539)))
	v548 = v543 + v544*int64(1000000000)
	goto L127
L126:
	;
	v548 = int64(0)
	goto L127
L127:
	;
	m.G0 = v539 + int32(16)
	goto L124
L128:
	;
	if v572 != 0 {
		goto L15
	} else {
		goto L135
	}
L129:
	;
	goto L128
L130:
	;
	goto L131
L131:
	;
	v563 = F_fsync(m, v112)
	mBase = m.M
	if v563 != int32(-1) {
		v572 = v563
		goto L129
	} else {
		goto L133
	}
L132:
	;
	v572 = int32(-1)
	goto L129
L133:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[5]))
	if v567 == int32(27) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = int32(0)
	v577 = int32(2)
	v579 = int32(1)
	v581 = int64(0)
	v585 = m.G0
	v587 = v585 - int32(16)
	m.G0 = v587
	if v548 != v581 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v703 = F_close(m, v112)
	mBase = m.M
	if v703 != 0 {
		goto L14
	} else {
		goto L153
	}
L137:
	;
	F___clock_gettime(m, int32(1), v587)
	mBase = m.M
	v593 = int64(*(*int32)(unsafe.Add(mBase, uint32(v587)+8)))
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v587)))
	v598 = v593 + (v594*int64(1000000000) - v548)
	goto L140
L138:
	;
	goto L139
L139:
	;
	v685 = int32(776)
	v686 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[16])) = v686 + base.I64_extend_i32_u(v579)
	v690 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[17])) = v690 + v581
	F_pgstat_count_backend_io_op(m, v577, v577, v579, v579, v581)
	mBase = m.M
	v695 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v695)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[12])) = uint8(v695)
	m.G0 = v587 + int32(16)
	goto L136
L140:
	;
	v648 = int32(776)
	v649 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[18])) = v649 + v598
	v653 = *(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[2]))
	v660 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v653))|base.B2i32(int32(1)<<(uint(v653)%32)&int32(_a_F_XLogFileInitInternal_20) == v660) == v660 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v665 = *(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[19]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[19])) = v665 + v598
	v669 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[11])) = uint8(v669)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[15])) = uint8(v669)
	goto L152
L151:
	;
	goto L152
L152:
	;
	goto L139
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = l0
	v711 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogFileInitInternal[20])))
	v713 = F_InstallXLogFileSegment(m, v12+int32(136), v12+int32(144), int32(1), l0+v711, l1)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	v750 = int32(-1)
	goto L19
L155:
	;
	F_errmsg_internal(m, v736, int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L164
	}
L156:
	;
	if v713 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v715 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v715)
	v719 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v727 = F_unlink(m, v12+int32(144))
	mBase = m.M
	v730 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	if v719 == int32(0) {
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v736 = int32(_a_F_XLogFileInitInternal_21)
	v737 = int32(3412)
	goto L155
L162:
	;
	if v730 == int32(0) {
		goto L154
	} else {
		goto L163
	}
L163:
	;
	v736 = int32(_a_F_XLogFileInitInternal_22)
	v737 = int32(3422)
	goto L155
L164:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), v737, int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L154
L166:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_23), v12+int32(48))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3379), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(144)
	F_errmsg(m, int32(_a_F_XLogFileInitInternal_24), v12+int32(32))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_XLogFileInitInternal_6), int32(3389), int32(_a_F_XLogFileInitInternal_12))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
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
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 int64
	_ = v418
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int64
	_ = v535
	var v536 int64
	_ = v536
	var v537 int64
	_ = v537
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int64
	_ = v623
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int64
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int64
	_ = v647
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int64
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v728 int32
	_ = v728
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v860 int64
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int64
	_ = v905
	var v906 int64
	_ = v906
	var v909 int64
	_ = v909
	var v911 int32
	_ = v911
	var v916 int64
	_ = v916
	var v917 int64
	_ = v917
	var v919 int64
	_ = v919
	var v925 int64
	_ = v925
	var v926 int64
	_ = v926
	var v927 int64
	_ = v927
	var v937 int64
	_ = v937
	var v939 int64
	_ = v939
	var v943 int64
	_ = v943
	var v946 int64
	_ = v946
	var v947 int64
	_ = v947
	var v949 int64
	_ = v949
	var v952 int64
	_ = v952
	var v957 int64
	_ = v957
	var v959 int64
	_ = v959
	var v960 int64
	_ = v960
	var v961 int64
	_ = v961
	var v963 int64
	_ = v963
	var v968 int64
	_ = v968
	var v977 int64
	_ = v977
	var v979 int64
	_ = v979
	var v983 int64
	_ = v983
	var v987 int64
	_ = v987
	var v988 int64
	_ = v988
	var v990 int64
	_ = v990
	var v996 int64
	_ = v996
	var v997 int64
	_ = v997
	var v998 int64
	_ = v998
	var v1008 int64
	_ = v1008
	var v1010 int64
	_ = v1010
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1032 int64
	_ = v1032
	var v1034 int64
	_ = v1034
	var v1035 int64
	_ = v1035
	var v1037 int64
	_ = v1037
	var v1040 int64
	_ = v1040
	var v1045 int64
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1049 int64
	_ = v1049
	var v1051 int64
	_ = v1051
	var v1056 int64
	_ = v1056
	var v1065 int64
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int64
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1072 int64
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int64
	_ = v1075
	var v1076 int64
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1092 int64
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1094 int64
	_ = v1094
	var v1104 int64
	_ = v1104
	var v1109 int64
	_ = v1109
	var v1111 int64
	_ = v1111
	var v1114 int64
	_ = v1114
	var v1119 int64
	_ = v1119
	var v1121 int64
	_ = v1121
	var v1122 int64
	_ = v1122
	var v1123 int64
	_ = v1123
	var v1125 int64
	_ = v1125
	var v1130 int64
	_ = v1130
	var v1139 int64
	_ = v1139
	var v1143 int64
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1151 int64
	_ = v1151
	var v1153 int64
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int64
	_ = v1157
	var v1162 int64
	_ = v1162
	var v1180 int64
	_ = v1180
	var v1187 int64
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1195 int64
	_ = v1195
	var v1197 int64
	_ = v1197
	var v1203 int64
	_ = v1203
	var v1204 int64
	_ = v1204
	var v1205 int64
	_ = v1205
	var v1215 int64
	_ = v1215
	var v1217 int64
	_ = v1217
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int64
	_ = v1263
	var v1264 int64
	_ = v1264
	var v1267 int64
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1273 int64
	_ = v1273
	var v1274 int64
	_ = v1274
	var v1276 int64
	_ = v1276
	var v1282 int64
	_ = v1282
	var v1283 int64
	_ = v1283
	var v1284 int64
	_ = v1284
	var v1294 int64
	_ = v1294
	var v1296 int64
	_ = v1296
	var v1300 int64
	_ = v1300
	var v1302 int64
	_ = v1302
	var v1304 int64
	_ = v1304
	var v1307 int64
	_ = v1307
	var v1312 int64
	_ = v1312
	var v1314 int64
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1318 int64
	_ = v1318
	var v1323 int64
	_ = v1323
	var v1332 int64
	_ = v1332
	var v1336 int64
	_ = v1336
	var v1338 int64
	_ = v1338
	var v1340 int64
	_ = v1340
	var v1346 int64
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1348 int64
	_ = v1348
	var v1358 int64
	_ = v1358
	var v1365 int64
	_ = v1365
	var v1366 int64
	_ = v1366
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1390 int64
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int64
	_ = v1401
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1443 int64
	_ = v1443
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1482 int64
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1510 int64
	_ = v1510
	var v1511 int64
	_ = v1511
	var v1513 int64
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1522 int64
	_ = v1522
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1561 int32
	_ = v1561
	var v1564 int64
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1577 int64
	_ = v1577
	var v1581 int64
	_ = v1581
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int64
	_ = v1620
	var v1627 int64
	_ = v1627
	var v1635 int64
	_ = v1635
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1705 int32
	_ = v1705
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int64
	_ = v1744
	var v1745 int64
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1759 int64
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1762 int64
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1769 int64
	_ = v1769
	var v1772 int64
	_ = v1772
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1784 int64
	_ = v1784
	var v1790 int64
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1795 int64
	_ = v1795
	var v1797 int64
	_ = v1797
	var v1805 int32
	_ = v1805
	var v1814 int64
	_ = v1814
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1826 int64
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int64
	_ = v1830
	var v1833 int64
	_ = v1833
	var v1837 int64
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1843 int64
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1848 int64
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1854 int64
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1860 int64
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1870 int64
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1874 int64
	_ = v1874
	var v1880 int64
	_ = v1880
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1967 int32
	_ = v1967
	var v1972 int32
	_ = v1972
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2032 int32
	_ = v2032
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2128 int32
	_ = v2128
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2145 int64
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2180 int64
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2208 int32
	_ = v2208
	var v2213 int32
	_ = v2213
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2273 int32
	_ = v2273
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2350 int64
	_ = v2350
	var v2392 int32
	_ = v2392
	v1 = l0
	v15 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(_a_F_XLogInsert_0)
	m.G0 = v42
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[0])))
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[2])) = int32(_a_F_XLogInsert_1)
	v2392 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3])) = v2392
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[4])) = v2392
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])) = uint8(v2392)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[0])) = uint8(v2392)
	m.G0 = v42 + int32(_a_F_XLogInsert_0)
	return v2350
L2:
	;
	v2180 = int64(40)
	v2182 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v2182 <= int32(0) {
		v2350 = v2180
		goto L1
	} else {
		goto L388
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L72
	} else {
		goto L384
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L72
	} else {
		goto L380
	}
L5:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L72
	} else {
		goto L377
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
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L72
	} else {
		goto L374
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
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[6]))
	if v49 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[7]))
	v82 = v15
	v83 = v15
	v85 = v15
	goto L14
L13:
	;
	goto L12
L14:
	;
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v42+int32(56)))) = v98
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])))
	*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(55)))) = uint8(v101)
	goto L16
L15:
	;
	v1938 = int32(0)
	v1940 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v1940 <= v1938 {
		v2350 = v1880
		goto L1
	} else {
		goto L363
	}
L16:
	;
	v103 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[11])) = v106
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[12])) = v103
	v112 = v106 + int32(24)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[13]))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114+v1))))
	v119 = v116<<(uint(int32(1))%32) | l1
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	if v121 <= v103 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])))
	if v573&int32(1) == int32(0) {
		v589 = v547
		goto L120
	} else {
		goto L121
	}
L18:
	;
	v124 = int64(0)
	v535 = v124
	v536 = v124
	v537 = v124
	v547 = v112
	v551 = int32(_a_F_XLogInsert_2)
	v561 = v82
	v562 = v83
	v564 = v85
	v565 = v103
	goto L17
L19:
	;
	goto L20
L20:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v42)+56))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	v134 = int64(0)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+55)))
	v140 = int32(0)
	v144 = v134
	v145 = v134
	v146 = v134
	v156 = v112
	v159 = v140
	v160 = int32(_a_F_XLogInsert_2)
	v162 = v121
	v163 = v132
	v167 = v140
	v170 = v82
	v171 = v83
	v173 = v85
	v174 = v103
	goto L21
L21:
	;
	v183 = v163 + v167*int32(_a_F_XLogInsert_3)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v184 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v535 = v509
	v536 = v510
	v537 = v511
	v547 = v513
	v551 = v516
	v561 = v523
	v562 = v524
	v564 = v525
	v565 = v526
	goto L17
L23:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v187&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v509 = v144
	v510 = v145
	v511 = v146
	v513 = v156
	v515 = v159
	v516 = v160
	v518 = v162
	v519 = v163
	v523 = v170
	v524 = v171
	v525 = v173
	v526 = v174
	goto L25
L25:
	;
	v531 = v167 + int32(1)
	if v531 < v518 {
		v144 = v509
		v145 = v510
		v146 = v511
		v156 = v513
		v159 = v515
		v160 = v516
		v162 = v518
		v163 = v519
		v167 = v531
		v170 = v523
		v171 = v524
		v173 = v525
		v174 = v526
		goto L21
	} else {
		goto L119
	}
L26:
	;
	v210 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	if v220 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v207 = v145
	v209 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v191 = int32(0)
	if base.B2i32(v135&int32(1) == v191)|v187&int32(2) != 0 {
		v207 = v145
		v209 = v191
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v183)+24))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v197)))
	v200 = base.I64_rotl(v198, int64(32))
	if base.Ui64(v200) <= base.Ui64(v128) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v207 = v145
	v209 = int32(1)
	goto L26
L32:
	;
	goto L33
L33:
	;
	if base.Ui64(v145-int64(1)) < base.Ui64(v200) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v206 = v145
	goto L36
L35:
	;
	v206 = v200
	goto L36
L36:
	;
	v207 = v206
	v209 = v191
	goto L26
L37:
	;
	v221 = v209 ^ int32(1) | int32(base.Ui32(v187&int32(16))>>(uint(int32(4))%32))
	goto L39
L38:
	;
	v221 = v210
	goto L39
L39:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+16)))
	v225 = int32(6)
	if v187&v225 == v225 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v229 = v222 | int32(64)
	goto L42
L41:
	;
	v229 = v222
	goto L42
L42:
	;
	v232 = base.B2i32(v119&int32(2) != int32(0)) | v209
	if v232 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v437 = int32(0)
	if v221 == v437 {
		goto L103
	} else {
		goto L104
	}
L44:
	;
	v421 = v144
	v422 = v146
	v425 = v160
	v427 = v229
	v430 = int32(0)
	v431 = v210
	v432 = v170
	v433 = v171
	v434 = v173
	v435 = v174
	goto L43
L45:
	;
	goto L46
L46:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v183)+24))
	v237 = int32(0)
	if v187&int32(8) == v237 {
		v257 = v210
		v258 = v237
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[15]))
	if v260 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L48:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236)+12)))
	if base.Ui32(v242) < base.Ui32(int32(24)) {
		v257 = v210
		v258 = v237
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236)+14)))
	v251 = base.B2i32(base.Ui32(v242) < base.Ui32(v245)) & base.B2i32(base.Ui32(v245) < base.Ui32(int32(_a_F_XLogInsert_4)))
	if v251 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v252 = v245 - v242
	goto L52
L51:
	;
	v252 = int32(0)
	goto L52
L52:
	;
	if v251 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v254 = v242
	goto L55
L54:
	;
	v254 = int32(0)
	goto L55
L55:
	;
	v257 = v252
	v258 = v254
	goto L47
L56:
	;
	v337 = v183 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v337
	v340 = v257 & int32(_a_F_XLogInsert_5)
	v342 = base.B2i32(v340 != int32(0))
	if v209 != 0 {
		goto L82
	} else {
		goto L83
	}
L57:
	;
	v265 = int32(0)
	v332 = v257 & int32(_a_F_XLogInsert_5)
	v333 = v265
	v335 = v265
	goto L56
L58:
	;
	goto L59
L59:
	;
	v267 = int32(0)
	v270 = v257 & int32(_a_F_XLogInsert_5)
	if v270 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v258 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v285 = v236
	v287 = v267
	goto L62
L62:
	;
	switch v260 - int32(1) {
	case 0:
		goto L69
	case 1:
		goto L71
	case 2:
		goto L70
	default:
		v332 = v270
		v333 = int32(0)
		v335 = v267
		goto L56
	}
L63:
	;
	base.MemoryCopy(m, v42-int32(-64), v236, v258)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v275 = v270 + v258
	v276 = int32(_a_F_XLogInsert_6) - v275
	if v276 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	base.MemoryCopy(m, v42-int32(-64)+v258, v275+v236, v276)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v285 = v42 - int32(-64)
	v287 = int32(2)
	goto L62
L69:
	;
	v320 = int32(_a_F_XLogInsert_6) - v270
	v323 = F_pglz_compress(m, v285, v320, v183-int32(-64), v53)
	mBase = m.M
	v324 = int32(0)
	v329 = base.B2i32(v323+v287 < v320) & base.B2i32(v324 <= v323)
	if v329 != 0 {
		goto L79
	} else {
		goto L80
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L72
	} else {
		goto L76
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
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
	v300 = m.ExcPending
	if v300 != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(1061), int32(_a_F_XLogInsert_9))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	v313 = m.ExcPending
	if v313 != 0 {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(1072), int32(_a_F_XLogInsert_9))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
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
	v330 = v323
	goto L81
L80:
	;
	v330 = v324
	goto L81
L81:
	;
	v332 = v270
	v333 = v329
	v335 = v330
	goto L56
L82:
	;
	v345 = v342 | int32(2)
	goto L84
L83:
	;
	v345 = v342
	goto L84
L84:
	;
	if v333 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183+v406))) = v409
	v418 = base.I64_extend_i32_u(v408) & int64(65535)
	v421 = v144 + v418
	v422 = v146 + v418
	v425 = v405
	v427 = v229 | int32(16)
	v430 = v333
	v431 = v257
	v432 = v258
	v433 = v407
	v434 = v408
	v435 = v174 + int32(1)
	goto L43
L86:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[15]))
	switch v347 - int32(1) {
	case 0:
		goto L90
	case 1:
		goto L92
	case 2:
		goto L91
	default:
		v378 = v345
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v340 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+44)) = v183 - int32(-64)
	v405 = v337
	v406 = int32(48)
	v407 = v378
	v408 = v335
	v409 = v335 & int32(_a_F_XLogInsert_5)
	goto L85
L90:
	;
	v378 = v345 | int32(4)
	goto L89
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L72
	} else {
		goto L96
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L72
	} else {
		goto L93
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_7), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L72
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(813), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
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
	v370 = m.ExcPending
	if v370 != 0 {
		goto L72
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(821), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v183)+44)) = v236
	v389 = int32(_a_F_XLogInsert_6)
	v405 = v337
	v406 = int32(48)
	v407 = v345
	v408 = v389
	v409 = v389
	goto L85
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v183)+44)) = v236
	v394 = v183 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+40)) = v394
	v396 = v332 + v258
	*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v236 + v396
	v399 = int32(_a_F_XLogInsert_6)
	v405 = v394
	v406 = int32(60)
	v407 = v345
	v408 = v399 - v257
	v409 = v399 - v396
	goto L85
L102:
	;
	if v159 == int32(0) {
		v467 = v437
		v468 = v450
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v449 = v422
	v450 = v427
	v451 = int32(0)
	v452 = v425
	goto L102
L104:
	;
	goto L105
L105:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v183)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v442
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v183)+36))
	v449 = v422 + base.I64_extend_i32_u(v441)
	v450 = v427 | int32(32)
	v451 = v441
	v452 = v448
	goto L102
L106:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+2)) = uint16(v451)
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)) = uint8(v468)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v167)
	if v232 == int32(0) {
		v491 = v156 + int32(4)
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v455 != v456 {
		v467 = v437
		v468 = v450
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v458 != v459 {
		v467 = v437
		v468 = v450
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v465 = base.B2i32(v463 == v464)
	if v463 == v464 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v466 = v450 | int32(-128)
	goto L112
L111:
	;
	v466 = v450
	goto L112
L112:
	;
	v467 = v465
	v468 = v466
	goto L106
L113:
	;
	if v467 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+8)) = uint8(v433)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)) = uint16(v432)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v434)
	if base.B2i32(v431&int32(_a_F_XLogInsert_5) == int32(0))|(v430^int32(1)) != 0 {
		v491 = v156 + int32(9)
		goto L113
	} else {
		goto L115
	}
L115:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+9)) = uint16(v431)
	v491 = v156 + int32(11)
	goto L113
L116:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v491)+8)) = v494
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v183)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v496
	v500 = v491 + int32(12)
	goto L118
L117:
	;
	v500 = v491
	goto L118
L118:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[3]))
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	v509 = v421
	v510 = v207
	v511 = v449
	v513 = v500 + int32(4)
	v515 = v183
	v516 = v452
	v518 = v506
	v519 = v508
	v523 = v432
	v524 = v433
	v525 = v434
	v526 = v435
	goto L25
L119:
	;
	goto L22
L120:
	;
	v590 = int32(0)
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592)+78)))
	if v593 != 0 {
		v613 = v590
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v579 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XLogInsert[17])))
	if v579 == int32(0) {
		v589 = v547
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v582 = int32(253)
	*(*uint8)(unsafe.Add(mBase, uint32(v547))) = uint8(v582)
	v585 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XLogInsert[17])))
	*(*uint16)(unsafe.Add(mBase, uint32(v547)+1)) = uint16(v585)
	v589 = v547 + int32(3)
	goto L120
L123:
	;
	if v613 != 0 {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[18]))
	if v595 <= int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[19])))
	if v599&int32(1) == int32(0) {
		v613 = v590
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v592)+20))
	if v604 != int32(2) {
		v613 = v590
		goto L123
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v592)+28))
	if v607 < int32(2) {
		v613 = v590
		goto L123
	} else {
		goto L130
	}
L130:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	v613 = base.B2i32(v610 != int32(0))
	goto L123
L131:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v589)+1)) = v615
	v617 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v589))) = uint8(v617)
	v621 = v589 + int32(5)
	goto L133
L132:
	;
	v621 = v589
	goto L133
L133:
	;
	v623 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	if v623 != int64(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if base.Ui64(int64(256)) <= base.Ui64(v623) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v651 = v537
	v652 = v621
	v653 = v551
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = int32(0)
	v658 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[10]))
	v659 = v652 - v658
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[21])) = v659
	v662 = v651 + base.I64_extend_i32_u(v659)
	v664 = int32(24)
	v668 = m.Env.Pgmem_crc32c(m, int32(-1), v658+v664, v659-v664)
	mBase = m.M
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[12]))
	if v670 != 0 {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v551))) = v644
	v647 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[2]))
	v651 = v647 + v537
	v652 = v642
	v653 = v650
	goto L136
L138:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v623) {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v635 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v635)
	v638 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)) = uint8(v638)
	v642 = v621 + int32(2)
	goto L137
L141:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v621)+1)) = uint32(v623)
	v631 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v631)
	v642 = v621 + int32(5)
	goto L137
L142:
	;
	v685 = v668
	v686 = v670
	goto L145
L143:
	;
	v728 = v668
	goto L144
L144:
	;
	if base.Ui64(int64(1069547521)) <= base.Ui64(v662) {
		goto L3
	} else {
		goto L148
	}
L145:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v686)+4))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v686)+8))
	v712 = m.Env.Pgmem_crc32c(m, v685, v710, v711)
	mBase = m.M
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	if v713 != 0 {
		v685 = v712
		v686 = v713
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v728 = v712
	goto L144
L147:
	;
	goto L146
L148:
	;
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	goto L149
L149:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+17)) = uint8(v1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+16)) = uint8(v119)
	*(*uint32)(unsafe.Add(mBase, uint32(v106))) = uint32(v662)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v757
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v728
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = int64(0)
	v766 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[5])))
	v767 = int32(0)
	v770 = m.G0
	v772 = v770 - int32(16)
	m.G0 = v772
	v776 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[11]))
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776)+17)))
	if v777 != 0 {
		v790 = v767
		v791 = int32(1)
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if v1880 == int64(0) {
		v82 = v561
		v83 = v562
		v85 = v564
		goto L14
	} else {
		goto L362
	}
L151:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v795 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])))
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24]))
	if v797 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L152:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776)+16)))
	v780 = v778 & int32(240)
	if v780 == int32(64) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v790 = int32(1)
	v791 = int32(0)
	goto L151
L154:
	;
	goto L155
L155:
	;
	if v780 == int32(224) {
		v790 = v767
		v791 = int32(0)
		goto L151
	} else {
		goto L156
	}
L156:
	;
	v790 = v767
	v791 = int32(1)
	goto L151
L157:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L72
	} else {
		goto L358
	}
L158:
	;
	m.G0 = v772 + int32(16)
	goto L150
L159:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L72
	} else {
		goto L323
	}
L160:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v776)+20))
	v1383 = m.Env.Pgmem_crc32c(m, v1381, v776, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v776)+20)) = v1383 ^ int32(-1)
	v1390 = v1366 & int64(8191)
	if v1390 == int64(0) {
		goto L280
	} else {
		goto L281
	}
L161:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	v1256 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v1259 = base.AtomicRmwXchg32(m, v1256, int32(0), int32(1))
	if v1259 != 0 {
		goto L260
	} else {
		goto L261
	}
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L72
	} else {
		goto L257
	}
L163:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v793)+300))
	v817 = int32(_a_F_XLogInsert_12)
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25])) = v819 + int32(1)
	if v791 != 0 {
		goto L172
	} else {
		goto L173
	}
L164:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[26])))
	if v801 == int32(1) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	if v797 == int32(0) {
		goto L162
	} else {
		goto L171
	}
L167:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v793)+308))
	v807 = base.B2i32(v805 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[26])) = uint8(v807)
	if v805 != int32(2) {
		goto L162
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[24])) = int32(1)
	goto L163
L170:
	;
	goto L169
L171:
	;
	goto L163
L172:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[27]))
	if v824 == int32(-1) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L72
	} else {
		goto L192
	}
L175:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[28]))
	v831 = base.I32_rem_s(v829, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[27])) = v831
	v833 = v831
	goto L177
L176:
	;
	v833 = v824
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[29])) = v833
	v837 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30]))
	v842 = F_LWLockAcquire(m, v837+v833<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L72
	} else {
		goto L178
	}
L178:
	;
	if v842 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v846 = int32(_a_F_XLogInsert_13)
	v848 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[27]))
	v852 = base.I32_rem_s(v848+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[27])) = v852
	goto L181
L180:
	;
	goto L181
L181:
	;
	v855 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8]))
	v856 = *(*int64)(unsafe.Add(mBase, uint32(v793)+152))
	if v855 != v856 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8])) = v856
	v860 = v856
	goto L184
L183:
	;
	v860 = v855
	goto L184
L184:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+160)))
	if v861 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v795&int32(1)&base.B2i32(base.Ui64(v860) <= base.Ui64(v536-int64(1))) != 0 {
		goto L161
	} else {
		goto L190
	}
L186:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v793)+164))
	v867 = base.B2i32(int32(0) < v865)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])) = uint8(v867)
	if int32(0) < v865 {
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v870 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[9])) = uint8(v870)
	goto L185
L189:
	;
	goto L161
L190:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L72
	} else {
		goto L191
	}
L191:
	;
	v881 = int32(_a_F_XLogInsert_12)
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25])) = v883 - int32(1)
	v1880 = int64(0)
	goto L158
L192:
	;
	if v790 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v901 = base.AtomicRmwXchg32(m, v898, int32(0), int32(1))
	if v901 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	v1019 = int32(0)
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v1024 = base.AtomicRmwXchg32(m, v1021, v1019, int32(1))
	if v1024 != 0 {
		goto L218
	} else {
		goto L219
	}
L196:
	;
	F_s_lock(m, v898, int32(_a_F_XLogInsert_14))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L72
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v905 = *(*int64)(unsafe.Add(mBase, uint32(v898)+16))
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v898)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v898)+16)) = v906
	v909 = v906 + base.I64_extend_i32_s((v892+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v898)+8)) = v909
	v911 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v898))), uint32(v911))
	v916 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v917 = base.I64_div_u_s(v906, v916)
	v919 = v906 - v917*v916
	if base.Ui64(v919) <= base.Ui64(int64(8151)) {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	goto L198
L200:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v772)+8)) = v943
	v946 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v947 = base.I64_div_u_s(v909, v946)
	v949 = v909 - v947*v946
	if base.Ui64(v949) <= base.Ui64(int64(8151)) {
		goto L206
	} else {
		goto L207
	}
L201:
	;
	v939 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	v943 = v917*v939 + v937&int64(4294967295)
	goto L200
L202:
	;
	v937 = v919 + int64(40)
	goto L201
L203:
	;
	goto L204
L204:
	;
	v925 = v919 - int64(8152)
	v926 = int64(8168)
	v927 = base.I64_div_u_s(v925, v926)
	v937 = v925 - v927*v926 + v927<<(uint(int64(13))%64) + int64(8216)
	goto L201
L205:
	;
	v979 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	v983 = v947*v979 + v977&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v983
	v987 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v988 = base.I64_div_u_s(v905, v987)
	v990 = v905 - v988*v987
	if base.Ui64(v990) <= base.Ui64(int64(8151)) {
		goto L215
	} else {
		goto L216
	}
L206:
	;
	v952 = int64(0)
	if v949 == v952 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L208
L208:
	;
	v959 = v949 - int64(8152)
	v960 = int64(8168)
	v961 = base.I64_div_u_s(v959, v960)
	v963 = v961 << (uint(int64(13)) % 64)
	v968 = v959 - v961*v960
	if v968 == int64(0) {
		v977 = v963 - int64(-8192)
		goto L205
	} else {
		goto L212
	}
L209:
	;
	v957 = v952
	goto L211
L210:
	;
	v957 = v949 + int64(40)
	goto L211
L211:
	;
	v977 = v957
	goto L205
L212:
	;
	v977 = v968 + v963 + int64(8216)
	goto L205
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v776)+8)) = v988*v1010 + v1008&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v793)+152)) = v943
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[8])) = v943
	v1365 = v983
	v1366 = v943
	goto L160
L214:
	;
	v1010 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	goto L213
L215:
	;
	v1008 = v990 + int64(40)
	goto L214
L216:
	;
	goto L217
L217:
	;
	v996 = v990 - int64(8152)
	v997 = int64(8168)
	v998 = base.I64_div_u_s(v996, v997)
	v1008 = v996 - v998*v997 + v998<<(uint(int64(13))%64) + int64(8216)
	goto L214
L218:
	;
	F_s_lock(m, v1021, int32(_a_F_XLogInsert_14))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L72
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1028 = int32(8)
	v1029 = v772 + v1028
	v1032 = *(*int64)(unsafe.Add(mBase, uint32(v1021)+8))
	v1034 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v1035 = base.I64_div_u_s(v1032, v1034)
	v1037 = v1032 - v1035*v1034
	if base.Ui64(v1037) <= base.Ui64(int64(8151)) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L220
L222:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32]))
	v1068 = base.I64_extend_i32_s(v1067)
	v1069 = v1035 * v1068
	v1072 = v1069 + v1065&int64(4294967295)
	v1074 = v1067 - int32(1)
	v1075 = base.I64_extend_i32_s(v1074)
	v1076 = v1072 & v1075
	if v1076 == int64(0) {
		goto L231
	} else {
		goto L232
	}
L223:
	;
	v1040 = int64(0)
	if v1037 == v1040 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v1047 = v1037 - int64(8152)
	v1048 = int64(8168)
	v1049 = base.I64_div_u_s(v1047, v1048)
	v1051 = v1049 << (uint(int64(13)) % 64)
	v1056 = v1047 - v1049*v1048
	if v1056 == int64(0) {
		v1065 = v1051 - int64(-8192)
		goto L222
	} else {
		goto L229
	}
L226:
	;
	v1045 = v1040
	goto L228
L227:
	;
	v1045 = v1037 + int64(40)
	goto L228
L228:
	;
	v1065 = v1045
	goto L222
L229:
	;
	v1065 = v1056 + v1051 + int64(8216)
	goto L222
L230:
	;
	if v1076 == int64(0) {
		v1705 = v1019
		goto L159
	} else {
		goto L256
	}
L231:
	;
	v1079 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1021))), uint32(v1079))
	*(*int64)(unsafe.Add(mBase, uint32(v1029))) = v1072
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1072
	goto L230
L232:
	;
	goto L233
L233:
	;
	v1085 = v1032 + int64(24)
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1021)+16))
	if base.Ui64(v1037) <= base.Ui64(int64(8151)) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1104 = v1037 + int64(40)
	goto L236
L235:
	;
	v1092 = v1037 - int64(8152)
	v1093 = int64(8168)
	v1094 = base.I64_div_u_s(v1092, v1093)
	v1104 = v1092 - v1094*v1093 + v1094<<(uint(int64(13))%64) + int64(8216)
	goto L236
L236:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1029))) = v1104&int64(4294967295) + v1069
	v1109 = base.I64_div_u_s(v1085, v1034)
	v1111 = v1085 - v1109*v1034
	if base.Ui64(v1111) <= base.Ui64(int64(8151)) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1143 = v1068*v1109 + v1139&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1143
	v1146 = v1074 & base.I32_wrap_i64(v1143)
	if v1146 == int32(0) {
		v1187 = v1085
		goto L245
	} else {
		goto L246
	}
L238:
	;
	v1114 = int64(0)
	if v1111 == v1114 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	v1121 = v1111 - int64(8152)
	v1122 = int64(8168)
	v1123 = base.I64_div_u_s(v1121, v1122)
	v1125 = v1123 << (uint(int64(13)) % 64)
	v1130 = v1121 - v1123*v1122
	if v1130 == int64(0) {
		v1139 = v1125 - int64(-8192)
		goto L237
	} else {
		goto L244
	}
L241:
	;
	v1119 = v1114
	goto L243
L242:
	;
	v1119 = v1111 + int64(40)
	goto L243
L243:
	;
	v1139 = v1119
	goto L237
L244:
	;
	v1139 = v1130 + v1125 + int64(8216)
	goto L237
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1021)+16)) = v1032
	*(*int64)(unsafe.Add(mBase, uint32(v1021)+8)) = v1187
	v1192 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1021))), uint32(v1192))
	v1195 = base.I64_div_u_s(v1086, v1034)
	v1197 = v1086 - v1034*v1195
	if base.Ui64(v1197) <= base.Ui64(int64(8151)) {
		goto L253
	} else {
		goto L254
	}
L246:
	;
	v1151 = v1143 + base.I64_extend_i32_u(v1067-v1146)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1151
	v1153 = base.I64_div_u_s(v1151, v1068)
	v1156 = base.I32_wrap_i64(v1151) & int32(_a_F_XLogInsert_15)
	v1157 = v1151 & v1075
	if v1157&int64(35184372080640) == int64(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1162 = v1153 * v1034
	if v1156 == int32(0) {
		v1187 = v1162
		goto L245
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1180 = v1153*v1034 + (int64(base.Ui64(v1157)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v1156 == int32(0) {
		v1187 = v1180
		goto L245
	} else {
		goto L251
	}
L250:
	;
	v1187 = v1162 + base.I64_extend_i32_u(v1156-int32(40))
	goto L245
L251:
	;
	v1187 = v1180 + base.I64_extend_i32_u(v1156-int32(24))
	goto L245
L252:
	;
	v1217 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	*(*int64)(unsafe.Add(mBase, uint32(v776+v1028))) = v1195*v1217 + v1215&int64(4294967295)
	goto L230
L253:
	;
	v1215 = v1197 + int64(40)
	goto L252
L254:
	;
	goto L255
L255:
	;
	v1203 = v1197 - int64(8152)
	v1204 = int64(8168)
	v1205 = base.I64_div_u_s(v1203, v1204)
	v1215 = v1203 - v1205*v1204 + v1205<<(uint(int64(13))%64) + int64(8216)
	goto L252
L256:
	;
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v772)+8))
	v1365 = v1233
	v1366 = v1234
	goto L160
L257:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_16), int32(0))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L72
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_17), int32(816), int32(_a_F_XLogInsert_18))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L72
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_s_lock(m, v1256, int32(_a_F_XLogInsert_14))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L72
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1263 = *(*int64)(unsafe.Add(mBase, uint32(v1256)+16))
	v1264 = *(*int64)(unsafe.Add(mBase, uint32(v1256)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1256)+16)) = v1264
	v1267 = v1264 + base.I64_extend_i32_s((v1250+int32(7))&int32(-8))
	*(*int64)(unsafe.Add(mBase, uint32(v1256)+8)) = v1267
	v1269 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1256))), uint32(v1269))
	v1273 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[31])))
	v1274 = base.I64_div_u_s(v1264, v1273)
	v1276 = v1264 - v1274*v1273
	if base.Ui64(v1276) <= base.Ui64(int64(8151)) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	goto L262
L264:
	;
	v1296 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32])))
	v1300 = v1274*v1296 + v1294&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v772)+8)) = v1300
	v1302 = base.I64_div_u_s(v1267, v1273)
	v1304 = v1267 - v1302*v1273
	if base.Ui64(v1304) <= base.Ui64(int64(8151)) {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	v1294 = v1276 + int64(40)
	goto L264
L266:
	;
	goto L267
L267:
	;
	v1282 = v1276 - int64(8152)
	v1283 = int64(8168)
	v1284 = base.I64_div_u_s(v1282, v1283)
	v1294 = v1282 - v1284*v1283 + v1284<<(uint(int64(13))%64) + int64(8216)
	goto L264
L268:
	;
	v1336 = v1302*v1296 + v1332&int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1336
	v1338 = base.I64_div_u_s(v1263, v1273)
	v1340 = v1263 - v1338*v1273
	if base.Ui64(v1340) <= base.Ui64(int64(8151)) {
		goto L277
	} else {
		goto L278
	}
L269:
	;
	v1307 = int64(0)
	if v1304 == v1307 {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	goto L271
L271:
	;
	v1314 = v1304 - int64(8152)
	v1315 = int64(8168)
	v1316 = base.I64_div_u_s(v1314, v1315)
	v1318 = v1316 << (uint(int64(13)) % 64)
	v1323 = v1314 - v1316*v1315
	if v1323 == int64(0) {
		v1332 = v1318 - int64(-8192)
		goto L268
	} else {
		goto L275
	}
L272:
	;
	v1312 = v1307
	goto L274
L273:
	;
	v1312 = v1304 + int64(40)
	goto L274
L274:
	;
	v1332 = v1312
	goto L268
L275:
	;
	v1332 = v1323 + v1318 + int64(8216)
	goto L268
L276:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v776)+8)) = v1338*v1296 + v1358&int64(4294967295)
	v1365 = v1336
	v1366 = v1300
	goto L160
L277:
	;
	v1358 = v1340 + int64(40)
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1346 = v1340 - int64(8152)
	v1347 = int64(8168)
	v1348 = base.I64_div_u_s(v1346, v1347)
	v1358 = v1346 - v1348*v1347 + v1348<<(uint(int64(13))%64) + int64(8216)
	goto L276
L280:
	;
	v1395 = int32(0)
	goto L282
L281:
	;
	v1395 = int32(_a_F_XLogInsert_6) - base.I32_wrap_i64(v1390)
	goto L282
L282:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	v1397 = F_GetXLogBuffer(m, v1366, v816)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L72
	} else {
		goto L283
	}
L283:
	;
	v1401 = v1366
	v1413 = v1397
	v1416 = v1395
	v1419 = int32(_a_F_XLogInsert_2)
	v1423 = v767
	goto L284
L284:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+4))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+8))
	if v1416 < v1439 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	if v790 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L286:
	;
	v1443 = v1401
	v1455 = v1413
	v1458 = v1416
	v1459 = v1439
	v1460 = v1438
	v1465 = v1423
	goto L289
L287:
	;
	v1522 = v1401
	v1534 = v1413
	v1537 = v1416
	v1538 = v1439
	v1539 = v1438
	v1544 = v1423
	goto L288
L288:
	;
	if v1538 != 0 {
		goto L305
	} else {
		goto L306
	}
L289:
	;
	if v1458 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1522 = v1511
	v1534 = v1503
	v1537 = v1518
	v1538 = v1505
	v1539 = v1504
	v1544 = v1485
	goto L288
L291:
	;
	base.MemoryCopy(m, v1455, v1460, v1458)
	goto L293
L292:
	;
	goto L293
L293:
	;
	v1482 = v1443 + base.I64_extend_i32_s(v1458)
	v1483 = F_GetXLogBuffer(m, v1482, v816)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L72
	} else {
		goto L294
	}
L294:
	;
	v1485 = v1458 + v1465
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+16)) = v1396 - v1485
	v1488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1483)+2)))
	v1489 = int32(1)
	v1490 = v1488 | v1489
	*(*uint16)(unsafe.Add(mBase, uint32(v1483)+2)) = uint16(v1490)
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32]))
	v1501 = base.B2i32(v1482&base.I64_extend_i32_s(v1495-v1489) == int64(0))
	if v1482&base.I64_extend_i32_s(v1495-v1489) == int64(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1502 = int32(40)
	goto L297
L296:
	;
	v1502 = int32(24)
	goto L297
L297:
	;
	v1503 = v1483 + v1502
	v1504 = v1458 + v1460
	v1505 = v1459 - v1458
	if v1482&base.I64_extend_i32_s(v1495-v1489) == int64(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1510 = int64(40)
	goto L300
L299:
	;
	v1510 = int64(24)
	goto L300
L300:
	;
	v1511 = v1510 + v1482
	v1513 = v1511 & int64(8191)
	if v1513 == int64(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1518 = int32(0)
	goto L303
L302:
	;
	v1518 = int32(_a_F_XLogInsert_6) - base.I32_wrap_i64(v1513)
	goto L303
L303:
	;
	if v1518 < v1505 {
		v1443 = v1511
		v1455 = v1503
		v1458 = v1518
		v1459 = v1505
		v1460 = v1504
		v1465 = v1485
		goto L289
	} else {
		goto L304
	}
L304:
	;
	goto L290
L305:
	;
	base.MemoryCopy(m, v1534, v1539, v1538)
	goto L307
L306:
	;
	goto L307
L307:
	;
	v1561 = v1537 - v1538
	v1564 = v1522 + base.I64_extend_i32_s(v1538)
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1419)))
	if v1565 != 0 {
		v1401 = v1564
		v1413 = v1534 + v1538
		v1416 = v1561
		v1419 = v1565
		v1423 = v1538 + v1544
		goto L284
	} else {
		goto L308
	}
L308:
	;
	goto L285
L309:
	;
	if v1635 != v1365 {
		goto L157
	} else {
		goto L318
	}
L310:
	;
	v1635 = (v1564 + int64(7)) & int64(-8)
	goto L309
L311:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32]))
	if v1564&base.I64_extend_i32_s(v1569-int32(1)) == int64(0) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1577 = v1564 + base.I64_extend_i32_s(v1561)
	if base.Ui64(v1365) <= base.Ui64(v1577) {
		v1635 = v1577
		goto L309
	} else {
		goto L313
	}
L313:
	;
	v1581 = v1577
	goto L314
L314:
	;
	v1618 = F_GetXLogBuffer(m, v1581, v816)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L72
	} else {
		goto L316
	}
L315:
	;
	v1635 = v1627
	goto L309
L316:
	;
	v1620 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1618)+16)) = v1620
	*(*int64)(unsafe.Add(mBase, uint32(v1618)+8)) = v1620
	*(*int64)(unsafe.Add(mBase, uint32(v1618))) = v1620
	v1627 = v1581 - int64(-8192)
	if base.Ui64(v1627) < base.Ui64(v1365) {
		v1581 = v1627
		goto L314
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1673 = int32(1)
	if v766&int32(2) != 0 {
		v1705 = v1673
		goto L159
	} else {
		goto L319
	}
L319:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[30]))
	v1680 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[29]))
	v1682 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[33])))
	if v1682 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1683 = int32(0)
	goto L322
L321:
	;
	v1683 = v1680
	goto L322
L322:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1677+v1683<<(uint(int32(7))%32))+24)) = v1366
	v1705 = v1673
	goto L159
L323:
	;
	v1729 = int32(_a_F_XLogInsert_12)
	v1731 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[25])) = v1731 - int32(1)
	v1736 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1736)))
	if v1737 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1738 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1736)+70)) = uint8(v1738)
	goto L326
L325:
	;
	goto L326
L326:
	;
	if v613 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[16]))
	v1742 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1741)+78)) = uint8(v1742)
	goto L329
L328:
	;
	goto L329
L329:
	;
	v1744 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	v1745 = *(*int64)(unsafe.Add(mBase, uint32(v772)+8))
	if base.Ui64(int64(8192)) <= base.Ui64(v1744^v1745) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v1753 = base.AtomicRmwXchg32(m, v1750, int32(440), int32(1))
	if v1753 != 0 {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	goto L332
L332:
	;
	if v790 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L333:
	;
	F_s_lock(m, v1750+int32(440), int32(_a_F_XLogInsert_14))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L72
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1759 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	v1761 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v1762 = *(*int64)(unsafe.Add(mBase, uint32(v1761)+184))
	if base.Ui64(v1762) < base.Ui64(v1759) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	goto L335
L337:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1761)+184)) = v1759
	goto L339
L338:
	;
	goto L339
L339:
	;
	v1765 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1761)+440)), uint32(v1765))
	v1769 = int64(0)
	v1772 = base.AtomicRmwCmpxchg64(m, v1761, int32(272), v1769, v1769)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[34])) = v1772
	v1777 = base.AtomicRmwOr32(m, v1765, int32(_a_F_XLogInsert_19), v1765)
	v1780 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[23]))
	v1784 = base.AtomicRmwCmpxchg64(m, v1780, int32(264), v1769, v1769)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[35])) = v1784
	goto L332
L340:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[36])))
	if v1864 != 0 {
		goto L354
	} else {
		goto L355
	}
L341:
	;
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[36])))
	if v1820 != 0 {
		goto L349
	} else {
		goto L350
	}
L342:
	;
	v1790 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	F_XLogFlush(m, v1790)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L72
	} else {
		goto L343
	}
L343:
	;
	if v1705 == int32(0) {
		goto L340
	} else {
		goto L344
	}
L344:
	;
	v1795 = *(*int64)(unsafe.Add(mBase, uint32(v772)+8))
	v1797 = v1795 + int64(24)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1797
	if base.Ui64(v1795^v1797) < base.Ui64(int64(8192)) {
		goto L341
	} else {
		goto L345
	}
L345:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[32]))
	if v1797&base.I64_extend_i32_s(v1805-int32(1)^int32(_a_F_XLogInsert_15)) == int64(0) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1814 = int64(64)
	goto L348
L347:
	;
	v1814 = int64(48)
	goto L348
L348:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v1814 + v1795
	goto L341
L349:
	;
	v1822 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[36])) = uint8(v1822)
	v1826 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[34]))
	F_WaitLSNWakeup(m, int32(3), v1826)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L72
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1830 = *(*int64)(unsafe.Add(mBase, uint32(v772)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[37])) = v1830
	v1833 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[38])) = v1833
	if v1705 == int32(0) {
		v1880 = v1833
		goto L158
	} else {
		goto L353
	}
L352:
	;
	goto L351
L353:
	;
	v1837 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v776))))
	v1839 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[39])) = uint8(v1839)
	v1841 = int32(_a_F_XLogInsert_20)
	v1843 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[40]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[40])) = v1837 + v1843
	v1846 = int32(_a_F_XLogInsert_21)
	v1848 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[41]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[41])) = v1848 + int64(1)
	v1852 = int32(_a_F_XLogInsert_22)
	v1854 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[42]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[42])) = v1854 + base.I64_extend_i32_s(v565)
	v1858 = int32(_a_F_XLogInsert_23)
	v1860 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[43]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[43])) = v1860 + v535
	v1880 = v1833
	goto L158
L354:
	;
	v1866 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogInsert[36])) = uint8(v1866)
	v1870 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[34]))
	F_WaitLSNWakeup(m, int32(3), v1870)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L72
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1874 = *(*int64)(unsafe.Add(mBase, uint32(v772)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[37])) = v1874
	*(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[38])) = v1790
	v1880 = v1790
	goto L158
L357:
	;
	goto L356
L358:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L72
	} else {
		goto L359
	}
L359:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_24), int32(0))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L72
	} else {
		goto L360
	}
L360:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_17), int32(1408), int32(_a_F_XLogInsert_25))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L72
	} else {
		goto L361
	}
L361:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L362:
	;
	goto L15
L363:
	;
	v1944 = v1940 & int32(7)
	v1946 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	if base.Ui32(int32(8)) <= base.Ui32(v1940) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1967 = v1938
	v1972 = int32(0)
	goto L367
L365:
	;
	v2032 = v1938
	goto L366
L366:
	;
	v2071 = int32(0)
	v2072 = v2032
	goto L371
L367:
	;
	v1993 = v1946 + v1967*int32(_a_F_XLogInsert_3)
	v1994 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[44]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[45]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[46]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[47]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[48]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[49]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_XLogInsert[50]))) = uint8(v1994)
	*(*uint8)(unsafe.Add(mBase, uint32(v1993))) = uint8(v1994)
	v2010 = int32(8)
	v2011 = v1967 + v2010
	v2013 = v1972 + v2010
	if v2013 != v1940&int32(2147483640) {
		v1967 = v2011
		v1972 = v2013
		goto L367
	} else {
		goto L369
	}
L368:
	;
	if v1944 == int32(0) {
		v2350 = v1880
		goto L1
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	v2032 = v2011
	goto L366
L371:
	;
	v2099 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1946+v2072*int32(_a_F_XLogInsert_3)))) = uint8(v2099)
	v2101 = int32(1)
	v2104 = v2071 + v2101
	if v2104 != v1944 {
		v2071 = v2104
		v2072 = v2072 + v2101
		goto L371
	} else {
		goto L373
	}
L372:
	;
	v2350 = v1880
	goto L1
L373:
	;
	goto L372
L374:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_26), int32(0))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L72
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(488), int32(_a_F_XLogInsert_27))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = l1
	F_errmsg_internal(m, int32(_a_F_XLogInsert_28), v42+int32(48))
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L72
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(497), int32(_a_F_XLogInsert_27))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L72
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L380:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_29), int32(0))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L72
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = int32(-1)
	v2145 = *(*int64)(unsafe.Add(mBase, _c_F_XLogInsert[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+32)) = v2145
	F_errdetail_internal(m, int32(_a_F_XLogInsert_30), v42+int32(32))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L72
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(951), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L72
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	F_errmsg_internal(m, int32(_a_F_XLogInsert_31), int32(0))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L72
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v119 & int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(1069547520)
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v662
	F_errdetail_internal(m, int32(_a_F_XLogInsert_32), v42)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L72
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_XLogInsert_8), int32(996), int32(_a_F_XLogInsert_11))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L72
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v2186 = v2182 & int32(7)
	v2188 = *(*int32)(unsafe.Add(mBase, _c_F_XLogInsert[14]))
	if base.Ui32(int32(8)) <= base.Ui32(v2182) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v2208 = v15
	v2213 = v15
	goto L392
L390:
	;
	v2273 = v15
	goto L391
L391:
	;
	v2311 = v15
	v2312 = v2273
	goto L396
L392:
	;
	v2234 = v2188 + v2208*int32(_a_F_XLogInsert_3)
	v2235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[44]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[45]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[46]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[47]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[48]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[49]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+uint32(_c_F_XLogInsert[50]))) = uint8(v2235)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234))) = uint8(v2235)
	v2251 = int32(8)
	v2252 = v2208 + v2251
	v2254 = v2213 + v2251
	if v2254 != v2182&int32(2147483640) {
		v2208 = v2252
		v2213 = v2254
		goto L392
	} else {
		goto L394
	}
L393:
	;
	if v2186 == int32(0) {
		v2350 = v2180
		goto L1
	} else {
		goto L395
	}
L394:
	;
	goto L393
L395:
	;
	v2273 = v2252
	goto L391
L396:
	;
	v2339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2188+v2312*int32(_a_F_XLogInsert_3)))) = uint8(v2339)
	v2341 = int32(1)
	v2344 = v2311 + v2341
	if v2344 != v2186 {
		v2311 = v2344
		v2312 = v2312 + v2341
		goto L396
	} else {
		goto L398
	}
L397:
	;
	v2350 = v2180
	goto L1
L398:
	;
	goto L397
}
func F_XLogPageRead(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v345 int64
	_ = v345
	var v354 int64
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int64
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v531 int64
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
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
	var v677 int32
	_ = v677
	var v681 int64
	_ = v681
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v686 int64
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int64
	_ = v733
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1037 int64
	_ = v1037
	var v1039 int64
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1048 int64
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1131 int32
	_ = v1131
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1153 int64
	_ = v1153
	var v1159 int64
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1168 int64
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int64
	_ = v1179
	var v1180 int64
	_ = v1180
	var v1188 int64
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int64
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1253 int32
	_ = v1253
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int64
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int64
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1307 int64
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1350 int32
	_ = v1350
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1381 int64
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1412 int64
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1435 int32
	_ = v1435
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1468 int64
	_ = v1468
	var v1469 int64
	_ = v1469
	var v1473 int64
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int64
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1501 int64
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1513 int64
	_ = v1513
	var v1514 int64
	_ = v1514
	var v1518 int64
	_ = v1518
	var v1568 int32
	_ = v1568
	var v1569 int64
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1580 int32
	_ = v1580
	var v1585 int64
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1605 int32
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1610 int64
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1624 int32
	_ = v1624
	var v1627 int64
	_ = v1627
	var v1630 int64
	_ = v1630
	var v1631 int64
	_ = v1631
	var v1632 int64
	_ = v1632
	var v1635 int64
	_ = v1635
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1653 int32
	_ = v1653
	var v1658 int64
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1693 int64
	_ = v1693
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1713 int32
	_ = v1713
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1738 int64
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1750 int64
	_ = v1750
	var v1751 int64
	_ = v1751
	var v1755 int64
	_ = v1755
	var v1805 int32
	_ = v1805
	var v1806 int64
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1817 int32
	_ = v1817
	var v1822 int64
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1842 int32
	_ = v1842
	var v1843 int64
	_ = v1843
	var v1847 int64
	_ = v1847
	var v1852 int32
	_ = v1852
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1868 int64
	_ = v1868
	var v1869 int64
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1881 int32
	_ = v1881
	var v1885 int64
	_ = v1885
	var v1887 int64
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1958 int32
	_ = v1958
	var v1968 int32
	_ = v1968
	v4 = l3
	v6 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(1216)
	m.G0 = v33
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v39 = base.I32_wrap_i64(l1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if v43 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v114 = (v36 - int32(1)) & v39
	v117 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v111))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2])) = v117
	v119 = int64(32)
	v121 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(v119) % 64)))
	v123 = l1 + base.I64_extend_i32_s(l2)
	if v112 != 0 {
		v142 = v6
		v143 = int32(0)
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v111 = v36
	v112 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v48 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v50 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v36))
	if v48 == v50 {
		v111 = v36
		v112 = v6
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v52 = int32(1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[3])))
	if v54 != v52 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v100 = int32(_a_F_XLogPageRead_0)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	v102 = F_close(m, v101)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v111 = v110
	v112 = v52
	goto L1
L7:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[5])))
	if v58&int32(1) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[6]))
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[7]))
	v71 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v72 = base.I64_div_u_s(v69, v71)
	goto L9
L9:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v64-int32(1))+v72) <= base.Ui64(v48)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v77 = F_GetRedoRecPtr(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[6]))
	v89 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[7]))
	v91 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v92 = base.I64_div_u_s(v89, v91)
	goto L13
L13:
	;
	if base.B2i32(base.Ui64(base.I64_extend_i32_s(v84-int32(1))+v92) <= base.Ui64(v82)) == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_RequestCheckpoint(m, int32(128))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	m.G0 = v33 + int32(1216)
	return v1968
L17:
	;
	v149 = v143
	v156 = v142
	v166 = v6
	goto L27
L18:
	;
	v129 = int32(_a_F_XLogPageRead_1)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v131 == int32(3) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v142 = v129
	v143 = int32(3)
	goto L17
L20:
	;
	v135 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(v123) <= base.Ui64(v135) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v142 = v129
	v143 = int32(2)
	goto L17
L23:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v137 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v1968 = int32(-2)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v142 = v129
	v143 = int32(1)
	goto L17
L27:
	;
	switch v149 {
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
	v1949 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1949 {
		goto L441
	} else {
		goto L442
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9])) = v114
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v156
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[11])))
	v1462 = m.G0
	v1464 = v1462 - int32(16)
	m.G0 = v1464
	if v1459 != 0 {
		goto L352
	} else {
		goto L353
	}
L31:
	;
	v149 = int32(2)
	v156 = v1435
	goto L27
L32:
	;
	v1412 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(int64(8191)) < base.Ui64(v1412^l1) {
		v1435 = int32(_a_F_XLogPageRead_1)
		goto L31
	} else {
		goto L350
	}
L33:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[12])))
	if v182 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	v149 = int32(1)
	v166 = v174
	goto L27
L35:
	;
	v201 = int32(1)
	v202 = v166 & v201
	v213 = v200
	v227 = int32(0)
	goto L45
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	if v186 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v197 = int32(2)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13])) = v197
	v200 = v197
	goto L35
L39:
	;
	if v186 != int32(3) {
		v200 = v186
		goto L35
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v194)
	v197 = int32(1)
	goto L38
L42:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v190&int32(1) != 0 {
		v200 = v186
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
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L11
	} else {
		goto L347
	}
L45:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])))
	if v240 != 0 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L11
	} else {
		goto L344
	}
L47:
	;
	v474 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v474)
	if base.Ui32(int32(2)) <= base.Ui32(v470-int32(1)) {
		goto L122
	} else {
		goto L123
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13])) = v434
	if v213 == v434 {
		v469 = v432
		v470 = v213
		goto L47
	} else {
		goto L105
	}
L49:
	;
	v432 = v429
	v434 = int32(1)
	goto L48
L50:
	;
	if v202 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v420 = int32(0)
	if v213 != int32(2) {
		v469 = v420
		v470 = v213
		goto L47
	} else {
		goto L103
	}
L53:
	;
	v1968 = int32(-2)
	goto L16
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v213-int32(1)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v277 = F_WalRcvStreaming(m)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L72
	}
L57:
	;
	if v213 == int32(3) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v264 != int32(1) {
		goto L29
	} else {
		goto L64
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v253
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_2), v33)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3743), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
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
	v267 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	if v267 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v272 == int32(0) {
		goto L29
	} else {
		goto L70
	}
L69:
	;
	goto L29
L70:
	;
	v432 = int32(1)
	v434 = int32(3)
	goto L48
L71:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[16]))
	if v299 != int32(1) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v277 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L11
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	v286 = F_LWLockAcquire(m, v282+int32(1152), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L11
	} else {
		goto L77
	}
L76:
	;
	goto L71
L77:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[18]))
	v290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+312)) = uint8(v290)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	F_LWLockRelease(m, v293+int32(1152))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L71
L79:
	;
	v310 = m.G0
	v311 = int32(16)
	v312 = v310 - v311
	m.G0 = v312
	F_gettimeofday(m, v312)
	mBase = m.M
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	v316 = int64(*(*int32)(unsafe.Add(mBase, uint32(v312)+8)))
	m.G0 = v312 + v311
	v324 = v316 + v315*int64(1000000) - int64(946684800000000)
	goto L83
L80:
	;
	v302 = F_rescanLatestTimeLine(m, v177, v176)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	if v302 == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v429 = int32(0)
	goto L49
L83:
	;
	v326 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19]))
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[20]))
	goto L84
L84:
	;
	if base.B2i32(base.I64_extend_i32_s(v328)*int64(1000) <= v324-v326) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[20]))
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19]))
	if v324 <= v339 {
		v357 = int32(0)
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v416 = v324
	goto L87
L87:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[19])) = v416
	v429 = int32(0)
	goto L49
L88:
	;
	v361 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v345 = v324 - v339
	if base.B2i32(int64(0) < v339)^base.B2i32(v345 < v324)|base.B2i32(int64(2147483646000) < v345) != 0 {
		v357 = int32(2147483647)
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v354 = base.I64_div_s(v345+int64(999), int64(1000))
	v357 = base.I32_wrap_i64(v354)
	goto L89
L92:
	;
	if v361 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = base.I32_wrap_i64(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = base.I32_wrap_i64(int64(base.Ui64(v123) >> (uint(v119) % 64)))
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_5), v33+int32(176))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L98
	}
L96:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3722), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v383 = F_WaitLatch(m, v378+int32(4), int32(41), v337-v357, int32(150994949))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v386+int32(4)))) = v389
	v394 = base.AtomicRmwOr32(m, v389, int32(_a_F_XLogPageRead_6), v389)
	goto L100
L100:
	;
	v398 = m.G0
	v399 = int32(16)
	v400 = v398 - v399
	m.G0 = v400
	F_gettimeofday(m, v400)
	mBase = m.M
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v400)))
	v404 = int64(*(*int32)(unsafe.Add(mBase, uint32(v400)+8)))
	m.G0 = v400 + v399
	goto L101
L101:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v416 = v404 + v403*int64(1000000) - int64(946684800000000)
	goto L87
L103:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[12])))
	if v424&int32(1) == int32(0) {
		v469 = v420
		v470 = v213
		goto L47
	} else {
		goto L104
	}
L104:
	;
	v429 = v420
	goto L49
L105:
	;
	v439 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	if v439 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v213<<(uint(int32(2))%32))+uint32(_c_F_XLogPageRead[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v443
	v448 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])))
	if v448 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v469 = v432
	v470 = v468
	goto L47
L110:
	;
	v449 = int32(_a_F_XLogPageRead_7)
	goto L112
L111:
	;
	v449 = int32(_a_F_XLogPageRead_8)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+168)) = v449
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v452<<(uint(int32(2))%32))+uint32(_c_F_XLogPageRead[22])))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+164)) = v455
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_9), v33+int32(160))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3760), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
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
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+80))
	if v1363 != 0 {
		goto L339
	} else {
		goto L340
	}
L117:
	;
	v1272 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L11
	} else {
		goto L324
	}
L118:
	;
	v149 = int32(3)
	v156 = v1253
	goto L27
L119:
	;
	v1145 = F_WalRcvStreaming(m)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L11
	} else {
		goto L300
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23])) = v732
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	v741 = F_LWLockAcquire(m, v737+int32(1152), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L11
	} else {
		goto L191
	}
L121:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	v724 = F_tliOfPointInHistory(m, v4, v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L11
	} else {
		goto L186
	}
L122:
	;
	if v470 != int32(3) {
		goto L44
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v518 {
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[25])))
	if (v469|(v483^int32(-1)))&int32(1) != 0 {
		v498 = v469
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v500 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[25])) = uint8(v500)
	if v498 == v500 {
		goto L119
	} else {
		goto L131
	}
L127:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	v491 = int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[16]))
	if v493 != v491 {
		v498 = v491
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v496 = F_rescanLatestTimeLine(m, v177, v176)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v498 = v491
	goto L126
L131:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[26]))
	if v505 == int32(0) {
		goto L119
	} else {
		goto L132
	}
L132:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	if v508 == int32(0) {
		goto L119
	} else {
		goto L133
	}
L133:
	;
	if v178&v201 == int32(0) {
		goto L121
	} else {
		goto L134
	}
L134:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[27]))
	v516 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[28]))
	v732 = v514
	v733 = v516
	goto L120
L135:
	;
	v521 = F_close(m, v518)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = int32(-1)
	goto L137
L136:
	;
	goto L137
L137:
	;
	if v179&v201 != 0 {
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
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v531 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	if v533 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v677
	v681 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v682 = base.I64_div_u_s(int64(4294967296), v681)
	v683 = base.I64_div_u_s(v531, v682)
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+52)) = uint32(v683)
	v686 = v531 - v682*v683
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+56)) = uint32(v686)
	v689 = v33 + int32(192)
	v694 = F_pg_snprintf(m, v689, int32(1024), int32(_a_F_XLogPageRead_10), v33+int32(48))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L11
	} else {
		goto L178
	}
L142:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	v538 = F_readTimeLineHistory(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L11
	} else {
		goto L145
	}
L143:
	;
	v542 = v533
	goto L144
L144:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v543 <= int32(0) {
		goto L141
	} else {
		goto L147
	}
L145:
	;
	if v538 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v542 = v538
	goto L144
L147:
	;
	if v529 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v549 = v529
	goto L150
L149:
	;
	v549 = int32(0)
	goto L150
L150:
	;
	v558 = int32(0)
	goto L151
L151:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v583+v558<<(uint(int32(2))%32))))
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
	v643 = v558 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v643 < v644 {
		v558 = v643
		goto L151
	} else {
		goto L177
	}
L155:
	;
	v596 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v597 = base.I64_div_u_s(v592, v596)
	if base.Ui64(v531) < base.Ui64(v597) {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if base.Ui32(int32(1)) < base.Ui32(v549) {
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
	if v549&int32(1) != 0 {
		goto L154
	} else {
		goto L170
	}
L161:
	;
	v601 = int32(1)
	v603 = F_XLogFileRead(m, v531, v588, v601, v601)
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
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_11), int32(0))
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
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(_a_F_XLogPageRead_12), int32(_a_F_XLogPageRead_13))
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
	v623 = F_XLogFileRead(m, v531, v588, int32(2), int32(1))
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
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24])) = v542
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
		v1253 = v636
		goto L118
	} else {
		goto L176
	}
L176:
	;
	v1435 = v636
	goto L31
L177:
	;
	goto L152
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30])) = int32(44)
	v701 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L11
	} else {
		goto L179
	}
L179:
	;
	if v701 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
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
	v720 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v720)
	v1350 = v227
	goto L116
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v689
	F_errmsg(m, int32(_a_F_XLogPageRead_14), v33+int32(32))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L11
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(_a_F_XLogPageRead_15), int32(_a_F_XLogPageRead_13))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	if base.Ui32(v724) < base.Ui32(v727) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v730 = v727
	goto L189
L188:
	;
	v730 = int32(0)
	goto L189
L189:
	;
	if v730 != 0 {
		goto L115
	} else {
		goto L190
	}
L190:
	;
	v732 = v724
	v733 = v123
	goto L120
L191:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[18]))
	v745 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v744)+312)) = uint8(v745)
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[17]))
	F_LWLockRelease(m, v748+int32(1152))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[26]))
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[31]))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[32])))
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[33]))
	v761 = F_time(m)
	mBase = m.M
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v766 = base.AtomicRmwXchg32(m, v760, int32(1456), int32(1))
	if v766 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	F_s_lock(m, v760+int32(1456), int32(_a_F_XLogPageRead_16))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L11
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v772 = int32(0)
	if v756 == v772 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L195
L197:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v760)+1452)) = uint8(v903)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v760)+8))
	if v906 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L198:
	;
	v901 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v760)+1388)) = uint8(v901)
	v903 = v758
	goto L197
L199:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	if v776 == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v780 = v760 + int32(1388)
	goto L204
L201:
	;
	v903 = int32(0)
	goto L197
L202:
	;
	v897 = F_strlen(m, v886)
	mBase = m.M
	goto L201
L204:
	;
	goto L205
L205:
	;
	v787 = int32(63)
	if (v780^v756)&int32(3) != 0 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v890 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v887))) = uint8(v890)
	goto L202
L207:
	;
	v871 = v866
	v872 = v867
	v873 = v868
	goto L228
L208:
	;
	if v861 == int32(0) {
		v886 = v859
		v887 = v860
		goto L206
	} else {
		goto L227
	}
L209:
	;
	v859 = v756
	v860 = v780
	v861 = v787
	goto L208
L210:
	;
	goto L211
L211:
	;
	v791 = int32(0)
	if base.B2i32(v756&int32(3) == v791)|int32(0) == v791 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	if v827 == int32(0) {
		v886 = v824
		v887 = v825
		goto L206
	} else {
		goto L221
	}
L213:
	;
	v803 = v756
	v804 = v780
	v805 = v787
	goto L216
L214:
	;
	goto L215
L215:
	;
	v824 = v756
	v825 = v780
	v826 = v787
	v827 = int32(1)
	goto L212
L216:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	*(*uint8)(unsafe.Add(mBase, uint32(v804))) = uint8(v807)
	if v807 == int32(0) {
		v866 = v803
		v867 = v804
		v868 = v805
		goto L207
	} else {
		goto L218
	}
L217:
	;
	v824 = v818
	v825 = v812
	v826 = v814
	v827 = v816
	goto L212
L218:
	;
	v811 = int32(1)
	v812 = v804 + v811
	v814 = v805 - v811
	v815 = int32(0)
	v816 = base.B2i32(v814 != v815)
	v818 = v803 + v811
	if v818&int32(3) == v815 {
		v824 = v818
		v825 = v812
		v826 = v814
		v827 = v816
		goto L212
	} else {
		goto L219
	}
L219:
	;
	if v814 != 0 {
		v803 = v818
		v804 = v812
		v805 = v814
		goto L216
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824))))
	if base.B2i32(v830 == int32(0))|base.B2i32(base.Ui32(v826) < base.Ui32(int32(4))) != 0 {
		v859 = v824
		v860 = v825
		v861 = v826
		goto L208
	} else {
		goto L222
	}
L222:
	;
	v837 = v824
	v838 = v825
	v839 = v826
	goto L223
L223:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v845 = int32(-2139062144)
	if (int32(16843008)-v842|v842)&v845 != v845 {
		v866 = v837
		v867 = v838
		v868 = v839
		goto L207
	} else {
		goto L225
	}
L224:
	;
	v859 = v853
	v860 = v851
	v861 = v855
	goto L208
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = v842
	v850 = int32(4)
	v851 = v838 + v850
	v853 = v837 + v850
	v855 = v839 - v850
	if base.Ui32(int32(3)) < base.Ui32(v855) {
		v837 = v853
		v838 = v851
		v839 = v855
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v866 = v859
	v867 = v860
	v868 = v861
	goto L207
L228:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	*(*uint8)(unsafe.Add(mBase, uint32(v872))) = uint8(v875)
	if v875 == int32(0) {
		v886 = v871
		v887 = v872
		goto L206
	} else {
		goto L230
	}
L229:
	;
	v886 = v882
	v887 = v880
	goto L206
L230:
	;
	v879 = int32(1)
	v880 = v872 + v879
	v882 = v871 + v879
	v884 = v873 - v879
	if v884 != 0 {
		v871 = v882
		v872 = v880
		v873 = v884
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1037 = v733 & base.I64_extend_i32_s(v772-v763)
	*(*int64)(unsafe.Add(mBase, uint32(v760)+24)) = v761
	v1039 = *(*int64)(unsafe.Add(mBase, uint32(v760)+32))
	if v1039 != int64(0) {
		goto L271
	} else {
		goto L272
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = int32(1)
	v912 = v760 + int32(104)
	if v754 != 0 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = int32(5)
	goto L232
L236:
	;
	goto L242
L237:
	;
	goto L238
L238:
	;
	v1032 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v1032)
	goto L232
L239:
	;
	goto L232
L240:
	;
	v1029 = F_strlen(m, v1018)
	mBase = m.M
	goto L239
L242:
	;
	goto L243
L243:
	;
	v919 = int32(1023)
	if (v912^v754)&int32(3) != 0 {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v1022 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019))) = uint8(v1022)
	goto L240
L245:
	;
	v1003 = v998
	v1004 = v999
	v1005 = v1000
	goto L266
L246:
	;
	if v993 == int32(0) {
		v1018 = v991
		v1019 = v992
		goto L244
	} else {
		goto L265
	}
L247:
	;
	v991 = v754
	v992 = v912
	v993 = v919
	goto L246
L248:
	;
	goto L249
L249:
	;
	v923 = int32(0)
	if base.B2i32(v754&int32(3) == v923)|int32(0) == v923 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	if v959 == int32(0) {
		v1018 = v956
		v1019 = v957
		goto L244
	} else {
		goto L259
	}
L251:
	;
	v935 = v754
	v936 = v912
	v937 = v919
	goto L254
L252:
	;
	goto L253
L253:
	;
	v956 = v754
	v957 = v912
	v958 = v919
	v959 = int32(1)
	goto L250
L254:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	*(*uint8)(unsafe.Add(mBase, uint32(v936))) = uint8(v939)
	if v939 == int32(0) {
		v998 = v935
		v999 = v936
		v1000 = v937
		goto L245
	} else {
		goto L256
	}
L255:
	;
	v956 = v950
	v957 = v944
	v958 = v946
	v959 = v948
	goto L250
L256:
	;
	v943 = int32(1)
	v944 = v936 + v943
	v946 = v937 - v943
	v947 = int32(0)
	v948 = base.B2i32(v946 != v947)
	v950 = v935 + v943
	if v950&int32(3) == v947 {
		v956 = v950
		v957 = v944
		v958 = v946
		v959 = v948
		goto L250
	} else {
		goto L257
	}
L257:
	;
	if v946 != 0 {
		v935 = v950
		v936 = v944
		v937 = v946
		goto L254
	} else {
		goto L258
	}
L258:
	;
	goto L255
L259:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	if base.B2i32(v962 == int32(0))|base.B2i32(base.Ui32(v958) < base.Ui32(int32(4))) != 0 {
		v991 = v956
		v992 = v957
		v993 = v958
		goto L246
	} else {
		goto L260
	}
L260:
	;
	v969 = v956
	v970 = v957
	v971 = v958
	goto L261
L261:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	v977 = int32(-2139062144)
	if (int32(16843008)-v974|v974)&v977 != v977 {
		v998 = v969
		v999 = v970
		v1000 = v971
		goto L245
	} else {
		goto L263
	}
L262:
	;
	v991 = v985
	v992 = v983
	v993 = v987
	goto L246
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v970))) = v974
	v982 = int32(4)
	v983 = v970 + v982
	v985 = v969 + v982
	v987 = v971 - v982
	if base.Ui32(int32(3)) < base.Ui32(v987) {
		v969 = v985
		v970 = v983
		v971 = v987
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	v998 = v991
	v999 = v992
	v1000 = v993
	goto L245
L266:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1004))) = uint8(v1007)
	if v1007 == int32(0) {
		v1018 = v1003
		v1019 = v1004
		goto L244
	} else {
		goto L268
	}
L267:
	;
	v1018 = v1014
	v1019 = v1012
	goto L244
L268:
	;
	v1011 = int32(1)
	v1012 = v1004 + v1011
	v1014 = v1003 + v1011
	v1016 = v1005 - v1011
	if v1016 != 0 {
		v1003 = v1014
		v1004 = v1012
		v1005 = v1016
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+40)) = v732
	*(*int64)(unsafe.Add(mBase, uint32(v760)+32)) = v1037
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v1052 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v760)+1456)), uint32(v1052))
	if v906 == v1052 {
		goto L276
	} else {
		goto L277
	}
L271:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v760)+56))
	if v1042 == v732 {
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v760)+64)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v760)+56)) = v732
	*(*int64)(unsafe.Add(mBase, uint32(v760)+48)) = v1037
	v1048 = base.AtomicRmwXchg64(m, v760, int32(1464), v1037)
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
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[5])))
	if v1059 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	goto L278
L278:
	;
	if v1051 != int32(-1) {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	goto L275
L280:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v1063+int32(32)))) = int32(1)
	v1070 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[35]))
	v1072 = F_pgmem_kill(m, v1070, int32(10))
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
	v1076 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[36]))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	v1082 = v1077 + v1051*int32(768) + int32(316)
	v1083 = int32(0)
	v1086 = base.AtomicRmwOr32(m, v1083, int32(_a_F_XLogPageRead_6), v1083)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	if v1087 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1082))) = int32(1)
	v1090 = int32(0)
	v1093 = base.AtomicRmwOr32(m, v1090, int32(_a_F_XLogPageRead_6), v1090)
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1094 == v1090 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	if v1097 == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[37]))
	if v1101 == v1097 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1103 = m.G0
	v1105 = v1103 - int32(16)
	m.G0 = v1105
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[38]))
	if v1108 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	goto L293
L293:
	;
	v1131 = F_pgmem_kill(m, v1097, int32(23))
	mBase = m.M
	goto L287
L294:
	;
	m.G0 = v1105 + int32(16)
	goto L286
L295:
	;
	v1111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1105)+15)) = uint8(v1111)
	goto L296
L296:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[39]))
	v1119 = F_write(m, v1115, v1105+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v1119 {
		goto L294
	} else {
		goto L298
	}
L297:
	;
	goto L294
L298:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30]))
	if v1123 == int32(27) {
		goto L296
	} else {
		goto L299
	}
L299:
	;
	goto L297
L300:
	;
	if v1145 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1150)
	v1350 = v227
	goto L116
L302:
	;
	goto L303
L303:
	;
	v1153 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8]))
	if base.Ui64(v123) < base.Ui64(v1153) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v1236 = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[40])) = v1236
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1236
	v1253 = v156
	goto L118
L305:
	;
	if v202 == int32(0) {
		goto L117
	} else {
		goto L323
	}
L306:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1209 {
		goto L304
	} else {
		goto L317
	}
L307:
	;
	v1159 = F_GetWalRcvFlushRecPtr(m, v33+int32(192), int32(_a_F_XLogPageRead_17))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L11
	} else {
		goto L308
	}
L308:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[8])) = v1159
	if base.Ui64(v1159) <= base.Ui64(v123) {
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[41]))
	v1166 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	if v1164 != v1166 {
		goto L305
	} else {
		goto L310
	}
L310:
	;
	v1168 = *(*int64)(unsafe.Add(mBase, uint32(v33)+192))
	if base.Ui64(v123) < base.Ui64(v1168) {
		goto L306
	} else {
		goto L311
	}
L311:
	;
	v1174 = m.G0
	v1175 = int32(16)
	v1176 = v1174 - v1175
	m.G0 = v1176
	F_gettimeofday(m, v1176)
	mBase = m.M
	v1179 = *(*int64)(unsafe.Add(mBase, uint32(v1176)))
	v1180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1176)+8)))
	m.G0 = v1176 + v1175
	v1188 = v1180 + v1179*int64(1000000) - int64(946684800000000)
	goto L312
L312:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[42])) = v1188
	v1191 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1194 = base.AtomicRmwXchg32(m, v1191, int32(96), int32(1))
	if v1194 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	F_s_lock(m, v1191+int32(96), int32(_a_F_XLogPageRead_18))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L11
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	*(*int64)(unsafe.Add(mBase, uint32(v1201)+72)) = v1188
	v1203 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1201)+96)), uint32(v1203))
	goto L306
L316:
	;
	goto L315
L317:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24]))
	if v1213 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[29]))
	v1219 = F_readTimeLineHistory(m, v1218)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L11
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1224 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[41]))
	v1229 = F_XLogFileRead(m, v1224, v1226, int32(3), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L11
	} else {
		goto L322
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[24])) = v1219
	goto L320
L322:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1229
	v1350 = v227
	goto L116
L323:
	;
	v1968 = int32(-2)
	goto L16
L324:
	;
	if v1272 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1275)
	v1350 = v227
	goto L116
L326:
	;
	goto L327
L327:
	;
	if v227 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_WalRcvRequestApplyReply(m)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
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
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L11
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[43]))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+124))
	if v1290 != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1318 = F_WaitLatch(m, v1312+int32(4), int32(33), int32(-1), int32(83886090))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L11
	} else {
		goto L337
	}
L334:
	;
	v1291 = *(*int64)(unsafe.Add(mBase, uint32(v1290)+16))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+120))
	v1293 = *(*int64)(unsafe.Add(mBase, uint32(v1292)+16))
	v1296 = base.I32_wrap_i64(v1291 - v1293)
	goto L336
L335:
	;
	v1296 = int32(0)
	goto L336
L336:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+112))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+16))
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[44]))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1300)+64)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v1300)+56)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v1300)+60)) = v1301 + v1298
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1307 = *(*int64)(unsafe.Add(mBase, uint32(v1306)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1284)+16)) = v1307 - int64(-8192)
	goto L333
L337:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[21]))
	v1324 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1321+int32(4)))) = v1324
	v1329 = base.AtomicRmwOr32(m, v1324, int32(_a_F_XLogPageRead_6), v1324)
	goto L338
L338:
	;
	v1350 = int32(1)
	goto L116
L339:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
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
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L11
	} else {
		goto L343
	}
L342:
	;
	goto L341
L343:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	v213 = v1370
	v227 = v1350
	goto L45
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+152)) = v724
	v1377 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+156)) = v1377
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+148)) = uint32(v4)
	v1381 = int64(base.Ui64(v4) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+144)) = uint32(v1381)
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_19), v33+int32(144))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L11
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3869), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
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
	v1398 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v1398
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_2), v33+int32(16))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L11
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(4011), int32(_a_F_XLogPageRead_4))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
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
	v1418 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0]))
	v1435 = base.I32_wrap_i64(v1412)&(v1418-int32(1)) - v114
	goto L31
L351:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v1478))) = int32(167772237)
	v1482 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	v1483 = int32(_a_F_XLogPageRead_1)
	v1485 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_XLogPageRead[9])))
	v1486 = F_pread(m, v1482, l4, v1483, v1485)
	mBase = m.M
	if v1486 != v1483 {
		goto L358
	} else {
		goto L359
	}
L352:
	;
	F___clock_gettime(m, int32(1), v1464)
	mBase = m.M
	v1468 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1464)+8)))
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1464)))
	v1473 = v1468 + v1469*int64(1000000000)
	goto L354
L353:
	;
	v1473 = int64(0)
	goto L354
L354:
	;
	m.G0 = v1464 + int32(16)
	goto L351
L355:
	;
	v149 = int32(0)
	goto L27
L356:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10]))
	v1968 = v1946
	goto L16
L357:
	;
	v1923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)))
	if v1923 != 0 {
		goto L434
	} else {
		goto L435
	}
L358:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30]))
	v1492 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	v1493 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1492))) = v1493
	if v1493 < v1486 {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	goto L360
L360:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v1731))) = int32(0)
	v1737 = int32(1)
	v1738 = int64(8192)
	v1742 = m.G0
	v1744 = v1742 - int32(16)
	m.G0 = v1744
	if v1473 != int64(0) {
		goto L404
	} else {
		goto L405
	}
L361:
	;
	v1500 = int32(1)
	v1501 = base.I64_extend_i32_u(v1486)
	v1505 = m.G0
	v1507 = v1505 - int32(16)
	m.G0 = v1507
	if v1473 != int64(0) {
		goto L365
	} else {
		goto L366
	}
L362:
	;
	goto L363
L363:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v1624
	v1627 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[2]))
	v1630 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v1631 = base.I64_div_u_s(int64(4294967296), v1630)
	v1632 = base.I64_div_u_s(v1627, v1631)
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+132)) = uint32(v1632)
	v1635 = v1627 - v1632*v1631
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+136)) = uint32(v1635)
	v1643 = F_pg_snprintf(m, v33+int32(192), int32(64), int32(_a_F_XLogPageRead_20), v33+int32(128))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L11
	} else {
		goto L381
	}
L364:
	;
	goto L363
L365:
	;
	F___clock_gettime(m, int32(1), v1507)
	mBase = m.M
	v1513 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1507)+8)))
	v1514 = *(*int64)(unsafe.Add(mBase, uint32(v1507)))
	v1518 = v1513 + (v1514*int64(1000000000) - v1473)
	goto L368
L366:
	;
	goto L367
L367:
	;
	v1605 = int32(880)
	v1606 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46])) = v1606 + base.I64_extend_i32_u(v1500)
	v1610 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47])) = v1610 + v1501
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1500, v1501)
	mBase = m.M
	v1615 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1615)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[49])) = uint8(v1615)
	m.G0 = v1507 + int32(16)
	goto L364
L368:
	;
	v1568 = int32(880)
	v1569 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50])) = v1569 + v1518
	v1573 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[51]))
	v1580 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1573))|base.B2i32(int32(1)<<(uint(v1573)%32)&int32(_a_F_XLogPageRead_21) == v1580) == v1580 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1585 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52])) = v1585 + v1518
	v1589 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1589)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[53])) = uint8(v1589)
	goto L380
L379:
	;
	goto L380
L380:
	;
	goto L367
L381:
	;
	if v1486 < int32(0) {
		goto L383
	} else {
		goto L384
	}
L382:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), v1726, int32(_a_F_XLogPageRead_22))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L11
	} else {
		goto L402
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[30])) = v1490
	if v41 != int32(15) {
		v1663 = v41
		goto L386
	} else {
		goto L387
	}
L384:
	;
	goto L385
L385:
	;
	if v41 != int32(15) {
		v1698 = v41
		goto L394
	} else {
		goto L395
	}
L386:
	;
	v1665 = F_errstart(m, v1663, int32(0))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L11
	} else {
		goto L390
	}
L387:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1653 != int32(2) {
		v1663 = v41
		goto L386
	} else {
		goto L388
	}
L388:
	;
	v1658 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v123 == v1658 {
		v1663 = int32(14)
		goto L386
	} else {
		goto L389
	}
L389:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v123
	v1663 = int32(15)
	goto L386
L390:
	;
	if v1665 == int32(0) {
		goto L357
	} else {
		goto L391
	}
L391:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L11
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v39
	v1674 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1674
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v33 + int32(192)
	F_errmsg(m, int32(_a_F_XLogPageRead_23), v33+int32(80))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L11
	} else {
		goto L393
	}
L393:
	;
	v1726 = int32(3415)
	goto L382
L394:
	;
	v1700 = F_errstart(m, v1698, int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L11
	} else {
		goto L398
	}
L395:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1688 != int32(2) {
		v1698 = v41
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1693 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v123 == v1693 {
		v1698 = int32(14)
		goto L394
	} else {
		goto L397
	}
L397:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v123
	v1698 = int32(15)
	goto L394
L398:
	;
	if v1700 == int32(0) {
		goto L357
	} else {
		goto L399
	}
L399:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = int32(_a_F_XLogPageRead_1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v39
	v1713 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1713
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v33 + int32(192)
	F_errmsg(m, int32(_a_F_XLogPageRead_24), v33+int32(96))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L11
	} else {
		goto L401
	}
L401:
	;
	v1726 = int32(3422)
	goto L382
L402:
	;
	goto L357
L403:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[23]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1184)) = v1861
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v1864 != int32(1) {
		goto L356
	} else {
		goto L420
	}
L404:
	;
	F___clock_gettime(m, int32(1), v1744)
	mBase = m.M
	v1750 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1744)+8)))
	v1751 = *(*int64)(unsafe.Add(mBase, uint32(v1744)))
	v1755 = v1750 + (v1751*int64(1000000000) - v1473)
	goto L407
L405:
	;
	goto L406
L406:
	;
	v1842 = int32(880)
	v1843 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[46])) = v1843 + base.I64_extend_i32_u(v1737)
	v1847 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[47])) = v1847 + v1738
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(6), v1737, v1738)
	mBase = m.M
	v1852 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1852)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[49])) = uint8(v1852)
	m.G0 = v1744 + int32(16)
	goto L403
L407:
	;
	v1805 = int32(880)
	v1806 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[50])) = v1806 + v1755
	v1810 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[51]))
	v1817 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1810))|base.B2i32(int32(1)<<(uint(v1810)%32)&int32(_a_F_XLogPageRead_21) == v1817) == v1817 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1822 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[52])) = v1822 + v1755
	v1826 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[48])) = uint8(v1826)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[53])) = uint8(v1826)
	goto L419
L418:
	;
	goto L419
L419:
	;
	goto L406
L420:
	;
	v1868 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[0])))
	v1869 = base.I64_rem_u_s(l1, v1868)
	if v1869 != int64(0) {
		goto L356
	} else {
		goto L421
	}
L421:
	;
	v1872 = F_XLogReaderValidatePageHeader(m, l0, l1, l4)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L11
	} else {
		goto L422
	}
L422:
	;
	if v1872 != 0 {
		goto L356
	} else {
		goto L423
	}
L423:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1874))))
	if v1875 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v1914 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1913))) = uint8(v1914)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v1914)
	goto L357
L425:
	;
	if v41 != int32(15) {
		v1892 = v41
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1895 = F_errstart(m, v1892, int32(0))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L11
	} else {
		goto L430
	}
L427:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4]))
	if v1881 != int32(2) {
		v1892 = v41
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v1885 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1887 = *(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54]))
	if v1885 == v1887 {
		v1892 = int32(14)
		goto L426
	} else {
		goto L429
	}
L429:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogPageRead[54])) = v1885
	v1892 = int32(15)
	goto L426
L430:
	;
	if v1895 == int32(0) {
		goto L424
	} else {
		goto L431
	}
L431:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v1899
	F_errmsg_internal(m, int32(_a_F_XLogPageRead_25), v33-int32(-64))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L11
	} else {
		goto L432
	}
L432:
	;
	F_errfinish(m, int32(_a_F_XLogPageRead_3), int32(3478), int32(_a_F_XLogPageRead_22))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L11
	} else {
		goto L433
	}
L433:
	;
	goto L424
L434:
	;
	v1968 = int32(-2)
	goto L16
L435:
	;
	goto L436
L436:
	;
	v1926 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[14])) = uint8(v1926)
	v1929 = *(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1]))
	if int32(0) <= v1929 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1932 = F_close(m, v1929)
	mBase = m.M
	goto L439
L438:
	;
	goto L439
L439:
	;
	v1933 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1933
	v1938 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v1938
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1938
	v1944 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogPageRead[15])))
	if v1944 != 0 {
		goto L355
	} else {
		goto L440
	}
L440:
	;
	v1968 = v1933
	goto L16
L441:
	;
	v1952 = F_close(m, v1949)
	mBase = m.M
	goto L443
L442:
	;
	goto L443
L443:
	;
	v1953 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[1])) = v1953
	v1958 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[10])) = v1958
	*(*int32)(unsafe.Add(mBase, _c_F_XLogPageRead[4])) = v1958
	v1968 = v1953
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int64
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int64
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v284 int64
	_ = v284
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int64
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int64
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v407 int64
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int64
	_ = v453
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v490 int64
	_ = v490
	var v491 int64
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
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
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v712 int32
	_ = v712
	var v741 int32
	_ = v741
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v776 int32
	_ = v776
	var v778 int64
	_ = v778
	var v782 int64
	_ = v782
	var v788 int32
	_ = v788
	var v814 int64
	_ = v814
	var v818 int64
	_ = v818
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v858 int64
	_ = v858
	var v861 int64
	_ = v861
	var v867 int32
	_ = v867
	var v870 int64
	_ = v870
	var v874 int64
	_ = v874
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v944 int64
	_ = v944
	var v952 int64
	_ = v952
	var v958 int32
	_ = v958
	var v962 int64
	_ = v962
	var v969 int64
	_ = v969
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v983 int64
	_ = v983
	var v988 int64
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1000 int64
	_ = v1000
	var v1004 int64
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1025 int32
	_ = v1025
	var v1027 int64
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1037 int64
	_ = v1037
	var v1040 int64
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int64
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1203 int32
	_ = v1203
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1237 int32
	_ = v1237
	var v1281 int64
	_ = v1281
	var v1284 int64
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1313 int32
	_ = v1313
	var v1340 int32
	_ = v1340
	var v1344 int64
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1385 int64
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	v2 = l1
	v3 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(_a_F_XLogReadAhead_0)
	m.G0 = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)))
	if v29 != 0 {
		v1430 = v3
		v1436 = v27
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v1436 + int32(_a_F_XLogReadAhead_0)
	return v1430
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v31)
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v42 = v38 & int64(-8192)
	v43 = int32(_a_F_XLogReadAhead_1)
	v44 = base.I32_wrap_i64(v38)
	v46 = v44 & int32(_a_F_XLogReadAhead_2)
	if base.Ui32(v43) <= base.Ui32(v46) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v49 = v43
	goto L5
L4:
	;
	v49 = v46
	goto L5
L5:
	;
	v52 = F_ReadPageInternal(m, l0, v42, v49+int32(24))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v52 == int32(-2) {
		v1430 = v3
		v1436 = v27
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v61 = v52
	v63 = v46
	v64 = v3
	v68 = v44
	v80 = v38
	v81 = v42
	goto L12
L9:
	;
	if v1396 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L10:
	;
	v1386 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1362)+1256)) = uint8(v1386)
	*(*int64)(unsafe.Add(mBase, uint32(v1362)+56)) = v1385
	*(*int64)(unsafe.Add(mBase, uint32(v1362)+48)) = v108
	v1390 = v1362
	v1396 = v1368
	v1398 = v1370
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v491
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+17)))
	if v493 != 0 {
		goto L145
	} else {
		goto L146
	}
L12:
	;
	if v61 < int32(0) {
		v1390 = l0
		v1396 = v64
		v1398 = v27
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v427 = int32(_a_F_XLogReadAhead_3)
	v428 = v107 + v128
	if base.Ui32(v427) <= base.Ui32(v428) {
		goto L135
	} else {
		goto L136
	}
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+2)))
	if v87&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v90 = int32(40)
	goto L17
L16:
	;
	v90 = int32(24)
	goto L17
L17:
	;
	if v63 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v111 = int32(0)
	if base.B2i32(v87&int32(1) == v111)|base.B2i32(v90 != v107) == v111 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v107 = v90
	v108 = v80 + base.I64_extend_i32_u(v90)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v90) <= base.Ui32(v63) {
		v107 = v63
		v108 = v80
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+124)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v27)+116)) = v68
	v100 = int64(base.Ui64(v80) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+112)) = uint32(v100)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_4), v27+int32(112))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v1390 = l0
	v1396 = v64
	v1398 = v27
	goto L9
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+4)) = uint32(v108)
	v119 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27))) = uint32(v119)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_5), v27)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v124 = base.I32_wrap_i64(v108)
	v126 = v124 & int32(_a_F_XLogReadAhead_2)
	v127 = v86 + v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if base.Ui32(v107) <= base.Ui32(int32(_a_F_XLogReadAhead_1)) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v1390 = l0
	v1396 = v64
	v1398 = v27
	goto L9
L28:
	;
	v153 = v128 + int32(2037)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v154 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v134 = F_ValidXLogRecordHeader(m, l0, v108, v131, v127, base.B2i32(v40 == int64(0)))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(23)) < base.Ui32(v128) {
		goto L28
	} else {
		goto L34
	}
L32:
	;
	if v134 == int32(0) {
		v1390 = l0
		v1396 = v64
		v1398 = v27
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v124
	v145 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+96)) = uint32(v145)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_6), v27+int32(96))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v1390 = l0
	v1396 = v64
	v1398 = v27
	goto L9
L36:
	;
	v200 = int32(_a_F_XLogReadAhead_3) - v126
	v201 = base.B2i32(base.Ui32(v128) <= base.Ui32(v200))
	if v201 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v194 = int32(0)
	if v2 != 0 {
		v1430 = v194
		v1436 = v27
		goto L1
	} else {
		goto L52
	}
L38:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)) = uint8(v186)
	v195 = v183
	v198 = base.B2i32(v183 == v186)
	goto L36
L39:
	;
	if base.Ui32(v170-v169) <= base.Ui32(v153) {
		goto L37
	} else {
		goto L51
	}
L40:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v153) <= base.Ui32(v175-v172+v173) {
		v183 = v172
		goto L38
	} else {
		goto L49
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v157 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v169) < base.Ui32(v170) {
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v161 = v157
	goto L46
L45:
	;
	v158 = int32(_a_F_XLogReadAhead_7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v158
	v161 = v158
	goto L46
L46:
	;
	v162 = F_palloc(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v162
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v167)
	v172 = v162
	v173 = v162
	v174 = v162
	goto L40
L48:
	;
	v172 = v169
	v173 = v154
	v174 = v170
	goto L40
L49:
	;
	if base.Ui32(v153) < base.Ui32(v174-v173) {
		v183 = v173
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	v183 = v169
	goto L38
L52:
	;
	v195 = v194
	v198 = int32(1)
	goto L36
L53:
	;
	if v200 != 0 {
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
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	base.MemoryCopy(m, v204, v205+v126, v200)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v210 = int32(_a_F_XLogReadAhead_8)
	v211 = int32(-8192)
	v214 = v128&v211 - v211
	if base.Ui32(v214) <= base.Ui32(v210) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = v210
	goto L61
L60:
	;
	v217 = v214
	goto L61
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v225 = v200
	v232 = base.B2i32(base.Ui32(v107) < base.Ui32(int32(_a_F_XLogReadAhead_9)))
	v236 = v220 + v200
	v245 = v81
	goto L63
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1257)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1216)) = v248
	v413 = int32(_a_F_XLogReadAhead_1)
	v414 = base.I32_wrap_i64(v248)
	v416 = v414 & int32(_a_F_XLogReadAhead_2)
	if base.Ui32(v413) <= base.Ui32(v416) {
		goto L130
	} else {
		goto L131
	}
L63:
	;
	v246 = int32(0)
	v248 = v245 - int64(-8192)
	v250 = F_ReadPageInternal(m, l0, v248, int32(24))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L65
	}
L64:
	;
	v371 = int32(-1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v373 = int32(24)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v378 = m.Env.Pgmem_crc32c(m, v371, v372+v373, v375-v373)
	mBase = m.M
	v380 = m.Env.Pgmem_crc32c(m, v378, v372, int32(20))
	mBase = m.M
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v372)+20))
	if v380^v381 != v371 {
		goto L123
	} else {
		goto L124
	}
L65:
	;
	if v250 == int32(-2) {
		v1430 = v246
		v1436 = v27
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if v250 < int32(0) {
		v1362 = l0
		v1368 = v195
		v1370 = v27
		v1385 = v248
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+2)))
	if v257&int32(4) != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v257&int32(1) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v124
	v266 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+16)) = uint32(v266)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_10), v27+int32(16))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	if v128 == v225+v273 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v1362 = l0
	v1368 = v195
	v1370 = v27
	v1385 = v248
	goto L10
L73:
	;
	v277 = v273
	goto L75
L74:
	;
	v277 = int32(0)
	goto L75
L75:
	;
	if v277 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v124
	v284 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27-int32(-64)))) = uint32(v284)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v273
	*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = base.I64_extend_i32_u(v128) - base.I64_extend_i32_u(v225)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_11), v27+int32(48))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v296 = int32(_a_F_XLogReadAhead_3)
	v297 = v128 + int32(24) - v225
	if base.Ui32(v296) <= base.Ui32(v297) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v1362 = l0
	v1368 = v195
	v1370 = v27
	v1385 = v248
	goto L10
L80:
	;
	v300 = v296
	goto L82
L81:
	;
	v300 = v297
	goto L82
L82:
	;
	v301 = F_ReadPageInternal(m, l0, v248, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	if v301 == int32(-2) {
		v1430 = v246
		v1436 = v27
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v301 < int32(0) {
		v1362 = l0
		v1368 = v195
		v1370 = v27
		v1385 = v248
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+2)))
	if v309&int32(2) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v312 = int32(40)
	goto L88
L87:
	;
	v312 = int32(24)
	goto L88
L88:
	;
	if base.Ui32(v301) < base.Ui32(v312) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v314 = F_ReadPageInternal(m, l0, v248, v312)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L92
	}
L90:
	;
	v316 = v301
	goto L91
L91:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v320 = int32(_a_F_XLogReadAhead_3) - v312
	if base.Ui32(v318) < base.Ui32(v320) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v316 = v314
	goto L91
L93:
	;
	v322 = v318
	goto L95
L94:
	;
	v322 = v320
	goto L95
L95:
	;
	v323 = v322 + v312
	if base.Ui32(v316) < base.Ui32(v323) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v325 = F_ReadPageInternal(m, l0, v248, v323)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v322 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	base.MemoryCopy(m, v236, v312+v317, v322)
	goto L102
L101:
	;
	goto L102
L102:
	;
	if v232&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v337 = F_ValidXLogRecordHeader(m, l0, v108, v333, v334, base.B2i32(v40 == int64(0)))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v341 = v225 + v322
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1248))
	if base.Ui32(v128) <= base.Ui32(v342) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	if v337 == int32(0) {
		v1362 = l0
		v1368 = v195
		v1370 = v27
		v1385 = v248
		goto L10
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v368 = v322 + v236
	goto L110
L109:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v346 = int32(0)
	v347 = base.B2i32(v341 == v346)
	if v347 == v346 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if base.Ui32(v341) < base.Ui32(v128) {
		v225 = v341
		v232 = int32(1)
		v236 = v368
		v245 = v248
		goto L63
	} else {
		goto L122
	}
L111:
	;
	base.MemoryCopy(m, v27+int32(128), v345, v341)
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v345 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_pfree(m, v345)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v355 = F_palloc(m, v217)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1248)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1244)) = v355
	if v347 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	base.MemoryCopy(m, v355, v27+int32(128), v341)
	goto L121
L120:
	;
	goto L121
L121:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1244))
	v368 = v364 + v341
	goto L110
L122:
	;
	goto L64
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v124
	v387 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+32)) = uint32(v387)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_12), v27+int32(32))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v394)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v108
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	if v395&int32(2) != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v1362 = l0
	v1368 = v195
	v1370 = v27
	v1385 = v248
	goto L10
L127:
	;
	v407 = int64(40)
	goto L129
L128:
	;
	v407 = int64(24)
	goto L129
L129:
	;
	v471 = v372
	v490 = v248
	v491 = base.I64_extend_i32_u((v397+int32(7))&int32(-8)) + (v407 + v248)
	goto L11
L130:
	;
	v419 = v413
	goto L132
L131:
	;
	v419 = v416
	goto L132
L132:
	;
	v422 = F_ReadPageInternal(m, l0, v248, v419+int32(24))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	if v422 != int32(-2) {
		v61 = v422
		v63 = v416
		v64 = v195
		v68 = v414
		v80 = v248
		v81 = v248
		goto L12
	} else {
		goto L134
	}
L134:
	;
	v1430 = v246
	v1436 = v27
	goto L1
L135:
	;
	v431 = v427
	goto L137
L136:
	;
	v431 = v428
	goto L137
L137:
	;
	v432 = F_ReadPageInternal(m, l0, v81, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	if v432 == int32(-2) {
		v1430 = int32(0)
		v1436 = v27
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v432 < int32(0) {
		v1390 = l0
		v1396 = v195
		v1398 = v27
		goto L9
	} else {
		goto L140
	}
L140:
	;
	v438 = int32(-1)
	v439 = int32(24)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v444 = m.Env.Pgmem_crc32c(m, v438, v127+v439, v441-v439)
	mBase = m.M
	v446 = m.Env.Pgmem_crc32c(m, v444, v127, int32(20))
	mBase = m.M
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v446^v447 != v438 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v124
	v453 = int64(base.Ui64(v108) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+80)) = uint32(v453)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_12), v27+int32(80))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v108
	v471 = v127
	v490 = v81
	v491 = v108 + base.I64_extend_i32_u((v128+int32(7))&int32(-8))
	goto L11
L144:
	;
	v1390 = l0
	v1396 = v195
	v1398 = v27
	goto L9
L145:
	;
	if v198 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+16)))
	if v494&int32(240) != int32(64) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = (v491 + base.I64_extend_i32_s(v499-int32(1))) & base.I64_extend_i32_s(int32(0)-v499)
	goto L145
L148:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v510 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L149:
	;
	v560 = v195
	goto L150
L150:
	;
	v564 = int32(0)
	v566 = m.G0
	v568 = v566 - int32(176)
	m.G0 = v568
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v471)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v560)+48)) = v570
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v471)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v560)+40)) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v471)))
	*(*int64)(unsafe.Add(mBase, uint32(v560)+32)) = v574
	*(*int64)(unsafe.Add(mBase, uint32(v560)+16)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v560)+68)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v560)+60)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+56)) = uint16(v564)
	*(*int32)(unsafe.Add(mBase, uint32(v560)+8)) = v564
	v586 = v560 + int32(76)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v589 = v587 - int32(24)
	if v589 == v564 {
		v1237 = v586
		goto L172
	} else {
		goto L173
	}
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+4)) = uint8(v553)
	v560 = v554
	goto L150
L152:
	;
	v549 = F_palloc(m, v153)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L168
	}
L153:
	;
	if base.Ui32(v153) < base.Ui32(v526-v525) {
		v553 = int32(0)
		v554 = v525
		goto L151
	} else {
		goto L167
	}
L154:
	;
	v533 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v153) <= base.Ui32(v534-v531+v532) {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v513 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v525) < base.Ui32(v526) {
		goto L153
	} else {
		goto L162
	}
L158:
	;
	v517 = v513
	goto L160
L159:
	;
	v514 = int32(_a_F_XLogReadAhead_7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v514
	v517 = v514
	goto L160
L160:
	;
	v518 = F_palloc(m, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v518
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v523)
	v529 = v518
	v531 = v518
	v532 = v518
	goto L154
L162:
	;
	v529 = v526
	v531 = v525
	v532 = v510
	goto L154
L163:
	;
	v553 = v533
	v554 = v531
	goto L151
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v529-v532) <= base.Ui32(v153) {
		goto L152
	} else {
		goto L166
	}
L166:
	;
	v553 = v533
	v554 = v532
	goto L151
L167:
	;
	goto L152
L168:
	;
	v553 = int32(1)
	v554 = v549
	goto L151
L169:
	;
	m.G0 = v568 + int32(176)
	if v1340 != 0 {
		goto L283
	} else {
		goto L284
	}
L170:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(128)))) = v1313
	v1340 = int32(0)
	goto L169
L171:
	;
	v1281 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+4)) = uint32(v1281)
	v1284 = int64(base.Ui64(v1281) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568))) = uint32(v1284)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_13), v568)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L6
	} else {
		goto L282
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = (v1237 - v560 + int32(7)) & int32(-8)
	v1340 = int32(1)
	goto L169
L173:
	;
	v597 = int32(-1)
	v600 = v471 + int32(24)
	v602 = v589
	v605 = v564
	v607 = v564
	goto L175
L174:
	;
	if v1096 != v1101 {
		goto L171
	} else {
		goto L259
	}
L175:
	;
	v620 = v602 - int32(1)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	switch v621 - int32(252) {
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
	v1091 = v1064
	v1093 = int32(0)
	v1094 = v1086
	v1096 = v1069
	v1101 = v1074
	goto L174
L177:
	;
	if base.Ui32(v1074) < base.Ui32(v1069) {
		v597 = v1064
		v600 = v1086
		v602 = v1069
		v605 = v1072
		v607 = v1074
		goto L175
	} else {
		goto L258
	}
L178:
	;
	if base.Ui32(v621) <= base.Ui32(int32(32)) {
		goto L188
	} else {
		goto L189
	}
L179:
	;
	if base.Ui32(v602) < base.Ui32(int32(5)) {
		goto L171
	} else {
		goto L186
	}
L180:
	;
	if base.Ui32(v602) < base.Ui32(int32(3)) {
		goto L171
	} else {
		goto L185
	}
L181:
	;
	if base.Ui32(v602) < base.Ui32(int32(5)) {
		goto L171
	} else {
		goto L184
	}
L182:
	;
	if v620 == int32(0) {
		goto L171
	} else {
		goto L183
	}
L183:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+68)) = v626
	v628 = int32(2)
	v1091 = v597
	v1093 = v626
	v1094 = v600 + v628
	v1096 = v602 - v628
	v1101 = v626 + v607
	goto L174
L184:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v600)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+68)) = v635
	v637 = int32(5)
	v1091 = v597
	v1093 = v635
	v1094 = v600 + v637
	v1096 = v602 - v637
	v1101 = v635 + v607
	goto L174
L185:
	;
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600)+1)))
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+56)) = uint16(v644)
	v646 = int32(3)
	v1064 = v597
	v1069 = v602 - v646
	v1072 = v605
	v1074 = v607
	v1086 = v600 + v646
	goto L177
L186:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v600)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+60)) = v652
	v654 = int32(5)
	v1064 = v597
	v1069 = v602 - v654
	v1072 = v605
	v1074 = v607
	v1086 = v600 + v654
	goto L177
L187:
	;
	if v621 <= v597 {
		goto L203
	} else {
		goto L204
	}
L188:
	;
	v661 = v597 + int32(1)
	if v621 <= v661 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v778 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+168)) = uint32(v778)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+160)) = v621
	v782 = int64(base.Ui64(v778) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+164)) = uint32(v782)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_14), v568+int32(160))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L202
	}
L191:
	;
	v668 = (v597 ^ int32(-1) + v621) & int32(7)
	if v668 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v673 = int32(0)
	v678 = v661
	goto L195
L193:
	;
	v712 = v661
	goto L194
L194:
	;
	if base.Ui32(v621-v597-int32(2)) <= base.Ui32(int32(6)) {
		goto L187
	} else {
		goto L198
	}
L195:
	;
	v696 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v586+v678*int32(52)))) = uint8(v696)
	v698 = int32(1)
	v699 = v678 + v698
	v701 = v673 + v698
	if v701 != v668 {
		v673 = v701
		v678 = v699
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v712 = v699
	goto L194
L197:
	;
	goto L196
L198:
	;
	v741 = v712
	goto L199
L199:
	;
	v758 = v586 + v741*int32(52)
	v759 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v758))) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+364)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+312)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+260)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+208)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+156)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+104)) = uint8(v759)
	*(*uint8)(unsafe.Add(mBase, uint32(v758)+52)) = uint8(v759)
	v776 = v741 + int32(8)
	if v776 != v621 {
		v741 = v776
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
	v814 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+152)) = uint32(v814)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+144)) = v621
	v818 = int64(base.Ui64(v814) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+148)) = uint32(v818)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_15), v568+int32(144))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L6
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560)+72)) = v621
	v828 = v586 + v621*int32(52)
	v829 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+30)) = uint8(v829)
	v831 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v828))) = uint8(v831)
	if v620 == v829 {
		goto L171
	} else {
		goto L207
	}
L206:
	;
	goto L170
L207:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+28)) = uint8(v835)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+16)) = v835 & int32(15)
	v844 = int32(1)
	v845 = int32(base.Ui32(v835)>>(uint(int32(5))%32)) & v844
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+43)) = uint8(v845)
	v850 = int32(base.Ui32(v835)>>(uint(int32(4))%32)) & v844
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+29)) = uint8(v850)
	v853 = v602 & int32(-2)
	if v853 == int32(2) {
		goto L171
	} else {
		goto L208
	}
L208:
	;
	v856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+48)) = uint16(v856)
	if v845 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v881 = int32(4)
	v882 = v602 - v881
	v884 = v600 + v881
	v885 = v856 + v607
	if v850 == int32(0) {
		v1013 = v884
		v1014 = v885
		v1015 = v882
		goto L217
	} else {
		goto L218
	}
L210:
	;
	if v856 != 0 {
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	if v856 == int32(0) {
		goto L209
	} else {
		goto L215
	}
L213:
	;
	v858 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+20)) = uint32(v858)
	v861 = int64(base.Ui64(v858) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+16)) = uint32(v861)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_16), v568+int32(16))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L6
	} else {
		goto L214
	}
L214:
	;
	goto L170
L215:
	;
	v870 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+136)) = uint32(v870)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+128)) = v856
	v874 = int64(base.Ui64(v870) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+132)) = uint32(v874)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_17), v568+int32(128))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	goto L170
L217:
	;
	if int32(0) <= base.I32_extend8_s(v835) {
		goto L249
	} else {
		goto L250
	}
L218:
	;
	if base.Ui32(v882) < base.Ui32(int32(2)) {
		goto L171
	} else {
		goto L219
	}
L219:
	;
	v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884))))
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+40)) = uint16(v890)
	if v853 == int32(6) {
		goto L171
	} else {
		goto L220
	}
L220:
	;
	v894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+36)) = uint16(v894)
	if v602 == int32(8) {
		goto L171
	} else {
		goto L221
	}
L221:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+42)) = uint8(v898)
	v900 = int32(1)
	v903 = int32(base.Ui32(v898)>>(uint(v900)%32)) & v900
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+30)) = uint8(v903)
	v905 = int32(9)
	v906 = v602 - v905
	v910 = v898 & int32(28)
	if v910 != 0 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v898&int32(1) != 0 {
		goto L232
	} else {
		goto L233
	}
L223:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+38)) = uint16(v924)
	v926 = v600 + v905
	v927 = v906
	v928 = v924
	goto L222
L224:
	;
	if v898&int32(1) != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	goto L226
L226:
	;
	v924 = int32(_a_F_XLogReadAhead_3) - v890
	goto L223
L227:
	;
	if base.Ui32(v906) < base.Ui32(int32(2)) {
		goto L171
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v924 = int32(0)
	goto L223
L230:
	;
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600)+9)))
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+38)) = uint16(v915)
	v917 = int32(11)
	v926 = v600 + v917
	v927 = v602 - v917
	v928 = v915
	goto L222
L231:
	;
	if v898&int32(29)|v995 != 0 {
		v1013 = v926
		v1014 = v890 + v885
		v1015 = v927
		goto L217
	} else {
		goto L246
	}
L232:
	;
	v931 = int32(0)
	if base.B2i32(v894 == v931)|base.B2i32(v928&int32(_a_F_XLogReadAhead_18) == v931) == v931 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	if (v928|v894)&int32(_a_F_XLogReadAhead_18) != 0 {
		goto L240
	} else {
		goto L241
	}
L235:
	;
	if v890 != int32(_a_F_XLogReadAhead_3) {
		v995 = int32(0)
		goto L231
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v944 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+112)) = uint32(v944)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+104)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v568)+100)) = v928 & int32(_a_F_XLogReadAhead_18)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+96)) = v894
	v952 = int64(base.Ui64(v944) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+108)) = uint32(v952)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_19), v568+int32(96))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
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
	v962 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+92)) = uint32(v962)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+84)) = v928 & int32(_a_F_XLogReadAhead_18)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+80)) = v894
	v969 = int64(base.Ui64(v962) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+88)) = uint32(v969)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_20), v568+int32(80))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L6
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v976 = int32(_a_F_XLogReadAhead_3)
	if base.B2i32(v910 == int32(0))|base.B2i32(v890 != v976) != 0 {
		v995 = base.B2i32(v890 == v976)
		goto L231
	} else {
		goto L244
	}
L243:
	;
	goto L170
L244:
	;
	v983 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+40)) = uint32(v983)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+32)) = int32(_a_F_XLogReadAhead_3)
	v988 = int64(base.Ui64(v983) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+36)) = uint32(v988)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_21), v568+int32(32))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L6
	} else {
		goto L245
	}
L245:
	;
	goto L170
L246:
	;
	v1000 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+72)) = uint32(v1000)
	*(*int32)(unsafe.Add(mBase, uint32(v568)+64)) = v890
	v1004 = int64(base.Ui64(v1000) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+68)) = uint32(v1004)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_22), v568-int32(-64))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L6
	} else {
		goto L247
	}
L247:
	;
	goto L170
L248:
	;
	if base.Ui32(v1053) < base.Ui32(int32(4)) {
		goto L171
	} else {
		goto L257
	}
L249:
	;
	if base.Ui32(v1015) < base.Ui32(int32(12)) {
		goto L171
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	if v605 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+12)) = v1025
	v1027 = *(*int64)(unsafe.Add(mBase, uint32(v1013)))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+4)) = v1027
	v1029 = int32(12)
	v1051 = v1013 + v1029
	v1052 = v828 + int32(4)
	v1053 = v1015 - v1029
	goto L248
L253:
	;
	v1037 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+52)) = uint32(v1037)
	v1040 = int64(base.Ui64(v1037) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v568)+48)) = uint32(v1040)
	F_report_invalid_record(m, l0, int32(_a_F_XLogReadAhead_23), v568+int32(48))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L6
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v605)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+12)) = v1047
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v605)))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+4)) = v1049
	v1051 = v1013
	v1052 = v605
	v1053 = v1015
	goto L248
L256:
	;
	goto L170
L257:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1051)))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+20)) = v1056
	v1058 = int32(4)
	v1064 = v621
	v1069 = v1053 - v1058
	v1072 = v1052
	v1074 = v1014
	v1086 = v1051 + v1058
	goto L177
L258:
	;
	goto L176
L259:
	;
	v1115 = v560 + int32(76)
	v1116 = int32(52)
	v1120 = v1115 + v1091*v1116 + v1116
	v1121 = int32(0)
	if v1121 <= v1091 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1129 = int32(0)
	v1130 = v1094
	v1134 = v1121
	v1136 = v1120
	goto L263
L261:
	;
	v1196 = v1093
	v1197 = v1094
	v1203 = v1120
	goto L262
L262:
	;
	if v1196 == int32(0) {
		v1237 = v1203
		goto L172
	} else {
		goto L278
	}
L263:
	;
	v1151 = v1115 + v1129*int32(52)
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	if v1152 != int32(1) {
		v1183 = v1130
		v1184 = v1136
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v560)+68))
	v1196 = v1191
	v1197 = v1183
	v1203 = v1184
	goto L262
L265:
	;
	v1186 = v1134 + int32(1)
	v1188 = v1186 & int32(255)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v560)+72))
	if v1188 <= v1189 {
		v1129 = v1188
		v1130 = v1183
		v1134 = v1186
		v1136 = v1184
		goto L263
	} else {
		goto L277
	}
L266:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+29)))
	if v1155 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+32)) = v1136
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1151)+40)))
	if v1159 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1165 = v1130
	v1166 = v1136
	goto L269
L269:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+43)))
	if v1167 != int32(1) {
		v1183 = v1165
		v1184 = v1166
		goto L265
	} else {
		goto L273
	}
L270:
	;
	base.MemoryCopy(m, v1136, v1130, v1159)
	goto L272
L271:
	;
	goto L272
L272:
	;
	v1161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1151)+40)))
	v1165 = v1161 + v1130
	v1166 = v1136 + v1161
	goto L269
L273:
	;
	v1173 = (v1166 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v1151)+44)) = v1173
	v1175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1151)+48)))
	if v1175 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	base.MemoryCopy(m, v1173, v1165, v1175)
	goto L276
L275:
	;
	goto L276
L276:
	;
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1151)+48)))
	v1183 = v1177 + v1165
	v1184 = v1173 + v1177
	goto L265
L277:
	;
	goto L264
L278:
	;
	v1221 = (v1203 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v560)+64)) = v1221
	if v1196 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	base.MemoryCopy(m, v1221, v1197, v1196)
	goto L281
L280:
	;
	goto L281
L281:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v560)+68))
	v1237 = v1221 + v1224
	goto L172
L282:
	;
	goto L170
L283:
	;
	v1344 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v560)+24)) = v1344
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+4)))
	if v1346 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	if base.Ui32(v128) <= base.Ui32(v200) {
		v1390 = l0
		v1396 = v560
		v1398 = v27
		goto L9
	} else {
		goto L298
	}
L286:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v1349 != v560 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	goto L288
L288:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v1357 != 0 {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v1352 = v1351
	goto L291
L290:
	;
	v1352 = v1349
	goto L291
L291:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1352 + v1353
	goto L288
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1357)+8)) = v560
	goto L294
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v560
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1360 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1430 = v560
	v1436 = v27
	goto L1
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v560
	v1430 = v560
	v1436 = v27
	goto L1
L298:
	;
	v1362 = l0
	v1368 = v560
	v1370 = v27
	v1385 = v490
	goto L10
L299:
	;
	v1421 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1390)+1192)) = v1421
	*(*int64)(unsafe.Add(mBase, uint32(v1390)+1176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1390)+132)) = v1421
	v1430 = v1421
	v1436 = v1398
	goto L1
L300:
	;
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396)+4)))
	if v1416 != int32(1) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	F_pfree(m, v1396)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
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
		if base.Ui32(int32(8)) <= base.Ui32(v60) {
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
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
							v79 = v34
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							if base.Ui64(v79) <= base.Ui64(v83) {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v89 == int32(0) {
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
							if v103 != 0 {
								F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
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
									v77 = v48
									*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
									v79 = v77
									v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
									v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
									if base.Ui64(v79) <= base.Ui64(v83) {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
										v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v93 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									} else {
										v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
										if v89 == int32(0) {
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
											if v93 == int32(0) {
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
											}
										}
									}
									v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
									v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
									if v103 != 0 {
										F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
											*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
											v112 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
											m.G0 = v6 + int32(16)
											return
										}
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
										v112 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
										m.G0 = v6 + int32(16)
										return
									}
								}
							} else {
								v50 = int32(0)
								v52 = int32(_a_F_XLogSendLogical_1)
								v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
								v54 = int64(0)
								v57 = base.AtomicRmwCmpxchg64(m, v53, int32(272), v54, v54)
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
								v62 = base.AtomicRmwOr32(m, v50, int32(_a_F_XLogSendLogical_2), v50)
								v65 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
								v69 = base.AtomicRmwCmpxchg64(m, v65, int32(264), v54, v54)
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v69
								v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
								v77 = v76
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
								v79 = v77
								v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
								if base.Ui64(v79) <= base.Ui64(v83) {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v89 == int32(0) {
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v93 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
								if v103 != 0 {
									F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
										v112 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
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
								v77 = v48
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
								v79 = v77
								v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
								if base.Ui64(v79) <= base.Ui64(v83) {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v89 == int32(0) {
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v93 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
								if v103 != 0 {
									F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
										v112 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v50 = int32(0)
							v52 = int32(_a_F_XLogSendLogical_1)
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v54 = int64(0)
							v57 = base.AtomicRmwCmpxchg64(m, v53, int32(272), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
							v62 = base.AtomicRmwOr32(m, v50, int32(_a_F_XLogSendLogical_2), v50)
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v69 = base.AtomicRmwCmpxchg64(m, v65, int32(264), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v69
							v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
							v77 = v76
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
							v79 = v77
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							if base.Ui64(v79) <= base.Ui64(v83) {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v89 == int32(0) {
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
							if v103 != 0 {
								F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
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
						v79 = v34
						v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
						if base.Ui64(v79) <= base.Ui64(v83) {
							v86 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
							v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
							if v93 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
							}
						} else {
							v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
							if v89 == int32(0) {
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							}
						}
						v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
						v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
						if v103 != 0 {
							F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
							*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
							v112 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
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
								v77 = v48
								*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
								v79 = v77
								v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
								if base.Ui64(v79) <= base.Ui64(v83) {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								} else {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
									if v89 == int32(0) {
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
										if v93 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
										}
									}
								}
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
								v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
								if v103 != 0 {
									F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
										v112 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v50 = int32(0)
							v52 = int32(_a_F_XLogSendLogical_1)
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v54 = int64(0)
							v57 = base.AtomicRmwCmpxchg64(m, v53, int32(272), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
							v62 = base.AtomicRmwOr32(m, v50, int32(_a_F_XLogSendLogical_2), v50)
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
							v69 = base.AtomicRmwCmpxchg64(m, v65, int32(264), v54, v54)
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v69
							v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
							v77 = v76
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
							v79 = v77
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							if base.Ui64(v79) <= base.Ui64(v83) {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v89 == int32(0) {
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
							if v103 != 0 {
								F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
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
							v77 = v48
							*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
							v79 = v77
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
							if base.Ui64(v79) <= base.Ui64(v83) {
								v86 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							} else {
								v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
								if v89 == int32(0) {
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
									if v93 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
									}
								}
							}
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
							v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
							if v103 != 0 {
								F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
									*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
									v112 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
								m.G0 = v6 + int32(16)
								return
							}
						}
					} else {
						v50 = int32(0)
						v52 = int32(_a_F_XLogSendLogical_1)
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
						v54 = int64(0)
						v57 = base.AtomicRmwCmpxchg64(m, v53, int32(272), v54, v54)
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9])) = v57
						v62 = base.AtomicRmwOr32(m, v50, int32(_a_F_XLogSendLogical_2), v50)
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[8]))
						v69 = base.AtomicRmwCmpxchg64(m, v65, int32(264), v54, v54)
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[10])) = v69
						v76 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[9]))
						v77 = v76
						*(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[3])) = v77
						v79 = v77
						v81 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[1]))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+40))
						if base.Ui64(v79) <= base.Ui64(v83) {
							v86 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])) = uint8(v86)
							v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
							if v93 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
							}
						} else {
							v89 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendLogical[0])))
							if v89 == int32(0) {
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[4]))
								if v93 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[5])) = int32(1)
								}
							}
						}
						v100 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendLogical[6]))
						v103 = base.AtomicRmwXchg32(m, v100, int32(76), int32(1))
						if v103 != 0 {
							F_s_lock(m, v100+int32(76), int32(_a_F_XLogSendLogical_0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
								v112 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v110 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendLogical[2]))
							*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v110
							v112 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v100)+76)), uint32(v112))
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v122
				F_errmsg_internal(m, int32(_a_F_XLogSendLogical_3), v6)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XLogSendLogical_4), int32(3705), int32(_a_F_XLogSendLogical_5))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int64
	_ = v271
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v280 int64
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v336 int64
	_ = v336
	var v338 int64
	_ = v338
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v387 int64
	_ = v387
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v396 int64
	_ = v396
	var v420 int32
	_ = v420
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
	var v441 int64
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v457 int64
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v492 int64
	_ = v492
	var v493 int64
	_ = v493
	var v496 int64
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v512 int64
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int64
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int64
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int64
	_ = v537
	var v540 int64
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int64
	_ = v564
	var v567 int64
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int64
	_ = v615
	var v616 int64
	_ = v616
	var v620 int64
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int64
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int64
	_ = v718
	var v719 int64
	_ = v719
	var v727 int64
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v741 int64
	_ = v741
	var v744 int64
	_ = v744
	var v746 int64
	_ = v746
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int64
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int64
	_ = v815
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[1])))
	if v45 != 0 {
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
	F_s_lock(m, v26+int32(76), int32(_a_F_XLogSendPhysical_0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v40 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+76)), uint32(v40))
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
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v47)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])))
	if v50 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v150 = m.G0
	v151 = int32(16)
	v152 = v150 - v151
	m.G0 = v152
	F_gettimeofday(m, v152)
	mBase = m.M
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
	v156 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+8)))
	m.G0 = v152 + v151
	goto L45
L14:
	;
	v54 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	v146 = v54
	goto L13
L15:
	;
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])))
	if v56 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v126 = F_readTimeLineHistory(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L42
	}
L18:
	;
	v62 = F_GetWalRcvFlushRecPtr(m, int32(0), v19+int32(88))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v98 = int32(0)
	v100 = int32(_a_F_XLogSendPhysical_1)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v102 = int64(0)
	v105 = base.AtomicRmwCmpxchg64(m, v101, int32(272), v102, v102)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[8])) = v105
	v110 = base.AtomicRmwOr32(m, v98, int32(_a_F_XLogSendPhysical_2), v98)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v117 = base.AtomicRmwCmpxchg64(m, v113, int32(264), v102, v102)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[9])) = v117
	goto L40
L21:
	;
	v66 = F_GetXLogReplayRecPtr(m, v19+int32(32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])))
	if v72 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+308))
	v80 = base.B2i32(v78 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])) = uint8(v80)
	v82 = v80
	goto L26
L25:
	;
	v82 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+300))
	goto L30
L28:
	;
	goto L29
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[11]))
	if v92 != v69 {
		v125 = v69
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])) = uint8(v89)
	v125 = v87
	goto L17
L31:
	;
	if base.Ui64(v66) < base.Ui64(v62) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = v62
	goto L34
L33:
	;
	v95 = v66
	goto L34
L34:
	;
	if v69 == v68 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v97 = v95
	goto L37
L36:
	;
	v97 = v66
	goto L37
L37:
	;
	v146 = v97
	goto L13
L38:
	;
	v146 = v124
	goto L13
L40:
	;
	goto L41
L41:
	;
	v124 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[8]))
	goto L38
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[11]))
	v132 = F_tliSwitchPoint(m, v130, v126, int32(_a_F_XLogSendPhysical_3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5])) = v132
	F_list_free_deep(m, v126)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])) = uint8(v138)
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	v146 = v141
	goto L13
L45:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[12])))
	if v166 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v233 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[4])))
	if v235 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[14]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	if v171 == v146 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v146
	v175 = v170 + int32(8)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[15])))
	v180 = base.I32_rem_s(v176+int32(1), int32(_a_F_XLogSendPhysical_4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[16])))
	if v180 == v181 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v185 = v175 + v180<<(uint(int32(4))%32)
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v185)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[17]))) = v186
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[18]))) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[16]))) = int32(-1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[19])))
	if v193 == v180 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v197 = v175 + v180<<(uint(int32(4))%32)
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v197)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[20]))) = v198
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v197)))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[21]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[19]))) = int32(-1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[22])))
	if v205 == v180 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v209 = v175 + v180<<(uint(int32(4))%32)
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v209)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[23]))) = v210
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v209)))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[24]))) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[22]))) = int32(-1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v217 = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v175+v176<<(uint(v217)%32)))) = v146
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[15])))
	*(*int64)(unsafe.Add(mBase, uint32(v170+v221<<(uint(v217)%32))+16)) = v156 + v155*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+uint32(_c_F_XLogSendPhysical[15]))) = v180
	goto L46
L58:
	;
	if base.Ui64(v146) <= base.Ui64(v233) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	v239 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	if base.Ui64(v233) < base.Ui64(v239) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+1168))
	if int32(0) <= v243 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+1168))
	v247 = F_close(m, v246)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v242)+1168)) = int32(-1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	v251 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[26]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	m.T0[v255].(func(*base.Module, int32, int32, int32))(m, int32(99), v251, v251)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v259)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[1])) = uint8(v259)
	v266 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	if v266 == int32(0) {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v271 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[5]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+4)) = uint32(v271)
	v274 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+12)) = uint32(v274)
	v276 = int64(32)
	v277 = int64(base.Ui64(v271) >> (uint(v276) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19))) = uint32(v277)
	v280 = int64(base.Ui64(v274) >> (uint(v276) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+8)) = uint32(v280)
	F_errmsg_internal(m, int32(_a_F_XLogSendPhysical_5), v19)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_XLogSendPhysical_6), int32(3530), int32(_a_F_XLogSendPhysical_7))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	goto L9
L70:
	;
	v292 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v292)
	goto L9
L71:
	;
	goto L72
L72:
	;
	v296 = v233 + int64(131072)
	v297 = base.B2i32(base.Ui64(v146) <= base.Ui64(v296))
	v300 = v297 & (v235 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[3])) = uint8(v300)
	v302 = int32(_a_F_XLogSendPhysical_8)
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v303))) = uint8(v304)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[28])) = v304
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v304
	goto L73
L73:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogSendPhysical_8), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v314 = int32(_a_F_XLogSendPhysical_9)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v316 = int32(_a_F_XLogSendPhysical_8)
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v319 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v315+v317))) = uint8(v319)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v315 + int32(1)
	F_enlargeStringInfo(m, v316, int32(8))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v329 = int32(_a_F_XLogSendPhysical_9)
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v331 = int32(_a_F_XLogSendPhysical_8)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v334 = int64(56)
	v336 = int64(65280)
	v338 = int64(40)
	v341 = int64(16711680)
	v343 = int64(24)
	v345 = int64(4278190080)
	v347 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v330+v332))) = v233<<(uint(v334)%64) | v233&v336<<(uint(v338)%64) | (v233&v341<<(uint(v343)%64) | v233&v345<<(uint(v347)%64)) | (int64(base.Ui64(v233)>>(uint(v347)%64))&v345 | int64(base.Ui64(v233)>>(uint(v343)%64))&v341 | (int64(base.Ui64(v233)>>(uint(v338)%64))&v336 | int64(base.Ui64(v233)>>(uint(v334)%64))))
	v371 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v330 + v371
	F_enlargeStringInfo(m, v331, v371)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v378 = int32(_a_F_XLogSendPhysical_9)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v380 = int32(_a_F_XLogSendPhysical_8)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v383 = int64(56)
	v385 = int64(65280)
	v387 = int64(40)
	v390 = int64(16711680)
	v392 = int64(24)
	v394 = int64(4278190080)
	v396 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v379+v381))) = v146<<(uint(v383)%64) | v146&v385<<(uint(v387)%64) | (v146&v390<<(uint(v392)%64) | v146&v394<<(uint(v396)%64)) | (int64(base.Ui64(v146)>>(uint(v396)%64))&v394 | int64(base.Ui64(v146)>>(uint(v392)%64))&v390 | (int64(base.Ui64(v146)>>(uint(v387)%64))&v385 | int64(base.Ui64(v146)>>(uint(v383)%64))))
	v420 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v379 + v420
	F_enlargeStringInfo(m, v380, v420)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v427 = int32(_a_F_XLogSendPhysical_9)
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v429 = int32(_a_F_XLogSendPhysical_8)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v428+v430))) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v428 + int32(8)
	if base.Ui64(v146) <= base.Ui64(v296) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v441 = v146
	goto L80
L79:
	;
	v441 = v296 & int64(-8192)
	goto L80
L80:
	;
	v443 = base.I32_wrap_i64(v441 - v233)
	F_enlargeStringInfo(m, v429, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v450 = v443
	v457 = v233
	goto L82
L82:
	;
	v462 = int32(_a_F_XLogSendPhysical_9)
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v467 = v464 + v466
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+1184))
	v471 = m.G0
	v473 = v471 - int32(16)
	m.G0 = v473
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])))
	if v478 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v692 = int32(_a_F_XLogSendPhysical_9)
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v695 = v694 + v636
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v695
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v700 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v698+v695))) = uint8(v700)
	v702 = int32(_a_F_XLogSendPhysical_10)
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v700)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[31])) = v700
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32])) = v700
	goto L124
L84:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v632 = v606 + v631
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29])) = v632
	v635 = v457 + base.I64_extend_i32_u(v606)
	v636 = v450 - v606
	if v636 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L7
	} else {
		goto L106
	}
L86:
	;
	m.G0 = v473 + int32(16)
	goto L84
L87:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v476)+308))
	v484 = base.B2i32(v482 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[10])) = uint8(v484)
	if v482 != int32(2) {
		v606 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v476)+300))
	if v470 != v489 {
		v606 = int32(0)
		goto L86
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v492 = v457 + base.I64_extend_i32_u(v450)
	v493 = int64(0)
	v496 = base.AtomicRmwCmpxchg64(m, v476, int32(256), v493, v493)
	if base.Ui64(v496) < base.Ui64(v492) {
		goto L85
	} else {
		goto L92
	}
L92:
	;
	if v450 == int32(0) {
		v573 = v467
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v606 = v573 - v467
	goto L86
L94:
	;
	v500 = v467
	v501 = v450
	v512 = v457
	goto L95
L95:
	;
	v519 = base.I32_wrap_i64(v512) & int32(_a_F_XLogSendPhysical_11)
	v520 = int32(_a_F_XLogSendPhysical_4) - v519
	v522 = v512 + base.I64_extend_i32_u(v520)
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+292))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524)+296))
	v532 = base.I64_rem_u_s(int64(base.Ui64(v512)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v528+int32(1)))
	v533 = base.I32_wrap_i64(v532)
	v535 = v533 << (uint(int32(3)) % 32)
	v537 = int64(0)
	v540 = base.AtomicRmwCmpxchg64(m, v525+v535, int32(0), v537, v537)
	if v522 != v540 {
		v573 = v500
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v573 = v569
	goto L93
L97:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+288))
	v545 = int32(0)
	v548 = base.AtomicRmwOr32(m, v545, int32(_a_F_XLogSendPhysical_2), v545)
	if base.Ui32(v501) < base.Ui32(v520) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v550 = v501
	goto L100
L99:
	;
	v550 = v520
	goto L100
L100:
	;
	if v550 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v500, v544+v533<<(uint(int32(13))%32)+v519, v550)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v556 = int32(0)
	v559 = base.AtomicRmwOr32(m, v556, int32(_a_F_XLogSendPhysical_2), v556)
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[7]))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+292))
	v564 = int64(0)
	v567 = base.AtomicRmwCmpxchg64(m, v562+v535, v556, v564, v564)
	if v522 != v567 {
		v573 = v500
		goto L93
	} else {
		goto L104
	}
L104:
	;
	v569 = v500 + v550
	v572 = v501 - v550
	if v572 != 0 {
		v500 = v569
		v501 = v572
		v512 = v512 + base.I64_extend_i32_u(v550)
		goto L95
	} else {
		goto L105
	}
L105:
	;
	goto L96
L106:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v473)+12)) = uint32(v496)
	v615 = int64(32)
	v616 = int64(base.Ui64(v496) >> (uint(v615) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v473)+8)) = uint32(v616)
	*(*uint32)(unsafe.Add(mBase, uint32(v473)+4)) = uint32(v492)
	v620 = int64(base.Ui64(v492) >> (uint(v615) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v473))) = uint32(v620)
	F_errmsg(m, int32(_a_F_XLogSendPhysical_12), v473)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_XLogSendPhysical_13), int32(1814), int32(_a_F_XLogSendPhysical_14))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
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
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v655 = int64(*(*int32)(unsafe.Add(mBase, uint32(v654)+1160)))
	v656 = base.I64_div_u_s(v635, v655)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v654)+1184))
	F_CheckXLogRemoved(m, v656, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L7
	} else {
		goto L114
	}
L110:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v640)+1184))
	v646 = v19 + int32(88)
	v647 = F_WALRead(m, v640, v642+v632, v635, v636, v644, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	if v647 != 0 {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	F_WALReadRaiseError(m, v646)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[6])))
	if v661 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	goto L83
L116:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[2]))
	v668 = base.AtomicRmwXchg32(m, v665, int32(76), int32(1))
	if v668 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_s_lock(m, v665+int32(76), int32(_a_F_XLogSendPhysical_0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665)+16)))
	v675 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v665)+16)) = uint8(v675)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v665)+76)), uint32(v675))
	if v674 != int32(1) {
		goto L115
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[25]))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+1168))
	if v684 < int32(0) {
		goto L115
	} else {
		goto L122
	}
L122:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v683)+1168))
	v688 = F_close(m, v687)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v683)+1168)) = int32(-1)
	goto L123
L123:
	;
	v450 = v636
	v457 = v635
	goto L82
L124:
	;
	v713 = m.G0
	v714 = int32(16)
	v715 = v713 - v714
	m.G0 = v715
	F_gettimeofday(m, v715)
	mBase = m.M
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v715)))
	v719 = int64(*(*int32)(unsafe.Add(mBase, uint32(v715)+8)))
	m.G0 = v715 + v714
	v727 = v719 + v718*int64(1000000) - int64(946684800000000)
	goto L125
L125:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogSendPhysical_10), int32(8))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v732 = int32(_a_F_XLogSendPhysical_15)
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32]))
	v734 = int32(_a_F_XLogSendPhysical_10)
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	v737 = int64(56)
	v739 = int64(65280)
	v741 = int64(40)
	v744 = int64(16711680)
	v746 = int64(24)
	v748 = int64(4278190080)
	v750 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v733+v735))) = v727<<(uint(v737)%64) | v727&v739<<(uint(v741)%64) | (v727&v744<<(uint(v746)%64) | v727&v748<<(uint(v750)%64)) | (int64(base.Ui64(v727)>>(uint(v750)%64))&v748 | int64(base.Ui64(v727)>>(uint(v746)%64))&v744 | (int64(base.Ui64(v727)>>(uint(v741)%64))&v739 | int64(base.Ui64(v727)>>(uint(v737)%64))))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[32])) = v733 + int32(8)
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[27]))
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[30]))
	v781 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
	*(*int64)(unsafe.Add(mBase, uint32(v778)+17)) = v781
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[29]))
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[26]))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)+20))
	m.T0[v788].(func(*base.Module, int32, int32, int32))(m, int32(100), v778, v785)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13])) = v441
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSendPhysical[2]))
	v797 = base.AtomicRmwXchg32(m, v794, int32(76), int32(1))
	if v797 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_s_lock(m, v794+int32(76), int32(_a_F_XLogSendPhysical_0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L7
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v804 = *(*int64)(unsafe.Add(mBase, _c_F_XLogSendPhysical[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v794)+8)) = v804
	v806 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v794)+76)), uint32(v806))
	v810 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogSendPhysical[33])))
	if v810 != int32(1) {
		goto L9
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+20)) = uint32(v804)
	v815 = int64(base.Ui64(v804) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+16)) = uint32(v815)
	v818 = v19 + int32(32)
	v823 = F_pg_snprintf(m, v818, int32(50), int32(_a_F_XLogSendPhysical_16), v19+int32(16))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v825 = F_strlen(m, v818)
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v8 = base.AtomicRmwXchg32(m, v5, int32(440), int32(1))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v5+int32(440), int32(_a_F_XLogSetAsyncXactLSN_0))
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)+208))
	if base.Ui64(l0) <= base.Ui64(v16) {
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
	v18 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15)+440)), uint32(v18))
	return
L7:
	;
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+208)) = l0
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+313)))
	v23 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15)+440)), uint32(v23))
	if v22 != 0 {
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+64))
	if v60 == int32(-1) {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v27 = int64(0)
	v30 = base.AtomicRmwCmpxchg64(m, v15, int32(272), v27, v27)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[2])) = v30
	v32 = int32(0)
	v35 = base.AtomicRmwOr32(m, v32, int32(_a_F_XLogSetAsyncXactLSN_1), v32)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[0]))
	v42 = base.AtomicRmwCmpxchg64(m, v38, int32(264), v27, v27)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[3])) = v42
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[4]))
	if v45 == v32 {
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
	v68 = v63 + v60*int32(768) + int32(316)
	v69 = int32(0)
	v72 = base.AtomicRmwOr32(m, v69, int32(_a_F_XLogSetAsyncXactLSN_2), v69)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v73 != 0 {
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
	v76 = int32(0)
	v79 = base.AtomicRmwOr32(m, v76, int32(_a_F_XLogSetAsyncXactLSN_2), v76)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v80 == v76 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v83 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[5]))
	if v87 == v83 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v89 = m.G0
	v91 = v89 - int32(16)
	m.G0 = v91
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[6]))
	if v94 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v117 = F_pgmem_kill(m, v83, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v91 + int32(16)
	goto L15
L24:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+15)) = uint8(v97)
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[7]))
	v105 = F_write(m, v101, v91+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v105 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_XLogSetAsyncXactLSN[8]))
	if v109 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
}
func F_XLogWalRcvSendReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v37 int64
	_ = v37
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v83 int64
	_ = v83
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v371 int32
	_ = v371
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v380 int64
	_ = v380
	var v383 int64
	_ = v383
	var v386 int64
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	v2 = l1
	v4 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0]))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v23 = m.G0
	v24 = int32(16)
	v25 = v23 - v24
	m.G0 = v25
	F_gettimeofday(m, v25)
	mBase = m.M
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
	m.G0 = v25 + v24
	v37 = v29 + v28*int64(1000000) - int64(946684800000000)
	goto L6
L5:
	;
	goto L4
L6:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3])) = v66
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5])) = v68
	v77 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[0])))
	if v77 <= int64(0) {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[1]))
	if v37 < v41 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
	v46 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
	v66 = v46
	v67 = v4
	v68 = v44
	goto L7
L12:
	;
	goto L11
L13:
	;
	v48 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v50 = v4
	goto L15
L15:
	;
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[2]))
	v54 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[4]))
	if v54 != v56 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return
L17:
	;
	v50 = v48
	goto L15
L18:
	;
	v66 = v56
	v67 = v50
	v68 = v52
	goto L7
L19:
	;
	goto L20
L20:
	;
	v59 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5]))
	if v59 != v52 {
		v66 = v54
		v67 = v50
		v68 = v52
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
	if v64 == v50 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v66 = v54
	v67 = v50
	v68 = v52
	goto L7
L24:
	;
	v83 = int64(9223372036854775807)
	goto L26
L25:
	;
	v83 = v77*int64(1000000) + v37
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[1])) = v83
	if v67 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	v91 = v67
	goto L29
L29:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6])) = v91
	v93 = int32(_a_F_XLogWalRcvSendReply_0)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v95)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[8])) = v95
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v95
	goto L31
L30:
	;
	v91 = v89
	goto L29
L31:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(_a_F_XLogWalRcvSendReply_0)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v107 = int32(_a_F_XLogWalRcvSendReply_1)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v110 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v106+v108))) = uint8(v110)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v114 + int32(1)
	v119 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
	F_enlargeStringInfo(m, v105, int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v124 = int32(_a_F_XLogWalRcvSendReply_0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v126 = int32(_a_F_XLogWalRcvSendReply_1)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v129 = int64(56)
	v131 = int64(65280)
	v133 = int64(40)
	v136 = int64(16711680)
	v138 = int64(24)
	v140 = int64(4278190080)
	v142 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v125+v127))) = v119<<(uint(v129)%64) | v119&v131<<(uint(v133)%64) | (v119&v136<<(uint(v138)%64) | v119&v140<<(uint(v142)%64)) | (int64(base.Ui64(v119)>>(uint(v142)%64))&v140 | int64(base.Ui64(v119)>>(uint(v138)%64))&v136 | (int64(base.Ui64(v119)>>(uint(v133)%64))&v131 | int64(base.Ui64(v119)>>(uint(v129)%64))))
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v168 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v167 + v168
	v172 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5]))
	F_enlargeStringInfo(m, v124, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v177 = int32(_a_F_XLogWalRcvSendReply_0)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v179 = int32(_a_F_XLogWalRcvSendReply_1)
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v182 = int64(56)
	v184 = int64(65280)
	v186 = int64(40)
	v189 = int64(16711680)
	v191 = int64(24)
	v193 = int64(4278190080)
	v195 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v178+v180))) = v172<<(uint(v182)%64) | v172&v184<<(uint(v186)%64) | (v172&v189<<(uint(v191)%64) | v172&v193<<(uint(v195)%64)) | (int64(base.Ui64(v172)>>(uint(v195)%64))&v193 | int64(base.Ui64(v172)>>(uint(v191)%64))&v189 | (int64(base.Ui64(v172)>>(uint(v186)%64))&v184 | int64(base.Ui64(v172)>>(uint(v182)%64))))
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v221 = int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v220 + v221
	v225 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
	F_enlargeStringInfo(m, v177, v221)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v232 = int32(_a_F_XLogWalRcvSendReply_1)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v235 = int64(56)
	v237 = int64(65280)
	v239 = int64(40)
	v242 = int64(16711680)
	v244 = int64(24)
	v246 = int64(4278190080)
	v248 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v231+v233))) = v225<<(uint(v235)%64) | v225&v237<<(uint(v239)%64) | (v225&v242<<(uint(v244)%64) | v225&v246<<(uint(v248)%64)) | (int64(base.Ui64(v225)>>(uint(v248)%64))&v246 | int64(base.Ui64(v225)>>(uint(v244)%64))&v242 | (int64(base.Ui64(v225)>>(uint(v239)%64))&v237 | int64(base.Ui64(v225)>>(uint(v235)%64))))
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v273 + int32(8)
	v280 = m.G0
	v281 = int32(16)
	v282 = v280 - v281
	m.G0 = v282
	F_gettimeofday(m, v282)
	mBase = m.M
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
	v286 = int64(*(*int32)(unsafe.Add(mBase, uint32(v282)+8)))
	m.G0 = v282 + v281
	v294 = v286 + v285*int64(1000000) - int64(946684800000000)
	goto L36
L36:
	;
	F_enlargeStringInfo(m, int32(_a_F_XLogWalRcvSendReply_0), int32(8))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	v299 = int32(_a_F_XLogWalRcvSendReply_0)
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v301 = int32(_a_F_XLogWalRcvSendReply_1)
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v304 = int64(56)
	v306 = int64(65280)
	v308 = int64(40)
	v311 = int64(16711680)
	v313 = int64(24)
	v315 = int64(4278190080)
	v317 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v300+v302))) = v294<<(uint(v304)%64) | v294&v306<<(uint(v308)%64) | (v294&v311<<(uint(v313)%64) | v294&v315<<(uint(v317)%64)) | (int64(base.Ui64(v294)>>(uint(v317)%64))&v315 | int64(base.Ui64(v294)>>(uint(v313)%64))&v311 | (int64(base.Ui64(v294)>>(uint(v308)%64))&v306 | int64(base.Ui64(v294)>>(uint(v304)%64))))
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v342 + int32(8)
	F_enlargeStringInfo(m, v299, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v352 = int32(_a_F_XLogWalRcvSendReply_1)
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	*(*uint8)(unsafe.Add(mBase, uint32(v351+v353))) = uint8(v2)
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9])) = v358 + int32(1)
	v364 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	if v364 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v367 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[6]))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v367)
	if v2 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[10]))
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[7]))
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[9]))
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[11]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+44))
	m.T0[v406].(func(*base.Module, int32, int32, int32))(m, v399, v401, v403)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L16
	} else {
		goto L48
	}
L43:
	;
	v371 = int32(_a_F_XLogWalRcvSendReply_2)
	goto L45
L44:
	;
	v371 = int32(_a_F_XLogWalRcvSendReply_3)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v371
	v373 = int64(32)
	v374 = int64(base.Ui64(v367) >> (uint(v373) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+16)) = uint32(v374)
	v377 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[3]))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v377)
	v380 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWalRcvSendReply[5]))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v380)
	v383 = int64(base.Ui64(v377) >> (uint(v373) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v383)
	v386 = int64(base.Ui64(v380) >> (uint(v373) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)) = uint32(v386)
	F_errmsg_internal(m, int32(_a_F_XLogWalRcvSendReply_4), v12)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_XLogWalRcvSendReply_5), int32(1273), int32(_a_F_XLogWalRcvSendReply_6))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	goto L1
}
func F_XLogWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v286 int64
	_ = v286
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v353 int64
	_ = v353
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v378 int64
	_ = v378
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v419 int64
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v432 int64
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int64
	_ = v450
	var v452 int32
	_ = v452
	var v455 int64
	_ = v455
	var v458 int32
	_ = v458
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v469 int64
	_ = v469
	var v471 int64
	_ = v471
	var v472 int64
	_ = v472
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v487 int32
	_ = v487
	var v491 int64
	_ = v491
	var v495 int64
	_ = v495
	var v496 int32
	_ = v496
	var v498 int64
	_ = v498
	var v500 int32
	_ = v500
	var v505 int64
	_ = v505
	var v506 int64
	_ = v506
	var v511 int32
	_ = v511
	var v517 int64
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v562 int64
	_ = v562
	var v563 int64
	_ = v563
	var v566 int64
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int64
	_ = v579
	var v583 int64
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int64
	_ = v592
	var v594 int32
	_ = v594
	var v595 int64
	_ = v595
	var v596 int32
	_ = v596
	var v604 int64
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int64
	_ = v619
	var v621 int32
	_ = v621
	var v623 int64
	_ = v623
	var v625 int64
	_ = v625
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v648 int64
	_ = v648
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v659 int32
	_ = v659
	var v663 int64
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int64
	_ = v671
	var v673 int64
	_ = v673
	var v675 int64
	_ = v675
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v688 int64
	_ = v688
	var v690 int64
	_ = v690
	var v691 int64
	_ = v691
	var v695 int64
	_ = v695
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	v4 = int64(0)
	v9 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(96)
	m.G0 = v20
	v22 = int32(_a_F_XLogWrite_0)
	v23 = int32(_a_F_XLogWrite_1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v28 = base.AtomicRmwCmpxchg64(m, v24, int32(272), v4, v4)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v28
	v33 = base.AtomicRmwOr32(m, v9, int32(_a_F_XLogWrite_2), v9)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v40 = base.AtomicRmwCmpxchg64(m, v36, int32(264), v4, v4)
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v40
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
	v50 = base.I64_rem_u_s(int64(base.Ui64(v40)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v46+int32(1)))
	v53 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v60 = v54
	v63 = v45
	v64 = v9
	v67 = base.I32_wrap_i64(v50)
	v68 = v9
	v69 = v9
	goto L2
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L13
	} else {
		goto L120
	}
L2:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v60) <= base.Ui64(v73) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v562 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v563 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v563) <= base.Ui64(v562) {
		goto L92
	} else {
		goto L93
	}
L4:
	;
	goto L3
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)+292))
	v79 = int64(0)
	v82 = base.AtomicRmwCmpxchg64(m, v75+v67<<(uint(int32(3))%32), int32(0), v79, v79)
	v84 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v82) <= base.Ui64(v84) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v82
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v94 = base.I64_div_u_s(v82-int64(1), base.I64_extend_i32_s(v92))
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	if v94 != v96 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v99 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v126 = v82
	v128 = v92
	goto L9
L9:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if v130 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v111 = v94
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v111
	v117 = F_XLogFileInit(m, v111, l1)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v109 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v110 = base.I64_div_u_s(v105-int64(1), v109)
	v111 = v110
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v117
	F_ReserveExternalFD(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v125 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v126 = v125
	v128 = v123
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	v139 = base.I64_div_u_s(v126-int64(1), base.I64_extend_i32_s(v128))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v139
	v142 = F_XLogFileOpen(m, v139, l1)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v151 = v126
	v152 = v128
	goto L19
L19:
	;
	if v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v142
	F_ReserveExternalFD(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v150 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v151 = v150
	v152 = v148
	goto L19
L22:
	;
	v153 = v69
	goto L24
L23:
	;
	v153 = v67
	goto L24
L24:
	;
	v154 = base.B2i32(base.Ui64(v82) <= base.Ui64(v88))
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v161 = v68
	goto L27
L26:
	;
	v161 = (base.I32_wrap_i64(v151) + int32(-8192)) & (v152 - int32(1))
	goto L27
L27:
	;
	v163 = v64 + int32(1)
	v165 = v163 << (uint(int32(13)) % 32)
	v168 = v154 & base.B2i32(base.Ui32(v152) <= base.Ui32(v161+v165))
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	if base.Ui64(v151) < base.Ui64(v88) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v154 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+296))
	if v168|base.B2i32(v67 == v172) == int32(0) {
		v517 = v88
		v521 = v163
		v525 = v161
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)+288))
	v189 = v165
	v191 = v177 + v153<<(uint(int32(13))%32)
	v194 = v161
	goto L33
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7])) = int32(0)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[8])))
	v205 = m.G0
	v207 = v205 - int32(16)
	m.G0 = v207
	if v202 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v398 = int32(0)
	if v168 == v398 {
		v517 = v88
		v521 = v398
		v525 = v397
		goto L28
	} else {
		goto L67
	}
L35:
	;
	v220 = int32(_a_F_XLogWrite_3)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = int32(167772242)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v227 = F_pwrite(m, v225, v191, v189, base.I64_extend_i32_u(v194))
	mBase = m.M
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[9]))
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v230
	if v227 <= v230 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F___clock_gettime(m, int32(1), v207)
	mBase = m.M
	v211 = int64(*(*int32)(unsafe.Add(mBase, uint32(v207)+8)))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	v216 = v211 + v212*int64(1000000000)
	goto L38
L37:
	;
	v216 = int64(0)
	goto L38
L38:
	;
	m.G0 = v207 + int32(16)
	goto L35
L39:
	;
	if v394 != 0 {
		v189 = v394
		v191 = v396
		v194 = v397
		goto L33
	} else {
		goto L66
	}
L40:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7]))
	if v235 == int32(27) {
		v394 = v189
		v396 = v191
		v397 = v194
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v268 = int32(1)
	v269 = base.I64_extend_i32_u(v227)
	v273 = m.G0
	v275 = v273 - int32(16)
	m.G0 = v275
	if v216 != int64(0) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v239 = v20 + int32(32)
	v241 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	F_XLogFileName(m, v239, l1, v241, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[7])) = v235
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v239
	F_errmsg(m, int32(_a_F_XLogWrite_4), v20)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_XLogWrite_5), int32(2478), int32(_a_F_XLogWrite_6))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v394 = v189 - v227
	v396 = v227 + v191
	v397 = v227 + v194
	goto L39
L50:
	;
	F___clock_gettime(m, int32(1), v275)
	mBase = m.M
	v281 = int64(*(*int32)(unsafe.Add(mBase, uint32(v275)+8)))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v275)))
	v286 = v281 + (v282*int64(1000000000) - v216)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v373 = int32(888)
	v374 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[10])) = v374 + base.I64_extend_i32_u(v268)
	v378 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[11])) = v378 + v269
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v268, v269)
	mBase = m.M
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[12])) = uint8(v383)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[13])) = uint8(v383)
	m.G0 = v275 + int32(16)
	goto L49
L53:
	;
	v336 = int32(888)
	v337 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[14])) = v337 + v286
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[15]))
	v348 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v341))|base.B2i32(int32(1)<<(uint(v341)%32)&int32(_a_F_XLogWrite_7) == v348) == v348 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v353 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[16])) = v353 + v286
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[12])) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[17])) = uint8(v357)
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L52
L66:
	;
	goto L34
L67:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v404 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	F_issue_xlog_fsync(m, v402, v404, l1)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[18])) = uint8(v408)
	v412 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v412
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[19]))
	if int32(0) < v415 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v419 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v420 = m.G0
	v422 = v420 - int32(80)
	m.G0 = v422
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = l1
	v427 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v428 = base.I64_div_u_s(int64(4294967296), v427)
	v429 = base.I64_div_u_s(v419, v428)
	*(*uint32)(unsafe.Add(mBase, uint32(v422)+4)) = uint32(v429)
	v432 = v419 - v428*v429
	*(*uint32)(unsafe.Add(mBase, uint32(v422)+8)) = uint32(v432)
	v435 = v422 + int32(16)
	v438 = F_pg_snprintf(m, v435, int32(64), int32(_a_F_XLogWrite_8), v422)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L13
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v450 = F_time(m)
	mBase = m.M
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v452)+240)) = v450
	v455 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v452)+248)) = v455
	v458 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[20])))
	if v458 != int32(1) {
		v517 = v88
		v521 = v398
		v525 = v397
		goto L28
	} else {
		goto L74
	}
L72:
	;
	F_XLogArchiveNotify(m, v435)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	m.G0 = v422 + int32(80)
	goto L71
L74:
	;
	v462 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[21]))
	v469 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22]))
	v471 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v472 = base.I64_div_u_s(v469, v471)
	if base.Ui64(v462) < base.Ui64(base.I64_extend_i32_s(v464-int32(1))+v472) {
		v517 = v88
		v521 = v398
		v525 = v397
		goto L28
	} else {
		goto L75
	}
L75:
	;
	v477 = base.AtomicRmwXchg32(m, v452, int32(440), int32(1))
	v478 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v477 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_s_lock(m, v452+int32(440), int32(_a_F_XLogWrite_9))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L13
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v485)+200))
	v487 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v485)+440)), uint32(v487))
	v491 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22]))
	if base.Ui64(v491) < base.Ui64(v486) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[22])) = v486
	v495 = v486
	goto L82
L81:
	;
	v495 = v491
	goto L82
L82:
	;
	v496 = int32(0)
	v498 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[21]))
	v505 = int64(*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3])))
	v506 = base.I64_div_u_s(v495, v505)
	if base.Ui64(v498) < base.Ui64(base.I64_extend_i32_s(v500-int32(1))+v506) {
		v517 = v478
		v521 = v496
		v525 = v397
		goto L28
	} else {
		goto L83
	}
L83:
	;
	F_RequestCheckpoint(m, int32(128))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v517 = v478
	v521 = v496
	v525 = v397
	goto L28
L85:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2])) = v517
	goto L4
L86:
	;
	goto L87
L87:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+296))
	if v67 != v538 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v540 = v67 + int32(1)
	goto L90
L89:
	;
	v540 = int32(0)
	goto L90
L90:
	;
	if base.B2i32(l2 == int32(0))|v521 != 0 {
		v60 = v517
		v63 = v537
		v64 = v521
		v67 = v540
		v68 = v525
		v69 = v153
		goto L2
	} else {
		goto L91
	}
L91:
	;
	goto L4
L92:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v641 = base.AtomicRmwXchg32(m, v638, int32(440), int32(1))
	if v641 != 0 {
		goto L107
	} else {
		goto L108
	}
L93:
	;
	v566 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	if base.Ui64(v566) <= base.Ui64(v562) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[23]))
	switch v569 - int32(2) {
	case 0, 2:
		v625 = v566
		goto L95
	default:
		goto L96
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1])) = v625
	v631 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[18])) = uint8(v631)
	goto L92
L96:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v575 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v619 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	F_issue_xlog_fsync(m, v617, v619, l1)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L13
	} else {
		goto L106
	}
L98:
	;
	v579 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4]))
	v583 = base.I64_div_u_s(v566-int64(1), base.I64_extend_i32_s(v573))
	if v579 == v583 {
		v617 = v575
		goto L97
	} else {
		goto L101
	}
L99:
	;
	v595 = v566
	v596 = v573
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[6])) = l1
	v604 = base.I64_div_u_s(v595-int64(1), base.I64_extend_i32_s(v596))
	*(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[4])) = v604
	v607 = F_XLogFileOpen(m, v604, l1)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L13
	} else {
		goto L104
	}
L101:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	if int32(0) <= v588 {
		v617 = v588
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v592 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[3]))
	v595 = v592
	v596 = v594
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5])) = v607
	F_ReserveExternalFD(m)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[5]))
	v617 = v613
	goto L97
L106:
	;
	v623 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v625 = v623
	goto L95
L107:
	;
	F_s_lock(m, v638+int32(440), int32(_a_F_XLogWrite_9))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L13
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v648 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v650)+184))
	if base.Ui64(v651) < base.Ui64(v648) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v650)+184)) = v648
	goto L113
L112:
	;
	goto L113
L113:
	;
	v655 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v650)+192))
	if base.Ui64(v656) < base.Ui64(v655) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v650)+192)) = v655
	goto L116
L115:
	;
	goto L116
L116:
	;
	v659 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v650)+440)), uint32(v659))
	v663 = base.AtomicRmwXchg64(m, v650, int32(264), v648)
	v667 = base.AtomicRmwOr32(m, v659, int32(_a_F_XLogWrite_2), v659)
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_XLogWrite[0]))
	v670 = int32(_a_F_XLogWrite_0)
	v671 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	v673 = base.AtomicRmwXchg64(m, v669, int32(272), v671)
	v675 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[1]))
	if base.Ui64(v53) < base.Ui64(v675) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v678 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XLogWrite[24])) = uint8(v678)
	goto L119
L118:
	;
	goto L119
L119:
	;
	m.G0 = v20 + int32(96)
	return
L120:
	;
	v688 = *(*int64)(unsafe.Add(mBase, _c_F_XLogWrite[2]))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+20)) = uint32(v688)
	v690 = int64(32)
	v691 = int64(base.Ui64(v688) >> (uint(v690) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+16)) = uint32(v691)
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+28)) = uint32(v82)
	v695 = int64(base.Ui64(v82) >> (uint(v690) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+24)) = uint32(v695)
	F_errmsg_internal(m, int32(_a_F_XLogWrite_10), v20+int32(16))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_XLogWrite_5), int32(2380), int32(_a_F_XLogWrite_6))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
