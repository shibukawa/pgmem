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
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltxtq_rexec_0), int32(0), v4, v5)
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
					F_errmsg(m, int32(_a_F_ltxtq_send_0), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(_a_F_ltxtq_send_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ltxtq_send_2), int32(614), int32(_a_F_ltxtq_send_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
			v37 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v37
			v40 = v10 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v40
			v43 = F_palloc(m, v37)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v43
				v47 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v47)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v50 = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v40 + v49*v50
				F_infix_2(m, v7+v50, int32(1))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v60 = v7 + int32(32)
					F_pq_begintypsend(m, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_enlargeStringInfo(m, v60, int32(1))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
							v69 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v66+v67))) = uint8(v69)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v66 + v69
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							v75 = F_strlen(m, v74)
							mBase = m.M
							F_pq_sendtext(m, v60, v74, v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								F_pfree(m, v78)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = v83 << (uint(int32(2)) % 32)
									m.G0 = v7 + int32(48)
									return v82
								}
							}
						}
					}
				}
			}
		}
	}
}
