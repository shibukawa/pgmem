package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_summarizer_state(m *base.Module, l0 int32) int32 {
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
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = v12 + int32(24)
	v17 = v12 + int32(16)
	v19 = v12 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v27 = F_LWLockAcquire(m, v23+int32(6272), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _consts[414]))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
		if v33 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
			v38 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v38
			*(*int64)(unsafe.Add(mBase, uint32(v19))) = v38
			v65 = int32(-1)
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v44
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v46
			v48 = int32(-1)
			if v43 == v48 {
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = v46
				v65 = v48
			} else {
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = v52
				v56 = *(*int32)(unsafe.Add(mBase, _consts[141]))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v43*int32(640))+44))
				if v61 <= int32(0) {
					v64 = int32(-1)
				} else {
					v64 = v61
				}
				v65 = v64
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)))) = v65
		v71 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		F_LWLockRelease(m, v71+int32(6272))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			v77 = F_get_call_result_type(m, l0, int32(0), v12)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				if v77 == int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(0)
					v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)))
					v84 = F_Int64GetDatum(m, v83)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v84
						v87 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v88 = F_Int64GetDatum(m, v87)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v88
							v91 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
							v92 = F_Int64GetDatum(m, v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v92
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								if v95 < int32(0) {
									v98 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v98)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v95
								}
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v106 = F_heap_form_tuple(m, v101, v12+int32(32), v12+int32(28))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
									v109 = F_HeapTupleHeaderGetDatum(m, v108)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(48)
										return v109
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(363732), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489535), int32(192), int32(348176))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
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
