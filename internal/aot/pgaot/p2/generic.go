package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generic_identify(m *base.Module, l0 int32) int32 {
	return int32(513572)
}
func F_generic_redo(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	if v17 < v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = v16
	v27 = v2
	v31 = v2
	goto L3
L3:
	;
	v40 = v14 + int32(16) + v27<<(uint(int32(2))%32)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v27*int32(52))+76)))
	if v44 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v166 = int32(0)
	if v160 < v166 {
		goto L1
	} else {
		goto L40
	}
L5:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	v162 = v31 + int32(1)
	v164 = v162 & int32(255)
	if v164 <= v160 {
		v26 = v159
		v27 = v164
		v31 = v162
		goto L3
	} else {
		goto L39
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v50 = v31 & int32(255)
	v51 = F_XLogReadBufferForRedo(m, l0, v50, v40)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v53 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v73 = v14 + int32(12)
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
	if v76 < v50 {
		v98 = v74
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v53^int32(-1))<<(uint(int32(2))%32))))
	v71 = v63
	goto L12
L14:
	;
	goto L15
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v71 = v65 + v53<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v103 = v101 + v102
	if base.Ui32(v101) < base.Ui32(v103) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v101 = v98
	goto L16
L18:
	;
	v82 = v75 + v50*int32(52) + int32(76)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 != int32(1) {
		v98 = v74
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+43)))
	if v86 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v73 == int32(0) {
		v98 = v74
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v91
	v101 = v91
	goto L16
L24:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v94
	goto L26
L25:
	;
	goto L26
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v98 = v96
	goto L17
L27:
	;
	v106 = v101
	goto L30
L28:
	;
	goto L29
L29:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+14)))
	v142 = F__emscripten_memset_bulkmem(m, v71+v136, base.I32_extend8_s(int32(0)), v139-v136)
	mBase = m.M
	goto L37
L30:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106))))
	v119 = v106 + int32(4)
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)))
	if v120 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L29
L32:
	;
	v123 = v120 + v119
	if base.Ui32(v123) < base.Ui32(v103) {
		v106 = v123
		goto L30
	} else {
		goto L36
	}
L33:
	;
	v121 = F__emscripten_memcpy_bulkmem(m, v71+v116, v119, v120)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L31
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = base.I32_wrap_i64(v20)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = base.I32_wrap_i64(int64(base.Ui64(v20) >> (uint(int64(32)) % 64)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_MarkBufferDirty(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L5
L39:
	;
	goto L4
L40:
	;
	v171 = v159
	v172 = v166
	v173 = int32(0)
	goto L41
L41:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(16)+v172<<(uint(int32(2))%32))))
	if v186 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L1
L43:
	;
	F_UnlockReleaseBuffer(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L9
	} else {
		goto L46
	}
L44:
	;
	v190 = v171
	goto L45
L45:
	;
	v192 = v173 + int32(1)
	v194 = v192 & int32(255)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
	if v194 <= v195 {
		v171 = v190
		v172 = v194
		v173 = v192
		goto L41
	} else {
		goto L47
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v190 = v189
	goto L45
L47:
	;
	goto L42
}
