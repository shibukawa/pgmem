package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_deserialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v41 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(435) {
		case 0:
			v41 = int32(1)
		case 1:
			v41 = int32(2)
		default:
			v41 = int32(0)
		}
	}
	if v41 != 0 {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v43 = F_pg_detoast_datum_packed(m, v42)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v76
			v80 = int32(1)
			if v47&v80 != 0 {
				v84 = v80
			} else {
				v84 = int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v43 + v84
			v88 = F_palloc0(m, int32(48))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v92 = F_pq_getmsgint64(m, v9)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v92
					v95 = F_pq_getmsgint64(m, v9)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int64(0)
					} else {
						v97 = F_pq_getmsgint64(m, v9)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v88)+24)) = v95
							*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v97
							F_pq_getmsgend(m, v9)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v88)
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_deserialize_0), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_deserialize_1), int32(_a_F_int8_avg_deserialize_2), int32(_a_F_int8_avg_deserialize_3))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
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
