package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltxtq_rexec(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltxtq_rexec_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_ltxtq_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_ltxtq_send_0), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v30 = F_errdetail(m, int32(_a_F_ltxtq_send_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_ltxtq_send_2), int32(650), int32(_a_F_ltxtq_send_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
		} else {
			v37 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v37
			v40 = v10 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v40
			v44 = F_palloc_mul(m, int32(1), v37)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v44
				v48 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v48)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v51 = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v40 + v50*v51
				F_infix_2(m, v7+v51, int32(1))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					v61 = v7 + int32(32)
					F_pq_begintypsend(m, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						F_enlargeStringInfo(m, v61, int32(1))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v67+v68))) = uint8(v70)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v67 + v70
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							v76 = F_strlen(m, v75)
							mBase = m.M
							F_pq_sendtext(m, v61, v75, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int64(0)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								F_pfree(m, v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84 << (uint(int32(2)) % 32)
									m.G0 = v7 + int32(48)
									return base.I64_extend_i32_u(v83)
								}
							}
						}
					}
				}
			}
		}
	}
}
