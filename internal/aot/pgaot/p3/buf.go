package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileDumpBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int64
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(1072)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2 < v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L19
	} else {
		goto L42
	}
L2:
	;
	v23 = v16
	v26 = v2
	goto L5
L3:
	;
	v218 = v16
	goto L4
L4:
	;
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v227)
	v229 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v233 = v229 - base.I64_extend_i32_s(v218-v230)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v233
	if v233 < int64(0) {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if int64(1073741824) <= v33 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v218 = v214
	goto L4
L7:
	;
	v37 = v32 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v38 <= v37 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v137 = v23
	v138 = v32
	v143 = v33
	goto L9
L9:
	;
	v147 = int64(1073741824) - v143
	v149 = base.I64_extend_i32_s(v137 - v26)
	if v147 < v149 {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	goto L13
L11:
	;
	v122 = v23
	v123 = v37
	goto L12
L12:
	;
	v131 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v123
	v137 = v122
	v138 = v123
	v143 = v131
	goto L9
L13:
	;
	v51 = int32(_a_F_BufFileDumpBuffer_0)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[0]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[0])) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v122 = v119
	v123 = v117
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[0])) = v52
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = F_repalloc(m, v97, v98<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L19
	} else {
		goto L25
	}
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v60 = F_OpenTemporaryFile(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v63 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v62
	v69 = v14 + int32(48)
	v74 = F_pg_snprintf(m, v69, int32(1024), int32(_a_F_BufFileDumpBuffer_1), v14+int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return
L20:
	;
	v94 = v60
	goto L15
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v77 = F_FileSetDelete(m, v76, v69)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v79
	v86 = F_pg_snprintf(m, v69, int32(1024), int32(_a_F_BufFileDumpBuffer_1), v14+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = F_FileSetCreate(m, v88, v69)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v94 = v89
	goto L15
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v106<<(uint(int32(2))%32)))) = v94
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = int32(1)
	v113 = v111 + v112
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v117 = v115 + v112
	if v113 <= v117 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	v151 = v147
	goto L29
L28:
	;
	v151 = v149
	goto L29
L29:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v138<<(uint(int32(2))%32))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[1])))
	if v160 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F___clock_gettime(m, int32(1), v14+int32(48))
	mBase = m.M
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v173 = v172
	v174 = v167*int64(-1000000000) - v170
	goto L32
L31:
	;
	v173 = v143
	v174 = int64(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = base.I32_wrap_i64(v151)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v26 + (l0 + int32(48))
	v179 = v14 + int32(48)
	v182 = F_FileWriteV(m, v157, v179, int32(1), v173, int32(167772168))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if v182 <= int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[1])))
	if v187 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F___clock_gettime(m, int32(1), v179)
	mBase = m.M
	v192 = int32(_a_F_BufFileDumpBuffer_2)
	v194 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[2]))
	v195 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[2])) = v194 + (v195 + (v196*int64(1000000000) + v174))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v203 + base.I64_extend_i32_u(v182)
	v207 = int32(_a_F_BufFileDumpBuffer_3)
	v209 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[3])) = v209 + int64(1)
	v213 = v182 + v26
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v213 < v214 {
		v23 = v214
		v26 = v213
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L6
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v233 + int64(1073741824)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v240 - int32(1)
	goto L41
L40:
	;
	goto L41
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	m.G0 = v14 + int32(1072)
	return
L42:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[4]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v157*int32(48))+32))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v260
	F_errmsg(m, int32(_a_F_BufFileDumpBuffer_4), v14)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_BufFileDumpBuffer_5), int32(546), int32(_a_F_BufFileDumpBuffer_6))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
