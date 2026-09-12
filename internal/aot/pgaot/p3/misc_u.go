package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	if l0 != 0 {
		F_ResourceOwnerForget(m, l1, l0, int32(1752900))
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_UnregisterSnapshotNoOwner(m, l0)
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_UnreservedPLKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	v3 = int32(0)
	if l1 == v3 {
		v81 = int32(1)
		v87 = int32(0)
	} else {
		v13 = int32(1)
		if l1 == v13 {
			v58 = l0
			v59 = l1
			v60 = v13
			v61 = int32(8191)
			v66 = int32(0)
		} else {
			v23 = l0
			v24 = int32(0)
			v25 = v13
			v27 = v3
			for {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v32 = int32(32)
				v33 = v31 | v32
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v36 = v34 | v32
				v37 = int32(8191)
				v42 = v33 + (v36+v25*v37)*v37
				v43 = int32(257)
				v48 = (v24*v43+v36)*v43 + v33
				v49 = int32(2)
				v50 = v23 + v49
				v52 = v27 + v49
				if v52 != l1&int32(-2) {
					v23 = v50
					v24 = v48
					v25 = v42
					v27 = v52
					continue
				} else {
					break
				}
				break
			}
			v58 = v50
			v59 = v48
			v60 = v42
			v61 = v42 * int32(8191)
			v66 = v48 * int32(257)
		}
		if l1&v13 != 0 {
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
			v69 = v67 | int32(32)
			v73 = v69 + v66
			v74 = v61 + v69
		} else {
			v73 = v59
			v74 = v60
		}
		v75 = int32(167)
		v76 = base.I32_rem_u_s(v74, v75)
		v78 = base.I32_rem_u_s(v73, v75)
		v81 = v76
		v87 = v78
	}
	v89 = int32(1)
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81<<(uint(v89)%32))+uint32(_consts[1277]))))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v89)%32))+uint32(_consts[1277]))))
	return v92 + v96
}
func F_UpdateDecodingStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	v20 = m.G0
	v22 = v20 - int32(144)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	if int64(0) < v25 {
		v36 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			if v36 != 0 {
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v24)+208))
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v24)+216))
				*(*int64)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = v47
				*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v41
				*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v40
				*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
				F_errmsg_internal(m, int32(429617), v22)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_errfinish(m, int32(496980), int32(1972), int32(126035))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
						v75 = v24 + int32(168)
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
						v79 = v24 + int32(176)
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
						v83 = v24 + int32(184)
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
						v87 = v24 + int32(192)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
						v91 = v24 + int32(200)
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
						v95 = v24 + int32(208)
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
						v99 = v24 + int32(216)
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
						v103 = v22 + int32(72)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
						v111 = base.I32_div_s(v106-v108, int32(288))
						v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
							v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
							v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
							v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
							F_pgstat_unlock_entry(m, v114)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								v151 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
								m.G0 = v22 + int32(144)
								return
							}
						}
					}
				}
			} else {
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
				v75 = v24 + int32(168)
				v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
				v79 = v24 + int32(176)
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
				v83 = v24 + int32(184)
				v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
				v87 = v24 + int32(192)
				v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
				v91 = v24 + int32(200)
				v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
				v95 = v24 + int32(208)
				v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
				v99 = v24 + int32(216)
				v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
				v103 = v22 + int32(72)
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
				v111 = base.I32_div_s(v106-v108, int32(288))
				v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
					v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
					v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
					v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
					v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
					v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
					v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
					v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
					F_pgstat_unlock_entry(m, v114)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						v151 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
						*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
						m.G0 = v22 + int32(144)
						return
					}
				}
			}
		}
	} else {
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
		if int64(0) < v28 {
			v36 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				if v36 != 0 {
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v24)+208))
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v24)+216))
					*(*int64)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = v47
					*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v43
					*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v40
					*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
					F_errmsg_internal(m, int32(429617), v22)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						F_errfinish(m, int32(496980), int32(1972), int32(126035))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
							v75 = v24 + int32(168)
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
							v79 = v24 + int32(176)
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
							v83 = v24 + int32(184)
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
							v87 = v24 + int32(192)
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
							v91 = v24 + int32(200)
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
							v95 = v24 + int32(208)
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
							v99 = v24 + int32(216)
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
							*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
							v103 = v22 + int32(72)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
							v111 = base.I32_div_s(v106-v108, int32(288))
							v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
								v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
								v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
								v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
								v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
								v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
								v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
								v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
								v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
								F_pgstat_unlock_entry(m, v114)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return
								} else {
									v151 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
									*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
									m.G0 = v22 + int32(144)
									return
								}
							}
						}
					}
				} else {
					v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
					v75 = v24 + int32(168)
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
					v79 = v24 + int32(176)
					v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
					v83 = v24 + int32(184)
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
					v87 = v24 + int32(192)
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
					v91 = v24 + int32(200)
					v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
					v95 = v24 + int32(208)
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
					v99 = v24 + int32(216)
					v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
					v103 = v22 + int32(72)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
					v111 = base.I32_div_s(v106-v108, int32(288))
					v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
						v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
						v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
						F_pgstat_unlock_entry(m, v114)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							v151 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
							*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
							m.G0 = v22 + int32(144)
							return
						}
					}
				}
			}
		} else {
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v24)+216))
			if v31 <= int64(0) {
				m.G0 = v22 + int32(144)
				return
			} else {
				v36 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if v36 != 0 {
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)+168))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+184))
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v24)+208))
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v24)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v44
						*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v43
						*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v40
						*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v39
						*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
						F_errmsg_internal(m, int32(429617), v22)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errfinish(m, int32(496980), int32(1972), int32(126035))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
								v75 = v24 + int32(168)
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
								v79 = v24 + int32(176)
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
								v83 = v24 + int32(184)
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
								v87 = v24 + int32(192)
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
								v91 = v24 + int32(200)
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
								v95 = v24 + int32(208)
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
								v99 = v24 + int32(216)
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
								v103 = v22 + int32(72)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
								v111 = base.I32_div_s(v106-v108, int32(288))
								v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
									v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
									v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
									v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
									v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
									v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
									v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
									v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
									v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
									v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
									v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
									v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
									v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
									v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
									v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
									F_pgstat_unlock_entry(m, v114)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										v151 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
										*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
										m.G0 = v22 + int32(144)
										return
									}
								}
							}
						}
					} else {
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v72
						v75 = v24 + int32(168)
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v76
						v79 = v24 + int32(176)
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v80
						v83 = v24 + int32(184)
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v84
						v87 = v24 + int32(192)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v88
						v91 = v24 + int32(200)
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v92
						v95 = v24 + int32(208)
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v96
						v99 = v24 + int32(216)
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v100
						v103 = v22 + int32(72)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v108 = *(*int32)(unsafe.Add(mBase, _consts[268]))
						v111 = base.I32_div_s(v106-v108, int32(288))
						v114 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v111), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+24))
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+24)) = v117 + v118
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v116)+32))
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+32)) = v121 + v122
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v116)+40))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v103)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+40)) = v125 + v126
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v103)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+48)) = v129 + v130
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v116)+56))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v103)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+56)) = v133 + v134
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v116)+64))
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v103)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+64)) = v137 + v138
							v141 = *(*int64)(unsafe.Add(mBase, uint32(v116)+72))
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v103)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+72)) = v141 + v142
							v145 = *(*int64)(unsafe.Add(mBase, uint32(v116)+80))
							v146 = *(*int64)(unsafe.Add(mBase, uint32(v103)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v116)+80)) = v145 + v146
							F_pgstat_unlock_entry(m, v114)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								v151 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v99))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v95))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v91))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v87))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v83))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v79))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v75))) = v151
								*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v151
								m.G0 = v22 + int32(144)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_UtilityContainsQuery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = l0
	goto L1
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	switch v5 - int32(241) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L7
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v3 = v27
	goto L1
L4:
	;
	return v24
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 == int32(6) {
		v26 = v20
		goto L3
	} else {
		goto L13
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 == int32(6) {
		v26 = v16
		goto L3
	} else {
		goto L12
	}
L7:
	;
	if v5 != int32(201) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v13 != int32(6) {
		v24 = v12
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = v12
	goto L3
L12:
	;
	v24 = v16
	goto L4
L13:
	;
	v24 = v20
	goto L4
}
func F_uint32in_subr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v5
	v20 = F_strtox_2(m, l0, v11+int32(44), v5, int64(4294967295))
	mBase = m.M
	v21 = base.I32_wrap_i64(v20)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v23 != 0 {
		v27 = base.B2i32(v23 != int32(68))
	} else {
		v27 = int32(0)
	}
	if v27 == int32(0) {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
		if v30 != l0 {
			if v23 == int32(68) {
				v55 = int32(0)
				v56 = F_errsave_start(m, l3)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v56 == int32(0) {
						v123 = v55
						m.G0 = v11 + int32(48)
						return v123
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
							F_errmsg(m, int32(189701), v11+int32(16))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l3, int32(493301), int32(924), int32(229075))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v123 = v55
									m.G0 = v11 + int32(48)
									return v123
								}
							}
						}
					}
				}
			} else {
				if l1 == int32(0) {
					v84 = v30
					for {
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
						if base.Ui32(v85-int32(9)) < base.Ui32(int32(5)) {
							v84 = v84 + int32(1)
							continue
						} else {
						}
						if v85 == int32(32) {
							v84 = v84 + int32(1)
							continue
						} else {
							break
						}
						break
					}
					if v85 == int32(0) {
						v123 = v21
						m.G0 = v11 + int32(48)
						return v123
					} else {
						v94 = int32(0)
						v95 = F_errsave_start(m, l3)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							if v95 == int32(0) {
								v123 = v94
								m.G0 = v11 + int32(48)
								return v123
							} else {
								F_errcode(m, int32(33685634))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l2
									F_errmsg(m, int32(709275), v11+int32(32))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l3, int32(493301), int32(940), int32(229075))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v123 = v94
											m.G0 = v11 + int32(48)
											return v123
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
					v123 = v21
					m.G0 = v11 + int32(48)
					return v123
				}
			}
		} else {
			v33 = int32(0)
			v34 = F_errsave_start(m, l3)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					v123 = v33
					m.G0 = v11 + int32(48)
					return v123
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
						F_errmsg(m, int32(709275), v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l3, int32(493301), int32(918), int32(229075))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v123 = v33
								m.G0 = v11 + int32(48)
								return v123
							}
						}
					}
				}
			}
		}
	} else {
		v33 = int32(0)
		v34 = F_errsave_start(m, l3)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if v34 == int32(0) {
				v123 = v33
				m.G0 = v11 + int32(48)
				return v123
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
					F_errmsg(m, int32(709275), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, l3, int32(493301), int32(918), int32(229075))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v123 = v33
							m.G0 = v11 + int32(48)
							return v123
						}
					}
				}
			}
		}
	}
}
func F_umask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	v3 = int32(4383200)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1385]))
	*(*int32)(unsafe.Add(mBase, _consts[1385])) = l0
	if base.Ui32(int32(-4095)) <= base.Ui32(v4) {
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v4
		v14 = int32(-1)
	} else {
		v14 = v4
	}
	return v14
}
func F_unaccent_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v2
	if v17 <= v2 {
		v128 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v128
L2:
	;
	v27 = v18
	v31 = v17
	goto L3
L3:
	;
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v117 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	v114 = v31 - v106
	if int32(0) < v114 {
		v27 = v27 + v106
		v31 = v114
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v95 = F_pg_mblen_range(m, v27, v17+v18)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L23
	} else {
		goto L28
	}
L7:
	;
	v38 = int32(0)
	v41 = v38
	v44 = v16
	v45 = v38
	v46 = v38
	goto L8
L8:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v27))))
	v56 = v44 + v53*int32(12)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v58 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L10:
	;
	goto L9
L11:
	;
	v58 = v56
	goto L13
L12:
	;
	v58 = v46
	goto L13
L13:
	;
	v60 = v41 + int32(1)
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v60
	goto L16
L15:
	;
	v61 = v45
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v62 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v60) < base.Ui32(v31) {
		v41 = v60
		v44 = v62
		v45 = v61
		v46 = v58
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v68 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	F_appendBinaryStringInfo(m, v14, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L27
	}
L22:
	;
	F_initStringInfo(m, v14)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v27 == v18 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F_appendBinaryStringInfo(m, v14, v18, v27-v18)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v106 = v61
	goto L5
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v97 == int32(0) {
		v106 = v95
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_appendBinaryStringInfo(m, v14, v27, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v106 = v95
	goto L5
L31:
	;
	goto L4
L32:
	;
	v128 = int32(0)
	goto L1
L33:
	;
	goto L34
L34:
	;
	v122 = F_palloc0(m, int32(16))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L35
	}
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v125 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v125)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v124
	v128 = v122
	goto L1
}
func F_union_tuples(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v21 = F_AllocSetContextCreateInternal(m, v16, int32(272623), v4, int32(8192), int32(8388608))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = int32(4489440)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
	v28 = F_brin_deform_tuple(m, l0, l2, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(0) < v34 {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	F_MemoryContextDelete(m, v21)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v34 <= int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(24)
	v48 = v4
	goto L9
L9:
	;
	v59 = v48 * int32(20)
	v60 = v28 + v40 + v59
	v61 = v59 + (l1 + v40)
	v63 = v48 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v63)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	if v66 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L5
L11:
	;
	v155 = v48 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v155 < v157 {
		v48 = v155
		goto L9
	} else {
		goto L29
	}
L12:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = F_index_getprocinfo(m, v127, base.I32_extend16_s(v48+int32(1)), int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L27
	}
L13:
	;
	v69 = int32(0)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v70 == v69 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+3)))
	v76 = v73 ^ int32(1)
	goto L16
L15:
	;
	v76 = v69
	goto L16
L16:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	if v77 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+3)))
	if v82 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	if v76&int32(1) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)) = uint8(v80)
	goto L17
L20:
	;
	if v77 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v85 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)) = uint16(v85)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	if v87 == int32(0) {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v95 = int32(0)
	goto L23
L23:
	;
	v108 = v95 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v65+int32(8)))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+10)))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+8)))
	v116 = F_datumCopy(m, v111, v114, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L11
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v118+v108))) = v116
	v122 = v95 + int32(1)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	if base.Ui32(v122) < base.Ui32(v123) {
		v95 = v122
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+248))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135+v63)))
	v138 = F_FunctionCall3Coll(m, v132, v137, l0, v61, v60)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L11
L29:
	;
	goto L10
L30:
	;
	return
L31:
	;
	v179 = int32(24)
	v191 = v4
	goto L34
L32:
	;
	goto L33
L33:
	;
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v281)
	F_MemoryContextDelete(m, v21)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v191<<(uint(int32(2))%32))))
	v202 = v191 * int32(20)
	v203 = l1 + v179 + v202
	v204 = v202 + (v28 + v179)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+3)) = uint8(v205)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+2)) = uint8(v207)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+3)))
	if v209 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v263 = v191 + int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v263 < v265 {
		v191 = v263
		goto L34
	} else {
		goto L43
	}
L37:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200))))
	if v210 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v218 = int32(0)
	goto L39
L39:
	;
	v231 = v218 << (uint(int32(2)) % 32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231+v232)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v200+int32(8)))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+10)))
	v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v236)+8)))
	v239 = F_datumCopy(m, v234, v237, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v231))) = v239
	v245 = v218 + int32(1)
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200))))
	if base.Ui32(v245) < base.Ui32(v246) {
		v218 = v245
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L35
L44:
	;
	return
}
func F_unique_key_recheck(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L54
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 != int32(442) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v18&int32(28) != int32(4) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	switch v18 & int32(3) {
	case 0:
		v48 = int32(24)
		goto L7
	default:
		goto L9
	case 2:
		goto L8
	}
L7:
	;
	v50 = v10 + int32(220)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12+v48)))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v53)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+216)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v59 = F_table_slot_create(m, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L15
	}
L8:
	;
	v48 = int32(28)
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(317355)
	F_errmsg(m, int32(537999), v10+int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(492474), int32(83), int32(317355))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+212)) = uint16(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+188))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+44))
	v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+80)) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[461]))
	if v73 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[462])))
	if v75&int32(1) == int32(0) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+188))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	v89 = m.T0[v88].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v68, v10+int32(208), int32(4156424), v59, v10+int32(80), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	m.G0 = v10 + int32(224)
	return int32(0)
L22:
	;
	if v89 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ExecDropSingleTupleTableSlot(m, v59)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+188))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+52))
	m.T0[v102].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+188))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	m.T0[v97].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v108 = F_index_open(m, v106, int32(3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	F_FormIndexDatum(m, v110, v59, v122, v10+int32(80), v10+int32(48))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L40
	}
L30:
	;
	v115 = F_CreateExecutorState(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L35
	}
L31:
	;
	v110 = F_BuildIndexInfo(m, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+76))
	if v112 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+92))
	if v113 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v122 = int32(0)
	goto L29
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+152))
	if v117 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v120 = v117
	goto L38
L37:
	;
	v118 = F_MakePerTupleExprContext(m, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v59
	v122 = v115
	goto L29
L39:
	;
	v120 = v118
	goto L38
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v110)+92))
	if v131 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v122 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v142 = F_index_insert(m, v108, v10+int32(80), v10+int32(48), v10+int32(216), v130, int32(3), int32(0), v110)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_check_exclusion_constraint(m, v130, v108, v110, v10+int32(208), v10+int32(80), v10+int32(48), v122, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L47
	}
L45:
	;
	F_index_insert_cleanup(m, v108, v110)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	goto L41
L48:
	;
	F_FreeExecutorState(m, v122)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L10
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_ExecDropSingleTupleTableSlot(m, v59)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	F_relation_close(m, v108, int32(3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L21
L54:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(317355)
	F_errmsg(m, int32(224975), v10)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(492474), int32(62), int32(317355))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(317355)
	F_errmsg(m, int32(515409), v10+int32(32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(492474), int32(69), int32(317355))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(336012), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(326201), int32(1218), int32(383171))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unistr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v425 int32
	_ = v425
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v581 int32
	_ = v581
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v691 int32
	_ = v691
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v893 int32
	_ = v893
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v942 int32
	_ = v942
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum_packed(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = int32(1)
	v31 = v26 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v34 = v32 & v30
	if v32 == v30 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v1223 = v1221 + int32(4)
	v1224 = F_palloc(m, v1223)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L246
	}
L4:
	;
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v65 = int32(base.Ui32(v59)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L7:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if v34 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v65 = base.B2i32(v37 == int32(18)) << (uint(int32(4)) % 32)
	goto L5
L12:
	;
	v50 = int32(1)
	v65 = int32(base.Ui32(v32)>>(uint(v50)%32)) - v50
	goto L5
L13:
	;
	v73 = int32(4)
	goto L4
L14:
	;
	if v65 <= int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v73 = v65
	goto L4
L16:
	;
	v76 = v31
	goto L18
L17:
	;
	v76 = v26 + int32(4)
	goto L18
L18:
	;
	v79 = v76
	v82 = v73
	v84 = int32(0)
	goto L26
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L242
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L239
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L235
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L232
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L228
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L225
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L221
	}
L26:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79))))
	if v97 == int32(92) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	if v1023 == int32(0) {
		goto L3
	} else {
		goto L220
	}
L28:
	;
	if int32(0) < v1021 {
		v79 = v1018
		v82 = v1021
		v84 = v1023
		goto L26
	} else {
		goto L219
	}
L29:
	;
	v1018 = v1014
	v1021 = v1012
	v1023 = int32(0)
	goto L28
L30:
	;
	if v82 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v84 != 0 {
		goto L25
	} else {
		goto L217
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L212
	}
L34:
	;
	v103 = v79 + int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 == int32(92) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v84 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(v82) < base.Ui32(int32(5)) {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	F_appendStringInfoChar(m, v23+int32(80), int32(92))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v112 = int32(2)
	v1012 = v82 - v112
	v1014 = v79 + v112
	goto L29
L40:
	;
	v119 = int32(0)
	goto L41
L41:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v103))))
	v151 = base.B2i32(base.Ui32(v140-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v140|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L43
L42:
	;
	if v151 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	if v151 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v153 = v119 + int32(1)
	if v153 != int32(4) {
		v119 = v153
		goto L41
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L42
L47:
	;
	goto L46
L48:
	;
	if base.Ui32(v82) < base.Ui32(int32(8)) {
		goto L33
	} else {
		goto L97
	}
L49:
	;
	if v82 == int32(5) {
		goto L33
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v226 = int32(-48)
	if v104 == int32(117) {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	if v104 != int32(117) {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v166 = int32(0)
	goto L54
L54:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+(v79+int32(2))))))
	v198 = base.B2i32(base.Ui32(v187-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v187|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L56
L55:
	;
	if v198 == int32(0) {
		goto L48
	} else {
		goto L61
	}
L56:
	;
	if v198 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v200 = v166 + int32(1)
	if v200 != int32(4) {
		v166 = v200
		goto L54
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L55
L60:
	;
	goto L59
L61:
	;
	goto L51
L62:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if base.Ui32((v256-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v277 = v226
		goto L69
	} else {
		goto L70
	}
L63:
	;
	v232 = int32(2)
	goto L65
L64:
	;
	v232 = int32(1)
	goto L65
L65:
	;
	v233 = v79 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if base.Ui32((v234-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v255 = v226
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if base.Ui32((v234-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v255 = int32(-87)
		goto L62
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v234-int32(65))&int32(255)) {
		goto L24
	} else {
		goto L68
	}
L68:
	;
	v255 = int32(-55)
	goto L62
L69:
	;
	v278 = int32(-48)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+2)))
	if base.Ui32((v280-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v301 = v278
		goto L75
	} else {
		goto L76
	}
L70:
	;
	if base.Ui32((v256-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v277 = int32(-87)
	goto L69
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(5)) < base.Ui32((v256-int32(65))&int32(255)) {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	v277 = int32(-55)
	goto L69
L75:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+3)))
	if base.Ui32((v302-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v323 = v278
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if base.Ui32((v280-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v301 = int32(-87)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(int32(5)) < base.Ui32((v280-int32(65))&int32(255)) {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	v301 = int32(-55)
	goto L75
L79:
	;
	v336 = v302 + v323 + ((v256+v277)<<(uint(int32(8))%32) + (v234+v255)<<(uint(int32(12))%32) + (v280+v301)<<(uint(int32(4))%32))
	if base.Ui32(int32(1114111)) <= base.Ui32(v336-int32(1)) {
		goto L23
	} else {
		goto L85
	}
L80:
	;
	if base.Ui32((v302-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v323 = int32(-87)
	goto L79
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(int32(5)) < base.Ui32((v302-int32(65))&int32(255)) {
		goto L24
	} else {
		goto L84
	}
L84:
	;
	v323 = int32(-55)
	goto L79
L85:
	;
	v342 = v336 & int32(2096128)
	if v84 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v356&int32(-1024) == int32(55296) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	if v342 != int32(56320) {
		goto L25
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v342 == int32(56320) {
		goto L25
	} else {
		goto L91
	}
L90:
	;
	v356 = v84<<(uint(int32(10))%32)&int32(1047552) | v336&int32(1023) + int32(65536)
	goto L86
L91:
	;
	v356 = v336
	goto L86
L92:
	;
	v372 = v356
	goto L94
L93:
	;
	F_pg_unicode_to_server(m, v356, v23+int32(48))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	v374 = v232 | int32(4)
	v1018 = v374 + v79
	v1021 = v82 - v374
	v1023 = v372
	goto L28
L95:
	;
	F_appendStringInfoString(m, v23+int32(80), v23+int32(48))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v372 = int32(0)
	goto L94
L97:
	;
	if v104 != int32(43) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if base.Ui32(v82) < base.Ui32(int32(10)) {
		goto L33
	} else {
		goto L150
	}
L99:
	;
	v402 = v79 + int32(2)
	v404 = int32(0)
	goto L100
L100:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v402))))
	v436 = base.B2i32(base.Ui32(v425-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v425|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L102
L101:
	;
	if v436 == int32(0) {
		goto L98
	} else {
		goto L107
	}
L102:
	;
	if v436 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v438 = v404 + int32(1)
	if v438 != int32(6) {
		v404 = v438
		goto L100
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	goto L101
L106:
	;
	goto L105
L107:
	;
	v444 = int32(-48)
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	if base.Ui32((v446-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v467 = v444
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)))
	if base.Ui32((v468-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v489 = v444
		goto L112
	} else {
		goto L113
	}
L109:
	;
	if base.Ui32((v446-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v467 = int32(-87)
		goto L108
	} else {
		goto L110
	}
L110:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v446-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v467 = int32(-55)
	goto L108
L112:
	;
	v490 = int32(-48)
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)))
	if base.Ui32((v492-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v513 = v490
		goto L118
	} else {
		goto L119
	}
L113:
	;
	if base.Ui32((v468-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v489 = int32(-87)
	goto L112
L115:
	;
	goto L116
L116:
	;
	if base.Ui32(int32(5)) < base.Ui32((v468-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L117
	}
L117:
	;
	v489 = int32(-55)
	goto L112
L118:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)))
	if base.Ui32((v514-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v535 = v490
		goto L122
	} else {
		goto L123
	}
L119:
	;
	if base.Ui32((v492-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v513 = int32(-87)
		goto L118
	} else {
		goto L120
	}
L120:
	;
	if base.Ui32(int32(5)) < base.Ui32((v492-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L121
	}
L121:
	;
	v513 = int32(-55)
	goto L118
L122:
	;
	v536 = int32(-48)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	if base.Ui32((v538-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v559 = v536
		goto L128
	} else {
		goto L129
	}
L123:
	;
	if base.Ui32((v514-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v535 = int32(-87)
	goto L122
L125:
	;
	goto L126
L126:
	;
	if base.Ui32(int32(5)) < base.Ui32((v514-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	v535 = int32(-55)
	goto L122
L128:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	if base.Ui32((v560-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v581 = v536
		goto L132
	} else {
		goto L133
	}
L129:
	;
	if base.Ui32((v538-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v559 = int32(-87)
		goto L128
	} else {
		goto L130
	}
L130:
	;
	if base.Ui32(int32(5)) < base.Ui32((v538-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L131
	}
L131:
	;
	v559 = int32(-55)
	goto L128
L132:
	;
	v602 = v560 + v581 + ((v468+v489)<<(uint(int32(16))%32) + (v446+v467)<<(uint(int32(20))%32) + (v492+v513)<<(uint(int32(12))%32) + (v514+v535)<<(uint(int32(8))%32) + (v538+v559)<<(uint(int32(4))%32))
	if base.Ui32(int32(1114111)) <= base.Ui32(v602-int32(1)) {
		goto L21
	} else {
		goto L138
	}
L133:
	;
	if base.Ui32((v560-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v581 = int32(-87)
	goto L132
L135:
	;
	goto L136
L136:
	;
	if base.Ui32(int32(5)) < base.Ui32((v560-int32(65))&int32(255)) {
		goto L22
	} else {
		goto L137
	}
L137:
	;
	v581 = int32(-55)
	goto L132
L138:
	;
	v608 = v602 & int32(2096128)
	if v84 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v622&int32(-1024) == int32(55296) {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	if v608 != int32(56320) {
		goto L25
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v608 == int32(56320) {
		goto L25
	} else {
		goto L144
	}
L143:
	;
	v622 = v84<<(uint(int32(10))%32)&int32(1047552) | v602&int32(1023) + int32(65536)
	goto L139
L144:
	;
	v622 = v602
	goto L139
L145:
	;
	v638 = v622
	goto L147
L146:
	;
	F_pg_unicode_to_server(m, v622, v23+int32(48))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L148
	}
L147:
	;
	v639 = int32(8)
	v1018 = v79 + v639
	v1021 = v82 - v639
	v1023 = v638
	goto L28
L148:
	;
	F_appendStringInfoString(m, v23+int32(80), v23+int32(48))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v638 = int32(0)
	goto L147
L150:
	;
	if v104 != int32(85) {
		goto L33
	} else {
		goto L151
	}
L151:
	;
	v668 = v79 + int32(2)
	v670 = int32(0)
	goto L152
L152:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670+v668))))
	v702 = base.B2i32(base.Ui32(v691-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v691|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L154
L153:
	;
	if v702 == int32(0) {
		goto L33
	} else {
		goto L159
	}
L154:
	;
	if v702 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v704 = v670 + int32(1)
	if v704 != int32(8) {
		v670 = v704
		goto L152
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	goto L153
L158:
	;
	goto L157
L159:
	;
	v710 = int32(-48)
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	if base.Ui32((v712-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v733 = v710
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)))
	if base.Ui32((v734-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v755 = v710
		goto L164
	} else {
		goto L165
	}
L161:
	;
	if base.Ui32((v712-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v733 = int32(-87)
		goto L160
	} else {
		goto L162
	}
L162:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v712-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L163
	}
L163:
	;
	v733 = int32(-55)
	goto L160
L164:
	;
	v756 = int32(-48)
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)))
	if base.Ui32((v758-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v779 = v756
		goto L170
	} else {
		goto L171
	}
L165:
	;
	if base.Ui32((v734-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v755 = int32(-87)
	goto L164
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(int32(5)) < base.Ui32((v734-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L169
	}
L169:
	;
	v755 = int32(-55)
	goto L164
L170:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)))
	if base.Ui32((v780-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v801 = v756
		goto L174
	} else {
		goto L175
	}
L171:
	;
	if base.Ui32((v758-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v779 = int32(-87)
		goto L170
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(int32(5)) < base.Ui32((v758-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L173
	}
L173:
	;
	v779 = int32(-55)
	goto L170
L174:
	;
	v802 = int32(-48)
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	if base.Ui32((v804-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v825 = v802
		goto L180
	} else {
		goto L181
	}
L175:
	;
	if base.Ui32((v780-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v801 = int32(-87)
	goto L174
L177:
	;
	goto L178
L178:
	;
	if base.Ui32(int32(5)) < base.Ui32((v780-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L179
	}
L179:
	;
	v801 = int32(-55)
	goto L174
L180:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	if base.Ui32((v826-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v847 = v802
		goto L184
	} else {
		goto L185
	}
L181:
	;
	if base.Ui32((v804-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v825 = int32(-87)
		goto L180
	} else {
		goto L182
	}
L182:
	;
	if base.Ui32(int32(5)) < base.Ui32((v804-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L183
	}
L183:
	;
	v825 = int32(-55)
	goto L180
L184:
	;
	v848 = int32(-48)
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	if base.Ui32((v850-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v871 = v848
		goto L190
	} else {
		goto L191
	}
L185:
	;
	if base.Ui32((v826-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v847 = int32(-87)
	goto L184
L187:
	;
	goto L188
L188:
	;
	if base.Ui32(int32(5)) < base.Ui32((v826-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L189
	}
L189:
	;
	v847 = int32(-55)
	goto L184
L190:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+9)))
	if base.Ui32((v872-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v893 = v848
		goto L194
	} else {
		goto L195
	}
L191:
	;
	if base.Ui32((v850-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v871 = int32(-87)
		goto L190
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(int32(5)) < base.Ui32((v850-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L193
	}
L193:
	;
	v871 = int32(-55)
	goto L190
L194:
	;
	v922 = v872 + v893 + ((v734+v755)<<(uint(int32(24))%32) + (v712+v733)<<(uint(int32(28))%32) + (v758+v779)<<(uint(int32(20))%32) + (v780+v801)<<(uint(int32(16))%32) + (v804+v825)<<(uint(int32(12))%32) + (v826+v847)<<(uint(int32(8))%32) + (v850+v871)<<(uint(int32(4))%32))
	if base.Ui32(int32(1114111)) <= base.Ui32(v922-int32(1)) {
		goto L19
	} else {
		goto L200
	}
L195:
	;
	if base.Ui32((v872-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v893 = int32(-87)
	goto L194
L197:
	;
	goto L198
L198:
	;
	if base.Ui32(int32(5)) < base.Ui32((v872-int32(65))&int32(255)) {
		goto L20
	} else {
		goto L199
	}
L199:
	;
	v893 = int32(-55)
	goto L194
L200:
	;
	v928 = v922 & int32(2096128)
	if v84 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if v942&int32(-1024) == int32(55296) {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	if v928 != int32(56320) {
		goto L25
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v928 == int32(56320) {
		goto L25
	} else {
		goto L206
	}
L205:
	;
	v942 = v84<<(uint(int32(10))%32)&int32(1047552) | v922&int32(1023) + int32(65536)
	goto L201
L206:
	;
	v942 = v922
	goto L201
L207:
	;
	v958 = v942
	goto L209
L208:
	;
	F_pg_unicode_to_server(m, v942, v23+int32(48))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v959 = int32(10)
	v1018 = v79 + v959
	v1021 = v82 - v959
	v1023 = v958
	goto L28
L210:
	;
	F_appendStringInfoString(m, v23+int32(80), v23+int32(48))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v958 = int32(0)
	goto L209
L212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(371503), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errhint(m, int32(640355), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(499506), int32(6902), int32(206184))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_appendStringInfoChar(m, v23+int32(80), v97)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v1007 = int32(1)
	v1012 = v82 - v1007
	v1014 = v79 + v1007
	goto L29
L219:
	;
	goto L27
L220:
	;
	goto L25
L221:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errmsg(m, int32(213149), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(499506), int32(6926), int32(206184))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_errmsg_internal(m, int32(102717), int32(0))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(499506), int32(6741), int32(308274))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v336
	F_errmsg(m, int32(509646), v23)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(499506), int32(6802), int32(206184))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errmsg_internal(m, int32(102717), int32(0))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(499506), int32(6741), int32(308274))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v602
	F_errmsg(m, int32(509646), v23+int32(16))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(499506), int32(6837), int32(206184))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errmsg_internal(m, int32(102717), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(499506), int32(6741), int32(308274))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v922
	F_errmsg(m, int32(509646), v23+int32(32))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(499506), int32(6872), int32(206184))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1224))) = v1223 << (uint(int32(2)) % 32)
	if v1221 != 0 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	F_pfree(m, v1233)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L1
	} else {
		goto L251
	}
L248:
	;
	v1231 = F__emscripten_memcpy_bulkmem(m, v1224+int32(4), v1220, v1221)
	mBase = m.M
	goto L250
L249:
	;
	goto L250
L250:
	;
	goto L247
L251:
	;
	m.G0 = v23 + int32(96)
	return v1224
}
func F_updateClosestMatch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 == int32(0) {
		return
	} else {
		if l1 == int32(0) {
			return
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v11 == int32(0) {
				return
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v14 == int32(0) {
					return
				} else {
					v17 = F_strlen(m, v6)
					mBase = m.M
					if base.Ui32(int32(255)) < base.Ui32(v17) {
						return
					} else {
						v20 = F_strlen(m, l1)
						mBase = m.M
						if base.Ui32(int32(255)) < base.Ui32(v20) {
							return
						} else {
							v23 = int32(1)
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v28 = F_varstr_levenshtein_less_equal(m, v6, v17, l1, v20, v23, v23, v23, v26, v23)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v30 < v28 {
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v33 = F_strlen(m, v32)
									mBase = m.M
									if base.Ui32(int32(base.Ui32(v33)>>(uint(int32(1))%32))) < base.Ui32(v28) {
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if base.B2i32(v37 != int32(-1))&base.B2i32(v37 <= v28) != 0 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
										}
									}
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_update_controlfile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v5 = m.G0
	v7 = v5 - int32(9296)
	m.G0 = v7
	v9 = F___time(m)
	mBase = m.M
	v10 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v9
	v15 = m.Env.Pgmem_crc32c(m, v10, l1, int32(292))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+292)) = v15 ^ v10
	v24 = F__emscripten_memset_bulkmem(m, v7+int32(1400), base.I32_extend8_s(int32(0)), int32(7896))
	mBase = m.M
	goto L1
L1:
	;
	goto L3
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = int32(300970)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = l0
	v39 = F_pg_snprintf(m, v7+int32(80), int32(1024), int32(176834), v7-int32(-64))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v28 = F__emscripten_memcpy_bulkmem(m, v7+int32(1104), l1, int32(296))
	mBase = m.M
	goto L5
L5:
	;
	goto L2
L6:
	;
	return
L7:
	;
	v44 = F_BasicOpenFile(m, v7+int32(80), int32(2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L41
	}
L9:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L37
	}
L10:
	;
	if int32(0) <= v44 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(167772173)
	v57 = int32(8192)
	v58 = F_write(m, v44, v7+int32(1104), v57)
	mBase = m.M
	if v58 != v57 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L33
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v62 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v87 = int32(4104780)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v89
	v92 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(167772171)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v97 != int32(1) {
		v111 = v89
		goto L25
	} else {
		goto L26
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(51)
	goto L19
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v7 + int32(80)
	F_errmsg(m, int32(298832), v7+int32(48))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(493404), int32(247), int32(385721))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	if v111 != 0 {
		goto L9
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	goto L27
L27:
	;
	v102 = F_fsync(m, v44)
	mBase = m.M
	if v102 != int32(-1) {
		v111 = v102
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v111 = int32(-1)
	goto L25
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v106 == int32(27) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(0)
	v116 = F_close(m, v44)
	mBase = m.M
	if v116 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v7 + int32(9296)
	return
L33:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(80)
	F_errmsg(m, int32(298247), v7)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(493404), int32(226), int32(385721))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v7 + int32(80)
	F_errmsg(m, int32(299195), v7+int32(32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(493404), int32(264), int32(385721))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(80)
	F_errmsg(m, int32(298990), v7+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(493404), int32(278), int32(385721))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_update_spins_per_delay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	v8 = base.I32_div_s(v3+l0*int32(15), int32(16))
	return v8
}
