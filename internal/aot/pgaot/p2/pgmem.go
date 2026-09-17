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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var __phi94 int32
	_ = __phi94
	var v97 int32
	_ = v97
	var __phi97 int32
	_ = __phi97
	var v98 int32
	_ = v98
	var __phi98 int32
	_ = __phi98
	var v99 int32
	_ = v99
	var __phi99 int32
	_ = __phi99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(288)
	m.G0 = v10
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v212 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(288)
	return v212
L2:
	;
	v17 = int32(_a_F_pgmem_dlsym_0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgmem_dlsym[0])))
	if base.B2i32(v20 == int32(0))|base.B2i32(v20 != v23) != 0 {
		v41 = v20
		v42 = v23
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v147 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L4:
	;
	if v41-v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	v26 = l1
	v27 = v17
	goto L7
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v31
		v42 = v30
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v41 = v31
	v42 = v30
	goto L5
L9:
	;
	v34 = int32(1)
	if v31 == v30 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v44 = int32(_a_F_pgmem_dlsym_1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgmem_dlsym[1])))
	if base.B2i32(v47 == int32(0))|base.B2i32(v47 != v50) != 0 {
		v68 = v47
		v69 = v50
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v71 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v79 = F_snprintf(m, v10+int32(32), int32(256), int32(_a_F_pgmem_dlsym_2), v10+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v68-v69 != 0 {
		v141 = l1
		goto L3
	} else {
		goto L21
	}
L15:
	;
	goto L14
L16:
	;
	v53 = l1
	v54 = v44
	goto L17
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v58
		v69 = v57
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v68 = v58
	v69 = v57
	goto L15
L19:
	;
	v61 = int32(1)
	if v58 == v57 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L13
L22:
	;
	return int32(0)
L23:
	;
	v84 = v71 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v86 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v136 = v10 + int32(32)
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136+v130))) = uint8(v138)
	v141 = v136
	goto L3
L25:
	;
	v130 = v84
	goto L24
L26:
	;
	goto L27
L27:
	;
	v90 = v71 + int32(2)
	if base.Ui32(int32(255)) < base.Ui32(v90) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v130 = v84
	goto L24
L29:
	;
	goto L30
L30:
	;
	__phi94 = v86
	__phi97 = v84
	__phi98 = v90
	__phi99 = v85
	v94 = __phi94
	v97 = __phi97
	v98 = __phi98
	v99 = __phi99
	goto L31
L31:
	;
	v100 = int32(32)
	v105 = v94 & int32(255)
	goto L33
L32:
	;
	v130 = v98
	goto L24
L33:
	;
	if base.B2i32(base.Ui32(v105-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v105|v100-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v117 = v94
	goto L36
L35:
	;
	v117 = int32(95)
	goto L36
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v100+v97))) = uint8(v117)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v119 == int32(0) {
		v130 = v98
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v122 = int32(1)
	if base.Ui32(v98) < base.Ui32(int32(255)) {
		__phi94 = v119
		__phi97 = v98
		__phi98 = v98 + v122
		__phi99 = v99 + v122
		v94 = __phi94
		v97 = __phi97
		v98 = __phi98
		v99 = __phi99
		goto L31
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v141
	v204 = F_snprintf(m, int32(_a_F_pgmem_dlsym_3), int32(512), int32(_a_F_pgmem_dlsym_4), v10)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L22
	} else {
		goto L54
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v154 = int32(0)
	goto L41
L41:
	;
	v161 = v150 + v154<<(uint(int32(3))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if base.B2i32(v165 == int32(0))|base.B2i32(v165 != v168) != 0 {
		v186 = v165
		v187 = v168
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v212 = v192
	goto L1
L43:
	;
	if v186-v187 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	v171 = v162
	v172 = v141
	goto L46
L46:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v176 == int32(0) {
		v186 = v176
		v187 = v175
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v186 = v176
	v187 = v175
	goto L44
L48:
	;
	v179 = int32(1)
	if v176 == v175 {
		v171 = v171 + v179
		v172 = v172 + v179
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v190 = v154 + int32(1)
	if v147 != v190 {
		v154 = v190
		goto L41
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L42
L53:
	;
	goto L39
L54:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgmem_dlsym[2])) = uint8(v208)
	v212 = int32(0)
	goto L1
}
func F_pgmem_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	v4 = int32(0)
	if l1 == v4 {
	} else {
		v11 = l1 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(l1) {
			v19 = v4
			v22 = v4
			for {
				v25 = l0 + v19<<(uint(int32(3))%32)
				v26 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+62)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+54)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+46)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+38)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+30)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+22)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+14)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v26)
				v42 = int32(8)
				v43 = v19 + v42
				v45 = v22 + v42
				if v45 != l1&int32(-8) {
					v19 = v43
					v22 = v45
					continue
				} else {
					break
				}
				break
			}
			if v11 == int32(0) {
			} else {
				v52 = v43
				v58 = int32(0)
				v60 = v52
				for {
					v67 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0+v60<<(uint(int32(3))%32))+6)) = uint16(v67)
					v69 = int32(1)
					v72 = v58 + v69
					if v72 != v11 {
						v58 = v72
						v60 = v60 + v69
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v52 = v4
			v58 = int32(0)
			v60 = v52
			for {
				v67 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0+v60<<(uint(int32(3))%32))+6)) = uint16(v67)
				v69 = int32(1)
				v72 = v58 + v69
				if v72 != v11 {
					v58 = v72
					v60 = v60 + v69
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v81 = m.Env.Pgmem_poll(m, l2)
	mBase = m.M
	return int32(0)
}
