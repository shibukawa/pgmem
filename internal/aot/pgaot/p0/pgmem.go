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
		F_ResourceOwnerForget(m, v9, v4, int32(_a_F_gen_pgmem_free_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			base.MemoryFill(m, v4, int32(0), int32(100))
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
		base.MemoryFill(m, v4, int32(0), int32(100))
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
func F_pgmem_shmctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v4
	v18 = m.Env.Pgmem_shmctl(m, l0, l1, v8+int32(12), v8+int32(8))
	mBase = m.M
	if v18 < v4 {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_shmctl[0])) = int32(0) - v18
		v38 = int32(-1)
	} else {
		if base.B2i32(l2 == int32(0))|base.B2i32(l1 != int32(2)) != 0 {
			v38 = v4
		} else {
			base.MemoryFill(m, l2, int32(0), int32(88))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v36
			v38 = v4
		}
	}
	m.G0 = v8 + int32(16)
	return v38
}
func F_pgmem_sigprocmask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v15
	goto L3
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0])) = v18
	goto L6
L4:
	;
	goto L10
L6:
	;
	goto L7
L7:
	;
	v41 = int32(_a_F_pgmem_sigprocmask_0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0])) = v42 & base.I32_rotl(int32(-2), int32(8))
	goto L4
L8:
	;
	v77 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v69 = int32(_a_F_pgmem_sigprocmask_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[0])) = v70 & base.I32_rotl(int32(-2), int32(18))
	goto L8
L12:
	;
	v84 = v77 - int32(1)
	if base.Ui32(v84) <= base.Ui32(int32(63)) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	m.G0 = v12 + int32(16)
	v152 = int32(_a_F_pgmem_sigprocmask_1)
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[1]))
	v155 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[1])) = v155
	if v153 == v155 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v146 = v77 + int32(1)
	if v146 != int32(65) {
		v77 = v146
		goto L12
	} else {
		goto L31
	}
L15:
	;
	if v97 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v84)>>(uint(int32(3))%32))&int32(536870908))+uint32(_c_F_pgmem_sigprocmask[2])))
	v97 = int32(base.Ui32(v92)>>(uint(v84)%32)) & int32(1)
	goto L18
L17:
	;
	v97 = int32(0)
	goto L18
L18:
	;
	goto L15
L19:
	;
	v102 = v77 - int32(1)
	if base.Ui32(v102) <= base.Ui32(int32(63)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v115 != 0 {
		goto L14
	} else {
		goto L24
	}
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v102)>>(uint(int32(3))%32))&int32(536870908))+uint32(_c_F_pgmem_sigprocmask[0])))
	v115 = int32(base.Ui32(v110)>>(uint(v102)%32)) & int32(1)
	goto L23
L22:
	;
	v115 = int32(0)
	goto L23
L23:
	;
	goto L20
L24:
	;
	v119 = v77 - int32(1)
	if base.B2i32(base.Ui32(v119) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v77-int32(32))) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_raise(m, v77)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[3])) = int32(28)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v135 = int32(base.Ui32(v119)>>(uint(int32(3))%32)) & int32(536870908)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pgmem_sigprocmask[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pgmem_sigprocmask[2]))) = v137 & base.I32_rotl(int32(-2), v119)
	goto L25
L29:
	;
	return
L30:
	;
	goto L14
L31:
	;
	goto L13
L32:
	;
	m.G0 = v8 + int32(32)
	return
L33:
	;
	v160 = v153
	v161 = int32(1)
	goto L34
L34:
	;
	v166 = int32(1) << (uint(v161) % 32)
	if v166&v160 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L32
L36:
	;
	v172 = v8 + int32(12)
	if base.Ui32(int32(65)) <= base.Ui32(v161) {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v214 = v160
	goto L38
L38:
	;
	if base.Ui32(int32(30)) < base.Ui32(v161) {
		goto L32
	} else {
		goto L55
	}
L39:
	;
	v214 = v160 & (v166 ^ int32(-1))
	goto L38
L40:
	;
	F_raise(m, v161)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L29
	} else {
		goto L54
	}
L41:
	;
	if v201 != 0 {
		goto L40
	} else {
		goto L51
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[3])) = int32(28)
	v201 = int32(-1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v172 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v181 = v161 * int32(20)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+uint32(_c_F_pgmem_sigprocmask[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+16)) = v182
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v181)+uint32(_c_F_pgmem_sigprocmask[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v172)+8)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v181)+uint32(_c_F_pgmem_sigprocmask[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v172))) = v186
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L49
L49:
	;
	goto L50
L50:
	;
	v201 = int32(0)
	goto L41
L51:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
	if v202&int32(4) != 0 {
		goto L40
	} else {
		goto L52
	}
L52:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v205 != 0 {
		goto L40
	} else {
		goto L53
	}
L53:
	;
	v206 = int32(_a_F_pgmem_sigprocmask_1)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_sigprocmask[1])) = v208 | v166
	goto L39
L54:
	;
	goto L39
L55:
	;
	if v214 != 0 {
		v160 = v214
		v161 = v161 + int32(1)
		goto L34
	} else {
		goto L56
	}
L56:
	;
	goto L35
}
