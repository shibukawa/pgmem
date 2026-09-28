package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileDumpBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
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
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v137 int32
	_ = v137
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v222 int64
	_ = v222
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v8 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(1072)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v8 < v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L19
	} else {
		goto L42
	}
L2:
	;
	v28 = v16
	v30 = v8
	goto L5
L3:
	;
	v222 = v16
	goto L4
L4:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v226)
	v228 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v231 = v228 + (v229 - v222)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v231
	if v231 < int64(0) {
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
	v222 = v213
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
	v137 = v32
	v142 = v28
	v143 = v33
	goto L9
L9:
	;
	v146 = v142 - v30
	v148 = int64(1073741824) - v143
	if v146 < v148 {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	goto L13
L11:
	;
	v122 = v37
	v127 = v28
	goto L12
L12:
	;
	v131 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
	v137 = v122
	v142 = v127
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
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v122 = v117
	v127 = v119
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
	v150 = v146
	goto L29
L28:
	;
	v150 = v148
	goto L29
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v137<<(uint(int32(2))%32))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[1])))
	if v158 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F___clock_gettime(m, int32(1), v14+int32(48))
	mBase = m.M
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	v169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
	v171 = v166*int64(-1000000000) - v169
	v172 = v165
	goto L32
L31:
	;
	v171 = int64(0)
	v172 = v143
	goto L32
L32:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+52)) = uint32(v150)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l0 + int32(56) + base.I32_wrap_i64(v30)
	v178 = v14 + int32(48)
	v181 = F_FileWriteV(m, v155, v178, int32(1), v172, int32(167772168))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if v181 <= int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[1])))
	if v186 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F___clock_gettime(m, int32(1), v178)
	mBase = m.M
	v191 = int32(_a_F_BufFileDumpBuffer_2)
	v193 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[2]))
	v194 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[2])) = v193 + (v194 + (v195*int64(1000000000) + v171))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v202 = base.I64_extend_i32_s(v181)
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v202 + v203
	v206 = int32(_a_F_BufFileDumpBuffer_3)
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[3])) = v208 + int64(1)
	v212 = v202 + v30
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v212 < v213 {
		v28 = v213
		v30 = v212
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L6
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v231 + int64(1073741824)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v238 - int32(1)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v243 = l0 + int32(40)
	v244 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v243)+8)) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v243))) = v244
	m.G0 = v14 + int32(1072)
	return
L42:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_BufFileDumpBuffer[4]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258+v155*int32(48))+32))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v262
	F_errmsg(m, int32(_a_F_BufFileDumpBuffer_4), v14)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_BufFileDumpBuffer_5), int32(547), int32(_a_F_BufFileDumpBuffer_6))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
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
