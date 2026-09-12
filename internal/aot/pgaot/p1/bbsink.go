package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bbsink_progress_archive_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(4294967298)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13 + base.I64_extend_i32_u(l1)
	F_bbsink_forward_archive_contents(m, l0, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v19
		v21 = int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
		if v22 != v21 {
			v29 = v21
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
			if base.Ui64(v19) <= base.Ui64(v25) {
				v29 = v21
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v19
				v29 = int32(2)
			}
		}
		v31 = v8 + int32(24)
		v32 = int32(0)
		v39 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if v39 == v32 {
		} else {
			if v29 == int32(0) {
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
				if v45&int32(1) == int32(0) {
				} else {
					v50 = int32(4510372)
					v52 = *(*int32)(unsafe.Add(mBase, _consts[11]))
					v53 = int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[11])) = v52 + v53
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v56 + v53
					if v29 <= int32(0) {
					} else {
						v63 = v29 & int32(3)
						v65 = v39 + int32(232)
						if base.Ui32(int32(4)) <= base.Ui32(v29) {
							v71 = int32(0)
							v74 = v32
							for {
								v80 = int32(2)
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v31+v74<<(uint(v80)%32))))
								v84 = int32(3)
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v8+v74<<(uint(v84)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v65+v83<<(uint(v84)%32)))) = v90
								v93 = v74 | int32(1)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v31+v93<<(uint(v80)%32))))
								v104 = *(*int64)(unsafe.Add(mBase, uint32(v8+v93<<(uint(v84)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v65+v97<<(uint(v84)%32)))) = v104
								v107 = v74 | v80
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v31+v107<<(uint(v80)%32))))
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v8+v107<<(uint(v84)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v65+v111<<(uint(v84)%32)))) = v118
								v121 = v74 | v84
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v31+v121<<(uint(v80)%32))))
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v8+v121<<(uint(v84)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v65+v125<<(uint(v84)%32)))) = v132
								v134 = int32(4)
								v135 = v74 + v134
								v137 = v71 + v134
								if v137 != v29&int32(2147483644) {
									v71 = v137
									v74 = v135
									continue
								} else {
									break
								}
								break
							}
							v142 = v135
						} else {
							v142 = v32
						}
						if v63 == int32(0) {
						} else {
							v151 = int32(0)
							v154 = v142
							for {
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v31+v154<<(uint(int32(2))%32))))
								v164 = int32(3)
								v170 = *(*int64)(unsafe.Add(mBase, uint32(v8+v154<<(uint(v164)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v65+v163<<(uint(v164)%32)))) = v170
								v172 = int32(1)
								v175 = v151 + v172
								if v175 != v63 {
									v151 = v175
									v154 = v154 + v172
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v186 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v187 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v186 + v187
					v190 = int32(4510372)
					v192 = *(*int32)(unsafe.Add(mBase, _consts[11]))
					*(*int32)(unsafe.Add(mBase, _consts[11])) = v192 - v187
				}
			}
		}
		m.G0 = v8 + int32(32)
		return
	}
}
func F_bbsink_server_end_manifest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_FileClose(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v14
		v19 = F_psprintf(m, int32(235512), v7+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v21
			v24 = F_psprintf(m, int32(77518), v7)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = F_durable_rename(m, v19, v24, int32(21))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_pfree(m, v24)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_pfree(m, v19)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_bbsink_forward_end_manifest(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
