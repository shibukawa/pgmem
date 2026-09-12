package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateSeqStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v27 = v3
			return v27
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v27 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v18 != v19 {
						v27 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						if v21 != v22 {
							v27 = v3
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
							v27 = base.B2i32(v24 == v25)
						}
					}
				}
				return v27
			}
		}
	}
}
func F_seq_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	v15 = v13 & int32(240)
	if v15 == int32(0) {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
		v21 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 < int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v21^int32(-1))<<(uint(int32(2))%32))))
				v40 = v32
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v40 = v34 + v21<<(uint(int32(13))%32) + int32(-8192)
			}
			v42 = F_palloc(m, int32(8192))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				if v42&int32(3) != 0 {
				} else {
				}
				v70 = F___memset(m, v42, int32(0), int32(8192))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v42)+10)) = int32(1572864)
				v76 = int32(8196)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+18)) = uint16(v76)
				v82 = int32(8184)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)) = uint16(v82)
				*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)) = uint16(v82)
				v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v42+v85))) = int32(5911)
				v89 = int32(12)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
				v97 = F_PageAddItemExtended(m, v42, v19+v89, v92-v89, int32(1), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					if v97 == int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(426601), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return
							} else {
								F_errfinish(m, int32(523052), int32(1930), int32(254552))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotr(v18, int64(32))
						v105 = F__emscripten_memcpy_bulkmem(m, v40, v42, int32(8192))
						mBase = m.M
						F_MarkBufferDirty(m, v21)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v21)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								F_pfree(m, v42)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
			F_errmsg_internal(m, int32(57294), v10)
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return
			} else {
				F_errfinish(m, int32(523052), int32(1905), int32(254552))
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
