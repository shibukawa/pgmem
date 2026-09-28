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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v16 < v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = v15
	v24 = v2
	v28 = v2
	goto L3
L3:
	;
	v36 = v13 + int32(16) + v24<<(uint(int32(2))%32)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v24*int32(52))+76)))
	if v40 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v155 = int32(0)
	if v149 < v155 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	v151 = v28 + int32(1)
	v153 = v151 & int32(255)
	if v153 <= v149 {
		v23 = v148
		v24 = v153
		v28 = v151
		goto L3
	} else {
		goto L40
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v46 = v28 & int32(255)
	v47 = F_XLogReadBufferForRedo(m, l0, v46, v36)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v47 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v49 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v69 = v13 + int32(12)
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	if v72 < v46 {
		v94 = v70
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_generic_redo[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v49^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L12
L14:
	;
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_generic_redo[1]))
	v67 = v61 + v49<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v98 != 0 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v97 = v94
	goto L16
L18:
	;
	v76 = v71 + v46*int32(52)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+76)))
	if v77 != int32(1) {
		v94 = v70
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v81 = v76 + int32(76)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+43)))
	if v82 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v69 == int32(0) {
		v94 = v70
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v69 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v87
	v97 = v87
	goto L16
L24:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v90
	goto L26
L25:
	;
	goto L26
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v94 = v92
	goto L17
L27:
	;
	v101 = v97
	goto L30
L28:
	;
	goto L29
L29:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+14)))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	v130 = v128 - v129
	if v130 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v111 = v101 + int32(4)
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+2)))
	if v112 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	base.MemoryCopy(m, v67+v113, v111, v112)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v116 = v111 + v112
	if base.Ui32(v116) < base.Ui32(v97+v98) {
		v101 = v116
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	base.MemoryFill(m, v129+v67, int32(0), v130)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotl(v19, int64(32))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_MarkBufferDirty(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	v160 = v148
	v161 = v155
	v162 = int32(0)
	goto L42
L42:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(16)+v162<<(uint(int32(2))%32))))
	if v174 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L1
L44:
	;
	F_UnlockReleaseBuffer(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	v178 = v160
	goto L46
L46:
	;
	v180 = v161 + int32(1)
	v182 = v180 & int32(255)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v178)+72))
	if v182 <= v183 {
		v160 = v178
		v161 = v180
		v162 = v182
		goto L42
	} else {
		goto L48
	}
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v178 = v177
	goto L46
L48:
	;
	goto L43
}
