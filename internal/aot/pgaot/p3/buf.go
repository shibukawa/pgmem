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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int64
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(1072)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L42
	}
L2:
	;
	v21 = v14
	v23 = v2
	goto L5
L3:
	;
	v215 = v14
	goto L4
L4:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v222)
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v228 = v224 - base.I64_extend_i32_s(v215-v225)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v228
	if v228 < int64(0) {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if int64(1073741824) <= v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v215 = v211
	goto L4
L7:
	;
	v33 = v28 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 <= v33 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v134 = v21
	v135 = v28
	v139 = v29
	goto L9
L9:
	;
	v142 = int64(1073741824) - v139
	v144 = v134 - v23
	if v142 < base.I64_extend_i32_s(v144) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	goto L13
L11:
	;
	v121 = v21
	v122 = v33
	goto L12
L12:
	;
	v128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
	v134 = v121
	v135 = v122
	v139 = v128
	goto L9
L13:
	;
	v45 = int32(4515540)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v121 = v118
	v122 = v116
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v46
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v102 = F_repalloc(m, v96, v97<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L25
	}
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v54 = F_OpenTemporaryFile(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v57 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v56
	v68 = F_pg_snprintf(m, v12+int32(48), int32(1024), int32(466808), v12+int32(32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return
L20:
	;
	v93 = v54
	goto L15
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = F_FileSetDelete(m, v70, v12+int32(48))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v75
	v84 = F_pg_snprintf(m, v12+int32(48), int32(1024), int32(466808), v12+int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = F_FileSetCreate(m, v86, v12+int32(48))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v93 = v89
	goto L15
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v102+v105<<(uint(int32(2))%32)))) = v93
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = int32(1)
	v112 = v110 + v111
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v116 = v114 + v111
	if v112 <= v116 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	v147 = base.I32_wrap_i64(v142)
	goto L29
L28:
	;
	v147 = v144
	goto L29
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v135<<(uint(int32(2))%32))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v155 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F___clock_gettime(m, int32(1), v12+int32(48))
	mBase = m.M
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+56)))
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v168 = v167
	v169 = v162*int64(-1000000000) - v165
	goto L32
L31:
	;
	v168 = v139
	v169 = int64(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v23 + (l0 + int32(48))
	v177 = F_FileWriteV(m, v152, v12+int32(48), int32(1), v168, int32(167772168))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if v177 <= int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v182 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F___clock_gettime(m, int32(1), v12+int32(48))
	mBase = m.M
	v189 = int32(4413872)
	v191 = *(*int64)(unsafe.Add(mBase, _consts[67]))
	v192 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+56)))
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int64)(unsafe.Add(mBase, _consts[67])) = v191 + (v192 + (v193*int64(1000000000) + v169))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v200 + base.I64_extend_i32_u(v177)
	v204 = int32(4413824)
	v206 = *(*int64)(unsafe.Add(mBase, _consts[61]))
	*(*int64)(unsafe.Add(mBase, _consts[61])) = v206 + int64(1)
	v210 = v177 + v23
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v210 < v211 {
		v21 = v211
		v23 = v210
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L6
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v228 + int64(1073741824)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v235 - int32(1)
	goto L41
L40:
	;
	goto L41
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	m.G0 = v12 + int32(1072)
	return
L42:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v152*int32(48))+32))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v255
	F_errmsg(m, int32(298450), v12)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(499138), int32(546), int32(226545))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
