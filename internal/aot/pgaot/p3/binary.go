package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_binary_quantize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v43 float32
	_ = v43
	var v47 float32
	_ = v47
	var v53 float32
	_ = v53
	var v59 float32
	_ = v59
	var v65 float32
	_ = v65
	var v71 float32
	_ = v71
	var v77 float32
	_ = v77
	var v83 float32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 float32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(8)
		v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
		v16 = F_InitBitVector(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(8)
			v19 = v16 + v18
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			v22 = base.I32_div_s(v20, v18)
			v23 = int32(0)
			if v18 <= v20 {
				v29 = v23
				for {
					v36 = int32(3)
					v39 = int32(2)
					v41 = v14 + v29<<(uint(v39)%32)
					v42 = *(*float32)(unsafe.Add(mBase, uint32(v41)))
					v43 = float32(0)
					v47 = *(*float32)(unsafe.Add(mBase, uint32(v41)+4))
					v53 = *(*float32)(unsafe.Add(mBase, uint32(v41)+8))
					v59 = *(*float32)(unsafe.Add(mBase, uint32(v41)+12))
					v65 = *(*float32)(unsafe.Add(mBase, uint32(v41)+16))
					v71 = *(*float32)(unsafe.Add(mBase, uint32(v41)+20))
					v77 = *(*float32)(unsafe.Add(mBase, uint32(v41)+24))
					v83 = *(*float32)(unsafe.Add(mBase, uint32(v41)+28))
					v86 = base.F32_gt(v42, v43)<<(uint(int32(7))%32) | base.F32_gt(v47, v43)<<(uint(int32(6))%32) | base.F32_gt(v53, v43)<<(uint(int32(5))%32) | base.F32_gt(v59, v43)<<(uint(int32(4))%32) | base.F32_gt(v65, v43)<<(uint(v36)%32) | base.F32_gt(v71, v43)<<(uint(v39)%32) | base.F32_gt(v77, v43)<<(uint(int32(1))%32) | base.F32_gt(v83, v43)
					*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v29)>>(uint(v36)%32))))) = uint8(v86)
					v89 = v29 + int32(8)
					if v89 < base.I32_extend16_s(v22)<<(uint(int32(3))%32) {
						v29 = v89
						continue
					} else {
						break
					}
					break
				}
				v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				v92 = v89
				v93 = v91
			} else {
				v92 = v23
				v93 = v20
			}
			if v92 < base.I32_extend16_s(v93) {
				v101 = v92
				for {
					v110 = v19 + int32(base.Ui32(v101)>>(uint(int32(3))%32))
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
					v115 = *(*float32)(unsafe.Add(mBase, uint32(v14+v101<<(uint(int32(2))%32))))
					v123 = v111 | base.F32_gt(v115, float32(0))<<(uint((v101^int32(-1))&int32(7))%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v123)
					v126 = v101 + int32(1)
					v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
					if v126 < v127 {
						v101 = v126
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v16
		}
	}
}
func F_binary_upgrade_set_next_pg_tablespace_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[710])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(432465), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516428), int32(46), int32(455983))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1037])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_pg_type_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[710])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(432465), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516428), int32(57), int32(455795))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1038])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_toast_relfilenode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[710])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(432465), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516428), int32(156), int32(430912))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[1039])) = v25
		return int32(0)
	}
}
