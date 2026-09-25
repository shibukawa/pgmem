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
	var v41 int32
	_ = v41
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
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
		if v29 == v32 {
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[0]))
			if v41 == int32(0) {
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[1])))
				if v45&int32(1) == int32(0) {
				} else {
					v50 = int32(_a_F_bbsink_progress_archive_contents_0)
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[2]))
					v53 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[2])) = v52 + v53
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v56 + v53
					v60 = int32(0)
					v63 = base.AtomicRmwOr32(m, v60, int32(_a_F_bbsink_progress_archive_contents_1), v60)
					if v29 <= v60 {
					} else {
						v67 = v29 & int32(3)
						v69 = v41 + int32(232)
						if base.Ui32(int32(4)) <= base.Ui32(v29) {
							v75 = int32(0)
							v78 = v32
							for {
								v84 = int32(2)
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v31+v78<<(uint(v84)%32))))
								v88 = int32(3)
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v8+v78<<(uint(v88)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v69+v87<<(uint(v88)%32)))) = v94
								v97 = v78 | int32(1)
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v31+v97<<(uint(v84)%32))))
								v108 = *(*int64)(unsafe.Add(mBase, uint32(v8+v97<<(uint(v88)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v69+v101<<(uint(v88)%32)))) = v108
								v111 = v78 | v84
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v31+v111<<(uint(v84)%32))))
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v8+v111<<(uint(v88)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v69+v115<<(uint(v88)%32)))) = v122
								v125 = v78 | v88
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v31+v125<<(uint(v84)%32))))
								v136 = *(*int64)(unsafe.Add(mBase, uint32(v8+v125<<(uint(v88)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v69+v129<<(uint(v88)%32)))) = v136
								v138 = int32(4)
								v139 = v78 + v138
								v141 = v75 + v138
								if v141 != v29&int32(2147483644) {
									v75 = v141
									v78 = v139
									continue
								} else {
									break
								}
								break
							}
							if v67 == int32(0) {
							} else {
								v148 = v139
								v155 = int32(0)
								v158 = v148
								for {
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v31+v158<<(uint(int32(2))%32))))
									v168 = int32(3)
									v174 = *(*int64)(unsafe.Add(mBase, uint32(v8+v158<<(uint(v168)%32))))
									*(*int64)(unsafe.Add(mBase, uint32(v69+v167<<(uint(v168)%32)))) = v174
									v176 = int32(1)
									v179 = v155 + v176
									if v179 != v67 {
										v155 = v179
										v158 = v158 + v176
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v148 = v32
							v155 = int32(0)
							v158 = v148
							for {
								v167 = *(*int32)(unsafe.Add(mBase, uint32(v31+v158<<(uint(int32(2))%32))))
								v168 = int32(3)
								v174 = *(*int64)(unsafe.Add(mBase, uint32(v8+v158<<(uint(v168)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v69+v167<<(uint(v168)%32)))) = v174
								v176 = int32(1)
								v179 = v155 + v176
								if v179 != v67 {
									v155 = v179
									v158 = v158 + v176
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v190 = int32(0)
					v193 = base.AtomicRmwOr32(m, v190, int32(_a_F_bbsink_progress_archive_contents_1), v190)
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					v195 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v194 + v195
					v198 = int32(_a_F_bbsink_progress_archive_contents_0)
					v200 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_archive_contents[2])) = v200 - v195
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
		v19 = F_psprintf(m, int32(_a_F_bbsink_server_end_manifest_0), v7+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v21
			v24 = F_psprintf(m, int32(_a_F_bbsink_server_end_manifest_1), v7)
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
