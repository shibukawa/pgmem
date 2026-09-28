package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_summarizer_state(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = v10 + int32(-40)
	v17 = v10 + int32(-48)
	v19 = v10 + int32(-56)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_summarizer_state[0]))
	v27 = F_LWLockAcquire(m, v23+int32(_a_F_pg_get_wal_summarizer_state_0), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_summarizer_state[1]))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
		if v33 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
			v38 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v38
			*(*int64)(unsafe.Add(mBase, uint32(v19))) = v38
			v67 = int32(-1)
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v44
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_summarizer_state[1]))
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v48
			v50 = int32(-1)
			if v43 == v50 {
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = v48
				v67 = v50
			} else {
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v47)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = v54
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_summarizer_state[2]))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v43*int32(768))+12))
				if v63 <= int32(0) {
					v66 = int32(-1)
				} else {
					v66 = v63
				}
				v67 = v66
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10+int32(-60)))) = v67
		v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_summarizer_state[0]))
		F_LWLockRelease(m, v74+int32(_a_F_pg_get_wal_summarizer_state_0))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return int64(0)
		} else {
			v80 = F_get_call_result_type(m, l0, int32(0), v12)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int64(0)
			} else {
				if v80 == int32(1) {
					v84 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v84
					v86 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v86
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v88
					v90 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v90
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					if v92 < v84 {
						v95 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v95)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = base.I64_extend_i32_u(v92)
					}
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v104 = F_heap_form_tuple(m, v99, v10+int32(-32), v10+int32(-36))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int64(0)
					} else {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
						v107 = F_HeapTupleHeaderGetDatum(m, v106)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int64(0)
						} else {
							m.G0 = v12 - int32(-64)
							return v107
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_get_wal_summarizer_state_1), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_wal_summarizer_state_2), int32(194), int32(_a_F_pg_get_wal_summarizer_state_3))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int64(0)
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
	}
}
