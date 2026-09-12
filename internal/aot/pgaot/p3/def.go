package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_defGetBoolean(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(0) {
		v221 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v221
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v13 == int32(465) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v221 = int32(1)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L66
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v17 {
	case 0:
		v221 = int32(0)
		goto L1
	case 1:
		goto L3
	default:
		goto L4
	}
L6:
	;
	goto L7
L7:
	;
	v18 = F_defGetString(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v25 = v18
	v26 = int32(342245)
	goto L11
L10:
	;
	if v63 == int32(0) {
		v221 = v9
		goto L1
	} else {
		goto L23
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v29 == v30 {
		v52 = v29
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v63 = int32(0)
	goto L10
L13:
	;
	v54 = int32(1)
	if v52 != 0 {
		v25 = v25 + v54
		v26 = v26 + v54
		goto L11
	} else {
		goto L22
	}
L14:
	;
	if base.Ui32((v29-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = v29 | int32(32)
	goto L17
L16:
	;
	v40 = v29
	goto L17
L17:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v30 | int32(32)
	goto L20
L19:
	;
	v49 = v30
	goto L20
L20:
	;
	if v40 == v49 {
		v52 = v40
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v63 = v40 - v49
	goto L10
L22:
	;
	goto L12
L23:
	;
	v70 = v18
	v71 = int32(359138)
	goto L25
L24:
	;
	if v108 == int32(0) {
		v221 = int32(0)
		goto L1
	} else {
		goto L37
	}
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == v75 {
		v97 = v74
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v108 = int32(0)
	goto L24
L27:
	;
	v99 = int32(1)
	if v97 != 0 {
		v70 = v70 + v99
		v71 = v71 + v99
		goto L25
	} else {
		goto L36
	}
L28:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v85 = v74 | int32(32)
	goto L31
L30:
	;
	v85 = v74
	goto L31
L31:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v94 = v75 | int32(32)
	goto L34
L33:
	;
	v94 = v75
	goto L34
L34:
	;
	if v85 == v94 {
		v97 = v85
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v108 = v85 - v94
	goto L24
L36:
	;
	goto L26
L37:
	;
	v115 = v18
	v116 = int32(271632)
	goto L39
L38:
	;
	if v153 == int32(0) {
		v221 = int32(1)
		goto L1
	} else {
		goto L51
	}
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v119 == v120 {
		v142 = v119
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v153 = int32(0)
	goto L38
L41:
	;
	v144 = int32(1)
	if v142 != 0 {
		v115 = v115 + v144
		v116 = v116 + v144
		goto L39
	} else {
		goto L50
	}
L42:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v119 | int32(32)
	goto L45
L44:
	;
	v130 = v119
	goto L45
L45:
	;
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v139 = v120 | int32(32)
	goto L48
L47:
	;
	v139 = v120
	goto L48
L48:
	;
	if v130 == v139 {
		v142 = v130
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v153 = v130 - v139
	goto L38
L50:
	;
	goto L40
L51:
	;
	v160 = v18
	v161 = int32(336934)
	goto L53
L52:
	;
	if v198 == int32(0) {
		v221 = int32(0)
		goto L1
	} else {
		goto L65
	}
L53:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == v165 {
		v187 = v164
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v198 = int32(0)
	goto L52
L55:
	;
	v189 = int32(1)
	if v187 != 0 {
		v160 = v160 + v189
		v161 = v161 + v189
		goto L53
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v175 = v164 | int32(32)
	goto L59
L58:
	;
	v175 = v164
	goto L59
L59:
	;
	if base.Ui32((v165-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v184 = v165 | int32(32)
	goto L62
L61:
	;
	v184 = v165
	goto L62
L62:
	;
	if v175 == v184 {
		v187 = v175
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v198 = v175 - v184
	goto L52
L64:
	;
	goto L54
L65:
	;
	goto L4
L66:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v210
	F_errmsg(m, int32(343697), v7)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(495712), int32(141), int32(282472))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_defGetInt32(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v9 == int32(465) {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			m.G0 = v6 + int32(32)
			return v50
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v21
					F_errmsg(m, int32(343628), v6+int32(16))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495712), int32(164), int32(550141))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v40
				F_errmsg(m, int32(343628), v6)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495712), int32(155), int32(550141))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
