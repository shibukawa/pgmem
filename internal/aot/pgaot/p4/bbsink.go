package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bbsink_copystream_end_archive(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v10
	v15 = m.G0
	v17 = v15 - v6
	m.G0 = v17
	F___gettimeofday(m, v17)
	mBase = m.M
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
	m.G0 = v17 + v6
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v21 + v20*int64(1000000) - int64(946684800000000)
	F_pq_beginmessage(m, v7, int32(100))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return
	} else {
		F_enlargeStringInfo(m, v7, int32(1))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v40 = int32(112)
			*(*uint8)(unsafe.Add(mBase, uint32(v37+v38))) = uint8(v40)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v37 + int32(1)
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
			F_enlargeStringInfo(m, v7, int32(8))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v52 = int64(56)
				v54 = int64(65280)
				v56 = int64(40)
				v59 = int64(16711680)
				v61 = int64(24)
				v63 = int64(4278190080)
				v65 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v49+v50))) = v45<<(uint(v52)%64) | v45&v54<<(uint(v56)%64) | (v45&v59<<(uint(v61)%64) | v45&v63<<(uint(v65)%64)) | (int64(base.Ui64(v45)>>(uint(v65)%64))&v63 | int64(base.Ui64(v45)>>(uint(v61)%64))&v59 | (int64(base.Ui64(v45)>>(uint(v56)%64))&v54 | int64(base.Ui64(v45)>>(uint(v52)%64))))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v49 + int32(8)
				F_pq_endmessage(m, v7)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, _consts[314]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
					v96 = m.T0[v95].(func(*base.Module) int32)(m)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_bbsink_copystream_end_backup(m *base.Module, l0 int32, l1 int64, l2 int32) {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	F_pq_putemptymessage(m, int32(99))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_SendXlogRecPtrResult(m, l1, l2)
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bbsink_forward_begin_backup(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v4
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	m.T0[v10].(func(*base.Module, int32))(m, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14
		return
	}
}
func F_bbsink_forward_cleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	m.T0[v4].(func(*base.Module, int32))(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_forward_manifest_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	m.T0[v5].(func(*base.Module, int32, int32))(m, v3, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_progress_end_archive(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		v9 = v7
	} else {
		v9 = int32(0)
	}
	if v5 < v9 {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[55]))
		if v17 == int32(0) {
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
			if v21 != int32(1) {
			} else {
				v24 = int32(4483812)
				v26 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				v27 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[26])) = v26 + v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v30 + v27
				*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32))+232)) = base.I64_extend_i32_s(v5 + int32(1))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v38 + v27
				v44 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				*(*int32)(unsafe.Add(mBase, _consts[26])) = v44 - v27
			}
		}
	} else {
	}
	F_bbsink_forward_end_archive(m, l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		return
	} else {
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v51 + int32(1)
		return
	}
}
func F_bbsink_server_begin_archive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v16 = F_psprintf(m, int32(176683), v8+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = F_PathNameOpenFile(m, v16, int32(193))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v19
			if v19 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
						F_errmsg(m, int32(298425), v8)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_errfinish(m, int32(493397), int32(149), int32(342845))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_pfree(m, v16)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_bbsink_forward_begin_archive(m, l0, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	}
}
