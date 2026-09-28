package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyLoadRawBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v9 = v7 - v8
	if base.B2i32(v8 <= v2)|base.B2i32(v9 <= v2) == v2 {
		if v9 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
			base.MemoryCopy(m, v17, v17+v8, v9)
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
		v24 = v21 - v22
	} else {
		v24 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = v24
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v28 == v29 {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = int32(0)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v34 - v31
	} else {
	}
	v41 = F_CopyGetData(m, l0, v24+v28, int32(_a_F_CopyLoadRawBuf_0)-v24)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return
	} else {
		v43 = v41 + v9
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
		v46 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v43+v44))) = uint8(v46)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = v43
		v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+360))
		v51 = v49 + base.I64_extend_i32_s(v41)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+360)) = v51
		v56 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[0]))
		if v56 == v46 {
		} else {
			v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[1])))
			if v60&int32(1) == int32(0) {
			} else {
				v65 = int32(_a_F_CopyLoadRawBuf_1)
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2]))
				v68 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2])) = v67 + v68
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v71 + v68
				v75 = int32(0)
				v77 = int32(_a_F_CopyLoadRawBuf_2)
				v78 = base.AtomicRmwOr32(m, v75, v77, v75)
				*(*int64)(unsafe.Add(mBase, uint32(v56+v75)+232)) = v51
				v86 = base.AtomicRmwOr32(m, v75, v77, v75)
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v87 + v68
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2])) = v93 - v68
			}
		}
		if v41 == int32(0) {
			v99 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)) = uint8(v99)
		} else {
		}
		return
	}
}
func F_get_raw_page_fork_1_9(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v14 = F_text_to_cstring(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = F_forkname_to_number(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					if base.Ui64(int64(4294967295)) <= base.Ui64(v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_get_raw_page_fork_1_9_0), int32(0))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_get_raw_page_fork_1_9_1), int32(113), int32(_a_F_get_raw_page_fork_1_9_2))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v37 = F_get_raw_page_internal(m, v6, v16, base.I32_wrap_i64(v13))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v37)
						}
					}
				}
			}
		}
	}
}
