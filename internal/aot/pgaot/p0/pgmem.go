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
		F_ResourceOwnerForget(m, v9, v4, int32(4432524))
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	v3 = int32(0)
	if l1 <= v3 {
		v82 = v3
	} else {
		v13 = l1 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(l1) {
			v20 = v3
			v21 = v3
			v25 = v3
			for {
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
					continue
				} else {
					break
				}
				break
			}
			v51 = v45
			v52 = v43
		} else {
			v51 = v3
			v52 = v3
		}
		if v13 == int32(0) {
			v82 = v52
		} else {
			v62 = v51
			v63 = v52
			v66 = v3
			for {
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v62))))
				v73 = v63 + base.B2i32(v70 == int32(0))
				v74 = int32(1)
				v77 = v66 + v74
				if v77 != v13 {
					v62 = v62 + v74
					v63 = v73
					v66 = v77
					continue
				} else {
					break
				}
				break
			}
			v82 = v73
		}
	}
	v89 = v82 << (uint(int32(2)) % 32)
	v92 = F_emscripten_builtin_malloc(m, v89+int32(4))
	mBase = m.M
	if v82 != 0 {
		v93 = l0
		v97 = v3
		for {
			*(*int32)(unsafe.Add(mBase, uint32(v92+v97<<(uint(int32(2))%32)))) = v93
			v106 = F_strlen(m, v93)
			mBase = m.M
			v108 = int32(1)
			v111 = v97 + v108
			if v111 != v82 {
				v93 = v106 + v93 + v108
				v97 = v111
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	*(*int32)(unsafe.Add(mBase, uint32(v89+v92))) = int32(0)
	v125 = F_main(m, v82, v92)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		return int32(0)
	} else {
		return v125
	}
}
