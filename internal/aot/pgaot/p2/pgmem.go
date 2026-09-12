package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgmem_dlsym(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var __phi91 int32
	_ = __phi91
	var v94 int32
	_ = v94
	var __phi94 int32
	_ = __phi94
	var v95 int32
	_ = v95
	var __phi95 int32
	_ = __phi95
	var v96 int32
	_ = v96
	var __phi96 int32
	_ = __phi96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(288)
	m.G0 = v10
	if l0 == v3 {
		v210 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(288)
	return v210
L2:
	;
	if l1 == int32(0) {
		v210 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(508661)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1302])))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		v39 = v19
		v40 = v20
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v146 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L5:
	;
	if v40-v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v19 != v20 {
		v39 = v19
		v40 = v20
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = l1
	v25 = v16
	goto L9
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v28
		v40 = v29
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v39 = v28
	v40 = v29
	goto L6
L11:
	;
	v32 = int32(1)
	if v28 == v29 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(106410)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1303])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v68 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v76 = F_snprintf(m, v10+int32(32), int32(256), int32(528748), v10+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	if v66-v65 != 0 {
		v140 = l1
		goto L4
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v50 = l1
	v51 = v42
	goto L20
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v65 = v54
	v66 = v55
	goto L17
L22:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L15
L25:
	;
	return int32(0)
L26:
	;
	v81 = v68 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v133 = v10 + int32(32)
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v127))) = uint8(v135)
	v140 = v133
	goto L4
L28:
	;
	v127 = v81
	goto L27
L29:
	;
	goto L30
L30:
	;
	v87 = v68 + int32(2)
	if base.Ui32(int32(255)) < base.Ui32(v87) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v127 = v81
	goto L27
L32:
	;
	goto L33
L33:
	;
	__phi91 = v83
	__phi94 = v81
	__phi95 = v82
	__phi96 = v87
	v91 = __phi91
	v94 = __phi94
	v95 = __phi95
	v96 = __phi96
	goto L34
L34:
	;
	v97 = int32(32)
	v102 = v91 & int32(255)
	goto L36
L35:
	;
	v127 = v96
	goto L27
L36:
	;
	if base.B2i32(base.Ui32(v102-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v102|v97-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v114 = v91
	goto L39
L38:
	;
	v114 = int32(95)
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v97+v94))) = uint8(v114)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v116 == int32(0) {
		v127 = v96
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v119 = int32(1)
	if base.Ui32(v96) < base.Ui32(int32(255)) {
		__phi91 = v116
		__phi94 = v96
		__phi95 = v95 + v119
		__phi96 = v96 + v119
		v91 = __phi91
		v94 = __phi94
		v95 = __phi95
		v96 = __phi96
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L35
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v140
	v202 = F_snprintf(m, int32(4641696), int32(512), int32(211318), v10)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L25
	} else {
		goto L58
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = int32(0)
	goto L44
L44:
	;
	v160 = v149 + v153<<(uint(int32(3))%32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v165 == int32(0) {
		v184 = v164
		v185 = v165
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v210 = v190
	goto L1
L46:
	;
	if v185-v184 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	goto L46
L48:
	;
	if v164 != v165 {
		v184 = v164
		v185 = v165
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v169 = v161
	v170 = v140
	goto L50
L50:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v174 == int32(0) {
		v184 = v173
		v185 = v174
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v184 = v173
	v185 = v174
	goto L47
L52:
	;
	v177 = int32(1)
	if v173 == v174 {
		v169 = v169 + v177
		v170 = v170 + v177
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v188 = v153 + int32(1)
	if v146 != v188 {
		v153 = v188
		goto L44
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L45
L57:
	;
	goto L42
L58:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1304])) = uint8(v206)
	v210 = int32(0)
	goto L1
}
func F_pgmem_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	v4 = int32(0)
	if l1 == v4 {
	} else {
		v12 = l1 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(l1) {
			v20 = v4
			v23 = v4
			for {
				v27 = l0 + v20<<(uint(int32(3))%32)
				v28 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+14)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+22)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+30)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+38)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+46)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+54)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+62)) = uint16(v28)
				v44 = int32(8)
				v45 = v20 + v44
				v47 = v23 + v44
				if v47 != l1&int32(-8) {
					v20 = v45
					v23 = v47
					continue
				} else {
					break
				}
				break
			}
			v52 = v45
		} else {
			v52 = v4
		}
		if v12 == int32(0) {
		} else {
			v62 = v52
			v64 = v4
			for {
				v70 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0+v62<<(uint(int32(3))%32))+6)) = uint16(v70)
				v72 = int32(1)
				v75 = v64 + v72
				if v75 != v12 {
					v62 = v62 + v72
					v64 = v75
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v85 = m.Env.Pgmem_poll(m, l2)
	mBase = m.M
	return int32(0)
}
