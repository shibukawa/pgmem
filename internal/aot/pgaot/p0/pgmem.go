package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if int32(0) < v5 {
		m.Env.Pgmem_cipher_free(m, v5)
		mBase = m.M
	} else {
	}
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+96))
	if v9 != 0 {
		F_ResourceOwnerForget(m, v9, v4, int32(4336172))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v15 = F___memset(m, v4, int32(0), int32(100))
			mBase = m.M
			F_pfree(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v15 = F___memset(m, v4, int32(0), int32(100))
		mBase = m.M
		F_pfree(m, v4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pgmem_main(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	v3 = int32(0)
	if l1 <= v3 {
		v82 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = v82 << (uint(int32(2)) % 32)
	v92 = F_emscripten_builtin_malloc(m, v89+int32(4))
	mBase = m.M
	if v82 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v13 = l1 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v3
	v21 = v3
	v25 = v3
	goto L6
L4:
	;
	v51 = v3
	v52 = v3
	goto L5
L5:
	;
	if v13 == int32(0) {
		v82 = v52
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v27 = l0 + v20
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v29 = int32(0)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+3)))
	v43 = v21 + base.B2i32(v28 == v29) + base.B2i32(v32 == v29) + base.B2i32(v36 == v29) + base.B2i32(v40 == v29)
	v44 = int32(4)
	v45 = v20 + v44
	v47 = v25 + v44
	if v47 != l1&int32(2147483644) {
		v20 = v45
		v21 = v43
		v25 = v47
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v51 = v45
	v52 = v43
	goto L5
L8:
	;
	goto L7
L9:
	;
	v62 = v51
	v63 = v52
	v66 = v3
	goto L10
L10:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v62))))
	v73 = v63 + base.B2i32(v70 == int32(0))
	v74 = int32(1)
	v77 = v66 + v74
	if v77 != v13 {
		v62 = v62 + v74
		v63 = v73
		v66 = v77
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v82 = v73
	goto L1
L12:
	;
	goto L11
L13:
	;
	v93 = l0
	v97 = v3
	goto L16
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v92))) = int32(0)
	v181 = F_main(m, v82, v92)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L36
	} else {
		goto L37
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92+v97<<(uint(int32(2))%32)))) = v93
	if v93&int32(3) == int32(0) {
		v129 = v93
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L15
L18:
	;
	v164 = int32(1)
	v167 = v97 + v164
	if v167 != v82 {
		v93 = v162 + v93 + v164
		v97 = v167
		goto L16
	} else {
		goto L35
	}
L19:
	;
	v162 = v154 - v93
	goto L18
L20:
	;
	v133 = v129
	goto L29
L21:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v113 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v162 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v118 = v93
	goto L25
L25:
	;
	v122 = v118 + int32(1)
	if v122&int32(3) == int32(0) {
		v129 = v122
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v154 = v122
	goto L19
L27:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v127 != 0 {
		v118 = v122
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v142 = int32(-2139062144)
	if (int32(16843008)-v139|v139)&v142 == v142 {
		v133 = v133 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v148 = v133
	goto L32
L31:
	;
	goto L30
L32:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 != 0 {
		v148 = v148 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v154 = v148
	goto L19
L34:
	;
	goto L33
L35:
	;
	goto L17
L36:
	;
	return int32(0)
L37:
	;
	return v181
}
