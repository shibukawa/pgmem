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
	v4 = int32(_a_F_tuplestore_puttuple_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple[0])) = v7
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
			*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple[0])) = v5
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
	v5 = int32(_a_F_tuplestore_puttupleslot_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttupleslot[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttupleslot[0])) = v8
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
				*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttupleslot[0])) = v6
				return
			}
		}
	}
}
func F_tuplestore_set_eflags(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_tuplestore_set_eflags_0), int32(0))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_tuplestore_set_eflags_1), int32(376), int32(_a_F_tuplestore_set_eflags_2))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v10 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_tuplestore_set_eflags_0), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_tuplestore_set_eflags_1), int32(376), int32(_a_F_tuplestore_set_eflags_2))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v13 < int32(2) {
				v83 = l1
			} else {
				v17 = v13 - int32(1)
				v18 = int32(3)
				v19 = v17 & v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				if base.Ui32(v13-int32(2)) < base.Ui32(v18) {
					v56 = l1
					v57 = int32(1)
					v65 = v56
					v66 = v57
					v67 = int32(0)
					for {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v20+v66*int32(24))))
						v76 = v75 | v65
						v77 = int32(1)
						v80 = v67 + v77
						if v80 != v19 {
							v65 = v76
							v66 = v66 + v77
							v67 = v80
							continue
						} else {
							break
						}
						break
					}
					v83 = v76
				} else {
					v30 = l1
					v31 = int32(1)
					v36 = int32(0)
					for {
						v39 = v20 + v31*int32(24)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						v47 = v40 | (v41 | (v42 | (v43 | v30)))
						v48 = int32(4)
						v49 = v31 + v48
						v51 = v36 + v48
						if v51 != v17&int32(-4) {
							v30 = v47
							v31 = v49
							v36 = v51
							continue
						} else {
							break
						}
						break
					}
					if v19 == int32(0) {
						v83 = v47
					} else {
						v56 = v47
						v57 = v49
						v65 = v56
						v66 = v57
						v67 = int32(0)
						for {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v20+v66*int32(24))))
							v76 = v75 | v65
							v77 = int32(1)
							v80 = v67 + v77
							if v80 != v19 {
								v65 = v76
								v66 = v66 + v77
								v67 = v80
								continue
							} else {
								break
							}
							break
						}
						v83 = v76
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v83
			return
		}
	}
}
