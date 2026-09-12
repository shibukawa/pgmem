package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v3 != 0 {
		F_BufFileClose(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			F_MemoryContextDelete(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				F_pfree(m, v9)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					F_pfree(m, v12)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v16 = m.ExcPending
						if v16 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		F_MemoryContextDelete(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			F_pfree(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_tuplestore_puttuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = int32(4549024)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = m.T0[v9].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_tuplestore_puttuple_common(m, l0, v10)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[10])) = v5
			return
		}
	}
}
func F_tuplestore_puttupleslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	v5 = int32(4549024)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_GetMemoryChunkSpace(m, v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v17 - base.I64_extend_i32_u(v15)
			F_tuplestore_puttuple_common(m, l0, v13)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v6
				return
			}
		}
	}
}
func F_tuplestore_set_eflags(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(164784), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				F_errfinish(m, int32(519198), int32(376), int32(164801))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(164784), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return
				} else {
					F_errfinish(m, int32(519198), int32(376), int32(164801))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v16 < int32(2) {
				v104 = l1
			} else {
				v20 = v16 - int32(1)
				v21 = int32(3)
				v22 = v20 & v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				if base.Ui32(v16-int32(2)) < base.Ui32(v21) {
					v69 = l1
					v70 = int32(1)
				} else {
					v39 = l1
					v40 = int32(1)
					v44 = int32(0)
					for {
						v50 = v40 * int32(24)
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(72)+v50)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v23+int32(48)))))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v23+int32(24)))))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v50+v23)))
						v62 = v52 | (v54 | (v56 | (v58 | v39)))
						v63 = int32(4)
						v64 = v40 + v63
						v66 = v44 + v63
						if v66 != v20&int32(-4) {
							v39 = v62
							v40 = v64
							v44 = v66
							continue
						} else {
							break
						}
						break
					}
					v69 = v62
					v70 = v64
				}
				if v22 == int32(0) {
					v104 = v69
				} else {
					v83 = v69
					v84 = v70
					v85 = int32(0)
					for {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v23+v84*int32(24))))
						v97 = v96 | v83
						v98 = int32(1)
						v101 = v85 + v98
						if v101 != v22 {
							v83 = v97
							v84 = v84 + v98
							v85 = v101
							continue
						} else {
							break
						}
						break
					}
					v104 = v97
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v104
			return
		}
	}
}
