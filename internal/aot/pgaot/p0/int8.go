package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
			v52 = int32(1)
			v53 = v42 + v52
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v56 = v54 & v52
			if v54 == v52 {
				v59 = int32(4)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v61&int32(254) == int32(2) {
					v70 = v59
				} else {
					v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
				}
				if v61 == int32(1) {
					v73 = v59
				} else {
					v73 = v70
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v84
			if v56 != 0 {
				v90 = v53
			} else {
				v90 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v90
			v93 = F_palloc0(m, int32(48))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v95)
				v99 = F_pq_getmsgint64(m, v8+int32(32))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v99
					F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						F_numericvar_to_int128(m, v8+int32(8), v93+int32(16))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_pq_getmsgend(m, v8+int32(32))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								if v118 != 0 {
									F_pfree(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v93
									}
								} else {
									m.G0 = v8 + int32(48)
									return v93
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v128 = m.ExcPending
		if v128 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(60782), int32(0))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494988), int32(6055), int32(338567))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
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
