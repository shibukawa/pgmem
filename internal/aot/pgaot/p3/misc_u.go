package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnregisterSnapshotFromOwner(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 != 0 {
		F_ResourceOwnerForget(m, l1, base.I64_extend_i32_u(l0), int32(_a_F_UnregisterSnapshotFromOwner_0))
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_UnregisterSnapshotNoOwner(m, l0)
			v8 = m.ExcPending
			if v8 != 0 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v3 = int32(0)
	if l1 != 0 {
		if l1 == int32(1) {
			v50 = l0
			v51 = int32(0)
			v52 = v3
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
			v60 = v58 | int32(32)
			v68 = v51*int32(257) + v60
			v69 = v60 + v52*int32(31)
		} else {
			v17 = l0
			v18 = int32(0)
			v19 = v3
			v24 = v3
			for {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				v26 = int32(32)
				v27 = v25 | v26
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v30 = v28 | v26
				v31 = int32(31)
				v36 = v27 + (v30+v19*v31)*v31
				v37 = int32(257)
				v42 = (v18*v37+v30)*v37 + v27
				v43 = int32(2)
				v44 = v17 + v43
				v46 = v24 + v43
				if v46 != l1&int32(-2) {
					v17 = v44
					v18 = v42
					v19 = v36
					v24 = v46
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v68 = v42
				v69 = v36
			} else {
				v50 = v44
				v51 = v42
				v52 = v36
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
				v60 = v58 | int32(32)
				v68 = v51*int32(257) + v60
				v69 = v60 + v52*int32(31)
			}
		}
		v75 = int32(171)
		v76 = base.I32_rem_u_s(v69, v75)
		v78 = base.I32_rem_u_s(v68, v75)
		v82 = v76
		v88 = v78
	} else {
		v82 = v3
		v88 = int32(0)
	}
	v89 = int32(1)
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82<<(uint(v89)%32))+uint32(_c_F_UnreservedPLKeywords_hash_func[0]))))
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88<<(uint(v89)%32))+uint32(_c_F_UnreservedPLKeywords_hash_func[0]))))
	return v93 + v98
}
func F_UpdateDecodingStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
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
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
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
	var v142 int32
	_ = v142
	v14 = m.G0
	v16 = v14 - int32(176)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	if int64(0) < v19 {
		v33 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			if v33 != 0 {
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v16-int32(-64)))) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v41
				*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v40
				*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v36
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
				F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v16)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(2060), int32(_a_F_UpdateDecodingStats_2))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
						v91 = v16 + int32(80)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
						v99 = base.I32_div_s(v94-v96, int32(296))
						v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
							F_pgstat_unlock_entry(m, v102)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
								m.G0 = v16 + int32(176)
								return
							}
						}
					}
				}
			} else {
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
				v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
				v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
				v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
				v91 = v16 + int32(80)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
				v99 = base.I32_div_s(v94-v96, int32(296))
				v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
					v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
					v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
					v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
					v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
					F_pgstat_unlock_entry(m, v102)
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return
					} else {
						base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
						m.G0 = v16 + int32(176)
						return
					}
				}
			}
		}
	} else {
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
		if int64(0) < v22 {
			v33 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				if v33 != 0 {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v43
					*(*int64)(unsafe.Add(mBase, uint32(v16-int32(-64)))) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v40
					*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v39
					*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v37
					*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
					F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v16)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(2060), int32(_a_F_UpdateDecodingStats_2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
							v91 = v16 + int32(80)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
							v99 = base.I32_div_s(v94-v96, int32(296))
							v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
								v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
								v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
								v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
								v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
								v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
								v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
								F_pgstat_unlock_entry(m, v102)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
									m.G0 = v16 + int32(176)
									return
								}
							}
						}
					}
				} else {
					v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
					v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
					v91 = v16 + int32(80)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
					v99 = base.I32_div_s(v94-v96, int32(296))
					v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
						F_pgstat_unlock_entry(m, v102)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return
						} else {
							base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
							m.G0 = v16 + int32(176)
							return
						}
					}
				}
			}
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
			if int64(0) < v25 {
				v33 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v33 != 0 {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v43
						*(*int64)(unsafe.Add(mBase, uint32(v16-int32(-64)))) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v40
						*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v39
						*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v38
						*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v37
						*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v36
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
						F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v16)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(2060), int32(_a_F_UpdateDecodingStats_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
								*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
								v91 = v16 + int32(80)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
								v99 = base.I32_div_s(v94-v96, int32(296))
								v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
									v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
									v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
									v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
									v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
									v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
									v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
									v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
									v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
									v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
									v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
									v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
									v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
									v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
									v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
									*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
									F_pgstat_unlock_entry(m, v102)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return
									} else {
										base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
										m.G0 = v16 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
						v91 = v16 + int32(80)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
						v99 = base.I32_div_s(v94-v96, int32(296))
						v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
							F_pgstat_unlock_entry(m, v102)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
								m.G0 = v16 + int32(176)
								return
							}
						}
					}
				}
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
				if v28 <= int64(0) {
					m.G0 = v16 + int32(176)
					return
				} else {
					v33 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v33 != 0 {
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
							v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
							v41 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v43
							*(*int64)(unsafe.Add(mBase, uint32(v16-int32(-64)))) = v42
							*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v41
							*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v40
							*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v39
							*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v38
							*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v36
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
							F_errmsg_internal(m, int32(_a_F_UpdateDecodingStats_0), v16)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_UpdateDecodingStats_1), int32(2060), int32(_a_F_UpdateDecodingStats_2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
									v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
									v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
									*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
									v91 = v16 + int32(80)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
									v99 = base.I32_div_s(v94-v96, int32(296))
									v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
										v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
										v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
										v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
										v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
										v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
										v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
										v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
										v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
										v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
										v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
										v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
										v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
										v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
										v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
										v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
										v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
										F_pgstat_unlock_entry(m, v102)
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return
										} else {
											base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
											m.G0 = v16 + int32(176)
											return
										}
									}
								}
							}
						} else {
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v18)+160))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v72
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v18)+168))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v76
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+200))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)+208))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v84
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v18)+216))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v86
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v18)+224))
							*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v88
							v91 = v16 + int32(80)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v96 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateDecodingStats[0]))
							v99 = base.I32_div_s(v94-v96, int32(296))
							v102 = F_pgstat_get_entry_ref_locked(m, int32(4), int32(0), base.I64_extend_i32_s(v99), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+24))
								v106 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+24)) = v105 + v106
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v104)+32))
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+32)) = v109 + v110
								v113 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+40)) = v113 + v114
								v117 = *(*int64)(unsafe.Add(mBase, uint32(v104)+48))
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = v117 + v118
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v104)+56))
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+56)) = v121 + v122
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v104)+64))
								v126 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+64)) = v125 + v126
								v129 = *(*int64)(unsafe.Add(mBase, uint32(v104)+72))
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v91)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+72)) = v129 + v130
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v104)+80))
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v91)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+80)) = v133 + v134
								v137 = *(*int64)(unsafe.Add(mBase, uint32(v104)+88))
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v91)+64))
								*(*int64)(unsafe.Add(mBase, uint32(v104)+88)) = v137 + v138
								F_pgstat_unlock_entry(m, v102)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									base.MemoryFill(m, v18+int32(160), int32(0), int32(72))
									m.G0 = v16 + int32(176)
									return
								}
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, _c_F_uint32in_subr[0])) = v5
	v20 = F_strtox_2(m, l0, v11+int32(44), v5, int64(4294967295))
	mBase = m.M
	v21 = base.I32_wrap_i64(v20)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_uint32in_subr[0]))
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
							F_errmsg(m, int32(_a_F_uint32in_subr_0), v11+int32(16))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(923), int32(_a_F_uint32in_subr_2))
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
					v83 = v30
					for {
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
						if base.B2i32(base.Ui32(v85-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v85 == int32(32)) != 0 {
							v83 = v83 + int32(1)
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
						v97 = int32(0)
						v98 = F_errsave_start(m, l3)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							if v98 == int32(0) {
								v123 = v97
								m.G0 = v11 + int32(48)
								return v123
							} else {
								F_errcode(m, int32(33685634))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l2
									F_errmsg(m, int32(_a_F_uint32in_subr_3), v11+int32(32))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(939), int32(_a_F_uint32in_subr_2))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											v123 = v97
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
						F_errmsg(m, int32(_a_F_uint32in_subr_3), v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(917), int32(_a_F_uint32in_subr_2))
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
					F_errmsg(m, int32(_a_F_uint32in_subr_3), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, l3, int32(_a_F_uint32in_subr_1), int32(917), int32(_a_F_uint32in_subr_2))
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
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_umask(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _c_F_umask[0])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_unaccent_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v144 int64
	_ = v144
	v2 = int32(0)
	v12 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2
	if v18 <= v2 {
		v144 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v144
L2:
	;
	v27 = v19
	v30 = v18
	goto L3
L3:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v121 == int32(0) {
		v144 = v12
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v118 = v30 - v108
	if int32(0) < v118 {
		v27 = v27 + v108
		v30 = v118
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v98 = F_pg_mblen_range(m, v27, v18+v19)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L23
	} else {
		goto L28
	}
L7:
	;
	v39 = int32(0)
	v42 = v39
	v45 = v39
	v46 = v39
	v49 = v17
	goto L8
L8:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v27))))
	v58 = v49 + v55*int32(12)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v60 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L10:
	;
	goto L9
L11:
	;
	v60 = v58
	goto L13
L12:
	;
	v60 = v46
	goto L13
L13:
	;
	v62 = v42 + int32(1)
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = v62
	goto L16
L15:
	;
	v63 = v45
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v64 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v62) < base.Ui32(v30) {
		v42 = v62
		v45 = v63
		v46 = v60
		v49 = v64
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v70 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v73 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_appendBinaryStringInfo(m, v15, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L23
	} else {
		goto L27
	}
L22:
	;
	F_initStringInfo(m, v15)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int64(0)
L24:
	;
	if v27 == v19 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F_appendBinaryStringInfo(m, v15, v19, v27-v19)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v108 = v63
	goto L5
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v100 == int32(0) {
		v108 = v98
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_appendBinaryStringInfo(m, v15, v27, v98)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v108 = v98
	goto L5
L31:
	;
	goto L4
L32:
	;
	v126 = F_palloc0_mul(m, int32(8), int32(2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v129 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+2)) = uint16(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v128
	v144 = base.I64_extend_i32_u(v126)
	goto L1
}
func F_union_tuples(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	v4 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_union_tuples_0), v4, int32(_a_F_union_tuples_1), int32(_a_F_union_tuples_2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = int32(_a_F_union_tuples_3)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0])) = v23
	v30 = F_brin_deform_tuple(m, l0, l2, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_union_tuples[0])) = v26
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(0) < v36 {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	F_MemoryContextDelete(m, v23)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v37 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v36 <= int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(24)
	v51 = v4
	goto L9
L9:
	;
	v64 = v51 * int32(24)
	v65 = v30 + v42 + v64
	v66 = v64 + (l1 + v42)
	v68 = v51 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v68)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	if v71 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L5
L11:
	;
	v169 = v51 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v169 < v171 {
		v51 = v169
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = F_index_getprocinfo(m, v137, base.I32_extend16_s(v51+int32(1)), int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L28
	}
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+3)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	if v77 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = int32(1)
	goto L16
L15:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+3)))
	v80 = v79
	goto L16
L16:
	;
	if base.B2i32(v74 == int32(0))&(v80&int32(1)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+3)))
	if v86 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+2)) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+3)))
	if v135 != 0 {
		goto L11
	} else {
		goto L27
	}
L20:
	;
	if v74 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v89 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v89)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	if v91 == int32(0) {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v99 = int32(0)
	goto L23
L23:
	;
	v114 = v99 << (uint(int32(3)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v114+v115)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(8)+v99<<(uint(int32(2))%32))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+10)))
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+8)))
	v124 = F_datumCopy(m, v117, v122, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L11
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v114))) = v124
	v130 = v99 + int32(1)
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui32(v130) < base.Ui32(v131) {
		v99 = v130
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L12
L28:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+248))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v145+v68)))
	v150 = F_FunctionCall3Coll(m, v142, v147, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(v66), base.I64_extend_i32_u(v65))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	goto L10
L31:
	;
	return
L32:
	;
	v195 = int32(24)
	v203 = v4
	goto L35
L33:
	;
	goto L34
L34:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v307)
	F_MemoryContextDelete(m, v23)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L45
	}
L35:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v203<<(uint(int32(2))%32))))
	v220 = v203 * int32(24)
	v221 = l1 + v195 + v220
	v222 = v220 + (v30 + v195)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v221)+3)) = uint8(v223)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v221)+2)) = uint8(v225)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+3)))
	if v227 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	v287 = v203 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v287 < v289 {
		v203 = v287
		goto L35
	} else {
		goto L44
	}
L38:
	;
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218))))
	if v228 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v236 = int32(0)
	goto L40
L40:
	;
	v251 = v236 << (uint(int32(3)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v251+v252)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v218+int32(8)+v236<<(uint(int32(2))%32))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+10)))
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v258)+8)))
	v261 = F_datumCopy(m, v254, v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L37
L42:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v263+v251))) = v261
	v267 = v236 + int32(1)
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218))))
	if base.Ui32(v267) < base.Ui32(v268) {
		v236 = v267
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L36
L45:
	;
	return
}
func F_unique_key_recheck(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(352)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L54
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v16 != int32(448) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19&int32(28) != int32(4) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	switch v19 & int32(3) {
	case 0:
		v49 = int32(24)
		goto L7
	default:
		goto L9
	case 2:
		goto L8
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13+v49)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+348)) = uint16(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+344)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v58 = F_table_slot_create(m, v56, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L15
	}
L8:
	;
	v49 = int32(28)
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int64(0)
L11:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_8), v11+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(83), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
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
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+348)))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+340)) = uint16(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+344))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+336)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_unique_key_recheck[0]))
	if v65 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_unique_key_recheck[1])))
	if v67&int32(1) == int32(0) {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+188))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	v76 = m.T0[v75].(func(*base.Module, int32, int32) int32)(m, v72, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+80)) = uint8(v78)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+188))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	v89 = m.T0[v88].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v76, v11+int32(336), int32(_a_F_unique_key_recheck_7), v58, v11+int32(80), v78)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L22
	}
L21:
	;
	m.G0 = v11 + int32(352)
	return int64(0)
L22:
	;
	if v89 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ExecDropSingleTupleTableSlot(m, v58)
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+188))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+52))
	m.T0[v102].(func(*base.Module, int32))(m, v76)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+188))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	m.T0[v97].(func(*base.Module, int32))(m, v76)
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
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
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
	v124 = v11 + int32(80)
	v126 = v11 + int32(48)
	F_FormIndexDatum(m, v110, v58, v121, v124, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L40
	}
L30:
	;
	v114 = F_CreateExecutorState(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	v121 = v2
	goto L29
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+152))
	if v116 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = v116
	goto L38
L37:
	;
	v117 = F_MakePerTupleExprContext(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v58
	v121 = v114
	goto L29
L39:
	;
	v119 = v117
	goto L38
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v110)+92))
	if v130 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v121 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v137 = F_index_insert(m, v108, v124, v126, v11+int32(344), v129, int32(3), int32(0), v110)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_check_exclusion_constraint(m, v129, v108, v110, v11+int32(336), v11+int32(80), v11+int32(48), v121, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L47
	}
L45:
	;
	F_index_insert_cleanup(m, v108, v110)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	F_FreeExecutorState(m, v121)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_ExecDropSingleTupleTableSlot(m, v58)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v156 = m.ExcPending
	if v156 != 0 {
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
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_1), v11)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(62), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_unique_key_recheck_0)
	F_errmsg(m, int32(_a_F_unique_key_recheck_3), v11+int32(32))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_2), int32(69), int32(_a_F_unique_key_recheck_0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_unique_key_recheck_4), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_unique_key_recheck_5), int32(1256), int32(_a_F_unique_key_recheck_6))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
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
func F_unistr(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v564 int32
	_ = v564
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v878 int32
	_ = v878
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	v21 = m.G0
	v23 = v21 - int32(96)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v31 = int32(1)
	v32 = v30 & v31
	if v30 == v31 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L242
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L238
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L234
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L230
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L226
	}
L8:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v1050 = v1048 + int32(4)
	v1051 = F_palloc(m, v1050)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L221
	}
L9:
	;
	if v32 != 0 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v64 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L12:
	;
	F_initStringInfo(m, v23+int32(80))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L21
	}
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v32 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L16:
	;
	if v35 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = int32(16)
	goto L19
L18:
	;
	v46 = int32(0)
	goto L19
L19:
	;
	v64 = v46
	goto L10
L20:
	;
	v49 = int32(1)
	v64 = int32(base.Ui32(v30)>>(uint(v49)%32)) - v49
	goto L10
L21:
	;
	v72 = int32(4)
	goto L9
L22:
	;
	if v64 <= int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v72 = v64
	goto L9
L24:
	;
	v75 = int32(1)
	goto L26
L25:
	;
	v75 = int32(4)
	goto L26
L26:
	;
	v78 = v26 + v75
	v82 = v72
	v84 = int32(0)
	goto L27
L27:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78))))
	if v97 == int32(92) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v1010&int32(_a_F_unistr_0) != 0 {
		goto L7
	} else {
		goto L220
	}
L29:
	;
	if int32(0) < v1008 {
		v78 = v1004
		v82 = v1008
		v84 = v1010
		goto L27
	} else {
		goto L219
	}
L30:
	;
	v1004 = v1001
	v1008 = v1000
	v1010 = int32(0)
	goto L29
L31:
	;
	if v82 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v84&int32(_a_F_unistr_0) != 0 {
		goto L7
	} else {
		goto L217
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L212
	}
L35:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v102 == int32(92) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v84&int32(_a_F_unistr_0) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v82) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L41
	}
L39:
	;
	F_appendStringInfoChar(m, v23+int32(80), int32(92))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v112 = int32(2)
	v1000 = v82 - v112
	v1001 = v78 + v112
	goto L30
L41:
	;
	v121 = int32(0)
	goto L42
L42:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+(v78+int32(1))))))
	v153 = base.B2i32(base.Ui32(v142-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v142|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L44
L43:
	;
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v153 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = v121 + int32(1)
	if v155 != int32(4) {
		v121 = v155
		goto L42
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L43
L48:
	;
	goto L47
L49:
	;
	if base.Ui32(v82) < base.Ui32(int32(8)) {
		goto L34
	} else {
		goto L98
	}
L50:
	;
	if v82 == int32(5) {
		goto L34
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v228 = int32(-48)
	if v102 == int32(117) {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	if v102 != int32(117) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v168 = int32(0)
	goto L55
L55:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+(v78+int32(2))))))
	v200 = base.B2i32(base.Ui32(v189-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v189|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L57
L56:
	;
	if v200 == int32(0) {
		goto L34
	} else {
		goto L62
	}
L57:
	;
	if v200 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v202 = v168 + int32(1)
	if v202 != int32(4) {
		v168 = v202
		goto L55
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L56
L61:
	;
	goto L60
L62:
	;
	goto L52
L63:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if base.Ui32((v258-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v279 = v228
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v234 = int32(2)
	goto L66
L65:
	;
	v234 = int32(1)
	goto L66
L66:
	;
	v235 = v78 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if base.Ui32((v236-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v257 = v228
		goto L63
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32((v236-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v257 = int32(-87)
		goto L63
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v236-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v257 = int32(-55)
	goto L63
L70:
	;
	v280 = int32(-48)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+2)))
	if base.Ui32((v282-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v303 = v280
		goto L76
	} else {
		goto L77
	}
L71:
	;
	if base.Ui32((v258-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v279 = int32(-87)
	goto L70
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(5)) < base.Ui32((v258-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v279 = int32(-55)
	goto L70
L76:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+3)))
	if base.Ui32((v304-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v325 = v280
		goto L80
	} else {
		goto L81
	}
L77:
	;
	if base.Ui32((v282-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v303 = int32(-87)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(int32(5)) < base.Ui32((v282-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v303 = int32(-55)
	goto L76
L80:
	;
	v338 = v325 + v304 + ((v279+v258)<<(uint(int32(8))%32) + (v236+v257)<<(uint(int32(12))%32) + (v282+v303)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_1)) <= base.Ui32(v338-int32(1)) {
		goto L6
	} else {
		goto L86
	}
L81:
	;
	if base.Ui32((v304-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v325 = int32(-87)
	goto L80
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(5)) < base.Ui32((v304-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v325 = int32(-55)
	goto L80
L86:
	;
	v344 = v338 & int32(_a_F_unistr_2)
	if v84&int32(_a_F_unistr_0) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v360&int32(-1024) == int32(_a_F_unistr_3) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	if v344 != int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v344 == int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L92
	}
L91:
	;
	v349 = int32(1023)
	v360 = v338&v349 | v84&v349<<(uint(int32(10))%32) + int32(_a_F_unistr_5)
	goto L87
L92:
	;
	v360 = v338
	goto L87
L93:
	;
	v375 = v360
	goto L95
L94:
	;
	v366 = v23 + int32(48)
	F_pg_unicode_to_server(m, v360, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v377 = v234 | int32(4)
	v1004 = v377 + v78
	v1008 = v82 - v377
	v1010 = v375
	goto L29
L96:
	;
	F_appendStringInfoString(m, v23+int32(80), v366)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v375 = int32(0)
	goto L95
L98:
	;
	if v102 != int32(43) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if base.B2i32(v102 != int32(85))|base.B2i32(base.Ui32(v82) < base.Ui32(int32(10))) != 0 {
		goto L34
	} else {
		goto L151
	}
L100:
	;
	v385 = v78 + int32(2)
	v387 = int32(0)
	goto L101
L101:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387+v385))))
	v419 = base.B2i32(base.Ui32(v408-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v408|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L103
L102:
	;
	if v419 == int32(0) {
		goto L99
	} else {
		goto L108
	}
L103:
	;
	if v419 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v421 = v387 + int32(1)
	if v421 != int32(6) {
		v387 = v421
		goto L101
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L102
L107:
	;
	goto L106
L108:
	;
	v427 = int32(-48)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if base.Ui32((v429-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v450 = v427
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+3)))
	if base.Ui32((v451-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v472 = v427
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if base.Ui32((v429-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v450 = int32(-87)
		goto L109
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v429-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v450 = int32(-55)
	goto L109
L113:
	;
	v473 = int32(-48)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
	if base.Ui32((v475-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v496 = v473
		goto L119
	} else {
		goto L120
	}
L114:
	;
	if base.Ui32((v451-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v472 = int32(-87)
	goto L113
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(int32(5)) < base.Ui32((v451-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v472 = int32(-55)
	goto L113
L119:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
	if base.Ui32((v497-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v518 = v473
		goto L123
	} else {
		goto L124
	}
L120:
	;
	if base.Ui32((v475-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v496 = int32(-87)
		goto L119
	} else {
		goto L121
	}
L121:
	;
	if base.Ui32(int32(5)) < base.Ui32((v475-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v496 = int32(-55)
	goto L119
L123:
	;
	v519 = int32(-48)
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
	if base.Ui32((v521-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v542 = v519
		goto L129
	} else {
		goto L130
	}
L124:
	;
	if base.Ui32((v497-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v518 = int32(-87)
	goto L123
L126:
	;
	goto L127
L127:
	;
	if base.Ui32(int32(5)) < base.Ui32((v497-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v518 = int32(-55)
	goto L123
L129:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+7)))
	if base.Ui32((v543-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v564 = v519
		goto L133
	} else {
		goto L134
	}
L130:
	;
	if base.Ui32((v521-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v542 = int32(-87)
		goto L129
	} else {
		goto L131
	}
L131:
	;
	if base.Ui32(int32(5)) < base.Ui32((v521-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v542 = int32(-55)
	goto L129
L133:
	;
	v585 = v564 + v543 + ((v472+v451)<<(uint(int32(16))%32) + (v429+v450)<<(uint(int32(20))%32) + (v475+v496)<<(uint(int32(12))%32) + (v518+v497)<<(uint(int32(8))%32) + (v521+v542)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_1)) <= base.Ui32(v585-int32(1)) {
		goto L5
	} else {
		goto L139
	}
L134:
	;
	if base.Ui32((v543-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v564 = int32(-87)
	goto L133
L136:
	;
	goto L137
L137:
	;
	if base.Ui32(int32(5)) < base.Ui32((v543-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	v564 = int32(-55)
	goto L133
L139:
	;
	v591 = v585 & int32(_a_F_unistr_2)
	if v84&int32(_a_F_unistr_0) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v607&int32(-1024) == int32(_a_F_unistr_3) {
		goto L146
	} else {
		goto L147
	}
L141:
	;
	if v591 != int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v591 == int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L145
	}
L144:
	;
	v596 = int32(1023)
	v607 = v585&v596 | v84&v596<<(uint(int32(10))%32) + int32(_a_F_unistr_5)
	goto L140
L145:
	;
	v607 = v585
	goto L140
L146:
	;
	v622 = v607
	goto L148
L147:
	;
	v613 = v23 + int32(48)
	F_pg_unicode_to_server(m, v607, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v623 = int32(8)
	v1004 = v78 + v623
	v1008 = v82 - v623
	v1010 = v622
	goto L29
L149:
	;
	F_appendStringInfoString(m, v23+int32(80), v613)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v622 = int32(0)
	goto L148
L151:
	;
	v653 = v78 + int32(2)
	v655 = int32(0)
	goto L152
L152:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655+v653))))
	v687 = base.B2i32(base.Ui32(v676-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v676|int32(32)-int32(97)) < base.Ui32(int32(6)))
	goto L154
L153:
	;
	if v687 == int32(0) {
		goto L34
	} else {
		goto L159
	}
L154:
	;
	if v687 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v689 = v655 + int32(1)
	if v689 != int32(8) {
		v655 = v689
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
	v695 = int32(-48)
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	if base.Ui32((v697-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v718 = v695
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+3)))
	if base.Ui32((v719-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v740 = v695
		goto L164
	} else {
		goto L165
	}
L161:
	;
	if base.Ui32((v697-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v718 = int32(-87)
		goto L160
	} else {
		goto L162
	}
L162:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v697-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	v718 = int32(-55)
	goto L160
L164:
	;
	v741 = int32(-48)
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)))
	if base.Ui32((v743-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v764 = v741
		goto L170
	} else {
		goto L171
	}
L165:
	;
	if base.Ui32((v719-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v740 = int32(-87)
	goto L164
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(int32(5)) < base.Ui32((v719-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	v740 = int32(-55)
	goto L164
L170:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
	if base.Ui32((v765-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v786 = v741
		goto L174
	} else {
		goto L175
	}
L171:
	;
	if base.Ui32((v743-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v764 = int32(-87)
		goto L170
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(int32(5)) < base.Ui32((v743-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v764 = int32(-55)
	goto L170
L174:
	;
	v787 = int32(-48)
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
	if base.Ui32((v789-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v810 = v787
		goto L180
	} else {
		goto L181
	}
L175:
	;
	if base.Ui32((v765-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v786 = int32(-87)
	goto L174
L177:
	;
	goto L178
L178:
	;
	if base.Ui32(int32(5)) < base.Ui32((v765-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	v786 = int32(-55)
	goto L174
L180:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+7)))
	if base.Ui32((v811-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v832 = v787
		goto L184
	} else {
		goto L185
	}
L181:
	;
	if base.Ui32((v789-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v810 = int32(-87)
		goto L180
	} else {
		goto L182
	}
L182:
	;
	if base.Ui32(int32(5)) < base.Ui32((v789-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	v810 = int32(-55)
	goto L180
L184:
	;
	v833 = int32(-48)
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
	if base.Ui32((v835-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v856 = v833
		goto L190
	} else {
		goto L191
	}
L185:
	;
	if base.Ui32((v811-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v832 = int32(-87)
	goto L184
L187:
	;
	goto L188
L188:
	;
	if base.Ui32(int32(5)) < base.Ui32((v811-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	v832 = int32(-55)
	goto L184
L190:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+9)))
	if base.Ui32((v857-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v878 = v833
		goto L194
	} else {
		goto L195
	}
L191:
	;
	if base.Ui32((v835-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v856 = int32(-87)
		goto L190
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(int32(5)) < base.Ui32((v835-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	v856 = int32(-55)
	goto L190
L194:
	;
	v907 = v878 + v857 + ((v740+v719)<<(uint(int32(24))%32) + (v697+v718)<<(uint(int32(28))%32) + (v743+v764)<<(uint(int32(20))%32) + (v786+v765)<<(uint(int32(16))%32) + (v789+v810)<<(uint(int32(12))%32) + (v832+v811)<<(uint(int32(8))%32) + (v835+v856)<<(uint(int32(4))%32))
	if base.Ui32(int32(_a_F_unistr_1)) <= base.Ui32(v907-int32(1)) {
		goto L4
	} else {
		goto L200
	}
L195:
	;
	if base.Ui32((v857-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v878 = int32(-87)
	goto L194
L197:
	;
	goto L198
L198:
	;
	if base.Ui32(int32(5)) < base.Ui32((v857-int32(65))&int32(255)) {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	v878 = int32(-55)
	goto L194
L200:
	;
	v913 = v907 & int32(_a_F_unistr_2)
	if v84&int32(_a_F_unistr_0) != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if v929&int32(-1024) == int32(_a_F_unistr_3) {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	if v913 != int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v913 == int32(_a_F_unistr_4) {
		goto L7
	} else {
		goto L206
	}
L205:
	;
	v918 = int32(1023)
	v929 = v907&v918 | v84&v918<<(uint(int32(10))%32) + int32(_a_F_unistr_5)
	goto L201
L206:
	;
	v929 = v907
	goto L201
L207:
	;
	v944 = v929
	goto L209
L208:
	;
	v935 = v23 + int32(48)
	F_pg_unicode_to_server(m, v929, v935)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v945 = int32(10)
	v1004 = v78 + v945
	v1008 = v82 - v945
	v1010 = v944
	goto L29
L210:
	;
	F_appendStringInfoString(m, v23+int32(80), v935)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v944 = int32(0)
	goto L209
L212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_unistr_6), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errhint(m, int32(_a_F_unistr_7), int32(0))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_9), int32(_a_F_unistr_10))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
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
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v995 = int32(1)
	v1000 = v82 - v995
	v1001 = v78 + v995
	goto L30
L219:
	;
	goto L28
L220:
	;
	goto L8
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1051))) = v1050 << (uint(int32(2)) % 32)
	if v1048 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	base.MemoryCopy(m, v1051+int32(4), v1047, v1048)
	goto L224
L223:
	;
	goto L224
L224:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	F_pfree(m, v1059)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	m.G0 = v23 + int32(96)
	return base.I64_extend_i32_u(v1051)
L226:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_unistr_11), int32(0))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_12), int32(_a_F_unistr_10))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v338
	F_errmsg(m, int32(_a_F_unistr_13), v23)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_14), int32(_a_F_unistr_10))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v585
	F_errmsg(m, int32(_a_F_unistr_13), v23+int32(16))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_15), int32(_a_F_unistr_10))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v907
	F_errmsg(m, int32(_a_F_unistr_13), v23+int32(32))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_16), int32(_a_F_unistr_10))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_unistr_17), int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_unistr_8), int32(_a_F_unistr_18), int32(_a_F_unistr_19))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_updateClosestMatch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v6 == v3)|base.B2i32(l1 == v3) != 0 {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v12 == int32(0) {
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v15 == int32(0) {
				return
			} else {
				v18 = F_strlen(m, v6)
				mBase = m.M
				if base.Ui32(int32(255)) < base.Ui32(v18) {
					return
				} else {
					v21 = F_strlen(m, l1)
					mBase = m.M
					if base.Ui32(int32(255)) < base.Ui32(v21) {
						return
					} else {
						v24 = int32(1)
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v29 = F_varstr_levenshtein_less_equal(m, v6, v18, l1, v21, v24, v24, v24, v27, v24)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v31 < v29 {
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v34 = F_strlen(m, v33)
								mBase = m.M
								if base.Ui32(int32(base.Ui32(v34)>>(uint(int32(1))%32))) < base.Ui32(v29) {
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if base.B2i32(v38 != int32(-1))&base.B2i32(v38 <= v29) != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
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
func F_update_controlfile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(_a_F_update_controlfile_0)
	m.G0 = v8
	v10 = F_time(m)
	mBase = m.M
	v11 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+308)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v10
	v16 = m.Env.Pgmem_crc32c(m, v11, l1, int32(308))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+308)) = v16 ^ v11
	base.MemoryFill(m, v8+int32(1416), int32(0), int32(_a_F_update_controlfile_1))
	v26 = v8 + int32(1104)
	base.MemoryCopy(m, v26, l1, int32(312))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = int32(_a_F_update_controlfile_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l0
	v33 = v8 + int32(80)
	v38 = F_pg_snprintf(m, v33, int32(1024), int32(_a_F_update_controlfile_3), v8-int32(-64))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v41 = F_BasicOpenFile(m, v33, int32(2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	if int32(0) <= v41 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0])) = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(167772173)
	v52 = int32(_a_F_update_controlfile_4)
	v53 = F_write(m, v41, v26, v52)
	mBase = m.M
	if v53 != v52 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0]))
	if v57 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v82 = int32(_a_F_update_controlfile_5)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(167772171)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_controlfile[2])))
	if v92 != int32(1) {
		v106 = v84
		goto L20
	} else {
		goto L21
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0])) = int32(51)
	goto L14
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_6), v8+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(248), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if v106 != 0 {
		goto L4
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	goto L22
L22:
	;
	v97 = F_fsync(m, v41)
	mBase = m.M
	if v97 != int32(-1) {
		v106 = v97
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v106 = int32(-1)
	goto L20
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[0]))
	if v101 == int32(27) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_update_controlfile[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(0)
	v111 = F_close(m, v41)
	mBase = m.M
	if v111 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v8 + int32(_a_F_update_controlfile_0)
	return
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_9), v8)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(227), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_10), v8+int32(32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(265), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(80)
	F_errmsg(m, int32(_a_F_update_controlfile_11), v8+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_update_controlfile_7), int32(279), int32(_a_F_update_controlfile_8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_update_most_recent_deletion_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_ExecFetchSlotHeapTuple(m, l0, v6, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		F_LockBufferInternal(m, v16, int32(1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = F_HeapTupleSatisfiesVacuum(m, v14, l1, v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_UnlockBuffer(m, v16)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if v20 != int32(2) {
						m.G0 = v10 + int32(16)
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+20)))
						if v27&int32(_a_F_update_most_recent_deletion_info_0) == int32(_a_F_update_most_recent_deletion_info_1) {
							v32 = F_HeapTupleGetUpdateXid(m, v26)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v35 = v32
								if v35 == int32(0) {
									m.G0 = v10 + int32(16)
									return
								} else {
									v42 = F_TransactionIdGetCommitTsData(m, v35, v10+int32(8), v10+int32(6))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										if v42 == int32(0) {
										} else {
											v46 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
											v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.B2i32(base.I64_extend_i32_s(int32(0))*int64(1000) <= v47-v46) == int32(0) {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
												v57 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l3))) = v57
												v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
												*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v59)
											}
										}
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
							v35 = v34
							if v35 == int32(0) {
								m.G0 = v10 + int32(16)
								return
							} else {
								v42 = F_TransactionIdGetCommitTsData(m, v35, v10+int32(8), v10+int32(6))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if v42 == int32(0) {
									} else {
										v46 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
										v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.B2i32(base.I64_extend_i32_s(int32(0))*int64(1000) <= v47-v46) == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
											v57 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l3))) = v57
											v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
											*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v59)
										}
									}
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_update_spins_per_delay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_update_spins_per_delay[0]))
	v8 = base.I32_div_s(v3+l0*int32(15), int32(16))
	return v8
}
