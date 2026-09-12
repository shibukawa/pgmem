package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BumpAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	v11 = (l1 + int32(7)) & int32(-8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v12) < base.Ui32(v11) {
		if base.Ui32(int32(1073741824)) <= base.Ui32(l1) {
			if l1 < int32(0) {
				F_MemoryContextSizeFailure(m, l1)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l2&int32(1) == int32(0) {
					F_MemoryContextSizeFailure(m, l1)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v27 = (l1+int32(7))&int32(-8) + int32(16)
					v28 = F_emscripten_builtin_malloc(m, v27)
					mBase = m.M
					if v28 == int32(0) {
						v31 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v57 = v31
							return v57
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v27
						v38 = v28 + v27
						*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v38
						v42 = l0 + int32(60)
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						if v43 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v46 = v44
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v42
							v46 = v42
						}
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v42))) = v28
						v57 = v28 + int32(16)
						return v57
					}
				}
			}
		} else {
			v27 = (l1+int32(7))&int32(-8) + int32(16)
			v28 = F_emscripten_builtin_malloc(m, v27)
			mBase = m.M
			if v28 == int32(0) {
				v31 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v57 = v31
					return v57
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v27
				v38 = v28 + v27
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v38
				v42 = l0 + int32(60)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v43 != 0 {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v46 = v44
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v42
					v46 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v46
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v42
				*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v42))) = v28
				v57 = v28 + int32(16)
				return v57
			}
		}
	} else {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
		if base.Ui32(v60-v61) < base.Ui32(v11) {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v66 = v64 << (uint(int32(1)) % 32)
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if base.Ui32(v66) < base.Ui32(v67) {
				v69 = v66
			} else {
				v69 = v67
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v69
			v72 = v11 + int32(16)
			if base.Ui32(v64) < base.Ui32(v72) {
				v74 = int32(1)
				if v72&(v72-v74) != 0 {
					v82 = v74 << (uint(int32(32)-base.I32_clz(v72)) % 32)
				} else {
					v82 = v72
				}
				v83 = v82
			} else {
				v83 = v64
			}
			v84 = F_emscripten_builtin_malloc(m, v83)
			mBase = m.M
			if v84 == int32(0) {
				v87 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					v114 = v87
					return v114
				}
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + v83
				*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v83 + v84
				*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v84 + int32(16)
				v98 = l0 + int32(60)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v99 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v98))) = v98
					v103 = v98
				} else {
					v103 = v99
				}
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = v98
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v103
				*(*int32)(unsafe.Add(mBase, uint32(v103))) = v84
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v84
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v108 + v11
				v114 = v108
				return v114
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v11 + v61
			return v61
		}
	}
}
func F_BumpDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v6 == int32(0) {
	} else {
		v10 = l0 + int32(60)
		if v6 == v10 {
		} else {
			v15 = v6
			for {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				if v15 == l0+int32(72) {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v15 + int32(16)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v19
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + (v15 - v29)
					F_emscripten_builtin_free(m, v15)
					mBase = m.M
				}
				if v19 != v10 {
					v15 = v19
					continue
				} else {
					break
				}
				break
			}
		}
	}
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return
}
func F_BumpReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v6 == int32(0) {
	} else {
		v10 = l0 + int32(60)
		if v6 == v10 {
		} else {
			v15 = v6
			for {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				if v15 == l0+int32(72) {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v15 + int32(16)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v19
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + (v15 - v29)
					F_emscripten_builtin_free(m, v15)
					mBase = m.M
				}
				if v19 != v10 {
					v15 = v19
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v40
	return
}
func F_BumpStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(224)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v17 == v6 {
		v53 = v6
		v54 = v6
		v55 = v6
	} else {
		v21 = l0 + int32(60)
		if v17 == v21 {
			v53 = v6
			v54 = v6
			v55 = v6
		} else {
			v29 = v17
			v30 = v6
			v31 = v6
			v32 = v6
			for {
				v36 = v32 + int32(1)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				v39 = v37 + (v30 - v29)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v42 = v31 + v37 - v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				if v43 != v21 {
					v29 = v43
					v30 = v39
					v31 = v42
					v32 = v36
					continue
				} else {
					break
				}
				break
			}
			v53 = v39
			v54 = v42
			v55 = v36
		}
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v55
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v54
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v53
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v53 - v54
		v67 = F_pg_snprintf(m, v15+int32(16), int32(200), int32(471202), v15)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v15+int32(16), l4)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				if l3 != 0 {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v73 + v55
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v76 + v53
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v79 + v54
				} else {
				}
				m.G0 = v15 + int32(224)
				return
			}
		}
	} else {
		if l3 != 0 {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v73 + v55
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v76 + v53
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v79 + v54
		} else {
		}
		m.G0 = v15 + int32(224)
		return
	}
}
