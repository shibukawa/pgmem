package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CreateDecodingContext(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v1 = l0
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[0]))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L23
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L23
	} else {
		goto L64
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L23
	} else {
		goto L60
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L23
	} else {
		goto L57
	}
L7:
	;
	v23 = v18 + int32(24)
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[1]))
	if v19 != v27 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	if v1 == int64(0) {
		v83 = v51
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v41 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+308))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[2])) = uint8(v39)
	v41 = v39
	goto L16
L15:
	;
	v41 = int32(0)
	goto L16
L16:
	;
	goto L13
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+201)))
	if v44 != int32(1) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDecodingContext[7])))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v84 = int32(0)
	v88 = F_StartupDecodingContext(m, l1, v83, v84, v84, l2, v84, v84, l3, l4, l5, l6)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L30
	}
L21:
	;
	if base.Ui64(v51) <= base.Ui64(v1) {
		v83 = v1
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v57 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v61)
	v63 = int64(32)
	v64 = int64(base.Ui64(v61) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v64)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v1)
	v68 = int64(base.Ui64(v1) >> (uint(v63) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v68)
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_10), v15+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v83 = v81
	goto L20
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(651), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v90 = int32(_a_F_CreateDecodingContext_4)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+24))
	if v95 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(_a_F_CreateDecodingContext_5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1058)
	v103 = int32(_a_F_CreateDecodingContext_6)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v15 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v15 + int32(80)
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+164)) = uint8(v113)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+147)) = uint8(v113)
	m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v88, v88+int32(108), v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L23
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[4])) = v91
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v128 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[5])) = v123
	goto L33
L35:
	;
	v131 = int32(1)
	goto L37
L36:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+146)))
	v131 = v130
	goto L37
L37:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+145)))
	v133 = v131 & v132
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+145)) = uint8(v133)
	if v133 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+116)) = uint8(v157)
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDecodingContext[6]))
	if v162 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)))
	if v137 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v140 = base.AtomicRmwXchg32(m, v18, int32(0), int32(1))
	if v140 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_s_lock(m, v18, int32(_a_F_CreateDecodingContext_9))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L23
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v83
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v145)
	v147 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18))), uint32(v147))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v83
	goto L38
L47:
	;
	v165 = int32(14)
	goto L49
L48:
	;
	v165 = int32(15)
	goto L49
L49:
	;
	v167 = F_errstart(m, v165, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	if v167 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_7), v15+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L23
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	m.G0 = v15 + int32(96)
	return v88
L54:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v18)+120))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v18)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v176)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v175)
	v179 = int64(32)
	v180 = int64(base.Ui64(v176) >> (uint(v179) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v180)
	v183 = int64(base.Ui64(v175) >> (uint(v179) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15))) = uint32(v183)
	v186 = F_errdetail(m, int32(_a_F_CreateDecodingContext_8), v15)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(695), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	F_errmsg_internal(m, int32(_a_F_CreateDecodingContext_11), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(594), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_CreateDecodingContext_0), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(600), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L23
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
	F_errcode(m, int32(325))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_3), v15-int32(-64))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(611), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v23
	F_errmsg(m, int32(_a_F_CreateDecodingContext_12), v15+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	v261 = F_errdetail(m, int32(_a_F_CreateDecodingContext_13), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	F_errhint(m, int32(_a_F_CreateDecodingContext_14), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_CreateDecodingContext_1), int32(624), int32(_a_F_CreateDecodingContext_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateDirAndVersionFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
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
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v8 = m.G0
	v10 = v8 - int32(1136)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(_a_F_CreateDirAndVersionFile_0)
	v19 = F_pg_sprintf(m, v10+int32(96), int32(_a_F_CreateDirAndVersionFile_1), v10+int32(80))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[0]))
	v23 = F_mkdir(m, l0, v22)
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L60
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L5:
	;
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(_a_F_CreateDirAndVersionFile_2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l0
	v36 = v10 + int32(112)
	v41 = F_pg_snprintf(m, v36, int32(1024), int32(_a_F_CreateDirAndVersionFile_3), v10+int32(48))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v29 != int32(20) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v44 = F_OpenTransientFile(m, v36, int32(193))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v44 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l3 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v59 = v44
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(167772228)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(0)
	v69 = int32(3)
	v70 = F_write(m, v59, v10+int32(96), v69)
	mBase = m.M
	if v70 != v69 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v51 != int32(20) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v55 = F_OpenTransientFile(m, v36, int32(513))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v55 < int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v59 = v55
	goto L15
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v74 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v99 = int32(_a_F_CreateDirAndVersionFile_4)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	v101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v101
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(167772227)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[3])))
	if v109 != int32(1) {
		v123 = v101
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1])) = int32(51)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_5), v10+int32(32))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(511), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	F_fsync_fname(m, l0, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L31:
	;
	if v123 == int32(0) {
		goto L30
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	goto L34
L34:
	;
	v114 = F_fsync(m, v59)
	mBase = m.M
	if v114 != int32(-1) {
		v123 = v114
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v123 = int32(-1)
	goto L32
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[1]))
	if v118 == int32(27) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[4])))
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v132 = F_errstart(m, v130, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v130 = int32(21)
	goto L42
L41:
	;
	v130 = int32(24)
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v132 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_8), v10+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(519), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L30
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(0)
	v158 = F_CloseTransientFile(m, v59)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if l3 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v162 = int32(_a_F_CreateDirAndVersionFile_9)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v164 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	m.G0 = v10 + int32(1136)
	return
L53:
	;
	F_XLogRegisterData(m, v10+int32(88), int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v179 = F_XLogInsert(m, int32(4), int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v181 = int32(_a_F_CreateDirAndVersionFile_9)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateDirAndVersionFile[5])) = v183 - int32(1)
	goto L52
L56:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l0
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_10), v10-int32(-64))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(482), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_CreateDirAndVersionFile_11), v10)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateDirAndVersionFile_6), int32(499), int32(_a_F_CreateDirAndVersionFile_7))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_drop_transactional_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v4 = l3
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
	v14 = F_MemoryContextAlloc(m, v12, int32(28))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v10 {
				v39 = v17
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = v39 + int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v47 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
				v58 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
				v23 = F_MemoryContextAlloc(m, v21, int32(24))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
					v29 = v23 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
					v32 = int32(_a_F_create_drop_transactional_internal_0)
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
					*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
					v39 = v23
					*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
					v46 = v39 + int32(8)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					if v47 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
					v58 = v14 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[1]))
			v23 = F_MemoryContextAlloc(m, v21, int32(24))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v10
				v29 = v23 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v29
				v32 = int32(_a_F_create_drop_transactional_internal_0)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v33
				*(*int32)(unsafe.Add(mBase, _c_F_create_drop_transactional_internal[2])) = v23
				v39 = v23
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v4)
				v46 = v39 + int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v47 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v46
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v55
				v58 = v14 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v58
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v61 + int32(1)
				return
			}
		}
	}
}
func F_create_plan(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	*(*int64)(unsafe.Add(mBase, uint32(l0)+368)) = int64(0)
	v6 = F_create_plan_recurse(m, l0, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v10 != int32(337) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v22 = int32(0)
	goto L7
L4:
	;
	goto L5
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	goto L5
L7:
	;
	v23 = int32(0)
	if v13 == v23 {
		v33 = v23
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v14 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= v22 {
		v33 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = v29 + v22<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+v22<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v52
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v54)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+26)) = uint8(v56)
	v22 = v22 + int32(1)
	goto L7
L13:
	;
	goto L6
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v33 == int32(0))|base.B2i32(v38 <= v22) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v41 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	return v6
L20:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_0), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_create_plan_1), int32(374), int32(_a_F_create_plan_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_syncrep_config(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v3 = l2
	v4 = int32(0)
	v8 = int32(16)
	if l1 == v4 {
		v41 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = F_palloc(m, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		v41 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v11
	goto L6
L5:
	;
	v17 = v14
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = v8
	v24 = v4
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v24<<(uint(int32(2))%32))))
	v30 = F_strlen(m, v29)
	mBase = m.M
	v32 = int32(1)
	v33 = v30 + v23 + v32
	v35 = v24 + v32
	if v35 != v17 {
		v23 = v33
		v24 = v35
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v41 = v33
	goto L1
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v41
	v52 = l0
	goto L13
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v96
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v57 = v52 + int32(1)
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v59 = F___isspace(m, v58)
	mBase = m.M
	if v59 != 0 {
		v52 = v57
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v60 = int32(1)
	switch v58&int32(255) - int32(43) {
	case 0:
		v66 = v60
		goto L17
	default:
		v68 = v58
		v69 = v52
		v70 = v60
		goto L16
	case 2:
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v71 = int32(0)
	v73 = v68 - int32(48)
	if base.Ui32(v73) <= base.Ui32(int32(9)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57))))
	v68 = v67
	v69 = v57
	v70 = v66
	goto L16
L18:
	;
	v66 = int32(0)
	goto L17
L19:
	;
	v76 = v71
	v77 = v73
	v78 = v69
	goto L22
L20:
	;
	v90 = v71
	goto L21
L21:
	;
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v80 = int32(10)
	v82 = v76*v80 - v77
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+1)))
	v87 = v83 - int32(48)
	if base.Ui32(v87) < base.Ui32(v80) {
		v76 = v82
		v77 = v87
		v78 = v78 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v90 = v82
	goto L21
L24:
	;
	goto L23
L25:
	;
	v96 = int32(0) - v90
	goto L27
L26:
	;
	v96 = v90
	goto L27
L27:
	;
	goto L12
L28:
	;
	return v44
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v101 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(0)
	goto L28
L32:
	;
	v111 = v44 + int32(16)
	v112 = int32(0)
	goto L33
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v112<<(uint(int32(2))%32))))
	if (v118^v111)&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	goto L28
L35:
	;
	v193 = F_strlen(m, v118)
	mBase = m.M
	v195 = int32(1)
	v198 = v112 + v195
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v198 < v199 {
		v111 = v111 + v193 + v195
		v112 = v198
		goto L33
	} else {
		goto L56
	}
L36:
	;
	goto L35
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v172)
	if v172&int32(255) == int32(0) {
		goto L36
	} else {
		goto L52
	}
L38:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v171 = v118
	v172 = v124
	v173 = v111
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v118&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v128 = v118
	v130 = v111
	goto L44
L42:
	;
	v142 = v118
	v144 = v111
	goto L43
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v149 = int32(-2139062144)
	if (int32(16843008)-v146|v146)&v149 != v149 {
		v171 = v142
		v172 = v146
		v173 = v144
		goto L37
	} else {
		goto L48
	}
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v131)
	if v131 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L45:
	;
	v142 = v138
	v144 = v136
	goto L43
L46:
	;
	v135 = int32(1)
	v136 = v130 + v135
	v138 = v128 + v135
	if v138&int32(3) != 0 {
		v128 = v138
		v130 = v136
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v154 = v142
	v155 = v146
	v156 = v144
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155
	v158 = int32(4)
	v159 = v156 + v158
	v161 = v154 + v158
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v166 = int32(-2139062144)
	if (int32(16843008)-v163|v163)&v166 == v166 {
		v154 = v161
		v155 = v163
		v156 = v159
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v171 = v161
	v172 = v163
	v173 = v159
	goto L37
L51:
	;
	goto L50
L52:
	;
	v180 = v171
	v182 = v173
	goto L53
L53:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)) = uint8(v183)
	v185 = int32(1)
	if v183 != 0 {
		v180 = v180 + v185
		v182 = v182 + v185
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L36
L55:
	;
	goto L54
L56:
	;
	goto L34
}
func F_create_tidrangescan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 int64
	_ = v61
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v117 float64
	_ = v117
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v189 float64
	_ = v189
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v221 float64
	_ = v221
	var v226 float64
	_ = v226
	var v227 int32
	_ = v227
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v232 float64
	_ = v232
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v241 float64
	_ = v241
	var v243 int32
	_ = v243
	var v249 float64
	_ = v249
	var v253 float64
	_ = v253
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v258 float64
	_ = v258
	var v267 float64
	_ = v267
	var v271 float64
	_ = v271
	var v275 float64
	_ = v275
	var v278 int64
	_ = v278
	var v280 float64
	_ = v280
	var v284 int64
	_ = v284
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v297 float64
	_ = v297
	v9 = float64(0)
	v21 = F_palloc0(m, int32(80))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(1503238553889)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v28
	v30 = F_get_baserel_parampathinfo(m, l0, l1, l3)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)) = uint8(base.B2i32(v32 < l4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v30
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+21)) = uint8(v36)
	v43 = m.G0
	v45 = v43 - int32(48)
	m.G0 = v45
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v51 = v30 + int32(8)
	goto L6
L5:
	;
	v51 = l1 + int32(16)
	goto L6
L6:
	;
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v55 = int32(0)
	v57 = F_clauselist_selectivity(m, l0, l2, v54, v55, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l1)+128))
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+32)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v61
	if l2 == int32(0) {
		v117 = v9
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	F_get_tablespace_page_costs(m, v124, v45+int32(16), v45+int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v68 <= int32(0) {
		v117 = v9
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v75 = v32
	goto L11
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v75<<(uint(int32(2))%32))))
	v97 = F_cost_qual_eval_walker(m, v94, v45+int32(24))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v45)+40))
	v117 = v103
	goto L8
L13:
	;
	v100 = v75 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v100 < v101 {
		v75 = v100
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v45)+16))
	if v30 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v227)+24))
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v232 = *(*float64)(unsafe.Add(mBase, _c_F_create_tidrangescan_path[0]))
	v236 = base.F64_add(base.F64_mul(v228, v229), base.F64_mul(base.F64_mul(v57, v60), base.F64_sub(base.F64_add(v226, v232), v117)))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v227)+16))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if int32(0) < v238 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v134 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+32)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v134
	if v133 == int32(0) {
		v189 = v9
		v200 = float64(0)
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v221 = v205
	v226 = v206
	goto L16
L20:
	;
	v201 = *(*float64)(unsafe.Add(mBase, uint32(l1)+208))
	v203 = *(*float64)(unsafe.Add(mBase, uint32(l1)+216))
	v221 = base.F64_add(v200, v201)
	v226 = base.F64_add(v189, v203)
	goto L16
L21:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v143 <= int32(0) {
		v189 = v9
		v200 = float64(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v151 = int32(0)
	goto L23
L23:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166+v151<<(uint(int32(2))%32))))
	v173 = F_cost_qual_eval_walker(m, v170, v45+int32(24))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v45)+40))
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v45)+32))
	v189 = v179
	v200 = v180
	goto L20
L25:
	;
	v176 = v151 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v176 < v177 {
		v151 = v176
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v241 = base.F64_convert_i32_u(v238)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tidrangescan_path[1])))
	if v243 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v275 = v236
	goto L29
L29:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v280 = base.F64_add(base.F64_add(v117, v221), v237)
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = v280
	if v238 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v249 = base.F64_add(base.F64_mul(v241, float64(-0.3)), float64(1))
	if base.F64_gt(v249, float64(0)) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v255 = v241
	goto L32
L32:
	;
	v257 = float64(1e+100)
	v258 = base.F64_div(v229, v255)
	if base.F64_gt(v258, v257)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v258)&int64(9223372036854775807))) != 0 {
		v271 = v257
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v253 = v249
	goto L35
L34:
	;
	v253 = math.Float64frombits(uint64(0x8000000000000000))
	goto L35
L35:
	;
	v255 = base.F64_add(v253, v241)
	goto L32
L36:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = v271
	v275 = base.F64_div(v236, v255)
	goto L29
L37:
	;
	v267 = float64(1)
	if base.F64_le(v258, v267) != 0 {
		v271 = v267
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v271 = base.F64_nearest(v258)
	goto L36
L39:
	;
	v284 = int64(-17)
	goto L41
L40:
	;
	v284 = int64(-262161)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = base.B2i32(v278|v284 != int64(-1))
	v289 = float64(0)
	v292 = base.F64_ceil(base.F64_mul(v57, base.F64_convert_i32_u(v59)))
	if base.F64_le(v292, v289) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v297 = v289
	goto L44
L43:
	;
	v297 = base.F64_add(v292, float64(-1))
	goto L44
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(base.F64_add(base.F64_mul(v131, v297), v132), base.F64_add(v280, v275))
	m.G0 = v45 + int32(48)
	return v21
}
func F_createarc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v3 = l2
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v12
		v88 = v11
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
		if v94 != 0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
			*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if v103 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
			if v110 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v114 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
			v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
			if v121 < int32(0) {
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
				v126 = v124 - int32(97)
				if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					if v136 != 0 {
					} else {
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
						v142 = v139 + v121*int32(24)
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
						if v143 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
							v146 = v145
						} else {
							v146 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
					}
				}
			}
		}
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if base.Ui32(v16) <= base.Ui32(v15) {
				v34 = l0 + int32(76)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+136))
				if base.Ui32(v36) < base.Ui32(int32(98000000)) {
					v50 = int32(1024)
					v52 = v16 << (uint(int32(1)) % 32)
					if base.Ui32(v50) <= base.Ui32(v52) {
						v55 = v50
					} else {
						v55 = v52
					}
					v57 = v34
					v58 = v55
					v62 = v58*int32(40) + int32(8)
					v64 = F_palloc_extended(m, v62, int32(2))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						if v64 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(101)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
							if v72 != 0 {
								v74 = v72
							} else {
								v74 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
							v88 = int32(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+136)) = v77 + v62
							*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v58
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = v81
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
							v88 = v64 + int32(8)
						}
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
						if v94 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
							*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if v103 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							if v110 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v114 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
							v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
							if v121 < int32(0) {
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
								v126 = v124 - int32(97)
								if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
									if v136 != 0 {
									} else {
										v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
										v142 = v139 + v121*int32(24)
										v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										if v143 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
											v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
											v146 = v145
										} else {
											v146 = int32(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
										*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
									}
								}
							}
						}
						return
					}
				} else {
					v39 = v35
					v40 = v34
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(101)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
					if v45 != 0 {
						v47 = v45
					} else {
						v47 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v47
					v88 = int32(0)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
					if v94 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v103 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v110 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v114 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
						v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
						if v121 < int32(0) {
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v126 = v124 - int32(97)
							if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v136 != 0 {
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
									v142 = v139 + v121*int32(24)
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									if v143 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										v146 = v145
									} else {
										v146 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
									*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
								}
							}
						}
					}
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				v88 = v14 + v15*int32(40) + int32(8)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
				if v94 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v103 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v110 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v114 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
					v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
					if v121 < int32(0) {
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v126 = v124 - int32(97)
						if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							if v136 != 0 {
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
								v142 = v139 + v121*int32(24)
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
								if v143 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									v146 = v145
								} else {
									v146 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
								*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
							}
						}
					}
				}
				return
			}
		} else {
			v27 = l0 + int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
			if base.Ui32(int32(97999999)) < base.Ui32(v30) {
				v39 = v29
				v40 = v27
				*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(101)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
				if v45 != 0 {
					v47 = v45
				} else {
					v47 = int32(19)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v47
				v88 = int32(0)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
				if v94 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
					*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if v103 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v110 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					v114 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
					v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
					if v121 < int32(0) {
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v126 = v124 - int32(97)
						if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							if v136 != 0 {
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
								v142 = v139 + v121*int32(24)
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
								if v143 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									v146 = v145
								} else {
									v146 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
								*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
							}
						}
					}
				}
				return
			} else {
				v57 = v27
				v58 = int32(64)
				v62 = v58*int32(40) + int32(8)
				v64 = F_palloc_extended(m, v62, int32(2))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v64 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(101)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
						if v72 != 0 {
							v74 = v72
						} else {
							v74 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
						v88 = int32(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+136))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+136)) = v77 + v62
						*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v58
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v81
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64
						v88 = v64 + int32(8)
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
					if v94 != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = l4
						*(*uint16)(unsafe.Add(mBase, uint32(v88)+4)) = uint16(v3)
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v99
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if v103 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v88
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v106
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						if v110 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = v88
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v88
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v114 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v113 + v114
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v117 + v114
						v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+4)))
						if v121 < int32(0) {
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v126 = v124 - int32(97)
							if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v126))|base.B2i32(int32(1)<<(uint(v126)%32)&int32(_a_F_createarc_0) == int32(0)) != 0 {
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								if v136 != 0 {
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
									v142 = v139 + v121*int32(24)
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
									if v143 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = v88
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
										v146 = v145
									} else {
										v146 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v146
									*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v88
								}
							}
						}
					}
					return
				}
			}
		}
	}
}
