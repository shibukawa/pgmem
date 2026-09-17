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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
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
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v49 = v42 + v48
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v50 == v48 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				if v58 == int32(18) {
					v61 = int32(16)
				} else {
					v61 = int32(0)
				}
				if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v68 = int32(4)
				} else {
					v68 = v61
				}
				v79 = v68
			} else {
				v69 = int32(1)
				if v52 != 0 {
					v79 = int32(base.Ui32(v50)>>(uint(v69)%32)) - v69
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v79
			if v52 != 0 {
				v85 = v49
			} else {
				v85 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v85
			v88 = F_palloc0(m, int32(48))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v93 = v8 + int32(32)
				v94 = F_pq_getmsgint64(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v94
					v98 = v8 + int32(8)
					F_numericvar_deserialize(m, v93, v98)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						F_numericvar_to_int128(m, v98, v88+int32(16))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_pq_getmsgend(m, v93)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								if v107 != 0 {
									F_pfree(m, v107)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v88
									}
								} else {
									m.G0 = v8 + int32(48)
									return v88
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
		v117 = m.ExcPending
		if v117 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_deserialize_0), int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_deserialize_1), int32(_a_F_int8_avg_deserialize_2), int32(_a_F_int8_avg_deserialize_3))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
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
