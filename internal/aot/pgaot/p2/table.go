package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_table_block_parallelscan_initialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v4
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v6
	v9 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v9
		v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_table_block_parallelscan_initialize[0])))
		if v16 != int32(1) {
			v29 = int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+118)))
			if v21 == int32(116) {
				v29 = int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_table_block_parallelscan_initialize[1]))
				v27 = base.I32_div_s(v25, int32(4))
				v29 = base.B2i32(base.Ui32(v27) < base.Ui32(v9))
			}
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v29)
		v31 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+24)), uint32(v31))
		*(*int64)(unsafe.Add(mBase, uint32(l1)+40)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(-1)
		return int32(48)
	}
}
func F_table_block_parallelscan_nextpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v8 == int32(-1) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		v12 = v11
	} else {
		v12 = v8
	}
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v13 != 0 {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v35 = v13
		v36 = v14 + int64(1)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if base.Ui32(v17) < base.Ui32(int32(2)) {
			v29 = v17
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui64(v20) <= base.Ui64(base.I64_extend_i32_u(v12-v17<<(uint(int32(6))%32))) {
				v29 = v17
			} else {
				v27 = int32(base.Ui32(v17) >> (uint(int32(1)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v27
				v29 = v27
			}
		}
		v32 = base.AtomicRmwAdd64(m, l2, int32(40), base.I64_extend_i32_u(v29))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v35 = v33
		v36 = v32
	}
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v36
	v38 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v35 - int32(1)
	if base.Ui64(v36) < base.Ui64(base.I64_extend_i32_u(v12)) {
		v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+28)))
		v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+20)))
		v47 = base.I64_rem_u_s(v36+v44, v46)
		v48 = base.I32_wrap_i64(v47)
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
		if v49 != 0 {
			v57 = v48
			v58 = v48
			F_ss_report_location(m, l0, v57)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = v58
				return v63
			}
		} else {
			v63 = v48
			return v63
		}
	} else {
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
		if v50 != int32(1) {
			v63 = v38
			return v63
		} else {
			v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+20)))
			if v36 != v53 {
				v63 = v38
				return v63
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
				v57 = v55
				v58 = int32(-1)
				F_ss_report_location(m, l0, v57)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = v58
					return v63
				}
			}
		}
	}
}
func F_table_tuple_get_latest_tid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+188))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
					v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
					v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v26 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v24 | v25<<(uint(int32(16))%32)
					F_errmsg(m, int32(_a_F_table_tuple_get_latest_tid_0), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_table_tuple_get_latest_tid_1), int32(284), int32(_a_F_table_tuple_get_latest_tid_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
			m.T0[v44].(func(*base.Module, int32, int32))(m, l0, l1)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
