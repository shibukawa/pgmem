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
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v9 = v7 - v8
	if base.B2i32(v8 <= v2)|base.B2i32(v9 <= v2) == v2 {
		if v9 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
			base.MemoryCopy(m, v17, v17+v8, v9)
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
		v24 = v21 - v22
	} else {
		v24 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v24
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v28 == v29 {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = int32(0)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v34 - v31
	} else {
	}
	v41 = F_CopyGetData(m, l0, v24+v28, int32(_a_F_CopyLoadRawBuf_0)-v24)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return
	} else {
		v43 = v41 + v9
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
		v46 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v43+v44))) = uint8(v46)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v43
		v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
		v51 = v49 + base.I64_extend_i32_s(v41)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v51
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
				*(*int64)(unsafe.Add(mBase, uint32(v56+int32(0))+232)) = v51
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v79 + v68
				v85 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2])) = v85 - v68
			}
		}
		if v41 == int32(0) {
			v91 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v91)
		} else {
		}
		return
	}
}
func F_get_raw_page_fork_1_9(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			v15 = F_text_to_cstring(m, v11)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_forkname_to_number(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if base.Ui64(int64(4294967295)) <= base.Ui64(v14) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_get_raw_page_fork_1_9_0), int32(0))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_raw_page_fork_1_9_1), int32(113), int32(_a_F_get_raw_page_fork_1_9_2))
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
					} else {
						v38 = F_get_raw_page_internal(m, v6, v17, base.I32_wrap_i64(v14))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							return v38
						}
					}
				}
			}
		}
	}
}
