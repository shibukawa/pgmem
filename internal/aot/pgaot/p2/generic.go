package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generic_identify(m *base.Module, l0 int32) int32 {
	return int32(_a_F_generic_identify_0)
}
func F_generic_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v18 < v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return
L2:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = v17
	v28 = v2
	v32 = v2
	goto L3
L3:
	;
	v42 = v15 + int32(16) + v28<<(uint(int32(2))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28*int32(52))+76)))
	if v46 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v168 = int32(0)
	if v162 < v168 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+72))
	v164 = v32 + int32(1)
	v166 = v164 & int32(255)
	if v166 <= v162 {
		v27 = v161
		v28 = v166
		v32 = v164
		goto L3
	} else {
		goto L40
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v52 = v32 & int32(255)
	v53 = F_XLogReadBufferForRedo(m, l0, v52, v42)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v53 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v55 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v75 = v15 + int32(12)
	v76 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
	if v78 < v52 {
		v100 = v76
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_generic_redo[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+(v55^int32(-1))<<(uint(int32(2))%32))))
	v73 = v65
	goto L12
L14:
	;
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_generic_redo[1]))
	v73 = v67 + v55<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v103 = v100
	goto L16
L18:
	;
	v82 = v77 + v52*int32(52)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+76)))
	if v83 != int32(1) {
		v100 = v76
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v87 = v82 + int32(76)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+43)))
	if v88 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v75 == int32(0) {
		v100 = v76
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v75 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v93
	v103 = v93
	goto L16
L24:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v96
	goto L26
L25:
	;
	goto L26
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+44))
	v100 = v98
	goto L17
L27:
	;
	v107 = v103
	goto L30
L28:
	;
	goto L29
L29:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+14)))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+12)))
	v140 = v138 - v139
	if v140 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v119 = v107 + int32(4)
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+2)))
	if v120 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107))))
	base.MemoryCopy(m, v73+v121, v119, v120)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v124 = v119 + v120
	if base.Ui32(v124) < base.Ui32(v103+v104) {
		v107 = v124
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	base.MemoryFill(m, v139+v73, int32(0), v140)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = base.I32_wrap_i64(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = base.I32_wrap_i64(int64(base.Ui64(v21) >> (uint(int64(32)) % 64)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	F_MarkBufferDirty(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	goto L5
L40:
	;
	goto L4
L41:
	;
	v173 = v161
	v174 = v168
	v175 = int32(0)
	goto L42
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)+v175<<(uint(int32(2))%32))))
	if v189 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L1
L44:
	;
	F_UnlockReleaseBuffer(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	v193 = v173
	goto L46
L46:
	;
	v195 = v174 + int32(1)
	v197 = v195 & int32(255)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+72))
	if v197 <= v198 {
		v173 = v193
		v174 = v195
		v175 = v197
		goto L42
	} else {
		goto L48
	}
L47:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v193 = v192
	goto L46
L48:
	;
	goto L43
}
