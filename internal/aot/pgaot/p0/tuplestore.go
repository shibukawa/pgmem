package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_trim(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
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
	var v129 int64
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	v2 = int32(0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v13&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v18 <= int32(0) {
		v87 = v17
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v98 = v87 - int32(1)
	if v98 <= int32(0) {
		goto L1
	} else {
		goto L29
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v18 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = v17
	v31 = v2
	v37 = v2
	goto L9
L7:
	;
	v68 = v17
	v69 = v2
	goto L8
L8:
	;
	v80 = v21 + v69*int32(24)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
	if v81 != 0 {
		v87 = v68
		goto L4
	} else {
		goto L25
	}
L9:
	;
	v42 = v21 + v31*int32(24)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v18&int32(1) == int32(0) {
		v87 = v58
		goto L4
	} else {
		goto L24
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v30 < v46 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v49 = v30
	goto L13
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+28)))
	if v51 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v48 = v30
	goto L16
L15:
	;
	v48 = v46
	goto L16
L16:
	;
	v49 = v48
	goto L13
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+32))
	if v49 < v54 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v58 = v49
	goto L19
L19:
	;
	v59 = int32(2)
	v60 = v31 + v59
	v62 = v37 + v59
	if v62 != v18&int32(2147483646) {
		v30 = v58
		v31 = v60
		v37 = v62
		goto L9
	} else {
		goto L23
	}
L20:
	;
	v56 = v49
	goto L22
L21:
	;
	v56 = v54
	goto L22
L22:
	;
	v58 = v56
	goto L19
L23:
	;
	goto L10
L24:
	;
	v68 = v58
	v69 = v60
	goto L8
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if v68 < v82 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = v68
	goto L28
L27:
	;
	v84 = v82
	goto L28
L28:
	;
	v87 = v84
	goto L4
L29:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v104 = v102 - v103
	if v104 < v101 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = v101
	goto L32
L31:
	;
	v106 = v104
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v108 < v98 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v113 = v108
	goto L36
L34:
	;
	v150 = v17
	goto L35
L35:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v158)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v98
	v162 = base.I32_div_s(v150, int32(8))
	if v98 < v162 {
		goto L1
	} else {
		goto L42
	}
L36:
	;
	v123 = v113 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123+v124)))
	v127 = F_GetMemoryChunkSpace(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v150 = v145
	goto L35
L38:
	;
	return
L39:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v129 + base.I64_extend_i32_u(v127)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133+v123)))
	F_pfree(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v123))) = int32(0)
	v143 = v113 + int32(1)
	if v143 != v98 {
		v113 = v143
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v167 = v164 + v98<<(uint(int32(2))%32)
	if v87 == v150 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v181 - v98
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v184 <= v178 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v169
	goto L43
L45:
	;
	goto L46
L46:
	;
	v173 = (v150 - v98) << (uint(int32(2)) % 32)
	if v173 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	base.MemoryCopy(m, v164, v167, v173)
	goto L43
L48:
	;
	v189 = v178
	v190 = v184
	goto L49
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v202 = v199 + v189*int32(24)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+4)))
	if v203 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L1
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v206 - v98
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v210 = v209
	goto L53
L52:
	;
	v210 = v190
	goto L53
L53:
	;
	v212 = v189 + int32(1)
	if v212 < v210 {
		v189 = v212
		v190 = v210
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
}
