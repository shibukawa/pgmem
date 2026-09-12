package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltxtq_rexec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(5602), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ltxtq_send(m *base.Module, l0 int32) int32 {
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(212002), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(573008), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495824), int32(614), int32(426026))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
		} else {
			v42 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v42
			v45 = v10 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v45
			v48 = F_palloc(m, v42)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v48
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v52)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v55 = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v45 + v54*v55
				F_infix_2(m, v7+v55, int32(1))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_pq_begintypsend(m, v7+int32(32))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_enlargeStringInfo(m, v7+int32(32), int32(1))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
							v76 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v73+v74))) = uint8(v76)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v73 + v76
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							v84 = F_strlen(m, v83)
							mBase = m.M
							F_pq_sendtext(m, v7+int32(32), v83, v84)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								F_pfree(m, v87)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v91 = v7 + int32(32)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v93))) = v94 << (uint(int32(2)) % 32)
									m.G0 = v7 + int32(48)
									return v93
								}
							}
						}
					}
				}
			}
		}
	}
}
